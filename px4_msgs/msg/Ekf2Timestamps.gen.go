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

#include <px4_msgs/msg/ekf2_timestamps.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/Ekf2Timestamps", Ekf2TimestampsTypeSupport)
}
const (
	Ekf2Timestamps_RELATIVE_TIMESTAMP_INVALID int16 = 32767// (0x7fff) If one of the relative timestamps
)

// Do not create instances of this type directly. Always use NewEkf2Timestamps
// function instead.
type Ekf2Timestamps struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	AirspeedTimestampRel int16 `yaml:"airspeed_timestamp_rel"`
	DistanceSensorTimestampRel int16 `yaml:"distance_sensor_timestamp_rel"`
	OpticalFlowTimestampRel int16 `yaml:"optical_flow_timestamp_rel"`
	VehicleAirDataTimestampRel int16 `yaml:"vehicle_air_data_timestamp_rel"`
	VehicleMagnetometerTimestampRel int16 `yaml:"vehicle_magnetometer_timestamp_rel"`
	VisualOdometryTimestampRel int16 `yaml:"visual_odometry_timestamp_rel"`
}

// NewEkf2Timestamps creates a new Ekf2Timestamps with default values.
func NewEkf2Timestamps() *Ekf2Timestamps {
	self := Ekf2Timestamps{}
	self.SetDefaults()
	return &self
}

func (t *Ekf2Timestamps) Clone() *Ekf2Timestamps {
	c := &Ekf2Timestamps{}
	c.Timestamp = t.Timestamp
	c.AirspeedTimestampRel = t.AirspeedTimestampRel
	c.DistanceSensorTimestampRel = t.DistanceSensorTimestampRel
	c.OpticalFlowTimestampRel = t.OpticalFlowTimestampRel
	c.VehicleAirDataTimestampRel = t.VehicleAirDataTimestampRel
	c.VehicleMagnetometerTimestampRel = t.VehicleMagnetometerTimestampRel
	c.VisualOdometryTimestampRel = t.VisualOdometryTimestampRel
	return c
}

func (t *Ekf2Timestamps) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Ekf2Timestamps) SetDefaults() {
	t.Timestamp = 0
	t.AirspeedTimestampRel = 0
	t.DistanceSensorTimestampRel = 0
	t.OpticalFlowTimestampRel = 0
	t.VehicleAirDataTimestampRel = 0
	t.VehicleMagnetometerTimestampRel = 0
	t.VisualOdometryTimestampRel = 0
}

// CloneEkf2TimestampsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEkf2TimestampsSlice(dst, src []Ekf2Timestamps) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Ekf2TimestampsTypeSupport types.MessageTypeSupport = _Ekf2TimestampsTypeSupport{}

type _Ekf2TimestampsTypeSupport struct{}

func (t _Ekf2TimestampsTypeSupport) New() types.Message {
	return NewEkf2Timestamps()
}

func (t _Ekf2TimestampsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Ekf2Timestamps
	return (unsafe.Pointer)(C.px4_msgs__msg__Ekf2Timestamps__create())
}

func (t _Ekf2TimestampsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Ekf2Timestamps__destroy((*C.px4_msgs__msg__Ekf2Timestamps)(pointer_to_free))
}

func (t _Ekf2TimestampsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Ekf2Timestamps)
	mem := (*C.px4_msgs__msg__Ekf2Timestamps)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.airspeed_timestamp_rel = C.int16_t(m.AirspeedTimestampRel)
	mem.distance_sensor_timestamp_rel = C.int16_t(m.DistanceSensorTimestampRel)
	mem.optical_flow_timestamp_rel = C.int16_t(m.OpticalFlowTimestampRel)
	mem.vehicle_air_data_timestamp_rel = C.int16_t(m.VehicleAirDataTimestampRel)
	mem.vehicle_magnetometer_timestamp_rel = C.int16_t(m.VehicleMagnetometerTimestampRel)
	mem.visual_odometry_timestamp_rel = C.int16_t(m.VisualOdometryTimestampRel)
}

func (t _Ekf2TimestampsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Ekf2Timestamps)
	mem := (*C.px4_msgs__msg__Ekf2Timestamps)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.AirspeedTimestampRel = int16(mem.airspeed_timestamp_rel)
	m.DistanceSensorTimestampRel = int16(mem.distance_sensor_timestamp_rel)
	m.OpticalFlowTimestampRel = int16(mem.optical_flow_timestamp_rel)
	m.VehicleAirDataTimestampRel = int16(mem.vehicle_air_data_timestamp_rel)
	m.VehicleMagnetometerTimestampRel = int16(mem.vehicle_magnetometer_timestamp_rel)
	m.VisualOdometryTimestampRel = int16(mem.visual_odometry_timestamp_rel)
}

func (t _Ekf2TimestampsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Ekf2Timestamps())
}

type CEkf2Timestamps = C.px4_msgs__msg__Ekf2Timestamps
type CEkf2Timestamps__Sequence = C.px4_msgs__msg__Ekf2Timestamps__Sequence

func Ekf2Timestamps__Sequence_to_Go(goSlice *[]Ekf2Timestamps, cSlice CEkf2Timestamps__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Ekf2Timestamps, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Ekf2Timestamps__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Ekf2Timestamps * uintptr(i)),
		))
		Ekf2TimestampsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Ekf2Timestamps__Sequence_to_C(cSlice *CEkf2Timestamps__Sequence, goSlice []Ekf2Timestamps) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Ekf2Timestamps)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Ekf2Timestamps * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Ekf2Timestamps)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Ekf2Timestamps * uintptr(i)),
		))
		Ekf2TimestampsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Ekf2Timestamps__Array_to_Go(goSlice []Ekf2Timestamps, cSlice []CEkf2Timestamps) {
	for i := 0; i < len(cSlice); i++ {
		Ekf2TimestampsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Ekf2Timestamps__Array_to_C(cSlice []CEkf2Timestamps, goSlice []Ekf2Timestamps) {
	for i := 0; i < len(goSlice); i++ {
		Ekf2TimestampsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
