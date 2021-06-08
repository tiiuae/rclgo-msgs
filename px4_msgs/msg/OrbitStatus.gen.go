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

#include <px4_msgs/msg/orbit_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/OrbitStatus", OrbitStatusTypeSupport)
}
const (
	OrbitStatus_ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER uint8 = 0// ORBIT_YAW_BEHAVIOUR
	OrbitStatus_ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING uint8 = 1// ORBIT_YAW_BEHAVIOUR
	OrbitStatus_ORBIT_YAW_BEHAVIOUR_UNCONTROLLED uint8 = 2// ORBIT_YAW_BEHAVIOUR
	OrbitStatus_ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE uint8 = 3// ORBIT_YAW_BEHAVIOUR
	OrbitStatus_ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED uint8 = 4// ORBIT_YAW_BEHAVIOUR
)

// Do not create instances of this type directly. Always use NewOrbitStatus
// function instead.
type OrbitStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Radius float32 `yaml:"radius"`// Radius of the orbit circle. Positive values orbit clockwise, negative values orbit counter-clockwise. [m]
	Frame uint8 `yaml:"frame"`// The coordinate system of the fields: x, y, z.
	X float64 `yaml:"x"`// X coordinate of center point. Coordinate system depends on frame field: local = x position in meters * 1e4, global = latitude in degrees * 1e7.
	Y float64 `yaml:"y"`// Y coordinate of center point. Coordinate system depends on frame field: local = y position in meters * 1e4, global = latitude in degrees * 1e7.
	Z float32 `yaml:"z"`// Altitude of center point. Coordinate system depends on frame field.
	YawBehaviour uint8 `yaml:"yaw_behaviour"`
}

// NewOrbitStatus creates a new OrbitStatus with default values.
func NewOrbitStatus() *OrbitStatus {
	self := OrbitStatus{}
	self.SetDefaults()
	return &self
}

func (t *OrbitStatus) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *OrbitStatus) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var OrbitStatusTypeSupport types.MessageTypeSupport = _OrbitStatusTypeSupport{}

type _OrbitStatusTypeSupport struct{}

func (t _OrbitStatusTypeSupport) New() types.Message {
	return NewOrbitStatus()
}

func (t _OrbitStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__OrbitStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__OrbitStatus__create())
}

func (t _OrbitStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__OrbitStatus__destroy((*C.px4_msgs__msg__OrbitStatus)(pointer_to_free))
}

func (t _OrbitStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*OrbitStatus)
	mem := (*C.px4_msgs__msg__OrbitStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.radius = C.float(m.Radius)
	mem.frame = C.uint8_t(m.Frame)
	mem.x = C.double(m.X)
	mem.y = C.double(m.Y)
	mem.z = C.float(m.Z)
	mem.yaw_behaviour = C.uint8_t(m.YawBehaviour)
}

func (t _OrbitStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*OrbitStatus)
	mem := (*C.px4_msgs__msg__OrbitStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Radius = float32(mem.radius)
	m.Frame = uint8(mem.frame)
	m.X = float64(mem.x)
	m.Y = float64(mem.y)
	m.Z = float32(mem.z)
	m.YawBehaviour = uint8(mem.yaw_behaviour)
}

func (t _OrbitStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__OrbitStatus())
}

type COrbitStatus = C.px4_msgs__msg__OrbitStatus
type COrbitStatus__Sequence = C.px4_msgs__msg__OrbitStatus__Sequence

func OrbitStatus__Sequence_to_Go(goSlice *[]OrbitStatus, cSlice COrbitStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]OrbitStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__OrbitStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OrbitStatus * uintptr(i)),
		))
		OrbitStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func OrbitStatus__Sequence_to_C(cSlice *COrbitStatus__Sequence, goSlice []OrbitStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__OrbitStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__OrbitStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__OrbitStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OrbitStatus * uintptr(i)),
		))
		OrbitStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func OrbitStatus__Array_to_Go(goSlice []OrbitStatus, cSlice []COrbitStatus) {
	for i := 0; i < len(cSlice); i++ {
		OrbitStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func OrbitStatus__Array_to_C(cSlice []COrbitStatus, goSlice []OrbitStatus) {
	for i := 0; i < len(goSlice); i++ {
		OrbitStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
