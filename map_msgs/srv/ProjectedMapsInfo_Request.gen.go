/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	map_msgs_msg "github.com/tiiuae/rclgo-msgs/map_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/projected_maps_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/ProjectedMapsInfo_Request", ProjectedMapsInfo_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewProjectedMapsInfo_Request
// function instead.
type ProjectedMapsInfo_Request struct {
	ProjectedMapsInfo []map_msgs_msg.ProjectedMapInfo `yaml:"projected_maps_info"`
}

// NewProjectedMapsInfo_Request creates a new ProjectedMapsInfo_Request with default values.
func NewProjectedMapsInfo_Request() *ProjectedMapsInfo_Request {
	self := ProjectedMapsInfo_Request{}
	self.SetDefaults()
	return &self
}

func (t *ProjectedMapsInfo_Request) Clone() *ProjectedMapsInfo_Request {
	c := &ProjectedMapsInfo_Request{}
	if t.ProjectedMapsInfo != nil {
		c.ProjectedMapsInfo = make([]map_msgs_msg.ProjectedMapInfo, len(t.ProjectedMapsInfo))
		map_msgs_msg.CloneProjectedMapInfoSlice(c.ProjectedMapsInfo, t.ProjectedMapsInfo)
	}
	return c
}

func (t *ProjectedMapsInfo_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ProjectedMapsInfo_Request) SetDefaults() {
	t.ProjectedMapsInfo = nil
}

// CloneProjectedMapsInfo_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneProjectedMapsInfo_RequestSlice(dst, src []ProjectedMapsInfo_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ProjectedMapsInfo_RequestTypeSupport types.MessageTypeSupport = _ProjectedMapsInfo_RequestTypeSupport{}

type _ProjectedMapsInfo_RequestTypeSupport struct{}

func (t _ProjectedMapsInfo_RequestTypeSupport) New() types.Message {
	return NewProjectedMapsInfo_Request()
}

func (t _ProjectedMapsInfo_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__ProjectedMapsInfo_Request
	return (unsafe.Pointer)(C.map_msgs__srv__ProjectedMapsInfo_Request__create())
}

func (t _ProjectedMapsInfo_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__ProjectedMapsInfo_Request__destroy((*C.map_msgs__srv__ProjectedMapsInfo_Request)(pointer_to_free))
}

func (t _ProjectedMapsInfo_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ProjectedMapsInfo_Request)
	mem := (*C.map_msgs__srv__ProjectedMapsInfo_Request)(dst)
	map_msgs_msg.ProjectedMapInfo__Sequence_to_C((*map_msgs_msg.CProjectedMapInfo__Sequence)(unsafe.Pointer(&mem.projected_maps_info)), m.ProjectedMapsInfo)
}

func (t _ProjectedMapsInfo_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ProjectedMapsInfo_Request)
	mem := (*C.map_msgs__srv__ProjectedMapsInfo_Request)(ros2_message_buffer)
	map_msgs_msg.ProjectedMapInfo__Sequence_to_Go(&m.ProjectedMapsInfo, *(*map_msgs_msg.CProjectedMapInfo__Sequence)(unsafe.Pointer(&mem.projected_maps_info)))
}

func (t _ProjectedMapsInfo_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__ProjectedMapsInfo_Request())
}

type CProjectedMapsInfo_Request = C.map_msgs__srv__ProjectedMapsInfo_Request
type CProjectedMapsInfo_Request__Sequence = C.map_msgs__srv__ProjectedMapsInfo_Request__Sequence

func ProjectedMapsInfo_Request__Sequence_to_Go(goSlice *[]ProjectedMapsInfo_Request, cSlice CProjectedMapsInfo_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ProjectedMapsInfo_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__srv__ProjectedMapsInfo_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__ProjectedMapsInfo_Request * uintptr(i)),
		))
		ProjectedMapsInfo_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ProjectedMapsInfo_Request__Sequence_to_C(cSlice *CProjectedMapsInfo_Request__Sequence, goSlice []ProjectedMapsInfo_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__srv__ProjectedMapsInfo_Request)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__srv__ProjectedMapsInfo_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__srv__ProjectedMapsInfo_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__ProjectedMapsInfo_Request * uintptr(i)),
		))
		ProjectedMapsInfo_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ProjectedMapsInfo_Request__Array_to_Go(goSlice []ProjectedMapsInfo_Request, cSlice []CProjectedMapsInfo_Request) {
	for i := 0; i < len(cSlice); i++ {
		ProjectedMapsInfo_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ProjectedMapsInfo_Request__Array_to_C(cSlice []CProjectedMapsInfo_Request, goSlice []ProjectedMapsInfo_Request) {
	for i := 0; i < len(goSlice); i++ {
		ProjectedMapsInfo_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
