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

#include <px4_msgs/msg/ekf_gps_drift.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EkfGpsDrift", EkfGpsDriftTypeSupport)
}

// Do not create instances of this type directly. Always use NewEkfGpsDrift
// function instead.
type EkfGpsDrift struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	HposDriftRate float32 `yaml:"hpos_drift_rate"`// Horizontal position rate magnitude checked using EKF2_REQ_HDRIFT (m/s)
	VposDriftRate float32 `yaml:"vpos_drift_rate"`// Vertical position rate magnitude checked using EKF2_REQ_VDRIFT (m/s)
	Hspd float32 `yaml:"hspd"`// Filtered horizontal velocity magnitude checked using EKF2_REQ_HDRIFT (m/s)
	Blocked bool `yaml:"blocked"`// true when drift calculation is blocked due to IMU movement check controlled by EKF2_MOVE_TEST
}

// NewEkfGpsDrift creates a new EkfGpsDrift with default values.
func NewEkfGpsDrift() *EkfGpsDrift {
	self := EkfGpsDrift{}
	self.SetDefaults()
	return &self
}

func (t *EkfGpsDrift) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *EkfGpsDrift) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var EkfGpsDriftTypeSupport types.MessageTypeSupport = _EkfGpsDriftTypeSupport{}

type _EkfGpsDriftTypeSupport struct{}

func (t _EkfGpsDriftTypeSupport) New() types.Message {
	return NewEkfGpsDrift()
}

func (t _EkfGpsDriftTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EkfGpsDrift
	return (unsafe.Pointer)(C.px4_msgs__msg__EkfGpsDrift__create())
}

func (t _EkfGpsDriftTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EkfGpsDrift__destroy((*C.px4_msgs__msg__EkfGpsDrift)(pointer_to_free))
}

func (t _EkfGpsDriftTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EkfGpsDrift)
	mem := (*C.px4_msgs__msg__EkfGpsDrift)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.hpos_drift_rate = C.float(m.HposDriftRate)
	mem.vpos_drift_rate = C.float(m.VposDriftRate)
	mem.hspd = C.float(m.Hspd)
	mem.blocked = C.bool(m.Blocked)
}

func (t _EkfGpsDriftTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EkfGpsDrift)
	mem := (*C.px4_msgs__msg__EkfGpsDrift)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.HposDriftRate = float32(mem.hpos_drift_rate)
	m.VposDriftRate = float32(mem.vpos_drift_rate)
	m.Hspd = float32(mem.hspd)
	m.Blocked = bool(mem.blocked)
}

func (t _EkfGpsDriftTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EkfGpsDrift())
}

type CEkfGpsDrift = C.px4_msgs__msg__EkfGpsDrift
type CEkfGpsDrift__Sequence = C.px4_msgs__msg__EkfGpsDrift__Sequence

func EkfGpsDrift__Sequence_to_Go(goSlice *[]EkfGpsDrift, cSlice CEkfGpsDrift__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EkfGpsDrift, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EkfGpsDrift__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EkfGpsDrift * uintptr(i)),
		))
		EkfGpsDriftTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EkfGpsDrift__Sequence_to_C(cSlice *CEkfGpsDrift__Sequence, goSlice []EkfGpsDrift) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EkfGpsDrift)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EkfGpsDrift * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EkfGpsDrift)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EkfGpsDrift * uintptr(i)),
		))
		EkfGpsDriftTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EkfGpsDrift__Array_to_Go(goSlice []EkfGpsDrift, cSlice []CEkfGpsDrift) {
	for i := 0; i < len(cSlice); i++ {
		EkfGpsDriftTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EkfGpsDrift__Array_to_C(cSlice []CEkfGpsDrift, goSlice []EkfGpsDrift) {
	for i := 0; i < len(goSlice); i++ {
		EkfGpsDriftTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
