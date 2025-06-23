// main.go - ConnectBot Plus Backend Installer (MVP Version)
package main

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "os/exec"
    "runtime"
    "strings"
    "time"
)

const (
    SSHPort  = 8822
    UserName = "terminal-user"
    Version  = "1.0.0-mvp"
)

type PairingInfo struct {
    Version     int      `json:"v"`
    Host        string   `json:"h"`
    Port        int      `json:"p"`
    Fingerprint string   `json:"fp"`
    Timestamp   int64    `json:"t"`
    Name        string   `json:"n,omitempty"`
    Methods     []string `json:"m"`
}

func main() {
    fmt.Printf("ConnectBot Plus Backend Installer v%s\n", Version)
    fmt.Println("=====================================\n")
    
    // Check OS
    if runtime.GOOS != "linux" {
        log.Fatal("This installer only supports Linux")
    }
    
    // Check root
    if os.Geteuid() != 0 {
        fmt.Println("⚠️  Please run with sudo:")
        fmt.Println("  curl -sSL https://cbplus.app/install | sudo bash")
        os.Exit(1)
    }
    
    fmt.Println("📦 Installing SSH server...")
    installSSH()
    
    fmt.Println("👤 Creating user account...")
    createUser()
    
    fmt.Println("🔧 Configuring SSH...")
    fingerprint := configureSSH()
    
    fmt.Println("🔑 Generating pairing code...")
    uid := generateUID(fingerprint)
    
    fmt.Println("🚀 Starting SSH service...")
    startSSH()
    
    displaySuccess(uid)
}

func installSSH() {
    // Try common package managers
    if _, err := exec.LookPath("apt-get"); err == nil {
        exec.Command("apt-get", "update").Run()
        cmd := exec.Command("apt-get", "install", "-y", "openssh-server")
        if err := cmd.Run(); err != nil {
            log.Fatal("Failed to install SSH: ", err)
        }
    } else if _, err := exec.LookPath("yum"); err == nil {
        cmd := exec.Command("yum", "install", "-y", "openssh-server")
        if err := cmd.Run(); err != nil {
            log.Fatal("Failed to install SSH: ", err)
        }
    } else {
        fmt.Println("⚠️  Could not install SSH automatically.")
        fmt.Println("Please install openssh-server manually and run again.")
        os.Exit(1)
    }
}

func createUser() {
    // Check if user exists
    if _, err := exec.Command("id", UserName).Output(); err == nil {
        fmt.Println("  User already exists ✓")
        return
    }
    
    // Create user
    cmd := exec.Command("useradd", "-m", "-s", "/bin/bash", UserName)
    if err := cmd.Run(); err != nil {
        log.Fatal("Failed to create user: ", err)
    }
    
    // Create .ssh directory
    sshDir := fmt.Sprintf("/home/%s/.ssh", UserName)
    os.MkdirAll(sshDir, 0700)
    
    // Create authorized_keys file
    keysFile := fmt.Sprintf("%s/authorized_keys", sshDir)
    os.OpenFile(keysFile, os.O_CREATE, 0600)
    
    // Fix ownership
    exec.Command("chown", "-R", UserName+":"+UserName, sshDir).Run()
}

func configureSSH() string {
    // Simple SSH config
    config := fmt.Sprintf(`
# ConnectBot Plus Configuration
Port %d
PermitRootLogin no
PasswordAuthentication no
PubkeyAuthentication yes
AllowUsers %s
`, SSHPort, UserName)
    
    // Append to sshd_config
    f, err := os.OpenFile("/etc/ssh/sshd_config", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("Failed to open sshd_config: ", err)
    }
    defer f.Close()
    
    f.WriteString("\n" + config)
    
    // Get host key fingerprint
    hostKeyPath := "/etc/ssh/ssh_host_rsa_key.pub"
    if _, err := os.Stat(hostKeyPath); os.IsNotExist(err) {
        // Generate host key
        exec.Command("ssh-keygen", "-A").Run()
    }
    
    // Read public key
    pubKeyData, err := os.ReadFile(hostKeyPath)
    if err != nil {
        return "unknown"
    }
    
    parts := strings.Fields(string(pubKeyData))
    if len(parts) < 2 {
        return "unknown"
    }
    
    // Decode base64 key
    keyBytes, _ := base64.StdEncoding.DecodeString(parts[1])
    hash := sha256.Sum256(keyBytes)
    fingerprint := fmt.Sprintf("SHA256:%s", base64.StdEncoding.EncodeToString(hash[:]))
    
    return fingerprint
}

func generateUID(fingerprint string) string {
    // Get IP address
    resp, err := http.Get("https://api.ipify.org")
    var hostIP string
    if err == nil {
        defer resp.Body.Close()
        ipBytes, _ := io.ReadAll(resp.Body)
        hostIP = string(ipBytes)
    } else {
        // Fallback to local IP
        hostIP = getLocalIP()
    }
    
    // Get hostname
    hostname, _ := os.Hostname()
    
    // Create pairing info
    info := PairingInfo{
        Version:     1,
        Host:        hostIP,
        Port:        SSHPort,
        Fingerprint: fingerprint,
        Timestamp:   time.Now().Unix(),
        Name:        hostname,
        Methods:     []string{"ssh"},
    }
    
    // Encode to JSON
    jsonData, _ := json.Marshal(info)
    
    // Encode to base64
    encoded := base64.StdEncoding.EncodeToString(jsonData)
    
    // Generate random suffix for uniqueness
    randomBytes := make([]byte, 6)
    rand.Read(randomBytes)
    randomSuffix := base64.URLEncoding.EncodeToString(randomBytes)[:6]
    
    // Format as UID
    uid := fmt.Sprintf("TXT-%s-%s-%s",
        encoded[0:6],
        encoded[6:12],
        randomSuffix,
    )
    
    return uid
}

func getLocalIP() string {
    out, err := exec.Command("hostname", "-I").Output()
    if err == nil {
        ips := strings.Fields(string(out))
        if len(ips) > 0 {
            return ips[0]
        }
    }
    return "localhost"
}

func startSSH() {
    // Restart SSH service
    if _, err := exec.LookPath("systemctl"); err == nil {
        exec.Command("systemctl", "restart", "sshd").Run()
        exec.Command("systemctl", "restart", "ssh").Run() // Ubuntu/Debian
    } else if _, err := exec.LookPath("service"); err == nil {
        exec.Command("service", "ssh", "restart").Run()
        exec.Command("service", "sshd", "restart").Run()
    }
    
    // Open firewall port
    if _, err := exec.LookPath("ufw"); err == nil {
        exec.Command("ufw", "allow", fmt.Sprintf("%d/tcp", SSHPort)).Run()
    }
}

func displaySuccess(uid string) {
    fmt.Println("\n✅ Setup Complete!")
    fmt.Println("==================")
    fmt.Println()
    fmt.Println("Your pairing code:")
    fmt.Println()
    fmt.Printf("    %s\n", uid)
    fmt.Println()
    fmt.Println("📱 Next steps:")
    fmt.Println("1. Install ConnectBot Plus on your Android device")
    fmt.Println("2. Tap 'Add Connection' → 'UID Pairing'")
    fmt.Println("3. Enter the code above")
    fmt.Println()
    fmt.Println("⏱️  This code expires in 5 minutes")
    fmt.Println()
    fmt.Printf("SSH is also available at: %s:%d\n", getLocalIP(), SSHPort)
    fmt.Printf("Username: %s\n", UserName)
    fmt.Println()
    fmt.Println("🔒 Security Note: Only paired devices can connect")
}
