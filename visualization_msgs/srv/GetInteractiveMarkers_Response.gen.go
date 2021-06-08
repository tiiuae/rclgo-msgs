/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	visualization_msgs_msg "github.com/tiiuae/rclgo-msgs/visualization_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/srv/get_interactive_markers.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/GetInteractiveMarkers_Response", GetInteractiveMarkers_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetInteractiveMarkers_Response
// function instead.
type GetInteractiveMarkers_Response struct {
	SequenceNumber uint64 `yaml:"sequence_number"`// Sequence number.Set to the sequence number of the latest update messageat the time the server received the request.Clients use this to detect if any updates were missed.
	Markers []visualization_msgs_msg.InteractiveMarker `yaml:"markers"`// All interactive markers provided by the server.
}

// NewGetInteractiveMarkers_Response creates a new GetInteractiveMarkers_Response with default values.
func NewGetInteractiveMarkers_Response() *GetInteractiveMarkers_Response {
	self := GetInteractiveMarkers_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetInteractiveMarkers_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *GetInteractiveMarkers_Response) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var GetInteractiveMarkers_ResponseTypeSupport types.MessageTypeSupport = _GetInteractiveMarkers_ResponseTypeSupport{}

type _GetInteractiveMarkers_ResponseTypeSupport struct{}

func (t _GetInteractiveMarkers_ResponseTypeSupport) New() types.Message {
	return NewGetInteractiveMarkers_Response()
}

func (t _GetInteractiveMarkers_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__srv__GetInteractiveMarkers_Response
	return (unsafe.Pointer)(C.visualization_msgs__srv__GetInteractiveMarkers_Response__create())
}

func (t _GetInteractiveMarkers_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__srv__GetInteractiveMarkers_Response__destroy((*C.visualization_msgs__srv__GetInteractiveMarkers_Response)(pointer_to_free))
}

func (t _GetInteractiveMarkers_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetInteractiveMarkers_Response)
	mem := (*C.visualization_msgs__srv__GetInteractiveMarkers_Response)(dst)
	mem.sequence_number = C.uint64_t(m.SequenceNumber)
	visualization_msgs_msg.InteractiveMarker__Sequence_to_C((*visualization_msgs_msg.CInteractiveMarker__Sequence)(unsafe.Pointer(&mem.markers)), m.Markers)
}

func (t _GetInteractiveMarkers_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetInteractiveMarkers_Response)
	mem := (*C.visualization_msgs__srv__GetInteractiveMarkers_Response)(ros2_message_buffer)
	m.SequenceNumber = uint64(mem.sequence_number)
	visualization_msgs_msg.InteractiveMarker__Sequence_to_Go(&m.Markers, *(*visualization_msgs_msg.CInteractiveMarker__Sequence)(unsafe.Pointer(&mem.markers)))
}

func (t _GetInteractiveMarkers_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__srv__GetInteractiveMarkers_Response())
}

type CGetInteractiveMarkers_Response = C.visualization_msgs__srv__GetInteractiveMarkers_Response
type CGetInteractiveMarkers_Response__Sequence = C.visualization_msgs__srv__GetInteractiveMarkers_Response__Sequence

func GetInteractiveMarkers_Response__Sequence_to_Go(goSlice *[]GetInteractiveMarkers_Response, cSlice CGetInteractiveMarkers_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetInteractiveMarkers_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__srv__GetInteractiveMarkers_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__srv__GetInteractiveMarkers_Response * uintptr(i)),
		))
		GetInteractiveMarkers_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetInteractiveMarkers_Response__Sequence_to_C(cSlice *CGetInteractiveMarkers_Response__Sequence, goSlice []GetInteractiveMarkers_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__srv__GetInteractiveMarkers_Response)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__srv__GetInteractiveMarkers_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__srv__GetInteractiveMarkers_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__srv__GetInteractiveMarkers_Response * uintptr(i)),
		))
		GetInteractiveMarkers_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetInteractiveMarkers_Response__Array_to_Go(goSlice []GetInteractiveMarkers_Response, cSlice []CGetInteractiveMarkers_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetInteractiveMarkers_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetInteractiveMarkers_Response__Array_to_C(cSlice []CGetInteractiveMarkers_Response, goSlice []GetInteractiveMarkers_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetInteractiveMarkers_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}