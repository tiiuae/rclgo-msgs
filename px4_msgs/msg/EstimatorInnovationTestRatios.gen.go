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

#include <px4_msgs/msg/estimator_innovation_test_ratios.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorInnovationTestRatios", EstimatorInnovationTestRatiosTypeSupport)
}

// Do not create instances of this type directly. Always use NewEstimatorInnovationTestRatios
// function instead.
type EstimatorInnovationTestRatios struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	GpsHvel [2]float32 `yaml:"gps_hvel"`// horizontal GPS velocity innovation (m/sec) and innovation variance ((m/sec)**2). GPS
	GpsVvel float32 `yaml:"gps_vvel"`// vertical GPS velocity innovation (m/sec) and innovation variance ((m/sec)**2). GPS
	GpsHpos [2]float32 `yaml:"gps_hpos"`// horizontal GPS position innovation (m) and innovation variance (m**2). GPS
	GpsVpos float32 `yaml:"gps_vpos"`// vertical GPS position innovation (m) and innovation variance (m**2). GPS
	EvHvel [2]float32 `yaml:"ev_hvel"`// horizontal external vision velocity innovation (m/sec) and innovation variance ((m/sec)**2). External Vision
	EvVvel float32 `yaml:"ev_vvel"`// vertical external vision velocity innovation (m/sec) and innovation variance ((m/sec)**2). External Vision
	EvHpos [2]float32 `yaml:"ev_hpos"`// horizontal external vision position innovation (m) and innovation variance (m**2). External Vision
	EvVpos float32 `yaml:"ev_vpos"`// vertical external vision position innovation (m) and innovation variance (m**2). External Vision
	RngVpos float32 `yaml:"rng_vpos"`// range sensor height innovation (m) and innovation variance (m**2). Height sensors
	BaroVpos float32 `yaml:"baro_vpos"`// barometer height innovation (m) and innovation variance (m**2). Height sensors
	AuxHvel [2]float32 `yaml:"aux_hvel"`// horizontal auxiliar velocity innovation from landing target measurement (m/sec) and innovation variance ((m/sec)**2). Auxiliary velocity
	AuxVvel float32 `yaml:"aux_vvel"`// vertical auxiliar velocity innovation from landing target measurement (m/sec) and innovation variance ((m/sec)**2). Auxiliary velocity
	Flow [2]float32 `yaml:"flow"`// flow innvoation (rad/sec) and innovation variance ((rad/sec)**2). Optical flow
	Heading float32 `yaml:"heading"`// heading innovation (rad) and innovation variance (rad**2). Various
	MagField [3]float32 `yaml:"mag_field"`// earth magnetic field innovation (Gauss) and innovation variance (Gauss**2). Various
	Drag [2]float32 `yaml:"drag"`// drag specific force innovation (m/sec**2) and innovation variance ((m/sec)**2). Various
	Airspeed float32 `yaml:"airspeed"`// airspeed innovation (m/sec) and innovation variance ((m/sec)**2). Various
	Beta float32 `yaml:"beta"`// synthetic sideslip innovation (rad) and innovation variance (rad**2). Various
	Hagl float32 `yaml:"hagl"`// height of ground innovation (m) and innovation variance (m**2). Various
}

// NewEstimatorInnovationTestRatios creates a new EstimatorInnovationTestRatios with default values.
func NewEstimatorInnovationTestRatios() *EstimatorInnovationTestRatios {
	self := EstimatorInnovationTestRatios{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorInnovationTestRatios) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *EstimatorInnovationTestRatios) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var EstimatorInnovationTestRatiosTypeSupport types.MessageTypeSupport = _EstimatorInnovationTestRatiosTypeSupport{}

type _EstimatorInnovationTestRatiosTypeSupport struct{}

func (t _EstimatorInnovationTestRatiosTypeSupport) New() types.Message {
	return NewEstimatorInnovationTestRatios()
}

func (t _EstimatorInnovationTestRatiosTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorInnovationTestRatios
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorInnovationTestRatios__create())
}

func (t _EstimatorInnovationTestRatiosTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorInnovationTestRatios__destroy((*C.px4_msgs__msg__EstimatorInnovationTestRatios)(pointer_to_free))
}

func (t _EstimatorInnovationTestRatiosTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorInnovationTestRatios)
	mem := (*C.px4_msgs__msg__EstimatorInnovationTestRatios)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_gps_hvel := mem.gps_hvel[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gps_hvel)), m.GpsHvel[:])
	mem.gps_vvel = C.float(m.GpsVvel)
	cSlice_gps_hpos := mem.gps_hpos[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gps_hpos)), m.GpsHpos[:])
	mem.gps_vpos = C.float(m.GpsVpos)
	cSlice_ev_hvel := mem.ev_hvel[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_ev_hvel)), m.EvHvel[:])
	mem.ev_vvel = C.float(m.EvVvel)
	cSlice_ev_hpos := mem.ev_hpos[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_ev_hpos)), m.EvHpos[:])
	mem.ev_vpos = C.float(m.EvVpos)
	mem.rng_vpos = C.float(m.RngVpos)
	mem.baro_vpos = C.float(m.BaroVpos)
	cSlice_aux_hvel := mem.aux_hvel[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_aux_hvel)), m.AuxHvel[:])
	mem.aux_vvel = C.float(m.AuxVvel)
	cSlice_flow := mem.flow[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_flow)), m.Flow[:])
	mem.heading = C.float(m.Heading)
	cSlice_mag_field := mem.mag_field[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_mag_field)), m.MagField[:])
	cSlice_drag := mem.drag[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_drag)), m.Drag[:])
	mem.airspeed = C.float(m.Airspeed)
	mem.beta = C.float(m.Beta)
	mem.hagl = C.float(m.Hagl)
}

func (t _EstimatorInnovationTestRatiosTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorInnovationTestRatios)
	mem := (*C.px4_msgs__msg__EstimatorInnovationTestRatios)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_gps_hvel := mem.gps_hvel[:]
	primitives.Float32__Array_to_Go(m.GpsHvel[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gps_hvel)))
	m.GpsVvel = float32(mem.gps_vvel)
	cSlice_gps_hpos := mem.gps_hpos[:]
	primitives.Float32__Array_to_Go(m.GpsHpos[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gps_hpos)))
	m.GpsVpos = float32(mem.gps_vpos)
	cSlice_ev_hvel := mem.ev_hvel[:]
	primitives.Float32__Array_to_Go(m.EvHvel[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_ev_hvel)))
	m.EvVvel = float32(mem.ev_vvel)
	cSlice_ev_hpos := mem.ev_hpos[:]
	primitives.Float32__Array_to_Go(m.EvHpos[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_ev_hpos)))
	m.EvVpos = float32(mem.ev_vpos)
	m.RngVpos = float32(mem.rng_vpos)
	m.BaroVpos = float32(mem.baro_vpos)
	cSlice_aux_hvel := mem.aux_hvel[:]
	primitives.Float32__Array_to_Go(m.AuxHvel[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_aux_hvel)))
	m.AuxVvel = float32(mem.aux_vvel)
	cSlice_flow := mem.flow[:]
	primitives.Float32__Array_to_Go(m.Flow[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_flow)))
	m.Heading = float32(mem.heading)
	cSlice_mag_field := mem.mag_field[:]
	primitives.Float32__Array_to_Go(m.MagField[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_mag_field)))
	cSlice_drag := mem.drag[:]
	primitives.Float32__Array_to_Go(m.Drag[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_drag)))
	m.Airspeed = float32(mem.airspeed)
	m.Beta = float32(mem.beta)
	m.Hagl = float32(mem.hagl)
}

func (t _EstimatorInnovationTestRatiosTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorInnovationTestRatios())
}

type CEstimatorInnovationTestRatios = C.px4_msgs__msg__EstimatorInnovationTestRatios
type CEstimatorInnovationTestRatios__Sequence = C.px4_msgs__msg__EstimatorInnovationTestRatios__Sequence

func EstimatorInnovationTestRatios__Sequence_to_Go(goSlice *[]EstimatorInnovationTestRatios, cSlice CEstimatorInnovationTestRatios__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorInnovationTestRatios, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorInnovationTestRatios__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorInnovationTestRatios * uintptr(i)),
		))
		EstimatorInnovationTestRatiosTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorInnovationTestRatios__Sequence_to_C(cSlice *CEstimatorInnovationTestRatios__Sequence, goSlice []EstimatorInnovationTestRatios) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorInnovationTestRatios)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorInnovationTestRatios * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorInnovationTestRatios)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorInnovationTestRatios * uintptr(i)),
		))
		EstimatorInnovationTestRatiosTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorInnovationTestRatios__Array_to_Go(goSlice []EstimatorInnovationTestRatios, cSlice []CEstimatorInnovationTestRatios) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorInnovationTestRatiosTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorInnovationTestRatios__Array_to_C(cSlice []CEstimatorInnovationTestRatios, goSlice []EstimatorInnovationTestRatios) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorInnovationTestRatiosTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}