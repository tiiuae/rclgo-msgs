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

#include <px4_msgs/msg/ping.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/Ping", PingTypeSupport)
}

// Do not create instances of this type directly. Always use NewPing
// function instead.
type Ping struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	PingTime uint64 `yaml:"ping_time"`// Timestamp of the ping packet
	PingSequence uint32 `yaml:"ping_sequence"`// Sequence number of the ping packet
	DroppedPackets uint32 `yaml:"dropped_packets"`// Number of dropped ping packets
	RttMs float32 `yaml:"rtt_ms"`// Round trip time (in ms)
	SystemId uint8 `yaml:"system_id"`// System ID of the remote system
	ComponentId uint8 `yaml:"component_id"`// Component ID of the remote system
}

// NewPing creates a new Ping with default values.
func NewPing() *Ping {
	self := Ping{}
	self.SetDefaults()
	return &self
}

func (t *Ping) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Ping) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var PingTypeSupport types.MessageTypeSupport = _PingTypeSupport{}

type _PingTypeSupport struct{}

func (t _PingTypeSupport) New() types.Message {
	return NewPing()
}

func (t _PingTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Ping
	return (unsafe.Pointer)(C.px4_msgs__msg__Ping__create())
}

func (t _PingTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Ping__destroy((*C.px4_msgs__msg__Ping)(pointer_to_free))
}

func (t _PingTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Ping)
	mem := (*C.px4_msgs__msg__Ping)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.ping_time = C.uint64_t(m.PingTime)
	mem.ping_sequence = C.uint32_t(m.PingSequence)
	mem.dropped_packets = C.uint32_t(m.DroppedPackets)
	mem.rtt_ms = C.float(m.RttMs)
	mem.system_id = C.uint8_t(m.SystemId)
	mem.component_id = C.uint8_t(m.ComponentId)
}

func (t _PingTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Ping)
	mem := (*C.px4_msgs__msg__Ping)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.PingTime = uint64(mem.ping_time)
	m.PingSequence = uint32(mem.ping_sequence)
	m.DroppedPackets = uint32(mem.dropped_packets)
	m.RttMs = float32(mem.rtt_ms)
	m.SystemId = uint8(mem.system_id)
	m.ComponentId = uint8(mem.component_id)
}

func (t _PingTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Ping())
}

type CPing = C.px4_msgs__msg__Ping
type CPing__Sequence = C.px4_msgs__msg__Ping__Sequence

func Ping__Sequence_to_Go(goSlice *[]Ping, cSlice CPing__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Ping, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Ping__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Ping * uintptr(i)),
		))
		PingTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Ping__Sequence_to_C(cSlice *CPing__Sequence, goSlice []Ping) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Ping)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Ping * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Ping)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Ping * uintptr(i)),
		))
		PingTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Ping__Array_to_Go(goSlice []Ping, cSlice []CPing) {
	for i := 0; i < len(cSlice); i++ {
		PingTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Ping__Array_to_C(cSlice []CPing, goSlice []Ping) {
	for i := 0; i < len(goSlice); i++ {
		PingTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
