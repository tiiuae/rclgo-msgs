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

#include <px4_msgs/msg/vehicle_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleStatus", VehicleStatusTypeSupport)
}
const (
	VehicleStatus_ARMING_STATE_INIT uint8 = 0
	VehicleStatus_ARMING_STATE_STANDBY uint8 = 1
	VehicleStatus_ARMING_STATE_ARMED uint8 = 2
	VehicleStatus_ARMING_STATE_STANDBY_ERROR uint8 = 3
	VehicleStatus_ARMING_STATE_SHUTDOWN uint8 = 4
	VehicleStatus_ARMING_STATE_IN_AIR_RESTORE uint8 = 5
	VehicleStatus_ARMING_STATE_MAX uint8 = 6
	VehicleStatus_FAILURE_NONE uint8 = 0// FailureDetector status
	VehicleStatus_FAILURE_ROLL uint8 = 1// (1 << 0). FailureDetector status
	VehicleStatus_FAILURE_PITCH uint8 = 2// (1 << 1). FailureDetector status
	VehicleStatus_FAILURE_ALT uint8 = 4// (1 << 2). FailureDetector status
	VehicleStatus_FAILURE_EXT uint8 = 8// (1 << 3). FailureDetector status
	VehicleStatus_FAILURE_ARM_ESC uint8 = 16// (1 << 4). FailureDetector status
	VehicleStatus_HIL_STATE_OFF uint8 = 0// HIL
	VehicleStatus_HIL_STATE_ON uint8 = 1// HIL
	VehicleStatus_NAVIGATION_STATE_MANUAL uint8 = 0// Manual mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_ALTCTL uint8 = 1// Altitude control mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_POSCTL uint8 = 2// Position control mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_MISSION uint8 = 3// Auto mission mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_LOITER uint8 = 4// Auto loiter mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_RTL uint8 = 5// Auto return to launch mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_LANDENGFAIL uint8 = 8// Auto land on engine failure. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_LANDGPSFAIL uint8 = 9// Auto land on gps failure (e.g. open loop loiter down). Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_ACRO uint8 = 10// Acro mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_UNUSED uint8 = 11// Free slot. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_DESCEND uint8 = 12// Descend mode (no position control). Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_TERMINATION uint8 = 13// Termination mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_OFFBOARD uint8 = 14// Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_STAB uint8 = 15// Stabilized mode. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_UNUSED2 uint8 = 16// Free slot. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_TAKEOFF uint8 = 17// Takeoff. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_LAND uint8 = 18// Land. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_FOLLOW_TARGET uint8 = 19// Auto Follow. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_AUTO_PRECLAND uint8 = 20// Precision land with landing target. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_ORBIT uint8 = 21// Orbit in a circle. Navigation state, i.e. "what should vehicle do".
	VehicleStatus_NAVIGATION_STATE_MAX uint8 = 22// Navigation state, i.e. "what should vehicle do".
	VehicleStatus_RC_IN_MODE_DEFAULT uint8 = 0
	VehicleStatus_RC_IN_MODE_OFF uint8 = 1
	VehicleStatus_RC_IN_MODE_GENERATED uint8 = 2
	VehicleStatus_VEHICLE_TYPE_UNKNOWN uint8 = 0
	VehicleStatus_VEHICLE_TYPE_ROTARY_WING uint8 = 1
	VehicleStatus_VEHICLE_TYPE_FIXED_WING uint8 = 2
	VehicleStatus_VEHICLE_TYPE_ROVER uint8 = 3
	VehicleStatus_VEHICLE_TYPE_AIRSHIP uint8 = 4
	VehicleStatus_ARM_DISARM_REASON_TRANSITION_TO_STANDBY uint8 = 0
	VehicleStatus_ARM_DISARM_REASON_RC_STICK uint8 = 1
	VehicleStatus_ARM_DISARM_REASON_RC_SWITCH uint8 = 2
	VehicleStatus_ARM_DISARM_REASON_COMMAND_INTERNAL uint8 = 3
	VehicleStatus_ARM_DISARM_REASON_COMMAND_EXTERNAL uint8 = 4
	VehicleStatus_ARM_DISARM_REASON_MISSION_START uint8 = 5
	VehicleStatus_ARM_DISARM_REASON_SAFETY_BUTTON uint8 = 6
	VehicleStatus_ARM_DISARM_REASON_AUTO_DISARM_LAND uint8 = 7
	VehicleStatus_ARM_DISARM_REASON_AUTO_DISARM_PREFLIGHT uint8 = 8
	VehicleStatus_ARM_DISARM_REASON_KILL_SWITCH uint8 = 9
	VehicleStatus_ARM_DISARM_REASON_LOCKDOWN uint8 = 10
	VehicleStatus_ARM_DISARM_REASON_FAILURE_DETECTOR uint8 = 11
	VehicleStatus_ARM_DISARM_REASON_SHUTDOWN uint8 = 12
	VehicleStatus_ARM_DISARM_REASON_UNIT_TEST uint8 = 13
)

// Do not create instances of this type directly. Always use NewVehicleStatus
// function instead.
type VehicleStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). If you change the order, add or remove arming_state_t states make sure to update the arraysin state_machine_helper.cpp as well.
	NavState uint8 `yaml:"nav_state"`// set navigation state machine to specified value
	NavStateTimestamp uint64 `yaml:"nav_state_timestamp"`// time when current nav_state activated
	ArmingState uint8 `yaml:"arming_state"`// current arming state
	HilState uint8 `yaml:"hil_state"`// current hil state
	Failsafe bool `yaml:"failsafe"`// true if system is in failsafe state (e.g.:RTL, Hover, Terminate, ...)
	FailsafeTimestamp uint64 `yaml:"failsafe_timestamp"`// time when failsafe was activated
	SystemType uint8 `yaml:"system_type"`// system type, contains mavlink MAV_TYPE
	SystemId uint8 `yaml:"system_id"`// system id, contains MAVLink's system ID field
	ComponentId uint8 `yaml:"component_id"`// subsystem / component id, contains MAVLink's component ID field
	VehicleType uint8 `yaml:"vehicle_type"`// Type of vehicle (fixed-wing, rotary wing, ground)
	IsVtol bool `yaml:"is_vtol"`// True if the system is VTOL capable
	IsVtolTailsitter bool `yaml:"is_vtol_tailsitter"`// True if the system performs a 90° pitch down rotation during transition from MC to FW
	VtolFwPermanentStab bool `yaml:"vtol_fw_permanent_stab"`// True if VTOL should stabilize attitude for fw in manual mode
	InTransitionMode bool `yaml:"in_transition_mode"`// True if VTOL is doing a transition
	InTransitionToFw bool `yaml:"in_transition_to_fw"`// True if VTOL is doing a transition from MC to FW
	RcSignalLost bool `yaml:"rc_signal_lost"`// true if RC reception lost
	RcInputMode uint8 `yaml:"rc_input_mode"`// set to 1 to disable the RC input, 2 to enable manual control to RC in mapping.
	DataLinkLost bool `yaml:"data_link_lost"`// datalink to GCS lost
	DataLinkLostCounter uint8 `yaml:"data_link_lost_counter"`// counts unique data link lost events
	HighLatencyDataLinkLost bool `yaml:"high_latency_data_link_lost"`// Set to true if the high latency data link (eg. RockBlock Iridium 9603 telemetry module) is lost
	EngineFailure bool `yaml:"engine_failure"`// Set to true if an engine failure is detected
	MissionFailure bool `yaml:"mission_failure"`// Set to true if mission could not continue/finish
	FailureDetectorStatus uint8 `yaml:"failure_detector_status"`// Bitmask containing FailureDetector status [0, 0, 0, 0, 0, FAILURE_ALT, FAILURE_PITCH, FAILURE_ROLL]
	OnboardControlSensorsPresent uint32 `yaml:"onboard_control_sensors_present"`// see SYS_STATUS mavlink message for the following
	OnboardControlSensorsEnabled uint32 `yaml:"onboard_control_sensors_enabled"`// see SYS_STATUS mavlink message for the following
	OnboardControlSensorsHealth uint32 `yaml:"onboard_control_sensors_health"`// see SYS_STATUS mavlink message for the following
	LatestArmingReason uint8 `yaml:"latest_arming_reason"`
	LatestDisarmingReason uint8 `yaml:"latest_disarming_reason"`
	ArmedTime uint64 `yaml:"armed_time"`// Arming timestamp (microseconds)
	TakeoffTime uint64 `yaml:"takeoff_time"`// Takeoff timestamp (microseconds)
}

// NewVehicleStatus creates a new VehicleStatus with default values.
func NewVehicleStatus() *VehicleStatus {
	self := VehicleStatus{}
	self.SetDefaults()
	return &self
}

func (t *VehicleStatus) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleStatus) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleStatusTypeSupport types.MessageTypeSupport = _VehicleStatusTypeSupport{}

type _VehicleStatusTypeSupport struct{}

func (t _VehicleStatusTypeSupport) New() types.Message {
	return NewVehicleStatus()
}

func (t _VehicleStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleStatus__create())
}

func (t _VehicleStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleStatus__destroy((*C.px4_msgs__msg__VehicleStatus)(pointer_to_free))
}

func (t _VehicleStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleStatus)
	mem := (*C.px4_msgs__msg__VehicleStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.nav_state = C.uint8_t(m.NavState)
	mem.nav_state_timestamp = C.uint64_t(m.NavStateTimestamp)
	mem.arming_state = C.uint8_t(m.ArmingState)
	mem.hil_state = C.uint8_t(m.HilState)
	mem.failsafe = C.bool(m.Failsafe)
	mem.failsafe_timestamp = C.uint64_t(m.FailsafeTimestamp)
	mem.system_type = C.uint8_t(m.SystemType)
	mem.system_id = C.uint8_t(m.SystemId)
	mem.component_id = C.uint8_t(m.ComponentId)
	mem.vehicle_type = C.uint8_t(m.VehicleType)
	mem.is_vtol = C.bool(m.IsVtol)
	mem.is_vtol_tailsitter = C.bool(m.IsVtolTailsitter)
	mem.vtol_fw_permanent_stab = C.bool(m.VtolFwPermanentStab)
	mem.in_transition_mode = C.bool(m.InTransitionMode)
	mem.in_transition_to_fw = C.bool(m.InTransitionToFw)
	mem.rc_signal_lost = C.bool(m.RcSignalLost)
	mem.rc_input_mode = C.uint8_t(m.RcInputMode)
	mem.data_link_lost = C.bool(m.DataLinkLost)
	mem.data_link_lost_counter = C.uint8_t(m.DataLinkLostCounter)
	mem.high_latency_data_link_lost = C.bool(m.HighLatencyDataLinkLost)
	mem.engine_failure = C.bool(m.EngineFailure)
	mem.mission_failure = C.bool(m.MissionFailure)
	mem.failure_detector_status = C.uint8_t(m.FailureDetectorStatus)
	mem.onboard_control_sensors_present = C.uint32_t(m.OnboardControlSensorsPresent)
	mem.onboard_control_sensors_enabled = C.uint32_t(m.OnboardControlSensorsEnabled)
	mem.onboard_control_sensors_health = C.uint32_t(m.OnboardControlSensorsHealth)
	mem.latest_arming_reason = C.uint8_t(m.LatestArmingReason)
	mem.latest_disarming_reason = C.uint8_t(m.LatestDisarmingReason)
	mem.armed_time = C.uint64_t(m.ArmedTime)
	mem.takeoff_time = C.uint64_t(m.TakeoffTime)
}

func (t _VehicleStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleStatus)
	mem := (*C.px4_msgs__msg__VehicleStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.NavState = uint8(mem.nav_state)
	m.NavStateTimestamp = uint64(mem.nav_state_timestamp)
	m.ArmingState = uint8(mem.arming_state)
	m.HilState = uint8(mem.hil_state)
	m.Failsafe = bool(mem.failsafe)
	m.FailsafeTimestamp = uint64(mem.failsafe_timestamp)
	m.SystemType = uint8(mem.system_type)
	m.SystemId = uint8(mem.system_id)
	m.ComponentId = uint8(mem.component_id)
	m.VehicleType = uint8(mem.vehicle_type)
	m.IsVtol = bool(mem.is_vtol)
	m.IsVtolTailsitter = bool(mem.is_vtol_tailsitter)
	m.VtolFwPermanentStab = bool(mem.vtol_fw_permanent_stab)
	m.InTransitionMode = bool(mem.in_transition_mode)
	m.InTransitionToFw = bool(mem.in_transition_to_fw)
	m.RcSignalLost = bool(mem.rc_signal_lost)
	m.RcInputMode = uint8(mem.rc_input_mode)
	m.DataLinkLost = bool(mem.data_link_lost)
	m.DataLinkLostCounter = uint8(mem.data_link_lost_counter)
	m.HighLatencyDataLinkLost = bool(mem.high_latency_data_link_lost)
	m.EngineFailure = bool(mem.engine_failure)
	m.MissionFailure = bool(mem.mission_failure)
	m.FailureDetectorStatus = uint8(mem.failure_detector_status)
	m.OnboardControlSensorsPresent = uint32(mem.onboard_control_sensors_present)
	m.OnboardControlSensorsEnabled = uint32(mem.onboard_control_sensors_enabled)
	m.OnboardControlSensorsHealth = uint32(mem.onboard_control_sensors_health)
	m.LatestArmingReason = uint8(mem.latest_arming_reason)
	m.LatestDisarmingReason = uint8(mem.latest_disarming_reason)
	m.ArmedTime = uint64(mem.armed_time)
	m.TakeoffTime = uint64(mem.takeoff_time)
}

func (t _VehicleStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleStatus())
}

type CVehicleStatus = C.px4_msgs__msg__VehicleStatus
type CVehicleStatus__Sequence = C.px4_msgs__msg__VehicleStatus__Sequence

func VehicleStatus__Sequence_to_Go(goSlice *[]VehicleStatus, cSlice CVehicleStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleStatus * uintptr(i)),
		))
		VehicleStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleStatus__Sequence_to_C(cSlice *CVehicleStatus__Sequence, goSlice []VehicleStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleStatus * uintptr(i)),
		))
		VehicleStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleStatus__Array_to_Go(goSlice []VehicleStatus, cSlice []CVehicleStatus) {
	for i := 0; i < len(cSlice); i++ {
		VehicleStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleStatus__Array_to_C(cSlice []CVehicleStatus, goSlice []VehicleStatus) {
	for i := 0; i < len(goSlice); i++ {
		VehicleStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
