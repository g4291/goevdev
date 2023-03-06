package evdev

/*
* source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
 */

const INPUT_PROP_POINTER uint16 = 0x00        /* needs a pointer */
const INPUT_PROP_DIRECT uint16 = 0x01         /* direct input devices */
const INPUT_PROP_BUTTONPAD uint16 = 0x02      /* has button(s) under pad */
const INPUT_PROP_SEMI_MT uint16 = 0x03        /* touch rectangle only */
const INPUT_PROP_TOPBUTTONPAD uint16 = 0x04   /* softbuttons at top of pad */
const INPUT_PROP_POINTING_STICK uint16 = 0x05 /* is a pointing stick */
const INPUT_PROP_ACCELEROMETER uint16 = 0x06  /* has accelerometer */

const INPUT_PROP_MAX uint16 = 0x1f
const INPUT_PROP_CNT uint16 = (INPUT_PROP_MAX + 1)

/*
 * Event types
 */

const EV_SYN uint16 = 0x00
const EV_KEY uint16 = 0x01
const EV_REL uint16 = 0x02
const EV_ABS uint16 = 0x03
const EV_MSC uint16 = 0x04
const EV_SW uint16 = 0x05
const EV_LED uint16 = 0x11
const EV_SND uint16 = 0x12
const EV_REP uint16 = 0x14
const EV_FF uint16 = 0x15
const EV_PWR uint16 = 0x16
const EV_FF_STATUS uint16 = 0x17
const EV_MAX uint16 = 0x1f
const EV_CNT uint16 = (EV_MAX + 1)

/*
 * Synchronization events.
 */

const SYN_REPORT uint16 = 0
const SYN_CONFIG uint16 = 1
const SYN_MT_REPORT uint16 = 2
const SYN_DROPPED uint16 = 3
const SYN_MAX uint16 = 0xf
const SYN_CNT uint16 = (SYN_MAX + 1)

/*
 * Keys and buttons
 *
 * Most of the keys/buttons are modeled after USB HUT 1.12
 * (see http://www.usb.org/developers/hidpage).
 * Abbreviations in the comments:
 * AC - Application Control
 * AL - Application Launch Button
 * SC - System Control
 */

const KEY_RESERVED uint16 = 0
const KEY_ESC uint16 = 1
const KEY_1 uint16 = 2
const KEY_2 uint16 = 3
const KEY_3 uint16 = 4
const KEY_4 uint16 = 5
const KEY_5 uint16 = 6
const KEY_6 uint16 = 7
const KEY_7 uint16 = 8
const KEY_8 uint16 = 9
const KEY_9 uint16 = 10
const KEY_0 uint16 = 11
const KEY_MINUS uint16 = 12
const KEY_EQUAL uint16 = 13
const KEY_BACKSPACE uint16 = 14
const KEY_TAB uint16 = 15
const KEY_Q uint16 = 16
const KEY_W uint16 = 17
const KEY_E uint16 = 18
const KEY_R uint16 = 19
const KEY_T uint16 = 20
const KEY_Y uint16 = 21
const KEY_U uint16 = 22
const KEY_I uint16 = 23
const KEY_O uint16 = 24
const KEY_P uint16 = 25
const KEY_LEFTBRACE uint16 = 26
const KEY_RIGHTBRACE uint16 = 27
const KEY_ENTER uint16 = 28
const KEY_LEFTCTRL uint16 = 29
const KEY_A uint16 = 30
const KEY_S uint16 = 31
const KEY_D uint16 = 32
const KEY_F uint16 = 33
const KEY_G uint16 = 34
const KEY_H uint16 = 35
const KEY_J uint16 = 36
const KEY_K uint16 = 37
const KEY_L uint16 = 38
const KEY_SEMICOLON uint16 = 39
const KEY_APOSTROPHE uint16 = 40
const KEY_GRAVE uint16 = 41
const KEY_LEFTSHIFT uint16 = 42
const KEY_BACKSLASH uint16 = 43
const KEY_Z uint16 = 44
const KEY_X uint16 = 45
const KEY_C uint16 = 46
const KEY_V uint16 = 47
const KEY_B uint16 = 48
const KEY_N uint16 = 49
const KEY_M uint16 = 50
const KEY_COMMA uint16 = 51
const KEY_DOT uint16 = 52
const KEY_SLASH uint16 = 53
const KEY_RIGHTSHIFT uint16 = 54
const KEY_KPASTERISK uint16 = 55
const KEY_LEFTALT uint16 = 56
const KEY_SPACE uint16 = 57
const KEY_CAPSLOCK uint16 = 58
const KEY_F1 uint16 = 59
const KEY_F2 uint16 = 60
const KEY_F3 uint16 = 61
const KEY_F4 uint16 = 62
const KEY_F5 uint16 = 63
const KEY_F6 uint16 = 64
const KEY_F7 uint16 = 65
const KEY_F8 uint16 = 66
const KEY_F9 uint16 = 67
const KEY_F10 uint16 = 68
const KEY_NUMLOCK uint16 = 69
const KEY_SCROLLLOCK uint16 = 70
const KEY_KP7 uint16 = 71
const KEY_KP8 uint16 = 72
const KEY_KP9 uint16 = 73
const KEY_KPMINUS uint16 = 74
const KEY_KP4 uint16 = 75
const KEY_KP5 uint16 = 76
const KEY_KP6 uint16 = 77
const KEY_KPPLUS uint16 = 78
const KEY_KP1 uint16 = 79
const KEY_KP2 uint16 = 80
const KEY_KP3 uint16 = 81
const KEY_KP0 uint16 = 82
const KEY_KPDOT uint16 = 83

const KEY_ZENKAKUHANKAKU uint16 = 85
const KEY_102ND uint16 = 86
const KEY_F11 uint16 = 87
const KEY_F12 uint16 = 88
const KEY_RO uint16 = 89
const KEY_KATAKANA uint16 = 90
const KEY_HIRAGANA uint16 = 91
const KEY_HENKAN uint16 = 92
const KEY_KATAKANAHIRAGANA uint16 = 93
const KEY_MUHENKAN uint16 = 94
const KEY_KPJPCOMMA uint16 = 95
const KEY_KPENTER uint16 = 96
const KEY_RIGHTCTRL uint16 = 97
const KEY_KPSLASH uint16 = 98
const KEY_SYSRQ uint16 = 99
const KEY_RIGHTALT uint16 = 100
const KEY_LINEFEED uint16 = 101
const KEY_HOME uint16 = 102
const KEY_UP uint16 = 103
const KEY_PAGEUP uint16 = 104
const KEY_LEFT uint16 = 105
const KEY_RIGHT uint16 = 106
const KEY_END uint16 = 107
const KEY_DOWN uint16 = 108
const KEY_PAGEDOWN uint16 = 109
const KEY_INSERT uint16 = 110
const KEY_DELETE uint16 = 111
const KEY_MACRO uint16 = 112
const KEY_MUTE uint16 = 113
const KEY_VOLUMEDOWN uint16 = 114
const KEY_VOLUMEUP uint16 = 115
const KEY_POWER uint16 = 116 /* SC System Power Down */
const KEY_KPEQUAL uint16 = 117
const KEY_KPPLUSMINUS uint16 = 118
const KEY_PAUSE uint16 = 119
const KEY_SCALE uint16 = 120 /* AL Compiz Scale (Expose) */

const KEY_KPCOMMA uint16 = 121
const KEY_HANGEUL uint16 = 122
const KEY_HANGUEL uint16 = KEY_HANGEUL
const KEY_HANJA uint16 = 123
const KEY_YEN uint16 = 124
const KEY_LEFTMETA uint16 = 125
const KEY_RIGHTMETA uint16 = 126
const KEY_COMPOSE uint16 = 127

const KEY_STOP uint16 = 128 /* AC Stop */
const KEY_AGAIN uint16 = 129
const KEY_PROPS uint16 = 130 /* AC Properties */
const KEY_UNDO uint16 = 131  /* AC Undo */
const KEY_FRONT uint16 = 132
const KEY_COPY uint16 = 133  /* AC Copy */
const KEY_OPEN uint16 = 134  /* AC Open */
const KEY_PASTE uint16 = 135 /* AC Paste */
const KEY_FIND uint16 = 136  /* AC Search */
const KEY_CUT uint16 = 137   /* AC Cut */
const KEY_HELP uint16 = 138  /* AL Integrated Help Center */
const KEY_MENU uint16 = 139  /* Menu (show menu) */
const KEY_CALC uint16 = 140  /* AL Calculator */
const KEY_SETUP uint16 = 141
const KEY_SLEEP uint16 = 142  /* SC System Sleep */
const KEY_WAKEUP uint16 = 143 /* System Wake Up */
const KEY_FILE uint16 = 144   /* AL Local Machine Browser */
const KEY_SENDFILE uint16 = 145
const KEY_DELETEFILE uint16 = 146
const KEY_XFER uint16 = 147
const KEY_PROG1 uint16 = 148
const KEY_PROG2 uint16 = 149
const KEY_WWW uint16 = 150 /* AL Internet Browser */
const KEY_MSDOS uint16 = 151
const KEY_COFFEE uint16 = 152 /* AL Terminal Lock/Screensaver */
const KEY_SCREENLOCK uint16 = KEY_COFFEE
const KEY_ROTATE_DISPLAY uint16 = 153 /* Display orientation for e.g. tablets */
const KEY_DIRECTION uint16 = KEY_ROTATE_DISPLAY
const KEY_CYCLEWINDOWS uint16 = 154
const KEY_MAIL uint16 = 155
const KEY_BOOKMARKS uint16 = 156 /* AC Bookmarks */
const KEY_COMPUTER uint16 = 157
const KEY_BACK uint16 = 158    /* AC Back */
const KEY_FORWARD uint16 = 159 /* AC Forward */
const KEY_CLOSECD uint16 = 160
const KEY_EJECTCD uint16 = 161
const KEY_EJECTCLOSECD uint16 = 162
const KEY_NEXTSONG uint16 = 163
const KEY_PLAYPAUSE uint16 = 164
const KEY_PREVIOUSSONG uint16 = 165
const KEY_STOPCD uint16 = 166
const KEY_RECORD uint16 = 167
const KEY_REWIND uint16 = 168
const KEY_PHONE uint16 = 169 /* Media Select Telephone */
const KEY_ISO uint16 = 170
const KEY_CONFIG uint16 = 171   /* AL Consumer Control Configuration */
const KEY_HOMEPAGE uint16 = 172 /* AC Home */
const KEY_REFRESH uint16 = 173  /* AC Refresh */
const KEY_EXIT uint16 = 174     /* AC Exit */
const KEY_MOVE uint16 = 175
const KEY_EDIT uint16 = 176
const KEY_SCROLLUP uint16 = 177
const KEY_SCROLLDOWN uint16 = 178
const KEY_KPLEFTPAREN uint16 = 179
const KEY_KPRIGHTPAREN uint16 = 180
const KEY_NEW uint16 = 181  /* AC New */
const KEY_REDO uint16 = 182 /* AC Redo/Repeat */

const KEY_F13 uint16 = 183
const KEY_F14 uint16 = 184
const KEY_F15 uint16 = 185
const KEY_F16 uint16 = 186
const KEY_F17 uint16 = 187
const KEY_F18 uint16 = 188
const KEY_F19 uint16 = 189
const KEY_F20 uint16 = 190
const KEY_F21 uint16 = 191
const KEY_F22 uint16 = 192
const KEY_F23 uint16 = 193
const KEY_F24 uint16 = 194

const KEY_PLAYCD uint16 = 200
const KEY_PAUSECD uint16 = 201
const KEY_PROG3 uint16 = 202
const KEY_PROG4 uint16 = 203
const KEY_ALL_APPLICATIONS uint16 = 204 /* AC Desktop Show All Applications */
const KEY_DASHBOARD uint16 = KEY_ALL_APPLICATIONS
const KEY_SUSPEND uint16 = 205
const KEY_CLOSE uint16 = 206 /* AC Close */
const KEY_PLAY uint16 = 207
const KEY_FASTFORWARD uint16 = 208
const KEY_BASSBOOST uint16 = 209
const KEY_PRINT uint16 = 210 /* AC Print */
const KEY_HP uint16 = 211
const KEY_CAMERA uint16 = 212
const KEY_SOUND uint16 = 213
const KEY_QUESTION uint16 = 214
const KEY_EMAIL uint16 = 215
const KEY_CHAT uint16 = 216
const KEY_SEARCH uint16 = 217
const KEY_CONNECT uint16 = 218
const KEY_FINANCE uint16 = 219 /* AL Checkbook/Finance */
const KEY_SPORT uint16 = 220
const KEY_SHOP uint16 = 221
const KEY_ALTERASE uint16 = 222
const KEY_CANCEL uint16 = 223 /* AC Cancel */
const KEY_BRIGHTNESSDOWN uint16 = 224
const KEY_BRIGHTNESSUP uint16 = 225
const KEY_MEDIA uint16 = 226

const KEY_SWITCHVIDEOMODE uint16 = 227 /* Cycle between available video outputs (Monitor/LCD/TV-out/etc) */
const KEY_KBDILLUMTOGGLE uint16 = 228
const KEY_KBDILLUMDOWN uint16 = 229
const KEY_KBDILLUMUP uint16 = 230

const KEY_SEND uint16 = 231        /* AC Send */
const KEY_REPLY uint16 = 232       /* AC Reply */
const KEY_FORWARDMAIL uint16 = 233 /* AC Forward Msg */
const KEY_SAVE uint16 = 234        /* AC Save */
const KEY_DOCUMENTS uint16 = 235

const KEY_BATTERY uint16 = 236

const KEY_BLUETOOTH uint16 = 237
const KEY_WLAN uint16 = 238
const KEY_UWB uint16 = 239

const KEY_UNKNOWN uint16 = 240

const KEY_VIDEO_NEXT uint16 = 241       /* drive next video source */
const KEY_VIDEO_PREV uint16 = 242       /* drive previous video source */
const KEY_BRIGHTNESS_CYCLE uint16 = 243 /* brightness up, after max is min */
const KEY_BRIGHTNESS_AUTO uint16 = 244
const KEY_BRIGHTNESS_ZERO uint16 = KEY_BRIGHTNESS_AUTO
const KEY_DISPLAY_OFF uint16 = 245 /* display device to off state */

const KEY_WWAN uint16 = 246 /* Wireless WAN (LTE, UMTS, GSM, etc.) */
const KEY_WIMAX uint16 = KEY_WWAN
const KEY_RFKILL uint16 = 247 /* Key that controls all radios */

const KEY_MICMUTE uint16 = 248 /* Mute / unmute the microphone */

/* Code 255 is reserved for special needs of AT keyboard driver */

const BTN_MISC uint16 = 0x100
const BTN_0 uint16 = 0x100
const BTN_1 uint16 = 0x101
const BTN_2 uint16 = 0x102
const BTN_3 uint16 = 0x103
const BTN_4 uint16 = 0x104
const BTN_5 uint16 = 0x105
const BTN_6 uint16 = 0x106
const BTN_7 uint16 = 0x107
const BTN_8 uint16 = 0x108
const BTN_9 uint16 = 0x109

const BTN_MOUSE uint16 = 0x110
const BTN_LEFT uint16 = 0x110
const BTN_RIGHT uint16 = 0x111
const BTN_MIDDLE uint16 = 0x112
const BTN_SIDE uint16 = 0x113
const BTN_EXTRA uint16 = 0x114
const BTN_FORWARD uint16 = 0x115
const BTN_BACK uint16 = 0x116
const BTN_TASK uint16 = 0x117

const BTN_JOYSTICK uint16 = 0x120
const BTN_TRIGGER uint16 = 0x120
const BTN_THUMB uint16 = 0x121
const BTN_THUMB2 uint16 = 0x122
const BTN_TOP uint16 = 0x123
const BTN_TOP2 uint16 = 0x124
const BTN_PINKIE uint16 = 0x125
const BTN_BASE uint16 = 0x126
const BTN_BASE2 uint16 = 0x127
const BTN_BASE3 uint16 = 0x128
const BTN_BASE4 uint16 = 0x129
const BTN_BASE5 uint16 = 0x12a
const BTN_BASE6 uint16 = 0x12b
const BTN_DEAD uint16 = 0x12f

const BTN_GAMEPAD uint16 = 0x130
const BTN_SOUTH uint16 = 0x130
const BTN_A uint16 = BTN_SOUTH
const BTN_EAST uint16 = 0x131
const BTN_B uint16 = BTN_EAST
const BTN_C uint16 = 0x132
const BTN_NORTH uint16 = 0x133
const BTN_X uint16 = BTN_NORTH
const BTN_WEST uint16 = 0x134
const BTN_Y uint16 = BTN_WEST
const BTN_Z uint16 = 0x135
const BTN_TL uint16 = 0x136
const BTN_TR uint16 = 0x137
const BTN_TL2 uint16 = 0x138
const BTN_TR2 uint16 = 0x139
const BTN_SELECT uint16 = 0x13a
const BTN_START uint16 = 0x13b
const BTN_MODE uint16 = 0x13c
const BTN_THUMBL uint16 = 0x13d
const BTN_THUMBR uint16 = 0x13e

const BTN_DIGI uint16 = 0x140
const BTN_TOOL_PEN uint16 = 0x140
const BTN_TOOL_RUBBER uint16 = 0x141
const BTN_TOOL_BRUSH uint16 = 0x142
const BTN_TOOL_PENCIL uint16 = 0x143
const BTN_TOOL_AIRBRUSH uint16 = 0x144
const BTN_TOOL_FINGER uint16 = 0x145
const BTN_TOOL_MOUSE uint16 = 0x146
const BTN_TOOL_LENS uint16 = 0x147
const BTN_TOOL_QUINTTAP uint16 = 0x148 /* Five fingers on trackpad */
const BTN_STYLUS3 uint16 = 0x149
const BTN_TOUCH uint16 = 0x14a
const BTN_STYLUS uint16 = 0x14b
const BTN_STYLUS2 uint16 = 0x14c
const BTN_TOOL_DOUBLETAP uint16 = 0x14d
const BTN_TOOL_TRIPLETAP uint16 = 0x14e
const BTN_TOOL_QUADTAP uint16 = 0x14f /* Four fingers on trackpad */

const BTN_WHEEL uint16 = 0x150
const BTN_GEAR_DOWN uint16 = 0x150
const BTN_GEAR_UP uint16 = 0x151

const KEY_OK uint16 = 0x160
const KEY_SELECT uint16 = 0x161
const KEY_GOTO uint16 = 0x162
const KEY_CLEAR uint16 = 0x163
const KEY_POWER2 uint16 = 0x164
const KEY_OPTION uint16 = 0x165
const KEY_INFO uint16 = 0x166 /* AL OEM Features/Tips/Tutorial */
const KEY_TIME uint16 = 0x167
const KEY_VENDOR uint16 = 0x168
const KEY_ARCHIVE uint16 = 0x169
const KEY_PROGRAM uint16 = 0x16a /* Media Select Program Guide */
const KEY_CHANNEL uint16 = 0x16b
const KEY_FAVORITES uint16 = 0x16c
const KEY_EPG uint16 = 0x16d
const KEY_PVR uint16 = 0x16e /* Media Select Home */
const KEY_MHP uint16 = 0x16f
const KEY_LANGUAGE uint16 = 0x170
const KEY_TITLE uint16 = 0x171
const KEY_SUBTITLE uint16 = 0x172
const KEY_ANGLE uint16 = 0x173
const KEY_FULL_SCREEN uint16 = 0x174 /* AC View Toggle */
const KEY_ZOOM uint16 = KEY_FULL_SCREEN
const KEY_MODE uint16 = 0x175
const KEY_KEYBOARD uint16 = 0x176
const KEY_ASPECT_RATIO uint16 = 0x177 /* HUTRR37: Aspect */
const KEY_SCREEN uint16 = KEY_ASPECT_RATIO
const KEY_PC uint16 = 0x178   /* Media Select Computer */
const KEY_TV uint16 = 0x179   /* Media Select TV */
const KEY_TV2 uint16 = 0x17a  /* Media Select Cable */
const KEY_VCR uint16 = 0x17b  /* Media Select VCR */
const KEY_VCR2 uint16 = 0x17c /* VCR Plus */
const KEY_SAT uint16 = 0x17d  /* Media Select Satellite */
const KEY_SAT2 uint16 = 0x17e
const KEY_CD uint16 = 0x17f   /* Media Select CD */
const KEY_TAPE uint16 = 0x180 /* Media Select Tape */
const KEY_RADIO uint16 = 0x181
const KEY_TUNER uint16 = 0x182 /* Media Select Tuner */
const KEY_PLAYER uint16 = 0x183
const KEY_TEXT uint16 = 0x184
const KEY_DVD uint16 = 0x185 /* Media Select DVD */
const KEY_AUX uint16 = 0x186
const KEY_MP3 uint16 = 0x187
const KEY_AUDIO uint16 = 0x188 /* AL Audio Browser */
const KEY_VIDEO uint16 = 0x189 /* AL Movie Browser */
const KEY_DIRECTORY uint16 = 0x18a
const KEY_LIST uint16 = 0x18b
const KEY_MEMO uint16 = 0x18c /* Media Select Messages */
const KEY_CALENDAR uint16 = 0x18d
const KEY_RED uint16 = 0x18e
const KEY_GREEN uint16 = 0x18f
const KEY_YELLOW uint16 = 0x190
const KEY_BLUE uint16 = 0x191
const KEY_CHANNELUP uint16 = 0x192   /* Channel Increment */
const KEY_CHANNELDOWN uint16 = 0x193 /* Channel Decrement */
const KEY_FIRST uint16 = 0x194
const KEY_LAST uint16 = 0x195 /* Recall Last */
const KEY_AB uint16 = 0x196
const KEY_NEXT uint16 = 0x197
const KEY_RESTART uint16 = 0x198
const KEY_SLOW uint16 = 0x199
const KEY_SHUFFLE uint16 = 0x19a
const KEY_BREAK uint16 = 0x19b
const KEY_PREVIOUS uint16 = 0x19c
const KEY_DIGITS uint16 = 0x19d
const KEY_TEEN uint16 = 0x19e
const KEY_TWEN uint16 = 0x19f
const KEY_VIDEOPHONE uint16 = 0x1a0     /* Media Select Video Phone */
const KEY_GAMES uint16 = 0x1a1          /* Media Select Games */
const KEY_ZOOMIN uint16 = 0x1a2         /* AC Zoom In */
const KEY_ZOOMOUT uint16 = 0x1a3        /* AC Zoom Out */
const KEY_ZOOMRESET uint16 = 0x1a4      /* AC Zoom */
const KEY_WORDPROCESSOR uint16 = 0x1a5  /* AL Word Processor */
const KEY_EDITOR uint16 = 0x1a6         /* AL Text Editor */
const KEY_SPREADSHEET uint16 = 0x1a7    /* AL Spreadsheet */
const KEY_GRAPHICSEDITOR uint16 = 0x1a8 /* AL Graphics Editor */
const KEY_PRESENTATION uint16 = 0x1a9   /* AL Presentation App */
const KEY_DATABASE uint16 = 0x1aa       /* AL Database App */
const KEY_NEWS uint16 = 0x1ab           /* AL Newsreader */
const KEY_VOICEMAIL uint16 = 0x1ac      /* AL Voicemail */
const KEY_ADDRESSBOOK uint16 = 0x1ad    /* AL Contacts/Address Book */
const KEY_MESSENGER uint16 = 0x1ae      /* AL Instant Messaging */
const KEY_DISPLAYTOGGLE uint16 = 0x1af  /* Turn display (LCD) on and off */
const KEY_BRIGHTNESS_TOGGLE uint16 = KEY_DISPLAYTOGGLE
const KEY_SPELLCHECK uint16 = 0x1b0 /* AL Spell Check */
const KEY_LOGOFF uint16 = 0x1b1     /* AL Logoff */

const KEY_DOLLAR uint16 = 0x1b2
const KEY_EURO uint16 = 0x1b3

const KEY_FRAMEBACK uint16 = 0x1b4 /* Consumer - transport controls */
const KEY_FRAMEFORWARD uint16 = 0x1b5
const KEY_CONTEXT_MENU uint16 = 0x1b6        /* GenDesc - system context menu */
const KEY_MEDIA_REPEAT uint16 = 0x1b7        /* Consumer - transport control */
const KEY_10CHANNELSUP uint16 = 0x1b8        /* 10 channels up (10+) */
const KEY_10CHANNELSDOWN uint16 = 0x1b9      /* 10 channels down (10-) */
const KEY_IMAGES uint16 = 0x1ba              /* AL Image Browser */
const KEY_NOTIFICATION_CENTER uint16 = 0x1bc /* Show/hide the notification center */
const KEY_PICKUP_PHONE uint16 = 0x1bd        /* Answer incoming call */
const KEY_HANGUP_PHONE uint16 = 0x1be        /* Decline incoming call */

const KEY_DEL_EOL uint16 = 0x1c0
const KEY_DEL_EOS uint16 = 0x1c1
const KEY_INS_LINE uint16 = 0x1c2
const KEY_DEL_LINE uint16 = 0x1c3

const KEY_FN uint16 = 0x1d0
const KEY_FN_ESC uint16 = 0x1d1
const KEY_FN_F1 uint16 = 0x1d2
const KEY_FN_F2 uint16 = 0x1d3
const KEY_FN_F3 uint16 = 0x1d4
const KEY_FN_F4 uint16 = 0x1d5
const KEY_FN_F5 uint16 = 0x1d6
const KEY_FN_F6 uint16 = 0x1d7
const KEY_FN_F7 uint16 = 0x1d8
const KEY_FN_F8 uint16 = 0x1d9
const KEY_FN_F9 uint16 = 0x1da
const KEY_FN_F10 uint16 = 0x1db
const KEY_FN_F11 uint16 = 0x1dc
const KEY_FN_F12 uint16 = 0x1dd
const KEY_FN_1 uint16 = 0x1de
const KEY_FN_2 uint16 = 0x1df
const KEY_FN_D uint16 = 0x1e0
const KEY_FN_E uint16 = 0x1e1
const KEY_FN_F uint16 = 0x1e2
const KEY_FN_S uint16 = 0x1e3
const KEY_FN_B uint16 = 0x1e4
const KEY_FN_RIGHT_SHIFT uint16 = 0x1e5

const KEY_BRL_DOT1 uint16 = 0x1f1
const KEY_BRL_DOT2 uint16 = 0x1f2
const KEY_BRL_DOT3 uint16 = 0x1f3
const KEY_BRL_DOT4 uint16 = 0x1f4
const KEY_BRL_DOT5 uint16 = 0x1f5
const KEY_BRL_DOT6 uint16 = 0x1f6
const KEY_BRL_DOT7 uint16 = 0x1f7
const KEY_BRL_DOT8 uint16 = 0x1f8
const KEY_BRL_DOT9 uint16 = 0x1f9
const KEY_BRL_DOT10 uint16 = 0x1fa

const KEY_NUMERIC_0 uint16 = 0x200 /* used by phones, remote controls, */
const KEY_NUMERIC_1 uint16 = 0x201 /* and other keypads */
const KEY_NUMERIC_2 uint16 = 0x202
const KEY_NUMERIC_3 uint16 = 0x203
const KEY_NUMERIC_4 uint16 = 0x204
const KEY_NUMERIC_5 uint16 = 0x205
const KEY_NUMERIC_6 uint16 = 0x206
const KEY_NUMERIC_7 uint16 = 0x207
const KEY_NUMERIC_8 uint16 = 0x208
const KEY_NUMERIC_9 uint16 = 0x209
const KEY_NUMERIC_STAR uint16 = 0x20a
const KEY_NUMERIC_POUND uint16 = 0x20b
const KEY_NUMERIC_A uint16 = 0x20c /* Phone key A - HUT Telephony 0xb9 */
const KEY_NUMERIC_B uint16 = 0x20d
const KEY_NUMERIC_C uint16 = 0x20e
const KEY_NUMERIC_D uint16 = 0x20f

const KEY_CAMERA_FOCUS uint16 = 0x210
const KEY_WPS_BUTTON uint16 = 0x211 /* WiFi Protected Setup key */

const KEY_TOUCHPAD_TOGGLE uint16 = 0x212 /* Request switch touchpad on or off */
const KEY_TOUCHPAD_ON uint16 = 0x213
const KEY_TOUCHPAD_OFF uint16 = 0x214

const KEY_CAMERA_ZOOMIN uint16 = 0x215
const KEY_CAMERA_ZOOMOUT uint16 = 0x216
const KEY_CAMERA_UP uint16 = 0x217
const KEY_CAMERA_DOWN uint16 = 0x218
const KEY_CAMERA_LEFT uint16 = 0x219
const KEY_CAMERA_RIGHT uint16 = 0x21a

const KEY_ATTENDANT_ON uint16 = 0x21b
const KEY_ATTENDANT_OFF uint16 = 0x21c
const KEY_ATTENDANT_TOGGLE uint16 = 0x21d /* Attendant call on or off */
const KEY_LIGHTS_TOGGLE uint16 = 0x21e    /* Reading light on or off */

const BTN_DPAD_UP uint16 = 0x220
const BTN_DPAD_DOWN uint16 = 0x221
const BTN_DPAD_LEFT uint16 = 0x222
const BTN_DPAD_RIGHT uint16 = 0x223

const KEY_ALS_TOGGLE uint16 = 0x230         /* Ambient light sensor */
const KEY_ROTATE_LOCK_TOGGLE uint16 = 0x231 /* Display rotation lock */

const KEY_BUTTONCONFIG uint16 = 0x240          /* AL Button Configuration */
const KEY_TASKMANAGER uint16 = 0x241           /* AL Task/Project Manager */
const KEY_JOURNAL uint16 = 0x242               /* AL Log/Journal/Timecard */
const KEY_CONTROLPANEL uint16 = 0x243          /* AL Control Panel */
const KEY_APPSELECT uint16 = 0x244             /* AL Select Task/Application */
const KEY_SCREENSAVER uint16 = 0x245           /* AL Screen Saver */
const KEY_VOICECOMMAND uint16 = 0x246          /* Listening Voice Command */
const KEY_ASSISTANT uint16 = 0x247             /* AL Context-aware desktop assistant */
const KEY_KBD_LAYOUT_NEXT uint16 = 0x248       /* AC Next Keyboard Layout Select */
const KEY_EMOJI_PICKER uint16 = 0x249          /* Show/hide emoji picker (HUTRR101) */
const KEY_DICTATE uint16 = 0x24a               /* Start or Stop Voice Dictation Session (HUTRR99) */
const KEY_CAMERA_ACCESS_ENABLE uint16 = 0x24b  /* Enables programmatic access to camera devices. (HUTRR72) */
const KEY_CAMERA_ACCESS_DISABLE uint16 = 0x24c /* Disables programmatic access to camera devices. (HUTRR72) */
const KEY_CAMERA_ACCESS_TOGGLE uint16 = 0x24d  /* Toggles the current state of the camera access control. (HUTRR72) */

const KEY_BRIGHTNESS_MIN uint16 = 0x250 /* Set Brightness to Minimum */
const KEY_BRIGHTNESS_MAX uint16 = 0x251 /* Set Brightness to Maximum */

const KEY_KBDINPUTASSIST_PREV uint16 = 0x260
const KEY_KBDINPUTASSIST_NEXT uint16 = 0x261
const KEY_KBDINPUTASSIST_PREVGROUP uint16 = 0x262
const KEY_KBDINPUTASSIST_NEXTGROUP uint16 = 0x263
const KEY_KBDINPUTASSIST_ACCEPT uint16 = 0x264
const KEY_KBDINPUTASSIST_CANCEL uint16 = 0x265

/* Diagonal movement keys */
const KEY_RIGHT_UP uint16 = 0x266
const KEY_RIGHT_DOWN uint16 = 0x267
const KEY_LEFT_UP uint16 = 0x268
const KEY_LEFT_DOWN uint16 = 0x269

const KEY_ROOT_MENU uint16 = 0x26a /* Show Device's Root Menu */
/* Show Top Menu of the Media (e.g. DVD) */
const KEY_MEDIA_TOP_MENU uint16 = 0x26b
const KEY_NUMERIC_11 uint16 = 0x26c
const KEY_NUMERIC_12 uint16 = 0x26d

/*
 * Toggle Audio Description: refers to an audio service that helps blind and
 * visually impaired consumers understand the action in a program. Note: in
 * some countries this is referred to as "Video Description".
 */
const KEY_AUDIO_DESC uint16 = 0x26e
const KEY_3D_MODE uint16 = 0x26f
const KEY_NEXT_FAVORITE uint16 = 0x270
const KEY_STOP_RECORD uint16 = 0x271
const KEY_PAUSE_RECORD uint16 = 0x272
const KEY_VOD uint16 = 0x273 /* Video on Demand */
const KEY_UNMUTE uint16 = 0x274
const KEY_FASTREVERSE uint16 = 0x275
const KEY_SLOWREVERSE uint16 = 0x276

/*
 * Control a data application associated with the currently viewed channel,
 * e.g. teletext or data broadcast application (MHEG, MHP, HbbTV, etc.)
 */
const KEY_DATA uint16 = 0x277
const KEY_ONSCREEN_KEYBOARD uint16 = 0x278

/* Electronic privacy screen control */
const KEY_PRIVACY_SCREEN_TOGGLE uint16 = 0x279

/* Select an area of screen to be copied */
const KEY_SELECTIVE_SCREENSHOT uint16 = 0x27a

/* Move the focus to the next or previous user controllable element within a UI container */
const KEY_NEXT_ELEMENT uint16 = 0x27b
const KEY_PREVIOUS_ELEMENT uint16 = 0x27c

/* Toggle Autopilot engagement */
const KEY_AUTOPILOT_ENGAGE_TOGGLE uint16 = 0x27d

/* Shortcut Keys */
const KEY_MARK_WAYPOINT uint16 = 0x27e
const KEY_SOS uint16 = 0x27f
const KEY_NAV_CHART uint16 = 0x280
const KEY_FISHING_CHART uint16 = 0x281
const KEY_SINGLE_RANGE_RADAR uint16 = 0x282
const KEY_DUAL_RANGE_RADAR uint16 = 0x283
const KEY_RADAR_OVERLAY uint16 = 0x284
const KEY_TRADITIONAL_SONAR uint16 = 0x285
const KEY_CLEARVU_SONAR uint16 = 0x286
const KEY_SIDEVU_SONAR uint16 = 0x287
const KEY_NAV_INFO uint16 = 0x288
const KEY_BRIGHTNESS_MENU uint16 = 0x289

/*
 * Some keyboards have keys which do not have a defined meaning, these keys
 * are intended to be programmed / bound to macros by the user. For most
 * keyboards with these macro-keys the key-sequence to inject, or action to
 * take, is all handled by software on the host side. So from the kernel's
 * point of view these are just normal keys.
 *
 * The KEY_MACRO# codes below are intended for such keys, which may be labeled
 * e.g. G1-G18, or S1 - S30. The KEY_MACRO# codes MUST NOT be used for keys
 * where the marking on the key does indicate a defined meaning / purpose.
 *
 * The KEY_MACRO# codes MUST also NOT be used as fallback for when no existing
 * KEY_FOO define matches the marking / purpose. In this case a new KEY_FOO
 * define MUST be added.
 */
const KEY_MACRO1 uint16 = 0x290
const KEY_MACRO2 uint16 = 0x291
const KEY_MACRO3 uint16 = 0x292
const KEY_MACRO4 uint16 = 0x293
const KEY_MACRO5 uint16 = 0x294
const KEY_MACRO6 uint16 = 0x295
const KEY_MACRO7 uint16 = 0x296
const KEY_MACRO8 uint16 = 0x297
const KEY_MACRO9 uint16 = 0x298
const KEY_MACRO10 uint16 = 0x299
const KEY_MACRO11 uint16 = 0x29a
const KEY_MACRO12 uint16 = 0x29b
const KEY_MACRO13 uint16 = 0x29c
const KEY_MACRO14 uint16 = 0x29d
const KEY_MACRO15 uint16 = 0x29e
const KEY_MACRO16 uint16 = 0x29f
const KEY_MACRO17 uint16 = 0x2a0
const KEY_MACRO18 uint16 = 0x2a1
const KEY_MACRO19 uint16 = 0x2a2
const KEY_MACRO20 uint16 = 0x2a3
const KEY_MACRO21 uint16 = 0x2a4
const KEY_MACRO22 uint16 = 0x2a5
const KEY_MACRO23 uint16 = 0x2a6
const KEY_MACRO24 uint16 = 0x2a7
const KEY_MACRO25 uint16 = 0x2a8
const KEY_MACRO26 uint16 = 0x2a9
const KEY_MACRO27 uint16 = 0x2aa
const KEY_MACRO28 uint16 = 0x2ab
const KEY_MACRO29 uint16 = 0x2ac
const KEY_MACRO30 uint16 = 0x2ad

/*
 * Some keyboards with the macro-keys described above have some extra keys
 * for controlling the host-side software responsible for the macro handling:
 * -A macro recording start/stop key. Note that not all keyboards which emit
 *  KEY_MACRO_RECORD_START will also emit KEY_MACRO_RECORD_STOP if
 *  KEY_MACRO_RECORD_STOP is not advertised, then KEY_MACRO_RECORD_START
 *  should be interpreted as a recording start/stop toggle;
 * -Keys for switching between different macro (pre)sets, either a key for
 *  cycling through the configured presets or keys to directly select a preset.
 */
const KEY_MACRO_RECORD_START uint16 = 0x2b0
const KEY_MACRO_RECORD_STOP uint16 = 0x2b1
const KEY_MACRO_PRESET_CYCLE uint16 = 0x2b2
const KEY_MACRO_PRESET1 uint16 = 0x2b3
const KEY_MACRO_PRESET2 uint16 = 0x2b4
const KEY_MACRO_PRESET3 uint16 = 0x2b5

/*
 * Some keyboards have a buildin LCD panel where the contents are controlled
 * by the host. Often these have a number of keys directly below the LCD
 * intended for controlling a menu shown on the LCD. These keys often don't
 * have any labeling so we just name them KEY_KBD_LCD_MENU#
 */
const KEY_KBD_LCD_MENU1 uint16 = 0x2b8
const KEY_KBD_LCD_MENU2 uint16 = 0x2b9
const KEY_KBD_LCD_MENU3 uint16 = 0x2ba
const KEY_KBD_LCD_MENU4 uint16 = 0x2bb
const KEY_KBD_LCD_MENU5 uint16 = 0x2bc

const BTN_TRIGGER_HAPPY uint16 = 0x2c0
const BTN_TRIGGER_HAPPY1 uint16 = 0x2c0
const BTN_TRIGGER_HAPPY2 uint16 = 0x2c1
const BTN_TRIGGER_HAPPY3 uint16 = 0x2c2
const BTN_TRIGGER_HAPPY4 uint16 = 0x2c3
const BTN_TRIGGER_HAPPY5 uint16 = 0x2c4
const BTN_TRIGGER_HAPPY6 uint16 = 0x2c5
const BTN_TRIGGER_HAPPY7 uint16 = 0x2c6
const BTN_TRIGGER_HAPPY8 uint16 = 0x2c7
const BTN_TRIGGER_HAPPY9 uint16 = 0x2c8
const BTN_TRIGGER_HAPPY10 uint16 = 0x2c9
const BTN_TRIGGER_HAPPY11 uint16 = 0x2ca
const BTN_TRIGGER_HAPPY12 uint16 = 0x2cb
const BTN_TRIGGER_HAPPY13 uint16 = 0x2cc
const BTN_TRIGGER_HAPPY14 uint16 = 0x2cd
const BTN_TRIGGER_HAPPY15 uint16 = 0x2ce
const BTN_TRIGGER_HAPPY16 uint16 = 0x2cf
const BTN_TRIGGER_HAPPY17 uint16 = 0x2d0
const BTN_TRIGGER_HAPPY18 uint16 = 0x2d1
const BTN_TRIGGER_HAPPY19 uint16 = 0x2d2
const BTN_TRIGGER_HAPPY20 uint16 = 0x2d3
const BTN_TRIGGER_HAPPY21 uint16 = 0x2d4
const BTN_TRIGGER_HAPPY22 uint16 = 0x2d5
const BTN_TRIGGER_HAPPY23 uint16 = 0x2d6
const BTN_TRIGGER_HAPPY24 uint16 = 0x2d7
const BTN_TRIGGER_HAPPY25 uint16 = 0x2d8
const BTN_TRIGGER_HAPPY26 uint16 = 0x2d9
const BTN_TRIGGER_HAPPY27 uint16 = 0x2da
const BTN_TRIGGER_HAPPY28 uint16 = 0x2db
const BTN_TRIGGER_HAPPY29 uint16 = 0x2dc
const BTN_TRIGGER_HAPPY30 uint16 = 0x2dd
const BTN_TRIGGER_HAPPY31 uint16 = 0x2de
const BTN_TRIGGER_HAPPY32 uint16 = 0x2df
const BTN_TRIGGER_HAPPY33 uint16 = 0x2e0
const BTN_TRIGGER_HAPPY34 uint16 = 0x2e1
const BTN_TRIGGER_HAPPY35 uint16 = 0x2e2
const BTN_TRIGGER_HAPPY36 uint16 = 0x2e3
const BTN_TRIGGER_HAPPY37 uint16 = 0x2e4
const BTN_TRIGGER_HAPPY38 uint16 = 0x2e5
const BTN_TRIGGER_HAPPY39 uint16 = 0x2e6
const BTN_TRIGGER_HAPPY40 uint16 = 0x2e7

/* We avoid low common keys in module aliases so they don't get huge. */
const KEY_MIN_INTERESTING uint16 = KEY_MUTE
const KEY_MAX uint16 = 0x2ff
const KEY_CNT uint16 = (KEY_MAX + 1)

/*
 * Relative axes
 */

const REL_X uint16 = 0x00
const REL_Y uint16 = 0x01
const REL_Z uint16 = 0x02
const REL_RX uint16 = 0x03
const REL_RY uint16 = 0x04
const REL_RZ uint16 = 0x05
const REL_HWHEEL uint16 = 0x06
const REL_DIAL uint16 = 0x07
const REL_WHEEL uint16 = 0x08
const REL_MISC uint16 = 0x09

/*
 * 0x0a is reserved and should not be used in input drivers.
 * It was used by HID as REL_MISC+1 and userspace needs to detect if
 * the next REL_* event is correct or is just REL_MISC + n.
 * We define here REL_RESERVED so userspace can rely on it and detect
 * the situation described above.
 */
const REL_RESERVED uint16 = 0x0a
const REL_WHEEL_HI_RES uint16 = 0x0b
const REL_HWHEEL_HI_RES uint16 = 0x0c
const REL_MAX uint16 = 0x0f
const REL_CNT uint16 = (REL_MAX + 1)

/*
 * Absolute axes
 */

const ABS_X uint16 = 0x00
const ABS_Y uint16 = 0x01
const ABS_Z uint16 = 0x02
const ABS_RX uint16 = 0x03
const ABS_RY uint16 = 0x04
const ABS_RZ uint16 = 0x05
const ABS_THROTTLE uint16 = 0x06
const ABS_RUDDER uint16 = 0x07
const ABS_WHEEL uint16 = 0x08
const ABS_GAS uint16 = 0x09
const ABS_BRAKE uint16 = 0x0a
const ABS_HAT0X uint16 = 0x10
const ABS_HAT0Y uint16 = 0x11
const ABS_HAT1X uint16 = 0x12
const ABS_HAT1Y uint16 = 0x13
const ABS_HAT2X uint16 = 0x14
const ABS_HAT2Y uint16 = 0x15
const ABS_HAT3X uint16 = 0x16
const ABS_HAT3Y uint16 = 0x17
const ABS_PRESSURE uint16 = 0x18
const ABS_DISTANCE uint16 = 0x19
const ABS_TILT_X uint16 = 0x1a
const ABS_TILT_Y uint16 = 0x1b
const ABS_TOOL_WIDTH uint16 = 0x1c

const ABS_VOLUME uint16 = 0x20
const ABS_PROFILE uint16 = 0x21

const ABS_MISC uint16 = 0x28

/*
 * 0x2e is reserved and should not be used in input drivers.
 * It was used by HID as ABS_MISC+6 and userspace needs to detect if
 * the next ABS_* event is correct or is just ABS_MISC + n.
 * We define here ABS_RESERVED so userspace can rely on it and detect
 * the situation described above.
 */
const ABS_RESERVED uint16 = 0x2e

const ABS_MT_SLOT uint16 = 0x2f        /* MT slot being modified */
const ABS_MT_TOUCH_MAJOR uint16 = 0x30 /* Major axis of touching ellipse */
const ABS_MT_TOUCH_MINOR uint16 = 0x31 /* Minor axis (omit if circular) */
const ABS_MT_WIDTH_MAJOR uint16 = 0x32 /* Major axis of approaching ellipse */
const ABS_MT_WIDTH_MINOR uint16 = 0x33 /* Minor axis (omit if circular) */
const ABS_MT_ORIENTATION uint16 = 0x34 /* Ellipse orientation */
const ABS_MT_POSITION_X uint16 = 0x35  /* Center X touch position */
const ABS_MT_POSITION_Y uint16 = 0x36  /* Center Y touch position */
const ABS_MT_TOOL_TYPE uint16 = 0x37   /* Type of touching device */
const ABS_MT_BLOB_ID uint16 = 0x38     /* Group a set of packets as a blob */
const ABS_MT_TRACKING_ID uint16 = 0x39 /* Unique ID of initiated contact */
const ABS_MT_PRESSURE uint16 = 0x3a    /* Pressure on contact area */
const ABS_MT_DISTANCE uint16 = 0x3b    /* Contact hover distance */
const ABS_MT_TOOL_X uint16 = 0x3c      /* Center X tool position */
const ABS_MT_TOOL_Y uint16 = 0x3d      /* Center Y tool position */

const ABS_MAX uint16 = 0x3f
const ABS_CNT uint16 = (ABS_MAX + 1)

/*
 * Switch events
 */

const SW_LID uint16 = 0x00              /* set = lid shut */
const SW_TABLET_MODE uint16 = 0x01      /* set = tablet mode */
const SW_HEADPHONE_INSERT uint16 = 0x02 /* set = inserted */
const SW_RFKILL_ALL uint16 = 0x03       /* rfkill master switch, type "any"
uint16 =  uint16 = 	 set = radio enabled */
const SW_RADIO uint16 = SW_RFKILL_ALL       /* deprecated */
const SW_MICROPHONE_INSERT uint16 = 0x04    /* set = inserted */
const SW_DOCK uint16 = 0x05                 /* set = plugged into dock */
const SW_LINEOUT_INSERT uint16 = 0x06       /* set = inserted */
const SW_JACK_PHYSICAL_INSERT uint16 = 0x07 /* set = mechanical switch set */
const SW_VIDEOOUT_INSERT uint16 = 0x08      /* set = inserted */
const SW_CAMERA_LENS_COVER uint16 = 0x09    /* set = lens covered */
const SW_KEYPAD_SLIDE uint16 = 0x0a         /* set = keypad slide out */
const SW_FRONT_PROXIMITY uint16 = 0x0b      /* set = front proximity sensor active */
const SW_ROTATE_LOCK uint16 = 0x0c          /* set = rotate locked/disabled */
const SW_LINEIN_INSERT uint16 = 0x0d        /* set = inserted */
const SW_MUTE_DEVICE uint16 = 0x0e          /* set = device disabled */
const SW_PEN_INSERTED uint16 = 0x0f         /* set = pen inserted */
const SW_MACHINE_COVER uint16 = 0x10        /* set = cover closed */
const SW_MAX uint16 = 0x10
const SW_CNT uint16 = (SW_MAX + 1)

/*
 * Misc events
 */

const MSC_SERIAL uint16 = 0x00
const MSC_PULSELED uint16 = 0x01
const MSC_GESTURE uint16 = 0x02
const MSC_RAW uint16 = 0x03
const MSC_SCAN uint16 = 0x04
const MSC_TIMESTAMP uint16 = 0x05
const MSC_MAX uint16 = 0x07
const MSC_CNT uint16 = (MSC_MAX + 1)

/*
 * LEDs
 */

const LED_NUML uint16 = 0x00
const LED_CAPSL uint16 = 0x01
const LED_SCROLLL uint16 = 0x02
const LED_COMPOSE uint16 = 0x03
const LED_KANA uint16 = 0x04
const LED_SLEEP uint16 = 0x05
const LED_SUSPEND uint16 = 0x06
const LED_MUTE uint16 = 0x07
const LED_MISC uint16 = 0x08
const LED_MAIL uint16 = 0x09
const LED_CHARGING uint16 = 0x0a
const LED_MAX uint16 = 0x0f
const LED_CNT uint16 = (LED_MAX + 1)

/*
 * Autorepeat values
 */

const REP_DELAY uint16 = 0x00
const REP_PERIOD uint16 = 0x01
const REP_MAX uint16 = 0x01
const REP_CNT uint16 = (REP_MAX + 1)

/*
 * Sounds
 */

const SND_CLICK uint16 = 0x00
const SND_BELL uint16 = 0x01
const SND_TONE uint16 = 0x02
const SND_MAX uint16 = 0x07
const SND_CNT uint16 = (SND_MAX + 1)
