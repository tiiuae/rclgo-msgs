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

#include <px4_msgs/msg/manual_control_switches.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ManualControlSwitches", ManualControlSwitchesTypeSupport)
}
const (
	ManualControlSwitches_SWITCH_POS_NONE uint8 = 0// switch is not mapped
	ManualControlSwitches_SWITCH_POS_ON uint8 = 1// switch activated (value = 1)
	ManualControlSwitches_SWITCH_POS_MIDDLE uint8 = 2// middle position (value = 0)
	ManualControlSwitches_SWITCH_POS_OFF uint8 = 3// switch not activated (value = -1)
	ManualControlSwitches_MODE_SLOT_NONE uint8 = 0// no mode slot assigned
	ManualControlSwitches_MODE_SLOT_1 uint8 = 1// mode slot 1 selected
	ManualControlSwitches_MODE_SLOT_2 uint8 = 2// mode slot 2 selected
	ManualControlSwitches_MODE_SLOT_3 uint8 = 3// mode slot 3 selected
	ManualControlSwitches_MODE_SLOT_4 uint8 = 4// mode slot 4 selected
	ManualControlSwitches_MODE_SLOT_5 uint8 = 5// mode slot 5 selected
	ManualControlSwitches_MODE_SLOT_6 uint8 = 6// mode slot 6 selected
	ManualControlSwitches_MODE_SLOT_NUM uint8 = 6// number of slots
)

// Do not create instances of this type directly. Always use NewManualControlSwitches
// function instead.
type ManualControlSwitches struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	ModeSlot uint8 `yaml:"mode_slot"`// the slot a specific model selector is in
	ArmSwitch uint8 `yaml:"arm_switch"`// arm/disarm switch: _DISARMED_, ARMED
	ReturnSwitch uint8 `yaml:"return_switch"`// return to launch 2 position switch (mandatory): _NORMAL_, RTL
	LoiterSwitch uint8 `yaml:"loiter_switch"`// loiter 2 position switch (optional): _MISSION_, LOITER
	OffboardSwitch uint8 `yaml:"offboard_switch"`// offboard 2 position switch (optional): _NORMAL_, OFFBOARD
	KillSwitch uint8 `yaml:"kill_switch"`// throttle kill: _NORMAL_, KILL
	GearSwitch uint8 `yaml:"gear_switch"`// landing gear switch: _DOWN_, UP
	TransitionSwitch uint8 `yaml:"transition_switch"`// VTOL transition switch: _HOVER, FORWARD_FLIGHT
	ModeSwitch uint8 `yaml:"mode_switch"`// main mode 3 position switch (mandatory): _MANUAL_, ASSIST, AUTO. legacy "advanced" switch configuration (will be removed soon)
	ManSwitch uint8 `yaml:"man_switch"`// manual switch (only relevant for fixed wings, optional): _STABILIZED_, MANUAL. legacy "advanced" switch configuration (will be removed soon)
	AcroSwitch uint8 `yaml:"acro_switch"`// acro 2 position switch (optional): _MANUAL_, ACRO. legacy "advanced" switch configuration (will be removed soon)
	StabSwitch uint8 `yaml:"stab_switch"`// stabilize switch (only relevant for fixed wings, optional): _MANUAL, STABILIZED. legacy "advanced" switch configuration (will be removed soon)
	PosctlSwitch uint8 `yaml:"posctl_switch"`// position control 2 position switch (optional): _ALTCTL_, POSCTL. legacy "advanced" switch configuration (will be removed soon)
	SwitchChanges uint32 `yaml:"switch_changes"`// number of switch changes
}

// NewManualControlSwitches creates a new ManualControlSwitches with default values.
func NewManualControlSwitches() *ManualControlSwitches {
	self := ManualControlSwitches{}
	self.SetDefaults()
	return &self
}

func (t *ManualControlSwitches) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *ManualControlSwitches) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var ManualControlSwitchesTypeSupport types.MessageTypeSupport = _ManualControlSwitchesTypeSupport{}

type _ManualControlSwitchesTypeSupport struct{}

func (t _ManualControlSwitchesTypeSupport) New() types.Message {
	return NewManualControlSwitches()
}

func (t _ManualControlSwitchesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ManualControlSwitches
	return (unsafe.Pointer)(C.px4_msgs__msg__ManualControlSwitches__create())
}

func (t _ManualControlSwitchesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ManualControlSwitches__destroy((*C.px4_msgs__msg__ManualControlSwitches)(pointer_to_free))
}

func (t _ManualControlSwitchesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ManualControlSwitches)
	mem := (*C.px4_msgs__msg__ManualControlSwitches)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.mode_slot = C.uint8_t(m.ModeSlot)
	mem.arm_switch = C.uint8_t(m.ArmSwitch)
	mem.return_switch = C.uint8_t(m.ReturnSwitch)
	mem.loiter_switch = C.uint8_t(m.LoiterSwitch)
	mem.offboard_switch = C.uint8_t(m.OffboardSwitch)
	mem.kill_switch = C.uint8_t(m.KillSwitch)
	mem.gear_switch = C.uint8_t(m.GearSwitch)
	mem.transition_switch = C.uint8_t(m.TransitionSwitch)
	mem.mode_switch = C.uint8_t(m.ModeSwitch)
	mem.man_switch = C.uint8_t(m.ManSwitch)
	mem.acro_switch = C.uint8_t(m.AcroSwitch)
	mem.stab_switch = C.uint8_t(m.StabSwitch)
	mem.posctl_switch = C.uint8_t(m.PosctlSwitch)
	mem.switch_changes = C.uint32_t(m.SwitchChanges)
}

func (t _ManualControlSwitchesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ManualControlSwitches)
	mem := (*C.px4_msgs__msg__ManualControlSwitches)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.ModeSlot = uint8(mem.mode_slot)
	m.ArmSwitch = uint8(mem.arm_switch)
	m.ReturnSwitch = uint8(mem.return_switch)
	m.LoiterSwitch = uint8(mem.loiter_switch)
	m.OffboardSwitch = uint8(mem.offboard_switch)
	m.KillSwitch = uint8(mem.kill_switch)
	m.GearSwitch = uint8(mem.gear_switch)
	m.TransitionSwitch = uint8(mem.transition_switch)
	m.ModeSwitch = uint8(mem.mode_switch)
	m.ManSwitch = uint8(mem.man_switch)
	m.AcroSwitch = uint8(mem.acro_switch)
	m.StabSwitch = uint8(mem.stab_switch)
	m.PosctlSwitch = uint8(mem.posctl_switch)
	m.SwitchChanges = uint32(mem.switch_changes)
}

func (t _ManualControlSwitchesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ManualControlSwitches())
}

type CManualControlSwitches = C.px4_msgs__msg__ManualControlSwitches
type CManualControlSwitches__Sequence = C.px4_msgs__msg__ManualControlSwitches__Sequence

func ManualControlSwitches__Sequence_to_Go(goSlice *[]ManualControlSwitches, cSlice CManualControlSwitches__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ManualControlSwitches, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ManualControlSwitches__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ManualControlSwitches * uintptr(i)),
		))
		ManualControlSwitchesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ManualControlSwitches__Sequence_to_C(cSlice *CManualControlSwitches__Sequence, goSlice []ManualControlSwitches) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ManualControlSwitches)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ManualControlSwitches * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ManualControlSwitches)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ManualControlSwitches * uintptr(i)),
		))
		ManualControlSwitchesTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ManualControlSwitches__Array_to_Go(goSlice []ManualControlSwitches, cSlice []CManualControlSwitches) {
	for i := 0; i < len(cSlice); i++ {
		ManualControlSwitchesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ManualControlSwitches__Array_to_C(cSlice []CManualControlSwitches, goSlice []ManualControlSwitches) {
	for i := 0; i < len(goSlice); i++ {
		ManualControlSwitchesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
