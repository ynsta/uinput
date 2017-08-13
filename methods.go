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

/*
UINPUT_IOCTL_BASE
#define UI_SET_EVBIT		_IOW(UINPUT_IOCTL_BASE, 100, int)
#define UI_SET_KEYBIT		_IOW(UINPUT_IOCTL_BASE, 101, int)
#define UI_SET_RELBIT		_IOW(UINPUT_IOCTL_BASE, 102, int)
#define UI_SET_ABSBIT		_IOW(UINPUT_IOCTL_BASE, 103, int)
#define UI_SET_MSCBIT		_IOW(UINPUT_IOCTL_BASE, 104, int)
#define UI_SET_LEDBIT		_IOW(UINPUT_IOCTL_BASE, 105, int)
#define UI_SET_SNDBIT		_IOW(UINPUT_IOCTL_BASE, 106, int)
#define UI_SET_FFBIT		_IOW(UINPUT_IOCTL_BASE, 107, int)
#define UI_SET_PHYS		_IOW(UINPUT_IOCTL_BASE, 108, char*)
#define UI_SET_SWBIT		_IOW(UINPUT_IOCTL_BASE, 109, int)
#define UI_SET_PROPBIT		_IOW(UINPUT_IOCTL_BASE, 110, int)
*/

const (
	UI_SET_EVBIT  = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((100) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
	UI_SET_KEYBIT = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((101) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
	UI_SET_RELBIT = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((102) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
	UI_SET_ABSBIT = (((_IOC_WRITE) << _IOC_DIRSHIFT) | ((UINPUT_IOCTL_BASE) << _IOC_TYPESHIFT) | ((103) << _IOC_NRSHIFT) | ((C.sizeof_int) << _IOC_SIZESHIFT))
)

type UInput struct {
	fd *os.File
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

func (ui *UInput) KeyEvent(code EventCode, value EventValue) error {

	ev := InputEvent{
		Type:  EV_KEY,
		Code:  uint16(code),
		Value: int32(value),
	}

	ref, _ := ev.PassRef()
	defer ev.Free()

	_, err := ui.fd.Write((*[sizeOfStructInputEventValue]byte)(unsafe.Pointer(ref))[:])

	return err
}

/*
void uinput_abs(int fd, __u16 abs, __s32 val)
{
    struct input_event ev;

    memset(&ev, 0, sizeof(ev));
    ev.type = EV_ABS;
    ev.code = abs;
    ev.value = val;
    write(fd, &ev, sizeof(ev));
}
*/
/*
void uinput_rel(int fd, __u16 rel, __s32 val)
{
    struct input_event ev;

    memset(&ev, 0, sizeof(ev));
    ev.type = EV_REL;
    ev.code = rel;
    ev.value = val;
    write(fd, &ev, sizeof(ev));
}
*/
/*
void uinput_scan(int fd, __s32 val)
{
    struct input_event ev;

    memset(&ev, 0, sizeof(ev));
    ev.type = EV_MSC;
    ev.code = MSC_SCAN;
    ev.value = val;
    write(fd, &ev, sizeof(ev));
}
*/
/*
void uinput_set_delay_period(int fd, __s32 delay, __s32 period)
{
    struct input_event ev;

    memset(&ev, 0, sizeof(ev));
    ev.type = EV_REP;
    ev.code = REP_DELAY;
    ev.value = delay;
    write(fd, &ev, sizeof(ev));
    ev.code = REP_PERIOD;
    ev.value = period;
    write(fd, &ev, sizeof(ev));
}
*/
/*
void uinput_syn(int fd)
{
    struct input_event ev;

    memset(&ev, 0, sizeof(ev));
    ev.type = EV_SYN;
    ev.code = SYN_REPORT;
    ev.value = 0;
    write(fd, &ev, sizeof(ev));
}
*/
func (ui *UInput) SynEvent() error {

	ev := InputEvent{
		Type:  EV_SYN,
		Code:  SYN_REPORT,
		Value: 0,
	}

	ref, _ := ev.PassRef()
	defer ev.Free()

	_, err := ui.fd.Write((*[sizeOfStructInputEventValue]byte)(unsafe.Pointer(ref))[:])

	return err
}
