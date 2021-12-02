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
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/polygon_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PolygonStamped", PolygonStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewPolygonStamped
// function instead.
type PolygonStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Polygon Polygon `yaml:"polygon"`
}

// NewPolygonStamped creates a new PolygonStamped with default values.
func NewPolygonStamped() *PolygonStamped {
	self := PolygonStamped{}
	self.SetDefaults()
	return &self
}

func (t *PolygonStamped) Clone() *PolygonStamped {
	c := &PolygonStamped{}
	c.Header = *t.Header.Clone()
	c.Polygon = *t.Polygon.Clone()
	return c
}

func (t *PolygonStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PolygonStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Polygon.SetDefaults()
}

// ClonePolygonStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePolygonStampedSlice(dst, src []PolygonStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PolygonStampedTypeSupport types.MessageTypeSupport = _PolygonStampedTypeSupport{}

type _PolygonStampedTypeSupport struct{}

func (t _PolygonStampedTypeSupport) New() types.Message {
	return NewPolygonStamped()
}

func (t _PolygonStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PolygonStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__PolygonStamped__create())
}

func (t _PolygonStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PolygonStamped__destroy((*C.geometry_msgs__msg__PolygonStamped)(pointer_to_free))
}

func (t _PolygonStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PolygonStamped)
	mem := (*C.geometry_msgs__msg__PolygonStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	PolygonTypeSupport.AsCStruct(unsafe.Pointer(&mem.polygon), &m.Polygon)
}

func (t _PolygonStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PolygonStamped)
	mem := (*C.geometry_msgs__msg__PolygonStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	PolygonTypeSupport.AsGoStruct(&m.Polygon, unsafe.Pointer(&mem.polygon))
}

func (t _PolygonStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PolygonStamped())
}

type CPolygonStamped = C.geometry_msgs__msg__PolygonStamped
type CPolygonStamped__Sequence = C.geometry_msgs__msg__PolygonStamped__Sequence

func PolygonStamped__Sequence_to_Go(goSlice *[]PolygonStamped, cSlice CPolygonStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PolygonStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__PolygonStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PolygonStamped * uintptr(i)),
		))
		PolygonStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PolygonStamped__Sequence_to_C(cSlice *CPolygonStamped__Sequence, goSlice []PolygonStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PolygonStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__PolygonStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__PolygonStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PolygonStamped * uintptr(i)),
		))
		PolygonStampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PolygonStamped__Array_to_Go(goSlice []PolygonStamped, cSlice []CPolygonStamped) {
	for i := 0; i < len(cSlice); i++ {
		PolygonStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PolygonStamped__Array_to_C(cSlice []CPolygonStamped, goSlice []PolygonStamped) {
	for i := 0; i < len(goSlice); i++ {
		PolygonStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
