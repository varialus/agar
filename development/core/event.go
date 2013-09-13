///*	Public domain	*/

package core

//
//#include <agar/core/begin.h>
//
//#define AG_EVENT_ARGS_MAX 16
//#define AG_EVENT_NAME_MAX 32
//
//#ifdef AG_DEBUG
//# define AG_PTR(v)	(event->argv[v].type==AG_VARIABLE_POINTER ? event->argv[v].data.p : AG_PtrMismatch())
//# define AG_OBJECT(v,t)	(AG_OfClass(event->argv[v].data.p,(t))) ? event->argv[v].data.p : AG_ObjectMismatch(AGOBJECT(event->argv[v].data.p)->cls->hier,(t))
//# define AG_STRING(v)	(event->argv[v].type==AG_VARIABLE_STRING ? event->argv[v].data.s : (char *)AG_PtrMismatch())
//# define AG_INT(v)	(event->argv[v].type==AG_VARIABLE_INT ? event->argv[v].data.i : AG_IntMismatch())
//# define AG_UINT(v)	(event->argv[v].type==AG_VARIABLE_UINT ? event->argv[v].data.u : (Uint)AG_IntMismatch())
//# define AG_LONG(v)	(event->argv[v].type==AG_VARIABLE_SINT32 ? event->argv[v].data.s32 : (Sint32)AG_IntMismatch())
//# define AG_ULONG(v)	(event->argv[v].type==AG_VARIABLE_UINT32 ? event->argv[v].data.u32 : (Uint32)AG_IntMismatch())
//# define AG_FLOAT(v)	(event->argv[v].type==AG_VARIABLE_FLOAT ? event->argv[v].data.flt : AG_FloatMismatch())
//# define AG_DOUBLE(v)	(event->argv[v].type==AG_VARIABLE_DOUBLE ? event->argv[v].data.dbl : (double)AG_FloatMismatch())
//#else /* !AG_DEBUG */
//# define AG_PTR(v)	(event->argv[v].data.p)
//# define AG_OBJECT(v,t)	(event->argv[v].data.p)
//# define AG_STRING(v)	(event->argv[v].data.s)
//# define AG_INT(v)	(event->argv[v].data.i)
//# define AG_UINT(v)	(event->argv[v].data.u)
//# define AG_LONG(v)	((long)event->argv[v].data.s32)
//# define AG_ULONG(v)	((Ulong)event->argv[v].data.u32)
//# define AG_FLOAT(v)	((float)event->argv[v].data.flt)
//# define AG_DOUBLE(v)	(event->argv[v].data.dbl)
//#endif /* AG_DEBUG */
//
//#define AG_SELF()	AG_PTR(0)
//#define AG_SENDER()	AG_PTR(event->argc)
//
//#define AG_PTR_NAMED(k)		AG_GetNamedPtr(event,(k))
//#define AG_OBJECT_NAMED(k,cls)	AG_GetNamedObject(event,(k),(cls))
//#define AG_STRING_NAMED(k)	AG_GetNamedString(event,(k))
//#define AG_INT_NAMED(k)		AG_GetNamedInt(event,(k))
//#define AG_UINT_NAMED(k)	AG_GetNamedUint(event,(k))
//#define AG_LONG_NAMED(k)	AG_GetNamedLong(event,(k))
//#define AG_ULONG_NAMED(k)	AG_GetNamedUlong(event,(k))
//#define AG_FLOAT_NAMED(k)	AG_GetNamedFlt(event,(k))
//#define AG_DOUBLE_NAMED(k)	AG_GetNamedDbl(event,(k))
//
//struct ag_timer;
//struct ag_event_sink;
//
///* Event structure */
//typedef struct ag_event {
//	char name[AG_EVENT_NAME_MAX];		/* String identifier */
//	Uint flags;
//#define	AG_EVENT_ASYNC     0x01			/* Service in separate thread */
//#define AG_EVENT_PROPAGATE 0x02			/* Forward to child objs */
//	void (*handler)(struct ag_event *);
//	int argc;				/* Argument count */
//	int argc0;				/* Arg. count (by SetEvent) */
//	AG_Variable argv[AG_EVENT_ARGS_MAX];	/* Argument values */
//	AG_TAILQ_ENTRY(ag_event) events;	/* Entry in Object */
//} AG_Event;
//
///* Low-level event sink */
//enum ag_event_sink_type {
//	AG_SINK_NONE,
//	AG_SINK_PROLOGUE,		/* Special event loop prologue */
//	AG_SINK_EPILOGUE,		/* Special event sink epilogue */
//	AG_SINK_SPINNER,		/* Special non-blocking sink */
//	AG_SINK_TERMINATOR,		/* Quit request */
//	AG_SINK_TIMER,			/* Timer expiration */
//	AG_SINK_READ,			/* Data available on fd */
//	AG_SINK_WRITE,			/* Write buffer available on fd */
//	AG_SINK_FSEVENT,		/* Filesystem event */
//	AG_SINK_PROCEVENT,		/* Process event */
//	AG_SINK_LAST
//};
//
//typedef int (*AG_EventSinkFn)(struct ag_event_sink *, AG_Event *);
//
//typedef struct ag_event_sink {
//	enum ag_event_sink_type type;		/* Event filter type */
//	int ident;				/* Identifier / fd */
//	Uint flags, flagsMatched;
//#define AG_FSEVENT_DELETE	0x0001		/* Referenced file deleted */
//#define AG_FSEVENT_WRITE	0x0002		/* Write occured */
//#define AG_FSEVENT_EXTEND	0x0004		/* File extended */
//#define AG_FSEVENT_ATTRIB	0x0008		/* File attributes changed */
//#define AG_FSEVENT_LINK		0x0010		/* Link count changed */
//#define AG_FSEVENT_RENAME	0x0020		/* Referenced file renamed */
//#define AG_FSEVENT_REVOKE	0x0040		/* Filesystem unmount / revoke() */
//#define AG_PROCEVENT_EXIT	0x1000		/* Process exited */
//#define AG_PROCEVENT_FORK	0x2000		/* Process forked */
//#define AG_PROCEVENT_EXEC	0x4000		/* Process exec'd */
//	AG_EventSinkFn fn;			/* Sink function */
//	AG_Event fnArgs;			/* Sink function arguments */
//	AG_TAILQ_ENTRY(ag_event_sink) sinks;    /* Epilogue "sinks" */
//} AG_EventSink;
//
///* Low-level event source */
//typedef struct ag_event_source {
type ag_event_source struct {
//	int  caps[AG_SINK_LAST];		/* Capabilities */
//	Uint flags;
//	int  breakReq;				/* Break from event loop */
//	int  returnCode;			/* AG_EventLoop() return code */
	returnCode int
//	int  (*sinkFn)(void);
//	int  (*addTimerFn)(struct ag_timer *, Uint32, int);
//	void (*delTimerFn)(struct ag_timer *);
//	int  (*resetTimerFn)(struct ag_timer *, Uint32, int);
//	AG_TAILQ_HEAD_(ag_event_sink) prologues;   /* Event prologues */
//	AG_TAILQ_HEAD_(ag_event_sink) epilogues;   /* Event sink epilogues */
//	AG_TAILQ_HEAD_(ag_event_sink) spinners;	   /* Spinning sinks */
//	AG_TAILQ_HEAD_(ag_event_sink) sinks;	   /* Normal event sinks */
//} AG_EventSource;
}

type aG_EventSource ag_event_source

//
///* Queue of events */
//typedef struct ag_event_queue {
//	Uint     nEvents;
//	AG_Event *events;
//} AG_EventQ;
//
//typedef void (*AG_EventFn)(AG_Event *);
//
//#ifdef AG_DEBUG
//#define AG_EVENT_BOUNDARY_CHECK(ev) \
//	if ((ev)->argc >= AG_EVENT_ARGS_MAX-1) \
//		AG_FatalError("Too many AG_Event(3) arguments");
//#else
//#define AG_EVENT_BOUNDARY_CHECK(ev)
//#endif
//
//#define AG_EVENT_INS_VAL(eev,tname,aname,member,val) {			\
//	AG_EVENT_BOUNDARY_CHECK(eev)					\
//	(eev)->argv[(eev)->argc].type = (tname);			\
//	if ((aname) != NULL) {						\
//		AG_Strlcpy((eev)->argv[(eev)->argc].name, (aname),	\
//		        AG_VARIABLE_NAME_MAX);				\
//	} else {							\
//		(eev)->argv[(eev)->argc].name[0] = '\0';		\
//	}								\
//	(eev)->argv[(eev)->argc].mutex = NULL;				\
//	(eev)->argv[(eev)->argc].data.member = (val);			\
//	(eev)->argv[(eev)->argc].fn.fnVoid = NULL;			\
//	(eev)->argc++;							\
//}
//#define AG_EVENT_INS_ARG(eev,ap,tname,member,t) { 			\
//	V = &(eev)->argv[(eev)->argc];					\
//	AG_EVENT_BOUNDARY_CHECK(eev)					\
//	V->type = (tname);						\
//	V->mutex = NULL;						\
//	V->data.member = va_arg(ap,t);					\
//	V->fn.fnVoid = NULL;						\
//	(eev)->argc++;							\
//}
//#define AG_EVENT_PUSH_ARG(ap,ev) {					\
//	AG_Variable *V;							\
//									\
//	switch (*c) {							\
//	case 'p':							\
//		AG_EVENT_INS_ARG((ev),ap,AG_VARIABLE_POINTER,p,void *);	\
//		break;							\
//	case 'i':							\
//		AG_EVENT_INS_ARG((ev),ap,AG_VARIABLE_INT,i,int);	\
//		break;							\
//	case 'u':							\
//		AG_EVENT_INS_ARG((ev),ap,AG_VARIABLE_UINT,i,int);	\
//		break;							\
//	case 'f':							\
//		AG_EVENT_INS_ARG((ev),ap,AG_VARIABLE_FLOAT,flt,double);	\
//		break;							\
//	case 'd':							\
//		AG_EVENT_INS_ARG((ev),ap,AG_VARIABLE_DOUBLE,dbl,double);\
//		break;							\
//	case 's':							\
//		AG_EVENT_INS_ARG((ev),ap,AG_VARIABLE_STRING,s,char *);	\
//		break;							\
//	case 'l':							\
//		switch (c[1]) {						\
//		case 'i':						\
//			AG_EVENT_INS_ARG((ev),ap,			\
//			    AG_VARIABLE_SINT32,				\
//			    s32,long);					\
//			break;						\
//		case 'u':						\
//			AG_EVENT_INS_ARG((ev),ap,			\
//			    AG_VARIABLE_UINT32,				\
//			    u32,unsigned long);				\
//			break;						\
//		default:						\
//			AG_FatalError("Bad AG_Event(3) arguments");	\
//			continue;					\
//		}							\
//		c++;							\
//		break;							\
//	case 'C':							\
//		switch (c[1]) {						\
//		case 's':						\
//			AG_EVENT_INS_ARG((ev),ap,			\
//			    AG_VARIABLE_CONST_STRING,			\
//			    Cs,const char *);				\
//			break;						\
//		case 'p':						\
//			AG_EVENT_INS_ARG((ev),ap,			\
//			    AG_VARIABLE_CONST_POINTER,			\
//			    Cp,const void *);				\
//			break;						\
//		default:						\
//			AG_FatalError("Bad AG_Event(3) arguments");	\
//			continue;					\
//		}							\
//		c++;							\
//		break;							\
//	case ' ':							\
//	case ',':							\
//	case '%':							\
//		c++;							\
//		continue;						\
//	default:							\
//		AG_FatalError("Bad AG_Event(3) argument: `%c'", *c);	\
//		c++;							\
//		continue;						\
//	}								\
//	c++;								\
//	if (*c == '(' && c[1] != '\0') {				\
//		char *cEnd;						\
//		AG_Strlcpy(V->name, &c[1], sizeof(V->name));		\
//		for (cEnd = V->name; *cEnd != '\0'; cEnd++) {		\
//			if (*cEnd == ')') {				\
//				*cEnd = '\0';				\
//				c+=2;					\
//				break;					\
//			}						\
//			c++;						\
//		}							\
//	} else {							\
//		V->name[0] = '\0';					\
//	}								\
//}
//
//#define AG_EVENT_GET_ARGS(ev, fmtp)					\
//	if (fmtp != NULL) {						\
//		const char *c = (const char *)fmtp;			\
//		va_list ap;						\
//									\
//		va_start(ap, fmtp);					\
//		while (*c != '\0') {					\
//			AG_EVENT_PUSH_ARG(ap,(ev));			\
//		}							\
//		va_end(ap);						\
//	}
//
//__BEGIN_DECLS
//int       AG_InitEventSubsystem(Uint);
//void      AG_DestroyEventSubsystem(void);
//void      AG_EventInit(AG_Event *);
//void      AG_EventArgs(AG_Event *, const char *, ...);
//AG_Event *AG_SetEvent(void *, const char *, AG_EventFn, const char *, ...);
//AG_Event *AG_AddEvent(void *, const char *, AG_EventFn, const char *, ...);
//void      AG_UnsetEvent(void *, const char *);
//void      AG_PostEvent(void *, void *, const char *, const char *, ...);
//AG_Event *AG_FindEventHandler(void *, const char *);
//
//void      AG_InitEventQ(AG_EventQ *);
//void      AG_FreeEventQ(AG_EventQ *);
//void      AG_QueueEvent(AG_EventQ *, const char *, const char *, ...);
//
//int       AG_SchedEvent(void *, void *, Uint32, const char *,
//                        const char *, ...);
//void      AG_ForwardEvent(void *, void *, AG_Event *);
//
//int             AG_EventLoop(void);
//AG_EventSource *AG_GetEventSource(void);
//AG_EventSink   *AG_AddEventPrologue(AG_EventSinkFn, const char *, ...);
//AG_EventSink   *AG_AddEventEpilogue(AG_EventSinkFn, const char *, ...);
//AG_EventSink   *AG_AddEventSpinner(AG_EventSinkFn, const char *, ...);
//AG_EventSink   *AG_AddEventSink(enum ag_event_sink_type, int, Uint,
//                                AG_EventSinkFn, const char *, ...);
//void            AG_DelEventPrologue(AG_EventSink *);
//void            AG_DelEventEpilogue(AG_EventSink *);
//void            AG_DelEventSpinner(AG_EventSink *);
//void            AG_DelEventSink(AG_EventSink *);
//void            AG_DelEventSinksByIdent(enum ag_event_sink_type, int, Uint);
//void            AG_Terminate(int);
//void            AG_TerminateEv(AG_Event *);
//
//int             AG_AddTimerKQUEUE(struct ag_timer *, Uint32, int);
//void            AG_DelTimerKQUEUE(struct ag_timer *);
//int             AG_AddTimerTIMERFD(struct ag_timer *, Uint32, int);
//void            AG_DelTimerTIMERFD(struct ag_timer *);
//int             AG_EventSinkKQUEUE(void);
//int             AG_EventSinkTIMERFD(void);
//int             AG_EventSinkTIMEDSELECT(void);
//int             AG_EventSinkSELECT(void);
//int             AG_EventSinkSPINNER(void);
//
///* Execute an event handler routine without processing any arguments. */
//static __inline__ void
//AG_ExecEventFn(void *obj, AG_Event *ev)
//{
//	if (ev->handler != NULL)
//		AG_PostEvent(NULL, obj, ev->name, NULL);
//}
//
///* Push arguments onto an Event structure. */
//static __inline__ void AG_EventPushPointer(AG_Event *ev, const char *key, void *val) { AG_EVENT_INS_VAL(ev, AG_VARIABLE_POINTER, key, p, val); }
//static __inline__ void AG_EventPushString(AG_Event *ev, const char *key, char *val)  { AG_EVENT_INS_VAL(ev, AG_VARIABLE_STRING, key, s, val); }
//static __inline__ void AG_EventPushInt(AG_Event *ev, const char *key, int val)       { AG_EVENT_INS_VAL(ev, AG_VARIABLE_INT, key, i, val); }
//static __inline__ void AG_EventPushUint(AG_Event *ev, const char *key, Uint val)     { AG_EVENT_INS_VAL(ev, AG_VARIABLE_UINT, key, i, (int)val); }
//static __inline__ void AG_EventPushLong(AG_Event *ev, const char *key, long val)     { AG_EVENT_INS_VAL(ev, AG_VARIABLE_SINT32, key, s32, (Sint32)val); }
//static __inline__ void AG_EventPushUlong(AG_Event *ev, const char *key, Ulong val)   { AG_EVENT_INS_VAL(ev, AG_VARIABLE_UINT32, key, u32, (Uint32)val); }
//static __inline__ void AG_EventPushFloat(AG_Event *ev, const char *key, float val)   { AG_EVENT_INS_VAL(ev, AG_VARIABLE_FLOAT, key, flt, val); }
//static __inline__ void AG_EventPushDouble(AG_Event *ev, const char *key, double val) { AG_EVENT_INS_VAL(ev, AG_VARIABLE_DOUBLE, key, dbl, val); }
//static __inline__ void AG_EventPopArgument(AG_Event *ev) { ev->argc--; }
//
///*
// * Accessor functions for AG_FOO_NAMED() macros.
// */
//static __inline__ AG_Variable *
//AG_GetNamedEventArg(AG_Event *ev, const char *key)
//{
//	int i;
//
//	for (i = 0; i < ev->argc; i++) {
//		if (strcmp(ev->argv[i].name, key) == 0)
//			return (&ev->argv[i]);
//	}
//	AG_FatalError("No such AG_Event argument: \"%s\"", key);
//	return (NULL);
//}
//static __inline__ void *
//AG_GetNamedPtr(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return (V->data.p);
//}
//static __inline__ char *
//AG_GetNamedString(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return (V->data.s);
//}
//static __inline__ int
//AG_GetNamedInt(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return (V->data.i);
//}
//static __inline__ Uint
//AG_GetNamedUint(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return (V->data.u);
//}
//static __inline__ long
//AG_GetNamedLong(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return ((long)V->data.s32);
//}
//static __inline__ Ulong
//AG_GetNamedUlong(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return ((Ulong)V->data.u32);
//}
//static __inline__ float
//AG_GetNamedFlt(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return (V->data.flt);
//}
//static __inline__ double
//AG_GetNamedDbl(AG_Event *event, const char *key)
//{
//	AG_Variable *V = AG_GetNamedEventArg(event, key);
//	return (V->data.dbl);
//}
//
//#ifdef AG_LEGACY
//# define AG_EVENT_SCHEDULED 0x04 /* Obsolete since AG_Timer(3) */
//# define AG_CHAR(v) ((char)event->argv[v].data.i)
//# define AG_UCHAR(v) ((Uchar)event->argv[v].data.u)
//# define AG_EventPushChar(ev,key,val) AG_EventPushInt((ev),(key),(int)(val))
//# define AG_EventPushUchar(ev,key,val) AG_EventPushUint((ev),(key),(Uint)(val))
//#endif
//__END_DECLS
//
//#include <agar/core/close.h>
//
///*
// * Copyright (c) 2001-2013 Hypertriton, Inc. <http://hypertriton.com/>
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
// * Implementation of the generic event system for AG_Object.
// */
//
//#include <core/core.h>
//
//#include <string.h>
//#include <stdarg.h>
//
//#include <config/have_kqueue.h>
//#include <config/have_timerfd.h>
//#include <config/have_select.h>
//#include <config/ag_objdebug.h>
//
//#if defined(HAVE_KQUEUE)
//# include <sys/types.h>
//# include <sys/event.h>
//# include <unistd.h>
//# include <errno.h>
//#endif
//#if defined(HAVE_TIMERFD)
//# include <sys/timerfd.h>
//# include <errno.h>
//#endif
//#if defined(HAVE_SELECT)
//# include <sys/types.h>
//# include <sys/time.h>
//# include <sys/select.h>
//# include <unistd.h>
//#endif
//
//AG_EventSource *agEventSource = NULL;	/* Event source (thread-local) */
//#ifdef AG_THREADS
//AG_ThreadKey    agEventSourceKey;
//#endif
//
//#ifdef HAVE_KQUEUE
//#define EVBUFSIZE 2
//typedef struct ag_event_source_kqueue {
//	struct ag_event_source _inherit;
//	int fd;					/* kqueue() fd */
//	struct kevent *changes;			/* Queued changes */
//	Uint          nChanges;
//	Uint        maxChanges;
//	struct kevent events[EVBUFSIZE];	/* Input event buffer */
//} AG_EventSourceKQUEUE;
//#endif /* HAVE_KQUEUE */
//
///* #define DEBUG_TIMERS */
//
///* Generate a unique event handler name. */
//static void
//GenEventName(AG_Event *ev, AG_Object *ob, const char *name)
//{
//	AG_Event *evOther;
//	int i = ob->nevents;
//
//	for (;;) {
//		ev->name[0] = '_';
//		StrlcpyUint(&ev->name[1], i++, sizeof(ev->name));
//
//		TAILQ_FOREACH(evOther, &ob->events, events) {
//			if (strcmp(evOther->name, ev->name) == 0)
//				break;
//		}
//		if (evOther == NULL)
//			break;
//	}
//}
//
///* Initialize a pointer argument. */
//static __inline__ void
//InitPointerArg(AG_Variable *V, void *p)
//{
//	V->name[0] = '\0';
//	V->type = AG_VARIABLE_POINTER;
//	V->fn.fnVoid = NULL;
//	V->mutex = NULL;
//	V->data.p = p;
//}
//
//static __inline__ void
//InitEvent(AG_Event *ev, AG_Object *ob)
//{
//	ev->flags = 0;
//	ev->argc = 1;
//	ev->argc0 = 1;
//	ev->handler = NULL;
//	InitPointerArg(&ev->argv[0], ob);
//}
//
///* Initialize an AG_Event structure. */
//void
//AG_EventInit(AG_Event *ev)
//{
//	InitEvent(ev, NULL);
//}
//
///* Initialize an AG_Event structure with the specified arguments. */
//void
//AG_EventArgs(AG_Event *ev, const char *fmt, ...)
//{
//	InitEvent(ev, NULL);
//	AG_EVENT_GET_ARGS(ev, fmt);
//	ev->argc0 = ev->argc;
//}
//
///* Set (or change) the event handler function for the named event. */
//AG_Event *
//AG_SetEvent(void *p, const char *name, AG_EventFn fn, const char *fmt, ...)
//{
//	AG_Object *ob = p;
//	AG_Event *ev;
//
//	AG_ObjectLock(ob);
//
//	if (name != NULL) {
//		TAILQ_FOREACH(ev, &ob->events, events)
//			if (strcmp(ev->name, name) == 0)
//				break;
//	} else {
//		ev = NULL;
//	}
//	if (ev == NULL) {
//		ev = Malloc(sizeof(AG_Event));
//		InitEvent(ev, ob);
//		if (name != NULL) {
//			Strlcpy(ev->name, name, sizeof(ev->name));
//		} else {
//			GenEventName(ev, ob, name);
//		}
//		TAILQ_INSERT_TAIL(&ob->events, ev, events);
//		ob->nevents++;
//	} else {				/* Preserve flags */
//		ev->argc = 1;
//		ev->argc0 = 1;
//	}
//	InitPointerArg(&ev->argv[0], ob);
//	ev->handler = fn;
//	AG_EVENT_GET_ARGS(ev, fmt);
//	ev->argc0 = ev->argc;
//
//	AG_ObjectUnlock(ob);
//	return (ev);
//}
//
///* Append an event handler function for the named event. */
//AG_Event *
//AG_AddEvent(void *p, const char *name, AG_EventFn fn, const char *fmt, ...)
//{
//	AG_Object *ob = p;
//	AG_Event *ev, *evOther;
//
//	AG_ObjectLock(ob);
//
//	ev = Malloc(sizeof(AG_Event));
//	InitEvent(ev, ob);
//
//	if (name != NULL) {
//		TAILQ_FOREACH(evOther, &ob->events, events) {
//			if (strcmp(evOther->name, name) == 0)
//				break;
//		}
//		if (evOther != NULL) {
//			ev->flags = evOther->flags;
//		}
//		Strlcpy(ev->name, name, sizeof(ev->name));
//	} else {
//		GenEventName(ev, ob, name);
//	}
//
//	ev->handler = fn;
//	AG_EVENT_GET_ARGS(ev, fmt);
//	ev->argc0 = ev->argc;
//
//	TAILQ_INSERT_TAIL(&ob->events, ev, events);
//	ob->nevents++;
//
//	AG_ObjectUnlock(ob);
//	return (ev);
//}
//
///* Remove the named event handler. */
//void
//AG_UnsetEvent(void *p, const char *name)
//{
//	AG_Object *ob = p;
//	AG_Event *ev;
//
//	AG_ObjectLock(ob);
//	TAILQ_FOREACH(ev, &ob->events, events) {
//		if (strcmp(name, ev->name) == 0)
//			break;
//	}
//	if (ev == NULL) {
//		goto out;
//	}
//	TAILQ_REMOVE(&ob->events, ev, events);
//	ob->nevents--;
//	Free(ev);
//out:
//	AG_ObjectUnlock(ob);
//}
//
///* Return the Event structure for the named event. */
//AG_Event *
//AG_FindEventHandler(void *p, const char *name)
//{
//	AG_Object *ob = p;
//	AG_Event *ev;
//	
//	AG_ObjectLock(ob);
//	TAILQ_FOREACH(ev, &ob->events, events) {
//		if (strcmp(name, ev->name) == 0)
//			break;
//	}
//	AG_ObjectUnlock(ob);
//	return (ev);
//}
//
///* Forward an event to an object's descendents. */
//static void
//PropagateEvent(AG_Object *sndr, AG_Object *rcvr, AG_Event *ev)
//{
//	AG_Object *chld;
//
//	OBJECT_FOREACH_CHILD(chld, rcvr, ag_object) {
//		PropagateEvent(rcvr, chld, ev);
//	}
//	AG_ForwardEvent(sndr, rcvr, ev);
//}
//
///* Timeout callback for scheduled events. */
//static Uint32
//EventTimeout(AG_Timer *to, AG_Event *event)
//{
//	AG_Object *ob = AG_SELF();
//	AG_Object *obSender = AG_PTR(1);
//	char *eventName = AG_STRING(2);
//	AG_Event *ev;
//
//#ifdef AG_OBJDEBUG
//	if (agDebugLvl >= 5)
//		Debug(ob, "Event <%s> timeout (%u ticks)\n", eventName,
//		    (Uint)to->ival);
//#endif
//	TAILQ_FOREACH(ev, &ob->events, events) {
//		if (strcmp(eventName, ev->name) == 0)
//			break;
//	}
//	if (ev == NULL) {
//		return (0);
//	}
//	InitPointerArg(&ev->argv[ev->argc], obSender);
//
//	/* Propagate event to children. */
//	if (ev->flags & AG_EVENT_PROPAGATE) {
//		AG_Object *child;
//#ifdef AG_OBJDEBUG
//		if (agDebugLvl >= 5)
//			Debug(ob, "Propagate <%s> (timeout)\n", ev->name);
//#endif
//		AG_LockVFS(ob);
//		OBJECT_FOREACH_CHILD(child, ob, ag_object) {
//			PropagateEvent(ob, child, ev);
//		}
//		AG_UnlockVFS(ob);
//	}
//
//	/* Invoke the event handler routine. */
//	if (ev->handler != NULL) {
//		ev->handler(ev);
//	}
//	return (0);
//}
//
//
//#ifdef AG_THREADS
///* Invoke an event handler routine asynchronously. */
//static void *
//EventThread(void *p)
//{
//	AG_Event *eev = p;
//	AG_Object *rcvr = eev->argv[0].data.p;
//	AG_Object *chld;
//
//	if (eev->flags & AG_EVENT_PROPAGATE) {
//#ifdef AG_OBJDEBUG
//		if (agDebugLvl >= 5)
//			Debug(rcvr, "Propagate <%s> (async)\n", eev->name);
//#endif
//		AG_LockVFS(rcvr);
//		OBJECT_FOREACH_CHILD(chld, rcvr, ag_object) {
//			PropagateEvent(rcvr, chld, eev);
//		}
//		AG_UnlockVFS(rcvr);
//	}
//#ifdef AG_OBJDEBUG
//	if (agDebugLvl >= 5)
//		Debug(rcvr, "BEGIN event thread for <%s>\n", eev->name);
//#endif
//	if (eev->handler != NULL) {
//		eev->handler(eev);
//	}
//#ifdef AG_OBJDEBUG
//	if (agDebugLvl >= 5)
//		Debug(rcvr, "CLOSE event thread for <%s>\n", eev->name);
//#endif
//	Free(eev);
//	return (NULL);
//}
//#endif /* AG_THREADS */
//
//void
//AG_InitEventQ(AG_EventQ *eq)
//{
//	eq->nEvents = 0;
//	eq->events = NULL;
//}
//
//void
//AG_FreeEventQ(AG_EventQ *eq)
//{
//	Free(eq->events);
//	eq->nEvents = 0;
//	eq->events = NULL;
//}
//
///* Add a new entry to an event queue. */
//void
//AG_QueueEvent(AG_EventQ *eq, const char *evname, const char *fmt, ...)
//{
//	AG_Event *ev;
//
//	eq->events = Realloc(eq->events, (eq->nEvents+1)*sizeof(AG_Event));
//	ev = &eq->events[eq->nEvents++];
//	InitEvent(ev, NULL);
//	AG_EVENT_GET_ARGS(ev, fmt);
//	ev->argc0 = ev->argc;
//}
//
///*
// * Raise the specified event. Configured event handler routines may be
// * called immediately, but they may also get called from a separate
// * thread, or queued for later execution.
// *
// * The argument vector passed to the event handler function contains
// * the AG_SetEvent() arguments, and any arguments specified here are
// * appended to that list.
// */
//void
//AG_PostEvent(void *sp, void *rp, const char *evname, const char *fmt, ...)
//{
//	AG_Object *sndr = sp;
//	AG_Object *rcvr = rp;
//	AG_Event *ev;
//	AG_Object *chld;
//	int propagated = 0;
//
//#ifdef AG_OBJDEBUG
//	if (agDebugLvl >= 5)
//		Debug(rcvr, "Event <%s> posted from %s\n", evname,
//		    sndr?sndr->name:"NULL");
//#endif
//	AG_ObjectLock(rcvr);
//	TAILQ_FOREACH(ev, &rcvr->events, events) {
//		if (strcmp(evname, ev->name) != 0)
//			continue;
//#ifdef AG_THREADS
//		if (ev->flags & AG_EVENT_ASYNC) {
//			AG_Thread th;
//			AG_Event *evNew;
//
//			/* TODO allocate from an per-object pool */
//			evNew = Malloc(sizeof(AG_Event));
//			memcpy(evNew, ev, sizeof(AG_Event));
//			AG_EVENT_GET_ARGS(evNew, fmt);
//			InitPointerArg(&evNew->argv[evNew->argc], sndr);
//			if (evNew->flags & AG_EVENT_PROPAGATE) { propagated = 1; }
//			if (propagated) {
//				evNew->flags &= ~(AG_EVENT_PROPAGATE);
//			}
//			AG_ThreadCreate(&th, EventThread, evNew);
//		} else
//#endif /* AG_THREADS */
//		{
//			AG_Event tmpev;
//
//			memcpy(&tmpev, ev, sizeof(AG_Event));
//			AG_EVENT_GET_ARGS(&tmpev, fmt);
//			InitPointerArg(&tmpev.argv[tmpev.argc], sndr);
//			if ((tmpev.flags & AG_EVENT_PROPAGATE) && !propagated) {
//#ifdef AG_OBJDEBUG
//				if (agDebugLvl >= 5)
//					Debug(rcvr, "Propagate <%s> (post)\n",
//					    evname);
//#endif
//				AG_LockVFS(rcvr);
//				OBJECT_FOREACH_CHILD(chld, rcvr, ag_object) {
//					PropagateEvent(rcvr, chld, &tmpev);
//				}
//				AG_UnlockVFS(rcvr);
//				propagated = 1;
//			}
//			if (tmpev.handler != NULL)
//				tmpev.handler(&tmpev);
//		}
//	}
//	AG_ObjectUnlock(rcvr);
//}
//
///*
// * Schedule the execution of the named event in the given number
// * of AG_Time(3) ticks.
// *
// * The argument vector passed to the event handler function contains
// * the AG_SetEvent() arguments, and any arguments specified here are
// * appended to that list.
// */
//int
//AG_SchedEvent(void *pSndr, void *pRcvr, Uint32 ticks, const char *evname,
//    const char *fmt, ...)
//{
//	AG_Object *sndr = pSndr;
//	AG_Object *rcvr = pRcvr;
//	AG_Event *ev;
//	AG_Timer *to;
//
//	if ((to = TryMalloc(sizeof(AG_Timer))) == NULL) {
//		return (-1);
//	}
//	AG_InitTimer(to, evname, AG_TIMER_AUTO_FREE);
//
//	AG_LockTiming();
//	AG_ObjectLock(rcvr);
//	
//	if (AG_AddTimer(rcvr, to, ticks,
//	    EventTimeout, "%p,%s", sndr, evname) == -1) {
//		free(to);
//		goto fail;
//	}
//	ev = &to->fnEvent;
//	AG_EventInit(ev);
//	ev->argv[0].data.p = rcvr;
//	AG_EVENT_GET_ARGS(ev, fmt);
//	ev->argc0 = ev->argc;
//
//	AG_UnlockTiming();
//	AG_ObjectUnlock(rcvr);
//	return (0);
//fail:
//	AG_UnlockTiming();
//	AG_ObjectUnlock(rcvr);
//	return (-1);
//}
//
///*
// * Forward an event, without modifying the original event structure, except
// * for the sender and receiver pointers.
// */
//void
//AG_ForwardEvent(void *pSndr, void *pRcvr, AG_Event *event)
//{
//	AG_Object *sndr = pSndr;
//	AG_Object *rcvr = pRcvr;
//	AG_Object *chld;
//	AG_Event *ev;
//
//#ifdef AG_OBJDEBUG
//	if (agDebugLvl >= 5)
//		Debug(rcvr, "Event <%s> forwarded from %s\n", event->name,
//		    sndr?sndr->name:"NULL");
//#endif
//	AG_ObjectLock(rcvr);
//	TAILQ_FOREACH(ev, &rcvr->events, events) {
//		if (strcmp(event->name, ev->name) != 0)
//			continue;
//#ifdef AG_THREADS
//		if (ev->flags & AG_EVENT_ASYNC) {
//			AG_Thread th;
//			AG_Event *evNew;
//
//			/* TODO allocate from an per-object pool */
//			evNew = Malloc(sizeof(AG_Event));
//			memcpy(evNew, ev, sizeof(AG_Event));
//			InitPointerArg(&evNew->argv[0], rcvr);
//			InitPointerArg(&evNew->argv[evNew->argc], sndr);
//			AG_ThreadCreate(&th, EventThread, evNew);
//		} else
//#endif /* AG_THREADS */
//		{
//			AG_Event tmpev;
//
//			memcpy(&tmpev, event, sizeof(AG_Event));
//			InitPointerArg(&tmpev.argv[0], rcvr);
//			InitPointerArg(&tmpev.argv[tmpev.argc], sndr);
//
//			if (ev->flags & AG_EVENT_PROPAGATE) {
//#ifdef AG_OBJDEBUG
//				if (agDebugLvl >= 5)
//					Debug(rcvr, "Propagate <%s> (forward)\n",
//					    event->name);
//#endif
//				AG_LockVFS(rcvr);
//				OBJECT_FOREACH_CHILD(chld, rcvr, ag_object) {
//					PropagateEvent(rcvr, chld, ev);
//				}
//				AG_UnlockVFS(rcvr);
//			}
//			/* XXX AG_EVENT_ASYNC.. */
//			if (ev->handler != NULL)
//				ev->handler(&tmpev);
//		}
//	}
//	AG_ObjectUnlock(rcvr);
//}
//
//#ifdef HAVE_KQUEUE
//static __inline__ int
//GrowKqChangelist(AG_EventSourceKQUEUE *kq, Uint n)
//{
//	struct kevent *changesNew;
//
//	if (n <= kq->maxChanges) {
//		return (0);
//	}
//	if ((changesNew = TryRealloc(kq->changes, n*sizeof(struct kevent)))
//	    == NULL) {
//		return (-1);
//	}
//	kq->changes = changesNew;
//	kq->maxChanges = n;
//	return (0);
//}
//#endif /* HAVE_KQUEUE */
//
///* Create a new event source. */
//static AG_EventSource *
//CreateEventSource(void)
//{
//#ifdef HAVE_KQUEUE
//	AG_EventSourceKQUEUE *kq = TryMalloc(sizeof(AG_EventSourceKQUEUE));
//	AG_EventSource *src = (AG_EventSource *)kq;
//#else
//	AG_EventSource *src = TryMalloc(sizeof(AG_EventSource));
//#endif
//	if (src == NULL) {
//		return (NULL);
//	}
//	src->flags = 0;
//	src->addTimerFn = NULL;
//	src->delTimerFn = NULL;
//	src->breakReq = 0;
//	src->returnCode = 0;
//	TAILQ_INIT(&src->prologues);
//	TAILQ_INIT(&src->epilogues);
//	TAILQ_INIT(&src->spinners);
//	TAILQ_INIT(&src->sinks);
//	memset(src->caps, 0, sizeof(src->caps));
//
//#if defined(HAVE_KQUEUE)
//	if ((kq->fd = kqueue()) == -1) {
//		AG_SetError("kqueue: %s", AG_Strerror(errno));
//		return (NULL);
//	}
//	kq->changes = NULL;
//	kq->nChanges = 0;
//	kq->maxChanges = 0;
//	memset(kq->events, 0, EVBUFSIZE*sizeof(struct kevent));
//	src->sinkFn = AG_EventSinkKQUEUE;
//	src->addTimerFn = AG_AddTimerKQUEUE;
//	src->delTimerFn = AG_DelTimerKQUEUE;
//	src->caps[AG_SINK_TIMER] = 1;		/* Provides timers internally */
//	src->caps[AG_SINK_READ] = 1;
//	src->caps[AG_SINK_WRITE] = 1;
//	src->caps[AG_SINK_FSEVENT] = 1;
//	src->caps[AG_SINK_PROCEVENT] = 1;
//#elif defined(HAVE_TIMERFD)
//	src->sinkFn = AG_EventSinkTIMERFD;
//	src->addTimerFn = AG_AddTimerTIMERFD;
//	src->delTimerFn = AG_DelTimerTIMERFD;
//	src->caps[AG_SINK_TIMER] = 1;		/* Provides timers internally */
//	src->caps[AG_SINK_READ] = 1;
//	src->caps[AG_SINK_WRITE] = 1;
//#elif defined(HAVE_SELECT) && !defined(AG_THREADS)
//	src->sinkFn = AG_EventSinkTIMEDSELECT;
//	src->caps[AG_SINK_READ] = 1;
//	src->caps[AG_SINK_WRITE] = 1;
//#elif defined(HAVE_SELECT) && defined(AG_THREADS)
//	src->sinkFn = AG_EventSinkSELECT;
//	src->caps[AG_SINK_READ] = 1;
//	src->caps[AG_SINK_WRITE] = 1;
//#else
//	src->sinkFn = AG_EventSinkSPINNER;
//#endif
//	return (src);
//}
//
//static void
//DestroyEventSource(void *pEventSource)
//{
//	AG_EventSource *src = pEventSource;
//	AG_EventSink *es, *esNext;
//#ifdef HAVE_KQUEUE
//	AG_EventSourceKQUEUE *kq = pEventSource;
//
//	if (kq->fd != -1) {
//		close(kq->fd);
//	}
//	Free(kq->changes);
//#endif
//	for (es = TAILQ_FIRST(&src->prologues); es != TAILQ_END(&src->prologues); es = esNext) {
//		esNext = TAILQ_NEXT(es, sinks);
//		free(es);
//	}
//	for (es = TAILQ_FIRST(&src->epilogues); es != TAILQ_END(&src->epilogues); es = esNext) {
//		esNext = TAILQ_NEXT(es, sinks);
//		free(es);
//	}
//	for (es = TAILQ_FIRST(&src->spinners); es != TAILQ_END(&src->spinners); es = esNext) {
//		esNext = TAILQ_NEXT(es, sinks);
//		free(es);
//	}
//	for (es = TAILQ_FIRST(&src->sinks); es != TAILQ_END(&src->sinks); es = esNext) {
//		esNext = TAILQ_NEXT(es, sinks);
//		free(es);
//	}
//	free(src);
//}
//
///* Return the calling thread's effective event source. */
//AG_EventSource *
//AG_GetEventSource(void)
//{
func aG_GetEventSource() aG_EventSource {
//	AG_EventSource *src;
	var src aG_EventSource
//
//#ifdef AG_THREADS
//	if ((src = AG_ThreadKeyGet(agEventSourceKey)) != NULL && src != NULL)
//		return (src);
//#else
//	if (agEventSource != NULL)
//		return (agEventSource);
//#endif
//	if ((src = CreateEventSource()) == NULL)
//		AG_FatalError(NULL);
//#ifdef AG_THREADS
//	AG_ThreadKeySet(agEventSourceKey, src);
//#else
//	agEventSource = src;
//#endif
//	return (src);
	return src
//}
}
//
//int
//AG_InitEventSubsystem(Uint flags)
//{
//	/* Initialize the main thread's event source. */
//	agEventSource = NULL;
//#ifdef AG_THREADS
//	if (AG_ThreadKeyTryCreate(&agEventSourceKey, DestroyEventSource) == -1)
//		return (-1);
//#endif
//	if ((agEventSource = AG_GetEventSource()) == NULL) {
//		return (-1);
//	}
//	return (0);
//}
//
//void
//AG_DestroyEventSubsystem(void)
//{
//	if (agEventSource != NULL) {
//		DestroyEventSource(agEventSource);
//		agEventSource = NULL;
//	}
//}
//
//#ifdef HAVE_KQUEUE
///*
// * Routines for translating between AG_EventSink and kqueue types.
// */
//static __inline__ enum ag_event_sink_type
//GetSinkType(int filter)
//{
//	switch (filter) {
//	case EVFILT_TIMER:	return (AG_SINK_TIMER);
//	case EVFILT_READ:	return (AG_SINK_READ);
//	case EVFILT_WRITE:	return (AG_SINK_WRITE);
//	case EVFILT_VNODE:	return (AG_SINK_FSEVENT);
//	case EVFILT_PROC:	return (AG_SINK_PROCEVENT);
//	default:		return (AG_SINK_NONE);
//	}
//}
//static Uint
//GetKqFilterFlags(Uint flags)
//{
//	Uint fflags = 0;
//	if (flags & AG_FSEVENT_DELETE) { fflags |= NOTE_DELETE; }
//	if (flags & AG_FSEVENT_WRITE)  { fflags |= NOTE_WRITE;  }
//	if (flags & AG_FSEVENT_EXTEND) { fflags |= NOTE_EXTEND; }
//	if (flags & AG_FSEVENT_ATTRIB) { fflags |= NOTE_ATTRIB; }
//	if (flags & AG_FSEVENT_LINK)   { fflags |= NOTE_LINK;   }
//	if (flags & AG_FSEVENT_RENAME) { fflags |= NOTE_RENAME; }
//	if (flags & AG_FSEVENT_REVOKE) { fflags |= NOTE_REVOKE; }
//	if (flags & AG_PROCEVENT_EXIT) { fflags |= NOTE_EXIT; }
//	if (flags & AG_PROCEVENT_FORK) { fflags |= NOTE_FORK; }
//	if (flags & AG_PROCEVENT_EXEC) { fflags |= NOTE_EXEC; }
//	return (fflags);
//}
//static Uint
//GetSinkFlags(Uint fflags)
//{
//	Uint flags = 0;
//	if (fflags & NOTE_DELETE) { fflags |= AG_FSEVENT_DELETE; }
//	if (fflags & NOTE_WRITE)  { fflags |= AG_FSEVENT_WRITE;  }
//	if (fflags & NOTE_EXTEND) { fflags |= AG_FSEVENT_EXTEND; }
//	if (fflags & NOTE_ATTRIB) { fflags |= AG_FSEVENT_ATTRIB; }
//	if (fflags & NOTE_LINK)   { fflags |= AG_FSEVENT_LINK;   }
//	if (fflags & NOTE_RENAME) { fflags |= AG_FSEVENT_RENAME; }
//	if (fflags & NOTE_REVOKE) { fflags |= AG_FSEVENT_REVOKE; }
//	if (fflags & NOTE_EXIT) { fflags |= AG_PROCEVENT_EXIT; }
//	if (fflags & NOTE_FORK) { fflags |= AG_PROCEVENT_FORK; }
//	if (fflags & NOTE_EXEC) { fflags |= AG_PROCEVENT_EXEC; }
//	return (flags);
//}
//#endif /* HAVE_KQUEUE */
//
///*
// * Add/remove an event processing prologue. The function will be invoked
// * only once at the beginning of AG_EventLoop().
// */
//AG_EventSink *
//AG_AddEventPrologue(AG_EventSinkFn fn, const char *fnArgs, ...)
//{
//	AG_EventSource *src = AG_GetEventSource();
//	AG_EventSink *es;
//
//	if ((es = TryMalloc(sizeof(AG_EventSink))) == NULL) {
//		return (NULL);
//	}
//	es->type = AG_SINK_PROLOGUE;
//	es->fn = fn;
//	InitEvent(&es->fnArgs, NULL);
//	AG_EVENT_GET_ARGS(&es->fnArgs, fnArgs);
//	es->fnArgs.argc0 = es->fnArgs.argc;
//	TAILQ_INSERT_TAIL(&src->prologues, es, sinks);
//	return (es);
//}
//void
//AG_DelEventPrologue(AG_EventSink *es)
//{
//	AG_EventSource *src = AG_GetEventSource();
//
//#ifdef AG_DEBUG
//	if (es->type != AG_SINK_PROLOGUE)
//		AG_FatalError("AG_DelEventPrologue");
//#endif
//	TAILQ_REMOVE(&src->prologues, es, sinks);
//	free(es);
//}
//
///*
// * Add/remove an event sink epilogue. The function will be invoked
// * after all incoming / queued events have been processed.
// */
//AG_EventSink *
//AG_AddEventEpilogue(AG_EventSinkFn fn, const char *fnArgs, ...)
//{
//	AG_EventSource *src = AG_GetEventSource();
//	AG_EventSink *es;
//
//	if ((es = TryMalloc(sizeof(AG_EventSink))) == NULL) {
//		return (NULL);
//	}
//	es->type = AG_SINK_EPILOGUE;
//	es->fn = fn;
//	InitEvent(&es->fnArgs, NULL);
//	AG_EVENT_GET_ARGS(&es->fnArgs, fnArgs);
//	es->fnArgs.argc0 = es->fnArgs.argc;
//	TAILQ_INSERT_TAIL(&src->epilogues, es, sinks);
//	return (es);
//}
//void
//AG_DelEventEpilogue(AG_EventSink *es)
//{
//	AG_EventSource *src = AG_GetEventSource();
//
//#ifdef AG_DEBUG
//	if (es->type != AG_SINK_EPILOGUE)
//		AG_FatalError("AG_DelEventEpilogue");
//#endif
//	TAILQ_REMOVE(&src->epilogues, es, sinks);
//	free(es);
//}
//
///*
// * Add/remove a spinning event sink. If at least one spinning sink exists,
// * AG_EventLoop() will invoke it repeatedly and force non-blocking operation
// * of subsequent polling methods.
// */
//AG_EventSink *
//AG_AddEventSpinner(AG_EventSinkFn fn, const char *fnArgs, ...)
//{
//	AG_EventSource *src = AG_GetEventSource();
//	AG_EventSink *es;
//
//	if ((es = TryMalloc(sizeof(AG_EventSink))) == NULL) {
//		return (NULL);
//	}
//	es->type = AG_SINK_SPINNER;
//	es->fn = fn;
//	InitEvent(&es->fnArgs, NULL);
//	AG_EVENT_GET_ARGS(&es->fnArgs, fnArgs);
//	es->fnArgs.argc0 = es->fnArgs.argc;
//	TAILQ_INSERT_TAIL(&src->spinners, es, sinks);
//	return (es);
//}
//void
//AG_DelEventSpinner(AG_EventSink *es)
//{
//	AG_EventSource *src = AG_GetEventSource();
//
//#ifdef AG_DEBUG
//	if (es->type != AG_SINK_SPINNER)
//		AG_FatalError("AG_DelEventSpinner");
//#endif
//	TAILQ_REMOVE(&src->spinners, es, sinks);
//	free(es);
//}
//
///*
// * Add/remove a low-level event sink. The function will be called
// * whenever the specified event occurs.
// */
//AG_EventSink *
//AG_AddEventSink(enum ag_event_sink_type type, int ident, Uint flags,
//    AG_EventSinkFn fn, const char *fnArgs, ...)
//{
//	AG_EventSource *src = AG_GetEventSource();
//	AG_EventSink *es;
//#ifdef HAVE_KQUEUE
//	AG_EventSourceKQUEUE *kq = (AG_EventSourceKQUEUE *)src;
//	struct kevent *kev;
//#endif
//	if (type >= AG_SINK_LAST || !src->caps[type]) {
//		AG_SetError("Unsupported event type: %u", (Uint)type);
//		return (NULL);
//	}
//	if ((es = TryMalloc(sizeof(AG_EventSink))) == NULL) {
//		return (NULL);
//	}
//	es->type = type;
//	es->ident = ident;
//	es->flags = flags;
//
//#ifdef HAVE_KQUEUE
//	if (GrowKqChangelist(kq, kq->nChanges+1) == -1) {
//		free(es);
//		return (NULL);
//	}
//	kev = &kq->changes[kq->nChanges++];
//	switch (type) {
//	case AG_SINK_READ:
//		EV_SET(kev, ident, EVFILT_READ, EV_ADD|EV_ENABLE, 0, 0, es);
//		break;
//	case AG_SINK_WRITE:
//		EV_SET(kev, ident, EVFILT_WRITE, EV_ADD|EV_ENABLE, 0, 0, es);
//		break;
//	case AG_SINK_FSEVENT:
//		EV_SET(kev, ident, EVFILT_VNODE, EV_ADD|EV_ENABLE,
//		    GetKqFilterFlags(flags), 0, es);
//		break;
//	case AG_SINK_PROCEVENT:
//		EV_SET(kev, ident, EVFILT_PROC, EV_ADD|EV_ENABLE,
//		    GetKqFilterFlags(flags), 0, es);
//		break;
//	default:
//		kq->nChanges--;
//		break;
//	}
//#endif /* HAVE_KQUEUE */
//
//	es->fn = fn;
//	InitEvent(&es->fnArgs, NULL);
//	AG_EVENT_GET_ARGS(&es->fnArgs, fnArgs);
//	es->fnArgs.argc0 = es->fnArgs.argc;
//	TAILQ_INSERT_TAIL(&src->sinks, es, sinks);
//	return (es);
//}
//void
//AG_DelEventSink(AG_EventSink *es)
//{
//	AG_EventSource *src = AG_GetEventSource();
//#ifdef HAVE_KQUEUE
//	AG_EventSourceKQUEUE *kq = (AG_EventSourceKQUEUE *)src;
//	struct kevent *kev;
//
//	if (GrowKqChangelist(kq, kq->nChanges+1) == -1) {
//		AG_FatalError(NULL);
//	}
//	kev = &kq->changes[kq->nChanges++];
//
//	switch (es->type) {
//	case AG_SINK_READ:
//		EV_SET(kev, es->ident, EVFILT_READ, EV_DELETE, 0, 0, NULL);
//		break;
//	case AG_SINK_WRITE:
//		EV_SET(kev, es->ident, EVFILT_WRITE, EV_DELETE, 0, 0, NULL);
//		break;
//	case AG_SINK_FSEVENT:
//		EV_SET(kev, es->ident, EVFILT_VNODE, EV_DELETE,
//		    GetKqFilterFlags(es->flags), 0, NULL);
//		break;
//	case AG_SINK_PROCEVENT:
//		EV_SET(kev, es->ident, EVFILT_PROC, EV_DELETE,
//		    GetKqFilterFlags(es->flags), 0, NULL);
//		break;
//	default:
//		kq->nChanges--;
//		break;
//	}
//#endif /* HAVE_KQUEUE */
//
//	TAILQ_REMOVE(&src->sinks, es, sinks);
//	free(es);
//}
//void
//AG_DelEventSinksByIdent(enum ag_event_sink_type type, int ident, Uint flags)
//{
//	AG_EventSource *src = AG_GetEventSource();
//	AG_EventSink *es, *esNext;
//
//	for (es = TAILQ_FIRST(&src->sinks); es != TAILQ_END(&src->sinks); es = esNext) {
//		esNext = TAILQ_NEXT(es, sinks);
//		if (es->type == type &&
//		    es->ident == ident &&
//		    es->flags == flags)
//			AG_DelEventSink(es);
//	}
//}
//
//#ifdef HAVE_KQUEUE
///*
// * Standard event sink using kqueue(2), commonly found on modern BSD
// * derived operating systems. 
// */
//int
//AG_EventSinkKQUEUE(void)
//{
//	AG_EventSourceKQUEUE *kq = (AG_EventSourceKQUEUE *)agEventSource;
//	int rv, i;
//	struct timespec timeo, *pTimeo;
//
//restart:
//	if (!TAILQ_EMPTY(&agEventSource->spinners)) {
//		timeo.tv_sec = 0;
//		timeo.tv_nsec = 0L;
//		pTimeo = &timeo;
//	} else {
//		pTimeo = NULL;
//	}
//#ifdef DEBUG_TIMERS
//	for (i = 0; i < kq->nChanges; i++) {
//		struct kevent *chg = &kq->changes[i];
//		Verbose("changes[%d]: f=%d i=%u f=0x%x ff=0x%x u=%p\n",
//		    i, (int)chg->filter,
//		    (Uint)chg->ident, chg->flags, chg->fflags,
//		    chg->udata);
//	}
//#endif
//	rv = kevent(kq->fd, kq->changes, kq->nChanges, kq->events, EVBUFSIZE,
//	    pTimeo);
//	if (rv < 0) {
//		if (errno == EINTR) {
//			goto restart;
//		}
//		AG_SetError("kevent(): %s", AG_Strerror(errno));
//		return (-1);
//	}
//	kq->nChanges = 0;
//	AG_LockTiming();
//	/* 1. Process timer expirations. */
//	for (i = 0; i < rv; i++) {
//		struct kevent *kev = &kq->events[i];
//		enum ag_event_sink_type esType = GetSinkType(kev->filter);
//		Uint32 rvt;
//		AG_Timer *to;
//
//		if (kev->flags & EV_ERROR) {
//			AG_SetError("kevent[%d]: %s", i, AG_Strerror(kev->data));
//			return (-1);
//		}
//		if (esType != AG_SINK_TIMER) {
//			continue;
//		}
//		to = (AG_Timer *)kev->udata;
//		rvt = to->fn(to, &to->fnEvent);
//		if (rvt > 0) {
//#ifdef DEBUG_TIMERS
//			Verbose("TIMER[%d] resetting t=+%u\n", to->id, (Uint)rvt);
//#endif
//			if (GrowKqChangelist(kq, kq->nChanges+1) == -1) {
//				return (-1);
//			}
//			EV_SET(&kq->changes[kq->nChanges++], to->id, EVFILT_TIMER,
//			    EV_ADD|EV_ENABLE|EV_ONESHOT, 0, (int)rvt, to);
//			to->ival = rvt;
//		} else {
//			AG_Object *ob = to->obj;
//#ifdef DEBUG_TIMERS
//			Verbose("TIMER[%d] expired\n", to->id);
//#endif
//			TAILQ_REMOVE(&ob->timers, to, timers);
//			if (TAILQ_EMPTY(&ob->timers)) {
//				TAILQ_REMOVE(&agTimerObjQ, ob, tobjs);
//			}
//			if (to->flags & AG_TIMER_AUTO_FREE) {
//				free(to);
//			} else {
//				to->ival = 0;
//				to->id = -1;
//				to->obj = NULL;
//			}
//			agTimerCount--;
//		}
//	}
//	/* 2. Process I/O and other events. */
//	for (i = 0; i < rv; i++) {
//		struct kevent *kev = &kq->events[i];
//		enum ag_event_sink_type esType = GetSinkType(kev->filter);
//		AG_EventSink *es;
//
//		switch (esType) {
//		case AG_SINK_READ:
//		case AG_SINK_WRITE:
//			es = (AG_EventSink *)kev->udata;
//			es->fn(es, &es->fnArgs);
//			break;
//		case AG_SINK_FSEVENT:
//		case AG_SINK_PROCEVENT:
//			es = (AG_EventSink *)kev->udata;
//			es->flagsMatched = GetSinkFlags(kev->fflags);
//			es->fn(es, &es->fnArgs);
//			break;
//		default:
//			break;
//		}
//	}
//	AG_UnlockTiming();
//	return (0);
//}
//
///*
// * Add/remove a kqueue(2) based timer.
// */
//static int
//GenerateTimerID(AG_Timer *to)
//{
//	AG_Object *obOther;
//	AG_Timer *toOther;
//	int id;
//
//gen_id:
//#ifdef AG_DEBUG
//	if (agTimerCount+1 >= (AG_INT_MAX-1))
//		AG_FatalError("agTimerCount");
//#endif
//	id = (int)++agTimerCount;			/* XXX */
//	TAILQ_FOREACH(obOther, &agTimerObjQ, tobjs) {
//		TAILQ_FOREACH(toOther, &obOther->timers, timers) {
//			if (toOther == to) { continue; }
//			if (toOther->id == id) {
//				id++;
//				goto gen_id;
//			}
//		}
//	}
//	return (id);
//}
//int
//AG_AddTimerKQUEUE(AG_Timer *to, Uint32 ival, int newTimer)
//{
//	AG_EventSourceKQUEUE *kq = (AG_EventSourceKQUEUE *)agEventSource;
//	
//	/* Create a kernel-based timer with kqueue. */
//	if (newTimer) {
//		to->id = GenerateTimerID(to);
//	}
//	if (newTimer || to->ival != ival) {
//#ifdef DEBUG_TIMERS
//		Verbose("kevent: creating timer ID=%d ival=%d\n", to->id, (int)ival);
//#endif
//		if (GrowKqChangelist(kq, kq->nChanges+1) == -1) {
//			return (-1);
//		}
//		EV_SET(&kq->changes[kq->nChanges++], to->id, EVFILT_TIMER,
//		    EV_ADD|EV_ENABLE|EV_ONESHOT, 0, (int)ival, to);
//		to->ival = ival;
//	}
//	return (0);
//}
//void
//AG_DelTimerKQUEUE(AG_Timer *to)
//{
//	AG_EventSourceKQUEUE *kq = (AG_EventSourceKQUEUE *)agEventSource;
//
//	if (GrowKqChangelist(kq, kq->nChanges+1) == -1) {
//		AG_FatalError(NULL);
//	}
//	EV_SET(&kq->changes[kq->nChanges++], to->id, EVFILT_TIMER, EV_DELETE,
//	    0, 0, NULL);
//	agTimerCount--;
//}
//
//#endif /* HAVE_KQUEUE */
//
//#ifdef HAVE_TIMERFD
///*
// * Standard event sink using select(2) and fd-based timers,
// * usually available on Linux.
// */
//int
//AG_EventSinkTIMERFD(void)
//{
//	fd_set rdFds, wrFds;
//	int nFds, rv;
//	AG_EventSink *es;
//	AG_Object *ob, *obNext;
//	AG_Timer *to, *toNext;
//	struct timeval timeo, *pTimeo;
//
//restart:
//	nFds = 0;
//	FD_ZERO(&rdFds);
//	FD_ZERO(&wrFds);
//	TAILQ_FOREACH(es, &agEventSource->sinks, sinks) {
//		switch (es->type) {
//		case AG_SINK_READ:
//			FD_SET(es->ident, &rdFds);
//			if (es->ident > nFds) { nFds = es->ident; }
//			break;
//		case AG_SINK_WRITE:
//			FD_SET(es->ident, &wrFds);
//			if (es->ident > nFds) { nFds = es->ident; }
//			break;
//		}
//	}
//	TAILQ_FOREACH(ob, &agTimerObjQ, tobjs) {
//		TAILQ_FOREACH(to, &ob->timers, timers) {
//			FD_SET(to->id, &rdFds);
//			if (to->id > nFds) { nFds = to->id; }
//		}
//	}
//	if (!TAILQ_EMPTY(&agEventSource->spinners)) {
//		timeo.tv_sec = 0;
//		timeo.tv_usec = 0;
//		pTimeo = &timeo;
//	} else {
//		pTimeo = NULL;
//	}
//	rv = select(nFds+1, &rdFds, &wrFds, NULL, pTimeo);
//	if (rv == -1) {
//		if (errno == EINTR) {
//			goto restart;
//		}
//		AG_SetError("select: %s", AG_Strerror(errno));
//		return (-1);
//	}
//	
//	AG_LockTiming();
//
//	/* 1. Process timer expirations. */
//	for (ob = TAILQ_FIRST(&agTimerObjQ);
//	     ob != TAILQ_END(&agTimerObjQ);
//	     ob = obNext) {
//		obNext = TAILQ_NEXT(ob, tobjs);
//		AG_ObjectLock(ob);
//		for (to = TAILQ_FIRST(&ob->timers);
//		     to != TAILQ_END(&ob->timers);
//		     to = toNext) {
//			struct itimerspec its;
//			Uint32 rvt;
//
//			toNext = TAILQ_NEXT(to, timers);
//			if (!FD_ISSET(to->id, &rdFds)) {
//				continue;
//			}
//			rvt = to->fn(to, &to->fnEvent);
//			if (rvt > 0) {
//				its.it_value.tv_sec = rvt/1000;
//				its.it_value.tv_nsec = (rvt % 1000)*1000000L;
//				its.it_interval.tv_sec = 0;
//				its.it_interval.tv_nsec = 0L;
//				if (timerfd_settime(to->id, 0, &its, NULL) == -1) {
//					Verbose("timerfd_settime: %s", AG_Strerror(errno));
//					FD_CLR(to->id, &rdFds);
//					AG_DelTimer(ob, to);
//				}
//			} else {
//				FD_CLR(to->id, &rdFds);
//				AG_DelTimer(ob, to);
//			}
//		}
//		AG_ObjectUnlock(ob);
//	}
//	
//	/* 2. Process I/O events. */
//	TAILQ_FOREACH(es, &agEventSource->sinks, sinks) {
//		switch (es->type) {
//		case AG_SINK_READ:
//			if (FD_ISSET(es->ident, &rdFds)) {
//				es->fn(es, &es->fnArgs);
//			}
//			break;
//		case AG_SINK_WRITE:
//			if (FD_ISSET(es->ident, &wrFds)) {
//				es->fn(es, &es->fnArgs);
//			}
//			break;
//		}
//	}
//
//	AG_UnlockTiming();
//	return (0);
//}
//
///*
// * Add/remove a fd-based timer.
// */
//int
//AG_AddTimerTIMERFD(AG_Timer *to, Uint32 ival, int newTimer)
//{
//	struct itimerspec its;
//
//	if (newTimer) {
//		/* Create a timerfd. Store the file descriptor as ID. */
//		if ((to->id = timerfd_create(CLOCK_MONOTONIC, 0)) == -1) {
//			AG_SetError("timerfd_create: %s", AG_Strerror(errno));
//			return (-1);
//		}
//	}
//	its.it_value.tv_sec = ival/1000;
//	its.it_value.tv_nsec = (ival % 1000)*1000000L;
//	its.it_interval.tv_sec = 0;
//	its.it_interval.tv_nsec = 0L;
//	if (timerfd_settime(to->id, 0, &its, NULL) == -1) {
//		close(to->id);
//		AG_SetError("timerfd_settime: %s", AG_Strerror(errno));
//		return (-1);
//	}
//	to->ival = ival;
//	return (0);
//}
//void
//AG_DelTimerTIMERFD(AG_Timer *to)
//{
//#ifdef AG_DEBUG
//	if (to->id == -1) { AG_FatalError("timerfd inconsistency"); }
//#endif
//	close(to->id);
//}
//#endif /* HAVE_TIMERFD */
//
//#if defined(HAVE_SELECT) && !defined(AG_THREADS)
///*
// * Standard event sink using select(2) with timers implemented using the
// * select() timeout. Only available in single-threaded builds, since timers
// * cannot be added, restarted or removed from different threads with this
// * method.
// */
//int
//AG_EventSinkTIMEDSELECT(void)
//{
//	fd_set rdFds, wrFds;
//	int i, nFds, rv;
//	AG_EventSink *es;
//	AG_Object *ob, *obNext;
//	AG_Timer *to, *toNext;
//	struct timeval timeo, *pTimeo;
//	Uint32 t, tSoonest;
//
//restart:
//	nFds = 0;
//	FD_ZERO(&rdFds);
//	FD_ZERO(&wrFds);
//	TAILQ_FOREACH(es, &agEventSource->sinks, sinks) {
//		switch (es->type) {
//		case AG_SINK_READ:
//			FD_SET(es->ident, &rdFds);
//			if (es->ident > nFds) { nFds = es->ident; }
//			break;
//		case AG_SINK_WRITE:
//			FD_SET(es->ident, &wrFds);
//			if (es->ident > nFds) { nFds = es->ident; }
//			break;
//		}
//	}
//
//	if (!TAILQ_EMPTY(&agEventSource->spinners)) {
//		timeo.tv_sec = 0;
//		timeo.tv_usec = 0;
//	} else {
//		AG_LockTiming();
//		t = AG_GetTicks();
//		tSoonest = 0xfffffffe;
//		TAILQ_FOREACH(ob, &agTimerObjQ, tobjs) {
//			TAILQ_FOREACH(to, &ob->timers, timers) {
//				if ((to->tSched - t) < tSoonest)
//					tSoonest = (to->tSched - t);
//			}
//		}
//		timeo.tv_sec = tSoonest/1000;
//		timeo.tv_usec = (tSoonest % 1000)*1000;
//		AG_UnlockTiming();
//	}
//	rv = select(nFds+1, &rdFds, &wrFds, NULL, &timeo);
//	if (rv == -1) {
//		if (errno == EINTR) {
//			goto restart;
//		}
//		AG_SetError("select: %s", AG_Strerror(errno));
//		return (-1);
//	}
//	
//	AG_LockTiming();
//	/* 1. Process timer expirations. */
//	AG_ProcessTimeouts(t);
//	if (rv > 0) {
//		/* 2. Process I/O events */
//		TAILQ_FOREACH(es, &agEventSource->sinks, sinks) {
//			switch (es->type) {
//			case AG_SINK_READ:
//				if (FD_ISSET(es->ident, &rdFds)) {
//					es->fn(es, &es->fnArgs);
//				}
//				break;
//			case AG_SINK_WRITE:
//				if (FD_ISSET(es->ident, &wrFds)) {
//					es->fn(es, &es->fnArgs);
//				}
//				break;
//			}
//		}
//	}
//	AG_UnlockTiming();
//	return (0);
//}
//#endif /* HAVE_SELECT and !AG_THREADS */
//
//#if defined(HAVE_SELECT) && defined(AG_THREADS)
///*
// * Standard event sink using non-blocking select(2) with timers implemented
// * with a delay loop. This is inefficient, but on some platforms, it is the
// * only thread-safe option.
// */
//int
//AG_EventSinkSELECT(void)
//{
//	fd_set rdFds, wrFds;
//	int nFds, rv;
//	AG_EventSink *es;
//	struct timeval timeo;
//
//	nFds = 0;
//	FD_ZERO(&rdFds);
//	FD_ZERO(&wrFds);
//	
//	TAILQ_FOREACH(es, &agEventSource->sinks, sinks) {
//		switch (es->type) {
//		case AG_SINK_READ:
//			FD_SET(es->ident, &rdFds);
//			if (es->ident > nFds) { nFds = es->ident; }
//			break;
//		case AG_SINK_WRITE:
//			FD_SET(es->ident, &wrFds);
//			if (es->ident > nFds) { nFds = es->ident; }
//			break;
//		}
//	}
//
//restart:
//	timeo.tv_sec = 0;
//	timeo.tv_usec = 0;
//	rv = select(nFds+1, &rdFds, &wrFds, NULL, &timeo);
//	if (rv == -1) {
//		if (errno == EINTR) {
//			goto restart;
//		}
//		AG_SetError("select: %s", AG_Strerror(errno));
//		return (-1);
//	}
//	
//	AG_LockTiming();
//	/* 1. Process timer expirations. */
//	AG_ProcessTimeouts(AG_GetTicks());
//	if (rv > 0) {
//		/* 2. Process I/O events. */
//		TAILQ_FOREACH(es, &agEventSource->sinks, sinks) {
//			switch (es->type) {
//			case AG_SINK_READ:
//				if (FD_ISSET(es->ident, &rdFds)) {
//					es->fn(es, &es->fnArgs);
//				}
//				break;
//			case AG_SINK_WRITE:
//				if (FD_ISSET(es->ident, &wrFds)) {
//					es->fn(es, &es->fnArgs);
//				}
//				break;
//			}
//		}
//	}
//	AG_UnlockTiming();
//	
//	if (TAILQ_EMPTY(&agEventSource->spinners)) {
//		AG_Delay(1);
//	}
//	return (0);
//}
//#endif /* HAVE_SELECT and AG_THREADS */
//
///*
// * Fallback "spinning" event sink using a delay loop. This is inefficient,
// * but is the only option on some platforms.
// */
//int
//AG_EventSinkSPINNER(void)
//{
//	AG_ProcessTimeouts(AG_GetTicks());
//	AG_Delay(1);
//	return (0);
//}
//
///*
// * Standard event loop routine. We loop over the event source and invoke
// * the related event sinks. AG_EventLoop() may be used outside of the
// * main thread (in which case, a new event source will be created).
// */
//int
//AG_EventLoop(void)
//{
func AG_EventLoop() int {
//	AG_EventSource *src = AG_GetEventSource();
	src := aG_GetEventSource()
//	AG_EventSink *es;
//	
//	TAILQ_FOREACH(es, &src->prologues, sinks) {
//		es->fn(es, &es->fnArgs);
//	}
//	for (;;) {
//		TAILQ_FOREACH(es, &src->spinners, sinks) {
//			es->fn(es, &es->fnArgs);
//		}
//		if (src->sinkFn() == -1) {
//			return (1);
//		}
//		TAILQ_FOREACH(es, &src->epilogues, sinks) {
//			es->fn(es, &es->fnArgs);
//		}
//		if (src->breakReq)
//			break;
//	}
//	return (src->returnCode);
	return src.returnCode
//}
}
//
///* Request that we break out of AG_EventLoop(). */
//void
//AG_Terminate(int retCode)
//{
//	AG_EventSource *src = AG_GetEventSource();
//
//	src->breakReq = 1;
//	src->returnCode = retCode;
//}
//void
//AG_TerminateEv(AG_Event *ev)
//{
//	AG_EventSource *src = AG_GetEventSource();
//
//	if (ev->argc > 1 &&
//	    ev->argv[1].type == AG_VARIABLE_INT) {
//		src->returnCode = ev->argv[1].data.i;
//	} else {
//		src->returnCode = 0;
//	}
//	src->breakReq = 1;
//}
