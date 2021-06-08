/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/quaternion.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Quaternion", QuaternionTypeSupport)
}

// Do not create instances of this type directly. Always use NewQuaternion
// function instead.
type Quaternion struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
	W float64 `yaml:"w"`
}

// NewQuaternion creates a new Quaternion with default values.
func NewQuaternion() *Quaternion {
	self := Quaternion{}
	self.SetDefaults()
	return &self
}

func (t *Quaternion) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Quaternion) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.W = 1
	
}

// Modifying this variable is undefined behavior.
var QuaternionTypeSupport types.MessageTypeSupport = _QuaternionTypeSupport{}

type _QuaternionTypeSupport struct{}

func (t _QuaternionTypeSupport) New() types.Message {
	return NewQuaternion()
}

func (t _QuaternionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Quaternion
	return (unsafe.Pointer)(C.geometry_msgs__msg__Quaternion__create())
}

func (t _QuaternionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Quaternion__destroy((*C.geometry_msgs__msg__Quaternion)(pointer_to_free))
}

func (t _QuaternionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Quaternion)
	mem := (*C.geometry_msgs__msg__Quaternion)(dst)
	mem.x = C.double(m.X)
	mem.y = C.double(m.Y)
	mem.z = C.double(m.Z)
	mem.w = C.double(m.W)
}

func (t _QuaternionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Quaternion)
	mem := (*C.geometry_msgs__msg__Quaternion)(ros2_message_buffer)
	m.X = float64(mem.x)
	m.Y = float64(mem.y)
	m.Z = float64(mem.z)
	m.W = float64(mem.w)
}

func (t _QuaternionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Quaternion())
}

type CQuaternion = C.geometry_msgs__msg__Quaternion
type CQuaternion__Sequence = C.geometry_msgs__msg__Quaternion__Sequence

func Quaternion__Sequence_to_Go(goSlice *[]Quaternion, cSlice CQuaternion__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Quaternion, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Quaternion__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Quaternion * uintptr(i)),
		))
		QuaternionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Quaternion__Sequence_to_C(cSlice *CQuaternion__Sequence, goSlice []Quaternion) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Quaternion)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Quaternion * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Quaternion)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Quaternion * uintptr(i)),
		))
		QuaternionTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Quaternion__Array_to_Go(goSlice []Quaternion, cSlice []CQuaternion) {
	for i := 0; i < len(cSlice); i++ {
		QuaternionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Quaternion__Array_to_C(cSlice []CQuaternion, goSlice []Quaternion) {
	for i := 0; i < len(goSlice); i++ {
		QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
