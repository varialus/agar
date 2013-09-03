// http://libagar.org/docs/inst/unix.html

package main

import (
	"fmt"
	"github.com/varialus/agar/development/core"
	"github.com/varialus/agar/development/gui"
	"os"
)
func main() {
	var win *gui.AG_Window
	fmt.Println("win ==", win)
	if core.AG_InitCore(nil, 0) == -1 || gui.AG_InitGraphics("") == -1 {
		os.Exit(1)
	}
	win = gui.AG_WindowNew(0)
	gui.AG_LabelNew(win, 0, "Hello, world!")
	gui.AG_WindowShow(win)
	gui.AG_EventLoop()
	os.Exit(0)
}
