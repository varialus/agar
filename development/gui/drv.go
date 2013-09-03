///*	Public domain	*/

package gui

///*
// * Generic graphics/input driver framework.
// */
//
//#ifndef _AGAR_GUI_DRV_H_
//#define _AGAR_GUI_DRV_H_
//#include <agar/gui/begin.h>

//import (
//	"unsafe"
//)

//
//enum ag_driver_type {
//	AG_FRAMEBUFFER,			/* Direct rendering to frame buffer */
//	AG_VECTOR			/* Vector drawing */
//};
//enum ag_driver_wm_type {
//	AG_WM_SINGLE,			/* Single display */
//	AG_WM_MULTIPLE			/* Multiple windows */
//};
//
//struct ag_widget;
//struct ag_window;
//struct ag_glyph;
//struct ag_glyph_cache;
//struct ag_cursor;
//struct ag_driver_event;
//
///* Generic graphics driver class */
//typedef struct ag_driver_class {
type ag_driver_class struct {
//	struct ag_object_class _inherit;
//	const char *name;			/* Short name */
//	enum ag_driver_type type;		/* Driver type */
//	enum ag_driver_wm_type wm;		/* Window manager type */
//	Uint flags;
//#define AG_DRIVER_OPENGL	0x01		/* Supports OpenGL calls */
//#define AG_DRIVER_SDL		0x02		/* Supports SDL calls */
//#define AG_DRIVER_TEXTURES	0x04		/* Support texture ops */
//
//	/* Initialization */
//	int  (*open)(void *drv, const char *spec);
//	void (*close)(void *drv);
//	int  (*getDisplaySize)(Uint *w, Uint *h);
//	/* Event processing */
//	void (*beginEventProcessing)(void *drv);
//	int  (*pendingEvents)(void *drv);
//	int  (*getNextEvent)(void *drv, struct ag_driver_event *dev);
//	int  (*processEvent)(void *drv, struct ag_driver_event *dev);
//	void (*genericEventLoop)(void *drv);
//	void (*endEventProcessing)(void *drv);
//	void (*terminate)(void);
//	/* GUI rendering */
//	void (*beginRendering)(void *drv);
//	void (*renderWindow)(struct ag_window *);
//	void (*endRendering)(void *drv);
//	/* Primitives */
//	void (*fillRect)(void *drv, AG_Rect r, AG_Color c);
//	/* Update video region (rendering context; FB driver specific) */
//	void (*updateRegion)(void *drv, AG_Rect r);
//	/* Texture operations (GL driver specific) */
//	void (*uploadTexture)(Uint *, AG_Surface *, AG_TexCoord *);
//	int  (*updateTexture)(Uint, AG_Surface *, AG_TexCoord *);
//	void (*deleteTexture)(void *drv, Uint);
//	/* Request a specific refresh rate (driver specific) */
//	int (*setRefreshRate)(void *drv, int fps);
//	/* Clipping and blending control (rendering context) */
//	void (*pushClipRect)(void *drv, AG_Rect r);
//	void (*popClipRect)(void *drv);
//	void (*pushBlendingMode)(void *drv, AG_BlendFn srcFn, AG_BlendFn dstFn);
//	void (*popBlendingMode)(void *drv);
//	/* Hardware cursor operations */
//	int  (*createCursor)(void *drv, struct ag_cursor *curs);
//	void (*freeCursor)(void *drv, struct ag_cursor *curs);
//	int  (*setCursor)(void *drv, struct ag_cursor *curs);
//	void (*unsetCursor)(void *drv);
//	int  (*getCursorVisibility)(void *drv);
//	void (*setCursorVisibility)(void *drv, int flag);
//	/* Widget surface operations (rendering context) */
//	void (*blitSurface)(void *drv, struct ag_widget *wid, AG_Surface *s, int x, int y);
//	void (*blitSurfaceFrom)(void *drv, struct ag_widget *wid, struct ag_widget *widSrc, int s, AG_Rect *r, int x, int y);
//	void (*blitSurfaceGL)(void *drv, struct ag_widget *wid, AG_Surface *s, float w, float h);
//	void (*blitSurfaceFromGL)(void *drv, struct ag_widget *wid, int s, float w, float h);
//	void (*blitSurfaceFlippedGL)(void *drv, struct ag_widget *wid, int s, float w, float h);
//	void (*backupSurfaces)(void *drv, struct ag_widget *wid);
//	void (*restoreSurfaces)(void *drv, struct ag_widget *wid);
//	int  (*renderToSurface)(void *drv, struct ag_widget *wid, AG_Surface **su);
//	/* Rendering operations (rendering context) */
//	void (*putPixel)(void *drv, int x, int y, AG_Color c);
//	void (*putPixel32)(void *drv, int x, int y, Uint32 c);
//	void (*putPixelRGB)(void *drv, int x, int y, Uint8 r, Uint8 g, Uint8 b);
//	void (*blendPixel)(void *drv, int x, int y, AG_Color c, AG_BlendFn fnSrc, AG_BlendFn fnDst);
//	void (*drawLine)(void *drv, int x1, int y1, int x2, int y2, AG_Color C);
//	void (*drawLineH)(void *drv, int x1, int x2, int y, AG_Color C);
//	void (*drawLineV)(void *drv, int x, int y1, int y2, AG_Color C);
//	void (*drawLineBlended)(void *drv, int x1, int y1, int x2, int y2, AG_Color C, AG_BlendFn fnSrc, AG_BlendFn fnDst);
//	void (*drawArrowUp)(void *drv, int x0, int y0, int h, AG_Color C[2]);
//	void (*drawArrowDown)(void *drv, int x0, int y0, int h, AG_Color C[2]);
//	void (*drawArrowLeft)(void *drv, int x0, int y0, int h, AG_Color C[2]);
//	void (*drawArrowRight)(void *drv, int x0, int y0, int h, AG_Color C[2]);
//	void (*drawBoxRounded)(void *drv, AG_Rect r, int z, int rad, AG_Color C[3]);
//	void (*drawBoxRoundedTop)(void *drv, AG_Rect r, int z, int rad, AG_Color C[3]);
//	void (*drawCircle)(void *drv, int x, int y, int r, AG_Color C);
//	void (*drawCircle2)(void *drv, int x, int y, int r, AG_Color C);
//	void (*drawRectFilled)(void *drv, AG_Rect r, AG_Color C);
//	void (*drawRectBlended)(void *drv, AG_Rect r, AG_Color C, AG_BlendFn fnSrc, AG_BlendFn fnDst);
//	void (*drawRectDithered)(void *drv, AG_Rect r, AG_Color C);
//	void (*updateGlyph)(void *drv, struct ag_glyph *);
//	void (*drawGlyph)(void *drv, const struct ag_glyph *, int x, int y);
//	/* Display list management (GL driver specific) */
//	void (*deleteList)(void *drv, Uint);
//} AG_DriverClass;
}

type aG_DriverClass ag_driver_class

func (d aG_DriverClass) genericEventLoop(drv *aG_DriverSw) {

}

//
///* Generic driver instance. */
//typedef struct ag_driver {
type ag_driver struct {
//	struct ag_object _inherit;
//	Uint id;			/* Numerical instance ID */
//	Uint flags;
	flags uint
//#define AG_DRIVER_FIXED_FPS	0x01	/* Invoked AG_EventLoop_FixedFPS() */
//
//	AG_Surface *sRef;		/* "Reference" surface */
//	AG_PixelFormat *videoFmt;	/* Video pixel format (for
//					   packed-pixel FB modes) */
//	struct ag_keyboard *kbd;	/* Main keyboard device */
//	struct ag_mouse *mouse;		/* Main mouse device */
//	struct ag_cursor *activeCursor;	/* Effective cursor */
//	struct ag_cursor *cursors;	/* Registered mouse cursors */
//	Uint             nCursors;
//	struct ag_glyph_cache *glyphCache; /* Cache of rendered glyphs */
//	Uint8 glStipple[128];		/* For GL drivers */
//} AG_Driver;
}

type aG_Driver ag_driver

const (
	aG_DRIVER_FIXED_FPS = 0x01
)
//
///* Generic driver event (for custom event loops). */
//enum ag_driver_event_type {
//	AG_DRIVER_UNKNOWN,		/* Unknown event */
//	AG_DRIVER_MOUSE_MOTION,		/* Cursor moved */
//	AG_DRIVER_MOUSE_BUTTON_DOWN,	/* Mouse button pressed */
//	AG_DRIVER_MOUSE_BUTTON_UP,	/* Mouse button released */
//	AG_DRIVER_MOUSE_ENTER,		/* Mouse entering window (MW) */
//	AG_DRIVER_MOUSE_LEAVE,		/* Mouse leaving window (MW) */
//	AG_DRIVER_FOCUS_IN,		/* Focus on window (MW) */
//	AG_DRIVER_FOCUS_OUT,		/* Focus out of window (MW) */
//	AG_DRIVER_KEY_DOWN,		/* Key pressed */
//	AG_DRIVER_KEY_UP,		/* Key released */
//	AG_DRIVER_EXPOSE,		/* Video update needed */
//	AG_DRIVER_VIDEORESIZE,		/* Video resize request */
//	AG_DRIVER_CLOSE			/* Window close request */
//};
//typedef struct ag_driver_event {
//	enum ag_driver_event_type type;	/* Type of event */
//	struct ag_window *win;		/* Associated window (AG_WM_MULTIPLE) */
//	union {
//		struct {
//			int x, y;
//		} motion;
//		struct {
//			AG_MouseButton which;
//			int x, y;
//		} button;
//		struct {
//			AG_KeySym ks;
//			Uint32 ucs;
//		} key;
//		struct {
//			int x, y, w, h;
//		} videoresize;
//	} data;
//	AG_TAILQ_ENTRY(ag_driver_event) events;
//} AG_DriverEvent;
//
//typedef AG_TAILQ_HEAD(ag_driver_eventq, ag_driver_event) AG_DriverEventQ;
//
//#define AGDRIVER(obj)		((AG_Driver *)(obj))
//#define AGDRIVER_CLASS(obj)	((struct ag_driver_class *)(AGOBJECT(obj)->cls))
//#define AGDRIVER_SINGLE(drv)	(AGDRIVER_CLASS(drv)->wm == AG_WM_SINGLE)
//#define AGDRIVER_MULTIPLE(drv)	(AGDRIVER_CLASS(drv)->wm == AG_WM_MULTIPLE)
//#define AGDRIVER_BOUNDED_WIDTH(win,x) (((x) < 0) ? 0 : \
//                                      ((x) > AGWIDGET(win)->w) ? (AGWIDGET(win)->w - 1) : (x))
//#define AGDRIVER_BOUNDED_HEIGHT(win,y) (((y) < 0) ? 0 : \
//                                       ((y) > AGWIDGET(win)->h) ? (AGWIDGET(win)->h - 1) : (y))
//__BEGIN_DECLS
//extern AG_ObjectClass agDriverClass;
//
//extern AG_Object       agDrivers;	/* Drivers VFS */
//extern AG_DriverClass *agDriverOps;	/* Current driver class */
var agDriverOps *aG_DriverClass
//extern void           *agDriverList[];	/* Available drivers (AG_DriverClass) */
//extern Uint            agDriverListSize;
/* Already Defined Below */
//extern int             agRenderingContext;
//
//void       AG_ListDriverNames(char *, size_t)
//                              BOUNDED_ATTRIBUTE(__string__, 1, 2);
//AG_Driver *AG_DriverOpen(AG_DriverClass *);
//void       AG_DriverClose(AG_Driver *);
//void       AG_ViewCapture(void);
//
///* Lookup a driver instance by ID */
//static __inline__ AG_Driver *
//AG_GetDriverByID(Uint id)
//{
//	AG_Driver *drv;
//
//	AGOBJECT_FOREACH_CHILD(drv, &agDrivers, ag_driver) {
//		if (drv->id == id)
//			return (drv);
//	}
//	return (NULL);
//}
//
///* Enter GUI rendering context. */
//static __inline__ void
//AG_BeginRendering(void *drv)
//{
//	agRenderingContext = 1;
//	AGDRIVER_CLASS(drv)->beginRendering(drv);
//}
//
///* Leave GUI rendering context. */
//static __inline__ void
//AG_EndRendering(void *drv)
//{
//	AGDRIVER_CLASS(drv)->endRendering(drv);
//	agRenderingContext = 0;
//}
//
///* Create a texture from a surface (GL drivers). */
//static __inline__ Uint
//AG_SurfaceTexture(AG_Surface *su, AG_TexCoord *tc)
//{
//	Uint texid;
//	agDriverOps->uploadTexture(&texid, su, tc);
//	return (texid);
//}
//
///* Update texture contents from a surface (GL drivers). */
//static __inline__ void
//AG_UpdateTexture(AG_Surface *su, int texid, AG_TexCoord *tc)
//{
//	if (agDriverOps->updateTexture((Uint)texid, su, tc) == -1)
//		AG_FatalError(NULL);
//}
//
//#ifdef AG_LEGACY
//extern AG_Driver *agView;  	/* Pre-1.4 */
//#endif /* AG_LEGACY */
//__END_DECLS
//
//#include <agar/gui/drv_mw.h>
//#include <agar/gui/drv_sw.h>
//
//__BEGIN_DECLS
///* Return whether Agar is using OpenGL. */
//static __inline__ int
//AG_UsingGL(void *drv)
//{
//	if (drv != NULL) {
//		return (AGDRIVER_CLASS(drv)->flags & AG_DRIVER_OPENGL);
//	} else {
//		return (agDriverOps->flags & AG_DRIVER_OPENGL);
//	}
//}
//
///* Return whether Agar is using SDL. */
//static __inline__ int
//AG_UsingSDL(void *drv)
//{
//	AG_DriverClass *dc = (drv != NULL) ? AGDRIVER_CLASS(drv) : agDriverOps;
//	return (dc->flags & AG_DRIVER_SDL);
//}
//
///* Query a driver for available display area in pixels. */
//static __inline__ int
//AG_GetDisplaySize(void *drv, Uint *w, Uint *h)
//{
//	AG_DriverClass *dc = (drv != NULL) ? AGDRIVER_CLASS(drv) : agDriverOps;
//
//	switch (dc->wm) {
//	case AG_WM_SINGLE:
//		*w = AGDRIVER_SW(drv)->w;
//		*h = AGDRIVER_SW(drv)->h;
//		return (0);
//	case AG_WM_MULTIPLE:
//		return dc->getDisplaySize(w, h);
//	}
//	return (-1);
//}
//__END_DECLS
//
//#include <agar/gui/close.h>
//#endif /* _AGAR_GUI_DRV_H_ */

///*
// * Copyright (c) 2009 Hypertriton, Inc. <http://hypertriton.com/>
// * All rights reserved.
// *
// * Redistribution and use in source and binary forms, with or without
// * modification, are permitted provided that the following conditions
// * are met:
// * 1. Redistributions of source code must retain the above copyright
// *    notice, this list of conditions and the following disclaimer.
// * 2. Redistributions in binary form must reproduce the above copyright
// *    notice, this list of conditions and the following disclaimer in the
// *    documentation and/or other materials provided with the distribution.
// * 
// * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// * ARE DISCLAIMED. IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE FOR
// * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// * DAMAGES (INCLUDING BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE
// * USE OF THIS SOFTWARE EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
// */
//
///*
// * Generic graphics/input driver framework.
// */
//
//#include <config/have_sdl.h>
//#include <config/have_opengl.h>
//#include <config/have_glx.h>
//#include <config/have_wgl.h>
//
//#include <core/core.h>
//#include <core/config.h>
//
//#include "window.h"
//#include "text.h"
//
//#if defined(HAVE_GLX)
//extern AG_Driver agDriverGLX;
var AgDriverGLX aG_Driver
//#endif
//#if defined(HAVE_SDL)
//extern AG_Driver agDriverSDLFB;
//#endif
//#if defined(HAVE_SDL) && defined(HAVE_OPENGL)
//extern AG_Driver agDriverSDLGL;
//#endif
//#if defined(HAVE_WGL)
//extern AG_Driver agDriverWGL;
//#endif
//
//AG_Object         agDrivers;			/* Drivers VFS */
//AG_DriverClass   *agDriverOps = NULL;		/* Current driver class */
//AG_DriverSw      *agDriverSw = NULL;		/* Driver instance (or NULL) */
//#ifdef AG_LEGACY
//AG_Driver        *agView = NULL;  		/* Pre-1.4 */
//#endif
//
//void *agDriverList[] = {
var agDriverList []*aG_Driver = []*aG_Driver{&AgDriverGLX}
//#if defined(HAVE_GLX)
//	&agDriverGLX,
//#endif
//#if defined(HAVE_WGL)
//	&agDriverWGL,
//#endif
//#if defined(HAVE_SDL)
//	&agDriverSDLFB,
//#endif
//#if defined(HAVE_SDL) && defined(HAVE_OPENGL)
//	&agDriverSDLGL,
//#endif
//};
//Uint agDriverListSize = sizeof(agDriverList) / sizeof(agDriverList[0]);
var agDriverListSize uint = uint(len(agDriverList))
//
///* Return a string with the available drivers. */
//void
//AG_ListDriverNames(char *buf, size_t buf_len)
//{
//	Uint i;
//
//	if (buf_len == 0) {
//		return;
//	}
//	buf[0] = '\0';
//
//	for (i = 0; i < agDriverListSize; i++) {
//		AG_DriverClass *drvClass = agDriverList[i];
//
//		Strlcat(buf, drvClass->name, buf_len);
//		if (i < agDriverListSize-1)
//			Strlcat(buf, " ", buf_len);
//	}
//}
//
///* Create an instance of the given driver class, if it opens successfully. */
//AG_Driver *
//AG_DriverOpen(AG_DriverClass *dc)
//{
//	AG_Driver *drv;
func aG_DriverOpen(dc *aG_DriverClass) (drv *aG_Driver) {
//
//	if ((drv = AG_ObjectNew(NULL, dc->name, AGCLASS(dc))) == NULL) {
//		return (NULL);
//	}
//	if (dc->open(drv, NULL) == -1) {
//		AG_ObjectDestroy(drv);
//		return (NULL);
//	}
//	for (drv->id = 1; ; drv->id++) {
//		if (AG_GetDriverByID(drv->id) == NULL)
//			break;
//	}
//	AG_ObjectSetName(drv, "%s%u", dc->name, drv->id);
//	AG_ObjectAttach(&agDrivers, drv);
//	return (drv);
	return drv
//}
}
//
///* Close and destroy a driver. */
//void
//AG_DriverClose(AG_Driver *drv)
//{
//	AG_ObjectDetach(drv);
//	AGDRIVER_CLASS(drv)->close(drv);
//	AG_ObjectDestroy(drv);
//}
//
///*
// * Dump the display surface(s) to a jpeg in ~/.appname/screenshot/.
// * Typically called via AG_GlobalKeys(3).
// * This only works under single-display drivers (use AG_WidgetSurface() to
// * capture windows instead).
// */
//void
//AG_ViewCapture(void)
//{
//	AG_Surface *s;
//	char dir[AG_PATHNAME_MAX];
//	char file[AG_PATHNAME_MAX];
//	Uint seq;
//
//	if (agDriverSw == NULL) {
//		Verbose("AG_ViewCapture() is not implemented under "
//		        "multiple-window drivers\n");
//		return;
//	}
//	if (AGDRIVER_SW_CLASS(agDriverSw)->videoCapture(agDriverSw, &s) == -1) {
//		Verbose("Capture failed: %s\n", AG_GetError());
//		return;
//	}
//
//	/* Save to a new file. */
//	AG_GetString(agConfig, "save-path", dir, sizeof(dir));
//	Strlcat(dir, AG_PATHSEP, sizeof(dir));
//	Strlcat(dir, "screenshot", sizeof(dir));
//	if (!AG_FileExists(dir) && AG_MkPath(dir) == -1) {
//		Verbose("Capture failed: %s\n", AG_GetError());
//		return;
//	}
//	for (seq = 0; ; seq++) {
//		Snprintf(file, sizeof(file), "%s%c%s%u.jpg",
//		    dir, AG_PATHSEPCHAR, agProgName, seq++);
//		if (!AG_FileExists(file))
//			break;			/* XXX race condition */
//	}
//	if (AG_SurfaceExportJPEG(s, file) == 0) {
//		Verbose("Saved capture to: %s\n", file);
//	} else {
//		Verbose("Capture failed: %s\n", AG_GetError());
//	}
//	AG_SurfaceFree(s);
//}
//
//static void
//Init(void *obj)
//{
//	AG_Driver *drv = obj;
//
//	drv->id = 0;
//	drv->flags = 0;
//	drv->sRef = AG_SurfaceRGBA(1,1, 32, 0,
//#if AG_BYTEORDER == AG_BIG_ENDIAN
// 	    0xff000000, 0x00ff0000, 0x0000ff00, 0x000000ff
//#else
//	    0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000
//#endif
//	);
//	if (drv->sRef == NULL) {
//		AG_FatalError(NULL);
//	}
//	drv->videoFmt = NULL;
//	drv->kbd = NULL;
//	drv->mouse = NULL;
//	drv->cursors = NULL;
//	drv->nCursors = 0;
//	drv->activeCursor = NULL;
//	AG_TextInitGlyphCache(drv);
//	memset(drv->glStipple, 0xaa, sizeof(drv->glStipple));
//}
//
//static void
//Destroy(void *obj)
//{
//	AG_Driver *drv = obj;
//
//	if (drv->sRef != NULL)
//		AG_SurfaceFree(drv->sRef);
//	if (drv->videoFmt != NULL)
//		AG_PixelFormatFree(drv->videoFmt);
//
//	AG_TextClearGlyphCache(drv);
//}
//
//AG_ObjectClass agDriverClass = {
//	"AG_Driver",
//	sizeof(AG_Driver),
//	{ 1,4 },
//	Init,
//	NULL,		/* reinit */
//	Destroy,
//	NULL,		/* load */
//	NULL,		/* save */
//	NULL		/* edit */
//};///*
// * Copyright (c) 2009 Hypertriton, Inc. <http://hypertriton.com/>
// * All rights reserved.
// *
// * Redistribution and use in source and binary forms, with or without
// * modification, are permitted provided that the following conditions
// * are met:
// * 1. Redistributions of source code must retain the above copyright
// *    notice, this list of conditions and the following disclaimer.
// * 2. Redistributions in binary form must reproduce the above copyright
// *    notice, this list of conditions and the following disclaimer in the
// *    documentation and/or other materials provided with the distribution.
// * 
// * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// * ARE DISCLAIMED. IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE FOR
// * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// * DAMAGES (INCLUDING BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE
// * USE OF THIS SOFTWARE EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
// */
//
///*
// * Generic graphics/input driver framework.
// */
//
//#include <config/have_sdl.h>
//#include <config/have_opengl.h>
//#include <config/have_glx.h>
//#include <config/have_wgl.h>
//
//#include <core/core.h>
//#include <core/config.h>
//
//#include "window.h"
//#include "text.h"
//
//#if defined(HAVE_GLX)
//extern AG_Driver agDriverGLX;
//#endif
//#if defined(HAVE_SDL)
//extern AG_Driver agDriverSDLFB;
//#endif
//#if defined(HAVE_SDL) && defined(HAVE_OPENGL)
//extern AG_Driver agDriverSDLGL;
//#endif
//#if defined(HAVE_WGL)
//extern AG_Driver agDriverWGL;
//#endif
//
//AG_Object         agDrivers;			/* Drivers VFS */
//AG_DriverClass   *agDriverOps = NULL;		/* Current driver class */
//AG_DriverSw      *agDriverSw = NULL;		/* Driver instance (or NULL) */
//#ifdef AG_LEGACY
//AG_Driver        *agView = NULL;  		/* Pre-1.4 */
//#endif
//
//void *agDriverList[] = {
//#if defined(HAVE_GLX)
//	&agDriverGLX,
//#endif
//#if defined(HAVE_WGL)
//	&agDriverWGL,
//#endif
//#if defined(HAVE_SDL)
//	&agDriverSDLFB,
//#endif
//#if defined(HAVE_SDL) && defined(HAVE_OPENGL)
//	&agDriverSDLGL,
//#endif
//};
//Uint agDriverListSize = sizeof(agDriverList) / sizeof(agDriverList[0]);
//
///* Return a string with the available drivers. */
//void
//AG_ListDriverNames(char *buf, size_t buf_len)
//{
//	Uint i;
//
//	if (buf_len == 0) {
//		return;
//	}
//	buf[0] = '\0';
//
//	for (i = 0; i < agDriverListSize; i++) {
//		AG_DriverClass *drvClass = agDriverList[i];
//
//		Strlcat(buf, drvClass->name, buf_len);
//		if (i < agDriverListSize-1)
//			Strlcat(buf, " ", buf_len);
//	}
//}
//
///* Create an instance of the given driver class, if it opens successfully. */
//AG_Driver *
//AG_DriverOpen(AG_DriverClass *dc)
//{
//	AG_Driver *drv;
//
//	if ((drv = AG_ObjectNew(NULL, dc->name, AGCLASS(dc))) == NULL) {
//		return (NULL);
//	}
//	if (dc->open(drv, NULL) == -1) {
//		AG_ObjectDestroy(drv);
//		return (NULL);
//	}
//	for (drv->id = 1; ; drv->id++) {
//		if (AG_GetDriverByID(drv->id) == NULL)
//			break;
//	}
//	AG_ObjectSetName(drv, "%s%u", dc->name, drv->id);
//	AG_ObjectAttach(&agDrivers, drv);
//	return (drv);
//}
//
///* Close and destroy a driver. */
//void
//AG_DriverClose(AG_Driver *drv)
//{
//	AG_ObjectDetach(drv);
//	AGDRIVER_CLASS(drv)->close(drv);
//	AG_ObjectDestroy(drv);
//}
//
///*
// * Dump the display surface(s) to a jpeg in ~/.appname/screenshot/.
// * Typically called via AG_GlobalKeys(3).
// * This only works under single-display drivers (use AG_WidgetSurface() to
// * capture windows instead).
// */
//void
//AG_ViewCapture(void)
//{
//	AG_Surface *s;
//	char dir[AG_PATHNAME_MAX];
//	char file[AG_PATHNAME_MAX];
//	Uint seq;
//
//	if (agDriverSw == NULL) {
//		Verbose("AG_ViewCapture() is not implemented under "
//		        "multiple-window drivers\n");
//		return;
//	}
//	if (AGDRIVER_SW_CLASS(agDriverSw)->videoCapture(agDriverSw, &s) == -1) {
//		Verbose("Capture failed: %s\n", AG_GetError());
//		return;
//	}
//
//	/* Save to a new file. */
//	AG_GetString(agConfig, "save-path", dir, sizeof(dir));
//	Strlcat(dir, AG_PATHSEP, sizeof(dir));
//	Strlcat(dir, "screenshot", sizeof(dir));
//	if (!AG_FileExists(dir) && AG_MkPath(dir) == -1) {
//		Verbose("Capture failed: %s\n", AG_GetError());
//		return;
//	}
//	for (seq = 0; ; seq++) {
//		Snprintf(file, sizeof(file), "%s%c%s%u.jpg",
//		    dir, AG_PATHSEPCHAR, agProgName, seq++);
//		if (!AG_FileExists(file))
//			break;			/* XXX race condition */
//	}
//	if (AG_SurfaceExportJPEG(s, file) == 0) {
//		Verbose("Saved capture to: %s\n", file);
//	} else {
//		Verbose("Capture failed: %s\n", AG_GetError());
//	}
//	AG_SurfaceFree(s);
//}
//
//static void
//Init(void *obj)
//{
//	AG_Driver *drv = obj;
//
//	drv->id = 0;
//	drv->flags = 0;
//	drv->sRef = AG_SurfaceRGBA(1,1, 32, 0,
//#if AG_BYTEORDER == AG_BIG_ENDIAN
// 	    0xff000000, 0x00ff0000, 0x0000ff00, 0x000000ff
//#else
//	    0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000
//#endif
//	);
//	if (drv->sRef == NULL) {
//		AG_FatalError(NULL);
//	}
//	drv->videoFmt = NULL;
//	drv->kbd = NULL;
//	drv->mouse = NULL;
//	drv->cursors = NULL;
//	drv->nCursors = 0;
//	drv->activeCursor = NULL;
//	AG_TextInitGlyphCache(drv);
//	memset(drv->glStipple, 0xaa, sizeof(drv->glStipple));
//}
//
//static void
//Destroy(void *obj)
//{
//	AG_Driver *drv = obj;
//
//	if (drv->sRef != NULL)
//		AG_SurfaceFree(drv->sRef);
//	if (drv->videoFmt != NULL)
//		AG_PixelFormatFree(drv->videoFmt);
//
//	AG_TextClearGlyphCache(drv);
//}
//
//AG_ObjectClass agDriverClass = {
//	"AG_Driver",
//	sizeof(AG_Driver),
//	{ 1,4 },
//	Init,
//	NULL,		/* reinit */
//	Destroy,
//	NULL,		/* load */
//	NULL,		/* save */
//	NULL		/* edit */
//};
