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

#include <px4_msgs/msg/distance_sensor.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/DistanceSensor", DistanceSensorTypeSupport)
}
const (
	DistanceSensor_MAV_DISTANCE_SENSOR_LASER uint8 = 0
	DistanceSensor_MAV_DISTANCE_SENSOR_ULTRASOUND uint8 = 1
	DistanceSensor_MAV_DISTANCE_SENSOR_INFRARED uint8 = 2
	DistanceSensor_MAV_DISTANCE_SENSOR_RADAR uint8 = 3
	DistanceSensor_ROTATION_YAW_0 uint8 = 0// MAV_SENSOR_ROTATION_NONE
	DistanceSensor_ROTATION_YAW_45 uint8 = 1// MAV_SENSOR_ROTATION_YAW_45
	DistanceSensor_ROTATION_YAW_90 uint8 = 2// MAV_SENSOR_ROTATION_YAW_90
	DistanceSensor_ROTATION_YAW_135 uint8 = 3// MAV_SENSOR_ROTATION_YAW_135
	DistanceSensor_ROTATION_YAW_180 uint8 = 4// MAV_SENSOR_ROTATION_YAW_180
	DistanceSensor_ROTATION_YAW_225 uint8 = 5// MAV_SENSOR_ROTATION_YAW_225
	DistanceSensor_ROTATION_YAW_270 uint8 = 6// MAV_SENSOR_ROTATION_YAW_270
	DistanceSensor_ROTATION_YAW_315 uint8 = 7// MAV_SENSOR_ROTATION_YAW_315
	DistanceSensor_ROTATION_FORWARD_FACING uint8 = 0// MAV_SENSOR_ROTATION_NONE
	DistanceSensor_ROTATION_RIGHT_FACING uint8 = 2// MAV_SENSOR_ROTATION_YAW_90
	DistanceSensor_ROTATION_BACKWARD_FACING uint8 = 4// MAV_SENSOR_ROTATION_YAW_180
	DistanceSensor_ROTATION_LEFT_FACING uint8 = 6// MAV_SENSOR_ROTATION_YAW_270
	DistanceSensor_ROTATION_UPWARD_FACING uint8 = 24// MAV_SENSOR_ROTATION_PITCH_90
	DistanceSensor_ROTATION_DOWNWARD_FACING uint8 = 25// MAV_SENSOR_ROTATION_PITCH_270
	DistanceSensor_ROTATION_CUSTOM uint8 = 100// MAV_SENSOR_ROTATION_CUSTOM
)

// Do not create instances of this type directly. Always use NewDistanceSensor
// function instead.
type DistanceSensor struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
	MinDistance float32 `yaml:"min_distance"`// Minimum distance the sensor can measure (in m)
	MaxDistance float32 `yaml:"max_distance"`// Maximum distance the sensor can measure (in m)
	CurrentDistance float32 `yaml:"current_distance"`// Current distance reading (in m)
	Variance float32 `yaml:"variance"`// Measurement variance (in m^2), 0 for unknown / invalid readings
	SignalQuality int8 `yaml:"signal_quality"`// Signal quality in percent (0...100%), where 0 = invalid signal, 100 = perfect signal, and -1 = unknown signal quality.
	Type uint8 `yaml:"type"`// Type from MAV_DISTANCE_SENSOR enum
	HFov float32 `yaml:"h_fov"`// Sensor horizontal field of view (rad)
	VFov float32 `yaml:"v_fov"`// Sensor vertical field of view (rad)
	Q [4]float32 `yaml:"q"`// Quaterion sensor orientation with respect to the vehicle body frame to specify the orientation ROTATION_CUSTOM
	Orientation uint8 `yaml:"orientation"`// Direction the sensor faces from MAV_SENSOR_ORIENTATION enum
}

// NewDistanceSensor creates a new DistanceSensor with default values.
func NewDistanceSensor() *DistanceSensor {
	self := DistanceSensor{}
	self.SetDefaults()
	return &self
}

func (t *DistanceSensor) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *DistanceSensor) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var DistanceSensorTypeSupport types.MessageTypeSupport = _DistanceSensorTypeSupport{}

type _DistanceSensorTypeSupport struct{}

func (t _DistanceSensorTypeSupport) New() types.Message {
	return NewDistanceSensor()
}

func (t _DistanceSensorTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__DistanceSensor
	return (unsafe.Pointer)(C.px4_msgs__msg__DistanceSensor__create())
}

func (t _DistanceSensorTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__DistanceSensor__destroy((*C.px4_msgs__msg__DistanceSensor)(pointer_to_free))
}

func (t _DistanceSensorTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DistanceSensor)
	mem := (*C.px4_msgs__msg__DistanceSensor)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.device_id = C.uint32_t(m.DeviceId)
	mem.min_distance = C.float(m.MinDistance)
	mem.max_distance = C.float(m.MaxDistance)
	mem.current_distance = C.float(m.CurrentDistance)
	mem.variance = C.float(m.Variance)
	mem.signal_quality = C.int8_t(m.SignalQuality)
	mem._type = C.uint8_t(m.Type)
	mem.h_fov = C.float(m.HFov)
	mem.v_fov = C.float(m.VFov)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	mem.orientation = C.uint8_t(m.Orientation)
}

func (t _DistanceSensorTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DistanceSensor)
	mem := (*C.px4_msgs__msg__DistanceSensor)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.DeviceId = uint32(mem.device_id)
	m.MinDistance = float32(mem.min_distance)
	m.MaxDistance = float32(mem.max_distance)
	m.CurrentDistance = float32(mem.current_distance)
	m.Variance = float32(mem.variance)
	m.SignalQuality = int8(mem.signal_quality)
	m.Type = uint8(mem._type)
	m.HFov = float32(mem.h_fov)
	m.VFov = float32(mem.v_fov)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_Go(m.Q[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)))
	m.Orientation = uint8(mem.orientation)
}

func (t _DistanceSensorTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__DistanceSensor())
}

type CDistanceSensor = C.px4_msgs__msg__DistanceSensor
type CDistanceSensor__Sequence = C.px4_msgs__msg__DistanceSensor__Sequence

func DistanceSensor__Sequence_to_Go(goSlice *[]DistanceSensor, cSlice CDistanceSensor__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DistanceSensor, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__DistanceSensor__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DistanceSensor * uintptr(i)),
		))
		DistanceSensorTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func DistanceSensor__Sequence_to_C(cSlice *CDistanceSensor__Sequence, goSlice []DistanceSensor) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__DistanceSensor)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__DistanceSensor * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__DistanceSensor)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DistanceSensor * uintptr(i)),
		))
		DistanceSensorTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func DistanceSensor__Array_to_Go(goSlice []DistanceSensor, cSlice []CDistanceSensor) {
	for i := 0; i < len(cSlice); i++ {
		DistanceSensorTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DistanceSensor__Array_to_C(cSlice []CDistanceSensor, goSlice []DistanceSensor) {
	for i := 0; i < len(goSlice); i++ {
		DistanceSensorTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
