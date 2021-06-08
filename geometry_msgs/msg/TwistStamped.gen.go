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
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/twist_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/TwistStamped", TwistStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewTwistStamped
// function instead.
type TwistStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Twist Twist `yaml:"twist"`
}

// NewTwistStamped creates a new TwistStamped with default values.
func NewTwistStamped() *TwistStamped {
	self := TwistStamped{}
	self.SetDefaults()
	return &self
}

func (t *TwistStamped) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *TwistStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Twist.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var TwistStampedTypeSupport types.MessageTypeSupport = _TwistStampedTypeSupport{}

type _TwistStampedTypeSupport struct{}

func (t _TwistStampedTypeSupport) New() types.Message {
	return NewTwistStamped()
}

func (t _TwistStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__TwistStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__TwistStamped__create())
}

func (t _TwistStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__TwistStamped__destroy((*C.geometry_msgs__msg__TwistStamped)(pointer_to_free))
}

func (t _TwistStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TwistStamped)
	mem := (*C.geometry_msgs__msg__TwistStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	TwistTypeSupport.AsCStruct(unsafe.Pointer(&mem.twist), &m.Twist)
}

func (t _TwistStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TwistStamped)
	mem := (*C.geometry_msgs__msg__TwistStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	TwistTypeSupport.AsGoStruct(&m.Twist, unsafe.Pointer(&mem.twist))
}

func (t _TwistStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__TwistStamped())
}

type CTwistStamped = C.geometry_msgs__msg__TwistStamped
type CTwistStamped__Sequence = C.geometry_msgs__msg__TwistStamped__Sequence

func TwistStamped__Sequence_to_Go(goSlice *[]TwistStamped, cSlice CTwistStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TwistStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__TwistStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__TwistStamped * uintptr(i)),
		))
		TwistStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TwistStamped__Sequence_to_C(cSlice *CTwistStamped__Sequence, goSlice []TwistStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__TwistStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__TwistStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__TwistStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__TwistStamped * uintptr(i)),
		))
		TwistStampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TwistStamped__Array_to_Go(goSlice []TwistStamped, cSlice []CTwistStamped) {
	for i := 0; i < len(cSlice); i++ {
		TwistStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TwistStamped__Array_to_C(cSlice []CTwistStamped, goSlice []TwistStamped) {
	for i := 0; i < len(goSlice); i++ {
		TwistStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
