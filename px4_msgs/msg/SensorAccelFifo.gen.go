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
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/sensor_accel_fifo.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorAccelFifo", SensorAccelFifoTypeSupport)
}

// Do not create instances of this type directly. Always use NewSensorAccelFifo
// function instead.
type SensorAccelFifo struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
	Dt float32 `yaml:"dt"`// delta time between samples (microseconds)
	Scale float32 `yaml:"scale"`
	Samples uint8 `yaml:"samples"`// number of valid samples
	X [32]int16 `yaml:"x"`// acceleration in the FRD board frame X-axis in m/s^2
	Y [32]int16 `yaml:"y"`// acceleration in the FRD board frame Y-axis in m/s^2
	Z [32]int16 `yaml:"z"`// acceleration in the FRD board frame Z-axis in m/s^2
}

// NewSensorAccelFifo creates a new SensorAccelFifo with default values.
func NewSensorAccelFifo() *SensorAccelFifo {
	self := SensorAccelFifo{}
	self.SetDefaults()
	return &self
}

func (t *SensorAccelFifo) Clone() *SensorAccelFifo {
	c := &SensorAccelFifo{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.DeviceId = t.DeviceId
	c.Dt = t.Dt
	c.Scale = t.Scale
	c.Samples = t.Samples
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	return c
}

func (t *SensorAccelFifo) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SensorAccelFifo) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.DeviceId = 0
	t.Dt = 0
	t.Scale = 0
	t.Samples = 0
	t.X = [32]int16{}
	t.Y = [32]int16{}
	t.Z = [32]int16{}
}

// CloneSensorAccelFifoSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSensorAccelFifoSlice(dst, src []SensorAccelFifo) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SensorAccelFifoTypeSupport types.MessageTypeSupport = _SensorAccelFifoTypeSupport{}

type _SensorAccelFifoTypeSupport struct{}

func (t _SensorAccelFifoTypeSupport) New() types.Message {
	return NewSensorAccelFifo()
}

func (t _SensorAccelFifoTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorAccelFifo
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorAccelFifo__create())
}

func (t _SensorAccelFifoTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorAccelFifo__destroy((*C.px4_msgs__msg__SensorAccelFifo)(pointer_to_free))
}

func (t _SensorAccelFifoTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorAccelFifo)
	mem := (*C.px4_msgs__msg__SensorAccelFifo)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.device_id = C.uint32_t(m.DeviceId)
	mem.dt = C.float(m.Dt)
	mem.scale = C.float(m.Scale)
	mem.samples = C.uint8_t(m.Samples)
	cSlice_x := mem.x[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_x)), m.X[:])
	cSlice_y := mem.y[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_y)), m.Y[:])
	cSlice_z := mem.z[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_z)), m.Z[:])
}

func (t _SensorAccelFifoTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorAccelFifo)
	mem := (*C.px4_msgs__msg__SensorAccelFifo)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.DeviceId = uint32(mem.device_id)
	m.Dt = float32(mem.dt)
	m.Scale = float32(mem.scale)
	m.Samples = uint8(mem.samples)
	cSlice_x := mem.x[:]
	primitives.Int16__Array_to_Go(m.X[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_x)))
	cSlice_y := mem.y[:]
	primitives.Int16__Array_to_Go(m.Y[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_y)))
	cSlice_z := mem.z[:]
	primitives.Int16__Array_to_Go(m.Z[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_z)))
}

func (t _SensorAccelFifoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorAccelFifo())
}

type CSensorAccelFifo = C.px4_msgs__msg__SensorAccelFifo
type CSensorAccelFifo__Sequence = C.px4_msgs__msg__SensorAccelFifo__Sequence

func SensorAccelFifo__Sequence_to_Go(goSlice *[]SensorAccelFifo, cSlice CSensorAccelFifo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorAccelFifo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorAccelFifo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorAccelFifo * uintptr(i)),
		))
		SensorAccelFifoTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorAccelFifo__Sequence_to_C(cSlice *CSensorAccelFifo__Sequence, goSlice []SensorAccelFifo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorAccelFifo)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorAccelFifo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorAccelFifo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorAccelFifo * uintptr(i)),
		))
		SensorAccelFifoTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorAccelFifo__Array_to_Go(goSlice []SensorAccelFifo, cSlice []CSensorAccelFifo) {
	for i := 0; i < len(cSlice); i++ {
		SensorAccelFifoTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorAccelFifo__Array_to_C(cSlice []CSensorAccelFifo, goSlice []SensorAccelFifo) {
	for i := 0; i < len(goSlice); i++ {
		SensorAccelFifoTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
