// +build mips mipsle mips64 mips64le

package serial

import (
	"syscall"
)

func cfSetIspeed(termios *syscall.Termios, speed uint64) {
	// no op
}

func cfSetOspeed(termios *syscall.Termios, speed uint64) {
	// no op
}
