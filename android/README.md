# Android App

This directory will contain the forked ConnectBot code with our pairing additions.

## Setup Instructions

1. Fork ConnectBot: https://github.com/connectbot/connectbot
2. Clone your fork here
3. Add the pairing functionality

## Key Modifications

### New Files to Add
- `PairingActivity.kt` - UID entry screen
- `PairingService.kt` - UID decoding and key exchange
- `QRScannerActivity.kt` - Future QR code scanning

### Files to Modify
- `HostListActivity.java` - Add pairing menu option
- `AndroidManifest.xml` - Register new activities

## Building

```bash
./gradlew assembleDebug
```

## Testing

```bash
./gradlew test
```