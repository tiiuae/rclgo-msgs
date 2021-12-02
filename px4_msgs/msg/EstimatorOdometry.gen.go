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

#include <px4_msgs/msg/estimator_odometry.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorOdometry", EstimatorOdometryTypeSupport)
}
const (
	EstimatorOdometry_COVARIANCE_MATRIX_X_VARIANCE uint8 = 0// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_Y_VARIANCE uint8 = 6// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_Z_VARIANCE uint8 = 11// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_ROLL_VARIANCE uint8 = 15// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_PITCH_VARIANCE uint8 = 18// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_YAW_VARIANCE uint8 = 20// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_VX_VARIANCE uint8 = 0// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_VY_VARIANCE uint8 = 6// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_VZ_VARIANCE uint8 = 11// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_ROLLRATE_VARIANCE uint8 = 15// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_PITCHRATE_VARIANCE uint8 = 18// Covariance matrix index constants
	EstimatorOdometry_COVARIANCE_MATRIX_YAWRATE_VARIANCE uint8 = 20// Covariance matrix index constants
	EstimatorOdometry_LOCAL_FRAME_NED uint8 = 0// NED earth-fixed frame. Position and linear velocity frame of reference constants
	EstimatorOdometry_LOCAL_FRAME_FRD uint8 = 1// FRD earth-fixed frame, arbitrary heading reference. Position and linear velocity frame of reference constants
	EstimatorOdometry_LOCAL_FRAME_OTHER uint8 = 2// Not aligned with the std frames of reference. Position and linear velocity frame of reference constants
	EstimatorOdometry_BODY_FRAME_FRD uint8 = 3// FRD body-fixed frame. Position and linear velocity frame of reference constants
)

// Do not create instances of this type directly. Always use NewEstimatorOdometry
// function instead.
type EstimatorOdometry struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). Vehicle odometry data. Fits ROS REP 147 for aerial vehicles
	TimestampSample uint64 `yaml:"timestamp_sample"`
	LocalFrame uint8 `yaml:"local_frame"`// Position and linear velocity local frame of reference
	X float32 `yaml:"x"`// North position. Position in meters. Frame of reference defined by local_frame. NaN if invalid/unknown
	Y float32 `yaml:"y"`// East position. Position in meters. Frame of reference defined by local_frame. NaN if invalid/unknown
	Z float32 `yaml:"z"`// Down position. Position in meters. Frame of reference defined by local_frame. NaN if invalid/unknown
	Q [4]float32 `yaml:"q"`// Quaternion rotation from FRD body frame to refernce frame. Orientation quaternion. First value NaN if invalid/unknown
	QOffset [4]float32 `yaml:"q_offset"`// Quaternion rotation from odometry reference frame to navigation frame. Orientation quaternion. First value NaN if invalid/unknown
	PoseCovariance [21]float32 `yaml:"pose_covariance"`// Row-major representation of 6x6 pose cross-covariance matrix URT.NED earth-fixed frame.Order: x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axisIf position covariance invalid/unknown, first cell is NaNIf orientation covariance invalid/unknown, 16th cell is NaN
	VelocityFrame uint8 `yaml:"velocity_frame"`// Reference frame of the velocity data
	Vx float32 `yaml:"vx"`// North velocity. Velocity in meters/sec. Frame of reference defined by velocity_frame variable. NaN if invalid/unknown
	Vy float32 `yaml:"vy"`// East velocity. Velocity in meters/sec. Frame of reference defined by velocity_frame variable. NaN if invalid/unknown
	Vz float32 `yaml:"vz"`// Down velocity. Velocity in meters/sec. Frame of reference defined by velocity_frame variable. NaN if invalid/unknown
	Rollspeed float32 `yaml:"rollspeed"`// Angular velocity about X body axis. Angular rate in body-fixed frame (rad/s). NaN if invalid/unknown
	Pitchspeed float32 `yaml:"pitchspeed"`// Angular velocity about Y body axis. Angular rate in body-fixed frame (rad/s). NaN if invalid/unknown
	Yawspeed float32 `yaml:"yawspeed"`// Angular velocity about Z body axis. Angular rate in body-fixed frame (rad/s). NaN if invalid/unknown
	VelocityCovariance [21]float32 `yaml:"velocity_covariance"`// Row-major representation of 6x6 velocity cross-covariance matrix URT.Linear velocity in NED earth-fixed frame. Angular velocity in body-fixed frame.Order: vx, vy, vz, rotation rate about X axis, rotation rate about Y axis, rotation rate about Z axisIf linear velocity covariance invalid/unknown, first cell is NaNIf angular velocity covariance invalid/unknown, 16th cell is NaN
}

// NewEstimatorOdometry creates a new EstimatorOdometry with default values.
func NewEstimatorOdometry() *EstimatorOdometry {
	self := EstimatorOdometry{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorOdometry) Clone() *EstimatorOdometry {
	c := &EstimatorOdometry{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.LocalFrame = t.LocalFrame
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	c.Q = t.Q
	c.QOffset = t.QOffset
	c.PoseCovariance = t.PoseCovariance
	c.VelocityFrame = t.VelocityFrame
	c.Vx = t.Vx
	c.Vy = t.Vy
	c.Vz = t.Vz
	c.Rollspeed = t.Rollspeed
	c.Pitchspeed = t.Pitchspeed
	c.Yawspeed = t.Yawspeed
	c.VelocityCovariance = t.VelocityCovariance
	return c
}

func (t *EstimatorOdometry) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorOdometry) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.LocalFrame = 0
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.Q = [4]float32{}
	t.QOffset = [4]float32{}
	t.PoseCovariance = [21]float32{}
	t.VelocityFrame = 0
	t.Vx = 0
	t.Vy = 0
	t.Vz = 0
	t.Rollspeed = 0
	t.Pitchspeed = 0
	t.Yawspeed = 0
	t.VelocityCovariance = [21]float32{}
}

// CloneEstimatorOdometrySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorOdometrySlice(dst, src []EstimatorOdometry) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorOdometryTypeSupport types.MessageTypeSupport = _EstimatorOdometryTypeSupport{}

type _EstimatorOdometryTypeSupport struct{}

func (t _EstimatorOdometryTypeSupport) New() types.Message {
	return NewEstimatorOdometry()
}

func (t _EstimatorOdometryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorOdometry
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorOdometry__create())
}

func (t _EstimatorOdometryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorOdometry__destroy((*C.px4_msgs__msg__EstimatorOdometry)(pointer_to_free))
}

func (t _EstimatorOdometryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorOdometry)
	mem := (*C.px4_msgs__msg__EstimatorOdometry)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.local_frame = C.uint8_t(m.LocalFrame)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	cSlice_q_offset := mem.q_offset[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q_offset)), m.QOffset[:])
	cSlice_pose_covariance := mem.pose_covariance[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_pose_covariance)), m.PoseCovariance[:])
	mem.velocity_frame = C.uint8_t(m.VelocityFrame)
	mem.vx = C.float(m.Vx)
	mem.vy = C.float(m.Vy)
	mem.vz = C.float(m.Vz)
	mem.rollspeed = C.float(m.Rollspeed)
	mem.pitchspeed = C.float(m.Pitchspeed)
	mem.yawspeed = C.float(m.Yawspeed)
	cSlice_velocity_covariance := mem.velocity_covariance[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_velocity_covariance)), m.VelocityCovariance[:])
}

func (t _EstimatorOdometryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorOdometry)
	mem := (*C.px4_msgs__msg__EstimatorOdometry)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.LocalFrame = uint8(mem.local_frame)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_Go(m.Q[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)))
	cSlice_q_offset := mem.q_offset[:]
	primitives.Float32__Array_to_Go(m.QOffset[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q_offset)))
	cSlice_pose_covariance := mem.pose_covariance[:]
	primitives.Float32__Array_to_Go(m.PoseCovariance[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_pose_covariance)))
	m.VelocityFrame = uint8(mem.velocity_frame)
	m.Vx = float32(mem.vx)
	m.Vy = float32(mem.vy)
	m.Vz = float32(mem.vz)
	m.Rollspeed = float32(mem.rollspeed)
	m.Pitchspeed = float32(mem.pitchspeed)
	m.Yawspeed = float32(mem.yawspeed)
	cSlice_velocity_covariance := mem.velocity_covariance[:]
	primitives.Float32__Array_to_Go(m.VelocityCovariance[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_velocity_covariance)))
}

func (t _EstimatorOdometryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorOdometry())
}

type CEstimatorOdometry = C.px4_msgs__msg__EstimatorOdometry
type CEstimatorOdometry__Sequence = C.px4_msgs__msg__EstimatorOdometry__Sequence

func EstimatorOdometry__Sequence_to_Go(goSlice *[]EstimatorOdometry, cSlice CEstimatorOdometry__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorOdometry, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorOdometry__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorOdometry * uintptr(i)),
		))
		EstimatorOdometryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorOdometry__Sequence_to_C(cSlice *CEstimatorOdometry__Sequence, goSlice []EstimatorOdometry) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorOdometry)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorOdometry * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorOdometry)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorOdometry * uintptr(i)),
		))
		EstimatorOdometryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorOdometry__Array_to_Go(goSlice []EstimatorOdometry, cSlice []CEstimatorOdometry) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorOdometryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorOdometry__Array_to_C(cSlice []CEstimatorOdometry, goSlice []EstimatorOdometry) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorOdometryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
