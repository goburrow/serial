package serial

// cribbed from: https://github.com/jacobsa/go-serial/blob/master/serial/open_linux.go#L172-L176

import (
	"errors"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	rs485_enabled        = (1 << 0)
	rs485_rts_on_send    = (1 << 1)
	rs485_rts_after_send = (1 << 2)
	rs485_rx_during_tx   = (1 << 4)
	tiocsrs485           = 0x542f
)

type rs485_ioctl_opts struct {
	flags                 uint32
	delay_rts_before_send uint32
	delay_rts_after_send  uint32
	padding               [5]uint32
}

func enable_rs485(fd int, config *RS485Config) error {
	if config.Enable {
		rs485 := rs485_ioctl_opts{
			rs485_enabled,
			uint32(config.DelayRtsBeforeSend / time.Millisecond),
			uint32(config.DelayRtsAfterSend / time.Millisecond),
			[5]uint32{0, 0, 0, 0, 0},
		}

		if config.RtsHighDuringSend {
			rs485.flags |= rs485_rts_on_send
		}

		if config.RtsHighAfterSend {
			rs485.flags |= rs485_rts_after_send
		}

		if config.RxDuringTx {
			rs485.flags |= rs485_rx_during_tx
		}

		r, _, errno := syscall.Syscall(
			syscall.SYS_IOCTL,
			uintptr(fd),
			uintptr(tiocsrs485),
			uintptr(unsafe.Pointer(&rs485)))

		if errno != 0 {
			return os.NewSyscallError("SYS_IOCTL (RS485)", errno)
		}

		if r != 0 {
			return errors.New("Unknown error from SYS_IOCTL (RS485)")
		}
	}
	return nil
}
