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

const (
	// UINPUT_VERSION as defined in linux/uinput.h:45
	UINPUT_VERSION = 5
	// UINPUT_MAX_NAME_SIZE as defined in linux/uinput.h:46
	UINPUT_MAX_NAME_SIZE = 80
	// UINPUT_IOCTL_BASE as defined in linux/uinput.h:62
	UINPUT_IOCTL_BASE = 85
	// UI_DEV_CREATE as defined in linux/uinput.h:63
	UI_DEV_CREATE = 21761
	// UI_DEV_DESTROY as defined in linux/uinput.h:64
	UI_DEV_DESTROY = 21762
	// EV_UINPUT as defined in linux/uinput.h:218
	EV_UINPUT = 257
	// UI_FF_UPLOAD as defined in linux/uinput.h:219
	UI_FF_UPLOAD = 1
	// UI_FF_ERASE as defined in linux/uinput.h:220
	UI_FF_ERASE = 2
	// EV_VERSION as defined in linux/input.h:34
	EV_VERSION = 65537
	// BUS_PCI as defined in linux/input.h:227
	BUS_PCI = 1
	// BUS_ISAPNP as defined in linux/input.h:228
	BUS_ISAPNP = 2
	// BUS_USB as defined in linux/input.h:229
	BUS_USB = 3
	// BUS_HIL as defined in linux/input.h:230
	BUS_HIL = 4
	// BUS_BLUETOOTH as defined in linux/input.h:231
	BUS_BLUETOOTH = 5
	// BUS_VIRTUAL as defined in linux/input.h:232
	BUS_VIRTUAL = 6
	// BUS_ISA as defined in linux/input.h:234
	BUS_ISA = 16
	// BUS_I8042 as defined in linux/input.h:235
	BUS_I8042 = 17
	// BUS_XTKBD as defined in linux/input.h:236
	BUS_XTKBD = 18
	// BUS_RS232 as defined in linux/input.h:237
	BUS_RS232 = 19
	// BUS_GAMEPORT as defined in linux/input.h:238
	BUS_GAMEPORT = 20
	// BUS_PARPORT as defined in linux/input.h:239
	BUS_PARPORT = 21
	// BUS_AMIGA as defined in linux/input.h:240
	BUS_AMIGA = 22
	// BUS_ADB as defined in linux/input.h:241
	BUS_ADB = 23
	// BUS_I2C as defined in linux/input.h:242
	BUS_I2C = 24
	// BUS_HOST as defined in linux/input.h:243
	BUS_HOST = 25
	// BUS_GSC as defined in linux/input.h:244
	BUS_GSC = 26
	// BUS_ATARI as defined in linux/input.h:245
	BUS_ATARI = 27
	// BUS_SPI as defined in linux/input.h:246
	BUS_SPI = 28
	// BUS_RMI as defined in linux/input.h:247
	BUS_RMI = 29
	// BUS_CEC as defined in linux/input.h:248
	BUS_CEC = 30
	// BUS_INTEL_ISHTP as defined in linux/input.h:249
	BUS_INTEL_ISHTP = 31
	// _IOC_NRBITS as defined in asm-generic/ioctl.h:22
	_IOC_NRBITS = 8
	// _IOC_TYPEBITS as defined in asm-generic/ioctl.h:23
	_IOC_TYPEBITS = 8
	// _IOC_SIZEBITS as defined in asm-generic/ioctl.h:31
	_IOC_SIZEBITS = 14
	// _IOC_DIRBITS as defined in asm-generic/ioctl.h:35
	_IOC_DIRBITS = 2
	// _IOC_NRMASK as defined in asm-generic/ioctl.h:38
	_IOC_NRMASK = 255
	// _IOC_TYPEMASK as defined in asm-generic/ioctl.h:39
	_IOC_TYPEMASK = 255
	// _IOC_SIZEMASK as defined in asm-generic/ioctl.h:40
	_IOC_SIZEMASK = 16383
	// _IOC_DIRMASK as defined in asm-generic/ioctl.h:41
	_IOC_DIRMASK = 3
	// _IOC_NRSHIFT as defined in asm-generic/ioctl.h:43
	_IOC_NRSHIFT = 0
	// _IOC_TYPESHIFT as defined in asm-generic/ioctl.h:44
	_IOC_TYPESHIFT = 8
	// _IOC_SIZESHIFT as defined in asm-generic/ioctl.h:45
	_IOC_SIZESHIFT = 16
	// _IOC_DIRSHIFT as defined in asm-generic/ioctl.h:46
	_IOC_DIRSHIFT = 30
	// _IOC_NONE as defined in asm-generic/ioctl.h:54
	_IOC_NONE = 0
	// _IOC_WRITE as defined in asm-generic/ioctl.h:58
	_IOC_WRITE = 1
	// _IOC_READ as defined in asm-generic/ioctl.h:62
	_IOC_READ = 2
	// EV_SYN as defined in linux/input-event-codes.h:37
	EV_SYN = 0
	// EV_KEY as defined in linux/input-event-codes.h:38
	EV_KEY = 1
	// EV_REL as defined in linux/input-event-codes.h:39
	EV_REL = 2
	// EV_ABS as defined in linux/input-event-codes.h:40
	EV_ABS = 3
	// EV_MSC as defined in linux/input-event-codes.h:41
	EV_MSC = 4
	// EV_SW as defined in linux/input-event-codes.h:42
	EV_SW = 5
	// EV_LED as defined in linux/input-event-codes.h:43
	EV_LED = 17
	// EV_SND as defined in linux/input-event-codes.h:44
	EV_SND = 18
	// EV_REP as defined in linux/input-event-codes.h:45
	EV_REP = 20
	// EV_FF as defined in linux/input-event-codes.h:46
	EV_FF = 21
	// EV_PWR as defined in linux/input-event-codes.h:47
	EV_PWR = 22
	// EV_FF_STATUS as defined in linux/input-event-codes.h:48
	EV_FF_STATUS = 23
	// EV_MAX as defined in linux/input-event-codes.h:49
	EV_MAX = 31
	// EV_CNT as defined in linux/input-event-codes.h:50
	EV_CNT = 32
	// SYN_REPORT as defined in linux/input-event-codes.h:56
	SYN_REPORT = 0
	// SYN_CONFIG as defined in linux/input-event-codes.h:57
	SYN_CONFIG = 1
	// SYN_MT_REPORT as defined in linux/input-event-codes.h:58
	SYN_MT_REPORT = 2
	// SYN_DROPPED as defined in linux/input-event-codes.h:59
	SYN_DROPPED = 3
	// SYN_MAX as defined in linux/input-event-codes.h:60
	SYN_MAX = 15
	// SYN_CNT as defined in linux/input-event-codes.h:61
	SYN_CNT = 16
	// KEY_RESERVED as defined in linux/input-event-codes.h:74
	KEY_RESERVED = 0
	// KEY_ESC as defined in linux/input-event-codes.h:75
	KEY_ESC = 1
	// KEY_1 as defined in linux/input-event-codes.h:76
	KEY_1 = 2
	// KEY_2 as defined in linux/input-event-codes.h:77
	KEY_2 = 3
	// KEY_3 as defined in linux/input-event-codes.h:78
	KEY_3 = 4
	// KEY_4 as defined in linux/input-event-codes.h:79
	KEY_4 = 5
	// KEY_5 as defined in linux/input-event-codes.h:80
	KEY_5 = 6
	// KEY_6 as defined in linux/input-event-codes.h:81
	KEY_6 = 7
	// KEY_7 as defined in linux/input-event-codes.h:82
	KEY_7 = 8
	// KEY_8 as defined in linux/input-event-codes.h:83
	KEY_8 = 9
	// KEY_9 as defined in linux/input-event-codes.h:84
	KEY_9 = 10
	// KEY_0 as defined in linux/input-event-codes.h:85
	KEY_0 = 11
	// KEY_MINUS as defined in linux/input-event-codes.h:86
	KEY_MINUS = 12
	// KEY_EQUAL as defined in linux/input-event-codes.h:87
	KEY_EQUAL = 13
	// KEY_BACKSPACE as defined in linux/input-event-codes.h:88
	KEY_BACKSPACE = 14
	// KEY_TAB as defined in linux/input-event-codes.h:89
	KEY_TAB = 15
	// KEY_Q as defined in linux/input-event-codes.h:90
	KEY_Q = 16
	// KEY_W as defined in linux/input-event-codes.h:91
	KEY_W = 17
	// KEY_E as defined in linux/input-event-codes.h:92
	KEY_E = 18
	// KEY_R as defined in linux/input-event-codes.h:93
	KEY_R = 19
	// KEY_T as defined in linux/input-event-codes.h:94
	KEY_T = 20
	// KEY_Y as defined in linux/input-event-codes.h:95
	KEY_Y = 21
	// KEY_U as defined in linux/input-event-codes.h:96
	KEY_U = 22
	// KEY_I as defined in linux/input-event-codes.h:97
	KEY_I = 23
	// KEY_O as defined in linux/input-event-codes.h:98
	KEY_O = 24
	// KEY_P as defined in linux/input-event-codes.h:99
	KEY_P = 25
	// KEY_LEFTBRACE as defined in linux/input-event-codes.h:100
	KEY_LEFTBRACE = 26
	// KEY_RIGHTBRACE as defined in linux/input-event-codes.h:101
	KEY_RIGHTBRACE = 27
	// KEY_ENTER as defined in linux/input-event-codes.h:102
	KEY_ENTER = 28
	// KEY_LEFTCTRL as defined in linux/input-event-codes.h:103
	KEY_LEFTCTRL = 29
	// KEY_A as defined in linux/input-event-codes.h:104
	KEY_A = 30
	// KEY_S as defined in linux/input-event-codes.h:105
	KEY_S = 31
	// KEY_D as defined in linux/input-event-codes.h:106
	KEY_D = 32
	// KEY_F as defined in linux/input-event-codes.h:107
	KEY_F = 33
	// KEY_G as defined in linux/input-event-codes.h:108
	KEY_G = 34
	// KEY_H as defined in linux/input-event-codes.h:109
	KEY_H = 35
	// KEY_J as defined in linux/input-event-codes.h:110
	KEY_J = 36
	// KEY_K as defined in linux/input-event-codes.h:111
	KEY_K = 37
	// KEY_L as defined in linux/input-event-codes.h:112
	KEY_L = 38
	// KEY_SEMICOLON as defined in linux/input-event-codes.h:113
	KEY_SEMICOLON = 39
	// KEY_APOSTROPHE as defined in linux/input-event-codes.h:114
	KEY_APOSTROPHE = 40
	// KEY_GRAVE as defined in linux/input-event-codes.h:115
	KEY_GRAVE = 41
	// KEY_LEFTSHIFT as defined in linux/input-event-codes.h:116
	KEY_LEFTSHIFT = 42
	// KEY_BACKSLASH as defined in linux/input-event-codes.h:117
	KEY_BACKSLASH = 43
	// KEY_Z as defined in linux/input-event-codes.h:118
	KEY_Z = 44
	// KEY_X as defined in linux/input-event-codes.h:119
	KEY_X = 45
	// KEY_C as defined in linux/input-event-codes.h:120
	KEY_C = 46
	// KEY_V as defined in linux/input-event-codes.h:121
	KEY_V = 47
	// KEY_B as defined in linux/input-event-codes.h:122
	KEY_B = 48
	// KEY_N as defined in linux/input-event-codes.h:123
	KEY_N = 49
	// KEY_M as defined in linux/input-event-codes.h:124
	KEY_M = 50
	// KEY_COMMA as defined in linux/input-event-codes.h:125
	KEY_COMMA = 51
	// KEY_DOT as defined in linux/input-event-codes.h:126
	KEY_DOT = 52
	// KEY_SLASH as defined in linux/input-event-codes.h:127
	KEY_SLASH = 53
	// KEY_RIGHTSHIFT as defined in linux/input-event-codes.h:128
	KEY_RIGHTSHIFT = 54
	// KEY_KPASTERISK as defined in linux/input-event-codes.h:129
	KEY_KPASTERISK = 55
	// KEY_LEFTALT as defined in linux/input-event-codes.h:130
	KEY_LEFTALT = 56
	// KEY_SPACE as defined in linux/input-event-codes.h:131
	KEY_SPACE = 57
	// KEY_CAPSLOCK as defined in linux/input-event-codes.h:132
	KEY_CAPSLOCK = 58
	// KEY_F1 as defined in linux/input-event-codes.h:133
	KEY_F1 = 59
	// KEY_F2 as defined in linux/input-event-codes.h:134
	KEY_F2 = 60
	// KEY_F3 as defined in linux/input-event-codes.h:135
	KEY_F3 = 61
	// KEY_F4 as defined in linux/input-event-codes.h:136
	KEY_F4 = 62
	// KEY_F5 as defined in linux/input-event-codes.h:137
	KEY_F5 = 63
	// KEY_F6 as defined in linux/input-event-codes.h:138
	KEY_F6 = 64
	// KEY_F7 as defined in linux/input-event-codes.h:139
	KEY_F7 = 65
	// KEY_F8 as defined in linux/input-event-codes.h:140
	KEY_F8 = 66
	// KEY_F9 as defined in linux/input-event-codes.h:141
	KEY_F9 = 67
	// KEY_F10 as defined in linux/input-event-codes.h:142
	KEY_F10 = 68
	// KEY_NUMLOCK as defined in linux/input-event-codes.h:143
	KEY_NUMLOCK = 69
	// KEY_SCROLLLOCK as defined in linux/input-event-codes.h:144
	KEY_SCROLLLOCK = 70
	// KEY_KP7 as defined in linux/input-event-codes.h:145
	KEY_KP7 = 71
	// KEY_KP8 as defined in linux/input-event-codes.h:146
	KEY_KP8 = 72
	// KEY_KP9 as defined in linux/input-event-codes.h:147
	KEY_KP9 = 73
	// KEY_KPMINUS as defined in linux/input-event-codes.h:148
	KEY_KPMINUS = 74
	// KEY_KP4 as defined in linux/input-event-codes.h:149
	KEY_KP4 = 75
	// KEY_KP5 as defined in linux/input-event-codes.h:150
	KEY_KP5 = 76
	// KEY_KP6 as defined in linux/input-event-codes.h:151
	KEY_KP6 = 77
	// KEY_KPPLUS as defined in linux/input-event-codes.h:152
	KEY_KPPLUS = 78
	// KEY_KP1 as defined in linux/input-event-codes.h:153
	KEY_KP1 = 79
	// KEY_KP2 as defined in linux/input-event-codes.h:154
	KEY_KP2 = 80
	// KEY_KP3 as defined in linux/input-event-codes.h:155
	KEY_KP3 = 81
	// KEY_KP0 as defined in linux/input-event-codes.h:156
	KEY_KP0 = 82
	// KEY_KPDOT as defined in linux/input-event-codes.h:157
	KEY_KPDOT = 83
	// KEY_ZENKAKUHANKAKU as defined in linux/input-event-codes.h:159
	KEY_ZENKAKUHANKAKU = 85
	// KEY_102ND as defined in linux/input-event-codes.h:160
	KEY_102ND = 86
	// KEY_F11 as defined in linux/input-event-codes.h:161
	KEY_F11 = 87
	// KEY_F12 as defined in linux/input-event-codes.h:162
	KEY_F12 = 88
	// KEY_RO as defined in linux/input-event-codes.h:163
	KEY_RO = 89
	// KEY_KATAKANA as defined in linux/input-event-codes.h:164
	KEY_KATAKANA = 90
	// KEY_HIRAGANA as defined in linux/input-event-codes.h:165
	KEY_HIRAGANA = 91
	// KEY_HENKAN as defined in linux/input-event-codes.h:166
	KEY_HENKAN = 92
	// KEY_KATAKANAHIRAGANA as defined in linux/input-event-codes.h:167
	KEY_KATAKANAHIRAGANA = 93
	// KEY_MUHENKAN as defined in linux/input-event-codes.h:168
	KEY_MUHENKAN = 94
	// KEY_KPJPCOMMA as defined in linux/input-event-codes.h:169
	KEY_KPJPCOMMA = 95
	// KEY_KPENTER as defined in linux/input-event-codes.h:170
	KEY_KPENTER = 96
	// KEY_RIGHTCTRL as defined in linux/input-event-codes.h:171
	KEY_RIGHTCTRL = 97
	// KEY_KPSLASH as defined in linux/input-event-codes.h:172
	KEY_KPSLASH = 98
	// KEY_SYSRQ as defined in linux/input-event-codes.h:173
	KEY_SYSRQ = 99
	// KEY_RIGHTALT as defined in linux/input-event-codes.h:174
	KEY_RIGHTALT = 100
	// KEY_LINEFEED as defined in linux/input-event-codes.h:175
	KEY_LINEFEED = 101
	// KEY_HOME as defined in linux/input-event-codes.h:176
	KEY_HOME = 102
	// KEY_UP as defined in linux/input-event-codes.h:177
	KEY_UP = 103
	// KEY_PAGEUP as defined in linux/input-event-codes.h:178
	KEY_PAGEUP = 104
	// KEY_LEFT as defined in linux/input-event-codes.h:179
	KEY_LEFT = 105
	// KEY_RIGHT as defined in linux/input-event-codes.h:180
	KEY_RIGHT = 106
	// KEY_END as defined in linux/input-event-codes.h:181
	KEY_END = 107
	// KEY_DOWN as defined in linux/input-event-codes.h:182
	KEY_DOWN = 108
	// KEY_PAGEDOWN as defined in linux/input-event-codes.h:183
	KEY_PAGEDOWN = 109
	// KEY_INSERT as defined in linux/input-event-codes.h:184
	KEY_INSERT = 110
	// KEY_DELETE as defined in linux/input-event-codes.h:185
	KEY_DELETE = 111
	// KEY_MACRO as defined in linux/input-event-codes.h:186
	KEY_MACRO = 112
	// KEY_MUTE as defined in linux/input-event-codes.h:187
	KEY_MUTE = 113
	// KEY_VOLUMEDOWN as defined in linux/input-event-codes.h:188
	KEY_VOLUMEDOWN = 114
	// KEY_VOLUMEUP as defined in linux/input-event-codes.h:189
	KEY_VOLUMEUP = 115
	// KEY_POWER as defined in linux/input-event-codes.h:190
	KEY_POWER = 116
	// KEY_KPEQUAL as defined in linux/input-event-codes.h:191
	KEY_KPEQUAL = 117
	// KEY_KPPLUSMINUS as defined in linux/input-event-codes.h:192
	KEY_KPPLUSMINUS = 118
	// KEY_PAUSE as defined in linux/input-event-codes.h:193
	KEY_PAUSE = 119
	// KEY_SCALE as defined in linux/input-event-codes.h:194
	KEY_SCALE = 120
	// KEY_KPCOMMA as defined in linux/input-event-codes.h:196
	KEY_KPCOMMA = 121
	// KEY_HANGEUL as defined in linux/input-event-codes.h:197
	KEY_HANGEUL = 122
	// KEY_HANGUEL as defined in linux/input-event-codes.h:198
	KEY_HANGUEL = 122
	// KEY_HANJA as defined in linux/input-event-codes.h:199
	KEY_HANJA = 123
	// KEY_YEN as defined in linux/input-event-codes.h:200
	KEY_YEN = 124
	// KEY_LEFTMETA as defined in linux/input-event-codes.h:201
	KEY_LEFTMETA = 125
	// KEY_RIGHTMETA as defined in linux/input-event-codes.h:202
	KEY_RIGHTMETA = 126
	// KEY_COMPOSE as defined in linux/input-event-codes.h:203
	KEY_COMPOSE = 127
	// KEY_STOP as defined in linux/input-event-codes.h:205
	KEY_STOP = 128
	// KEY_AGAIN as defined in linux/input-event-codes.h:206
	KEY_AGAIN = 129
	// KEY_PROPS as defined in linux/input-event-codes.h:207
	KEY_PROPS = 130
	// KEY_UNDO as defined in linux/input-event-codes.h:208
	KEY_UNDO = 131
	// KEY_FRONT as defined in linux/input-event-codes.h:209
	KEY_FRONT = 132
	// KEY_COPY as defined in linux/input-event-codes.h:210
	KEY_COPY = 133
	// KEY_OPEN as defined in linux/input-event-codes.h:211
	KEY_OPEN = 134
	// KEY_PASTE as defined in linux/input-event-codes.h:212
	KEY_PASTE = 135
	// KEY_FIND as defined in linux/input-event-codes.h:213
	KEY_FIND = 136
	// KEY_CUT as defined in linux/input-event-codes.h:214
	KEY_CUT = 137
	// KEY_HELP as defined in linux/input-event-codes.h:215
	KEY_HELP = 138
	// KEY_MENU as defined in linux/input-event-codes.h:216
	KEY_MENU = 139
	// KEY_CALC as defined in linux/input-event-codes.h:217
	KEY_CALC = 140
	// KEY_SETUP as defined in linux/input-event-codes.h:218
	KEY_SETUP = 141
	// KEY_SLEEP as defined in linux/input-event-codes.h:219
	KEY_SLEEP = 142
	// KEY_WAKEUP as defined in linux/input-event-codes.h:220
	KEY_WAKEUP = 143
	// KEY_FILE as defined in linux/input-event-codes.h:221
	KEY_FILE = 144
	// KEY_SENDFILE as defined in linux/input-event-codes.h:222
	KEY_SENDFILE = 145
	// KEY_DELETEFILE as defined in linux/input-event-codes.h:223
	KEY_DELETEFILE = 146
	// KEY_XFER as defined in linux/input-event-codes.h:224
	KEY_XFER = 147
	// KEY_PROG1 as defined in linux/input-event-codes.h:225
	KEY_PROG1 = 148
	// KEY_PROG2 as defined in linux/input-event-codes.h:226
	KEY_PROG2 = 149
	// KEY_WWW as defined in linux/input-event-codes.h:227
	KEY_WWW = 150
	// KEY_MSDOS as defined in linux/input-event-codes.h:228
	KEY_MSDOS = 151
	// KEY_COFFEE as defined in linux/input-event-codes.h:229
	KEY_COFFEE = 152
	// KEY_SCREENLOCK as defined in linux/input-event-codes.h:230
	KEY_SCREENLOCK = 152
	// KEY_ROTATE_DISPLAY as defined in linux/input-event-codes.h:231
	KEY_ROTATE_DISPLAY = 153
	// KEY_DIRECTION as defined in linux/input-event-codes.h:232
	KEY_DIRECTION = 153
	// KEY_CYCLEWINDOWS as defined in linux/input-event-codes.h:233
	KEY_CYCLEWINDOWS = 154
	// KEY_MAIL as defined in linux/input-event-codes.h:234
	KEY_MAIL = 155
	// KEY_BOOKMARKS as defined in linux/input-event-codes.h:235
	KEY_BOOKMARKS = 156
	// KEY_COMPUTER as defined in linux/input-event-codes.h:236
	KEY_COMPUTER = 157
	// KEY_BACK as defined in linux/input-event-codes.h:237
	KEY_BACK = 158
	// KEY_FORWARD as defined in linux/input-event-codes.h:238
	KEY_FORWARD = 159
	// KEY_CLOSECD as defined in linux/input-event-codes.h:239
	KEY_CLOSECD = 160
	// KEY_EJECTCD as defined in linux/input-event-codes.h:240
	KEY_EJECTCD = 161
	// KEY_EJECTCLOSECD as defined in linux/input-event-codes.h:241
	KEY_EJECTCLOSECD = 162
	// KEY_NEXTSONG as defined in linux/input-event-codes.h:242
	KEY_NEXTSONG = 163
	// KEY_PLAYPAUSE as defined in linux/input-event-codes.h:243
	KEY_PLAYPAUSE = 164
	// KEY_PREVIOUSSONG as defined in linux/input-event-codes.h:244
	KEY_PREVIOUSSONG = 165
	// KEY_STOPCD as defined in linux/input-event-codes.h:245
	KEY_STOPCD = 166
	// KEY_RECORD as defined in linux/input-event-codes.h:246
	KEY_RECORD = 167
	// KEY_REWIND as defined in linux/input-event-codes.h:247
	KEY_REWIND = 168
	// KEY_PHONE as defined in linux/input-event-codes.h:248
	KEY_PHONE = 169
	// KEY_ISO as defined in linux/input-event-codes.h:249
	KEY_ISO = 170
	// KEY_CONFIG as defined in linux/input-event-codes.h:250
	KEY_CONFIG = 171
	// KEY_HOMEPAGE as defined in linux/input-event-codes.h:251
	KEY_HOMEPAGE = 172
	// KEY_REFRESH as defined in linux/input-event-codes.h:252
	KEY_REFRESH = 173
	// KEY_EXIT as defined in linux/input-event-codes.h:253
	KEY_EXIT = 174
	// KEY_MOVE as defined in linux/input-event-codes.h:254
	KEY_MOVE = 175
	// KEY_EDIT as defined in linux/input-event-codes.h:255
	KEY_EDIT = 176
	// KEY_SCROLLUP as defined in linux/input-event-codes.h:256
	KEY_SCROLLUP = 177
	// KEY_SCROLLDOWN as defined in linux/input-event-codes.h:257
	KEY_SCROLLDOWN = 178
	// KEY_KPLEFTPAREN as defined in linux/input-event-codes.h:258
	KEY_KPLEFTPAREN = 179
	// KEY_KPRIGHTPAREN as defined in linux/input-event-codes.h:259
	KEY_KPRIGHTPAREN = 180
	// KEY_NEW as defined in linux/input-event-codes.h:260
	KEY_NEW = 181
	// KEY_REDO as defined in linux/input-event-codes.h:261
	KEY_REDO = 182
	// KEY_F13 as defined in linux/input-event-codes.h:263
	KEY_F13 = 183
	// KEY_F14 as defined in linux/input-event-codes.h:264
	KEY_F14 = 184
	// KEY_F15 as defined in linux/input-event-codes.h:265
	KEY_F15 = 185
	// KEY_F16 as defined in linux/input-event-codes.h:266
	KEY_F16 = 186
	// KEY_F17 as defined in linux/input-event-codes.h:267
	KEY_F17 = 187
	// KEY_F18 as defined in linux/input-event-codes.h:268
	KEY_F18 = 188
	// KEY_F19 as defined in linux/input-event-codes.h:269
	KEY_F19 = 189
	// KEY_F20 as defined in linux/input-event-codes.h:270
	KEY_F20 = 190
	// KEY_F21 as defined in linux/input-event-codes.h:271
	KEY_F21 = 191
	// KEY_F22 as defined in linux/input-event-codes.h:272
	KEY_F22 = 192
	// KEY_F23 as defined in linux/input-event-codes.h:273
	KEY_F23 = 193
	// KEY_F24 as defined in linux/input-event-codes.h:274
	KEY_F24 = 194
	// KEY_PLAYCD as defined in linux/input-event-codes.h:276
	KEY_PLAYCD = 200
	// KEY_PAUSECD as defined in linux/input-event-codes.h:277
	KEY_PAUSECD = 201
	// KEY_PROG3 as defined in linux/input-event-codes.h:278
	KEY_PROG3 = 202
	// KEY_PROG4 as defined in linux/input-event-codes.h:279
	KEY_PROG4 = 203
	// KEY_DASHBOARD as defined in linux/input-event-codes.h:280
	KEY_DASHBOARD = 204
	// KEY_SUSPEND as defined in linux/input-event-codes.h:281
	KEY_SUSPEND = 205
	// KEY_CLOSE as defined in linux/input-event-codes.h:282
	KEY_CLOSE = 206
	// KEY_PLAY as defined in linux/input-event-codes.h:283
	KEY_PLAY = 207
	// KEY_FASTFORWARD as defined in linux/input-event-codes.h:284
	KEY_FASTFORWARD = 208
	// KEY_BASSBOOST as defined in linux/input-event-codes.h:285
	KEY_BASSBOOST = 209
	// KEY_PRINT as defined in linux/input-event-codes.h:286
	KEY_PRINT = 210
	// KEY_HP as defined in linux/input-event-codes.h:287
	KEY_HP = 211
	// KEY_CAMERA as defined in linux/input-event-codes.h:288
	KEY_CAMERA = 212
	// KEY_SOUND as defined in linux/input-event-codes.h:289
	KEY_SOUND = 213
	// KEY_QUESTION as defined in linux/input-event-codes.h:290
	KEY_QUESTION = 214
	// KEY_EMAIL as defined in linux/input-event-codes.h:291
	KEY_EMAIL = 215
	// KEY_CHAT as defined in linux/input-event-codes.h:292
	KEY_CHAT = 216
	// KEY_SEARCH as defined in linux/input-event-codes.h:293
	KEY_SEARCH = 217
	// KEY_CONNECT as defined in linux/input-event-codes.h:294
	KEY_CONNECT = 218
	// KEY_FINANCE as defined in linux/input-event-codes.h:295
	KEY_FINANCE = 219
	// KEY_SPORT as defined in linux/input-event-codes.h:296
	KEY_SPORT = 220
	// KEY_SHOP as defined in linux/input-event-codes.h:297
	KEY_SHOP = 221
	// KEY_ALTERASE as defined in linux/input-event-codes.h:298
	KEY_ALTERASE = 222
	// KEY_CANCEL as defined in linux/input-event-codes.h:299
	KEY_CANCEL = 223
	// KEY_BRIGHTNESSDOWN as defined in linux/input-event-codes.h:300
	KEY_BRIGHTNESSDOWN = 224
	// KEY_BRIGHTNESSUP as defined in linux/input-event-codes.h:301
	KEY_BRIGHTNESSUP = 225
	// KEY_MEDIA as defined in linux/input-event-codes.h:302
	KEY_MEDIA = 226
	// KEY_SWITCHVIDEOMODE as defined in linux/input-event-codes.h:304
	KEY_SWITCHVIDEOMODE = 227
	// KEY_KBDILLUMTOGGLE as defined in linux/input-event-codes.h:306
	KEY_KBDILLUMTOGGLE = 228
	// KEY_KBDILLUMDOWN as defined in linux/input-event-codes.h:307
	KEY_KBDILLUMDOWN = 229
	// KEY_KBDILLUMUP as defined in linux/input-event-codes.h:308
	KEY_KBDILLUMUP = 230
	// KEY_SEND as defined in linux/input-event-codes.h:310
	KEY_SEND = 231
	// KEY_REPLY as defined in linux/input-event-codes.h:311
	KEY_REPLY = 232
	// KEY_FORWARDMAIL as defined in linux/input-event-codes.h:312
	KEY_FORWARDMAIL = 233
	// KEY_SAVE as defined in linux/input-event-codes.h:313
	KEY_SAVE = 234
	// KEY_DOCUMENTS as defined in linux/input-event-codes.h:314
	KEY_DOCUMENTS = 235
	// KEY_BATTERY as defined in linux/input-event-codes.h:316
	KEY_BATTERY = 236
	// KEY_BLUETOOTH as defined in linux/input-event-codes.h:318
	KEY_BLUETOOTH = 237
	// KEY_WLAN as defined in linux/input-event-codes.h:319
	KEY_WLAN = 238
	// KEY_UWB as defined in linux/input-event-codes.h:320
	KEY_UWB = 239
	// KEY_UNKNOWN as defined in linux/input-event-codes.h:322
	KEY_UNKNOWN = 240
	// KEY_VIDEO_NEXT as defined in linux/input-event-codes.h:324
	KEY_VIDEO_NEXT = 241
	// KEY_VIDEO_PREV as defined in linux/input-event-codes.h:325
	KEY_VIDEO_PREV = 242
	// KEY_BRIGHTNESS_CYCLE as defined in linux/input-event-codes.h:326
	KEY_BRIGHTNESS_CYCLE = 243
	// KEY_BRIGHTNESS_AUTO as defined in linux/input-event-codes.h:327
	KEY_BRIGHTNESS_AUTO = 244
	// KEY_BRIGHTNESS_ZERO as defined in linux/input-event-codes.h:330
	KEY_BRIGHTNESS_ZERO = 244
	// KEY_DISPLAY_OFF as defined in linux/input-event-codes.h:331
	KEY_DISPLAY_OFF = 245
	// KEY_WWAN as defined in linux/input-event-codes.h:333
	KEY_WWAN = 246
	// KEY_WIMAX as defined in linux/input-event-codes.h:334
	KEY_WIMAX = 246
	// KEY_RFKILL as defined in linux/input-event-codes.h:335
	KEY_RFKILL = 247
	// KEY_MICMUTE as defined in linux/input-event-codes.h:337
	KEY_MICMUTE = 248
	// BTN_MISC as defined in linux/input-event-codes.h:341
	BTN_MISC = 256
	// BTN_0 as defined in linux/input-event-codes.h:342
	BTN_0 = 256
	// BTN_1 as defined in linux/input-event-codes.h:343
	BTN_1 = 257
	// BTN_2 as defined in linux/input-event-codes.h:344
	BTN_2 = 258
	// BTN_3 as defined in linux/input-event-codes.h:345
	BTN_3 = 259
	// BTN_4 as defined in linux/input-event-codes.h:346
	BTN_4 = 260
	// BTN_5 as defined in linux/input-event-codes.h:347
	BTN_5 = 261
	// BTN_6 as defined in linux/input-event-codes.h:348
	BTN_6 = 262
	// BTN_7 as defined in linux/input-event-codes.h:349
	BTN_7 = 263
	// BTN_8 as defined in linux/input-event-codes.h:350
	BTN_8 = 264
	// BTN_9 as defined in linux/input-event-codes.h:351
	BTN_9 = 265
	// BTN_MOUSE as defined in linux/input-event-codes.h:353
	BTN_MOUSE = 272
	// BTN_LEFT as defined in linux/input-event-codes.h:354
	BTN_LEFT = 272
	// BTN_RIGHT as defined in linux/input-event-codes.h:355
	BTN_RIGHT = 273
	// BTN_MIDDLE as defined in linux/input-event-codes.h:356
	BTN_MIDDLE = 274
	// BTN_SIDE as defined in linux/input-event-codes.h:357
	BTN_SIDE = 275
	// BTN_EXTRA as defined in linux/input-event-codes.h:358
	BTN_EXTRA = 276
	// BTN_FORWARD as defined in linux/input-event-codes.h:359
	BTN_FORWARD = 277
	// BTN_BACK as defined in linux/input-event-codes.h:360
	BTN_BACK = 278
	// BTN_TASK as defined in linux/input-event-codes.h:361
	BTN_TASK = 279
	// BTN_JOYSTICK as defined in linux/input-event-codes.h:363
	BTN_JOYSTICK = 288
	// BTN_TRIGGER as defined in linux/input-event-codes.h:364
	BTN_TRIGGER = 288
	// BTN_THUMB as defined in linux/input-event-codes.h:365
	BTN_THUMB = 289
	// BTN_THUMB2 as defined in linux/input-event-codes.h:366
	BTN_THUMB2 = 290
	// BTN_TOP as defined in linux/input-event-codes.h:367
	BTN_TOP = 291
	// BTN_TOP2 as defined in linux/input-event-codes.h:368
	BTN_TOP2 = 292
	// BTN_PINKIE as defined in linux/input-event-codes.h:369
	BTN_PINKIE = 293
	// BTN_BASE as defined in linux/input-event-codes.h:370
	BTN_BASE = 294
	// BTN_BASE2 as defined in linux/input-event-codes.h:371
	BTN_BASE2 = 295
	// BTN_BASE3 as defined in linux/input-event-codes.h:372
	BTN_BASE3 = 296
	// BTN_BASE4 as defined in linux/input-event-codes.h:373
	BTN_BASE4 = 297
	// BTN_BASE5 as defined in linux/input-event-codes.h:374
	BTN_BASE5 = 298
	// BTN_BASE6 as defined in linux/input-event-codes.h:375
	BTN_BASE6 = 299
	// BTN_DEAD as defined in linux/input-event-codes.h:376
	BTN_DEAD = 303
	// BTN_GAMEPAD as defined in linux/input-event-codes.h:378
	BTN_GAMEPAD = 304
	// BTN_SOUTH as defined in linux/input-event-codes.h:379
	BTN_SOUTH = 304
	// BTN_A as defined in linux/input-event-codes.h:380
	BTN_A = 304
	// BTN_EAST as defined in linux/input-event-codes.h:381
	BTN_EAST = 305
	// BTN_B as defined in linux/input-event-codes.h:382
	BTN_B = 305
	// BTN_C as defined in linux/input-event-codes.h:383
	BTN_C = 306
	// BTN_NORTH as defined in linux/input-event-codes.h:384
	BTN_NORTH = 307
	// BTN_X as defined in linux/input-event-codes.h:385
	BTN_X = 307
	// BTN_WEST as defined in linux/input-event-codes.h:386
	BTN_WEST = 308
	// BTN_Y as defined in linux/input-event-codes.h:387
	BTN_Y = 308
	// BTN_Z as defined in linux/input-event-codes.h:388
	BTN_Z = 309
	// BTN_TL as defined in linux/input-event-codes.h:389
	BTN_TL = 310
	// BTN_TR as defined in linux/input-event-codes.h:390
	BTN_TR = 311
	// BTN_TL2 as defined in linux/input-event-codes.h:391
	BTN_TL2 = 312
	// BTN_TR2 as defined in linux/input-event-codes.h:392
	BTN_TR2 = 313
	// BTN_SELECT as defined in linux/input-event-codes.h:393
	BTN_SELECT = 314
	// BTN_START as defined in linux/input-event-codes.h:394
	BTN_START = 315
	// BTN_MODE as defined in linux/input-event-codes.h:395
	BTN_MODE = 316
	// BTN_THUMBL as defined in linux/input-event-codes.h:396
	BTN_THUMBL = 317
	// BTN_THUMBR as defined in linux/input-event-codes.h:397
	BTN_THUMBR = 318
	// BTN_DIGI as defined in linux/input-event-codes.h:399
	BTN_DIGI = 320
	// BTN_TOOL_PEN as defined in linux/input-event-codes.h:400
	BTN_TOOL_PEN = 320
	// BTN_TOOL_RUBBER as defined in linux/input-event-codes.h:401
	BTN_TOOL_RUBBER = 321
	// BTN_TOOL_BRUSH as defined in linux/input-event-codes.h:402
	BTN_TOOL_BRUSH = 322
	// BTN_TOOL_PENCIL as defined in linux/input-event-codes.h:403
	BTN_TOOL_PENCIL = 323
	// BTN_TOOL_AIRBRUSH as defined in linux/input-event-codes.h:404
	BTN_TOOL_AIRBRUSH = 324
	// BTN_TOOL_FINGER as defined in linux/input-event-codes.h:405
	BTN_TOOL_FINGER = 325
	// BTN_TOOL_MOUSE as defined in linux/input-event-codes.h:406
	BTN_TOOL_MOUSE = 326
	// BTN_TOOL_LENS as defined in linux/input-event-codes.h:407
	BTN_TOOL_LENS = 327
	// BTN_TOOL_QUINTTAP as defined in linux/input-event-codes.h:408
	BTN_TOOL_QUINTTAP = 328
	// BTN_TOUCH as defined in linux/input-event-codes.h:409
	BTN_TOUCH = 330
	// BTN_STYLUS as defined in linux/input-event-codes.h:410
	BTN_STYLUS = 331
	// BTN_STYLUS2 as defined in linux/input-event-codes.h:411
	BTN_STYLUS2 = 332
	// BTN_TOOL_DOUBLETAP as defined in linux/input-event-codes.h:412
	BTN_TOOL_DOUBLETAP = 333
	// BTN_TOOL_TRIPLETAP as defined in linux/input-event-codes.h:413
	BTN_TOOL_TRIPLETAP = 334
	// BTN_TOOL_QUADTAP as defined in linux/input-event-codes.h:414
	BTN_TOOL_QUADTAP = 335
	// BTN_WHEEL as defined in linux/input-event-codes.h:416
	BTN_WHEEL = 336
	// BTN_GEAR_DOWN as defined in linux/input-event-codes.h:417
	BTN_GEAR_DOWN = 336
	// BTN_GEAR_UP as defined in linux/input-event-codes.h:418
	BTN_GEAR_UP = 337
	// KEY_OK as defined in linux/input-event-codes.h:420
	KEY_OK = 352
	// KEY_SELECT as defined in linux/input-event-codes.h:421
	KEY_SELECT = 353
	// KEY_GOTO as defined in linux/input-event-codes.h:422
	KEY_GOTO = 354
	// KEY_CLEAR as defined in linux/input-event-codes.h:423
	KEY_CLEAR = 355
	// KEY_POWER2 as defined in linux/input-event-codes.h:424
	KEY_POWER2 = 356
	// KEY_OPTION as defined in linux/input-event-codes.h:425
	KEY_OPTION = 357
	// KEY_INFO as defined in linux/input-event-codes.h:426
	KEY_INFO = 358
	// KEY_TIME as defined in linux/input-event-codes.h:427
	KEY_TIME = 359
	// KEY_VENDOR as defined in linux/input-event-codes.h:428
	KEY_VENDOR = 360
	// KEY_ARCHIVE as defined in linux/input-event-codes.h:429
	KEY_ARCHIVE = 361
	// KEY_PROGRAM as defined in linux/input-event-codes.h:430
	KEY_PROGRAM = 362
	// KEY_CHANNEL as defined in linux/input-event-codes.h:431
	KEY_CHANNEL = 363
	// KEY_FAVORITES as defined in linux/input-event-codes.h:432
	KEY_FAVORITES = 364
	// KEY_EPG as defined in linux/input-event-codes.h:433
	KEY_EPG = 365
	// KEY_PVR as defined in linux/input-event-codes.h:434
	KEY_PVR = 366
	// KEY_MHP as defined in linux/input-event-codes.h:435
	KEY_MHP = 367
	// KEY_LANGUAGE as defined in linux/input-event-codes.h:436
	KEY_LANGUAGE = 368
	// KEY_TITLE as defined in linux/input-event-codes.h:437
	KEY_TITLE = 369
	// KEY_SUBTITLE as defined in linux/input-event-codes.h:438
	KEY_SUBTITLE = 370
	// KEY_ANGLE as defined in linux/input-event-codes.h:439
	KEY_ANGLE = 371
	// KEY_ZOOM as defined in linux/input-event-codes.h:440
	KEY_ZOOM = 372
	// KEY_MODE as defined in linux/input-event-codes.h:441
	KEY_MODE = 373
	// KEY_KEYBOARD as defined in linux/input-event-codes.h:442
	KEY_KEYBOARD = 374
	// KEY_SCREEN as defined in linux/input-event-codes.h:443
	KEY_SCREEN = 375
	// KEY_PC as defined in linux/input-event-codes.h:444
	KEY_PC = 376
	// KEY_TV as defined in linux/input-event-codes.h:445
	KEY_TV = 377
	// KEY_TV2 as defined in linux/input-event-codes.h:446
	KEY_TV2 = 378
	// KEY_VCR as defined in linux/input-event-codes.h:447
	KEY_VCR = 379
	// KEY_VCR2 as defined in linux/input-event-codes.h:448
	KEY_VCR2 = 380
	// KEY_SAT as defined in linux/input-event-codes.h:449
	KEY_SAT = 381
	// KEY_SAT2 as defined in linux/input-event-codes.h:450
	KEY_SAT2 = 382
	// KEY_CD as defined in linux/input-event-codes.h:451
	KEY_CD = 383
	// KEY_TAPE as defined in linux/input-event-codes.h:452
	KEY_TAPE = 384
	// KEY_RADIO as defined in linux/input-event-codes.h:453
	KEY_RADIO = 385
	// KEY_TUNER as defined in linux/input-event-codes.h:454
	KEY_TUNER = 386
	// KEY_PLAYER as defined in linux/input-event-codes.h:455
	KEY_PLAYER = 387
	// KEY_TEXT as defined in linux/input-event-codes.h:456
	KEY_TEXT = 388
	// KEY_DVD as defined in linux/input-event-codes.h:457
	KEY_DVD = 389
	// KEY_AUX as defined in linux/input-event-codes.h:458
	KEY_AUX = 390
	// KEY_MP3 as defined in linux/input-event-codes.h:459
	KEY_MP3 = 391
	// KEY_AUDIO as defined in linux/input-event-codes.h:460
	KEY_AUDIO = 392
	// KEY_VIDEO as defined in linux/input-event-codes.h:461
	KEY_VIDEO = 393
	// KEY_DIRECTORY as defined in linux/input-event-codes.h:462
	KEY_DIRECTORY = 394
	// KEY_LIST as defined in linux/input-event-codes.h:463
	KEY_LIST = 395
	// KEY_MEMO as defined in linux/input-event-codes.h:464
	KEY_MEMO = 396
	// KEY_CALENDAR as defined in linux/input-event-codes.h:465
	KEY_CALENDAR = 397
	// KEY_RED as defined in linux/input-event-codes.h:466
	KEY_RED = 398
	// KEY_GREEN as defined in linux/input-event-codes.h:467
	KEY_GREEN = 399
	// KEY_YELLOW as defined in linux/input-event-codes.h:468
	KEY_YELLOW = 400
	// KEY_BLUE as defined in linux/input-event-codes.h:469
	KEY_BLUE = 401
	// KEY_CHANNELUP as defined in linux/input-event-codes.h:470
	KEY_CHANNELUP = 402
	// KEY_CHANNELDOWN as defined in linux/input-event-codes.h:471
	KEY_CHANNELDOWN = 403
	// KEY_FIRST as defined in linux/input-event-codes.h:472
	KEY_FIRST = 404
	// KEY_LAST as defined in linux/input-event-codes.h:473
	KEY_LAST = 405
	// KEY_AB as defined in linux/input-event-codes.h:474
	KEY_AB = 406
	// KEY_NEXT as defined in linux/input-event-codes.h:475
	KEY_NEXT = 407
	// KEY_RESTART as defined in linux/input-event-codes.h:476
	KEY_RESTART = 408
	// KEY_SLOW as defined in linux/input-event-codes.h:477
	KEY_SLOW = 409
	// KEY_SHUFFLE as defined in linux/input-event-codes.h:478
	KEY_SHUFFLE = 410
	// KEY_BREAK as defined in linux/input-event-codes.h:479
	KEY_BREAK = 411
	// KEY_PREVIOUS as defined in linux/input-event-codes.h:480
	KEY_PREVIOUS = 412
	// KEY_DIGITS as defined in linux/input-event-codes.h:481
	KEY_DIGITS = 413
	// KEY_TEEN as defined in linux/input-event-codes.h:482
	KEY_TEEN = 414
	// KEY_TWEN as defined in linux/input-event-codes.h:483
	KEY_TWEN = 415
	// KEY_VIDEOPHONE as defined in linux/input-event-codes.h:484
	KEY_VIDEOPHONE = 416
	// KEY_GAMES as defined in linux/input-event-codes.h:485
	KEY_GAMES = 417
	// KEY_ZOOMIN as defined in linux/input-event-codes.h:486
	KEY_ZOOMIN = 418
	// KEY_ZOOMOUT as defined in linux/input-event-codes.h:487
	KEY_ZOOMOUT = 419
	// KEY_ZOOMRESET as defined in linux/input-event-codes.h:488
	KEY_ZOOMRESET = 420
	// KEY_WORDPROCESSOR as defined in linux/input-event-codes.h:489
	KEY_WORDPROCESSOR = 421
	// KEY_EDITOR as defined in linux/input-event-codes.h:490
	KEY_EDITOR = 422
	// KEY_SPREADSHEET as defined in linux/input-event-codes.h:491
	KEY_SPREADSHEET = 423
	// KEY_GRAPHICSEDITOR as defined in linux/input-event-codes.h:492
	KEY_GRAPHICSEDITOR = 424
	// KEY_PRESENTATION as defined in linux/input-event-codes.h:493
	KEY_PRESENTATION = 425
	// KEY_DATABASE as defined in linux/input-event-codes.h:494
	KEY_DATABASE = 426
	// KEY_NEWS as defined in linux/input-event-codes.h:495
	KEY_NEWS = 427
	// KEY_VOICEMAIL as defined in linux/input-event-codes.h:496
	KEY_VOICEMAIL = 428
	// KEY_ADDRESSBOOK as defined in linux/input-event-codes.h:497
	KEY_ADDRESSBOOK = 429
	// KEY_MESSENGER as defined in linux/input-event-codes.h:498
	KEY_MESSENGER = 430
	// KEY_DISPLAYTOGGLE as defined in linux/input-event-codes.h:499
	KEY_DISPLAYTOGGLE = 431
	// KEY_BRIGHTNESS_TOGGLE as defined in linux/input-event-codes.h:500
	KEY_BRIGHTNESS_TOGGLE = 431
	// KEY_SPELLCHECK as defined in linux/input-event-codes.h:501
	KEY_SPELLCHECK = 432
	// KEY_LOGOFF as defined in linux/input-event-codes.h:502
	KEY_LOGOFF = 433
	// KEY_DOLLAR as defined in linux/input-event-codes.h:504
	KEY_DOLLAR = 434
	// KEY_EURO as defined in linux/input-event-codes.h:505
	KEY_EURO = 435
	// KEY_FRAMEBACK as defined in linux/input-event-codes.h:507
	KEY_FRAMEBACK = 436
	// KEY_FRAMEFORWARD as defined in linux/input-event-codes.h:508
	KEY_FRAMEFORWARD = 437
	// KEY_CONTEXT_MENU as defined in linux/input-event-codes.h:509
	KEY_CONTEXT_MENU = 438
	// KEY_MEDIA_REPEAT as defined in linux/input-event-codes.h:510
	KEY_MEDIA_REPEAT = 439
	// KEY_10CHANNELSUP as defined in linux/input-event-codes.h:511
	KEY_10CHANNELSUP = 440
	// KEY_10CHANNELSDOWN as defined in linux/input-event-codes.h:512
	KEY_10CHANNELSDOWN = 441
	// KEY_IMAGES as defined in linux/input-event-codes.h:513
	KEY_IMAGES = 442
	// KEY_DEL_EOL as defined in linux/input-event-codes.h:515
	KEY_DEL_EOL = 448
	// KEY_DEL_EOS as defined in linux/input-event-codes.h:516
	KEY_DEL_EOS = 449
	// KEY_INS_LINE as defined in linux/input-event-codes.h:517
	KEY_INS_LINE = 450
	// KEY_DEL_LINE as defined in linux/input-event-codes.h:518
	KEY_DEL_LINE = 451
	// KEY_FN as defined in linux/input-event-codes.h:520
	KEY_FN = 464
	// KEY_FN_ESC as defined in linux/input-event-codes.h:521
	KEY_FN_ESC = 465
	// KEY_FN_F1 as defined in linux/input-event-codes.h:522
	KEY_FN_F1 = 466
	// KEY_FN_F2 as defined in linux/input-event-codes.h:523
	KEY_FN_F2 = 467
	// KEY_FN_F3 as defined in linux/input-event-codes.h:524
	KEY_FN_F3 = 468
	// KEY_FN_F4 as defined in linux/input-event-codes.h:525
	KEY_FN_F4 = 469
	// KEY_FN_F5 as defined in linux/input-event-codes.h:526
	KEY_FN_F5 = 470
	// KEY_FN_F6 as defined in linux/input-event-codes.h:527
	KEY_FN_F6 = 471
	// KEY_FN_F7 as defined in linux/input-event-codes.h:528
	KEY_FN_F7 = 472
	// KEY_FN_F8 as defined in linux/input-event-codes.h:529
	KEY_FN_F8 = 473
	// KEY_FN_F9 as defined in linux/input-event-codes.h:530
	KEY_FN_F9 = 474
	// KEY_FN_F10 as defined in linux/input-event-codes.h:531
	KEY_FN_F10 = 475
	// KEY_FN_F11 as defined in linux/input-event-codes.h:532
	KEY_FN_F11 = 476
	// KEY_FN_F12 as defined in linux/input-event-codes.h:533
	KEY_FN_F12 = 477
	// KEY_FN_1 as defined in linux/input-event-codes.h:534
	KEY_FN_1 = 478
	// KEY_FN_2 as defined in linux/input-event-codes.h:535
	KEY_FN_2 = 479
	// KEY_FN_D as defined in linux/input-event-codes.h:536
	KEY_FN_D = 480
	// KEY_FN_E as defined in linux/input-event-codes.h:537
	KEY_FN_E = 481
	// KEY_FN_F as defined in linux/input-event-codes.h:538
	KEY_FN_F = 482
	// KEY_FN_S as defined in linux/input-event-codes.h:539
	KEY_FN_S = 483
	// KEY_FN_B as defined in linux/input-event-codes.h:540
	KEY_FN_B = 484
	// KEY_BRL_DOT1 as defined in linux/input-event-codes.h:542
	KEY_BRL_DOT1 = 497
	// KEY_BRL_DOT2 as defined in linux/input-event-codes.h:543
	KEY_BRL_DOT2 = 498
	// KEY_BRL_DOT3 as defined in linux/input-event-codes.h:544
	KEY_BRL_DOT3 = 499
	// KEY_BRL_DOT4 as defined in linux/input-event-codes.h:545
	KEY_BRL_DOT4 = 500
	// KEY_BRL_DOT5 as defined in linux/input-event-codes.h:546
	KEY_BRL_DOT5 = 501
	// KEY_BRL_DOT6 as defined in linux/input-event-codes.h:547
	KEY_BRL_DOT6 = 502
	// KEY_BRL_DOT7 as defined in linux/input-event-codes.h:548
	KEY_BRL_DOT7 = 503
	// KEY_BRL_DOT8 as defined in linux/input-event-codes.h:549
	KEY_BRL_DOT8 = 504
	// KEY_BRL_DOT9 as defined in linux/input-event-codes.h:550
	KEY_BRL_DOT9 = 505
	// KEY_BRL_DOT10 as defined in linux/input-event-codes.h:551
	KEY_BRL_DOT10 = 506
	// KEY_NUMERIC_0 as defined in linux/input-event-codes.h:553
	KEY_NUMERIC_0 = 512
	// KEY_NUMERIC_1 as defined in linux/input-event-codes.h:554
	KEY_NUMERIC_1 = 513
	// KEY_NUMERIC_2 as defined in linux/input-event-codes.h:555
	KEY_NUMERIC_2 = 514
	// KEY_NUMERIC_3 as defined in linux/input-event-codes.h:556
	KEY_NUMERIC_3 = 515
	// KEY_NUMERIC_4 as defined in linux/input-event-codes.h:557
	KEY_NUMERIC_4 = 516
	// KEY_NUMERIC_5 as defined in linux/input-event-codes.h:558
	KEY_NUMERIC_5 = 517
	// KEY_NUMERIC_6 as defined in linux/input-event-codes.h:559
	KEY_NUMERIC_6 = 518
	// KEY_NUMERIC_7 as defined in linux/input-event-codes.h:560
	KEY_NUMERIC_7 = 519
	// KEY_NUMERIC_8 as defined in linux/input-event-codes.h:561
	KEY_NUMERIC_8 = 520
	// KEY_NUMERIC_9 as defined in linux/input-event-codes.h:562
	KEY_NUMERIC_9 = 521
	// KEY_NUMERIC_STAR as defined in linux/input-event-codes.h:563
	KEY_NUMERIC_STAR = 522
	// KEY_NUMERIC_POUND as defined in linux/input-event-codes.h:564
	KEY_NUMERIC_POUND = 523
	// KEY_NUMERIC_A as defined in linux/input-event-codes.h:565
	KEY_NUMERIC_A = 524
	// KEY_NUMERIC_B as defined in linux/input-event-codes.h:566
	KEY_NUMERIC_B = 525
	// KEY_NUMERIC_C as defined in linux/input-event-codes.h:567
	KEY_NUMERIC_C = 526
	// KEY_NUMERIC_D as defined in linux/input-event-codes.h:568
	KEY_NUMERIC_D = 527
	// KEY_CAMERA_FOCUS as defined in linux/input-event-codes.h:570
	KEY_CAMERA_FOCUS = 528
	// KEY_WPS_BUTTON as defined in linux/input-event-codes.h:571
	KEY_WPS_BUTTON = 529
	// KEY_TOUCHPAD_TOGGLE as defined in linux/input-event-codes.h:573
	KEY_TOUCHPAD_TOGGLE = 530
	// KEY_TOUCHPAD_ON as defined in linux/input-event-codes.h:574
	KEY_TOUCHPAD_ON = 531
	// KEY_TOUCHPAD_OFF as defined in linux/input-event-codes.h:575
	KEY_TOUCHPAD_OFF = 532
	// KEY_CAMERA_ZOOMIN as defined in linux/input-event-codes.h:577
	KEY_CAMERA_ZOOMIN = 533
	// KEY_CAMERA_ZOOMOUT as defined in linux/input-event-codes.h:578
	KEY_CAMERA_ZOOMOUT = 534
	// KEY_CAMERA_UP as defined in linux/input-event-codes.h:579
	KEY_CAMERA_UP = 535
	// KEY_CAMERA_DOWN as defined in linux/input-event-codes.h:580
	KEY_CAMERA_DOWN = 536
	// KEY_CAMERA_LEFT as defined in linux/input-event-codes.h:581
	KEY_CAMERA_LEFT = 537
	// KEY_CAMERA_RIGHT as defined in linux/input-event-codes.h:582
	KEY_CAMERA_RIGHT = 538
	// KEY_ATTENDANT_ON as defined in linux/input-event-codes.h:584
	KEY_ATTENDANT_ON = 539
	// KEY_ATTENDANT_OFF as defined in linux/input-event-codes.h:585
	KEY_ATTENDANT_OFF = 540
	// KEY_ATTENDANT_TOGGLE as defined in linux/input-event-codes.h:586
	KEY_ATTENDANT_TOGGLE = 541
	// KEY_LIGHTS_TOGGLE as defined in linux/input-event-codes.h:587
	KEY_LIGHTS_TOGGLE = 542
	// BTN_DPAD_UP as defined in linux/input-event-codes.h:589
	BTN_DPAD_UP = 544
	// BTN_DPAD_DOWN as defined in linux/input-event-codes.h:590
	BTN_DPAD_DOWN = 545
	// BTN_DPAD_LEFT as defined in linux/input-event-codes.h:591
	BTN_DPAD_LEFT = 546
	// BTN_DPAD_RIGHT as defined in linux/input-event-codes.h:592
	BTN_DPAD_RIGHT = 547
	// KEY_ALS_TOGGLE as defined in linux/input-event-codes.h:594
	KEY_ALS_TOGGLE = 560
	// KEY_BUTTONCONFIG as defined in linux/input-event-codes.h:596
	KEY_BUTTONCONFIG = 576
	// KEY_TASKMANAGER as defined in linux/input-event-codes.h:597
	KEY_TASKMANAGER = 577
	// KEY_JOURNAL as defined in linux/input-event-codes.h:598
	KEY_JOURNAL = 578
	// KEY_CONTROLPANEL as defined in linux/input-event-codes.h:599
	KEY_CONTROLPANEL = 579
	// KEY_APPSELECT as defined in linux/input-event-codes.h:600
	KEY_APPSELECT = 580
	// KEY_SCREENSAVER as defined in linux/input-event-codes.h:601
	KEY_SCREENSAVER = 581
	// KEY_VOICECOMMAND as defined in linux/input-event-codes.h:602
	KEY_VOICECOMMAND = 582
	// KEY_BRIGHTNESS_MIN as defined in linux/input-event-codes.h:604
	KEY_BRIGHTNESS_MIN = 592
	// KEY_BRIGHTNESS_MAX as defined in linux/input-event-codes.h:605
	KEY_BRIGHTNESS_MAX = 593
	// KEY_KBDINPUTASSIST_PREV as defined in linux/input-event-codes.h:607
	KEY_KBDINPUTASSIST_PREV = 608
	// KEY_KBDINPUTASSIST_NEXT as defined in linux/input-event-codes.h:608
	KEY_KBDINPUTASSIST_NEXT = 609
	// KEY_KBDINPUTASSIST_PREVGROUP as defined in linux/input-event-codes.h:609
	KEY_KBDINPUTASSIST_PREVGROUP = 610
	// KEY_KBDINPUTASSIST_NEXTGROUP as defined in linux/input-event-codes.h:610
	KEY_KBDINPUTASSIST_NEXTGROUP = 611
	// KEY_KBDINPUTASSIST_ACCEPT as defined in linux/input-event-codes.h:611
	KEY_KBDINPUTASSIST_ACCEPT = 612
	// KEY_KBDINPUTASSIST_CANCEL as defined in linux/input-event-codes.h:612
	KEY_KBDINPUTASSIST_CANCEL = 613
	// KEY_RIGHT_UP as defined in linux/input-event-codes.h:615
	KEY_RIGHT_UP = 614
	// KEY_RIGHT_DOWN as defined in linux/input-event-codes.h:616
	KEY_RIGHT_DOWN = 615
	// KEY_LEFT_UP as defined in linux/input-event-codes.h:617
	KEY_LEFT_UP = 616
	// KEY_LEFT_DOWN as defined in linux/input-event-codes.h:618
	KEY_LEFT_DOWN = 617
	// KEY_ROOT_MENU as defined in linux/input-event-codes.h:620
	KEY_ROOT_MENU = 618
	// KEY_MEDIA_TOP_MENU as defined in linux/input-event-codes.h:622
	KEY_MEDIA_TOP_MENU = 619
	// KEY_NUMERIC_11 as defined in linux/input-event-codes.h:623
	KEY_NUMERIC_11 = 620
	// KEY_NUMERIC_12 as defined in linux/input-event-codes.h:624
	KEY_NUMERIC_12 = 621
	// KEY_AUDIO_DESC as defined in linux/input-event-codes.h:630
	KEY_AUDIO_DESC = 622
	// KEY_3D_MODE as defined in linux/input-event-codes.h:631
	KEY_3D_MODE = 623
	// KEY_NEXT_FAVORITE as defined in linux/input-event-codes.h:632
	KEY_NEXT_FAVORITE = 624
	// KEY_STOP_RECORD as defined in linux/input-event-codes.h:633
	KEY_STOP_RECORD = 625
	// KEY_PAUSE_RECORD as defined in linux/input-event-codes.h:634
	KEY_PAUSE_RECORD = 626
	// KEY_VOD as defined in linux/input-event-codes.h:635
	KEY_VOD = 627
	// KEY_UNMUTE as defined in linux/input-event-codes.h:636
	KEY_UNMUTE = 628
	// KEY_FASTREVERSE as defined in linux/input-event-codes.h:637
	KEY_FASTREVERSE = 629
	// KEY_SLOWREVERSE as defined in linux/input-event-codes.h:638
	KEY_SLOWREVERSE = 630
	// KEY_DATA as defined in linux/input-event-codes.h:643
	KEY_DATA = 631
	// BTN_TRIGGER_HAPPY as defined in linux/input-event-codes.h:645
	BTN_TRIGGER_HAPPY = 704
	// BTN_TRIGGER_HAPPY1 as defined in linux/input-event-codes.h:646
	BTN_TRIGGER_HAPPY1 = 704
	// BTN_TRIGGER_HAPPY2 as defined in linux/input-event-codes.h:647
	BTN_TRIGGER_HAPPY2 = 705
	// BTN_TRIGGER_HAPPY3 as defined in linux/input-event-codes.h:648
	BTN_TRIGGER_HAPPY3 = 706
	// BTN_TRIGGER_HAPPY4 as defined in linux/input-event-codes.h:649
	BTN_TRIGGER_HAPPY4 = 707
	// BTN_TRIGGER_HAPPY5 as defined in linux/input-event-codes.h:650
	BTN_TRIGGER_HAPPY5 = 708
	// BTN_TRIGGER_HAPPY6 as defined in linux/input-event-codes.h:651
	BTN_TRIGGER_HAPPY6 = 709
	// BTN_TRIGGER_HAPPY7 as defined in linux/input-event-codes.h:652
	BTN_TRIGGER_HAPPY7 = 710
	// BTN_TRIGGER_HAPPY8 as defined in linux/input-event-codes.h:653
	BTN_TRIGGER_HAPPY8 = 711
	// BTN_TRIGGER_HAPPY9 as defined in linux/input-event-codes.h:654
	BTN_TRIGGER_HAPPY9 = 712
	// BTN_TRIGGER_HAPPY10 as defined in linux/input-event-codes.h:655
	BTN_TRIGGER_HAPPY10 = 713
	// BTN_TRIGGER_HAPPY11 as defined in linux/input-event-codes.h:656
	BTN_TRIGGER_HAPPY11 = 714
	// BTN_TRIGGER_HAPPY12 as defined in linux/input-event-codes.h:657
	BTN_TRIGGER_HAPPY12 = 715
	// BTN_TRIGGER_HAPPY13 as defined in linux/input-event-codes.h:658
	BTN_TRIGGER_HAPPY13 = 716
	// BTN_TRIGGER_HAPPY14 as defined in linux/input-event-codes.h:659
	BTN_TRIGGER_HAPPY14 = 717
	// BTN_TRIGGER_HAPPY15 as defined in linux/input-event-codes.h:660
	BTN_TRIGGER_HAPPY15 = 718
	// BTN_TRIGGER_HAPPY16 as defined in linux/input-event-codes.h:661
	BTN_TRIGGER_HAPPY16 = 719
	// BTN_TRIGGER_HAPPY17 as defined in linux/input-event-codes.h:662
	BTN_TRIGGER_HAPPY17 = 720
	// BTN_TRIGGER_HAPPY18 as defined in linux/input-event-codes.h:663
	BTN_TRIGGER_HAPPY18 = 721
	// BTN_TRIGGER_HAPPY19 as defined in linux/input-event-codes.h:664
	BTN_TRIGGER_HAPPY19 = 722
	// BTN_TRIGGER_HAPPY20 as defined in linux/input-event-codes.h:665
	BTN_TRIGGER_HAPPY20 = 723
	// BTN_TRIGGER_HAPPY21 as defined in linux/input-event-codes.h:666
	BTN_TRIGGER_HAPPY21 = 724
	// BTN_TRIGGER_HAPPY22 as defined in linux/input-event-codes.h:667
	BTN_TRIGGER_HAPPY22 = 725
	// BTN_TRIGGER_HAPPY23 as defined in linux/input-event-codes.h:668
	BTN_TRIGGER_HAPPY23 = 726
	// BTN_TRIGGER_HAPPY24 as defined in linux/input-event-codes.h:669
	BTN_TRIGGER_HAPPY24 = 727
	// BTN_TRIGGER_HAPPY25 as defined in linux/input-event-codes.h:670
	BTN_TRIGGER_HAPPY25 = 728
	// BTN_TRIGGER_HAPPY26 as defined in linux/input-event-codes.h:671
	BTN_TRIGGER_HAPPY26 = 729
	// BTN_TRIGGER_HAPPY27 as defined in linux/input-event-codes.h:672
	BTN_TRIGGER_HAPPY27 = 730
	// BTN_TRIGGER_HAPPY28 as defined in linux/input-event-codes.h:673
	BTN_TRIGGER_HAPPY28 = 731
	// BTN_TRIGGER_HAPPY29 as defined in linux/input-event-codes.h:674
	BTN_TRIGGER_HAPPY29 = 732
	// BTN_TRIGGER_HAPPY30 as defined in linux/input-event-codes.h:675
	BTN_TRIGGER_HAPPY30 = 733
	// BTN_TRIGGER_HAPPY31 as defined in linux/input-event-codes.h:676
	BTN_TRIGGER_HAPPY31 = 734
	// BTN_TRIGGER_HAPPY32 as defined in linux/input-event-codes.h:677
	BTN_TRIGGER_HAPPY32 = 735
	// BTN_TRIGGER_HAPPY33 as defined in linux/input-event-codes.h:678
	BTN_TRIGGER_HAPPY33 = 736
	// BTN_TRIGGER_HAPPY34 as defined in linux/input-event-codes.h:679
	BTN_TRIGGER_HAPPY34 = 737
	// BTN_TRIGGER_HAPPY35 as defined in linux/input-event-codes.h:680
	BTN_TRIGGER_HAPPY35 = 738
	// BTN_TRIGGER_HAPPY36 as defined in linux/input-event-codes.h:681
	BTN_TRIGGER_HAPPY36 = 739
	// BTN_TRIGGER_HAPPY37 as defined in linux/input-event-codes.h:682
	BTN_TRIGGER_HAPPY37 = 740
	// BTN_TRIGGER_HAPPY38 as defined in linux/input-event-codes.h:683
	BTN_TRIGGER_HAPPY38 = 741
	// BTN_TRIGGER_HAPPY39 as defined in linux/input-event-codes.h:684
	BTN_TRIGGER_HAPPY39 = 742
	// BTN_TRIGGER_HAPPY40 as defined in linux/input-event-codes.h:685
	BTN_TRIGGER_HAPPY40 = 743
	// KEY_MIN_INTERESTING as defined in linux/input-event-codes.h:688
	KEY_MIN_INTERESTING = 113
	// KEY_MAX as defined in linux/input-event-codes.h:689
	KEY_MAX = 767
	// KEY_CNT as defined in linux/input-event-codes.h:690
	KEY_CNT = 768
	// REL_X as defined in linux/input-event-codes.h:696
	REL_X = 0
	// REL_Y as defined in linux/input-event-codes.h:697
	REL_Y = 1
	// REL_Z as defined in linux/input-event-codes.h:698
	REL_Z = 2
	// REL_RX as defined in linux/input-event-codes.h:699
	REL_RX = 3
	// REL_RY as defined in linux/input-event-codes.h:700
	REL_RY = 4
	// REL_RZ as defined in linux/input-event-codes.h:701
	REL_RZ = 5
	// REL_HWHEEL as defined in linux/input-event-codes.h:702
	REL_HWHEEL = 6
	// REL_DIAL as defined in linux/input-event-codes.h:703
	REL_DIAL = 7
	// REL_WHEEL as defined in linux/input-event-codes.h:704
	REL_WHEEL = 8
	// REL_MISC as defined in linux/input-event-codes.h:705
	REL_MISC = 9
	// REL_MAX as defined in linux/input-event-codes.h:706
	REL_MAX = 15
	// REL_CNT as defined in linux/input-event-codes.h:707
	REL_CNT = 16
	// ABS_X as defined in linux/input-event-codes.h:713
	ABS_X = 0
	// ABS_Y as defined in linux/input-event-codes.h:714
	ABS_Y = 1
	// ABS_Z as defined in linux/input-event-codes.h:715
	ABS_Z = 2
	// ABS_RX as defined in linux/input-event-codes.h:716
	ABS_RX = 3
	// ABS_RY as defined in linux/input-event-codes.h:717
	ABS_RY = 4
	// ABS_RZ as defined in linux/input-event-codes.h:718
	ABS_RZ = 5
	// ABS_THROTTLE as defined in linux/input-event-codes.h:719
	ABS_THROTTLE = 6
	// ABS_RUDDER as defined in linux/input-event-codes.h:720
	ABS_RUDDER = 7
	// ABS_WHEEL as defined in linux/input-event-codes.h:721
	ABS_WHEEL = 8
	// ABS_GAS as defined in linux/input-event-codes.h:722
	ABS_GAS = 9
	// ABS_BRAKE as defined in linux/input-event-codes.h:723
	ABS_BRAKE = 10
	// ABS_HAT0X as defined in linux/input-event-codes.h:724
	ABS_HAT0X = 16
	// ABS_HAT0Y as defined in linux/input-event-codes.h:725
	ABS_HAT0Y = 17
	// ABS_HAT1X as defined in linux/input-event-codes.h:726
	ABS_HAT1X = 18
	// ABS_HAT1Y as defined in linux/input-event-codes.h:727
	ABS_HAT1Y = 19
	// ABS_HAT2X as defined in linux/input-event-codes.h:728
	ABS_HAT2X = 20
	// ABS_HAT2Y as defined in linux/input-event-codes.h:729
	ABS_HAT2Y = 21
	// ABS_HAT3X as defined in linux/input-event-codes.h:730
	ABS_HAT3X = 22
	// ABS_HAT3Y as defined in linux/input-event-codes.h:731
	ABS_HAT3Y = 23
	// ABS_PRESSURE as defined in linux/input-event-codes.h:732
	ABS_PRESSURE = 24
	// ABS_DISTANCE as defined in linux/input-event-codes.h:733
	ABS_DISTANCE = 25
	// ABS_TILT_X as defined in linux/input-event-codes.h:734
	ABS_TILT_X = 26
	// ABS_TILT_Y as defined in linux/input-event-codes.h:735
	ABS_TILT_Y = 27
	// ABS_TOOL_WIDTH as defined in linux/input-event-codes.h:736
	ABS_TOOL_WIDTH = 28
	// ABS_VOLUME as defined in linux/input-event-codes.h:738
	ABS_VOLUME = 32
	// ABS_MISC as defined in linux/input-event-codes.h:740
	ABS_MISC = 40
	// ABS_MT_SLOT as defined in linux/input-event-codes.h:742
	ABS_MT_SLOT = 47
	// ABS_MT_TOUCH_MAJOR as defined in linux/input-event-codes.h:743
	ABS_MT_TOUCH_MAJOR = 48
	// ABS_MT_TOUCH_MINOR as defined in linux/input-event-codes.h:744
	ABS_MT_TOUCH_MINOR = 49
	// ABS_MT_WIDTH_MAJOR as defined in linux/input-event-codes.h:745
	ABS_MT_WIDTH_MAJOR = 50
	// ABS_MT_WIDTH_MINOR as defined in linux/input-event-codes.h:746
	ABS_MT_WIDTH_MINOR = 51
	// ABS_MT_ORIENTATION as defined in linux/input-event-codes.h:747
	ABS_MT_ORIENTATION = 52
	// ABS_MT_POSITION_X as defined in linux/input-event-codes.h:748
	ABS_MT_POSITION_X = 53
	// ABS_MT_POSITION_Y as defined in linux/input-event-codes.h:749
	ABS_MT_POSITION_Y = 54
	// ABS_MT_TOOL_TYPE as defined in linux/input-event-codes.h:750
	ABS_MT_TOOL_TYPE = 55
	// ABS_MT_BLOB_ID as defined in linux/input-event-codes.h:751
	ABS_MT_BLOB_ID = 56
	// ABS_MT_TRACKING_ID as defined in linux/input-event-codes.h:752
	ABS_MT_TRACKING_ID = 57
	// ABS_MT_PRESSURE as defined in linux/input-event-codes.h:753
	ABS_MT_PRESSURE = 58
	// ABS_MT_DISTANCE as defined in linux/input-event-codes.h:754
	ABS_MT_DISTANCE = 59
	// ABS_MT_TOOL_X as defined in linux/input-event-codes.h:755
	ABS_MT_TOOL_X = 60
	// ABS_MT_TOOL_Y as defined in linux/input-event-codes.h:756
	ABS_MT_TOOL_Y = 61
	// ABS_MAX as defined in linux/input-event-codes.h:759
	ABS_MAX = 63
	// ABS_CNT as defined in linux/input-event-codes.h:760
	ABS_CNT = 64
	// MSC_SERIAL as defined in linux/input-event-codes.h:791
	MSC_SERIAL = 0
	// MSC_PULSELED as defined in linux/input-event-codes.h:792
	MSC_PULSELED = 1
	// MSC_GESTURE as defined in linux/input-event-codes.h:793
	MSC_GESTURE = 2
	// MSC_RAW as defined in linux/input-event-codes.h:794
	MSC_RAW = 3
	// MSC_SCAN as defined in linux/input-event-codes.h:795
	MSC_SCAN = 4
	// MSC_TIMESTAMP as defined in linux/input-event-codes.h:796
	MSC_TIMESTAMP = 5
	// MSC_MAX as defined in linux/input-event-codes.h:797
	MSC_MAX = 7
	// MSC_CNT as defined in linux/input-event-codes.h:798
	MSC_CNT = 8
)

// _ItimerWhich as declared in sys/time.h:88
type _ItimerWhich int32

// _ItimerWhich enumeration from sys/time.h:88
const ()
