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

#include <geometry_msgs/msg/wrench.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Wrench", WrenchTypeSupport)
}

// Do not create instances of this type directly. Always use NewWrench
// function instead.
type Wrench struct {
	Force Vector3 `yaml:"force"`
	Torque Vector3 `yaml:"torque"`
}

// NewWrench creates a new Wrench with default values.
func NewWrench() *Wrench {
	self := Wrench{}
	self.SetDefaults()
	return &self
}

func (t *Wrench) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Wrench) SetDefaults() {
	t.Force.SetDefaults()
	t.Torque.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var WrenchTypeSupport types.MessageTypeSupport = _WrenchTypeSupport{}

type _WrenchTypeSupport struct{}

func (t _WrenchTypeSupport) New() types.Message {
	return NewWrench()
}

func (t _WrenchTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Wrench
	return (unsafe.Pointer)(C.geometry_msgs__msg__Wrench__create())
}

func (t _WrenchTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Wrench__destroy((*C.geometry_msgs__msg__Wrench)(pointer_to_free))
}

func (t _WrenchTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Wrench)
	mem := (*C.geometry_msgs__msg__Wrench)(dst)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.force), &m.Force)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.torque), &m.Torque)
}

func (t _WrenchTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Wrench)
	mem := (*C.geometry_msgs__msg__Wrench)(ros2_message_buffer)
	Vector3TypeSupport.AsGoStruct(&m.Force, unsafe.Pointer(&mem.force))
	Vector3TypeSupport.AsGoStruct(&m.Torque, unsafe.Pointer(&mem.torque))
}

func (t _WrenchTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Wrench())
}

type CWrench = C.geometry_msgs__msg__Wrench
type CWrench__Sequence = C.geometry_msgs__msg__Wrench__Sequence

func Wrench__Sequence_to_Go(goSlice *[]Wrench, cSlice CWrench__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Wrench, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Wrench__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Wrench * uintptr(i)),
		))
		WrenchTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Wrench__Sequence_to_C(cSlice *CWrench__Sequence, goSlice []Wrench) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Wrench)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Wrench * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Wrench)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Wrench * uintptr(i)),
		))
		WrenchTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Wrench__Array_to_Go(goSlice []Wrench, cSlice []CWrench) {
	for i := 0; i < len(cSlice); i++ {
		WrenchTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Wrench__Array_to_C(cSlice []CWrench, goSlice []Wrench) {
	for i := 0; i < len(goSlice); i++ {
		WrenchTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
