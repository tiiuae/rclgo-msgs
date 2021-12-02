/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/srv/set_map.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("nav_msgs/SetMap_Response", SetMap_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetMap_Response
// function instead.
type SetMap_Response struct {
	Success bool `yaml:"success"`// True if the map was successfully set, false otherwise.
}

// NewSetMap_Response creates a new SetMap_Response with default values.
func NewSetMap_Response() *SetMap_Response {
	self := SetMap_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetMap_Response) Clone() *SetMap_Response {
	c := &SetMap_Response{}
	c.Success = t.Success
	return c
}

func (t *SetMap_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetMap_Response) SetDefaults() {
	t.Success = false
}

// CloneSetMap_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetMap_ResponseSlice(dst, src []SetMap_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetMap_ResponseTypeSupport types.MessageTypeSupport = _SetMap_ResponseTypeSupport{}

type _SetMap_ResponseTypeSupport struct{}

func (t _SetMap_ResponseTypeSupport) New() types.Message {
	return NewSetMap_Response()
}

func (t _SetMap_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__srv__SetMap_Response
	return (unsafe.Pointer)(C.nav_msgs__srv__SetMap_Response__create())
}

func (t _SetMap_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__srv__SetMap_Response__destroy((*C.nav_msgs__srv__SetMap_Response)(pointer_to_free))
}

func (t _SetMap_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetMap_Response)
	mem := (*C.nav_msgs__srv__SetMap_Response)(dst)
	mem.success = C.bool(m.Success)
}

func (t _SetMap_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetMap_Response)
	mem := (*C.nav_msgs__srv__SetMap_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
}

func (t _SetMap_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__srv__SetMap_Response())
}

type CSetMap_Response = C.nav_msgs__srv__SetMap_Response
type CSetMap_Response__Sequence = C.nav_msgs__srv__SetMap_Response__Sequence

func SetMap_Response__Sequence_to_Go(goSlice *[]SetMap_Response, cSlice CSetMap_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetMap_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__srv__SetMap_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__SetMap_Response * uintptr(i)),
		))
		SetMap_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetMap_Response__Sequence_to_C(cSlice *CSetMap_Response__Sequence, goSlice []SetMap_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__srv__SetMap_Response)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__srv__SetMap_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__srv__SetMap_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__SetMap_Response * uintptr(i)),
		))
		SetMap_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetMap_Response__Array_to_Go(goSlice []SetMap_Response, cSlice []CSetMap_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetMap_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetMap_Response__Array_to_C(cSlice []CSetMap_Response, goSlice []SetMap_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetMap_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
