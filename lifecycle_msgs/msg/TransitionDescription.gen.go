/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package lifecycle_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <lifecycle_msgs/msg/transition_description.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/TransitionDescription", TransitionDescriptionTypeSupport)
}

// Do not create instances of this type directly. Always use NewTransitionDescription
// function instead.
type TransitionDescription struct {
	Transition Transition `yaml:"transition"`// The transition id and label of this description.
	StartState State `yaml:"start_state"`// The current state from which this transition transitions.
	GoalState State `yaml:"goal_state"`// The desired target state of this transition.
}

// NewTransitionDescription creates a new TransitionDescription with default values.
func NewTransitionDescription() *TransitionDescription {
	self := TransitionDescription{}
	self.SetDefaults()
	return &self
}

func (t *TransitionDescription) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *TransitionDescription) SetDefaults() {
	t.Transition.SetDefaults()
	t.StartState.SetDefaults()
	t.GoalState.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var TransitionDescriptionTypeSupport types.MessageTypeSupport = _TransitionDescriptionTypeSupport{}

type _TransitionDescriptionTypeSupport struct{}

func (t _TransitionDescriptionTypeSupport) New() types.Message {
	return NewTransitionDescription()
}

func (t _TransitionDescriptionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__msg__TransitionDescription
	return (unsafe.Pointer)(C.lifecycle_msgs__msg__TransitionDescription__create())
}

func (t _TransitionDescriptionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__msg__TransitionDescription__destroy((*C.lifecycle_msgs__msg__TransitionDescription)(pointer_to_free))
}

func (t _TransitionDescriptionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TransitionDescription)
	mem := (*C.lifecycle_msgs__msg__TransitionDescription)(dst)
	TransitionTypeSupport.AsCStruct(unsafe.Pointer(&mem.transition), &m.Transition)
	StateTypeSupport.AsCStruct(unsafe.Pointer(&mem.start_state), &m.StartState)
	StateTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_state), &m.GoalState)
}

func (t _TransitionDescriptionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TransitionDescription)
	mem := (*C.lifecycle_msgs__msg__TransitionDescription)(ros2_message_buffer)
	TransitionTypeSupport.AsGoStruct(&m.Transition, unsafe.Pointer(&mem.transition))
	StateTypeSupport.AsGoStruct(&m.StartState, unsafe.Pointer(&mem.start_state))
	StateTypeSupport.AsGoStruct(&m.GoalState, unsafe.Pointer(&mem.goal_state))
}

func (t _TransitionDescriptionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__msg__TransitionDescription())
}

type CTransitionDescription = C.lifecycle_msgs__msg__TransitionDescription
type CTransitionDescription__Sequence = C.lifecycle_msgs__msg__TransitionDescription__Sequence

func TransitionDescription__Sequence_to_Go(goSlice *[]TransitionDescription, cSlice CTransitionDescription__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TransitionDescription, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.lifecycle_msgs__msg__TransitionDescription__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__msg__TransitionDescription * uintptr(i)),
		))
		TransitionDescriptionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TransitionDescription__Sequence_to_C(cSlice *CTransitionDescription__Sequence, goSlice []TransitionDescription) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.lifecycle_msgs__msg__TransitionDescription)(C.malloc((C.size_t)(C.sizeof_struct_lifecycle_msgs__msg__TransitionDescription * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.lifecycle_msgs__msg__TransitionDescription)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__msg__TransitionDescription * uintptr(i)),
		))
		TransitionDescriptionTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TransitionDescription__Array_to_Go(goSlice []TransitionDescription, cSlice []CTransitionDescription) {
	for i := 0; i < len(cSlice); i++ {
		TransitionDescriptionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TransitionDescription__Array_to_C(cSlice []CTransitionDescription, goSlice []TransitionDescription) {
	for i := 0; i < len(goSlice); i++ {
		TransitionDescriptionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
