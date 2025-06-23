# ConnectBot Plus

> Android SSH client with automated Linux server setup via UID pairing

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## 🚀 What is ConnectBot Plus?

ConnectBot Plus extends the excellent [ConnectBot](https://github.com/connectbot/connectbot) SSH client with a revolutionary feature: **zero-configuration server setup**. 

Instead of manually configuring SSH keys, IP addresses, and server settings, just:

1. Run our installer on your Linux server: `curl -sSL https://cbplus.app/install | sudo bash`
2. Get a pairing code: `TXT-ABC123-DEF456-GHI789`
3. Enter it in the app
4. You're connected! 🎉

## ✨ Features

- **30-second setup**: From zero to SSH connection in under a minute
- **UID Pairing**: Time-limited pairing codes for secure, easy setup
- **No manual configuration**: Automatically configures SSH, creates users, sets up keys
- **Based on ConnectBot**: Battle-tested SSH client with 15+ years of development
- **FOSS**: Dual-licensed under Apache 2.0 and GPL 3.0
- **Self-hosted**: Your servers, your control

## 📱 Installation

### Android App
1. Download the latest APK from [Releases](https://github.com/machjesusmoto/connectbot-plus/releases)
2. Or build from source (see below)

### Server Installer
```bash
# Quick install (Ubuntu/Debian)
curl -sSL https://cbplus.app/install | sudo bash

# Manual download
wget https://github.com/machjesusmoto/connectbot-plus/releases/latest/download/installer-linux-amd64
chmod +x installer-linux-amd64
sudo ./installer-linux-amd64
```

## 🛠️ Building from Source

### Prerequisites
- Android Studio Arctic Fox or newer
- JDK 11+
- Go 1.19+ (for backend installer)

### Android App
```bash
git clone https://github.com/machjesusmoto/connectbot-plus.git
cd connectbot-plus
./gradlew assembleDebug
```

### Backend Installer
```bash
cd backend-installer
go build -o installer cmd/installer/main.go
```

## 🎯 Project Status

**Current Phase: MVP Development**

- [x] Project structure
- [ ] Fork ConnectBot
- [ ] Add UID pairing UI
- [ ] Implement pairing service
- [ ] Create backend installer
- [ ] Integration testing
- [ ] Release v1.0.0-mvp

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Quick Start for Contributors
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is dual-licensed:
- **Apache License 2.0** - See [LICENSE-APACHE](LICENSE-APACHE)
- **GNU General Public License v3.0** - See [LICENSE-GPL](LICENSE-GPL)

You may choose either license for your use.

### Why Dual Licensing?
- **Apache 2.0**: Allows commercial use and proprietary derivatives
- **GPL 3.0**: Ensures improvements return to the community

## 🏗️ Architecture

```
┌─────────────────────────────┐
│   Android App (Client)       │
├─────────────────────────────┤
│  - ConnectBot Base           │
│  - UID Pairing System        │
│  - QR Scanner (future)       │
└──────────────┬──────────────┘
               │
               │ SSH Connection
               │
┌──────────────┴──────────────┐
│   User's Linux System        │
├─────────────────────────────┤
│  - SSH Server (port 8822)    │
│  - terminal-user account     │
│  - Auto-configured by installer │
└─────────────────────────────┘
```

## 🙏 Credits

- [ConnectBot](https://github.com/connectbot/connectbot) - The excellent SSH client we're building upon
- [Termux](https://github.com/termux) - Terminal emulator components (planned for v2)
- All our [contributors](https://github.com/machjesusmoto/connectbot-plus/graphs/contributors)

## 📞 Support

- 🐛 [Issue Tracker](https://github.com/machjesusmoto/connectbot-plus/issues)
- 💬 [Discussions](https://github.com/machjesusmoto/connectbot-plus/discussions)
- 📧 Email: support@cbplus.app (coming soon)

---

**ConnectBot Plus**: SSH setup so easy, your non-technical friends can do it! 🚀