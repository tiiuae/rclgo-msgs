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

#include <px4_msgs/msg/event.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/Event", EventTypeSupport)
}
const (
	Event_ORB_QUEUE_LENGTH uint8 = 8
)

// Do not create instances of this type directly. Always use NewEvent
// function instead.
type Event struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). Events interface
	Id uint32 `yaml:"id"`// Event ID
	EventSequence uint16 `yaml:"event_sequence"`// Event sequence number
	Arguments [25]uint8 `yaml:"arguments"`// (optional) arguments, depend on event id
	LogLevels uint8 `yaml:"log_levels"`// Log levels: 4 bits MSB: internal, 4 bits LSB: external
}

// NewEvent creates a new Event with default values.
func NewEvent() *Event {
	self := Event{}
	self.SetDefaults()
	return &self
}

func (t *Event) Clone() *Event {
	c := &Event{}
	c.Timestamp = t.Timestamp
	c.Id = t.Id
	c.EventSequence = t.EventSequence
	c.Arguments = t.Arguments
	c.LogLevels = t.LogLevels
	return c
}

func (t *Event) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Event) SetDefaults() {
	t.Timestamp = 0
	t.Id = 0
	t.EventSequence = 0
	t.Arguments = [25]uint8{}
	t.LogLevels = 0
}

// CloneEventSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEventSlice(dst, src []Event) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EventTypeSupport types.MessageTypeSupport = _EventTypeSupport{}

type _EventTypeSupport struct{}

func (t _EventTypeSupport) New() types.Message {
	return NewEvent()
}

func (t _EventTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Event
	return (unsafe.Pointer)(C.px4_msgs__msg__Event__create())
}

func (t _EventTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Event__destroy((*C.px4_msgs__msg__Event)(pointer_to_free))
}

func (t _EventTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Event)
	mem := (*C.px4_msgs__msg__Event)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.id = C.uint32_t(m.Id)
	mem.event_sequence = C.uint16_t(m.EventSequence)
	cSlice_arguments := mem.arguments[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_arguments)), m.Arguments[:])
	mem.log_levels = C.uint8_t(m.LogLevels)
}

func (t _EventTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Event)
	mem := (*C.px4_msgs__msg__Event)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Id = uint32(mem.id)
	m.EventSequence = uint16(mem.event_sequence)
	cSlice_arguments := mem.arguments[:]
	primitives.Uint8__Array_to_Go(m.Arguments[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_arguments)))
	m.LogLevels = uint8(mem.log_levels)
}

func (t _EventTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Event())
}

type CEvent = C.px4_msgs__msg__Event
type CEvent__Sequence = C.px4_msgs__msg__Event__Sequence

func Event__Sequence_to_Go(goSlice *[]Event, cSlice CEvent__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Event, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Event__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Event * uintptr(i)),
		))
		EventTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Event__Sequence_to_C(cSlice *CEvent__Sequence, goSlice []Event) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Event)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Event * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Event)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Event * uintptr(i)),
		))
		EventTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Event__Array_to_Go(goSlice []Event, cSlice []CEvent) {
	for i := 0; i < len(cSlice); i++ {
		EventTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Event__Array_to_C(cSlice []CEvent, goSlice []Event) {
	for i := 0; i < len(goSlice); i++ {
		EventTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}