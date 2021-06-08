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

#include <px4_msgs/msg/vehicle_trajectory_waypoint.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleTrajectoryWaypoint", VehicleTrajectoryWaypointTypeSupport)
}
const (
	VehicleTrajectoryWaypoint_MAV_TRAJECTORY_REPRESENTATION_WAYPOINTS uint8 = 0
	VehicleTrajectoryWaypoint_POINT_0 uint8 = 0
	VehicleTrajectoryWaypoint_POINT_1 uint8 = 1
	VehicleTrajectoryWaypoint_POINT_2 uint8 = 2
	VehicleTrajectoryWaypoint_POINT_3 uint8 = 3
	VehicleTrajectoryWaypoint_POINT_4 uint8 = 4
	VehicleTrajectoryWaypoint_NUMBER_POINTS uint8 = 5
)

// Do not create instances of this type directly. Always use NewVehicleTrajectoryWaypoint
// function instead.
type VehicleTrajectoryWaypoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Type uint8 `yaml:"type"`// Type from MAV_TRAJECTORY_REPRESENTATION enum.
	Waypoints [5]TrajectoryWaypoint `yaml:"waypoints"`
}

// NewVehicleTrajectoryWaypoint creates a new VehicleTrajectoryWaypoint with default values.
func NewVehicleTrajectoryWaypoint() *VehicleTrajectoryWaypoint {
	self := VehicleTrajectoryWaypoint{}
	self.SetDefaults()
	return &self
}

func (t *VehicleTrajectoryWaypoint) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleTrajectoryWaypoint) SetDefaults() {
	t.Waypoints[0].SetDefaults()
	t.Waypoints[1].SetDefaults()
	t.Waypoints[2].SetDefaults()
	t.Waypoints[3].SetDefaults()
	t.Waypoints[4].SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var VehicleTrajectoryWaypointTypeSupport types.MessageTypeSupport = _VehicleTrajectoryWaypointTypeSupport{}

type _VehicleTrajectoryWaypointTypeSupport struct{}

func (t _VehicleTrajectoryWaypointTypeSupport) New() types.Message {
	return NewVehicleTrajectoryWaypoint()
}

func (t _VehicleTrajectoryWaypointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleTrajectoryWaypoint
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleTrajectoryWaypoint__create())
}

func (t _VehicleTrajectoryWaypointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleTrajectoryWaypoint__destroy((*C.px4_msgs__msg__VehicleTrajectoryWaypoint)(pointer_to_free))
}

func (t _VehicleTrajectoryWaypointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleTrajectoryWaypoint)
	mem := (*C.px4_msgs__msg__VehicleTrajectoryWaypoint)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem._type = C.uint8_t(m.Type)
	TrajectoryWaypoint__Array_to_C(mem.waypoints[:], m.Waypoints[:])
}

func (t _VehicleTrajectoryWaypointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleTrajectoryWaypoint)
	mem := (*C.px4_msgs__msg__VehicleTrajectoryWaypoint)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Type = uint8(mem._type)
	TrajectoryWaypoint__Array_to_Go(m.Waypoints[:], mem.waypoints[:])
}

func (t _VehicleTrajectoryWaypointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleTrajectoryWaypoint())
}

type CVehicleTrajectoryWaypoint = C.px4_msgs__msg__VehicleTrajectoryWaypoint
type CVehicleTrajectoryWaypoint__Sequence = C.px4_msgs__msg__VehicleTrajectoryWaypoint__Sequence

func VehicleTrajectoryWaypoint__Sequence_to_Go(goSlice *[]VehicleTrajectoryWaypoint, cSlice CVehicleTrajectoryWaypoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleTrajectoryWaypoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleTrajectoryWaypoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryWaypoint * uintptr(i)),
		))
		VehicleTrajectoryWaypointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleTrajectoryWaypoint__Sequence_to_C(cSlice *CVehicleTrajectoryWaypoint__Sequence, goSlice []VehicleTrajectoryWaypoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleTrajectoryWaypoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryWaypoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleTrajectoryWaypoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryWaypoint * uintptr(i)),
		))
		VehicleTrajectoryWaypointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleTrajectoryWaypoint__Array_to_Go(goSlice []VehicleTrajectoryWaypoint, cSlice []CVehicleTrajectoryWaypoint) {
	for i := 0; i < len(cSlice); i++ {
		VehicleTrajectoryWaypointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleTrajectoryWaypoint__Array_to_C(cSlice []CVehicleTrajectoryWaypoint, goSlice []VehicleTrajectoryWaypoint) {
	for i := 0; i < len(goSlice); i++ {
		VehicleTrajectoryWaypointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}