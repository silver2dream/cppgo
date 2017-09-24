package cdecl

import "errors"

func Call(addr uintptr, a ...uintptr) (uintptr, error) {
	switch l := len(a); l {
	case 0:
		return Call0(addr), nil
	case 1:
		return Call1(addr, a[0]), nil
	case 2:
		return Call2(addr, a[0], a[1]), nil
	case 3:
		return Call3(addr, a[0], a[1], a[2]), nil
	case 4:
		return Call4(addr, a[0], a[1], a[2], a[3]), nil
	case 5:
		return Call5(addr, a[0], a[1], a[2], a[3], a[4]), nil
	case 6:
		return Call6(addr, a[0], a[1], a[2], a[3], a[4], a[5]), nil
	default:
		return 0, errors.New("too many arguments")
	}
}