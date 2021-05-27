package serial

import (
	"syscall"
)

func cfSetIspeed(termios *syscall.Termios, speed uint32) {
	termios.Ispeed = int32(speed)
}

func cfSetOspeed(termios *syscall.Termios, speed uint32) {
	termios.Ospeed = int32(speed)
}
