/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2017 Stany MARCEL <stanypub@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package uinput

import (
	"C"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

const (
	UI_SET_EVBIT  = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((100) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
	UI_SET_KEYBIT = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((101) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
	UI_SET_RELBIT = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((102) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
	UI_SET_ABSBIT = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((103) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
)

type UInput struct {
	fd            *os.File
	keyPressedMap map[EventCode]bool
}

type EventCode uint16

type EventValue int32

type AxisSetup struct {
	Code                 EventCode
	Min, Max, Fuzz, Flat EventValue
}

func (ui *UInput) Open() error {
	// open uinput device
	var err error
	if ui.fd, err = os.OpenFile("/dev/uinput", os.O_WRONLY, 0); err != nil {
		log.Fatal(err)
	}
	return err
}

func (ui *UInput) Init(name string,
	vendor, product, version uint16,
	keys, rels []EventCode,
	axes []AxisSetup,
	keyboard bool) error {

	if ui.fd == nil {
		if err := ui.Open(); err != nil {
			return nil
		}
	}

	var uidev UInputUserDev

	// copy Name
	for i, v := range name {
		if i >= len(uidev.Name)-1 {
			break
		}
		uidev.Name[i] = byte(v)
	}

	uidev.Id.Bustype = BUS_USB
	uidev.Id.Vendor = vendor
	uidev.Id.Product = product
	uidev.Id.Version = version

	if err := ui.setupKeys(keys); err != nil {
		return err
	}
	if err := ui.setupAxes(axes, &uidev); err != nil {
		return err
	}
	if err := ui.setupRels(rels); err != nil {
		return err
	}
	if keyboard {
		if err := ui.setupKeyboard(); err != nil {
			return err
		}
	}
	ref, _ := uidev.PassRef()
	defer uidev.Free()

	// Write
	ui.fd.Write((*[sizeOfStructUinputUserDevValue]byte)(unsafe.Pointer(ref))[:])

	if err := ioctl(ui.fd, UI_DEV_CREATE, 0); err != nil {
		return err
	}
	return nil
}

func (ui *UInput) setupKeys(keys []EventCode) error {

	if len(keys) > 0 {
		if err := ioctl(ui.fd, UI_SET_EVBIT, EV_KEY); err != nil {
			return err
		}
	}
	for _, key := range keys {
		if err := ioctl(ui.fd, UI_SET_KEYBIT, key); err != nil {
			return err
		}
	}
	return nil
}

func (ui *UInput) setupRels(rels []EventCode) error {

	if len(rels) > 0 {
		if err := ioctl(ui.fd, UI_SET_EVBIT, EV_REL); err != nil {
			return err
		}
	}
	for _, rel := range rels {
		if err := ioctl(ui.fd, UI_SET_RELBIT, rel); err != nil {
			return err
		}
	}
	return nil
}

func (ui *UInput) setupAxes(axes []AxisSetup, uidev *UInputUserDev) error {
	if len(axes) > 0 {
		if err := ioctl(ui.fd, UI_SET_EVBIT, EV_ABS); err != nil {
			return err
		}
	}
	for _, axis := range axes {
		if err := ioctl(ui.fd, UI_SET_ABSBIT, axis.Code); err != nil {
			return err
		}
		uidev.Absmin[axis.Code] = int32(axis.Min)
		uidev.Absmax[axis.Code] = int32(axis.Max)
		uidev.Absfuzz[axis.Code] = int32(axis.Fuzz)
		uidev.Absflat[axis.Code] = int32(axis.Flat)
	}
	return nil
}

func (ui *UInput) setupKeyboard() error {
	if err := ioctl(ui.fd, UI_SET_EVBIT, EV_MSC); err != nil {
		return err
	}
	if err := ioctl(ui.fd, UI_SET_EVBIT, MSC_SCAN); err != nil {
		return err
	}
	if err := ioctl(ui.fd, UI_SET_EVBIT, EV_REP); err != nil {
		return err
	}
	return nil
}

func (ui *UInput) Close() {

	if ui.fd != nil {
		ioctl(ui.fd, UI_DEV_DESTROY, 0)
		ui.fd.Close()
	}
}

func ioctl(fd *os.File, request uintptr, data interface{}) error {
	var arg uintptr

	switch dt := data.(type) {
	case uintptr:
		arg = dt
	case EventCode:
		arg = uintptr(dt)
	case int:
		arg = uintptr(dt)
	default:
		return fmt.Errorf("ioctl: Unable to convert: %T", data)
	}

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), request, arg); err != 0 {
		return fmt.Errorf("ioctl: %s", err.Error())
	}

	return nil
}

func (ui *UInput) genEvent(eventType uint16, code EventCode, value EventValue) error {

	ev := InputEvent{
		Type:  eventType,
		Code:  uint16(code),
		Value: int32(value),
	}

	ref, _ := ev.PassRef()
	defer ev.Free()

	_, err := ui.fd.Write((*[sizeOfStructInputEventValue]byte)(unsafe.Pointer(ref))[:])

	return err
}

func (ui *UInput) KeyEvent(code EventCode, value EventValue) error {
	return ui.genEvent(EV_KEY, code, value)
}

func (ui *UInput) AbsEvent(code EventCode, value EventValue) error {
	return ui.genEvent(EV_ABS, code, value)
}

func (ui *UInput) RelEvent(code EventCode, value EventValue) error {
	return ui.genEvent(EV_REL, code, value)
}

func (ui *UInput) ScanEvent(value EventValue) error {
	return ui.genEvent(EV_MSC, MSC_SCAN, value)
}

func (ui *UInput) SynEvent() error {
	return ui.genEvent(EV_SYN, SYN_REPORT, 0)
}

func (ui *UInput) SetDelayPeriod(delay, period EventValue) error {
	if err := ui.genEvent(EV_REP, REP_DELAY, delay); err != nil {
		return err
	}
	if err := ui.genEvent(EV_REP, REP_PERIOD, period); err != nil {
		return err
	}
	return nil
}

// ScanEventValues by keys recorded from a logitech keyboard
var scans = map[EventCode]EventValue{
	KEY_ESC:          0x70029,
	KEY_F1:           0x7003a,
	KEY_F2:           0x7003b,
	KEY_F3:           0x7003c,
	KEY_F4:           0x7003d,
	KEY_F5:           0x7003e,
	KEY_F6:           0x7003f,
	KEY_F7:           0x70040,
	KEY_F8:           0x70041,
	KEY_F9:           0x70042,
	KEY_F10:          0x70043,
	KEY_F11:          0x70044,
	KEY_F12:          0x70045,
	KEY_SYSRQ:        0x70046,
	KEY_SCROLLLOCK:   0x70047,
	KEY_PAUSE:        0x70048,
	KEY_GRAVE:        0x70035,
	KEY_1:            0x7001e,
	KEY_2:            0x7001f,
	KEY_3:            0x70020,
	KEY_4:            0x70021,
	KEY_5:            0x70022,
	KEY_6:            0x70023,
	KEY_7:            0x70024,
	KEY_8:            0x70025,
	KEY_9:            0x70026,
	KEY_0:            0x70027,
	KEY_MINUS:        0x7002d,
	KEY_EQUAL:        0x7002e,
	KEY_BACKSPACE:    0x7002a,
	KEY_TAB:          0x7002b,
	KEY_Q:            0x70014,
	KEY_W:            0x7001a,
	KEY_E:            0x70008,
	KEY_R:            0x70015,
	KEY_T:            0x70017,
	KEY_Y:            0x7001c,
	KEY_U:            0x70018,
	KEY_I:            0x7000c,
	KEY_O:            0x70012,
	KEY_P:            0x70013,
	KEY_LEFTBRACE:    0x7002f,
	KEY_RIGHTBRACE:   0x70030,
	KEY_ENTER:        0x70028,
	KEY_CAPSLOCK:     0x70039,
	KEY_A:            0x70004,
	KEY_S:            0x70016,
	KEY_D:            0x70007,
	KEY_F:            0x70009,
	KEY_G:            0x7000a,
	KEY_H:            0x7000b,
	KEY_J:            0x7000d,
	KEY_K:            0x7000e,
	KEY_L:            0x7000f,
	KEY_SEMICOLON:    0x70033,
	KEY_APOSTROPHE:   0x70034,
	KEY_BACKSLASH:    0x70032,
	KEY_LEFTSHIFT:    0x700e1,
	KEY_102ND:        0x70064,
	KEY_Z:            0x7001d,
	KEY_X:            0x7001b,
	KEY_C:            0x70006,
	KEY_V:            0x70019,
	KEY_B:            0x70005,
	KEY_N:            0x70011,
	KEY_M:            0x70010,
	KEY_COMMA:        0x70036,
	KEY_DOT:          0x70037,
	KEY_SLASH:        0x70038,
	KEY_RIGHTSHIFT:   0x700e5,
	KEY_LEFTCTRL:     0x700e0,
	KEY_LEFTMETA:     0x700e3,
	KEY_LEFTALT:      0x700e2,
	KEY_SPACE:        0x7002c,
	KEY_RIGHTALT:     0x700e6,
	KEY_RIGHTMETA:    0x700e7,
	KEY_COMPOSE:      0x70065,
	KEY_RIGHTCTRL:    0x700e4,
	KEY_INSERT:       0x70049,
	KEY_HOME:         0x7004a,
	KEY_PAGEUP:       0x7004b,
	KEY_DELETE:       0x7004c,
	KEY_END:          0x7004d,
	KEY_PAGEDOWN:     0x7004e,
	KEY_UP:           0x70052,
	KEY_LEFT:         0x70050,
	KEY_DOWN:         0x70051,
	KEY_RIGHT:        0x7004f,
	KEY_NUMLOCK:      0x70053,
	KEY_KPSLASH:      0x70054,
	KEY_KPASTERISK:   0x70055,
	KEY_KPMINUS:      0x70056,
	KEY_KP7:          0x7005f,
	KEY_KP8:          0x70060,
	KEY_KP9:          0x70061,
	KEY_KPPLUS:       0x70057,
	KEY_KP4:          0x7005c,
	KEY_KP5:          0x7005d,
	KEY_KP6:          0x7005e,
	KEY_KP1:          0x70059,
	KEY_KP2:          0x7005a,
	KEY_KP3:          0x7005b,
	KEY_KPENTER:      0x70058,
	KEY_KP0:          0x70062,
	KEY_KPDOT:        0x70063,
	KEY_CONFIG:       0xc0183,
	KEY_PLAYPAUSE:    0xc00cd,
	KEY_MUTE:         0xc00e2,
	KEY_VOLUMEDOWN:   0xc00ea,
	KEY_VOLUMEUP:     0xc00e9,
	KEY_HOMEPAGE:     0xc0223,
	KEY_PREVIOUSSONG: 0xc00f0,
	KEY_NEXTSONG:     0xc00f1,
	KEY_BACK:         0xc00f2,
	KEY_FORWARD:      0xc00f3,
}

func (ui *UInput) KeyPressed(codes []EventCode) error {
	gensyn := false
	for _, k := range codes {
		if _, isin := ui.keyPressedMap[k]; !isin {
			gensyn = true
			ui.keyPressedMap[k] = true
			if err := ui.ScanEvent(scans[k]); err != nil {
				return err
			}
			if err := ui.KeyEvent(k, 1); err != nil {
				return err
			}
		}
	}
	if gensyn {
		return ui.SynEvent()
	}
	return nil
}

func (ui *UInput) KeyReleased(codes []EventCode) error {
	gensyn := false

	// with no argument release all pressed keys
	if len(codes) <= 0 {
		for k, _ := range ui.keyPressedMap {
			gensyn = true
			delete(ui.keyPressedMap, k)
			if err := ui.ScanEvent(scans[k]); err != nil {
				return err
			}
			if err := ui.KeyEvent(k, 0); err != nil {
				return err
			}
		}
	}

	for _, k := range codes {
		if _, isin := ui.keyPressedMap[k]; isin {
			gensyn = true
			delete(ui.keyPressedMap, k)
			if err := ui.ScanEvent(scans[k]); err != nil {
				return err
			}
			if err := ui.KeyEvent(k, 0); err != nil {
				return err
			}
		}
	}
	if gensyn {
		return ui.SynEvent()
	}
	return nil
}
