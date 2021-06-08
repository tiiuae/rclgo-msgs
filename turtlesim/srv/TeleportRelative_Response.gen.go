/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package turtlesim_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/teleport_relative.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/TeleportRelative_Response", TeleportRelative_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewTeleportRelative_Response
// function instead.
type TeleportRelative_Response struct {
}

// NewTeleportRelative_Response creates a new TeleportRelative_Response with default values.
func NewTeleportRelative_Response() *TeleportRelative_Response {
	self := TeleportRelative_Response{}
	self.SetDefaults()
	return &self
}

func (t *TeleportRelative_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *TeleportRelative_Response) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var TeleportRelative_ResponseTypeSupport types.MessageTypeSupport = _TeleportRelative_ResponseTypeSupport{}

type _TeleportRelative_ResponseTypeSupport struct{}

func (t _TeleportRelative_ResponseTypeSupport) New() types.Message {
	return NewTeleportRelative_Response()
}

func (t _TeleportRelative_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__TeleportRelative_Response
	return (unsafe.Pointer)(C.turtlesim__srv__TeleportRelative_Response__create())
}

func (t _TeleportRelative_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__TeleportRelative_Response__destroy((*C.turtlesim__srv__TeleportRelative_Response)(pointer_to_free))
}

func (t _TeleportRelative_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _TeleportRelative_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _TeleportRelative_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__TeleportRelative_Response())
}

type CTeleportRelative_Response = C.turtlesim__srv__TeleportRelative_Response
type CTeleportRelative_Response__Sequence = C.turtlesim__srv__TeleportRelative_Response__Sequence

func TeleportRelative_Response__Sequence_to_Go(goSlice *[]TeleportRelative_Response, cSlice CTeleportRelative_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TeleportRelative_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__TeleportRelative_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__TeleportRelative_Response * uintptr(i)),
		))
		TeleportRelative_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TeleportRelative_Response__Sequence_to_C(cSlice *CTeleportRelative_Response__Sequence, goSlice []TeleportRelative_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__TeleportRelative_Response)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__TeleportRelative_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__TeleportRelative_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__TeleportRelative_Response * uintptr(i)),
		))
		TeleportRelative_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TeleportRelative_Response__Array_to_Go(goSlice []TeleportRelative_Response, cSlice []CTeleportRelative_Response) {
	for i := 0; i < len(cSlice); i++ {
		TeleportRelative_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TeleportRelative_Response__Array_to_C(cSlice []CTeleportRelative_Response, goSlice []TeleportRelative_Response) {
	for i := 0; i < len(goSlice); i++ {
		TeleportRelative_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
