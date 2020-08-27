package term

import (
	"errors"
	"syscall"

	"github.com/pkg/term/termios"
)

const (
	//ParNONE paryty off
	ParNONE = iota
	//ParPAR paryty
	ParPAR
	//ParODD parity is ODD
	ParODD
	//ParMARK parity is MARK
	ParMARK
	//ParSPACE parity is SPACE
	ParSPACE
	//ParIGN parity is ParIGN
	ParIGN
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
		a.Cflag &^= syscall.IGNPAR
		a.Cflag &^= syscall.PARENB
		a.Cflag &^= syscall.PARODD
		a.Cflag &^= cmspar
	case ParPAR:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag &^= cmspar
		a.Cflag |= syscall.PARENB
		a.Cflag &^= syscall.PARODD
	case ParODD:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag &^= cmspar
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
	case ParMARK:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
		a.Cflag |= cmspar
	case ParSPACE:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag |= syscall.PARENB
		a.Cflag |= cmspar
		a.Cflag &^= syscall.PARODD
	case ParIGN:
		a.Cflag |= syscall.IGNPAR
		a.Cflag &^= syscall.PARENB
		a.Cflag &^= syscall.PARODD
		a.Cflag &^= cmspar
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
			return ParODD, nil
		} else {
			return ParPAR, nil
		}
	} else {
		return ParNONE, nil
	}
}

//config.c_cflag &= ~(CSIZE | PARENB);
//config.c_cflag |= CS8;
