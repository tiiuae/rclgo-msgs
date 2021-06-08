/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/offboard_control_mode.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/OffboardControlMode", OffboardControlModeTypeSupport)
}

// Do not create instances of this type directly. Always use NewOffboardControlMode
// function instead.
type OffboardControlMode struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Position bool `yaml:"position"`
	Velocity bool `yaml:"velocity"`
	Acceleration bool `yaml:"acceleration"`
	Attitude bool `yaml:"attitude"`
	BodyRate bool `yaml:"body_rate"`
}

// NewOffboardControlMode creates a new OffboardControlMode with default values.
func NewOffboardControlMode() *OffboardControlMode {
	self := OffboardControlMode{}
	self.SetDefaults()
	return &self
}

func (t *OffboardControlMode) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *OffboardControlMode) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var OffboardControlModeTypeSupport types.MessageTypeSupport = _OffboardControlModeTypeSupport{}

type _OffboardControlModeTypeSupport struct{}

func (t _OffboardControlModeTypeSupport) New() types.Message {
	return NewOffboardControlMode()
}

func (t _OffboardControlModeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__OffboardControlMode
	return (unsafe.Pointer)(C.px4_msgs__msg__OffboardControlMode__create())
}

func (t _OffboardControlModeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__OffboardControlMode__destroy((*C.px4_msgs__msg__OffboardControlMode)(pointer_to_free))
}

func (t _OffboardControlModeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*OffboardControlMode)
	mem := (*C.px4_msgs__msg__OffboardControlMode)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.position = C.bool(m.Position)
	mem.velocity = C.bool(m.Velocity)
	mem.acceleration = C.bool(m.Acceleration)
	mem.attitude = C.bool(m.Attitude)
	mem.body_rate = C.bool(m.BodyRate)
}

func (t _OffboardControlModeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*OffboardControlMode)
	mem := (*C.px4_msgs__msg__OffboardControlMode)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Position = bool(mem.position)
	m.Velocity = bool(mem.velocity)
	m.Acceleration = bool(mem.acceleration)
	m.Attitude = bool(mem.attitude)
	m.BodyRate = bool(mem.body_rate)
}

func (t _OffboardControlModeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__OffboardControlMode())
}

type COffboardControlMode = C.px4_msgs__msg__OffboardControlMode
type COffboardControlMode__Sequence = C.px4_msgs__msg__OffboardControlMode__Sequence

func OffboardControlMode__Sequence_to_Go(goSlice *[]OffboardControlMode, cSlice COffboardControlMode__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]OffboardControlMode, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__OffboardControlMode__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OffboardControlMode * uintptr(i)),
		))
		OffboardControlModeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func OffboardControlMode__Sequence_to_C(cSlice *COffboardControlMode__Sequence, goSlice []OffboardControlMode) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__OffboardControlMode)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__OffboardControlMode * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__OffboardControlMode)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OffboardControlMode * uintptr(i)),
		))
		OffboardControlModeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func OffboardControlMode__Array_to_Go(goSlice []OffboardControlMode, cSlice []COffboardControlMode) {
	for i := 0; i < len(cSlice); i++ {
		OffboardControlModeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func OffboardControlMode__Array_to_C(cSlice []COffboardControlMode, goSlice []OffboardControlMode) {
	for i := 0; i < len(goSlice); i++ {
		OffboardControlModeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
