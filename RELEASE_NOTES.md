# DegenMon v1.0.0 Release Notes

## Release Overview

**DegenMon v1.0.0** is the initial stable release of our crypto token price monitoring application. This release introduces a comprehensive monitoring system that tracks token prices in real-time and provides instant desktop notifications when significant price movements occur.

## Key Features

### Real-Time Price Monitoring
- **Jupiter API Integration**: Leverages Jupiter's robust price API for accurate, real-time token price data
- **Continuous Monitoring**: Runs 24/7 with configurable check intervals (default: 5 seconds)
- **Multi-Token Support**: Monitor multiple tokens simultaneously with individual configurations

### Smart Notifications
- **Cross-Platform Desktop Notifications**: Works seamlessly on Windows, macOS, and Linux
- **Configurable Thresholds**: Set custom price change percentages for each token
- **Instant Alerts**: Get notified immediately when price movements exceed your specified thresholds

### Easy Configuration
- **YAML-Based Config**: Simple, human-readable configuration file
- **Token-Specific Settings**: Individual price change thresholds for each monitored token
- **Contract Address Support**: Direct support for Solana token contract addresses

## Current Configuration

The default configuration monitors three popular tokens:

```yaml
tokens:
  - ticker: troll
    ca: 5UUH9RTDiSpq6HKS6bp4NdU9PNJpXRXuiw6ShBTBhgH2
    priceChange: 0.01  # 1% change threshold
  - ticker: fartcoin
    ca: 9BB6NFEcjBCtnNLFko2FqVQBq8HHM13kCyYcdQbgpump
    priceChange: 0.01  # 1% change threshold
  - ticker: bonk
    ca: DezXAZ8z7PnrnRJjz3wXBoRgixCa6xjnB7YaB1pPB263
    priceChange: 0.01  # 1% change threshold
```

## Technical Details

### Built With
- **Go 1.24.2**: Modern Go runtime with latest features and optimizations
- **Jupiter API**: Professional-grade price data from Jupiter
- **beeep**: Cross-platform desktop notification library
- **YAML v2**: Industry-standard configuration parsing

### System Requirements
- **Go 1.16+** (for building from source)
- **Internet Connection** (for API calls)
- **Notification Daemon** (Linux only - `dunst` or `notification-daemon`)

## Installation & Usage

### Quick Start
1. **Download** the appropriate binary for your platform
2. **Configure** your tokens in `config.yaml`
3. **Run** the application:
   ```bash
   ./degenmon
   ```

### Building from Source
```bash
# Install dependencies
go mod tidy

# Build the application
go build -o degenmon

# Run the application
./degenmon
```

## Use Cases

DegenMon is perfect for:
- **Active Traders**: Get instant notifications on significant price movements
- **Portfolio Monitoring**: Track multiple tokens with custom alert thresholds
- **Market Research**: Monitor token performance with real-time data
- **Risk Management**: Set up alerts for price movements that matter to you

## Configuration Options

### Price Change Thresholds
- `0.01` = 1% change triggers notification
- `0.05` = 5% change triggers notification
- `0.10` = 10% change triggers notification

### Adding New Tokens
Simply add new entries to your `config.yaml`:
```yaml
- ticker: yourtoken
  ca: YOUR_TOKEN_CONTRACT_ADDRESS
  priceChange: 0.02  # 2% threshold
```
## Support & Feedback

- **Issues**: Report bugs and request features on GitHub
- **Documentation**: Full documentation available in README.md
- **Configuration Help**: Example configurations provided

## License

This project is released under the MIT License - see LICENSE file for details.

---

**Happy Trading!**

*For questions or support, please refer to the README.md file or create an issue on GitHub.*
