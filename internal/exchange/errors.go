package exchange

import "fmt"

// UnsupportedExchangeError represents an unsupported exchange type error
type UnsupportedExchangeError struct {
	ExchangeType ExchangeType
}

func (e *UnsupportedExchangeError) Error() string {
	return fmt.Sprintf("unsupported exchange type: %s", e.ExchangeType)
}

// NewUnsupportedExchangeError creates a new unsupported exchange error
func NewUnsupportedExchangeError(exchangeType ExchangeType) *UnsupportedExchangeError {
	return &UnsupportedExchangeError{ExchangeType: exchangeType}
}
