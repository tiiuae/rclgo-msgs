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

#include <px4_msgs/msg/actuator_armed.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorArmed", ActuatorArmedTypeSupport)
}

// Do not create instances of this type directly. Always use NewActuatorArmed
// function instead.
type ActuatorArmed struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Armed bool `yaml:"armed"`// Set to true if system is armed
	Prearmed bool `yaml:"prearmed"`// Set to true if the actuator safety is disabled but motors are not armed
	ReadyToArm bool `yaml:"ready_to_arm"`// Set to true if system is ready to be armed
	Lockdown bool `yaml:"lockdown"`// Set to true if actuators are forced to being disabled (due to emergency or HIL)
	ManualLockdown bool `yaml:"manual_lockdown"`// Set to true if manual throttle kill switch is engaged
	ForceFailsafe bool `yaml:"force_failsafe"`// Set to true if the actuators are forced to the failsafe position
	InEscCalibrationMode bool `yaml:"in_esc_calibration_mode"`// IO/FMU should ignore messages from the actuator controls topics
	SoftStop bool `yaml:"soft_stop"`// Set to true if we need to ESCs to remove the idle constraint
}

// NewActuatorArmed creates a new ActuatorArmed with default values.
func NewActuatorArmed() *ActuatorArmed {
	self := ActuatorArmed{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorArmed) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *ActuatorArmed) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var ActuatorArmedTypeSupport types.MessageTypeSupport = _ActuatorArmedTypeSupport{}

type _ActuatorArmedTypeSupport struct{}

func (t _ActuatorArmedTypeSupport) New() types.Message {
	return NewActuatorArmed()
}

func (t _ActuatorArmedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorArmed
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorArmed__create())
}

func (t _ActuatorArmedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorArmed__destroy((*C.px4_msgs__msg__ActuatorArmed)(pointer_to_free))
}

func (t _ActuatorArmedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorArmed)
	mem := (*C.px4_msgs__msg__ActuatorArmed)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.armed = C.bool(m.Armed)
	mem.prearmed = C.bool(m.Prearmed)
	mem.ready_to_arm = C.bool(m.ReadyToArm)
	mem.lockdown = C.bool(m.Lockdown)
	mem.manual_lockdown = C.bool(m.ManualLockdown)
	mem.force_failsafe = C.bool(m.ForceFailsafe)
	mem.in_esc_calibration_mode = C.bool(m.InEscCalibrationMode)
	mem.soft_stop = C.bool(m.SoftStop)
}

func (t _ActuatorArmedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorArmed)
	mem := (*C.px4_msgs__msg__ActuatorArmed)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Armed = bool(mem.armed)
	m.Prearmed = bool(mem.prearmed)
	m.ReadyToArm = bool(mem.ready_to_arm)
	m.Lockdown = bool(mem.lockdown)
	m.ManualLockdown = bool(mem.manual_lockdown)
	m.ForceFailsafe = bool(mem.force_failsafe)
	m.InEscCalibrationMode = bool(mem.in_esc_calibration_mode)
	m.SoftStop = bool(mem.soft_stop)
}

func (t _ActuatorArmedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorArmed())
}

type CActuatorArmed = C.px4_msgs__msg__ActuatorArmed
type CActuatorArmed__Sequence = C.px4_msgs__msg__ActuatorArmed__Sequence

func ActuatorArmed__Sequence_to_Go(goSlice *[]ActuatorArmed, cSlice CActuatorArmed__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorArmed, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorArmed__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorArmed * uintptr(i)),
		))
		ActuatorArmedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorArmed__Sequence_to_C(cSlice *CActuatorArmed__Sequence, goSlice []ActuatorArmed) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorArmed)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorArmed * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorArmed)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorArmed * uintptr(i)),
		))
		ActuatorArmedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorArmed__Array_to_Go(goSlice []ActuatorArmed, cSlice []CActuatorArmed) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorArmedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorArmed__Array_to_C(cSlice []CActuatorArmed, goSlice []ActuatorArmed) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorArmedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}