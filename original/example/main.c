// http://libagar.org/docs/inst/unix.html

#include <agar/core.h>
#include <agar/gui.h>

int
main(int argc, char *argv)
{
	AG_Window *win;

	if (AG_InitCore(NULL, 0) == -1 ||
		AG_InitGraphics(0) == -1) {
		return (1);
	}
	win = AG_WindowNew(0);
	AG_LabelNew(win, 0, "Hello, world!");
	AG_WindowShow(win);
	AG_EventLoop();
	return (0);
}
