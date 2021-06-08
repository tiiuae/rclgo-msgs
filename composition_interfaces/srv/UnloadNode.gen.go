/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package composition_interfaces_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lcomposition_interfaces__rosidl_typesupport_c -lcomposition_interfaces__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <composition_interfaces/srv/unload_node.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("composition_interfaces/UnloadNode", UnloadNodeTypeSupport)
}

type _UnloadNodeTypeSupport struct {}

func (s _UnloadNodeTypeSupport) Request() types.MessageTypeSupport {
	return UnloadNode_RequestTypeSupport
}

func (s _UnloadNodeTypeSupport) Response() types.MessageTypeSupport {
	return UnloadNode_ResponseTypeSupport
}

func (s _UnloadNodeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__composition_interfaces__srv__UnloadNode())
}

// Modifying this variable is undefined behavior.
var UnloadNodeTypeSupport types.ServiceTypeSupport = _UnloadNodeTypeSupport{}
