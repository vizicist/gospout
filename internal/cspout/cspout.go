package cspout

// #include <stdlib.h>
// #include <stdbool.h>
// #include <cspout_go.h>
// #cgo CFLAGS: -I/users/tjt/documents/github/gospout/internal/cspout
// #cgo LDFLAGS: /users/tjt/documents/github/gospout/internal/cspout/libspout.a
// #cgo LDFLAGS: -lstdc++
// #cgo LDFLAGS: -lopengl32
// #cgo LDFLAGS: -ldxgi
// #cgo LDFLAGS: -lshlwapi
// #cgo LDFLAGS: -lgdi32
// #cgo LDFLAGS: -lD3D11
// #cgo LDFLAGS: -lD3D9
// #cgo LDFLAGS: -lVersion
import "C"
import (
        "unsafe"
)

// Sender sends a testure to a spout receiver
type Sender struct {
        sender C.GoSpoutSender
}

// CreateSender returns a SpoutSender
func CreateSender(name string, width int, height int) Sender {
        var ret Sender
        var cname = C.CString(name)
        ret.sender = C.GoCreateSender(cname, C.int(width), C.int(height))
        C.free(unsafe.Pointer(cname))
        return ret
}

// SendTexture sends a texture
func SendTexture(s Sender, texture uint32, width int, height int) bool {
        cb := C.GoSendTexture(s.sender, C.uint(texture), C.int(width), C.int(height))
        var b bool = bool(cb)
        return b
}

// CreateReceiver creates a receiver
func CreateReceiver(sendername string, width *int, height *int, bUseActive bool) bool {
        var cname = C.CString(sendername)
        var w uint
        var h uint
        wp := (*C.uint)(unsafe.Pointer(&w))
        hp := (*C.uint)(unsafe.Pointer(&h))
        cb := C.GoCreateReceiver(cname, wp, hp, C.bool(bUseActive));
        C.free(unsafe.Pointer(cname))
        *width = int(w)
        *height = int(h)
        var b bool = bool(cb)
        return b
}

// ReceiveTexture creates a receiver
func ReceiveTexture(sendername string, width *int, height *int, textureID int, textureTarget int, bInvert bool, hostFBO int) bool {
        var cname = C.CString(sendername)
        var w uint = uint(*width)
        var h uint = uint(*height)
        wp := (*C.uint)(unsafe.Pointer(&w))
        hp := (*C.uint)(unsafe.Pointer(&h))
        cb := C.GoReceiveTexture(cname, wp, hp, C.int(textureID), C.int(textureTarget), C.bool(bInvert), C.int(hostFBO));
        *width = int(w)
        *height = int(h)
        C.free(unsafe.Pointer(cname))
        var b bool = bool(cb)
        return b
}