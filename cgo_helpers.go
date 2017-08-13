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
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocStructInputEventMemory allocates memory for type C.struct_input_event in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructInputEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructInputEventValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructInputEventValue = unsafe.Sizeof([1]C.struct_input_event{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *InputEvent) Ref() *C.struct_input_event {
	if x == nil {
		return nil
	}
	return x.ref44438961
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *InputEvent) Free() {
	if x != nil && x.allocs44438961 != nil {
		x.allocs44438961.(*cgoAllocMap).Free()
		x.ref44438961 = nil
	}
}

// NewInputEventRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewInputEventRef(ref unsafe.Pointer) *InputEvent {
	if ref == nil {
		return nil
	}
	obj := new(InputEvent)
	obj.ref44438961 = (*C.struct_input_event)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *InputEvent) PassRef() (*C.struct_input_event, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref44438961 != nil {
		return x.ref44438961, nil
	}
	mem44438961 := allocStructInputEventMemory(1)
	ref44438961 := (*C.struct_input_event)(mem44438961)
	allocs44438961 := new(cgoAllocMap)
	var c_type_allocs *cgoAllocMap
	ref44438961._type, c_type_allocs = (C.__u16)(x.Type), cgoAllocsUnknown
	allocs44438961.Borrow(c_type_allocs)

	var ccode_allocs *cgoAllocMap
	ref44438961.code, ccode_allocs = (C.__u16)(x.Code), cgoAllocsUnknown
	allocs44438961.Borrow(ccode_allocs)

	var cvalue_allocs *cgoAllocMap
	ref44438961.value, cvalue_allocs = (C.__s32)(x.Value), cgoAllocsUnknown
	allocs44438961.Borrow(cvalue_allocs)

	x.ref44438961 = ref44438961
	x.allocs44438961 = allocs44438961
	return ref44438961, allocs44438961

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x InputEvent) PassValue() (C.struct_input_event, *cgoAllocMap) {
	if x.ref44438961 != nil {
		return *x.ref44438961, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *InputEvent) Deref() {
	if x.ref44438961 == nil {
		return
	}
	x.Type = (uint16)(x.ref44438961._type)
	x.Code = (uint16)(x.ref44438961.code)
	x.Value = (int32)(x.ref44438961.value)
}

// allocStructInputIdMemory allocates memory for type C.struct_input_id in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructInputIdMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructInputIdValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructInputIdValue = unsafe.Sizeof([1]C.struct_input_id{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *InputId) Ref() *C.struct_input_id {
	if x == nil {
		return nil
	}
	return x.refb2752280
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *InputId) Free() {
	if x != nil && x.allocsb2752280 != nil {
		x.allocsb2752280.(*cgoAllocMap).Free()
		x.refb2752280 = nil
	}
}

// NewInputIdRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewInputIdRef(ref unsafe.Pointer) *InputId {
	if ref == nil {
		return nil
	}
	obj := new(InputId)
	obj.refb2752280 = (*C.struct_input_id)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *InputId) PassRef() (*C.struct_input_id, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb2752280 != nil {
		return x.refb2752280, nil
	}
	memb2752280 := allocStructInputIdMemory(1)
	refb2752280 := (*C.struct_input_id)(memb2752280)
	allocsb2752280 := new(cgoAllocMap)
	var cbustype_allocs *cgoAllocMap
	refb2752280.bustype, cbustype_allocs = (C.__u16)(x.Bustype), cgoAllocsUnknown
	allocsb2752280.Borrow(cbustype_allocs)

	var cvendor_allocs *cgoAllocMap
	refb2752280.vendor, cvendor_allocs = (C.__u16)(x.Vendor), cgoAllocsUnknown
	allocsb2752280.Borrow(cvendor_allocs)

	var cproduct_allocs *cgoAllocMap
	refb2752280.product, cproduct_allocs = (C.__u16)(x.Product), cgoAllocsUnknown
	allocsb2752280.Borrow(cproduct_allocs)

	var cversion_allocs *cgoAllocMap
	refb2752280.version, cversion_allocs = (C.__u16)(x.Version), cgoAllocsUnknown
	allocsb2752280.Borrow(cversion_allocs)

	x.refb2752280 = refb2752280
	x.allocsb2752280 = allocsb2752280
	return refb2752280, allocsb2752280

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x InputId) PassValue() (C.struct_input_id, *cgoAllocMap) {
	if x.refb2752280 != nil {
		return *x.refb2752280, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *InputId) Deref() {
	if x.refb2752280 == nil {
		return
	}
	x.Bustype = (uint16)(x.refb2752280.bustype)
	x.Vendor = (uint16)(x.refb2752280.vendor)
	x.Product = (uint16)(x.refb2752280.product)
	x.Version = (uint16)(x.refb2752280.version)
}

// allocStructUinputAbsSetupMemory allocates memory for type C.struct_uinput_abs_setup in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructUinputAbsSetupMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructUinputAbsSetupValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructUinputAbsSetupValue = unsafe.Sizeof([1]C.struct_uinput_abs_setup{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *UInputAbsSetup) Ref() *C.struct_uinput_abs_setup {
	if x == nil {
		return nil
	}
	return x.ref751a0800
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *UInputAbsSetup) Free() {
	if x != nil && x.allocs751a0800 != nil {
		x.allocs751a0800.(*cgoAllocMap).Free()
		x.ref751a0800 = nil
	}
}

// NewUInputAbsSetupRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewUInputAbsSetupRef(ref unsafe.Pointer) *UInputAbsSetup {
	if ref == nil {
		return nil
	}
	obj := new(UInputAbsSetup)
	obj.ref751a0800 = (*C.struct_uinput_abs_setup)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *UInputAbsSetup) PassRef() (*C.struct_uinput_abs_setup, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref751a0800 != nil {
		return x.ref751a0800, nil
	}
	mem751a0800 := allocStructUinputAbsSetupMemory(1)
	ref751a0800 := (*C.struct_uinput_abs_setup)(mem751a0800)
	allocs751a0800 := new(cgoAllocMap)
	var ccode_allocs *cgoAllocMap
	ref751a0800.code, ccode_allocs = (C.__u16)(x.Code), cgoAllocsUnknown
	allocs751a0800.Borrow(ccode_allocs)

	x.ref751a0800 = ref751a0800
	x.allocs751a0800 = allocs751a0800
	return ref751a0800, allocs751a0800

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x UInputAbsSetup) PassValue() (C.struct_uinput_abs_setup, *cgoAllocMap) {
	if x.ref751a0800 != nil {
		return *x.ref751a0800, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *UInputAbsSetup) Deref() {
	if x.ref751a0800 == nil {
		return
	}
	x.Code = (uint16)(x.ref751a0800.code)
}

// allocStructUinputFfEraseMemory allocates memory for type C.struct_uinput_ff_erase in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructUinputFfEraseMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructUinputFfEraseValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructUinputFfEraseValue = unsafe.Sizeof([1]C.struct_uinput_ff_erase{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *UInputFfErase) Ref() *C.struct_uinput_ff_erase {
	if x == nil {
		return nil
	}
	return x.reff7f7edee
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *UInputFfErase) Free() {
	if x != nil && x.allocsf7f7edee != nil {
		x.allocsf7f7edee.(*cgoAllocMap).Free()
		x.reff7f7edee = nil
	}
}

// NewUInputFfEraseRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewUInputFfEraseRef(ref unsafe.Pointer) *UInputFfErase {
	if ref == nil {
		return nil
	}
	obj := new(UInputFfErase)
	obj.reff7f7edee = (*C.struct_uinput_ff_erase)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *UInputFfErase) PassRef() (*C.struct_uinput_ff_erase, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff7f7edee != nil {
		return x.reff7f7edee, nil
	}
	memf7f7edee := allocStructUinputFfEraseMemory(1)
	reff7f7edee := (*C.struct_uinput_ff_erase)(memf7f7edee)
	allocsf7f7edee := new(cgoAllocMap)
	var crequest_id_allocs *cgoAllocMap
	reff7f7edee.request_id, crequest_id_allocs = (C.__u32)(x.RequestId), cgoAllocsUnknown
	allocsf7f7edee.Borrow(crequest_id_allocs)

	var cretval_allocs *cgoAllocMap
	reff7f7edee.retval, cretval_allocs = (C.__s32)(x.Retval), cgoAllocsUnknown
	allocsf7f7edee.Borrow(cretval_allocs)

	var ceffect_id_allocs *cgoAllocMap
	reff7f7edee.effect_id, ceffect_id_allocs = (C.__u32)(x.EffectId), cgoAllocsUnknown
	allocsf7f7edee.Borrow(ceffect_id_allocs)

	x.reff7f7edee = reff7f7edee
	x.allocsf7f7edee = allocsf7f7edee
	return reff7f7edee, allocsf7f7edee

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x UInputFfErase) PassValue() (C.struct_uinput_ff_erase, *cgoAllocMap) {
	if x.reff7f7edee != nil {
		return *x.reff7f7edee, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *UInputFfErase) Deref() {
	if x.reff7f7edee == nil {
		return
	}
	x.RequestId = (uint32)(x.reff7f7edee.request_id)
	x.Retval = (int32)(x.reff7f7edee.retval)
	x.EffectId = (uint32)(x.reff7f7edee.effect_id)
}

// allocStructUinputFfUploadMemory allocates memory for type C.struct_uinput_ff_upload in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructUinputFfUploadMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructUinputFfUploadValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructUinputFfUploadValue = unsafe.Sizeof([1]C.struct_uinput_ff_upload{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *UInputFfUpload) Ref() *C.struct_uinput_ff_upload {
	if x == nil {
		return nil
	}
	return x.ref25b46468
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *UInputFfUpload) Free() {
	if x != nil && x.allocs25b46468 != nil {
		x.allocs25b46468.(*cgoAllocMap).Free()
		x.ref25b46468 = nil
	}
}

// NewUInputFfUploadRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewUInputFfUploadRef(ref unsafe.Pointer) *UInputFfUpload {
	if ref == nil {
		return nil
	}
	obj := new(UInputFfUpload)
	obj.ref25b46468 = (*C.struct_uinput_ff_upload)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *UInputFfUpload) PassRef() (*C.struct_uinput_ff_upload, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref25b46468 != nil {
		return x.ref25b46468, nil
	}
	mem25b46468 := allocStructUinputFfUploadMemory(1)
	ref25b46468 := (*C.struct_uinput_ff_upload)(mem25b46468)
	allocs25b46468 := new(cgoAllocMap)
	var crequest_id_allocs *cgoAllocMap
	ref25b46468.request_id, crequest_id_allocs = (C.__u32)(x.RequestId), cgoAllocsUnknown
	allocs25b46468.Borrow(crequest_id_allocs)

	var cretval_allocs *cgoAllocMap
	ref25b46468.retval, cretval_allocs = (C.__s32)(x.Retval), cgoAllocsUnknown
	allocs25b46468.Borrow(cretval_allocs)

	x.ref25b46468 = ref25b46468
	x.allocs25b46468 = allocs25b46468
	return ref25b46468, allocs25b46468

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x UInputFfUpload) PassValue() (C.struct_uinput_ff_upload, *cgoAllocMap) {
	if x.ref25b46468 != nil {
		return *x.ref25b46468, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *UInputFfUpload) Deref() {
	if x.ref25b46468 == nil {
		return
	}
	x.RequestId = (uint32)(x.ref25b46468.request_id)
	x.Retval = (int32)(x.ref25b46468.retval)
}

// allocStructUinputSetupMemory allocates memory for type C.struct_uinput_setup in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructUinputSetupMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructUinputSetupValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructUinputSetupValue = unsafe.Sizeof([1]C.struct_uinput_setup{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *UInputSetup) Ref() *C.struct_uinput_setup {
	if x == nil {
		return nil
	}
	return x.ref8f941cb8
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *UInputSetup) Free() {
	if x != nil && x.allocs8f941cb8 != nil {
		x.allocs8f941cb8.(*cgoAllocMap).Free()
		x.ref8f941cb8 = nil
	}
}

// NewUInputSetupRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewUInputSetupRef(ref unsafe.Pointer) *UInputSetup {
	if ref == nil {
		return nil
	}
	obj := new(UInputSetup)
	obj.ref8f941cb8 = (*C.struct_uinput_setup)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *UInputSetup) PassRef() (*C.struct_uinput_setup, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8f941cb8 != nil {
		return x.ref8f941cb8, nil
	}
	mem8f941cb8 := allocStructUinputSetupMemory(1)
	ref8f941cb8 := (*C.struct_uinput_setup)(mem8f941cb8)
	allocs8f941cb8 := new(cgoAllocMap)
	var cid_allocs *cgoAllocMap
	ref8f941cb8.id, cid_allocs = x.Id.PassValue()
	allocs8f941cb8.Borrow(cid_allocs)

	var cname_allocs *cgoAllocMap
	ref8f941cb8.name, cname_allocs = *(*[80]C.char)(unsafe.Pointer(&x.Name)), cgoAllocsUnknown
	allocs8f941cb8.Borrow(cname_allocs)

	var cff_effects_max_allocs *cgoAllocMap
	ref8f941cb8.ff_effects_max, cff_effects_max_allocs = (C.__u32)(x.FfEffectsMax), cgoAllocsUnknown
	allocs8f941cb8.Borrow(cff_effects_max_allocs)

	x.ref8f941cb8 = ref8f941cb8
	x.allocs8f941cb8 = allocs8f941cb8
	return ref8f941cb8, allocs8f941cb8

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x UInputSetup) PassValue() (C.struct_uinput_setup, *cgoAllocMap) {
	if x.ref8f941cb8 != nil {
		return *x.ref8f941cb8, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *UInputSetup) Deref() {
	if x.ref8f941cb8 == nil {
		return
	}
	x.Id = *NewInputIdRef(unsafe.Pointer(&x.ref8f941cb8.id))
	x.Name = *(*[80]byte)(unsafe.Pointer(&x.ref8f941cb8.name))
	x.FfEffectsMax = (uint32)(x.ref8f941cb8.ff_effects_max)
}

// allocStructUinputUserDevMemory allocates memory for type C.struct_uinput_user_dev in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructUinputUserDevMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructUinputUserDevValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructUinputUserDevValue = unsafe.Sizeof([1]C.struct_uinput_user_dev{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *UInputUserDev) Ref() *C.struct_uinput_user_dev {
	if x == nil {
		return nil
	}
	return x.ref8bb14617
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *UInputUserDev) Free() {
	if x != nil && x.allocs8bb14617 != nil {
		x.allocs8bb14617.(*cgoAllocMap).Free()
		x.ref8bb14617 = nil
	}
}

// NewUInputUserDevRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewUInputUserDevRef(ref unsafe.Pointer) *UInputUserDev {
	if ref == nil {
		return nil
	}
	obj := new(UInputUserDev)
	obj.ref8bb14617 = (*C.struct_uinput_user_dev)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *UInputUserDev) PassRef() (*C.struct_uinput_user_dev, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8bb14617 != nil {
		return x.ref8bb14617, nil
	}
	mem8bb14617 := allocStructUinputUserDevMemory(1)
	ref8bb14617 := (*C.struct_uinput_user_dev)(mem8bb14617)
	allocs8bb14617 := new(cgoAllocMap)
	var cname_allocs *cgoAllocMap
	ref8bb14617.name, cname_allocs = *(*[80]C.char)(unsafe.Pointer(&x.Name)), cgoAllocsUnknown
	allocs8bb14617.Borrow(cname_allocs)

	var cid_allocs *cgoAllocMap
	ref8bb14617.id, cid_allocs = x.Id.PassValue()
	allocs8bb14617.Borrow(cid_allocs)

	var cff_effects_max_allocs *cgoAllocMap
	ref8bb14617.ff_effects_max, cff_effects_max_allocs = (C.__u32)(x.FfEffectsMax), cgoAllocsUnknown
	allocs8bb14617.Borrow(cff_effects_max_allocs)

	var cabsmax_allocs *cgoAllocMap
	ref8bb14617.absmax, cabsmax_allocs = *(*[64]C.__s32)(unsafe.Pointer(&x.Absmax)), cgoAllocsUnknown
	allocs8bb14617.Borrow(cabsmax_allocs)

	var cabsmin_allocs *cgoAllocMap
	ref8bb14617.absmin, cabsmin_allocs = *(*[64]C.__s32)(unsafe.Pointer(&x.Absmin)), cgoAllocsUnknown
	allocs8bb14617.Borrow(cabsmin_allocs)

	var cabsfuzz_allocs *cgoAllocMap
	ref8bb14617.absfuzz, cabsfuzz_allocs = *(*[64]C.__s32)(unsafe.Pointer(&x.Absfuzz)), cgoAllocsUnknown
	allocs8bb14617.Borrow(cabsfuzz_allocs)

	var cabsflat_allocs *cgoAllocMap
	ref8bb14617.absflat, cabsflat_allocs = *(*[64]C.__s32)(unsafe.Pointer(&x.Absflat)), cgoAllocsUnknown
	allocs8bb14617.Borrow(cabsflat_allocs)

	x.ref8bb14617 = ref8bb14617
	x.allocs8bb14617 = allocs8bb14617
	return ref8bb14617, allocs8bb14617

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x UInputUserDev) PassValue() (C.struct_uinput_user_dev, *cgoAllocMap) {
	if x.ref8bb14617 != nil {
		return *x.ref8bb14617, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *UInputUserDev) Deref() {
	if x.ref8bb14617 == nil {
		return
	}
	x.Name = *(*[80]byte)(unsafe.Pointer(&x.ref8bb14617.name))
	x.Id = *NewInputIdRef(unsafe.Pointer(&x.ref8bb14617.id))
	x.FfEffectsMax = (uint32)(x.ref8bb14617.ff_effects_max)
	x.Absmax = *(*[64]int32)(unsafe.Pointer(&x.ref8bb14617.absmax))
	x.Absmin = *(*[64]int32)(unsafe.Pointer(&x.ref8bb14617.absmin))
	x.Absfuzz = *(*[64]int32)(unsafe.Pointer(&x.ref8bb14617.absfuzz))
	x.Absflat = *(*[64]int32)(unsafe.Pointer(&x.ref8bb14617.absflat))
}
