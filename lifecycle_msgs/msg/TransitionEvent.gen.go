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

#include <lifecycle_msgs/msg/transition_event.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/TransitionEvent", TransitionEventTypeSupport)
}

// Do not create instances of this type directly. Always use NewTransitionEvent
// function instead.
type TransitionEvent struct {
	Timestamp uint64 `yaml:"timestamp"`// The time point at which this event occurred.
	Transition Transition `yaml:"transition"`// The id and label of this transition event.
	StartState State `yaml:"start_state"`// The starting state from which this event transitioned.
	GoalState State `yaml:"goal_state"`// The end state of this transition event.
}

// NewTransitionEvent creates a new TransitionEvent with default values.
func NewTransitionEvent() *TransitionEvent {
	self := TransitionEvent{}
	self.SetDefaults()
	return &self
}

func (t *TransitionEvent) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *TransitionEvent) SetDefaults() {
	t.Transition.SetDefaults()
	t.StartState.SetDefaults()
	t.GoalState.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var TransitionEventTypeSupport types.MessageTypeSupport = _TransitionEventTypeSupport{}

type _TransitionEventTypeSupport struct{}

func (t _TransitionEventTypeSupport) New() types.Message {
	return NewTransitionEvent()
}

func (t _TransitionEventTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__msg__TransitionEvent
	return (unsafe.Pointer)(C.lifecycle_msgs__msg__TransitionEvent__create())
}

func (t _TransitionEventTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__msg__TransitionEvent__destroy((*C.lifecycle_msgs__msg__TransitionEvent)(pointer_to_free))
}

func (t _TransitionEventTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TransitionEvent)
	mem := (*C.lifecycle_msgs__msg__TransitionEvent)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	TransitionTypeSupport.AsCStruct(unsafe.Pointer(&mem.transition), &m.Transition)
	StateTypeSupport.AsCStruct(unsafe.Pointer(&mem.start_state), &m.StartState)
	StateTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_state), &m.GoalState)
}

func (t _TransitionEventTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TransitionEvent)
	mem := (*C.lifecycle_msgs__msg__TransitionEvent)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	TransitionTypeSupport.AsGoStruct(&m.Transition, unsafe.Pointer(&mem.transition))
	StateTypeSupport.AsGoStruct(&m.StartState, unsafe.Pointer(&mem.start_state))
	StateTypeSupport.AsGoStruct(&m.GoalState, unsafe.Pointer(&mem.goal_state))
}

func (t _TransitionEventTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__msg__TransitionEvent())
}

type CTransitionEvent = C.lifecycle_msgs__msg__TransitionEvent
type CTransitionEvent__Sequence = C.lifecycle_msgs__msg__TransitionEvent__Sequence

func TransitionEvent__Sequence_to_Go(goSlice *[]TransitionEvent, cSlice CTransitionEvent__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TransitionEvent, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.lifecycle_msgs__msg__TransitionEvent__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__msg__TransitionEvent * uintptr(i)),
		))
		TransitionEventTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TransitionEvent__Sequence_to_C(cSlice *CTransitionEvent__Sequence, goSlice []TransitionEvent) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.lifecycle_msgs__msg__TransitionEvent)(C.malloc((C.size_t)(C.sizeof_struct_lifecycle_msgs__msg__TransitionEvent * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.lifecycle_msgs__msg__TransitionEvent)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__msg__TransitionEvent * uintptr(i)),
		))
		TransitionEventTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TransitionEvent__Array_to_Go(goSlice []TransitionEvent, cSlice []CTransitionEvent) {
	for i := 0; i < len(cSlice); i++ {
		TransitionEventTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TransitionEvent__Array_to_C(cSlice []CTransitionEvent, goSlice []TransitionEvent) {
	for i := 0; i < len(goSlice); i++ {
		TransitionEventTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
