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

#include <px4_msgs/msg/vehicle_thrust_setpoint.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleThrustSetpoint", VehicleThrustSetpointTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleThrustSetpoint
// function instead.
type VehicleThrustSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// timestamp of the data sample on which this message is based (microseconds)
	Xyz [3]float32 `yaml:"xyz"`// thrust setpoint along X, Y, Z body axis (in N)
}

// NewVehicleThrustSetpoint creates a new VehicleThrustSetpoint with default values.
func NewVehicleThrustSetpoint() *VehicleThrustSetpoint {
	self := VehicleThrustSetpoint{}
	self.SetDefaults()
	return &self
}

func (t *VehicleThrustSetpoint) Clone() *VehicleThrustSetpoint {
	c := &VehicleThrustSetpoint{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Xyz = t.Xyz
	return c
}

func (t *VehicleThrustSetpoint) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleThrustSetpoint) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Xyz = [3]float32{}
}

// CloneVehicleThrustSetpointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleThrustSetpointSlice(dst, src []VehicleThrustSetpoint) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleThrustSetpointTypeSupport types.MessageTypeSupport = _VehicleThrustSetpointTypeSupport{}

type _VehicleThrustSetpointTypeSupport struct{}

func (t _VehicleThrustSetpointTypeSupport) New() types.Message {
	return NewVehicleThrustSetpoint()
}

func (t _VehicleThrustSetpointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleThrustSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleThrustSetpoint__create())
}

func (t _VehicleThrustSetpointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleThrustSetpoint__destroy((*C.px4_msgs__msg__VehicleThrustSetpoint)(pointer_to_free))
}

func (t _VehicleThrustSetpointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleThrustSetpoint)
	mem := (*C.px4_msgs__msg__VehicleThrustSetpoint)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_xyz := mem.xyz[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_xyz)), m.Xyz[:])
}

func (t _VehicleThrustSetpointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleThrustSetpoint)
	mem := (*C.px4_msgs__msg__VehicleThrustSetpoint)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_xyz := mem.xyz[:]
	primitives.Float32__Array_to_Go(m.Xyz[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_xyz)))
}

func (t _VehicleThrustSetpointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleThrustSetpoint())
}

type CVehicleThrustSetpoint = C.px4_msgs__msg__VehicleThrustSetpoint
type CVehicleThrustSetpoint__Sequence = C.px4_msgs__msg__VehicleThrustSetpoint__Sequence

func VehicleThrustSetpoint__Sequence_to_Go(goSlice *[]VehicleThrustSetpoint, cSlice CVehicleThrustSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleThrustSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleThrustSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleThrustSetpoint * uintptr(i)),
		))
		VehicleThrustSetpointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleThrustSetpoint__Sequence_to_C(cSlice *CVehicleThrustSetpoint__Sequence, goSlice []VehicleThrustSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleThrustSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleThrustSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleThrustSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleThrustSetpoint * uintptr(i)),
		))
		VehicleThrustSetpointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleThrustSetpoint__Array_to_Go(goSlice []VehicleThrustSetpoint, cSlice []CVehicleThrustSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		VehicleThrustSetpointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleThrustSetpoint__Array_to_C(cSlice []CVehicleThrustSetpoint, goSlice []VehicleThrustSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		VehicleThrustSetpointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
