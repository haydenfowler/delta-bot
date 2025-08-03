# Extending Delta Bot - Adding New Exchanges

This guide demonstrates how Delta Bot's extensible architecture allows you to add new cryptocurrency exchanges without modifying existing code, following the **Open/Closed Principle**.

## Overview

The exchange system is built around interfaces and factories, making it easy to add support for new exchanges while maintaining code stability and testability.

## Architecture Components

### 1. Core Interfaces (`internal/exchange/interfaces.go`)

All exchanges must implement the `Exchange` interface:

```go
type Exchange interface {
    // Metadata
    GetInfo() ExchangeInfo
    GetType() ExchangeType
    IsConnected() bool

    // Market Data
    GetTicker(ctx context.Context, symbol Symbol) (*Ticker, error)
    GetOrderBook(ctx context.Context, symbol Symbol, limit int) (*OrderBook, error)
    GetSymbols(ctx context.Context) ([]Symbol, error)

    // Account
    GetBalances(ctx context.Context) ([]Balance, error)
    GetBalance(ctx context.Context, asset string) (*Balance, error)

    // Trading
    PlaceOrder(ctx context.Context, order *Order) (*Order, error)
    GetOrder(ctx context.Context, orderID string) (*Order, error)
    CancelOrder(ctx context.Context, orderID string) error

    // WebSocket
    SubscribeToTickers(ctx context.Context, symbols []Symbol, callback TickerCallback) error
    SubscribeToOrderBook(ctx context.Context, symbol Symbol, callback OrderBookCallback) error
    UnsubscribeAll(ctx context.Context) error

    // Connection management
    Connect(ctx context.Context) error
    Disconnect(ctx context.Context) error
    Ping(ctx context.Context) error
}
```

### 2. Factory Pattern (`internal/exchange/factory.go`)

The factory creates exchange instances based on configuration:

```go
func (f *factory) CreateExchange(exchangeType ExchangeType, config ExchangeConfig) (Exchange, error) {
    switch exchangeType {
    case ExchangeTypeBinance:
        return NewBinanceExchange(config, f.logger)
    case ExchangeTypeKucoin:
        return NewKucoinExchange(config, f.logger)
    // Add new exchanges here
    default:
        return nil, fmt.Errorf("unsupported exchange type: %s", exchangeType)
    }
}
```

### 3. Registry Pattern (`internal/exchange/registry.go`)

The registry manages multiple exchange instances and provides lifecycle management.

## Adding a New Exchange

Here's a step-by-step guide to add support for **OKX Exchange**:

### Step 1: Define Exchange Type

Add the new exchange type to `internal/exchange/types.go`:

```go
const (
    ExchangeTypeBinance ExchangeType = "binance"
    ExchangeTypeKucoin  ExchangeType = "kucoin"
    ExchangeTypeCoinbase ExchangeType = "coinbase"
    ExchangeTypeKraken  ExchangeType = "kraken"
    ExchangeTypeOKX     ExchangeType = "okx"     // ← Add this
)
```

### Step 2: Update Configuration

Add OKX configuration support in `internal/config/config.go`:

```go
func loadExchangesConfig() exchange.ExchangesConfig {
    exchanges := []exchange.ExchangeConfig{
        // ... existing exchanges ...
        {
            Type:      exchange.ExchangeTypeOKX,
            Name:      "OKX",
            APIKey:    getEnv("OKX_API_KEY", ""),
            SecretKey: getEnv("OKX_SECRET_KEY", ""),
            Passphrase: getEnv("OKX_PASSPHRASE", ""),
            Sandbox:   getBoolEnv("OKX_SANDBOX", false),
            RateLimit: getIntEnv("OKX_RATE_LIMIT", 600),
            Timeout:   getIntEnv("OKX_TIMEOUT", 30),
            Enabled:   getEnv("OKX_API_KEY", "") != "",
        },
    }
    // ...
}
```

### Step 3: Implement Exchange

Create `internal/exchange/okx.go`:

```go
package exchange

import (
    "context"
    "fmt"
    "math/big"
    "time"
    
    "github.com/haydenfowler/delta-bot/internal/logger"
)

type okxExchange struct {
    config     ExchangeConfig
    logger     *logger.Logger
    connected  bool
    // Add OKX-specific client fields
}

func NewOKXExchange(config ExchangeConfig, logger *logger.Logger) (Exchange, error) {
    if config.APIKey == "" || config.SecretKey == "" {
        return nil, fmt.Errorf("OKX API credentials are required")
    }

    return &okxExchange{
        config:    config,
        logger:    logger,
        connected: false,
    }, nil
}

// Implement all Exchange interface methods
func (o *okxExchange) GetInfo() ExchangeInfo {
    return ExchangeInfo{
        Name:    "OKX",
        Type:    ExchangeTypeOKX,
        Symbols: []Symbol{"BTC-USDT", "ETH-USDT", "BTC-ETH"}, // OKX uses dashes
        MinOrderSize: map[Symbol]string{
            "BTC-USDT": "0.00001",
            "ETH-USDT": "0.001",
            "BTC-ETH":  "0.001",
        },
        Fees: TradingFees{
            MakerFee: big.NewFloat(0.0008), // 0.08%
            TakerFee: big.NewFloat(0.001),  // 0.1%
        },
    }
}

func (o *okxExchange) GetType() ExchangeType {
    return ExchangeTypeOKX
}

func (o *okxExchange) IsConnected() bool {
    return o.connected
}

func (o *okxExchange) Connect(ctx context.Context) error {
    // Implement OKX connection logic
    o.connected = true
    o.logger.Info("Connected to OKX")
    return nil
}

func (o *okxExchange) Disconnect(ctx context.Context) error {
    // Implement OKX disconnection logic
    o.connected = false
    o.logger.Info("Disconnected from OKX")
    return nil
}

// Implement remaining interface methods...
func (o *okxExchange) GetTicker(ctx context.Context, symbol Symbol) (*Ticker, error) {
    // Implement OKX ticker API call
    return &Ticker{
        Symbol:    symbol,
        BidPrice:  big.NewFloat(45000.0),  // Replace with actual API call
        AskPrice:  big.NewFloat(45001.0),
        LastPrice: big.NewFloat(45000.5),
        Volume:    big.NewFloat(1234.567),
        Timestamp: time.Now(),
    }, nil
}

// ... implement all other interface methods
```

### Step 4: Update Factory

Add OKX to the factory in `internal/exchange/factory.go`:

```go
func (f *factory) CreateExchange(exchangeType ExchangeType, config ExchangeConfig) (Exchange, error) {
    switch exchangeType {
    case ExchangeTypeBinance:
        return NewBinanceExchange(config, f.logger)
    case ExchangeTypeOKX:
        return NewOKXExchange(config, f.logger)  // ← Add this
    // ... other cases
    default:
        return nil, fmt.Errorf("unsupported exchange type: %s", exchangeType)
    }
}

func (f *factory) SupportedExchanges() []ExchangeType {
    return []ExchangeType{
        ExchangeTypeBinance,
        ExchangeTypeKucoin,
        ExchangeTypeCoinbase,
        ExchangeTypeKraken,
        ExchangeTypeOKX,  // ← Add this
    }
}
```

### Step 5: Update Environment Configuration

Add to `.env.example`:

```bash
# OKX
OKX_API_KEY=your_okx_api_key
OKX_SECRET_KEY=your_okx_secret_key
OKX_PASSPHRASE=your_okx_passphrase
OKX_SANDBOX=false
OKX_RATE_LIMIT=600
OKX_TIMEOUT=30
```

### Step 6: Test the Implementation

Create tests in `internal/exchange/okx_test.go`:

```go
package exchange

import (
    "context"
    "testing"
    
    "github.com/haydenfowler/delta-bot/internal/logger"
)

func TestOKXExchange(t *testing.T) {
    logger := logger.New("", "test")
    config := ExchangeConfig{
        Type:      ExchangeTypeOKX,
        APIKey:    "test_key",
        SecretKey: "test_secret",
        Enabled:   true,
    }

    exchange, err := NewOKXExchange(config, logger)
    if err != nil {
        t.Fatalf("Failed to create OKX exchange: %v", err)
    }

    if exchange.GetType() != ExchangeTypeOKX {
        t.Errorf("Expected OKX exchange type, got %s", exchange.GetType())
    }

    // Test connection
    ctx := context.Background()
    if err := exchange.Connect(ctx); err != nil {
        t.Errorf("Failed to connect: %v", err)
    }

    if !exchange.IsConnected() {
        t.Error("Exchange should be connected")
    }
}
```

## Real-World Implementation Tips

### 1. API Client Integration

For real exchanges, you'll typically integrate their official SDK:

```go
import "github.com/okx/okx-go-sdk"

type okxExchange struct {
    client    *okx.Client
    // ... other fields
}

func NewOKXExchange(config ExchangeConfig, logger *logger.Logger) (Exchange, error) {
    client := okx.NewClient(config.APIKey, config.SecretKey, config.Passphrase)
    
    return &okxExchange{
        client: client,
        // ... initialize other fields
    }, nil
}
```

### 2. Error Handling

Implement robust error handling and retries:

```go
func (o *okxExchange) GetTicker(ctx context.Context, symbol Symbol) (*Ticker, error) {
    var ticker *Ticker
    var err error
    
    for attempts := 0; attempts < 3; attempts++ {
        ticker, err = o.fetchTickerWithRetry(ctx, symbol)
        if err == nil {
            break
        }
        
        if isRateLimitError(err) {
            time.Sleep(time.Duration(attempts+1) * time.Second)
            continue
        }
        
        return nil, fmt.Errorf("failed to get ticker after %d attempts: %w", attempts+1, err)
    }
    
    return ticker, nil
}
```

### 3. Rate Limiting

Implement proper rate limiting:

```go
import "golang.org/x/time/rate"

type okxExchange struct {
    rateLimiter *rate.Limiter
    // ... other fields
}

func (o *okxExchange) makeAPICall(ctx context.Context) error {
    if err := o.rateLimiter.Wait(ctx); err != nil {
        return fmt.Errorf("rate limit wait failed: %w", err)
    }
    
    // Make actual API call
    return nil
}
```

### 4. WebSocket Implementation

For real-time data:

```go
func (o *okxExchange) SubscribeToTickers(ctx context.Context, symbols []Symbol, callback TickerCallback) error {
    conn, err := o.connectWebSocket(ctx)
    if err != nil {
        return err
    }
    
    // Subscribe to ticker updates
    for _, symbol := range symbols {
        if err := o.subscribeToSymbol(conn, symbol); err != nil {
            return err
        }
    }
    
    // Handle incoming messages in goroutine
    go o.handleWebSocketMessages(conn, callback)
    
    return nil
}
```

## Benefits of This Architecture

### 1. **Open/Closed Principle**
- ✅ Open for extension (add new exchanges)
- ✅ Closed for modification (existing code unchanged)

### 2. **Zero Impact on Existing Code**
- ✅ No changes to arbitrage logic
- ✅ No changes to execution engine
- ✅ No changes to API endpoints

### 3. **Consistent Interface**
- ✅ All exchanges work the same way
- ✅ Easy to switch between exchanges
- ✅ Uniform error handling

### 4. **Testing Friendly**
- ✅ Mock exchanges for testing
- ✅ Each exchange can be tested independently
- ✅ Integration tests work with any exchange

### 5. **Configuration Driven**
- ✅ Enable/disable exchanges via environment variables
- ✅ No code changes needed for different deployments
- ✅ Easy to manage multiple exchange credentials

## Example Usage

Once implemented, using multiple exchanges is automatic:

```bash
# Enable multiple exchanges
BINANCE_API_KEY=xxx
KUCOIN_API_KEY=yyy  
OKX_API_KEY=zzz

# Start the bot
make run

# The bot will automatically:
# 1. Connect to all configured exchanges
# 2. Detect arbitrage opportunities across all exchanges
# 3. Execute trades on the most profitable exchange
```

The arbitrage detection and execution logic works seamlessly across all exchanges without any modifications!

## Conclusion

This extensible architecture demonstrates how proper interface design and the Open/Closed Principle create maintainable, scalable systems. Adding a new exchange requires only:

1. ✅ Implementing the `Exchange` interface
2. ✅ Adding one case to the factory
3. ✅ Adding configuration variables

**Zero modifications** to existing arbitrage, execution, or API code! 🚀