/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/multi_nested.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/MultiNested", MultiNestedTypeSupport)
}

// Do not create instances of this type directly. Always use NewMultiNested
// function instead.
type MultiNested struct {
	ArrayOfArrays [3]Arrays `yaml:"array_of_arrays"`// Mulitple levels of nested messages
	ArrayOfBoundedSequences [3]BoundedSequences `yaml:"array_of_bounded_sequences"`// Mulitple levels of nested messages
	ArrayOfUnboundedSequences [3]UnboundedSequences `yaml:"array_of_unbounded_sequences"`// Mulitple levels of nested messages
	BoundedSequenceOfArrays []Arrays `yaml:"bounded_sequence_of_arrays"`// Mulitple levels of nested messages
	BoundedSequenceOfBoundedSequences []BoundedSequences `yaml:"bounded_sequence_of_bounded_sequences"`// Mulitple levels of nested messages
	BoundedSequenceOfUnboundedSequences []UnboundedSequences `yaml:"bounded_sequence_of_unbounded_sequences"`// Mulitple levels of nested messages
	UnboundedSequenceOfArrays []Arrays `yaml:"unbounded_sequence_of_arrays"`// Mulitple levels of nested messages
	UnboundedSequenceOfBoundedSequences []BoundedSequences `yaml:"unbounded_sequence_of_bounded_sequences"`// Mulitple levels of nested messages
	UnboundedSequenceOfUnboundedSequences []UnboundedSequences `yaml:"unbounded_sequence_of_unbounded_sequences"`// Mulitple levels of nested messages
}

// NewMultiNested creates a new MultiNested with default values.
func NewMultiNested() *MultiNested {
	self := MultiNested{}
	self.SetDefaults()
	return &self
}

func (t *MultiNested) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *MultiNested) SetDefaults() {
	t.ArrayOfArrays[0].SetDefaults()
	t.ArrayOfArrays[1].SetDefaults()
	t.ArrayOfArrays[2].SetDefaults()
	t.ArrayOfBoundedSequences[0].SetDefaults()
	t.ArrayOfBoundedSequences[1].SetDefaults()
	t.ArrayOfBoundedSequences[2].SetDefaults()
	t.ArrayOfUnboundedSequences[0].SetDefaults()
	t.ArrayOfUnboundedSequences[1].SetDefaults()
	t.ArrayOfUnboundedSequences[2].SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var MultiNestedTypeSupport types.MessageTypeSupport = _MultiNestedTypeSupport{}

type _MultiNestedTypeSupport struct{}

func (t _MultiNestedTypeSupport) New() types.Message {
	return NewMultiNested()
}

func (t _MultiNestedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__MultiNested
	return (unsafe.Pointer)(C.test_msgs__msg__MultiNested__create())
}

func (t _MultiNestedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__MultiNested__destroy((*C.test_msgs__msg__MultiNested)(pointer_to_free))
}

func (t _MultiNestedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiNested)
	mem := (*C.test_msgs__msg__MultiNested)(dst)
	Arrays__Array_to_C(mem.array_of_arrays[:], m.ArrayOfArrays[:])
	BoundedSequences__Array_to_C(mem.array_of_bounded_sequences[:], m.ArrayOfBoundedSequences[:])
	UnboundedSequences__Array_to_C(mem.array_of_unbounded_sequences[:], m.ArrayOfUnboundedSequences[:])
	Arrays__Sequence_to_C(&mem.bounded_sequence_of_arrays, m.BoundedSequenceOfArrays)
	BoundedSequences__Sequence_to_C(&mem.bounded_sequence_of_bounded_sequences, m.BoundedSequenceOfBoundedSequences)
	UnboundedSequences__Sequence_to_C(&mem.bounded_sequence_of_unbounded_sequences, m.BoundedSequenceOfUnboundedSequences)
	Arrays__Sequence_to_C(&mem.unbounded_sequence_of_arrays, m.UnboundedSequenceOfArrays)
	BoundedSequences__Sequence_to_C(&mem.unbounded_sequence_of_bounded_sequences, m.UnboundedSequenceOfBoundedSequences)
	UnboundedSequences__Sequence_to_C(&mem.unbounded_sequence_of_unbounded_sequences, m.UnboundedSequenceOfUnboundedSequences)
}

func (t _MultiNestedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiNested)
	mem := (*C.test_msgs__msg__MultiNested)(ros2_message_buffer)
	Arrays__Array_to_Go(m.ArrayOfArrays[:], mem.array_of_arrays[:])
	BoundedSequences__Array_to_Go(m.ArrayOfBoundedSequences[:], mem.array_of_bounded_sequences[:])
	UnboundedSequences__Array_to_Go(m.ArrayOfUnboundedSequences[:], mem.array_of_unbounded_sequences[:])
	Arrays__Sequence_to_Go(&m.BoundedSequenceOfArrays, mem.bounded_sequence_of_arrays)
	BoundedSequences__Sequence_to_Go(&m.BoundedSequenceOfBoundedSequences, mem.bounded_sequence_of_bounded_sequences)
	UnboundedSequences__Sequence_to_Go(&m.BoundedSequenceOfUnboundedSequences, mem.bounded_sequence_of_unbounded_sequences)
	Arrays__Sequence_to_Go(&m.UnboundedSequenceOfArrays, mem.unbounded_sequence_of_arrays)
	BoundedSequences__Sequence_to_Go(&m.UnboundedSequenceOfBoundedSequences, mem.unbounded_sequence_of_bounded_sequences)
	UnboundedSequences__Sequence_to_Go(&m.UnboundedSequenceOfUnboundedSequences, mem.unbounded_sequence_of_unbounded_sequences)
}

func (t _MultiNestedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__MultiNested())
}

type CMultiNested = C.test_msgs__msg__MultiNested
type CMultiNested__Sequence = C.test_msgs__msg__MultiNested__Sequence

func MultiNested__Sequence_to_Go(goSlice *[]MultiNested, cSlice CMultiNested__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiNested, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__MultiNested__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__MultiNested * uintptr(i)),
		))
		MultiNestedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MultiNested__Sequence_to_C(cSlice *CMultiNested__Sequence, goSlice []MultiNested) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__MultiNested)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__MultiNested * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__MultiNested)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__MultiNested * uintptr(i)),
		))
		MultiNestedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MultiNested__Array_to_Go(goSlice []MultiNested, cSlice []CMultiNested) {
	for i := 0; i < len(cSlice); i++ {
		MultiNestedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiNested__Array_to_C(cSlice []CMultiNested, goSlice []MultiNested) {
	for i := 0; i < len(goSlice); i++ {
		MultiNestedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
