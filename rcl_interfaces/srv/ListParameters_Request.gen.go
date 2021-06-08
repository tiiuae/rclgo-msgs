/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/list_parameters.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/ListParameters_Request", ListParameters_RequestTypeSupport)
}
const (
	ListParameters_Request_DEPTH_RECURSIVE uint64 = 0// Recursively get parameters with unlimited depth.
)

// Do not create instances of this type directly. Always use NewListParameters_Request
// function instead.
type ListParameters_Request struct {
	Prefixes []string `yaml:"prefixes"`// The list of parameter prefixes to query.
	Depth uint64 `yaml:"depth"`// Relative depth from given prefixes to return.Use DEPTH_RECURSIVE to get the recursive parameters and prefixes for each prefix.
}

// NewListParameters_Request creates a new ListParameters_Request with default values.
func NewListParameters_Request() *ListParameters_Request {
	self := ListParameters_Request{}
	self.SetDefaults()
	return &self
}

func (t *ListParameters_Request) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *ListParameters_Request) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var ListParameters_RequestTypeSupport types.MessageTypeSupport = _ListParameters_RequestTypeSupport{}

type _ListParameters_RequestTypeSupport struct{}

func (t _ListParameters_RequestTypeSupport) New() types.Message {
	return NewListParameters_Request()
}

func (t _ListParameters_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__ListParameters_Request
	return (unsafe.Pointer)(C.rcl_interfaces__srv__ListParameters_Request__create())
}

func (t _ListParameters_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__ListParameters_Request__destroy((*C.rcl_interfaces__srv__ListParameters_Request)(pointer_to_free))
}

func (t _ListParameters_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ListParameters_Request)
	mem := (*C.rcl_interfaces__srv__ListParameters_Request)(dst)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.prefixes)), m.Prefixes)
	mem.depth = C.uint64_t(m.Depth)
}

func (t _ListParameters_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ListParameters_Request)
	mem := (*C.rcl_interfaces__srv__ListParameters_Request)(ros2_message_buffer)
	primitives.String__Sequence_to_Go(&m.Prefixes, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.prefixes)))
	m.Depth = uint64(mem.depth)
}

func (t _ListParameters_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__ListParameters_Request())
}

type CListParameters_Request = C.rcl_interfaces__srv__ListParameters_Request
type CListParameters_Request__Sequence = C.rcl_interfaces__srv__ListParameters_Request__Sequence

func ListParameters_Request__Sequence_to_Go(goSlice *[]ListParameters_Request, cSlice CListParameters_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ListParameters_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__srv__ListParameters_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__ListParameters_Request * uintptr(i)),
		))
		ListParameters_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ListParameters_Request__Sequence_to_C(cSlice *CListParameters_Request__Sequence, goSlice []ListParameters_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__ListParameters_Request)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__srv__ListParameters_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__srv__ListParameters_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__ListParameters_Request * uintptr(i)),
		))
		ListParameters_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ListParameters_Request__Array_to_Go(goSlice []ListParameters_Request, cSlice []CListParameters_Request) {
	for i := 0; i < len(cSlice); i++ {
		ListParameters_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ListParameters_Request__Array_to_C(cSlice []CListParameters_Request, goSlice []ListParameters_Request) {
	for i := 0; i < len(goSlice); i++ {
		ListParameters_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
