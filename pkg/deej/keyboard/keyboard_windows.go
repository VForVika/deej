package keyboard

import (
	"syscall"
	"unsafe"
)

type KEYBDINPUT struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type INPUT struct {
	t        uint32
	keyboard KEYBDINPUT
	padding  uint64
}

var mod = syscall.NewLazyDLL("user32.dll")
var proc = mod.NewProc("SendInput")

func PlayPause() {

	inputs := [2]INPUT{
		{t: 1, keyboard: KEYBDINPUT{
			wVk: 0xB3,
		}},

		{t: 1, keyboard: KEYBDINPUT{
			wVk:     0xB3,
			dwFlags: 0x0002,
		}},
	}
	proc.Call(uintptr(len(inputs)), uintptr(unsafe.Pointer(&inputs)), unsafe.Sizeof(INPUT{}))
}

func Next() {

	inputs := [2]INPUT{
		{t: 1, keyboard: KEYBDINPUT{
			wVk: 0xB0,
		}},

		{t: 1, keyboard: KEYBDINPUT{
			wVk:     0xB0,
			dwFlags: 0x0002,
		}},
	}
	proc.Call(uintptr(len(inputs)), uintptr(unsafe.Pointer(&inputs)), unsafe.Sizeof(INPUT{}))
}

func Previous() {

	inputs := [2]INPUT{
		{t: 1, keyboard: KEYBDINPUT{
			wVk: 0xB1,
		}},

		{t: 1, keyboard: KEYBDINPUT{
			wVk:     0xB1,
			dwFlags: 0x0002,
		}},
	}
	proc.Call(uintptr(len(inputs)), uintptr(unsafe.Pointer(&inputs)), unsafe.Sizeof(INPUT{}))
}
