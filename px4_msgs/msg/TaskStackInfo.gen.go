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

#include <px4_msgs/msg/task_stack_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/TaskStackInfo", TaskStackInfoTypeSupport)
}
const (
	TaskStackInfo_ORB_QUEUE_LENGTH uint8 = 2
)

// Do not create instances of this type directly. Always use NewTaskStackInfo
// function instead.
type TaskStackInfo struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	StackFree uint16 `yaml:"stack_free"`
	TaskName [24]byte `yaml:"task_name"`
}

// NewTaskStackInfo creates a new TaskStackInfo with default values.
func NewTaskStackInfo() *TaskStackInfo {
	self := TaskStackInfo{}
	self.SetDefaults()
	return &self
}

func (t *TaskStackInfo) Clone() *TaskStackInfo {
	c := &TaskStackInfo{}
	c.Timestamp = t.Timestamp
	c.StackFree = t.StackFree
	c.TaskName = t.TaskName
	return c
}

func (t *TaskStackInfo) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TaskStackInfo) SetDefaults() {
	t.Timestamp = 0
	t.StackFree = 0
	t.TaskName = [24]byte{}
}

// CloneTaskStackInfoSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTaskStackInfoSlice(dst, src []TaskStackInfo) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TaskStackInfoTypeSupport types.MessageTypeSupport = _TaskStackInfoTypeSupport{}

type _TaskStackInfoTypeSupport struct{}

func (t _TaskStackInfoTypeSupport) New() types.Message {
	return NewTaskStackInfo()
}

func (t _TaskStackInfoTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TaskStackInfo
	return (unsafe.Pointer)(C.px4_msgs__msg__TaskStackInfo__create())
}

func (t _TaskStackInfoTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TaskStackInfo__destroy((*C.px4_msgs__msg__TaskStackInfo)(pointer_to_free))
}

func (t _TaskStackInfoTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TaskStackInfo)
	mem := (*C.px4_msgs__msg__TaskStackInfo)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.stack_free = C.uint16_t(m.StackFree)
	cSlice_task_name := mem.task_name[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_task_name)), m.TaskName[:])
}

func (t _TaskStackInfoTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TaskStackInfo)
	mem := (*C.px4_msgs__msg__TaskStackInfo)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.StackFree = uint16(mem.stack_free)
	cSlice_task_name := mem.task_name[:]
	primitives.Char__Array_to_Go(m.TaskName[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_task_name)))
}

func (t _TaskStackInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TaskStackInfo())
}

type CTaskStackInfo = C.px4_msgs__msg__TaskStackInfo
type CTaskStackInfo__Sequence = C.px4_msgs__msg__TaskStackInfo__Sequence

func TaskStackInfo__Sequence_to_Go(goSlice *[]TaskStackInfo, cSlice CTaskStackInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TaskStackInfo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TaskStackInfo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TaskStackInfo * uintptr(i)),
		))
		TaskStackInfoTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TaskStackInfo__Sequence_to_C(cSlice *CTaskStackInfo__Sequence, goSlice []TaskStackInfo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TaskStackInfo)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TaskStackInfo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TaskStackInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TaskStackInfo * uintptr(i)),
		))
		TaskStackInfoTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TaskStackInfo__Array_to_Go(goSlice []TaskStackInfo, cSlice []CTaskStackInfo) {
	for i := 0; i < len(cSlice); i++ {
		TaskStackInfoTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TaskStackInfo__Array_to_C(cSlice []CTaskStackInfo, goSlice []TaskStackInfo) {
	for i := 0; i < len(goSlice); i++ {
		TaskStackInfoTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
