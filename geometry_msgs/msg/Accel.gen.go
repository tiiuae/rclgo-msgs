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

#include <geometry_msgs/msg/accel.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Accel", AccelTypeSupport)
}

// Do not create instances of this type directly. Always use NewAccel
// function instead.
type Accel struct {
	Linear Vector3 `yaml:"linear"`// This expresses acceleration in free space broken into its linear and angular parts.
	Angular Vector3 `yaml:"angular"`// This expresses acceleration in free space broken into its linear and angular parts.
}

// NewAccel creates a new Accel with default values.
func NewAccel() *Accel {
	self := Accel{}
	self.SetDefaults()
	return &self
}

func (t *Accel) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Accel) SetDefaults() {
	t.Linear.SetDefaults()
	t.Angular.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var AccelTypeSupport types.MessageTypeSupport = _AccelTypeSupport{}

type _AccelTypeSupport struct{}

func (t _AccelTypeSupport) New() types.Message {
	return NewAccel()
}

func (t _AccelTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Accel
	return (unsafe.Pointer)(C.geometry_msgs__msg__Accel__create())
}

func (t _AccelTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Accel__destroy((*C.geometry_msgs__msg__Accel)(pointer_to_free))
}

func (t _AccelTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Accel)
	mem := (*C.geometry_msgs__msg__Accel)(dst)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.linear), &m.Linear)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.angular), &m.Angular)
}

func (t _AccelTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Accel)
	mem := (*C.geometry_msgs__msg__Accel)(ros2_message_buffer)
	Vector3TypeSupport.AsGoStruct(&m.Linear, unsafe.Pointer(&mem.linear))
	Vector3TypeSupport.AsGoStruct(&m.Angular, unsafe.Pointer(&mem.angular))
}

func (t _AccelTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Accel())
}

type CAccel = C.geometry_msgs__msg__Accel
type CAccel__Sequence = C.geometry_msgs__msg__Accel__Sequence

func Accel__Sequence_to_Go(goSlice *[]Accel, cSlice CAccel__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Accel, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Accel__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Accel * uintptr(i)),
		))
		AccelTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Accel__Sequence_to_C(cSlice *CAccel__Sequence, goSlice []Accel) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Accel)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Accel * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Accel)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Accel * uintptr(i)),
		))
		AccelTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Accel__Array_to_Go(goSlice []Accel, cSlice []CAccel) {
	for i := 0; i < len(cSlice); i++ {
		AccelTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Accel__Array_to_C(cSlice []CAccel, goSlice []Accel) {
	for i := 0; i < len(goSlice); i++ {
		AccelTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
