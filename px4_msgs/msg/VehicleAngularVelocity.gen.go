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

#include <px4_msgs/msg/vehicle_angular_velocity.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleAngularVelocity", VehicleAngularVelocityTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleAngularVelocity
// function instead.
type VehicleAngularVelocity struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// timestamp of the data sample on which this message is based (microseconds)
	Xyz [3]float32 `yaml:"xyz"`// Bias corrected angular velocity about the FRD body frame XYZ-axis in rad/s
}

// NewVehicleAngularVelocity creates a new VehicleAngularVelocity with default values.
func NewVehicleAngularVelocity() *VehicleAngularVelocity {
	self := VehicleAngularVelocity{}
	self.SetDefaults()
	return &self
}

func (t *VehicleAngularVelocity) Clone() *VehicleAngularVelocity {
	c := &VehicleAngularVelocity{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Xyz = t.Xyz
	return c
}

func (t *VehicleAngularVelocity) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleAngularVelocity) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Xyz = [3]float32{}
}

// CloneVehicleAngularVelocitySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleAngularVelocitySlice(dst, src []VehicleAngularVelocity) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleAngularVelocityTypeSupport types.MessageTypeSupport = _VehicleAngularVelocityTypeSupport{}

type _VehicleAngularVelocityTypeSupport struct{}

func (t _VehicleAngularVelocityTypeSupport) New() types.Message {
	return NewVehicleAngularVelocity()
}

func (t _VehicleAngularVelocityTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAngularVelocity
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAngularVelocity__create())
}

func (t _VehicleAngularVelocityTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAngularVelocity__destroy((*C.px4_msgs__msg__VehicleAngularVelocity)(pointer_to_free))
}

func (t _VehicleAngularVelocityTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleAngularVelocity)
	mem := (*C.px4_msgs__msg__VehicleAngularVelocity)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_xyz := mem.xyz[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_xyz)), m.Xyz[:])
}

func (t _VehicleAngularVelocityTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleAngularVelocity)
	mem := (*C.px4_msgs__msg__VehicleAngularVelocity)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_xyz := mem.xyz[:]
	primitives.Float32__Array_to_Go(m.Xyz[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_xyz)))
}

func (t _VehicleAngularVelocityTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAngularVelocity())
}

type CVehicleAngularVelocity = C.px4_msgs__msg__VehicleAngularVelocity
type CVehicleAngularVelocity__Sequence = C.px4_msgs__msg__VehicleAngularVelocity__Sequence

func VehicleAngularVelocity__Sequence_to_Go(goSlice *[]VehicleAngularVelocity, cSlice CVehicleAngularVelocity__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAngularVelocity, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAngularVelocity__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocity * uintptr(i)),
		))
		VehicleAngularVelocityTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleAngularVelocity__Sequence_to_C(cSlice *CVehicleAngularVelocity__Sequence, goSlice []VehicleAngularVelocity) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAngularVelocity)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocity * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAngularVelocity)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocity * uintptr(i)),
		))
		VehicleAngularVelocityTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleAngularVelocity__Array_to_Go(goSlice []VehicleAngularVelocity, cSlice []CVehicleAngularVelocity) {
	for i := 0; i < len(cSlice); i++ {
		VehicleAngularVelocityTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAngularVelocity__Array_to_C(cSlice []CVehicleAngularVelocity, goSlice []VehicleAngularVelocity) {
	for i := 0; i < len(goSlice); i++ {
		VehicleAngularVelocityTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
