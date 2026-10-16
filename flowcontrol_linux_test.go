package serial

import (
	"os"
	"strconv"
	"syscall"
	"testing"

	"golang.org/x/sys/unix"
)

// TestOpenSetsFlowControlOnPty opens the slave side of a pseudo-terminal
// through Open and reads its termios back from the kernel, so it proves the
// flags reach the device rather than only the struct newTermios returns.
func TestOpenSetsFlowControlOnPty(t *testing.T) {
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
			slave := openPty(t)
			p, err := Open(&Config{Address: slave, BaudRate: 9600, FlowControl: tt.flowControl})
			if err != nil {
				t.Fatal(err)
			}
			defer p.Close()
			termios := &syscall.Termios{}
			if err := tcgetattr(p.(*port).fd, termios); err != nil {
				t.Fatal(err)
			}
			checkFlowControl(t, termios, tt.rtscts, tt.xonxoff)
		})
	}
}

// openPty opens a new pseudo-terminal master and returns its slave's path.
func openPty(t *testing.T) string {
	t.Helper()
	master, err := os.OpenFile("/dev/ptmx", os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		t.Skipf("no pseudo-terminal: %v", err)
	}
	t.Cleanup(func() { master.Close() })
	fd := int(master.Fd())
	if err := unix.IoctlSetPointerInt(fd, unix.TIOCSPTLCK, 0); err != nil {
		t.Fatal(err)
	}
	n, err := unix.IoctlGetInt(fd, unix.TIOCGPTN)
	if err != nil {
		t.Fatal(err)
	}
	return "/dev/pts/" + strconv.Itoa(n)
}
