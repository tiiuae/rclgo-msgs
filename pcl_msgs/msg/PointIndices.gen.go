/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pcl_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpcl_msgs__rosidl_typesupport_c -lpcl_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pcl_msgs/msg/point_indices.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pcl_msgs/PointIndices", PointIndicesTypeSupport)
}

// Do not create instances of this type directly. Always use NewPointIndices
// function instead.
type PointIndices struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Indices []int32 `yaml:"indices"`
}

// NewPointIndices creates a new PointIndices with default values.
func NewPointIndices() *PointIndices {
	self := PointIndices{}
	self.SetDefaults()
	return &self
}

func (t *PointIndices) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *PointIndices) SetDefaults() {
	t.Header.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var PointIndicesTypeSupport types.MessageTypeSupport = _PointIndicesTypeSupport{}

type _PointIndicesTypeSupport struct{}

func (t _PointIndicesTypeSupport) New() types.Message {
	return NewPointIndices()
}

func (t _PointIndicesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pcl_msgs__msg__PointIndices
	return (unsafe.Pointer)(C.pcl_msgs__msg__PointIndices__create())
}

func (t _PointIndicesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pcl_msgs__msg__PointIndices__destroy((*C.pcl_msgs__msg__PointIndices)(pointer_to_free))
}

func (t _PointIndicesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PointIndices)
	mem := (*C.pcl_msgs__msg__PointIndices)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.indices)), m.Indices)
}

func (t _PointIndicesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PointIndices)
	mem := (*C.pcl_msgs__msg__PointIndices)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.Int32__Sequence_to_Go(&m.Indices, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.indices)))
}

func (t _PointIndicesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pcl_msgs__msg__PointIndices())
}

type CPointIndices = C.pcl_msgs__msg__PointIndices
type CPointIndices__Sequence = C.pcl_msgs__msg__PointIndices__Sequence

func PointIndices__Sequence_to_Go(goSlice *[]PointIndices, cSlice CPointIndices__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointIndices, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pcl_msgs__msg__PointIndices__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__PointIndices * uintptr(i)),
		))
		PointIndicesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PointIndices__Sequence_to_C(cSlice *CPointIndices__Sequence, goSlice []PointIndices) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pcl_msgs__msg__PointIndices)(C.malloc((C.size_t)(C.sizeof_struct_pcl_msgs__msg__PointIndices * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pcl_msgs__msg__PointIndices)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__PointIndices * uintptr(i)),
		))
		PointIndicesTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PointIndices__Array_to_Go(goSlice []PointIndices, cSlice []CPointIndices) {
	for i := 0; i < len(cSlice); i++ {
		PointIndicesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PointIndices__Array_to_C(cSlice []CPointIndices, goSlice []PointIndices) {
	for i := 0; i < len(goSlice); i++ {
		PointIndicesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
