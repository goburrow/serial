package serial

import (
	"time"
)

// platform independent RS485 config. Thie structure is ignored unless Enable is true.
type RS485Config struct {
	// Enable RS485 support
	Enable bool
	// Delay RTS prior to send
	DelayRtsBeforeSend time.Duration
	// Delay RTS after send
	DelayRtsAfterSend time.Duration
	// Set RTS high during send
	RtsHighDuringSend bool
	// Set RTS high after send
	RtsHighAfterSend bool
	// Rx during Tx
	RxDuringTx bool
}
