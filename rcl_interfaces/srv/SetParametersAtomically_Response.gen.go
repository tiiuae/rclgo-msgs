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
	typemap.RegisterMessage("rcl_interfaces/SetParametersAtomically_Response", SetParametersAtomically_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetParametersAtomically_Response
// function instead.
type SetParametersAtomically_Response struct {
	Result rcl_interfaces_msg.SetParametersResult `yaml:"result"`// Indicates whether setting all of the parameters succeeded or not and why.
}

// NewSetParametersAtomically_Response creates a new SetParametersAtomically_Response with default values.
func NewSetParametersAtomically_Response() *SetParametersAtomically_Response {
	self := SetParametersAtomically_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetParametersAtomically_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *SetParametersAtomically_Response) SetDefaults() {
	t.Result.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var SetParametersAtomically_ResponseTypeSupport types.MessageTypeSupport = _SetParametersAtomically_ResponseTypeSupport{}

type _SetParametersAtomically_ResponseTypeSupport struct{}

func (t _SetParametersAtomically_ResponseTypeSupport) New() types.Message {
	return NewSetParametersAtomically_Response()
}

func (t _SetParametersAtomically_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__SetParametersAtomically_Response
	return (unsafe.Pointer)(C.rcl_interfaces__srv__SetParametersAtomically_Response__create())
}

func (t _SetParametersAtomically_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__SetParametersAtomically_Response__destroy((*C.rcl_interfaces__srv__SetParametersAtomically_Response)(pointer_to_free))
}

func (t _SetParametersAtomically_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParametersAtomically_Response)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(dst)
	rcl_interfaces_msg.SetParametersResultTypeSupport.AsCStruct(unsafe.Pointer(&mem.result), &m.Result)
}

func (t _SetParametersAtomically_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParametersAtomically_Response)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(ros2_message_buffer)
	rcl_interfaces_msg.SetParametersResultTypeSupport.AsGoStruct(&m.Result, unsafe.Pointer(&mem.result))
}

func (t _SetParametersAtomically_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__SetParametersAtomically_Response())
}

type CSetParametersAtomically_Response = C.rcl_interfaces__srv__SetParametersAtomically_Response
type CSetParametersAtomically_Response__Sequence = C.rcl_interfaces__srv__SetParametersAtomically_Response__Sequence

func SetParametersAtomically_Response__Sequence_to_Go(goSlice *[]SetParametersAtomically_Response, cSlice CSetParametersAtomically_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParametersAtomically_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__srv__SetParametersAtomically_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Response * uintptr(i)),
		))
		SetParametersAtomically_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetParametersAtomically_Response__Sequence_to_C(cSlice *CSetParametersAtomically_Response__Sequence, goSlice []SetParametersAtomically_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Response * uintptr(i)),
		))
		SetParametersAtomically_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetParametersAtomically_Response__Array_to_Go(goSlice []SetParametersAtomically_Response, cSlice []CSetParametersAtomically_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetParametersAtomically_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParametersAtomically_Response__Array_to_C(cSlice []CSetParametersAtomically_Response, goSlice []SetParametersAtomically_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetParametersAtomically_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
