/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <map_msgs/srv/save_map.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("map_msgs/SaveMap", SaveMapTypeSupport)
}

type _SaveMapTypeSupport struct {}

func (s _SaveMapTypeSupport) Request() types.MessageTypeSupport {
	return SaveMap_RequestTypeSupport
}

func (s _SaveMapTypeSupport) Response() types.MessageTypeSupport {
	return SaveMap_ResponseTypeSupport
}

func (s _SaveMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__map_msgs__srv__SaveMap())
}

// Modifying this variable is undefined behavior.
var SaveMapTypeSupport types.ServiceTypeSupport = _SaveMapTypeSupport{}