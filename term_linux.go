package term

import (
	"errors"
	"syscall"

	"github.com/pkg/term/termios"
)

type attr syscall.Termios

func (a *attr) setSpeed(baud int) error {
	var rate uint32
	switch baud {
	case 50:
		rate = syscall.B50
	case 75:
		rate = syscall.B75
	case 110:
		rate = syscall.B110
	case 134:
		rate = syscall.B134
	case 150:
		rate = syscall.B150
	case 200:
		rate = syscall.B200
	case 300:
		rate = syscall.B300
	case 600:
		rate = syscall.B600
	case 1200:
		rate = syscall.B1200
	case 1800:
		rate = syscall.B1800
	case 2400:
		rate = syscall.B2400
	case 4800:
		rate = syscall.B4800
	case 9600:
		rate = syscall.B9600
	case 19200:
		rate = syscall.B19200
	case 38400:
		rate = syscall.B38400
	case 57600:
		rate = syscall.B57600
	case 115200:
		rate = syscall.B115200
	case 230400:
		rate = syscall.B230400
	case 460800:
		rate = syscall.B460800
	case 500000:
		rate = syscall.B500000
	case 576000:
		rate = syscall.B576000
	case 921600:
		rate = syscall.B921600
	case 1000000:
		rate = syscall.B1000000
	case 1152000:
		rate = syscall.B1152000
	case 1500000:
		rate = syscall.B1500000
	case 2000000:
		rate = syscall.B2000000
	case 2500000:
		rate = syscall.B2500000
	case 3000000:
		rate = syscall.B3000000
	case 3500000:
		rate = syscall.B3500000
	case 4000000:
		rate = syscall.B4000000
	default:
		return syscall.EINVAL
	}
	(*syscall.Termios)(a).Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	(*syscall.Termios)(a).Ispeed = rate
	(*syscall.Termios)(a).Ospeed = rate
	return nil
}

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

const CMSPAR uint32 = 0x40000000

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
