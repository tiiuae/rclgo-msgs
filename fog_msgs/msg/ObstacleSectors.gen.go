/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/msg/obstacle_sectors.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/ObstacleSectors", ObstacleSectorsTypeSupport)
}
const (
	ObstacleSectors_OBSTACLE_NOT_DETECTED float64 = -1// # Obstacle distance special values.
	ObstacleSectors_OBSTACLE_NO_DATA float64 = -2// # Obstacle distance special values.
	ObstacleSectors_SENSOR_NONE int8 = -1// # Sensor types enum.
	ObstacleSectors_SENSOR_DEPTH int8 = 0// # Sensor types enum.
	ObstacleSectors_SENSOR_LIDAR1D int8 = 1// # Sensor types enum.
	ObstacleSectors_SENSOR_LIDAR2D int8 = 2// # Sensor types enum.
	ObstacleSectors_SENSOR_LIDAR3D int8 = 3// # Sensor types enum.
)

// Do not create instances of this type directly. Always use NewObstacleSectors
// function instead.
type ObstacleSectors struct {
	Header std_msgs_msg.Header `yaml:"header"`// Time at which the data in this message was generated, coordinate frame ID.
	NHorizontalSectors uint32 `yaml:"n_horizontal_sectors"`// Number of horizontal bumper sectors (total number of sectors -2).
	SectorsVerticalFov float64 `yaml:"sectors_vertical_fov"`// Vertical FOV of the individual horizontal sectors.
	Sectors []float64 `yaml:"sectors"`// Distance to closest detected obstacle in each sector.
	SectorSensors []int8 `yaml:"sector_sensors"`// Sensor ID of the sensor, which was used for obstacle detection in each sector.
}

// NewObstacleSectors creates a new ObstacleSectors with default values.
func NewObstacleSectors() *ObstacleSectors {
	self := ObstacleSectors{}
	self.SetDefaults()
	return &self
}

func (t *ObstacleSectors) Clone() *ObstacleSectors {
	c := &ObstacleSectors{}
	c.Header = *t.Header.Clone()
	c.NHorizontalSectors = t.NHorizontalSectors
	c.SectorsVerticalFov = t.SectorsVerticalFov
	if t.Sectors != nil {
		c.Sectors = make([]float64, len(t.Sectors))
		copy(c.Sectors, t.Sectors)
	}
	if t.SectorSensors != nil {
		c.SectorSensors = make([]int8, len(t.SectorSensors))
		copy(c.SectorSensors, t.SectorSensors)
	}
	return c
}

func (t *ObstacleSectors) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ObstacleSectors) SetDefaults() {
	t.Header.SetDefaults()
	t.NHorizontalSectors = 0
	t.SectorsVerticalFov = 0
	t.Sectors = nil
	t.SectorSensors = nil
}

// CloneObstacleSectorsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneObstacleSectorsSlice(dst, src []ObstacleSectors) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ObstacleSectorsTypeSupport types.MessageTypeSupport = _ObstacleSectorsTypeSupport{}

type _ObstacleSectorsTypeSupport struct{}

func (t _ObstacleSectorsTypeSupport) New() types.Message {
	return NewObstacleSectors()
}

func (t _ObstacleSectorsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__msg__ObstacleSectors
	return (unsafe.Pointer)(C.fog_msgs__msg__ObstacleSectors__create())
}

func (t _ObstacleSectorsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__msg__ObstacleSectors__destroy((*C.fog_msgs__msg__ObstacleSectors)(pointer_to_free))
}

func (t _ObstacleSectorsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ObstacleSectors)
	mem := (*C.fog_msgs__msg__ObstacleSectors)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.n_horizontal_sectors = C.uint32_t(m.NHorizontalSectors)
	mem.sectors_vertical_fov = C.double(m.SectorsVerticalFov)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.sectors)), m.Sectors)
	primitives.Int8__Sequence_to_C((*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.sector_sensors)), m.SectorSensors)
}

func (t _ObstacleSectorsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ObstacleSectors)
	mem := (*C.fog_msgs__msg__ObstacleSectors)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.NHorizontalSectors = uint32(mem.n_horizontal_sectors)
	m.SectorsVerticalFov = float64(mem.sectors_vertical_fov)
	primitives.Float64__Sequence_to_Go(&m.Sectors, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.sectors)))
	primitives.Int8__Sequence_to_Go(&m.SectorSensors, *(*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.sector_sensors)))
}

func (t _ObstacleSectorsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__msg__ObstacleSectors())
}

type CObstacleSectors = C.fog_msgs__msg__ObstacleSectors
type CObstacleSectors__Sequence = C.fog_msgs__msg__ObstacleSectors__Sequence

func ObstacleSectors__Sequence_to_Go(goSlice *[]ObstacleSectors, cSlice CObstacleSectors__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ObstacleSectors, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__msg__ObstacleSectors__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__msg__ObstacleSectors * uintptr(i)),
		))
		ObstacleSectorsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ObstacleSectors__Sequence_to_C(cSlice *CObstacleSectors__Sequence, goSlice []ObstacleSectors) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__msg__ObstacleSectors)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__msg__ObstacleSectors * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__msg__ObstacleSectors)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__msg__ObstacleSectors * uintptr(i)),
		))
		ObstacleSectorsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ObstacleSectors__Array_to_Go(goSlice []ObstacleSectors, cSlice []CObstacleSectors) {
	for i := 0; i < len(cSlice); i++ {
		ObstacleSectorsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ObstacleSectors__Array_to_C(cSlice []CObstacleSectors, goSlice []ObstacleSectors) {
	for i := 0; i < len(goSlice); i++ {
		ObstacleSectorsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
