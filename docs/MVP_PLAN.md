# MVP Implementation Plan

## Goal
Create a working Android SSH client with automated server setup in 4 weeks.

## Week 1: Fork & Setup

### Day 1-2: Fork ConnectBot
- [ ] Fork ConnectBot repository
- [ ] Clone and open in Android Studio
- [ ] Verify build works
- [ ] Run on device/emulator

### Day 3-4: Add Pairing Classes
- [ ] Create pairing package
- [ ] Implement PairingActivity
- [ ] Implement PairingService
- [ ] Add to AndroidManifest

### Day 5: UI Integration
- [ ] Add pairing option to main menu
- [ ] Create pairing layout
- [ ] Test navigation flow

## Week 2: Backend Installer

### Day 1: Basic Installer
- [ ] Create Go project structure
- [ ] Implement SSH installation
- [ ] User creation logic
- [ ] UID generation

### Day 2: Cross-compilation
- [ ] Build for linux-amd64
- [ ] Build for linux-arm64
- [ ] Create install script

### Day 3: Testing
- [ ] Test on Ubuntu
- [ ] Test on Debian
- [ ] Document any issues

## Week 3: Integration

### Day 1-2: Connect Systems
- [ ] Test UID decode in app
- [ ] Implement key exchange
- [ ] Save connection to database

### Day 3-4: SSH Connection
- [ ] Test SSH connection
- [ ] Handle connection errors
- [ ] Verify terminal works

### Day 5: Edge Cases
- [ ] Expired UID handling
- [ ] Network error handling
- [ ] Invalid UID format

## Week 4: Polish & Release

### Day 1-2: Testing
- [ ] End-to-end testing
- [ ] Fix critical bugs
- [ ] Update documentation

### Day 3-4: Release Prep
- [ ] Create signed APK
- [ ] Write release notes
- [ ] Update README

### Day 5: Launch
- [ ] Create GitHub release
- [ ] Upload artifacts
- [ ] Announce on Reddit/HN

## Success Criteria

- Setup time: <2 minutes ✅
- Zero manual configuration ✅
- Works on Ubuntu/Debian ✅
- Clear error messages ✅
- Basic but functional UI ✅