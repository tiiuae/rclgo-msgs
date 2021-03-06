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

#include <px4_msgs/msg/ulog_stream_ack.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/UlogStreamAck", UlogStreamAckTypeSupport)
}
const (
	UlogStreamAck_ACK_TIMEOUT int32 = 50// timeout waiting for an ack until we retry to send the message [ms]
	UlogStreamAck_ACK_MAX_TRIES int32 = 50// maximum amount of tries to (re-)send a message, each time waiting ACK_TIMEOUT ms
)

// Do not create instances of this type directly. Always use NewUlogStreamAck
// function instead.
type UlogStreamAck struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	MsgSequence uint16 `yaml:"msg_sequence"`
}

// NewUlogStreamAck creates a new UlogStreamAck with default values.
func NewUlogStreamAck() *UlogStreamAck {
	self := UlogStreamAck{}
	self.SetDefaults()
	return &self
}

func (t *UlogStreamAck) Clone() *UlogStreamAck {
	c := &UlogStreamAck{}
	c.Timestamp = t.Timestamp
	c.MsgSequence = t.MsgSequence
	return c
}

func (t *UlogStreamAck) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UlogStreamAck) SetDefaults() {
	t.Timestamp = 0
	t.MsgSequence = 0
}

// CloneUlogStreamAckSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUlogStreamAckSlice(dst, src []UlogStreamAck) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UlogStreamAckTypeSupport types.MessageTypeSupport = _UlogStreamAckTypeSupport{}

type _UlogStreamAckTypeSupport struct{}

func (t _UlogStreamAckTypeSupport) New() types.Message {
	return NewUlogStreamAck()
}

func (t _UlogStreamAckTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__UlogStreamAck
	return (unsafe.Pointer)(C.px4_msgs__msg__UlogStreamAck__create())
}

func (t _UlogStreamAckTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__UlogStreamAck__destroy((*C.px4_msgs__msg__UlogStreamAck)(pointer_to_free))
}

func (t _UlogStreamAckTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UlogStreamAck)
	mem := (*C.px4_msgs__msg__UlogStreamAck)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.msg_sequence = C.uint16_t(m.MsgSequence)
}

func (t _UlogStreamAckTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UlogStreamAck)
	mem := (*C.px4_msgs__msg__UlogStreamAck)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.MsgSequence = uint16(mem.msg_sequence)
}

func (t _UlogStreamAckTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__UlogStreamAck())
}

type CUlogStreamAck = C.px4_msgs__msg__UlogStreamAck
type CUlogStreamAck__Sequence = C.px4_msgs__msg__UlogStreamAck__Sequence

func UlogStreamAck__Sequence_to_Go(goSlice *[]UlogStreamAck, cSlice CUlogStreamAck__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UlogStreamAck, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__UlogStreamAck__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UlogStreamAck * uintptr(i)),
		))
		UlogStreamAckTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UlogStreamAck__Sequence_to_C(cSlice *CUlogStreamAck__Sequence, goSlice []UlogStreamAck) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__UlogStreamAck)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__UlogStreamAck * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__UlogStreamAck)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UlogStreamAck * uintptr(i)),
		))
		UlogStreamAckTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UlogStreamAck__Array_to_Go(goSlice []UlogStreamAck, cSlice []CUlogStreamAck) {
	for i := 0; i < len(cSlice); i++ {
		UlogStreamAckTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UlogStreamAck__Array_to_C(cSlice []CUlogStreamAck, goSlice []UlogStreamAck) {
	for i := 0; i < len(goSlice); i++ {
		UlogStreamAckTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
