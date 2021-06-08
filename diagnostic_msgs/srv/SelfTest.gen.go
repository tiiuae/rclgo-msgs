/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package diagnostic_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <diagnostic_msgs/srv/self_test.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("diagnostic_msgs/SelfTest", SelfTestTypeSupport)
}

type _SelfTestTypeSupport struct {}

func (s _SelfTestTypeSupport) Request() types.MessageTypeSupport {
	return SelfTest_RequestTypeSupport
}

func (s _SelfTestTypeSupport) Response() types.MessageTypeSupport {
	return SelfTest_ResponseTypeSupport
}

func (s _SelfTestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__diagnostic_msgs__srv__SelfTest())
}

// Modifying this variable is undefined behavior.
var SelfTestTypeSupport types.ServiceTypeSupport = _SelfTestTypeSupport{}
