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

#include <px4_msgs/msg/actuator_controls0.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorControls0", ActuatorControls0TypeSupport)
}
const (
	ActuatorControls0_NUM_ACTUATOR_CONTROLS uint8 = 8
	ActuatorControls0_NUM_ACTUATOR_CONTROL_GROUPS uint8 = 6
	ActuatorControls0_INDEX_ROLL uint8 = 0
	ActuatorControls0_INDEX_PITCH uint8 = 1
	ActuatorControls0_INDEX_YAW uint8 = 2
	ActuatorControls0_INDEX_THROTTLE uint8 = 3
	ActuatorControls0_INDEX_FLAPS uint8 = 4
	ActuatorControls0_INDEX_SPOILERS uint8 = 5
	ActuatorControls0_INDEX_AIRBRAKES uint8 = 6
	ActuatorControls0_INDEX_LANDING_GEAR uint8 = 7
	ActuatorControls0_INDEX_GIMBAL_SHUTTER uint8 = 3
	ActuatorControls0_INDEX_CAMERA_ZOOM uint8 = 4
	ActuatorControls0_GROUP_INDEX_ATTITUDE uint8 = 0
	ActuatorControls0_GROUP_INDEX_ATTITUDE_ALTERNATE uint8 = 1
	ActuatorControls0_GROUP_INDEX_GIMBAL uint8 = 2
	ActuatorControls0_GROUP_INDEX_MANUAL_PASSTHROUGH uint8 = 3
	ActuatorControls0_GROUP_INDEX_ALLOCATED_PART1 uint8 = 4
	ActuatorControls0_GROUP_INDEX_ALLOCATED_PART2 uint8 = 5
	ActuatorControls0_GROUP_INDEX_PAYLOAD uint8 = 6
)

// Do not create instances of this type directly. Always use NewActuatorControls0
// function instead.
type ActuatorControls0 struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp the data this control response is based on was sampled
	Control [8]float32 `yaml:"control"`
}

// NewActuatorControls0 creates a new ActuatorControls0 with default values.
func NewActuatorControls0() *ActuatorControls0 {
	self := ActuatorControls0{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorControls0) Clone() *ActuatorControls0 {
	c := &ActuatorControls0{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Control = t.Control
	return c
}

func (t *ActuatorControls0) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ActuatorControls0) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Control = [8]float32{}
}

// CloneActuatorControls0Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneActuatorControls0Slice(dst, src []ActuatorControls0) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ActuatorControls0TypeSupport types.MessageTypeSupport = _ActuatorControls0TypeSupport{}

type _ActuatorControls0TypeSupport struct{}

func (t _ActuatorControls0TypeSupport) New() types.Message {
	return NewActuatorControls0()
}

func (t _ActuatorControls0TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControls0
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControls0__create())
}

func (t _ActuatorControls0TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControls0__destroy((*C.px4_msgs__msg__ActuatorControls0)(pointer_to_free))
}

func (t _ActuatorControls0TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorControls0)
	mem := (*C.px4_msgs__msg__ActuatorControls0)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_control := mem.control[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control)), m.Control[:])
}

func (t _ActuatorControls0TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorControls0)
	mem := (*C.px4_msgs__msg__ActuatorControls0)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_control := mem.control[:]
	primitives.Float32__Array_to_Go(m.Control[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control)))
}

func (t _ActuatorControls0TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControls0())
}

type CActuatorControls0 = C.px4_msgs__msg__ActuatorControls0
type CActuatorControls0__Sequence = C.px4_msgs__msg__ActuatorControls0__Sequence

func ActuatorControls0__Sequence_to_Go(goSlice *[]ActuatorControls0, cSlice CActuatorControls0__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControls0, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControls0__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControls0 * uintptr(i)),
		))
		ActuatorControls0TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorControls0__Sequence_to_C(cSlice *CActuatorControls0__Sequence, goSlice []ActuatorControls0) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControls0)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControls0 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControls0)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControls0 * uintptr(i)),
		))
		ActuatorControls0TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorControls0__Array_to_Go(goSlice []ActuatorControls0, cSlice []CActuatorControls0) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorControls0TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControls0__Array_to_C(cSlice []CActuatorControls0, goSlice []ActuatorControls0) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorControls0TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
