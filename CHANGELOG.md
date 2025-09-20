# Changelog

All notable changes to DegenMon will be documented in this file.

## [1.0.0] - 2024-12-19

### Initial Release

#### Features Added
- **Real-time token price monitoring** using Jupiter API
- **Cross-platform desktop notifications** (Windows, macOS, Linux)
- **YAML-based configuration system** for easy token management
- **Configurable price change thresholds** per token
- **Continuous monitoring** with 5-second check intervals
- **Multi-token support** with individual alert settings

#### Technical Implementation
- Built with **Go 1.24.2** for optimal performance
- Integrated **Jupiter API** for accurate price data
- Uses **beeep** library for cross-platform notifications
- **YAML v2** for configuration parsing

#### Default Configuration
- Monitors 3 tokens: troll, fartcoin, bonk
- 1% price change threshold for all tokens
- 5-second monitoring interval

#### Platform Support
- ✅ **Windows**: Toast notifications
- ✅ **macOS**: Native notifications
- ✅ **Linux**: D-Bus notifications (requires notification daemon)

### 📝 Notes
- This is the first stable release of DegenMon
- Focus on core monitoring and notification functionality
- Future releases will include web interface and advanced features

---

## Installation
```bash
# Download binary for your platform
# Configure tokens in config.yaml
./degenmon
```

## Configuration Example
```yaml
tokens:
  - ticker: troll
    ca: 5UUH9RTDiSpq6HKS6bp4NdU9PNJpXRXuiw6ShBTBhgH2
    priceChange: 0.01
  - ticker: fartcoin
    ca: 9BB6NFEcjBCtnNLFko2FqVQBq8HHM13kCyYcdQbgpump
    priceChange: 0.01
  - ticker: bonk
    ca: DezXAZ8z7PnrnRJjz3wXBoRgixCa6xjnB7YaB1pPB263
    priceChange: 0.01
