// MIT License (MIT)

// WARNING: This file has automatically been generated on Sun, 13 Aug 2017 23:59:44 CEST.
// By https://git.io/c-for-go. DO NOT EDIT.

package uinput

/*
#include <linux/uinput.h>
#include <linux/input-event-codes.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// InputEvent as declared in linux/input.h:23
type InputEvent struct {
	Type           uint16
	Code           uint16
	Value          int32
	ref44438961    *C.struct_input_event
	allocs44438961 interface{}
}

// InputId as declared in linux/input.h:40
type InputId struct {
	Bustype        uint16
	Vendor         uint16
	Product        uint16
	Version        uint16
	refb2752280    *C.struct_input_id
	allocsb2752280 interface{}
}

// UInputAbsSetup as declared in linux/uinput.h:101
type UInputAbsSetup struct {
	Code           uint16
	ref751a0800    *C.struct_uinput_abs_setup
	allocs751a0800 interface{}
}

// UInputFfErase as declared in linux/uinput.h:55
type UInputFfErase struct {
	RequestId      uint32
	Retval         int32
	EffectId       uint32
	reff7f7edee    *C.struct_uinput_ff_erase
	allocsf7f7edee interface{}
}

// UInputFfUpload as declared in linux/uinput.h:48
type UInputFfUpload struct {
	RequestId      uint32
	Retval         int32
	ref25b46468    *C.struct_uinput_ff_upload
	allocs25b46468 interface{}
}

// UInputSetup as declared in linux/uinput.h:66
type UInputSetup struct {
	Id             InputId
	Name           [80]byte
	FfEffectsMax   uint32
	ref8f941cb8    *C.struct_uinput_setup
	allocs8f941cb8 interface{}
}

// UInputUserDev as declared in linux/uinput.h:222
type UInputUserDev struct {
	Name           [80]byte
	Id             InputId
	FfEffectsMax   uint32
	Absmax         [64]int32
	Absmin         [64]int32
	Absfuzz        [64]int32
	Absflat        [64]int32
	ref8bb14617    *C.struct_uinput_user_dev
	allocs8bb14617 interface{}
}
