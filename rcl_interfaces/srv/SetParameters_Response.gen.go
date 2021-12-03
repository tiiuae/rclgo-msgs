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
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/set_parameters.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/SetParameters_Response", SetParameters_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetParameters_Response
// function instead.
type SetParameters_Response struct {
	Results []rcl_interfaces_msg.SetParametersResult `yaml:"results"`// Indicates whether setting each parameter succeeded or not and why.
}

// NewSetParameters_Response creates a new SetParameters_Response with default values.
func NewSetParameters_Response() *SetParameters_Response {
	self := SetParameters_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetParameters_Response) Clone() *SetParameters_Response {
	c := &SetParameters_Response{}
	if t.Results != nil {
		c.Results = make([]rcl_interfaces_msg.SetParametersResult, len(t.Results))
		rcl_interfaces_msg.CloneSetParametersResultSlice(c.Results, t.Results)
	}
	return c
}

func (t *SetParameters_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetParameters_Response) SetDefaults() {
	t.Results = nil
}

// CloneSetParameters_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetParameters_ResponseSlice(dst, src []SetParameters_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetParameters_ResponseTypeSupport types.MessageTypeSupport = _SetParameters_ResponseTypeSupport{}

type _SetParameters_ResponseTypeSupport struct{}

func (t _SetParameters_ResponseTypeSupport) New() types.Message {
	return NewSetParameters_Response()
}

func (t _SetParameters_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__SetParameters_Response
	return (unsafe.Pointer)(C.rcl_interfaces__srv__SetParameters_Response__create())
}

func (t _SetParameters_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__SetParameters_Response__destroy((*C.rcl_interfaces__srv__SetParameters_Response)(pointer_to_free))
}

func (t _SetParameters_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParameters_Response)
	mem := (*C.rcl_interfaces__srv__SetParameters_Response)(dst)
	rcl_interfaces_msg.SetParametersResult__Sequence_to_C((*rcl_interfaces_msg.CSetParametersResult__Sequence)(unsafe.Pointer(&mem.results)), m.Results)
}

func (t _SetParameters_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParameters_Response)
	mem := (*C.rcl_interfaces__srv__SetParameters_Response)(ros2_message_buffer)
	rcl_interfaces_msg.SetParametersResult__Sequence_to_Go(&m.Results, *(*rcl_interfaces_msg.CSetParametersResult__Sequence)(unsafe.Pointer(&mem.results)))
}

func (t _SetParameters_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__SetParameters_Response())
}

type CSetParameters_Response = C.rcl_interfaces__srv__SetParameters_Response
type CSetParameters_Response__Sequence = C.rcl_interfaces__srv__SetParameters_Response__Sequence

func SetParameters_Response__Sequence_to_Go(goSlice *[]SetParameters_Response, cSlice CSetParameters_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParameters_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__srv__SetParameters_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__SetParameters_Response * uintptr(i)),
		))
		SetParameters_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetParameters_Response__Sequence_to_C(cSlice *CSetParameters_Response__Sequence, goSlice []SetParameters_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__SetParameters_Response)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__srv__SetParameters_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__srv__SetParameters_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__SetParameters_Response * uintptr(i)),
		))
		SetParameters_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetParameters_Response__Array_to_Go(goSlice []SetParameters_Response, cSlice []CSetParameters_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetParameters_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParameters_Response__Array_to_C(cSlice []CSetParameters_Response, goSlice []SetParameters_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetParameters_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
