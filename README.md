# DegenMon - Token Price Monitor

A Go application that monitors cryptocurrency token prices using the Jupiter API and sends desktop notifications when price changes exceed specified thresholds.

## Features

- **YAML Configuration**: Easy-to-edit configuration file for tokens and thresholds
- **Jupiter API Integration**: Uses Jupiter's price API for real-time token prices
- **Cross-Platform Notifications**: Desktop notifications on Windows, macOS, and Linux
- **Configurable Thresholds**: Set individual price change thresholds for each token
- **Continuous Monitoring**: Runs continuously with configurable check intervals

## Installation

1. Make sure you have Go installed (version 1.16 or later)
2. Clone or download this project
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build -o degenmon
   ```

## Configuration

Edit the `config.yaml` file to configure the tokens you want to monitor:

```yaml
tokens:
  - ticker: troll
    ca: 5UUH9RTDiSpq6HKS6bp4NdU9PNJpXRXuiw6ShBTBhgH2
    priceChange: 0.01  # 1% change threshold
  - ticker: sol
    ca: 9BB6NFEcjBCtnNLFko2FqVQBq8HHM13kCyYcdQbgpump
    priceChange: 0.05  # 5% change threshold
  - ticker: bonk
    ca: DezXAZ8z7PnrnRJjz3wXBoRgixCa6xjnB7YaB1pPB263
    priceChange: 0.02  # 2% change threshold
```

### Configuration Parameters

- `ticker`: The token symbol/ticker to monitor
- `ca` : The token contract address
- `priceChange`: The minimum price change percentage (as decimal) that triggers a notification
  - `0.01` = 1% change
  - `0.05` = 5% change
  - `0.10` = 10% change

## Usage

Run the application:

```bash
./degenmon
```

The application will:
1. Load the configuration from `config.yaml`
2. Start monitoring the specified tokens
3. Check prices every 5 seconds
4. Send desktop notifications when price changes exceed the configured thresholds
5. Log all price changes and notifications to the console

## Dependencies

- `gopkg.in/yaml.v2`: YAML configuration parsing
- `github.com/gen2brain/beeep`: Cross-platform desktop notifications

## Platform Support

- **Windows**: Uses Windows Toast notifications
- **macOS**: Uses native macOS notifications
- **Linux**: Uses D-Bus notifications (requires notification daemon)

## Troubleshooting

### No notifications appearing
- **Linux**: Make sure you have a notification daemon running (like `dunst` or `notification-daemon`)
- **macOS**: Check System Preferences > Notifications to ensure notifications are enabled
- **Windows**: Notifications should work out of the box on Windows 10/11

### API Errors
- Check your internet connection
- Verify that the token tickers and CAs in your config file are valid
- The Jupiter API may have rate limits or temporary outages

### Configuration Issues
- Ensure `config.yaml` is in the same directory as the executable
- Check YAML syntax (proper indentation, no tabs)
- Verify that `priceChange` values are decimal numbers (0.01 for 1%, not 1)

## Customization

You can modify the check interval by changing this line in `main.go`:
```go
checkInterval := 30 * time.Second  // Change to desired interval
```

## License

This project is open source and available under the MIT License.
