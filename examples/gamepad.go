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

package main

import (
	"fmt"
	"log"
	"time"

	uinput "github.com/ynsta/uinput"
)

func main() {

	var ui uinput.UInput

	if err := ui.Open(); err != nil {
		log.Fatal(err)
	}
	if err := ui.Init("Microsoft X-Box 360 pad",
		0x045e,
		0x028e,
		0x110,
		[]uinput.EventCode{uinput.BTN_START,
			uinput.BTN_MODE,
			uinput.BTN_SELECT,
			uinput.BTN_A,
			uinput.BTN_B,
			uinput.BTN_X,
			uinput.BTN_Y,
			uinput.BTN_TL,
			uinput.BTN_TR,
			uinput.BTN_THUMBL,
			uinput.BTN_THUMBR,
		},
		[]uinput.EventCode{},
		[]uinput.AxisSetup{
			uinput.AxisSetup{Code: uinput.ABS_X, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128},
			uinput.AxisSetup{Code: uinput.ABS_Y, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128},
			uinput.AxisSetup{Code: uinput.ABS_RX, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128},
			uinput.AxisSetup{Code: uinput.ABS_RY, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128},
			uinput.AxisSetup{Code: uinput.ABS_Z, Min: 0, Max: 255, Fuzz: 0, Flat: 0},
			uinput.AxisSetup{Code: uinput.ABS_RZ, Min: 0, Max: 255, Fuzz: 0, Flat: 0},
			uinput.AxisSetup{Code: uinput.ABS_HAT0X, Min: -1, Max: 1, Fuzz: 0, Flat: 0},
			uinput.AxisSetup{Code: uinput.ABS_HAT0Y, Min: -1, Max: 1, Fuzz: 0, Flat: 0},
		},
		false,
	); err != nil {
		log.Fatal(err)
	}

	quit := make(chan int)

	go func() {
		toggle := uinput.EventValue(1)

		for {
			select {
			case <-quit:
				fmt.Println("quit")
				return
			default:
				ui.KeyEvent(uinput.BTN_A, toggle)
				if toggle == 1 {
					toggle = 0
				} else {
					toggle = 1
				}
				ui.SynEvent()
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()

	fmt.Println("Press enter to exit")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)
	quit <- 0

	ui.Close()
	return
}
