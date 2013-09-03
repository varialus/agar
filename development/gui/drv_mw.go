///*	Public domain	*/

package gui

//
///*
// * Multiple-window graphics driver framework. In this mode, Agar offers an
// * interface to an existing window manager, as opposed to providing one
// * internally (i.e., each Agar window corresponds to a "native" window).
// */
//
//struct ag_size_alloc;
//
//typedef struct ag_driver_mw_class {
type ag_driver_mw_class struct {
//	struct ag_driver_class _inherit;
//
//	/* Open/close native windows */
//	int  (*openWindow)(struct ag_window *, AG_Rect r, int bpp, Uint flags);
//	void (*closeWindow)(struct ag_window *);
//
//	/* Show and hide window */
//	int (*mapWindow)(struct ag_window *);
//	int (*unmapWindow)(struct ag_window *);
//
//	/* Configure stacking order and parenting */
//	int (*raiseWindow)(struct ag_window *);
//	int (*lowerWindow)(struct ag_window *);
//	int (*reparentWindow)(struct ag_window *, struct ag_window *, int, int);
//
//	/* Change and query input focus state */
//	int (*getInputFocus)(struct ag_window **);
//	int (*setInputFocus)(struct ag_window *);
//
//	/* Move and resize windows */
//	int  (*moveWindow)(struct ag_window *, int x, int y);
//	int  (*resizeWindow)(struct ag_window *, Uint w, Uint h);
//	int  (*moveResizeWindow)(struct ag_window *, struct ag_size_alloc *a);
//	void (*preResizeCallback)(struct ag_window *);
//	void (*postResizeCallback)(struct ag_window *, struct ag_size_alloc *a);
//
//	/* Capture window framebuffer contents (unlike renderToSurface) */
//	int (*captureWindow)(struct ag_window *, AG_Surface **s);
//	
//	/* Configure window parameters */
//	int  (*setBorderWidth)(struct ag_window *, Uint w);
//	int  (*setWindowCaption)(struct ag_window *, const char *);
//	void (*setTransientFor)(struct ag_window *, struct ag_window *);
//} AG_DriverMwClass;
}

type aG_DriverMwClass ag_driver_mw_class

//
//typedef struct ag_driver_mw {
//	struct ag_driver _inherit;
//	Uint flags;
//#define AG_DRIVER_MW_OPEN	0x01		/* Window is open */
//	struct ag_window *win;			/* Back pointer to window */
//} AG_DriverMw;
//
//#define AGDRIVER_MW(obj) ((AG_DriverMw *)(obj))
//#define AGDRIVER_MW_CLASS(obj) ((struct ag_driver_mw_class *)(AGOBJECT(obj)->cls))
//
///* Flags to openWindow */
//#define AG_DRIVER_MW_ANYPOS	0x01		/* No preferred position */
//
//__BEGIN_DECLS
//extern AG_ObjectClass    agDriverMwClass;
//__END_DECLS
//
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
// * Multiple-window graphics driver framework. Under this framework, each
// * Agar window has a corresponding "native" window managed by the driver.
// */
//
//#include <config/have_sdl.h>
//#include <config/have_opengl.h>
//#include <config/have_glx.h>
//
//#include <core/core.h>
//
//#include "window.h"
//
//static void
//Init(void *obj)
//{
//	AG_DriverMw *dmw = obj;
//
//	dmw->flags = 0;
//	dmw->win = NULL;
//}
//
//AG_ObjectClass agDriverMwClass = {
//	"AG_Driver:AG_DriverMw",
//	sizeof(AG_DriverMw),
//	{ 1,4 },
//	Init,
//	NULL,		/* reinit */
//	NULL,		/* destroy */
//	NULL,		/* load */
//	NULL,		/* save */
//	NULL		/* edit */
//};
