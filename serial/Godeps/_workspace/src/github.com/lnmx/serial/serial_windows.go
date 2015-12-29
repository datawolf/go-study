package serial

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

type Port struct {
	config Config
	handle syscall.Handle
}

func (p *Port) configure(cfg Config) (err error) {
	p.config = cfg

	return
}

func (p *Port) open() (err error) {

	// open serial port handle
	//
	err = p.openHandle()

	if err != nil {
		return fmt.Errorf("error opening serial device %s: %s", p.config.Device, err)
	}

	// configure baud, data, parity
	//
	err = p.setCommState()

	if err != nil {
		return fmt.Errorf("error applying serial settings: %s", err)
	}

	// set read/write timeouts
	//
	err = p.setTimeouts()

	if err != nil {
		return fmt.Errorf("error setting timeouts: %s", err)
	}

	// reset
	//
	err = p.reset()

	if err != nil {
		return fmt.Errorf("error during reset: %s", err)
	}

	return
}

func (p *Port) openHandle() (err error) {
	device := p.config.Device

	// add a "\\.\" prefix, e.g. "\\.\COM2"
	// optional for COM1-9, required for COM10 and up
	//
	if !strings.HasPrefix(device, `\\.\`) {
		device = `\\.\` + device
	}

	device_utf16, err := syscall.UTF16PtrFromString(device)

	if err != nil {
		return
	}

	p.handle, err = syscall.CreateFile(device_utf16,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		0,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0)

	return
}

func (p *Port) setCommState() (err error) {

	dcb := &_DCB{
		DCBlength: 28,
	}

	err = _GetCommState(p.handle, dcb)

	if err != nil {
		return
	}

	f := _DCB_Flags{
		Binary:           true,
		DtrControl:       0x01, // enable
		DsrSensitivity:   false,
		TXContinueOnXoff: false,
		OutX:             false,
		InX:              false,
		ErrorChar:        false,
		Null:             false,
		RtsControl:       0x01, // enable
		AbortOnError:     false,
		OutxCtsFlow:      false,
		OutxDsrFlow:      false,
	}

	dcb.Flags = f.Get()

	if v, ok := settings[p.config.BaudRate]; ok {
		dcb.BaudRate = uint32(v)
	} else {
		return fmt.Errorf("unsupported baud rate: %d", p.config.BaudRate)
	}

	if v, ok := settings[p.config.DataBits]; ok {
		dcb.ByteSize = byte(v)
	} else {
		return fmt.Errorf("unsupported data bits: %d", p.config.DataBits)
	}

	if v, ok := settings[p.config.Parity]; ok {
		dcb.Parity = byte(v)
	} else {
		return fmt.Errorf("unsupported parity: %d", p.config.Parity)
	}

	if v, ok := settings[p.config.StopBits]; ok {
		dcb.StopBits = byte(v)
	} else {
		return fmt.Errorf("unsupported stop bits: %d", p.config.StopBits)
	}

	err = _SetCommState(p.handle, dcb)

	return
}

const MAXDWORD = 0xffffffff

func (p *Port) setTimeouts() (err error) {

	timeouts := &_COMMTIMEOUTS{
		// ms to wait for next byte
		ReadIntervalTimeout: p.config.ReadIntervalTimeout,

		// read total timeout = constant + (multiplier * byte count)
		// zero for both disables total time-outs
		ReadTotalTimeoutConstant:   p.config.ReadTotalTimeoutConstant,
		ReadTotalTimeoutMultiplier: p.config.ReadTotalTimeoutMultiplier,

		// write total timeout = constant + (multiplier * byte count)
		// zero for both disables total time-outs
		WriteTotalTimeoutConstant:   p.config.WriteTotalTimeoutConstant,
		WriteTotalTimeoutMultiplier: p.config.WriteTotalTimeoutMultiplier,
	}

	err = _SetCommTimeouts(p.handle, timeouts)

	if err != nil {
		return
	}

	// Verify that the COMMTIMEOUTS were applied.  At least one USB-to-Serial
	// driver was found to report success after silently overwriting some
	// values.
	//
	readback := new(_COMMTIMEOUTS)

	err = _GetCommTimeouts(p.handle, readback)

	if err != nil {
		return
	}

	if *timeouts != *readback {
		err = fmt.Errorf("timeout settings overridden by serial device:\nrequested:\n%v\nactual:\n%v", *timeouts, *readback)

		return
	}

	return
}

func (p *Port) reset() (err error) {

	err = syscall.FlushFileBuffers(p.handle)

	if err != nil {
		return
	}

	err = _PurgeComm(p.handle, _PURGE_TXABORT|_PURGE_RXABORT|_PURGE_TXCLEAR|_PURGE_RXCLEAR)

	if err != nil {
		return
	}

	err = _ClearCommError(p.handle)

	if err != nil {
		return
	}

	return
}

func (p *Port) close() (err error) {
	return syscall.CloseHandle(p.handle)
}

func (p *Port) read(b []byte) (n int, err error) {
	var done uint32
	err = syscall.ReadFile(p.handle, b, &done, nil)
	n = int(done)

	return
}

func (p *Port) write(b []byte) (n int, err error) {
	for {
		var done uint32
		err = syscall.WriteFile(p.handle, b[n:], &done, nil)
		n += int(done)

		if err != nil || n == len(b) {
			break
		}
	}

	if err != nil {
		err = p.flush()
	}

	return
}

func (p *Port) flush() (err error) {
	return syscall.FlushFileBuffers(p.handle)
}

func (p *Port) signal(s Signal, value bool) (err error) {
	switch {

	// DTR
	//
	case s == DTR && value == false:
		return _EscapeCommFunction(p.handle, _CLRDTR)
	case s == DTR && value == true:
		return _EscapeCommFunction(p.handle, _SETDTR)

	// RTS
	//
	case s == RTS && value == false:
		return _EscapeCommFunction(p.handle, _CLRRTS)
	case s == RTS && value == true:
		return _EscapeCommFunction(p.handle, _SETRTS)

	default:
		return fmt.Errorf("Unreconized signal: %v %v", s, value)
	}
}

var (
	// handles to some serial-related functions not supported by syscall
	//
	kernel32, _           = syscall.LoadLibrary("kernel32.dll")
	getCommState, _       = syscall.GetProcAddress(kernel32, "GetCommState")
	setCommState, _       = syscall.GetProcAddress(kernel32, "SetCommState")
	setCommTimeouts, _    = syscall.GetProcAddress(kernel32, "SetCommTimeouts")
	getCommTimeouts, _    = syscall.GetProcAddress(kernel32, "GetCommTimeouts")
	escapeCommFunction, _ = syscall.GetProcAddress(kernel32, "EscapeCommFunction")
	flushFileBuffers, _   = syscall.GetProcAddress(kernel32, "FlushFileBuffers")
	purgeComm, _          = syscall.GetProcAddress(kernel32, "PurgeComm")
	clearCommError, _     = syscall.GetProcAddress(kernel32, "ClearCommError")

	// map the generic constants in serial.go to values for DCB
	// (ref. WinBase.h)
	//
	settings = map[int]int{
		DataBits_5:      5,
		DataBits_6:      6,
		DataBits_7:      7,
		DataBits_8:      8,
		StopBits_1:      0,
		StopBits_1_5:    1,
		StopBits_2:      2,
		Parity_None:     0,
		Parity_Odd:      1,
		Parity_Even:     2,
		Parity_Mark:     3,
		Parity_Space:    4,
		BaudRate_9600:   9600,
		BaudRate_19200:  19200,
		BaudRate_38400:  38400,
		BaudRate_57600:  57600,
		BaudRate_115200: 115200,
	}
)

type _DCB_Flags struct {
	Binary           bool
	Parity           bool
	OutxCtsFlow      bool
	OutxDsrFlow      bool
	DtrControl       byte
	DsrSensitivity   bool
	TXContinueOnXoff bool
	OutX             bool
	InX              bool
	ErrorChar        bool
	Null             bool
	RtsControl       byte
	AbortOnError     bool
	// 17 reserved bits
}

func (f *_DCB_Flags) Set(v uint32) {

	bit := func(n uint32) bool {
		var mask uint32 = 1 << n

		return (mask & v) != 0
	}

	bit2 := func(n uint32) byte {
		x := v >> n
		x = x & 0x3

		return byte(x)
	}

	f.Binary = bit(0)
	f.Parity = bit(1)
	f.OutxCtsFlow = bit(2)
	f.OutxDsrFlow = bit(3)
	f.DtrControl = bit2(4)
	f.DsrSensitivity = bit(6)
	f.TXContinueOnXoff = bit(7)
	f.OutX = bit(8)
	f.InX = bit(9)
	f.ErrorChar = bit(10)
	f.Null = bit(11)
	f.RtsControl = bit2(12)
	f.AbortOnError = bit(14)
}

func (f _DCB_Flags) Get() (v uint32) {

	bit := func(b bool, n uint32) (v uint32) {
		if b {
			v = 1
		}

		v = v << n

		return
	}

	bit2 := func(b byte, n uint32) (v uint32) {
		v = uint32(b & 0x03)

		v = v << n

		return
	}

	v = v | bit(f.Binary, 0)
	v = v | bit(f.Parity, 1)
	v = v | bit(f.OutxCtsFlow, 2)
	v = v | bit(f.OutxDsrFlow, 3)
	v = v | bit2(f.DtrControl, 4)
	v = v | bit(f.DsrSensitivity, 6)
	v = v | bit(f.TXContinueOnXoff, 7)
	v = v | bit(f.OutX, 8)
	v = v | bit(f.InX, 9)
	v = v | bit(f.ErrorChar, 10)
	v = v | bit(f.Null, 11)
	v = v | bit2(f.RtsControl, 12)
	v = v | bit(f.AbortOnError, 14)

	return
}

// sizeof(DCB) = 28 bytes

type _DCB struct {
	DCBlength uint32
	BaudRate  uint32

	// Flags:
	//   1 binary
	//   1 enable parity
	//   1 out cts flow control
	//   1 out dtr flow control
	//   2 dtr flow control
	//   1 dsr sensitivity
	//   1 continue tx on xoff
	//   1 enable output x-on/x-off
	//   1 enable input x-on/x-off
	//   1 enable err replacement
	//   1 enable null stripping
	//   2 rts flow control
	//   1 abort reads/writes on error
	//  17 reserved
	Flags     uint32
	Reserved1 uint16
	XonLim    uint16
	XoffLim   uint16
	ByteSize  byte
	Parity    byte
	StopBits  byte
	XonChar   byte
	XoffChar  byte
	ErrorChar byte
	EofChar   byte
	EvtChar   byte
	Reserved2 uint16
}

// COMMTIMEOUTS
//
type _COMMTIMEOUTS struct {
	ReadIntervalTimeout         uint32
	ReadTotalTimeoutMultiplier  uint32
	ReadTotalTimeoutConstant    uint32
	WriteTotalTimeoutMultiplier uint32
	WriteTotalTimeoutConstant   uint32
}

func (c _COMMTIMEOUTS) String() string {
	return fmt.Sprintf(
		"\n  ReadIntervalTimeout:         %08x"+
			"\n  ReadTotalTimeoutMultiplier:  %08x"+
			"\n  ReadTotalTimeoutConstant:    %08x"+
			"\n  WriteTotalTimeoutMultiplier: %08x"+
			"\n  WriteTotalTimeoutConstant:   %08x"+
			"\n",
		c.ReadIntervalTimeout, c.ReadTotalTimeoutMultiplier, c.ReadTotalTimeoutConstant,
		c.WriteTotalTimeoutMultiplier, c.WriteTotalTimeoutConstant)
}

func _GetCommTimeouts(handle syscall.Handle, timeouts *_COMMTIMEOUTS) (err error) {
	// BOOL GetCommTimeouts( HANDLE hFile, LPCOMMTIMEOUTS lpCommTimeouts);

	r0, _, e1 := syscall.Syscall(getCommTimeouts, 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

func _SetCommTimeouts(handle syscall.Handle, timeouts *_COMMTIMEOUTS) (err error) {
	// BOOL SetCommTimeouts( HANDLE hFile, LPCOMMTIMEOUTS lpCommTimeouts);

	r0, _, e1 := syscall.Syscall(setCommTimeouts, 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

func _GetCommState(handle syscall.Handle, dcb *_DCB) (err error) {
	// BOOL GetCommState( HANDLE hFile, LPDCB lpDCB );

	r0, _, e1 := syscall.Syscall(getCommState, 2, uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

func _SetCommState(handle syscall.Handle, dcb *_DCB) (err error) {
	// BOOL SetCommState( HANDLE hFile, LPDCB lpDCB );

	r0, _, e1 := syscall.Syscall(setCommState, 2, uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

type purgeFlag int

const (
	_PURGE_TXABORT purgeFlag = 0x01
	_PURGE_RXABORT           = 0x02
	_PURGE_TXCLEAR           = 0x04
	_PURGE_RXCLEAR           = 0x08
)

func _PurgeComm(handle syscall.Handle, purge purgeFlag) (err error) {
	// BOOL PurgeComm( HANDLE hFile, DWORD dwFlags )

	r0, _, e1 := syscall.Syscall(purgeComm, 2, uintptr(handle), uintptr(purge), 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

// operations for EscapeCommFunction
//
type escapeFn int

const (
	_SETXOFF  escapeFn = 1
	_SETXON            = 2
	_SETRTS            = 3
	_CLRRTS            = 4
	_SETDTR            = 5
	_CLRDTR            = 6
	_RESETDEV          = 7
	_SETBREAK          = 8
	_CLRBREAK          = 9
)

func _EscapeCommFunction(handle syscall.Handle, escape escapeFn) (err error) {
	// BOOL EscapeCommFunction( HANDLE hFile, DWORD dwFunc );

	r0, _, e1 := syscall.Syscall(escapeCommFunction, 2, uintptr(handle), uintptr(escape), 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

func _ClearCommError(handle syscall.Handle) (err error) {
	// BOOL ClearCommError( HANDLE hFile, LPDWORD lpErrors, LPCOMSTAT lpStat )

	r0, _, e1 := syscall.Syscall(clearCommError, 3, uintptr(handle), 0, 0)

	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return

}
