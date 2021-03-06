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
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

// fakeWriteDevice is used for unit testing purpose
type fakeWriteDevice struct {
	iOCtlErrors           bool
	iOCtlCountBeforeError int
	writeErrors           bool
	writeCountBeforeError int
	writeBList            [][]byte
	iOCtlRequestList      []uintptr
	iOCtlDataList         []interface{}
}

func (wd *fakeWriteDevice) Open() error {
	return nil
}

func (wd *fakeWriteDevice) Close() {
}

func (wd *fakeWriteDevice) Write(b []byte) (int, error) {
	if wd.writeErrors {
		if wd.writeCountBeforeError == 0 {
			return 0, fmt.Errorf("Write: test error")
		}
		wd.writeCountBeforeError--
	}
	wd.writeBList = append(wd.writeBList, b)
	return len(b), nil
}

func (wd *fakeWriteDevice) IOCtl(request uintptr, data interface{}) error {
	if wd.iOCtlErrors {
		if wd.iOCtlCountBeforeError == 0 {
			return fmt.Errorf("IOCtl: test error")
		}
		wd.iOCtlCountBeforeError--
	}
	wd.iOCtlRequestList = append(wd.iOCtlRequestList, request)
	wd.iOCtlDataList = append(wd.iOCtlDataList, data)
	return nil
}

func logWdOperations(t *testing.T, wd *fakeWriteDevice) {
	t.Logf("nb write = %d", len(wd.writeBList))
	t.Logf("nb ioctl = %d", len(wd.iOCtlRequestList))
	for i, b := range wd.writeBList {
		t.Logf("write number %d : \n%s", i, hex.Dump(b))
	}
	for i, r := range wd.iOCtlRequestList {
		var val uint32
		switch wd.iOCtlDataList[i].(type) {
		case int:
			val = uint32(wd.iOCtlDataList[i].(int))
		case uint32:
			val = wd.iOCtlDataList[i].(uint32)
		case uintptr:
			val = uint32(wd.iOCtlDataList[i].(uintptr))
		case EventCode:
			val = uint32(wd.iOCtlDataList[i].(EventCode))
		case EventValue:
			val = uint32(wd.iOCtlDataList[i].(EventValue))
		default:
			val = uint32(0xdeadbeaf)
		}
		t.Logf("ioctl number %d : req=%d(0x%08x), dat=%d(0x%08x)", i, r, r, val, val)
	}
}

func testWdOperations(
	t *testing.T,
	wd *fakeWriteDevice,
	awaitedWRBufList [][]byte,
	awaitedIOReqList []uintptr,
	awaitedIODatList []uint32,
) {
	if len(wd.writeBList) != len(awaitedWRBufList) {
		t.Errorf("Number of writes operations differs from %d awaited: %d ",
			len(awaitedWRBufList),
			len(wd.writeBList))
	} else {
		for i := range wd.writeBList {
			if !reflect.DeepEqual(wd.writeBList[i], awaitedWRBufList[i]) {
				t.Errorf("Write operation %d differs\nWanted:\n%s\nGot:\n%s",
					i,
					hex.Dump(awaitedWRBufList[i]),
					hex.Dump(wd.writeBList[i]))
			}
		}
	}
	if len(wd.iOCtlRequestList) != len(awaitedIOReqList) {
		t.Errorf("Number of IOCtl operations differs from %d awaited: %d ",
			len(awaitedIOReqList),
			len(wd.iOCtlRequestList))
	} else {
		for i, r := range wd.iOCtlRequestList {
			var val uint32
			switch wd.iOCtlDataList[i].(type) {
			case int:
				val = uint32(wd.iOCtlDataList[i].(int))
			case uint32:
				val = wd.iOCtlDataList[i].(uint32)
			case uintptr:
				val = uint32(wd.iOCtlDataList[i].(uintptr))
			case EventCode:
				val = uint32(wd.iOCtlDataList[i].(EventCode))
			case EventValue:
				val = uint32(wd.iOCtlDataList[i].(EventValue))
			default:
				val = uint32(0xdeadbeaf)
			}
			if r != awaitedIOReqList[i] || val != awaitedIODatList[i] {
				t.Errorf("ioctl operation %d differs\nWanted: ioctl(%08x, %08x)\nGot:    ioctl(%08x, %08x))",
					i, r, val, awaitedIOReqList[i], awaitedIODatList[i])
			}
		}
	}
}

func TestUInput_Init(t *testing.T) {
	type args struct {
		wd       WriteDeviceInterface
		name     string
		vendor   uint16
		product  uint16
		version  uint16
		keys     []EventCode
		rels     []EventCode
		axes     []AxisSetup
		keyboard bool
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		awaitedWRBufList [][]byte
		awaitedIOReqList []uintptr
		awaitedIODatList []uint32
	}{
		{
			name: "Nominal",
			args: args{
				wd: &fakeWriteDevice{
					iOCtlErrors:           false,
					iOCtlCountBeforeError: 0,
					writeErrors:           false,
					writeCountBeforeError: 0,
				},
				name:     "My Test Device",
				vendor:   uint16(0x1234),
				product:  uint16(0x5678),
				version:  uint16(0x9abc),
				keys:     []EventCode{BTN_START, BTN_A},
				rels:     []EventCode{KEY_ENTER},
				axes:     []AxisSetup{AxisSetup{Code: ABS_X, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128}},
				keyboard: false,
			},
			wantErr: false,
			awaitedWRBufList: [][]byte{{
				/* 00000: */ 0x4d, 0x79, 0x20, 0x54, 0x65, 0x73, 0x74, 0x20, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x00, 0x00,
				/* 0x010: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x020: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x030: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x040: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x050: */ 0x03, 0x00, 0x34, 0x12, 0x78, 0x56, 0xbc, 0x9a, 0x00, 0x00, 0x00, 0x00, 0xff, 0x7f, 0x00, 0x00,
				/* 0x060: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x070: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x080: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x090: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x0a0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x0b0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x0c0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x0d0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x0e0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x0f0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x100: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x110: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x120: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x130: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x140: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x150: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0xff, 0xff,
				/* 0x160: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x170: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x180: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x190: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x1a0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x1b0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x1c0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x1d0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x1e0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x1f0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x200: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x210: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x220: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x230: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x240: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x250: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00,
				/* 0x260: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x270: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x280: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x290: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x2a0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x2b0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x2c0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x2d0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x2e0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x2f0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x300: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x310: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x320: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x330: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x340: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x350: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
				/* 0x360: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x370: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x380: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x390: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x3a0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x3b0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x3c0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x3d0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x3e0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x3f0: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x400: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x410: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x420: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x430: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x440: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				/* 0x450: */ 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}},
			awaitedIOReqList: []uintptr{0x40045564, 0x40045565, 0x40045565, 0x40045564, 0x40045567, 0x40045564, 0x40045566, 0x00005501},
			awaitedIODatList: []uint32{0x00000001, 0x0000013b, 0x00000130, 0x00000003, 0x00000000, 0x00000002, 0x0000001c, 0x00000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ui UInput
			if err := ui.Init(tt.args.wd, tt.args.name, tt.args.vendor, tt.args.product, tt.args.version, tt.args.keys, tt.args.rels, tt.args.axes, tt.args.keyboard); (err != nil) != tt.wantErr {
				t.Errorf("UInput.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.args.wd.(*fakeWriteDevice))
			testWdOperations(t,
				tt.args.wd.(*fakeWriteDevice),
				tt.awaitedWRBufList,
				tt.awaitedIOReqList,
				tt.awaitedIODatList,
			)
		})
	}
}

func TestUInput_setupKeys(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		keys []EventCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.setupKeys(tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("UInput.setupKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_setupRels(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		rels []EventCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.setupRels(tt.args.rels); (err != nil) != tt.wantErr {
				t.Errorf("UInput.setupRels() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_setupAxes(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		axes  []AxisSetup
		uidev *UInputUserDev
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.setupAxes(tt.args.axes, tt.args.uidev); (err != nil) != tt.wantErr {
				t.Errorf("UInput.setupAxes() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_setupKeyboard(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.setupKeyboard(); (err != nil) != tt.wantErr {
				t.Errorf("UInput.setupKeyboard() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_genEvent(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		eventType uint16
		code      EventCode
		value     EventValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.genEvent(tt.args.eventType, tt.args.code, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UInput.genEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_KeyEvent(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		code  EventCode
		value EventValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.KeyEvent(tt.args.code, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UInput.KeyEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_AbsEvent(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		code  EventCode
		value EventValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.AbsEvent(tt.args.code, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UInput.AbsEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_RelEvent(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		code  EventCode
		value EventValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.RelEvent(tt.args.code, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UInput.RelEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_ScanEvent(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		value EventValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.ScanEvent(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UInput.ScanEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_SynEvent(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.SynEvent(); (err != nil) != tt.wantErr {
				t.Errorf("UInput.SynEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_SetDelayPeriod(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		delay  EventValue
		period EventValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.SetDelayPeriod(tt.args.delay, tt.args.period); (err != nil) != tt.wantErr {
				t.Errorf("UInput.SetDelayPeriod() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_KeyPressed(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		codes []EventCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.KeyPressed(tt.args.codes); (err != nil) != tt.wantErr {
				t.Errorf("UInput.KeyPressed() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}

func TestUInput_KeyReleased(t *testing.T) {
	type fields struct {
		wd            WriteDeviceInterface
		keyPressedMap map[EventCode]bool
	}
	type args struct {
		codes []EventCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UInput{
				wd:            tt.fields.wd,
				keyPressedMap: tt.fields.keyPressedMap,
			}
			if err := ui.KeyReleased(tt.args.codes); (err != nil) != tt.wantErr {
				t.Errorf("UInput.KeyReleased() error = %v, wantErr %v", err, tt.wantErr)
			}
			logWdOperations(t, tt.fields.wd.(*fakeWriteDevice))
		})
	}
}
