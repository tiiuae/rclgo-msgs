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
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/float32.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Float32", Float32TypeSupport)
}

// Do not create instances of this type directly. Always use NewFloat32
// function instead.
type Float32 struct {
	Data float32 `yaml:"data"`
}

// NewFloat32 creates a new Float32 with default values.
func NewFloat32() *Float32 {
	self := Float32{}
	self.SetDefaults()
	return &self
}

func (t *Float32) Clone() *Float32 {
	c := &Float32{}
	c.Data = t.Data
	return c
}

func (t *Float32) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Float32) SetDefaults() {
	t.Data = 0
}

// CloneFloat32Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFloat32Slice(dst, src []Float32) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Float32TypeSupport types.MessageTypeSupport = _Float32TypeSupport{}

type _Float32TypeSupport struct{}

func (t _Float32TypeSupport) New() types.Message {
	return NewFloat32()
}

func (t _Float32TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Float32
	return (unsafe.Pointer)(C.std_msgs__msg__Float32__create())
}

func (t _Float32TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Float32__destroy((*C.std_msgs__msg__Float32)(pointer_to_free))
}

func (t _Float32TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Float32)
	mem := (*C.std_msgs__msg__Float32)(dst)
	mem.data = C.float(m.Data)
}

func (t _Float32TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Float32)
	mem := (*C.std_msgs__msg__Float32)(ros2_message_buffer)
	m.Data = float32(mem.data)
}

func (t _Float32TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Float32())
}

type CFloat32 = C.std_msgs__msg__Float32
type CFloat32__Sequence = C.std_msgs__msg__Float32__Sequence

func Float32__Sequence_to_Go(goSlice *[]Float32, cSlice CFloat32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Float32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Float32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Float32 * uintptr(i)),
		))
		Float32TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Float32__Sequence_to_C(cSlice *CFloat32__Sequence, goSlice []Float32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Float32)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Float32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Float32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Float32 * uintptr(i)),
		))
		Float32TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Float32__Array_to_Go(goSlice []Float32, cSlice []CFloat32) {
	for i := 0; i < len(cSlice); i++ {
		Float32TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Float32__Array_to_C(cSlice []CFloat32, goSlice []Float32) {
	for i := 0; i < len(goSlice); i++ {
		Float32TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
