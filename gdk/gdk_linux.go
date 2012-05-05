package gdk
/*
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
// #cgo pkg-config: gdk-2.0 gthread-2.0
*/
import (
	"C"
	"unsafe"
)

func (v *GdkWindow) GetNativeWindowID() int32 {
	return int32(C.gdk_x11_drawable_get_xid((*C.GdkDrawable)(unsafe.Pointer(v.Window))))
}

