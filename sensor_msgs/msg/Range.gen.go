/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/range.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Range", RangeTypeSupport)
}
const (
	Range_ULTRASOUND uint8 = 0// Radiation type enumsIf you want a value added to this list, send an email to the ros-users list
	Range_INFRARED uint8 = 1// Radiation type enumsIf you want a value added to this list, send an email to the ros-users list
)

// Do not create instances of this type directly. Always use NewRange
// function instead.
type Range struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp in the header is the time the ranger
	RadiationType uint8 `yaml:"radiation_type"`// the type of radiation used by the sensor
	FieldOfView float32 `yaml:"field_of_view"`// the size of the arc that the distance reading is
	MinRange float32 `yaml:"min_range"`// minimum range value [m]
	MaxRange float32 `yaml:"max_range"`// maximum range value [m]
	Range float32 `yaml:"range"`// range data [m]
}

// NewRange creates a new Range with default values.
func NewRange() *Range {
	self := Range{}
	self.SetDefaults()
	return &self
}

func (t *Range) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Range) SetDefaults() {
	t.Header.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var RangeTypeSupport types.MessageTypeSupport = _RangeTypeSupport{}

type _RangeTypeSupport struct{}

func (t _RangeTypeSupport) New() types.Message {
	return NewRange()
}

func (t _RangeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Range
	return (unsafe.Pointer)(C.sensor_msgs__msg__Range__create())
}

func (t _RangeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Range__destroy((*C.sensor_msgs__msg__Range)(pointer_to_free))
}

func (t _RangeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Range)
	mem := (*C.sensor_msgs__msg__Range)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.radiation_type = C.uint8_t(m.RadiationType)
	mem.field_of_view = C.float(m.FieldOfView)
	mem.min_range = C.float(m.MinRange)
	mem.max_range = C.float(m.MaxRange)
	mem._range = C.float(m.Range)
}

func (t _RangeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Range)
	mem := (*C.sensor_msgs__msg__Range)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.RadiationType = uint8(mem.radiation_type)
	m.FieldOfView = float32(mem.field_of_view)
	m.MinRange = float32(mem.min_range)
	m.MaxRange = float32(mem.max_range)
	m.Range = float32(mem._range)
}

func (t _RangeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Range())
}

type CRange = C.sensor_msgs__msg__Range
type CRange__Sequence = C.sensor_msgs__msg__Range__Sequence

func Range__Sequence_to_Go(goSlice *[]Range, cSlice CRange__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Range, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__Range__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Range * uintptr(i)),
		))
		RangeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Range__Sequence_to_C(cSlice *CRange__Sequence, goSlice []Range) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Range)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__Range * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__Range)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Range * uintptr(i)),
		))
		RangeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Range__Array_to_Go(goSlice []Range, cSlice []CRange) {
	for i := 0; i < len(cSlice); i++ {
		RangeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Range__Array_to_C(cSlice []CRange, goSlice []Range) {
	for i := 0; i < len(goSlice); i++ {
		RangeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
