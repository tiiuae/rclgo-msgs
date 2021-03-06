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

#include <px4_msgs/msg/sensors_status_imu.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorsStatusImu", SensorsStatusImuTypeSupport)
}

// Do not create instances of this type directly. Always use NewSensorsStatusImu
// function instead.
type SensorsStatusImu struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). Sensor check metrics. This will be zero for a sensor that's primary or unpopulated.
	AccelDeviceIdPrimary uint32 `yaml:"accel_device_id_primary"`// current primary accel device id for reference
	AccelDeviceIds [4]uint32 `yaml:"accel_device_ids"`
	AccelInconsistencyMSS [4]float32 `yaml:"accel_inconsistency_m_s_s"`// magnitude of acceleration difference between IMU instance and mean in m/s^2.
	AccelHealthy [4]bool `yaml:"accel_healthy"`
	AccelPriority [4]uint8 `yaml:"accel_priority"`
	GyroDeviceIdPrimary uint32 `yaml:"gyro_device_id_primary"`// current primary gyro device id for reference
	GyroDeviceIds [4]uint32 `yaml:"gyro_device_ids"`
	GyroInconsistencyRadS [4]float32 `yaml:"gyro_inconsistency_rad_s"`// magnitude of angular rate difference between IMU instance and mean in (rad/s).
	GyroHealthy [4]bool `yaml:"gyro_healthy"`
	GyroPriority [4]uint8 `yaml:"gyro_priority"`
}

// NewSensorsStatusImu creates a new SensorsStatusImu with default values.
func NewSensorsStatusImu() *SensorsStatusImu {
	self := SensorsStatusImu{}
	self.SetDefaults()
	return &self
}

func (t *SensorsStatusImu) Clone() *SensorsStatusImu {
	c := &SensorsStatusImu{}
	c.Timestamp = t.Timestamp
	c.AccelDeviceIdPrimary = t.AccelDeviceIdPrimary
	c.AccelDeviceIds = t.AccelDeviceIds
	c.AccelInconsistencyMSS = t.AccelInconsistencyMSS
	c.AccelHealthy = t.AccelHealthy
	c.AccelPriority = t.AccelPriority
	c.GyroDeviceIdPrimary = t.GyroDeviceIdPrimary
	c.GyroDeviceIds = t.GyroDeviceIds
	c.GyroInconsistencyRadS = t.GyroInconsistencyRadS
	c.GyroHealthy = t.GyroHealthy
	c.GyroPriority = t.GyroPriority
	return c
}

func (t *SensorsStatusImu) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SensorsStatusImu) SetDefaults() {
	t.Timestamp = 0
	t.AccelDeviceIdPrimary = 0
	t.AccelDeviceIds = [4]uint32{}
	t.AccelInconsistencyMSS = [4]float32{}
	t.AccelHealthy = [4]bool{}
	t.AccelPriority = [4]uint8{}
	t.GyroDeviceIdPrimary = 0
	t.GyroDeviceIds = [4]uint32{}
	t.GyroInconsistencyRadS = [4]float32{}
	t.GyroHealthy = [4]bool{}
	t.GyroPriority = [4]uint8{}
}

// CloneSensorsStatusImuSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSensorsStatusImuSlice(dst, src []SensorsStatusImu) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SensorsStatusImuTypeSupport types.MessageTypeSupport = _SensorsStatusImuTypeSupport{}

type _SensorsStatusImuTypeSupport struct{}

func (t _SensorsStatusImuTypeSupport) New() types.Message {
	return NewSensorsStatusImu()
}

func (t _SensorsStatusImuTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorsStatusImu
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorsStatusImu__create())
}

func (t _SensorsStatusImuTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorsStatusImu__destroy((*C.px4_msgs__msg__SensorsStatusImu)(pointer_to_free))
}

func (t _SensorsStatusImuTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorsStatusImu)
	mem := (*C.px4_msgs__msg__SensorsStatusImu)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.accel_device_id_primary = C.uint32_t(m.AccelDeviceIdPrimary)
	cSlice_accel_device_ids := mem.accel_device_ids[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_accel_device_ids)), m.AccelDeviceIds[:])
	cSlice_accel_inconsistency_m_s_s := mem.accel_inconsistency_m_s_s[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_inconsistency_m_s_s)), m.AccelInconsistencyMSS[:])
	cSlice_accel_healthy := mem.accel_healthy[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_accel_healthy)), m.AccelHealthy[:])
	cSlice_accel_priority := mem.accel_priority[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_accel_priority)), m.AccelPriority[:])
	mem.gyro_device_id_primary = C.uint32_t(m.GyroDeviceIdPrimary)
	cSlice_gyro_device_ids := mem.gyro_device_ids[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_gyro_device_ids)), m.GyroDeviceIds[:])
	cSlice_gyro_inconsistency_rad_s := mem.gyro_inconsistency_rad_s[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_inconsistency_rad_s)), m.GyroInconsistencyRadS[:])
	cSlice_gyro_healthy := mem.gyro_healthy[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_gyro_healthy)), m.GyroHealthy[:])
	cSlice_gyro_priority := mem.gyro_priority[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_gyro_priority)), m.GyroPriority[:])
}

func (t _SensorsStatusImuTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorsStatusImu)
	mem := (*C.px4_msgs__msg__SensorsStatusImu)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.AccelDeviceIdPrimary = uint32(mem.accel_device_id_primary)
	cSlice_accel_device_ids := mem.accel_device_ids[:]
	primitives.Uint32__Array_to_Go(m.AccelDeviceIds[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_accel_device_ids)))
	cSlice_accel_inconsistency_m_s_s := mem.accel_inconsistency_m_s_s[:]
	primitives.Float32__Array_to_Go(m.AccelInconsistencyMSS[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_inconsistency_m_s_s)))
	cSlice_accel_healthy := mem.accel_healthy[:]
	primitives.Bool__Array_to_Go(m.AccelHealthy[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_accel_healthy)))
	cSlice_accel_priority := mem.accel_priority[:]
	primitives.Uint8__Array_to_Go(m.AccelPriority[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_accel_priority)))
	m.GyroDeviceIdPrimary = uint32(mem.gyro_device_id_primary)
	cSlice_gyro_device_ids := mem.gyro_device_ids[:]
	primitives.Uint32__Array_to_Go(m.GyroDeviceIds[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_gyro_device_ids)))
	cSlice_gyro_inconsistency_rad_s := mem.gyro_inconsistency_rad_s[:]
	primitives.Float32__Array_to_Go(m.GyroInconsistencyRadS[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_inconsistency_rad_s)))
	cSlice_gyro_healthy := mem.gyro_healthy[:]
	primitives.Bool__Array_to_Go(m.GyroHealthy[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_gyro_healthy)))
	cSlice_gyro_priority := mem.gyro_priority[:]
	primitives.Uint8__Array_to_Go(m.GyroPriority[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_gyro_priority)))
}

func (t _SensorsStatusImuTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorsStatusImu())
}

type CSensorsStatusImu = C.px4_msgs__msg__SensorsStatusImu
type CSensorsStatusImu__Sequence = C.px4_msgs__msg__SensorsStatusImu__Sequence

func SensorsStatusImu__Sequence_to_Go(goSlice *[]SensorsStatusImu, cSlice CSensorsStatusImu__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorsStatusImu, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorsStatusImu__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorsStatusImu * uintptr(i)),
		))
		SensorsStatusImuTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorsStatusImu__Sequence_to_C(cSlice *CSensorsStatusImu__Sequence, goSlice []SensorsStatusImu) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorsStatusImu)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorsStatusImu * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorsStatusImu)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorsStatusImu * uintptr(i)),
		))
		SensorsStatusImuTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorsStatusImu__Array_to_Go(goSlice []SensorsStatusImu, cSlice []CSensorsStatusImu) {
	for i := 0; i < len(cSlice); i++ {
		SensorsStatusImuTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorsStatusImu__Array_to_C(cSlice []CSensorsStatusImu, goSlice []SensorsStatusImu) {
	for i := 0; i < len(goSlice); i++ {
		SensorsStatusImuTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
