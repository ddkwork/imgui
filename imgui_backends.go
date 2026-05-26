package imgui

import (
	"fmt"
	"unsafe"
)

// Source: imgui_backends.h
type (
	ImGuiID                = uint32
	ImS8                   = int8
	ImU8                   = uint8
	ImS16                  = int16
	ImU16                  = uint16
	ImS32                  = int32
	ImU32                  = uint32
	ImS64                  = int64
	ImU64                  = uint64
	ImGuiCol               = int32
	ImGuiCond              = int32
	ImGuiDataType          = int32
	ImGuiMouseButton       = int32
	ImGuiMouseCursor       = int32
	ImGuiStyleVar          = int32
	ImGuiTableBgTarget     = int32
	ImDrawFlags            = int32
	ImDrawListFlags        = int32
	ImDrawTextFlags        = int32
	ImFontFlags            = int32
	ImFontAtlasFlags       = int32
	ImGuiBackendFlags      = int32
	ImGuiButtonFlags       = int32
	ImGuiChildFlags        = int32
	ImGuiColorEditFlags    = int32
	ImGuiConfigFlags       = int32
	ImGuiComboFlags        = int32
	ImGuiDragDropFlags     = int32
	ImGuiFocusedFlags      = int32
	ImGuiHoveredFlags      = int32
	ImGuiInputFlags        = int32
	ImGuiInputTextFlags    = int32
	ImGuiItemFlags         = int32
	ImGuiKeyChord          = int32
	ImGuiListClipperFlags  = int32
	ImGuiPopupFlags        = int32
	ImGuiMultiSelectFlags  = int32
	ImGuiSelectableFlags   = int32
	ImGuiSliderFlags       = int32
	ImGuiTabBarFlags       = int32
	ImGuiTabItemFlags      = int32
	ImGuiTableFlags        = int32
	ImGuiTableColumnFlags  = int32
	ImGuiTableRowFlags     = int32
	ImGuiTreeNodeFlags     = int32
	ImGuiViewportFlags     = int32
	ImGuiWindowFlags       = int32
	ImWchar32              = uint32
	ImWchar16              = uint16
	ImWchar                = uint16
	ImGuiSelectionUserData = int64
	ImTextureID            = uint64
	ImDrawIdx              = uint16
	ImFontAtlasRectId      = int32
	ImFontAtlasCustomRect  = ImFontAtlasRect
)

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiDir
type ImGuiDir int32

const (
	ImguidirNone ImGuiDir = -1 + iota
	ImguidirLeft
	ImguidirRight
	ImguidirUp
	ImguidirDown
	ImguidirCount
)

func (i ImGuiDir) String() string {
	switch i {
	case ImguidirNone:
		return "Imguidir None"
	case ImguidirLeft:
		return "Imguidir Left"
	case ImguidirRight:
		return "Imguidir Right"
	case ImguidirUp:
		return "Imguidir Up"
	case ImguidirDown:
		return "Imguidir Down"
	case ImguidirCount:
		return "Imguidir Count"
	default:
		return fmt.Sprintf("ImGuiDir(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiKey
type ImGuiKey uint32

const (
	ImguikeyNone                ImGuiKey = 0
	ImguikeyNamedkeyBegin       ImGuiKey = 512
	ImguikeyTab                 ImGuiKey = 512
	ImguikeyLeftarrow           ImGuiKey = 513
	ImguikeyRightarrow          ImGuiKey = 514
	ImguikeyUparrow             ImGuiKey = 515
	ImguikeyDownarrow           ImGuiKey = 516
	ImguikeyPageup              ImGuiKey = 517
	ImguikeyPagedown            ImGuiKey = 518
	ImguikeyHome                ImGuiKey = 519
	ImguikeyEnd                 ImGuiKey = 520
	ImguikeyInsert              ImGuiKey = 521
	ImguikeyDelete              ImGuiKey = 522
	ImguikeyBackspace           ImGuiKey = 523
	ImguikeySpace               ImGuiKey = 524
	ImguikeyEnter               ImGuiKey = 525
	ImguikeyEscape              ImGuiKey = 526
	ImguikeyLeftctrl            ImGuiKey = 527
	ImguikeyLeftshift           ImGuiKey = 528
	ImguikeyLeftalt             ImGuiKey = 529
	ImguikeyLeftsuper           ImGuiKey = 530
	ImguikeyRightctrl           ImGuiKey = 531
	ImguikeyRightshift          ImGuiKey = 532
	ImguikeyRightalt            ImGuiKey = 533
	ImguikeyRightsuper          ImGuiKey = 534
	ImguikeyMenu                ImGuiKey = 535
	Imguikey0                   ImGuiKey = 536
	Imguikey1                   ImGuiKey = 537
	Imguikey2                   ImGuiKey = 538
	Imguikey3                   ImGuiKey = 539
	Imguikey4                   ImGuiKey = 540
	Imguikey5                   ImGuiKey = 541
	Imguikey6                   ImGuiKey = 542
	Imguikey7                   ImGuiKey = 543
	Imguikey8                   ImGuiKey = 544
	Imguikey9                   ImGuiKey = 545
	ImguikeyA                   ImGuiKey = 546
	ImguikeyB                   ImGuiKey = 547
	ImguikeyC                   ImGuiKey = 548
	ImguikeyD                   ImGuiKey = 549
	ImguikeyE                   ImGuiKey = 550
	ImguikeyF                   ImGuiKey = 551
	ImguikeyG                   ImGuiKey = 552
	ImguikeyH                   ImGuiKey = 553
	ImguikeyI                   ImGuiKey = 554
	ImguikeyJ                   ImGuiKey = 555
	ImguikeyK                   ImGuiKey = 556
	ImguikeyL                   ImGuiKey = 557
	ImguikeyM                   ImGuiKey = 558
	ImguikeyN                   ImGuiKey = 559
	ImguikeyO                   ImGuiKey = 560
	ImguikeyP                   ImGuiKey = 561
	ImguikeyQ                   ImGuiKey = 562
	ImguikeyR                   ImGuiKey = 563
	ImguikeyS                   ImGuiKey = 564
	ImguikeyT                   ImGuiKey = 565
	ImguikeyU                   ImGuiKey = 566
	ImguikeyV                   ImGuiKey = 567
	ImguikeyW                   ImGuiKey = 568
	ImguikeyX                   ImGuiKey = 569
	ImguikeyY                   ImGuiKey = 570
	ImguikeyZ                   ImGuiKey = 571
	ImguikeyF1                  ImGuiKey = 572
	ImguikeyF2                  ImGuiKey = 573
	ImguikeyF3                  ImGuiKey = 574
	ImguikeyF4                  ImGuiKey = 575
	ImguikeyF5                  ImGuiKey = 576
	ImguikeyF6                  ImGuiKey = 577
	ImguikeyF7                  ImGuiKey = 578
	ImguikeyF8                  ImGuiKey = 579
	ImguikeyF9                  ImGuiKey = 580
	ImguikeyF10                 ImGuiKey = 581
	ImguikeyF11                 ImGuiKey = 582
	ImguikeyF12                 ImGuiKey = 583
	ImguikeyF13                 ImGuiKey = 584
	ImguikeyF14                 ImGuiKey = 585
	ImguikeyF15                 ImGuiKey = 586
	ImguikeyF16                 ImGuiKey = 587
	ImguikeyF17                 ImGuiKey = 588
	ImguikeyF18                 ImGuiKey = 589
	ImguikeyF19                 ImGuiKey = 590
	ImguikeyF20                 ImGuiKey = 591
	ImguikeyF21                 ImGuiKey = 592
	ImguikeyF22                 ImGuiKey = 593
	ImguikeyF23                 ImGuiKey = 594
	ImguikeyF24                 ImGuiKey = 595
	ImguikeyApostrophe          ImGuiKey = 596
	ImguikeyComma               ImGuiKey = 597
	ImguikeyMinus               ImGuiKey = 598
	ImguikeyPeriod              ImGuiKey = 599
	ImguikeySlash               ImGuiKey = 600
	ImguikeySemicolon           ImGuiKey = 601
	ImguikeyEqual               ImGuiKey = 602
	ImguikeyLeftbracket         ImGuiKey = 603
	ImguikeyBackslash           ImGuiKey = 604
	ImguikeyRightbracket        ImGuiKey = 605
	ImguikeyGraveaccent         ImGuiKey = 606
	ImguikeyCapslock            ImGuiKey = 607
	ImguikeyScrolllock          ImGuiKey = 608
	ImguikeyNumlock             ImGuiKey = 609
	ImguikeyPrintscreen         ImGuiKey = 610
	ImguikeyPause               ImGuiKey = 611
	ImguikeyKeypad0             ImGuiKey = 612
	ImguikeyKeypad1             ImGuiKey = 613
	ImguikeyKeypad2             ImGuiKey = 614
	ImguikeyKeypad3             ImGuiKey = 615
	ImguikeyKeypad4             ImGuiKey = 616
	ImguikeyKeypad5             ImGuiKey = 617
	ImguikeyKeypad6             ImGuiKey = 618
	ImguikeyKeypad7             ImGuiKey = 619
	ImguikeyKeypad8             ImGuiKey = 620
	ImguikeyKeypad9             ImGuiKey = 621
	ImguikeyKeypaddecimal       ImGuiKey = 622
	ImguikeyKeypaddivide        ImGuiKey = 623
	ImguikeyKeypadmultiply      ImGuiKey = 624
	ImguikeyKeypadsubtract      ImGuiKey = 625
	ImguikeyKeypadadd           ImGuiKey = 626
	ImguikeyKeypadenter         ImGuiKey = 627
	ImguikeyKeypadequal         ImGuiKey = 628
	ImguikeyAppback             ImGuiKey = 629
	ImguikeyAppforward          ImGuiKey = 630
	ImguikeyOem102              ImGuiKey = 631
	ImguikeyGamepadstart        ImGuiKey = 632
	ImguikeyGamepadback         ImGuiKey = 633
	ImguikeyGamepadfaceleft     ImGuiKey = 634
	ImguikeyGamepadfaceright    ImGuiKey = 635
	ImguikeyGamepadfaceup       ImGuiKey = 636
	ImguikeyGamepadfacedown     ImGuiKey = 637
	ImguikeyGamepaddpadleft     ImGuiKey = 638
	ImguikeyGamepaddpadright    ImGuiKey = 639
	ImguikeyGamepaddpadup       ImGuiKey = 640
	ImguikeyGamepaddpaddown     ImGuiKey = 641
	ImguikeyGamepadl1           ImGuiKey = 642
	ImguikeyGamepadr1           ImGuiKey = 643
	ImguikeyGamepadl2           ImGuiKey = 644
	ImguikeyGamepadr2           ImGuiKey = 645
	ImguikeyGamepadl3           ImGuiKey = 646
	ImguikeyGamepadr3           ImGuiKey = 647
	ImguikeyGamepadlstickleft   ImGuiKey = 648
	ImguikeyGamepadlstickright  ImGuiKey = 649
	ImguikeyGamepadlstickup     ImGuiKey = 650
	ImguikeyGamepadlstickdown   ImGuiKey = 651
	ImguikeyGamepadrstickleft   ImGuiKey = 652
	ImguikeyGamepadrstickright  ImGuiKey = 653
	ImguikeyGamepadrstickup     ImGuiKey = 654
	ImguikeyGamepadrstickdown   ImGuiKey = 655
	ImguikeyMouseleft           ImGuiKey = 656
	ImguikeyMouseright          ImGuiKey = 657
	ImguikeyMousemiddle         ImGuiKey = 658
	ImguikeyMousex1             ImGuiKey = 659
	ImguikeyMousex2             ImGuiKey = 660
	ImguikeyMousewheelx         ImGuiKey = 661
	ImguikeyMousewheely         ImGuiKey = 662
	ImguikeyReservedformodctrl  ImGuiKey = 663
	ImguikeyReservedformodshift ImGuiKey = 664
	ImguikeyReservedformodalt   ImGuiKey = 665
	ImguikeyReservedformodsuper ImGuiKey = 666
	ImguikeyNamedkeyEnd         ImGuiKey = 667
	ImguikeyNamedkeyCount       ImGuiKey = 155
	ImguimodNone                ImGuiKey = 0
	ImguimodCtrl                ImGuiKey = 4096
	ImguimodShift               ImGuiKey = 8192
	ImguimodAlt                 ImGuiKey = 16384
	ImguimodSuper               ImGuiKey = 32768
	ImguimodMask                ImGuiKey = 61440
	ImguikeyCount               ImGuiKey = 667
	ImguimodShortcut            ImGuiKey = 4096
)

func (i ImGuiKey) String() string {
	switch i {
	case ImguikeyNone:
		return "Imguikey None"
	case ImguikeyNamedkeyBegin:
		return "Imguikey Namedkey Begin"
	case ImguikeyLeftarrow:
		return "Imguikey Leftarrow"
	case ImguikeyRightarrow:
		return "Imguikey Rightarrow"
	case ImguikeyUparrow:
		return "Imguikey Uparrow"
	case ImguikeyDownarrow:
		return "Imguikey Downarrow"
	case ImguikeyPageup:
		return "Imguikey Pageup"
	case ImguikeyPagedown:
		return "Imguikey Pagedown"
	case ImguikeyHome:
		return "Imguikey Home"
	case ImguikeyEnd:
		return "Imguikey End"
	case ImguikeyInsert:
		return "Imguikey Insert"
	case ImguikeyDelete:
		return "Imguikey Delete"
	case ImguikeyBackspace:
		return "Imguikey Backspace"
	case ImguikeySpace:
		return "Imguikey Space"
	case ImguikeyEnter:
		return "Imguikey Enter"
	case ImguikeyEscape:
		return "Imguikey Escape"
	case ImguikeyLeftctrl:
		return "Imguikey Leftctrl"
	case ImguikeyLeftshift:
		return "Imguikey Leftshift"
	case ImguikeyLeftalt:
		return "Imguikey Leftalt"
	case ImguikeyLeftsuper:
		return "Imguikey Leftsuper"
	case ImguikeyRightctrl:
		return "Imguikey Rightctrl"
	case ImguikeyRightshift:
		return "Imguikey Rightshift"
	case ImguikeyRightalt:
		return "Imguikey Rightalt"
	case ImguikeyRightsuper:
		return "Imguikey Rightsuper"
	case ImguikeyMenu:
		return "Imguikey Menu"
	case Imguikey0:
		return "Imguikey 0"
	case Imguikey1:
		return "Imguikey 1"
	case Imguikey2:
		return "Imguikey 2"
	case Imguikey3:
		return "Imguikey 3"
	case Imguikey4:
		return "Imguikey 4"
	case Imguikey5:
		return "Imguikey 5"
	case Imguikey6:
		return "Imguikey 6"
	case Imguikey7:
		return "Imguikey 7"
	case Imguikey8:
		return "Imguikey 8"
	case Imguikey9:
		return "Imguikey 9"
	case ImguikeyA:
		return "Imguikey A"
	case ImguikeyB:
		return "Imguikey B"
	case ImguikeyC:
		return "Imguikey C"
	case ImguikeyD:
		return "Imguikey D"
	case ImguikeyE:
		return "Imguikey E"
	case ImguikeyF:
		return "Imguikey F"
	case ImguikeyG:
		return "Imguikey G"
	case ImguikeyH:
		return "Imguikey H"
	case ImguikeyI:
		return "Imguikey I"
	case ImguikeyJ:
		return "Imguikey J"
	case ImguikeyK:
		return "Imguikey K"
	case ImguikeyL:
		return "Imguikey L"
	case ImguikeyM:
		return "Imguikey M"
	case ImguikeyN:
		return "Imguikey N"
	case ImguikeyO:
		return "Imguikey O"
	case ImguikeyP:
		return "Imguikey P"
	case ImguikeyQ:
		return "Imguikey Q"
	case ImguikeyR:
		return "Imguikey R"
	case ImguikeyS:
		return "Imguikey S"
	case ImguikeyT:
		return "Imguikey T"
	case ImguikeyU:
		return "Imguikey U"
	case ImguikeyV:
		return "Imguikey V"
	case ImguikeyW:
		return "Imguikey W"
	case ImguikeyX:
		return "Imguikey X"
	case ImguikeyY:
		return "Imguikey Y"
	case ImguikeyZ:
		return "Imguikey Z"
	case ImguikeyF1:
		return "Imguikey F1"
	case ImguikeyF2:
		return "Imguikey F2"
	case ImguikeyF3:
		return "Imguikey F3"
	case ImguikeyF4:
		return "Imguikey F4"
	case ImguikeyF5:
		return "Imguikey F5"
	case ImguikeyF6:
		return "Imguikey F6"
	case ImguikeyF7:
		return "Imguikey F7"
	case ImguikeyF8:
		return "Imguikey F8"
	case ImguikeyF9:
		return "Imguikey F9"
	case ImguikeyF10:
		return "Imguikey F10"
	case ImguikeyF11:
		return "Imguikey F11"
	case ImguikeyF12:
		return "Imguikey F12"
	case ImguikeyF13:
		return "Imguikey F13"
	case ImguikeyF14:
		return "Imguikey F14"
	case ImguikeyF15:
		return "Imguikey F15"
	case ImguikeyF16:
		return "Imguikey F16"
	case ImguikeyF17:
		return "Imguikey F17"
	case ImguikeyF18:
		return "Imguikey F18"
	case ImguikeyF19:
		return "Imguikey F19"
	case ImguikeyF20:
		return "Imguikey F20"
	case ImguikeyF21:
		return "Imguikey F21"
	case ImguikeyF22:
		return "Imguikey F22"
	case ImguikeyF23:
		return "Imguikey F23"
	case ImguikeyF24:
		return "Imguikey F24"
	case ImguikeyApostrophe:
		return "Imguikey Apostrophe"
	case ImguikeyComma:
		return "Imguikey Comma"
	case ImguikeyMinus:
		return "Imguikey Minus"
	case ImguikeyPeriod:
		return "Imguikey Period"
	case ImguikeySlash:
		return "Imguikey Slash"
	case ImguikeySemicolon:
		return "Imguikey Semicolon"
	case ImguikeyEqual:
		return "Imguikey Equal"
	case ImguikeyLeftbracket:
		return "Imguikey Leftbracket"
	case ImguikeyBackslash:
		return "Imguikey Backslash"
	case ImguikeyRightbracket:
		return "Imguikey Rightbracket"
	case ImguikeyGraveaccent:
		return "Imguikey Graveaccent"
	case ImguikeyCapslock:
		return "Imguikey Capslock"
	case ImguikeyScrolllock:
		return "Imguikey Scrolllock"
	case ImguikeyNumlock:
		return "Imguikey Numlock"
	case ImguikeyPrintscreen:
		return "Imguikey Printscreen"
	case ImguikeyPause:
		return "Imguikey Pause"
	case ImguikeyKeypad0:
		return "Imguikey Keypad 0"
	case ImguikeyKeypad1:
		return "Imguikey Keypad 1"
	case ImguikeyKeypad2:
		return "Imguikey Keypad 2"
	case ImguikeyKeypad3:
		return "Imguikey Keypad 3"
	case ImguikeyKeypad4:
		return "Imguikey Keypad 4"
	case ImguikeyKeypad5:
		return "Imguikey Keypad 5"
	case ImguikeyKeypad6:
		return "Imguikey Keypad 6"
	case ImguikeyKeypad7:
		return "Imguikey Keypad 7"
	case ImguikeyKeypad8:
		return "Imguikey Keypad 8"
	case ImguikeyKeypad9:
		return "Imguikey Keypad 9"
	case ImguikeyKeypaddecimal:
		return "Imguikey Keypaddecimal"
	case ImguikeyKeypaddivide:
		return "Imguikey Keypaddivide"
	case ImguikeyKeypadmultiply:
		return "Imguikey Keypadmultiply"
	case ImguikeyKeypadsubtract:
		return "Imguikey Keypadsubtract"
	case ImguikeyKeypadadd:
		return "Imguikey Keypadadd"
	case ImguikeyKeypadenter:
		return "Imguikey Keypadenter"
	case ImguikeyKeypadequal:
		return "Imguikey Keypadequal"
	case ImguikeyAppback:
		return "Imguikey Appback"
	case ImguikeyAppforward:
		return "Imguikey Appforward"
	case ImguikeyOem102:
		return "Imguikey Oem 102"
	case ImguikeyGamepadstart:
		return "Imguikey Gamepadstart"
	case ImguikeyGamepadback:
		return "Imguikey Gamepadback"
	case ImguikeyGamepadfaceleft:
		return "Imguikey Gamepadfaceleft"
	case ImguikeyGamepadfaceright:
		return "Imguikey Gamepadfaceright"
	case ImguikeyGamepadfaceup:
		return "Imguikey Gamepadfaceup"
	case ImguikeyGamepadfacedown:
		return "Imguikey Gamepadfacedown"
	case ImguikeyGamepaddpadleft:
		return "Imguikey Gamepaddpadleft"
	case ImguikeyGamepaddpadright:
		return "Imguikey Gamepaddpadright"
	case ImguikeyGamepaddpadup:
		return "Imguikey Gamepaddpadup"
	case ImguikeyGamepaddpaddown:
		return "Imguikey Gamepaddpaddown"
	case ImguikeyGamepadl1:
		return "Imguikey Gamepadl 1"
	case ImguikeyGamepadr1:
		return "Imguikey Gamepadr 1"
	case ImguikeyGamepadl2:
		return "Imguikey Gamepadl 2"
	case ImguikeyGamepadr2:
		return "Imguikey Gamepadr 2"
	case ImguikeyGamepadl3:
		return "Imguikey Gamepadl 3"
	case ImguikeyGamepadr3:
		return "Imguikey Gamepadr 3"
	case ImguikeyGamepadlstickleft:
		return "Imguikey Gamepadlstickleft"
	case ImguikeyGamepadlstickright:
		return "Imguikey Gamepadlstickright"
	case ImguikeyGamepadlstickup:
		return "Imguikey Gamepadlstickup"
	case ImguikeyGamepadlstickdown:
		return "Imguikey Gamepadlstickdown"
	case ImguikeyGamepadrstickleft:
		return "Imguikey Gamepadrstickleft"
	case ImguikeyGamepadrstickright:
		return "Imguikey Gamepadrstickright"
	case ImguikeyGamepadrstickup:
		return "Imguikey Gamepadrstickup"
	case ImguikeyGamepadrstickdown:
		return "Imguikey Gamepadrstickdown"
	case ImguikeyMouseleft:
		return "Imguikey Mouseleft"
	case ImguikeyMouseright:
		return "Imguikey Mouseright"
	case ImguikeyMousemiddle:
		return "Imguikey Mousemiddle"
	case ImguikeyMousex1:
		return "Imguikey Mousex 1"
	case ImguikeyMousex2:
		return "Imguikey Mousex 2"
	case ImguikeyMousewheelx:
		return "Imguikey Mousewheelx"
	case ImguikeyMousewheely:
		return "Imguikey Mousewheely"
	case ImguikeyReservedformodctrl:
		return "Imguikey Reservedformodctrl"
	case ImguikeyReservedformodshift:
		return "Imguikey Reservedformodshift"
	case ImguikeyReservedformodalt:
		return "Imguikey Reservedformodalt"
	case ImguikeyReservedformodsuper:
		return "Imguikey Reservedformodsuper"
	case ImguikeyNamedkeyEnd:
		return "Imguikey Namedkey End"
	case ImguikeyNamedkeyCount:
		return "Imguikey Namedkey Count"
	case ImguimodCtrl:
		return "Imguimod Ctrl"
	case ImguimodShift:
		return "Imguimod Shift"
	case ImguimodAlt:
		return "Imguimod Alt"
	case ImguimodSuper:
		return "Imguimod Super"
	case ImguimodMask:
		return "Imguimod Mask"
	default:
		return fmt.Sprintf("ImGuiKey(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiMouseSource
type ImGuiMouseSource uint32

const (
	ImguimousesourceMouse ImGuiMouseSource = iota
	ImguimousesourceTouchscreen
	ImguimousesourcePen
	ImguimousesourceCount
)

func (i ImGuiMouseSource) String() string {
	switch i {
	case ImguimousesourceMouse:
		return "Imguimousesource Mouse"
	case ImguimousesourceTouchscreen:
		return "Imguimousesource Touchscreen"
	case ImguimousesourcePen:
		return "Imguimousesource Pen"
	case ImguimousesourceCount:
		return "Imguimousesource Count"
	default:
		return fmt.Sprintf("ImGuiMouseSource(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiSortDirection
type ImGuiSortDirection uint32

const (
	ImguisortdirectionNone ImGuiSortDirection = iota
	ImguisortdirectionAscending
	ImguisortdirectionDescending
)

func (i ImGuiSortDirection) String() string {
	switch i {
	case ImguisortdirectionNone:
		return "Imguisortdirection None"
	case ImguisortdirectionAscending:
		return "Imguisortdirection Ascending"
	case ImguisortdirectionDescending:
		return "Imguisortdirection Descending"
	default:
		return fmt.Sprintf("ImGuiSortDirection(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiWindowFlags_
type ImGuiWindowFlags_ uint32

const (
	ImguiwindowflagsNone                      ImGuiWindowFlags_ = 0
	ImguiwindowflagsNotitlebar                ImGuiWindowFlags_ = 1
	ImguiwindowflagsNoresize                  ImGuiWindowFlags_ = 2
	ImguiwindowflagsNomove                    ImGuiWindowFlags_ = 4
	ImguiwindowflagsNoscrollbar               ImGuiWindowFlags_ = 8
	ImguiwindowflagsNoscrollwithmouse         ImGuiWindowFlags_ = 16
	ImguiwindowflagsNocollapse                ImGuiWindowFlags_ = 32
	ImguiwindowflagsAlwaysautoresize          ImGuiWindowFlags_ = 64
	ImguiwindowflagsNobackground              ImGuiWindowFlags_ = 128
	ImguiwindowflagsNosavedsettings           ImGuiWindowFlags_ = 256
	ImguiwindowflagsNomouseinputs             ImGuiWindowFlags_ = 512
	ImguiwindowflagsMenubar                   ImGuiWindowFlags_ = 1024
	ImguiwindowflagsHorizontalscrollbar       ImGuiWindowFlags_ = 2048
	ImguiwindowflagsNofocusonappearing        ImGuiWindowFlags_ = 4096
	ImguiwindowflagsNobringtofrontonfocus     ImGuiWindowFlags_ = 8192
	ImguiwindowflagsAlwaysverticalscrollbar   ImGuiWindowFlags_ = 16384
	ImguiwindowflagsAlwayshorizontalscrollbar ImGuiWindowFlags_ = 32768
	ImguiwindowflagsNonavinputs               ImGuiWindowFlags_ = 65536
	ImguiwindowflagsNonavfocus                ImGuiWindowFlags_ = 131072
	ImguiwindowflagsUnsaveddocument           ImGuiWindowFlags_ = 262144
	ImguiwindowflagsNonav                     ImGuiWindowFlags_ = 196608
	ImguiwindowflagsNodecoration              ImGuiWindowFlags_ = 43
	ImguiwindowflagsNoinputs                  ImGuiWindowFlags_ = 197120
	ImguiwindowflagsChildwindow               ImGuiWindowFlags_ = 16777216
	ImguiwindowflagsTooltip                   ImGuiWindowFlags_ = 33554432
	ImguiwindowflagsPopup                     ImGuiWindowFlags_ = 67108864
	ImguiwindowflagsModal                     ImGuiWindowFlags_ = 134217728
	ImguiwindowflagsChildmenu                 ImGuiWindowFlags_ = 268435456
)

func (i ImGuiWindowFlags_) String() string {
	switch i {
	case ImguiwindowflagsNone:
		return "Imguiwindowflags None"
	case ImguiwindowflagsNotitlebar:
		return "Imguiwindowflags Notitlebar"
	case ImguiwindowflagsNoresize:
		return "Imguiwindowflags Noresize"
	case ImguiwindowflagsNomove:
		return "Imguiwindowflags Nomove"
	case ImguiwindowflagsNoscrollbar:
		return "Imguiwindowflags Noscrollbar"
	case ImguiwindowflagsNoscrollwithmouse:
		return "Imguiwindowflags Noscrollwithmouse"
	case ImguiwindowflagsNocollapse:
		return "Imguiwindowflags Nocollapse"
	case ImguiwindowflagsAlwaysautoresize:
		return "Imguiwindowflags Alwaysautoresize"
	case ImguiwindowflagsNobackground:
		return "Imguiwindowflags Nobackground"
	case ImguiwindowflagsNosavedsettings:
		return "Imguiwindowflags Nosavedsettings"
	case ImguiwindowflagsNomouseinputs:
		return "Imguiwindowflags Nomouseinputs"
	case ImguiwindowflagsMenubar:
		return "Imguiwindowflags Menubar"
	case ImguiwindowflagsHorizontalscrollbar:
		return "Imguiwindowflags Horizontalscrollbar"
	case ImguiwindowflagsNofocusonappearing:
		return "Imguiwindowflags Nofocusonappearing"
	case ImguiwindowflagsNobringtofrontonfocus:
		return "Imguiwindowflags Nobringtofrontonfocus"
	case ImguiwindowflagsAlwaysverticalscrollbar:
		return "Imguiwindowflags Alwaysverticalscrollbar"
	case ImguiwindowflagsAlwayshorizontalscrollbar:
		return "Imguiwindowflags Alwayshorizontalscrollbar"
	case ImguiwindowflagsNonavinputs:
		return "Imguiwindowflags Nonavinputs"
	case ImguiwindowflagsNonavfocus:
		return "Imguiwindowflags Nonavfocus"
	case ImguiwindowflagsUnsaveddocument:
		return "Imguiwindowflags Unsaveddocument"
	case ImguiwindowflagsNonav:
		return "Imguiwindowflags Nonav"
	case ImguiwindowflagsNodecoration:
		return "Imguiwindowflags Nodecoration"
	case ImguiwindowflagsNoinputs:
		return "Imguiwindowflags Noinputs"
	case ImguiwindowflagsChildwindow:
		return "Imguiwindowflags Childwindow"
	case ImguiwindowflagsTooltip:
		return "Imguiwindowflags Tooltip"
	case ImguiwindowflagsPopup:
		return "Imguiwindowflags Popup"
	case ImguiwindowflagsModal:
		return "Imguiwindowflags Modal"
	case ImguiwindowflagsChildmenu:
		return "Imguiwindowflags Childmenu"
	default:
		return fmt.Sprintf("ImGuiWindowFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiChildFlags_
type ImGuiChildFlags_ uint32

const (
	ImguichildflagsNone                   ImGuiChildFlags_ = 0
	ImguichildflagsBorders                ImGuiChildFlags_ = 1
	ImguichildflagsAlwaysusewindowpadding ImGuiChildFlags_ = 2
	ImguichildflagsResizex                ImGuiChildFlags_ = 4
	ImguichildflagsResizey                ImGuiChildFlags_ = 8
	ImguichildflagsAutoresizex            ImGuiChildFlags_ = 16
	ImguichildflagsAutoresizey            ImGuiChildFlags_ = 32
	ImguichildflagsAlwaysautoresize       ImGuiChildFlags_ = 64
	ImguichildflagsFramestyle             ImGuiChildFlags_ = 128
	ImguichildflagsNavflattened           ImGuiChildFlags_ = 256
)

func (i ImGuiChildFlags_) String() string {
	switch i {
	case ImguichildflagsNone:
		return "Imguichildflags None"
	case ImguichildflagsBorders:
		return "Imguichildflags Borders"
	case ImguichildflagsAlwaysusewindowpadding:
		return "Imguichildflags Alwaysusewindowpadding"
	case ImguichildflagsResizex:
		return "Imguichildflags Resizex"
	case ImguichildflagsResizey:
		return "Imguichildflags Resizey"
	case ImguichildflagsAutoresizex:
		return "Imguichildflags Autoresizex"
	case ImguichildflagsAutoresizey:
		return "Imguichildflags Autoresizey"
	case ImguichildflagsAlwaysautoresize:
		return "Imguichildflags Alwaysautoresize"
	case ImguichildflagsFramestyle:
		return "Imguichildflags Framestyle"
	case ImguichildflagsNavflattened:
		return "Imguichildflags Navflattened"
	default:
		return fmt.Sprintf("ImGuiChildFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiItemFlags_
type ImGuiItemFlags_ uint32

const (
	ImguiitemflagsNone              ImGuiItemFlags_ = 0
	ImguiitemflagsNotabstop         ImGuiItemFlags_ = 1
	ImguiitemflagsNonav             ImGuiItemFlags_ = 2
	ImguiitemflagsNonavdefaultfocus ImGuiItemFlags_ = 4
	ImguiitemflagsButtonrepeat      ImGuiItemFlags_ = 8
	ImguiitemflagsAutoclosepopups   ImGuiItemFlags_ = 16
	ImguiitemflagsAllowduplicateid  ImGuiItemFlags_ = 32
	ImguiitemflagsDisabled          ImGuiItemFlags_ = 64
)

func (i ImGuiItemFlags_) String() string {
	switch i {
	case ImguiitemflagsNone:
		return "Imguiitemflags None"
	case ImguiitemflagsNotabstop:
		return "Imguiitemflags Notabstop"
	case ImguiitemflagsNonav:
		return "Imguiitemflags Nonav"
	case ImguiitemflagsNonavdefaultfocus:
		return "Imguiitemflags Nonavdefaultfocus"
	case ImguiitemflagsButtonrepeat:
		return "Imguiitemflags Buttonrepeat"
	case ImguiitemflagsAutoclosepopups:
		return "Imguiitemflags Autoclosepopups"
	case ImguiitemflagsAllowduplicateid:
		return "Imguiitemflags Allowduplicateid"
	case ImguiitemflagsDisabled:
		return "Imguiitemflags Disabled"
	default:
		return fmt.Sprintf("ImGuiItemFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiInputTextFlags_
type ImGuiInputTextFlags_ uint32

const (
	ImguiinputtextflagsNone                ImGuiInputTextFlags_ = 0
	ImguiinputtextflagsCharsdecimal        ImGuiInputTextFlags_ = 1
	ImguiinputtextflagsCharshexadecimal    ImGuiInputTextFlags_ = 2
	ImguiinputtextflagsCharsscientific     ImGuiInputTextFlags_ = 4
	ImguiinputtextflagsCharsuppercase      ImGuiInputTextFlags_ = 8
	ImguiinputtextflagsCharsnoblank        ImGuiInputTextFlags_ = 16
	ImguiinputtextflagsAllowtabinput       ImGuiInputTextFlags_ = 32
	ImguiinputtextflagsEnterreturnstrue    ImGuiInputTextFlags_ = 64
	ImguiinputtextflagsEscapeclearsall     ImGuiInputTextFlags_ = 128
	ImguiinputtextflagsCtrlenterfornewline ImGuiInputTextFlags_ = 256
	ImguiinputtextflagsReadonly            ImGuiInputTextFlags_ = 512
	ImguiinputtextflagsPassword            ImGuiInputTextFlags_ = 1024
	ImguiinputtextflagsAlwaysoverwrite     ImGuiInputTextFlags_ = 2048
	ImguiinputtextflagsAutoselectall       ImGuiInputTextFlags_ = 4096
	ImguiinputtextflagsParseemptyrefval    ImGuiInputTextFlags_ = 8192
	ImguiinputtextflagsDisplayemptyrefval  ImGuiInputTextFlags_ = 16384
	ImguiinputtextflagsNohorizontalscroll  ImGuiInputTextFlags_ = 32768
	ImguiinputtextflagsNoundoredo          ImGuiInputTextFlags_ = 65536
	ImguiinputtextflagsElideleft           ImGuiInputTextFlags_ = 131072
	ImguiinputtextflagsCallbackcompletion  ImGuiInputTextFlags_ = 262144
	ImguiinputtextflagsCallbackhistory     ImGuiInputTextFlags_ = 524288
	ImguiinputtextflagsCallbackalways      ImGuiInputTextFlags_ = 1048576
	ImguiinputtextflagsCallbackcharfilter  ImGuiInputTextFlags_ = 2097152
	ImguiinputtextflagsCallbackresize      ImGuiInputTextFlags_ = 4194304
	ImguiinputtextflagsCallbackedit        ImGuiInputTextFlags_ = 8388608
	ImguiinputtextflagsWordwrap            ImGuiInputTextFlags_ = 16777216
)

func (i ImGuiInputTextFlags_) String() string {
	switch i {
	case ImguiinputtextflagsNone:
		return "Imguiinputtextflags None"
	case ImguiinputtextflagsCharsdecimal:
		return "Imguiinputtextflags Charsdecimal"
	case ImguiinputtextflagsCharshexadecimal:
		return "Imguiinputtextflags Charshexadecimal"
	case ImguiinputtextflagsCharsscientific:
		return "Imguiinputtextflags Charsscientific"
	case ImguiinputtextflagsCharsuppercase:
		return "Imguiinputtextflags Charsuppercase"
	case ImguiinputtextflagsCharsnoblank:
		return "Imguiinputtextflags Charsnoblank"
	case ImguiinputtextflagsAllowtabinput:
		return "Imguiinputtextflags Allowtabinput"
	case ImguiinputtextflagsEnterreturnstrue:
		return "Imguiinputtextflags Enterreturnstrue"
	case ImguiinputtextflagsEscapeclearsall:
		return "Imguiinputtextflags Escapeclearsall"
	case ImguiinputtextflagsCtrlenterfornewline:
		return "Imguiinputtextflags Ctrlenterfornewline"
	case ImguiinputtextflagsReadonly:
		return "Imguiinputtextflags Readonly"
	case ImguiinputtextflagsPassword:
		return "Imguiinputtextflags Password"
	case ImguiinputtextflagsAlwaysoverwrite:
		return "Imguiinputtextflags Alwaysoverwrite"
	case ImguiinputtextflagsAutoselectall:
		return "Imguiinputtextflags Autoselectall"
	case ImguiinputtextflagsParseemptyrefval:
		return "Imguiinputtextflags Parseemptyrefval"
	case ImguiinputtextflagsDisplayemptyrefval:
		return "Imguiinputtextflags Displayemptyrefval"
	case ImguiinputtextflagsNohorizontalscroll:
		return "Imguiinputtextflags Nohorizontalscroll"
	case ImguiinputtextflagsNoundoredo:
		return "Imguiinputtextflags Noundoredo"
	case ImguiinputtextflagsElideleft:
		return "Imguiinputtextflags Elideleft"
	case ImguiinputtextflagsCallbackcompletion:
		return "Imguiinputtextflags Callbackcompletion"
	case ImguiinputtextflagsCallbackhistory:
		return "Imguiinputtextflags Callbackhistory"
	case ImguiinputtextflagsCallbackalways:
		return "Imguiinputtextflags Callbackalways"
	case ImguiinputtextflagsCallbackcharfilter:
		return "Imguiinputtextflags Callbackcharfilter"
	case ImguiinputtextflagsCallbackresize:
		return "Imguiinputtextflags Callbackresize"
	case ImguiinputtextflagsCallbackedit:
		return "Imguiinputtextflags Callbackedit"
	case ImguiinputtextflagsWordwrap:
		return "Imguiinputtextflags Wordwrap"
	default:
		return fmt.Sprintf("ImGuiInputTextFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTreeNodeFlags_
type ImGuiTreeNodeFlags_ uint32

const (
	ImguitreenodeflagsNone                 ImGuiTreeNodeFlags_ = 0
	ImguitreenodeflagsSelected             ImGuiTreeNodeFlags_ = 1
	ImguitreenodeflagsFramed               ImGuiTreeNodeFlags_ = 2
	ImguitreenodeflagsAllowoverlap         ImGuiTreeNodeFlags_ = 4
	ImguitreenodeflagsNotreepushonopen     ImGuiTreeNodeFlags_ = 8
	ImguitreenodeflagsNoautoopenonlog      ImGuiTreeNodeFlags_ = 16
	ImguitreenodeflagsDefaultopen          ImGuiTreeNodeFlags_ = 32
	ImguitreenodeflagsOpenondoubleclick    ImGuiTreeNodeFlags_ = 64
	ImguitreenodeflagsOpenonarrow          ImGuiTreeNodeFlags_ = 128
	ImguitreenodeflagsLeaf                 ImGuiTreeNodeFlags_ = 256
	ImguitreenodeflagsBullet               ImGuiTreeNodeFlags_ = 512
	ImguitreenodeflagsFramepadding         ImGuiTreeNodeFlags_ = 1024
	ImguitreenodeflagsSpanavailwidth       ImGuiTreeNodeFlags_ = 2048
	ImguitreenodeflagsSpanfullwidth        ImGuiTreeNodeFlags_ = 4096
	ImguitreenodeflagsSpanlabelwidth       ImGuiTreeNodeFlags_ = 8192
	ImguitreenodeflagsSpanallcolumns       ImGuiTreeNodeFlags_ = 16384
	ImguitreenodeflagsLabelspanallcolumns  ImGuiTreeNodeFlags_ = 32768
	ImguitreenodeflagsNavleftjumpstoparent ImGuiTreeNodeFlags_ = 131072
	ImguitreenodeflagsCollapsingheader     ImGuiTreeNodeFlags_ = 26
	ImguitreenodeflagsDrawlinesnone        ImGuiTreeNodeFlags_ = 262144
	ImguitreenodeflagsDrawlinesfull        ImGuiTreeNodeFlags_ = 524288
	ImguitreenodeflagsDrawlinestonodes     ImGuiTreeNodeFlags_ = 1048576
	ImguitreenodeflagsNavleftjumpsbackhere ImGuiTreeNodeFlags_ = 131072
	ImguitreenodeflagsSpantextwidth        ImGuiTreeNodeFlags_ = 8192
)

func (i ImGuiTreeNodeFlags_) String() string {
	switch i {
	case ImguitreenodeflagsNone:
		return "Imguitreenodeflags None"
	case ImguitreenodeflagsSelected:
		return "Imguitreenodeflags Selected"
	case ImguitreenodeflagsFramed:
		return "Imguitreenodeflags Framed"
	case ImguitreenodeflagsAllowoverlap:
		return "Imguitreenodeflags Allowoverlap"
	case ImguitreenodeflagsNotreepushonopen:
		return "Imguitreenodeflags Notreepushonopen"
	case ImguitreenodeflagsNoautoopenonlog:
		return "Imguitreenodeflags Noautoopenonlog"
	case ImguitreenodeflagsDefaultopen:
		return "Imguitreenodeflags Defaultopen"
	case ImguitreenodeflagsOpenondoubleclick:
		return "Imguitreenodeflags Openondoubleclick"
	case ImguitreenodeflagsOpenonarrow:
		return "Imguitreenodeflags Openonarrow"
	case ImguitreenodeflagsLeaf:
		return "Imguitreenodeflags Leaf"
	case ImguitreenodeflagsBullet:
		return "Imguitreenodeflags Bullet"
	case ImguitreenodeflagsFramepadding:
		return "Imguitreenodeflags Framepadding"
	case ImguitreenodeflagsSpanavailwidth:
		return "Imguitreenodeflags Spanavailwidth"
	case ImguitreenodeflagsSpanfullwidth:
		return "Imguitreenodeflags Spanfullwidth"
	case ImguitreenodeflagsSpanlabelwidth:
		return "Imguitreenodeflags Spanlabelwidth"
	case ImguitreenodeflagsSpanallcolumns:
		return "Imguitreenodeflags Spanallcolumns"
	case ImguitreenodeflagsLabelspanallcolumns:
		return "Imguitreenodeflags Labelspanallcolumns"
	case ImguitreenodeflagsNavleftjumpstoparent:
		return "Imguitreenodeflags Navleftjumpstoparent"
	case ImguitreenodeflagsCollapsingheader:
		return "Imguitreenodeflags Collapsingheader"
	case ImguitreenodeflagsDrawlinesnone:
		return "Imguitreenodeflags Drawlinesnone"
	case ImguitreenodeflagsDrawlinesfull:
		return "Imguitreenodeflags Drawlinesfull"
	case ImguitreenodeflagsDrawlinestonodes:
		return "Imguitreenodeflags Drawlinestonodes"
	default:
		return fmt.Sprintf("ImGuiTreeNodeFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiPopupFlags_
type ImGuiPopupFlags_ uint32

const (
	ImguipopupflagsNone                    ImGuiPopupFlags_ = 0
	ImguipopupflagsMousebuttonleft         ImGuiPopupFlags_ = 4
	ImguipopupflagsMousebuttonright        ImGuiPopupFlags_ = 8
	ImguipopupflagsMousebuttonmiddle       ImGuiPopupFlags_ = 12
	ImguipopupflagsNoreopen                ImGuiPopupFlags_ = 32
	ImguipopupflagsNoopenoverexistingpopup ImGuiPopupFlags_ = 128
	ImguipopupflagsNoopenoveritems         ImGuiPopupFlags_ = 256
	ImguipopupflagsAnypopupid              ImGuiPopupFlags_ = 1024
	ImguipopupflagsAnypopuplevel           ImGuiPopupFlags_ = 2048
	ImguipopupflagsAnypopup                ImGuiPopupFlags_ = 3072
	ImguipopupflagsMousebuttonshift        ImGuiPopupFlags_ = 2
	ImguipopupflagsMousebuttonmask         ImGuiPopupFlags_ = 12
	ImguipopupflagsInvalidmask             ImGuiPopupFlags_ = 3
)

func (i ImGuiPopupFlags_) String() string {
	switch i {
	case ImguipopupflagsNone:
		return "Imguipopupflags None"
	case ImguipopupflagsMousebuttonleft:
		return "Imguipopupflags Mousebuttonleft"
	case ImguipopupflagsMousebuttonright:
		return "Imguipopupflags Mousebuttonright"
	case ImguipopupflagsMousebuttonmiddle:
		return "Imguipopupflags Mousebuttonmiddle"
	case ImguipopupflagsNoreopen:
		return "Imguipopupflags Noreopen"
	case ImguipopupflagsNoopenoverexistingpopup:
		return "Imguipopupflags Noopenoverexistingpopup"
	case ImguipopupflagsNoopenoveritems:
		return "Imguipopupflags Noopenoveritems"
	case ImguipopupflagsAnypopupid:
		return "Imguipopupflags Anypopupid"
	case ImguipopupflagsAnypopuplevel:
		return "Imguipopupflags Anypopuplevel"
	case ImguipopupflagsAnypopup:
		return "Imguipopupflags Anypopup"
	case ImguipopupflagsMousebuttonshift:
		return "Imguipopupflags Mousebuttonshift"
	case ImguipopupflagsInvalidmask:
		return "Imguipopupflags Invalidmask"
	default:
		return fmt.Sprintf("ImGuiPopupFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiSelectableFlags_
type ImGuiSelectableFlags_ uint32

const (
	ImguiselectableflagsNone              ImGuiSelectableFlags_ = 0
	ImguiselectableflagsNoautoclosepopups ImGuiSelectableFlags_ = 1
	ImguiselectableflagsSpanallcolumns    ImGuiSelectableFlags_ = 2
	ImguiselectableflagsAllowdoubleclick  ImGuiSelectableFlags_ = 4
	ImguiselectableflagsDisabled          ImGuiSelectableFlags_ = 8
	ImguiselectableflagsAllowoverlap      ImGuiSelectableFlags_ = 16
	ImguiselectableflagsHighlight         ImGuiSelectableFlags_ = 32
	ImguiselectableflagsSelectonnav       ImGuiSelectableFlags_ = 64
	ImguiselectableflagsDontclosepopups   ImGuiSelectableFlags_ = 1
)

func (i ImGuiSelectableFlags_) String() string {
	switch i {
	case ImguiselectableflagsNone:
		return "Imguiselectableflags None"
	case ImguiselectableflagsNoautoclosepopups:
		return "Imguiselectableflags Noautoclosepopups"
	case ImguiselectableflagsSpanallcolumns:
		return "Imguiselectableflags Spanallcolumns"
	case ImguiselectableflagsAllowdoubleclick:
		return "Imguiselectableflags Allowdoubleclick"
	case ImguiselectableflagsDisabled:
		return "Imguiselectableflags Disabled"
	case ImguiselectableflagsAllowoverlap:
		return "Imguiselectableflags Allowoverlap"
	case ImguiselectableflagsHighlight:
		return "Imguiselectableflags Highlight"
	case ImguiselectableflagsSelectonnav:
		return "Imguiselectableflags Selectonnav"
	default:
		return fmt.Sprintf("ImGuiSelectableFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiComboFlags_
type ImGuiComboFlags_ uint32

const (
	ImguicomboflagsNone            ImGuiComboFlags_ = 0
	ImguicomboflagsPopupalignleft  ImGuiComboFlags_ = 1
	ImguicomboflagsHeightsmall     ImGuiComboFlags_ = 2
	ImguicomboflagsHeightregular   ImGuiComboFlags_ = 4
	ImguicomboflagsHeightlarge     ImGuiComboFlags_ = 8
	ImguicomboflagsHeightlargest   ImGuiComboFlags_ = 16
	ImguicomboflagsNoarrowbutton   ImGuiComboFlags_ = 32
	ImguicomboflagsNopreview       ImGuiComboFlags_ = 64
	ImguicomboflagsWidthfitpreview ImGuiComboFlags_ = 128
	ImguicomboflagsHeightmask      ImGuiComboFlags_ = 30
)

func (i ImGuiComboFlags_) String() string {
	switch i {
	case ImguicomboflagsNone:
		return "Imguicomboflags None"
	case ImguicomboflagsPopupalignleft:
		return "Imguicomboflags Popupalignleft"
	case ImguicomboflagsHeightsmall:
		return "Imguicomboflags Heightsmall"
	case ImguicomboflagsHeightregular:
		return "Imguicomboflags Heightregular"
	case ImguicomboflagsHeightlarge:
		return "Imguicomboflags Heightlarge"
	case ImguicomboflagsHeightlargest:
		return "Imguicomboflags Heightlargest"
	case ImguicomboflagsNoarrowbutton:
		return "Imguicomboflags Noarrowbutton"
	case ImguicomboflagsNopreview:
		return "Imguicomboflags Nopreview"
	case ImguicomboflagsWidthfitpreview:
		return "Imguicomboflags Widthfitpreview"
	case ImguicomboflagsHeightmask:
		return "Imguicomboflags Heightmask"
	default:
		return fmt.Sprintf("ImGuiComboFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTabBarFlags_
type ImGuiTabBarFlags_ uint32

const (
	ImguitabbarflagsNone                         ImGuiTabBarFlags_ = 0
	ImguitabbarflagsReorderable                  ImGuiTabBarFlags_ = 1
	ImguitabbarflagsAutoselectnewtabs            ImGuiTabBarFlags_ = 2
	ImguitabbarflagsTablistpopupbutton           ImGuiTabBarFlags_ = 4
	ImguitabbarflagsNoclosewithmiddlemousebutton ImGuiTabBarFlags_ = 8
	ImguitabbarflagsNotablistscrollingbuttons    ImGuiTabBarFlags_ = 16
	ImguitabbarflagsNotooltip                    ImGuiTabBarFlags_ = 32
	ImguitabbarflagsDrawselectedoverline         ImGuiTabBarFlags_ = 64
	ImguitabbarflagsFittingpolicymixed           ImGuiTabBarFlags_ = 128
	ImguitabbarflagsFittingpolicyshrink          ImGuiTabBarFlags_ = 256
	ImguitabbarflagsFittingpolicyscroll          ImGuiTabBarFlags_ = 512
	ImguitabbarflagsFittingpolicymask            ImGuiTabBarFlags_ = 896
	ImguitabbarflagsFittingpolicydefault         ImGuiTabBarFlags_ = 128
	ImguitabbarflagsFittingpolicyresizedown      ImGuiTabBarFlags_ = 256
)

func (i ImGuiTabBarFlags_) String() string {
	switch i {
	case ImguitabbarflagsNone:
		return "Imguitabbarflags None"
	case ImguitabbarflagsReorderable:
		return "Imguitabbarflags Reorderable"
	case ImguitabbarflagsAutoselectnewtabs:
		return "Imguitabbarflags Autoselectnewtabs"
	case ImguitabbarflagsTablistpopupbutton:
		return "Imguitabbarflags Tablistpopupbutton"
	case ImguitabbarflagsNoclosewithmiddlemousebutton:
		return "Imguitabbarflags Noclosewithmiddlemousebutton"
	case ImguitabbarflagsNotablistscrollingbuttons:
		return "Imguitabbarflags Notablistscrollingbuttons"
	case ImguitabbarflagsNotooltip:
		return "Imguitabbarflags Notooltip"
	case ImguitabbarflagsDrawselectedoverline:
		return "Imguitabbarflags Drawselectedoverline"
	case ImguitabbarflagsFittingpolicymixed:
		return "Imguitabbarflags Fittingpolicymixed"
	case ImguitabbarflagsFittingpolicyshrink:
		return "Imguitabbarflags Fittingpolicyshrink"
	case ImguitabbarflagsFittingpolicyscroll:
		return "Imguitabbarflags Fittingpolicyscroll"
	case ImguitabbarflagsFittingpolicymask:
		return "Imguitabbarflags Fittingpolicymask"
	default:
		return fmt.Sprintf("ImGuiTabBarFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTabItemFlags_
type ImGuiTabItemFlags_ uint32

const (
	ImguitabitemflagsNone                         ImGuiTabItemFlags_ = 0
	ImguitabitemflagsUnsaveddocument              ImGuiTabItemFlags_ = 1
	ImguitabitemflagsSetselected                  ImGuiTabItemFlags_ = 2
	ImguitabitemflagsNoclosewithmiddlemousebutton ImGuiTabItemFlags_ = 4
	ImguitabitemflagsNopushid                     ImGuiTabItemFlags_ = 8
	ImguitabitemflagsNotooltip                    ImGuiTabItemFlags_ = 16
	ImguitabitemflagsNoreorder                    ImGuiTabItemFlags_ = 32
	ImguitabitemflagsLeading                      ImGuiTabItemFlags_ = 64
	ImguitabitemflagsTrailing                     ImGuiTabItemFlags_ = 128
	ImguitabitemflagsNoassumedclosure             ImGuiTabItemFlags_ = 256
)

func (i ImGuiTabItemFlags_) String() string {
	switch i {
	case ImguitabitemflagsNone:
		return "Imguitabitemflags None"
	case ImguitabitemflagsUnsaveddocument:
		return "Imguitabitemflags Unsaveddocument"
	case ImguitabitemflagsSetselected:
		return "Imguitabitemflags Setselected"
	case ImguitabitemflagsNoclosewithmiddlemousebutton:
		return "Imguitabitemflags Noclosewithmiddlemousebutton"
	case ImguitabitemflagsNopushid:
		return "Imguitabitemflags Nopushid"
	case ImguitabitemflagsNotooltip:
		return "Imguitabitemflags Notooltip"
	case ImguitabitemflagsNoreorder:
		return "Imguitabitemflags Noreorder"
	case ImguitabitemflagsLeading:
		return "Imguitabitemflags Leading"
	case ImguitabitemflagsTrailing:
		return "Imguitabitemflags Trailing"
	case ImguitabitemflagsNoassumedclosure:
		return "Imguitabitemflags Noassumedclosure"
	default:
		return fmt.Sprintf("ImGuiTabItemFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiFocusedFlags_
type ImGuiFocusedFlags_ uint32

const (
	ImguifocusedflagsNone                ImGuiFocusedFlags_ = 0
	ImguifocusedflagsChildwindows        ImGuiFocusedFlags_ = 1
	ImguifocusedflagsRootwindow          ImGuiFocusedFlags_ = 2
	ImguifocusedflagsAnywindow           ImGuiFocusedFlags_ = 4
	ImguifocusedflagsNopopuphierarchy    ImGuiFocusedFlags_ = 8
	ImguifocusedflagsRootandchildwindows ImGuiFocusedFlags_ = 3
)

func (i ImGuiFocusedFlags_) String() string {
	switch i {
	case ImguifocusedflagsNone:
		return "Imguifocusedflags None"
	case ImguifocusedflagsChildwindows:
		return "Imguifocusedflags Childwindows"
	case ImguifocusedflagsRootwindow:
		return "Imguifocusedflags Rootwindow"
	case ImguifocusedflagsAnywindow:
		return "Imguifocusedflags Anywindow"
	case ImguifocusedflagsNopopuphierarchy:
		return "Imguifocusedflags Nopopuphierarchy"
	case ImguifocusedflagsRootandchildwindows:
		return "Imguifocusedflags Rootandchildwindows"
	default:
		return fmt.Sprintf("ImGuiFocusedFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiHoveredFlags_
type ImGuiHoveredFlags_ uint32

const (
	ImguihoveredflagsNone                         ImGuiHoveredFlags_ = 0
	ImguihoveredflagsChildwindows                 ImGuiHoveredFlags_ = 1
	ImguihoveredflagsRootwindow                   ImGuiHoveredFlags_ = 2
	ImguihoveredflagsAnywindow                    ImGuiHoveredFlags_ = 4
	ImguihoveredflagsNopopuphierarchy             ImGuiHoveredFlags_ = 8
	ImguihoveredflagsAllowwhenblockedbypopup      ImGuiHoveredFlags_ = 32
	ImguihoveredflagsAllowwhenblockedbyactiveitem ImGuiHoveredFlags_ = 128
	ImguihoveredflagsAllowwhenoverlappedbyitem    ImGuiHoveredFlags_ = 256
	ImguihoveredflagsAllowwhenoverlappedbywindow  ImGuiHoveredFlags_ = 512
	ImguihoveredflagsAllowwhendisabled            ImGuiHoveredFlags_ = 1024
	ImguihoveredflagsNonavoverride                ImGuiHoveredFlags_ = 2048
	ImguihoveredflagsAllowwhenoverlapped          ImGuiHoveredFlags_ = 768
	ImguihoveredflagsRectonly                     ImGuiHoveredFlags_ = 928
	ImguihoveredflagsRootandchildwindows          ImGuiHoveredFlags_ = 3
	ImguihoveredflagsFortooltip                   ImGuiHoveredFlags_ = 4096
	ImguihoveredflagsStationary                   ImGuiHoveredFlags_ = 8192
	ImguihoveredflagsDelaynone                    ImGuiHoveredFlags_ = 16384
	ImguihoveredflagsDelayshort                   ImGuiHoveredFlags_ = 32768
	ImguihoveredflagsDelaynormal                  ImGuiHoveredFlags_ = 65536
	ImguihoveredflagsNoshareddelay                ImGuiHoveredFlags_ = 131072
)

func (i ImGuiHoveredFlags_) String() string {
	switch i {
	case ImguihoveredflagsNone:
		return "Imguihoveredflags None"
	case ImguihoveredflagsChildwindows:
		return "Imguihoveredflags Childwindows"
	case ImguihoveredflagsRootwindow:
		return "Imguihoveredflags Rootwindow"
	case ImguihoveredflagsAnywindow:
		return "Imguihoveredflags Anywindow"
	case ImguihoveredflagsNopopuphierarchy:
		return "Imguihoveredflags Nopopuphierarchy"
	case ImguihoveredflagsAllowwhenblockedbypopup:
		return "Imguihoveredflags Allowwhenblockedbypopup"
	case ImguihoveredflagsAllowwhenblockedbyactiveitem:
		return "Imguihoveredflags Allowwhenblockedbyactiveitem"
	case ImguihoveredflagsAllowwhenoverlappedbyitem:
		return "Imguihoveredflags Allowwhenoverlappedbyitem"
	case ImguihoveredflagsAllowwhenoverlappedbywindow:
		return "Imguihoveredflags Allowwhenoverlappedbywindow"
	case ImguihoveredflagsAllowwhendisabled:
		return "Imguihoveredflags Allowwhendisabled"
	case ImguihoveredflagsNonavoverride:
		return "Imguihoveredflags Nonavoverride"
	case ImguihoveredflagsAllowwhenoverlapped:
		return "Imguihoveredflags Allowwhenoverlapped"
	case ImguihoveredflagsRectonly:
		return "Imguihoveredflags Rectonly"
	case ImguihoveredflagsRootandchildwindows:
		return "Imguihoveredflags Rootandchildwindows"
	case ImguihoveredflagsFortooltip:
		return "Imguihoveredflags Fortooltip"
	case ImguihoveredflagsStationary:
		return "Imguihoveredflags Stationary"
	case ImguihoveredflagsDelaynone:
		return "Imguihoveredflags Delaynone"
	case ImguihoveredflagsDelayshort:
		return "Imguihoveredflags Delayshort"
	case ImguihoveredflagsDelaynormal:
		return "Imguihoveredflags Delaynormal"
	case ImguihoveredflagsNoshareddelay:
		return "Imguihoveredflags Noshareddelay"
	default:
		return fmt.Sprintf("ImGuiHoveredFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiDragDropFlags_
type ImGuiDragDropFlags_ uint32

const (
	ImguidragdropflagsNone                     ImGuiDragDropFlags_ = 0
	ImguidragdropflagsSourcenopreviewtooltip   ImGuiDragDropFlags_ = 1
	ImguidragdropflagsSourcenodisablehover     ImGuiDragDropFlags_ = 2
	ImguidragdropflagsSourcenoholdtoopenothers ImGuiDragDropFlags_ = 4
	ImguidragdropflagsSourceallownullid        ImGuiDragDropFlags_ = 8
	ImguidragdropflagsSourceextern             ImGuiDragDropFlags_ = 16
	ImguidragdropflagsPayloadautoexpire        ImGuiDragDropFlags_ = 32
	ImguidragdropflagsPayloadnocrosscontext    ImGuiDragDropFlags_ = 64
	ImguidragdropflagsPayloadnocrossprocess    ImGuiDragDropFlags_ = 128
	ImguidragdropflagsAcceptbeforedelivery     ImGuiDragDropFlags_ = 1024
	ImguidragdropflagsAcceptnodrawdefaultrect  ImGuiDragDropFlags_ = 2048
	ImguidragdropflagsAcceptnopreviewtooltip   ImGuiDragDropFlags_ = 4096
	ImguidragdropflagsAcceptdrawashovered      ImGuiDragDropFlags_ = 8192
	ImguidragdropflagsAcceptpeekonly           ImGuiDragDropFlags_ = 3072
	ImguidragdropflagsSourceautoexpirepayload  ImGuiDragDropFlags_ = 32
)

func (i ImGuiDragDropFlags_) String() string {
	switch i {
	case ImguidragdropflagsNone:
		return "Imguidragdropflags None"
	case ImguidragdropflagsSourcenopreviewtooltip:
		return "Imguidragdropflags Sourcenopreviewtooltip"
	case ImguidragdropflagsSourcenodisablehover:
		return "Imguidragdropflags Sourcenodisablehover"
	case ImguidragdropflagsSourcenoholdtoopenothers:
		return "Imguidragdropflags Sourcenoholdtoopenothers"
	case ImguidragdropflagsSourceallownullid:
		return "Imguidragdropflags Sourceallownullid"
	case ImguidragdropflagsSourceextern:
		return "Imguidragdropflags Sourceextern"
	case ImguidragdropflagsPayloadautoexpire:
		return "Imguidragdropflags Payloadautoexpire"
	case ImguidragdropflagsPayloadnocrosscontext:
		return "Imguidragdropflags Payloadnocrosscontext"
	case ImguidragdropflagsPayloadnocrossprocess:
		return "Imguidragdropflags Payloadnocrossprocess"
	case ImguidragdropflagsAcceptbeforedelivery:
		return "Imguidragdropflags Acceptbeforedelivery"
	case ImguidragdropflagsAcceptnodrawdefaultrect:
		return "Imguidragdropflags Acceptnodrawdefaultrect"
	case ImguidragdropflagsAcceptnopreviewtooltip:
		return "Imguidragdropflags Acceptnopreviewtooltip"
	case ImguidragdropflagsAcceptdrawashovered:
		return "Imguidragdropflags Acceptdrawashovered"
	case ImguidragdropflagsAcceptpeekonly:
		return "Imguidragdropflags Acceptpeekonly"
	default:
		return fmt.Sprintf("ImGuiDragDropFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiDataType_
type ImGuiDataType_ uint32

const (
	ImguidatatypeS8 ImGuiDataType_ = iota
	ImguidatatypeU8
	ImguidatatypeS16
	ImguidatatypeU16
	ImguidatatypeS32
	ImguidatatypeU32
	ImguidatatypeS64
	ImguidatatypeU64
	ImguidatatypeFloat
	ImguidatatypeDouble
	ImguidatatypeBool
	ImguidatatypeString
	ImguidatatypeCount
)

func (i ImGuiDataType_) String() string {
	switch i {
	case ImguidatatypeS8:
		return "Imguidatatype S8"
	case ImguidatatypeU8:
		return "Imguidatatype U8"
	case ImguidatatypeS16:
		return "Imguidatatype S16"
	case ImguidatatypeU16:
		return "Imguidatatype U16"
	case ImguidatatypeS32:
		return "Imguidatatype S32"
	case ImguidatatypeU32:
		return "Imguidatatype U32"
	case ImguidatatypeS64:
		return "Imguidatatype S64"
	case ImguidatatypeU64:
		return "Imguidatatype U64"
	case ImguidatatypeFloat:
		return "Imguidatatype Float"
	case ImguidatatypeDouble:
		return "Imguidatatype Double"
	case ImguidatatypeBool:
		return "Imguidatatype Bool"
	case ImguidatatypeString:
		return "Imguidatatype String"
	case ImguidatatypeCount:
		return "Imguidatatype Count"
	default:
		return fmt.Sprintf("ImGuiDataType_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiInputFlags_
type ImGuiInputFlags_ uint32

const (
	ImguiinputflagsNone                 ImGuiInputFlags_ = 0
	ImguiinputflagsRepeat               ImGuiInputFlags_ = 1
	ImguiinputflagsRouteactive          ImGuiInputFlags_ = 1024
	ImguiinputflagsRoutefocused         ImGuiInputFlags_ = 2048
	ImguiinputflagsRouteglobal          ImGuiInputFlags_ = 4096
	ImguiinputflagsRoutealways          ImGuiInputFlags_ = 8192
	ImguiinputflagsRouteoverfocused     ImGuiInputFlags_ = 16384
	ImguiinputflagsRouteoveractive      ImGuiInputFlags_ = 32768
	ImguiinputflagsRouteunlessbgfocused ImGuiInputFlags_ = 65536
	ImguiinputflagsRoutefromrootwindow  ImGuiInputFlags_ = 131072
	ImguiinputflagsTooltip              ImGuiInputFlags_ = 262144
)

func (i ImGuiInputFlags_) String() string {
	switch i {
	case ImguiinputflagsNone:
		return "Imguiinputflags None"
	case ImguiinputflagsRepeat:
		return "Imguiinputflags Repeat"
	case ImguiinputflagsRouteactive:
		return "Imguiinputflags Routeactive"
	case ImguiinputflagsRoutefocused:
		return "Imguiinputflags Routefocused"
	case ImguiinputflagsRouteglobal:
		return "Imguiinputflags Routeglobal"
	case ImguiinputflagsRoutealways:
		return "Imguiinputflags Routealways"
	case ImguiinputflagsRouteoverfocused:
		return "Imguiinputflags Routeoverfocused"
	case ImguiinputflagsRouteoveractive:
		return "Imguiinputflags Routeoveractive"
	case ImguiinputflagsRouteunlessbgfocused:
		return "Imguiinputflags Routeunlessbgfocused"
	case ImguiinputflagsRoutefromrootwindow:
		return "Imguiinputflags Routefromrootwindow"
	case ImguiinputflagsTooltip:
		return "Imguiinputflags Tooltip"
	default:
		return fmt.Sprintf("ImGuiInputFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiConfigFlags_
type ImGuiConfigFlags_ uint32

const (
	ImguiconfigflagsNone                 ImGuiConfigFlags_ = 0
	ImguiconfigflagsNavenablekeyboard    ImGuiConfigFlags_ = 1
	ImguiconfigflagsNavenablegamepad     ImGuiConfigFlags_ = 2
	ImguiconfigflagsNomouse              ImGuiConfigFlags_ = 16
	ImguiconfigflagsNomousecursorchange  ImGuiConfigFlags_ = 32
	ImguiconfigflagsNokeyboard           ImGuiConfigFlags_ = 64
	ImguiconfigflagsIssrgb               ImGuiConfigFlags_ = 1048576
	ImguiconfigflagsIstouchscreen        ImGuiConfigFlags_ = 2097152
	ImguiconfigflagsNavenablesetmousepos ImGuiConfigFlags_ = 4
	ImguiconfigflagsNavnocapturekeyboard ImGuiConfigFlags_ = 8
)

func (i ImGuiConfigFlags_) String() string {
	switch i {
	case ImguiconfigflagsNone:
		return "Imguiconfigflags None"
	case ImguiconfigflagsNavenablekeyboard:
		return "Imguiconfigflags Navenablekeyboard"
	case ImguiconfigflagsNavenablegamepad:
		return "Imguiconfigflags Navenablegamepad"
	case ImguiconfigflagsNomouse:
		return "Imguiconfigflags Nomouse"
	case ImguiconfigflagsNomousecursorchange:
		return "Imguiconfigflags Nomousecursorchange"
	case ImguiconfigflagsNokeyboard:
		return "Imguiconfigflags Nokeyboard"
	case ImguiconfigflagsIssrgb:
		return "Imguiconfigflags Issrgb"
	case ImguiconfigflagsIstouchscreen:
		return "Imguiconfigflags Istouchscreen"
	case ImguiconfigflagsNavenablesetmousepos:
		return "Imguiconfigflags Navenablesetmousepos"
	case ImguiconfigflagsNavnocapturekeyboard:
		return "Imguiconfigflags Navnocapturekeyboard"
	default:
		return fmt.Sprintf("ImGuiConfigFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiBackendFlags_
type ImGuiBackendFlags_ uint32

const (
	ImguibackendflagsNone                 ImGuiBackendFlags_ = 0
	ImguibackendflagsHasgamepad           ImGuiBackendFlags_ = 1
	ImguibackendflagsHasmousecursors      ImGuiBackendFlags_ = 2
	ImguibackendflagsHassetmousepos       ImGuiBackendFlags_ = 4
	ImguibackendflagsRendererhasvtxoffset ImGuiBackendFlags_ = 8
	ImguibackendflagsRendererhastextures  ImGuiBackendFlags_ = 16
)

func (i ImGuiBackendFlags_) String() string {
	switch i {
	case ImguibackendflagsNone:
		return "Imguibackendflags None"
	case ImguibackendflagsHasgamepad:
		return "Imguibackendflags Hasgamepad"
	case ImguibackendflagsHasmousecursors:
		return "Imguibackendflags Hasmousecursors"
	case ImguibackendflagsHassetmousepos:
		return "Imguibackendflags Hassetmousepos"
	case ImguibackendflagsRendererhasvtxoffset:
		return "Imguibackendflags Rendererhasvtxoffset"
	case ImguibackendflagsRendererhastextures:
		return "Imguibackendflags Rendererhastextures"
	default:
		return fmt.Sprintf("ImGuiBackendFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiCol_
type ImGuiCol_ uint32

const (
	ImguicolText                      ImGuiCol_ = 0
	ImguicolTextdisabled              ImGuiCol_ = 1
	ImguicolWindowbg                  ImGuiCol_ = 2
	ImguicolChildbg                   ImGuiCol_ = 3
	ImguicolPopupbg                   ImGuiCol_ = 4
	ImguicolBorder                    ImGuiCol_ = 5
	ImguicolBordershadow              ImGuiCol_ = 6
	ImguicolFramebg                   ImGuiCol_ = 7
	ImguicolFramebghovered            ImGuiCol_ = 8
	ImguicolFramebgactive             ImGuiCol_ = 9
	ImguicolTitlebg                   ImGuiCol_ = 10
	ImguicolTitlebgactive             ImGuiCol_ = 11
	ImguicolTitlebgcollapsed          ImGuiCol_ = 12
	ImguicolMenubarbg                 ImGuiCol_ = 13
	ImguicolScrollbarbg               ImGuiCol_ = 14
	ImguicolScrollbargrab             ImGuiCol_ = 15
	ImguicolScrollbargrabhovered      ImGuiCol_ = 16
	ImguicolScrollbargrabactive       ImGuiCol_ = 17
	ImguicolCheckmark                 ImGuiCol_ = 18
	ImguicolCheckboxselectedbg        ImGuiCol_ = 19
	ImguicolSlidergrab                ImGuiCol_ = 20
	ImguicolSlidergrabactive          ImGuiCol_ = 21
	ImguicolButton                    ImGuiCol_ = 22
	ImguicolButtonhovered             ImGuiCol_ = 23
	ImguicolButtonactive              ImGuiCol_ = 24
	ImguicolHeader                    ImGuiCol_ = 25
	ImguicolHeaderhovered             ImGuiCol_ = 26
	ImguicolHeaderactive              ImGuiCol_ = 27
	ImguicolSeparator                 ImGuiCol_ = 28
	ImguicolSeparatorhovered          ImGuiCol_ = 29
	ImguicolSeparatoractive           ImGuiCol_ = 30
	ImguicolResizegrip                ImGuiCol_ = 31
	ImguicolResizegriphovered         ImGuiCol_ = 32
	ImguicolResizegripactive          ImGuiCol_ = 33
	ImguicolInputtextcursor           ImGuiCol_ = 34
	ImguicolTabhovered                ImGuiCol_ = 35
	ImguicolTab                       ImGuiCol_ = 36
	ImguicolTabselected               ImGuiCol_ = 37
	ImguicolTabselectedoverline       ImGuiCol_ = 38
	ImguicolTabdimmed                 ImGuiCol_ = 39
	ImguicolTabdimmedselected         ImGuiCol_ = 40
	ImguicolTabdimmedselectedoverline ImGuiCol_ = 41
	ImguicolPlotlines                 ImGuiCol_ = 42
	ImguicolPlotlineshovered          ImGuiCol_ = 43
	ImguicolPlothistogram             ImGuiCol_ = 44
	ImguicolPlothistogramhovered      ImGuiCol_ = 45
	ImguicolTableheaderbg             ImGuiCol_ = 46
	ImguicolTableborderstrong         ImGuiCol_ = 47
	ImguicolTableborderlight          ImGuiCol_ = 48
	ImguicolTablerowbg                ImGuiCol_ = 49
	ImguicolTablerowbgalt             ImGuiCol_ = 50
	ImguicolTextlink                  ImGuiCol_ = 51
	ImguicolTextselectedbg            ImGuiCol_ = 52
	ImguicolTreelines                 ImGuiCol_ = 53
	ImguicolDragdroptarget            ImGuiCol_ = 54
	ImguicolDragdroptargetbg          ImGuiCol_ = 55
	ImguicolUnsavedmarker             ImGuiCol_ = 56
	ImguicolNavcursor                 ImGuiCol_ = 57
	ImguicolNavwindowinghighlight     ImGuiCol_ = 58
	ImguicolNavwindowingdimbg         ImGuiCol_ = 59
	ImguicolModalwindowdimbg          ImGuiCol_ = 60
	ImguicolCount                     ImGuiCol_ = 61
	ImguicolTabactive                 ImGuiCol_ = 37
	ImguicolTabunfocused              ImGuiCol_ = 39
	ImguicolTabunfocusedactive        ImGuiCol_ = 40
	ImguicolNavhighlight              ImGuiCol_ = 57
)

func (i ImGuiCol_) String() string {
	switch i {
	case ImguicolText:
		return "Imguicol Text"
	case ImguicolTextdisabled:
		return "Imguicol Textdisabled"
	case ImguicolWindowbg:
		return "Imguicol Windowbg"
	case ImguicolChildbg:
		return "Imguicol Childbg"
	case ImguicolPopupbg:
		return "Imguicol Popupbg"
	case ImguicolBorder:
		return "Imguicol Border"
	case ImguicolBordershadow:
		return "Imguicol Bordershadow"
	case ImguicolFramebg:
		return "Imguicol Framebg"
	case ImguicolFramebghovered:
		return "Imguicol Framebghovered"
	case ImguicolFramebgactive:
		return "Imguicol Framebgactive"
	case ImguicolTitlebg:
		return "Imguicol Titlebg"
	case ImguicolTitlebgactive:
		return "Imguicol Titlebgactive"
	case ImguicolTitlebgcollapsed:
		return "Imguicol Titlebgcollapsed"
	case ImguicolMenubarbg:
		return "Imguicol Menubarbg"
	case ImguicolScrollbarbg:
		return "Imguicol Scrollbarbg"
	case ImguicolScrollbargrab:
		return "Imguicol Scrollbargrab"
	case ImguicolScrollbargrabhovered:
		return "Imguicol Scrollbargrabhovered"
	case ImguicolScrollbargrabactive:
		return "Imguicol Scrollbargrabactive"
	case ImguicolCheckmark:
		return "Imguicol Checkmark"
	case ImguicolCheckboxselectedbg:
		return "Imguicol Checkboxselectedbg"
	case ImguicolSlidergrab:
		return "Imguicol Slidergrab"
	case ImguicolSlidergrabactive:
		return "Imguicol Slidergrabactive"
	case ImguicolButton:
		return "Imguicol Button"
	case ImguicolButtonhovered:
		return "Imguicol Buttonhovered"
	case ImguicolButtonactive:
		return "Imguicol Buttonactive"
	case ImguicolHeader:
		return "Imguicol Header"
	case ImguicolHeaderhovered:
		return "Imguicol Headerhovered"
	case ImguicolHeaderactive:
		return "Imguicol Headeractive"
	case ImguicolSeparator:
		return "Imguicol Separator"
	case ImguicolSeparatorhovered:
		return "Imguicol Separatorhovered"
	case ImguicolSeparatoractive:
		return "Imguicol Separatoractive"
	case ImguicolResizegrip:
		return "Imguicol Resizegrip"
	case ImguicolResizegriphovered:
		return "Imguicol Resizegriphovered"
	case ImguicolResizegripactive:
		return "Imguicol Resizegripactive"
	case ImguicolInputtextcursor:
		return "Imguicol Inputtextcursor"
	case ImguicolTabhovered:
		return "Imguicol Tabhovered"
	case ImguicolTab:
		return "Imguicol Tab"
	case ImguicolTabselected:
		return "Imguicol Tabselected"
	case ImguicolTabselectedoverline:
		return "Imguicol Tabselectedoverline"
	case ImguicolTabdimmed:
		return "Imguicol Tabdimmed"
	case ImguicolTabdimmedselected:
		return "Imguicol Tabdimmedselected"
	case ImguicolTabdimmedselectedoverline:
		return "Imguicol Tabdimmedselectedoverline"
	case ImguicolPlotlines:
		return "Imguicol Plotlines"
	case ImguicolPlotlineshovered:
		return "Imguicol Plotlineshovered"
	case ImguicolPlothistogram:
		return "Imguicol Plothistogram"
	case ImguicolPlothistogramhovered:
		return "Imguicol Plothistogramhovered"
	case ImguicolTableheaderbg:
		return "Imguicol Tableheaderbg"
	case ImguicolTableborderstrong:
		return "Imguicol Tableborderstrong"
	case ImguicolTableborderlight:
		return "Imguicol Tableborderlight"
	case ImguicolTablerowbg:
		return "Imguicol Tablerowbg"
	case ImguicolTablerowbgalt:
		return "Imguicol Tablerowbgalt"
	case ImguicolTextlink:
		return "Imguicol Textlink"
	case ImguicolTextselectedbg:
		return "Imguicol Textselectedbg"
	case ImguicolTreelines:
		return "Imguicol Treelines"
	case ImguicolDragdroptarget:
		return "Imguicol Dragdroptarget"
	case ImguicolDragdroptargetbg:
		return "Imguicol Dragdroptargetbg"
	case ImguicolUnsavedmarker:
		return "Imguicol Unsavedmarker"
	case ImguicolNavcursor:
		return "Imguicol Navcursor"
	case ImguicolNavwindowinghighlight:
		return "Imguicol Navwindowinghighlight"
	case ImguicolNavwindowingdimbg:
		return "Imguicol Navwindowingdimbg"
	case ImguicolModalwindowdimbg:
		return "Imguicol Modalwindowdimbg"
	case ImguicolCount:
		return "Imguicol Count"
	default:
		return fmt.Sprintf("ImGuiCol_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiStyleVar_
type ImGuiStyleVar_ uint32

const (
	ImguistylevarAlpha ImGuiStyleVar_ = iota
	ImguistylevarDisabledalpha
	ImguistylevarWindowpadding
	ImguistylevarWindowrounding
	ImguistylevarWindowbordersize
	ImguistylevarWindowminsize
	ImguistylevarWindowtitlealign
	ImguistylevarChildrounding
	ImguistylevarChildbordersize
	ImguistylevarPopuprounding
	ImguistylevarPopupbordersize
	ImguistylevarFramepadding
	ImguistylevarFramerounding
	ImguistylevarFramebordersize
	ImguistylevarItemspacing
	ImguistylevarIteminnerspacing
	ImguistylevarIndentspacing
	ImguistylevarCellpadding
	ImguistylevarScrollbarsize
	ImguistylevarScrollbarrounding
	ImguistylevarScrollbarpadding
	ImguistylevarGrabminsize
	ImguistylevarGrabrounding
	ImguistylevarImagerounding
	ImguistylevarImagebordersize
	ImguistylevarTabrounding
	ImguistylevarTabbordersize
	ImguistylevarTabminwidthbase
	ImguistylevarTabminwidthshrink
	ImguistylevarTabbarbordersize
	ImguistylevarTabbaroverlinesize
	ImguistylevarTableangledheadersangle
	ImguistylevarTableangledheaderstextalign
	ImguistylevarTreelinessize
	ImguistylevarTreelinesrounding
	ImguistylevarDragdroptargetrounding
	ImguistylevarButtontextalign
	ImguistylevarSelectabletextalign
	ImguistylevarSeparatorsize
	ImguistylevarSeparatortextbordersize
	ImguistylevarSeparatortextalign
	ImguistylevarSeparatortextpadding
	ImguistylevarCount
)

func (i ImGuiStyleVar_) String() string {
	switch i {
	case ImguistylevarAlpha:
		return "Imguistylevar Alpha"
	case ImguistylevarDisabledalpha:
		return "Imguistylevar Disabledalpha"
	case ImguistylevarWindowpadding:
		return "Imguistylevar Windowpadding"
	case ImguistylevarWindowrounding:
		return "Imguistylevar Windowrounding"
	case ImguistylevarWindowbordersize:
		return "Imguistylevar Windowbordersize"
	case ImguistylevarWindowminsize:
		return "Imguistylevar Windowminsize"
	case ImguistylevarWindowtitlealign:
		return "Imguistylevar Windowtitlealign"
	case ImguistylevarChildrounding:
		return "Imguistylevar Childrounding"
	case ImguistylevarChildbordersize:
		return "Imguistylevar Childbordersize"
	case ImguistylevarPopuprounding:
		return "Imguistylevar Popuprounding"
	case ImguistylevarPopupbordersize:
		return "Imguistylevar Popupbordersize"
	case ImguistylevarFramepadding:
		return "Imguistylevar Framepadding"
	case ImguistylevarFramerounding:
		return "Imguistylevar Framerounding"
	case ImguistylevarFramebordersize:
		return "Imguistylevar Framebordersize"
	case ImguistylevarItemspacing:
		return "Imguistylevar Itemspacing"
	case ImguistylevarIteminnerspacing:
		return "Imguistylevar Iteminnerspacing"
	case ImguistylevarIndentspacing:
		return "Imguistylevar Indentspacing"
	case ImguistylevarCellpadding:
		return "Imguistylevar Cellpadding"
	case ImguistylevarScrollbarsize:
		return "Imguistylevar Scrollbarsize"
	case ImguistylevarScrollbarrounding:
		return "Imguistylevar Scrollbarrounding"
	case ImguistylevarScrollbarpadding:
		return "Imguistylevar Scrollbarpadding"
	case ImguistylevarGrabminsize:
		return "Imguistylevar Grabminsize"
	case ImguistylevarGrabrounding:
		return "Imguistylevar Grabrounding"
	case ImguistylevarImagerounding:
		return "Imguistylevar Imagerounding"
	case ImguistylevarImagebordersize:
		return "Imguistylevar Imagebordersize"
	case ImguistylevarTabrounding:
		return "Imguistylevar Tabrounding"
	case ImguistylevarTabbordersize:
		return "Imguistylevar Tabbordersize"
	case ImguistylevarTabminwidthbase:
		return "Imguistylevar Tabminwidthbase"
	case ImguistylevarTabminwidthshrink:
		return "Imguistylevar Tabminwidthshrink"
	case ImguistylevarTabbarbordersize:
		return "Imguistylevar Tabbarbordersize"
	case ImguistylevarTabbaroverlinesize:
		return "Imguistylevar Tabbaroverlinesize"
	case ImguistylevarTableangledheadersangle:
		return "Imguistylevar Tableangledheadersangle"
	case ImguistylevarTableangledheaderstextalign:
		return "Imguistylevar Tableangledheaderstextalign"
	case ImguistylevarTreelinessize:
		return "Imguistylevar Treelinessize"
	case ImguistylevarTreelinesrounding:
		return "Imguistylevar Treelinesrounding"
	case ImguistylevarDragdroptargetrounding:
		return "Imguistylevar Dragdroptargetrounding"
	case ImguistylevarButtontextalign:
		return "Imguistylevar Buttontextalign"
	case ImguistylevarSelectabletextalign:
		return "Imguistylevar Selectabletextalign"
	case ImguistylevarSeparatorsize:
		return "Imguistylevar Separatorsize"
	case ImguistylevarSeparatortextbordersize:
		return "Imguistylevar Separatortextbordersize"
	case ImguistylevarSeparatortextalign:
		return "Imguistylevar Separatortextalign"
	case ImguistylevarSeparatortextpadding:
		return "Imguistylevar Separatortextpadding"
	case ImguistylevarCount:
		return "Imguistylevar Count"
	default:
		return fmt.Sprintf("ImGuiStyleVar_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiButtonFlags_
type ImGuiButtonFlags_ uint32

const (
	ImguibuttonflagsNone              ImGuiButtonFlags_ = 0
	ImguibuttonflagsMousebuttonleft   ImGuiButtonFlags_ = 1
	ImguibuttonflagsMousebuttonright  ImGuiButtonFlags_ = 2
	ImguibuttonflagsMousebuttonmiddle ImGuiButtonFlags_ = 4
	ImguibuttonflagsMousebuttonmask   ImGuiButtonFlags_ = 7
	ImguibuttonflagsEnablenav         ImGuiButtonFlags_ = 8
	ImguibuttonflagsAllowoverlap      ImGuiButtonFlags_ = 4096
)

func (i ImGuiButtonFlags_) String() string {
	switch i {
	case ImguibuttonflagsNone:
		return "Imguibuttonflags None"
	case ImguibuttonflagsMousebuttonleft:
		return "Imguibuttonflags Mousebuttonleft"
	case ImguibuttonflagsMousebuttonright:
		return "Imguibuttonflags Mousebuttonright"
	case ImguibuttonflagsMousebuttonmiddle:
		return "Imguibuttonflags Mousebuttonmiddle"
	case ImguibuttonflagsMousebuttonmask:
		return "Imguibuttonflags Mousebuttonmask"
	case ImguibuttonflagsEnablenav:
		return "Imguibuttonflags Enablenav"
	case ImguibuttonflagsAllowoverlap:
		return "Imguibuttonflags Allowoverlap"
	default:
		return fmt.Sprintf("ImGuiButtonFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiColorEditFlags_
type ImGuiColorEditFlags_ uint32

const (
	ImguicoloreditflagsNone             ImGuiColorEditFlags_ = 0
	ImguicoloreditflagsNoalpha          ImGuiColorEditFlags_ = 2
	ImguicoloreditflagsNopicker         ImGuiColorEditFlags_ = 4
	ImguicoloreditflagsNooptions        ImGuiColorEditFlags_ = 8
	ImguicoloreditflagsNosmallpreview   ImGuiColorEditFlags_ = 16
	ImguicoloreditflagsNoinputs         ImGuiColorEditFlags_ = 32
	ImguicoloreditflagsNotooltip        ImGuiColorEditFlags_ = 64
	ImguicoloreditflagsNolabel          ImGuiColorEditFlags_ = 128
	ImguicoloreditflagsNosidepreview    ImGuiColorEditFlags_ = 256
	ImguicoloreditflagsNodragdrop       ImGuiColorEditFlags_ = 512
	ImguicoloreditflagsNoborder         ImGuiColorEditFlags_ = 1024
	ImguicoloreditflagsNocolormarkers   ImGuiColorEditFlags_ = 2048
	ImguicoloreditflagsAlphaopaque      ImGuiColorEditFlags_ = 4096
	ImguicoloreditflagsAlphanobg        ImGuiColorEditFlags_ = 8192
	ImguicoloreditflagsAlphapreviewhalf ImGuiColorEditFlags_ = 16384
	ImguicoloreditflagsAlphabar         ImGuiColorEditFlags_ = 262144
	ImguicoloreditflagsHdr              ImGuiColorEditFlags_ = 524288
	ImguicoloreditflagsDisplayrgb       ImGuiColorEditFlags_ = 1048576
	ImguicoloreditflagsDisplayhsv       ImGuiColorEditFlags_ = 2097152
	ImguicoloreditflagsDisplayhex       ImGuiColorEditFlags_ = 4194304
	ImguicoloreditflagsUint8            ImGuiColorEditFlags_ = 8388608
	ImguicoloreditflagsFloat            ImGuiColorEditFlags_ = 16777216
	ImguicoloreditflagsPickerhuebar     ImGuiColorEditFlags_ = 33554432
	ImguicoloreditflagsPickerhuewheel   ImGuiColorEditFlags_ = 67108864
	ImguicoloreditflagsInputrgb         ImGuiColorEditFlags_ = 134217728
	ImguicoloreditflagsInputhsv         ImGuiColorEditFlags_ = 268435456
	ImguicoloreditflagsDefaultoptions   ImGuiColorEditFlags_ = 177209344
	ImguicoloreditflagsAlphamask        ImGuiColorEditFlags_ = 28674
	ImguicoloreditflagsDisplaymask      ImGuiColorEditFlags_ = 7340032
	ImguicoloreditflagsDatatypemask     ImGuiColorEditFlags_ = 25165824
	ImguicoloreditflagsPickermask       ImGuiColorEditFlags_ = 100663296
	ImguicoloreditflagsInputmask        ImGuiColorEditFlags_ = 402653184
	ImguicoloreditflagsAlphapreview     ImGuiColorEditFlags_ = 0
)

func (i ImGuiColorEditFlags_) String() string {
	switch i {
	case ImguicoloreditflagsNone:
		return "Imguicoloreditflags None"
	case ImguicoloreditflagsNoalpha:
		return "Imguicoloreditflags Noalpha"
	case ImguicoloreditflagsNopicker:
		return "Imguicoloreditflags Nopicker"
	case ImguicoloreditflagsNooptions:
		return "Imguicoloreditflags Nooptions"
	case ImguicoloreditflagsNosmallpreview:
		return "Imguicoloreditflags Nosmallpreview"
	case ImguicoloreditflagsNoinputs:
		return "Imguicoloreditflags Noinputs"
	case ImguicoloreditflagsNotooltip:
		return "Imguicoloreditflags Notooltip"
	case ImguicoloreditflagsNolabel:
		return "Imguicoloreditflags Nolabel"
	case ImguicoloreditflagsNosidepreview:
		return "Imguicoloreditflags Nosidepreview"
	case ImguicoloreditflagsNodragdrop:
		return "Imguicoloreditflags Nodragdrop"
	case ImguicoloreditflagsNoborder:
		return "Imguicoloreditflags Noborder"
	case ImguicoloreditflagsNocolormarkers:
		return "Imguicoloreditflags Nocolormarkers"
	case ImguicoloreditflagsAlphaopaque:
		return "Imguicoloreditflags Alphaopaque"
	case ImguicoloreditflagsAlphanobg:
		return "Imguicoloreditflags Alphanobg"
	case ImguicoloreditflagsAlphapreviewhalf:
		return "Imguicoloreditflags Alphapreviewhalf"
	case ImguicoloreditflagsAlphabar:
		return "Imguicoloreditflags Alphabar"
	case ImguicoloreditflagsHdr:
		return "Imguicoloreditflags Hdr"
	case ImguicoloreditflagsDisplayrgb:
		return "Imguicoloreditflags Displayrgb"
	case ImguicoloreditflagsDisplayhsv:
		return "Imguicoloreditflags Displayhsv"
	case ImguicoloreditflagsDisplayhex:
		return "Imguicoloreditflags Displayhex"
	case ImguicoloreditflagsUint8:
		return "Imguicoloreditflags Uint 8"
	case ImguicoloreditflagsFloat:
		return "Imguicoloreditflags Float"
	case ImguicoloreditflagsPickerhuebar:
		return "Imguicoloreditflags Pickerhuebar"
	case ImguicoloreditflagsPickerhuewheel:
		return "Imguicoloreditflags Pickerhuewheel"
	case ImguicoloreditflagsInputrgb:
		return "Imguicoloreditflags Inputrgb"
	case ImguicoloreditflagsInputhsv:
		return "Imguicoloreditflags Inputhsv"
	case ImguicoloreditflagsDefaultoptions:
		return "Imguicoloreditflags Defaultoptions"
	case ImguicoloreditflagsAlphamask:
		return "Imguicoloreditflags Alphamask"
	case ImguicoloreditflagsDisplaymask:
		return "Imguicoloreditflags Displaymask"
	case ImguicoloreditflagsDatatypemask:
		return "Imguicoloreditflags Datatypemask"
	case ImguicoloreditflagsPickermask:
		return "Imguicoloreditflags Pickermask"
	case ImguicoloreditflagsInputmask:
		return "Imguicoloreditflags Inputmask"
	default:
		return fmt.Sprintf("ImGuiColorEditFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiSliderFlags_
type ImGuiSliderFlags_ uint32

const (
	ImguisliderflagsNone            ImGuiSliderFlags_ = 0
	ImguisliderflagsLogarithmic     ImGuiSliderFlags_ = 32
	ImguisliderflagsNoroundtoformat ImGuiSliderFlags_ = 64
	ImguisliderflagsNoinput         ImGuiSliderFlags_ = 128
	ImguisliderflagsWraparound      ImGuiSliderFlags_ = 256
	ImguisliderflagsClamponinput    ImGuiSliderFlags_ = 512
	ImguisliderflagsClampzerorange  ImGuiSliderFlags_ = 1024
	ImguisliderflagsNospeedtweaks   ImGuiSliderFlags_ = 2048
	ImguisliderflagsColormarkers    ImGuiSliderFlags_ = 4096
	ImguisliderflagsAlwaysclamp     ImGuiSliderFlags_ = 1536
	ImguisliderflagsInvalidmask     ImGuiSliderFlags_ = 1879048207
)

func (i ImGuiSliderFlags_) String() string {
	switch i {
	case ImguisliderflagsNone:
		return "Imguisliderflags None"
	case ImguisliderflagsLogarithmic:
		return "Imguisliderflags Logarithmic"
	case ImguisliderflagsNoroundtoformat:
		return "Imguisliderflags Noroundtoformat"
	case ImguisliderflagsNoinput:
		return "Imguisliderflags Noinput"
	case ImguisliderflagsWraparound:
		return "Imguisliderflags Wraparound"
	case ImguisliderflagsClamponinput:
		return "Imguisliderflags Clamponinput"
	case ImguisliderflagsClampzerorange:
		return "Imguisliderflags Clampzerorange"
	case ImguisliderflagsNospeedtweaks:
		return "Imguisliderflags Nospeedtweaks"
	case ImguisliderflagsColormarkers:
		return "Imguisliderflags Colormarkers"
	case ImguisliderflagsAlwaysclamp:
		return "Imguisliderflags Alwaysclamp"
	case ImguisliderflagsInvalidmask:
		return "Imguisliderflags Invalidmask"
	default:
		return fmt.Sprintf("ImGuiSliderFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiMouseButton_
type ImGuiMouseButton_ uint32

const (
	ImguimousebuttonLeft   ImGuiMouseButton_ = 0
	ImguimousebuttonRight  ImGuiMouseButton_ = 1
	ImguimousebuttonMiddle ImGuiMouseButton_ = 2
	ImguimousebuttonCount  ImGuiMouseButton_ = 5
)

func (i ImGuiMouseButton_) String() string {
	switch i {
	case ImguimousebuttonLeft:
		return "Imguimousebutton Left"
	case ImguimousebuttonRight:
		return "Imguimousebutton Right"
	case ImguimousebuttonMiddle:
		return "Imguimousebutton Middle"
	case ImguimousebuttonCount:
		return "Imguimousebutton Count"
	default:
		return fmt.Sprintf("ImGuiMouseButton_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiMouseCursor_
type ImGuiMouseCursor_ int32

const (
	ImguimousecursorNone ImGuiMouseCursor_ = -1 + iota
	ImguimousecursorArrow
	ImguimousecursorTextinput
	ImguimousecursorResizeall
	ImguimousecursorResizens
	ImguimousecursorResizeew
	ImguimousecursorResizenesw
	ImguimousecursorResizenwse
	ImguimousecursorHand
	ImguimousecursorWait
	ImguimousecursorProgress
	ImguimousecursorNotallowed
	ImguimousecursorCount
)

func (i ImGuiMouseCursor_) String() string {
	switch i {
	case ImguimousecursorNone:
		return "Imguimousecursor None"
	case ImguimousecursorArrow:
		return "Imguimousecursor Arrow"
	case ImguimousecursorTextinput:
		return "Imguimousecursor Textinput"
	case ImguimousecursorResizeall:
		return "Imguimousecursor Resizeall"
	case ImguimousecursorResizens:
		return "Imguimousecursor Resizens"
	case ImguimousecursorResizeew:
		return "Imguimousecursor Resizeew"
	case ImguimousecursorResizenesw:
		return "Imguimousecursor Resizenesw"
	case ImguimousecursorResizenwse:
		return "Imguimousecursor Resizenwse"
	case ImguimousecursorHand:
		return "Imguimousecursor Hand"
	case ImguimousecursorWait:
		return "Imguimousecursor Wait"
	case ImguimousecursorProgress:
		return "Imguimousecursor Progress"
	case ImguimousecursorNotallowed:
		return "Imguimousecursor Notallowed"
	case ImguimousecursorCount:
		return "Imguimousecursor Count"
	default:
		return fmt.Sprintf("ImGuiMouseCursor_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiCond_
type ImGuiCond_ uint32

const (
	ImguicondNone         ImGuiCond_ = 0
	ImguicondAlways       ImGuiCond_ = 1
	ImguicondOnce         ImGuiCond_ = 2
	ImguicondFirstuseever ImGuiCond_ = 4
	ImguicondAppearing    ImGuiCond_ = 8
)

func (i ImGuiCond_) String() string {
	switch i {
	case ImguicondNone:
		return "Imguicond None"
	case ImguicondAlways:
		return "Imguicond Always"
	case ImguicondOnce:
		return "Imguicond Once"
	case ImguicondFirstuseever:
		return "Imguicond Firstuseever"
	case ImguicondAppearing:
		return "Imguicond Appearing"
	default:
		return fmt.Sprintf("ImGuiCond_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTableFlags_
type ImGuiTableFlags_ uint32

const (
	ImguitableflagsNone                       ImGuiTableFlags_ = 0
	ImguitableflagsResizable                  ImGuiTableFlags_ = 1
	ImguitableflagsReorderable                ImGuiTableFlags_ = 2
	ImguitableflagsHideable                   ImGuiTableFlags_ = 4
	ImguitableflagsSortable                   ImGuiTableFlags_ = 8
	ImguitableflagsNosavedsettings            ImGuiTableFlags_ = 16
	ImguitableflagsContextmenuinbody          ImGuiTableFlags_ = 32
	ImguitableflagsRowbg                      ImGuiTableFlags_ = 64
	ImguitableflagsBordersinnerh              ImGuiTableFlags_ = 128
	ImguitableflagsBordersouterh              ImGuiTableFlags_ = 256
	ImguitableflagsBordersinnerv              ImGuiTableFlags_ = 512
	ImguitableflagsBordersouterv              ImGuiTableFlags_ = 1024
	ImguitableflagsBordersh                   ImGuiTableFlags_ = 384
	ImguitableflagsBordersv                   ImGuiTableFlags_ = 1536
	ImguitableflagsBordersinner               ImGuiTableFlags_ = 640
	ImguitableflagsBordersouter               ImGuiTableFlags_ = 1280
	ImguitableflagsBorders                    ImGuiTableFlags_ = 1920
	ImguitableflagsNobordersinbody            ImGuiTableFlags_ = 2048
	ImguitableflagsNobordersinbodyuntilresize ImGuiTableFlags_ = 4096
	ImguitableflagsSizingfixedfit             ImGuiTableFlags_ = 8192
	ImguitableflagsSizingfixedsame            ImGuiTableFlags_ = 16384
	ImguitableflagsSizingstretchprop          ImGuiTableFlags_ = 24576
	ImguitableflagsSizingstretchsame          ImGuiTableFlags_ = 32768
	ImguitableflagsNohostextendx              ImGuiTableFlags_ = 65536
	ImguitableflagsNohostextendy              ImGuiTableFlags_ = 131072
	ImguitableflagsNokeepcolumnsvisible       ImGuiTableFlags_ = 262144
	ImguitableflagsPrecisewidths              ImGuiTableFlags_ = 524288
	ImguitableflagsNoclip                     ImGuiTableFlags_ = 1048576
	ImguitableflagsPadouterx                  ImGuiTableFlags_ = 2097152
	ImguitableflagsNopadouterx                ImGuiTableFlags_ = 4194304
	ImguitableflagsNopadinnerx                ImGuiTableFlags_ = 8388608
	ImguitableflagsScrollx                    ImGuiTableFlags_ = 16777216
	ImguitableflagsScrolly                    ImGuiTableFlags_ = 33554432
	ImguitableflagsSortmulti                  ImGuiTableFlags_ = 67108864
	ImguitableflagsSorttristate               ImGuiTableFlags_ = 134217728
	ImguitableflagsHighlighthoveredcolumn     ImGuiTableFlags_ = 268435456
	ImguitableflagsSizingmask                 ImGuiTableFlags_ = 57344
)

func (i ImGuiTableFlags_) String() string {
	switch i {
	case ImguitableflagsNone:
		return "Imguitableflags None"
	case ImguitableflagsResizable:
		return "Imguitableflags Resizable"
	case ImguitableflagsReorderable:
		return "Imguitableflags Reorderable"
	case ImguitableflagsHideable:
		return "Imguitableflags Hideable"
	case ImguitableflagsSortable:
		return "Imguitableflags Sortable"
	case ImguitableflagsNosavedsettings:
		return "Imguitableflags Nosavedsettings"
	case ImguitableflagsContextmenuinbody:
		return "Imguitableflags Contextmenuinbody"
	case ImguitableflagsRowbg:
		return "Imguitableflags Rowbg"
	case ImguitableflagsBordersinnerh:
		return "Imguitableflags Bordersinnerh"
	case ImguitableflagsBordersouterh:
		return "Imguitableflags Bordersouterh"
	case ImguitableflagsBordersinnerv:
		return "Imguitableflags Bordersinnerv"
	case ImguitableflagsBordersouterv:
		return "Imguitableflags Bordersouterv"
	case ImguitableflagsBordersh:
		return "Imguitableflags Bordersh"
	case ImguitableflagsBordersv:
		return "Imguitableflags Bordersv"
	case ImguitableflagsBordersinner:
		return "Imguitableflags Bordersinner"
	case ImguitableflagsBordersouter:
		return "Imguitableflags Bordersouter"
	case ImguitableflagsBorders:
		return "Imguitableflags Borders"
	case ImguitableflagsNobordersinbody:
		return "Imguitableflags Nobordersinbody"
	case ImguitableflagsNobordersinbodyuntilresize:
		return "Imguitableflags Nobordersinbodyuntilresize"
	case ImguitableflagsSizingfixedfit:
		return "Imguitableflags Sizingfixedfit"
	case ImguitableflagsSizingfixedsame:
		return "Imguitableflags Sizingfixedsame"
	case ImguitableflagsSizingstretchprop:
		return "Imguitableflags Sizingstretchprop"
	case ImguitableflagsSizingstretchsame:
		return "Imguitableflags Sizingstretchsame"
	case ImguitableflagsNohostextendx:
		return "Imguitableflags Nohostextendx"
	case ImguitableflagsNohostextendy:
		return "Imguitableflags Nohostextendy"
	case ImguitableflagsNokeepcolumnsvisible:
		return "Imguitableflags Nokeepcolumnsvisible"
	case ImguitableflagsPrecisewidths:
		return "Imguitableflags Precisewidths"
	case ImguitableflagsNoclip:
		return "Imguitableflags Noclip"
	case ImguitableflagsPadouterx:
		return "Imguitableflags Padouterx"
	case ImguitableflagsNopadouterx:
		return "Imguitableflags Nopadouterx"
	case ImguitableflagsNopadinnerx:
		return "Imguitableflags Nopadinnerx"
	case ImguitableflagsScrollx:
		return "Imguitableflags Scrollx"
	case ImguitableflagsScrolly:
		return "Imguitableflags Scrolly"
	case ImguitableflagsSortmulti:
		return "Imguitableflags Sortmulti"
	case ImguitableflagsSorttristate:
		return "Imguitableflags Sorttristate"
	case ImguitableflagsHighlighthoveredcolumn:
		return "Imguitableflags Highlighthoveredcolumn"
	case ImguitableflagsSizingmask:
		return "Imguitableflags Sizingmask"
	default:
		return fmt.Sprintf("ImGuiTableFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTableColumnFlags_
type ImGuiTableColumnFlags_ uint32

const (
	ImguitablecolumnflagsNone                 ImGuiTableColumnFlags_ = 0
	ImguitablecolumnflagsDisabled             ImGuiTableColumnFlags_ = 1
	ImguitablecolumnflagsDefaulthide          ImGuiTableColumnFlags_ = 2
	ImguitablecolumnflagsDefaultsort          ImGuiTableColumnFlags_ = 4
	ImguitablecolumnflagsWidthstretch         ImGuiTableColumnFlags_ = 8
	ImguitablecolumnflagsWidthfixed           ImGuiTableColumnFlags_ = 16
	ImguitablecolumnflagsNoresize             ImGuiTableColumnFlags_ = 32
	ImguitablecolumnflagsNoreorder            ImGuiTableColumnFlags_ = 64
	ImguitablecolumnflagsNohide               ImGuiTableColumnFlags_ = 128
	ImguitablecolumnflagsNoclip               ImGuiTableColumnFlags_ = 256
	ImguitablecolumnflagsNosort               ImGuiTableColumnFlags_ = 512
	ImguitablecolumnflagsNosortascending      ImGuiTableColumnFlags_ = 1024
	ImguitablecolumnflagsNosortdescending     ImGuiTableColumnFlags_ = 2048
	ImguitablecolumnflagsNoheaderlabel        ImGuiTableColumnFlags_ = 4096
	ImguitablecolumnflagsNoheaderwidth        ImGuiTableColumnFlags_ = 8192
	ImguitablecolumnflagsPrefersortascending  ImGuiTableColumnFlags_ = 16384
	ImguitablecolumnflagsPrefersortdescending ImGuiTableColumnFlags_ = 32768
	ImguitablecolumnflagsIndentenable         ImGuiTableColumnFlags_ = 65536
	ImguitablecolumnflagsIndentdisable        ImGuiTableColumnFlags_ = 131072
	ImguitablecolumnflagsAngledheader         ImGuiTableColumnFlags_ = 262144
	ImguitablecolumnflagsIsenabled            ImGuiTableColumnFlags_ = 16777216
	ImguitablecolumnflagsIsvisible            ImGuiTableColumnFlags_ = 33554432
	ImguitablecolumnflagsIssorted             ImGuiTableColumnFlags_ = 67108864
	ImguitablecolumnflagsIshovered            ImGuiTableColumnFlags_ = 134217728
	ImguitablecolumnflagsWidthmask            ImGuiTableColumnFlags_ = 24
	ImguitablecolumnflagsIndentmask           ImGuiTableColumnFlags_ = 196608
	ImguitablecolumnflagsStatusmask           ImGuiTableColumnFlags_ = 251658240
	ImguitablecolumnflagsNodirectresize       ImGuiTableColumnFlags_ = 1073741824
)

func (i ImGuiTableColumnFlags_) String() string {
	switch i {
	case ImguitablecolumnflagsNone:
		return "Imguitablecolumnflags None"
	case ImguitablecolumnflagsDisabled:
		return "Imguitablecolumnflags Disabled"
	case ImguitablecolumnflagsDefaulthide:
		return "Imguitablecolumnflags Defaulthide"
	case ImguitablecolumnflagsDefaultsort:
		return "Imguitablecolumnflags Defaultsort"
	case ImguitablecolumnflagsWidthstretch:
		return "Imguitablecolumnflags Widthstretch"
	case ImguitablecolumnflagsWidthfixed:
		return "Imguitablecolumnflags Widthfixed"
	case ImguitablecolumnflagsNoresize:
		return "Imguitablecolumnflags Noresize"
	case ImguitablecolumnflagsNoreorder:
		return "Imguitablecolumnflags Noreorder"
	case ImguitablecolumnflagsNohide:
		return "Imguitablecolumnflags Nohide"
	case ImguitablecolumnflagsNoclip:
		return "Imguitablecolumnflags Noclip"
	case ImguitablecolumnflagsNosort:
		return "Imguitablecolumnflags Nosort"
	case ImguitablecolumnflagsNosortascending:
		return "Imguitablecolumnflags Nosortascending"
	case ImguitablecolumnflagsNosortdescending:
		return "Imguitablecolumnflags Nosortdescending"
	case ImguitablecolumnflagsNoheaderlabel:
		return "Imguitablecolumnflags Noheaderlabel"
	case ImguitablecolumnflagsNoheaderwidth:
		return "Imguitablecolumnflags Noheaderwidth"
	case ImguitablecolumnflagsPrefersortascending:
		return "Imguitablecolumnflags Prefersortascending"
	case ImguitablecolumnflagsPrefersortdescending:
		return "Imguitablecolumnflags Prefersortdescending"
	case ImguitablecolumnflagsIndentenable:
		return "Imguitablecolumnflags Indentenable"
	case ImguitablecolumnflagsIndentdisable:
		return "Imguitablecolumnflags Indentdisable"
	case ImguitablecolumnflagsAngledheader:
		return "Imguitablecolumnflags Angledheader"
	case ImguitablecolumnflagsIsenabled:
		return "Imguitablecolumnflags Isenabled"
	case ImguitablecolumnflagsIsvisible:
		return "Imguitablecolumnflags Isvisible"
	case ImguitablecolumnflagsIssorted:
		return "Imguitablecolumnflags Issorted"
	case ImguitablecolumnflagsIshovered:
		return "Imguitablecolumnflags Ishovered"
	case ImguitablecolumnflagsWidthmask:
		return "Imguitablecolumnflags Widthmask"
	case ImguitablecolumnflagsIndentmask:
		return "Imguitablecolumnflags Indentmask"
	case ImguitablecolumnflagsStatusmask:
		return "Imguitablecolumnflags Statusmask"
	case ImguitablecolumnflagsNodirectresize:
		return "Imguitablecolumnflags Nodirectresize"
	default:
		return fmt.Sprintf("ImGuiTableColumnFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTableRowFlags_
type ImGuiTableRowFlags_ uint32

const (
	ImguitablerowflagsNone ImGuiTableRowFlags_ = iota
	ImguitablerowflagsHeaders
)

func (i ImGuiTableRowFlags_) String() string {
	switch i {
	case ImguitablerowflagsNone:
		return "Imguitablerowflags None"
	case ImguitablerowflagsHeaders:
		return "Imguitablerowflags Headers"
	default:
		return fmt.Sprintf("ImGuiTableRowFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiTableBgTarget_
type ImGuiTableBgTarget_ uint32

const (
	ImguitablebgtargetNone ImGuiTableBgTarget_ = iota
	ImguitablebgtargetRowbg0
	ImguitablebgtargetRowbg1
	ImguitablebgtargetCellbg
)

func (i ImGuiTableBgTarget_) String() string {
	switch i {
	case ImguitablebgtargetNone:
		return "Imguitablebgtarget None"
	case ImguitablebgtargetRowbg0:
		return "Imguitablebgtarget Rowbg 0"
	case ImguitablebgtargetRowbg1:
		return "Imguitablebgtarget Rowbg 1"
	case ImguitablebgtargetCellbg:
		return "Imguitablebgtarget Cellbg"
	default:
		return fmt.Sprintf("ImGuiTableBgTarget_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiListClipperFlags_
type ImGuiListClipperFlags_ uint32

const (
	ImguilistclipperflagsNone ImGuiListClipperFlags_ = iota
	ImguilistclipperflagsNosettablerowcounters
)

func (i ImGuiListClipperFlags_) String() string {
	switch i {
	case ImguilistclipperflagsNone:
		return "Imguilistclipperflags None"
	case ImguilistclipperflagsNosettablerowcounters:
		return "Imguilistclipperflags Nosettablerowcounters"
	default:
		return fmt.Sprintf("ImGuiListClipperFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiMultiSelectFlags_
type ImGuiMultiSelectFlags_ uint32

const (
	ImguimultiselectflagsNone                  ImGuiMultiSelectFlags_ = 0
	ImguimultiselectflagsSingleselect          ImGuiMultiSelectFlags_ = 1
	ImguimultiselectflagsNoselectall           ImGuiMultiSelectFlags_ = 2
	ImguimultiselectflagsNorangeselect         ImGuiMultiSelectFlags_ = 4
	ImguimultiselectflagsNoautoselect          ImGuiMultiSelectFlags_ = 8
	ImguimultiselectflagsNoautoclear           ImGuiMultiSelectFlags_ = 16
	ImguimultiselectflagsNoautoclearonreselect ImGuiMultiSelectFlags_ = 32
	ImguimultiselectflagsBoxselect1d           ImGuiMultiSelectFlags_ = 64
	ImguimultiselectflagsBoxselect2d           ImGuiMultiSelectFlags_ = 128
	ImguimultiselectflagsBoxselectnoscroll     ImGuiMultiSelectFlags_ = 256
	ImguimultiselectflagsClearonescape         ImGuiMultiSelectFlags_ = 512
	ImguimultiselectflagsClearonclickvoid      ImGuiMultiSelectFlags_ = 1024
	ImguimultiselectflagsScopewindow           ImGuiMultiSelectFlags_ = 2048
	ImguimultiselectflagsScoperect             ImGuiMultiSelectFlags_ = 4096
	ImguimultiselectflagsSelectonauto          ImGuiMultiSelectFlags_ = 8192
	ImguimultiselectflagsSelectonclickalways   ImGuiMultiSelectFlags_ = 16384
	ImguimultiselectflagsSelectonclickrelease  ImGuiMultiSelectFlags_ = 32768
	ImguimultiselectflagsNavwrapx              ImGuiMultiSelectFlags_ = 65536
	ImguimultiselectflagsNoselectonrightclick  ImGuiMultiSelectFlags_ = 131072
	ImguimultiselectflagsSelectonmask          ImGuiMultiSelectFlags_ = 57344
	ImguimultiselectflagsSelectonclick         ImGuiMultiSelectFlags_ = 8192
)

func (i ImGuiMultiSelectFlags_) String() string {
	switch i {
	case ImguimultiselectflagsNone:
		return "Imguimultiselectflags None"
	case ImguimultiselectflagsSingleselect:
		return "Imguimultiselectflags Singleselect"
	case ImguimultiselectflagsNoselectall:
		return "Imguimultiselectflags Noselectall"
	case ImguimultiselectflagsNorangeselect:
		return "Imguimultiselectflags Norangeselect"
	case ImguimultiselectflagsNoautoselect:
		return "Imguimultiselectflags Noautoselect"
	case ImguimultiselectflagsNoautoclear:
		return "Imguimultiselectflags Noautoclear"
	case ImguimultiselectflagsNoautoclearonreselect:
		return "Imguimultiselectflags Noautoclearonreselect"
	case ImguimultiselectflagsBoxselect1d:
		return "Imguimultiselectflags Boxselect 1d"
	case ImguimultiselectflagsBoxselect2d:
		return "Imguimultiselectflags Boxselect 2d"
	case ImguimultiselectflagsBoxselectnoscroll:
		return "Imguimultiselectflags Boxselectnoscroll"
	case ImguimultiselectflagsClearonescape:
		return "Imguimultiselectflags Clearonescape"
	case ImguimultiselectflagsClearonclickvoid:
		return "Imguimultiselectflags Clearonclickvoid"
	case ImguimultiselectflagsScopewindow:
		return "Imguimultiselectflags Scopewindow"
	case ImguimultiselectflagsScoperect:
		return "Imguimultiselectflags Scoperect"
	case ImguimultiselectflagsSelectonauto:
		return "Imguimultiselectflags Selectonauto"
	case ImguimultiselectflagsSelectonclickalways:
		return "Imguimultiselectflags Selectonclickalways"
	case ImguimultiselectflagsSelectonclickrelease:
		return "Imguimultiselectflags Selectonclickrelease"
	case ImguimultiselectflagsNavwrapx:
		return "Imguimultiselectflags Navwrapx"
	case ImguimultiselectflagsNoselectonrightclick:
		return "Imguimultiselectflags Noselectonrightclick"
	case ImguimultiselectflagsSelectonmask:
		return "Imguimultiselectflags Selectonmask"
	default:
		return fmt.Sprintf("ImGuiMultiSelectFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiSelectionRequestType
type ImGuiSelectionRequestType uint32

const (
	ImguiselectionrequesttypeNone ImGuiSelectionRequestType = iota
	ImguiselectionrequesttypeSetall
	ImguiselectionrequesttypeSetrange
)

func (i ImGuiSelectionRequestType) String() string {
	switch i {
	case ImguiselectionrequesttypeNone:
		return "Imguiselectionrequesttype None"
	case ImguiselectionrequesttypeSetall:
		return "Imguiselectionrequesttype Setall"
	case ImguiselectionrequesttypeSetrange:
		return "Imguiselectionrequesttype Setrange"
	default:
		return fmt.Sprintf("ImGuiSelectionRequestType(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImDrawFlags_
type ImDrawFlags_ int32

const (
	ImdrawflagsNone                    ImDrawFlags_ = 0
	ImdrawflagsRoundcornerstopleft     ImDrawFlags_ = 16
	ImdrawflagsRoundcornerstopright    ImDrawFlags_ = 32
	ImdrawflagsRoundcornersbottomleft  ImDrawFlags_ = 64
	ImdrawflagsRoundcornersbottomright ImDrawFlags_ = 128
	ImdrawflagsRoundcornersnone        ImDrawFlags_ = 256
	ImdrawflagsClosed                  ImDrawFlags_ = 512
	ImdrawflagsRoundcornerstop         ImDrawFlags_ = 48
	ImdrawflagsRoundcornersbottom      ImDrawFlags_ = 192
	ImdrawflagsRoundcornersleft        ImDrawFlags_ = 80
	ImdrawflagsRoundcornersright       ImDrawFlags_ = 160
	ImdrawflagsRoundcornersall         ImDrawFlags_ = 240
	ImdrawflagsRoundcornersdefault     ImDrawFlags_ = 240
	ImdrawflagsRoundcornersmask        ImDrawFlags_ = 496
	ImdrawflagsInvalidmask             ImDrawFlags_ = -2147483633
)

func (i ImDrawFlags_) String() string {
	switch i {
	case ImdrawflagsNone:
		return "Imdrawflags None"
	case ImdrawflagsRoundcornerstopleft:
		return "Imdrawflags Roundcornerstopleft"
	case ImdrawflagsRoundcornerstopright:
		return "Imdrawflags Roundcornerstopright"
	case ImdrawflagsRoundcornersbottomleft:
		return "Imdrawflags Roundcornersbottomleft"
	case ImdrawflagsRoundcornersbottomright:
		return "Imdrawflags Roundcornersbottomright"
	case ImdrawflagsRoundcornersnone:
		return "Imdrawflags Roundcornersnone"
	case ImdrawflagsClosed:
		return "Imdrawflags Closed"
	case ImdrawflagsRoundcornerstop:
		return "Imdrawflags Roundcornerstop"
	case ImdrawflagsRoundcornersbottom:
		return "Imdrawflags Roundcornersbottom"
	case ImdrawflagsRoundcornersleft:
		return "Imdrawflags Roundcornersleft"
	case ImdrawflagsRoundcornersright:
		return "Imdrawflags Roundcornersright"
	case ImdrawflagsRoundcornersall:
		return "Imdrawflags Roundcornersall"
	case ImdrawflagsRoundcornersmask:
		return "Imdrawflags Roundcornersmask"
	case ImdrawflagsInvalidmask:
		return "Imdrawflags Invalidmask"
	default:
		return fmt.Sprintf("ImDrawFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImDrawListFlags_
type ImDrawListFlags_ uint32

const (
	ImdrawlistflagsNone                   ImDrawListFlags_ = 0
	ImdrawlistflagsAntialiasedlines       ImDrawListFlags_ = 1
	ImdrawlistflagsAntialiasedlinesusetex ImDrawListFlags_ = 2
	ImdrawlistflagsAntialiasedfill        ImDrawListFlags_ = 4
	ImdrawlistflagsAllowvtxoffset         ImDrawListFlags_ = 8
)

func (i ImDrawListFlags_) String() string {
	switch i {
	case ImdrawlistflagsNone:
		return "Imdrawlistflags None"
	case ImdrawlistflagsAntialiasedlines:
		return "Imdrawlistflags Antialiasedlines"
	case ImdrawlistflagsAntialiasedlinesusetex:
		return "Imdrawlistflags Antialiasedlinesusetex"
	case ImdrawlistflagsAntialiasedfill:
		return "Imdrawlistflags Antialiasedfill"
	case ImdrawlistflagsAllowvtxoffset:
		return "Imdrawlistflags Allowvtxoffset"
	default:
		return fmt.Sprintf("ImDrawListFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImTextureFormat
type ImTextureFormat uint32

const (
	ImtextureformatRgba32 ImTextureFormat = iota
	ImtextureformatAlpha8
)

func (i ImTextureFormat) String() string {
	switch i {
	case ImtextureformatRgba32:
		return "Imtextureformat Rgba 32"
	case ImtextureformatAlpha8:
		return "Imtextureformat Alpha 8"
	default:
		return fmt.Sprintf("ImTextureFormat(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImTextureStatus
type ImTextureStatus uint32

const (
	ImtexturestatusOk ImTextureStatus = iota
	ImtexturestatusDestroyed
	ImtexturestatusWantcreate
	ImtexturestatusWantupdates
	ImtexturestatusWantdestroy
)

func (i ImTextureStatus) String() string {
	switch i {
	case ImtexturestatusOk:
		return "Imtexturestatus Ok"
	case ImtexturestatusDestroyed:
		return "Imtexturestatus Destroyed"
	case ImtexturestatusWantcreate:
		return "Imtexturestatus Wantcreate"
	case ImtexturestatusWantupdates:
		return "Imtexturestatus Wantupdates"
	case ImtexturestatusWantdestroy:
		return "Imtexturestatus Wantdestroy"
	default:
		return fmt.Sprintf("ImTextureStatus(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImFontAtlasFlags_
type ImFontAtlasFlags_ uint32

const (
	ImfontatlasflagsNone               ImFontAtlasFlags_ = 0
	ImfontatlasflagsNopoweroftwoheight ImFontAtlasFlags_ = 1
	ImfontatlasflagsNomousecursors     ImFontAtlasFlags_ = 2
	ImfontatlasflagsNobakedlines       ImFontAtlasFlags_ = 4
)

func (i ImFontAtlasFlags_) String() string {
	switch i {
	case ImfontatlasflagsNone:
		return "Imfontatlasflags None"
	case ImfontatlasflagsNopoweroftwoheight:
		return "Imfontatlasflags Nopoweroftwoheight"
	case ImfontatlasflagsNomousecursors:
		return "Imfontatlasflags Nomousecursors"
	case ImfontatlasflagsNobakedlines:
		return "Imfontatlasflags Nobakedlines"
	default:
		return fmt.Sprintf("ImFontAtlasFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImFontFlags_
type ImFontFlags_ uint32

const (
	ImfontflagsNone            ImFontFlags_ = 0
	ImfontflagsNoloaderror     ImFontFlags_ = 2
	ImfontflagsNoloadglyphs    ImFontFlags_ = 4
	ImfontflagsLockbakedsizes  ImFontFlags_ = 8
	ImfontflagsImplicitrefsize ImFontFlags_ = 16
)

func (i ImFontFlags_) String() string {
	switch i {
	case ImfontflagsNone:
		return "Imfontflags None"
	case ImfontflagsNoloaderror:
		return "Imfontflags Noloaderror"
	case ImfontflagsNoloadglyphs:
		return "Imfontflags Noloadglyphs"
	case ImfontflagsLockbakedsizes:
		return "Imfontflags Lockbakedsizes"
	case ImfontflagsImplicitrefsize:
		return "Imfontflags Implicitrefsize"
	default:
		return fmt.Sprintf("ImFontFlags_(0x%X)", uint32(i))
	}
}

// Source: gen_cabi_imgui_backends.h:0 -> ImGuiViewportFlags_
type ImGuiViewportFlags_ uint32

const (
	ImguiviewportflagsNone              ImGuiViewportFlags_ = 0
	ImguiviewportflagsIsplatformwindow  ImGuiViewportFlags_ = 1
	ImguiviewportflagsIsplatformmonitor ImGuiViewportFlags_ = 2
	ImguiviewportflagsOwnedbyapp        ImGuiViewportFlags_ = 4
)

func (i ImGuiViewportFlags_) String() string {
	switch i {
	case ImguiviewportflagsNone:
		return "Imguiviewportflags None"
	case ImguiviewportflagsIsplatformwindow:
		return "Imguiviewportflags Isplatformwindow"
	case ImguiviewportflagsIsplatformmonitor:
		return "Imguiviewportflags Isplatformmonitor"
	case ImguiviewportflagsOwnedbyapp:
		return "Imguiviewportflags Ownedbyapp"
	default:
		return fmt.Sprintf("ImGuiViewportFlags_(0x%X)", uint32(i))
	}
}

type (
	ID3D11Buffer                  struct{} // gen_cabi_imgui_backends.h:7 -> ID3D11Buffer
	ID3D11Device                  struct{} // gen_cabi_imgui_backends.h:8 -> ID3D11Device
	ID3D11DeviceContext           struct{} // gen_cabi_imgui_backends.h:9 -> ID3D11DeviceContext
	ImDrawChannel                 struct{} // gen_cabi_imgui_backends.h:10 -> ImDrawChannel
	ImDrawCmd                     struct{} // gen_cabi_imgui_backends.h:11 -> ImDrawCmd
	ImDrawCmdHeader               struct{} // gen_cabi_imgui_backends.h:12 -> ImDrawCmdHeader
	ImDrawData                    struct{} // gen_cabi_imgui_backends.h:13 -> ImDrawData
	ImDrawList                    struct{} // gen_cabi_imgui_backends.h:14 -> ImDrawList
	ImDrawListSharedData          struct{} // gen_cabi_imgui_backends.h:15 -> ImDrawListSharedData
	ImDrawListSplitter            struct{} // gen_cabi_imgui_backends.h:16 -> ImDrawListSplitter
	ImDrawVert                    struct{} // gen_cabi_imgui_backends.h:17 -> ImDrawVert
	ImFont                        struct{} // gen_cabi_imgui_backends.h:18 -> ImFont
	ImFontAtlas                   struct{} // gen_cabi_imgui_backends.h:19 -> ImFontAtlas
	ImFontAtlasBuilder            struct{} // gen_cabi_imgui_backends.h:20 -> ImFontAtlasBuilder
	ImFontAtlasRect               struct{} // gen_cabi_imgui_backends.h:21 -> ImFontAtlasRect
	ImFontBaked                   struct{} // gen_cabi_imgui_backends.h:22 -> ImFontBaked
	ImFontConfig                  struct{} // gen_cabi_imgui_backends.h:23 -> ImFontConfig
	ImFontGlyph                   struct{} // gen_cabi_imgui_backends.h:24 -> ImFontGlyph
	ImFontGlyphRangesBuilder      struct{} // gen_cabi_imgui_backends.h:25 -> ImFontGlyphRangesBuilder
	ImFontLoader                  struct{} // gen_cabi_imgui_backends.h:26 -> ImFontLoader
	ImGuiContext                  struct{} // gen_cabi_imgui_backends.h:27 -> ImGuiContext
	ImGuiIO                       struct{} // gen_cabi_imgui_backends.h:28 -> ImGuiIO
	ImGuiInputTextCallbackData    struct{} // gen_cabi_imgui_backends.h:29 -> ImGuiInputTextCallbackData
	ImGuiKeyData                  struct{} // gen_cabi_imgui_backends.h:30 -> ImGuiKeyData
	ImGuiListClipper              struct{} // gen_cabi_imgui_backends.h:31 -> ImGuiListClipper
	ImGuiMultiSelectIO            struct{} // gen_cabi_imgui_backends.h:32 -> ImGuiMultiSelectIO
	ImGuiOnceUponAFrame           struct{} // gen_cabi_imgui_backends.h:33 -> ImGuiOnceUponAFrame
	ImGuiPayload                  struct{} // gen_cabi_imgui_backends.h:34 -> ImGuiPayload
	ImGuiPlatformIO               struct{} // gen_cabi_imgui_backends.h:35 -> ImGuiPlatformIO
	ImGuiPlatformImeData          struct{} // gen_cabi_imgui_backends.h:36 -> ImGuiPlatformImeData
	ImGuiSelectionBasicStorage    struct{} // gen_cabi_imgui_backends.h:37 -> ImGuiSelectionBasicStorage
	ImGuiSelectionExternalStorage struct{} // gen_cabi_imgui_backends.h:38 -> ImGuiSelectionExternalStorage
	ImGuiSelectionRequest         struct{} // gen_cabi_imgui_backends.h:39 -> ImGuiSelectionRequest
	ImGuiSizeCallbackData         struct{} // gen_cabi_imgui_backends.h:40 -> ImGuiSizeCallbackData
	ImGuiStorage                  struct{} // gen_cabi_imgui_backends.h:41 -> ImGuiStorage
	ImGuiStoragePair              struct{} // gen_cabi_imgui_backends.h:42 -> ImGuiStoragePair
	ImGuiStyle                    struct{} // gen_cabi_imgui_backends.h:43 -> ImGuiStyle
	ImGuiTableColumnSortSpecs     struct{} // gen_cabi_imgui_backends.h:44 -> ImGuiTableColumnSortSpecs
	ImGuiTableSortSpecs           struct{} // gen_cabi_imgui_backends.h:45 -> ImGuiTableSortSpecs
	ImGuiTextBuffer               struct{} // gen_cabi_imgui_backends.h:46 -> ImGuiTextBuffer
	ImGuiTextFilter               struct{} // gen_cabi_imgui_backends.h:47 -> ImGuiTextFilter
	ImGuiViewport                 struct{} // gen_cabi_imgui_backends.h:48 -> ImGuiViewport
	ImGui_ImplDX11_RenderState    struct{} // gen_cabi_imgui_backends.h:49 -> ImGui_ImplDX11_RenderState
	ImNewWrapper                  struct{} // gen_cabi_imgui_backends.h:50 -> ImNewWrapper
	ImTextureData                 struct{} // gen_cabi_imgui_backends.h:51 -> ImTextureData
	ImTextureRect                 struct{} // gen_cabi_imgui_backends.h:52 -> ImTextureRect
	ImTextureRef                  struct{} // gen_cabi_imgui_backends.h:53 -> ImTextureRef
	ImVec2                        struct {
		X float32
		Y float32
	} // gen_cabi_imgui_backends.h:921 -> ImVec2
	ImVec4 struct {
		X float32
		Y float32
		Z float32
		W float32
	} // gen_cabi_imgui_backends.h:922 -> ImVec4
	ImColor struct{} // gen_cabi_imgui_backends.h:926 -> ImColor
)

func (i *Imgui) ImVec2New() *ImVec2 {
	r1, _, _ := getProc("ImVec2_new").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec2New2(X float32, Y float32) *ImVec2 {
	r1, _, _ := getProc("ImVec2_new2").Call(uintptr(X), uintptr(Y))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec2New3(Param1 *ImVec2) *ImVec2 {
	r1, _, _ := getProc("ImVec2_new3").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec2X(Self *ImVec2) float32 {
	r1, _, _ := getProc("ImVec2_x").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec2SetX(Self *ImVec2, X float32) {
	getProc("ImVec2_setX").Call(uintptr(unsafe.Pointer(Self)), uintptr(X))
}

func (i *Imgui) ImVec2Y(Self *ImVec2) float32 {
	r1, _, _ := getProc("ImVec2_y").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec2SetY(Self *ImVec2, Y float32) {
	getProc("ImVec2_setY").Call(uintptr(unsafe.Pointer(Self)), uintptr(Y))
}

func (i *Imgui) ImVec2OperatorSubscript(Self *ImVec2, Idx uint64) *float32 {
	r1, _, _ := getProc("ImVec2_operatorSubscript").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&Idx)))
	return (*float32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec2OperatorSubscriptWithIdx(Self *ImVec2, Idx uint64) float32 {
	r1, _, _ := getProc("ImVec2_operatorSubscriptWithIdx").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&Idx)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec2OperatorAssign(Self *ImVec2, Param1 *ImVec2) {
	getProc("ImVec2_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImVec2Delete(Self *ImVec2) {
	getProc("ImVec2_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImVec4New() *ImVec4 {
	r1, _, _ := getProc("ImVec4_new").Call()
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec4New2(X float32, Y float32, Z float32, W float32) *ImVec4 {
	r1, _, _ := getProc("ImVec4_new2").Call(uintptr(X), uintptr(Y), uintptr(Z), uintptr(W))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec4New3(Param1 *ImVec4) *ImVec4 {
	r1, _, _ := getProc("ImVec4_new3").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImVec4X(Self *ImVec4) float32 {
	r1, _, _ := getProc("ImVec4_x").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec4SetX(Self *ImVec4, X float32) {
	getProc("ImVec4_setX").Call(uintptr(unsafe.Pointer(Self)), uintptr(X))
}

func (i *Imgui) ImVec4Y(Self *ImVec4) float32 {
	r1, _, _ := getProc("ImVec4_y").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec4SetY(Self *ImVec4, Y float32) {
	getProc("ImVec4_setY").Call(uintptr(unsafe.Pointer(Self)), uintptr(Y))
}

func (i *Imgui) ImVec4Z(Self *ImVec4) float32 {
	r1, _, _ := getProc("ImVec4_z").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec4SetZ(Self *ImVec4, Z float32) {
	getProc("ImVec4_setZ").Call(uintptr(unsafe.Pointer(Self)), uintptr(Z))
}

func (i *Imgui) ImVec4W(Self *ImVec4) float32 {
	r1, _, _ := getProc("ImVec4_w").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImVec4SetW(Self *ImVec4, W float32) {
	getProc("ImVec4_setW").Call(uintptr(unsafe.Pointer(Self)), uintptr(W))
}

func (i *Imgui) ImVec4OperatorAssign(Self *ImVec4, Param1 *ImVec4) {
	getProc("ImVec4_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImVec4Delete(Self *ImVec4) {
	getProc("ImVec4_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImTextureRefNew() *ImTextureRef {
	r1, _, _ := getProc("ImTextureRef_new").Call()
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRefNew2(Tex_id uint64) *ImTextureRef {
	r1, _, _ := getProc("ImTextureRef_new2").Call(*(*uintptr)(unsafe.Pointer(&Tex_id)))
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRefNew3(Tex_id unsafe.Pointer) *ImTextureRef {
	r1, _, _ := getProc("ImTextureRef_new3").Call(uintptr(Tex_id))
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRefNew4(Param1 *ImTextureRef) *ImTextureRef {
	r1, _, _ := getProc("ImTextureRef_new4").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRefGetTexID(Self *ImTextureRef) uint64 {
	r1, _, _ := getProc("ImTextureRef_GetTexID").Call(uintptr(unsafe.Pointer(Self)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImTextureRefTexData(Self *ImTextureRef) *ImTextureData {
	r1, _, _ := getProc("ImTextureRef__TexData").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRefSetTexData(Self *ImTextureRef, TexData *ImTextureData) {
	getProc("ImTextureRef_set_TexData").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexData)))
}

func (i *Imgui) ImTextureRefTexID(Self *ImTextureRef) uint64 {
	r1, _, _ := getProc("ImTextureRef__TexID").Call(uintptr(unsafe.Pointer(Self)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImTextureRefSetTexID(Self *ImTextureRef, TexID uint64) {
	getProc("ImTextureRef_set_TexID").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&TexID)))
}

func (i *Imgui) ImTextureRefOperatorAssign(Self *ImTextureRef, Param1 *ImTextureRef) {
	getProc("ImTextureRef_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImTextureRefDelete(Self *ImTextureRef) {
	getProc("ImTextureRef_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTableSortSpecsNew() *ImGuiTableSortSpecs {
	r1, _, _ := getProc("ImGuiTableSortSpecs_new").Call()
	return (*ImGuiTableSortSpecs)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTableSortSpecsSpecs(Self *ImGuiTableSortSpecs) *ImGuiTableColumnSortSpecs {
	r1, _, _ := getProc("ImGuiTableSortSpecs_Specs").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImGuiTableColumnSortSpecs)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTableSortSpecsSetSpecs(Self *ImGuiTableSortSpecs, Specs *ImGuiTableColumnSortSpecs) {
	getProc("ImGuiTableSortSpecs_setSpecs").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Specs)))
}

func (i *Imgui) ImGuiTableSortSpecsSpecsCount(Self *ImGuiTableSortSpecs) int32 {
	r1, _, _ := getProc("ImGuiTableSortSpecs_SpecsCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiTableSortSpecsSetSpecsCount(Self *ImGuiTableSortSpecs, SpecsCount int32) {
	getProc("ImGuiTableSortSpecs_setSpecsCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(SpecsCount))
}

func (i *Imgui) ImGuiTableSortSpecsSpecsDirty(Self *ImGuiTableSortSpecs) bool {
	r1, _, _ := getProc("ImGuiTableSortSpecs_SpecsDirty").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiTableSortSpecsSetSpecsDirty(Self *ImGuiTableSortSpecs, SpecsDirty bool) {
	getProc("ImGuiTableSortSpecs_setSpecsDirty").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if SpecsDirty {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiTableSortSpecsDelete(Self *ImGuiTableSortSpecs) {
	getProc("ImGuiTableSortSpecs_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTableColumnSortSpecsNew() *ImGuiTableColumnSortSpecs {
	r1, _, _ := getProc("ImGuiTableColumnSortSpecs_new").Call()
	return (*ImGuiTableColumnSortSpecs)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTableColumnSortSpecsColumnUserID(Self *ImGuiTableColumnSortSpecs) uint32 {
	r1, _, _ := getProc("ImGuiTableColumnSortSpecs_ColumnUserID").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiTableColumnSortSpecsSetColumnUserID(Self *ImGuiTableColumnSortSpecs, ColumnUserID uint32) {
	getProc("ImGuiTableColumnSortSpecs_setColumnUserID").Call(uintptr(unsafe.Pointer(Self)), uintptr(ColumnUserID))
}

func (i *Imgui) ImGuiTableColumnSortSpecsColumnIndex(Self *ImGuiTableColumnSortSpecs) int16 {
	r1, _, _ := getProc("ImGuiTableColumnSortSpecs_ColumnIndex").Call(uintptr(unsafe.Pointer(Self)))
	return int16(r1)
}

func (i *Imgui) ImGuiTableColumnSortSpecsSetColumnIndex(Self *ImGuiTableColumnSortSpecs, ColumnIndex int16) {
	getProc("ImGuiTableColumnSortSpecs_setColumnIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(ColumnIndex))
}

func (i *Imgui) ImGuiTableColumnSortSpecsSortOrder(Self *ImGuiTableColumnSortSpecs) int16 {
	r1, _, _ := getProc("ImGuiTableColumnSortSpecs_SortOrder").Call(uintptr(unsafe.Pointer(Self)))
	return int16(r1)
}

func (i *Imgui) ImGuiTableColumnSortSpecsSetSortOrder(Self *ImGuiTableColumnSortSpecs, SortOrder int16) {
	getProc("ImGuiTableColumnSortSpecs_setSortOrder").Call(uintptr(unsafe.Pointer(Self)), uintptr(SortOrder))
}

func (i *Imgui) ImGuiTableColumnSortSpecsSortDirection(Self *ImGuiTableColumnSortSpecs) int32 {
	r1, _, _ := getProc("ImGuiTableColumnSortSpecs_SortDirection").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiTableColumnSortSpecsSetSortDirection(Self *ImGuiTableColumnSortSpecs, SortDirection int32) {
	getProc("ImGuiTableColumnSortSpecs_setSortDirection").Call(uintptr(unsafe.Pointer(Self)), uintptr(SortDirection))
}

func (i *Imgui) ImGuiTableColumnSortSpecsDelete(Self *ImGuiTableColumnSortSpecs) {
	getProc("ImGuiTableColumnSortSpecs_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImNewWrapperDelete(Self *ImNewWrapper) {
	getProc("ImNewWrapper_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiStyleNew() *ImGuiStyle {
	r1, _, _ := getProc("ImGuiStyle_new").Call()
	return (*ImGuiStyle)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleFontSizeBase(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_FontSizeBase").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetFontSizeBase(Self *ImGuiStyle, FontSizeBase float32) {
	getProc("ImGuiStyle_setFontSizeBase").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontSizeBase))
}

func (i *Imgui) ImGuiStyleFontScaleMain(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_FontScaleMain").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetFontScaleMain(Self *ImGuiStyle, FontScaleMain float32) {
	getProc("ImGuiStyle_setFontScaleMain").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontScaleMain))
}

func (i *Imgui) ImGuiStyleFontScaleDpi(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_FontScaleDpi").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetFontScaleDpi(Self *ImGuiStyle, FontScaleDpi float32) {
	getProc("ImGuiStyle_setFontScaleDpi").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontScaleDpi))
}

func (i *Imgui) ImGuiStyleAlpha(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_Alpha").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetAlpha(Self *ImGuiStyle, Alpha float32) {
	getProc("ImGuiStyle_setAlpha").Call(uintptr(unsafe.Pointer(Self)), uintptr(Alpha))
}

func (i *Imgui) ImGuiStyleDisabledAlpha(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_DisabledAlpha").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetDisabledAlpha(Self *ImGuiStyle, DisabledAlpha float32) {
	getProc("ImGuiStyle_setDisabledAlpha").Call(uintptr(unsafe.Pointer(Self)), uintptr(DisabledAlpha))
}

func (i *Imgui) ImGuiStyleWindowPadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_WindowPadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetWindowPadding(Self *ImGuiStyle, WindowPadding *ImVec2) {
	getProc("ImGuiStyle_setWindowPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(WindowPadding)))
}

func (i *Imgui) ImGuiStyleWindowRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_WindowRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetWindowRounding(Self *ImGuiStyle, WindowRounding float32) {
	getProc("ImGuiStyle_setWindowRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(WindowRounding))
}

func (i *Imgui) ImGuiStyleWindowBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_WindowBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetWindowBorderSize(Self *ImGuiStyle, WindowBorderSize float32) {
	getProc("ImGuiStyle_setWindowBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(WindowBorderSize))
}

func (i *Imgui) ImGuiStyleWindowBorderHoverPadding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_WindowBorderHoverPadding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetWindowBorderHoverPadding(Self *ImGuiStyle, WindowBorderHoverPadding float32) {
	getProc("ImGuiStyle_setWindowBorderHoverPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(WindowBorderHoverPadding))
}

func (i *Imgui) ImGuiStyleWindowMinSize(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_WindowMinSize").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetWindowMinSize(Self *ImGuiStyle, WindowMinSize *ImVec2) {
	getProc("ImGuiStyle_setWindowMinSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(WindowMinSize)))
}

func (i *Imgui) ImGuiStyleWindowTitleAlign(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_WindowTitleAlign").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetWindowTitleAlign(Self *ImGuiStyle, WindowTitleAlign *ImVec2) {
	getProc("ImGuiStyle_setWindowTitleAlign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(WindowTitleAlign)))
}

func (i *Imgui) ImGuiStyleWindowMenuButtonPosition(Self *ImGuiStyle) int32 {
	r1, _, _ := getProc("ImGuiStyle_WindowMenuButtonPosition").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiStyleSetWindowMenuButtonPosition(Self *ImGuiStyle, WindowMenuButtonPosition int32) {
	getProc("ImGuiStyle_setWindowMenuButtonPosition").Call(uintptr(unsafe.Pointer(Self)), uintptr(WindowMenuButtonPosition))
}

func (i *Imgui) ImGuiStyleChildRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ChildRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetChildRounding(Self *ImGuiStyle, ChildRounding float32) {
	getProc("ImGuiStyle_setChildRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(ChildRounding))
}

func (i *Imgui) ImGuiStyleChildBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ChildBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetChildBorderSize(Self *ImGuiStyle, ChildBorderSize float32) {
	getProc("ImGuiStyle_setChildBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(ChildBorderSize))
}

func (i *Imgui) ImGuiStylePopupRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_PopupRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetPopupRounding(Self *ImGuiStyle, PopupRounding float32) {
	getProc("ImGuiStyle_setPopupRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(PopupRounding))
}

func (i *Imgui) ImGuiStylePopupBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_PopupBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetPopupBorderSize(Self *ImGuiStyle, PopupBorderSize float32) {
	getProc("ImGuiStyle_setPopupBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(PopupBorderSize))
}

func (i *Imgui) ImGuiStyleFramePadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_FramePadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetFramePadding(Self *ImGuiStyle, FramePadding *ImVec2) {
	getProc("ImGuiStyle_setFramePadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FramePadding)))
}

func (i *Imgui) ImGuiStyleFrameRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_FrameRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetFrameRounding(Self *ImGuiStyle, FrameRounding float32) {
	getProc("ImGuiStyle_setFrameRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(FrameRounding))
}

func (i *Imgui) ImGuiStyleFrameBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_FrameBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetFrameBorderSize(Self *ImGuiStyle, FrameBorderSize float32) {
	getProc("ImGuiStyle_setFrameBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(FrameBorderSize))
}

func (i *Imgui) ImGuiStyleItemSpacing(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_ItemSpacing").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetItemSpacing(Self *ImGuiStyle, ItemSpacing *ImVec2) {
	getProc("ImGuiStyle_setItemSpacing").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(ItemSpacing)))
}

func (i *Imgui) ImGuiStyleItemInnerSpacing(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_ItemInnerSpacing").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetItemInnerSpacing(Self *ImGuiStyle, ItemInnerSpacing *ImVec2) {
	getProc("ImGuiStyle_setItemInnerSpacing").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(ItemInnerSpacing)))
}

func (i *Imgui) ImGuiStyleCellPadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_CellPadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetCellPadding(Self *ImGuiStyle, CellPadding *ImVec2) {
	getProc("ImGuiStyle_setCellPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(CellPadding)))
}

func (i *Imgui) ImGuiStyleTouchExtraPadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_TouchExtraPadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetTouchExtraPadding(Self *ImGuiStyle, TouchExtraPadding *ImVec2) {
	getProc("ImGuiStyle_setTouchExtraPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TouchExtraPadding)))
}

func (i *Imgui) ImGuiStyleIndentSpacing(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_IndentSpacing").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetIndentSpacing(Self *ImGuiStyle, IndentSpacing float32) {
	getProc("ImGuiStyle_setIndentSpacing").Call(uintptr(unsafe.Pointer(Self)), uintptr(IndentSpacing))
}

func (i *Imgui) ImGuiStyleColumnsMinSpacing(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ColumnsMinSpacing").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetColumnsMinSpacing(Self *ImGuiStyle, ColumnsMinSpacing float32) {
	getProc("ImGuiStyle_setColumnsMinSpacing").Call(uintptr(unsafe.Pointer(Self)), uintptr(ColumnsMinSpacing))
}

func (i *Imgui) ImGuiStyleScrollbarSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ScrollbarSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetScrollbarSize(Self *ImGuiStyle, ScrollbarSize float32) {
	getProc("ImGuiStyle_setScrollbarSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(ScrollbarSize))
}

func (i *Imgui) ImGuiStyleScrollbarRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ScrollbarRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetScrollbarRounding(Self *ImGuiStyle, ScrollbarRounding float32) {
	getProc("ImGuiStyle_setScrollbarRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(ScrollbarRounding))
}

func (i *Imgui) ImGuiStyleScrollbarPadding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ScrollbarPadding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetScrollbarPadding(Self *ImGuiStyle, ScrollbarPadding float32) {
	getProc("ImGuiStyle_setScrollbarPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(ScrollbarPadding))
}

func (i *Imgui) ImGuiStyleGrabMinSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_GrabMinSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetGrabMinSize(Self *ImGuiStyle, GrabMinSize float32) {
	getProc("ImGuiStyle_setGrabMinSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(GrabMinSize))
}

func (i *Imgui) ImGuiStyleGrabRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_GrabRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetGrabRounding(Self *ImGuiStyle, GrabRounding float32) {
	getProc("ImGuiStyle_setGrabRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(GrabRounding))
}

func (i *Imgui) ImGuiStyleLogSliderDeadzone(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_LogSliderDeadzone").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetLogSliderDeadzone(Self *ImGuiStyle, LogSliderDeadzone float32) {
	getProc("ImGuiStyle_setLogSliderDeadzone").Call(uintptr(unsafe.Pointer(Self)), uintptr(LogSliderDeadzone))
}

func (i *Imgui) ImGuiStyleImageRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ImageRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetImageRounding(Self *ImGuiStyle, ImageRounding float32) {
	getProc("ImGuiStyle_setImageRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(ImageRounding))
}

func (i *Imgui) ImGuiStyleImageBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ImageBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetImageBorderSize(Self *ImGuiStyle, ImageBorderSize float32) {
	getProc("ImGuiStyle_setImageBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(ImageBorderSize))
}

func (i *Imgui) ImGuiStyleTabRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabRounding(Self *ImGuiStyle, TabRounding float32) {
	getProc("ImGuiStyle_setTabRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabRounding))
}

func (i *Imgui) ImGuiStyleTabBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabBorderSize(Self *ImGuiStyle, TabBorderSize float32) {
	getProc("ImGuiStyle_setTabBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabBorderSize))
}

func (i *Imgui) ImGuiStyleTabMinWidthBase(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabMinWidthBase").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabMinWidthBase(Self *ImGuiStyle, TabMinWidthBase float32) {
	getProc("ImGuiStyle_setTabMinWidthBase").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabMinWidthBase))
}

func (i *Imgui) ImGuiStyleTabMinWidthShrink(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabMinWidthShrink").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabMinWidthShrink(Self *ImGuiStyle, TabMinWidthShrink float32) {
	getProc("ImGuiStyle_setTabMinWidthShrink").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabMinWidthShrink))
}

func (i *Imgui) ImGuiStyleTabCloseButtonMinWidthSelected(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabCloseButtonMinWidthSelected").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabCloseButtonMinWidthSelected(Self *ImGuiStyle, TabCloseButtonMinWidthSelected float32) {
	getProc("ImGuiStyle_setTabCloseButtonMinWidthSelected").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabCloseButtonMinWidthSelected))
}

func (i *Imgui) ImGuiStyleTabCloseButtonMinWidthUnselected(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabCloseButtonMinWidthUnselected").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabCloseButtonMinWidthUnselected(Self *ImGuiStyle, TabCloseButtonMinWidthUnselected float32) {
	getProc("ImGuiStyle_setTabCloseButtonMinWidthUnselected").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabCloseButtonMinWidthUnselected))
}

func (i *Imgui) ImGuiStyleTabBarBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabBarBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabBarBorderSize(Self *ImGuiStyle, TabBarBorderSize float32) {
	getProc("ImGuiStyle_setTabBarBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabBarBorderSize))
}

func (i *Imgui) ImGuiStyleTabBarOverlineSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TabBarOverlineSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTabBarOverlineSize(Self *ImGuiStyle, TabBarOverlineSize float32) {
	getProc("ImGuiStyle_setTabBarOverlineSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(TabBarOverlineSize))
}

func (i *Imgui) ImGuiStyleTableAngledHeadersAngle(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TableAngledHeadersAngle").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTableAngledHeadersAngle(Self *ImGuiStyle, TableAngledHeadersAngle float32) {
	getProc("ImGuiStyle_setTableAngledHeadersAngle").Call(uintptr(unsafe.Pointer(Self)), uintptr(TableAngledHeadersAngle))
}

func (i *Imgui) ImGuiStyleTableAngledHeadersTextAlign(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_TableAngledHeadersTextAlign").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetTableAngledHeadersTextAlign(Self *ImGuiStyle, TableAngledHeadersTextAlign *ImVec2) {
	getProc("ImGuiStyle_setTableAngledHeadersTextAlign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TableAngledHeadersTextAlign)))
}

func (i *Imgui) ImGuiStyleTreeLinesFlags(Self *ImGuiStyle) int32 {
	r1, _, _ := getProc("ImGuiStyle_TreeLinesFlags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiStyleSetTreeLinesFlags(Self *ImGuiStyle, TreeLinesFlags int32) {
	getProc("ImGuiStyle_setTreeLinesFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(TreeLinesFlags))
}

func (i *Imgui) ImGuiStyleTreeLinesSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TreeLinesSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTreeLinesSize(Self *ImGuiStyle, TreeLinesSize float32) {
	getProc("ImGuiStyle_setTreeLinesSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(TreeLinesSize))
}

func (i *Imgui) ImGuiStyleTreeLinesRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_TreeLinesRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetTreeLinesRounding(Self *ImGuiStyle, TreeLinesRounding float32) {
	getProc("ImGuiStyle_setTreeLinesRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(TreeLinesRounding))
}

func (i *Imgui) ImGuiStyleDragDropTargetRounding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_DragDropTargetRounding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetDragDropTargetRounding(Self *ImGuiStyle, DragDropTargetRounding float32) {
	getProc("ImGuiStyle_setDragDropTargetRounding").Call(uintptr(unsafe.Pointer(Self)), uintptr(DragDropTargetRounding))
}

func (i *Imgui) ImGuiStyleDragDropTargetBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_DragDropTargetBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetDragDropTargetBorderSize(Self *ImGuiStyle, DragDropTargetBorderSize float32) {
	getProc("ImGuiStyle_setDragDropTargetBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(DragDropTargetBorderSize))
}

func (i *Imgui) ImGuiStyleDragDropTargetPadding(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_DragDropTargetPadding").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetDragDropTargetPadding(Self *ImGuiStyle, DragDropTargetPadding float32) {
	getProc("ImGuiStyle_setDragDropTargetPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(DragDropTargetPadding))
}

func (i *Imgui) ImGuiStyleColorMarkerSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_ColorMarkerSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetColorMarkerSize(Self *ImGuiStyle, ColorMarkerSize float32) {
	getProc("ImGuiStyle_setColorMarkerSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(ColorMarkerSize))
}

func (i *Imgui) ImGuiStyleColorButtonPosition(Self *ImGuiStyle) int32 {
	r1, _, _ := getProc("ImGuiStyle_ColorButtonPosition").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiStyleSetColorButtonPosition(Self *ImGuiStyle, ColorButtonPosition int32) {
	getProc("ImGuiStyle_setColorButtonPosition").Call(uintptr(unsafe.Pointer(Self)), uintptr(ColorButtonPosition))
}

func (i *Imgui) ImGuiStyleButtonTextAlign(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_ButtonTextAlign").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetButtonTextAlign(Self *ImGuiStyle, ButtonTextAlign *ImVec2) {
	getProc("ImGuiStyle_setButtonTextAlign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(ButtonTextAlign)))
}

func (i *Imgui) ImGuiStyleSelectableTextAlign(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_SelectableTextAlign").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetSelectableTextAlign(Self *ImGuiStyle, SelectableTextAlign *ImVec2) {
	getProc("ImGuiStyle_setSelectableTextAlign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(SelectableTextAlign)))
}

func (i *Imgui) ImGuiStyleSeparatorSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_SeparatorSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetSeparatorSize(Self *ImGuiStyle, SeparatorSize float32) {
	getProc("ImGuiStyle_setSeparatorSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(SeparatorSize))
}

func (i *Imgui) ImGuiStyleSeparatorTextBorderSize(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_SeparatorTextBorderSize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetSeparatorTextBorderSize(Self *ImGuiStyle, SeparatorTextBorderSize float32) {
	getProc("ImGuiStyle_setSeparatorTextBorderSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(SeparatorTextBorderSize))
}

func (i *Imgui) ImGuiStyleSeparatorTextAlign(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_SeparatorTextAlign").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetSeparatorTextAlign(Self *ImGuiStyle, SeparatorTextAlign *ImVec2) {
	getProc("ImGuiStyle_setSeparatorTextAlign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(SeparatorTextAlign)))
}

func (i *Imgui) ImGuiStyleSeparatorTextPadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_SeparatorTextPadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetSeparatorTextPadding(Self *ImGuiStyle, SeparatorTextPadding *ImVec2) {
	getProc("ImGuiStyle_setSeparatorTextPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(SeparatorTextPadding)))
}

func (i *Imgui) ImGuiStyleDisplayWindowPadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_DisplayWindowPadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetDisplayWindowPadding(Self *ImGuiStyle, DisplayWindowPadding *ImVec2) {
	getProc("ImGuiStyle_setDisplayWindowPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DisplayWindowPadding)))
}

func (i *Imgui) ImGuiStyleDisplaySafeAreaPadding(Self *ImGuiStyle) *ImVec2 {
	r1, _, _ := getProc("ImGuiStyle_DisplaySafeAreaPadding").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetDisplaySafeAreaPadding(Self *ImGuiStyle, DisplaySafeAreaPadding *ImVec2) {
	getProc("ImGuiStyle_setDisplaySafeAreaPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DisplaySafeAreaPadding)))
}

func (i *Imgui) ImGuiStyleMouseCursorScale(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_MouseCursorScale").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetMouseCursorScale(Self *ImGuiStyle, MouseCursorScale float32) {
	getProc("ImGuiStyle_setMouseCursorScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseCursorScale))
}

func (i *Imgui) ImGuiStyleAntiAliasedLines(Self *ImGuiStyle) bool {
	r1, _, _ := getProc("ImGuiStyle_AntiAliasedLines").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiStyleSetAntiAliasedLines(Self *ImGuiStyle, AntiAliasedLines bool) {
	getProc("ImGuiStyle_setAntiAliasedLines").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if AntiAliasedLines {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiStyleAntiAliasedLinesUseTex(Self *ImGuiStyle) bool {
	r1, _, _ := getProc("ImGuiStyle_AntiAliasedLinesUseTex").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiStyleSetAntiAliasedLinesUseTex(Self *ImGuiStyle, AntiAliasedLinesUseTex bool) {
	getProc("ImGuiStyle_setAntiAliasedLinesUseTex").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if AntiAliasedLinesUseTex {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiStyleAntiAliasedFill(Self *ImGuiStyle) bool {
	r1, _, _ := getProc("ImGuiStyle_AntiAliasedFill").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiStyleSetAntiAliasedFill(Self *ImGuiStyle, AntiAliasedFill bool) {
	getProc("ImGuiStyle_setAntiAliasedFill").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if AntiAliasedFill {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiStyleCurveTessellationTol(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_CurveTessellationTol").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetCurveTessellationTol(Self *ImGuiStyle, CurveTessellationTol float32) {
	getProc("ImGuiStyle_setCurveTessellationTol").Call(uintptr(unsafe.Pointer(Self)), uintptr(CurveTessellationTol))
}

func (i *Imgui) ImGuiStyleCircleTessellationMaxError(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_CircleTessellationMaxError").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetCircleTessellationMaxError(Self *ImGuiStyle, CircleTessellationMaxError float32) {
	getProc("ImGuiStyle_setCircleTessellationMaxError").Call(uintptr(unsafe.Pointer(Self)), uintptr(CircleTessellationMaxError))
}

func (i *Imgui) ImGuiStyleColors(Self *ImGuiStyle) *ImVec4 {
	r1, _, _ := getProc("ImGuiStyle_Colors").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStyleSetColors(Self *ImGuiStyle, Colors *ImVec4) {
	getProc("ImGuiStyle_setColors").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Colors)))
}

func (i *Imgui) ImGuiStyleHoverStationaryDelay(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_HoverStationaryDelay").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetHoverStationaryDelay(Self *ImGuiStyle, HoverStationaryDelay float32) {
	getProc("ImGuiStyle_setHoverStationaryDelay").Call(uintptr(unsafe.Pointer(Self)), uintptr(HoverStationaryDelay))
}

func (i *Imgui) ImGuiStyleHoverDelayShort(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_HoverDelayShort").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetHoverDelayShort(Self *ImGuiStyle, HoverDelayShort float32) {
	getProc("ImGuiStyle_setHoverDelayShort").Call(uintptr(unsafe.Pointer(Self)), uintptr(HoverDelayShort))
}

func (i *Imgui) ImGuiStyleHoverDelayNormal(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle_HoverDelayNormal").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetHoverDelayNormal(Self *ImGuiStyle, HoverDelayNormal float32) {
	getProc("ImGuiStyle_setHoverDelayNormal").Call(uintptr(unsafe.Pointer(Self)), uintptr(HoverDelayNormal))
}

func (i *Imgui) ImGuiStyleHoverFlagsForTooltipMouse(Self *ImGuiStyle) int32 {
	r1, _, _ := getProc("ImGuiStyle_HoverFlagsForTooltipMouse").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiStyleSetHoverFlagsForTooltipMouse(Self *ImGuiStyle, HoverFlagsForTooltipMouse int32) {
	getProc("ImGuiStyle_setHoverFlagsForTooltipMouse").Call(uintptr(unsafe.Pointer(Self)), uintptr(HoverFlagsForTooltipMouse))
}

func (i *Imgui) ImGuiStyleHoverFlagsForTooltipNav(Self *ImGuiStyle) int32 {
	r1, _, _ := getProc("ImGuiStyle_HoverFlagsForTooltipNav").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiStyleSetHoverFlagsForTooltipNav(Self *ImGuiStyle, HoverFlagsForTooltipNav int32) {
	getProc("ImGuiStyle_setHoverFlagsForTooltipNav").Call(uintptr(unsafe.Pointer(Self)), uintptr(HoverFlagsForTooltipNav))
}

func (i *Imgui) ImGuiStyleMainScale(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle__MainScale").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetMainScale(Self *ImGuiStyle, MainScale float32) {
	getProc("ImGuiStyle_set_MainScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(MainScale))
}

func (i *Imgui) ImGuiStyleNextFrameFontSizeBase(Self *ImGuiStyle) float32 {
	r1, _, _ := getProc("ImGuiStyle__NextFrameFontSizeBase").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStyleSetNextFrameFontSizeBase(Self *ImGuiStyle, NextFrameFontSizeBase float32) {
	getProc("ImGuiStyle_set_NextFrameFontSizeBase").Call(uintptr(unsafe.Pointer(Self)), uintptr(NextFrameFontSizeBase))
}

func (i *Imgui) ImGuiStyleScaleAllSizes(Self *ImGuiStyle, Scale_factor float32) {
	getProc("ImGuiStyle_ScaleAllSizes").Call(uintptr(unsafe.Pointer(Self)), uintptr(Scale_factor))
}

func (i *Imgui) ImGuiStyleDelete(Self *ImGuiStyle) {
	getProc("ImGuiStyle_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiKeyDataNew(Param1 *ImGuiKeyData) *ImGuiKeyData {
	r1, _, _ := getProc("ImGuiKeyData_new").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiKeyData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiKeyDataDown(Self *ImGuiKeyData) bool {
	r1, _, _ := getProc("ImGuiKeyData_Down").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiKeyDataSetDown(Self *ImGuiKeyData, Down bool) {
	getProc("ImGuiKeyData_setDown").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Down {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiKeyDataDownDuration(Self *ImGuiKeyData) float32 {
	r1, _, _ := getProc("ImGuiKeyData_DownDuration").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiKeyDataSetDownDuration(Self *ImGuiKeyData, DownDuration float32) {
	getProc("ImGuiKeyData_setDownDuration").Call(uintptr(unsafe.Pointer(Self)), uintptr(DownDuration))
}

func (i *Imgui) ImGuiKeyDataDownDurationPrev(Self *ImGuiKeyData) float32 {
	r1, _, _ := getProc("ImGuiKeyData_DownDurationPrev").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiKeyDataSetDownDurationPrev(Self *ImGuiKeyData, DownDurationPrev float32) {
	getProc("ImGuiKeyData_setDownDurationPrev").Call(uintptr(unsafe.Pointer(Self)), uintptr(DownDurationPrev))
}

func (i *Imgui) ImGuiKeyDataAnalogValue(Self *ImGuiKeyData) float32 {
	r1, _, _ := getProc("ImGuiKeyData_AnalogValue").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiKeyDataSetAnalogValue(Self *ImGuiKeyData, AnalogValue float32) {
	getProc("ImGuiKeyData_setAnalogValue").Call(uintptr(unsafe.Pointer(Self)), uintptr(AnalogValue))
}

func (i *Imgui) ImGuiKeyDataOperatorAssign(Self *ImGuiKeyData, Param1 *ImGuiKeyData) {
	getProc("ImGuiKeyData_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiKeyDataDelete(Self *ImGuiKeyData) {
	getProc("ImGuiKeyData_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiIONew() *ImGuiIO {
	r1, _, _ := getProc("ImGuiIO_new").Call()
	return (*ImGuiIO)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIONew2(Param1 *ImGuiIO) *ImGuiIO {
	r1, _, _ := getProc("ImGuiIO_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiIO)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOConfigFlags(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_ConfigFlags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetConfigFlags(Self *ImGuiIO, ConfigFlags int32) {
	getProc("ImGuiIO_setConfigFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(ConfigFlags))
}

func (i *Imgui) ImGuiIOBackendFlags(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_BackendFlags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetBackendFlags(Self *ImGuiIO, BackendFlags int32) {
	getProc("ImGuiIO_setBackendFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(BackendFlags))
}

func (i *Imgui) ImGuiIODisplaySize(Self *ImGuiIO) *ImVec2 {
	r1, _, _ := getProc("ImGuiIO_DisplaySize").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetDisplaySize(Self *ImGuiIO, DisplaySize *ImVec2) {
	getProc("ImGuiIO_setDisplaySize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DisplaySize)))
}

func (i *Imgui) ImGuiIODisplayFramebufferScale(Self *ImGuiIO) *ImVec2 {
	r1, _, _ := getProc("ImGuiIO_DisplayFramebufferScale").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetDisplayFramebufferScale(Self *ImGuiIO, DisplayFramebufferScale *ImVec2) {
	getProc("ImGuiIO_setDisplayFramebufferScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DisplayFramebufferScale)))
}

func (i *Imgui) ImGuiIODeltaTime(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_DeltaTime").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetDeltaTime(Self *ImGuiIO, DeltaTime float32) {
	getProc("ImGuiIO_setDeltaTime").Call(uintptr(unsafe.Pointer(Self)), uintptr(DeltaTime))
}

func (i *Imgui) ImGuiIOIniSavingRate(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_IniSavingRate").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetIniSavingRate(Self *ImGuiIO, IniSavingRate float32) {
	getProc("ImGuiIO_setIniSavingRate").Call(uintptr(unsafe.Pointer(Self)), uintptr(IniSavingRate))
}

func (i *Imgui) ImGuiIOIniFilename(Self *ImGuiIO) *int8 {
	r1, _, _ := getProc("ImGuiIO_IniFilename").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetIniFilename(Self *ImGuiIO, IniFilename *int8) {
	getProc("ImGuiIO_setIniFilename").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(IniFilename)))
}

func (i *Imgui) ImGuiIOLogFilename(Self *ImGuiIO) *int8 {
	r1, _, _ := getProc("ImGuiIO_LogFilename").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetLogFilename(Self *ImGuiIO, LogFilename *int8) {
	getProc("ImGuiIO_setLogFilename").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(LogFilename)))
}

func (i *Imgui) ImGuiIOFonts(Self *ImGuiIO) *ImFontAtlas {
	r1, _, _ := getProc("ImGuiIO_Fonts").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFontAtlas)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetFonts(Self *ImGuiIO, Fonts *ImFontAtlas) {
	getProc("ImGuiIO_setFonts").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Fonts)))
}

func (i *Imgui) ImGuiIOFontDefault(Self *ImGuiIO) *ImFont {
	r1, _, _ := getProc("ImGuiIO_FontDefault").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetFontDefault(Self *ImGuiIO, FontDefault *ImFont) {
	getProc("ImGuiIO_setFontDefault").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FontDefault)))
}

func (i *Imgui) ImGuiIOFontAllowUserScaling(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_FontAllowUserScaling").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetFontAllowUserScaling(Self *ImGuiIO, FontAllowUserScaling bool) {
	getProc("ImGuiIO_setFontAllowUserScaling").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if FontAllowUserScaling {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavSwapGamepadButtons(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavSwapGamepadButtons").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavSwapGamepadButtons(Self *ImGuiIO, ConfigNavSwapGamepadButtons bool) {
	getProc("ImGuiIO_setConfigNavSwapGamepadButtons").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavSwapGamepadButtons {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavMoveSetMousePos(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavMoveSetMousePos").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavMoveSetMousePos(Self *ImGuiIO, ConfigNavMoveSetMousePos bool) {
	getProc("ImGuiIO_setConfigNavMoveSetMousePos").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavMoveSetMousePos {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavCaptureKeyboard(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavCaptureKeyboard").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavCaptureKeyboard(Self *ImGuiIO, ConfigNavCaptureKeyboard bool) {
	getProc("ImGuiIO_setConfigNavCaptureKeyboard").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavCaptureKeyboard {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavEscapeClearFocusItem(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavEscapeClearFocusItem").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavEscapeClearFocusItem(Self *ImGuiIO, ConfigNavEscapeClearFocusItem bool) {
	getProc("ImGuiIO_setConfigNavEscapeClearFocusItem").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavEscapeClearFocusItem {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavEscapeClearFocusWindow(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavEscapeClearFocusWindow").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavEscapeClearFocusWindow(Self *ImGuiIO, ConfigNavEscapeClearFocusWindow bool) {
	getProc("ImGuiIO_setConfigNavEscapeClearFocusWindow").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavEscapeClearFocusWindow {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavCursorVisibleAuto(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavCursorVisibleAuto").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavCursorVisibleAuto(Self *ImGuiIO, ConfigNavCursorVisibleAuto bool) {
	getProc("ImGuiIO_setConfigNavCursorVisibleAuto").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavCursorVisibleAuto {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigNavCursorVisibleAlways(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigNavCursorVisibleAlways").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigNavCursorVisibleAlways(Self *ImGuiIO, ConfigNavCursorVisibleAlways bool) {
	getProc("ImGuiIO_setConfigNavCursorVisibleAlways").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigNavCursorVisibleAlways {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOMouseDrawCursor(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_MouseDrawCursor").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetMouseDrawCursor(Self *ImGuiIO, MouseDrawCursor bool) {
	getProc("ImGuiIO_setMouseDrawCursor").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if MouseDrawCursor {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigMacOSXBehaviors(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigMacOSXBehaviors").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigMacOSXBehaviors(Self *ImGuiIO, ConfigMacOSXBehaviors bool) {
	getProc("ImGuiIO_setConfigMacOSXBehaviors").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigMacOSXBehaviors {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigInputTrickleEventQueue(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigInputTrickleEventQueue").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigInputTrickleEventQueue(Self *ImGuiIO, ConfigInputTrickleEventQueue bool) {
	getProc("ImGuiIO_setConfigInputTrickleEventQueue").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigInputTrickleEventQueue {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigInputTextCursorBlink(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigInputTextCursorBlink").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigInputTextCursorBlink(Self *ImGuiIO, ConfigInputTextCursorBlink bool) {
	getProc("ImGuiIO_setConfigInputTextCursorBlink").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigInputTextCursorBlink {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigInputTextEnterKeepActive(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigInputTextEnterKeepActive").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigInputTextEnterKeepActive(Self *ImGuiIO, ConfigInputTextEnterKeepActive bool) {
	getProc("ImGuiIO_setConfigInputTextEnterKeepActive").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigInputTextEnterKeepActive {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDragClickToInputText(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDragClickToInputText").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDragClickToInputText(Self *ImGuiIO, ConfigDragClickToInputText bool) {
	getProc("ImGuiIO_setConfigDragClickToInputText").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDragClickToInputText {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigWindowsResizeFromEdges(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigWindowsResizeFromEdges").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigWindowsResizeFromEdges(Self *ImGuiIO, ConfigWindowsResizeFromEdges bool) {
	getProc("ImGuiIO_setConfigWindowsResizeFromEdges").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigWindowsResizeFromEdges {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigWindowsMoveFromTitleBarOnly(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigWindowsMoveFromTitleBarOnly").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigWindowsMoveFromTitleBarOnly(Self *ImGuiIO, ConfigWindowsMoveFromTitleBarOnly bool) {
	getProc("ImGuiIO_setConfigWindowsMoveFromTitleBarOnly").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigWindowsMoveFromTitleBarOnly {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigWindowsCopyContentsWithCtrlC(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigWindowsCopyContentsWithCtrlC").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigWindowsCopyContentsWithCtrlC(Self *ImGuiIO, ConfigWindowsCopyContentsWithCtrlC bool) {
	getProc("ImGuiIO_setConfigWindowsCopyContentsWithCtrlC").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigWindowsCopyContentsWithCtrlC {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigScrollbarScrollByPage(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigScrollbarScrollByPage").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigScrollbarScrollByPage(Self *ImGuiIO, ConfigScrollbarScrollByPage bool) {
	getProc("ImGuiIO_setConfigScrollbarScrollByPage").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigScrollbarScrollByPage {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigMemoryCompactTimer(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_ConfigMemoryCompactTimer").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetConfigMemoryCompactTimer(Self *ImGuiIO, ConfigMemoryCompactTimer float32) {
	getProc("ImGuiIO_setConfigMemoryCompactTimer").Call(uintptr(unsafe.Pointer(Self)), uintptr(ConfigMemoryCompactTimer))
}

func (i *Imgui) ImGuiIOMouseDoubleClickTime(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_MouseDoubleClickTime").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetMouseDoubleClickTime(Self *ImGuiIO, MouseDoubleClickTime float32) {
	getProc("ImGuiIO_setMouseDoubleClickTime").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseDoubleClickTime))
}

func (i *Imgui) ImGuiIOMouseDoubleClickMaxDist(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_MouseDoubleClickMaxDist").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetMouseDoubleClickMaxDist(Self *ImGuiIO, MouseDoubleClickMaxDist float32) {
	getProc("ImGuiIO_setMouseDoubleClickMaxDist").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseDoubleClickMaxDist))
}

func (i *Imgui) ImGuiIOMouseDragThreshold(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_MouseDragThreshold").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetMouseDragThreshold(Self *ImGuiIO, MouseDragThreshold float32) {
	getProc("ImGuiIO_setMouseDragThreshold").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseDragThreshold))
}

func (i *Imgui) ImGuiIOKeyRepeatDelay(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_KeyRepeatDelay").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetKeyRepeatDelay(Self *ImGuiIO, KeyRepeatDelay float32) {
	getProc("ImGuiIO_setKeyRepeatDelay").Call(uintptr(unsafe.Pointer(Self)), uintptr(KeyRepeatDelay))
}

func (i *Imgui) ImGuiIOKeyRepeatRate(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_KeyRepeatRate").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetKeyRepeatRate(Self *ImGuiIO, KeyRepeatRate float32) {
	getProc("ImGuiIO_setKeyRepeatRate").Call(uintptr(unsafe.Pointer(Self)), uintptr(KeyRepeatRate))
}

func (i *Imgui) ImGuiIOConfigErrorRecovery(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigErrorRecovery").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigErrorRecovery(Self *ImGuiIO, ConfigErrorRecovery bool) {
	getProc("ImGuiIO_setConfigErrorRecovery").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigErrorRecovery {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigErrorRecoveryEnableAssert(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigErrorRecoveryEnableAssert").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigErrorRecoveryEnableAssert(Self *ImGuiIO, ConfigErrorRecoveryEnableAssert bool) {
	getProc("ImGuiIO_setConfigErrorRecoveryEnableAssert").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigErrorRecoveryEnableAssert {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigErrorRecoveryEnableDebugLog(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigErrorRecoveryEnableDebugLog").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigErrorRecoveryEnableDebugLog(Self *ImGuiIO, ConfigErrorRecoveryEnableDebugLog bool) {
	getProc("ImGuiIO_setConfigErrorRecoveryEnableDebugLog").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigErrorRecoveryEnableDebugLog {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigErrorRecoveryEnableTooltip(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigErrorRecoveryEnableTooltip").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigErrorRecoveryEnableTooltip(Self *ImGuiIO, ConfigErrorRecoveryEnableTooltip bool) {
	getProc("ImGuiIO_setConfigErrorRecoveryEnableTooltip").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigErrorRecoveryEnableTooltip {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugIsDebuggerPresent(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugIsDebuggerPresent").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugIsDebuggerPresent(Self *ImGuiIO, ConfigDebugIsDebuggerPresent bool) {
	getProc("ImGuiIO_setConfigDebugIsDebuggerPresent").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugIsDebuggerPresent {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugHighlightIdConflicts(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugHighlightIdConflicts").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugHighlightIdConflicts(Self *ImGuiIO, ConfigDebugHighlightIdConflicts bool) {
	getProc("ImGuiIO_setConfigDebugHighlightIdConflicts").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugHighlightIdConflicts {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugHighlightIdConflictsShowItemPicker(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugHighlightIdConflictsShowItemPicker").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugHighlightIdConflictsShowItemPicker(Self *ImGuiIO, ConfigDebugHighlightIdConflictsShowItemPicker bool) {
	getProc("ImGuiIO_setConfigDebugHighlightIdConflictsShowItemPicker").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugHighlightIdConflictsShowItemPicker {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugBeginReturnValueOnce(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugBeginReturnValueOnce").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugBeginReturnValueOnce(Self *ImGuiIO, ConfigDebugBeginReturnValueOnce bool) {
	getProc("ImGuiIO_setConfigDebugBeginReturnValueOnce").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugBeginReturnValueOnce {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugBeginReturnValueLoop(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugBeginReturnValueLoop").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugBeginReturnValueLoop(Self *ImGuiIO, ConfigDebugBeginReturnValueLoop bool) {
	getProc("ImGuiIO_setConfigDebugBeginReturnValueLoop").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugBeginReturnValueLoop {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugIgnoreFocusLoss(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugIgnoreFocusLoss").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugIgnoreFocusLoss(Self *ImGuiIO, ConfigDebugIgnoreFocusLoss bool) {
	getProc("ImGuiIO_setConfigDebugIgnoreFocusLoss").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugIgnoreFocusLoss {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOConfigDebugIniSettings(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_ConfigDebugIniSettings").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetConfigDebugIniSettings(Self *ImGuiIO, ConfigDebugIniSettings bool) {
	getProc("ImGuiIO_setConfigDebugIniSettings").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if ConfigDebugIniSettings {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOBackendPlatformName(Self *ImGuiIO) *int8 {
	r1, _, _ := getProc("ImGuiIO_BackendPlatformName").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetBackendPlatformName(Self *ImGuiIO, BackendPlatformName *int8) {
	getProc("ImGuiIO_setBackendPlatformName").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(BackendPlatformName)))
}

func (i *Imgui) ImGuiIOBackendRendererName(Self *ImGuiIO) *int8 {
	r1, _, _ := getProc("ImGuiIO_BackendRendererName").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetBackendRendererName(Self *ImGuiIO, BackendRendererName *int8) {
	getProc("ImGuiIO_setBackendRendererName").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(BackendRendererName)))
}

func (i *Imgui) ImGuiIOAddKeyEvent(Self *ImGuiIO, Key int32, Down bool) {
	getProc("ImGuiIO_AddKeyEvent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), func() uintptr {
		if Down {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOAddKeyAnalogEvent(Self *ImGuiIO, Key int32, Down bool, V float32) {
	getProc("ImGuiIO_AddKeyAnalogEvent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), func() uintptr {
		if Down {
			return 1
		}
		return 0
	}(), uintptr(V))
}

func (i *Imgui) ImGuiIOAddMousePosEvent(Self *ImGuiIO, X float32, Y float32) {
	getProc("ImGuiIO_AddMousePosEvent").Call(uintptr(unsafe.Pointer(Self)), uintptr(X), uintptr(Y))
}

func (i *Imgui) ImGuiIOAddMouseButtonEvent(Self *ImGuiIO, Button int32, Down bool) {
	getProc("ImGuiIO_AddMouseButtonEvent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Button), func() uintptr {
		if Down {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOAddMouseWheelEvent(Self *ImGuiIO, Wheel_x float32, Wheel_y float32) {
	getProc("ImGuiIO_AddMouseWheelEvent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Wheel_x), uintptr(Wheel_y))
}

func (i *Imgui) ImGuiIOAddMouseSourceEvent(Self *ImGuiIO, Source int32) {
	getProc("ImGuiIO_AddMouseSourceEvent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Source))
}

func (i *Imgui) ImGuiIOAddFocusEvent(Self *ImGuiIO, Focused bool) {
	getProc("ImGuiIO_AddFocusEvent").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Focused {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOAddInputCharacter(Self *ImGuiIO, C uint32) {
	getProc("ImGuiIO_AddInputCharacter").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
}

func (i *Imgui) ImGuiIOAddInputCharacterUTF16(Self *ImGuiIO, C uint16) {
	getProc("ImGuiIO_AddInputCharacterUTF16").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
}

func (i *Imgui) ImGuiIOAddInputCharactersUTF8(Self *ImGuiIO, Str *int8) {
	getProc("ImGuiIO_AddInputCharactersUTF8").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Str)))
}

func (i *Imgui) ImGuiIOSetKeyEventNativeData(Self *ImGuiIO, Key int32, Native_keycode int32, Native_scancode int32) {
	getProc("ImGuiIO_SetKeyEventNativeData").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Native_keycode), uintptr(Native_scancode))
}

func (i *Imgui) ImGuiIOSetAppAcceptingEvents(Self *ImGuiIO, Accepting_events bool) {
	getProc("ImGuiIO_SetAppAcceptingEvents").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Accepting_events {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOClearEventsQueue(Self *ImGuiIO) {
	getProc("ImGuiIO_ClearEventsQueue").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiIOClearInputKeys(Self *ImGuiIO) {
	getProc("ImGuiIO_ClearInputKeys").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiIOClearInputMouse(Self *ImGuiIO) {
	getProc("ImGuiIO_ClearInputMouse").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiIOWantCaptureMouse(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_WantCaptureMouse").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetWantCaptureMouse(Self *ImGuiIO, WantCaptureMouse bool) {
	getProc("ImGuiIO_setWantCaptureMouse").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantCaptureMouse {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOWantCaptureKeyboard(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_WantCaptureKeyboard").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetWantCaptureKeyboard(Self *ImGuiIO, WantCaptureKeyboard bool) {
	getProc("ImGuiIO_setWantCaptureKeyboard").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantCaptureKeyboard {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOWantTextInput(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_WantTextInput").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetWantTextInput(Self *ImGuiIO, WantTextInput bool) {
	getProc("ImGuiIO_setWantTextInput").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantTextInput {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOWantSetMousePos(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_WantSetMousePos").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetWantSetMousePos(Self *ImGuiIO, WantSetMousePos bool) {
	getProc("ImGuiIO_setWantSetMousePos").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantSetMousePos {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOWantSaveIniSettings(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_WantSaveIniSettings").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetWantSaveIniSettings(Self *ImGuiIO, WantSaveIniSettings bool) {
	getProc("ImGuiIO_setWantSaveIniSettings").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantSaveIniSettings {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIONavActive(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_NavActive").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetNavActive(Self *ImGuiIO, NavActive bool) {
	getProc("ImGuiIO_setNavActive").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if NavActive {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIONavVisible(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_NavVisible").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetNavVisible(Self *ImGuiIO, NavVisible bool) {
	getProc("ImGuiIO_setNavVisible").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if NavVisible {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOFramerate(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_Framerate").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetFramerate(Self *ImGuiIO, Framerate float32) {
	getProc("ImGuiIO_setFramerate").Call(uintptr(unsafe.Pointer(Self)), uintptr(Framerate))
}

func (i *Imgui) ImGuiIOMetricsRenderVertices(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_MetricsRenderVertices").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetMetricsRenderVertices(Self *ImGuiIO, MetricsRenderVertices int32) {
	getProc("ImGuiIO_setMetricsRenderVertices").Call(uintptr(unsafe.Pointer(Self)), uintptr(MetricsRenderVertices))
}

func (i *Imgui) ImGuiIOMetricsRenderIndices(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_MetricsRenderIndices").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetMetricsRenderIndices(Self *ImGuiIO, MetricsRenderIndices int32) {
	getProc("ImGuiIO_setMetricsRenderIndices").Call(uintptr(unsafe.Pointer(Self)), uintptr(MetricsRenderIndices))
}

func (i *Imgui) ImGuiIOMetricsRenderWindows(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_MetricsRenderWindows").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetMetricsRenderWindows(Self *ImGuiIO, MetricsRenderWindows int32) {
	getProc("ImGuiIO_setMetricsRenderWindows").Call(uintptr(unsafe.Pointer(Self)), uintptr(MetricsRenderWindows))
}

func (i *Imgui) ImGuiIOMetricsActiveWindows(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_MetricsActiveWindows").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetMetricsActiveWindows(Self *ImGuiIO, MetricsActiveWindows int32) {
	getProc("ImGuiIO_setMetricsActiveWindows").Call(uintptr(unsafe.Pointer(Self)), uintptr(MetricsActiveWindows))
}

func (i *Imgui) ImGuiIOMouseDelta(Self *ImGuiIO) *ImVec2 {
	r1, _, _ := getProc("ImGuiIO_MouseDelta").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDelta(Self *ImGuiIO, MouseDelta *ImVec2) {
	getProc("ImGuiIO_setMouseDelta").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDelta)))
}

func (i *Imgui) ImGuiIOCtx(Self *ImGuiIO) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiIO_Ctx").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiIOSetCtx(Self *ImGuiIO, Ctx unsafe.Pointer) {
	getProc("ImGuiIO_setCtx").Call(uintptr(unsafe.Pointer(Self)), uintptr(Ctx))
}

func (i *Imgui) ImGuiIOMousePos(Self *ImGuiIO) *ImVec2 {
	r1, _, _ := getProc("ImGuiIO_MousePos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMousePos(Self *ImGuiIO, MousePos *ImVec2) {
	getProc("ImGuiIO_setMousePos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MousePos)))
}

func (i *Imgui) ImGuiIOMouseDown(Self *ImGuiIO) *bool {
	r1, _, _ := getProc("ImGuiIO_MouseDown").Call(uintptr(unsafe.Pointer(Self)))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDown(Self *ImGuiIO, MouseDown *bool) {
	getProc("ImGuiIO_setMouseDown").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDown)))
}

func (i *Imgui) ImGuiIOMouseWheel(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_MouseWheel").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetMouseWheel(Self *ImGuiIO, MouseWheel float32) {
	getProc("ImGuiIO_setMouseWheel").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseWheel))
}

func (i *Imgui) ImGuiIOMouseWheelH(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_MouseWheelH").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetMouseWheelH(Self *ImGuiIO, MouseWheelH float32) {
	getProc("ImGuiIO_setMouseWheelH").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseWheelH))
}

func (i *Imgui) ImGuiIOMouseSource(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_MouseSource").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetMouseSource(Self *ImGuiIO, MouseSource int32) {
	getProc("ImGuiIO_setMouseSource").Call(uintptr(unsafe.Pointer(Self)), uintptr(MouseSource))
}

func (i *Imgui) ImGuiIOKeyCtrl(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_KeyCtrl").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetKeyCtrl(Self *ImGuiIO, KeyCtrl bool) {
	getProc("ImGuiIO_setKeyCtrl").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if KeyCtrl {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOKeyShift(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_KeyShift").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetKeyShift(Self *ImGuiIO, KeyShift bool) {
	getProc("ImGuiIO_setKeyShift").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if KeyShift {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOKeyAlt(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_KeyAlt").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetKeyAlt(Self *ImGuiIO, KeyAlt bool) {
	getProc("ImGuiIO_setKeyAlt").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if KeyAlt {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOKeySuper(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_KeySuper").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetKeySuper(Self *ImGuiIO, KeySuper bool) {
	getProc("ImGuiIO_setKeySuper").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if KeySuper {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOKeyMods(Self *ImGuiIO) int32 {
	r1, _, _ := getProc("ImGuiIO_KeyMods").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiIOSetKeyMods(Self *ImGuiIO, KeyMods int32) {
	getProc("ImGuiIO_setKeyMods").Call(uintptr(unsafe.Pointer(Self)), uintptr(KeyMods))
}

func (i *Imgui) ImGuiIOKeysData(Self *ImGuiIO) *ImGuiKeyData {
	r1, _, _ := getProc("ImGuiIO_KeysData").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImGuiKeyData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetKeysData(Self *ImGuiIO, KeysData *ImGuiKeyData) {
	getProc("ImGuiIO_setKeysData").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(KeysData)))
}

func (i *Imgui) ImGuiIOWantCaptureMouseUnlessPopupClose(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_WantCaptureMouseUnlessPopupClose").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetWantCaptureMouseUnlessPopupClose(Self *ImGuiIO, WantCaptureMouseUnlessPopupClose bool) {
	getProc("ImGuiIO_setWantCaptureMouseUnlessPopupClose").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantCaptureMouseUnlessPopupClose {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOMousePosPrev(Self *ImGuiIO) *ImVec2 {
	r1, _, _ := getProc("ImGuiIO_MousePosPrev").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMousePosPrev(Self *ImGuiIO, MousePosPrev *ImVec2) {
	getProc("ImGuiIO_setMousePosPrev").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MousePosPrev)))
}

func (i *Imgui) ImGuiIOMouseClickedPos(Self *ImGuiIO) *ImVec2 {
	r1, _, _ := getProc("ImGuiIO_MouseClickedPos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseClickedPos(Self *ImGuiIO, MouseClickedPos *ImVec2) {
	getProc("ImGuiIO_setMouseClickedPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseClickedPos)))
}

func (i *Imgui) ImGuiIOMouseClickedTime(Self *ImGuiIO) *float64 {
	r1, _, _ := getProc("ImGuiIO_MouseClickedTime").Call(uintptr(unsafe.Pointer(Self)))
	return (*float64)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseClickedTime(Self *ImGuiIO, MouseClickedTime *float64) {
	getProc("ImGuiIO_setMouseClickedTime").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseClickedTime)))
}

func (i *Imgui) ImGuiIOMouseClicked(Self *ImGuiIO) *bool {
	r1, _, _ := getProc("ImGuiIO_MouseClicked").Call(uintptr(unsafe.Pointer(Self)))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseClicked(Self *ImGuiIO, MouseClicked *bool) {
	getProc("ImGuiIO_setMouseClicked").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseClicked)))
}

func (i *Imgui) ImGuiIOMouseDoubleClicked(Self *ImGuiIO) *bool {
	r1, _, _ := getProc("ImGuiIO_MouseDoubleClicked").Call(uintptr(unsafe.Pointer(Self)))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDoubleClicked(Self *ImGuiIO, MouseDoubleClicked *bool) {
	getProc("ImGuiIO_setMouseDoubleClicked").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDoubleClicked)))
}

func (i *Imgui) ImGuiIOMouseClickedCount(Self *ImGuiIO) *ImU16 {
	r1, _, _ := getProc("ImGuiIO_MouseClickedCount").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImU16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseClickedCount(Self *ImGuiIO, MouseClickedCount *ImU16) {
	getProc("ImGuiIO_setMouseClickedCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseClickedCount)))
}

func (i *Imgui) ImGuiIOMouseClickedLastCount(Self *ImGuiIO) *ImU16 {
	r1, _, _ := getProc("ImGuiIO_MouseClickedLastCount").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImU16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseClickedLastCount(Self *ImGuiIO, MouseClickedLastCount *ImU16) {
	getProc("ImGuiIO_setMouseClickedLastCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseClickedLastCount)))
}

func (i *Imgui) ImGuiIOMouseReleased(Self *ImGuiIO) *bool {
	r1, _, _ := getProc("ImGuiIO_MouseReleased").Call(uintptr(unsafe.Pointer(Self)))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseReleased(Self *ImGuiIO, MouseReleased *bool) {
	getProc("ImGuiIO_setMouseReleased").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseReleased)))
}

func (i *Imgui) ImGuiIOMouseReleasedTime(Self *ImGuiIO) *float64 {
	r1, _, _ := getProc("ImGuiIO_MouseReleasedTime").Call(uintptr(unsafe.Pointer(Self)))
	return (*float64)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseReleasedTime(Self *ImGuiIO, MouseReleasedTime *float64) {
	getProc("ImGuiIO_setMouseReleasedTime").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseReleasedTime)))
}

func (i *Imgui) ImGuiIOMouseDownOwned(Self *ImGuiIO) *bool {
	r1, _, _ := getProc("ImGuiIO_MouseDownOwned").Call(uintptr(unsafe.Pointer(Self)))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDownOwned(Self *ImGuiIO, MouseDownOwned *bool) {
	getProc("ImGuiIO_setMouseDownOwned").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDownOwned)))
}

func (i *Imgui) ImGuiIOMouseDownOwnedUnlessPopupClose(Self *ImGuiIO) *bool {
	r1, _, _ := getProc("ImGuiIO_MouseDownOwnedUnlessPopupClose").Call(uintptr(unsafe.Pointer(Self)))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDownOwnedUnlessPopupClose(Self *ImGuiIO, MouseDownOwnedUnlessPopupClose *bool) {
	getProc("ImGuiIO_setMouseDownOwnedUnlessPopupClose").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDownOwnedUnlessPopupClose)))
}

func (i *Imgui) ImGuiIOMouseWheelRequestAxisSwap(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_MouseWheelRequestAxisSwap").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetMouseWheelRequestAxisSwap(Self *ImGuiIO, MouseWheelRequestAxisSwap bool) {
	getProc("ImGuiIO_setMouseWheelRequestAxisSwap").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if MouseWheelRequestAxisSwap {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOMouseCtrlLeftAsRightClick(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_MouseCtrlLeftAsRightClick").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetMouseCtrlLeftAsRightClick(Self *ImGuiIO, MouseCtrlLeftAsRightClick bool) {
	getProc("ImGuiIO_setMouseCtrlLeftAsRightClick").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if MouseCtrlLeftAsRightClick {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOMouseDownDuration(Self *ImGuiIO) *float32 {
	r1, _, _ := getProc("ImGuiIO_MouseDownDuration").Call(uintptr(unsafe.Pointer(Self)))
	return (*float32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDownDuration(Self *ImGuiIO, MouseDownDuration *float32) {
	getProc("ImGuiIO_setMouseDownDuration").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDownDuration)))
}

func (i *Imgui) ImGuiIOMouseDownDurationPrev(Self *ImGuiIO) *float32 {
	r1, _, _ := getProc("ImGuiIO_MouseDownDurationPrev").Call(uintptr(unsafe.Pointer(Self)))
	return (*float32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDownDurationPrev(Self *ImGuiIO, MouseDownDurationPrev *float32) {
	getProc("ImGuiIO_setMouseDownDurationPrev").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDownDurationPrev)))
}

func (i *Imgui) ImGuiIOMouseDragMaxDistanceSqr(Self *ImGuiIO) *float32 {
	r1, _, _ := getProc("ImGuiIO_MouseDragMaxDistanceSqr").Call(uintptr(unsafe.Pointer(Self)))
	return (*float32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiIOSetMouseDragMaxDistanceSqr(Self *ImGuiIO, MouseDragMaxDistanceSqr *float32) {
	getProc("ImGuiIO_setMouseDragMaxDistanceSqr").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(MouseDragMaxDistanceSqr)))
}

func (i *Imgui) ImGuiIOPenPressure(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_PenPressure").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetPenPressure(Self *ImGuiIO, PenPressure float32) {
	getProc("ImGuiIO_setPenPressure").Call(uintptr(unsafe.Pointer(Self)), uintptr(PenPressure))
}

func (i *Imgui) ImGuiIOAppFocusLost(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_AppFocusLost").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetAppFocusLost(Self *ImGuiIO, AppFocusLost bool) {
	getProc("ImGuiIO_setAppFocusLost").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if AppFocusLost {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOAppAcceptingEvents(Self *ImGuiIO) bool {
	r1, _, _ := getProc("ImGuiIO_AppAcceptingEvents").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiIOSetAppAcceptingEvents_1(Self *ImGuiIO, AppAcceptingEvents bool) {
	getProc("ImGuiIO_setAppAcceptingEvents").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if AppAcceptingEvents {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiIOInputQueueSurrogate(Self *ImGuiIO) uint16 {
	r1, _, _ := getProc("ImGuiIO_InputQueueSurrogate").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImGuiIOSetInputQueueSurrogate(Self *ImGuiIO, InputQueueSurrogate uint16) {
	getProc("ImGuiIO_setInputQueueSurrogate").Call(uintptr(unsafe.Pointer(Self)), uintptr(InputQueueSurrogate))
}

func (i *Imgui) ImGuiIOInputQueueCharacters(Self *ImGuiIO) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiIO_InputQueueCharacters").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiIOSetInputQueueCharacters(Self *ImGuiIO, InputQueueCharacters unsafe.Pointer) {
	getProc("ImGuiIO_setInputQueueCharacters").Call(uintptr(unsafe.Pointer(Self)), uintptr(InputQueueCharacters))
}

func (i *Imgui) ImGuiIOFontGlobalScale(Self *ImGuiIO) float32 {
	r1, _, _ := getProc("ImGuiIO_FontGlobalScale").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiIOSetFontGlobalScale(Self *ImGuiIO, FontGlobalScale float32) {
	getProc("ImGuiIO_setFontGlobalScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontGlobalScale))
}

func (i *Imgui) ImGuiIOOperatorAssign(Self *ImGuiIO, Param1 *ImGuiIO) {
	getProc("ImGuiIO_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiIOSetKeyEventNativeData2(Self *ImGuiIO, Key int32, Native_keycode int32, Native_scancode int32, Native_legacy_index int32) {
	getProc("ImGuiIO_SetKeyEventNativeData2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Native_keycode), uintptr(Native_scancode), uintptr(Native_legacy_index))
}

func (i *Imgui) ImGuiIODelete(Self *ImGuiIO) {
	getProc("ImGuiIO_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiInputTextCallbackDataNew() *ImGuiInputTextCallbackData {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_new").Call()
	return (*ImGuiInputTextCallbackData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiInputTextCallbackDataCtx(Self *ImGuiInputTextCallbackData) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_Ctx").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetCtx(Self *ImGuiInputTextCallbackData, Ctx unsafe.Pointer) {
	getProc("ImGuiInputTextCallbackData_setCtx").Call(uintptr(unsafe.Pointer(Self)), uintptr(Ctx))
}

func (i *Imgui) ImGuiInputTextCallbackDataEventFlag(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_EventFlag").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetEventFlag(Self *ImGuiInputTextCallbackData, EventFlag int32) {
	getProc("ImGuiInputTextCallbackData_setEventFlag").Call(uintptr(unsafe.Pointer(Self)), uintptr(EventFlag))
}

func (i *Imgui) ImGuiInputTextCallbackDataFlags(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetFlags(Self *ImGuiInputTextCallbackData, Flags int32) {
	getProc("ImGuiInputTextCallbackData_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImGuiInputTextCallbackDataID(Self *ImGuiInputTextCallbackData) uint32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_ID").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetID(Self *ImGuiInputTextCallbackData, ID uint32) {
	getProc("ImGuiInputTextCallbackData_setID").Call(uintptr(unsafe.Pointer(Self)), uintptr(ID))
}

func (i *Imgui) ImGuiInputTextCallbackDataEventKey(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_EventKey").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetEventKey(Self *ImGuiInputTextCallbackData, EventKey int32) {
	getProc("ImGuiInputTextCallbackData_setEventKey").Call(uintptr(unsafe.Pointer(Self)), uintptr(EventKey))
}

func (i *Imgui) ImGuiInputTextCallbackDataEventChar(Self *ImGuiInputTextCallbackData) uint16 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_EventChar").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetEventChar(Self *ImGuiInputTextCallbackData, EventChar uint16) {
	getProc("ImGuiInputTextCallbackData_setEventChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(EventChar))
}

func (i *Imgui) ImGuiInputTextCallbackDataEventActivated(Self *ImGuiInputTextCallbackData) bool {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_EventActivated").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiInputTextCallbackDataSetEventActivated(Self *ImGuiInputTextCallbackData, EventActivated bool) {
	getProc("ImGuiInputTextCallbackData_setEventActivated").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if EventActivated {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiInputTextCallbackDataBufDirty(Self *ImGuiInputTextCallbackData) bool {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_BufDirty").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiInputTextCallbackDataSetBufDirty(Self *ImGuiInputTextCallbackData, BufDirty bool) {
	getProc("ImGuiInputTextCallbackData_setBufDirty").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if BufDirty {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiInputTextCallbackDataBuf(Self *ImGuiInputTextCallbackData) *int8 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_Buf").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiInputTextCallbackDataSetBuf(Self *ImGuiInputTextCallbackData, Buf *int8) {
	getProc("ImGuiInputTextCallbackData_setBuf").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Buf)))
}

func (i *Imgui) ImGuiInputTextCallbackDataBufTextLen(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_BufTextLen").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetBufTextLen(Self *ImGuiInputTextCallbackData, BufTextLen int32) {
	getProc("ImGuiInputTextCallbackData_setBufTextLen").Call(uintptr(unsafe.Pointer(Self)), uintptr(BufTextLen))
}

func (i *Imgui) ImGuiInputTextCallbackDataBufSize(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_BufSize").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetBufSize(Self *ImGuiInputTextCallbackData, BufSize int32) {
	getProc("ImGuiInputTextCallbackData_setBufSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(BufSize))
}

func (i *Imgui) ImGuiInputTextCallbackDataCursorPos(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_CursorPos").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetCursorPos(Self *ImGuiInputTextCallbackData, CursorPos int32) {
	getProc("ImGuiInputTextCallbackData_setCursorPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(CursorPos))
}

func (i *Imgui) ImGuiInputTextCallbackDataSelectionStart(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_SelectionStart").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetSelectionStart(Self *ImGuiInputTextCallbackData, SelectionStart int32) {
	getProc("ImGuiInputTextCallbackData_setSelectionStart").Call(uintptr(unsafe.Pointer(Self)), uintptr(SelectionStart))
}

func (i *Imgui) ImGuiInputTextCallbackDataSelectionEnd(Self *ImGuiInputTextCallbackData) int32 {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_SelectionEnd").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiInputTextCallbackDataSetSelectionEnd(Self *ImGuiInputTextCallbackData, SelectionEnd int32) {
	getProc("ImGuiInputTextCallbackData_setSelectionEnd").Call(uintptr(unsafe.Pointer(Self)), uintptr(SelectionEnd))
}

func (i *Imgui) ImGuiInputTextCallbackDataDeleteChars(Self *ImGuiInputTextCallbackData, Pos int32, Bytes_count int32) {
	getProc("ImGuiInputTextCallbackData_DeleteChars").Call(uintptr(unsafe.Pointer(Self)), uintptr(Pos), uintptr(Bytes_count))
}

func (i *Imgui) ImGuiInputTextCallbackDataInsertChars(Self *ImGuiInputTextCallbackData, Pos int32, Text *int8) {
	getProc("ImGuiInputTextCallbackData_InsertChars").Call(uintptr(unsafe.Pointer(Self)), uintptr(Pos), uintptr(unsafe.Pointer(Text)))
}

func (i *Imgui) ImGuiInputTextCallbackDataSelectAll(Self *ImGuiInputTextCallbackData) {
	getProc("ImGuiInputTextCallbackData_SelectAll").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiInputTextCallbackDataSetSelection(Self *ImGuiInputTextCallbackData, S int32, E int32) {
	getProc("ImGuiInputTextCallbackData_SetSelection").Call(uintptr(unsafe.Pointer(Self)), uintptr(S), uintptr(E))
}

func (i *Imgui) ImGuiInputTextCallbackDataClearSelection(Self *ImGuiInputTextCallbackData) {
	getProc("ImGuiInputTextCallbackData_ClearSelection").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiInputTextCallbackDataHasSelection(Self *ImGuiInputTextCallbackData) bool {
	r1, _, _ := getProc("ImGuiInputTextCallbackData_HasSelection").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiInputTextCallbackDataInsertChars2(Self *ImGuiInputTextCallbackData, Pos int32, Text *int8, Text_end *int8) {
	getProc("ImGuiInputTextCallbackData_InsertChars2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Pos), uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)))
}

func (i *Imgui) ImGuiInputTextCallbackDataDelete(Self *ImGuiInputTextCallbackData) {
	getProc("ImGuiInputTextCallbackData_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiSizeCallbackDataPos(Self *ImGuiSizeCallbackData) *ImVec2 {
	r1, _, _ := getProc("ImGuiSizeCallbackData_Pos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiSizeCallbackDataSetPos(Self *ImGuiSizeCallbackData, Pos *ImVec2) {
	getProc("ImGuiSizeCallbackData_setPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)))
}

func (i *Imgui) ImGuiSizeCallbackDataCurrentSize(Self *ImGuiSizeCallbackData) *ImVec2 {
	r1, _, _ := getProc("ImGuiSizeCallbackData_CurrentSize").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiSizeCallbackDataSetCurrentSize(Self *ImGuiSizeCallbackData, CurrentSize *ImVec2) {
	getProc("ImGuiSizeCallbackData_setCurrentSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(CurrentSize)))
}

func (i *Imgui) ImGuiSizeCallbackDataDesiredSize(Self *ImGuiSizeCallbackData) *ImVec2 {
	r1, _, _ := getProc("ImGuiSizeCallbackData_DesiredSize").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiSizeCallbackDataSetDesiredSize(Self *ImGuiSizeCallbackData, DesiredSize *ImVec2) {
	getProc("ImGuiSizeCallbackData_setDesiredSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DesiredSize)))
}

func (i *Imgui) ImGuiSizeCallbackDataDelete(Self *ImGuiSizeCallbackData) {
	getProc("ImGuiSizeCallbackData_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiPayloadNew() *ImGuiPayload {
	r1, _, _ := getProc("ImGuiPayload_new").Call()
	return (*ImGuiPayload)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiPayloadDataSize(Self *ImGuiPayload) int32 {
	r1, _, _ := getProc("ImGuiPayload_DataSize").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiPayloadSetDataSize(Self *ImGuiPayload, DataSize int32) {
	getProc("ImGuiPayload_setDataSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(DataSize))
}

func (i *Imgui) ImGuiPayloadSourceId(Self *ImGuiPayload) uint32 {
	r1, _, _ := getProc("ImGuiPayload_SourceId").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiPayloadSetSourceId(Self *ImGuiPayload, SourceId uint32) {
	getProc("ImGuiPayload_setSourceId").Call(uintptr(unsafe.Pointer(Self)), uintptr(SourceId))
}

func (i *Imgui) ImGuiPayloadSourceParentId(Self *ImGuiPayload) uint32 {
	r1, _, _ := getProc("ImGuiPayload_SourceParentId").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiPayloadSetSourceParentId(Self *ImGuiPayload, SourceParentId uint32) {
	getProc("ImGuiPayload_setSourceParentId").Call(uintptr(unsafe.Pointer(Self)), uintptr(SourceParentId))
}

func (i *Imgui) ImGuiPayloadDataFrameCount(Self *ImGuiPayload) int32 {
	r1, _, _ := getProc("ImGuiPayload_DataFrameCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiPayloadSetDataFrameCount(Self *ImGuiPayload, DataFrameCount int32) {
	getProc("ImGuiPayload_setDataFrameCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(DataFrameCount))
}

func (i *Imgui) ImGuiPayloadPreview(Self *ImGuiPayload) bool {
	r1, _, _ := getProc("ImGuiPayload_Preview").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiPayloadSetPreview(Self *ImGuiPayload, Preview bool) {
	getProc("ImGuiPayload_setPreview").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Preview {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiPayloadDelivery(Self *ImGuiPayload) bool {
	r1, _, _ := getProc("ImGuiPayload_Delivery").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiPayloadSetDelivery(Self *ImGuiPayload, Delivery bool) {
	getProc("ImGuiPayload_setDelivery").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Delivery {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiPayloadClear(Self *ImGuiPayload) {
	getProc("ImGuiPayload_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiPayloadIsDataType(Self *ImGuiPayload, Type *int8) bool {
	r1, _, _ := getProc("ImGuiPayload_IsDataType").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Type)))
	return r1 != 0
}

func (i *Imgui) ImGuiPayloadIsPreview(Self *ImGuiPayload) bool {
	r1, _, _ := getProc("ImGuiPayload_IsPreview").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiPayloadIsDelivery(Self *ImGuiPayload) bool {
	r1, _, _ := getProc("ImGuiPayload_IsDelivery").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiPayloadDelete(Self *ImGuiPayload) {
	getProc("ImGuiPayload_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiOnceUponAFrameNew() *ImGuiOnceUponAFrame {
	r1, _, _ := getProc("ImGuiOnceUponAFrame_new").Call()
	return (*ImGuiOnceUponAFrame)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiOnceUponAFrameRefFrame(Self *ImGuiOnceUponAFrame) int32 {
	r1, _, _ := getProc("ImGuiOnceUponAFrame_RefFrame").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiOnceUponAFrameSetRefFrame(Self *ImGuiOnceUponAFrame, RefFrame int32) {
	getProc("ImGuiOnceUponAFrame_setRefFrame").Call(uintptr(unsafe.Pointer(Self)), uintptr(RefFrame))
}

func (i *Imgui) ImGuiOnceUponAFrameToBool(Self *ImGuiOnceUponAFrame) bool {
	r1, _, _ := getProc("ImGuiOnceUponAFrame_ToBool").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiOnceUponAFrameDelete(Self *ImGuiOnceUponAFrame) {
	getProc("ImGuiOnceUponAFrame_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTextFilterNew() *ImGuiTextFilter {
	r1, _, _ := getProc("ImGuiTextFilter_new").Call()
	return (*ImGuiTextFilter)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextFilterNew2(Param1 *ImGuiTextFilter) *ImGuiTextFilter {
	r1, _, _ := getProc("ImGuiTextFilter_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiTextFilter)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextFilterNew3(Default_filter *int8) *ImGuiTextFilter {
	r1, _, _ := getProc("ImGuiTextFilter_new3").Call(uintptr(unsafe.Pointer(Default_filter)))
	return (*ImGuiTextFilter)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextFilterDraw(Self *ImGuiTextFilter) bool {
	r1, _, _ := getProc("ImGuiTextFilter_Draw").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiTextFilterPassFilter(Self *ImGuiTextFilter, Text *int8) bool {
	r1, _, _ := getProc("ImGuiTextFilter_PassFilter").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Text)))
	return r1 != 0
}

func (i *Imgui) ImGuiTextFilterBuild(Self *ImGuiTextFilter) {
	getProc("ImGuiTextFilter_Build").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTextFilterClear(Self *ImGuiTextFilter) {
	getProc("ImGuiTextFilter_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTextFilterIsActive(Self *ImGuiTextFilter) bool {
	r1, _, _ := getProc("ImGuiTextFilter_IsActive").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiTextFilterCountGrep(Self *ImGuiTextFilter) int32 {
	r1, _, _ := getProc("ImGuiTextFilter_CountGrep").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiTextFilterSetCountGrep(Self *ImGuiTextFilter, CountGrep int32) {
	getProc("ImGuiTextFilter_setCountGrep").Call(uintptr(unsafe.Pointer(Self)), uintptr(CountGrep))
}

func (i *Imgui) ImGuiTextFilterOperatorAssign(Self *ImGuiTextFilter, Param1 *ImGuiTextFilter) {
	getProc("ImGuiTextFilter_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiTextFilterDrawWithLabel(Self *ImGuiTextFilter, Label *int8) bool {
	r1, _, _ := getProc("ImGuiTextFilter_DrawWithLabel").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Label)))
	return r1 != 0
}

func (i *Imgui) ImGuiTextFilterDraw2(Self *ImGuiTextFilter, Label *int8, Width float32) bool {
	r1, _, _ := getProc("ImGuiTextFilter_Draw2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Label)), uintptr(Width))
	return r1 != 0
}

func (i *Imgui) ImGuiTextFilterPassFilter2(Self *ImGuiTextFilter, Text *int8, Text_end *int8) bool {
	r1, _, _ := getProc("ImGuiTextFilter_PassFilter2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)))
	return r1 != 0
}

func (i *Imgui) ImGuiTextFilterDelete(Self *ImGuiTextFilter) {
	getProc("ImGuiTextFilter_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTextBufferNew() *ImGuiTextBuffer {
	r1, _, _ := getProc("ImGuiTextBuffer_new").Call()
	return (*ImGuiTextBuffer)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextBufferNew2(Param1 *ImGuiTextBuffer) *ImGuiTextBuffer {
	r1, _, _ := getProc("ImGuiTextBuffer_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiTextBuffer)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextBufferBuf(Self *ImGuiTextBuffer) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiTextBuffer_Buf").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiTextBufferSetBuf(Self *ImGuiTextBuffer, Buf unsafe.Pointer) {
	getProc("ImGuiTextBuffer_setBuf").Call(uintptr(unsafe.Pointer(Self)), uintptr(Buf))
}

func (i *Imgui) ImGuiTextBufferOperatorSubscript(Self *ImGuiTextBuffer, I int32) int8 {
	r1, _, _ := getProc("ImGuiTextBuffer_operatorSubscript").Call(uintptr(unsafe.Pointer(Self)), uintptr(I))
	return int8(r1)
}

func (i *Imgui) ImGuiTextBufferBegin(Self *ImGuiTextBuffer) *int8 {
	r1, _, _ := getProc("ImGuiTextBuffer_begin").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextBufferEnd(Self *ImGuiTextBuffer) *int8 {
	r1, _, _ := getProc("ImGuiTextBuffer_end").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextBufferSize(Self *ImGuiTextBuffer) int32 {
	r1, _, _ := getProc("ImGuiTextBuffer_size").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiTextBufferEmpty(Self *ImGuiTextBuffer) bool {
	r1, _, _ := getProc("ImGuiTextBuffer_empty").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiTextBufferClear(Self *ImGuiTextBuffer) {
	getProc("ImGuiTextBuffer_clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiTextBufferResize(Self *ImGuiTextBuffer, Size int32) {
	getProc("ImGuiTextBuffer_resize").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size))
}

func (i *Imgui) ImGuiTextBufferReserve(Self *ImGuiTextBuffer, Capacity int32) {
	getProc("ImGuiTextBuffer_reserve").Call(uintptr(unsafe.Pointer(Self)), uintptr(Capacity))
}

func (i *Imgui) ImGuiTextBufferCStr(Self *ImGuiTextBuffer) *int8 {
	r1, _, _ := getProc("ImGuiTextBuffer_cStr").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiTextBufferAppend(Self *ImGuiTextBuffer, Str *int8) {
	getProc("ImGuiTextBuffer_append").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Str)))
}

func (i *Imgui) ImGuiTextBufferOperatorAssign(Self *ImGuiTextBuffer, Param1 *ImGuiTextBuffer) {
	getProc("ImGuiTextBuffer_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiTextBufferAppend2(Self *ImGuiTextBuffer, Str *int8, Str_end *int8) {
	getProc("ImGuiTextBuffer_append2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Str)), uintptr(unsafe.Pointer(Str_end)))
}

func (i *Imgui) ImGuiTextBufferDelete(Self *ImGuiTextBuffer) {
	getProc("ImGuiTextBuffer_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiStoragePairNew(Key uint32, Val int32) *ImGuiStoragePair {
	r1, _, _ := getProc("ImGuiStoragePair_new").Call(uintptr(Key), uintptr(Val))
	return (*ImGuiStoragePair)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStoragePairNew2(Key uint32, Val float32) *ImGuiStoragePair {
	r1, _, _ := getProc("ImGuiStoragePair_new2").Call(uintptr(Key), uintptr(Val))
	return (*ImGuiStoragePair)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStoragePairNew3(Key uint32, Val unsafe.Pointer) *ImGuiStoragePair {
	r1, _, _ := getProc("ImGuiStoragePair_new3").Call(uintptr(Key), uintptr(Val))
	return (*ImGuiStoragePair)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStoragePairKey(Self *ImGuiStoragePair) uint32 {
	r1, _, _ := getProc("ImGuiStoragePair_key").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiStoragePairSetKey(Self *ImGuiStoragePair, Key uint32) {
	getProc("ImGuiStoragePair_setKey").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
}

func (i *Imgui) ImGuiStoragePairDelete(Self *ImGuiStoragePair) {
	getProc("ImGuiStoragePair_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiStorageNew(Param1 *ImGuiStorage) *ImGuiStorage {
	r1, _, _ := getProc("ImGuiStorage_new").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiStorage)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageData(Self *ImGuiStorage) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiStorage_Data").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiStorageSetData(Self *ImGuiStorage, Data unsafe.Pointer) {
	getProc("ImGuiStorage_setData").Call(uintptr(unsafe.Pointer(Self)), uintptr(Data))
}

func (i *Imgui) ImGuiStorageClear(Self *ImGuiStorage) {
	getProc("ImGuiStorage_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiStorageGetInt(Self *ImGuiStorage, Key uint32) int32 {
	r1, _, _ := getProc("ImGuiStorage_GetInt").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return int32(r1)
}

func (i *Imgui) ImGuiStorageSetInt(Self *ImGuiStorage, Key uint32, Val int32) {
	getProc("ImGuiStorage_SetInt").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Val))
}

func (i *Imgui) ImGuiStorageGetBool(Self *ImGuiStorage, Key uint32) bool {
	r1, _, _ := getProc("ImGuiStorage_GetBool").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return r1 != 0
}

func (i *Imgui) ImGuiStorageSetBool(Self *ImGuiStorage, Key uint32, Val bool) {
	getProc("ImGuiStorage_SetBool").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), func() uintptr {
		if Val {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiStorageGetFloat(Self *ImGuiStorage, Key uint32) float32 {
	r1, _, _ := getProc("ImGuiStorage_GetFloat").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStorageSetFloat(Self *ImGuiStorage, Key uint32, Val float32) {
	getProc("ImGuiStorage_SetFloat").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Val))
}

func (i *Imgui) ImGuiStorageGetVoidPtr(Self *ImGuiStorage, Key uint32) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiStorage_GetVoidPtr").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiStorageGetIntRef(Self *ImGuiStorage, Key uint32) *int32 {
	r1, _, _ := getProc("ImGuiStorage_GetIntRef").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return (*int32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageGetBoolRef(Self *ImGuiStorage, Key uint32) *bool {
	r1, _, _ := getProc("ImGuiStorage_GetBoolRef").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageGetFloatRef(Self *ImGuiStorage, Key uint32) *float32 {
	r1, _, _ := getProc("ImGuiStorage_GetFloatRef").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key))
	return (*float32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageBuildSortByKey(Self *ImGuiStorage) {
	getProc("ImGuiStorage_BuildSortByKey").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiStorageSetAllInt(Self *ImGuiStorage, Val int32) {
	getProc("ImGuiStorage_SetAllInt").Call(uintptr(unsafe.Pointer(Self)), uintptr(Val))
}

func (i *Imgui) ImGuiStorageOperatorAssign(Self *ImGuiStorage, Param1 *ImGuiStorage) {
	getProc("ImGuiStorage_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiStorageGetInt2(Self *ImGuiStorage, Key uint32, Default_val int32) int32 {
	r1, _, _ := getProc("ImGuiStorage_GetInt2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Default_val))
	return int32(r1)
}

func (i *Imgui) ImGuiStorageGetBool2(Self *ImGuiStorage, Key uint32, Default_val bool) bool {
	r1, _, _ := getProc("ImGuiStorage_GetBool2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), func() uintptr {
		if Default_val {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) ImGuiStorageGetFloat2(Self *ImGuiStorage, Key uint32, Default_val float32) float32 {
	r1, _, _ := getProc("ImGuiStorage_GetFloat2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Default_val))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiStorageGetIntRef2(Self *ImGuiStorage, Key uint32, Default_val int32) *int32 {
	r1, _, _ := getProc("ImGuiStorage_GetIntRef2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Default_val))
	return (*int32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageGetBoolRef2(Self *ImGuiStorage, Key uint32, Default_val bool) *bool {
	r1, _, _ := getProc("ImGuiStorage_GetBoolRef2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), func() uintptr {
		if Default_val {
			return 1
		}
		return 0
	}())
	return (*bool)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageGetFloatRef2(Self *ImGuiStorage, Key uint32, Default_val float32) *float32 {
	r1, _, _ := getProc("ImGuiStorage_GetFloatRef2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Key), uintptr(Default_val))
	return (*float32)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiStorageDelete(Self *ImGuiStorage) {
	getProc("ImGuiStorage_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiListClipperNew() *ImGuiListClipper {
	r1, _, _ := getProc("ImGuiListClipper_new").Call()
	return (*ImGuiListClipper)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiListClipperDisplayStart(Self *ImGuiListClipper) int32 {
	r1, _, _ := getProc("ImGuiListClipper_DisplayStart").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiListClipperSetDisplayStart(Self *ImGuiListClipper, DisplayStart int32) {
	getProc("ImGuiListClipper_setDisplayStart").Call(uintptr(unsafe.Pointer(Self)), uintptr(DisplayStart))
}

func (i *Imgui) ImGuiListClipperDisplayEnd(Self *ImGuiListClipper) int32 {
	r1, _, _ := getProc("ImGuiListClipper_DisplayEnd").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiListClipperSetDisplayEnd(Self *ImGuiListClipper, DisplayEnd int32) {
	getProc("ImGuiListClipper_setDisplayEnd").Call(uintptr(unsafe.Pointer(Self)), uintptr(DisplayEnd))
}

func (i *Imgui) ImGuiListClipperUserIndex(Self *ImGuiListClipper) int32 {
	r1, _, _ := getProc("ImGuiListClipper_UserIndex").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiListClipperSetUserIndex(Self *ImGuiListClipper, UserIndex int32) {
	getProc("ImGuiListClipper_setUserIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(UserIndex))
}

func (i *Imgui) ImGuiListClipperItemsCount(Self *ImGuiListClipper) int32 {
	r1, _, _ := getProc("ImGuiListClipper_ItemsCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiListClipperSetItemsCount(Self *ImGuiListClipper, ItemsCount int32) {
	getProc("ImGuiListClipper_setItemsCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(ItemsCount))
}

func (i *Imgui) ImGuiListClipperItemsHeight(Self *ImGuiListClipper) float32 {
	r1, _, _ := getProc("ImGuiListClipper_ItemsHeight").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiListClipperSetItemsHeight(Self *ImGuiListClipper, ItemsHeight float32) {
	getProc("ImGuiListClipper_setItemsHeight").Call(uintptr(unsafe.Pointer(Self)), uintptr(ItemsHeight))
}

func (i *Imgui) ImGuiListClipperFlags(Self *ImGuiListClipper) int32 {
	r1, _, _ := getProc("ImGuiListClipper_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiListClipperSetFlags(Self *ImGuiListClipper, Flags int32) {
	getProc("ImGuiListClipper_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImGuiListClipperStartPosY(Self *ImGuiListClipper) float64 {
	r1, _, _ := getProc("ImGuiListClipper_StartPosY").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiListClipperSetStartPosY(Self *ImGuiListClipper, StartPosY float64) {
	getProc("ImGuiListClipper_setStartPosY").Call(uintptr(unsafe.Pointer(Self)), uintptr(StartPosY))
}

func (i *Imgui) ImGuiListClipperStartSeekOffsetY(Self *ImGuiListClipper) float64 {
	r1, _, _ := getProc("ImGuiListClipper_StartSeekOffsetY").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiListClipperSetStartSeekOffsetY(Self *ImGuiListClipper, StartSeekOffsetY float64) {
	getProc("ImGuiListClipper_setStartSeekOffsetY").Call(uintptr(unsafe.Pointer(Self)), uintptr(StartSeekOffsetY))
}

func (i *Imgui) ImGuiListClipperCtx(Self *ImGuiListClipper) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiListClipper_Ctx").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiListClipperSetCtx(Self *ImGuiListClipper, Ctx unsafe.Pointer) {
	getProc("ImGuiListClipper_setCtx").Call(uintptr(unsafe.Pointer(Self)), uintptr(Ctx))
}

func (i *Imgui) ImGuiListClipperBegin(Self *ImGuiListClipper, Items_count int32) {
	getProc("ImGuiListClipper_Begin").Call(uintptr(unsafe.Pointer(Self)), uintptr(Items_count))
}

func (i *Imgui) ImGuiListClipperEnd(Self *ImGuiListClipper) {
	getProc("ImGuiListClipper_End").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiListClipperStep(Self *ImGuiListClipper) bool {
	r1, _, _ := getProc("ImGuiListClipper_Step").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiListClipperIncludeItemByIndex(Self *ImGuiListClipper, Item_index int32) {
	getProc("ImGuiListClipper_IncludeItemByIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(Item_index))
}

func (i *Imgui) ImGuiListClipperIncludeItemsByIndex(Self *ImGuiListClipper, Item_begin int32, Item_end int32) {
	getProc("ImGuiListClipper_IncludeItemsByIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(Item_begin), uintptr(Item_end))
}

func (i *Imgui) ImGuiListClipperSeekCursorForItem(Self *ImGuiListClipper, Item_index int32) {
	getProc("ImGuiListClipper_SeekCursorForItem").Call(uintptr(unsafe.Pointer(Self)), uintptr(Item_index))
}

func (i *Imgui) ImGuiListClipperBegin2(Self *ImGuiListClipper, Items_count int32, Items_height float32) {
	getProc("ImGuiListClipper_Begin2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Items_count), uintptr(Items_height))
}

func (i *Imgui) ImGuiListClipperDelete(Self *ImGuiListClipper) {
	getProc("ImGuiListClipper_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImColorNew() *ImColor {
	r1, _, _ := getProc("ImColor_new").Call()
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew2(R float32, G float32, B float32) *ImColor {
	r1, _, _ := getProc("ImColor_new2").Call(uintptr(R), uintptr(G), uintptr(B))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew3(Col *ImVec4) *ImColor {
	r1, _, _ := getProc("ImColor_new3").Call(uintptr(unsafe.Pointer(Col)))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew4(R int32, G int32, B int32) *ImColor {
	r1, _, _ := getProc("ImColor_new4").Call(uintptr(R), uintptr(G), uintptr(B))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew5(Rgba uint32) *ImColor {
	r1, _, _ := getProc("ImColor_new5").Call(uintptr(Rgba))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew6(Param1 *ImColor) *ImColor {
	r1, _, _ := getProc("ImColor_new6").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew7(R float32, G float32, B float32, A float32) *ImColor {
	r1, _, _ := getProc("ImColor_new7").Call(uintptr(R), uintptr(G), uintptr(B), uintptr(A))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorNew8(R int32, G int32, B int32, A int32) *ImColor {
	r1, _, _ := getProc("ImColor_new8").Call(uintptr(R), uintptr(G), uintptr(B), uintptr(A))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorValue(Self *ImColor) *ImVec4 {
	r1, _, _ := getProc("ImColor_Value").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorSetValue(Self *ImColor, Value *ImVec4) {
	getProc("ImColor_setValue").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Value)))
}

func (i *Imgui) ImColorToUnsignedInt(Self *ImColor) uint32 {
	r1, _, _ := getProc("ImColor_ToUnsignedInt").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImColorToImVec4(Self *ImColor) *ImVec4 {
	r1, _, _ := getProc("ImColor_ToImVec4").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorSetHSV(Self *ImColor, H float32, S float32, V float32) {
	getProc("ImColor_SetHSV").Call(uintptr(unsafe.Pointer(Self)), uintptr(H), uintptr(S), uintptr(V))
}

func (i *Imgui) ImColorHSV(H float32, S float32, V float32) *ImColor {
	r1, _, _ := getProc("ImColor_HSV").Call(uintptr(H), uintptr(S), uintptr(V))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorSetHSV2(Self *ImColor, H float32, S float32, V float32, A float32) {
	getProc("ImColor_SetHSV2").Call(uintptr(unsafe.Pointer(Self)), uintptr(H), uintptr(S), uintptr(V), uintptr(A))
}

func (i *Imgui) ImColorHSV2(H float32, S float32, V float32, A float32) *ImColor {
	r1, _, _ := getProc("ImColor_HSV2").Call(uintptr(H), uintptr(S), uintptr(V), uintptr(A))
	return (*ImColor)(unsafe.Pointer(r1))
}

func (i *Imgui) ImColorDelete(Self *ImColor) {
	getProc("ImColor_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiMultiSelectIONew(Param1 *ImGuiMultiSelectIO) *ImGuiMultiSelectIO {
	r1, _, _ := getProc("ImGuiMultiSelectIO_new").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiMultiSelectIO)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiMultiSelectIORequests(Self *ImGuiMultiSelectIO) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiMultiSelectIO_Requests").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiMultiSelectIOSetRequests(Self *ImGuiMultiSelectIO, Requests unsafe.Pointer) {
	getProc("ImGuiMultiSelectIO_setRequests").Call(uintptr(unsafe.Pointer(Self)), uintptr(Requests))
}

func (i *Imgui) ImGuiMultiSelectIORangeSrcItem(Self *ImGuiMultiSelectIO) int64 {
	r1, _, _ := getProc("ImGuiMultiSelectIO_RangeSrcItem").Call(uintptr(unsafe.Pointer(Self)))
	return *(*int64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiMultiSelectIOSetRangeSrcItem(Self *ImGuiMultiSelectIO, RangeSrcItem int64) {
	getProc("ImGuiMultiSelectIO_setRangeSrcItem").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&RangeSrcItem)))
}

func (i *Imgui) ImGuiMultiSelectIONavIdItem(Self *ImGuiMultiSelectIO) int64 {
	r1, _, _ := getProc("ImGuiMultiSelectIO_NavIdItem").Call(uintptr(unsafe.Pointer(Self)))
	return *(*int64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiMultiSelectIOSetNavIdItem(Self *ImGuiMultiSelectIO, NavIdItem int64) {
	getProc("ImGuiMultiSelectIO_setNavIdItem").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&NavIdItem)))
}

func (i *Imgui) ImGuiMultiSelectIONavIdSelected(Self *ImGuiMultiSelectIO) bool {
	r1, _, _ := getProc("ImGuiMultiSelectIO_NavIdSelected").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiMultiSelectIOSetNavIdSelected(Self *ImGuiMultiSelectIO, NavIdSelected bool) {
	getProc("ImGuiMultiSelectIO_setNavIdSelected").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if NavIdSelected {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiMultiSelectIORangeSrcReset(Self *ImGuiMultiSelectIO) bool {
	r1, _, _ := getProc("ImGuiMultiSelectIO_RangeSrcReset").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiMultiSelectIOSetRangeSrcReset(Self *ImGuiMultiSelectIO, RangeSrcReset bool) {
	getProc("ImGuiMultiSelectIO_setRangeSrcReset").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if RangeSrcReset {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiMultiSelectIOItemsCount(Self *ImGuiMultiSelectIO) int32 {
	r1, _, _ := getProc("ImGuiMultiSelectIO_ItemsCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiMultiSelectIOSetItemsCount(Self *ImGuiMultiSelectIO, ItemsCount int32) {
	getProc("ImGuiMultiSelectIO_setItemsCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(ItemsCount))
}

func (i *Imgui) ImGuiMultiSelectIOOperatorAssign(Self *ImGuiMultiSelectIO, Param1 *ImGuiMultiSelectIO) {
	getProc("ImGuiMultiSelectIO_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiMultiSelectIODelete(Self *ImGuiMultiSelectIO) {
	getProc("ImGuiMultiSelectIO_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiSelectionRequestType(Self *ImGuiSelectionRequest) int32 {
	r1, _, _ := getProc("ImGuiSelectionRequest_Type").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiSelectionRequestSetType(Self *ImGuiSelectionRequest, Type int32) {
	getProc("ImGuiSelectionRequest_setType").Call(uintptr(unsafe.Pointer(Self)), uintptr(Type))
}

func (i *Imgui) ImGuiSelectionRequestSelected(Self *ImGuiSelectionRequest) bool {
	r1, _, _ := getProc("ImGuiSelectionRequest_Selected").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiSelectionRequestSetSelected(Self *ImGuiSelectionRequest, Selected bool) {
	getProc("ImGuiSelectionRequest_setSelected").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Selected {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiSelectionRequestRangeDirection(Self *ImGuiSelectionRequest) int8 {
	r1, _, _ := getProc("ImGuiSelectionRequest_RangeDirection").Call(uintptr(unsafe.Pointer(Self)))
	return int8(r1)
}

func (i *Imgui) ImGuiSelectionRequestSetRangeDirection(Self *ImGuiSelectionRequest, RangeDirection int8) {
	getProc("ImGuiSelectionRequest_setRangeDirection").Call(uintptr(unsafe.Pointer(Self)), uintptr(RangeDirection))
}

func (i *Imgui) ImGuiSelectionRequestRangeFirstItem(Self *ImGuiSelectionRequest) int64 {
	r1, _, _ := getProc("ImGuiSelectionRequest_RangeFirstItem").Call(uintptr(unsafe.Pointer(Self)))
	return *(*int64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiSelectionRequestSetRangeFirstItem(Self *ImGuiSelectionRequest, RangeFirstItem int64) {
	getProc("ImGuiSelectionRequest_setRangeFirstItem").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&RangeFirstItem)))
}

func (i *Imgui) ImGuiSelectionRequestRangeLastItem(Self *ImGuiSelectionRequest) int64 {
	r1, _, _ := getProc("ImGuiSelectionRequest_RangeLastItem").Call(uintptr(unsafe.Pointer(Self)))
	return *(*int64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiSelectionRequestSetRangeLastItem(Self *ImGuiSelectionRequest, RangeLastItem int64) {
	getProc("ImGuiSelectionRequest_setRangeLastItem").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&RangeLastItem)))
}

func (i *Imgui) ImGuiSelectionRequestDelete(Self *ImGuiSelectionRequest) {
	getProc("ImGuiSelectionRequest_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiSelectionBasicStorageNew() *ImGuiSelectionBasicStorage {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage_new").Call()
	return (*ImGuiSelectionBasicStorage)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiSelectionBasicStorageSize(Self *ImGuiSelectionBasicStorage) int32 {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage_Size").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiSelectionBasicStorageSetSize(Self *ImGuiSelectionBasicStorage, Size int32) {
	getProc("ImGuiSelectionBasicStorage_setSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size))
}

func (i *Imgui) ImGuiSelectionBasicStoragePreserveOrder(Self *ImGuiSelectionBasicStorage) bool {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage_PreserveOrder").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiSelectionBasicStorageSetPreserveOrder(Self *ImGuiSelectionBasicStorage, PreserveOrder bool) {
	getProc("ImGuiSelectionBasicStorage_setPreserveOrder").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if PreserveOrder {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiSelectionBasicStorageSelectionOrder(Self *ImGuiSelectionBasicStorage) int32 {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage__SelectionOrder").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiSelectionBasicStorageSetSelectionOrder(Self *ImGuiSelectionBasicStorage, SelectionOrder int32) {
	getProc("ImGuiSelectionBasicStorage_set_SelectionOrder").Call(uintptr(unsafe.Pointer(Self)), uintptr(SelectionOrder))
}

func (i *Imgui) ImGuiSelectionBasicStorageStorage(Self *ImGuiSelectionBasicStorage) *ImGuiStorage {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage__Storage").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImGuiStorage)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiSelectionBasicStorageSetStorage(Self *ImGuiSelectionBasicStorage, Storage *ImGuiStorage) {
	getProc("ImGuiSelectionBasicStorage_set_Storage").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Storage)))
}

func (i *Imgui) ImGuiSelectionBasicStorageApplyRequests(Self *ImGuiSelectionBasicStorage, Ms_io *ImGuiMultiSelectIO) {
	getProc("ImGuiSelectionBasicStorage_ApplyRequests").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Ms_io)))
}

func (i *Imgui) ImGuiSelectionBasicStorageContains(Self *ImGuiSelectionBasicStorage, Id uint32) bool {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage_Contains").Call(uintptr(unsafe.Pointer(Self)), uintptr(Id))
	return r1 != 0
}

func (i *Imgui) ImGuiSelectionBasicStorageClear(Self *ImGuiSelectionBasicStorage) {
	getProc("ImGuiSelectionBasicStorage_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiSelectionBasicStorageSwap(Self *ImGuiSelectionBasicStorage, R *ImGuiSelectionBasicStorage) {
	getProc("ImGuiSelectionBasicStorage_Swap").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(R)))
}

func (i *Imgui) ImGuiSelectionBasicStorageSetItemSelected(Self *ImGuiSelectionBasicStorage, Id uint32, Selected bool) {
	getProc("ImGuiSelectionBasicStorage_SetItemSelected").Call(uintptr(unsafe.Pointer(Self)), uintptr(Id), func() uintptr {
		if Selected {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiSelectionBasicStorageGetStorageIdFromIndex(Self *ImGuiSelectionBasicStorage, Idx int32) uint32 {
	r1, _, _ := getProc("ImGuiSelectionBasicStorage_GetStorageIdFromIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(Idx))
	return uint32(r1)
}

func (i *Imgui) ImGuiSelectionBasicStorageDelete(Self *ImGuiSelectionBasicStorage) {
	getProc("ImGuiSelectionBasicStorage_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiSelectionExternalStorageNew() *ImGuiSelectionExternalStorage {
	r1, _, _ := getProc("ImGuiSelectionExternalStorage_new").Call()
	return (*ImGuiSelectionExternalStorage)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiSelectionExternalStorageApplyRequests(Self *ImGuiSelectionExternalStorage, Ms_io *ImGuiMultiSelectIO) {
	getProc("ImGuiSelectionExternalStorage_ApplyRequests").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Ms_io)))
}

func (i *Imgui) ImGuiSelectionExternalStorageDelete(Self *ImGuiSelectionExternalStorage) {
	getProc("ImGuiSelectionExternalStorage_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawCmdNew() *ImDrawCmd {
	r1, _, _ := getProc("ImDrawCmd_new").Call()
	return (*ImDrawCmd)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawCmdClipRect(Self *ImDrawCmd) *ImVec4 {
	r1, _, _ := getProc("ImDrawCmd_ClipRect").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawCmdSetClipRect(Self *ImDrawCmd, ClipRect *ImVec4) {
	getProc("ImDrawCmd_setClipRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(ClipRect)))
}

func (i *Imgui) ImDrawCmdTexRef(Self *ImDrawCmd) *ImTextureRef {
	r1, _, _ := getProc("ImDrawCmd_TexRef").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawCmdSetTexRef(Self *ImDrawCmd, TexRef *ImTextureRef) {
	getProc("ImDrawCmd_setTexRef").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexRef)))
}

func (i *Imgui) ImDrawCmdVtxOffset(Self *ImDrawCmd) uint32 {
	r1, _, _ := getProc("ImDrawCmd_VtxOffset").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImDrawCmdSetVtxOffset(Self *ImDrawCmd, VtxOffset uint32) {
	getProc("ImDrawCmd_setVtxOffset").Call(uintptr(unsafe.Pointer(Self)), uintptr(VtxOffset))
}

func (i *Imgui) ImDrawCmdIdxOffset(Self *ImDrawCmd) uint32 {
	r1, _, _ := getProc("ImDrawCmd_IdxOffset").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImDrawCmdSetIdxOffset(Self *ImDrawCmd, IdxOffset uint32) {
	getProc("ImDrawCmd_setIdxOffset").Call(uintptr(unsafe.Pointer(Self)), uintptr(IdxOffset))
}

func (i *Imgui) ImDrawCmdElemCount(Self *ImDrawCmd) uint32 {
	r1, _, _ := getProc("ImDrawCmd_ElemCount").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImDrawCmdSetElemCount(Self *ImDrawCmd, ElemCount uint32) {
	getProc("ImDrawCmd_setElemCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(ElemCount))
}

func (i *Imgui) ImDrawCmdUserCallbackDataSize(Self *ImDrawCmd) int32 {
	r1, _, _ := getProc("ImDrawCmd_UserCallbackDataSize").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawCmdSetUserCallbackDataSize(Self *ImDrawCmd, UserCallbackDataSize int32) {
	getProc("ImDrawCmd_setUserCallbackDataSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(UserCallbackDataSize))
}

func (i *Imgui) ImDrawCmdUserCallbackDataOffset(Self *ImDrawCmd) int32 {
	r1, _, _ := getProc("ImDrawCmd_UserCallbackDataOffset").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawCmdSetUserCallbackDataOffset(Self *ImDrawCmd, UserCallbackDataOffset int32) {
	getProc("ImDrawCmd_setUserCallbackDataOffset").Call(uintptr(unsafe.Pointer(Self)), uintptr(UserCallbackDataOffset))
}

func (i *Imgui) ImDrawCmdGetTexID(Self *ImDrawCmd) uint64 {
	r1, _, _ := getProc("ImDrawCmd_GetTexID").Call(uintptr(unsafe.Pointer(Self)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImDrawCmdDelete(Self *ImDrawCmd) {
	getProc("ImDrawCmd_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawVertPos(Self *ImDrawVert) *ImVec2 {
	r1, _, _ := getProc("ImDrawVert_pos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawVertSetPos(Self *ImDrawVert, Pos *ImVec2) {
	getProc("ImDrawVert_setPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)))
}

func (i *Imgui) ImDrawVertUv(Self *ImDrawVert) *ImVec2 {
	r1, _, _ := getProc("ImDrawVert_uv").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawVertSetUv(Self *ImDrawVert, Uv *ImVec2) {
	getProc("ImDrawVert_setUv").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Uv)))
}

func (i *Imgui) ImDrawVertCol(Self *ImDrawVert) uint32 {
	r1, _, _ := getProc("ImDrawVert_col").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImDrawVertSetCol(Self *ImDrawVert, Col uint32) {
	getProc("ImDrawVert_setCol").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col))
}

func (i *Imgui) ImDrawVertDelete(Self *ImDrawVert) {
	getProc("ImDrawVert_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawCmdHeaderNew(Param1 *ImDrawCmdHeader) *ImDrawCmdHeader {
	r1, _, _ := getProc("ImDrawCmdHeader_new").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImDrawCmdHeader)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawCmdHeaderClipRect(Self *ImDrawCmdHeader) *ImVec4 {
	r1, _, _ := getProc("ImDrawCmdHeader_ClipRect").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawCmdHeaderSetClipRect(Self *ImDrawCmdHeader, ClipRect *ImVec4) {
	getProc("ImDrawCmdHeader_setClipRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(ClipRect)))
}

func (i *Imgui) ImDrawCmdHeaderTexRef(Self *ImDrawCmdHeader) *ImTextureRef {
	r1, _, _ := getProc("ImDrawCmdHeader_TexRef").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawCmdHeaderSetTexRef(Self *ImDrawCmdHeader, TexRef *ImTextureRef) {
	getProc("ImDrawCmdHeader_setTexRef").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexRef)))
}

func (i *Imgui) ImDrawCmdHeaderVtxOffset(Self *ImDrawCmdHeader) uint32 {
	r1, _, _ := getProc("ImDrawCmdHeader_VtxOffset").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImDrawCmdHeaderSetVtxOffset(Self *ImDrawCmdHeader, VtxOffset uint32) {
	getProc("ImDrawCmdHeader_setVtxOffset").Call(uintptr(unsafe.Pointer(Self)), uintptr(VtxOffset))
}

func (i *Imgui) ImDrawCmdHeaderOperatorAssign(Self *ImDrawCmdHeader, Param1 *ImDrawCmdHeader) {
	getProc("ImDrawCmdHeader_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImDrawCmdHeaderDelete(Self *ImDrawCmdHeader) {
	getProc("ImDrawCmdHeader_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawChannelNew(Param1 *ImDrawChannel) *ImDrawChannel {
	r1, _, _ := getProc("ImDrawChannel_new").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImDrawChannel)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawChannelCmdBuffer(Self *ImDrawChannel) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawChannel__CmdBuffer").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawChannelSetCmdBuffer(Self *ImDrawChannel, CmdBuffer unsafe.Pointer) {
	getProc("ImDrawChannel_set_CmdBuffer").Call(uintptr(unsafe.Pointer(Self)), uintptr(CmdBuffer))
}

func (i *Imgui) ImDrawChannelIdxBuffer(Self *ImDrawChannel) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawChannel__IdxBuffer").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawChannelSetIdxBuffer(Self *ImDrawChannel, IdxBuffer unsafe.Pointer) {
	getProc("ImDrawChannel_set_IdxBuffer").Call(uintptr(unsafe.Pointer(Self)), uintptr(IdxBuffer))
}

func (i *Imgui) ImDrawChannelOperatorAssign(Self *ImDrawChannel, Param1 *ImDrawChannel) {
	getProc("ImDrawChannel_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImDrawChannelDelete(Self *ImDrawChannel) {
	getProc("ImDrawChannel_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListSplitterNew() *ImDrawListSplitter {
	r1, _, _ := getProc("ImDrawListSplitter_new").Call()
	return (*ImDrawListSplitter)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSplitterNew2(Param1 *ImDrawListSplitter) *ImDrawListSplitter {
	r1, _, _ := getProc("ImDrawListSplitter_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImDrawListSplitter)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSplitterCurrent(Self *ImDrawListSplitter) int32 {
	r1, _, _ := getProc("ImDrawListSplitter__Current").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawListSplitterSetCurrent(Self *ImDrawListSplitter, Current int32) {
	getProc("ImDrawListSplitter_set_Current").Call(uintptr(unsafe.Pointer(Self)), uintptr(Current))
}

func (i *Imgui) ImDrawListSplitterCount(Self *ImDrawListSplitter) int32 {
	r1, _, _ := getProc("ImDrawListSplitter__Count").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawListSplitterSetCount(Self *ImDrawListSplitter, Count int32) {
	getProc("ImDrawListSplitter_set_Count").Call(uintptr(unsafe.Pointer(Self)), uintptr(Count))
}

func (i *Imgui) ImDrawListSplitterChannels(Self *ImDrawListSplitter) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawListSplitter__Channels").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSplitterSetChannels(Self *ImDrawListSplitter, Channels unsafe.Pointer) {
	getProc("ImDrawListSplitter_set_Channels").Call(uintptr(unsafe.Pointer(Self)), uintptr(Channels))
}

func (i *Imgui) ImDrawListSplitterClear(Self *ImDrawListSplitter) {
	getProc("ImDrawListSplitter_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListSplitterClearFreeMemory(Self *ImDrawListSplitter) {
	getProc("ImDrawListSplitter_ClearFreeMemory").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListSplitterSplit(Self *ImDrawListSplitter, Draw_list *ImDrawList, Count int32) {
	getProc("ImDrawListSplitter_Split").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Count))
}

func (i *Imgui) ImDrawListSplitterMerge(Self *ImDrawListSplitter, Draw_list *ImDrawList) {
	getProc("ImDrawListSplitter_Merge").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)))
}

func (i *Imgui) ImDrawListSplitterSetCurrentChannel(Self *ImDrawListSplitter, Draw_list *ImDrawList, Channel_idx int32) {
	getProc("ImDrawListSplitter_SetCurrentChannel").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Channel_idx))
}

func (i *Imgui) ImDrawListSplitterOperatorAssign(Self *ImDrawListSplitter, Param1 *ImDrawListSplitter) {
	getProc("ImDrawListSplitter_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImDrawListSplitterDelete(Self *ImDrawListSplitter) {
	getProc("ImDrawListSplitter_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListNew(Shared_data unsafe.Pointer) *ImDrawList {
	r1, _, _ := getProc("ImDrawList_new").Call(uintptr(Shared_data))
	return (*ImDrawList)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListNew2(Param1 *ImDrawList) *ImDrawList {
	r1, _, _ := getProc("ImDrawList_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImDrawList)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListCmdBuffer(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList_CmdBuffer").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetCmdBuffer(Self *ImDrawList, CmdBuffer unsafe.Pointer) {
	getProc("ImDrawList_setCmdBuffer").Call(uintptr(unsafe.Pointer(Self)), uintptr(CmdBuffer))
}

func (i *Imgui) ImDrawListIdxBuffer(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList_IdxBuffer").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetIdxBuffer(Self *ImDrawList, IdxBuffer unsafe.Pointer) {
	getProc("ImDrawList_setIdxBuffer").Call(uintptr(unsafe.Pointer(Self)), uintptr(IdxBuffer))
}

func (i *Imgui) ImDrawListVtxBuffer(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList_VtxBuffer").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetVtxBuffer(Self *ImDrawList, VtxBuffer unsafe.Pointer) {
	getProc("ImDrawList_setVtxBuffer").Call(uintptr(unsafe.Pointer(Self)), uintptr(VtxBuffer))
}

func (i *Imgui) ImDrawListFlags(Self *ImDrawList) int32 {
	r1, _, _ := getProc("ImDrawList_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawListSetFlags(Self *ImDrawList, Flags int32) {
	getProc("ImDrawList_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImDrawListVtxCurrentIdx(Self *ImDrawList) uint32 {
	r1, _, _ := getProc("ImDrawList__VtxCurrentIdx").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImDrawListSetVtxCurrentIdx(Self *ImDrawList, VtxCurrentIdx uint32) {
	getProc("ImDrawList_set_VtxCurrentIdx").Call(uintptr(unsafe.Pointer(Self)), uintptr(VtxCurrentIdx))
}

func (i *Imgui) ImDrawListData(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList__Data").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetData(Self *ImDrawList, Data unsafe.Pointer) {
	getProc("ImDrawList_set_Data").Call(uintptr(unsafe.Pointer(Self)), uintptr(Data))
}

func (i *Imgui) ImDrawListVtxWritePtr(Self *ImDrawList) *ImDrawVert {
	r1, _, _ := getProc("ImDrawList__VtxWritePtr").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImDrawVert)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSetVtxWritePtr(Self *ImDrawList, VtxWritePtr *ImDrawVert) {
	getProc("ImDrawList_set_VtxWritePtr").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(VtxWritePtr)))
}

func (i *Imgui) ImDrawListIdxWritePtr(Self *ImDrawList) *uint16 {
	r1, _, _ := getProc("ImDrawList__IdxWritePtr").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSetIdxWritePtr(Self *ImDrawList, IdxWritePtr *uint16) {
	getProc("ImDrawList_set_IdxWritePtr").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(IdxWritePtr)))
}

func (i *Imgui) ImDrawListPath(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList__Path").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetPath(Self *ImDrawList, Path unsafe.Pointer) {
	getProc("ImDrawList_set_Path").Call(uintptr(unsafe.Pointer(Self)), uintptr(Path))
}

func (i *Imgui) ImDrawListCmdHeader(Self *ImDrawList) *ImDrawCmdHeader {
	r1, _, _ := getProc("ImDrawList__CmdHeader").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImDrawCmdHeader)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSetCmdHeader(Self *ImDrawList, CmdHeader *ImDrawCmdHeader) {
	getProc("ImDrawList_set_CmdHeader").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(CmdHeader)))
}

func (i *Imgui) ImDrawListSplitter(Self *ImDrawList) *ImDrawListSplitter {
	r1, _, _ := getProc("ImDrawList__Splitter").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImDrawListSplitter)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSetSplitter(Self *ImDrawList, Splitter *ImDrawListSplitter) {
	getProc("ImDrawList_set_Splitter").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Splitter)))
}

func (i *Imgui) ImDrawListClipRectStack(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList__ClipRectStack").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetClipRectStack(Self *ImDrawList, ClipRectStack unsafe.Pointer) {
	getProc("ImDrawList_set_ClipRectStack").Call(uintptr(unsafe.Pointer(Self)), uintptr(ClipRectStack))
}

func (i *Imgui) ImDrawListTextureStack(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList__TextureStack").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetTextureStack(Self *ImDrawList, TextureStack unsafe.Pointer) {
	getProc("ImDrawList_set_TextureStack").Call(uintptr(unsafe.Pointer(Self)), uintptr(TextureStack))
}

func (i *Imgui) ImDrawListCallbacksDataBuf(Self *ImDrawList) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawList__CallbacksDataBuf").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawListSetCallbacksDataBuf(Self *ImDrawList, CallbacksDataBuf unsafe.Pointer) {
	getProc("ImDrawList_set_CallbacksDataBuf").Call(uintptr(unsafe.Pointer(Self)), uintptr(CallbacksDataBuf))
}

func (i *Imgui) ImDrawListFringeScale(Self *ImDrawList) float32 {
	r1, _, _ := getProc("ImDrawList__FringeScale").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImDrawListSetFringeScale(Self *ImDrawList, FringeScale float32) {
	getProc("ImDrawList_set_FringeScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(FringeScale))
}

func (i *Imgui) ImDrawListOwnerName(Self *ImDrawList) *int8 {
	r1, _, _ := getProc("ImDrawList__OwnerName").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListSetOwnerName(Self *ImDrawList, OwnerName *int8) {
	getProc("ImDrawList_set_OwnerName").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(OwnerName)))
}

func (i *Imgui) ImDrawListPushClipRect(Self *ImDrawList, Clip_rect_min *ImVec2, Clip_rect_max *ImVec2) {
	getProc("ImDrawList_PushClipRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Clip_rect_min)), uintptr(unsafe.Pointer(Clip_rect_max)))
}

func (i *Imgui) ImDrawListPushClipRectFullScreen(Self *ImDrawList) {
	getProc("ImDrawList_PushClipRectFullScreen").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListPopClipRect(Self *ImDrawList) {
	getProc("ImDrawList_PopClipRect").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListPushTexture(Self *ImDrawList, Tex_ref *ImTextureRef) {
	getProc("ImDrawList_PushTexture").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)))
}

func (i *Imgui) ImDrawListPopTexture(Self *ImDrawList) {
	getProc("ImDrawList_PopTexture").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListGetClipRectMin(Self *ImDrawList) *ImVec2 {
	r1, _, _ := getProc("ImDrawList_GetClipRectMin").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListGetClipRectMax(Self *ImDrawList) *ImVec2 {
	r1, _, _ := getProc("ImDrawList_GetClipRectMax").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListAddLine(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, Col uint32) {
	getProc("ImDrawList_AddLine").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddLineH(Self *ImDrawList, Min_x float32, Max_x float32, Y float32, Col uint32) {
	getProc("ImDrawList_AddLineH").Call(uintptr(unsafe.Pointer(Self)), uintptr(Min_x), uintptr(Max_x), uintptr(Y), uintptr(Col))
}

func (i *Imgui) ImDrawListAddLineV(Self *ImDrawList, X float32, Min_y float32, Max_y float32, Col uint32) {
	getProc("ImDrawList_AddLineV").Call(uintptr(unsafe.Pointer(Self)), uintptr(X), uintptr(Min_y), uintptr(Max_y), uintptr(Col))
}

func (i *Imgui) ImDrawListAddRect(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32) {
	getProc("ImDrawList_AddRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddRectFilled(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32) {
	getProc("ImDrawList_AddRectFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddRectFilledMultiColor(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col_upr_left uint32, Col_upr_right uint32, Col_bot_right uint32, Col_bot_left uint32) {
	getProc("ImDrawList_AddRectFilledMultiColor").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col_upr_left), uintptr(Col_upr_right), uintptr(Col_bot_right), uintptr(Col_bot_left))
}

func (i *Imgui) ImDrawListAddQuad(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Col uint32) {
	getProc("ImDrawList_AddQuad").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddQuadFilled(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Col uint32) {
	getProc("ImDrawList_AddQuadFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddTriangle(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, Col uint32) {
	getProc("ImDrawList_AddTriangle").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddTriangleFilled(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, Col uint32) {
	getProc("ImDrawList_AddTriangleFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddCircle(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32) {
	getProc("ImDrawList_AddCircle").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col))
}

func (i *Imgui) ImDrawListAddCircleFilled(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32) {
	getProc("ImDrawList_AddCircleFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col))
}

func (i *Imgui) ImDrawListAddNgon(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32, Num_segments int32) {
	getProc("ImDrawList_AddNgon").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddNgonFilled(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32, Num_segments int32) {
	getProc("ImDrawList_AddNgonFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddEllipse(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32) {
	getProc("ImDrawList_AddEllipse").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddEllipseFilled(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32) {
	getProc("ImDrawList_AddEllipseFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddText(Self *ImDrawList, Pos *ImVec2, Col uint32, Text_begin *int8) {
	getProc("ImDrawList_AddText").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Text_begin)))
}

func (i *Imgui) ImDrawListAddText2(Self *ImDrawList, Font *ImFont, Font_size float32, Pos *ImVec2, Col uint32, Text_begin *int8) {
	getProc("ImDrawList_AddText2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Font_size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Text_begin)))
}

func (i *Imgui) ImDrawListAddBezierCubic(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddBezierCubic").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddBezierQuadratic(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddBezierQuadratic").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddPolyline(Self *ImDrawList, Points *ImVec2, Num_points int32, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddPolyline").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Points)), uintptr(Num_points), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddConvexPolyFilled(Self *ImDrawList, Points *ImVec2, Num_points int32, Col uint32) {
	getProc("ImDrawList_AddConvexPolyFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Points)), uintptr(Num_points), uintptr(Col))
}

func (i *Imgui) ImDrawListAddConcavePolyFilled(Self *ImDrawList, Points *ImVec2, Num_points int32, Col uint32) {
	getProc("ImDrawList_AddConcavePolyFilled").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Points)), uintptr(Num_points), uintptr(Col))
}

func (i *Imgui) ImDrawListAddImage(Self *ImDrawList, Tex_ref *ImTextureRef, P_min *ImVec2, P_max *ImVec2) {
	getProc("ImDrawList_AddImage").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)))
}

func (i *Imgui) ImDrawListAddImageQuad(Self *ImDrawList, Tex_ref *ImTextureRef, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2) {
	getProc("ImDrawList_AddImageQuad").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)))
}

func (i *Imgui) ImDrawListAddImageRounded(Self *ImDrawList, Tex_ref *ImTextureRef, P_min *ImVec2, P_max *ImVec2, Uv_min *ImVec2, Uv_max *ImVec2, Col uint32, Rounding float32) {
	getProc("ImDrawList_AddImageRounded").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(unsafe.Pointer(Uv_min)), uintptr(unsafe.Pointer(Uv_max)), uintptr(Col), uintptr(Rounding))
}

func (i *Imgui) ImDrawListPathClear(Self *ImDrawList) {
	getProc("ImDrawList_PathClear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListPathLineTo(Self *ImDrawList, Pos *ImVec2) {
	getProc("ImDrawList_PathLineTo").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)))
}

func (i *Imgui) ImDrawListPathLineToMergeDuplicate(Self *ImDrawList, Pos *ImVec2) {
	getProc("ImDrawList_PathLineToMergeDuplicate").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)))
}

func (i *Imgui) ImDrawListPathFillConvex(Self *ImDrawList, Col uint32) {
	getProc("ImDrawList_PathFillConvex").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col))
}

func (i *Imgui) ImDrawListPathFillConcave(Self *ImDrawList, Col uint32) {
	getProc("ImDrawList_PathFillConcave").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col))
}

func (i *Imgui) ImDrawListPathStroke(Self *ImDrawList, Col uint32) {
	getProc("ImDrawList_PathStroke").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col))
}

func (i *Imgui) ImDrawListPathArcTo(Self *ImDrawList, Center *ImVec2, Radius float32, A_min float32, A_max float32) {
	getProc("ImDrawList_PathArcTo").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(A_min), uintptr(A_max))
}

func (i *Imgui) ImDrawListPathArcToFast(Self *ImDrawList, Center *ImVec2, Radius float32, A_min_of_12 int32, A_max_of_12 int32) {
	getProc("ImDrawList_PathArcToFast").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(A_min_of_12), uintptr(A_max_of_12))
}

func (i *Imgui) ImDrawListPathEllipticalArcTo(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Rot float32, A_min float32, A_max float32) {
	getProc("ImDrawList_PathEllipticalArcTo").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Rot), uintptr(A_min), uintptr(A_max))
}

func (i *Imgui) ImDrawListPathBezierCubicCurveTo(Self *ImDrawList, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2) {
	getProc("ImDrawList_PathBezierCubicCurveTo").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)))
}

func (i *Imgui) ImDrawListPathBezierQuadraticCurveTo(Self *ImDrawList, P2 *ImVec2, P3 *ImVec2) {
	getProc("ImDrawList_PathBezierQuadraticCurveTo").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)))
}

func (i *Imgui) ImDrawListPathRect(Self *ImDrawList, Rect_min *ImVec2, Rect_max *ImVec2) {
	getProc("ImDrawList_PathRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Rect_min)), uintptr(unsafe.Pointer(Rect_max)))
}

func (i *Imgui) ImDrawListAddDrawCmd(Self *ImDrawList) {
	getProc("ImDrawList_AddDrawCmd").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListCloneOutput(Self *ImDrawList) *ImDrawList {
	r1, _, _ := getProc("ImDrawList_CloneOutput").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImDrawList)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawListChannelsSplit(Self *ImDrawList, Count int32) {
	getProc("ImDrawList_ChannelsSplit").Call(uintptr(unsafe.Pointer(Self)), uintptr(Count))
}

func (i *Imgui) ImDrawListChannelsMerge(Self *ImDrawList) {
	getProc("ImDrawList_ChannelsMerge").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListChannelsSetCurrent(Self *ImDrawList, N int32) {
	getProc("ImDrawList_ChannelsSetCurrent").Call(uintptr(unsafe.Pointer(Self)), uintptr(N))
}

func (i *Imgui) ImDrawListPrimReserve(Self *ImDrawList, Idx_count int32, Vtx_count int32) {
	getProc("ImDrawList_PrimReserve").Call(uintptr(unsafe.Pointer(Self)), uintptr(Idx_count), uintptr(Vtx_count))
}

func (i *Imgui) ImDrawListPrimUnreserve(Self *ImDrawList, Idx_count int32, Vtx_count int32) {
	getProc("ImDrawList_PrimUnreserve").Call(uintptr(unsafe.Pointer(Self)), uintptr(Idx_count), uintptr(Vtx_count))
}

func (i *Imgui) ImDrawListPrimRect(Self *ImDrawList, A *ImVec2, B *ImVec2, Col uint32) {
	getProc("ImDrawList_PrimRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(Col))
}

func (i *Imgui) ImDrawListPrimRectUV(Self *ImDrawList, A *ImVec2, B *ImVec2, Uv_a *ImVec2, Uv_b *ImVec2, Col uint32) {
	getProc("ImDrawList_PrimRectUV").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(unsafe.Pointer(Uv_a)), uintptr(unsafe.Pointer(Uv_b)), uintptr(Col))
}

func (i *Imgui) ImDrawListPrimQuadUV(Self *ImDrawList, A *ImVec2, B *ImVec2, C *ImVec2, D *ImVec2, Uv_a *ImVec2, Uv_b *ImVec2, Uv_c *ImVec2, Uv_d *ImVec2, Col uint32) {
	getProc("ImDrawList_PrimQuadUV").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(unsafe.Pointer(C)), uintptr(unsafe.Pointer(D)), uintptr(unsafe.Pointer(Uv_a)), uintptr(unsafe.Pointer(Uv_b)), uintptr(unsafe.Pointer(Uv_c)), uintptr(unsafe.Pointer(Uv_d)), uintptr(Col))
}

func (i *Imgui) ImDrawListPrimWriteVtx(Self *ImDrawList, Pos *ImVec2, Uv *ImVec2, Col uint32) {
	getProc("ImDrawList_PrimWriteVtx").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)), uintptr(unsafe.Pointer(Uv)), uintptr(Col))
}

func (i *Imgui) ImDrawListPrimWriteIdx(Self *ImDrawList, Idx uint16) {
	getProc("ImDrawList_PrimWriteIdx").Call(uintptr(unsafe.Pointer(Self)), uintptr(Idx))
}

func (i *Imgui) ImDrawListPrimVtx(Self *ImDrawList, Pos *ImVec2, Uv *ImVec2, Col uint32) {
	getProc("ImDrawList_PrimVtx").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)), uintptr(unsafe.Pointer(Uv)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddRect2(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32, Rounding float32, Flags int32, Thickness float32) {
	getProc("ImDrawList_AddRect2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col), uintptr(Rounding), uintptr(Flags), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddPolyline2(Self *ImDrawList, Points *ImVec2, Num_points int32, Col uint32, Flags int32, Thickness float32) {
	getProc("ImDrawList_AddPolyline2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Points)), uintptr(Num_points), uintptr(Col), uintptr(Flags), uintptr(Thickness))
}

func (i *Imgui) ImDrawListPathStroke2(Self *ImDrawList, Col uint32, Flags int32, Thickness float32) {
	getProc("ImDrawList_PathStroke2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col), uintptr(Flags), uintptr(Thickness))
}

func (i *Imgui) ImDrawListPushTextureID(Self *ImDrawList, Tex_ref *ImTextureRef) {
	getProc("ImDrawList_PushTextureID").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)))
}

func (i *Imgui) ImDrawListPopTextureID(Self *ImDrawList) {
	getProc("ImDrawList_PopTextureID").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListSetDrawListSharedData(Self *ImDrawList, Data unsafe.Pointer) {
	getProc("ImDrawList__SetDrawListSharedData").Call(uintptr(unsafe.Pointer(Self)), uintptr(Data))
}

func (i *Imgui) ImDrawListResetForNewFrame(Self *ImDrawList) {
	getProc("ImDrawList__ResetForNewFrame").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListClearFreeMemory(Self *ImDrawList) {
	getProc("ImDrawList__ClearFreeMemory").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListPopUnusedDrawCmd(Self *ImDrawList) {
	getProc("ImDrawList__PopUnusedDrawCmd").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListTryMergeDrawCmds(Self *ImDrawList) {
	getProc("ImDrawList__TryMergeDrawCmds").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListOnChangedClipRect(Self *ImDrawList) {
	getProc("ImDrawList__OnChangedClipRect").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListOnChangedTexture(Self *ImDrawList) {
	getProc("ImDrawList__OnChangedTexture").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListOnChangedVtxOffset(Self *ImDrawList) {
	getProc("ImDrawList__OnChangedVtxOffset").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawListSetTexture(Self *ImDrawList, Tex_ref *ImTextureRef) {
	getProc("ImDrawList__SetTexture").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)))
}

func (i *Imgui) ImDrawListCalcCircleAutoSegmentCount(Self *ImDrawList, Radius float32) int32 {
	r1, _, _ := getProc("ImDrawList__CalcCircleAutoSegmentCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(Radius))
	return int32(r1)
}

func (i *Imgui) ImDrawListPathArcToFastEx(Self *ImDrawList, Center *ImVec2, Radius float32, A_min_sample int32, A_max_sample int32, A_step int32) {
	getProc("ImDrawList__PathArcToFastEx").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(A_min_sample), uintptr(A_max_sample), uintptr(A_step))
}

func (i *Imgui) ImDrawListPathArcToN(Self *ImDrawList, Center *ImVec2, Radius float32, A_min float32, A_max float32, Num_segments int32) {
	getProc("ImDrawList__PathArcToN").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(A_min), uintptr(A_max), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListOperatorAssign(Self *ImDrawList, Param1 *ImDrawList) {
	getProc("ImDrawList_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImDrawListPushClipRect2(Self *ImDrawList, Clip_rect_min *ImVec2, Clip_rect_max *ImVec2, Intersect_with_current_clip_rect bool) {
	getProc("ImDrawList_PushClipRect2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Clip_rect_min)), uintptr(unsafe.Pointer(Clip_rect_max)), func() uintptr {
		if Intersect_with_current_clip_rect {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImDrawListAddLine2(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddLine2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddLineH2(Self *ImDrawList, Min_x float32, Max_x float32, Y float32, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddLineH2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Min_x), uintptr(Max_x), uintptr(Y), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddLineV2(Self *ImDrawList, X float32, Min_y float32, Max_y float32, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddLineV2").Call(uintptr(unsafe.Pointer(Self)), uintptr(X), uintptr(Min_y), uintptr(Max_y), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddRect3(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32, Rounding float32) {
	getProc("ImDrawList_AddRect3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col), uintptr(Rounding))
}

func (i *Imgui) ImDrawListAddRect4(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32, Rounding float32, Thickness float32) {
	getProc("ImDrawList_AddRect4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col), uintptr(Rounding), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddRect5(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32, Rounding float32, Thickness float32, Flags int32) {
	getProc("ImDrawList_AddRect5").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col), uintptr(Rounding), uintptr(Thickness), uintptr(Flags))
}

func (i *Imgui) ImDrawListAddRectFilled2(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32, Rounding float32) {
	getProc("ImDrawList_AddRectFilled2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col), uintptr(Rounding))
}

func (i *Imgui) ImDrawListAddRectFilled3(Self *ImDrawList, P_min *ImVec2, P_max *ImVec2, Col uint32, Rounding float32, Flags int32) {
	getProc("ImDrawList_AddRectFilled3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(Col), uintptr(Rounding), uintptr(Flags))
}

func (i *Imgui) ImDrawListAddQuad2(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddQuad2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddTriangle2(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, Col uint32, Thickness float32) {
	getProc("ImDrawList_AddTriangle2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddCircle2(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32, Num_segments int32) {
	getProc("ImDrawList_AddCircle2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddCircle3(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32, Num_segments int32, Thickness float32) {
	getProc("ImDrawList_AddCircle3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col), uintptr(Num_segments), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddCircleFilled2(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32, Num_segments int32) {
	getProc("ImDrawList_AddCircleFilled2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddNgon2(Self *ImDrawList, Center *ImVec2, Radius float32, Col uint32, Num_segments int32, Thickness float32) {
	getProc("ImDrawList_AddNgon2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(Col), uintptr(Num_segments), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddEllipse2(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32, Rot float32) {
	getProc("ImDrawList_AddEllipse2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col), uintptr(Rot))
}

func (i *Imgui) ImDrawListAddEllipse3(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32, Rot float32, Num_segments int32) {
	getProc("ImDrawList_AddEllipse3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col), uintptr(Rot), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddEllipse4(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32, Rot float32, Num_segments int32, Thickness float32) {
	getProc("ImDrawList_AddEllipse4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col), uintptr(Rot), uintptr(Num_segments), uintptr(Thickness))
}

func (i *Imgui) ImDrawListAddEllipseFilled2(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32, Rot float32) {
	getProc("ImDrawList_AddEllipseFilled2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col), uintptr(Rot))
}

func (i *Imgui) ImDrawListAddEllipseFilled3(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Col uint32, Rot float32, Num_segments int32) {
	getProc("ImDrawList_AddEllipseFilled3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Col), uintptr(Rot), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddText3(Self *ImDrawList, Pos *ImVec2, Col uint32, Text_begin *int8, Text_end *int8) {
	getProc("ImDrawList_AddText3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)))
}

func (i *Imgui) ImDrawListAddText4(Self *ImDrawList, Font *ImFont, Font_size float32, Pos *ImVec2, Col uint32, Text_begin *int8, Text_end *int8) {
	getProc("ImDrawList_AddText4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Font_size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)))
}

func (i *Imgui) ImDrawListAddText5(Self *ImDrawList, Font *ImFont, Font_size float32, Pos *ImVec2, Col uint32, Text_begin *int8, Text_end *int8, Wrap_width float32) {
	getProc("ImDrawList_AddText5").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Font_size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)), uintptr(Wrap_width))
}

func (i *Imgui) ImDrawListAddText6(Self *ImDrawList, Font *ImFont, Font_size float32, Pos *ImVec2, Col uint32, Text_begin *int8, Text_end *int8, Wrap_width float32, Cpu_fine_clip_rect *ImVec4) {
	getProc("ImDrawList_AddText6").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Font_size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)), uintptr(Wrap_width), uintptr(unsafe.Pointer(Cpu_fine_clip_rect)))
}

func (i *Imgui) ImDrawListAddBezierCubic2(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Col uint32, Thickness float32, Num_segments int32) {
	getProc("ImDrawList_AddBezierCubic2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(Col), uintptr(Thickness), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddBezierQuadratic2(Self *ImDrawList, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, Col uint32, Thickness float32, Num_segments int32) {
	getProc("ImDrawList_AddBezierQuadratic2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(Col), uintptr(Thickness), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListAddPolyline3(Self *ImDrawList, Points *ImVec2, Num_points int32, Col uint32, Thickness float32, Flags int32) {
	getProc("ImDrawList_AddPolyline3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Points)), uintptr(Num_points), uintptr(Col), uintptr(Thickness), uintptr(Flags))
}

func (i *Imgui) ImDrawListAddImage2(Self *ImDrawList, Tex_ref *ImTextureRef, P_min *ImVec2, P_max *ImVec2, Uv_min *ImVec2) {
	getProc("ImDrawList_AddImage2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(unsafe.Pointer(Uv_min)))
}

func (i *Imgui) ImDrawListAddImage3(Self *ImDrawList, Tex_ref *ImTextureRef, P_min *ImVec2, P_max *ImVec2, Uv_min *ImVec2, Uv_max *ImVec2) {
	getProc("ImDrawList_AddImage3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(unsafe.Pointer(Uv_min)), uintptr(unsafe.Pointer(Uv_max)))
}

func (i *Imgui) ImDrawListAddImage4(Self *ImDrawList, Tex_ref *ImTextureRef, P_min *ImVec2, P_max *ImVec2, Uv_min *ImVec2, Uv_max *ImVec2, Col uint32) {
	getProc("ImDrawList_AddImage4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(unsafe.Pointer(Uv_min)), uintptr(unsafe.Pointer(Uv_max)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddImageQuad2(Self *ImDrawList, Tex_ref *ImTextureRef, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Uv1 *ImVec2) {
	getProc("ImDrawList_AddImageQuad2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(unsafe.Pointer(Uv1)))
}

func (i *Imgui) ImDrawListAddImageQuad3(Self *ImDrawList, Tex_ref *ImTextureRef, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Uv1 *ImVec2, Uv2 *ImVec2) {
	getProc("ImDrawList_AddImageQuad3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Uv2)))
}

func (i *Imgui) ImDrawListAddImageQuad4(Self *ImDrawList, Tex_ref *ImTextureRef, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Uv1 *ImVec2, Uv2 *ImVec2, Uv3 *ImVec2) {
	getProc("ImDrawList_AddImageQuad4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Uv2)), uintptr(unsafe.Pointer(Uv3)))
}

func (i *Imgui) ImDrawListAddImageQuad5(Self *ImDrawList, Tex_ref *ImTextureRef, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Uv1 *ImVec2, Uv2 *ImVec2, Uv3 *ImVec2, Uv4 *ImVec2) {
	getProc("ImDrawList_AddImageQuad5").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Uv2)), uintptr(unsafe.Pointer(Uv3)), uintptr(unsafe.Pointer(Uv4)))
}

func (i *Imgui) ImDrawListAddImageQuad6(Self *ImDrawList, Tex_ref *ImTextureRef, P1 *ImVec2, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Uv1 *ImVec2, Uv2 *ImVec2, Uv3 *ImVec2, Uv4 *ImVec2, Col uint32) {
	getProc("ImDrawList_AddImageQuad6").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P1)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Uv2)), uintptr(unsafe.Pointer(Uv3)), uintptr(unsafe.Pointer(Uv4)), uintptr(Col))
}

func (i *Imgui) ImDrawListAddImageRounded2(Self *ImDrawList, Tex_ref *ImTextureRef, P_min *ImVec2, P_max *ImVec2, Uv_min *ImVec2, Uv_max *ImVec2, Col uint32, Rounding float32, Flags int32) {
	getProc("ImDrawList_AddImageRounded2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(P_min)), uintptr(unsafe.Pointer(P_max)), uintptr(unsafe.Pointer(Uv_min)), uintptr(unsafe.Pointer(Uv_max)), uintptr(Col), uintptr(Rounding), uintptr(Flags))
}

func (i *Imgui) ImDrawListPathStroke3(Self *ImDrawList, Col uint32, Thickness float32) {
	getProc("ImDrawList_PathStroke3").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col), uintptr(Thickness))
}

func (i *Imgui) ImDrawListPathStroke4(Self *ImDrawList, Col uint32, Thickness float32, Flags int32) {
	getProc("ImDrawList_PathStroke4").Call(uintptr(unsafe.Pointer(Self)), uintptr(Col), uintptr(Thickness), uintptr(Flags))
}

func (i *Imgui) ImDrawListPathArcTo2(Self *ImDrawList, Center *ImVec2, Radius float32, A_min float32, A_max float32, Num_segments int32) {
	getProc("ImDrawList_PathArcTo2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(Radius), uintptr(A_min), uintptr(A_max), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListPathEllipticalArcTo2(Self *ImDrawList, Center *ImVec2, Radius *ImVec2, Rot float32, A_min float32, A_max float32, Num_segments int32) {
	getProc("ImDrawList_PathEllipticalArcTo2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Center)), uintptr(unsafe.Pointer(Radius)), uintptr(Rot), uintptr(A_min), uintptr(A_max), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListPathBezierCubicCurveTo2(Self *ImDrawList, P2 *ImVec2, P3 *ImVec2, P4 *ImVec2, Num_segments int32) {
	getProc("ImDrawList_PathBezierCubicCurveTo2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(unsafe.Pointer(P4)), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListPathBezierQuadraticCurveTo2(Self *ImDrawList, P2 *ImVec2, P3 *ImVec2, Num_segments int32) {
	getProc("ImDrawList_PathBezierQuadraticCurveTo2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(P2)), uintptr(unsafe.Pointer(P3)), uintptr(Num_segments))
}

func (i *Imgui) ImDrawListPathRect2(Self *ImDrawList, Rect_min *ImVec2, Rect_max *ImVec2, Rounding float32) {
	getProc("ImDrawList_PathRect2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Rect_min)), uintptr(unsafe.Pointer(Rect_max)), uintptr(Rounding))
}

func (i *Imgui) ImDrawListPathRect3(Self *ImDrawList, Rect_min *ImVec2, Rect_max *ImVec2, Rounding float32, Flags int32) {
	getProc("ImDrawList_PathRect3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Rect_min)), uintptr(unsafe.Pointer(Rect_max)), uintptr(Rounding), uintptr(Flags))
}

func (i *Imgui) ImDrawListDelete(Self *ImDrawList) {
	getProc("ImDrawList_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawDataNew() *ImDrawData {
	r1, _, _ := getProc("ImDrawData_new").Call()
	return (*ImDrawData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawDataNew2(Param1 *ImDrawData) *ImDrawData {
	r1, _, _ := getProc("ImDrawData_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImDrawData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawDataValid(Self *ImDrawData) bool {
	r1, _, _ := getProc("ImDrawData_Valid").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImDrawDataSetValid(Self *ImDrawData, Valid bool) {
	getProc("ImDrawData_setValid").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Valid {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImDrawDataCmdListsCount(Self *ImDrawData) int32 {
	r1, _, _ := getProc("ImDrawData_CmdListsCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawDataSetCmdListsCount(Self *ImDrawData, CmdListsCount int32) {
	getProc("ImDrawData_setCmdListsCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(CmdListsCount))
}

func (i *Imgui) ImDrawDataTotalIdxCount(Self *ImDrawData) int32 {
	r1, _, _ := getProc("ImDrawData_TotalIdxCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawDataSetTotalIdxCount(Self *ImDrawData, TotalIdxCount int32) {
	getProc("ImDrawData_setTotalIdxCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(TotalIdxCount))
}

func (i *Imgui) ImDrawDataTotalVtxCount(Self *ImDrawData) int32 {
	r1, _, _ := getProc("ImDrawData_TotalVtxCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImDrawDataSetTotalVtxCount(Self *ImDrawData, TotalVtxCount int32) {
	getProc("ImDrawData_setTotalVtxCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(TotalVtxCount))
}

func (i *Imgui) ImDrawDataCmdLists(Self *ImDrawData) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawData_CmdLists").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawDataSetCmdLists(Self *ImDrawData, CmdLists unsafe.Pointer) {
	getProc("ImDrawData_setCmdLists").Call(uintptr(unsafe.Pointer(Self)), uintptr(CmdLists))
}

func (i *Imgui) ImDrawDataDisplayPos(Self *ImDrawData) *ImVec2 {
	r1, _, _ := getProc("ImDrawData_DisplayPos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawDataSetDisplayPos(Self *ImDrawData, DisplayPos *ImVec2) {
	getProc("ImDrawData_setDisplayPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DisplayPos)))
}

func (i *Imgui) ImDrawDataDisplaySize(Self *ImDrawData) *ImVec2 {
	r1, _, _ := getProc("ImDrawData_DisplaySize").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawDataSetDisplaySize(Self *ImDrawData, DisplaySize *ImVec2) {
	getProc("ImDrawData_setDisplaySize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DisplaySize)))
}

func (i *Imgui) ImDrawDataFramebufferScale(Self *ImDrawData) *ImVec2 {
	r1, _, _ := getProc("ImDrawData_FramebufferScale").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawDataSetFramebufferScale(Self *ImDrawData, FramebufferScale *ImVec2) {
	getProc("ImDrawData_setFramebufferScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FramebufferScale)))
}

func (i *Imgui) ImDrawDataOwnerViewport(Self *ImDrawData) *ImGuiViewport {
	r1, _, _ := getProc("ImDrawData_OwnerViewport").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImGuiViewport)(unsafe.Pointer(r1))
}

func (i *Imgui) ImDrawDataSetOwnerViewport(Self *ImDrawData, OwnerViewport *ImGuiViewport) {
	getProc("ImDrawData_setOwnerViewport").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(OwnerViewport)))
}

func (i *Imgui) ImDrawDataTextures(Self *ImDrawData) unsafe.Pointer {
	r1, _, _ := getProc("ImDrawData_Textures").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImDrawDataSetTextures(Self *ImDrawData, Textures unsafe.Pointer) {
	getProc("ImDrawData_setTextures").Call(uintptr(unsafe.Pointer(Self)), uintptr(Textures))
}

func (i *Imgui) ImDrawDataClear(Self *ImDrawData) {
	getProc("ImDrawData_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawDataAddDrawList(Self *ImDrawData, Draw_list *ImDrawList) {
	getProc("ImDrawData_AddDrawList").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)))
}

func (i *Imgui) ImDrawDataDeIndexAllBuffers(Self *ImDrawData) {
	getProc("ImDrawData_DeIndexAllBuffers").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImDrawDataScaleClipRects(Self *ImDrawData, Fb_scale *ImVec2) {
	getProc("ImDrawData_ScaleClipRects").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Fb_scale)))
}

func (i *Imgui) ImDrawDataOperatorAssign(Self *ImDrawData, Param1 *ImDrawData) {
	getProc("ImDrawData_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImDrawDataDelete(Self *ImDrawData) {
	getProc("ImDrawData_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImTextureRectNew(Param1 *ImTextureRect) *ImTextureRect {
	r1, _, _ := getProc("ImTextureRect_new").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImTextureRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRectNew2() *ImTextureRect {
	r1, _, _ := getProc("ImTextureRect_new2").Call()
	return (*ImTextureRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureRectX(Self *ImTextureRect) uint16 {
	r1, _, _ := getProc("ImTextureRect_x").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImTextureRectSetX(Self *ImTextureRect, X uint16) {
	getProc("ImTextureRect_setX").Call(uintptr(unsafe.Pointer(Self)), uintptr(X))
}

func (i *Imgui) ImTextureRectY(Self *ImTextureRect) uint16 {
	r1, _, _ := getProc("ImTextureRect_y").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImTextureRectSetY(Self *ImTextureRect, Y uint16) {
	getProc("ImTextureRect_setY").Call(uintptr(unsafe.Pointer(Self)), uintptr(Y))
}

func (i *Imgui) ImTextureRectW(Self *ImTextureRect) uint16 {
	r1, _, _ := getProc("ImTextureRect_w").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImTextureRectSetW(Self *ImTextureRect, W uint16) {
	getProc("ImTextureRect_setW").Call(uintptr(unsafe.Pointer(Self)), uintptr(W))
}

func (i *Imgui) ImTextureRectH(Self *ImTextureRect) uint16 {
	r1, _, _ := getProc("ImTextureRect_h").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImTextureRectSetH(Self *ImTextureRect, H uint16) {
	getProc("ImTextureRect_setH").Call(uintptr(unsafe.Pointer(Self)), uintptr(H))
}

func (i *Imgui) ImTextureRectOperatorAssign(Self *ImTextureRect, Param1 *ImTextureRect) {
	getProc("ImTextureRect_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImTextureRectDelete(Self *ImTextureRect) {
	getProc("ImTextureRect_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImTextureDataNew() *ImTextureData {
	r1, _, _ := getProc("ImTextureData_new").Call()
	return (*ImTextureData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureDataNew2(Param1 *ImTextureData) *ImTextureData {
	r1, _, _ := getProc("ImTextureData_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImTextureData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureDataUniqueID(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_UniqueID").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetUniqueID(Self *ImTextureData, UniqueID int32) {
	getProc("ImTextureData_setUniqueID").Call(uintptr(unsafe.Pointer(Self)), uintptr(UniqueID))
}

func (i *Imgui) ImTextureDataStatus(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_Status").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetStatus(Self *ImTextureData, Status int32) {
	getProc("ImTextureData_setStatus").Call(uintptr(unsafe.Pointer(Self)), uintptr(Status))
}

func (i *Imgui) ImTextureDataTexID(Self *ImTextureData) uint64 {
	r1, _, _ := getProc("ImTextureData_TexID").Call(uintptr(unsafe.Pointer(Self)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImTextureDataSetTexID(Self *ImTextureData, TexID uint64) {
	getProc("ImTextureData_setTexID").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&TexID)))
}

func (i *Imgui) ImTextureDataFormat(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_Format").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetFormat(Self *ImTextureData, Format int32) {
	getProc("ImTextureData_setFormat").Call(uintptr(unsafe.Pointer(Self)), uintptr(Format))
}

func (i *Imgui) ImTextureDataWidth(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_Width").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetWidth(Self *ImTextureData, Width int32) {
	getProc("ImTextureData_setWidth").Call(uintptr(unsafe.Pointer(Self)), uintptr(Width))
}

func (i *Imgui) ImTextureDataHeight(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_Height").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetHeight(Self *ImTextureData, Height int32) {
	getProc("ImTextureData_setHeight").Call(uintptr(unsafe.Pointer(Self)), uintptr(Height))
}

func (i *Imgui) ImTextureDataBytesPerPixel(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_BytesPerPixel").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetBytesPerPixel(Self *ImTextureData, BytesPerPixel int32) {
	getProc("ImTextureData_setBytesPerPixel").Call(uintptr(unsafe.Pointer(Self)), uintptr(BytesPerPixel))
}

func (i *Imgui) ImTextureDataPixels(Self *ImTextureData) *uint8 {
	r1, _, _ := getProc("ImTextureData_Pixels").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureDataSetPixels(Self *ImTextureData, Pixels *uint8) {
	getProc("ImTextureData_setPixels").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pixels)))
}

func (i *Imgui) ImTextureDataUsedRect(Self *ImTextureData) *ImTextureRect {
	r1, _, _ := getProc("ImTextureData_UsedRect").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureDataSetUsedRect(Self *ImTextureData, UsedRect *ImTextureRect) {
	getProc("ImTextureData_setUsedRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(UsedRect)))
}

func (i *Imgui) ImTextureDataUpdateRect(Self *ImTextureData) *ImTextureRect {
	r1, _, _ := getProc("ImTextureData_UpdateRect").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureDataSetUpdateRect(Self *ImTextureData, UpdateRect *ImTextureRect) {
	getProc("ImTextureData_setUpdateRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(UpdateRect)))
}

func (i *Imgui) ImTextureDataUpdates(Self *ImTextureData) unsafe.Pointer {
	r1, _, _ := getProc("ImTextureData_Updates").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImTextureDataSetUpdates(Self *ImTextureData, Updates unsafe.Pointer) {
	getProc("ImTextureData_setUpdates").Call(uintptr(unsafe.Pointer(Self)), uintptr(Updates))
}

func (i *Imgui) ImTextureDataUnusedFrames(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_UnusedFrames").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataSetUnusedFrames(Self *ImTextureData, UnusedFrames int32) {
	getProc("ImTextureData_setUnusedFrames").Call(uintptr(unsafe.Pointer(Self)), uintptr(UnusedFrames))
}

func (i *Imgui) ImTextureDataRefCount(Self *ImTextureData) uint16 {
	r1, _, _ := getProc("ImTextureData_RefCount").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImTextureDataSetRefCount(Self *ImTextureData, RefCount uint16) {
	getProc("ImTextureData_setRefCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(RefCount))
}

func (i *Imgui) ImTextureDataUseColors(Self *ImTextureData) bool {
	r1, _, _ := getProc("ImTextureData_UseColors").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImTextureDataSetUseColors(Self *ImTextureData, UseColors bool) {
	getProc("ImTextureData_setUseColors").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if UseColors {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImTextureDataWantDestroyNextFrame(Self *ImTextureData) bool {
	r1, _, _ := getProc("ImTextureData_WantDestroyNextFrame").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImTextureDataSetWantDestroyNextFrame(Self *ImTextureData, WantDestroyNextFrame bool) {
	getProc("ImTextureData_setWantDestroyNextFrame").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantDestroyNextFrame {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImTextureDataCreate(Self *ImTextureData, Format int32, W int32, H int32) {
	getProc("ImTextureData_Create").Call(uintptr(unsafe.Pointer(Self)), uintptr(Format), uintptr(W), uintptr(H))
}

func (i *Imgui) ImTextureDataDestroyPixels(Self *ImTextureData) {
	getProc("ImTextureData_DestroyPixels").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImTextureDataGetPixels(Self *ImTextureData) unsafe.Pointer {
	r1, _, _ := getProc("ImTextureData_GetPixels").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImTextureDataGetPixelsAt(Self *ImTextureData, X int32, Y int32) unsafe.Pointer {
	r1, _, _ := getProc("ImTextureData_GetPixelsAt").Call(uintptr(unsafe.Pointer(Self)), uintptr(X), uintptr(Y))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImTextureDataGetSizeInBytes(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_GetSizeInBytes").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataGetPitch(Self *ImTextureData) int32 {
	r1, _, _ := getProc("ImTextureData_GetPitch").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImTextureDataGetTexRef(Self *ImTextureData) *ImTextureRef {
	r1, _, _ := getProc("ImTextureData_GetTexRef").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureRef)(unsafe.Pointer(r1))
}

func (i *Imgui) ImTextureDataGetTexID(Self *ImTextureData) uint64 {
	r1, _, _ := getProc("ImTextureData_GetTexID").Call(uintptr(unsafe.Pointer(Self)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImTextureDataSetTexID_1(Self *ImTextureData, Tex_id uint64) {
	getProc("ImTextureData_SetTexID").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&Tex_id)))
}

func (i *Imgui) ImTextureDataSetStatus_1(Self *ImTextureData, Status int32) {
	getProc("ImTextureData_SetStatus").Call(uintptr(unsafe.Pointer(Self)), uintptr(Status))
}

func (i *Imgui) ImTextureDataOperatorAssign(Self *ImTextureData, Param1 *ImTextureData) {
	getProc("ImTextureData_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImTextureDataDelete(Self *ImTextureData) {
	getProc("ImTextureData_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontConfigNew() *ImFontConfig {
	r1, _, _ := getProc("ImFontConfig_new").Call()
	return (*ImFontConfig)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontConfigFontDataSize(Self *ImFontConfig) int32 {
	r1, _, _ := getProc("ImFontConfig_FontDataSize").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontConfigSetFontDataSize(Self *ImFontConfig, FontDataSize int32) {
	getProc("ImFontConfig_setFontDataSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontDataSize))
}

func (i *Imgui) ImFontConfigFontDataOwnedByAtlas(Self *ImFontConfig) bool {
	r1, _, _ := getProc("ImFontConfig_FontDataOwnedByAtlas").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontConfigSetFontDataOwnedByAtlas(Self *ImFontConfig, FontDataOwnedByAtlas bool) {
	getProc("ImFontConfig_setFontDataOwnedByAtlas").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if FontDataOwnedByAtlas {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontConfigMergeMode(Self *ImFontConfig) bool {
	r1, _, _ := getProc("ImFontConfig_MergeMode").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontConfigSetMergeMode(Self *ImFontConfig, MergeMode bool) {
	getProc("ImFontConfig_setMergeMode").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if MergeMode {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontConfigPixelSnapH(Self *ImFontConfig) bool {
	r1, _, _ := getProc("ImFontConfig_PixelSnapH").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontConfigSetPixelSnapH(Self *ImFontConfig, PixelSnapH bool) {
	getProc("ImFontConfig_setPixelSnapH").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if PixelSnapH {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontConfigOversampleH(Self *ImFontConfig) int8 {
	r1, _, _ := getProc("ImFontConfig_OversampleH").Call(uintptr(unsafe.Pointer(Self)))
	return int8(r1)
}

func (i *Imgui) ImFontConfigSetOversampleH(Self *ImFontConfig, OversampleH int8) {
	getProc("ImFontConfig_setOversampleH").Call(uintptr(unsafe.Pointer(Self)), uintptr(OversampleH))
}

func (i *Imgui) ImFontConfigOversampleV(Self *ImFontConfig) int8 {
	r1, _, _ := getProc("ImFontConfig_OversampleV").Call(uintptr(unsafe.Pointer(Self)))
	return int8(r1)
}

func (i *Imgui) ImFontConfigSetOversampleV(Self *ImFontConfig, OversampleV int8) {
	getProc("ImFontConfig_setOversampleV").Call(uintptr(unsafe.Pointer(Self)), uintptr(OversampleV))
}

func (i *Imgui) ImFontConfigEllipsisChar(Self *ImFontConfig) uint16 {
	r1, _, _ := getProc("ImFontConfig_EllipsisChar").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontConfigSetEllipsisChar(Self *ImFontConfig, EllipsisChar uint16) {
	getProc("ImFontConfig_setEllipsisChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(EllipsisChar))
}

func (i *Imgui) ImFontConfigSizePixels(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_SizePixels").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetSizePixels(Self *ImFontConfig, SizePixels float32) {
	getProc("ImFontConfig_setSizePixels").Call(uintptr(unsafe.Pointer(Self)), uintptr(SizePixels))
}

func (i *Imgui) ImFontConfigGlyphRanges(Self *ImFontConfig) *uint16 {
	r1, _, _ := getProc("ImFontConfig_GlyphRanges").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontConfigSetGlyphRanges(Self *ImFontConfig, GlyphRanges *uint16) {
	getProc("ImFontConfig_setGlyphRanges").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(GlyphRanges)))
}

func (i *Imgui) ImFontConfigGlyphExcludeRanges(Self *ImFontConfig) *uint16 {
	r1, _, _ := getProc("ImFontConfig_GlyphExcludeRanges").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontConfigSetGlyphExcludeRanges(Self *ImFontConfig, GlyphExcludeRanges *uint16) {
	getProc("ImFontConfig_setGlyphExcludeRanges").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(GlyphExcludeRanges)))
}

func (i *Imgui) ImFontConfigGlyphOffset(Self *ImFontConfig) *ImVec2 {
	r1, _, _ := getProc("ImFontConfig_GlyphOffset").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontConfigSetGlyphOffset(Self *ImFontConfig, GlyphOffset *ImVec2) {
	getProc("ImFontConfig_setGlyphOffset").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(GlyphOffset)))
}

func (i *Imgui) ImFontConfigGlyphMinAdvanceX(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_GlyphMinAdvanceX").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetGlyphMinAdvanceX(Self *ImFontConfig, GlyphMinAdvanceX float32) {
	getProc("ImFontConfig_setGlyphMinAdvanceX").Call(uintptr(unsafe.Pointer(Self)), uintptr(GlyphMinAdvanceX))
}

func (i *Imgui) ImFontConfigGlyphMaxAdvanceX(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_GlyphMaxAdvanceX").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetGlyphMaxAdvanceX(Self *ImFontConfig, GlyphMaxAdvanceX float32) {
	getProc("ImFontConfig_setGlyphMaxAdvanceX").Call(uintptr(unsafe.Pointer(Self)), uintptr(GlyphMaxAdvanceX))
}

func (i *Imgui) ImFontConfigGlyphExtraAdvanceX(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_GlyphExtraAdvanceX").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetGlyphExtraAdvanceX(Self *ImFontConfig, GlyphExtraAdvanceX float32) {
	getProc("ImFontConfig_setGlyphExtraAdvanceX").Call(uintptr(unsafe.Pointer(Self)), uintptr(GlyphExtraAdvanceX))
}

func (i *Imgui) ImFontConfigFontNo(Self *ImFontConfig) uint32 {
	r1, _, _ := getProc("ImFontConfig_FontNo").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontConfigSetFontNo(Self *ImFontConfig, FontNo uint32) {
	getProc("ImFontConfig_setFontNo").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontNo))
}

func (i *Imgui) ImFontConfigFontLoaderFlags(Self *ImFontConfig) uint32 {
	r1, _, _ := getProc("ImFontConfig_FontLoaderFlags").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontConfigSetFontLoaderFlags(Self *ImFontConfig, FontLoaderFlags uint32) {
	getProc("ImFontConfig_setFontLoaderFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontLoaderFlags))
}

func (i *Imgui) ImFontConfigRasterizerMultiply(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_RasterizerMultiply").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetRasterizerMultiply(Self *ImFontConfig, RasterizerMultiply float32) {
	getProc("ImFontConfig_setRasterizerMultiply").Call(uintptr(unsafe.Pointer(Self)), uintptr(RasterizerMultiply))
}

func (i *Imgui) ImFontConfigRasterizerDensity(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_RasterizerDensity").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetRasterizerDensity(Self *ImFontConfig, RasterizerDensity float32) {
	getProc("ImFontConfig_setRasterizerDensity").Call(uintptr(unsafe.Pointer(Self)), uintptr(RasterizerDensity))
}

func (i *Imgui) ImFontConfigExtraSizeScale(Self *ImFontConfig) float32 {
	r1, _, _ := getProc("ImFontConfig_ExtraSizeScale").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontConfigSetExtraSizeScale(Self *ImFontConfig, ExtraSizeScale float32) {
	getProc("ImFontConfig_setExtraSizeScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(ExtraSizeScale))
}

func (i *Imgui) ImFontConfigFlags(Self *ImFontConfig) int32 {
	r1, _, _ := getProc("ImFontConfig_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontConfigSetFlags(Self *ImFontConfig, Flags int32) {
	getProc("ImFontConfig_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImFontConfigDstFont(Self *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontConfig_DstFont").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontConfigSetDstFont(Self *ImFontConfig, DstFont *ImFont) {
	getProc("ImFontConfig_setDstFont").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(DstFont)))
}

func (i *Imgui) ImFontConfigFontLoader(Self *ImFontConfig) *ImFontLoader {
	r1, _, _ := getProc("ImFontConfig_FontLoader").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFontLoader)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontConfigSetFontLoader(Self *ImFontConfig, FontLoader *ImFontLoader) {
	getProc("ImFontConfig_setFontLoader").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FontLoader)))
}

func (i *Imgui) ImFontConfigPixelSnapV(Self *ImFontConfig) bool {
	r1, _, _ := getProc("ImFontConfig_PixelSnapV").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontConfigSetPixelSnapV(Self *ImFontConfig, PixelSnapV bool) {
	getProc("ImFontConfig_setPixelSnapV").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if PixelSnapV {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontConfigDelete(Self *ImFontConfig) {
	getProc("ImFontConfig_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontGlyphNew() *ImFontGlyph {
	r1, _, _ := getProc("ImFontGlyph_new").Call()
	return (*ImFontGlyph)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontGlyphColored(Self *ImFontGlyph) uint32 {
	r1, _, _ := getProc("ImFontGlyph_Colored").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontGlyphSetColored(Self *ImFontGlyph, Colored uint32) {
	getProc("ImFontGlyph_setColored").Call(uintptr(unsafe.Pointer(Self)), uintptr(Colored))
}

func (i *Imgui) ImFontGlyphVisible(Self *ImFontGlyph) uint32 {
	r1, _, _ := getProc("ImFontGlyph_Visible").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontGlyphSetVisible(Self *ImFontGlyph, Visible uint32) {
	getProc("ImFontGlyph_setVisible").Call(uintptr(unsafe.Pointer(Self)), uintptr(Visible))
}

func (i *Imgui) ImFontGlyphSourceIdx(Self *ImFontGlyph) uint32 {
	r1, _, _ := getProc("ImFontGlyph_SourceIdx").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontGlyphSetSourceIdx(Self *ImFontGlyph, SourceIdx uint32) {
	getProc("ImFontGlyph_setSourceIdx").Call(uintptr(unsafe.Pointer(Self)), uintptr(SourceIdx))
}

func (i *Imgui) ImFontGlyphCodepoint(Self *ImFontGlyph) uint32 {
	r1, _, _ := getProc("ImFontGlyph_Codepoint").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontGlyphSetCodepoint(Self *ImFontGlyph, Codepoint uint32) {
	getProc("ImFontGlyph_setCodepoint").Call(uintptr(unsafe.Pointer(Self)), uintptr(Codepoint))
}

func (i *Imgui) ImFontGlyphAdvanceX(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_AdvanceX").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetAdvanceX(Self *ImFontGlyph, AdvanceX float32) {
	getProc("ImFontGlyph_setAdvanceX").Call(uintptr(unsafe.Pointer(Self)), uintptr(AdvanceX))
}

func (i *Imgui) ImFontGlyphX0(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_X0").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetX0(Self *ImFontGlyph, X0 float32) {
	getProc("ImFontGlyph_setX0").Call(uintptr(unsafe.Pointer(Self)), uintptr(X0))
}

func (i *Imgui) ImFontGlyphY0(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_Y0").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetY0(Self *ImFontGlyph, Y0 float32) {
	getProc("ImFontGlyph_setY0").Call(uintptr(unsafe.Pointer(Self)), uintptr(Y0))
}

func (i *Imgui) ImFontGlyphX1(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_X1").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetX1(Self *ImFontGlyph, X1 float32) {
	getProc("ImFontGlyph_setX1").Call(uintptr(unsafe.Pointer(Self)), uintptr(X1))
}

func (i *Imgui) ImFontGlyphY1(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_Y1").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetY1(Self *ImFontGlyph, Y1 float32) {
	getProc("ImFontGlyph_setY1").Call(uintptr(unsafe.Pointer(Self)), uintptr(Y1))
}

func (i *Imgui) ImFontGlyphU0(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_U0").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetU0(Self *ImFontGlyph, U0 float32) {
	getProc("ImFontGlyph_setU0").Call(uintptr(unsafe.Pointer(Self)), uintptr(U0))
}

func (i *Imgui) ImFontGlyphV0(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_V0").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetV0(Self *ImFontGlyph, V0 float32) {
	getProc("ImFontGlyph_setV0").Call(uintptr(unsafe.Pointer(Self)), uintptr(V0))
}

func (i *Imgui) ImFontGlyphU1(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_U1").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetU1(Self *ImFontGlyph, U1 float32) {
	getProc("ImFontGlyph_setU1").Call(uintptr(unsafe.Pointer(Self)), uintptr(U1))
}

func (i *Imgui) ImFontGlyphV1(Self *ImFontGlyph) float32 {
	r1, _, _ := getProc("ImFontGlyph_V1").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontGlyphSetV1(Self *ImFontGlyph, V1 float32) {
	getProc("ImFontGlyph_setV1").Call(uintptr(unsafe.Pointer(Self)), uintptr(V1))
}

func (i *Imgui) ImFontGlyphPackId(Self *ImFontGlyph) int32 {
	r1, _, _ := getProc("ImFontGlyph_PackId").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontGlyphSetPackId(Self *ImFontGlyph, PackId int32) {
	getProc("ImFontGlyph_setPackId").Call(uintptr(unsafe.Pointer(Self)), uintptr(PackId))
}

func (i *Imgui) ImFontGlyphDelete(Self *ImFontGlyph) {
	getProc("ImFontGlyph_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontGlyphRangesBuilderNew() *ImFontGlyphRangesBuilder {
	r1, _, _ := getProc("ImFontGlyphRangesBuilder_new").Call()
	return (*ImFontGlyphRangesBuilder)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontGlyphRangesBuilderNew2(Param1 *ImFontGlyphRangesBuilder) *ImFontGlyphRangesBuilder {
	r1, _, _ := getProc("ImFontGlyphRangesBuilder_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImFontGlyphRangesBuilder)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontGlyphRangesBuilderUsedChars(Self *ImFontGlyphRangesBuilder) unsafe.Pointer {
	r1, _, _ := getProc("ImFontGlyphRangesBuilder_UsedChars").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontGlyphRangesBuilderSetUsedChars(Self *ImFontGlyphRangesBuilder, UsedChars unsafe.Pointer) {
	getProc("ImFontGlyphRangesBuilder_setUsedChars").Call(uintptr(unsafe.Pointer(Self)), uintptr(UsedChars))
}

func (i *Imgui) ImFontGlyphRangesBuilderClear(Self *ImFontGlyphRangesBuilder) {
	getProc("ImFontGlyphRangesBuilder_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontGlyphRangesBuilderGetBit(Self *ImFontGlyphRangesBuilder, N uint64) bool {
	r1, _, _ := getProc("ImFontGlyphRangesBuilder_GetBit").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&N)))
	return r1 != 0
}

func (i *Imgui) ImFontGlyphRangesBuilderSetBit(Self *ImFontGlyphRangesBuilder, N uint64) {
	getProc("ImFontGlyphRangesBuilder_SetBit").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&N)))
}

func (i *Imgui) ImFontGlyphRangesBuilderAddChar(Self *ImFontGlyphRangesBuilder, C uint16) {
	getProc("ImFontGlyphRangesBuilder_AddChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
}

func (i *Imgui) ImFontGlyphRangesBuilderAddText(Self *ImFontGlyphRangesBuilder, Text *int8) {
	getProc("ImFontGlyphRangesBuilder_AddText").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Text)))
}

func (i *Imgui) ImFontGlyphRangesBuilderAddRanges(Self *ImFontGlyphRangesBuilder, Ranges *uint16) {
	getProc("ImFontGlyphRangesBuilder_AddRanges").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Ranges)))
}

func (i *Imgui) ImFontGlyphRangesBuilderBuildRanges(Self *ImFontGlyphRangesBuilder, Out_ranges unsafe.Pointer) {
	getProc("ImFontGlyphRangesBuilder_BuildRanges").Call(uintptr(unsafe.Pointer(Self)), uintptr(Out_ranges))
}

func (i *Imgui) ImFontGlyphRangesBuilderOperatorAssign(Self *ImFontGlyphRangesBuilder, Param1 *ImFontGlyphRangesBuilder) {
	getProc("ImFontGlyphRangesBuilder_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImFontGlyphRangesBuilderAddText2(Self *ImFontGlyphRangesBuilder, Text *int8, Text_end *int8) {
	getProc("ImFontGlyphRangesBuilder_AddText2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)))
}

func (i *Imgui) ImFontGlyphRangesBuilderDelete(Self *ImFontGlyphRangesBuilder) {
	getProc("ImFontGlyphRangesBuilder_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasRectNew() *ImFontAtlasRect {
	r1, _, _ := getProc("ImFontAtlasRect_new").Call()
	return (*ImFontAtlasRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasRectNew2(Param1 *ImFontAtlasRect) *ImFontAtlasRect {
	r1, _, _ := getProc("ImFontAtlasRect_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImFontAtlasRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasRectX(Self *ImFontAtlasRect) uint16 {
	r1, _, _ := getProc("ImFontAtlasRect_x").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontAtlasRectSetX(Self *ImFontAtlasRect, X uint16) {
	getProc("ImFontAtlasRect_setX").Call(uintptr(unsafe.Pointer(Self)), uintptr(X))
}

func (i *Imgui) ImFontAtlasRectY(Self *ImFontAtlasRect) uint16 {
	r1, _, _ := getProc("ImFontAtlasRect_y").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontAtlasRectSetY(Self *ImFontAtlasRect, Y uint16) {
	getProc("ImFontAtlasRect_setY").Call(uintptr(unsafe.Pointer(Self)), uintptr(Y))
}

func (i *Imgui) ImFontAtlasRectW(Self *ImFontAtlasRect) uint16 {
	r1, _, _ := getProc("ImFontAtlasRect_w").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontAtlasRectSetW(Self *ImFontAtlasRect, W uint16) {
	getProc("ImFontAtlasRect_setW").Call(uintptr(unsafe.Pointer(Self)), uintptr(W))
}

func (i *Imgui) ImFontAtlasRectH(Self *ImFontAtlasRect) uint16 {
	r1, _, _ := getProc("ImFontAtlasRect_h").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontAtlasRectSetH(Self *ImFontAtlasRect, H uint16) {
	getProc("ImFontAtlasRect_setH").Call(uintptr(unsafe.Pointer(Self)), uintptr(H))
}

func (i *Imgui) ImFontAtlasRectUv0(Self *ImFontAtlasRect) *ImVec2 {
	r1, _, _ := getProc("ImFontAtlasRect_uv0").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasRectSetUv0(Self *ImFontAtlasRect, Uv0 *ImVec2) {
	getProc("ImFontAtlasRect_setUv0").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Uv0)))
}

func (i *Imgui) ImFontAtlasRectUv1(Self *ImFontAtlasRect) *ImVec2 {
	r1, _, _ := getProc("ImFontAtlasRect_uv1").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasRectSetUv1(Self *ImFontAtlasRect, Uv1 *ImVec2) {
	getProc("ImFontAtlasRect_setUv1").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Uv1)))
}

func (i *Imgui) ImFontAtlasRectOperatorAssign(Self *ImFontAtlasRect, Param1 *ImFontAtlasRect) {
	getProc("ImFontAtlasRect_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImFontAtlasRectDelete(Self *ImFontAtlasRect) {
	getProc("ImFontAtlasRect_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasNew() *ImFontAtlas {
	r1, _, _ := getProc("ImFontAtlas_new").Call()
	return (*ImFontAtlas)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasNew2(Param1 *ImFontAtlas) *ImFontAtlas {
	r1, _, _ := getProc("ImFontAtlas_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImFontAtlas)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFont(Self *ImFontAtlas, Font_cfg *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFont").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font_cfg)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontDefault(Self *ImFontAtlas) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontDefault").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontDefaultVector(Self *ImFontAtlas) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontDefaultVector").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontDefaultBitmap(Self *ImFontAtlas) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontDefaultBitmap").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromFileTTF(Self *ImFontAtlas, Filename *int8) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromFileTTF").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Filename)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromMemoryCompressedBase85TTF(Self *ImFontAtlas, Compressed_font_data_base85 *int8) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromMemoryCompressedBase85TTF").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Compressed_font_data_base85)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasRemoveFont(Self *ImFontAtlas, Font *ImFont) {
	getProc("ImFontAtlas_RemoveFont").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)))
}

func (i *Imgui) ImFontAtlasClear(Self *ImFontAtlas) {
	getProc("ImFontAtlas_Clear").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasClearFonts(Self *ImFontAtlas) {
	getProc("ImFontAtlas_ClearFonts").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasCompactCache(Self *ImFontAtlas) {
	getProc("ImFontAtlas_CompactCache").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasSetFontLoader(Self *ImFontAtlas, Font_loader *ImFontLoader) {
	getProc("ImFontAtlas_SetFontLoader").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font_loader)))
}

func (i *Imgui) ImFontAtlasClearInputData(Self *ImFontAtlas) {
	getProc("ImFontAtlas_ClearInputData").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasClearTexData(Self *ImFontAtlas) {
	getProc("ImFontAtlas_ClearTexData").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAtlasBuild(Self *ImFontAtlas) bool {
	r1, _, _ := getProc("ImFontAtlas_Build").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasGetTexDataAsAlpha8(Self *ImFontAtlas, Out_pixels **uint8, Out_width *int32, Out_height *int32) {
	getProc("ImFontAtlas_GetTexDataAsAlpha8").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Out_pixels)), uintptr(unsafe.Pointer(Out_width)), uintptr(unsafe.Pointer(Out_height)))
}

func (i *Imgui) ImFontAtlasGetTexDataAsRGBA32(Self *ImFontAtlas, Out_pixels **uint8, Out_width *int32, Out_height *int32) {
	getProc("ImFontAtlas_GetTexDataAsRGBA32").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Out_pixels)), uintptr(unsafe.Pointer(Out_width)), uintptr(unsafe.Pointer(Out_height)))
}

func (i *Imgui) ImFontAtlasSetTexID(Self *ImFontAtlas, Id uint64) {
	getProc("ImFontAtlas_SetTexID").Call(uintptr(unsafe.Pointer(Self)), *(*uintptr)(unsafe.Pointer(&Id)))
}

func (i *Imgui) ImFontAtlasSetTexIDWithId(Self *ImFontAtlas, Id *ImTextureRef) {
	getProc("ImFontAtlas_SetTexIDWithId").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Id)))
}

func (i *Imgui) ImFontAtlasIsBuilt(Self *ImFontAtlas) bool {
	r1, _, _ := getProc("ImFontAtlas_IsBuilt").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasGetGlyphRangesDefault(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesDefault").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesGreek(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesGreek").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesKorean(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesKorean").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesJapanese(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesJapanese").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesChineseFull(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesChineseFull").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesChineseSimplifiedCommon(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesCyrillic(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesCyrillic").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesThai(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesThai").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetGlyphRangesVietnamese(Self *ImFontAtlas) *uint16 {
	r1, _, _ := getProc("ImFontAtlas_GetGlyphRangesVietnamese").Call(uintptr(unsafe.Pointer(Self)))
	return (*uint16)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddCustomRect(Self *ImFontAtlas, Width int32, Height int32) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(Width), uintptr(Height))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasRemoveCustomRect(Self *ImFontAtlas, Id int32) {
	getProc("ImFontAtlas_RemoveCustomRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(Id))
}

func (i *Imgui) ImFontAtlasGetCustomRect(Self *ImFontAtlas, Id int32, Out_r *ImFontAtlasRect) bool {
	r1, _, _ := getProc("ImFontAtlas_GetCustomRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(Id), uintptr(unsafe.Pointer(Out_r)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasFlags(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetFlags(Self *ImFontAtlas, Flags int32) {
	getProc("ImFontAtlas_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImFontAtlasTexDesiredFormat(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexDesiredFormat").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexDesiredFormat(Self *ImFontAtlas, TexDesiredFormat int32) {
	getProc("ImFontAtlas_setTexDesiredFormat").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexDesiredFormat))
}

func (i *Imgui) ImFontAtlasTexGlyphPadding(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexGlyphPadding").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexGlyphPadding(Self *ImFontAtlas, TexGlyphPadding int32) {
	getProc("ImFontAtlas_setTexGlyphPadding").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexGlyphPadding))
}

func (i *Imgui) ImFontAtlasTexMinWidth(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexMinWidth").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexMinWidth(Self *ImFontAtlas, TexMinWidth int32) {
	getProc("ImFontAtlas_setTexMinWidth").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexMinWidth))
}

func (i *Imgui) ImFontAtlasTexMinHeight(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexMinHeight").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexMinHeight(Self *ImFontAtlas, TexMinHeight int32) {
	getProc("ImFontAtlas_setTexMinHeight").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexMinHeight))
}

func (i *Imgui) ImFontAtlasTexMaxWidth(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexMaxWidth").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexMaxWidth(Self *ImFontAtlas, TexMaxWidth int32) {
	getProc("ImFontAtlas_setTexMaxWidth").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexMaxWidth))
}

func (i *Imgui) ImFontAtlasTexMaxHeight(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexMaxHeight").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexMaxHeight(Self *ImFontAtlas, TexMaxHeight int32) {
	getProc("ImFontAtlas_setTexMaxHeight").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexMaxHeight))
}

func (i *Imgui) ImFontAtlasTexData(Self *ImFontAtlas) *ImTextureData {
	r1, _, _ := getProc("ImFontAtlas_TexData").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImTextureData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetTexData(Self *ImFontAtlas, TexData *ImTextureData) {
	getProc("ImFontAtlas_setTexData").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexData)))
}

func (i *Imgui) ImFontAtlasTexList(Self *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("ImFontAtlas_TexList").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontAtlasSetTexList(Self *ImFontAtlas, TexList unsafe.Pointer) {
	getProc("ImFontAtlas_setTexList").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexList))
}

func (i *Imgui) ImFontAtlasLocked(Self *ImFontAtlas) bool {
	r1, _, _ := getProc("ImFontAtlas_Locked").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasSetLocked(Self *ImFontAtlas, Locked bool) {
	getProc("ImFontAtlas_setLocked").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if Locked {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontAtlasRendererHasTextures(Self *ImFontAtlas) bool {
	r1, _, _ := getProc("ImFontAtlas_RendererHasTextures").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasSetRendererHasTextures(Self *ImFontAtlas, RendererHasTextures bool) {
	getProc("ImFontAtlas_setRendererHasTextures").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if RendererHasTextures {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontAtlasTexIsBuilt(Self *ImFontAtlas) bool {
	r1, _, _ := getProc("ImFontAtlas_TexIsBuilt").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasSetTexIsBuilt(Self *ImFontAtlas, TexIsBuilt bool) {
	getProc("ImFontAtlas_setTexIsBuilt").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if TexIsBuilt {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontAtlasTexPixelsUseColors(Self *ImFontAtlas) bool {
	r1, _, _ := getProc("ImFontAtlas_TexPixelsUseColors").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontAtlasSetTexPixelsUseColors(Self *ImFontAtlas, TexPixelsUseColors bool) {
	getProc("ImFontAtlas_setTexPixelsUseColors").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if TexPixelsUseColors {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontAtlasTexUvScale(Self *ImFontAtlas) *ImVec2 {
	r1, _, _ := getProc("ImFontAtlas_TexUvScale").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetTexUvScale(Self *ImFontAtlas, TexUvScale *ImVec2) {
	getProc("ImFontAtlas_setTexUvScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexUvScale)))
}

func (i *Imgui) ImFontAtlasTexUvWhitePixel(Self *ImFontAtlas) *ImVec2 {
	r1, _, _ := getProc("ImFontAtlas_TexUvWhitePixel").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetTexUvWhitePixel(Self *ImFontAtlas, TexUvWhitePixel *ImVec2) {
	getProc("ImFontAtlas_setTexUvWhitePixel").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexUvWhitePixel)))
}

func (i *Imgui) ImFontAtlasFonts(Self *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("ImFontAtlas_Fonts").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontAtlasSetFonts(Self *ImFontAtlas, Fonts unsafe.Pointer) {
	getProc("ImFontAtlas_setFonts").Call(uintptr(unsafe.Pointer(Self)), uintptr(Fonts))
}

func (i *Imgui) ImFontAtlasSources(Self *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("ImFontAtlas_Sources").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontAtlasSetSources(Self *ImFontAtlas, Sources unsafe.Pointer) {
	getProc("ImFontAtlas_setSources").Call(uintptr(unsafe.Pointer(Self)), uintptr(Sources))
}

func (i *Imgui) ImFontAtlasTexUvLines(Self *ImFontAtlas) *ImVec4 {
	r1, _, _ := getProc("ImFontAtlas_TexUvLines").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetTexUvLines(Self *ImFontAtlas, TexUvLines *ImVec4) {
	getProc("ImFontAtlas_setTexUvLines").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TexUvLines)))
}

func (i *Imgui) ImFontAtlasTexNextUniqueID(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_TexNextUniqueID").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetTexNextUniqueID(Self *ImFontAtlas, TexNextUniqueID int32) {
	getProc("ImFontAtlas_setTexNextUniqueID").Call(uintptr(unsafe.Pointer(Self)), uintptr(TexNextUniqueID))
}

func (i *Imgui) ImFontAtlasFontNextUniqueID(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_FontNextUniqueID").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetFontNextUniqueID(Self *ImFontAtlas, FontNextUniqueID int32) {
	getProc("ImFontAtlas_setFontNextUniqueID").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontNextUniqueID))
}

func (i *Imgui) ImFontAtlasDrawListSharedDatas(Self *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("ImFontAtlas_DrawListSharedDatas").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontAtlasSetDrawListSharedDatas(Self *ImFontAtlas, DrawListSharedDatas unsafe.Pointer) {
	getProc("ImFontAtlas_setDrawListSharedDatas").Call(uintptr(unsafe.Pointer(Self)), uintptr(DrawListSharedDatas))
}

func (i *Imgui) ImFontAtlasBuilder(Self *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("ImFontAtlas_Builder").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontAtlasSetBuilder(Self *ImFontAtlas, Builder unsafe.Pointer) {
	getProc("ImFontAtlas_setBuilder").Call(uintptr(unsafe.Pointer(Self)), uintptr(Builder))
}

func (i *Imgui) ImFontAtlasFontLoader(Self *ImFontAtlas) *ImFontLoader {
	r1, _, _ := getProc("ImFontAtlas_FontLoader").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFontLoader)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetFontLoader_1(Self *ImFontAtlas, FontLoader *ImFontLoader) {
	getProc("ImFontAtlas_setFontLoader").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FontLoader)))
}

func (i *Imgui) ImFontAtlasFontLoaderName(Self *ImFontAtlas) *int8 {
	r1, _, _ := getProc("ImFontAtlas_FontLoaderName").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetFontLoaderName(Self *ImFontAtlas, FontLoaderName *int8) {
	getProc("ImFontAtlas_setFontLoaderName").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FontLoaderName)))
}

func (i *Imgui) ImFontAtlasFontLoaderFlags(Self *ImFontAtlas) uint32 {
	r1, _, _ := getProc("ImFontAtlas_FontLoaderFlags").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontAtlasSetFontLoaderFlags(Self *ImFontAtlas, FontLoaderFlags uint32) {
	getProc("ImFontAtlas_setFontLoaderFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontLoaderFlags))
}

func (i *Imgui) ImFontAtlasRefCount(Self *ImFontAtlas) int32 {
	r1, _, _ := getProc("ImFontAtlas_RefCount").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasSetRefCount(Self *ImFontAtlas, RefCount int32) {
	getProc("ImFontAtlas_setRefCount").Call(uintptr(unsafe.Pointer(Self)), uintptr(RefCount))
}

func (i *Imgui) ImFontAtlasOwnerContext(Self *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("ImFontAtlas_OwnerContext").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontAtlasSetOwnerContext(Self *ImFontAtlas, OwnerContext unsafe.Pointer) {
	getProc("ImFontAtlas_setOwnerContext").Call(uintptr(unsafe.Pointer(Self)), uintptr(OwnerContext))
}

func (i *Imgui) ImFontAtlasTempRect(Self *ImFontAtlas) *ImFontAtlasRect {
	r1, _, _ := getProc("ImFontAtlas_TempRect").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFontAtlasRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasSetTempRect(Self *ImFontAtlas, TempRect *ImFontAtlasRect) {
	getProc("ImFontAtlas_setTempRect").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(TempRect)))
}

func (i *Imgui) ImFontAtlasAddCustomRectRegular(Self *ImFontAtlas, W int32, H int32) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRectRegular").Call(uintptr(unsafe.Pointer(Self)), uintptr(W), uintptr(H))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasGetCustomRectByIndex(Self *ImFontAtlas, Id int32) *ImFontAtlasRect {
	r1, _, _ := getProc("ImFontAtlas_GetCustomRectByIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(Id))
	return (*ImFontAtlasRect)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasCalcCustomRectUV(Self *ImFontAtlas, R *ImFontAtlasRect, Out_uv_min *ImVec2, Out_uv_max *ImVec2) {
	getProc("ImFontAtlas_CalcCustomRectUV").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(R)), uintptr(unsafe.Pointer(Out_uv_min)), uintptr(unsafe.Pointer(Out_uv_max)))
}

func (i *Imgui) ImFontAtlasAddCustomRectFontGlyph(Self *ImFontAtlas, Font *ImFont, Codepoint uint16, W int32, H int32, Advance_x float32) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRectFontGlyph").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Codepoint), uintptr(W), uintptr(H), uintptr(Advance_x))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasAddCustomRectFontGlyphForSize(Self *ImFontAtlas, Font *ImFont, Font_size float32, Codepoint uint16, W int32, H int32, Advance_x float32) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRectFontGlyphForSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Font_size), uintptr(Codepoint), uintptr(W), uintptr(H), uintptr(Advance_x))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasOperatorAssign(Self *ImFontAtlas, Param1 *ImFontAtlas) {
	getProc("ImFontAtlas_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImFontAtlasAddFontDefaultWithFontCfg(Self *ImFontAtlas, Font_cfg *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontDefaultWithFontCfg").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font_cfg)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontDefaultVectorWithFontCfg(Self *ImFontAtlas, Font_cfg *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontDefaultVectorWithFontCfg").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font_cfg)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontDefaultBitmapWithFontCfg(Self *ImFontAtlas, Font_cfg *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontDefaultBitmapWithFontCfg").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font_cfg)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromFileTTF2(Self *ImFontAtlas, Filename *int8, Size_pixels float32) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromFileTTF2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Filename)), uintptr(Size_pixels))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromFileTTF3(Self *ImFontAtlas, Filename *int8, Size_pixels float32, Font_cfg *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromFileTTF3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Filename)), uintptr(Size_pixels), uintptr(unsafe.Pointer(Font_cfg)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromFileTTF4(Self *ImFontAtlas, Filename *int8, Size_pixels float32, Font_cfg *ImFontConfig, Glyph_ranges *uint16) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromFileTTF4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Filename)), uintptr(Size_pixels), uintptr(unsafe.Pointer(Font_cfg)), uintptr(unsafe.Pointer(Glyph_ranges)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromMemoryCompressedBase85TTF2(Self *ImFontAtlas, Compressed_font_data_base85 *int8, Size_pixels float32) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromMemoryCompressedBase85TTF2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Compressed_font_data_base85)), uintptr(Size_pixels))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromMemoryCompressedBase85TTF3(Self *ImFontAtlas, Compressed_font_data_base85 *int8, Size_pixels float32, Font_cfg *ImFontConfig) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromMemoryCompressedBase85TTF3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Compressed_font_data_base85)), uintptr(Size_pixels), uintptr(unsafe.Pointer(Font_cfg)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasAddFontFromMemoryCompressedBase85TTF4(Self *ImFontAtlas, Compressed_font_data_base85 *int8, Size_pixels float32, Font_cfg *ImFontConfig, Glyph_ranges *uint16) *ImFont {
	r1, _, _ := getProc("ImFontAtlas_AddFontFromMemoryCompressedBase85TTF4").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Compressed_font_data_base85)), uintptr(Size_pixels), uintptr(unsafe.Pointer(Font_cfg)), uintptr(unsafe.Pointer(Glyph_ranges)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontAtlasGetTexDataAsAlpha82(Self *ImFontAtlas, Out_pixels **uint8, Out_width *int32, Out_height *int32, Out_bytes_per_pixel *int32) {
	getProc("ImFontAtlas_GetTexDataAsAlpha82").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Out_pixels)), uintptr(unsafe.Pointer(Out_width)), uintptr(unsafe.Pointer(Out_height)), uintptr(unsafe.Pointer(Out_bytes_per_pixel)))
}

func (i *Imgui) ImFontAtlasGetTexDataAsRGBA322(Self *ImFontAtlas, Out_pixels **uint8, Out_width *int32, Out_height *int32, Out_bytes_per_pixel *int32) {
	getProc("ImFontAtlas_GetTexDataAsRGBA322").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Out_pixels)), uintptr(unsafe.Pointer(Out_width)), uintptr(unsafe.Pointer(Out_height)), uintptr(unsafe.Pointer(Out_bytes_per_pixel)))
}

func (i *Imgui) ImFontAtlasAddCustomRect2(Self *ImFontAtlas, Width int32, Height int32, Out_r *ImFontAtlasRect) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRect2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Width), uintptr(Height), uintptr(unsafe.Pointer(Out_r)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasAddCustomRectFontGlyph2(Self *ImFontAtlas, Font *ImFont, Codepoint uint16, W int32, H int32, Advance_x float32, Offset *ImVec2) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRectFontGlyph2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Codepoint), uintptr(W), uintptr(H), uintptr(Advance_x), uintptr(unsafe.Pointer(Offset)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasAddCustomRectFontGlyphForSize2(Self *ImFontAtlas, Font *ImFont, Font_size float32, Codepoint uint16, W int32, H int32, Advance_x float32, Offset *ImVec2) int32 {
	r1, _, _ := getProc("ImFontAtlas_AddCustomRectFontGlyphForSize2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Font)), uintptr(Font_size), uintptr(Codepoint), uintptr(W), uintptr(H), uintptr(Advance_x), uintptr(unsafe.Pointer(Offset)))
	return int32(r1)
}

func (i *Imgui) ImFontAtlasDelete(Self *ImFontAtlas) {
	getProc("ImFontAtlas_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontBakedNew() *ImFontBaked {
	r1, _, _ := getProc("ImFontBaked_new").Call()
	return (*ImFontBaked)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontBakedNew2(Param1 *ImFontBaked) *ImFontBaked {
	r1, _, _ := getProc("ImFontBaked_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImFontBaked)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontBakedIndexAdvanceX(Self *ImFontBaked) unsafe.Pointer {
	r1, _, _ := getProc("ImFontBaked_IndexAdvanceX").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontBakedSetIndexAdvanceX(Self *ImFontBaked, IndexAdvanceX unsafe.Pointer) {
	getProc("ImFontBaked_setIndexAdvanceX").Call(uintptr(unsafe.Pointer(Self)), uintptr(IndexAdvanceX))
}

func (i *Imgui) ImFontBakedFallbackAdvanceX(Self *ImFontBaked) float32 {
	r1, _, _ := getProc("ImFontBaked_FallbackAdvanceX").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontBakedSetFallbackAdvanceX(Self *ImFontBaked, FallbackAdvanceX float32) {
	getProc("ImFontBaked_setFallbackAdvanceX").Call(uintptr(unsafe.Pointer(Self)), uintptr(FallbackAdvanceX))
}

func (i *Imgui) ImFontBakedSize(Self *ImFontBaked) float32 {
	r1, _, _ := getProc("ImFontBaked_Size").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontBakedSetSize(Self *ImFontBaked, Size float32) {
	getProc("ImFontBaked_setSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size))
}

func (i *Imgui) ImFontBakedRasterizerDensity(Self *ImFontBaked) float32 {
	r1, _, _ := getProc("ImFontBaked_RasterizerDensity").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontBakedSetRasterizerDensity(Self *ImFontBaked, RasterizerDensity float32) {
	getProc("ImFontBaked_setRasterizerDensity").Call(uintptr(unsafe.Pointer(Self)), uintptr(RasterizerDensity))
}

func (i *Imgui) ImFontBakedIndexLookup(Self *ImFontBaked) unsafe.Pointer {
	r1, _, _ := getProc("ImFontBaked_IndexLookup").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontBakedSetIndexLookup(Self *ImFontBaked, IndexLookup unsafe.Pointer) {
	getProc("ImFontBaked_setIndexLookup").Call(uintptr(unsafe.Pointer(Self)), uintptr(IndexLookup))
}

func (i *Imgui) ImFontBakedGlyphs(Self *ImFontBaked) unsafe.Pointer {
	r1, _, _ := getProc("ImFontBaked_Glyphs").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontBakedSetGlyphs(Self *ImFontBaked, Glyphs unsafe.Pointer) {
	getProc("ImFontBaked_setGlyphs").Call(uintptr(unsafe.Pointer(Self)), uintptr(Glyphs))
}

func (i *Imgui) ImFontBakedFallbackGlyphIndex(Self *ImFontBaked) int32 {
	r1, _, _ := getProc("ImFontBaked_FallbackGlyphIndex").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontBakedSetFallbackGlyphIndex(Self *ImFontBaked, FallbackGlyphIndex int32) {
	getProc("ImFontBaked_setFallbackGlyphIndex").Call(uintptr(unsafe.Pointer(Self)), uintptr(FallbackGlyphIndex))
}

func (i *Imgui) ImFontBakedAscent(Self *ImFontBaked) float32 {
	r1, _, _ := getProc("ImFontBaked_Ascent").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontBakedSetAscent(Self *ImFontBaked, Ascent float32) {
	getProc("ImFontBaked_setAscent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Ascent))
}

func (i *Imgui) ImFontBakedDescent(Self *ImFontBaked) float32 {
	r1, _, _ := getProc("ImFontBaked_Descent").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontBakedSetDescent(Self *ImFontBaked, Descent float32) {
	getProc("ImFontBaked_setDescent").Call(uintptr(unsafe.Pointer(Self)), uintptr(Descent))
}

func (i *Imgui) ImFontBakedMetricsTotalSurface(Self *ImFontBaked) uint32 {
	r1, _, _ := getProc("ImFontBaked_MetricsTotalSurface").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontBakedSetMetricsTotalSurface(Self *ImFontBaked, MetricsTotalSurface uint32) {
	getProc("ImFontBaked_setMetricsTotalSurface").Call(uintptr(unsafe.Pointer(Self)), uintptr(MetricsTotalSurface))
}

func (i *Imgui) ImFontBakedWantDestroy(Self *ImFontBaked) uint32 {
	r1, _, _ := getProc("ImFontBaked_WantDestroy").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontBakedSetWantDestroy(Self *ImFontBaked, WantDestroy uint32) {
	getProc("ImFontBaked_setWantDestroy").Call(uintptr(unsafe.Pointer(Self)), uintptr(WantDestroy))
}

func (i *Imgui) ImFontBakedLoadNoFallback(Self *ImFontBaked) uint32 {
	r1, _, _ := getProc("ImFontBaked_LoadNoFallback").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontBakedSetLoadNoFallback(Self *ImFontBaked, LoadNoFallback uint32) {
	getProc("ImFontBaked_setLoadNoFallback").Call(uintptr(unsafe.Pointer(Self)), uintptr(LoadNoFallback))
}

func (i *Imgui) ImFontBakedLoadNoRenderOnLayout(Self *ImFontBaked) uint32 {
	r1, _, _ := getProc("ImFontBaked_LoadNoRenderOnLayout").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontBakedSetLoadNoRenderOnLayout(Self *ImFontBaked, LoadNoRenderOnLayout uint32) {
	getProc("ImFontBaked_setLoadNoRenderOnLayout").Call(uintptr(unsafe.Pointer(Self)), uintptr(LoadNoRenderOnLayout))
}

func (i *Imgui) ImFontBakedLastUsedFrame(Self *ImFontBaked) int32 {
	r1, _, _ := getProc("ImFontBaked_LastUsedFrame").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontBakedSetLastUsedFrame(Self *ImFontBaked, LastUsedFrame int32) {
	getProc("ImFontBaked_setLastUsedFrame").Call(uintptr(unsafe.Pointer(Self)), uintptr(LastUsedFrame))
}

func (i *Imgui) ImFontBakedBakedId(Self *ImFontBaked) uint32 {
	r1, _, _ := getProc("ImFontBaked_BakedId").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontBakedSetBakedId(Self *ImFontBaked, BakedId uint32) {
	getProc("ImFontBaked_setBakedId").Call(uintptr(unsafe.Pointer(Self)), uintptr(BakedId))
}

func (i *Imgui) ImFontBakedOwnerFont(Self *ImFontBaked) *ImFont {
	r1, _, _ := getProc("ImFontBaked_OwnerFont").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontBakedSetOwnerFont(Self *ImFontBaked, OwnerFont *ImFont) {
	getProc("ImFontBaked_setOwnerFont").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(OwnerFont)))
}

func (i *Imgui) ImFontBakedClearOutputData(Self *ImFontBaked) {
	getProc("ImFontBaked_ClearOutputData").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontBakedFindGlyph(Self *ImFontBaked, C uint16) *ImFontGlyph {
	r1, _, _ := getProc("ImFontBaked_FindGlyph").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
	return (*ImFontGlyph)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontBakedFindGlyphNoFallback(Self *ImFontBaked, C uint16) *ImFontGlyph {
	r1, _, _ := getProc("ImFontBaked_FindGlyphNoFallback").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
	return (*ImFontGlyph)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontBakedGetCharAdvance(Self *ImFontBaked, C uint16) float32 {
	r1, _, _ := getProc("ImFontBaked_GetCharAdvance").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontBakedIsGlyphLoaded(Self *ImFontBaked, C uint16) bool {
	r1, _, _ := getProc("ImFontBaked_IsGlyphLoaded").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
	return r1 != 0
}

func (i *Imgui) ImFontBakedOperatorAssign(Self *ImFontBaked, Param1 *ImFontBaked) {
	getProc("ImFontBaked_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImFontBakedDelete(Self *ImFontBaked) {
	getProc("ImFontBaked_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontNew() *ImFont {
	r1, _, _ := getProc("ImFont_new").Call()
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontNew2(Param1 *ImFont) *ImFont {
	r1, _, _ := getProc("ImFont_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontLastBaked(Self *ImFont) *ImFontBaked {
	r1, _, _ := getProc("ImFont_LastBaked").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFontBaked)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontSetLastBaked(Self *ImFont, LastBaked *ImFontBaked) {
	getProc("ImFont_setLastBaked").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(LastBaked)))
}

func (i *Imgui) ImFontOwnerAtlas(Self *ImFont) *ImFontAtlas {
	r1, _, _ := getProc("ImFont_OwnerAtlas").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImFontAtlas)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontSetOwnerAtlas(Self *ImFont, OwnerAtlas *ImFontAtlas) {
	getProc("ImFont_setOwnerAtlas").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(OwnerAtlas)))
}

func (i *Imgui) ImFontFlags(Self *ImFont) int32 {
	r1, _, _ := getProc("ImFont_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImFontSetFlags(Self *ImFont, Flags int32) {
	getProc("ImFont_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImFontCurrentRasterizerDensity(Self *ImFont) float32 {
	r1, _, _ := getProc("ImFont_CurrentRasterizerDensity").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontSetCurrentRasterizerDensity(Self *ImFont, CurrentRasterizerDensity float32) {
	getProc("ImFont_setCurrentRasterizerDensity").Call(uintptr(unsafe.Pointer(Self)), uintptr(CurrentRasterizerDensity))
}

func (i *Imgui) ImFontFontId(Self *ImFont) uint32 {
	r1, _, _ := getProc("ImFont_FontId").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImFontSetFontId(Self *ImFont, FontId uint32) {
	getProc("ImFont_setFontId").Call(uintptr(unsafe.Pointer(Self)), uintptr(FontId))
}

func (i *Imgui) ImFontLegacySize(Self *ImFont) float32 {
	r1, _, _ := getProc("ImFont_LegacySize").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontSetLegacySize(Self *ImFont, LegacySize float32) {
	getProc("ImFont_setLegacySize").Call(uintptr(unsafe.Pointer(Self)), uintptr(LegacySize))
}

func (i *Imgui) ImFontSources(Self *ImFont) unsafe.Pointer {
	r1, _, _ := getProc("ImFont_Sources").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImFontSetSources(Self *ImFont, Sources unsafe.Pointer) {
	getProc("ImFont_setSources").Call(uintptr(unsafe.Pointer(Self)), uintptr(Sources))
}

func (i *Imgui) ImFontEllipsisChar(Self *ImFont) uint16 {
	r1, _, _ := getProc("ImFont_EllipsisChar").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontSetEllipsisChar(Self *ImFont, EllipsisChar uint16) {
	getProc("ImFont_setEllipsisChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(EllipsisChar))
}

func (i *Imgui) ImFontFallbackChar(Self *ImFont) uint16 {
	r1, _, _ := getProc("ImFont_FallbackChar").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImFontSetFallbackChar(Self *ImFont, FallbackChar uint16) {
	getProc("ImFont_setFallbackChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(FallbackChar))
}

func (i *Imgui) ImFontUsed8kPagesMap(Self *ImFont) *ImU8 {
	r1, _, _ := getProc("ImFont_Used8kPagesMap").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImU8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontSetUsed8kPagesMap(Self *ImFont, Used8kPagesMap *ImU8) {
	getProc("ImFont_setUsed8kPagesMap").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Used8kPagesMap)))
}

func (i *Imgui) ImFontEllipsisAutoBake(Self *ImFont) bool {
	r1, _, _ := getProc("ImFont_EllipsisAutoBake").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontSetEllipsisAutoBake(Self *ImFont, EllipsisAutoBake bool) {
	getProc("ImFont_setEllipsisAutoBake").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if EllipsisAutoBake {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImFontRemapPairs(Self *ImFont) *ImGuiStorage {
	r1, _, _ := getProc("ImFont_RemapPairs").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImGuiStorage)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontSetRemapPairs(Self *ImFont, RemapPairs *ImGuiStorage) {
	getProc("ImFont_setRemapPairs").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(RemapPairs)))
}

func (i *Imgui) ImFontScale(Self *ImFont) float32 {
	r1, _, _ := getProc("ImFont_Scale").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImFontSetScale(Self *ImFont, Scale float32) {
	getProc("ImFont_setScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(Scale))
}

func (i *Imgui) ImFontIsGlyphInFont(Self *ImFont, C uint16) bool {
	r1, _, _ := getProc("ImFont_IsGlyphInFont").Call(uintptr(unsafe.Pointer(Self)), uintptr(C))
	return r1 != 0
}

func (i *Imgui) ImFontIsLoaded(Self *ImFont) bool {
	r1, _, _ := getProc("ImFont_IsLoaded").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImFontGetDebugName(Self *ImFont) *int8 {
	r1, _, _ := getProc("ImFont_GetDebugName").Call(uintptr(unsafe.Pointer(Self)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontGetFontBaked(Self *ImFont, Font_size float32) *ImFontBaked {
	r1, _, _ := getProc("ImFont_GetFontBaked").Call(uintptr(unsafe.Pointer(Self)), uintptr(Font_size))
	return (*ImFontBaked)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontCalcTextSizeA(Self *ImFont, Size float32, Max_width float32, Wrap_width float32, Text_begin *int8) *ImVec2 {
	r1, _, _ := getProc("ImFont_CalcTextSizeA").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size), uintptr(Max_width), uintptr(Wrap_width), uintptr(unsafe.Pointer(Text_begin)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontCalcWordWrapPosition(Self *ImFont, Size float32, Text *int8, Text_end *int8, Wrap_width float32) *int8 {
	r1, _, _ := getProc("ImFont_CalcWordWrapPosition").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size), uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)), uintptr(Wrap_width))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontRenderChar(Self *ImFont, Draw_list *ImDrawList, Size float32, Pos *ImVec2, Col uint32, C uint16) {
	getProc("ImFont_RenderChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(C))
}

func (i *Imgui) ImFontRenderText(Self *ImFont, Draw_list *ImDrawList, Size float32, Pos *ImVec2, Col uint32, Clip_rect *ImVec4, Text_begin *int8, Text_end *int8) {
	getProc("ImFont_RenderText").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Clip_rect)), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)))
}

func (i *Imgui) ImFontCalcWordWrapPositionA(Self *ImFont, Scale float32, Text *int8, Text_end *int8, Wrap_width float32) *int8 {
	r1, _, _ := getProc("ImFont_CalcWordWrapPositionA").Call(uintptr(unsafe.Pointer(Self)), uintptr(Scale), uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)), uintptr(Wrap_width))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontClearOutputData(Self *ImFont) {
	getProc("ImFont_ClearOutputData").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImFontAddRemapChar(Self *ImFont, From_codepoint uint16, To_codepoint uint16) {
	getProc("ImFont_AddRemapChar").Call(uintptr(unsafe.Pointer(Self)), uintptr(From_codepoint), uintptr(To_codepoint))
}

func (i *Imgui) ImFontIsGlyphRangeUnused(Self *ImFont, C_begin uint32, C_last uint32) bool {
	r1, _, _ := getProc("ImFont_IsGlyphRangeUnused").Call(uintptr(unsafe.Pointer(Self)), uintptr(C_begin), uintptr(C_last))
	return r1 != 0
}

func (i *Imgui) ImFontOperatorAssign(Self *ImFont, Param1 *ImFont) {
	getProc("ImFont_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImFontGetFontBaked2(Self *ImFont, Font_size float32, Density float32) *ImFontBaked {
	r1, _, _ := getProc("ImFont_GetFontBaked2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Font_size), uintptr(Density))
	return (*ImFontBaked)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontCalcTextSizeA2(Self *ImFont, Size float32, Max_width float32, Wrap_width float32, Text_begin *int8, Text_end *int8) *ImVec2 {
	r1, _, _ := getProc("ImFont_CalcTextSizeA2").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size), uintptr(Max_width), uintptr(Wrap_width), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontCalcTextSizeA3(Self *ImFont, Size float32, Max_width float32, Wrap_width float32, Text_begin *int8, Text_end *int8, Out_remaining **int8) *ImVec2 {
	r1, _, _ := getProc("ImFont_CalcTextSizeA3").Call(uintptr(unsafe.Pointer(Self)), uintptr(Size), uintptr(Max_width), uintptr(Wrap_width), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)), uintptr(unsafe.Pointer(Out_remaining)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImFontRenderChar2(Self *ImFont, Draw_list *ImDrawList, Size float32, Pos *ImVec2, Col uint32, C uint16, Cpu_fine_clip *ImVec4) {
	getProc("ImFont_RenderChar2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(C), uintptr(unsafe.Pointer(Cpu_fine_clip)))
}

func (i *Imgui) ImFontRenderText2(Self *ImFont, Draw_list *ImDrawList, Size float32, Pos *ImVec2, Col uint32, Clip_rect *ImVec4, Text_begin *int8, Text_end *int8, Wrap_width float32) {
	getProc("ImFont_RenderText2").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Clip_rect)), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)), uintptr(Wrap_width))
}

func (i *Imgui) ImFontRenderText3(Self *ImFont, Draw_list *ImDrawList, Size float32, Pos *ImVec2, Col uint32, Clip_rect *ImVec4, Text_begin *int8, Text_end *int8, Wrap_width float32, Flags int32) {
	getProc("ImFont_RenderText3").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Draw_list)), uintptr(Size), uintptr(unsafe.Pointer(Pos)), uintptr(Col), uintptr(unsafe.Pointer(Clip_rect)), uintptr(unsafe.Pointer(Text_begin)), uintptr(unsafe.Pointer(Text_end)), uintptr(Wrap_width), uintptr(Flags))
}

func (i *Imgui) ImFontDelete(Self *ImFont) {
	getProc("ImFont_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiViewportNew() *ImGuiViewport {
	r1, _, _ := getProc("ImGuiViewport_new").Call()
	return (*ImGuiViewport)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportID(Self *ImGuiViewport) uint32 {
	r1, _, _ := getProc("ImGuiViewport_ID").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiViewportSetID(Self *ImGuiViewport, ID uint32) {
	getProc("ImGuiViewport_setID").Call(uintptr(unsafe.Pointer(Self)), uintptr(ID))
}

func (i *Imgui) ImGuiViewportFlags(Self *ImGuiViewport) int32 {
	r1, _, _ := getProc("ImGuiViewport_Flags").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiViewportSetFlags(Self *ImGuiViewport, Flags int32) {
	getProc("ImGuiViewport_setFlags").Call(uintptr(unsafe.Pointer(Self)), uintptr(Flags))
}

func (i *Imgui) ImGuiViewportPos(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_Pos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportSetPos(Self *ImGuiViewport, Pos *ImVec2) {
	getProc("ImGuiViewport_setPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Pos)))
}

func (i *Imgui) ImGuiViewportSize(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_Size").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportSetSize(Self *ImGuiViewport, Size *ImVec2) {
	getProc("ImGuiViewport_setSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Size)))
}

func (i *Imgui) ImGuiViewportFramebufferScale(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_FramebufferScale").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportSetFramebufferScale(Self *ImGuiViewport, FramebufferScale *ImVec2) {
	getProc("ImGuiViewport_setFramebufferScale").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(FramebufferScale)))
}

func (i *Imgui) ImGuiViewportWorkPos(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_WorkPos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportSetWorkPos(Self *ImGuiViewport, WorkPos *ImVec2) {
	getProc("ImGuiViewport_setWorkPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(WorkPos)))
}

func (i *Imgui) ImGuiViewportWorkSize(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_WorkSize").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportSetWorkSize(Self *ImGuiViewport, WorkSize *ImVec2) {
	getProc("ImGuiViewport_setWorkSize").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(WorkSize)))
}

func (i *Imgui) ImGuiViewportGetCenter(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_GetCenter").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportGetWorkCenter(Self *ImGuiViewport) *ImVec2 {
	r1, _, _ := getProc("ImGuiViewport_GetWorkCenter").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiViewportDelete(Self *ImGuiViewport) {
	getProc("ImGuiViewport_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiPlatformIONew() *ImGuiPlatformIO {
	r1, _, _ := getProc("ImGuiPlatformIO_new").Call()
	return (*ImGuiPlatformIO)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiPlatformIONew2(Param1 *ImGuiPlatformIO) *ImGuiPlatformIO {
	r1, _, _ := getProc("ImGuiPlatformIO_new2").Call(uintptr(unsafe.Pointer(Param1)))
	return (*ImGuiPlatformIO)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiPlatformIOPlatformLocaleDecimalPoint(Self *ImGuiPlatformIO) uint16 {
	r1, _, _ := getProc("ImGuiPlatformIO_Platform_LocaleDecimalPoint").Call(uintptr(unsafe.Pointer(Self)))
	return uint16(r1)
}

func (i *Imgui) ImGuiPlatformIOSetPlatformLocaleDecimalPoint(Self *ImGuiPlatformIO, Platform_LocaleDecimalPoint uint16) {
	getProc("ImGuiPlatformIO_setPlatform_LocaleDecimalPoint").Call(uintptr(unsafe.Pointer(Self)), uintptr(Platform_LocaleDecimalPoint))
}

func (i *Imgui) ImGuiPlatformIORendererTextureMaxWidth(Self *ImGuiPlatformIO) int32 {
	r1, _, _ := getProc("ImGuiPlatformIO_Renderer_TextureMaxWidth").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiPlatformIOSetRendererTextureMaxWidth(Self *ImGuiPlatformIO, Renderer_TextureMaxWidth int32) {
	getProc("ImGuiPlatformIO_setRenderer_TextureMaxWidth").Call(uintptr(unsafe.Pointer(Self)), uintptr(Renderer_TextureMaxWidth))
}

func (i *Imgui) ImGuiPlatformIORendererTextureMaxHeight(Self *ImGuiPlatformIO) int32 {
	r1, _, _ := getProc("ImGuiPlatformIO_Renderer_TextureMaxHeight").Call(uintptr(unsafe.Pointer(Self)))
	return int32(r1)
}

func (i *Imgui) ImGuiPlatformIOSetRendererTextureMaxHeight(Self *ImGuiPlatformIO, Renderer_TextureMaxHeight int32) {
	getProc("ImGuiPlatformIO_setRenderer_TextureMaxHeight").Call(uintptr(unsafe.Pointer(Self)), uintptr(Renderer_TextureMaxHeight))
}

func (i *Imgui) ImGuiPlatformIOTextures(Self *ImGuiPlatformIO) unsafe.Pointer {
	r1, _, _ := getProc("ImGuiPlatformIO_Textures").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiPlatformIOSetTextures(Self *ImGuiPlatformIO, Textures unsafe.Pointer) {
	getProc("ImGuiPlatformIO_setTextures").Call(uintptr(unsafe.Pointer(Self)), uintptr(Textures))
}

func (i *Imgui) ImGuiPlatformIOClearPlatformHandlers(Self *ImGuiPlatformIO) {
	getProc("ImGuiPlatformIO_ClearPlatformHandlers").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiPlatformIOClearRendererHandlers(Self *ImGuiPlatformIO) {
	getProc("ImGuiPlatformIO_ClearRendererHandlers").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiPlatformIOOperatorAssign(Self *ImGuiPlatformIO, Param1 *ImGuiPlatformIO) {
	getProc("ImGuiPlatformIO_operatorAssign").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(Param1)))
}

func (i *Imgui) ImGuiPlatformIODelete(Self *ImGuiPlatformIO) {
	getProc("ImGuiPlatformIO_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiPlatformImeDataNew() *ImGuiPlatformImeData {
	r1, _, _ := getProc("ImGuiPlatformImeData_new").Call()
	return (*ImGuiPlatformImeData)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiPlatformImeDataWantVisible(Self *ImGuiPlatformImeData) bool {
	r1, _, _ := getProc("ImGuiPlatformImeData_WantVisible").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiPlatformImeDataSetWantVisible(Self *ImGuiPlatformImeData, WantVisible bool) {
	getProc("ImGuiPlatformImeData_setWantVisible").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantVisible {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiPlatformImeDataWantTextInput(Self *ImGuiPlatformImeData) bool {
	r1, _, _ := getProc("ImGuiPlatformImeData_WantTextInput").Call(uintptr(unsafe.Pointer(Self)))
	return r1 != 0
}

func (i *Imgui) ImGuiPlatformImeDataSetWantTextInput(Self *ImGuiPlatformImeData, WantTextInput bool) {
	getProc("ImGuiPlatformImeData_setWantTextInput").Call(uintptr(unsafe.Pointer(Self)), func() uintptr {
		if WantTextInput {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) ImGuiPlatformImeDataInputPos(Self *ImGuiPlatformImeData) *ImVec2 {
	r1, _, _ := getProc("ImGuiPlatformImeData_InputPos").Call(uintptr(unsafe.Pointer(Self)))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) ImGuiPlatformImeDataSetInputPos(Self *ImGuiPlatformImeData, InputPos *ImVec2) {
	getProc("ImGuiPlatformImeData_setInputPos").Call(uintptr(unsafe.Pointer(Self)), uintptr(unsafe.Pointer(InputPos)))
}

func (i *Imgui) ImGuiPlatformImeDataInputLineHeight(Self *ImGuiPlatformImeData) float32 {
	r1, _, _ := getProc("ImGuiPlatformImeData_InputLineHeight").Call(uintptr(unsafe.Pointer(Self)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) ImGuiPlatformImeDataSetInputLineHeight(Self *ImGuiPlatformImeData, InputLineHeight float32) {
	getProc("ImGuiPlatformImeData_setInputLineHeight").Call(uintptr(unsafe.Pointer(Self)), uintptr(InputLineHeight))
}

func (i *Imgui) ImGuiPlatformImeDataViewportId(Self *ImGuiPlatformImeData) uint32 {
	r1, _, _ := getProc("ImGuiPlatformImeData_ViewportId").Call(uintptr(unsafe.Pointer(Self)))
	return uint32(r1)
}

func (i *Imgui) ImGuiPlatformImeDataSetViewportId(Self *ImGuiPlatformImeData, ViewportId uint32) {
	getProc("ImGuiPlatformImeData_setViewportId").Call(uintptr(unsafe.Pointer(Self)), uintptr(ViewportId))
}

func (i *Imgui) ImGuiPlatformImeDataDelete(Self *ImGuiPlatformImeData) {
	getProc("ImGuiPlatformImeData_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) ImGuiImplDX11RenderStateDevice(Self *ImGui_ImplDX11_RenderState) unsafe.Pointer {
	r1, _, _ := getProc("ImGui_ImplDX11_RenderState_Device").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiImplDX11RenderStateSetDevice(Self *ImGui_ImplDX11_RenderState, Device unsafe.Pointer) {
	getProc("ImGui_ImplDX11_RenderState_setDevice").Call(uintptr(unsafe.Pointer(Self)), uintptr(Device))
}

func (i *Imgui) ImGuiImplDX11RenderStateDeviceContext(Self *ImGui_ImplDX11_RenderState) unsafe.Pointer {
	r1, _, _ := getProc("ImGui_ImplDX11_RenderState_DeviceContext").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiImplDX11RenderStateSetDeviceContext(Self *ImGui_ImplDX11_RenderState, DeviceContext unsafe.Pointer) {
	getProc("ImGui_ImplDX11_RenderState_setDeviceContext").Call(uintptr(unsafe.Pointer(Self)), uintptr(DeviceContext))
}

func (i *Imgui) ImGuiImplDX11RenderStateVertexConstantBuffer(Self *ImGui_ImplDX11_RenderState) unsafe.Pointer {
	r1, _, _ := getProc("ImGui_ImplDX11_RenderState_VertexConstantBuffer").Call(uintptr(unsafe.Pointer(Self)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) ImGuiImplDX11RenderStateSetVertexConstantBuffer(Self *ImGui_ImplDX11_RenderState, VertexConstantBuffer unsafe.Pointer) {
	getProc("ImGui_ImplDX11_RenderState_setVertexConstantBuffer").Call(uintptr(unsafe.Pointer(Self)), uintptr(VertexConstantBuffer))
}

func (i *Imgui) ImGuiImplDX11RenderStateDelete(Self *ImGui_ImplDX11_RenderState) {
	getProc("ImGui_ImplDX11_RenderState_delete").Call(uintptr(unsafe.Pointer(Self)))
}

func (i *Imgui) CabiImGuiCreateContext(Shared_font_atlas *ImFontAtlas) unsafe.Pointer {
	r1, _, _ := getProc("cabi_ImGui__CreateContext").Call(uintptr(unsafe.Pointer(Shared_font_atlas)))
	return unsafe.Pointer(r1)
}

func (i *Imgui) CabiImGuiDestroyContext(Ctx unsafe.Pointer) {
	getProc("cabi_ImGui__DestroyContext").Call(uintptr(Ctx))
}

func (i *Imgui) CabiImGuiGetCurrentContext() unsafe.Pointer {
	r1, _, _ := getProc("cabi_ImGui__GetCurrentContext").Call()
	return unsafe.Pointer(r1)
}

func (i *Imgui) CabiImGuiSetCurrentContext(Ctx unsafe.Pointer) {
	getProc("cabi_ImGui__SetCurrentContext").Call(uintptr(Ctx))
}

func (i *Imgui) CabiImGuiGetIO() *ImGuiIO {
	r1, _, _ := getProc("cabi_ImGui__GetIO").Call()
	return (*ImGuiIO)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetPlatformIO() *ImGuiPlatformIO {
	r1, _, _ := getProc("cabi_ImGui__GetPlatformIO").Call()
	return (*ImGuiPlatformIO)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetStyle() *ImGuiStyle {
	r1, _, _ := getProc("cabi_ImGui__GetStyle").Call()
	return (*ImGuiStyle)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiNewFrame() {
	getProc("cabi_ImGui__NewFrame").Call()
}

func (i *Imgui) CabiImGuiEndFrame() {
	getProc("cabi_ImGui__EndFrame").Call()
}

func (i *Imgui) CabiImGuiRender() {
	getProc("cabi_ImGui__Render").Call()
}

func (i *Imgui) CabiImGuiGetDrawData() *ImDrawData {
	r1, _, _ := getProc("cabi_ImGui__GetDrawData").Call()
	return (*ImDrawData)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiShowDemoWindow(P_open *bool) {
	getProc("cabi_ImGui__ShowDemoWindow").Call(uintptr(unsafe.Pointer(P_open)))
}

func (i *Imgui) CabiImGuiShowMetricsWindow(P_open *bool) {
	getProc("cabi_ImGui__ShowMetricsWindow").Call(uintptr(unsafe.Pointer(P_open)))
}

func (i *Imgui) CabiImGuiShowDebugLogWindow(P_open *bool) {
	getProc("cabi_ImGui__ShowDebugLogWindow").Call(uintptr(unsafe.Pointer(P_open)))
}

func (i *Imgui) CabiImGuiShowIDStackToolWindow(P_open *bool) {
	getProc("cabi_ImGui__ShowIDStackToolWindow").Call(uintptr(unsafe.Pointer(P_open)))
}

func (i *Imgui) CabiImGuiShowAboutWindow(P_open *bool) {
	getProc("cabi_ImGui__ShowAboutWindow").Call(uintptr(unsafe.Pointer(P_open)))
}

func (i *Imgui) CabiImGuiShowStyleEditor(Ref *ImGuiStyle) {
	getProc("cabi_ImGui__ShowStyleEditor").Call(uintptr(unsafe.Pointer(Ref)))
}

func (i *Imgui) CabiImGuiShowStyleSelector(Label *int8) bool {
	r1, _, _ := getProc("cabi_ImGui__ShowStyleSelector").Call(uintptr(unsafe.Pointer(Label)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiShowFontSelector(Label *int8) {
	getProc("cabi_ImGui__ShowFontSelector").Call(uintptr(unsafe.Pointer(Label)))
}

func (i *Imgui) CabiImGuiShowUserGuide() {
	getProc("cabi_ImGui__ShowUserGuide").Call()
}

func (i *Imgui) CabiImGuiGetVersion() *int8 {
	r1, _, _ := getProc("cabi_ImGui__GetVersion").Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiStyleColorsDark(Dst *ImGuiStyle) {
	getProc("cabi_ImGui__StyleColorsDark").Call(uintptr(unsafe.Pointer(Dst)))
}

func (i *Imgui) CabiImGuiStyleColorsLight(Dst *ImGuiStyle) {
	getProc("cabi_ImGui__StyleColorsLight").Call(uintptr(unsafe.Pointer(Dst)))
}

func (i *Imgui) CabiImGuiStyleColorsClassic(Dst *ImGuiStyle) {
	getProc("cabi_ImGui__StyleColorsClassic").Call(uintptr(unsafe.Pointer(Dst)))
}

func (i *Imgui) CabiImGuiBegin(Name *int8, P_open *bool, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__Begin").Call(uintptr(unsafe.Pointer(Name)), uintptr(unsafe.Pointer(P_open)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEnd() {
	getProc("cabi_ImGui__End").Call()
}

func (i *Imgui) CabiImGuiBeginChild1(Str_id *int8, Size *ImVec2, Child_flags int32, Window_flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginChild_1").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(unsafe.Pointer(Size)), uintptr(Child_flags), uintptr(Window_flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginChild2(Id int32, Size *ImVec2, Child_flags int32, Window_flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginChild_2").Call(uintptr(Id), uintptr(unsafe.Pointer(Size)), uintptr(Child_flags), uintptr(Window_flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndChild() {
	getProc("cabi_ImGui__EndChild").Call()
}

func (i *Imgui) CabiImGuiIsWindowAppearing() bool {
	r1, _, _ := getProc("cabi_ImGui__IsWindowAppearing").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsWindowCollapsed() bool {
	r1, _, _ := getProc("cabi_ImGui__IsWindowCollapsed").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsWindowFocused(Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsWindowFocused").Call(uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsWindowHovered(Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsWindowHovered").Call(uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetWindowDrawList() *ImDrawList {
	r1, _, _ := getProc("cabi_ImGui__GetWindowDrawList").Call()
	return (*ImDrawList)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetWindowPos() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetWindowPos").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetWindowSize() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetWindowSize").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetWindowWidth() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetWindowWidth").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetWindowHeight() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetWindowHeight").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiSetNextWindowPos(Pos *ImVec2, Cond int32, Pivot *ImVec2) {
	getProc("cabi_ImGui__SetNextWindowPos").Call(uintptr(unsafe.Pointer(Pos)), uintptr(Cond), uintptr(unsafe.Pointer(Pivot)))
}

func (i *Imgui) CabiImGuiSetNextWindowSize(Size *ImVec2, Cond int32) {
	getProc("cabi_ImGui__SetNextWindowSize").Call(uintptr(unsafe.Pointer(Size)), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetNextWindowSizeConstraints(Size_min *ImVec2, Size_max *ImVec2, Custom_callback unsafe.Pointer, Custom_callback_data unsafe.Pointer) {
	getProc("cabi_ImGui__SetNextWindowSizeConstraints").Call(uintptr(unsafe.Pointer(Size_min)), uintptr(unsafe.Pointer(Size_max)), uintptr(Custom_callback), uintptr(Custom_callback_data))
}

func (i *Imgui) CabiImGuiSetNextWindowContentSize(Size *ImVec2) {
	getProc("cabi_ImGui__SetNextWindowContentSize").Call(uintptr(unsafe.Pointer(Size)))
}

func (i *Imgui) CabiImGuiSetNextWindowCollapsed(Collapsed bool, Cond int32) {
	getProc("cabi_ImGui__SetNextWindowCollapsed").Call(func() uintptr {
		if Collapsed {
			return 1
		}
		return 0
	}(), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetNextWindowFocus() {
	getProc("cabi_ImGui__SetNextWindowFocus").Call()
}

func (i *Imgui) CabiImGuiSetNextWindowScroll(Scroll *ImVec2) {
	getProc("cabi_ImGui__SetNextWindowScroll").Call(uintptr(unsafe.Pointer(Scroll)))
}

func (i *Imgui) CabiImGuiSetNextWindowBgAlpha(Alpha float32) {
	getProc("cabi_ImGui__SetNextWindowBgAlpha").Call(uintptr(Alpha))
}

func (i *Imgui) CabiImGuiSetWindowPos1(Pos *ImVec2, Cond int32) {
	getProc("cabi_ImGui__SetWindowPos_1").Call(uintptr(unsafe.Pointer(Pos)), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetWindowSize1(Size *ImVec2, Cond int32) {
	getProc("cabi_ImGui__SetWindowSize_1").Call(uintptr(unsafe.Pointer(Size)), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetWindowCollapsed1(Collapsed bool, Cond int32) {
	getProc("cabi_ImGui__SetWindowCollapsed_1").Call(func() uintptr {
		if Collapsed {
			return 1
		}
		return 0
	}(), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetWindowFocus1() {
	getProc("cabi_ImGui__SetWindowFocus_1").Call()
}

func (i *Imgui) CabiImGuiSetWindowPos2(Name *int8, Pos *ImVec2, Cond int32) {
	getProc("cabi_ImGui__SetWindowPos_2").Call(uintptr(unsafe.Pointer(Name)), uintptr(unsafe.Pointer(Pos)), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetWindowSize2(Name *int8, Size *ImVec2, Cond int32) {
	getProc("cabi_ImGui__SetWindowSize_2").Call(uintptr(unsafe.Pointer(Name)), uintptr(unsafe.Pointer(Size)), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetWindowCollapsed2(Name *int8, Collapsed bool, Cond int32) {
	getProc("cabi_ImGui__SetWindowCollapsed_2").Call(uintptr(unsafe.Pointer(Name)), func() uintptr {
		if Collapsed {
			return 1
		}
		return 0
	}(), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetWindowFocus2(Name *int8) {
	getProc("cabi_ImGui__SetWindowFocus_2").Call(uintptr(unsafe.Pointer(Name)))
}

func (i *Imgui) CabiImGuiGetScrollX() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetScrollX").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetScrollY() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetScrollY").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiSetScrollX(Scroll_x float32) {
	getProc("cabi_ImGui__SetScrollX").Call(uintptr(Scroll_x))
}

func (i *Imgui) CabiImGuiSetScrollY(Scroll_y float32) {
	getProc("cabi_ImGui__SetScrollY").Call(uintptr(Scroll_y))
}

func (i *Imgui) CabiImGuiGetScrollMaxX() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetScrollMaxX").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetScrollMaxY() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetScrollMaxY").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiSetScrollHereX(Center_x_ratio float32) {
	getProc("cabi_ImGui__SetScrollHereX").Call(uintptr(Center_x_ratio))
}

func (i *Imgui) CabiImGuiSetScrollHereY(Center_y_ratio float32) {
	getProc("cabi_ImGui__SetScrollHereY").Call(uintptr(Center_y_ratio))
}

func (i *Imgui) CabiImGuiSetScrollFromPosX(Local_x float32, Center_x_ratio float32) {
	getProc("cabi_ImGui__SetScrollFromPosX").Call(uintptr(Local_x), uintptr(Center_x_ratio))
}

func (i *Imgui) CabiImGuiSetScrollFromPosY(Local_y float32, Center_y_ratio float32) {
	getProc("cabi_ImGui__SetScrollFromPosY").Call(uintptr(Local_y), uintptr(Center_y_ratio))
}

func (i *Imgui) CabiImGuiPushFont1(Font *ImFont, Font_size_base_unscaled float32) {
	getProc("cabi_ImGui__PushFont_1").Call(uintptr(unsafe.Pointer(Font)), uintptr(Font_size_base_unscaled))
}

func (i *Imgui) CabiImGuiPopFont() {
	getProc("cabi_ImGui__PopFont").Call()
}

func (i *Imgui) CabiImGuiGetFont() *ImFont {
	r1, _, _ := getProc("cabi_ImGui__GetFont").Call()
	return (*ImFont)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetFontSize() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetFontSize").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetFontBaked() *ImFontBaked {
	r1, _, _ := getProc("cabi_ImGui__GetFontBaked").Call()
	return (*ImFontBaked)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiPushStyleColor1(Idx int32, Col int32) {
	getProc("cabi_ImGui__PushStyleColor_1").Call(uintptr(Idx), uintptr(Col))
}

func (i *Imgui) CabiImGuiPushStyleColor2(Idx int32, Col *ImVec4) {
	getProc("cabi_ImGui__PushStyleColor_2").Call(uintptr(Idx), uintptr(unsafe.Pointer(Col)))
}

func (i *Imgui) CabiImGuiPopStyleColor(Count int32) {
	getProc("cabi_ImGui__PopStyleColor").Call(uintptr(Count))
}

func (i *Imgui) CabiImGuiPushStyleVar1(Idx int32, Val float32) {
	getProc("cabi_ImGui__PushStyleVar_1").Call(uintptr(Idx), uintptr(Val))
}

func (i *Imgui) CabiImGuiPushStyleVar2(Idx int32, Val *ImVec2) {
	getProc("cabi_ImGui__PushStyleVar_2").Call(uintptr(Idx), uintptr(unsafe.Pointer(Val)))
}

func (i *Imgui) CabiImGuiPushStyleVarX(Idx int32, Val_x float32) {
	getProc("cabi_ImGui__PushStyleVarX").Call(uintptr(Idx), uintptr(Val_x))
}

func (i *Imgui) CabiImGuiPushStyleVarY(Idx int32, Val_y float32) {
	getProc("cabi_ImGui__PushStyleVarY").Call(uintptr(Idx), uintptr(Val_y))
}

func (i *Imgui) CabiImGuiPopStyleVar(Count int32) {
	getProc("cabi_ImGui__PopStyleVar").Call(uintptr(Count))
}

func (i *Imgui) CabiImGuiPushItemFlag(Option int32, Enabled bool) {
	getProc("cabi_ImGui__PushItemFlag").Call(uintptr(Option), func() uintptr {
		if Enabled {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiPopItemFlag() {
	getProc("cabi_ImGui__PopItemFlag").Call()
}

func (i *Imgui) CabiImGuiPushItemWidth(Item_width float32) {
	getProc("cabi_ImGui__PushItemWidth").Call(uintptr(Item_width))
}

func (i *Imgui) CabiImGuiPopItemWidth() {
	getProc("cabi_ImGui__PopItemWidth").Call()
}

func (i *Imgui) CabiImGuiSetNextItemWidth(Item_width float32) {
	getProc("cabi_ImGui__SetNextItemWidth").Call(uintptr(Item_width))
}

func (i *Imgui) CabiImGuiCalcItemWidth() float32 {
	r1, _, _ := getProc("cabi_ImGui__CalcItemWidth").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiPushTextWrapPos(Wrap_local_pos_x float32) {
	getProc("cabi_ImGui__PushTextWrapPos").Call(uintptr(Wrap_local_pos_x))
}

func (i *Imgui) CabiImGuiPopTextWrapPos() {
	getProc("cabi_ImGui__PopTextWrapPos").Call()
}

func (i *Imgui) CabiImGuiGetFontTexUvWhitePixel() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetFontTexUvWhitePixel").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetColorU321(Idx int32, Alpha_mul float32) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetColorU32_1").Call(uintptr(Idx), uintptr(Alpha_mul))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetColorU322(Col *ImVec4) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetColorU32_2").Call(uintptr(unsafe.Pointer(Col)))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetColorU323(Col int32, Alpha_mul float32) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetColorU32_3").Call(uintptr(Col), uintptr(Alpha_mul))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetStyleColorVec4(Idx int32) *ImVec4 {
	r1, _, _ := getProc("cabi_ImGui__GetStyleColorVec4").Call(uintptr(Idx))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetCursorScreenPos() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetCursorScreenPos").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiSetCursorScreenPos(Pos *ImVec2) {
	getProc("cabi_ImGui__SetCursorScreenPos").Call(uintptr(unsafe.Pointer(Pos)))
}

func (i *Imgui) CabiImGuiGetContentRegionAvail() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetContentRegionAvail").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetCursorPos() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetCursorPos").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetCursorPosX() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetCursorPosX").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetCursorPosY() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetCursorPosY").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiSetCursorPos(Local_pos *ImVec2) {
	getProc("cabi_ImGui__SetCursorPos").Call(uintptr(unsafe.Pointer(Local_pos)))
}

func (i *Imgui) CabiImGuiSetCursorPosX(Local_x float32) {
	getProc("cabi_ImGui__SetCursorPosX").Call(uintptr(Local_x))
}

func (i *Imgui) CabiImGuiSetCursorPosY(Local_y float32) {
	getProc("cabi_ImGui__SetCursorPosY").Call(uintptr(Local_y))
}

func (i *Imgui) CabiImGuiGetCursorStartPos() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetCursorStartPos").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiSeparator() {
	getProc("cabi_ImGui__Separator").Call()
}

func (i *Imgui) CabiImGuiSameLine(Offset_from_start_x float32, Spacing float32) {
	getProc("cabi_ImGui__SameLine").Call(uintptr(Offset_from_start_x), uintptr(Spacing))
}

func (i *Imgui) CabiImGuiNewLine() {
	getProc("cabi_ImGui__NewLine").Call()
}

func (i *Imgui) CabiImGuiSpacing() {
	getProc("cabi_ImGui__Spacing").Call()
}

func (i *Imgui) CabiImGuiDummy(Size *ImVec2) {
	getProc("cabi_ImGui__Dummy").Call(uintptr(unsafe.Pointer(Size)))
}

func (i *Imgui) CabiImGuiIndent(Indent_w float32) {
	getProc("cabi_ImGui__Indent").Call(uintptr(Indent_w))
}

func (i *Imgui) CabiImGuiUnindent(Indent_w float32) {
	getProc("cabi_ImGui__Unindent").Call(uintptr(Indent_w))
}

func (i *Imgui) CabiImGuiBeginGroup() {
	getProc("cabi_ImGui__BeginGroup").Call()
}

func (i *Imgui) CabiImGuiEndGroup() {
	getProc("cabi_ImGui__EndGroup").Call()
}

func (i *Imgui) CabiImGuiAlignTextToFramePadding() {
	getProc("cabi_ImGui__AlignTextToFramePadding").Call()
}

func (i *Imgui) CabiImGuiGetTextLineHeight() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetTextLineHeight").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetTextLineHeightWithSpacing() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetTextLineHeightWithSpacing").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetFrameHeight() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetFrameHeight").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetFrameHeightWithSpacing() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetFrameHeightWithSpacing").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiPushID1(Str_id *int8) {
	getProc("cabi_ImGui__PushID_1").Call(uintptr(unsafe.Pointer(Str_id)))
}

func (i *Imgui) CabiImGuiPushID2(Str_id_begin *int8, Str_id_end *int8) {
	getProc("cabi_ImGui__PushID_2").Call(uintptr(unsafe.Pointer(Str_id_begin)), uintptr(unsafe.Pointer(Str_id_end)))
}

func (i *Imgui) CabiImGuiPushID3(Ptr_id unsafe.Pointer) {
	getProc("cabi_ImGui__PushID_3").Call(uintptr(Ptr_id))
}

func (i *Imgui) CabiImGuiPushID4(Int_id int32) {
	getProc("cabi_ImGui__PushID_4").Call(uintptr(Int_id))
}

func (i *Imgui) CabiImGuiPopID() {
	getProc("cabi_ImGui__PopID").Call()
}

func (i *Imgui) CabiImGuiGetID1(Str_id *int8) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetID_1").Call(uintptr(unsafe.Pointer(Str_id)))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetID2(Str_id_begin *int8, Str_id_end *int8) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetID_2").Call(uintptr(unsafe.Pointer(Str_id_begin)), uintptr(unsafe.Pointer(Str_id_end)))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetID3(Ptr_id unsafe.Pointer) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetID_3").Call(uintptr(Ptr_id))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetID4(Int_id int32) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetID_4").Call(uintptr(Int_id))
	return int32(r1)
}

func (i *Imgui) CabiImGuiTextUnformatted(Text *int8, Text_end *int8) {
	getProc("cabi_ImGui__TextUnformatted").Call(uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)))
}

func (i *Imgui) CabiImGuiTextV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__TextV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiTextColoredV(Col *ImVec4, Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__TextColoredV").Call(uintptr(unsafe.Pointer(Col)), uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiTextDisabledV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__TextDisabledV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiTextWrappedV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__TextWrappedV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiLabelTextV(Label *int8, Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__LabelTextV").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiBulletTextV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__BulletTextV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiSeparatorText(Label *int8) {
	getProc("cabi_ImGui__SeparatorText").Call(uintptr(unsafe.Pointer(Label)))
}

func (i *Imgui) CabiImGuiButton(Label *int8, Size *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__Button").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Size)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSmallButton(Label *int8) bool {
	r1, _, _ := getProc("cabi_ImGui__SmallButton").Call(uintptr(unsafe.Pointer(Label)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInvisibleButton(Str_id *int8, Size *ImVec2, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InvisibleButton").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(unsafe.Pointer(Size)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiArrowButton(Str_id *int8, Dir int32) bool {
	r1, _, _ := getProc("cabi_ImGui__ArrowButton").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Dir))
	return r1 != 0
}

func (i *Imgui) CabiImGuiCheckbox(Label *int8, V *bool) bool {
	r1, _, _ := getProc("cabi_ImGui__Checkbox").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiCheckboxFlags1(Label *int8, Flags *int32, Flags_value int32) bool {
	r1, _, _ := getProc("cabi_ImGui__CheckboxFlags_1").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Flags)), uintptr(Flags_value))
	return r1 != 0
}

func (i *Imgui) CabiImGuiCheckboxFlags2(Label *int8, Flags *uint32, Flags_value uint32) bool {
	r1, _, _ := getProc("cabi_ImGui__CheckboxFlags_2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Flags)), uintptr(Flags_value))
	return r1 != 0
}

func (i *Imgui) CabiImGuiRadioButton1(Label *int8, Active bool) bool {
	r1, _, _ := getProc("cabi_ImGui__RadioButton_1").Call(uintptr(unsafe.Pointer(Label)), func() uintptr {
		if Active {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiRadioButton2(Label *int8, V *int32, V_button int32) bool {
	r1, _, _ := getProc("cabi_ImGui__RadioButton_2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_button))
	return r1 != 0
}

func (i *Imgui) CabiImGuiProgressBar(Fraction float32, Size_arg *ImVec2, Overlay *int8) {
	getProc("cabi_ImGui__ProgressBar").Call(uintptr(Fraction), uintptr(unsafe.Pointer(Size_arg)), uintptr(unsafe.Pointer(Overlay)))
}

func (i *Imgui) CabiImGuiBullet() {
	getProc("cabi_ImGui__Bullet").Call()
}

func (i *Imgui) CabiImGuiTextLink(Label *int8) bool {
	r1, _, _ := getProc("cabi_ImGui__TextLink").Call(uintptr(unsafe.Pointer(Label)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTextLinkOpenURL(Label *int8, Url *int8) bool {
	r1, _, _ := getProc("cabi_ImGui__TextLinkOpenURL").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Url)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiImage1(Tex_ref *ImTextureRef, Image_size *ImVec2, Uv0 *ImVec2, Uv1 *ImVec2) {
	getProc("cabi_ImGui__Image_1").Call(uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(Image_size)), uintptr(unsafe.Pointer(Uv0)), uintptr(unsafe.Pointer(Uv1)))
}

func (i *Imgui) CabiImGuiImageWithBg(Tex_ref *ImTextureRef, Image_size *ImVec2, Uv0 *ImVec2, Uv1 *ImVec2, Bg_col *ImVec4, Tint_col *ImVec4) {
	getProc("cabi_ImGui__ImageWithBg").Call(uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(Image_size)), uintptr(unsafe.Pointer(Uv0)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Bg_col)), uintptr(unsafe.Pointer(Tint_col)))
}

func (i *Imgui) CabiImGuiImageButton(Str_id *int8, Tex_ref *ImTextureRef, Image_size *ImVec2, Uv0 *ImVec2, Uv1 *ImVec2, Bg_col *ImVec4, Tint_col *ImVec4) bool {
	r1, _, _ := getProc("cabi_ImGui__ImageButton").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(Image_size)), uintptr(unsafe.Pointer(Uv0)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Bg_col)), uintptr(unsafe.Pointer(Tint_col)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginCombo(Label *int8, Preview_value *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginCombo").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Preview_value)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndCombo() {
	getProc("cabi_ImGui__EndCombo").Call()
}

func (i *Imgui) CabiImGuiCombo1(Label *int8, Current_item *int32, Items **int8, Items_count int32, Popup_max_height_in_items int32) bool {
	r1, _, _ := getProc("cabi_ImGui__Combo_1").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Current_item)), uintptr(unsafe.Pointer(Items)), uintptr(Items_count), uintptr(Popup_max_height_in_items))
	return r1 != 0
}

func (i *Imgui) CabiImGuiCombo2(Label *int8, Current_item *int32, Items_separated_by_zeros *int8, Popup_max_height_in_items int32) bool {
	r1, _, _ := getProc("cabi_ImGui__Combo_2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Current_item)), uintptr(unsafe.Pointer(Items_separated_by_zeros)), uintptr(Popup_max_height_in_items))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragFloat(Label *int8, V *float32, V_speed float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragFloat").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragFloat2(Label *int8, V *float32, V_speed float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragFloat2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragFloat3(Label *int8, V *float32, V_speed float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragFloat3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragFloat4(Label *int8, V *float32, V_speed float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragFloat4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragFloatRange2(Label *int8, V_current_min *float32, V_current_max *float32, V_speed float32, V_min float32, V_max float32, Format *int8, Format_max *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragFloatRange2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V_current_min)), uintptr(unsafe.Pointer(V_current_max)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(unsafe.Pointer(Format_max)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragInt(Label *int8, V *int32, V_speed float32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragInt").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragInt2(Label *int8, V *int32, V_speed float32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragInt2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragInt3(Label *int8, V *int32, V_speed float32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragInt3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragInt4(Label *int8, V *int32, V_speed float32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragInt4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragIntRange2(Label *int8, V_current_min *int32, V_current_max *int32, V_speed float32, V_min int32, V_max int32, Format *int8, Format_max *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragIntRange2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V_current_min)), uintptr(unsafe.Pointer(V_current_max)), uintptr(V_speed), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(unsafe.Pointer(Format_max)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragScalar(Label *int8, Data_type int32, P_data unsafe.Pointer, V_speed float32, P_min unsafe.Pointer, P_max unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragScalar").Call(uintptr(unsafe.Pointer(Label)), uintptr(Data_type), uintptr(P_data), uintptr(V_speed), uintptr(P_min), uintptr(P_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiDragScalarN(Label *int8, Data_type int32, P_data unsafe.Pointer, Components int32, V_speed float32, P_min unsafe.Pointer, P_max unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__DragScalarN").Call(uintptr(unsafe.Pointer(Label)), uintptr(Data_type), uintptr(P_data), uintptr(Components), uintptr(V_speed), uintptr(P_min), uintptr(P_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderFloat(Label *int8, V *float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderFloat").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderFloat2(Label *int8, V *float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderFloat2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderFloat3(Label *int8, V *float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderFloat3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderFloat4(Label *int8, V *float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderFloat4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderAngle(Label *int8, V_rad *float32, V_degrees_min float32, V_degrees_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderAngle").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V_rad)), uintptr(V_degrees_min), uintptr(V_degrees_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderInt(Label *int8, V *int32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderInt").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderInt2(Label *int8, V *int32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderInt2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderInt3(Label *int8, V *int32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderInt3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderInt4(Label *int8, V *int32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderInt4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderScalar(Label *int8, Data_type int32, P_data unsafe.Pointer, P_min unsafe.Pointer, P_max unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderScalar").Call(uintptr(unsafe.Pointer(Label)), uintptr(Data_type), uintptr(P_data), uintptr(P_min), uintptr(P_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSliderScalarN(Label *int8, Data_type int32, P_data unsafe.Pointer, Components int32, P_min unsafe.Pointer, P_max unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SliderScalarN").Call(uintptr(unsafe.Pointer(Label)), uintptr(Data_type), uintptr(P_data), uintptr(Components), uintptr(P_min), uintptr(P_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiVSliderFloat(Label *int8, Size *ImVec2, V *float32, V_min float32, V_max float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__VSliderFloat").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Size)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiVSliderInt(Label *int8, Size *ImVec2, V *int32, V_min int32, V_max int32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__VSliderInt").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Size)), uintptr(unsafe.Pointer(V)), uintptr(V_min), uintptr(V_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiVSliderScalar(Label *int8, Size *ImVec2, Data_type int32, P_data unsafe.Pointer, P_min unsafe.Pointer, P_max unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__VSliderScalar").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Size)), uintptr(Data_type), uintptr(P_data), uintptr(P_min), uintptr(P_max), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputText(Label *int8, Buf *int8, Buf_size uintptr, Flags int32, Callback unsafe.Pointer, User_data unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__InputText").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Buf)), Buf_size, uintptr(Flags), uintptr(Callback), uintptr(User_data))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputTextMultiline(Label *int8, Buf *int8, Buf_size uintptr, Size *ImVec2, Flags int32, Callback unsafe.Pointer, User_data unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__InputTextMultiline").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Buf)), Buf_size, uintptr(unsafe.Pointer(Size)), uintptr(Flags), uintptr(Callback), uintptr(User_data))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputTextWithHint(Label *int8, Hint *int8, Buf *int8, Buf_size uintptr, Flags int32, Callback unsafe.Pointer, User_data unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__InputTextWithHint").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Hint)), uintptr(unsafe.Pointer(Buf)), Buf_size, uintptr(Flags), uintptr(Callback), uintptr(User_data))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputFloat(Label *int8, V *float32, Step float32, Step_fast float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputFloat").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(Step), uintptr(Step_fast), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputFloat2(Label *int8, V *float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputFloat2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputFloat3(Label *int8, V *float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputFloat3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputFloat4(Label *int8, V *float32, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputFloat4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputInt(Label *int8, V *int32, Step int32, Step_fast int32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputInt").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(Step), uintptr(Step_fast), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputInt2(Label *int8, V *int32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputInt2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputInt3(Label *int8, V *int32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputInt3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputInt4(Label *int8, V *int32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputInt4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputDouble(Label *int8, V *float64, Step float64, Step_fast float64, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputDouble").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(V)), uintptr(Step), uintptr(Step_fast), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputScalar(Label *int8, Data_type int32, P_data unsafe.Pointer, P_step unsafe.Pointer, P_step_fast unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputScalar").Call(uintptr(unsafe.Pointer(Label)), uintptr(Data_type), uintptr(P_data), uintptr(P_step), uintptr(P_step_fast), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiInputScalarN(Label *int8, Data_type int32, P_data unsafe.Pointer, Components int32, P_step unsafe.Pointer, P_step_fast unsafe.Pointer, Format *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__InputScalarN").Call(uintptr(unsafe.Pointer(Label)), uintptr(Data_type), uintptr(P_data), uintptr(Components), uintptr(P_step), uintptr(P_step_fast), uintptr(unsafe.Pointer(Format)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiColorEdit3(Label *int8, Col *float32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__ColorEdit3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Col)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiColorEdit4(Label *int8, Col *float32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__ColorEdit4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Col)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiColorPicker3(Label *int8, Col *float32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__ColorPicker3").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Col)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiColorPicker4(Label *int8, Col *float32, Flags int32, Ref_col *float32) bool {
	r1, _, _ := getProc("cabi_ImGui__ColorPicker4").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Col)), uintptr(Flags), uintptr(unsafe.Pointer(Ref_col)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiColorButton(Desc_id *int8, Col *ImVec4, Flags int32, Size *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__ColorButton").Call(uintptr(unsafe.Pointer(Desc_id)), uintptr(unsafe.Pointer(Col)), uintptr(Flags), uintptr(unsafe.Pointer(Size)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSetColorEditOptions(Flags int32) {
	getProc("cabi_ImGui__SetColorEditOptions").Call(uintptr(Flags))
}

func (i *Imgui) CabiImGuiTreeNode(Label *int8) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNode").Call(uintptr(unsafe.Pointer(Label)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTreeNodeV1(Str_id *int8, Fmt *int8, Args unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNodeV_1").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTreeNodeV2(Ptr_id unsafe.Pointer, Fmt *int8, Args unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNodeV_2").Call(uintptr(Ptr_id), uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTreeNodeEx(Label *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNodeEx").Call(uintptr(unsafe.Pointer(Label)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTreeNodeExV1(Str_id *int8, Flags int32, Fmt *int8, Args unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNodeExV_1").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Flags), uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTreeNodeExV2(Ptr_id unsafe.Pointer, Flags int32, Fmt *int8, Args unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNodeExV_2").Call(uintptr(Ptr_id), uintptr(Flags), uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTreePush1(Str_id *int8) {
	getProc("cabi_ImGui__TreePush_1").Call(uintptr(unsafe.Pointer(Str_id)))
}

func (i *Imgui) CabiImGuiTreePush2(Ptr_id unsafe.Pointer) {
	getProc("cabi_ImGui__TreePush_2").Call(uintptr(Ptr_id))
}

func (i *Imgui) CabiImGuiTreePop() {
	getProc("cabi_ImGui__TreePop").Call()
}

func (i *Imgui) CabiImGuiGetTreeNodeToLabelSpacing() float32 {
	r1, _, _ := getProc("cabi_ImGui__GetTreeNodeToLabelSpacing").Call()
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiCollapsingHeader1(Label *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__CollapsingHeader_1").Call(uintptr(unsafe.Pointer(Label)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiCollapsingHeader2(Label *int8, P_visible *bool, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__CollapsingHeader_2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(P_visible)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSetNextItemOpen(Is_open bool, Cond int32) {
	getProc("cabi_ImGui__SetNextItemOpen").Call(func() uintptr {
		if Is_open {
			return 1
		}
		return 0
	}(), uintptr(Cond))
}

func (i *Imgui) CabiImGuiSetNextItemStorageID(Storage_id int32) {
	getProc("cabi_ImGui__SetNextItemStorageID").Call(uintptr(Storage_id))
}

func (i *Imgui) CabiImGuiTreeNodeGetOpen(Storage_id int32) bool {
	r1, _, _ := getProc("cabi_ImGui__TreeNodeGetOpen").Call(uintptr(Storage_id))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSelectable1(Label *int8, Selected bool, Flags int32, Size *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__Selectable_1").Call(uintptr(unsafe.Pointer(Label)), func() uintptr {
		if Selected {
			return 1
		}
		return 0
	}(), uintptr(Flags), uintptr(unsafe.Pointer(Size)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSelectable2(Label *int8, P_selected *bool, Flags int32, Size *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__Selectable_2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(P_selected)), uintptr(Flags), uintptr(unsafe.Pointer(Size)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginMultiSelect(Flags int32, Selection_size int32, Items_count int32) *ImGuiMultiSelectIO {
	r1, _, _ := getProc("cabi_ImGui__BeginMultiSelect").Call(uintptr(Flags), uintptr(Selection_size), uintptr(Items_count))
	return (*ImGuiMultiSelectIO)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiEndMultiSelect() *ImGuiMultiSelectIO {
	r1, _, _ := getProc("cabi_ImGui__EndMultiSelect").Call()
	return (*ImGuiMultiSelectIO)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiSetNextItemSelectionUserData(Selection_user_data int32) {
	getProc("cabi_ImGui__SetNextItemSelectionUserData").Call(uintptr(Selection_user_data))
}

func (i *Imgui) CabiImGuiIsItemToggledSelection() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemToggledSelection").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginListBox(Label *int8, Size *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginListBox").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Size)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndListBox() {
	getProc("cabi_ImGui__EndListBox").Call()
}

func (i *Imgui) CabiImGuiListBox(Label *int8, Current_item *int32, Items **int8, Items_count int32, Height_in_items int32) bool {
	r1, _, _ := getProc("cabi_ImGui__ListBox").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Current_item)), uintptr(unsafe.Pointer(Items)), uintptr(Items_count), uintptr(Height_in_items))
	return r1 != 0
}

func (i *Imgui) CabiImGuiPlotLines(Label *int8, Values *float32, Values_count int32, Values_offset int32, Overlay_text *int8, Scale_min float32, Scale_max float32, Graph_size *ImVec2, Stride int32) {
	getProc("cabi_ImGui__PlotLines").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Values)), uintptr(Values_count), uintptr(Values_offset), uintptr(unsafe.Pointer(Overlay_text)), uintptr(Scale_min), uintptr(Scale_max), uintptr(unsafe.Pointer(Graph_size)), uintptr(Stride))
}

func (i *Imgui) CabiImGuiPlotHistogram(Label *int8, Values *float32, Values_count int32, Values_offset int32, Overlay_text *int8, Scale_min float32, Scale_max float32, Graph_size *ImVec2, Stride int32) {
	getProc("cabi_ImGui__PlotHistogram").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Values)), uintptr(Values_count), uintptr(Values_offset), uintptr(unsafe.Pointer(Overlay_text)), uintptr(Scale_min), uintptr(Scale_max), uintptr(unsafe.Pointer(Graph_size)), uintptr(Stride))
}

func (i *Imgui) CabiImGuiValue1(Prefix *int8, B bool) {
	getProc("cabi_ImGui__Value_1").Call(uintptr(unsafe.Pointer(Prefix)), func() uintptr {
		if B {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiValue2(Prefix *int8, V int32) {
	getProc("cabi_ImGui__Value_2").Call(uintptr(unsafe.Pointer(Prefix)), uintptr(V))
}

func (i *Imgui) CabiImGuiValue3(Prefix *int8, V uint32) {
	getProc("cabi_ImGui__Value_3").Call(uintptr(unsafe.Pointer(Prefix)), uintptr(V))
}

func (i *Imgui) CabiImGuiValue4(Prefix *int8, V float32, Float_format *int8) {
	getProc("cabi_ImGui__Value_4").Call(uintptr(unsafe.Pointer(Prefix)), uintptr(V), uintptr(unsafe.Pointer(Float_format)))
}

func (i *Imgui) CabiImGuiBeginMenuBar() bool {
	r1, _, _ := getProc("cabi_ImGui__BeginMenuBar").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndMenuBar() {
	getProc("cabi_ImGui__EndMenuBar").Call()
}

func (i *Imgui) CabiImGuiBeginMainMenuBar() bool {
	r1, _, _ := getProc("cabi_ImGui__BeginMainMenuBar").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndMainMenuBar() {
	getProc("cabi_ImGui__EndMainMenuBar").Call()
}

func (i *Imgui) CabiImGuiBeginMenu(Label *int8, Enabled bool) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginMenu").Call(uintptr(unsafe.Pointer(Label)), func() uintptr {
		if Enabled {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndMenu() {
	getProc("cabi_ImGui__EndMenu").Call()
}

func (i *Imgui) CabiImGuiMenuItem1(Label *int8, Shortcut *int8, Selected bool, Enabled bool) bool {
	r1, _, _ := getProc("cabi_ImGui__MenuItem_1").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Shortcut)), func() uintptr {
		if Selected {
			return 1
		}
		return 0
	}(), func() uintptr {
		if Enabled {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiMenuItem2(Label *int8, Shortcut *int8, P_selected *bool, Enabled bool) bool {
	r1, _, _ := getProc("cabi_ImGui__MenuItem_2").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(Shortcut)), uintptr(unsafe.Pointer(P_selected)), func() uintptr {
		if Enabled {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginTooltip() bool {
	r1, _, _ := getProc("cabi_ImGui__BeginTooltip").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndTooltip() {
	getProc("cabi_ImGui__EndTooltip").Call()
}

func (i *Imgui) CabiImGuiSetTooltipV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__SetTooltipV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiBeginItemTooltip() bool {
	r1, _, _ := getProc("cabi_ImGui__BeginItemTooltip").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiSetItemTooltipV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__SetItemTooltipV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiBeginPopup(Str_id *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginPopup").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginPopupModal(Name *int8, P_open *bool, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginPopupModal").Call(uintptr(unsafe.Pointer(Name)), uintptr(unsafe.Pointer(P_open)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndPopup() {
	getProc("cabi_ImGui__EndPopup").Call()
}

func (i *Imgui) CabiImGuiOpenPopup1(Str_id *int8, Popup_flags int32) {
	getProc("cabi_ImGui__OpenPopup_1").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Popup_flags))
}

func (i *Imgui) CabiImGuiOpenPopup2(Id int32, Popup_flags int32) {
	getProc("cabi_ImGui__OpenPopup_2").Call(uintptr(Id), uintptr(Popup_flags))
}

func (i *Imgui) CabiImGuiOpenPopupOnItemClick(Str_id *int8, Popup_flags int32) {
	getProc("cabi_ImGui__OpenPopupOnItemClick").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Popup_flags))
}

func (i *Imgui) CabiImGuiCloseCurrentPopup() {
	getProc("cabi_ImGui__CloseCurrentPopup").Call()
}

func (i *Imgui) CabiImGuiBeginPopupContextItem(Str_id *int8, Popup_flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginPopupContextItem").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Popup_flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginPopupContextWindow(Str_id *int8, Popup_flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginPopupContextWindow").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Popup_flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginPopupContextVoid(Str_id *int8, Popup_flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginPopupContextVoid").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Popup_flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsPopupOpen(Str_id *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsPopupOpen").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiBeginTable(Str_id *int8, Columns int32, Flags int32, Outer_size *ImVec2, Inner_width float32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginTable").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Columns), uintptr(Flags), uintptr(unsafe.Pointer(Outer_size)), uintptr(Inner_width))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndTable() {
	getProc("cabi_ImGui__EndTable").Call()
}

func (i *Imgui) CabiImGuiTableNextRow(Row_flags int32, Min_row_height float32) {
	getProc("cabi_ImGui__TableNextRow").Call(uintptr(Row_flags), uintptr(Min_row_height))
}

func (i *Imgui) CabiImGuiTableNextColumn() bool {
	r1, _, _ := getProc("cabi_ImGui__TableNextColumn").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiTableSetColumnIndex(Column_n int32) bool {
	r1, _, _ := getProc("cabi_ImGui__TableSetColumnIndex").Call(uintptr(Column_n))
	return r1 != 0
}

func (i *Imgui) CabiImGuiTableSetupColumn(Label *int8, Flags int32, Init_width_or_weight float32, User_id int32) {
	getProc("cabi_ImGui__TableSetupColumn").Call(uintptr(unsafe.Pointer(Label)), uintptr(Flags), uintptr(Init_width_or_weight), uintptr(User_id))
}

func (i *Imgui) CabiImGuiTableSetupScrollFreeze(Cols int32, Rows int32) {
	getProc("cabi_ImGui__TableSetupScrollFreeze").Call(uintptr(Cols), uintptr(Rows))
}

func (i *Imgui) CabiImGuiTableHeader(Label *int8) {
	getProc("cabi_ImGui__TableHeader").Call(uintptr(unsafe.Pointer(Label)))
}

func (i *Imgui) CabiImGuiTableHeadersRow() {
	getProc("cabi_ImGui__TableHeadersRow").Call()
}

func (i *Imgui) CabiImGuiTableAngledHeadersRow() {
	getProc("cabi_ImGui__TableAngledHeadersRow").Call()
}

func (i *Imgui) CabiImGuiTableGetSortSpecs() *ImGuiTableSortSpecs {
	r1, _, _ := getProc("cabi_ImGui__TableGetSortSpecs").Call()
	return (*ImGuiTableSortSpecs)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiTableGetColumnCount() int32 {
	r1, _, _ := getProc("cabi_ImGui__TableGetColumnCount").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiTableGetColumnIndex() int32 {
	r1, _, _ := getProc("cabi_ImGui__TableGetColumnIndex").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiTableGetRowIndex() int32 {
	r1, _, _ := getProc("cabi_ImGui__TableGetRowIndex").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiTableGetColumnName(Column_n int32) *int8 {
	r1, _, _ := getProc("cabi_ImGui__TableGetColumnName").Call(uintptr(Column_n))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiTableGetColumnFlags(Column_n int32) int32 {
	r1, _, _ := getProc("cabi_ImGui__TableGetColumnFlags").Call(uintptr(Column_n))
	return int32(r1)
}

func (i *Imgui) CabiImGuiTableSetColumnEnabled(Column_n int32, V bool) {
	getProc("cabi_ImGui__TableSetColumnEnabled").Call(uintptr(Column_n), func() uintptr {
		if V {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiTableGetHoveredColumn() int32 {
	r1, _, _ := getProc("cabi_ImGui__TableGetHoveredColumn").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiTableSetBgColor(Target int32, Color int32, Column_n int32) {
	getProc("cabi_ImGui__TableSetBgColor").Call(uintptr(Target), uintptr(Color), uintptr(Column_n))
}

func (i *Imgui) CabiImGuiColumns(Count int32, Id *int8, Borders bool) {
	getProc("cabi_ImGui__Columns").Call(uintptr(Count), uintptr(unsafe.Pointer(Id)), func() uintptr {
		if Borders {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiNextColumn() {
	getProc("cabi_ImGui__NextColumn").Call()
}

func (i *Imgui) CabiImGuiGetColumnIndex() int32 {
	r1, _, _ := getProc("cabi_ImGui__GetColumnIndex").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetColumnWidth(Column_index int32) float32 {
	r1, _, _ := getProc("cabi_ImGui__GetColumnWidth").Call(uintptr(Column_index))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiSetColumnWidth(Column_index int32, Width float32) {
	getProc("cabi_ImGui__SetColumnWidth").Call(uintptr(Column_index), uintptr(Width))
}

func (i *Imgui) CabiImGuiGetColumnOffset(Column_index int32) float32 {
	r1, _, _ := getProc("cabi_ImGui__GetColumnOffset").Call(uintptr(Column_index))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiSetColumnOffset(Column_index int32, Offset_x float32) {
	getProc("cabi_ImGui__SetColumnOffset").Call(uintptr(Column_index), uintptr(Offset_x))
}

func (i *Imgui) CabiImGuiGetColumnsCount() int32 {
	r1, _, _ := getProc("cabi_ImGui__GetColumnsCount").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiBeginTabBar(Str_id *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginTabBar").Call(uintptr(unsafe.Pointer(Str_id)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndTabBar() {
	getProc("cabi_ImGui__EndTabBar").Call()
}

func (i *Imgui) CabiImGuiBeginTabItem(Label *int8, P_open *bool, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginTabItem").Call(uintptr(unsafe.Pointer(Label)), uintptr(unsafe.Pointer(P_open)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndTabItem() {
	getProc("cabi_ImGui__EndTabItem").Call()
}

func (i *Imgui) CabiImGuiTabItemButton(Label *int8, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__TabItemButton").Call(uintptr(unsafe.Pointer(Label)), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSetTabItemClosed(Tab_or_docked_window_label *int8) {
	getProc("cabi_ImGui__SetTabItemClosed").Call(uintptr(unsafe.Pointer(Tab_or_docked_window_label)))
}

func (i *Imgui) CabiImGuiLogToTTY(Auto_open_depth int32) {
	getProc("cabi_ImGui__LogToTTY").Call(uintptr(Auto_open_depth))
}

func (i *Imgui) CabiImGuiLogToFile(Auto_open_depth int32, Filename *int8) {
	getProc("cabi_ImGui__LogToFile").Call(uintptr(Auto_open_depth), uintptr(unsafe.Pointer(Filename)))
}

func (i *Imgui) CabiImGuiLogToClipboard(Auto_open_depth int32) {
	getProc("cabi_ImGui__LogToClipboard").Call(uintptr(Auto_open_depth))
}

func (i *Imgui) CabiImGuiLogFinish() {
	getProc("cabi_ImGui__LogFinish").Call()
}

func (i *Imgui) CabiImGuiLogButtons() {
	getProc("cabi_ImGui__LogButtons").Call()
}

func (i *Imgui) CabiImGuiLogTextV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__LogTextV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiBeginDragDropSource(Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__BeginDragDropSource").Call(uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSetDragDropPayload(Type *int8, Data unsafe.Pointer, Sz uintptr, Cond int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SetDragDropPayload").Call(uintptr(unsafe.Pointer(Type)), uintptr(Data), Sz, uintptr(Cond))
	return r1 != 0
}

func (i *Imgui) CabiImGuiEndDragDropSource() {
	getProc("cabi_ImGui__EndDragDropSource").Call()
}

func (i *Imgui) CabiImGuiBeginDragDropTarget() bool {
	r1, _, _ := getProc("cabi_ImGui__BeginDragDropTarget").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiAcceptDragDropPayload(Type *int8, Flags int32) *ImGuiPayload {
	r1, _, _ := getProc("cabi_ImGui__AcceptDragDropPayload").Call(uintptr(unsafe.Pointer(Type)), uintptr(Flags))
	return (*ImGuiPayload)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiEndDragDropTarget() {
	getProc("cabi_ImGui__EndDragDropTarget").Call()
}

func (i *Imgui) CabiImGuiGetDragDropPayload() *ImGuiPayload {
	r1, _, _ := getProc("cabi_ImGui__GetDragDropPayload").Call()
	return (*ImGuiPayload)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiBeginDisabled(Disabled bool) {
	getProc("cabi_ImGui__BeginDisabled").Call(func() uintptr {
		if Disabled {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiEndDisabled() {
	getProc("cabi_ImGui__EndDisabled").Call()
}

func (i *Imgui) CabiImGuiPushClipRect(Clip_rect_min *ImVec2, Clip_rect_max *ImVec2, Intersect_with_current_clip_rect bool) {
	getProc("cabi_ImGui__PushClipRect").Call(uintptr(unsafe.Pointer(Clip_rect_min)), uintptr(unsafe.Pointer(Clip_rect_max)), func() uintptr {
		if Intersect_with_current_clip_rect {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiPopClipRect() {
	getProc("cabi_ImGui__PopClipRect").Call()
}

func (i *Imgui) CabiImGuiSetItemDefaultFocus() {
	getProc("cabi_ImGui__SetItemDefaultFocus").Call()
}

func (i *Imgui) CabiImGuiSetKeyboardFocusHere(Offset int32) {
	getProc("cabi_ImGui__SetKeyboardFocusHere").Call(uintptr(Offset))
}

func (i *Imgui) CabiImGuiSetNavCursorVisible(Visible bool) {
	getProc("cabi_ImGui__SetNavCursorVisible").Call(func() uintptr {
		if Visible {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiSetNextItemAllowOverlap() {
	getProc("cabi_ImGui__SetNextItemAllowOverlap").Call()
}

func (i *Imgui) CabiImGuiIsItemHovered(Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemHovered").Call(uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemActive() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemActive").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemFocused() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemFocused").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemClicked(Mouse_button int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemClicked").Call(uintptr(Mouse_button))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemVisible() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemVisible").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemEdited() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemEdited").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemActivated() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemActivated").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemDeactivated() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemDeactivated").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemDeactivatedAfterEdit() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemDeactivatedAfterEdit").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsItemToggledOpen() bool {
	r1, _, _ := getProc("cabi_ImGui__IsItemToggledOpen").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsAnyItemHovered() bool {
	r1, _, _ := getProc("cabi_ImGui__IsAnyItemHovered").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsAnyItemActive() bool {
	r1, _, _ := getProc("cabi_ImGui__IsAnyItemActive").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsAnyItemFocused() bool {
	r1, _, _ := getProc("cabi_ImGui__IsAnyItemFocused").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetItemID() int32 {
	r1, _, _ := getProc("cabi_ImGui__GetItemID").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetItemRectMin() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetItemRectMin").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetItemRectMax() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetItemRectMax").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetItemRectSize() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetItemRectSize").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetItemFlags() int32 {
	r1, _, _ := getProc("cabi_ImGui__GetItemFlags").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetMainViewport() *ImGuiViewport {
	r1, _, _ := getProc("cabi_ImGui__GetMainViewport").Call()
	return (*ImGuiViewport)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetBackgroundDrawList() *ImDrawList {
	r1, _, _ := getProc("cabi_ImGui__GetBackgroundDrawList").Call()
	return (*ImDrawList)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetForegroundDrawList() *ImDrawList {
	r1, _, _ := getProc("cabi_ImGui__GetForegroundDrawList").Call()
	return (*ImDrawList)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiIsRectVisible1(Size *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__IsRectVisible_1").Call(uintptr(unsafe.Pointer(Size)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsRectVisible2(Rect_min *ImVec2, Rect_max *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__IsRectVisible_2").Call(uintptr(unsafe.Pointer(Rect_min)), uintptr(unsafe.Pointer(Rect_max)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetTime() float64 {
	r1, _, _ := getProc("cabi_ImGui__GetTime").Call()
	return *(*float64)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiGetFrameCount() int32 {
	r1, _, _ := getProc("cabi_ImGui__GetFrameCount").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetDrawListSharedData() unsafe.Pointer {
	r1, _, _ := getProc("cabi_ImGui__GetDrawListSharedData").Call()
	return unsafe.Pointer(r1)
}

func (i *Imgui) CabiImGuiGetStyleColorName(Idx int32) *int8 {
	r1, _, _ := getProc("cabi_ImGui__GetStyleColorName").Call(uintptr(Idx))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiSetStateStorage(Storage *ImGuiStorage) {
	getProc("cabi_ImGui__SetStateStorage").Call(uintptr(unsafe.Pointer(Storage)))
}

func (i *Imgui) CabiImGuiGetStateStorage() *ImGuiStorage {
	r1, _, _ := getProc("cabi_ImGui__GetStateStorage").Call()
	return (*ImGuiStorage)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiCalcTextSize(Text *int8, Text_end *int8, Hide_text_after_double_hash bool, Wrap_width float32) *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__CalcTextSize").Call(uintptr(unsafe.Pointer(Text)), uintptr(unsafe.Pointer(Text_end)), func() uintptr {
		if Hide_text_after_double_hash {
			return 1
		}
		return 0
	}(), uintptr(Wrap_width))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiColorConvertU32ToFloat4(In int32) *ImVec4 {
	r1, _, _ := getProc("cabi_ImGui__ColorConvertU32ToFloat4").Call(uintptr(In))
	return (*ImVec4)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiColorConvertFloat4ToU32(In *ImVec4) int32 {
	r1, _, _ := getProc("cabi_ImGui__ColorConvertFloat4ToU32").Call(uintptr(unsafe.Pointer(In)))
	return int32(r1)
}

func (i *Imgui) CabiImGuiColorConvertRGBtoHSV(R float32, G float32, B float32, Out_h *float32, Out_s *float32, Out_v *float32) {
	getProc("cabi_ImGui__ColorConvertRGBtoHSV").Call(uintptr(R), uintptr(G), uintptr(B), uintptr(unsafe.Pointer(Out_h)), uintptr(unsafe.Pointer(Out_s)), uintptr(unsafe.Pointer(Out_v)))
}

func (i *Imgui) CabiImGuiColorConvertHSVtoRGB(H float32, S float32, V float32, Out_r *float32, Out_g *float32, Out_b *float32) {
	getProc("cabi_ImGui__ColorConvertHSVtoRGB").Call(uintptr(H), uintptr(S), uintptr(V), uintptr(unsafe.Pointer(Out_r)), uintptr(unsafe.Pointer(Out_g)), uintptr(unsafe.Pointer(Out_b)))
}

func (i *Imgui) CabiImGuiIsKeyDown(Key int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsKeyDown").Call(uintptr(Key))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsKeyPressed(Key int32, Repeat bool) bool {
	r1, _, _ := getProc("cabi_ImGui__IsKeyPressed").Call(uintptr(Key), func() uintptr {
		if Repeat {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsKeyReleased(Key int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsKeyReleased").Call(uintptr(Key))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsKeyChordPressed(Key_chord int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsKeyChordPressed").Call(uintptr(Key_chord))
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetKeyPressedAmount(Key int32, Repeat_delay float32, Rate float32) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetKeyPressedAmount").Call(uintptr(Key), uintptr(Repeat_delay), uintptr(Rate))
	return int32(r1)
}

func (i *Imgui) CabiImGuiGetKeyName(Key int32) *int8 {
	r1, _, _ := getProc("cabi_ImGui__GetKeyName").Call(uintptr(Key))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiSetNextFrameWantCaptureKeyboard(Want_capture_keyboard bool) {
	getProc("cabi_ImGui__SetNextFrameWantCaptureKeyboard").Call(func() uintptr {
		if Want_capture_keyboard {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiShortcut(Key_chord int32, Flags int32) bool {
	r1, _, _ := getProc("cabi_ImGui__Shortcut").Call(uintptr(Key_chord), uintptr(Flags))
	return r1 != 0
}

func (i *Imgui) CabiImGuiSetNextItemShortcut(Key_chord int32, Flags int32) {
	getProc("cabi_ImGui__SetNextItemShortcut").Call(uintptr(Key_chord), uintptr(Flags))
}

func (i *Imgui) CabiImGuiSetItemKeyOwner(Key int32) bool {
	r1, _, _ := getProc("cabi_ImGui__SetItemKeyOwner").Call(uintptr(Key))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsMouseDown(Button int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseDown").Call(uintptr(Button))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsMouseClicked(Button int32, Repeat bool) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseClicked").Call(uintptr(Button), func() uintptr {
		if Repeat {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsMouseReleased(Button int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseReleased").Call(uintptr(Button))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsMouseDoubleClicked(Button int32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseDoubleClicked").Call(uintptr(Button))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsMouseReleasedWithDelay(Button int32, Delay float32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseReleasedWithDelay").Call(uintptr(Button), uintptr(Delay))
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetMouseClickedCount(Button int32) int32 {
	r1, _, _ := getProc("cabi_ImGui__GetMouseClickedCount").Call(uintptr(Button))
	return int32(r1)
}

func (i *Imgui) CabiImGuiIsMouseHoveringRect(R_min *ImVec2, R_max *ImVec2, Clip bool) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseHoveringRect").Call(uintptr(unsafe.Pointer(R_min)), uintptr(unsafe.Pointer(R_max)), func() uintptr {
		if Clip {
			return 1
		}
		return 0
	}())
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsMousePosValid(Mouse_pos *ImVec2) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMousePosValid").Call(uintptr(unsafe.Pointer(Mouse_pos)))
	return r1 != 0
}

func (i *Imgui) CabiImGuiIsAnyMouseDown() bool {
	r1, _, _ := getProc("cabi_ImGui__IsAnyMouseDown").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetMousePos() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetMousePos").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetMousePosOnOpeningCurrentPopup() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetMousePosOnOpeningCurrentPopup").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiIsMouseDragging(Button int32, Lock_threshold float32) bool {
	r1, _, _ := getProc("cabi_ImGui__IsMouseDragging").Call(uintptr(Button), uintptr(Lock_threshold))
	return r1 != 0
}

func (i *Imgui) CabiImGuiGetMouseDragDelta(Button int32, Lock_threshold float32) *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetMouseDragDelta").Call(uintptr(Button), uintptr(Lock_threshold))
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiResetMouseDragDelta(Button int32) {
	getProc("cabi_ImGui__ResetMouseDragDelta").Call(uintptr(Button))
}

func (i *Imgui) CabiImGuiGetMouseCursor() int32 {
	r1, _, _ := getProc("cabi_ImGui__GetMouseCursor").Call()
	return int32(r1)
}

func (i *Imgui) CabiImGuiSetMouseCursor(Cursor_type int32) {
	getProc("cabi_ImGui__SetMouseCursor").Call(uintptr(Cursor_type))
}

func (i *Imgui) CabiImGuiSetNextFrameWantCaptureMouse(Want_capture_mouse bool) {
	getProc("cabi_ImGui__SetNextFrameWantCaptureMouse").Call(func() uintptr {
		if Want_capture_mouse {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiGetClipboardText() *int8 {
	r1, _, _ := getProc("cabi_ImGui__GetClipboardText").Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiSetClipboardText(Text *int8) {
	getProc("cabi_ImGui__SetClipboardText").Call(uintptr(unsafe.Pointer(Text)))
}

func (i *Imgui) CabiImGuiLoadIniSettingsFromDisk(Ini_filename *int8) {
	getProc("cabi_ImGui__LoadIniSettingsFromDisk").Call(uintptr(unsafe.Pointer(Ini_filename)))
}

func (i *Imgui) CabiImGuiLoadIniSettingsFromMemory(Ini_data *int8, Ini_size uintptr) {
	getProc("cabi_ImGui__LoadIniSettingsFromMemory").Call(uintptr(unsafe.Pointer(Ini_data)), Ini_size)
}

func (i *Imgui) CabiImGuiSaveIniSettingsToDisk(Ini_filename *int8) {
	getProc("cabi_ImGui__SaveIniSettingsToDisk").Call(uintptr(unsafe.Pointer(Ini_filename)))
}

func (i *Imgui) CabiImGuiSaveIniSettingsToMemory(Out_ini_size *uintptr) *int8 {
	r1, _, _ := getProc("cabi_ImGui__SaveIniSettingsToMemory").Call(uintptr(unsafe.Pointer(Out_ini_size)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiDebugTextEncoding(Text *int8) {
	getProc("cabi_ImGui__DebugTextEncoding").Call(uintptr(unsafe.Pointer(Text)))
}

func (i *Imgui) CabiImGuiDebugFlashStyleColor(Idx int32) {
	getProc("cabi_ImGui__DebugFlashStyleColor").Call(uintptr(Idx))
}

func (i *Imgui) CabiImGuiDebugStartItemPicker() {
	getProc("cabi_ImGui__DebugStartItemPicker").Call()
}

func (i *Imgui) CabiImGuiDebugCheckVersionAndDataLayout(Version_str *int8, Sz_io uintptr, Sz_style uintptr, Sz_vec2 uintptr, Sz_vec4 uintptr, Sz_drawvert uintptr, Sz_drawidx uintptr) bool {
	r1, _, _ := getProc("cabi_ImGui__DebugCheckVersionAndDataLayout").Call(uintptr(unsafe.Pointer(Version_str)), Sz_io, Sz_style, Sz_vec2, Sz_vec4, Sz_drawvert, Sz_drawidx)
	return r1 != 0
}

func (i *Imgui) CabiImGuiDebugLogV(Fmt *int8, Args unsafe.Pointer) {
	getProc("cabi_ImGui__DebugLogV").Call(uintptr(unsafe.Pointer(Fmt)), uintptr(Args))
}

func (i *Imgui) CabiImGuiSetAllocatorFunctions(Alloc_func unsafe.Pointer, Free_func unsafe.Pointer, User_data unsafe.Pointer) {
	getProc("cabi_ImGui__SetAllocatorFunctions").Call(uintptr(Alloc_func), uintptr(Free_func), uintptr(User_data))
}

func (i *Imgui) CabiImGuiGetAllocatorFunctions(P_alloc_func *unsafe.Pointer, P_free_func *unsafe.Pointer, P_user_data *unsafe.Pointer) {
	getProc("cabi_ImGui__GetAllocatorFunctions").Call(uintptr(unsafe.Pointer(P_alloc_func)), uintptr(unsafe.Pointer(P_free_func)), uintptr(unsafe.Pointer(P_user_data)))
}

func (i *Imgui) CabiImGuiMemAlloc(Size uintptr) unsafe.Pointer {
	r1, _, _ := getProc("cabi_ImGui__MemAlloc").Call(Size)
	return unsafe.Pointer(r1)
}

func (i *Imgui) CabiImGuiMemFree(Ptr unsafe.Pointer) {
	getProc("cabi_ImGui__MemFree").Call(uintptr(Ptr))
}

func (i *Imgui) CabiImGuiPushFont2(Font *ImFont) {
	getProc("cabi_ImGui__PushFont_2").Call(uintptr(unsafe.Pointer(Font)))
}

func (i *Imgui) CabiImGuiSetWindowFontScale(Scale float32) {
	getProc("cabi_ImGui__SetWindowFontScale").Call(uintptr(Scale))
}

func (i *Imgui) CabiImGuiImage2(Tex_ref *ImTextureRef, Image_size *ImVec2, Uv0 *ImVec2, Uv1 *ImVec2, Tint_col *ImVec4, Border_col *ImVec4) {
	getProc("cabi_ImGui__Image_2").Call(uintptr(unsafe.Pointer(Tex_ref)), uintptr(unsafe.Pointer(Image_size)), uintptr(unsafe.Pointer(Uv0)), uintptr(unsafe.Pointer(Uv1)), uintptr(unsafe.Pointer(Tint_col)), uintptr(unsafe.Pointer(Border_col)))
}

func (i *Imgui) CabiImGuiPushButtonRepeat(Repeat bool) {
	getProc("cabi_ImGui__PushButtonRepeat").Call(func() uintptr {
		if Repeat {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiPopButtonRepeat() {
	getProc("cabi_ImGui__PopButtonRepeat").Call()
}

func (i *Imgui) CabiImGuiPushTabStop(Tab_stop bool) {
	getProc("cabi_ImGui__PushTabStop").Call(func() uintptr {
		if Tab_stop {
			return 1
		}
		return 0
	}())
}

func (i *Imgui) CabiImGuiPopTabStop() {
	getProc("cabi_ImGui__PopTabStop").Call()
}

func (i *Imgui) CabiImGuiGetContentRegionMax() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetContentRegionMax").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetWindowContentRegionMin() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetWindowContentRegionMin").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiGetWindowContentRegionMax() *ImVec2 {
	r1, _, _ := getProc("cabi_ImGui__GetWindowContentRegionMax").Call()
	return (*ImVec2)(unsafe.Pointer(r1))
}

func (i *Imgui) CabiImGuiImplWin32Init(Hwnd unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui_ImplWin32_Init").Call(uintptr(Hwnd))
	return r1 != 0
}

func (i *Imgui) CabiImGuiImplWin32InitForOpenGL(Hwnd unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui_ImplWin32_InitForOpenGL").Call(uintptr(Hwnd))
	return r1 != 0
}

func (i *Imgui) CabiImGuiImplWin32Shutdown() {
	getProc("cabi_ImGui_ImplWin32_Shutdown").Call()
}

func (i *Imgui) CabiImGuiImplWin32NewFrame() {
	getProc("cabi_ImGui_ImplWin32_NewFrame").Call()
}

func (i *Imgui) CabiImGuiImplWin32EnableDpiAwareness() {
	getProc("cabi_ImGui_ImplWin32_EnableDpiAwareness").Call()
}

func (i *Imgui) CabiImGuiImplWin32GetDpiScaleForHwnd(Hwnd unsafe.Pointer) float32 {
	r1, _, _ := getProc("cabi_ImGui_ImplWin32_GetDpiScaleForHwnd").Call(uintptr(Hwnd))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiImplWin32GetDpiScaleForMonitor(Monitor unsafe.Pointer) float32 {
	r1, _, _ := getProc("cabi_ImGui_ImplWin32_GetDpiScaleForMonitor").Call(uintptr(Monitor))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (i *Imgui) CabiImGuiImplWin32EnableAlphaCompositing(Hwnd unsafe.Pointer) {
	getProc("cabi_ImGui_ImplWin32_EnableAlphaCompositing").Call(uintptr(Hwnd))
}

func (i *Imgui) CabiImGuiImplDX11Init(Device unsafe.Pointer, Device_context unsafe.Pointer) bool {
	r1, _, _ := getProc("cabi_ImGui_ImplDX11_Init").Call(uintptr(Device), uintptr(Device_context))
	return r1 != 0
}

func (i *Imgui) CabiImGuiImplDX11Shutdown() {
	getProc("cabi_ImGui_ImplDX11_Shutdown").Call()
}

func (i *Imgui) CabiImGuiImplDX11NewFrame() {
	getProc("cabi_ImGui_ImplDX11_NewFrame").Call()
}

func (i *Imgui) CabiImGuiImplDX11RenderDrawData(Draw_data *ImDrawData) {
	getProc("cabi_ImGui_ImplDX11_RenderDrawData").Call(uintptr(unsafe.Pointer(Draw_data)))
}

func (i *Imgui) CabiImGuiImplDX11CreateDeviceObjects() bool {
	r1, _, _ := getProc("cabi_ImGui_ImplDX11_CreateDeviceObjects").Call()
	return r1 != 0
}

func (i *Imgui) CabiImGuiImplDX11InvalidateDeviceObjects() {
	getProc("cabi_ImGui_ImplDX11_InvalidateDeviceObjects").Call()
}

func (i *Imgui) CabiImGuiImplDX11UpdateTexture(Tex *ImTextureData) {
	getProc("cabi_ImGui_ImplDX11_UpdateTexture").Call(uintptr(unsafe.Pointer(Tex)))
}

func (i *Imgui) CabiImGuiImplWin32WndProcHandler(HWnd unsafe.Pointer, Msg int32, WParam uintptr, LParam int) int {
	r1, _, _ := getProc("cabi_ImGui_ImplWin32_WndProcHandler").Call(uintptr(HWnd), uintptr(Msg), WParam, uintptr(LParam))
	return *(*int)(unsafe.Pointer(&r1))
}
