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

#include <shape_msgs/msg/mesh_triangle.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("shape_msgs/MeshTriangle", MeshTriangleTypeSupport)
}

// Do not create instances of this type directly. Always use NewMeshTriangle
// function instead.
type MeshTriangle struct {
	VertexIndices [3]uint32 `yaml:"vertex_indices"`
}

// NewMeshTriangle creates a new MeshTriangle with default values.
func NewMeshTriangle() *MeshTriangle {
	self := MeshTriangle{}
	self.SetDefaults()
	return &self
}

func (t *MeshTriangle) Clone() *MeshTriangle {
	c := &MeshTriangle{}
	c.VertexIndices = t.VertexIndices
	return c
}

func (t *MeshTriangle) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MeshTriangle) SetDefaults() {
	t.VertexIndices = [3]uint32{}
}

// CloneMeshTriangleSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMeshTriangleSlice(dst, src []MeshTriangle) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MeshTriangleTypeSupport types.MessageTypeSupport = _MeshTriangleTypeSupport{}

type _MeshTriangleTypeSupport struct{}

func (t _MeshTriangleTypeSupport) New() types.Message {
	return NewMeshTriangle()
}

func (t _MeshTriangleTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.shape_msgs__msg__MeshTriangle
	return (unsafe.Pointer)(C.shape_msgs__msg__MeshTriangle__create())
}

func (t _MeshTriangleTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.shape_msgs__msg__MeshTriangle__destroy((*C.shape_msgs__msg__MeshTriangle)(pointer_to_free))
}

func (t _MeshTriangleTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MeshTriangle)
	mem := (*C.shape_msgs__msg__MeshTriangle)(dst)
	cSlice_vertex_indices := mem.vertex_indices[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_vertex_indices)), m.VertexIndices[:])
}

func (t _MeshTriangleTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MeshTriangle)
	mem := (*C.shape_msgs__msg__MeshTriangle)(ros2_message_buffer)
	cSlice_vertex_indices := mem.vertex_indices[:]
	primitives.Uint32__Array_to_Go(m.VertexIndices[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_vertex_indices)))
}

func (t _MeshTriangleTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__shape_msgs__msg__MeshTriangle())
}

type CMeshTriangle = C.shape_msgs__msg__MeshTriangle
type CMeshTriangle__Sequence = C.shape_msgs__msg__MeshTriangle__Sequence

func MeshTriangle__Sequence_to_Go(goSlice *[]MeshTriangle, cSlice CMeshTriangle__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MeshTriangle, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.shape_msgs__msg__MeshTriangle__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_shape_msgs__msg__MeshTriangle * uintptr(i)),
		))
		MeshTriangleTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MeshTriangle__Sequence_to_C(cSlice *CMeshTriangle__Sequence, goSlice []MeshTriangle) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.shape_msgs__msg__MeshTriangle)(C.malloc((C.size_t)(C.sizeof_struct_shape_msgs__msg__MeshTriangle * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.shape_msgs__msg__MeshTriangle)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_shape_msgs__msg__MeshTriangle * uintptr(i)),
		))
		MeshTriangleTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MeshTriangle__Array_to_Go(goSlice []MeshTriangle, cSlice []CMeshTriangle) {
	for i := 0; i < len(cSlice); i++ {
		MeshTriangleTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MeshTriangle__Array_to_C(cSlice []CMeshTriangle, goSlice []MeshTriangle) {
	for i := 0; i < len(goSlice); i++ {
		MeshTriangleTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
