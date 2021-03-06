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

#include <px4_msgs/msg/camera_capture.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/CameraCapture", CameraCaptureTypeSupport)
}

// Do not create instances of this type directly. Always use NewCameraCapture
// function instead.
type CameraCapture struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampUtc uint64 `yaml:"timestamp_utc"`// Capture time in UTC / GPS time
	Seq uint32 `yaml:"seq"`// Image sequence number
	Lat float64 `yaml:"lat"`// Latitude in degrees (WGS84)
	Lon float64 `yaml:"lon"`// Longitude in degrees (WGS84)
	Alt float32 `yaml:"alt"`// Altitude (AMSL)
	GroundDistance float32 `yaml:"ground_distance"`// Altitude above ground (meters)
	Q [4]float32 `yaml:"q"`// Attitude of the camera, zero rotation is facing towards front of vehicle
	Result int8 `yaml:"result"`// 1 for success, 0 for failure, -1 if camera does not provide feedback
}

// NewCameraCapture creates a new CameraCapture with default values.
func NewCameraCapture() *CameraCapture {
	self := CameraCapture{}
	self.SetDefaults()
	return &self
}

func (t *CameraCapture) Clone() *CameraCapture {
	c := &CameraCapture{}
	c.Timestamp = t.Timestamp
	c.TimestampUtc = t.TimestampUtc
	c.Seq = t.Seq
	c.Lat = t.Lat
	c.Lon = t.Lon
	c.Alt = t.Alt
	c.GroundDistance = t.GroundDistance
	c.Q = t.Q
	c.Result = t.Result
	return c
}

func (t *CameraCapture) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CameraCapture) SetDefaults() {
	t.Timestamp = 0
	t.TimestampUtc = 0
	t.Seq = 0
	t.Lat = 0
	t.Lon = 0
	t.Alt = 0
	t.GroundDistance = 0
	t.Q = [4]float32{}
	t.Result = 0
}

// CloneCameraCaptureSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCameraCaptureSlice(dst, src []CameraCapture) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CameraCaptureTypeSupport types.MessageTypeSupport = _CameraCaptureTypeSupport{}

type _CameraCaptureTypeSupport struct{}

func (t _CameraCaptureTypeSupport) New() types.Message {
	return NewCameraCapture()
}

func (t _CameraCaptureTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CameraCapture
	return (unsafe.Pointer)(C.px4_msgs__msg__CameraCapture__create())
}

func (t _CameraCaptureTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CameraCapture__destroy((*C.px4_msgs__msg__CameraCapture)(pointer_to_free))
}

func (t _CameraCaptureTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CameraCapture)
	mem := (*C.px4_msgs__msg__CameraCapture)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_utc = C.uint64_t(m.TimestampUtc)
	mem.seq = C.uint32_t(m.Seq)
	mem.lat = C.double(m.Lat)
	mem.lon = C.double(m.Lon)
	mem.alt = C.float(m.Alt)
	mem.ground_distance = C.float(m.GroundDistance)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	mem.result = C.int8_t(m.Result)
}

func (t _CameraCaptureTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CameraCapture)
	mem := (*C.px4_msgs__msg__CameraCapture)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampUtc = uint64(mem.timestamp_utc)
	m.Seq = uint32(mem.seq)
	m.Lat = float64(mem.lat)
	m.Lon = float64(mem.lon)
	m.Alt = float32(mem.alt)
	m.GroundDistance = float32(mem.ground_distance)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_Go(m.Q[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)))
	m.Result = int8(mem.result)
}

func (t _CameraCaptureTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CameraCapture())
}

type CCameraCapture = C.px4_msgs__msg__CameraCapture
type CCameraCapture__Sequence = C.px4_msgs__msg__CameraCapture__Sequence

func CameraCapture__Sequence_to_Go(goSlice *[]CameraCapture, cSlice CCameraCapture__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CameraCapture, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CameraCapture__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraCapture * uintptr(i)),
		))
		CameraCaptureTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func CameraCapture__Sequence_to_C(cSlice *CCameraCapture__Sequence, goSlice []CameraCapture) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CameraCapture)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CameraCapture * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CameraCapture)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraCapture * uintptr(i)),
		))
		CameraCaptureTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CameraCapture__Array_to_Go(goSlice []CameraCapture, cSlice []CCameraCapture) {
	for i := 0; i < len(cSlice); i++ {
		CameraCaptureTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CameraCapture__Array_to_C(cSlice []CCameraCapture, goSlice []CameraCapture) {
	for i := 0; i < len(goSlice); i++ {
		CameraCaptureTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
