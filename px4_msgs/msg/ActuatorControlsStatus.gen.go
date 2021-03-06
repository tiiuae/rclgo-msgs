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

#include <px4_msgs/msg/actuator_controls_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorControlsStatus", ActuatorControlsStatusTypeSupport)
}
const (
	ActuatorControlsStatus_INDEX_ROLL uint8 = 0
	ActuatorControlsStatus_INDEX_PITCH uint8 = 1
	ActuatorControlsStatus_INDEX_YAW uint8 = 2
	ActuatorControlsStatus_INDEX_THROTTLE uint8 = 3
)

// Do not create instances of this type directly. Always use NewActuatorControlsStatus
// function instead.
type ActuatorControlsStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	ControlPower [4]float32 `yaml:"control_power"`
}

// NewActuatorControlsStatus creates a new ActuatorControlsStatus with default values.
func NewActuatorControlsStatus() *ActuatorControlsStatus {
	self := ActuatorControlsStatus{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorControlsStatus) Clone() *ActuatorControlsStatus {
	c := &ActuatorControlsStatus{}
	c.Timestamp = t.Timestamp
	c.ControlPower = t.ControlPower
	return c
}

func (t *ActuatorControlsStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ActuatorControlsStatus) SetDefaults() {
	t.Timestamp = 0
	t.ControlPower = [4]float32{}
}

// CloneActuatorControlsStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneActuatorControlsStatusSlice(dst, src []ActuatorControlsStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ActuatorControlsStatusTypeSupport types.MessageTypeSupport = _ActuatorControlsStatusTypeSupport{}

type _ActuatorControlsStatusTypeSupport struct{}

func (t _ActuatorControlsStatusTypeSupport) New() types.Message {
	return NewActuatorControlsStatus()
}

func (t _ActuatorControlsStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControlsStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControlsStatus__create())
}

func (t _ActuatorControlsStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControlsStatus__destroy((*C.px4_msgs__msg__ActuatorControlsStatus)(pointer_to_free))
}

func (t _ActuatorControlsStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorControlsStatus)
	mem := (*C.px4_msgs__msg__ActuatorControlsStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_control_power := mem.control_power[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control_power)), m.ControlPower[:])
}

func (t _ActuatorControlsStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorControlsStatus)
	mem := (*C.px4_msgs__msg__ActuatorControlsStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_control_power := mem.control_power[:]
	primitives.Float32__Array_to_Go(m.ControlPower[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control_power)))
}

func (t _ActuatorControlsStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControlsStatus())
}

type CActuatorControlsStatus = C.px4_msgs__msg__ActuatorControlsStatus
type CActuatorControlsStatus__Sequence = C.px4_msgs__msg__ActuatorControlsStatus__Sequence

func ActuatorControlsStatus__Sequence_to_Go(goSlice *[]ActuatorControlsStatus, cSlice CActuatorControlsStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControlsStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus * uintptr(i)),
		))
		ActuatorControlsStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorControlsStatus__Sequence_to_C(cSlice *CActuatorControlsStatus__Sequence, goSlice []ActuatorControlsStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControlsStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus * uintptr(i)),
		))
		ActuatorControlsStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorControlsStatus__Array_to_Go(goSlice []ActuatorControlsStatus, cSlice []CActuatorControlsStatus) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorControlsStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControlsStatus__Array_to_C(cSlice []CActuatorControlsStatus, goSlice []ActuatorControlsStatus) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorControlsStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
