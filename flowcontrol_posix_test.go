//go:build darwin || linux || freebsd || openbsd || netbsd
// +build darwin linux freebsd openbsd netbsd

package serial

import (
	"syscall"
	"testing"

	"golang.org/x/sys/unix"
)

func TestNewTermiosFlowControl(t *testing.T) {
	tests := []struct {
		name        string
		flowControl FlowControl
		rtscts      bool
		xonxoff     bool
	}{
		{"none", FlowControlNone, false, false},
		{"rts/cts", FlowControlRTSCTS, true, false},
		{"xon/xoff", FlowControlXONXOFF, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			termios, err := newTermios(&Config{FlowControl: tt.flowControl})
			if err != nil {
				t.Fatal(err)
			}
			checkFlowControl(t, termios, tt.rtscts, tt.xonxoff)
		})
	}
}

func TestNewTermiosUnsupportedFlowControl(t *testing.T) {
	if _, err := newTermios(&Config{FlowControl: FlowControl(99)}); err == nil {
		t.Fatal("expected an error for an unknown flow control")
	}
}

func checkFlowControl(t *testing.T, termios *syscall.Termios, rtscts, xonxoff bool) {
	t.Helper()
	if got := termios.Cflag&unix.CRTSCTS == unix.CRTSCTS; got != rtscts {
		t.Errorf("CRTSCTS set = %v, want %v (cflag %#x)", got, rtscts, termios.Cflag)
	}
	if got := termios.Iflag&(syscall.IXON|syscall.IXOFF) == syscall.IXON|syscall.IXOFF; got != xonxoff {
		t.Errorf("IXON|IXOFF set = %v, want %v (iflag %#x)", got, xonxoff, termios.Iflag)
	}
	if xonxoff && (termios.Cc[unix.VSTART] != 0x11 || termios.Cc[unix.VSTOP] != 0x13) {
		t.Errorf("VSTART/VSTOP = %#x/%#x, want 0x11/0x13", termios.Cc[unix.VSTART], termios.Cc[unix.VSTOP])
	}
}
