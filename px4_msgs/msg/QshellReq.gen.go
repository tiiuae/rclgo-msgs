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

#include <px4_msgs/msg/qshell_req.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/QshellReq", QshellReqTypeSupport)
}
const (
	QshellReq_MAX_STRLEN uint32 = 100
)

// Do not create instances of this type directly. Always use NewQshellReq
// function instead.
type QshellReq struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Cmd [100]byte `yaml:"cmd"`
	Strlen uint32 `yaml:"strlen"`
	RequestSequence uint32 `yaml:"request_sequence"`
}

// NewQshellReq creates a new QshellReq with default values.
func NewQshellReq() *QshellReq {
	self := QshellReq{}
	self.SetDefaults()
	return &self
}

func (t *QshellReq) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *QshellReq) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var QshellReqTypeSupport types.MessageTypeSupport = _QshellReqTypeSupport{}

type _QshellReqTypeSupport struct{}

func (t _QshellReqTypeSupport) New() types.Message {
	return NewQshellReq()
}

func (t _QshellReqTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__QshellReq
	return (unsafe.Pointer)(C.px4_msgs__msg__QshellReq__create())
}

func (t _QshellReqTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__QshellReq__destroy((*C.px4_msgs__msg__QshellReq)(pointer_to_free))
}

func (t _QshellReqTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*QshellReq)
	mem := (*C.px4_msgs__msg__QshellReq)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_cmd := mem.cmd[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_cmd)), m.Cmd[:])
	mem.strlen = C.uint32_t(m.Strlen)
	mem.request_sequence = C.uint32_t(m.RequestSequence)
}

func (t _QshellReqTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*QshellReq)
	mem := (*C.px4_msgs__msg__QshellReq)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_cmd := mem.cmd[:]
	primitives.Char__Array_to_Go(m.Cmd[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_cmd)))
	m.Strlen = uint32(mem.strlen)
	m.RequestSequence = uint32(mem.request_sequence)
}

func (t _QshellReqTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__QshellReq())
}

type CQshellReq = C.px4_msgs__msg__QshellReq
type CQshellReq__Sequence = C.px4_msgs__msg__QshellReq__Sequence

func QshellReq__Sequence_to_Go(goSlice *[]QshellReq, cSlice CQshellReq__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]QshellReq, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__QshellReq__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__QshellReq * uintptr(i)),
		))
		QshellReqTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func QshellReq__Sequence_to_C(cSlice *CQshellReq__Sequence, goSlice []QshellReq) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__QshellReq)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__QshellReq * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__QshellReq)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__QshellReq * uintptr(i)),
		))
		QshellReqTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func QshellReq__Array_to_Go(goSlice []QshellReq, cSlice []CQshellReq) {
	for i := 0; i < len(cSlice); i++ {
		QshellReqTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func QshellReq__Array_to_C(cSlice []CQshellReq, goSlice []QshellReq) {
	for i := 0; i < len(goSlice); i++ {
		QshellReqTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
