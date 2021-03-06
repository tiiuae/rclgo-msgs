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

#include <px4_msgs/msg/sensor_mag.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorMag", SensorMagTypeSupport)
}
const (
	SensorMag_ORB_QUEUE_LENGTH uint8 = 4
)

// Do not create instances of this type directly. Always use NewSensorMag
// function instead.
type SensorMag struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
	X float32 `yaml:"x"`// magnetic field in the FRD board frame X-axis in Gauss
	Y float32 `yaml:"y"`// magnetic field in the FRD board frame Y-axis in Gauss
	Z float32 `yaml:"z"`// magnetic field in the FRD board frame Z-axis in Gauss
	Temperature float32 `yaml:"temperature"`// temperature in degrees Celsius
	ErrorCount uint32 `yaml:"error_count"`
	IsExternal bool `yaml:"is_external"`// if true the mag is external (i.e. not built into the board)
}

// NewSensorMag creates a new SensorMag with default values.
func NewSensorMag() *SensorMag {
	self := SensorMag{}
	self.SetDefaults()
	return &self
}

func (t *SensorMag) Clone() *SensorMag {
	c := &SensorMag{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.DeviceId = t.DeviceId
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	c.Temperature = t.Temperature
	c.ErrorCount = t.ErrorCount
	c.IsExternal = t.IsExternal
	return c
}

func (t *SensorMag) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SensorMag) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.DeviceId = 0
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.Temperature = 0
	t.ErrorCount = 0
	t.IsExternal = false
}

// CloneSensorMagSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSensorMagSlice(dst, src []SensorMag) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SensorMagTypeSupport types.MessageTypeSupport = _SensorMagTypeSupport{}

type _SensorMagTypeSupport struct{}

func (t _SensorMagTypeSupport) New() types.Message {
	return NewSensorMag()
}

func (t _SensorMagTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorMag
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorMag__create())
}

func (t _SensorMagTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorMag__destroy((*C.px4_msgs__msg__SensorMag)(pointer_to_free))
}

func (t _SensorMagTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorMag)
	mem := (*C.px4_msgs__msg__SensorMag)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.device_id = C.uint32_t(m.DeviceId)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	mem.temperature = C.float(m.Temperature)
	mem.error_count = C.uint32_t(m.ErrorCount)
	mem.is_external = C.bool(m.IsExternal)
}

func (t _SensorMagTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorMag)
	mem := (*C.px4_msgs__msg__SensorMag)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.DeviceId = uint32(mem.device_id)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	m.Temperature = float32(mem.temperature)
	m.ErrorCount = uint32(mem.error_count)
	m.IsExternal = bool(mem.is_external)
}

func (t _SensorMagTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorMag())
}

type CSensorMag = C.px4_msgs__msg__SensorMag
type CSensorMag__Sequence = C.px4_msgs__msg__SensorMag__Sequence

func SensorMag__Sequence_to_Go(goSlice *[]SensorMag, cSlice CSensorMag__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorMag, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorMag__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorMag * uintptr(i)),
		))
		SensorMagTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorMag__Sequence_to_C(cSlice *CSensorMag__Sequence, goSlice []SensorMag) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorMag)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorMag * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorMag)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorMag * uintptr(i)),
		))
		SensorMagTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorMag__Array_to_Go(goSlice []SensorMag, cSlice []CSensorMag) {
	for i := 0; i < len(cSlice); i++ {
		SensorMagTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorMag__Array_to_C(cSlice []CSensorMag, goSlice []SensorMag) {
	for i := 0; i < len(goSlice); i++ {
		SensorMagTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
