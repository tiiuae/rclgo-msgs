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

#include <px4_msgs/msg/obstacle_distance.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ObstacleDistance", ObstacleDistanceTypeSupport)
}
const (
	ObstacleDistance_MAV_FRAME_GLOBAL uint8 = 0
	ObstacleDistance_MAV_FRAME_LOCAL_NED uint8 = 1
	ObstacleDistance_MAV_FRAME_BODY_FRD uint8 = 12
	ObstacleDistance_MAV_DISTANCE_SENSOR_LASER uint8 = 0
	ObstacleDistance_MAV_DISTANCE_SENSOR_ULTRASOUND uint8 = 1
	ObstacleDistance_MAV_DISTANCE_SENSOR_INFRARED uint8 = 2
	ObstacleDistance_MAV_DISTANCE_SENSOR_RADAR uint8 = 3
)

// Do not create instances of this type directly. Always use NewObstacleDistance
// function instead.
type ObstacleDistance struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). Obstacle distances in front of the sensor.
	Frame uint8 `yaml:"frame"`// Coordinate frame of reference for the yaw rotation and offset of the sensor data. Defaults to MAV_FRAME_GLOBAL, which is North aligned. For body-mounted sensors use MAV_FRAME_BODY_FRD, which is vehicle front aligned.
	SensorType uint8 `yaml:"sensor_type"`// Type from MAV_DISTANCE_SENSOR enum.
	Distances [72]uint16 `yaml:"distances"`// Distance of obstacles around the UAV with index 0 corresponding to local North. A value of 0 means that the obstacle is right in front of the sensor. A value of max_distance +1 means no obstacle is present. A value of UINT16_MAX for unknown/not used. In a array element, one unit corresponds to 1cm.
	Increment float32 `yaml:"increment"`// Angular width in degrees of each array element.
	MinDistance uint16 `yaml:"min_distance"`// Minimum distance the sensor can measure in centimeters.
	MaxDistance uint16 `yaml:"max_distance"`// Maximum distance the sensor can measure in centimeters.
	AngleOffset float32 `yaml:"angle_offset"`// Relative angle offset of the 0-index element in the distances array. Value of 0 corresponds to forward. Positive values are offsets to the right.
}

// NewObstacleDistance creates a new ObstacleDistance with default values.
func NewObstacleDistance() *ObstacleDistance {
	self := ObstacleDistance{}
	self.SetDefaults()
	return &self
}

func (t *ObstacleDistance) Clone() *ObstacleDistance {
	c := &ObstacleDistance{}
	c.Timestamp = t.Timestamp
	c.Frame = t.Frame
	c.SensorType = t.SensorType
	c.Distances = t.Distances
	c.Increment = t.Increment
	c.MinDistance = t.MinDistance
	c.MaxDistance = t.MaxDistance
	c.AngleOffset = t.AngleOffset
	return c
}

func (t *ObstacleDistance) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ObstacleDistance) SetDefaults() {
	t.Timestamp = 0
	t.Frame = 0
	t.SensorType = 0
	t.Distances = [72]uint16{}
	t.Increment = 0
	t.MinDistance = 0
	t.MaxDistance = 0
	t.AngleOffset = 0
}

// CloneObstacleDistanceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneObstacleDistanceSlice(dst, src []ObstacleDistance) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ObstacleDistanceTypeSupport types.MessageTypeSupport = _ObstacleDistanceTypeSupport{}

type _ObstacleDistanceTypeSupport struct{}

func (t _ObstacleDistanceTypeSupport) New() types.Message {
	return NewObstacleDistance()
}

func (t _ObstacleDistanceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ObstacleDistance
	return (unsafe.Pointer)(C.px4_msgs__msg__ObstacleDistance__create())
}

func (t _ObstacleDistanceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ObstacleDistance__destroy((*C.px4_msgs__msg__ObstacleDistance)(pointer_to_free))
}

func (t _ObstacleDistanceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ObstacleDistance)
	mem := (*C.px4_msgs__msg__ObstacleDistance)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.frame = C.uint8_t(m.Frame)
	mem.sensor_type = C.uint8_t(m.SensorType)
	cSlice_distances := mem.distances[:]
	primitives.Uint16__Array_to_C(*(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_distances)), m.Distances[:])
	mem.increment = C.float(m.Increment)
	mem.min_distance = C.uint16_t(m.MinDistance)
	mem.max_distance = C.uint16_t(m.MaxDistance)
	mem.angle_offset = C.float(m.AngleOffset)
}

func (t _ObstacleDistanceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ObstacleDistance)
	mem := (*C.px4_msgs__msg__ObstacleDistance)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Frame = uint8(mem.frame)
	m.SensorType = uint8(mem.sensor_type)
	cSlice_distances := mem.distances[:]
	primitives.Uint16__Array_to_Go(m.Distances[:], *(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_distances)))
	m.Increment = float32(mem.increment)
	m.MinDistance = uint16(mem.min_distance)
	m.MaxDistance = uint16(mem.max_distance)
	m.AngleOffset = float32(mem.angle_offset)
}

func (t _ObstacleDistanceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ObstacleDistance())
}

type CObstacleDistance = C.px4_msgs__msg__ObstacleDistance
type CObstacleDistance__Sequence = C.px4_msgs__msg__ObstacleDistance__Sequence

func ObstacleDistance__Sequence_to_Go(goSlice *[]ObstacleDistance, cSlice CObstacleDistance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ObstacleDistance, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ObstacleDistance__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ObstacleDistance * uintptr(i)),
		))
		ObstacleDistanceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ObstacleDistance__Sequence_to_C(cSlice *CObstacleDistance__Sequence, goSlice []ObstacleDistance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ObstacleDistance)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ObstacleDistance * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ObstacleDistance)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ObstacleDistance * uintptr(i)),
		))
		ObstacleDistanceTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ObstacleDistance__Array_to_Go(goSlice []ObstacleDistance, cSlice []CObstacleDistance) {
	for i := 0; i < len(cSlice); i++ {
		ObstacleDistanceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ObstacleDistance__Array_to_C(cSlice []CObstacleDistance, goSlice []ObstacleDistance) {
	for i := 0; i < len(goSlice); i++ {
		ObstacleDistanceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
