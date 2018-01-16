package x11

import (
	"fmt"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func MoveWindow(xc *xgb.Conn, xw xproto.Window, x, y, width, height int32) (int32, int32, int32, int32) {
	vals := []uint32{}

	flags := xproto.ConfigWindowHeight |
		xproto.ConfigWindowWidth |
		xproto.ConfigWindowX |
		xproto.ConfigWindowY

	vals = append(vals, uint32(x))
	vals = append(vals, uint32(y))

	if int16(width) <= 0 {
		width = 1
	}
	vals = append(vals, uint32(width))

	if int16(height) <= 0 {
		height = 1
	}
	vals = append(vals, uint32(height))

	cook := xproto.ConfigureWindowChecked(xc, xw, uint16(flags), vals)
	if err := cook.Check(); err != nil {
		fmt.Println("X11 configure window failed: ", err)
	}
	return x, y, width, height
}