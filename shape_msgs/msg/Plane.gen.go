/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package shape_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lshape_msgs__rosidl_typesupport_c -lshape_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <shape_msgs/msg/plane.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("shape_msgs/Plane", PlaneTypeSupport)
}

// Do not create instances of this type directly. Always use NewPlane
// function instead.
type Plane struct {
	Coef [4]float64 `yaml:"coef"`// Representation of a plane, using the plane equation ax + by + cz + d = 0.a := coef[0]b := coef[1]c := coef[2]d := coef[3]
}

// NewPlane creates a new Plane with default values.
func NewPlane() *Plane {
	self := Plane{}
	self.SetDefaults()
	return &self
}

func (t *Plane) Clone() *Plane {
	c := &Plane{}
	c.Coef = t.Coef
	return c
}

func (t *Plane) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Plane) SetDefaults() {
	t.Coef = [4]float64{}
}

// ClonePlaneSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePlaneSlice(dst, src []Plane) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PlaneTypeSupport types.MessageTypeSupport = _PlaneTypeSupport{}

type _PlaneTypeSupport struct{}

func (t _PlaneTypeSupport) New() types.Message {
	return NewPlane()
}

func (t _PlaneTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.shape_msgs__msg__Plane
	return (unsafe.Pointer)(C.shape_msgs__msg__Plane__create())
}

func (t _PlaneTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.shape_msgs__msg__Plane__destroy((*C.shape_msgs__msg__Plane)(pointer_to_free))
}

func (t _PlaneTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Plane)
	mem := (*C.shape_msgs__msg__Plane)(dst)
	cSlice_coef := mem.coef[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_coef)), m.Coef[:])
}

func (t _PlaneTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Plane)
	mem := (*C.shape_msgs__msg__Plane)(ros2_message_buffer)
	cSlice_coef := mem.coef[:]
	primitives.Float64__Array_to_Go(m.Coef[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_coef)))
}

func (t _PlaneTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__shape_msgs__msg__Plane())
}

type CPlane = C.shape_msgs__msg__Plane
type CPlane__Sequence = C.shape_msgs__msg__Plane__Sequence

func Plane__Sequence_to_Go(goSlice *[]Plane, cSlice CPlane__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Plane, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.shape_msgs__msg__Plane__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_shape_msgs__msg__Plane * uintptr(i)),
		))
		PlaneTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Plane__Sequence_to_C(cSlice *CPlane__Sequence, goSlice []Plane) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.shape_msgs__msg__Plane)(C.malloc((C.size_t)(C.sizeof_struct_shape_msgs__msg__Plane * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.shape_msgs__msg__Plane)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_shape_msgs__msg__Plane * uintptr(i)),
		))
		PlaneTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Plane__Array_to_Go(goSlice []Plane, cSlice []CPlane) {
	for i := 0; i < len(cSlice); i++ {
		PlaneTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Plane__Array_to_C(cSlice []CPlane, goSlice []Plane) {
	for i := 0; i < len(goSlice); i++ {
		PlaneTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
