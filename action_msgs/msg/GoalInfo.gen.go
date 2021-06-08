/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package action_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo-msgs/builtin_interfaces/msg"
	unique_identifier_msgs_msg "github.com/tiiuae/rclgo-msgs/unique_identifier_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -laction_msgs__rosidl_typesupport_c -laction_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lunique_identifier_msgs__rosidl_typesupport_c -lunique_identifier_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <action_msgs/msg/goal_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("action_msgs/GoalInfo", GoalInfoTypeSupport)
}

// Do not create instances of this type directly. Always use NewGoalInfo
// function instead.
type GoalInfo struct {
	GoalId unique_identifier_msgs_msg.UUID `yaml:"goal_id"`// Goal ID
	Stamp builtin_interfaces_msg.Time `yaml:"stamp"`// Time when the goal was accepted
}

// NewGoalInfo creates a new GoalInfo with default values.
func NewGoalInfo() *GoalInfo {
	self := GoalInfo{}
	self.SetDefaults()
	return &self
}

func (t *GoalInfo) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *GoalInfo) SetDefaults() {
	t.GoalId.SetDefaults()
	t.Stamp.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var GoalInfoTypeSupport types.MessageTypeSupport = _GoalInfoTypeSupport{}

type _GoalInfoTypeSupport struct{}

func (t _GoalInfoTypeSupport) New() types.Message {
	return NewGoalInfo()
}

func (t _GoalInfoTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.action_msgs__msg__GoalInfo
	return (unsafe.Pointer)(C.action_msgs__msg__GoalInfo__create())
}

func (t _GoalInfoTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.action_msgs__msg__GoalInfo__destroy((*C.action_msgs__msg__GoalInfo)(pointer_to_free))
}

func (t _GoalInfoTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GoalInfo)
	mem := (*C.action_msgs__msg__GoalInfo)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_id), &m.GoalId)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.stamp), &m.Stamp)
}

func (t _GoalInfoTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GoalInfo)
	mem := (*C.action_msgs__msg__GoalInfo)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalId, unsafe.Pointer(&mem.goal_id))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Stamp, unsafe.Pointer(&mem.stamp))
}

func (t _GoalInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__action_msgs__msg__GoalInfo())
}

type CGoalInfo = C.action_msgs__msg__GoalInfo
type CGoalInfo__Sequence = C.action_msgs__msg__GoalInfo__Sequence

func GoalInfo__Sequence_to_Go(goSlice *[]GoalInfo, cSlice CGoalInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalInfo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.action_msgs__msg__GoalInfo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_action_msgs__msg__GoalInfo * uintptr(i)),
		))
		GoalInfoTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GoalInfo__Sequence_to_C(cSlice *CGoalInfo__Sequence, goSlice []GoalInfo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.action_msgs__msg__GoalInfo)(C.malloc((C.size_t)(C.sizeof_struct_action_msgs__msg__GoalInfo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.action_msgs__msg__GoalInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_action_msgs__msg__GoalInfo * uintptr(i)),
		))
		GoalInfoTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GoalInfo__Array_to_Go(goSlice []GoalInfo, cSlice []CGoalInfo) {
	for i := 0; i < len(cSlice); i++ {
		GoalInfoTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GoalInfo__Array_to_C(cSlice []CGoalInfo, goSlice []GoalInfo) {
	for i := 0; i < len(goSlice); i++ {
		GoalInfoTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}