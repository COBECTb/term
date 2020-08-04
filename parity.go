package term

import (
	"errors"
	"syscall"

	"github.com/pkg/term/termios"
)

const (
	//ParNONE paryty off
	ParNONE = iota
	//PAR paryty
	PAR
	//ODD parity is ODD
	ODD
	//MARK parity is MARK
	MARK
	//SPACE parity is SPACE
	SPACE
)

const cmspar uint32 = 0x40000000

//SetParity - parity
func (t *Term) SetParity(parity int) error {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return err
	}
	switch parity {
	case ParNONE:
		a.Cflag &^= syscall.PARENB
		a.Cflag &^= cmspar
	case PAR:
		a.Cflag &^= cmspar
		a.Cflag |= syscall.PARENB
		a.Cflag &^= syscall.PARODD
	case ODD:
		a.Cflag &^= cmspar
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
	case MARK:
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
		a.Cflag |= cmspar
	case SPACE:
		a.Cflag |= syscall.PARENB
		a.Cflag |= cmspar
		a.Cflag &^= syscall.PARODD
	default:
		return errors.New("Unknown parity option")
	}
	return termios.Tcsetattr(uintptr(t.fd), termios.TCSANOW, (*syscall.Termios)(&a))
}

//GetParity - get parity
func (t *Term) GetParity() (int, error) {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return 0, err
	}

	if a.Cflag&syscall.PARENB > 0 {
		if a.Cflag&syscall.PARODD > 0 {
			return ODD, nil
		} else {
			return PAR, nil
		}
	} else {
		return ParNONE, nil
	}
}

//config.c_cflag &= ~(CSIZE | PARENB);
//config.c_cflag |= CS8;
