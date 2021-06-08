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
	rcl_interfaces_msg "github.com/tiiuae/rclgo-msgs/rcl_interfaces/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/set_parameters_atomically.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/SetParametersAtomically_Request", SetParametersAtomically_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetParametersAtomically_Request
// function instead.
type SetParametersAtomically_Request struct {
	Parameters []rcl_interfaces_msg.Parameter `yaml:"parameters"`// A list of parameters to set atomically.This call will either set all values, or none of the values.
}

// NewSetParametersAtomically_Request creates a new SetParametersAtomically_Request with default values.
func NewSetParametersAtomically_Request() *SetParametersAtomically_Request {
	self := SetParametersAtomically_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetParametersAtomically_Request) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *SetParametersAtomically_Request) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var SetParametersAtomically_RequestTypeSupport types.MessageTypeSupport = _SetParametersAtomically_RequestTypeSupport{}

type _SetParametersAtomically_RequestTypeSupport struct{}

func (t _SetParametersAtomically_RequestTypeSupport) New() types.Message {
	return NewSetParametersAtomically_Request()
}

func (t _SetParametersAtomically_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__SetParametersAtomically_Request
	return (unsafe.Pointer)(C.rcl_interfaces__srv__SetParametersAtomically_Request__create())
}

func (t _SetParametersAtomically_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__SetParametersAtomically_Request__destroy((*C.rcl_interfaces__srv__SetParametersAtomically_Request)(pointer_to_free))
}

func (t _SetParametersAtomically_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParametersAtomically_Request)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(dst)
	rcl_interfaces_msg.Parameter__Sequence_to_C((*rcl_interfaces_msg.CParameter__Sequence)(unsafe.Pointer(&mem.parameters)), m.Parameters)
}

func (t _SetParametersAtomically_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParametersAtomically_Request)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(ros2_message_buffer)
	rcl_interfaces_msg.Parameter__Sequence_to_Go(&m.Parameters, *(*rcl_interfaces_msg.CParameter__Sequence)(unsafe.Pointer(&mem.parameters)))
}

func (t _SetParametersAtomically_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__SetParametersAtomically_Request())
}

type CSetParametersAtomically_Request = C.rcl_interfaces__srv__SetParametersAtomically_Request
type CSetParametersAtomically_Request__Sequence = C.rcl_interfaces__srv__SetParametersAtomically_Request__Sequence

func SetParametersAtomically_Request__Sequence_to_Go(goSlice *[]SetParametersAtomically_Request, cSlice CSetParametersAtomically_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParametersAtomically_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__srv__SetParametersAtomically_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Request * uintptr(i)),
		))
		SetParametersAtomically_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetParametersAtomically_Request__Sequence_to_C(cSlice *CSetParametersAtomically_Request__Sequence, goSlice []SetParametersAtomically_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Request * uintptr(i)),
		))
		SetParametersAtomically_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetParametersAtomically_Request__Array_to_Go(goSlice []SetParametersAtomically_Request, cSlice []CSetParametersAtomically_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetParametersAtomically_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParametersAtomically_Request__Array_to_C(cSlice []CSetParametersAtomically_Request, goSlice []SetParametersAtomically_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetParametersAtomically_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
