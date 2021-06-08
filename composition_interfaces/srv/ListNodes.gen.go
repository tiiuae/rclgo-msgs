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
#include <composition_interfaces/srv/list_nodes.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("composition_interfaces/ListNodes", ListNodesTypeSupport)
}

type _ListNodesTypeSupport struct {}

func (s _ListNodesTypeSupport) Request() types.MessageTypeSupport {
	return ListNodes_RequestTypeSupport
}

func (s _ListNodesTypeSupport) Response() types.MessageTypeSupport {
	return ListNodes_ResponseTypeSupport
}

func (s _ListNodesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__composition_interfaces__srv__ListNodes())
}

// Modifying this variable is undefined behavior.
var ListNodesTypeSupport types.ServiceTypeSupport = _ListNodesTypeSupport{}
