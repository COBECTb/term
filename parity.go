package term

import (
	"errors"
	"syscall"

	"github.com/pkg/term/termios"
)

//#define CMSPAR   010000000000
//import "C"

const (
	//ParNONE paryty off
	ParNONE = iota
	//ParEVEN paryty
	ParEVEN
	//ParODD parity is ODD
	ParODD
	//ParMARK parity is MARK
	ParMARK
	//ParSPACE parity is SPACE
	ParSPACE
	//ParIGN parity is ParIGN
	ParIGN
)

const CMSPAR uint64 = 0x40000000

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
		a.Cflag &^= CMSPAR
	case ParEVEN:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag &^= CMSPAR
		a.Cflag |= syscall.PARENB
		a.Cflag &^= syscall.PARODD
	case ParODD:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag &^= CMSPAR
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
	case ParMARK:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag |= syscall.PARENB
		a.Cflag |= syscall.PARODD
		a.Cflag |= CMSPAR
	case ParSPACE:
		a.Cflag &^= syscall.IGNPAR
		a.Cflag |= syscall.PARENB
		a.Cflag |= CMSPAR
		a.Cflag &^= syscall.PARODD
	case ParIGN:
		a.Cflag |= syscall.IGNPAR
		a.Cflag &^= syscall.PARENB
		a.Cflag &^= syscall.PARODD
		a.Cflag &^= CMSPAR
	default:
		return errors.New("Unknown parity option")
	}
	return termios.Tcsetattr(uintptr(t.fd), termios.TCSANOW, (*syscall.Termios)(&a))
}

//SetRecvParityCheckOn - parity
func (t *Term) SetRecvParityCheckOn(onIgn bool) error {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return err
	}
	if onIgn {
		a.Cflag |= syscall.IGNPAR
	} else {
		a.Cflag &^= syscall.IGNPAR
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
			return ParEVEN, nil
		}
	} else {
		return ParNONE, nil
	}
}

//config.c_cflag &= ~(CSIZE | PARENB);
//config.c_cflag |= CS8;
