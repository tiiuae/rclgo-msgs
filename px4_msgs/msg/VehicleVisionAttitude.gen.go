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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_vision_attitude.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleVisionAttitude", VehicleVisionAttitudeTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleVisionAttitude
// function instead.
type VehicleVisionAttitude struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	Q [4]float32 `yaml:"q"`// Quaternion rotation from the FRD body frame to the NED earth frame
	DeltaQReset [4]float32 `yaml:"delta_q_reset"`// Amount by which quaternion has changed during last reset
	QuatResetCounter uint8 `yaml:"quat_reset_counter"`// Quaternion reset counter
}

// NewVehicleVisionAttitude creates a new VehicleVisionAttitude with default values.
func NewVehicleVisionAttitude() *VehicleVisionAttitude {
	self := VehicleVisionAttitude{}
	self.SetDefaults()
	return &self
}

func (t *VehicleVisionAttitude) Clone() *VehicleVisionAttitude {
	c := &VehicleVisionAttitude{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Q = t.Q
	c.DeltaQReset = t.DeltaQReset
	c.QuatResetCounter = t.QuatResetCounter
	return c
}

func (t *VehicleVisionAttitude) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleVisionAttitude) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Q = [4]float32{}
	t.DeltaQReset = [4]float32{}
	t.QuatResetCounter = 0
}

// CloneVehicleVisionAttitudeSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleVisionAttitudeSlice(dst, src []VehicleVisionAttitude) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleVisionAttitudeTypeSupport types.MessageTypeSupport = _VehicleVisionAttitudeTypeSupport{}

type _VehicleVisionAttitudeTypeSupport struct{}

func (t _VehicleVisionAttitudeTypeSupport) New() types.Message {
	return NewVehicleVisionAttitude()
}

func (t _VehicleVisionAttitudeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleVisionAttitude
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleVisionAttitude__create())
}

func (t _VehicleVisionAttitudeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleVisionAttitude__destroy((*C.px4_msgs__msg__VehicleVisionAttitude)(pointer_to_free))
}

func (t _VehicleVisionAttitudeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleVisionAttitude)
	mem := (*C.px4_msgs__msg__VehicleVisionAttitude)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	cSlice_delta_q_reset := mem.delta_q_reset[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_q_reset)), m.DeltaQReset[:])
	mem.quat_reset_counter = C.uint8_t(m.QuatResetCounter)
}

func (t _VehicleVisionAttitudeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleVisionAttitude)
	mem := (*C.px4_msgs__msg__VehicleVisionAttitude)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_Go(m.Q[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)))
	cSlice_delta_q_reset := mem.delta_q_reset[:]
	primitives.Float32__Array_to_Go(m.DeltaQReset[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_q_reset)))
	m.QuatResetCounter = uint8(mem.quat_reset_counter)
}

func (t _VehicleVisionAttitudeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleVisionAttitude())
}

type CVehicleVisionAttitude = C.px4_msgs__msg__VehicleVisionAttitude
type CVehicleVisionAttitude__Sequence = C.px4_msgs__msg__VehicleVisionAttitude__Sequence

func VehicleVisionAttitude__Sequence_to_Go(goSlice *[]VehicleVisionAttitude, cSlice CVehicleVisionAttitude__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleVisionAttitude, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleVisionAttitude__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleVisionAttitude * uintptr(i)),
		))
		VehicleVisionAttitudeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleVisionAttitude__Sequence_to_C(cSlice *CVehicleVisionAttitude__Sequence, goSlice []VehicleVisionAttitude) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleVisionAttitude)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleVisionAttitude * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleVisionAttitude)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleVisionAttitude * uintptr(i)),
		))
		VehicleVisionAttitudeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleVisionAttitude__Array_to_Go(goSlice []VehicleVisionAttitude, cSlice []CVehicleVisionAttitude) {
	for i := 0; i < len(cSlice); i++ {
		VehicleVisionAttitudeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleVisionAttitude__Array_to_C(cSlice []CVehicleVisionAttitude, goSlice []VehicleVisionAttitude) {
	for i := 0; i < len(goSlice); i++ {
		VehicleVisionAttitudeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
