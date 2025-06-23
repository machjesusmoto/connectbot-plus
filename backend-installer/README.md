# Backend Installer

This directory contains the automated installer that configures Linux systems for ConnectBot Plus.

## What it does

1. Installs OpenSSH server
2. Creates a dedicated `terminal-user` account
3. Configures SSH on port 8822
4. Generates a time-limited pairing UID
5. Opens firewall ports

## Building

```bash
go build -o installer cmd/installer/main.go
```

## Cross-compilation

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o dist/installer-linux-amd64 cmd/installer/main.go

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o dist/installer-linux-arm64 cmd/installer/main.go
```

## Testing

```bash
# Test in Docker
docker run -it --rm -p 8822:8822 ubuntu:22.04 bash
# Inside container:
apt update && apt install -y curl sudo
curl -sSL http://host.docker.internal:8000/install.sh | bash
```

## Supported Systems

- Ubuntu 20.04+
- Debian 10+
- More coming soon...

## Security Notes

- Creates non-root user
- Disables password authentication
- Uses key-based auth only
- Configures on non-standard port