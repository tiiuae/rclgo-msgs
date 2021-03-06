/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	nav_msgs_msg "github.com/tiiuae/rclgo-msgs/nav_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/srv/path_to_local.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/PathToLocal_Request", PathToLocal_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewPathToLocal_Request
// function instead.
type PathToLocal_Request struct {
	Path nav_msgs_msg.Path `yaml:"path"`
}

// NewPathToLocal_Request creates a new PathToLocal_Request with default values.
func NewPathToLocal_Request() *PathToLocal_Request {
	self := PathToLocal_Request{}
	self.SetDefaults()
	return &self
}

func (t *PathToLocal_Request) Clone() *PathToLocal_Request {
	c := &PathToLocal_Request{}
	c.Path = *t.Path.Clone()
	return c
}

func (t *PathToLocal_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PathToLocal_Request) SetDefaults() {
	t.Path.SetDefaults()
}

// ClonePathToLocal_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePathToLocal_RequestSlice(dst, src []PathToLocal_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PathToLocal_RequestTypeSupport types.MessageTypeSupport = _PathToLocal_RequestTypeSupport{}

type _PathToLocal_RequestTypeSupport struct{}

func (t _PathToLocal_RequestTypeSupport) New() types.Message {
	return NewPathToLocal_Request()
}

func (t _PathToLocal_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__PathToLocal_Request
	return (unsafe.Pointer)(C.fog_msgs__srv__PathToLocal_Request__create())
}

func (t _PathToLocal_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__PathToLocal_Request__destroy((*C.fog_msgs__srv__PathToLocal_Request)(pointer_to_free))
}

func (t _PathToLocal_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PathToLocal_Request)
	mem := (*C.fog_msgs__srv__PathToLocal_Request)(dst)
	nav_msgs_msg.PathTypeSupport.AsCStruct(unsafe.Pointer(&mem.path), &m.Path)
}

func (t _PathToLocal_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PathToLocal_Request)
	mem := (*C.fog_msgs__srv__PathToLocal_Request)(ros2_message_buffer)
	nav_msgs_msg.PathTypeSupport.AsGoStruct(&m.Path, unsafe.Pointer(&mem.path))
}

func (t _PathToLocal_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__PathToLocal_Request())
}

type CPathToLocal_Request = C.fog_msgs__srv__PathToLocal_Request
type CPathToLocal_Request__Sequence = C.fog_msgs__srv__PathToLocal_Request__Sequence

func PathToLocal_Request__Sequence_to_Go(goSlice *[]PathToLocal_Request, cSlice CPathToLocal_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PathToLocal_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__PathToLocal_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__PathToLocal_Request * uintptr(i)),
		))
		PathToLocal_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PathToLocal_Request__Sequence_to_C(cSlice *CPathToLocal_Request__Sequence, goSlice []PathToLocal_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__PathToLocal_Request)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__PathToLocal_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__PathToLocal_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__PathToLocal_Request * uintptr(i)),
		))
		PathToLocal_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PathToLocal_Request__Array_to_Go(goSlice []PathToLocal_Request, cSlice []CPathToLocal_Request) {
	for i := 0; i < len(cSlice); i++ {
		PathToLocal_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PathToLocal_Request__Array_to_C(cSlice []CPathToLocal_Request, goSlice []PathToLocal_Request) {
	for i := 0; i < len(goSlice); i++ {
		PathToLocal_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
