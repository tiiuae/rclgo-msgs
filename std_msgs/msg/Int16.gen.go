/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/int16.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Int16", Int16TypeSupport)
}

// Do not create instances of this type directly. Always use NewInt16
// function instead.
type Int16 struct {
	Data int16 `yaml:"data"`
}

// NewInt16 creates a new Int16 with default values.
func NewInt16() *Int16 {
	self := Int16{}
	self.SetDefaults()
	return &self
}

func (t *Int16) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Int16) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var Int16TypeSupport types.MessageTypeSupport = _Int16TypeSupport{}

type _Int16TypeSupport struct{}

func (t _Int16TypeSupport) New() types.Message {
	return NewInt16()
}

func (t _Int16TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Int16
	return (unsafe.Pointer)(C.std_msgs__msg__Int16__create())
}

func (t _Int16TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Int16__destroy((*C.std_msgs__msg__Int16)(pointer_to_free))
}

func (t _Int16TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int16)
	mem := (*C.std_msgs__msg__Int16)(dst)
	mem.data = C.int16_t(m.Data)
}

func (t _Int16TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int16)
	mem := (*C.std_msgs__msg__Int16)(ros2_message_buffer)
	m.Data = int16(mem.data)
}

func (t _Int16TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Int16())
}

type CInt16 = C.std_msgs__msg__Int16
type CInt16__Sequence = C.std_msgs__msg__Int16__Sequence

func Int16__Sequence_to_Go(goSlice *[]Int16, cSlice CInt16__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int16, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Int16__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int16 * uintptr(i)),
		))
		Int16TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int16__Sequence_to_C(cSlice *CInt16__Sequence, goSlice []Int16) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Int16)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Int16 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Int16)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int16 * uintptr(i)),
		))
		Int16TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int16__Array_to_Go(goSlice []Int16, cSlice []CInt16) {
	for i := 0; i < len(cSlice); i++ {
		Int16TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int16__Array_to_C(cSlice []CInt16, goSlice []Int16) {
	for i := 0; i < len(goSlice); i++ {
		Int16TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}