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

#include <px4_msgs/msg/vehicle_command.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleCommand", VehicleCommandTypeSupport)
}
const (
	VehicleCommand_VEHICLE_CMD_CUSTOM_0 uint16 = 0// test command
	VehicleCommand_VEHICLE_CMD_CUSTOM_1 uint16 = 1// test command
	VehicleCommand_VEHICLE_CMD_CUSTOM_2 uint16 = 2// test command
	VehicleCommand_VEHICLE_CMD_NAV_WAYPOINT uint16 = 16// Navigate to MISSION. |Hold time in decimal seconds. (ignored by fixed wing, time to stay at MISSION for rotary wing)| Acceptance radius in meters (if the sphere with this radius is hit, the MISSION counts as reached)| 0 to pass through the WP, if > 0 radius in meters to pass by WP. Positive value for clockwise orbit, negative value for counter-clockwise orbit. Allows trajectory control.| Desired yaw angle at MISSION (rotary wing)| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_LOITER_UNLIM uint16 = 17// Loiter around this MISSION an unlimited amount of time |Empty| Empty| Radius around MISSION, in meters. If positive loiter clockwise, else counter-clockwise| Desired yaw angle.| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_LOITER_TURNS uint16 = 18// Loiter around this MISSION for X turns |Turns| Empty| Radius around MISSION, in meters. If positive loiter clockwise, else counter-clockwise| Desired yaw angle.| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_LOITER_TIME uint16 = 19// Loiter around this MISSION for X seconds |Seconds (decimal)| Empty| Radius around MISSION, in meters. If positive loiter clockwise, else counter-clockwise| Desired yaw angle.| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_RETURN_TO_LAUNCH uint16 = 20// Return to launch location |Empty| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_NAV_LAND uint16 = 21// Land at location |Empty| Empty| Empty| Desired yaw angle.| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_TAKEOFF uint16 = 22// Takeoff from ground / hand |Minimum pitch (if airspeed sensor present), desired pitch without sensor| Empty| Empty| Yaw angle (if magnetometer present), ignored without magnetometer| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_PRECLAND uint16 = 23// Attempt a precision landing
	VehicleCommand_VEHICLE_CMD_DO_ORBIT uint16 = 34// Start orbiting on the circumference of a circle defined by the parameters. |Radius [m] |Velocity [m/s] |Yaw behaviour |Empty |Latitude/X |Longitude/Y |Altitude/Z |
	VehicleCommand_VEHICLE_CMD_NAV_ROI uint16 = 80// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. |Region of interest mode. (see MAV_ROI enum)| MISSION index/ target ID. (see MAV_ROI enum)| ROI index (allows a vehicle to manage multiple ROI's)| Empty| x the location of the fixed ROI (see MAV_FRAME)| y| z|
	VehicleCommand_VEHICLE_CMD_NAV_PATHPLANNING uint16 = 81// Control autonomous path planning on the MAV. |0: Disable local obstacle avoidance / local path planning (without resetting map), 1: Enable local path planning, 2: Enable and reset local path planning| 0: Disable full path planning (without resetting map), 1: Enable, 2: Enable and reset map/occupancy grid, 3: Enable and reset planned route, but not occupancy grid| Empty| Yaw angle at goal, in compass degrees, [0..360]| Latitude/X of goal| Longitude/Y of goal| Altitude/Z of goal|
	VehicleCommand_VEHICLE_CMD_NAV_VTOL_TAKEOFF uint16 = 84// Takeoff from ground / hand and transition to fixed wing |Minimum pitch (if airspeed sensor present), desired pitch without sensor| Empty| Empty| Yaw angle (if magnetometer present), ignored without magnetometer| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_VTOL_LAND uint16 = 85// Transition to MC and land at location |Empty| Empty| Empty| Desired yaw angle.| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_NAV_GUIDED_LIMITS uint16 = 90// set limits for external control |timeout - maximum time (in seconds) that external controller will be allowed to control vehicle. 0 means no timeout| absolute altitude min (in meters, AMSL) - if vehicle moves below this alt, the command will be aborted and the mission will continue.  0 means no lower altitude limit| absolute altitude max (in meters)- if vehicle moves above this alt, the command will be aborted and the mission will continue.  0 means no upper altitude limit| horizontal move limit (in meters, AMSL) - if vehicle moves more than this distance from it's location at the moment the command was executed, the command will be aborted and the mission will continue. 0 means no horizontal altitude limit| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_NAV_GUIDED_MASTER uint16 = 91// set id of master controller |System ID| Component ID| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_NAV_DELAY uint16 = 93// Delay the next navigation command a number of seconds or until a specified time |Delay in seconds (decimal, -1 to enable time-of-day fields)| hour (24h format, UTC, -1 to ignore)| minute (24h format, UTC, -1 to ignore)| second (24h format, UTC)| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_NAV_LAST uint16 = 95// NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration |Empty| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_CONDITION_DELAY uint16 = 112// Delay mission state machine. |Delay in seconds (decimal)| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_CONDITION_CHANGE_ALT uint16 = 113// Ascend/descend at rate.  Delay mission state machine until desired altitude reached. |Descent / Ascend rate (m/s)| Empty| Empty| Empty| Empty| Empty| Finish Altitude|
	VehicleCommand_VEHICLE_CMD_CONDITION_DISTANCE uint16 = 114// Delay mission state machine until within desired distance of next NAV point. |Distance (meters)| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_CONDITION_YAW uint16 = 115// Reach a certain target angle. |target angle: [0-360], 0 is north| speed during yaw change:[deg per second]| direction: negative: counter clockwise, positive: clockwise [-1,1]| relative offset or absolute angle: [ 1,0]| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_CONDITION_LAST uint16 = 159// NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration |Empty| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_CONDITION_GATE uint16 = 4501// Wait until passing a threshold |2D coord mode: 0: Orthogonal to planned route | Altitude mode: 0: Ignore altitude| Empty| Empty| Lat| Lon| Alt|
	VehicleCommand_VEHICLE_CMD_DO_SET_MODE uint16 = 176// Set system mode. |Mode, as defined by ENUM MAV_MODE| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_JUMP uint16 = 177// Jump to the desired command in the mission list.  Repeat this action only the specified number of times |Sequence number| Repeat count| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_CHANGE_SPEED uint16 = 178// Change speed and/or throttle set points. |Speed type (0=Airspeed, 1=Ground Speed)| Speed  (m/s, -1 indicates no change)| Throttle  ( Percent, -1 indicates no change)| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_SET_HOME uint16 = 179// Changes the home location either to the current location or a specified location. |Use current (1=use current location, 0=use specified location)| Empty| Empty| Empty| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_DO_SET_PARAMETER uint16 = 180// Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter. |Parameter number| Parameter value| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_SET_RELAY uint16 = 181// Set a relay to a condition. |Relay number| Setting (1=on, 0=off, others possible depending on system hardware)| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_REPEAT_RELAY uint16 = 182// Cycle a relay on and off for a desired number of cyles with a desired period. |Relay number| Cycle count| Cycle time (seconds, decimal)| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_SET_SERVO uint16 = 183// Set a servo to a desired PWM value. |Servo number| PWM (microseconds, 1000 to 2000 typical)| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_REPEAT_SERVO uint16 = 184// Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period. |Servo number| PWM (microseconds, 1000 to 2000 typical)| Cycle count| Cycle time (seconds)| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_FLIGHTTERMINATION uint16 = 185// Terminate flight immediately |Flight termination activated if > 0.5| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_LAND_START uint16 = 189// Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0/0 if not needed. If specified then it will be used to help find the closest landing sequence. |Empty| Empty| Empty| Empty| Latitude| Longitude| Empty|  */
	VehicleCommand_VEHICLE_CMD_DO_GO_AROUND uint16 = 191// Mission command to safely abort an autonmous landing. |Altitude (meters)| Empty| Empty| Empty| Empty| Empty| Empty|  */
	VehicleCommand_VEHICLE_CMD_DO_REPOSITION uint16 = 192
	VehicleCommand_VEHICLE_CMD_DO_PAUSE_CONTINUE uint16 = 193
	VehicleCommand_VEHICLE_CMD_DO_SET_ROI_LOCATION uint16 = 195// Sets the region of interest (ROI) to a location. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. |Empty| Empty| Empty| Empty| Latitude| Longitude| Altitude|
	VehicleCommand_VEHICLE_CMD_DO_SET_ROI_WPNEXT_OFFSET uint16 = 196// Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. |Empty| Empty| Empty| Empty| pitch offset from next waypoint| roll offset from next waypoint| yaw offset from next waypoint|
	VehicleCommand_VEHICLE_CMD_DO_SET_ROI_NONE uint16 = 197// Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. |Empty| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_CONTROL_VIDEO uint16 = 200// Control onboard camera system. |Camera ID (-1 for all)| Transmission: 0: disabled, 1: enabled compressed, 2: enabled raw| Transmission mode: 0: video stream, >0: single images every n seconds (decimal)| Recording: 0: disabled, 1: enabled compressed, 2: enabled raw| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_SET_ROI uint16 = 201// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. |Region of interest mode. (see MAV_ROI enum)| MISSION index/ target ID. (see MAV_ROI enum)| ROI index (allows a vehicle to manage multiple ROI's)| Empty| x the location of the fixed ROI (see MAV_FRAME)| y| z|
	VehicleCommand_VEHICLE_CMD_DO_DIGICAM_CONTROL uint16 = 203
	VehicleCommand_VEHICLE_CMD_DO_MOUNT_CONFIGURE uint16 = 204// Mission command to configure a camera or antenna mount |Mount operation mode (see MAV_MOUNT_MODE enum)| stabilize roll? (1 = yes, 0 = no)| stabilize pitch? (1 = yes, 0 = no)| stabilize yaw? (1 = yes, 0 = no)| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_MOUNT_CONTROL uint16 = 205// Mission command to control a camera or antenna mount |pitch or lat in degrees, depending on mount mode.| roll or lon in degrees depending on mount mode| yaw or alt (in meters) depending on mount mode| reserved| reserved| reserved| MAV_MOUNT_MODE enum value|
	VehicleCommand_VEHICLE_CMD_DO_SET_CAM_TRIGG_DIST uint16 = 206// Mission command to set TRIG_DIST for this flight |Camera trigger distance (meters)| Shutter integration time (ms)| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_FENCE_ENABLE uint16 = 207// Mission command to enable the geofence |enable? (0=disable, 1=enable)| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_PARACHUTE uint16 = 208// Mission command to trigger a parachute |action (0=disable, 1=enable, 2=release, for some systems see PARACHUTE_ACTION enum, not in general message set.)| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_MOTOR_TEST uint16 = 209// motor test command |Instance (1, ...)| throttle type| throttle| timeout [s]| Motor count | Test order| Empty|
	VehicleCommand_VEHICLE_CMD_DO_INVERTED_FLIGHT uint16 = 210// Change to/from inverted flight |inverted (0=normal, 1=inverted)| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_SET_CAM_TRIGG_INTERVAL uint16 = 214// Mission command to set TRIG_INTERVAL for this flight |Camera trigger distance (meters)| Shutter integration time (ms)| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_MOUNT_CONTROL_QUAT uint16 = 220// Mission command to control a camera or antenna mount, using a quaternion as reference. |q1 - quaternion param #1, w (1 in null-rotation)| q2 - quaternion param #2, x (0 in null-rotation)| q3 - quaternion param #3, y (0 in null-rotation)| q4 - quaternion param #4, z (0 in null-rotation)| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_GUIDED_MASTER uint16 = 221// set id of master controller |System ID| Component ID| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_GUIDED_LIMITS uint16 = 222// set limits for external control |timeout - maximum time (in seconds) that external controller will be allowed to control vehicle. 0 means no timeout| absolute altitude min (in meters, AMSL) - if vehicle moves below this alt, the command will be aborted and the mission will continue.  0 means no lower altitude limit| absolute altitude max (in meters)- if vehicle moves above this alt, the command will be aborted and the mission will continue.  0 means no upper altitude limit| horizontal move limit (in meters, AMSL) - if vehicle moves more than this distance from it's location at the moment the command was executed, the command will be aborted and the mission will continue. 0 means no horizontal altitude limit| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_DO_LAST uint16 = 240// NOP - This command is only used to mark the upper limit of the DO commands in the enumeration |Empty| Empty| Empty| Empty| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_PREFLIGHT_CALIBRATION uint16 = 241// Trigger calibration. This command will be only accepted if in pre-flight mode. See mavlink spec MAV_CMD_PREFLIGHT_CALIBRATION
	VehicleCommand_PREFLIGHT_CALIBRATION_TEMPERATURE_CALIBRATION uint16 = 3// param value for VEHICLE_CMD_PREFLIGHT_CALIBRATION to start temperature calibration
	VehicleCommand_VEHICLE_CMD_PREFLIGHT_SET_SENSOR_OFFSETS uint16 = 242// Set sensor offsets. This command will be only accepted if in pre-flight mode. |Sensor to adjust the offsets for: 0: gyros, 1: accelerometer, 2: magnetometer, 3: barometer, 4: optical flow| X axis offset (or generic dimension 1), in the sensor's raw units| Y axis offset (or generic dimension 2), in the sensor's raw units| Z axis offset (or generic dimension 3), in the sensor's raw units| Generic dimension 4, in the sensor's raw units| Generic dimension 5, in the sensor's raw units| Generic dimension 6, in the sensor's raw units|
	VehicleCommand_VEHICLE_CMD_PREFLIGHT_UAVCAN uint16 = 243// UAVCAN configuration. If param 1 == 1 actuator mapping and direction assignment should be started
	VehicleCommand_VEHICLE_CMD_PREFLIGHT_STORAGE uint16 = 245// Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode. |Parameter storage: 0: READ FROM FLASH/EEPROM, 1: WRITE CURRENT TO FLASH/EEPROM| Mission storage: 0: READ FROM FLASH/EEPROM, 1: WRITE CURRENT TO FLASH/EEPROM| Reserved| Reserved| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_PREFLIGHT_REBOOT_SHUTDOWN uint16 = 246// Request the reboot or shutdown of system components. |0: Do nothing for autopilot, 1: Reboot autopilot, 2: Shutdown autopilot.| 0: Do nothing for onboard computer, 1: Reboot onboard computer, 2: Shutdown onboard computer.| Reserved| Reserved| Empty| Empty| Empty|
	VehicleCommand_VEHICLE_CMD_OBLIQUE_SURVEY uint16 = 260// Mission command to set a Camera Auto Mount Pivoting Oblique Survey for this flight|Camera trigger distance (meters)| Shutter integration time (ms)| Camera minimum trigger interval| Number of positions| Roll| Pitch| Empty|
	VehicleCommand_VEHICLE_CMD_GIMBAL_DEVICE_INFORMATION uint16 = 283// Command to ask information about a low level gimbal
	VehicleCommand_VEHICLE_CMD_MISSION_START uint16 = 300// start running a mission |first_item: the first mission item to run| last_item:  the last mission item to run (after this item is run, the mission ends)|
	VehicleCommand_VEHICLE_CMD_COMPONENT_ARM_DISARM uint16 = 400// Arms / Disarms a component |1 to arm, 0 to disarm|
	VehicleCommand_VEHICLE_CMD_INJECT_FAILURE uint16 = 420// Inject artificial failure for testing purposes
	VehicleCommand_VEHICLE_CMD_START_RX_PAIR uint16 = 500// Starts receiver pairing |0:Spektrum| 0:Spektrum DSM2, 1:Spektrum DSMX|
	VehicleCommand_VEHICLE_CMD_REQUEST_MESSAGE uint16 = 512// Request to send a single instance of the specified message
	VehicleCommand_VEHICLE_CMD_SET_CAMERA_MODE uint16 = 530// Set camera capture mode (photo, video, etc.)
	VehicleCommand_VEHICLE_CMD_SET_CAMERA_ZOOM uint16 = 531// Set camera zoom
	VehicleCommand_VEHICLE_CMD_SET_CAMERA_FOCUS uint16 = 532
	VehicleCommand_VEHICLE_CMD_DO_GIMBAL_MANAGER_PITCHYAW uint16 = 1000// Setpoint to be sent to a gimbal manager to set a gimbal pitch and yaw
	VehicleCommand_VEHICLE_CMD_DO_GIMBAL_MANAGER_CONFIGURE uint16 = 1001// Gimbal configuration to set which sysid/compid is in primary and secondary control
	VehicleCommand_VEHICLE_CMD_DO_TRIGGER_CONTROL uint16 = 2003// Enable or disable on-board camera triggering system
	VehicleCommand_VEHICLE_CMD_LOGGING_START uint16 = 2510// start streaming ULog data
	VehicleCommand_VEHICLE_CMD_LOGGING_STOP uint16 = 2511// stop streaming ULog data
	VehicleCommand_VEHICLE_CMD_CONTROL_HIGH_LATENCY uint16 = 2600// control starting/stopping transmitting data over the high latency link
	VehicleCommand_VEHICLE_CMD_DO_VTOL_TRANSITION uint16 = 3000// Command VTOL transition
	VehicleCommand_VEHICLE_CMD_ARM_AUTHORIZATION_REQUEST uint16 = 3001// Request arm authorization
	VehicleCommand_VEHICLE_CMD_PAYLOAD_PREPARE_DEPLOY uint16 = 30001// Prepare a payload deployment in the flight plan
	VehicleCommand_VEHICLE_CMD_PAYLOAD_CONTROL_DEPLOY uint16 = 30002// Control a pre-programmed payload deployment
	VehicleCommand_VEHICLE_CMD_FIXED_MAG_CAL_YAW uint16 = 42006// Magnetometer calibration based on provided known yaw. This allows for fast calibration using WMM field tables in the vehicle, given only the known yaw of the vehicle. If Latitude and longitude are both zero then use the current vehicle location.
	VehicleCommand_VEHICLE_CMD_PX4_INTERNAL_START uint32 = 65537// start of PX4 internal only vehicle commands (> UINT16_MAX). PX4 vehicle commands (beyond 16 bit mavlink commands)
	VehicleCommand_VEHICLE_CMD_SET_GPS_GLOBAL_ORIGIN uint32 = 100000// Sets the GPS co-ordinates of the vehicle local origin (0,0,0) position. |Empty|Empty|Empty|Empty|Latitude|Longitude|Altitude|. PX4 vehicle commands (beyond 16 bit mavlink commands)
	VehicleCommand_VEHICLE_CMD_RESULT_ACCEPTED uint8 = 0// Command ACCEPTED and EXECUTED |
	VehicleCommand_VEHICLE_CMD_RESULT_TEMPORARILY_REJECTED uint8 = 1// Command TEMPORARY REJECTED/DENIED |
	VehicleCommand_VEHICLE_CMD_RESULT_DENIED uint8 = 2// Command PERMANENTLY DENIED |
	VehicleCommand_VEHICLE_CMD_RESULT_UNSUPPORTED uint8 = 3// Command UNKNOWN/UNSUPPORTED |
	VehicleCommand_VEHICLE_CMD_RESULT_FAILED uint8 = 4// Command executed, but failed |
	VehicleCommand_VEHICLE_CMD_RESULT_IN_PROGRESS uint8 = 5// Command being executed |
	VehicleCommand_VEHICLE_CMD_RESULT_ENUM_END uint8 = 6
	VehicleCommand_VEHICLE_MOUNT_MODE_RETRACT uint8 = 0// Load and keep safe position (Roll,Pitch,Yaw) from permanent memory and stop stabilization |
	VehicleCommand_VEHICLE_MOUNT_MODE_NEUTRAL uint8 = 1// Load and keep neutral position (Roll,Pitch,Yaw) from permanent memory. |
	VehicleCommand_VEHICLE_MOUNT_MODE_MAVLINK_TARGETING uint8 = 2// Load neutral position and start MAVLink Roll,Pitch,Yaw control with stabilization |
	VehicleCommand_VEHICLE_MOUNT_MODE_RC_TARGETING uint8 = 3// Load neutral position and start RC Roll,Pitch,Yaw control with stabilization |
	VehicleCommand_VEHICLE_MOUNT_MODE_GPS_POINT uint8 = 4// Load neutral position and start to point to Lat,Lon,Alt |
	VehicleCommand_VEHICLE_MOUNT_MODE_ENUM_END uint8 = 5
	VehicleCommand_VEHICLE_ROI_NONE uint8 = 0// No region of interest |
	VehicleCommand_VEHICLE_ROI_WPNEXT uint8 = 1// Point toward next MISSION |
	VehicleCommand_VEHICLE_ROI_WPINDEX uint8 = 2// Point toward given MISSION |
	VehicleCommand_VEHICLE_ROI_LOCATION uint8 = 3// Point toward fixed location |
	VehicleCommand_VEHICLE_ROI_TARGET uint8 = 4// Point toward target
	VehicleCommand_VEHICLE_ROI_ENUM_END uint8 = 5
	VehicleCommand_VEHICLE_CAMERA_ZOOM_TYPE_STEP uint8 = 0// Zoom one step increment
	VehicleCommand_VEHICLE_CAMERA_ZOOM_TYPE_CONTINUOUS uint8 = 1// Continuous zoom up/down until stopped
	VehicleCommand_VEHICLE_CAMERA_ZOOM_TYPE_RANGE uint8 = 2// Zoom value as proportion of full camera range
	VehicleCommand_VEHICLE_CAMERA_ZOOM_TYPE_FOCAL_LENGTH uint8 = 3// Zoom to a focal length
	VehicleCommand_FAILURE_UNIT_SENSOR_GYRO uint8 = 0
	VehicleCommand_FAILURE_UNIT_SENSOR_ACCEL uint8 = 1
	VehicleCommand_FAILURE_UNIT_SENSOR_MAG uint8 = 2
	VehicleCommand_FAILURE_UNIT_SENSOR_BARO uint8 = 3
	VehicleCommand_FAILURE_UNIT_SENSOR_GPS uint8 = 4
	VehicleCommand_FAILURE_UNIT_SENSOR_OPTICAL_FLOW uint8 = 5
	VehicleCommand_FAILURE_UNIT_SENSOR_VIO uint8 = 6
	VehicleCommand_FAILURE_UNIT_SENSOR_DISTANCE_SENSOR uint8 = 7
	VehicleCommand_FAILURE_UNIT_SENSOR_AIRSPEED uint8 = 8
	VehicleCommand_FAILURE_UNIT_SYSTEM_BATTERY uint8 = 100
	VehicleCommand_FAILURE_UNIT_SYSTEM_MOTOR uint8 = 101
	VehicleCommand_FAILURE_UNIT_SYSTEM_SERVO uint8 = 102
	VehicleCommand_FAILURE_UNIT_SYSTEM_AVOIDANCE uint8 = 103
	VehicleCommand_FAILURE_UNIT_SYSTEM_RC_SIGNAL uint8 = 104
	VehicleCommand_FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL uint8 = 105
	VehicleCommand_FAILURE_TYPE_OK uint8 = 0
	VehicleCommand_FAILURE_TYPE_OFF uint8 = 1
	VehicleCommand_FAILURE_TYPE_STUCK uint8 = 2
	VehicleCommand_FAILURE_TYPE_GARBAGE uint8 = 3
	VehicleCommand_FAILURE_TYPE_WRONG uint8 = 4
	VehicleCommand_FAILURE_TYPE_SLOW uint8 = 5
	VehicleCommand_FAILURE_TYPE_DELAYED uint8 = 6
	VehicleCommand_FAILURE_TYPE_INTERMITTENT uint8 = 7
	VehicleCommand_ORB_QUEUE_LENGTH uint8 = 8
)

// Do not create instances of this type directly. Always use NewVehicleCommand
// function instead.
type VehicleCommand struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Param1 float32 `yaml:"param1"`// Parameter 1, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Param2 float32 `yaml:"param2"`// Parameter 2, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Param3 float32 `yaml:"param3"`// Parameter 3, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Param4 float32 `yaml:"param4"`// Parameter 4, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Param5 float64 `yaml:"param5"`// Parameter 5, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Param6 float64 `yaml:"param6"`// Parameter 6, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Param7 float32 `yaml:"param7"`// Parameter 7, as defined by MAVLink uint16 VEHICLE_CMD enum.
	Command uint32 `yaml:"command"`// Command ID
	TargetSystem uint8 `yaml:"target_system"`// System which should execute the command
	TargetComponent uint8 `yaml:"target_component"`// Component which should execute the command, 0 for all components
	SourceSystem uint8 `yaml:"source_system"`// System sending the command
	SourceComponent uint8 `yaml:"source_component"`// Component sending the command
	Confirmation uint8 `yaml:"confirmation"`// 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)
	FromExternal bool `yaml:"from_external"`
}

// NewVehicleCommand creates a new VehicleCommand with default values.
func NewVehicleCommand() *VehicleCommand {
	self := VehicleCommand{}
	self.SetDefaults()
	return &self
}

func (t *VehicleCommand) Clone() *VehicleCommand {
	c := &VehicleCommand{}
	c.Timestamp = t.Timestamp
	c.Param1 = t.Param1
	c.Param2 = t.Param2
	c.Param3 = t.Param3
	c.Param4 = t.Param4
	c.Param5 = t.Param5
	c.Param6 = t.Param6
	c.Param7 = t.Param7
	c.Command = t.Command
	c.TargetSystem = t.TargetSystem
	c.TargetComponent = t.TargetComponent
	c.SourceSystem = t.SourceSystem
	c.SourceComponent = t.SourceComponent
	c.Confirmation = t.Confirmation
	c.FromExternal = t.FromExternal
	return c
}

func (t *VehicleCommand) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleCommand) SetDefaults() {
	t.Timestamp = 0
	t.Param1 = 0
	t.Param2 = 0
	t.Param3 = 0
	t.Param4 = 0
	t.Param5 = 0
	t.Param6 = 0
	t.Param7 = 0
	t.Command = 0
	t.TargetSystem = 0
	t.TargetComponent = 0
	t.SourceSystem = 0
	t.SourceComponent = 0
	t.Confirmation = 0
	t.FromExternal = false
}

// CloneVehicleCommandSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleCommandSlice(dst, src []VehicleCommand) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleCommandTypeSupport types.MessageTypeSupport = _VehicleCommandTypeSupport{}

type _VehicleCommandTypeSupport struct{}

func (t _VehicleCommandTypeSupport) New() types.Message {
	return NewVehicleCommand()
}

func (t _VehicleCommandTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleCommand
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleCommand__create())
}

func (t _VehicleCommandTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleCommand__destroy((*C.px4_msgs__msg__VehicleCommand)(pointer_to_free))
}

func (t _VehicleCommandTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleCommand)
	mem := (*C.px4_msgs__msg__VehicleCommand)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.param1 = C.float(m.Param1)
	mem.param2 = C.float(m.Param2)
	mem.param3 = C.float(m.Param3)
	mem.param4 = C.float(m.Param4)
	mem.param5 = C.double(m.Param5)
	mem.param6 = C.double(m.Param6)
	mem.param7 = C.float(m.Param7)
	mem.command = C.uint32_t(m.Command)
	mem.target_system = C.uint8_t(m.TargetSystem)
	mem.target_component = C.uint8_t(m.TargetComponent)
	mem.source_system = C.uint8_t(m.SourceSystem)
	mem.source_component = C.uint8_t(m.SourceComponent)
	mem.confirmation = C.uint8_t(m.Confirmation)
	mem.from_external = C.bool(m.FromExternal)
}

func (t _VehicleCommandTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleCommand)
	mem := (*C.px4_msgs__msg__VehicleCommand)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Param1 = float32(mem.param1)
	m.Param2 = float32(mem.param2)
	m.Param3 = float32(mem.param3)
	m.Param4 = float32(mem.param4)
	m.Param5 = float64(mem.param5)
	m.Param6 = float64(mem.param6)
	m.Param7 = float32(mem.param7)
	m.Command = uint32(mem.command)
	m.TargetSystem = uint8(mem.target_system)
	m.TargetComponent = uint8(mem.target_component)
	m.SourceSystem = uint8(mem.source_system)
	m.SourceComponent = uint8(mem.source_component)
	m.Confirmation = uint8(mem.confirmation)
	m.FromExternal = bool(mem.from_external)
}

func (t _VehicleCommandTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleCommand())
}

type CVehicleCommand = C.px4_msgs__msg__VehicleCommand
type CVehicleCommand__Sequence = C.px4_msgs__msg__VehicleCommand__Sequence

func VehicleCommand__Sequence_to_Go(goSlice *[]VehicleCommand, cSlice CVehicleCommand__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleCommand, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleCommand__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleCommand * uintptr(i)),
		))
		VehicleCommandTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleCommand__Sequence_to_C(cSlice *CVehicleCommand__Sequence, goSlice []VehicleCommand) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleCommand)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleCommand * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleCommand)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleCommand * uintptr(i)),
		))
		VehicleCommandTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleCommand__Array_to_Go(goSlice []VehicleCommand, cSlice []CVehicleCommand) {
	for i := 0; i < len(cSlice); i++ {
		VehicleCommandTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleCommand__Array_to_C(cSlice []CVehicleCommand, goSlice []VehicleCommand) {
	for i := 0; i < len(goSlice); i++ {
		VehicleCommandTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
