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
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/estimator_status_flags.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorStatusFlags", EstimatorStatusFlagsTypeSupport)
}

// Do not create instances of this type directly. Always use NewEstimatorStatusFlags
// function instead.
type EstimatorStatusFlags struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	ControlStatusChanges uint32 `yaml:"control_status_changes"`// number of filter control status (cs) changes. filter control status
	CsTiltAlign bool `yaml:"cs_tilt_align"`// 0 - true if the filter tilt alignment is complete. filter control status
	CsYawAlign bool `yaml:"cs_yaw_align"`// 1 - true if the filter yaw alignment is complete. filter control status
	CsGps bool `yaml:"cs_gps"`// 2 - true if GPS measurement fusion is intended. filter control status
	CsOptFlow bool `yaml:"cs_opt_flow"`// 3 - true if optical flow measurements fusion is intended. filter control status
	CsMagHdg bool `yaml:"cs_mag_hdg"`// 4 - true if a simple magnetic yaw heading fusion is intended. filter control status
	CsMag3d bool `yaml:"cs_mag_3d"`// 5 - true if 3-axis magnetometer measurement fusion is inteded. filter control status
	CsMagDec bool `yaml:"cs_mag_dec"`// 6 - true if synthetic magnetic declination measurements fusion is intended. filter control status
	CsInAir bool `yaml:"cs_in_air"`// 7 - true when the vehicle is airborne. filter control status
	CsWind bool `yaml:"cs_wind"`// 8 - true when wind velocity is being estimated. filter control status
	CsBaroHgt bool `yaml:"cs_baro_hgt"`// 9 - true when baro height is being fused as a primary height reference. filter control status
	CsRngHgt bool `yaml:"cs_rng_hgt"`// 10 - true when range finder height is being fused as a primary height reference. filter control status
	CsGpsHgt bool `yaml:"cs_gps_hgt"`// 11 - true when GPS height is being fused as a primary height reference. filter control status
	CsEvPos bool `yaml:"cs_ev_pos"`// 12 - true when local position data fusion from external vision is intended. filter control status
	CsEvYaw bool `yaml:"cs_ev_yaw"`// 13 - true when yaw data from external vision measurements fusion is intended. filter control status
	CsEvHgt bool `yaml:"cs_ev_hgt"`// 14 - true when height data from external vision measurements is being fused. filter control status
	CsFuseBeta bool `yaml:"cs_fuse_beta"`// 15 - true when synthetic sideslip measurements are being fused. filter control status
	CsMagFieldDisturbed bool `yaml:"cs_mag_field_disturbed"`// 16 - true when the mag field does not match the expected strength. filter control status
	CsFixedWing bool `yaml:"cs_fixed_wing"`// 17 - true when the vehicle is operating as a fixed wing vehicle. filter control status
	CsMagFault bool `yaml:"cs_mag_fault"`// 18 - true when the magnetometer has been declared faulty and is no longer being used. filter control status
	CsFuseAspd bool `yaml:"cs_fuse_aspd"`// 19 - true when airspeed measurements are being fused. filter control status
	CsGndEffect bool `yaml:"cs_gnd_effect"`// 20 - true when protection from ground effect induced static pressure rise is active. filter control status
	CsRngStuck bool `yaml:"cs_rng_stuck"`// 21 - true when rng data wasn't ready for more than 10s and new rng values haven't changed enough. filter control status
	CsGpsYaw bool `yaml:"cs_gps_yaw"`// 22 - true when yaw (not ground course) data fusion from a GPS receiver is intended. filter control status
	CsMagAlignedInFlight bool `yaml:"cs_mag_aligned_in_flight"`// 23 - true when the in-flight mag field alignment has been completed. filter control status
	CsEvVel bool `yaml:"cs_ev_vel"`// 24 - true when local frame velocity data fusion from external vision measurements is intended. filter control status
	CsSyntheticMagZ bool `yaml:"cs_synthetic_mag_z"`// 25 - true when we are using a synthesized measurement for the magnetometer Z component. filter control status
	CsVehicleAtRest bool `yaml:"cs_vehicle_at_rest"`// 26 - true when the vehicle is at rest. filter control status
	FaultStatusChanges uint32 `yaml:"fault_status_changes"`// number of filter fault status (fs) changes. fault status
	FsBadMagX bool `yaml:"fs_bad_mag_x"`// 0 - true if the fusion of the magnetometer X-axis has encountered a numerical error. fault status
	FsBadMagY bool `yaml:"fs_bad_mag_y"`// 1 - true if the fusion of the magnetometer Y-axis has encountered a numerical error. fault status
	FsBadMagZ bool `yaml:"fs_bad_mag_z"`// 2 - true if the fusion of the magnetometer Z-axis has encountered a numerical error. fault status
	FsBadHdg bool `yaml:"fs_bad_hdg"`// 3 - true if the fusion of the heading angle has encountered a numerical error. fault status
	FsBadMagDecl bool `yaml:"fs_bad_mag_decl"`// 4 - true if the fusion of the magnetic declination has encountered a numerical error. fault status
	FsBadAirspeed bool `yaml:"fs_bad_airspeed"`// 5 - true if fusion of the airspeed has encountered a numerical error. fault status
	FsBadSideslip bool `yaml:"fs_bad_sideslip"`// 6 - true if fusion of the synthetic sideslip constraint has encountered a numerical error. fault status
	FsBadOptflowX bool `yaml:"fs_bad_optflow_x"`// 7 - true if fusion of the optical flow X axis has encountered a numerical error. fault status
	FsBadOptflowY bool `yaml:"fs_bad_optflow_y"`// 8 - true if fusion of the optical flow Y axis has encountered a numerical error. fault status
	FsBadVelN bool `yaml:"fs_bad_vel_n"`// 9 - true if fusion of the North velocity has encountered a numerical error. fault status
	FsBadVelE bool `yaml:"fs_bad_vel_e"`// 10 - true if fusion of the East velocity has encountered a numerical error. fault status
	FsBadVelD bool `yaml:"fs_bad_vel_d"`// 11 - true if fusion of the Down velocity has encountered a numerical error. fault status
	FsBadPosN bool `yaml:"fs_bad_pos_n"`// 12 - true if fusion of the North position has encountered a numerical error. fault status
	FsBadPosE bool `yaml:"fs_bad_pos_e"`// 13 - true if fusion of the East position has encountered a numerical error. fault status
	FsBadPosD bool `yaml:"fs_bad_pos_d"`// 14 - true if fusion of the Down position has encountered a numerical error. fault status
	FsBadAccBias bool `yaml:"fs_bad_acc_bias"`// 15 - true if bad delta velocity bias estimates have been detected. fault status
	FsBadAccVertical bool `yaml:"fs_bad_acc_vertical"`// 16 - true if bad vertical accelerometer data has been detected. fault status
	FsBadAccClipping bool `yaml:"fs_bad_acc_clipping"`// 17 - true if delta velocity data contains clipping (asymmetric railing). fault status
	InnovationFaultStatusChanges uint32 `yaml:"innovation_fault_status_changes"`// number of innovation fault status (reject) changes. innovation test failures
	RejectHorVel bool `yaml:"reject_hor_vel"`// 0 - true if horizontal velocity observations have been rejected. innovation test failures
	RejectVerVel bool `yaml:"reject_ver_vel"`// 1 - true if vertical velocity observations have been rejected. innovation test failures
	RejectHorPos bool `yaml:"reject_hor_pos"`// 2 - true if horizontal position observations have been rejected. innovation test failures
	RejectVerPos bool `yaml:"reject_ver_pos"`// 3 - true if vertical position observations have been rejected. innovation test failures
	RejectMagX bool `yaml:"reject_mag_x"`// 4 - true if the X magnetometer observation has been rejected. innovation test failures
	RejectMagY bool `yaml:"reject_mag_y"`// 5 - true if the Y magnetometer observation has been rejected. innovation test failures
	RejectMagZ bool `yaml:"reject_mag_z"`// 6 - true if the Z magnetometer observation has been rejected. innovation test failures
	RejectYaw bool `yaml:"reject_yaw"`// 7 - true if the yaw observation has been rejected. innovation test failures
	RejectAirspeed bool `yaml:"reject_airspeed"`// 8 - true if the airspeed observation has been rejected. innovation test failures
	RejectSideslip bool `yaml:"reject_sideslip"`// 9 - true if the synthetic sideslip observation has been rejected. innovation test failures
	RejectHagl bool `yaml:"reject_hagl"`// 10 - true if the height above ground observation has been rejected. innovation test failures
	RejectOptflowX bool `yaml:"reject_optflow_x"`// 11 - true if the X optical flow observation has been rejected. innovation test failures
	RejectOptflowY bool `yaml:"reject_optflow_y"`// 12 - true if the Y optical flow observation has been rejected. innovation test failures
}

// NewEstimatorStatusFlags creates a new EstimatorStatusFlags with default values.
func NewEstimatorStatusFlags() *EstimatorStatusFlags {
	self := EstimatorStatusFlags{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorStatusFlags) Clone() *EstimatorStatusFlags {
	c := &EstimatorStatusFlags{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.ControlStatusChanges = t.ControlStatusChanges
	c.CsTiltAlign = t.CsTiltAlign
	c.CsYawAlign = t.CsYawAlign
	c.CsGps = t.CsGps
	c.CsOptFlow = t.CsOptFlow
	c.CsMagHdg = t.CsMagHdg
	c.CsMag3d = t.CsMag3d
	c.CsMagDec = t.CsMagDec
	c.CsInAir = t.CsInAir
	c.CsWind = t.CsWind
	c.CsBaroHgt = t.CsBaroHgt
	c.CsRngHgt = t.CsRngHgt
	c.CsGpsHgt = t.CsGpsHgt
	c.CsEvPos = t.CsEvPos
	c.CsEvYaw = t.CsEvYaw
	c.CsEvHgt = t.CsEvHgt
	c.CsFuseBeta = t.CsFuseBeta
	c.CsMagFieldDisturbed = t.CsMagFieldDisturbed
	c.CsFixedWing = t.CsFixedWing
	c.CsMagFault = t.CsMagFault
	c.CsFuseAspd = t.CsFuseAspd
	c.CsGndEffect = t.CsGndEffect
	c.CsRngStuck = t.CsRngStuck
	c.CsGpsYaw = t.CsGpsYaw
	c.CsMagAlignedInFlight = t.CsMagAlignedInFlight
	c.CsEvVel = t.CsEvVel
	c.CsSyntheticMagZ = t.CsSyntheticMagZ
	c.CsVehicleAtRest = t.CsVehicleAtRest
	c.FaultStatusChanges = t.FaultStatusChanges
	c.FsBadMagX = t.FsBadMagX
	c.FsBadMagY = t.FsBadMagY
	c.FsBadMagZ = t.FsBadMagZ
	c.FsBadHdg = t.FsBadHdg
	c.FsBadMagDecl = t.FsBadMagDecl
	c.FsBadAirspeed = t.FsBadAirspeed
	c.FsBadSideslip = t.FsBadSideslip
	c.FsBadOptflowX = t.FsBadOptflowX
	c.FsBadOptflowY = t.FsBadOptflowY
	c.FsBadVelN = t.FsBadVelN
	c.FsBadVelE = t.FsBadVelE
	c.FsBadVelD = t.FsBadVelD
	c.FsBadPosN = t.FsBadPosN
	c.FsBadPosE = t.FsBadPosE
	c.FsBadPosD = t.FsBadPosD
	c.FsBadAccBias = t.FsBadAccBias
	c.FsBadAccVertical = t.FsBadAccVertical
	c.FsBadAccClipping = t.FsBadAccClipping
	c.InnovationFaultStatusChanges = t.InnovationFaultStatusChanges
	c.RejectHorVel = t.RejectHorVel
	c.RejectVerVel = t.RejectVerVel
	c.RejectHorPos = t.RejectHorPos
	c.RejectVerPos = t.RejectVerPos
	c.RejectMagX = t.RejectMagX
	c.RejectMagY = t.RejectMagY
	c.RejectMagZ = t.RejectMagZ
	c.RejectYaw = t.RejectYaw
	c.RejectAirspeed = t.RejectAirspeed
	c.RejectSideslip = t.RejectSideslip
	c.RejectHagl = t.RejectHagl
	c.RejectOptflowX = t.RejectOptflowX
	c.RejectOptflowY = t.RejectOptflowY
	return c
}

func (t *EstimatorStatusFlags) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorStatusFlags) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.ControlStatusChanges = 0
	t.CsTiltAlign = false
	t.CsYawAlign = false
	t.CsGps = false
	t.CsOptFlow = false
	t.CsMagHdg = false
	t.CsMag3d = false
	t.CsMagDec = false
	t.CsInAir = false
	t.CsWind = false
	t.CsBaroHgt = false
	t.CsRngHgt = false
	t.CsGpsHgt = false
	t.CsEvPos = false
	t.CsEvYaw = false
	t.CsEvHgt = false
	t.CsFuseBeta = false
	t.CsMagFieldDisturbed = false
	t.CsFixedWing = false
	t.CsMagFault = false
	t.CsFuseAspd = false
	t.CsGndEffect = false
	t.CsRngStuck = false
	t.CsGpsYaw = false
	t.CsMagAlignedInFlight = false
	t.CsEvVel = false
	t.CsSyntheticMagZ = false
	t.CsVehicleAtRest = false
	t.FaultStatusChanges = 0
	t.FsBadMagX = false
	t.FsBadMagY = false
	t.FsBadMagZ = false
	t.FsBadHdg = false
	t.FsBadMagDecl = false
	t.FsBadAirspeed = false
	t.FsBadSideslip = false
	t.FsBadOptflowX = false
	t.FsBadOptflowY = false
	t.FsBadVelN = false
	t.FsBadVelE = false
	t.FsBadVelD = false
	t.FsBadPosN = false
	t.FsBadPosE = false
	t.FsBadPosD = false
	t.FsBadAccBias = false
	t.FsBadAccVertical = false
	t.FsBadAccClipping = false
	t.InnovationFaultStatusChanges = 0
	t.RejectHorVel = false
	t.RejectVerVel = false
	t.RejectHorPos = false
	t.RejectVerPos = false
	t.RejectMagX = false
	t.RejectMagY = false
	t.RejectMagZ = false
	t.RejectYaw = false
	t.RejectAirspeed = false
	t.RejectSideslip = false
	t.RejectHagl = false
	t.RejectOptflowX = false
	t.RejectOptflowY = false
}

// CloneEstimatorStatusFlagsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorStatusFlagsSlice(dst, src []EstimatorStatusFlags) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorStatusFlagsTypeSupport types.MessageTypeSupport = _EstimatorStatusFlagsTypeSupport{}

type _EstimatorStatusFlagsTypeSupport struct{}

func (t _EstimatorStatusFlagsTypeSupport) New() types.Message {
	return NewEstimatorStatusFlags()
}

func (t _EstimatorStatusFlagsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorStatusFlags
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorStatusFlags__create())
}

func (t _EstimatorStatusFlagsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorStatusFlags__destroy((*C.px4_msgs__msg__EstimatorStatusFlags)(pointer_to_free))
}

func (t _EstimatorStatusFlagsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorStatusFlags)
	mem := (*C.px4_msgs__msg__EstimatorStatusFlags)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.control_status_changes = C.uint32_t(m.ControlStatusChanges)
	mem.cs_tilt_align = C.bool(m.CsTiltAlign)
	mem.cs_yaw_align = C.bool(m.CsYawAlign)
	mem.cs_gps = C.bool(m.CsGps)
	mem.cs_opt_flow = C.bool(m.CsOptFlow)
	mem.cs_mag_hdg = C.bool(m.CsMagHdg)
	mem.cs_mag_3d = C.bool(m.CsMag3d)
	mem.cs_mag_dec = C.bool(m.CsMagDec)
	mem.cs_in_air = C.bool(m.CsInAir)
	mem.cs_wind = C.bool(m.CsWind)
	mem.cs_baro_hgt = C.bool(m.CsBaroHgt)
	mem.cs_rng_hgt = C.bool(m.CsRngHgt)
	mem.cs_gps_hgt = C.bool(m.CsGpsHgt)
	mem.cs_ev_pos = C.bool(m.CsEvPos)
	mem.cs_ev_yaw = C.bool(m.CsEvYaw)
	mem.cs_ev_hgt = C.bool(m.CsEvHgt)
	mem.cs_fuse_beta = C.bool(m.CsFuseBeta)
	mem.cs_mag_field_disturbed = C.bool(m.CsMagFieldDisturbed)
	mem.cs_fixed_wing = C.bool(m.CsFixedWing)
	mem.cs_mag_fault = C.bool(m.CsMagFault)
	mem.cs_fuse_aspd = C.bool(m.CsFuseAspd)
	mem.cs_gnd_effect = C.bool(m.CsGndEffect)
	mem.cs_rng_stuck = C.bool(m.CsRngStuck)
	mem.cs_gps_yaw = C.bool(m.CsGpsYaw)
	mem.cs_mag_aligned_in_flight = C.bool(m.CsMagAlignedInFlight)
	mem.cs_ev_vel = C.bool(m.CsEvVel)
	mem.cs_synthetic_mag_z = C.bool(m.CsSyntheticMagZ)
	mem.cs_vehicle_at_rest = C.bool(m.CsVehicleAtRest)
	mem.fault_status_changes = C.uint32_t(m.FaultStatusChanges)
	mem.fs_bad_mag_x = C.bool(m.FsBadMagX)
	mem.fs_bad_mag_y = C.bool(m.FsBadMagY)
	mem.fs_bad_mag_z = C.bool(m.FsBadMagZ)
	mem.fs_bad_hdg = C.bool(m.FsBadHdg)
	mem.fs_bad_mag_decl = C.bool(m.FsBadMagDecl)
	mem.fs_bad_airspeed = C.bool(m.FsBadAirspeed)
	mem.fs_bad_sideslip = C.bool(m.FsBadSideslip)
	mem.fs_bad_optflow_x = C.bool(m.FsBadOptflowX)
	mem.fs_bad_optflow_y = C.bool(m.FsBadOptflowY)
	mem.fs_bad_vel_n = C.bool(m.FsBadVelN)
	mem.fs_bad_vel_e = C.bool(m.FsBadVelE)
	mem.fs_bad_vel_d = C.bool(m.FsBadVelD)
	mem.fs_bad_pos_n = C.bool(m.FsBadPosN)
	mem.fs_bad_pos_e = C.bool(m.FsBadPosE)
	mem.fs_bad_pos_d = C.bool(m.FsBadPosD)
	mem.fs_bad_acc_bias = C.bool(m.FsBadAccBias)
	mem.fs_bad_acc_vertical = C.bool(m.FsBadAccVertical)
	mem.fs_bad_acc_clipping = C.bool(m.FsBadAccClipping)
	mem.innovation_fault_status_changes = C.uint32_t(m.InnovationFaultStatusChanges)
	mem.reject_hor_vel = C.bool(m.RejectHorVel)
	mem.reject_ver_vel = C.bool(m.RejectVerVel)
	mem.reject_hor_pos = C.bool(m.RejectHorPos)
	mem.reject_ver_pos = C.bool(m.RejectVerPos)
	mem.reject_mag_x = C.bool(m.RejectMagX)
	mem.reject_mag_y = C.bool(m.RejectMagY)
	mem.reject_mag_z = C.bool(m.RejectMagZ)
	mem.reject_yaw = C.bool(m.RejectYaw)
	mem.reject_airspeed = C.bool(m.RejectAirspeed)
	mem.reject_sideslip = C.bool(m.RejectSideslip)
	mem.reject_hagl = C.bool(m.RejectHagl)
	mem.reject_optflow_x = C.bool(m.RejectOptflowX)
	mem.reject_optflow_y = C.bool(m.RejectOptflowY)
}

func (t _EstimatorStatusFlagsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorStatusFlags)
	mem := (*C.px4_msgs__msg__EstimatorStatusFlags)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.ControlStatusChanges = uint32(mem.control_status_changes)
	m.CsTiltAlign = bool(mem.cs_tilt_align)
	m.CsYawAlign = bool(mem.cs_yaw_align)
	m.CsGps = bool(mem.cs_gps)
	m.CsOptFlow = bool(mem.cs_opt_flow)
	m.CsMagHdg = bool(mem.cs_mag_hdg)
	m.CsMag3d = bool(mem.cs_mag_3d)
	m.CsMagDec = bool(mem.cs_mag_dec)
	m.CsInAir = bool(mem.cs_in_air)
	m.CsWind = bool(mem.cs_wind)
	m.CsBaroHgt = bool(mem.cs_baro_hgt)
	m.CsRngHgt = bool(mem.cs_rng_hgt)
	m.CsGpsHgt = bool(mem.cs_gps_hgt)
	m.CsEvPos = bool(mem.cs_ev_pos)
	m.CsEvYaw = bool(mem.cs_ev_yaw)
	m.CsEvHgt = bool(mem.cs_ev_hgt)
	m.CsFuseBeta = bool(mem.cs_fuse_beta)
	m.CsMagFieldDisturbed = bool(mem.cs_mag_field_disturbed)
	m.CsFixedWing = bool(mem.cs_fixed_wing)
	m.CsMagFault = bool(mem.cs_mag_fault)
	m.CsFuseAspd = bool(mem.cs_fuse_aspd)
	m.CsGndEffect = bool(mem.cs_gnd_effect)
	m.CsRngStuck = bool(mem.cs_rng_stuck)
	m.CsGpsYaw = bool(mem.cs_gps_yaw)
	m.CsMagAlignedInFlight = bool(mem.cs_mag_aligned_in_flight)
	m.CsEvVel = bool(mem.cs_ev_vel)
	m.CsSyntheticMagZ = bool(mem.cs_synthetic_mag_z)
	m.CsVehicleAtRest = bool(mem.cs_vehicle_at_rest)
	m.FaultStatusChanges = uint32(mem.fault_status_changes)
	m.FsBadMagX = bool(mem.fs_bad_mag_x)
	m.FsBadMagY = bool(mem.fs_bad_mag_y)
	m.FsBadMagZ = bool(mem.fs_bad_mag_z)
	m.FsBadHdg = bool(mem.fs_bad_hdg)
	m.FsBadMagDecl = bool(mem.fs_bad_mag_decl)
	m.FsBadAirspeed = bool(mem.fs_bad_airspeed)
	m.FsBadSideslip = bool(mem.fs_bad_sideslip)
	m.FsBadOptflowX = bool(mem.fs_bad_optflow_x)
	m.FsBadOptflowY = bool(mem.fs_bad_optflow_y)
	m.FsBadVelN = bool(mem.fs_bad_vel_n)
	m.FsBadVelE = bool(mem.fs_bad_vel_e)
	m.FsBadVelD = bool(mem.fs_bad_vel_d)
	m.FsBadPosN = bool(mem.fs_bad_pos_n)
	m.FsBadPosE = bool(mem.fs_bad_pos_e)
	m.FsBadPosD = bool(mem.fs_bad_pos_d)
	m.FsBadAccBias = bool(mem.fs_bad_acc_bias)
	m.FsBadAccVertical = bool(mem.fs_bad_acc_vertical)
	m.FsBadAccClipping = bool(mem.fs_bad_acc_clipping)
	m.InnovationFaultStatusChanges = uint32(mem.innovation_fault_status_changes)
	m.RejectHorVel = bool(mem.reject_hor_vel)
	m.RejectVerVel = bool(mem.reject_ver_vel)
	m.RejectHorPos = bool(mem.reject_hor_pos)
	m.RejectVerPos = bool(mem.reject_ver_pos)
	m.RejectMagX = bool(mem.reject_mag_x)
	m.RejectMagY = bool(mem.reject_mag_y)
	m.RejectMagZ = bool(mem.reject_mag_z)
	m.RejectYaw = bool(mem.reject_yaw)
	m.RejectAirspeed = bool(mem.reject_airspeed)
	m.RejectSideslip = bool(mem.reject_sideslip)
	m.RejectHagl = bool(mem.reject_hagl)
	m.RejectOptflowX = bool(mem.reject_optflow_x)
	m.RejectOptflowY = bool(mem.reject_optflow_y)
}

func (t _EstimatorStatusFlagsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorStatusFlags())
}

type CEstimatorStatusFlags = C.px4_msgs__msg__EstimatorStatusFlags
type CEstimatorStatusFlags__Sequence = C.px4_msgs__msg__EstimatorStatusFlags__Sequence

func EstimatorStatusFlags__Sequence_to_Go(goSlice *[]EstimatorStatusFlags, cSlice CEstimatorStatusFlags__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorStatusFlags, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorStatusFlags__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorStatusFlags * uintptr(i)),
		))
		EstimatorStatusFlagsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorStatusFlags__Sequence_to_C(cSlice *CEstimatorStatusFlags__Sequence, goSlice []EstimatorStatusFlags) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorStatusFlags)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorStatusFlags * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorStatusFlags)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorStatusFlags * uintptr(i)),
		))
		EstimatorStatusFlagsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorStatusFlags__Array_to_Go(goSlice []EstimatorStatusFlags, cSlice []CEstimatorStatusFlags) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorStatusFlagsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorStatusFlags__Array_to_C(cSlice []CEstimatorStatusFlags, goSlice []EstimatorStatusFlags) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorStatusFlagsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
