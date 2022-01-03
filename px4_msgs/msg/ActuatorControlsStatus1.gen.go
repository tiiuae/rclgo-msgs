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

#include <px4_msgs/msg/actuator_controls_status1.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorControlsStatus1", ActuatorControlsStatus1TypeSupport)
}
const (
	ActuatorControlsStatus1_INDEX_ROLL uint8 = 0
	ActuatorControlsStatus1_INDEX_PITCH uint8 = 1
	ActuatorControlsStatus1_INDEX_YAW uint8 = 2
	ActuatorControlsStatus1_INDEX_THROTTLE uint8 = 3
)

// Do not create instances of this type directly. Always use NewActuatorControlsStatus1
// function instead.
type ActuatorControlsStatus1 struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	ControlPower [4]float32 `yaml:"control_power"`
}

// NewActuatorControlsStatus1 creates a new ActuatorControlsStatus1 with default values.
func NewActuatorControlsStatus1() *ActuatorControlsStatus1 {
	self := ActuatorControlsStatus1{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorControlsStatus1) Clone() *ActuatorControlsStatus1 {
	c := &ActuatorControlsStatus1{}
	c.Timestamp = t.Timestamp
	c.ControlPower = t.ControlPower
	return c
}

func (t *ActuatorControlsStatus1) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ActuatorControlsStatus1) SetDefaults() {
	t.Timestamp = 0
	t.ControlPower = [4]float32{}
}

// CloneActuatorControlsStatus1Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneActuatorControlsStatus1Slice(dst, src []ActuatorControlsStatus1) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ActuatorControlsStatus1TypeSupport types.MessageTypeSupport = _ActuatorControlsStatus1TypeSupport{}

type _ActuatorControlsStatus1TypeSupport struct{}

func (t _ActuatorControlsStatus1TypeSupport) New() types.Message {
	return NewActuatorControlsStatus1()
}

func (t _ActuatorControlsStatus1TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControlsStatus1
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControlsStatus1__create())
}

func (t _ActuatorControlsStatus1TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControlsStatus1__destroy((*C.px4_msgs__msg__ActuatorControlsStatus1)(pointer_to_free))
}

func (t _ActuatorControlsStatus1TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorControlsStatus1)
	mem := (*C.px4_msgs__msg__ActuatorControlsStatus1)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_control_power := mem.control_power[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control_power)), m.ControlPower[:])
}

func (t _ActuatorControlsStatus1TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorControlsStatus1)
	mem := (*C.px4_msgs__msg__ActuatorControlsStatus1)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_control_power := mem.control_power[:]
	primitives.Float32__Array_to_Go(m.ControlPower[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control_power)))
}

func (t _ActuatorControlsStatus1TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControlsStatus1())
}

type CActuatorControlsStatus1 = C.px4_msgs__msg__ActuatorControlsStatus1
type CActuatorControlsStatus1__Sequence = C.px4_msgs__msg__ActuatorControlsStatus1__Sequence

func ActuatorControlsStatus1__Sequence_to_Go(goSlice *[]ActuatorControlsStatus1, cSlice CActuatorControlsStatus1__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControlsStatus1, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsStatus1__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus1 * uintptr(i)),
		))
		ActuatorControlsStatus1TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorControlsStatus1__Sequence_to_C(cSlice *CActuatorControlsStatus1__Sequence, goSlice []ActuatorControlsStatus1) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControlsStatus1)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus1 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsStatus1)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus1 * uintptr(i)),
		))
		ActuatorControlsStatus1TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorControlsStatus1__Array_to_Go(goSlice []ActuatorControlsStatus1, cSlice []CActuatorControlsStatus1) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorControlsStatus1TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControlsStatus1__Array_to_C(cSlice []CActuatorControlsStatus1, goSlice []ActuatorControlsStatus1) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorControlsStatus1TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}