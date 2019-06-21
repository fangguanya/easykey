// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey

// These are virtual key codes.
// See also, http://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
const (
	LBUTTON             = uint32(0x01)
	RBUTTON             = uint32(0x02)
	CANCEL              = uint32(0x03)
	MBUTTON             = uint32(0x04)
	XBUTTON1            = uint32(0x05)
	XBUTTON2            = uint32(0x06)
	BACK                = uint32(0x08)
	TAB                 = uint32(0x09)
	CLEAR               = uint32(0x0C)
	RETURN              = uint32(0x0D)
	SHIFT               = uint32(0x10)
	CONTROL             = uint32(0x11)
	MENU                = uint32(0x12)
	PAUSE               = uint32(0x13)
	CAPITAL             = uint32(0x14)
	KANA                = uint32(0x15)
	HANGUEL             = uint32(0x15)
	HANGUL              = uint32(0x15)
	JUNJA               = uint32(0x17)
	FINAL               = uint32(0x18)
	HANJA               = uint32(0x19)
	KANJI               = uint32(0x19)
	ESCAPE              = uint32(0x1B)
	CONVERT             = uint32(0x1C)
	NONCONVERT          = uint32(0x1D)
	ACCEPT              = uint32(0x1E)
	MODECHANGE          = uint32(0x1F)
	SPACE               = uint32(0x20)
	PRIOR               = uint32(0x21)
	NEXT                = uint32(0x22)
	END                 = uint32(0x23)
	HOME                = uint32(0x24)
	LEFT                = uint32(0x25)
	UP                  = uint32(0x26)
	RIGHT               = uint32(0x27)
	DOWN                = uint32(0x28)
	SELECT              = uint32(0x29)
	PRINT               = uint32(0x2A)
	EXECUTE             = uint32(0x2B)
	SNAPSHOT            = uint32(0x2C)
	INSERT              = uint32(0x2D)
	DELETE              = uint32(0x2E)
	HELP                = uint32(0x2F)
	LWIN                = uint32(0x5B)
	RWIN                = uint32(0x5C)
	APPS                = uint32(0x5D)
	SLEEP               = uint32(0x5F)
	NUMPAD0             = uint32(0x60)
	NUMPAD1             = uint32(0x61)
	NUMPAD2             = uint32(0x62)
	NUMPAD3             = uint32(0x63)
	NUMPAD4             = uint32(0x64)
	NUMPAD5             = uint32(0x65)
	NUMPAD6             = uint32(0x66)
	NUMPAD7             = uint32(0x67)
	NUMPAD8             = uint32(0x68)
	NUMPAD9             = uint32(0x69)
	MULTIPLY            = uint32(0x6A)
	ADD                 = uint32(0x6B)
	SEPARATOR           = uint32(0x6C)
	SUBTRACT            = uint32(0x6D)
	DECIMAL             = uint32(0x6E)
	DIVIDE              = uint32(0x6F)
	F1                  = uint32(0x70)
	F2                  = uint32(0x71)
	F3                  = uint32(0x72)
	F4                  = uint32(0x73)
	F5                  = uint32(0x74)
	F6                  = uint32(0x75)
	F7                  = uint32(0x76)
	F8                  = uint32(0x77)
	F9                  = uint32(0x78)
	F10                 = uint32(0x79)
	F11                 = uint32(0x7A)
	F12                 = uint32(0x7B)
	F13                 = uint32(0x7C)
	F14                 = uint32(0x7D)
	F15                 = uint32(0x7E)
	F16                 = uint32(0x7F)
	F17                 = uint32(0x80)
	F18                 = uint32(0x81)
	F19                 = uint32(0x82)
	F20                 = uint32(0x83)
	F21                 = uint32(0x84)
	F22                 = uint32(0x85)
	F23                 = uint32(0x86)
	F24                 = uint32(0x87)
	NUMLOCK             = uint32(0x90)
	SCROLL              = uint32(0x91)
	LSHIFT              = uint32(0xA0)
	RSHIFT              = uint32(0xA1)
	LCONTROL            = uint32(0xA2)
	RCONTROL            = uint32(0xA3)
	LMENU               = uint32(0xA4)
	RMENU               = uint32(0xA5)
	BROWSER_BACK        = uint32(0xA6)
	BROWSER_FORWARD     = uint32(0xA7)
	BROWSER_REFRESH     = uint32(0xA8)
	BROWSER_STOP        = uint32(0xA9)
	BROWSER_SEARCH      = uint32(0xAA)
	BROWSER_FAVORITES   = uint32(0xAB)
	BROWSER_HOME        = uint32(0xAC)
	VOLUME_MUTE         = uint32(0xAD)
	VOLUME_DOWN         = uint32(0xAE)
	VOLUME_UP           = uint32(0xAF)
	MEDIA_NEXT_TRACK    = uint32(0xB0)
	MEDIA_PREV_TRACK    = uint32(0xB1)
	MEDIA_STOP          = uint32(0xB2)
	MEDIA_PLAY_PAUSE    = uint32(0xB3)
	LAUNCH_MAIL         = uint32(0xB4)
	LAUNCH_MEDIA_SELECT = uint32(0xB5)
	LAUNCH_APP1         = uint32(0xB6)
	LAUNCH_APP2         = uint32(0xB7)
	OEM_1               = uint32(0xBA)
	OEM_PLUS            = uint32(0xBB)
	OEM_COMMA           = uint32(0xBC)
	OEM_MINUS           = uint32(0xBD)
	OEM_PERIOD          = uint32(0xBE)
	OEM_2               = uint32(0xBF)
	OEM_3               = uint32(0xC0)
	OEM_4               = uint32(0xDB)
	OEM_5               = uint32(0xDC)
	OEM_6               = uint32(0xDD)
	OEM_7               = uint32(0xDE)
	OEM_8               = uint32(0xDF)
	OEM_102             = uint32(0xE2)
	PROCESSKEY          = uint32(0xE5)
	PACKET              = uint32(0xE7)
	ATTN                = uint32(0xF6)
	CRSEL               = uint32(0xF7)
	EXSEL               = uint32(0xF8)
	EREOF               = uint32(0xF9)
	PLAY                = uint32(0xFA)
	ZOOM                = uint32(0xFB)
	NONAME              = uint32(0xFC)
	PA1                 = uint32(0xFD)
	OEM_CLEAR           = uint32(0xFE)
)