// Generated automatically.  DO NOT HAND-EDIT.

package xfce

import "github.com/go-curses/terminfo"

func init() {

	// Xfce Terminal
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "xfce",
		Columns:      80,
		Lines:        24,
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[2J",
		EnterCA:      "\x1b7\x1b[?47h",
		ExitCA:       "\x1b[2J\x1b[?47l\x1b8",
		ShowCursor:   "\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[0m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b[?1h\x1b=",
		ExitKeypad:   "\x1b[?1l\x1b>",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		ResetFgBg:    "\x1b[39;49m",
		PadChar:      "\x00",
		AltChars:     "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b)0",
		Mouse:        "\x1b[M",
		MouseMode:    "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1bOA",
		KeyDown:      "\x1bOB",
		KeyRight:     "\x1bOC",
		KeyLeft:      "\x1bOD",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\u007f",
		KeyHome:      "\x1bOH",
		KeyEnd:       "\x1bOF",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1bOP",
		KeyF2:        "\x1bOQ",
		KeyF3:        "\x1bOR",
		KeyF4:        "\x1bOS",
		KeyF5:        "\x1b[15~",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyBacktab:   "\x1b[Z",
		Modifiers:    1,
	})
}
