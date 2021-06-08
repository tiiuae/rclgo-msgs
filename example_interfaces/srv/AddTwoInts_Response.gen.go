/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/add_two_ints.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/AddTwoInts_Response", AddTwoInts_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewAddTwoInts_Response
// function instead.
type AddTwoInts_Response struct {
	Sum int64 `yaml:"sum"`
}

// NewAddTwoInts_Response creates a new AddTwoInts_Response with default values.
func NewAddTwoInts_Response() *AddTwoInts_Response {
	self := AddTwoInts_Response{}
	self.SetDefaults()
	return &self
}

func (t *AddTwoInts_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *AddTwoInts_Response) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var AddTwoInts_ResponseTypeSupport types.MessageTypeSupport = _AddTwoInts_ResponseTypeSupport{}

type _AddTwoInts_ResponseTypeSupport struct{}

func (t _AddTwoInts_ResponseTypeSupport) New() types.Message {
	return NewAddTwoInts_Response()
}

func (t _AddTwoInts_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__AddTwoInts_Response
	return (unsafe.Pointer)(C.example_interfaces__srv__AddTwoInts_Response__create())
}

func (t _AddTwoInts_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__AddTwoInts_Response__destroy((*C.example_interfaces__srv__AddTwoInts_Response)(pointer_to_free))
}

func (t _AddTwoInts_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AddTwoInts_Response)
	mem := (*C.example_interfaces__srv__AddTwoInts_Response)(dst)
	mem.sum = C.int64_t(m.Sum)
}

func (t _AddTwoInts_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AddTwoInts_Response)
	mem := (*C.example_interfaces__srv__AddTwoInts_Response)(ros2_message_buffer)
	m.Sum = int64(mem.sum)
}

func (t _AddTwoInts_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__AddTwoInts_Response())
}

type CAddTwoInts_Response = C.example_interfaces__srv__AddTwoInts_Response
type CAddTwoInts_Response__Sequence = C.example_interfaces__srv__AddTwoInts_Response__Sequence

func AddTwoInts_Response__Sequence_to_Go(goSlice *[]AddTwoInts_Response, cSlice CAddTwoInts_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AddTwoInts_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__srv__AddTwoInts_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__AddTwoInts_Response * uintptr(i)),
		))
		AddTwoInts_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func AddTwoInts_Response__Sequence_to_C(cSlice *CAddTwoInts_Response__Sequence, goSlice []AddTwoInts_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__srv__AddTwoInts_Response)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__srv__AddTwoInts_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__srv__AddTwoInts_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__AddTwoInts_Response * uintptr(i)),
		))
		AddTwoInts_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func AddTwoInts_Response__Array_to_Go(goSlice []AddTwoInts_Response, cSlice []CAddTwoInts_Response) {
	for i := 0; i < len(cSlice); i++ {
		AddTwoInts_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AddTwoInts_Response__Array_to_C(cSlice []CAddTwoInts_Response, goSlice []AddTwoInts_Response) {
	for i := 0; i < len(goSlice); i++ {
		AddTwoInts_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
