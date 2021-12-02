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

#include <px4_msgs/msg/estimator_innovations.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorInnovations", EstimatorInnovationsTypeSupport)
}

// Do not create instances of this type directly. Always use NewEstimatorInnovations
// function instead.
type EstimatorInnovations struct {
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

// NewEstimatorInnovations creates a new EstimatorInnovations with default values.
func NewEstimatorInnovations() *EstimatorInnovations {
	self := EstimatorInnovations{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorInnovations) Clone() *EstimatorInnovations {
	c := &EstimatorInnovations{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.GpsHvel = t.GpsHvel
	c.GpsVvel = t.GpsVvel
	c.GpsHpos = t.GpsHpos
	c.GpsVpos = t.GpsVpos
	c.EvHvel = t.EvHvel
	c.EvVvel = t.EvVvel
	c.EvHpos = t.EvHpos
	c.EvVpos = t.EvVpos
	c.RngVpos = t.RngVpos
	c.BaroVpos = t.BaroVpos
	c.AuxHvel = t.AuxHvel
	c.AuxVvel = t.AuxVvel
	c.Flow = t.Flow
	c.Heading = t.Heading
	c.MagField = t.MagField
	c.Drag = t.Drag
	c.Airspeed = t.Airspeed
	c.Beta = t.Beta
	c.Hagl = t.Hagl
	return c
}

func (t *EstimatorInnovations) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorInnovations) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.GpsHvel = [2]float32{}
	t.GpsVvel = 0
	t.GpsHpos = [2]float32{}
	t.GpsVpos = 0
	t.EvHvel = [2]float32{}
	t.EvVvel = 0
	t.EvHpos = [2]float32{}
	t.EvVpos = 0
	t.RngVpos = 0
	t.BaroVpos = 0
	t.AuxHvel = [2]float32{}
	t.AuxVvel = 0
	t.Flow = [2]float32{}
	t.Heading = 0
	t.MagField = [3]float32{}
	t.Drag = [2]float32{}
	t.Airspeed = 0
	t.Beta = 0
	t.Hagl = 0
}

// CloneEstimatorInnovationsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorInnovationsSlice(dst, src []EstimatorInnovations) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorInnovationsTypeSupport types.MessageTypeSupport = _EstimatorInnovationsTypeSupport{}

type _EstimatorInnovationsTypeSupport struct{}

func (t _EstimatorInnovationsTypeSupport) New() types.Message {
	return NewEstimatorInnovations()
}

func (t _EstimatorInnovationsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorInnovations
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorInnovations__create())
}

func (t _EstimatorInnovationsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorInnovations__destroy((*C.px4_msgs__msg__EstimatorInnovations)(pointer_to_free))
}

func (t _EstimatorInnovationsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorInnovations)
	mem := (*C.px4_msgs__msg__EstimatorInnovations)(dst)
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

func (t _EstimatorInnovationsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorInnovations)
	mem := (*C.px4_msgs__msg__EstimatorInnovations)(ros2_message_buffer)
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

func (t _EstimatorInnovationsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorInnovations())
}

type CEstimatorInnovations = C.px4_msgs__msg__EstimatorInnovations
type CEstimatorInnovations__Sequence = C.px4_msgs__msg__EstimatorInnovations__Sequence

func EstimatorInnovations__Sequence_to_Go(goSlice *[]EstimatorInnovations, cSlice CEstimatorInnovations__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorInnovations, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorInnovations__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorInnovations * uintptr(i)),
		))
		EstimatorInnovationsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorInnovations__Sequence_to_C(cSlice *CEstimatorInnovations__Sequence, goSlice []EstimatorInnovations) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorInnovations)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorInnovations * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorInnovations)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorInnovations * uintptr(i)),
		))
		EstimatorInnovationsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorInnovations__Array_to_Go(goSlice []EstimatorInnovations, cSlice []CEstimatorInnovations) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorInnovationsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorInnovations__Array_to_C(cSlice []CEstimatorInnovations, goSlice []EstimatorInnovations) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorInnovationsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
