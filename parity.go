package term

import (
	"errors"
	"syscall"

	"github.com/pkg/term/termios"
)

const (
	//NONE paryty off
	parNONE = iota
	//PAR paryty
	PAR
	//ODD parity is ODD
	ODD
)

//SetParity - parity
func (t *Term) SetParity(parity int) error {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return err
	}
	switch parity {
	case parNONE:
		a.Cflag &^= syscall.PARENB
	case PAR:
		a.Cflag |= syscall.PARENB
		a.Cflag &^= syscall.PARODD
	case ODD:
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
	default:
		return errors.New("Unknown parity option")
	}
	return termios.Tcsetattr(uintptr(t.fd), termios.TCSANOW, (*syscall.Termios)(&a))
}

//config.c_cflag &= ~(CSIZE | PARENB);
//config.c_cflag |= CS8;
