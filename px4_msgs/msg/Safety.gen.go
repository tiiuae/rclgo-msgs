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

#include <px4_msgs/msg/safety.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/Safety", SafetyTypeSupport)
}

// Do not create instances of this type directly. Always use NewSafety
// function instead.
type Safety struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	SafetySwitchAvailable bool `yaml:"safety_switch_available"`// Set to true if a safety switch is connected
	SafetyOff bool `yaml:"safety_off"`// Set to true if safety is off
}

// NewSafety creates a new Safety with default values.
func NewSafety() *Safety {
	self := Safety{}
	self.SetDefaults()
	return &self
}

func (t *Safety) Clone() *Safety {
	c := &Safety{}
	c.Timestamp = t.Timestamp
	c.SafetySwitchAvailable = t.SafetySwitchAvailable
	c.SafetyOff = t.SafetyOff
	return c
}

func (t *Safety) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Safety) SetDefaults() {
	t.Timestamp = 0
	t.SafetySwitchAvailable = false
	t.SafetyOff = false
}

// CloneSafetySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSafetySlice(dst, src []Safety) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SafetyTypeSupport types.MessageTypeSupport = _SafetyTypeSupport{}

type _SafetyTypeSupport struct{}

func (t _SafetyTypeSupport) New() types.Message {
	return NewSafety()
}

func (t _SafetyTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Safety
	return (unsafe.Pointer)(C.px4_msgs__msg__Safety__create())
}

func (t _SafetyTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Safety__destroy((*C.px4_msgs__msg__Safety)(pointer_to_free))
}

func (t _SafetyTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Safety)
	mem := (*C.px4_msgs__msg__Safety)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.safety_switch_available = C.bool(m.SafetySwitchAvailable)
	mem.safety_off = C.bool(m.SafetyOff)
}

func (t _SafetyTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Safety)
	mem := (*C.px4_msgs__msg__Safety)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.SafetySwitchAvailable = bool(mem.safety_switch_available)
	m.SafetyOff = bool(mem.safety_off)
}

func (t _SafetyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Safety())
}

type CSafety = C.px4_msgs__msg__Safety
type CSafety__Sequence = C.px4_msgs__msg__Safety__Sequence

func Safety__Sequence_to_Go(goSlice *[]Safety, cSlice CSafety__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Safety, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Safety__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Safety * uintptr(i)),
		))
		SafetyTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Safety__Sequence_to_C(cSlice *CSafety__Sequence, goSlice []Safety) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Safety)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Safety * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Safety)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Safety * uintptr(i)),
		))
		SafetyTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Safety__Array_to_Go(goSlice []Safety, cSlice []CSafety) {
	for i := 0; i < len(cSlice); i++ {
		SafetyTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Safety__Array_to_C(cSlice []CSafety, goSlice []Safety) {
	for i := 0; i < len(goSlice); i++ {
		SafetyTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
