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

#include <px4_msgs/msg/estimator_selector_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorSelectorStatus", EstimatorSelectorStatusTypeSupport)
}

// Do not create instances of this type directly. Always use NewEstimatorSelectorStatus
// function instead.
type EstimatorSelectorStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	PrimaryInstance uint8 `yaml:"primary_instance"`
	InstancesAvailable uint8 `yaml:"instances_available"`
	InstanceChangedCount uint32 `yaml:"instance_changed_count"`
	LastInstanceChange uint64 `yaml:"last_instance_change"`
	AccelDeviceId uint32 `yaml:"accel_device_id"`
	BaroDeviceId uint32 `yaml:"baro_device_id"`
	GyroDeviceId uint32 `yaml:"gyro_device_id"`
	MagDeviceId uint32 `yaml:"mag_device_id"`
	CombinedTestRatio [9]float32 `yaml:"combined_test_ratio"`
	RelativeTestRatio [9]float32 `yaml:"relative_test_ratio"`
	Healthy [9]bool `yaml:"healthy"`
	AccumulatedGyroError [4]float32 `yaml:"accumulated_gyro_error"`
	AccumulatedAccelError [4]float32 `yaml:"accumulated_accel_error"`
	GyroFaultDetected bool `yaml:"gyro_fault_detected"`
	AccelFaultDetected bool `yaml:"accel_fault_detected"`
}

// NewEstimatorSelectorStatus creates a new EstimatorSelectorStatus with default values.
func NewEstimatorSelectorStatus() *EstimatorSelectorStatus {
	self := EstimatorSelectorStatus{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorSelectorStatus) Clone() *EstimatorSelectorStatus {
	c := &EstimatorSelectorStatus{}
	c.Timestamp = t.Timestamp
	c.PrimaryInstance = t.PrimaryInstance
	c.InstancesAvailable = t.InstancesAvailable
	c.InstanceChangedCount = t.InstanceChangedCount
	c.LastInstanceChange = t.LastInstanceChange
	c.AccelDeviceId = t.AccelDeviceId
	c.BaroDeviceId = t.BaroDeviceId
	c.GyroDeviceId = t.GyroDeviceId
	c.MagDeviceId = t.MagDeviceId
	c.CombinedTestRatio = t.CombinedTestRatio
	c.RelativeTestRatio = t.RelativeTestRatio
	c.Healthy = t.Healthy
	c.AccumulatedGyroError = t.AccumulatedGyroError
	c.AccumulatedAccelError = t.AccumulatedAccelError
	c.GyroFaultDetected = t.GyroFaultDetected
	c.AccelFaultDetected = t.AccelFaultDetected
	return c
}

func (t *EstimatorSelectorStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorSelectorStatus) SetDefaults() {
	t.Timestamp = 0
	t.PrimaryInstance = 0
	t.InstancesAvailable = 0
	t.InstanceChangedCount = 0
	t.LastInstanceChange = 0
	t.AccelDeviceId = 0
	t.BaroDeviceId = 0
	t.GyroDeviceId = 0
	t.MagDeviceId = 0
	t.CombinedTestRatio = [9]float32{}
	t.RelativeTestRatio = [9]float32{}
	t.Healthy = [9]bool{}
	t.AccumulatedGyroError = [4]float32{}
	t.AccumulatedAccelError = [4]float32{}
	t.GyroFaultDetected = false
	t.AccelFaultDetected = false
}

// CloneEstimatorSelectorStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorSelectorStatusSlice(dst, src []EstimatorSelectorStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorSelectorStatusTypeSupport types.MessageTypeSupport = _EstimatorSelectorStatusTypeSupport{}

type _EstimatorSelectorStatusTypeSupport struct{}

func (t _EstimatorSelectorStatusTypeSupport) New() types.Message {
	return NewEstimatorSelectorStatus()
}

func (t _EstimatorSelectorStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorSelectorStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorSelectorStatus__create())
}

func (t _EstimatorSelectorStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorSelectorStatus__destroy((*C.px4_msgs__msg__EstimatorSelectorStatus)(pointer_to_free))
}

func (t _EstimatorSelectorStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorSelectorStatus)
	mem := (*C.px4_msgs__msg__EstimatorSelectorStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.primary_instance = C.uint8_t(m.PrimaryInstance)
	mem.instances_available = C.uint8_t(m.InstancesAvailable)
	mem.instance_changed_count = C.uint32_t(m.InstanceChangedCount)
	mem.last_instance_change = C.uint64_t(m.LastInstanceChange)
	mem.accel_device_id = C.uint32_t(m.AccelDeviceId)
	mem.baro_device_id = C.uint32_t(m.BaroDeviceId)
	mem.gyro_device_id = C.uint32_t(m.GyroDeviceId)
	mem.mag_device_id = C.uint32_t(m.MagDeviceId)
	cSlice_combined_test_ratio := mem.combined_test_ratio[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_combined_test_ratio)), m.CombinedTestRatio[:])
	cSlice_relative_test_ratio := mem.relative_test_ratio[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_relative_test_ratio)), m.RelativeTestRatio[:])
	cSlice_healthy := mem.healthy[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_healthy)), m.Healthy[:])
	cSlice_accumulated_gyro_error := mem.accumulated_gyro_error[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accumulated_gyro_error)), m.AccumulatedGyroError[:])
	cSlice_accumulated_accel_error := mem.accumulated_accel_error[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accumulated_accel_error)), m.AccumulatedAccelError[:])
	mem.gyro_fault_detected = C.bool(m.GyroFaultDetected)
	mem.accel_fault_detected = C.bool(m.AccelFaultDetected)
}

func (t _EstimatorSelectorStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorSelectorStatus)
	mem := (*C.px4_msgs__msg__EstimatorSelectorStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.PrimaryInstance = uint8(mem.primary_instance)
	m.InstancesAvailable = uint8(mem.instances_available)
	m.InstanceChangedCount = uint32(mem.instance_changed_count)
	m.LastInstanceChange = uint64(mem.last_instance_change)
	m.AccelDeviceId = uint32(mem.accel_device_id)
	m.BaroDeviceId = uint32(mem.baro_device_id)
	m.GyroDeviceId = uint32(mem.gyro_device_id)
	m.MagDeviceId = uint32(mem.mag_device_id)
	cSlice_combined_test_ratio := mem.combined_test_ratio[:]
	primitives.Float32__Array_to_Go(m.CombinedTestRatio[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_combined_test_ratio)))
	cSlice_relative_test_ratio := mem.relative_test_ratio[:]
	primitives.Float32__Array_to_Go(m.RelativeTestRatio[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_relative_test_ratio)))
	cSlice_healthy := mem.healthy[:]
	primitives.Bool__Array_to_Go(m.Healthy[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_healthy)))
	cSlice_accumulated_gyro_error := mem.accumulated_gyro_error[:]
	primitives.Float32__Array_to_Go(m.AccumulatedGyroError[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accumulated_gyro_error)))
	cSlice_accumulated_accel_error := mem.accumulated_accel_error[:]
	primitives.Float32__Array_to_Go(m.AccumulatedAccelError[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accumulated_accel_error)))
	m.GyroFaultDetected = bool(mem.gyro_fault_detected)
	m.AccelFaultDetected = bool(mem.accel_fault_detected)
}

func (t _EstimatorSelectorStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorSelectorStatus())
}

type CEstimatorSelectorStatus = C.px4_msgs__msg__EstimatorSelectorStatus
type CEstimatorSelectorStatus__Sequence = C.px4_msgs__msg__EstimatorSelectorStatus__Sequence

func EstimatorSelectorStatus__Sequence_to_Go(goSlice *[]EstimatorSelectorStatus, cSlice CEstimatorSelectorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorSelectorStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorSelectorStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorSelectorStatus * uintptr(i)),
		))
		EstimatorSelectorStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorSelectorStatus__Sequence_to_C(cSlice *CEstimatorSelectorStatus__Sequence, goSlice []EstimatorSelectorStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorSelectorStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorSelectorStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorSelectorStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorSelectorStatus * uintptr(i)),
		))
		EstimatorSelectorStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorSelectorStatus__Array_to_Go(goSlice []EstimatorSelectorStatus, cSlice []CEstimatorSelectorStatus) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorSelectorStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorSelectorStatus__Array_to_C(cSlice []CEstimatorSelectorStatus, goSlice []EstimatorSelectorStatus) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorSelectorStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
