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

#include <px4_msgs/msg/internal_combustion_engine_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/InternalCombustionEngineStatus", InternalCombustionEngineStatusTypeSupport)
}
const (
	InternalCombustionEngineStatus_STATE_STOPPED uint8 = 0// The engine is not running. This is the default state.
	InternalCombustionEngineStatus_STATE_STARTING uint8 = 1// The engine is starting. This is a transient state.
	InternalCombustionEngineStatus_STATE_RUNNING uint8 = 2// The engine is running normally.
	InternalCombustionEngineStatus_STATE_FAULT uint8 = 3// The engine can no longer function.
	InternalCombustionEngineStatus_FLAG_GENERAL_ERROR uint32 = 1// General error.
	InternalCombustionEngineStatus_FLAG_CRANKSHAFT_SENSOR_ERROR_SUPPORTED uint32 = 2// Error of the crankshaft sensor. This flag is optional.
	InternalCombustionEngineStatus_FLAG_CRANKSHAFT_SENSOR_ERROR uint32 = 4
	InternalCombustionEngineStatus_FLAG_TEMPERATURE_SUPPORTED uint32 = 8// Temperature levels. These flags are optional
	InternalCombustionEngineStatus_FLAG_TEMPERATURE_BELOW_NOMINAL uint32 = 16// Under-temperature warning
	InternalCombustionEngineStatus_FLAG_TEMPERATURE_ABOVE_NOMINAL uint32 = 32// Over-temperature warning
	InternalCombustionEngineStatus_FLAG_TEMPERATURE_OVERHEATING uint32 = 64// Critical overheating
	InternalCombustionEngineStatus_FLAG_TEMPERATURE_EGT_ABOVE_NOMINAL uint32 = 128// Exhaust gas over-temperature warning
	InternalCombustionEngineStatus_FLAG_FUEL_PRESSURE_SUPPORTED uint32 = 256// Fuel pressure. These flags are optional
	InternalCombustionEngineStatus_FLAG_FUEL_PRESSURE_BELOW_NOMINAL uint32 = 512// Under-pressure warning
	InternalCombustionEngineStatus_FLAG_FUEL_PRESSURE_ABOVE_NOMINAL uint32 = 1024// Over-pressure warning
	InternalCombustionEngineStatus_FLAG_DETONATION_SUPPORTED uint32 = 2048// Detonation warning. This flag is optional.
	InternalCombustionEngineStatus_FLAG_DETONATION_OBSERVED uint32 = 4096// Detonation condition observed warning
	InternalCombustionEngineStatus_FLAG_MISFIRE_SUPPORTED uint32 = 8192// Misfire warning. This flag is optional.
	InternalCombustionEngineStatus_FLAG_MISFIRE_OBSERVED uint32 = 16384// Misfire condition observed warning
	InternalCombustionEngineStatus_FLAG_OIL_PRESSURE_SUPPORTED uint32 = 32768// Oil pressure. These flags are optional
	InternalCombustionEngineStatus_FLAG_OIL_PRESSURE_BELOW_NOMINAL uint32 = 65536// Under-pressure warning
	InternalCombustionEngineStatus_FLAG_OIL_PRESSURE_ABOVE_NOMINAL uint32 = 131072// Over-pressure warning
	InternalCombustionEngineStatus_FLAG_DEBRIS_SUPPORTED uint32 = 262144// Debris warning. This flag is optional
	InternalCombustionEngineStatus_FLAG_DEBRIS_DETECTED uint32 = 524288// Detection of debris warning
	InternalCombustionEngineStatus_SPARK_PLUG_SINGLE uint8 = 0
	InternalCombustionEngineStatus_SPARK_PLUG_FIRST_ACTIVE uint8 = 1
	InternalCombustionEngineStatus_SPARK_PLUG_SECOND_ACTIVE uint8 = 2
	InternalCombustionEngineStatus_SPARK_PLUG_BOTH_ACTIVE uint8 = 3
)

// Do not create instances of this type directly. Always use NewInternalCombustionEngineStatus
// function instead.
type InternalCombustionEngineStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	State uint8 `yaml:"state"`
	Flags uint32 `yaml:"flags"`
	EngineLoadPercent uint8 `yaml:"engine_load_percent"`// Engine load estimate, percent, [0, 127]
	EngineSpeedRpm uint32 `yaml:"engine_speed_rpm"`// Engine speed, revolutions per minute
	SparkDwellTimeMs float32 `yaml:"spark_dwell_time_ms"`// Spark dwell time, millisecond
	AtmosphericPressureKpa float32 `yaml:"atmospheric_pressure_kpa"`// Atmospheric (barometric) pressure, kilopascal
	IntakeManifoldPressureKpa float32 `yaml:"intake_manifold_pressure_kpa"`// Engine intake manifold pressure, kilopascal
	IntakeManifoldTemperature float32 `yaml:"intake_manifold_temperature"`// Engine intake manifold temperature, kelvin
	CoolantTemperature float32 `yaml:"coolant_temperature"`// Engine coolant temperature, kelvin
	OilPressure float32 `yaml:"oil_pressure"`// Oil pressure, kilopascal
	OilTemperature float32 `yaml:"oil_temperature"`// Oil temperature, kelvin
	FuelPressure float32 `yaml:"fuel_pressure"`// Fuel pressure, kilopascal
	FuelConsumptionRateCm3pm float32 `yaml:"fuel_consumption_rate_cm3pm"`// Instant fuel consumption estimate, (centimeter^3)/minute
	EstimatedConsumedFuelVolumeCm3 float32 `yaml:"estimated_consumed_fuel_volume_cm3"`// Estimate of the consumed fuel since the start of the engine, centimeter^3
	ThrottlePositionPercent uint8 `yaml:"throttle_position_percent"`// Throttle position, percent
	EcuIndex uint8 `yaml:"ecu_index"`// The index of the publishing ECU
	SparkPlugUsage uint8 `yaml:"spark_plug_usage"`// Spark plug activity report.
	IgnitionTimingDeg float32 `yaml:"ignition_timing_deg"`// Cylinder ignition timing, angular degrees of the crankshaft
	InjectionTimeMs float32 `yaml:"injection_time_ms"`// Fuel injection time, millisecond
	CylinderHeadTemperature float32 `yaml:"cylinder_head_temperature"`// Cylinder head temperature (CHT), kelvin
	ExhaustGasTemperature float32 `yaml:"exhaust_gas_temperature"`// Exhaust gas temperature (EGT), kelvin
	LambdaCoefficient float32 `yaml:"lambda_coefficient"`// Estimated lambda coefficient, dimensionless ratio
}

// NewInternalCombustionEngineStatus creates a new InternalCombustionEngineStatus with default values.
func NewInternalCombustionEngineStatus() *InternalCombustionEngineStatus {
	self := InternalCombustionEngineStatus{}
	self.SetDefaults()
	return &self
}

func (t *InternalCombustionEngineStatus) Clone() *InternalCombustionEngineStatus {
	c := &InternalCombustionEngineStatus{}
	c.Timestamp = t.Timestamp
	c.State = t.State
	c.Flags = t.Flags
	c.EngineLoadPercent = t.EngineLoadPercent
	c.EngineSpeedRpm = t.EngineSpeedRpm
	c.SparkDwellTimeMs = t.SparkDwellTimeMs
	c.AtmosphericPressureKpa = t.AtmosphericPressureKpa
	c.IntakeManifoldPressureKpa = t.IntakeManifoldPressureKpa
	c.IntakeManifoldTemperature = t.IntakeManifoldTemperature
	c.CoolantTemperature = t.CoolantTemperature
	c.OilPressure = t.OilPressure
	c.OilTemperature = t.OilTemperature
	c.FuelPressure = t.FuelPressure
	c.FuelConsumptionRateCm3pm = t.FuelConsumptionRateCm3pm
	c.EstimatedConsumedFuelVolumeCm3 = t.EstimatedConsumedFuelVolumeCm3
	c.ThrottlePositionPercent = t.ThrottlePositionPercent
	c.EcuIndex = t.EcuIndex
	c.SparkPlugUsage = t.SparkPlugUsage
	c.IgnitionTimingDeg = t.IgnitionTimingDeg
	c.InjectionTimeMs = t.InjectionTimeMs
	c.CylinderHeadTemperature = t.CylinderHeadTemperature
	c.ExhaustGasTemperature = t.ExhaustGasTemperature
	c.LambdaCoefficient = t.LambdaCoefficient
	return c
}

func (t *InternalCombustionEngineStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InternalCombustionEngineStatus) SetDefaults() {
	t.Timestamp = 0
	t.State = 0
	t.Flags = 0
	t.EngineLoadPercent = 0
	t.EngineSpeedRpm = 0
	t.SparkDwellTimeMs = 0
	t.AtmosphericPressureKpa = 0
	t.IntakeManifoldPressureKpa = 0
	t.IntakeManifoldTemperature = 0
	t.CoolantTemperature = 0
	t.OilPressure = 0
	t.OilTemperature = 0
	t.FuelPressure = 0
	t.FuelConsumptionRateCm3pm = 0
	t.EstimatedConsumedFuelVolumeCm3 = 0
	t.ThrottlePositionPercent = 0
	t.EcuIndex = 0
	t.SparkPlugUsage = 0
	t.IgnitionTimingDeg = 0
	t.InjectionTimeMs = 0
	t.CylinderHeadTemperature = 0
	t.ExhaustGasTemperature = 0
	t.LambdaCoefficient = 0
}

// CloneInternalCombustionEngineStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInternalCombustionEngineStatusSlice(dst, src []InternalCombustionEngineStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InternalCombustionEngineStatusTypeSupport types.MessageTypeSupport = _InternalCombustionEngineStatusTypeSupport{}

type _InternalCombustionEngineStatusTypeSupport struct{}

func (t _InternalCombustionEngineStatusTypeSupport) New() types.Message {
	return NewInternalCombustionEngineStatus()
}

func (t _InternalCombustionEngineStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__InternalCombustionEngineStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__InternalCombustionEngineStatus__create())
}

func (t _InternalCombustionEngineStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__InternalCombustionEngineStatus__destroy((*C.px4_msgs__msg__InternalCombustionEngineStatus)(pointer_to_free))
}

func (t _InternalCombustionEngineStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InternalCombustionEngineStatus)
	mem := (*C.px4_msgs__msg__InternalCombustionEngineStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.state = C.uint8_t(m.State)
	mem.flags = C.uint32_t(m.Flags)
	mem.engine_load_percent = C.uint8_t(m.EngineLoadPercent)
	mem.engine_speed_rpm = C.uint32_t(m.EngineSpeedRpm)
	mem.spark_dwell_time_ms = C.float(m.SparkDwellTimeMs)
	mem.atmospheric_pressure_kpa = C.float(m.AtmosphericPressureKpa)
	mem.intake_manifold_pressure_kpa = C.float(m.IntakeManifoldPressureKpa)
	mem.intake_manifold_temperature = C.float(m.IntakeManifoldTemperature)
	mem.coolant_temperature = C.float(m.CoolantTemperature)
	mem.oil_pressure = C.float(m.OilPressure)
	mem.oil_temperature = C.float(m.OilTemperature)
	mem.fuel_pressure = C.float(m.FuelPressure)
	mem.fuel_consumption_rate_cm3pm = C.float(m.FuelConsumptionRateCm3pm)
	mem.estimated_consumed_fuel_volume_cm3 = C.float(m.EstimatedConsumedFuelVolumeCm3)
	mem.throttle_position_percent = C.uint8_t(m.ThrottlePositionPercent)
	mem.ecu_index = C.uint8_t(m.EcuIndex)
	mem.spark_plug_usage = C.uint8_t(m.SparkPlugUsage)
	mem.ignition_timing_deg = C.float(m.IgnitionTimingDeg)
	mem.injection_time_ms = C.float(m.InjectionTimeMs)
	mem.cylinder_head_temperature = C.float(m.CylinderHeadTemperature)
	mem.exhaust_gas_temperature = C.float(m.ExhaustGasTemperature)
	mem.lambda_coefficient = C.float(m.LambdaCoefficient)
}

func (t _InternalCombustionEngineStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InternalCombustionEngineStatus)
	mem := (*C.px4_msgs__msg__InternalCombustionEngineStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.State = uint8(mem.state)
	m.Flags = uint32(mem.flags)
	m.EngineLoadPercent = uint8(mem.engine_load_percent)
	m.EngineSpeedRpm = uint32(mem.engine_speed_rpm)
	m.SparkDwellTimeMs = float32(mem.spark_dwell_time_ms)
	m.AtmosphericPressureKpa = float32(mem.atmospheric_pressure_kpa)
	m.IntakeManifoldPressureKpa = float32(mem.intake_manifold_pressure_kpa)
	m.IntakeManifoldTemperature = float32(mem.intake_manifold_temperature)
	m.CoolantTemperature = float32(mem.coolant_temperature)
	m.OilPressure = float32(mem.oil_pressure)
	m.OilTemperature = float32(mem.oil_temperature)
	m.FuelPressure = float32(mem.fuel_pressure)
	m.FuelConsumptionRateCm3pm = float32(mem.fuel_consumption_rate_cm3pm)
	m.EstimatedConsumedFuelVolumeCm3 = float32(mem.estimated_consumed_fuel_volume_cm3)
	m.ThrottlePositionPercent = uint8(mem.throttle_position_percent)
	m.EcuIndex = uint8(mem.ecu_index)
	m.SparkPlugUsage = uint8(mem.spark_plug_usage)
	m.IgnitionTimingDeg = float32(mem.ignition_timing_deg)
	m.InjectionTimeMs = float32(mem.injection_time_ms)
	m.CylinderHeadTemperature = float32(mem.cylinder_head_temperature)
	m.ExhaustGasTemperature = float32(mem.exhaust_gas_temperature)
	m.LambdaCoefficient = float32(mem.lambda_coefficient)
}

func (t _InternalCombustionEngineStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__InternalCombustionEngineStatus())
}

type CInternalCombustionEngineStatus = C.px4_msgs__msg__InternalCombustionEngineStatus
type CInternalCombustionEngineStatus__Sequence = C.px4_msgs__msg__InternalCombustionEngineStatus__Sequence

func InternalCombustionEngineStatus__Sequence_to_Go(goSlice *[]InternalCombustionEngineStatus, cSlice CInternalCombustionEngineStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InternalCombustionEngineStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__InternalCombustionEngineStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__InternalCombustionEngineStatus * uintptr(i)),
		))
		InternalCombustionEngineStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InternalCombustionEngineStatus__Sequence_to_C(cSlice *CInternalCombustionEngineStatus__Sequence, goSlice []InternalCombustionEngineStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__InternalCombustionEngineStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__InternalCombustionEngineStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__InternalCombustionEngineStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__InternalCombustionEngineStatus * uintptr(i)),
		))
		InternalCombustionEngineStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InternalCombustionEngineStatus__Array_to_Go(goSlice []InternalCombustionEngineStatus, cSlice []CInternalCombustionEngineStatus) {
	for i := 0; i < len(cSlice); i++ {
		InternalCombustionEngineStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InternalCombustionEngineStatus__Array_to_C(cSlice []CInternalCombustionEngineStatus, goSlice []InternalCombustionEngineStatus) {
	for i := 0; i < len(goSlice); i++ {
		InternalCombustionEngineStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
