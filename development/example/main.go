// http://libagar.org/docs/inst/unix.html

package main

//#include <agar/core.h>
//#include <agar/gui.h>
import (
	"fmt"
	"github.com/varialus/agar/development/core"
	"github.com/varialus/agar/development/gui"
	"os"
)
//
//int
//main(int argc, char *argv)
//{
func main() {
//	AG_Window *win;
	var win *gui.AG_Window
	fmt.Println("win ==", win)
//
//	if (AG_InitCore(NULL, 0) == -1 ||
//		AG_InitGraphics(0) == -1) {
	if core.AG_InitCore(nil, 0) == -1 || gui.AG_InitGraphics(0) == -1 {
//		return (1);
		os.Exit(1)
//	}
	}
//	win = AG_WindowNew(0);
	win = gui.AG_WindowNew(0)
//	AG_LabelNew(win, 0, "Hello, world!");
	gui.AG_LabelNew(win, 0, "Hello, world!")
//	AG_WindowShow(win);
	gui.AG_WindowShow(win)
//	AG_EventLoop();
	gui.AG_EventLoop()
//	return (0);
	os.Exit(0)
//}
}
