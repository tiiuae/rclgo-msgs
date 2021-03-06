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

#include <px4_msgs/msg/vehicle_land_detected.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleLandDetected", VehicleLandDetectedTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleLandDetected
// function instead.
type VehicleLandDetected struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	AltMax float32 `yaml:"alt_max"`// maximum altitude in [m] that can be reached
	Freefall bool `yaml:"freefall"`// true if vehicle is currently in free-fall
	GroundContact bool `yaml:"ground_contact"`// true if vehicle has ground contact but is not landed (1. stage)
	MaybeLanded bool `yaml:"maybe_landed"`// true if the vehicle might have landed (2. stage)
	Landed bool `yaml:"landed"`// true if vehicle is currently landed on the ground (3. stage)
	InGroundEffect bool `yaml:"in_ground_effect"`// indicates if from the perspective of the landing detector the vehicle might be in ground effect (baro). This flag will become true if the vehicle is not moving horizontally and is descending (crude assumption that user is landing).
	InDescend bool `yaml:"in_descend"`
	HasLowThrottle bool `yaml:"has_low_throttle"`
	VerticalMovement bool `yaml:"vertical_movement"`
	HorizontalMovement bool `yaml:"horizontal_movement"`
	CloseToGroundOrSkippedCheck bool `yaml:"close_to_ground_or_skipped_check"`
}

// NewVehicleLandDetected creates a new VehicleLandDetected with default values.
func NewVehicleLandDetected() *VehicleLandDetected {
	self := VehicleLandDetected{}
	self.SetDefaults()
	return &self
}

func (t *VehicleLandDetected) Clone() *VehicleLandDetected {
	c := &VehicleLandDetected{}
	c.Timestamp = t.Timestamp
	c.AltMax = t.AltMax
	c.Freefall = t.Freefall
	c.GroundContact = t.GroundContact
	c.MaybeLanded = t.MaybeLanded
	c.Landed = t.Landed
	c.InGroundEffect = t.InGroundEffect
	c.InDescend = t.InDescend
	c.HasLowThrottle = t.HasLowThrottle
	c.VerticalMovement = t.VerticalMovement
	c.HorizontalMovement = t.HorizontalMovement
	c.CloseToGroundOrSkippedCheck = t.CloseToGroundOrSkippedCheck
	return c
}

func (t *VehicleLandDetected) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleLandDetected) SetDefaults() {
	t.Timestamp = 0
	t.AltMax = 0
	t.Freefall = false
	t.GroundContact = false
	t.MaybeLanded = false
	t.Landed = false
	t.InGroundEffect = false
	t.InDescend = false
	t.HasLowThrottle = false
	t.VerticalMovement = false
	t.HorizontalMovement = false
	t.CloseToGroundOrSkippedCheck = false
}

// CloneVehicleLandDetectedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleLandDetectedSlice(dst, src []VehicleLandDetected) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleLandDetectedTypeSupport types.MessageTypeSupport = _VehicleLandDetectedTypeSupport{}

type _VehicleLandDetectedTypeSupport struct{}

func (t _VehicleLandDetectedTypeSupport) New() types.Message {
	return NewVehicleLandDetected()
}

func (t _VehicleLandDetectedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleLandDetected
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleLandDetected__create())
}

func (t _VehicleLandDetectedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleLandDetected__destroy((*C.px4_msgs__msg__VehicleLandDetected)(pointer_to_free))
}

func (t _VehicleLandDetectedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleLandDetected)
	mem := (*C.px4_msgs__msg__VehicleLandDetected)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.alt_max = C.float(m.AltMax)
	mem.freefall = C.bool(m.Freefall)
	mem.ground_contact = C.bool(m.GroundContact)
	mem.maybe_landed = C.bool(m.MaybeLanded)
	mem.landed = C.bool(m.Landed)
	mem.in_ground_effect = C.bool(m.InGroundEffect)
	mem.in_descend = C.bool(m.InDescend)
	mem.has_low_throttle = C.bool(m.HasLowThrottle)
	mem.vertical_movement = C.bool(m.VerticalMovement)
	mem.horizontal_movement = C.bool(m.HorizontalMovement)
	mem.close_to_ground_or_skipped_check = C.bool(m.CloseToGroundOrSkippedCheck)
}

func (t _VehicleLandDetectedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleLandDetected)
	mem := (*C.px4_msgs__msg__VehicleLandDetected)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.AltMax = float32(mem.alt_max)
	m.Freefall = bool(mem.freefall)
	m.GroundContact = bool(mem.ground_contact)
	m.MaybeLanded = bool(mem.maybe_landed)
	m.Landed = bool(mem.landed)
	m.InGroundEffect = bool(mem.in_ground_effect)
	m.InDescend = bool(mem.in_descend)
	m.HasLowThrottle = bool(mem.has_low_throttle)
	m.VerticalMovement = bool(mem.vertical_movement)
	m.HorizontalMovement = bool(mem.horizontal_movement)
	m.CloseToGroundOrSkippedCheck = bool(mem.close_to_ground_or_skipped_check)
}

func (t _VehicleLandDetectedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleLandDetected())
}

type CVehicleLandDetected = C.px4_msgs__msg__VehicleLandDetected
type CVehicleLandDetected__Sequence = C.px4_msgs__msg__VehicleLandDetected__Sequence

func VehicleLandDetected__Sequence_to_Go(goSlice *[]VehicleLandDetected, cSlice CVehicleLandDetected__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleLandDetected, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleLandDetected__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLandDetected * uintptr(i)),
		))
		VehicleLandDetectedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleLandDetected__Sequence_to_C(cSlice *CVehicleLandDetected__Sequence, goSlice []VehicleLandDetected) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleLandDetected)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleLandDetected * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleLandDetected)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLandDetected * uintptr(i)),
		))
		VehicleLandDetectedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleLandDetected__Array_to_Go(goSlice []VehicleLandDetected, cSlice []CVehicleLandDetected) {
	for i := 0; i < len(cSlice); i++ {
		VehicleLandDetectedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleLandDetected__Array_to_C(cSlice []CVehicleLandDetected, goSlice []VehicleLandDetected) {
	for i := 0; i < len(goSlice); i++ {
		VehicleLandDetectedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
