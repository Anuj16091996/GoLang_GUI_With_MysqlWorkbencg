package View

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

const (
	SIZE_W = 600
	SIZE_H = 400
)

type MyMainWindow struct {
	*walk.MainWindow
}

func StartGui() {

	mw := new(MyMainWindow)

	MainWindow{
		Visible:  false,
		AssignTo: &mw.MainWindow,
	}.Create()

	defaultStyle := win.GetWindowLong(mw.Handle(), win.GWL_STYLE) // Gets current style
	newStyle := defaultStyle &^ win.WS_THICKFRAME                 // Remove WS_THICKFRAME
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, newStyle)

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)

	mw.Run()

}
