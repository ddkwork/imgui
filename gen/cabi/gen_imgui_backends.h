#pragma once
#ifndef MIQT_IMGUI_GEN_IMGUI_BACKENDS_H
#define MIQT_IMGUI_GEN_IMGUI_BACKENDS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "miqt_export.h"

#ifdef __cplusplus
class ID3D11Buffer;
class ID3D11Device;
class ID3D11DeviceContext;
class ImDrawChannel;
class ImDrawCmd;
class ImDrawCmdHeader;
class ImDrawData;
class ImDrawList;
class ImDrawListSharedData;
class ImDrawListSplitter;
class ImDrawVert;
class ImFont;
class ImFontAtlas;
class ImFontAtlasBuilder;
class ImFontAtlasRect;
class ImFontBaked;
class ImFontConfig;
class ImFontGlyph;
class ImFontGlyphRangesBuilder;
class ImFontLoader;
class ImGuiContext;
class ImGuiIO;
class ImGuiInputTextCallbackData;
class ImGuiKeyData;
class ImGuiListClipper;
class ImGuiMultiSelectIO;
class ImGuiOnceUponAFrame;
class ImGuiPayload;
class ImGuiPlatformIO;
class ImGuiPlatformImeData;
class ImGuiSelectionBasicStorage;
class ImGuiSelectionExternalStorage;
class ImGuiSelectionRequest;
class ImGuiSizeCallbackData;
class ImGuiStorage;
class ImGuiStoragePair;
class ImGuiStyle;
class ImGuiTableColumnSortSpecs;
class ImGuiTableSortSpecs;
class ImGuiTextBuffer;
class ImGuiTextFilter;
class ImGuiViewport;
class ImGui_ImplDX11_RenderState;
class ImNewWrapper;
class ImTextureData;
class ImTextureRect;
class ImTextureRef;
#else
typedef struct ID3D11Buffer ID3D11Buffer;
typedef struct ID3D11Device ID3D11Device;
typedef struct ID3D11DeviceContext ID3D11DeviceContext;
typedef struct ImDrawChannel ImDrawChannel;
typedef struct ImDrawCmd ImDrawCmd;
typedef struct ImDrawCmdHeader ImDrawCmdHeader;
typedef struct ImDrawData ImDrawData;
typedef struct ImDrawList ImDrawList;
typedef struct ImDrawListSharedData ImDrawListSharedData;
typedef struct ImDrawListSplitter ImDrawListSplitter;
typedef struct ImDrawVert ImDrawVert;
typedef struct ImFont ImFont;
typedef struct ImFontAtlas ImFontAtlas;
typedef struct ImFontAtlasBuilder ImFontAtlasBuilder;
typedef struct ImFontAtlasRect ImFontAtlasRect;
typedef struct ImFontBaked ImFontBaked;
typedef struct ImFontConfig ImFontConfig;
typedef struct ImFontGlyph ImFontGlyph;
typedef struct ImFontGlyphRangesBuilder ImFontGlyphRangesBuilder;
typedef struct ImFontLoader ImFontLoader;
typedef struct ImGuiContext ImGuiContext;
typedef struct ImGuiIO ImGuiIO;
typedef struct ImGuiInputTextCallbackData ImGuiInputTextCallbackData;
typedef struct ImGuiKeyData ImGuiKeyData;
typedef struct ImGuiListClipper ImGuiListClipper;
typedef struct ImGuiMultiSelectIO ImGuiMultiSelectIO;
typedef struct ImGuiOnceUponAFrame ImGuiOnceUponAFrame;
typedef struct ImGuiPayload ImGuiPayload;
typedef struct ImGuiPlatformIO ImGuiPlatformIO;
typedef struct ImGuiPlatformImeData ImGuiPlatformImeData;
typedef struct ImGuiSelectionBasicStorage ImGuiSelectionBasicStorage;
typedef struct ImGuiSelectionExternalStorage ImGuiSelectionExternalStorage;
typedef struct ImGuiSelectionRequest ImGuiSelectionRequest;
typedef struct ImGuiSizeCallbackData ImGuiSizeCallbackData;
typedef struct ImGuiStorage ImGuiStorage;
typedef struct ImGuiStoragePair ImGuiStoragePair;
typedef struct ImGuiStyle ImGuiStyle;
typedef struct ImGuiTableColumnSortSpecs ImGuiTableColumnSortSpecs;
typedef struct ImGuiTableSortSpecs ImGuiTableSortSpecs;
typedef struct ImGuiTextBuffer ImGuiTextBuffer;
typedef struct ImGuiTextFilter ImGuiTextFilter;
typedef struct ImGuiViewport ImGuiViewport;
typedef struct ImGui_ImplDX11_RenderState ImGui_ImplDX11_RenderState;
typedef struct ImNewWrapper ImNewWrapper;
typedef struct ImTextureData ImTextureData;
typedef struct ImTextureRect ImTextureRect;
typedef struct ImTextureRef ImTextureRef;
#endif
#ifdef __cplusplus
extern "C" {
#endif

#ifndef IMGUI_VERSION_NUM

// 来源: 枚举 (ImGuiDir)
typedef enum {
	ImGuiDir_None = -1,
	ImGuiDir_Left = 0,
	ImGuiDir_Right = 1,
	ImGuiDir_Up = 2,
	ImGuiDir_Down = 3,
	ImGuiDir_COUNT = 4,
} ImGuiDir;

// 来源: 枚举 (ImGuiKey)
typedef enum {
	ImGuiKey_None = 0,
	ImGuiKey_NamedKey_BEGIN = 512,
	ImGuiKey_Tab = 512,
	ImGuiKey_LeftArrow = 513,
	ImGuiKey_RightArrow = 514,
	ImGuiKey_UpArrow = 515,
	ImGuiKey_DownArrow = 516,
	ImGuiKey_PageUp = 517,
	ImGuiKey_PageDown = 518,
	ImGuiKey_Home = 519,
	ImGuiKey_End = 520,
	ImGuiKey_Insert = 521,
	ImGuiKey_Delete = 522,
	ImGuiKey_Backspace = 523,
	ImGuiKey_Space = 524,
	ImGuiKey_Enter = 525,
	ImGuiKey_Escape = 526,
	ImGuiKey_LeftCtrl = 527,
	ImGuiKey_LeftShift = 528,
	ImGuiKey_LeftAlt = 529,
	ImGuiKey_LeftSuper = 530,
	ImGuiKey_RightCtrl = 531,
	ImGuiKey_RightShift = 532,
	ImGuiKey_RightAlt = 533,
	ImGuiKey_RightSuper = 534,
	ImGuiKey_Menu = 535,
	ImGuiKey_0 = 536,
	ImGuiKey_1 = 537,
	ImGuiKey_2 = 538,
	ImGuiKey_3 = 539,
	ImGuiKey_4 = 540,
	ImGuiKey_5 = 541,
	ImGuiKey_6 = 542,
	ImGuiKey_7 = 543,
	ImGuiKey_8 = 544,
	ImGuiKey_9 = 545,
	ImGuiKey_A = 546,
	ImGuiKey_B = 547,
	ImGuiKey_C = 548,
	ImGuiKey_D = 549,
	ImGuiKey_E = 550,
	ImGuiKey_F = 551,
	ImGuiKey_G = 552,
	ImGuiKey_H = 553,
	ImGuiKey_I = 554,
	ImGuiKey_J = 555,
	ImGuiKey_K = 556,
	ImGuiKey_L = 557,
	ImGuiKey_M = 558,
	ImGuiKey_N = 559,
	ImGuiKey_O = 560,
	ImGuiKey_P = 561,
	ImGuiKey_Q = 562,
	ImGuiKey_R = 563,
	ImGuiKey_S = 564,
	ImGuiKey_T = 565,
	ImGuiKey_U = 566,
	ImGuiKey_V = 567,
	ImGuiKey_W = 568,
	ImGuiKey_X = 569,
	ImGuiKey_Y = 570,
	ImGuiKey_Z = 571,
	ImGuiKey_F1 = 572,
	ImGuiKey_F2 = 573,
	ImGuiKey_F3 = 574,
	ImGuiKey_F4 = 575,
	ImGuiKey_F5 = 576,
	ImGuiKey_F6 = 577,
	ImGuiKey_F7 = 578,
	ImGuiKey_F8 = 579,
	ImGuiKey_F9 = 580,
	ImGuiKey_F10 = 581,
	ImGuiKey_F11 = 582,
	ImGuiKey_F12 = 583,
	ImGuiKey_F13 = 584,
	ImGuiKey_F14 = 585,
	ImGuiKey_F15 = 586,
	ImGuiKey_F16 = 587,
	ImGuiKey_F17 = 588,
	ImGuiKey_F18 = 589,
	ImGuiKey_F19 = 590,
	ImGuiKey_F20 = 591,
	ImGuiKey_F21 = 592,
	ImGuiKey_F22 = 593,
	ImGuiKey_F23 = 594,
	ImGuiKey_F24 = 595,
	ImGuiKey_Apostrophe = 596,
	ImGuiKey_Comma = 597,
	ImGuiKey_Minus = 598,
	ImGuiKey_Period = 599,
	ImGuiKey_Slash = 600,
	ImGuiKey_Semicolon = 601,
	ImGuiKey_Equal = 602,
	ImGuiKey_LeftBracket = 603,
	ImGuiKey_Backslash = 604,
	ImGuiKey_RightBracket = 605,
	ImGuiKey_GraveAccent = 606,
	ImGuiKey_CapsLock = 607,
	ImGuiKey_ScrollLock = 608,
	ImGuiKey_NumLock = 609,
	ImGuiKey_PrintScreen = 610,
	ImGuiKey_Pause = 611,
	ImGuiKey_Keypad0 = 612,
	ImGuiKey_Keypad1 = 613,
	ImGuiKey_Keypad2 = 614,
	ImGuiKey_Keypad3 = 615,
	ImGuiKey_Keypad4 = 616,
	ImGuiKey_Keypad5 = 617,
	ImGuiKey_Keypad6 = 618,
	ImGuiKey_Keypad7 = 619,
	ImGuiKey_Keypad8 = 620,
	ImGuiKey_Keypad9 = 621,
	ImGuiKey_KeypadDecimal = 622,
	ImGuiKey_KeypadDivide = 623,
	ImGuiKey_KeypadMultiply = 624,
	ImGuiKey_KeypadSubtract = 625,
	ImGuiKey_KeypadAdd = 626,
	ImGuiKey_KeypadEnter = 627,
	ImGuiKey_KeypadEqual = 628,
	ImGuiKey_AppBack = 629,
	ImGuiKey_AppForward = 630,
	ImGuiKey_Oem102 = 631,
	ImGuiKey_GamepadStart = 632,
	ImGuiKey_GamepadBack = 633,
	ImGuiKey_GamepadFaceLeft = 634,
	ImGuiKey_GamepadFaceRight = 635,
	ImGuiKey_GamepadFaceUp = 636,
	ImGuiKey_GamepadFaceDown = 637,
	ImGuiKey_GamepadDpadLeft = 638,
	ImGuiKey_GamepadDpadRight = 639,
	ImGuiKey_GamepadDpadUp = 640,
	ImGuiKey_GamepadDpadDown = 641,
	ImGuiKey_GamepadL1 = 642,
	ImGuiKey_GamepadR1 = 643,
	ImGuiKey_GamepadL2 = 644,
	ImGuiKey_GamepadR2 = 645,
	ImGuiKey_GamepadL3 = 646,
	ImGuiKey_GamepadR3 = 647,
	ImGuiKey_GamepadLStickLeft = 648,
	ImGuiKey_GamepadLStickRight = 649,
	ImGuiKey_GamepadLStickUp = 650,
	ImGuiKey_GamepadLStickDown = 651,
	ImGuiKey_GamepadRStickLeft = 652,
	ImGuiKey_GamepadRStickRight = 653,
	ImGuiKey_GamepadRStickUp = 654,
	ImGuiKey_GamepadRStickDown = 655,
	ImGuiKey_MouseLeft = 656,
	ImGuiKey_MouseRight = 657,
	ImGuiKey_MouseMiddle = 658,
	ImGuiKey_MouseX1 = 659,
	ImGuiKey_MouseX2 = 660,
	ImGuiKey_MouseWheelX = 661,
	ImGuiKey_MouseWheelY = 662,
	ImGuiKey_ReservedForModCtrl = 663,
	ImGuiKey_ReservedForModShift = 664,
	ImGuiKey_ReservedForModAlt = 665,
	ImGuiKey_ReservedForModSuper = 666,
	ImGuiKey_NamedKey_END = 667,
	ImGuiKey_NamedKey_COUNT = 155,
	ImGuiMod_None = 0,
	ImGuiMod_Ctrl = 4096,
	ImGuiMod_Shift = 8192,
	ImGuiMod_Alt = 16384,
	ImGuiMod_Super = 32768,
	ImGuiMod_Mask_ = 61440,
	ImGuiKey_COUNT = 667,
	ImGuiMod_Shortcut = 4096,
} ImGuiKey;

// 来源: 枚举 (ImGuiMouseSource)
typedef enum {
	ImGuiMouseSource_Mouse = 0,
	ImGuiMouseSource_TouchScreen = 1,
	ImGuiMouseSource_Pen = 2,
	ImGuiMouseSource_COUNT = 3,
} ImGuiMouseSource;

// 来源: 枚举 (ImGuiSortDirection)
typedef enum {
	ImGuiSortDirection_None = 0,
	ImGuiSortDirection_Ascending = 1,
	ImGuiSortDirection_Descending = 2,
} ImGuiSortDirection;

// 来源: 枚举 (ImGuiWindowFlags_)
typedef enum {
	ImGuiWindowFlags_None = 0,
	ImGuiWindowFlags_NoTitleBar = 1,
	ImGuiWindowFlags_NoResize = 2,
	ImGuiWindowFlags_NoMove = 4,
	ImGuiWindowFlags_NoScrollbar = 8,
	ImGuiWindowFlags_NoScrollWithMouse = 16,
	ImGuiWindowFlags_NoCollapse = 32,
	ImGuiWindowFlags_AlwaysAutoResize = 64,
	ImGuiWindowFlags_NoBackground = 128,
	ImGuiWindowFlags_NoSavedSettings = 256,
	ImGuiWindowFlags_NoMouseInputs = 512,
	ImGuiWindowFlags_MenuBar = 1024,
	ImGuiWindowFlags_HorizontalScrollbar = 2048,
	ImGuiWindowFlags_NoFocusOnAppearing = 4096,
	ImGuiWindowFlags_NoBringToFrontOnFocus = 8192,
	ImGuiWindowFlags_AlwaysVerticalScrollbar = 16384,
	ImGuiWindowFlags_AlwaysHorizontalScrollbar = 32768,
	ImGuiWindowFlags_NoNavInputs = 65536,
	ImGuiWindowFlags_NoNavFocus = 131072,
	ImGuiWindowFlags_UnsavedDocument = 262144,
	ImGuiWindowFlags_NoNav = 196608,
	ImGuiWindowFlags_NoDecoration = 43,
	ImGuiWindowFlags_NoInputs = 197120,
	ImGuiWindowFlags_ChildWindow = 16777216,
	ImGuiWindowFlags_Tooltip = 33554432,
	ImGuiWindowFlags_Popup = 67108864,
	ImGuiWindowFlags_Modal = 134217728,
	ImGuiWindowFlags_ChildMenu = 268435456,
} ImGuiWindowFlags_;

// 来源: 枚举 (ImGuiChildFlags_)
typedef enum {
	ImGuiChildFlags_None = 0,
	ImGuiChildFlags_Borders = 1,
	ImGuiChildFlags_AlwaysUseWindowPadding = 2,
	ImGuiChildFlags_ResizeX = 4,
	ImGuiChildFlags_ResizeY = 8,
	ImGuiChildFlags_AutoResizeX = 16,
	ImGuiChildFlags_AutoResizeY = 32,
	ImGuiChildFlags_AlwaysAutoResize = 64,
	ImGuiChildFlags_FrameStyle = 128,
	ImGuiChildFlags_NavFlattened = 256,
} ImGuiChildFlags_;

// 来源: 枚举 (ImGuiItemFlags_)
typedef enum {
	ImGuiItemFlags_None = 0,
	ImGuiItemFlags_NoTabStop = 1,
	ImGuiItemFlags_NoNav = 2,
	ImGuiItemFlags_NoNavDefaultFocus = 4,
	ImGuiItemFlags_ButtonRepeat = 8,
	ImGuiItemFlags_AutoClosePopups = 16,
	ImGuiItemFlags_AllowDuplicateId = 32,
	ImGuiItemFlags_Disabled = 64,
} ImGuiItemFlags_;

// 来源: 枚举 (ImGuiInputTextFlags_)
typedef enum {
	ImGuiInputTextFlags_None = 0,
	ImGuiInputTextFlags_CharsDecimal = 1,
	ImGuiInputTextFlags_CharsHexadecimal = 2,
	ImGuiInputTextFlags_CharsScientific = 4,
	ImGuiInputTextFlags_CharsUppercase = 8,
	ImGuiInputTextFlags_CharsNoBlank = 16,
	ImGuiInputTextFlags_AllowTabInput = 32,
	ImGuiInputTextFlags_EnterReturnsTrue = 64,
	ImGuiInputTextFlags_EscapeClearsAll = 128,
	ImGuiInputTextFlags_CtrlEnterForNewLine = 256,
	ImGuiInputTextFlags_ReadOnly = 512,
	ImGuiInputTextFlags_Password = 1024,
	ImGuiInputTextFlags_AlwaysOverwrite = 2048,
	ImGuiInputTextFlags_AutoSelectAll = 4096,
	ImGuiInputTextFlags_ParseEmptyRefVal = 8192,
	ImGuiInputTextFlags_DisplayEmptyRefVal = 16384,
	ImGuiInputTextFlags_NoHorizontalScroll = 32768,
	ImGuiInputTextFlags_NoUndoRedo = 65536,
	ImGuiInputTextFlags_ElideLeft = 131072,
	ImGuiInputTextFlags_CallbackCompletion = 262144,
	ImGuiInputTextFlags_CallbackHistory = 524288,
	ImGuiInputTextFlags_CallbackAlways = 1048576,
	ImGuiInputTextFlags_CallbackCharFilter = 2097152,
	ImGuiInputTextFlags_CallbackResize = 4194304,
	ImGuiInputTextFlags_CallbackEdit = 8388608,
	ImGuiInputTextFlags_WordWrap = 16777216,
} ImGuiInputTextFlags_;

// 来源: 枚举 (ImGuiTreeNodeFlags_)
typedef enum {
	ImGuiTreeNodeFlags_None = 0,
	ImGuiTreeNodeFlags_Selected = 1,
	ImGuiTreeNodeFlags_Framed = 2,
	ImGuiTreeNodeFlags_AllowOverlap = 4,
	ImGuiTreeNodeFlags_NoTreePushOnOpen = 8,
	ImGuiTreeNodeFlags_NoAutoOpenOnLog = 16,
	ImGuiTreeNodeFlags_DefaultOpen = 32,
	ImGuiTreeNodeFlags_OpenOnDoubleClick = 64,
	ImGuiTreeNodeFlags_OpenOnArrow = 128,
	ImGuiTreeNodeFlags_Leaf = 256,
	ImGuiTreeNodeFlags_Bullet = 512,
	ImGuiTreeNodeFlags_FramePadding = 1024,
	ImGuiTreeNodeFlags_SpanAvailWidth = 2048,
	ImGuiTreeNodeFlags_SpanFullWidth = 4096,
	ImGuiTreeNodeFlags_SpanLabelWidth = 8192,
	ImGuiTreeNodeFlags_SpanAllColumns = 16384,
	ImGuiTreeNodeFlags_LabelSpanAllColumns = 32768,
	ImGuiTreeNodeFlags_NavLeftJumpsToParent = 131072,
	ImGuiTreeNodeFlags_CollapsingHeader = 26,
	ImGuiTreeNodeFlags_DrawLinesNone = 262144,
	ImGuiTreeNodeFlags_DrawLinesFull = 524288,
	ImGuiTreeNodeFlags_DrawLinesToNodes = 1048576,
	ImGuiTreeNodeFlags_NavLeftJumpsBackHere = 131072,
	ImGuiTreeNodeFlags_SpanTextWidth = 8192,
} ImGuiTreeNodeFlags_;

// 来源: 枚举 (ImGuiPopupFlags_)
typedef enum {
	ImGuiPopupFlags_None = 0,
	ImGuiPopupFlags_MouseButtonLeft = 4,
	ImGuiPopupFlags_MouseButtonRight = 8,
	ImGuiPopupFlags_MouseButtonMiddle = 12,
	ImGuiPopupFlags_NoReopen = 32,
	ImGuiPopupFlags_NoOpenOverExistingPopup = 128,
	ImGuiPopupFlags_NoOpenOverItems = 256,
	ImGuiPopupFlags_AnyPopupId = 1024,
	ImGuiPopupFlags_AnyPopupLevel = 2048,
	ImGuiPopupFlags_AnyPopup = 3072,
	ImGuiPopupFlags_MouseButtonShift_ = 2,
	ImGuiPopupFlags_MouseButtonMask_ = 12,
	ImGuiPopupFlags_InvalidMask_ = 3,
} ImGuiPopupFlags_;

// 来源: 枚举 (ImGuiSelectableFlags_)
typedef enum {
	ImGuiSelectableFlags_None = 0,
	ImGuiSelectableFlags_NoAutoClosePopups = 1,
	ImGuiSelectableFlags_SpanAllColumns = 2,
	ImGuiSelectableFlags_AllowDoubleClick = 4,
	ImGuiSelectableFlags_Disabled = 8,
	ImGuiSelectableFlags_AllowOverlap = 16,
	ImGuiSelectableFlags_Highlight = 32,
	ImGuiSelectableFlags_SelectOnNav = 64,
	ImGuiSelectableFlags_DontClosePopups = 1,
} ImGuiSelectableFlags_;

// 来源: 枚举 (ImGuiComboFlags_)
typedef enum {
	ImGuiComboFlags_None = 0,
	ImGuiComboFlags_PopupAlignLeft = 1,
	ImGuiComboFlags_HeightSmall = 2,
	ImGuiComboFlags_HeightRegular = 4,
	ImGuiComboFlags_HeightLarge = 8,
	ImGuiComboFlags_HeightLargest = 16,
	ImGuiComboFlags_NoArrowButton = 32,
	ImGuiComboFlags_NoPreview = 64,
	ImGuiComboFlags_WidthFitPreview = 128,
	ImGuiComboFlags_HeightMask_ = 30,
} ImGuiComboFlags_;

// 来源: 枚举 (ImGuiTabBarFlags_)
typedef enum {
	ImGuiTabBarFlags_None = 0,
	ImGuiTabBarFlags_Reorderable = 1,
	ImGuiTabBarFlags_AutoSelectNewTabs = 2,
	ImGuiTabBarFlags_TabListPopupButton = 4,
	ImGuiTabBarFlags_NoCloseWithMiddleMouseButton = 8,
	ImGuiTabBarFlags_NoTabListScrollingButtons = 16,
	ImGuiTabBarFlags_NoTooltip = 32,
	ImGuiTabBarFlags_DrawSelectedOverline = 64,
	ImGuiTabBarFlags_FittingPolicyMixed = 128,
	ImGuiTabBarFlags_FittingPolicyShrink = 256,
	ImGuiTabBarFlags_FittingPolicyScroll = 512,
	ImGuiTabBarFlags_FittingPolicyMask_ = 896,
	ImGuiTabBarFlags_FittingPolicyDefault_ = 128,
	ImGuiTabBarFlags_FittingPolicyResizeDown = 256,
} ImGuiTabBarFlags_;

// 来源: 枚举 (ImGuiTabItemFlags_)
typedef enum {
	ImGuiTabItemFlags_None = 0,
	ImGuiTabItemFlags_UnsavedDocument = 1,
	ImGuiTabItemFlags_SetSelected = 2,
	ImGuiTabItemFlags_NoCloseWithMiddleMouseButton = 4,
	ImGuiTabItemFlags_NoPushId = 8,
	ImGuiTabItemFlags_NoTooltip = 16,
	ImGuiTabItemFlags_NoReorder = 32,
	ImGuiTabItemFlags_Leading = 64,
	ImGuiTabItemFlags_Trailing = 128,
	ImGuiTabItemFlags_NoAssumedClosure = 256,
} ImGuiTabItemFlags_;

// 来源: 枚举 (ImGuiFocusedFlags_)
typedef enum {
	ImGuiFocusedFlags_None = 0,
	ImGuiFocusedFlags_ChildWindows = 1,
	ImGuiFocusedFlags_RootWindow = 2,
	ImGuiFocusedFlags_AnyWindow = 4,
	ImGuiFocusedFlags_NoPopupHierarchy = 8,
	ImGuiFocusedFlags_RootAndChildWindows = 3,
} ImGuiFocusedFlags_;

// 来源: 枚举 (ImGuiHoveredFlags_)
typedef enum {
	ImGuiHoveredFlags_None = 0,
	ImGuiHoveredFlags_ChildWindows = 1,
	ImGuiHoveredFlags_RootWindow = 2,
	ImGuiHoveredFlags_AnyWindow = 4,
	ImGuiHoveredFlags_NoPopupHierarchy = 8,
	ImGuiHoveredFlags_AllowWhenBlockedByPopup = 32,
	ImGuiHoveredFlags_AllowWhenBlockedByActiveItem = 128,
	ImGuiHoveredFlags_AllowWhenOverlappedByItem = 256,
	ImGuiHoveredFlags_AllowWhenOverlappedByWindow = 512,
	ImGuiHoveredFlags_AllowWhenDisabled = 1024,
	ImGuiHoveredFlags_NoNavOverride = 2048,
	ImGuiHoveredFlags_AllowWhenOverlapped = 768,
	ImGuiHoveredFlags_RectOnly = 928,
	ImGuiHoveredFlags_RootAndChildWindows = 3,
	ImGuiHoveredFlags_ForTooltip = 4096,
	ImGuiHoveredFlags_Stationary = 8192,
	ImGuiHoveredFlags_DelayNone = 16384,
	ImGuiHoveredFlags_DelayShort = 32768,
	ImGuiHoveredFlags_DelayNormal = 65536,
	ImGuiHoveredFlags_NoSharedDelay = 131072,
} ImGuiHoveredFlags_;

// 来源: 枚举 (ImGuiDragDropFlags_)
typedef enum {
	ImGuiDragDropFlags_None = 0,
	ImGuiDragDropFlags_SourceNoPreviewTooltip = 1,
	ImGuiDragDropFlags_SourceNoDisableHover = 2,
	ImGuiDragDropFlags_SourceNoHoldToOpenOthers = 4,
	ImGuiDragDropFlags_SourceAllowNullID = 8,
	ImGuiDragDropFlags_SourceExtern = 16,
	ImGuiDragDropFlags_PayloadAutoExpire = 32,
	ImGuiDragDropFlags_PayloadNoCrossContext = 64,
	ImGuiDragDropFlags_PayloadNoCrossProcess = 128,
	ImGuiDragDropFlags_AcceptBeforeDelivery = 1024,
	ImGuiDragDropFlags_AcceptNoDrawDefaultRect = 2048,
	ImGuiDragDropFlags_AcceptNoPreviewTooltip = 4096,
	ImGuiDragDropFlags_AcceptDrawAsHovered = 8192,
	ImGuiDragDropFlags_AcceptPeekOnly = 3072,
	ImGuiDragDropFlags_SourceAutoExpirePayload = 32,
} ImGuiDragDropFlags_;

// 来源: 枚举 (ImGuiDataType_)
typedef enum {
	ImGuiDataType_S8 = 0,
	ImGuiDataType_U8 = 1,
	ImGuiDataType_S16 = 2,
	ImGuiDataType_U16 = 3,
	ImGuiDataType_S32 = 4,
	ImGuiDataType_U32 = 5,
	ImGuiDataType_S64 = 6,
	ImGuiDataType_U64 = 7,
	ImGuiDataType_Float = 8,
	ImGuiDataType_Double = 9,
	ImGuiDataType_Bool = 10,
	ImGuiDataType_String = 11,
	ImGuiDataType_COUNT = 12,
} ImGuiDataType_;

// 来源: 枚举 (ImGuiInputFlags_)
typedef enum {
	ImGuiInputFlags_None = 0,
	ImGuiInputFlags_Repeat = 1,
	ImGuiInputFlags_RouteActive = 1024,
	ImGuiInputFlags_RouteFocused = 2048,
	ImGuiInputFlags_RouteGlobal = 4096,
	ImGuiInputFlags_RouteAlways = 8192,
	ImGuiInputFlags_RouteOverFocused = 16384,
	ImGuiInputFlags_RouteOverActive = 32768,
	ImGuiInputFlags_RouteUnlessBgFocused = 65536,
	ImGuiInputFlags_RouteFromRootWindow = 131072,
	ImGuiInputFlags_Tooltip = 262144,
} ImGuiInputFlags_;

// 来源: 枚举 (ImGuiConfigFlags_)
typedef enum {
	ImGuiConfigFlags_None = 0,
	ImGuiConfigFlags_NavEnableKeyboard = 1,
	ImGuiConfigFlags_NavEnableGamepad = 2,
	ImGuiConfigFlags_NoMouse = 16,
	ImGuiConfigFlags_NoMouseCursorChange = 32,
	ImGuiConfigFlags_NoKeyboard = 64,
	ImGuiConfigFlags_IsSRGB = 1048576,
	ImGuiConfigFlags_IsTouchScreen = 2097152,
	ImGuiConfigFlags_NavEnableSetMousePos = 4,
	ImGuiConfigFlags_NavNoCaptureKeyboard = 8,
} ImGuiConfigFlags_;

// 来源: 枚举 (ImGuiBackendFlags_)
typedef enum {
	ImGuiBackendFlags_None = 0,
	ImGuiBackendFlags_HasGamepad = 1,
	ImGuiBackendFlags_HasMouseCursors = 2,
	ImGuiBackendFlags_HasSetMousePos = 4,
	ImGuiBackendFlags_RendererHasVtxOffset = 8,
	ImGuiBackendFlags_RendererHasTextures = 16,
} ImGuiBackendFlags_;

// 来源: 枚举 (ImGuiCol_)
typedef enum {
	ImGuiCol_Text = 0,
	ImGuiCol_TextDisabled = 1,
	ImGuiCol_WindowBg = 2,
	ImGuiCol_ChildBg = 3,
	ImGuiCol_PopupBg = 4,
	ImGuiCol_Border = 5,
	ImGuiCol_BorderShadow = 6,
	ImGuiCol_FrameBg = 7,
	ImGuiCol_FrameBgHovered = 8,
	ImGuiCol_FrameBgActive = 9,
	ImGuiCol_TitleBg = 10,
	ImGuiCol_TitleBgActive = 11,
	ImGuiCol_TitleBgCollapsed = 12,
	ImGuiCol_MenuBarBg = 13,
	ImGuiCol_ScrollbarBg = 14,
	ImGuiCol_ScrollbarGrab = 15,
	ImGuiCol_ScrollbarGrabHovered = 16,
	ImGuiCol_ScrollbarGrabActive = 17,
	ImGuiCol_CheckMark = 18,
	ImGuiCol_CheckboxSelectedBg = 19,
	ImGuiCol_SliderGrab = 20,
	ImGuiCol_SliderGrabActive = 21,
	ImGuiCol_Button = 22,
	ImGuiCol_ButtonHovered = 23,
	ImGuiCol_ButtonActive = 24,
	ImGuiCol_Header = 25,
	ImGuiCol_HeaderHovered = 26,
	ImGuiCol_HeaderActive = 27,
	ImGuiCol_Separator = 28,
	ImGuiCol_SeparatorHovered = 29,
	ImGuiCol_SeparatorActive = 30,
	ImGuiCol_ResizeGrip = 31,
	ImGuiCol_ResizeGripHovered = 32,
	ImGuiCol_ResizeGripActive = 33,
	ImGuiCol_InputTextCursor = 34,
	ImGuiCol_TabHovered = 35,
	ImGuiCol_Tab = 36,
	ImGuiCol_TabSelected = 37,
	ImGuiCol_TabSelectedOverline = 38,
	ImGuiCol_TabDimmed = 39,
	ImGuiCol_TabDimmedSelected = 40,
	ImGuiCol_TabDimmedSelectedOverline = 41,
	ImGuiCol_PlotLines = 42,
	ImGuiCol_PlotLinesHovered = 43,
	ImGuiCol_PlotHistogram = 44,
	ImGuiCol_PlotHistogramHovered = 45,
	ImGuiCol_TableHeaderBg = 46,
	ImGuiCol_TableBorderStrong = 47,
	ImGuiCol_TableBorderLight = 48,
	ImGuiCol_TableRowBg = 49,
	ImGuiCol_TableRowBgAlt = 50,
	ImGuiCol_TextLink = 51,
	ImGuiCol_TextSelectedBg = 52,
	ImGuiCol_TreeLines = 53,
	ImGuiCol_DragDropTarget = 54,
	ImGuiCol_DragDropTargetBg = 55,
	ImGuiCol_UnsavedMarker = 56,
	ImGuiCol_NavCursor = 57,
	ImGuiCol_NavWindowingHighlight = 58,
	ImGuiCol_NavWindowingDimBg = 59,
	ImGuiCol_ModalWindowDimBg = 60,
	ImGuiCol_COUNT = 61,
	ImGuiCol_TabActive = 37,
	ImGuiCol_TabUnfocused = 39,
	ImGuiCol_TabUnfocusedActive = 40,
	ImGuiCol_NavHighlight = 57,
} ImGuiCol_;

// 来源: 枚举 (ImGuiStyleVar_)
typedef enum {
	ImGuiStyleVar_Alpha = 0,
	ImGuiStyleVar_DisabledAlpha = 1,
	ImGuiStyleVar_WindowPadding = 2,
	ImGuiStyleVar_WindowRounding = 3,
	ImGuiStyleVar_WindowBorderSize = 4,
	ImGuiStyleVar_WindowMinSize = 5,
	ImGuiStyleVar_WindowTitleAlign = 6,
	ImGuiStyleVar_ChildRounding = 7,
	ImGuiStyleVar_ChildBorderSize = 8,
	ImGuiStyleVar_PopupRounding = 9,
	ImGuiStyleVar_PopupBorderSize = 10,
	ImGuiStyleVar_FramePadding = 11,
	ImGuiStyleVar_FrameRounding = 12,
	ImGuiStyleVar_FrameBorderSize = 13,
	ImGuiStyleVar_ItemSpacing = 14,
	ImGuiStyleVar_ItemInnerSpacing = 15,
	ImGuiStyleVar_IndentSpacing = 16,
	ImGuiStyleVar_CellPadding = 17,
	ImGuiStyleVar_ScrollbarSize = 18,
	ImGuiStyleVar_ScrollbarRounding = 19,
	ImGuiStyleVar_ScrollbarPadding = 20,
	ImGuiStyleVar_GrabMinSize = 21,
	ImGuiStyleVar_GrabRounding = 22,
	ImGuiStyleVar_ImageRounding = 23,
	ImGuiStyleVar_ImageBorderSize = 24,
	ImGuiStyleVar_TabRounding = 25,
	ImGuiStyleVar_TabBorderSize = 26,
	ImGuiStyleVar_TabMinWidthBase = 27,
	ImGuiStyleVar_TabMinWidthShrink = 28,
	ImGuiStyleVar_TabBarBorderSize = 29,
	ImGuiStyleVar_TabBarOverlineSize = 30,
	ImGuiStyleVar_TableAngledHeadersAngle = 31,
	ImGuiStyleVar_TableAngledHeadersTextAlign = 32,
	ImGuiStyleVar_TreeLinesSize = 33,
	ImGuiStyleVar_TreeLinesRounding = 34,
	ImGuiStyleVar_DragDropTargetRounding = 35,
	ImGuiStyleVar_ButtonTextAlign = 36,
	ImGuiStyleVar_SelectableTextAlign = 37,
	ImGuiStyleVar_SeparatorSize = 38,
	ImGuiStyleVar_SeparatorTextBorderSize = 39,
	ImGuiStyleVar_SeparatorTextAlign = 40,
	ImGuiStyleVar_SeparatorTextPadding = 41,
	ImGuiStyleVar_COUNT = 42,
} ImGuiStyleVar_;

// 来源: 枚举 (ImGuiButtonFlags_)
typedef enum {
	ImGuiButtonFlags_None = 0,
	ImGuiButtonFlags_MouseButtonLeft = 1,
	ImGuiButtonFlags_MouseButtonRight = 2,
	ImGuiButtonFlags_MouseButtonMiddle = 4,
	ImGuiButtonFlags_MouseButtonMask_ = 7,
	ImGuiButtonFlags_EnableNav = 8,
	ImGuiButtonFlags_AllowOverlap = 4096,
} ImGuiButtonFlags_;

// 来源: 枚举 (ImGuiColorEditFlags_)
typedef enum {
	ImGuiColorEditFlags_None = 0,
	ImGuiColorEditFlags_NoAlpha = 2,
	ImGuiColorEditFlags_NoPicker = 4,
	ImGuiColorEditFlags_NoOptions = 8,
	ImGuiColorEditFlags_NoSmallPreview = 16,
	ImGuiColorEditFlags_NoInputs = 32,
	ImGuiColorEditFlags_NoTooltip = 64,
	ImGuiColorEditFlags_NoLabel = 128,
	ImGuiColorEditFlags_NoSidePreview = 256,
	ImGuiColorEditFlags_NoDragDrop = 512,
	ImGuiColorEditFlags_NoBorder = 1024,
	ImGuiColorEditFlags_NoColorMarkers = 2048,
	ImGuiColorEditFlags_AlphaOpaque = 4096,
	ImGuiColorEditFlags_AlphaNoBg = 8192,
	ImGuiColorEditFlags_AlphaPreviewHalf = 16384,
	ImGuiColorEditFlags_AlphaBar = 262144,
	ImGuiColorEditFlags_HDR = 524288,
	ImGuiColorEditFlags_DisplayRGB = 1048576,
	ImGuiColorEditFlags_DisplayHSV = 2097152,
	ImGuiColorEditFlags_DisplayHex = 4194304,
	ImGuiColorEditFlags_Uint8 = 8388608,
	ImGuiColorEditFlags_Float = 16777216,
	ImGuiColorEditFlags_PickerHueBar = 33554432,
	ImGuiColorEditFlags_PickerHueWheel = 67108864,
	ImGuiColorEditFlags_InputRGB = 134217728,
	ImGuiColorEditFlags_InputHSV = 268435456,
	ImGuiColorEditFlags_DefaultOptions_ = 177209344,
	ImGuiColorEditFlags_AlphaMask_ = 28674,
	ImGuiColorEditFlags_DisplayMask_ = 7340032,
	ImGuiColorEditFlags_DataTypeMask_ = 25165824,
	ImGuiColorEditFlags_PickerMask_ = 100663296,
	ImGuiColorEditFlags_InputMask_ = 402653184,
	ImGuiColorEditFlags_AlphaPreview = 0,
} ImGuiColorEditFlags_;

// 来源: 枚举 (ImGuiSliderFlags_)
typedef enum {
	ImGuiSliderFlags_None = 0,
	ImGuiSliderFlags_Logarithmic = 32,
	ImGuiSliderFlags_NoRoundToFormat = 64,
	ImGuiSliderFlags_NoInput = 128,
	ImGuiSliderFlags_WrapAround = 256,
	ImGuiSliderFlags_ClampOnInput = 512,
	ImGuiSliderFlags_ClampZeroRange = 1024,
	ImGuiSliderFlags_NoSpeedTweaks = 2048,
	ImGuiSliderFlags_ColorMarkers = 4096,
	ImGuiSliderFlags_AlwaysClamp = 1536,
	ImGuiSliderFlags_InvalidMask_ = 1879048207,
} ImGuiSliderFlags_;

// 来源: 枚举 (ImGuiMouseButton_)
typedef enum {
	ImGuiMouseButton_Left = 0,
	ImGuiMouseButton_Right = 1,
	ImGuiMouseButton_Middle = 2,
	ImGuiMouseButton_COUNT = 5,
} ImGuiMouseButton_;

// 来源: 枚举 (ImGuiMouseCursor_)
typedef enum {
	ImGuiMouseCursor_None = -1,
	ImGuiMouseCursor_Arrow = 0,
	ImGuiMouseCursor_TextInput = 1,
	ImGuiMouseCursor_ResizeAll = 2,
	ImGuiMouseCursor_ResizeNS = 3,
	ImGuiMouseCursor_ResizeEW = 4,
	ImGuiMouseCursor_ResizeNESW = 5,
	ImGuiMouseCursor_ResizeNWSE = 6,
	ImGuiMouseCursor_Hand = 7,
	ImGuiMouseCursor_Wait = 8,
	ImGuiMouseCursor_Progress = 9,
	ImGuiMouseCursor_NotAllowed = 10,
	ImGuiMouseCursor_COUNT = 11,
} ImGuiMouseCursor_;

// 来源: 枚举 (ImGuiCond_)
typedef enum {
	ImGuiCond_None = 0,
	ImGuiCond_Always = 1,
	ImGuiCond_Once = 2,
	ImGuiCond_FirstUseEver = 4,
	ImGuiCond_Appearing = 8,
} ImGuiCond_;

// 来源: 枚举 (ImGuiTableFlags_)
typedef enum {
	ImGuiTableFlags_None = 0,
	ImGuiTableFlags_Resizable = 1,
	ImGuiTableFlags_Reorderable = 2,
	ImGuiTableFlags_Hideable = 4,
	ImGuiTableFlags_Sortable = 8,
	ImGuiTableFlags_NoSavedSettings = 16,
	ImGuiTableFlags_ContextMenuInBody = 32,
	ImGuiTableFlags_RowBg = 64,
	ImGuiTableFlags_BordersInnerH = 128,
	ImGuiTableFlags_BordersOuterH = 256,
	ImGuiTableFlags_BordersInnerV = 512,
	ImGuiTableFlags_BordersOuterV = 1024,
	ImGuiTableFlags_BordersH = 384,
	ImGuiTableFlags_BordersV = 1536,
	ImGuiTableFlags_BordersInner = 640,
	ImGuiTableFlags_BordersOuter = 1280,
	ImGuiTableFlags_Borders = 1920,
	ImGuiTableFlags_NoBordersInBody = 2048,
	ImGuiTableFlags_NoBordersInBodyUntilResize = 4096,
	ImGuiTableFlags_SizingFixedFit = 8192,
	ImGuiTableFlags_SizingFixedSame = 16384,
	ImGuiTableFlags_SizingStretchProp = 24576,
	ImGuiTableFlags_SizingStretchSame = 32768,
	ImGuiTableFlags_NoHostExtendX = 65536,
	ImGuiTableFlags_NoHostExtendY = 131072,
	ImGuiTableFlags_NoKeepColumnsVisible = 262144,
	ImGuiTableFlags_PreciseWidths = 524288,
	ImGuiTableFlags_NoClip = 1048576,
	ImGuiTableFlags_PadOuterX = 2097152,
	ImGuiTableFlags_NoPadOuterX = 4194304,
	ImGuiTableFlags_NoPadInnerX = 8388608,
	ImGuiTableFlags_ScrollX = 16777216,
	ImGuiTableFlags_ScrollY = 33554432,
	ImGuiTableFlags_SortMulti = 67108864,
	ImGuiTableFlags_SortTristate = 134217728,
	ImGuiTableFlags_HighlightHoveredColumn = 268435456,
	ImGuiTableFlags_SizingMask_ = 57344,
} ImGuiTableFlags_;

// 来源: 枚举 (ImGuiTableColumnFlags_)
typedef enum {
	ImGuiTableColumnFlags_None = 0,
	ImGuiTableColumnFlags_Disabled = 1,
	ImGuiTableColumnFlags_DefaultHide = 2,
	ImGuiTableColumnFlags_DefaultSort = 4,
	ImGuiTableColumnFlags_WidthStretch = 8,
	ImGuiTableColumnFlags_WidthFixed = 16,
	ImGuiTableColumnFlags_NoResize = 32,
	ImGuiTableColumnFlags_NoReorder = 64,
	ImGuiTableColumnFlags_NoHide = 128,
	ImGuiTableColumnFlags_NoClip = 256,
	ImGuiTableColumnFlags_NoSort = 512,
	ImGuiTableColumnFlags_NoSortAscending = 1024,
	ImGuiTableColumnFlags_NoSortDescending = 2048,
	ImGuiTableColumnFlags_NoHeaderLabel = 4096,
	ImGuiTableColumnFlags_NoHeaderWidth = 8192,
	ImGuiTableColumnFlags_PreferSortAscending = 16384,
	ImGuiTableColumnFlags_PreferSortDescending = 32768,
	ImGuiTableColumnFlags_IndentEnable = 65536,
	ImGuiTableColumnFlags_IndentDisable = 131072,
	ImGuiTableColumnFlags_AngledHeader = 262144,
	ImGuiTableColumnFlags_IsEnabled = 16777216,
	ImGuiTableColumnFlags_IsVisible = 33554432,
	ImGuiTableColumnFlags_IsSorted = 67108864,
	ImGuiTableColumnFlags_IsHovered = 134217728,
	ImGuiTableColumnFlags_WidthMask_ = 24,
	ImGuiTableColumnFlags_IndentMask_ = 196608,
	ImGuiTableColumnFlags_StatusMask_ = 251658240,
	ImGuiTableColumnFlags_NoDirectResize_ = 1073741824,
} ImGuiTableColumnFlags_;

// 来源: 枚举 (ImGuiTableRowFlags_)
typedef enum {
	ImGuiTableRowFlags_None = 0,
	ImGuiTableRowFlags_Headers = 1,
} ImGuiTableRowFlags_;

// 来源: 枚举 (ImGuiTableBgTarget_)
typedef enum {
	ImGuiTableBgTarget_None = 0,
	ImGuiTableBgTarget_RowBg0 = 1,
	ImGuiTableBgTarget_RowBg1 = 2,
	ImGuiTableBgTarget_CellBg = 3,
} ImGuiTableBgTarget_;

// 来源: 枚举 (ImGuiListClipperFlags_)
typedef enum {
	ImGuiListClipperFlags_None = 0,
	ImGuiListClipperFlags_NoSetTableRowCounters = 1,
} ImGuiListClipperFlags_;

// 来源: 枚举 (ImGuiMultiSelectFlags_)
typedef enum {
	ImGuiMultiSelectFlags_None = 0,
	ImGuiMultiSelectFlags_SingleSelect = 1,
	ImGuiMultiSelectFlags_NoSelectAll = 2,
	ImGuiMultiSelectFlags_NoRangeSelect = 4,
	ImGuiMultiSelectFlags_NoAutoSelect = 8,
	ImGuiMultiSelectFlags_NoAutoClear = 16,
	ImGuiMultiSelectFlags_NoAutoClearOnReselect = 32,
	ImGuiMultiSelectFlags_BoxSelect1d = 64,
	ImGuiMultiSelectFlags_BoxSelect2d = 128,
	ImGuiMultiSelectFlags_BoxSelectNoScroll = 256,
	ImGuiMultiSelectFlags_ClearOnEscape = 512,
	ImGuiMultiSelectFlags_ClearOnClickVoid = 1024,
	ImGuiMultiSelectFlags_ScopeWindow = 2048,
	ImGuiMultiSelectFlags_ScopeRect = 4096,
	ImGuiMultiSelectFlags_SelectOnAuto = 8192,
	ImGuiMultiSelectFlags_SelectOnClickAlways = 16384,
	ImGuiMultiSelectFlags_SelectOnClickRelease = 32768,
	ImGuiMultiSelectFlags_NavWrapX = 65536,
	ImGuiMultiSelectFlags_NoSelectOnRightClick = 131072,
	ImGuiMultiSelectFlags_SelectOnMask_ = 57344,
	ImGuiMultiSelectFlags_SelectOnClick = 8192,
} ImGuiMultiSelectFlags_;

// 来源: 枚举 (ImGuiSelectionRequestType)
typedef enum {
	ImGuiSelectionRequestType_None = 0,
	ImGuiSelectionRequestType_SetAll = 1,
	ImGuiSelectionRequestType_SetRange = 2,
} ImGuiSelectionRequestType;

// 来源: 枚举 (ImDrawFlags_)
typedef enum {
	ImDrawFlags_None = 0,
	ImDrawFlags_RoundCornersTopLeft = 16,
	ImDrawFlags_RoundCornersTopRight = 32,
	ImDrawFlags_RoundCornersBottomLeft = 64,
	ImDrawFlags_RoundCornersBottomRight = 128,
	ImDrawFlags_RoundCornersNone = 256,
	ImDrawFlags_Closed = 512,
	ImDrawFlags_RoundCornersTop = 48,
	ImDrawFlags_RoundCornersBottom = 192,
	ImDrawFlags_RoundCornersLeft = 80,
	ImDrawFlags_RoundCornersRight = 160,
	ImDrawFlags_RoundCornersAll = 240,
	ImDrawFlags_RoundCornersDefault_ = 240,
	ImDrawFlags_RoundCornersMask_ = 496,
	ImDrawFlags_InvalidMask_ = -2147483633,
} ImDrawFlags_;

// 来源: 枚举 (ImDrawListFlags_)
typedef enum {
	ImDrawListFlags_None = 0,
	ImDrawListFlags_AntiAliasedLines = 1,
	ImDrawListFlags_AntiAliasedLinesUseTex = 2,
	ImDrawListFlags_AntiAliasedFill = 4,
	ImDrawListFlags_AllowVtxOffset = 8,
} ImDrawListFlags_;

// 来源: 枚举 (ImTextureFormat)
typedef enum {
	ImTextureFormat_RGBA32 = 0,
	ImTextureFormat_Alpha8 = 1,
} ImTextureFormat;

// 来源: 枚举 (ImTextureStatus)
typedef enum {
	ImTextureStatus_OK = 0,
	ImTextureStatus_Destroyed = 1,
	ImTextureStatus_WantCreate = 2,
	ImTextureStatus_WantUpdates = 3,
	ImTextureStatus_WantDestroy = 4,
} ImTextureStatus;

// 来源: 枚举 (ImFontAtlasFlags_)
typedef enum {
	ImFontAtlasFlags_None = 0,
	ImFontAtlasFlags_NoPowerOfTwoHeight = 1,
	ImFontAtlasFlags_NoMouseCursors = 2,
	ImFontAtlasFlags_NoBakedLines = 4,
} ImFontAtlasFlags_;

// 来源: 枚举 (ImFontFlags_)
typedef enum {
	ImFontFlags_None = 0,
	ImFontFlags_NoLoadError = 2,
	ImFontFlags_NoLoadGlyphs = 4,
	ImFontFlags_LockBakedSizes = 8,
	ImFontFlags_ImplicitRefSize = 16,
} ImFontFlags_;

// 来源: 枚举 (ImGuiViewportFlags_)
typedef enum {
	ImGuiViewportFlags_None = 0,
	ImGuiViewportFlags_IsPlatformWindow = 1,
	ImGuiViewportFlags_IsPlatformMonitor = 2,
	ImGuiViewportFlags_OwnedByApp = 4,
} ImGuiViewportFlags_;

#endif

#ifndef IMGUI_VERSION_NUM
#ifdef __cplusplus
struct ImVec2 { float x; float y; };
struct ImVec4 { float x; float y; float z; float w; };
#else
typedef struct ImVec2 { float x; float y; } ImVec2;
typedef struct ImVec4 { float x; float y; float z; float w; } ImVec4;
#endif
#endif

#ifndef IMGUI_VERSION_NUM
#ifdef __cplusplus
class ImColor;
#else
typedef struct ImColor ImColor;
#endif
#endif

#ifndef IMGUI_VERSION_NUM
typedef unsigned int ImGuiID;
typedef signed char ImS8;
typedef unsigned char ImU8;
typedef short ImS16;
typedef unsigned short ImU16;
typedef int ImS32;
typedef unsigned int ImU32;
typedef long long ImS64;
typedef unsigned long long ImU64;
typedef int ImGuiCol;
typedef int ImGuiCond;
typedef int ImGuiDataType;
typedef int ImGuiMouseButton;
typedef int ImGuiMouseCursor;
typedef int ImGuiStyleVar;
typedef int ImGuiTableBgTarget;
typedef int ImDrawFlags;
typedef int ImDrawListFlags;
typedef int ImDrawTextFlags;
typedef int ImFontFlags;
typedef int ImFontAtlasFlags;
typedef int ImGuiBackendFlags;
typedef int ImGuiButtonFlags;
typedef int ImGuiChildFlags;
typedef int ImGuiColorEditFlags;
typedef int ImGuiConfigFlags;
typedef int ImGuiComboFlags;
typedef int ImGuiDragDropFlags;
typedef int ImGuiFocusedFlags;
typedef int ImGuiHoveredFlags;
typedef int ImGuiInputFlags;
typedef int ImGuiInputTextFlags;
typedef int ImGuiItemFlags;
typedef int ImGuiKeyChord;
typedef int ImGuiListClipperFlags;
typedef int ImGuiPopupFlags;
typedef int ImGuiMultiSelectFlags;
typedef int ImGuiSelectableFlags;
typedef int ImGuiSliderFlags;
typedef int ImGuiTabBarFlags;
typedef int ImGuiTabItemFlags;
typedef int ImGuiTableFlags;
typedef int ImGuiTableColumnFlags;
typedef int ImGuiTableRowFlags;
typedef int ImGuiTreeNodeFlags;
typedef int ImGuiViewportFlags;
typedef int ImGuiWindowFlags;
typedef unsigned int ImWchar32;
typedef unsigned short ImWchar16;
typedef unsigned short ImWchar;
typedef long long ImGuiSelectionUserData;
typedef unsigned long long ImTextureID;
typedef unsigned short ImDrawIdx;
typedef int ImFontAtlasRectId;
typedef ImFontAtlasRect ImFontAtlasCustomRect;
#endif

#ifndef IMGUI_VERSION_NUM
#ifdef __cplusplus
class ImColor;
#else
typedef struct ImColor ImColor;
#endif
#endif

#ifndef MIQT_TYPES_ONLY

// 来源: 类 (ImVec2)
// 来源: 构造函数
MIQT_EXPORT ImVec2* ImVec2_new();
// 来源: 构造函数
MIQT_EXPORT ImVec2* ImVec2_new2(float _x, float _y);
// 来源: 构造函数
MIQT_EXPORT ImVec2* ImVec2_new3(ImVec2* param1);
// 来源: 类方法
MIQT_EXPORT float ImVec2_x(const ImVec2* self);
// 来源: 类方法
MIQT_EXPORT void ImVec2_setX(ImVec2* self, float x);
// 来源: 类方法
MIQT_EXPORT float ImVec2_y(const ImVec2* self);
// 来源: 类方法
MIQT_EXPORT void ImVec2_setY(ImVec2* self, float y);
// 来源: 类方法
MIQT_EXPORT float* ImVec2_operatorSubscript(ImVec2* self, unsigned long long idx);
// 来源: 类方法
MIQT_EXPORT float ImVec2_operatorSubscriptWithIdx(const ImVec2* self, unsigned long long idx);
// 来源: 类方法
MIQT_EXPORT void ImVec2_operatorAssign(ImVec2* self, ImVec2* param1);

// 来源: 析构函数
MIQT_EXPORT void ImVec2_delete(ImVec2* self);

// 来源: 类 (ImVec4)
// 来源: 构造函数
MIQT_EXPORT ImVec4* ImVec4_new();
// 来源: 构造函数
MIQT_EXPORT ImVec4* ImVec4_new2(float _x, float _y, float _z, float _w);
// 来源: 构造函数
MIQT_EXPORT ImVec4* ImVec4_new3(ImVec4* param1);
// 来源: 类方法
MIQT_EXPORT float ImVec4_x(const ImVec4* self);
// 来源: 类方法
MIQT_EXPORT void ImVec4_setX(ImVec4* self, float x);
// 来源: 类方法
MIQT_EXPORT float ImVec4_y(const ImVec4* self);
// 来源: 类方法
MIQT_EXPORT void ImVec4_setY(ImVec4* self, float y);
// 来源: 类方法
MIQT_EXPORT float ImVec4_z(const ImVec4* self);
// 来源: 类方法
MIQT_EXPORT void ImVec4_setZ(ImVec4* self, float z);
// 来源: 类方法
MIQT_EXPORT float ImVec4_w(const ImVec4* self);
// 来源: 类方法
MIQT_EXPORT void ImVec4_setW(ImVec4* self, float w);
// 来源: 类方法
MIQT_EXPORT void ImVec4_operatorAssign(ImVec4* self, ImVec4* param1);

// 来源: 析构函数
MIQT_EXPORT void ImVec4_delete(ImVec4* self);

// 来源: 类 (ImTextureRef)
// 来源: 构造函数
MIQT_EXPORT ImTextureRef* ImTextureRef_new();
// 来源: 构造函数
MIQT_EXPORT ImTextureRef* ImTextureRef_new2(unsigned long long tex_id);
// 来源: 构造函数
MIQT_EXPORT ImTextureRef* ImTextureRef_new3(void* tex_id);
// 来源: 构造函数
MIQT_EXPORT ImTextureRef* ImTextureRef_new4(ImTextureRef* param1);
// 来源: 类方法
MIQT_EXPORT unsigned long long ImTextureRef_GetTexID(const ImTextureRef* self);
// 来源: 类方法
MIQT_EXPORT ImTextureData* ImTextureRef__TexData(const ImTextureRef* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureRef_set_TexData(ImTextureRef* self, ImTextureData* _TexData);
// 来源: 类方法
MIQT_EXPORT unsigned long long ImTextureRef__TexID(const ImTextureRef* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureRef_set_TexID(ImTextureRef* self, unsigned long long _TexID);
// 来源: 类方法
MIQT_EXPORT void ImTextureRef_operatorAssign(ImTextureRef* self, ImTextureRef* param1);

// 来源: 析构函数
MIQT_EXPORT void ImTextureRef_delete(ImTextureRef* self);

// 来源: 类 (ImGuiTableSortSpecs)
// 来源: 构造函数
MIQT_EXPORT ImGuiTableSortSpecs* ImGuiTableSortSpecs_new();
// 来源: 类方法
MIQT_EXPORT ImGuiTableColumnSortSpecs* ImGuiTableSortSpecs_Specs(const ImGuiTableSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableSortSpecs_setSpecs(ImGuiTableSortSpecs* self, ImGuiTableColumnSortSpecs* Specs);
// 来源: 类方法
MIQT_EXPORT int ImGuiTableSortSpecs_SpecsCount(const ImGuiTableSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableSortSpecs_setSpecsCount(ImGuiTableSortSpecs* self, int SpecsCount);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTableSortSpecs_SpecsDirty(const ImGuiTableSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableSortSpecs_setSpecsDirty(ImGuiTableSortSpecs* self, bool SpecsDirty);

// 来源: 析构函数
MIQT_EXPORT void ImGuiTableSortSpecs_delete(ImGuiTableSortSpecs* self);

// 来源: 类 (ImGuiTableColumnSortSpecs)
// 来源: 构造函数
MIQT_EXPORT ImGuiTableColumnSortSpecs* ImGuiTableColumnSortSpecs_new();
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiTableColumnSortSpecs_ColumnUserID(const ImGuiTableColumnSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableColumnSortSpecs_setColumnUserID(ImGuiTableColumnSortSpecs* self, unsigned int ColumnUserID);
// 来源: 类方法
MIQT_EXPORT short ImGuiTableColumnSortSpecs_ColumnIndex(const ImGuiTableColumnSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableColumnSortSpecs_setColumnIndex(ImGuiTableColumnSortSpecs* self, short ColumnIndex);
// 来源: 类方法
MIQT_EXPORT short ImGuiTableColumnSortSpecs_SortOrder(const ImGuiTableColumnSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableColumnSortSpecs_setSortOrder(ImGuiTableColumnSortSpecs* self, short SortOrder);
// 来源: 类方法
MIQT_EXPORT int ImGuiTableColumnSortSpecs_SortDirection(const ImGuiTableColumnSortSpecs* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTableColumnSortSpecs_setSortDirection(ImGuiTableColumnSortSpecs* self, int SortDirection);

// 来源: 析构函数
MIQT_EXPORT void ImGuiTableColumnSortSpecs_delete(ImGuiTableColumnSortSpecs* self);

// 来源: 类 (ImNewWrapper)
// 来源: 析构函数
MIQT_EXPORT void ImNewWrapper_delete(ImNewWrapper* self);

// 来源: 类 (ImGuiStyle)
// 来源: 构造函数
MIQT_EXPORT ImGuiStyle* ImGuiStyle_new();
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_FontSizeBase(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setFontSizeBase(ImGuiStyle* self, float FontSizeBase);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_FontScaleMain(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setFontScaleMain(ImGuiStyle* self, float FontScaleMain);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_FontScaleDpi(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setFontScaleDpi(ImGuiStyle* self, float FontScaleDpi);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_Alpha(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setAlpha(ImGuiStyle* self, float Alpha);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_DisabledAlpha(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setDisabledAlpha(ImGuiStyle* self, float DisabledAlpha);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_WindowPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowPadding(ImGuiStyle* self, ImVec2* WindowPadding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_WindowRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowRounding(ImGuiStyle* self, float WindowRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_WindowBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowBorderSize(ImGuiStyle* self, float WindowBorderSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_WindowBorderHoverPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowBorderHoverPadding(ImGuiStyle* self, float WindowBorderHoverPadding);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_WindowMinSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowMinSize(ImGuiStyle* self, ImVec2* WindowMinSize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_WindowTitleAlign(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowTitleAlign(ImGuiStyle* self, ImVec2* WindowTitleAlign);
// 来源: 类方法
MIQT_EXPORT int ImGuiStyle_WindowMenuButtonPosition(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setWindowMenuButtonPosition(ImGuiStyle* self, int WindowMenuButtonPosition);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ChildRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setChildRounding(ImGuiStyle* self, float ChildRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ChildBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setChildBorderSize(ImGuiStyle* self, float ChildBorderSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_PopupRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setPopupRounding(ImGuiStyle* self, float PopupRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_PopupBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setPopupBorderSize(ImGuiStyle* self, float PopupBorderSize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_FramePadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setFramePadding(ImGuiStyle* self, ImVec2* FramePadding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_FrameRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setFrameRounding(ImGuiStyle* self, float FrameRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_FrameBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setFrameBorderSize(ImGuiStyle* self, float FrameBorderSize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_ItemSpacing(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setItemSpacing(ImGuiStyle* self, ImVec2* ItemSpacing);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_ItemInnerSpacing(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setItemInnerSpacing(ImGuiStyle* self, ImVec2* ItemInnerSpacing);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_CellPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setCellPadding(ImGuiStyle* self, ImVec2* CellPadding);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_TouchExtraPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTouchExtraPadding(ImGuiStyle* self, ImVec2* TouchExtraPadding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_IndentSpacing(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setIndentSpacing(ImGuiStyle* self, float IndentSpacing);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ColumnsMinSpacing(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setColumnsMinSpacing(ImGuiStyle* self, float ColumnsMinSpacing);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ScrollbarSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setScrollbarSize(ImGuiStyle* self, float ScrollbarSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ScrollbarRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setScrollbarRounding(ImGuiStyle* self, float ScrollbarRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ScrollbarPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setScrollbarPadding(ImGuiStyle* self, float ScrollbarPadding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_GrabMinSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setGrabMinSize(ImGuiStyle* self, float GrabMinSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_GrabRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setGrabRounding(ImGuiStyle* self, float GrabRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_LogSliderDeadzone(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setLogSliderDeadzone(ImGuiStyle* self, float LogSliderDeadzone);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ImageRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setImageRounding(ImGuiStyle* self, float ImageRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ImageBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setImageBorderSize(ImGuiStyle* self, float ImageBorderSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabRounding(ImGuiStyle* self, float TabRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabBorderSize(ImGuiStyle* self, float TabBorderSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabMinWidthBase(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabMinWidthBase(ImGuiStyle* self, float TabMinWidthBase);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabMinWidthShrink(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabMinWidthShrink(ImGuiStyle* self, float TabMinWidthShrink);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabCloseButtonMinWidthSelected(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabCloseButtonMinWidthSelected(ImGuiStyle* self, float TabCloseButtonMinWidthSelected);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabCloseButtonMinWidthUnselected(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabCloseButtonMinWidthUnselected(ImGuiStyle* self, float TabCloseButtonMinWidthUnselected);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabBarBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabBarBorderSize(ImGuiStyle* self, float TabBarBorderSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TabBarOverlineSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTabBarOverlineSize(ImGuiStyle* self, float TabBarOverlineSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TableAngledHeadersAngle(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTableAngledHeadersAngle(ImGuiStyle* self, float TableAngledHeadersAngle);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_TableAngledHeadersTextAlign(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTableAngledHeadersTextAlign(ImGuiStyle* self, ImVec2* TableAngledHeadersTextAlign);
// 来源: 类方法
MIQT_EXPORT int ImGuiStyle_TreeLinesFlags(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTreeLinesFlags(ImGuiStyle* self, int TreeLinesFlags);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TreeLinesSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTreeLinesSize(ImGuiStyle* self, float TreeLinesSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_TreeLinesRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setTreeLinesRounding(ImGuiStyle* self, float TreeLinesRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_DragDropTargetRounding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setDragDropTargetRounding(ImGuiStyle* self, float DragDropTargetRounding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_DragDropTargetBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setDragDropTargetBorderSize(ImGuiStyle* self, float DragDropTargetBorderSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_DragDropTargetPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setDragDropTargetPadding(ImGuiStyle* self, float DragDropTargetPadding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_ColorMarkerSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setColorMarkerSize(ImGuiStyle* self, float ColorMarkerSize);
// 来源: 类方法
MIQT_EXPORT int ImGuiStyle_ColorButtonPosition(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setColorButtonPosition(ImGuiStyle* self, int ColorButtonPosition);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_ButtonTextAlign(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setButtonTextAlign(ImGuiStyle* self, ImVec2* ButtonTextAlign);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_SelectableTextAlign(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setSelectableTextAlign(ImGuiStyle* self, ImVec2* SelectableTextAlign);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_SeparatorSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setSeparatorSize(ImGuiStyle* self, float SeparatorSize);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_SeparatorTextBorderSize(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setSeparatorTextBorderSize(ImGuiStyle* self, float SeparatorTextBorderSize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_SeparatorTextAlign(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setSeparatorTextAlign(ImGuiStyle* self, ImVec2* SeparatorTextAlign);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_SeparatorTextPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setSeparatorTextPadding(ImGuiStyle* self, ImVec2* SeparatorTextPadding);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_DisplayWindowPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setDisplayWindowPadding(ImGuiStyle* self, ImVec2* DisplayWindowPadding);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiStyle_DisplaySafeAreaPadding(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setDisplaySafeAreaPadding(ImGuiStyle* self, ImVec2* DisplaySafeAreaPadding);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_MouseCursorScale(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setMouseCursorScale(ImGuiStyle* self, float MouseCursorScale);
// 来源: 类方法
MIQT_EXPORT bool ImGuiStyle_AntiAliasedLines(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setAntiAliasedLines(ImGuiStyle* self, bool AntiAliasedLines);
// 来源: 类方法
MIQT_EXPORT bool ImGuiStyle_AntiAliasedLinesUseTex(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setAntiAliasedLinesUseTex(ImGuiStyle* self, bool AntiAliasedLinesUseTex);
// 来源: 类方法
MIQT_EXPORT bool ImGuiStyle_AntiAliasedFill(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setAntiAliasedFill(ImGuiStyle* self, bool AntiAliasedFill);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_CurveTessellationTol(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setCurveTessellationTol(ImGuiStyle* self, float CurveTessellationTol);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_CircleTessellationMaxError(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setCircleTessellationMaxError(ImGuiStyle* self, float CircleTessellationMaxError);
MIQT_EXPORT ImVec4* ImGuiStyle_Colors(const ImGuiStyle* self);
MIQT_EXPORT void ImGuiStyle_setColors(ImGuiStyle* self, const ImVec4* Colors);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_HoverStationaryDelay(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setHoverStationaryDelay(ImGuiStyle* self, float HoverStationaryDelay);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_HoverDelayShort(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setHoverDelayShort(ImGuiStyle* self, float HoverDelayShort);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle_HoverDelayNormal(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setHoverDelayNormal(ImGuiStyle* self, float HoverDelayNormal);
// 来源: 类方法
MIQT_EXPORT int ImGuiStyle_HoverFlagsForTooltipMouse(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setHoverFlagsForTooltipMouse(ImGuiStyle* self, int HoverFlagsForTooltipMouse);
// 来源: 类方法
MIQT_EXPORT int ImGuiStyle_HoverFlagsForTooltipNav(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_setHoverFlagsForTooltipNav(ImGuiStyle* self, int HoverFlagsForTooltipNav);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle__MainScale(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_set_MainScale(ImGuiStyle* self, float _MainScale);
// 来源: 类方法
MIQT_EXPORT float ImGuiStyle__NextFrameFontSizeBase(const ImGuiStyle* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_set_NextFrameFontSizeBase(ImGuiStyle* self, float _NextFrameFontSizeBase);
// 来源: 类方法
MIQT_EXPORT void ImGuiStyle_ScaleAllSizes(ImGuiStyle* self, float scale_factor);

// 来源: 析构函数
MIQT_EXPORT void ImGuiStyle_delete(ImGuiStyle* self);

// 来源: 类 (ImGuiKeyData)
// 来源: 构造函数
MIQT_EXPORT ImGuiKeyData* ImGuiKeyData_new(ImGuiKeyData* param1);
// 来源: 类方法
MIQT_EXPORT bool ImGuiKeyData_Down(const ImGuiKeyData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiKeyData_setDown(ImGuiKeyData* self, bool Down);
// 来源: 类方法
MIQT_EXPORT float ImGuiKeyData_DownDuration(const ImGuiKeyData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiKeyData_setDownDuration(ImGuiKeyData* self, float DownDuration);
// 来源: 类方法
MIQT_EXPORT float ImGuiKeyData_DownDurationPrev(const ImGuiKeyData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiKeyData_setDownDurationPrev(ImGuiKeyData* self, float DownDurationPrev);
// 来源: 类方法
MIQT_EXPORT float ImGuiKeyData_AnalogValue(const ImGuiKeyData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiKeyData_setAnalogValue(ImGuiKeyData* self, float AnalogValue);
// 来源: 类方法
MIQT_EXPORT void ImGuiKeyData_operatorAssign(ImGuiKeyData* self, ImGuiKeyData* param1);

// 来源: 析构函数
MIQT_EXPORT void ImGuiKeyData_delete(ImGuiKeyData* self);

// 来源: 类 (ImGuiIO)
// 来源: 构造函数
MIQT_EXPORT ImGuiIO* ImGuiIO_new();
// 来源: 构造函数
MIQT_EXPORT ImGuiIO* ImGuiIO_new2(ImGuiIO* param1);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_ConfigFlags(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigFlags(ImGuiIO* self, int ConfigFlags);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_BackendFlags(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setBackendFlags(ImGuiIO* self, int BackendFlags);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiIO_DisplaySize(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setDisplaySize(ImGuiIO* self, ImVec2* DisplaySize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiIO_DisplayFramebufferScale(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setDisplayFramebufferScale(ImGuiIO* self, ImVec2* DisplayFramebufferScale);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_DeltaTime(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setDeltaTime(ImGuiIO* self, float DeltaTime);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_IniSavingRate(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setIniSavingRate(ImGuiIO* self, float IniSavingRate);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiIO_IniFilename(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setIniFilename(ImGuiIO* self, const char* IniFilename);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiIO_LogFilename(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setLogFilename(ImGuiIO* self, const char* LogFilename);
// 来源: 类方法
MIQT_EXPORT ImFontAtlas* ImGuiIO_Fonts(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setFonts(ImGuiIO* self, ImFontAtlas* Fonts);
// 来源: 类方法
MIQT_EXPORT ImFont* ImGuiIO_FontDefault(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setFontDefault(ImGuiIO* self, ImFont* FontDefault);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_FontAllowUserScaling(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setFontAllowUserScaling(ImGuiIO* self, bool FontAllowUserScaling);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavSwapGamepadButtons(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavSwapGamepadButtons(ImGuiIO* self, bool ConfigNavSwapGamepadButtons);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavMoveSetMousePos(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavMoveSetMousePos(ImGuiIO* self, bool ConfigNavMoveSetMousePos);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavCaptureKeyboard(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavCaptureKeyboard(ImGuiIO* self, bool ConfigNavCaptureKeyboard);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavEscapeClearFocusItem(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavEscapeClearFocusItem(ImGuiIO* self, bool ConfigNavEscapeClearFocusItem);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavEscapeClearFocusWindow(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavEscapeClearFocusWindow(ImGuiIO* self, bool ConfigNavEscapeClearFocusWindow);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavCursorVisibleAuto(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavCursorVisibleAuto(ImGuiIO* self, bool ConfigNavCursorVisibleAuto);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigNavCursorVisibleAlways(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigNavCursorVisibleAlways(ImGuiIO* self, bool ConfigNavCursorVisibleAlways);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_MouseDrawCursor(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseDrawCursor(ImGuiIO* self, bool MouseDrawCursor);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigMacOSXBehaviors(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigMacOSXBehaviors(ImGuiIO* self, bool ConfigMacOSXBehaviors);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigInputTrickleEventQueue(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigInputTrickleEventQueue(ImGuiIO* self, bool ConfigInputTrickleEventQueue);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigInputTextCursorBlink(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigInputTextCursorBlink(ImGuiIO* self, bool ConfigInputTextCursorBlink);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigInputTextEnterKeepActive(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigInputTextEnterKeepActive(ImGuiIO* self, bool ConfigInputTextEnterKeepActive);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDragClickToInputText(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDragClickToInputText(ImGuiIO* self, bool ConfigDragClickToInputText);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigWindowsResizeFromEdges(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigWindowsResizeFromEdges(ImGuiIO* self, bool ConfigWindowsResizeFromEdges);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigWindowsMoveFromTitleBarOnly(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigWindowsMoveFromTitleBarOnly(ImGuiIO* self, bool ConfigWindowsMoveFromTitleBarOnly);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigWindowsCopyContentsWithCtrlC(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigWindowsCopyContentsWithCtrlC(ImGuiIO* self, bool ConfigWindowsCopyContentsWithCtrlC);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigScrollbarScrollByPage(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigScrollbarScrollByPage(ImGuiIO* self, bool ConfigScrollbarScrollByPage);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_ConfigMemoryCompactTimer(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigMemoryCompactTimer(ImGuiIO* self, float ConfigMemoryCompactTimer);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_MouseDoubleClickTime(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseDoubleClickTime(ImGuiIO* self, float MouseDoubleClickTime);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_MouseDoubleClickMaxDist(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseDoubleClickMaxDist(ImGuiIO* self, float MouseDoubleClickMaxDist);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_MouseDragThreshold(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseDragThreshold(ImGuiIO* self, float MouseDragThreshold);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_KeyRepeatDelay(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeyRepeatDelay(ImGuiIO* self, float KeyRepeatDelay);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_KeyRepeatRate(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeyRepeatRate(ImGuiIO* self, float KeyRepeatRate);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigErrorRecovery(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigErrorRecovery(ImGuiIO* self, bool ConfigErrorRecovery);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigErrorRecoveryEnableAssert(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigErrorRecoveryEnableAssert(ImGuiIO* self, bool ConfigErrorRecoveryEnableAssert);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigErrorRecoveryEnableDebugLog(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigErrorRecoveryEnableDebugLog(ImGuiIO* self, bool ConfigErrorRecoveryEnableDebugLog);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigErrorRecoveryEnableTooltip(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigErrorRecoveryEnableTooltip(ImGuiIO* self, bool ConfigErrorRecoveryEnableTooltip);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugIsDebuggerPresent(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugIsDebuggerPresent(ImGuiIO* self, bool ConfigDebugIsDebuggerPresent);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugHighlightIdConflicts(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugHighlightIdConflicts(ImGuiIO* self, bool ConfigDebugHighlightIdConflicts);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugHighlightIdConflictsShowItemPicker(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugHighlightIdConflictsShowItemPicker(ImGuiIO* self, bool ConfigDebugHighlightIdConflictsShowItemPicker);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugBeginReturnValueOnce(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugBeginReturnValueOnce(ImGuiIO* self, bool ConfigDebugBeginReturnValueOnce);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugBeginReturnValueLoop(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugBeginReturnValueLoop(ImGuiIO* self, bool ConfigDebugBeginReturnValueLoop);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugIgnoreFocusLoss(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugIgnoreFocusLoss(ImGuiIO* self, bool ConfigDebugIgnoreFocusLoss);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_ConfigDebugIniSettings(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setConfigDebugIniSettings(ImGuiIO* self, bool ConfigDebugIniSettings);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiIO_BackendPlatformName(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setBackendPlatformName(ImGuiIO* self, const char* BackendPlatformName);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiIO_BackendRendererName(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setBackendRendererName(ImGuiIO* self, const char* BackendRendererName);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddKeyEvent(ImGuiIO* self, int key, bool down);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddKeyAnalogEvent(ImGuiIO* self, int key, bool down, float v);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddMousePosEvent(ImGuiIO* self, float x, float y);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddMouseButtonEvent(ImGuiIO* self, int button, bool down);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddMouseWheelEvent(ImGuiIO* self, float wheel_x, float wheel_y);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddMouseSourceEvent(ImGuiIO* self, int source);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddFocusEvent(ImGuiIO* self, bool focused);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddInputCharacter(ImGuiIO* self, unsigned int c);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddInputCharacterUTF16(ImGuiIO* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_AddInputCharactersUTF8(ImGuiIO* self, const char* str);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_SetKeyEventNativeData(ImGuiIO* self, int key, int native_keycode, int native_scancode);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_SetAppAcceptingEvents(ImGuiIO* self, bool accepting_events);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_ClearEventsQueue(ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_ClearInputKeys(ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_ClearInputMouse(ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_WantCaptureMouse(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setWantCaptureMouse(ImGuiIO* self, bool WantCaptureMouse);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_WantCaptureKeyboard(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setWantCaptureKeyboard(ImGuiIO* self, bool WantCaptureKeyboard);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_WantTextInput(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setWantTextInput(ImGuiIO* self, bool WantTextInput);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_WantSetMousePos(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setWantSetMousePos(ImGuiIO* self, bool WantSetMousePos);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_WantSaveIniSettings(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setWantSaveIniSettings(ImGuiIO* self, bool WantSaveIniSettings);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_NavActive(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setNavActive(ImGuiIO* self, bool NavActive);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_NavVisible(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setNavVisible(ImGuiIO* self, bool NavVisible);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_Framerate(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setFramerate(ImGuiIO* self, float Framerate);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_MetricsRenderVertices(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMetricsRenderVertices(ImGuiIO* self, int MetricsRenderVertices);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_MetricsRenderIndices(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMetricsRenderIndices(ImGuiIO* self, int MetricsRenderIndices);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_MetricsRenderWindows(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMetricsRenderWindows(ImGuiIO* self, int MetricsRenderWindows);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_MetricsActiveWindows(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMetricsActiveWindows(ImGuiIO* self, int MetricsActiveWindows);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiIO_MouseDelta(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseDelta(ImGuiIO* self, ImVec2* MouseDelta);
// 来源: 类方法
MIQT_EXPORT void* ImGuiIO_Ctx(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setCtx(ImGuiIO* self, void* Ctx);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiIO_MousePos(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMousePos(ImGuiIO* self, ImVec2* MousePos);
MIQT_EXPORT bool* ImGuiIO_MouseDown(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDown(ImGuiIO* self, const bool* MouseDown);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_MouseWheel(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseWheel(ImGuiIO* self, float MouseWheel);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_MouseWheelH(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseWheelH(ImGuiIO* self, float MouseWheelH);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_MouseSource(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseSource(ImGuiIO* self, int MouseSource);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_KeyCtrl(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeyCtrl(ImGuiIO* self, bool KeyCtrl);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_KeyShift(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeyShift(ImGuiIO* self, bool KeyShift);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_KeyAlt(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeyAlt(ImGuiIO* self, bool KeyAlt);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_KeySuper(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeySuper(ImGuiIO* self, bool KeySuper);
// 来源: 类方法
MIQT_EXPORT int ImGuiIO_KeyMods(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setKeyMods(ImGuiIO* self, int KeyMods);
MIQT_EXPORT ImGuiKeyData* ImGuiIO_KeysData(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setKeysData(ImGuiIO* self, const ImGuiKeyData* KeysData);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_WantCaptureMouseUnlessPopupClose(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setWantCaptureMouseUnlessPopupClose(ImGuiIO* self, bool WantCaptureMouseUnlessPopupClose);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiIO_MousePosPrev(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMousePosPrev(ImGuiIO* self, ImVec2* MousePosPrev);
MIQT_EXPORT ImVec2* ImGuiIO_MouseClickedPos(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseClickedPos(ImGuiIO* self, const ImVec2* MouseClickedPos);
MIQT_EXPORT double* ImGuiIO_MouseClickedTime(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseClickedTime(ImGuiIO* self, const double* MouseClickedTime);
MIQT_EXPORT bool* ImGuiIO_MouseClicked(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseClicked(ImGuiIO* self, const bool* MouseClicked);
MIQT_EXPORT bool* ImGuiIO_MouseDoubleClicked(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDoubleClicked(ImGuiIO* self, const bool* MouseDoubleClicked);
MIQT_EXPORT ImU16* ImGuiIO_MouseClickedCount(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseClickedCount(ImGuiIO* self, const ImU16* MouseClickedCount);
MIQT_EXPORT ImU16* ImGuiIO_MouseClickedLastCount(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseClickedLastCount(ImGuiIO* self, const ImU16* MouseClickedLastCount);
MIQT_EXPORT bool* ImGuiIO_MouseReleased(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseReleased(ImGuiIO* self, const bool* MouseReleased);
MIQT_EXPORT double* ImGuiIO_MouseReleasedTime(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseReleasedTime(ImGuiIO* self, const double* MouseReleasedTime);
MIQT_EXPORT bool* ImGuiIO_MouseDownOwned(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDownOwned(ImGuiIO* self, const bool* MouseDownOwned);
MIQT_EXPORT bool* ImGuiIO_MouseDownOwnedUnlessPopupClose(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDownOwnedUnlessPopupClose(ImGuiIO* self, const bool* MouseDownOwnedUnlessPopupClose);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_MouseWheelRequestAxisSwap(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseWheelRequestAxisSwap(ImGuiIO* self, bool MouseWheelRequestAxisSwap);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_MouseCtrlLeftAsRightClick(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setMouseCtrlLeftAsRightClick(ImGuiIO* self, bool MouseCtrlLeftAsRightClick);
MIQT_EXPORT float* ImGuiIO_MouseDownDuration(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDownDuration(ImGuiIO* self, const float* MouseDownDuration);
MIQT_EXPORT float* ImGuiIO_MouseDownDurationPrev(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDownDurationPrev(ImGuiIO* self, const float* MouseDownDurationPrev);
MIQT_EXPORT float* ImGuiIO_MouseDragMaxDistanceSqr(const ImGuiIO* self);
MIQT_EXPORT void ImGuiIO_setMouseDragMaxDistanceSqr(ImGuiIO* self, const float* MouseDragMaxDistanceSqr);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_PenPressure(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setPenPressure(ImGuiIO* self, float PenPressure);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_AppFocusLost(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setAppFocusLost(ImGuiIO* self, bool AppFocusLost);
// 来源: 类方法
MIQT_EXPORT bool ImGuiIO_AppAcceptingEvents(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setAppAcceptingEvents(ImGuiIO* self, bool AppAcceptingEvents);
// 来源: 类方法
MIQT_EXPORT unsigned short ImGuiIO_InputQueueSurrogate(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setInputQueueSurrogate(ImGuiIO* self, unsigned short InputQueueSurrogate);
// 来源: 类方法
MIQT_EXPORT void* ImGuiIO_InputQueueCharacters(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setInputQueueCharacters(ImGuiIO* self, void* InputQueueCharacters);
// 来源: 类方法
MIQT_EXPORT float ImGuiIO_FontGlobalScale(const ImGuiIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_setFontGlobalScale(ImGuiIO* self, float FontGlobalScale);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_operatorAssign(ImGuiIO* self, ImGuiIO* param1);
// 来源: 类方法
MIQT_EXPORT void ImGuiIO_SetKeyEventNativeData2(ImGuiIO* self, int key, int native_keycode, int native_scancode, int native_legacy_index);

// 来源: 析构函数
MIQT_EXPORT void ImGuiIO_delete(ImGuiIO* self);

// 来源: 类 (ImGuiInputTextCallbackData)
// 来源: 构造函数
MIQT_EXPORT ImGuiInputTextCallbackData* ImGuiInputTextCallbackData_new();
// 来源: 类方法
MIQT_EXPORT void* ImGuiInputTextCallbackData_Ctx(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setCtx(ImGuiInputTextCallbackData* self, void* Ctx);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_EventFlag(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setEventFlag(ImGuiInputTextCallbackData* self, int EventFlag);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_Flags(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setFlags(ImGuiInputTextCallbackData* self, int Flags);
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiInputTextCallbackData_ID(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setID(ImGuiInputTextCallbackData* self, unsigned int ID);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_EventKey(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setEventKey(ImGuiInputTextCallbackData* self, int EventKey);
// 来源: 类方法
MIQT_EXPORT unsigned short ImGuiInputTextCallbackData_EventChar(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setEventChar(ImGuiInputTextCallbackData* self, unsigned short EventChar);
// 来源: 类方法
MIQT_EXPORT bool ImGuiInputTextCallbackData_EventActivated(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setEventActivated(ImGuiInputTextCallbackData* self, bool EventActivated);
// 来源: 类方法
MIQT_EXPORT bool ImGuiInputTextCallbackData_BufDirty(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setBufDirty(ImGuiInputTextCallbackData* self, bool BufDirty);
// 来源: 类方法
MIQT_EXPORT char* ImGuiInputTextCallbackData_Buf(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setBuf(ImGuiInputTextCallbackData* self, char* Buf);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_BufTextLen(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setBufTextLen(ImGuiInputTextCallbackData* self, int BufTextLen);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_BufSize(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setBufSize(ImGuiInputTextCallbackData* self, int BufSize);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_CursorPos(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setCursorPos(ImGuiInputTextCallbackData* self, int CursorPos);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_SelectionStart(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setSelectionStart(ImGuiInputTextCallbackData* self, int SelectionStart);
// 来源: 类方法
MIQT_EXPORT int ImGuiInputTextCallbackData_SelectionEnd(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_setSelectionEnd(ImGuiInputTextCallbackData* self, int SelectionEnd);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self, int pos, int bytes_count);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self, int pos, const char* text);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_SetSelection(ImGuiInputTextCallbackData* self, int s, int e);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiInputTextCallbackData_HasSelection(const ImGuiInputTextCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiInputTextCallbackData_InsertChars2(ImGuiInputTextCallbackData* self, int pos, const char* text, const char* text_end);

// 来源: 析构函数
MIQT_EXPORT void ImGuiInputTextCallbackData_delete(ImGuiInputTextCallbackData* self);

// 来源: 类 (ImGuiSizeCallbackData)
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiSizeCallbackData_Pos(const ImGuiSizeCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSizeCallbackData_setPos(ImGuiSizeCallbackData* self, ImVec2* Pos);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiSizeCallbackData_CurrentSize(const ImGuiSizeCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSizeCallbackData_setCurrentSize(ImGuiSizeCallbackData* self, ImVec2* CurrentSize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiSizeCallbackData_DesiredSize(const ImGuiSizeCallbackData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSizeCallbackData_setDesiredSize(ImGuiSizeCallbackData* self, ImVec2* DesiredSize);

// 来源: 析构函数
MIQT_EXPORT void ImGuiSizeCallbackData_delete(ImGuiSizeCallbackData* self);

// 来源: 类 (ImGuiPayload)
// 来源: 构造函数
MIQT_EXPORT ImGuiPayload* ImGuiPayload_new();
// 来源: 类方法
MIQT_EXPORT int ImGuiPayload_DataSize(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_setDataSize(ImGuiPayload* self, int DataSize);
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiPayload_SourceId(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_setSourceId(ImGuiPayload* self, unsigned int SourceId);
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiPayload_SourceParentId(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_setSourceParentId(ImGuiPayload* self, unsigned int SourceParentId);
// 来源: 类方法
MIQT_EXPORT int ImGuiPayload_DataFrameCount(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_setDataFrameCount(ImGuiPayload* self, int DataFrameCount);
// 来源: 类方法
MIQT_EXPORT bool ImGuiPayload_Preview(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_setPreview(ImGuiPayload* self, bool Preview);
// 来源: 类方法
MIQT_EXPORT bool ImGuiPayload_Delivery(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_setDelivery(ImGuiPayload* self, bool Delivery);
// 来源: 类方法
MIQT_EXPORT void ImGuiPayload_Clear(ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiPayload_IsDataType(const ImGuiPayload* self, const char* type);
// 来源: 类方法
MIQT_EXPORT bool ImGuiPayload_IsPreview(const ImGuiPayload* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiPayload_IsDelivery(const ImGuiPayload* self);

// 来源: 析构函数
MIQT_EXPORT void ImGuiPayload_delete(ImGuiPayload* self);

// 来源: 类 (ImGuiOnceUponAFrame)
// 来源: 构造函数
MIQT_EXPORT ImGuiOnceUponAFrame* ImGuiOnceUponAFrame_new();
// 来源: 类方法
MIQT_EXPORT int ImGuiOnceUponAFrame_RefFrame(const ImGuiOnceUponAFrame* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiOnceUponAFrame_setRefFrame(ImGuiOnceUponAFrame* self, int RefFrame);
// 来源: 类方法
MIQT_EXPORT bool ImGuiOnceUponAFrame_ToBool(const ImGuiOnceUponAFrame* self);

// 来源: 析构函数
MIQT_EXPORT void ImGuiOnceUponAFrame_delete(ImGuiOnceUponAFrame* self);

// 来源: 类 (ImGuiTextFilter)
// 来源: 构造函数
MIQT_EXPORT ImGuiTextFilter* ImGuiTextFilter_new();
// 来源: 构造函数
MIQT_EXPORT ImGuiTextFilter* ImGuiTextFilter_new2(ImGuiTextFilter* param1);
// 来源: 构造函数
MIQT_EXPORT ImGuiTextFilter* ImGuiTextFilter_new3(const char* default_filter);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextFilter_Draw(ImGuiTextFilter* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextFilter_PassFilter(const ImGuiTextFilter* self, const char* text);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextFilter_Build(ImGuiTextFilter* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextFilter_Clear(ImGuiTextFilter* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextFilter_IsActive(const ImGuiTextFilter* self);
// 来源: 类方法
MIQT_EXPORT int ImGuiTextFilter_CountGrep(const ImGuiTextFilter* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextFilter_setCountGrep(ImGuiTextFilter* self, int CountGrep);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextFilter_operatorAssign(ImGuiTextFilter* self, ImGuiTextFilter* param1);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextFilter_DrawWithLabel(ImGuiTextFilter* self, const char* label);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextFilter_Draw2(ImGuiTextFilter* self, const char* label, float width);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextFilter_PassFilter2(const ImGuiTextFilter* self, const char* text, const char* text_end);

// 来源: 析构函数
MIQT_EXPORT void ImGuiTextFilter_delete(ImGuiTextFilter* self);

// 来源: 类 (ImGuiTextBuffer)
// 来源: 构造函数
MIQT_EXPORT ImGuiTextBuffer* ImGuiTextBuffer_new();
// 来源: 构造函数
MIQT_EXPORT ImGuiTextBuffer* ImGuiTextBuffer_new2(ImGuiTextBuffer* param1);
// 来源: 类方法
MIQT_EXPORT void* ImGuiTextBuffer_Buf(const ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_setBuf(ImGuiTextBuffer* self, void* Buf);
// 来源: 类方法
MIQT_EXPORT char ImGuiTextBuffer_operatorSubscript(const ImGuiTextBuffer* self, int i);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiTextBuffer_begin(const ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiTextBuffer_end(const ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT int ImGuiTextBuffer_size(const ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiTextBuffer_empty(const ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_clear(ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_resize(ImGuiTextBuffer* self, int size);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_reserve(ImGuiTextBuffer* self, int capacity);
// 来源: 类方法
MIQT_EXPORT const char* ImGuiTextBuffer_cStr(const ImGuiTextBuffer* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_append(ImGuiTextBuffer* self, const char* str);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_operatorAssign(ImGuiTextBuffer* self, ImGuiTextBuffer* param1);
// 来源: 类方法
MIQT_EXPORT void ImGuiTextBuffer_append2(ImGuiTextBuffer* self, const char* str, const char* str_end);

// 来源: 析构函数
MIQT_EXPORT void ImGuiTextBuffer_delete(ImGuiTextBuffer* self);

// 来源: 类 (ImGuiStoragePair)
// 来源: 构造函数
MIQT_EXPORT ImGuiStoragePair* ImGuiStoragePair_new(unsigned int _key, int _val);
// 来源: 构造函数
MIQT_EXPORT ImGuiStoragePair* ImGuiStoragePair_new2(unsigned int _key, float _val);
// 来源: 构造函数
MIQT_EXPORT ImGuiStoragePair* ImGuiStoragePair_new3(unsigned int _key, void* _val);
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiStoragePair_key(const ImGuiStoragePair* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStoragePair_setKey(ImGuiStoragePair* self, unsigned int key);

// 来源: 析构函数
MIQT_EXPORT void ImGuiStoragePair_delete(ImGuiStoragePair* self);

// 来源: 类 (ImGuiStorage)
// 来源: 构造函数
MIQT_EXPORT ImGuiStorage* ImGuiStorage_new(ImGuiStorage* param1);
// 来源: 类方法
MIQT_EXPORT void* ImGuiStorage_Data(const ImGuiStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_setData(ImGuiStorage* self, void* Data);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_Clear(ImGuiStorage* self);
// 来源: 类方法
MIQT_EXPORT int ImGuiStorage_GetInt(const ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_SetInt(ImGuiStorage* self, unsigned int key, int val);
// 来源: 类方法
MIQT_EXPORT bool ImGuiStorage_GetBool(const ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_SetBool(ImGuiStorage* self, unsigned int key, bool val);
// 来源: 类方法
MIQT_EXPORT float ImGuiStorage_GetFloat(const ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_SetFloat(ImGuiStorage* self, unsigned int key, float val);
// 来源: 类方法
MIQT_EXPORT void* ImGuiStorage_GetVoidPtr(const ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT int* ImGuiStorage_GetIntRef(ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT bool* ImGuiStorage_GetBoolRef(ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT float* ImGuiStorage_GetFloatRef(ImGuiStorage* self, unsigned int key);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_BuildSortByKey(ImGuiStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_SetAllInt(ImGuiStorage* self, int val);
// 来源: 类方法
MIQT_EXPORT void ImGuiStorage_operatorAssign(ImGuiStorage* self, ImGuiStorage* param1);
// 来源: 类方法
MIQT_EXPORT int ImGuiStorage_GetInt2(const ImGuiStorage* self, unsigned int key, int default_val);
// 来源: 类方法
MIQT_EXPORT bool ImGuiStorage_GetBool2(const ImGuiStorage* self, unsigned int key, bool default_val);
// 来源: 类方法
MIQT_EXPORT float ImGuiStorage_GetFloat2(const ImGuiStorage* self, unsigned int key, float default_val);
// 来源: 类方法
MIQT_EXPORT int* ImGuiStorage_GetIntRef2(ImGuiStorage* self, unsigned int key, int default_val);
// 来源: 类方法
MIQT_EXPORT bool* ImGuiStorage_GetBoolRef2(ImGuiStorage* self, unsigned int key, bool default_val);
// 来源: 类方法
MIQT_EXPORT float* ImGuiStorage_GetFloatRef2(ImGuiStorage* self, unsigned int key, float default_val);

// 来源: 析构函数
MIQT_EXPORT void ImGuiStorage_delete(ImGuiStorage* self);

// 来源: 类 (ImGuiListClipper)
// 来源: 构造函数
MIQT_EXPORT ImGuiListClipper* ImGuiListClipper_new();
// 来源: 类方法
MIQT_EXPORT int ImGuiListClipper_DisplayStart(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setDisplayStart(ImGuiListClipper* self, int DisplayStart);
// 来源: 类方法
MIQT_EXPORT int ImGuiListClipper_DisplayEnd(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setDisplayEnd(ImGuiListClipper* self, int DisplayEnd);
// 来源: 类方法
MIQT_EXPORT int ImGuiListClipper_UserIndex(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setUserIndex(ImGuiListClipper* self, int UserIndex);
// 来源: 类方法
MIQT_EXPORT int ImGuiListClipper_ItemsCount(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setItemsCount(ImGuiListClipper* self, int ItemsCount);
// 来源: 类方法
MIQT_EXPORT float ImGuiListClipper_ItemsHeight(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setItemsHeight(ImGuiListClipper* self, float ItemsHeight);
// 来源: 类方法
MIQT_EXPORT int ImGuiListClipper_Flags(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setFlags(ImGuiListClipper* self, int Flags);
// 来源: 类方法
MIQT_EXPORT double ImGuiListClipper_StartPosY(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setStartPosY(ImGuiListClipper* self, double StartPosY);
// 来源: 类方法
MIQT_EXPORT double ImGuiListClipper_StartSeekOffsetY(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setStartSeekOffsetY(ImGuiListClipper* self, double StartSeekOffsetY);
// 来源: 类方法
MIQT_EXPORT void* ImGuiListClipper_Ctx(const ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_setCtx(ImGuiListClipper* self, void* Ctx);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_Begin(ImGuiListClipper* self, int items_count);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_End(ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT bool ImGuiListClipper_Step(ImGuiListClipper* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_IncludeItemByIndex(ImGuiListClipper* self, int item_index);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_IncludeItemsByIndex(ImGuiListClipper* self, int item_begin, int item_end);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_SeekCursorForItem(ImGuiListClipper* self, int item_index);
// 来源: 类方法
MIQT_EXPORT void ImGuiListClipper_Begin2(ImGuiListClipper* self, int items_count, float items_height);

// 来源: 析构函数
MIQT_EXPORT void ImGuiListClipper_delete(ImGuiListClipper* self);

// 来源: 类 (ImColor)
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new();
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new2(float r, float g, float b);
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new3(ImVec4* col);
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new4(int r, int g, int b);
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new5(unsigned int rgba);
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new6(ImColor* param1);
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new7(float r, float g, float b, float a);
// 来源: 构造函数
MIQT_EXPORT ImColor* ImColor_new8(int r, int g, int b, int a);
// 来源: 类方法
MIQT_EXPORT ImVec4* ImColor_Value(const ImColor* self);
// 来源: 类方法
MIQT_EXPORT void ImColor_setValue(ImColor* self, ImVec4* Value);
// 来源: 类方法
MIQT_EXPORT unsigned int ImColor_ToUnsignedInt(const ImColor* self);
// 来源: 类方法
MIQT_EXPORT ImVec4* ImColor_ToImVec4(const ImColor* self);
// 来源: 类方法
MIQT_EXPORT void ImColor_SetHSV(ImColor* self, float h, float s, float v);
// 来源: 类方法
MIQT_EXPORT ImColor* ImColor_HSV(float h, float s, float v);
// 来源: 类方法
MIQT_EXPORT void ImColor_SetHSV2(ImColor* self, float h, float s, float v, float a);
// 来源: 类方法
MIQT_EXPORT ImColor* ImColor_HSV2(float h, float s, float v, float a);

// 来源: 析构函数
MIQT_EXPORT void ImColor_delete(ImColor* self);

// 来源: 类 (ImGuiMultiSelectIO)
// 来源: 构造函数
MIQT_EXPORT ImGuiMultiSelectIO* ImGuiMultiSelectIO_new(ImGuiMultiSelectIO* param1);
// 来源: 类方法
MIQT_EXPORT void* ImGuiMultiSelectIO_Requests(const ImGuiMultiSelectIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_setRequests(ImGuiMultiSelectIO* self, void* Requests);
// 来源: 类方法
MIQT_EXPORT long long ImGuiMultiSelectIO_RangeSrcItem(const ImGuiMultiSelectIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_setRangeSrcItem(ImGuiMultiSelectIO* self, long long RangeSrcItem);
// 来源: 类方法
MIQT_EXPORT long long ImGuiMultiSelectIO_NavIdItem(const ImGuiMultiSelectIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_setNavIdItem(ImGuiMultiSelectIO* self, long long NavIdItem);
// 来源: 类方法
MIQT_EXPORT bool ImGuiMultiSelectIO_NavIdSelected(const ImGuiMultiSelectIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_setNavIdSelected(ImGuiMultiSelectIO* self, bool NavIdSelected);
// 来源: 类方法
MIQT_EXPORT bool ImGuiMultiSelectIO_RangeSrcReset(const ImGuiMultiSelectIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_setRangeSrcReset(ImGuiMultiSelectIO* self, bool RangeSrcReset);
// 来源: 类方法
MIQT_EXPORT int ImGuiMultiSelectIO_ItemsCount(const ImGuiMultiSelectIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_setItemsCount(ImGuiMultiSelectIO* self, int ItemsCount);
// 来源: 类方法
MIQT_EXPORT void ImGuiMultiSelectIO_operatorAssign(ImGuiMultiSelectIO* self, ImGuiMultiSelectIO* param1);

// 来源: 析构函数
MIQT_EXPORT void ImGuiMultiSelectIO_delete(ImGuiMultiSelectIO* self);

// 来源: 类 (ImGuiSelectionRequest)
// 来源: 类方法
MIQT_EXPORT int ImGuiSelectionRequest_Type(const ImGuiSelectionRequest* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionRequest_setType(ImGuiSelectionRequest* self, int Type);
// 来源: 类方法
MIQT_EXPORT bool ImGuiSelectionRequest_Selected(const ImGuiSelectionRequest* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionRequest_setSelected(ImGuiSelectionRequest* self, bool Selected);
// 来源: 类方法
MIQT_EXPORT signed char ImGuiSelectionRequest_RangeDirection(const ImGuiSelectionRequest* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionRequest_setRangeDirection(ImGuiSelectionRequest* self, signed char RangeDirection);
// 来源: 类方法
MIQT_EXPORT long long ImGuiSelectionRequest_RangeFirstItem(const ImGuiSelectionRequest* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionRequest_setRangeFirstItem(ImGuiSelectionRequest* self, long long RangeFirstItem);
// 来源: 类方法
MIQT_EXPORT long long ImGuiSelectionRequest_RangeLastItem(const ImGuiSelectionRequest* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionRequest_setRangeLastItem(ImGuiSelectionRequest* self, long long RangeLastItem);

// 来源: 析构函数
MIQT_EXPORT void ImGuiSelectionRequest_delete(ImGuiSelectionRequest* self);

// 来源: 类 (ImGuiSelectionBasicStorage)
// 来源: 构造函数
MIQT_EXPORT ImGuiSelectionBasicStorage* ImGuiSelectionBasicStorage_new();
// 来源: 类方法
MIQT_EXPORT int ImGuiSelectionBasicStorage_Size(const ImGuiSelectionBasicStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_setSize(ImGuiSelectionBasicStorage* self, int Size);
// 来源: 类方法
MIQT_EXPORT bool ImGuiSelectionBasicStorage_PreserveOrder(const ImGuiSelectionBasicStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_setPreserveOrder(ImGuiSelectionBasicStorage* self, bool PreserveOrder);
// 来源: 类方法
MIQT_EXPORT int ImGuiSelectionBasicStorage__SelectionOrder(const ImGuiSelectionBasicStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_set_SelectionOrder(ImGuiSelectionBasicStorage* self, int _SelectionOrder);
// 来源: 类方法
MIQT_EXPORT ImGuiStorage* ImGuiSelectionBasicStorage__Storage(const ImGuiSelectionBasicStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_set_Storage(ImGuiSelectionBasicStorage* self, ImGuiStorage* _Storage);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_ApplyRequests(ImGuiSelectionBasicStorage* self, ImGuiMultiSelectIO* ms_io);
// 来源: 类方法
MIQT_EXPORT bool ImGuiSelectionBasicStorage_Contains(const ImGuiSelectionBasicStorage* self, unsigned int id);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_Clear(ImGuiSelectionBasicStorage* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_Swap(ImGuiSelectionBasicStorage* self, ImGuiSelectionBasicStorage* r);
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionBasicStorage_SetItemSelected(ImGuiSelectionBasicStorage* self, unsigned int id, bool selected);
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiSelectionBasicStorage_GetStorageIdFromIndex(ImGuiSelectionBasicStorage* self, int idx);

// 来源: 析构函数
MIQT_EXPORT void ImGuiSelectionBasicStorage_delete(ImGuiSelectionBasicStorage* self);

// 来源: 类 (ImGuiSelectionExternalStorage)
// 来源: 构造函数
MIQT_EXPORT ImGuiSelectionExternalStorage* ImGuiSelectionExternalStorage_new();
// 来源: 类方法
MIQT_EXPORT void ImGuiSelectionExternalStorage_ApplyRequests(ImGuiSelectionExternalStorage* self, ImGuiMultiSelectIO* ms_io);

// 来源: 析构函数
MIQT_EXPORT void ImGuiSelectionExternalStorage_delete(ImGuiSelectionExternalStorage* self);

// 来源: 类 (ImDrawCmd)
// 来源: 构造函数
MIQT_EXPORT ImDrawCmd* ImDrawCmd_new();
// 来源: 类方法
MIQT_EXPORT ImVec4* ImDrawCmd_ClipRect(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setClipRect(ImDrawCmd* self, ImVec4* ClipRect);
// 来源: 类方法
MIQT_EXPORT ImTextureRef* ImDrawCmd_TexRef(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setTexRef(ImDrawCmd* self, ImTextureRef* TexRef);
// 来源: 类方法
MIQT_EXPORT unsigned int ImDrawCmd_VtxOffset(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setVtxOffset(ImDrawCmd* self, unsigned int VtxOffset);
// 来源: 类方法
MIQT_EXPORT unsigned int ImDrawCmd_IdxOffset(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setIdxOffset(ImDrawCmd* self, unsigned int IdxOffset);
// 来源: 类方法
MIQT_EXPORT unsigned int ImDrawCmd_ElemCount(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setElemCount(ImDrawCmd* self, unsigned int ElemCount);
// 来源: 类方法
MIQT_EXPORT int ImDrawCmd_UserCallbackDataSize(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setUserCallbackDataSize(ImDrawCmd* self, int UserCallbackDataSize);
// 来源: 类方法
MIQT_EXPORT int ImDrawCmd_UserCallbackDataOffset(const ImDrawCmd* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmd_setUserCallbackDataOffset(ImDrawCmd* self, int UserCallbackDataOffset);
// 来源: 类方法
MIQT_EXPORT unsigned long long ImDrawCmd_GetTexID(const ImDrawCmd* self);

// 来源: 析构函数
MIQT_EXPORT void ImDrawCmd_delete(ImDrawCmd* self);

// 来源: 类 (ImDrawVert)
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawVert_pos(const ImDrawVert* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawVert_setPos(ImDrawVert* self, ImVec2* pos);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawVert_uv(const ImDrawVert* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawVert_setUv(ImDrawVert* self, ImVec2* uv);
// 来源: 类方法
MIQT_EXPORT unsigned int ImDrawVert_col(const ImDrawVert* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawVert_setCol(ImDrawVert* self, unsigned int col);

// 来源: 析构函数
MIQT_EXPORT void ImDrawVert_delete(ImDrawVert* self);

// 来源: 类 (ImDrawCmdHeader)
// 来源: 构造函数
MIQT_EXPORT ImDrawCmdHeader* ImDrawCmdHeader_new(ImDrawCmdHeader* param1);
// 来源: 类方法
MIQT_EXPORT ImVec4* ImDrawCmdHeader_ClipRect(const ImDrawCmdHeader* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmdHeader_setClipRect(ImDrawCmdHeader* self, ImVec4* ClipRect);
// 来源: 类方法
MIQT_EXPORT ImTextureRef* ImDrawCmdHeader_TexRef(const ImDrawCmdHeader* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmdHeader_setTexRef(ImDrawCmdHeader* self, ImTextureRef* TexRef);
// 来源: 类方法
MIQT_EXPORT unsigned int ImDrawCmdHeader_VtxOffset(const ImDrawCmdHeader* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmdHeader_setVtxOffset(ImDrawCmdHeader* self, unsigned int VtxOffset);
// 来源: 类方法
MIQT_EXPORT void ImDrawCmdHeader_operatorAssign(ImDrawCmdHeader* self, ImDrawCmdHeader* param1);

// 来源: 析构函数
MIQT_EXPORT void ImDrawCmdHeader_delete(ImDrawCmdHeader* self);

// 来源: 类 (ImDrawChannel)
// 来源: 构造函数
MIQT_EXPORT ImDrawChannel* ImDrawChannel_new(ImDrawChannel* param1);
// 来源: 类方法
MIQT_EXPORT void* ImDrawChannel__CmdBuffer(const ImDrawChannel* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawChannel_set_CmdBuffer(ImDrawChannel* self, void* _CmdBuffer);
// 来源: 类方法
MIQT_EXPORT void* ImDrawChannel__IdxBuffer(const ImDrawChannel* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawChannel_set_IdxBuffer(ImDrawChannel* self, void* _IdxBuffer);
// 来源: 类方法
MIQT_EXPORT void ImDrawChannel_operatorAssign(ImDrawChannel* self, ImDrawChannel* param1);

// 来源: 析构函数
MIQT_EXPORT void ImDrawChannel_delete(ImDrawChannel* self);

// 来源: 类 (ImDrawListSplitter)
// 来源: 构造函数
MIQT_EXPORT ImDrawListSplitter* ImDrawListSplitter_new();
// 来源: 构造函数
MIQT_EXPORT ImDrawListSplitter* ImDrawListSplitter_new2(ImDrawListSplitter* param1);
// 来源: 类方法
MIQT_EXPORT int ImDrawListSplitter__Current(const ImDrawListSplitter* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_set_Current(ImDrawListSplitter* self, int _Current);
// 来源: 类方法
MIQT_EXPORT int ImDrawListSplitter__Count(const ImDrawListSplitter* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_set_Count(ImDrawListSplitter* self, int _Count);
// 来源: 类方法
MIQT_EXPORT void* ImDrawListSplitter__Channels(const ImDrawListSplitter* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_set_Channels(ImDrawListSplitter* self, void* _Channels);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_Clear(ImDrawListSplitter* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_Split(ImDrawListSplitter* self, ImDrawList* draw_list, int count);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_Merge(ImDrawListSplitter* self, ImDrawList* draw_list);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self, ImDrawList* draw_list, int channel_idx);
// 来源: 类方法
MIQT_EXPORT void ImDrawListSplitter_operatorAssign(ImDrawListSplitter* self, ImDrawListSplitter* param1);

// 来源: 析构函数
MIQT_EXPORT void ImDrawListSplitter_delete(ImDrawListSplitter* self);

// 来源: 类 (ImDrawList)
// 来源: 构造函数
MIQT_EXPORT ImDrawList* ImDrawList_new(void* shared_data);
// 来源: 构造函数
MIQT_EXPORT ImDrawList* ImDrawList_new2(ImDrawList* param1);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList_CmdBuffer(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_setCmdBuffer(ImDrawList* self, void* CmdBuffer);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList_IdxBuffer(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_setIdxBuffer(ImDrawList* self, void* IdxBuffer);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList_VtxBuffer(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_setVtxBuffer(ImDrawList* self, void* VtxBuffer);
// 来源: 类方法
MIQT_EXPORT int ImDrawList_Flags(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_setFlags(ImDrawList* self, int Flags);
// 来源: 类方法
MIQT_EXPORT unsigned int ImDrawList__VtxCurrentIdx(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_VtxCurrentIdx(ImDrawList* self, unsigned int _VtxCurrentIdx);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList__Data(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_Data(ImDrawList* self, void* _Data);
// 来源: 类方法
MIQT_EXPORT ImDrawVert* ImDrawList__VtxWritePtr(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_VtxWritePtr(ImDrawList* self, ImDrawVert* _VtxWritePtr);
// 来源: 类方法
MIQT_EXPORT unsigned short* ImDrawList__IdxWritePtr(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_IdxWritePtr(ImDrawList* self, unsigned short* _IdxWritePtr);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList__Path(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_Path(ImDrawList* self, void* _Path);
// 来源: 类方法
MIQT_EXPORT ImDrawCmdHeader* ImDrawList__CmdHeader(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_CmdHeader(ImDrawList* self, ImDrawCmdHeader* _CmdHeader);
// 来源: 类方法
MIQT_EXPORT ImDrawListSplitter* ImDrawList__Splitter(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_Splitter(ImDrawList* self, ImDrawListSplitter* _Splitter);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList__ClipRectStack(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_ClipRectStack(ImDrawList* self, void* _ClipRectStack);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList__TextureStack(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_TextureStack(ImDrawList* self, void* _TextureStack);
// 来源: 类方法
MIQT_EXPORT void* ImDrawList__CallbacksDataBuf(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_CallbacksDataBuf(ImDrawList* self, void* _CallbacksDataBuf);
// 来源: 类方法
MIQT_EXPORT float ImDrawList__FringeScale(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_FringeScale(ImDrawList* self, float _FringeScale);
// 来源: 类方法
MIQT_EXPORT const char* ImDrawList__OwnerName(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_set_OwnerName(ImDrawList* self, const char* _OwnerName);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PushClipRect(ImDrawList* self, ImVec2* clip_rect_min, ImVec2* clip_rect_max);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PushClipRectFullScreen(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PopClipRect(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PushTexture(ImDrawList* self, ImTextureRef* tex_ref);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PopTexture(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawList_GetClipRectMin(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawList_GetClipRectMax(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddLine(ImDrawList* self, ImVec2* p1, ImVec2* p2, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddLineH(ImDrawList* self, float min_x, float max_x, float y, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddLineV(ImDrawList* self, float x, float min_y, float max_y, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRect(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRectFilled(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRectFilledMultiColor(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col_upr_left, unsigned int col_upr_right, unsigned int col_bot_right, unsigned int col_bot_left);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddQuad(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddQuadFilled(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddTriangle(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddTriangleFilled(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddCircle(ImDrawList* self, ImVec2* center, float radius, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddCircleFilled(ImDrawList* self, ImVec2* center, float radius, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddNgon(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddNgonFilled(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipse(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipseFilled(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddText(ImDrawList* self, ImVec2* pos, unsigned int col, const char* text_begin);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddText2(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddBezierCubic(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddBezierQuadratic(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddPolyline(ImDrawList* self, ImVec2* points, int num_points, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddConvexPolyFilled(ImDrawList* self, ImVec2* points, int num_points, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddConcavePolyFilled(ImDrawList* self, ImVec2* points, int num_points, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImage(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageQuad(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageRounded(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max, unsigned int col, float rounding);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathClear(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathLineTo(ImDrawList* self, ImVec2* pos);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathLineToMergeDuplicate(ImDrawList* self, ImVec2* pos);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathFillConvex(ImDrawList* self, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathFillConcave(ImDrawList* self, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathStroke(ImDrawList* self, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathArcTo(ImDrawList* self, ImVec2* center, float radius, float a_min, float a_max);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathArcToFast(ImDrawList* self, ImVec2* center, float radius, int a_min_of_12, int a_max_of_12);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathEllipticalArcTo(ImDrawList* self, ImVec2* center, ImVec2* radius, float rot, float a_min, float a_max);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathBezierCubicCurveTo(ImDrawList* self, ImVec2* p2, ImVec2* p3, ImVec2* p4);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathBezierQuadraticCurveTo(ImDrawList* self, ImVec2* p2, ImVec2* p3);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathRect(ImDrawList* self, ImVec2* rect_min, ImVec2* rect_max);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddDrawCmd(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT ImDrawList* ImDrawList_CloneOutput(const ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_ChannelsSplit(ImDrawList* self, int count);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_ChannelsMerge(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_ChannelsSetCurrent(ImDrawList* self, int n);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimReserve(ImDrawList* self, int idx_count, int vtx_count);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimUnreserve(ImDrawList* self, int idx_count, int vtx_count);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimRect(ImDrawList* self, ImVec2* a, ImVec2* b, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimRectUV(ImDrawList* self, ImVec2* a, ImVec2* b, ImVec2* uv_a, ImVec2* uv_b, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimQuadUV(ImDrawList* self, ImVec2* a, ImVec2* b, ImVec2* c, ImVec2* d, ImVec2* uv_a, ImVec2* uv_b, ImVec2* uv_c, ImVec2* uv_d, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimWriteVtx(ImDrawList* self, ImVec2* pos, ImVec2* uv, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimWriteIdx(ImDrawList* self, unsigned short idx);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PrimVtx(ImDrawList* self, ImVec2* pos, ImVec2* uv, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRect2(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, int flags, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddPolyline2(ImDrawList* self, ImVec2* points, int num_points, unsigned int col, int flags, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathStroke2(ImDrawList* self, unsigned int col, int flags, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PushTextureID(ImDrawList* self, ImTextureRef* tex_ref);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PopTextureID(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__SetDrawListSharedData(ImDrawList* self, void* data);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__ResetForNewFrame(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__ClearFreeMemory(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__PopUnusedDrawCmd(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__TryMergeDrawCmds(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__OnChangedClipRect(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__OnChangedTexture(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__OnChangedVtxOffset(ImDrawList* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__SetTexture(ImDrawList* self, ImTextureRef* tex_ref);
// 来源: 类方法
MIQT_EXPORT int ImDrawList__CalcCircleAutoSegmentCount(const ImDrawList* self, float radius);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__PathArcToFastEx(ImDrawList* self, ImVec2* center, float radius, int a_min_sample, int a_max_sample, int a_step);
// 来源: 类方法
MIQT_EXPORT void ImDrawList__PathArcToN(ImDrawList* self, ImVec2* center, float radius, float a_min, float a_max, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_operatorAssign(ImDrawList* self, ImDrawList* param1);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PushClipRect2(ImDrawList* self, ImVec2* clip_rect_min, ImVec2* clip_rect_max, bool intersect_with_current_clip_rect);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddLine2(ImDrawList* self, ImVec2* p1, ImVec2* p2, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddLineH2(ImDrawList* self, float min_x, float max_x, float y, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddLineV2(ImDrawList* self, float x, float min_y, float max_y, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRect3(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRect4(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRect5(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, float thickness, int flags);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRectFilled2(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddRectFilled3(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, int flags);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddQuad2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddTriangle2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddCircle2(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddCircle3(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddCircleFilled2(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddNgon2(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipse2(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipse3(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipse4(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot, int num_segments, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipseFilled2(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddEllipseFilled3(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddText3(ImDrawList* self, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddText4(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddText5(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end, float wrap_width);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddText6(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end, float wrap_width, ImVec4* cpu_fine_clip_rect);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddBezierCubic2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col, float thickness, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddBezierQuadratic2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col, float thickness, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddPolyline3(ImDrawList* self, ImVec2* points, int num_points, unsigned int col, float thickness, int flags);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImage2(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImage3(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImage4(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageQuad2(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageQuad3(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageQuad4(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2, ImVec2* uv3);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageQuad5(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2, ImVec2* uv3, ImVec2* uv4);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageQuad6(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2, ImVec2* uv3, ImVec2* uv4, unsigned int col);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_AddImageRounded2(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max, unsigned int col, float rounding, int flags);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathStroke3(ImDrawList* self, unsigned int col, float thickness);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathStroke4(ImDrawList* self, unsigned int col, float thickness, int flags);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathArcTo2(ImDrawList* self, ImVec2* center, float radius, float a_min, float a_max, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathEllipticalArcTo2(ImDrawList* self, ImVec2* center, ImVec2* radius, float rot, float a_min, float a_max, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathBezierCubicCurveTo2(ImDrawList* self, ImVec2* p2, ImVec2* p3, ImVec2* p4, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathBezierQuadraticCurveTo2(ImDrawList* self, ImVec2* p2, ImVec2* p3, int num_segments);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathRect2(ImDrawList* self, ImVec2* rect_min, ImVec2* rect_max, float rounding);
// 来源: 类方法
MIQT_EXPORT void ImDrawList_PathRect3(ImDrawList* self, ImVec2* rect_min, ImVec2* rect_max, float rounding, int flags);

// 来源: 析构函数
MIQT_EXPORT void ImDrawList_delete(ImDrawList* self);

// 来源: 类 (ImDrawData)
// 来源: 构造函数
MIQT_EXPORT ImDrawData* ImDrawData_new();
// 来源: 构造函数
MIQT_EXPORT ImDrawData* ImDrawData_new2(ImDrawData* param1);
// 来源: 类方法
MIQT_EXPORT bool ImDrawData_Valid(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setValid(ImDrawData* self, bool Valid);
// 来源: 类方法
MIQT_EXPORT int ImDrawData_CmdListsCount(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setCmdListsCount(ImDrawData* self, int CmdListsCount);
// 来源: 类方法
MIQT_EXPORT int ImDrawData_TotalIdxCount(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setTotalIdxCount(ImDrawData* self, int TotalIdxCount);
// 来源: 类方法
MIQT_EXPORT int ImDrawData_TotalVtxCount(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setTotalVtxCount(ImDrawData* self, int TotalVtxCount);
// 来源: 类方法
MIQT_EXPORT void* ImDrawData_CmdLists(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setCmdLists(ImDrawData* self, void* CmdLists);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawData_DisplayPos(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setDisplayPos(ImDrawData* self, ImVec2* DisplayPos);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawData_DisplaySize(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setDisplaySize(ImDrawData* self, ImVec2* DisplaySize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImDrawData_FramebufferScale(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setFramebufferScale(ImDrawData* self, ImVec2* FramebufferScale);
// 来源: 类方法
MIQT_EXPORT ImGuiViewport* ImDrawData_OwnerViewport(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setOwnerViewport(ImDrawData* self, ImGuiViewport* OwnerViewport);
// 来源: 类方法
MIQT_EXPORT void* ImDrawData_Textures(const ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_setTextures(ImDrawData* self, void* Textures);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_Clear(ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_AddDrawList(ImDrawData* self, ImDrawList* draw_list);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_DeIndexAllBuffers(ImDrawData* self);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_ScaleClipRects(ImDrawData* self, ImVec2* fb_scale);
// 来源: 类方法
MIQT_EXPORT void ImDrawData_operatorAssign(ImDrawData* self, ImDrawData* param1);

// 来源: 析构函数
MIQT_EXPORT void ImDrawData_delete(ImDrawData* self);

// 来源: 类 (ImTextureRect)
// 来源: 构造函数
MIQT_EXPORT ImTextureRect* ImTextureRect_new(ImTextureRect* param1);
// 来源: 构造函数
MIQT_EXPORT ImTextureRect* ImTextureRect_new2();
// 来源: 类方法
MIQT_EXPORT unsigned short ImTextureRect_x(const ImTextureRect* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureRect_setX(ImTextureRect* self, unsigned short x);
// 来源: 类方法
MIQT_EXPORT unsigned short ImTextureRect_y(const ImTextureRect* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureRect_setY(ImTextureRect* self, unsigned short y);
// 来源: 类方法
MIQT_EXPORT unsigned short ImTextureRect_w(const ImTextureRect* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureRect_setW(ImTextureRect* self, unsigned short w);
// 来源: 类方法
MIQT_EXPORT unsigned short ImTextureRect_h(const ImTextureRect* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureRect_setH(ImTextureRect* self, unsigned short h);
// 来源: 类方法
MIQT_EXPORT void ImTextureRect_operatorAssign(ImTextureRect* self, ImTextureRect* param1);

// 来源: 析构函数
MIQT_EXPORT void ImTextureRect_delete(ImTextureRect* self);

// 来源: 类 (ImTextureData)
// 来源: 构造函数
MIQT_EXPORT ImTextureData* ImTextureData_new();
// 来源: 构造函数
MIQT_EXPORT ImTextureData* ImTextureData_new2(ImTextureData* param1);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_UniqueID(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setUniqueID(ImTextureData* self, int UniqueID);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_Status(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setStatus(ImTextureData* self, int Status);
// 来源: 类方法
MIQT_EXPORT unsigned long long ImTextureData_TexID(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setTexID(ImTextureData* self, unsigned long long TexID);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_Format(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setFormat(ImTextureData* self, int Format);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_Width(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setWidth(ImTextureData* self, int Width);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_Height(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setHeight(ImTextureData* self, int Height);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_BytesPerPixel(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setBytesPerPixel(ImTextureData* self, int BytesPerPixel);
// 来源: 类方法
MIQT_EXPORT unsigned char* ImTextureData_Pixels(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setPixels(ImTextureData* self, unsigned char* Pixels);
// 来源: 类方法
MIQT_EXPORT ImTextureRect* ImTextureData_UsedRect(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setUsedRect(ImTextureData* self, ImTextureRect* UsedRect);
// 来源: 类方法
MIQT_EXPORT ImTextureRect* ImTextureData_UpdateRect(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setUpdateRect(ImTextureData* self, ImTextureRect* UpdateRect);
// 来源: 类方法
MIQT_EXPORT void* ImTextureData_Updates(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setUpdates(ImTextureData* self, void* Updates);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_UnusedFrames(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setUnusedFrames(ImTextureData* self, int UnusedFrames);
// 来源: 类方法
MIQT_EXPORT unsigned short ImTextureData_RefCount(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setRefCount(ImTextureData* self, unsigned short RefCount);
// 来源: 类方法
MIQT_EXPORT bool ImTextureData_UseColors(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setUseColors(ImTextureData* self, bool UseColors);
// 来源: 类方法
MIQT_EXPORT bool ImTextureData_WantDestroyNextFrame(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_setWantDestroyNextFrame(ImTextureData* self, bool WantDestroyNextFrame);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_Create(ImTextureData* self, int format, int w, int h);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_DestroyPixels(ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void* ImTextureData_GetPixels(ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void* ImTextureData_GetPixelsAt(ImTextureData* self, int x, int y);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_GetSizeInBytes(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT int ImTextureData_GetPitch(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT ImTextureRef* ImTextureData_GetTexRef(ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT unsigned long long ImTextureData_GetTexID(const ImTextureData* self);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_SetTexID(ImTextureData* self, unsigned long long tex_id);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_SetStatus(ImTextureData* self, int status);
// 来源: 类方法
MIQT_EXPORT void ImTextureData_operatorAssign(ImTextureData* self, ImTextureData* param1);

// 来源: 析构函数
MIQT_EXPORT void ImTextureData_delete(ImTextureData* self);

// 来源: 类 (ImFontConfig)
// 来源: 构造函数
MIQT_EXPORT ImFontConfig* ImFontConfig_new();
// 来源: 类方法
MIQT_EXPORT int ImFontConfig_FontDataSize(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setFontDataSize(ImFontConfig* self, int FontDataSize);
// 来源: 类方法
MIQT_EXPORT bool ImFontConfig_FontDataOwnedByAtlas(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setFontDataOwnedByAtlas(ImFontConfig* self, bool FontDataOwnedByAtlas);
// 来源: 类方法
MIQT_EXPORT bool ImFontConfig_MergeMode(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setMergeMode(ImFontConfig* self, bool MergeMode);
// 来源: 类方法
MIQT_EXPORT bool ImFontConfig_PixelSnapH(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setPixelSnapH(ImFontConfig* self, bool PixelSnapH);
// 来源: 类方法
MIQT_EXPORT signed char ImFontConfig_OversampleH(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setOversampleH(ImFontConfig* self, signed char OversampleH);
// 来源: 类方法
MIQT_EXPORT signed char ImFontConfig_OversampleV(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setOversampleV(ImFontConfig* self, signed char OversampleV);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFontConfig_EllipsisChar(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setEllipsisChar(ImFontConfig* self, unsigned short EllipsisChar);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_SizePixels(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setSizePixels(ImFontConfig* self, float SizePixels);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontConfig_GlyphRanges(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setGlyphRanges(ImFontConfig* self, const unsigned short* GlyphRanges);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontConfig_GlyphExcludeRanges(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setGlyphExcludeRanges(ImFontConfig* self, const unsigned short* GlyphExcludeRanges);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFontConfig_GlyphOffset(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setGlyphOffset(ImFontConfig* self, ImVec2* GlyphOffset);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_GlyphMinAdvanceX(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setGlyphMinAdvanceX(ImFontConfig* self, float GlyphMinAdvanceX);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_GlyphMaxAdvanceX(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setGlyphMaxAdvanceX(ImFontConfig* self, float GlyphMaxAdvanceX);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_GlyphExtraAdvanceX(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setGlyphExtraAdvanceX(ImFontConfig* self, float GlyphExtraAdvanceX);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontConfig_FontNo(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setFontNo(ImFontConfig* self, unsigned int FontNo);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontConfig_FontLoaderFlags(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setFontLoaderFlags(ImFontConfig* self, unsigned int FontLoaderFlags);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_RasterizerMultiply(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setRasterizerMultiply(ImFontConfig* self, float RasterizerMultiply);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_RasterizerDensity(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setRasterizerDensity(ImFontConfig* self, float RasterizerDensity);
// 来源: 类方法
MIQT_EXPORT float ImFontConfig_ExtraSizeScale(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setExtraSizeScale(ImFontConfig* self, float ExtraSizeScale);
// 来源: 类方法
MIQT_EXPORT int ImFontConfig_Flags(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setFlags(ImFontConfig* self, int Flags);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontConfig_DstFont(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setDstFont(ImFontConfig* self, ImFont* DstFont);
// 来源: 类方法
MIQT_EXPORT const ImFontLoader* ImFontConfig_FontLoader(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setFontLoader(ImFontConfig* self, const ImFontLoader* FontLoader);
// 来源: 类方法
MIQT_EXPORT bool ImFontConfig_PixelSnapV(const ImFontConfig* self);
// 来源: 类方法
MIQT_EXPORT void ImFontConfig_setPixelSnapV(ImFontConfig* self, bool PixelSnapV);

// 来源: 析构函数
MIQT_EXPORT void ImFontConfig_delete(ImFontConfig* self);

// 来源: 类 (ImFontGlyph)
// 来源: 构造函数
MIQT_EXPORT ImFontGlyph* ImFontGlyph_new();
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontGlyph_Colored(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setColored(ImFontGlyph* self, unsigned int Colored);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontGlyph_Visible(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setVisible(ImFontGlyph* self, unsigned int Visible);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontGlyph_SourceIdx(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setSourceIdx(ImFontGlyph* self, unsigned int SourceIdx);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontGlyph_Codepoint(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setCodepoint(ImFontGlyph* self, unsigned int Codepoint);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_AdvanceX(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setAdvanceX(ImFontGlyph* self, float AdvanceX);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_X0(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setX0(ImFontGlyph* self, float X0);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_Y0(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setY0(ImFontGlyph* self, float Y0);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_X1(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setX1(ImFontGlyph* self, float X1);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_Y1(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setY1(ImFontGlyph* self, float Y1);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_U0(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setU0(ImFontGlyph* self, float U0);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_V0(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setV0(ImFontGlyph* self, float V0);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_U1(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setU1(ImFontGlyph* self, float U1);
// 来源: 类方法
MIQT_EXPORT float ImFontGlyph_V1(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setV1(ImFontGlyph* self, float V1);
// 来源: 类方法
MIQT_EXPORT int ImFontGlyph_PackId(const ImFontGlyph* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyph_setPackId(ImFontGlyph* self, int PackId);

// 来源: 析构函数
MIQT_EXPORT void ImFontGlyph_delete(ImFontGlyph* self);

// 来源: 类 (ImFontGlyphRangesBuilder)
// 来源: 构造函数
MIQT_EXPORT ImFontGlyphRangesBuilder* ImFontGlyphRangesBuilder_new();
// 来源: 构造函数
MIQT_EXPORT ImFontGlyphRangesBuilder* ImFontGlyphRangesBuilder_new2(ImFontGlyphRangesBuilder* param1);
// 来源: 类方法
MIQT_EXPORT void* ImFontGlyphRangesBuilder_UsedChars(const ImFontGlyphRangesBuilder* self);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_setUsedChars(ImFontGlyphRangesBuilder* self, void* UsedChars);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self);
// 来源: 类方法
MIQT_EXPORT bool ImFontGlyphRangesBuilder_GetBit(const ImFontGlyphRangesBuilder* self, unsigned long long n);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self, unsigned long long n);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self, const char* text);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self, const unsigned short* ranges);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self, void* out_ranges);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_operatorAssign(ImFontGlyphRangesBuilder* self, ImFontGlyphRangesBuilder* param1);
// 来源: 类方法
MIQT_EXPORT void ImFontGlyphRangesBuilder_AddText2(ImFontGlyphRangesBuilder* self, const char* text, const char* text_end);

// 来源: 析构函数
MIQT_EXPORT void ImFontGlyphRangesBuilder_delete(ImFontGlyphRangesBuilder* self);

// 来源: 类 (ImFontAtlasRect)
// 来源: 构造函数
MIQT_EXPORT ImFontAtlasRect* ImFontAtlasRect_new();
// 来源: 构造函数
MIQT_EXPORT ImFontAtlasRect* ImFontAtlasRect_new2(ImFontAtlasRect* param1);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFontAtlasRect_x(const ImFontAtlasRect* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_setX(ImFontAtlasRect* self, unsigned short x);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFontAtlasRect_y(const ImFontAtlasRect* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_setY(ImFontAtlasRect* self, unsigned short y);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFontAtlasRect_w(const ImFontAtlasRect* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_setW(ImFontAtlasRect* self, unsigned short w);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFontAtlasRect_h(const ImFontAtlasRect* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_setH(ImFontAtlasRect* self, unsigned short h);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFontAtlasRect_uv0(const ImFontAtlasRect* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_setUv0(ImFontAtlasRect* self, ImVec2* uv0);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFontAtlasRect_uv1(const ImFontAtlasRect* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_setUv1(ImFontAtlasRect* self, ImVec2* uv1);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlasRect_operatorAssign(ImFontAtlasRect* self, ImFontAtlasRect* param1);

// 来源: 析构函数
MIQT_EXPORT void ImFontAtlasRect_delete(ImFontAtlasRect* self);

// 来源: 类 (ImFontAtlas)
// 来源: 构造函数
MIQT_EXPORT ImFontAtlas* ImFontAtlas_new();
// 来源: 构造函数
MIQT_EXPORT ImFontAtlas* ImFontAtlas_new2(ImFontAtlas* param1);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFont(ImFontAtlas* self, ImFontConfig* font_cfg);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontDefault(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontDefaultVector(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontDefaultBitmap(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromFileTTF(ImFontAtlas* self, const char* filename);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self, const char* compressed_font_data_base85);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_RemoveFont(ImFontAtlas* self, ImFont* font);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_Clear(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_ClearFonts(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_CompactCache(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_SetFontLoader(ImFontAtlas* self, const ImFontLoader* font_loader);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_ClearInputData(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_ClearTexData(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_Build(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_SetTexID(ImFontAtlas* self, unsigned long long id);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_SetTexIDWithId(ImFontAtlas* self, ImTextureRef* id);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_IsBuilt(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesDefault(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesGreek(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesKorean(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesThai(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT const unsigned short* ImFontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRect(ImFontAtlas* self, int width, int height);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_RemoveCustomRect(ImFontAtlas* self, int id);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_GetCustomRect(const ImFontAtlas* self, int id, ImFontAtlasRect* out_r);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_Flags(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setFlags(ImFontAtlas* self, int Flags);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexDesiredFormat(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexDesiredFormat(ImFontAtlas* self, int TexDesiredFormat);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexGlyphPadding(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexGlyphPadding(ImFontAtlas* self, int TexGlyphPadding);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexMinWidth(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexMinWidth(ImFontAtlas* self, int TexMinWidth);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexMinHeight(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexMinHeight(ImFontAtlas* self, int TexMinHeight);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexMaxWidth(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexMaxWidth(ImFontAtlas* self, int TexMaxWidth);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexMaxHeight(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexMaxHeight(ImFontAtlas* self, int TexMaxHeight);
// 来源: 类方法
MIQT_EXPORT ImTextureData* ImFontAtlas_TexData(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexData(ImFontAtlas* self, ImTextureData* TexData);
// 来源: 类方法
MIQT_EXPORT void* ImFontAtlas_TexList(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexList(ImFontAtlas* self, void* TexList);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_Locked(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setLocked(ImFontAtlas* self, bool Locked);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_RendererHasTextures(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setRendererHasTextures(ImFontAtlas* self, bool RendererHasTextures);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_TexIsBuilt(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexIsBuilt(ImFontAtlas* self, bool TexIsBuilt);
// 来源: 类方法
MIQT_EXPORT bool ImFontAtlas_TexPixelsUseColors(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexPixelsUseColors(ImFontAtlas* self, bool TexPixelsUseColors);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFontAtlas_TexUvScale(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexUvScale(ImFontAtlas* self, ImVec2* TexUvScale);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFontAtlas_TexUvWhitePixel(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexUvWhitePixel(ImFontAtlas* self, ImVec2* TexUvWhitePixel);
// 来源: 类方法
MIQT_EXPORT void* ImFontAtlas_Fonts(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setFonts(ImFontAtlas* self, void* Fonts);
// 来源: 类方法
MIQT_EXPORT void* ImFontAtlas_Sources(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setSources(ImFontAtlas* self, void* Sources);
MIQT_EXPORT ImVec4* ImFontAtlas_TexUvLines(const ImFontAtlas* self);
MIQT_EXPORT void ImFontAtlas_setTexUvLines(ImFontAtlas* self, const ImVec4* TexUvLines);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_TexNextUniqueID(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTexNextUniqueID(ImFontAtlas* self, int TexNextUniqueID);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_FontNextUniqueID(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setFontNextUniqueID(ImFontAtlas* self, int FontNextUniqueID);
// 来源: 类方法
MIQT_EXPORT void* ImFontAtlas_DrawListSharedDatas(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setDrawListSharedDatas(ImFontAtlas* self, void* DrawListSharedDatas);
// 来源: 类方法
MIQT_EXPORT void* ImFontAtlas_Builder(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setBuilder(ImFontAtlas* self, void* Builder);
// 来源: 类方法
MIQT_EXPORT const ImFontLoader* ImFontAtlas_FontLoader(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setFontLoader(ImFontAtlas* self, const ImFontLoader* FontLoader);
// 来源: 类方法
MIQT_EXPORT const char* ImFontAtlas_FontLoaderName(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setFontLoaderName(ImFontAtlas* self, const char* FontLoaderName);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontAtlas_FontLoaderFlags(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setFontLoaderFlags(ImFontAtlas* self, unsigned int FontLoaderFlags);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_RefCount(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setRefCount(ImFontAtlas* self, int RefCount);
// 来源: 类方法
MIQT_EXPORT void* ImFontAtlas_OwnerContext(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setOwnerContext(ImFontAtlas* self, void* OwnerContext);
// 来源: 类方法
MIQT_EXPORT ImFontAtlasRect* ImFontAtlas_TempRect(const ImFontAtlas* self);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_setTempRect(ImFontAtlas* self, ImFontAtlasRect* TempRect);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRectRegular(ImFontAtlas* self, int w, int h);
// 来源: 类方法
MIQT_EXPORT ImFontAtlasRect* ImFontAtlas_GetCustomRectByIndex(ImFontAtlas* self, int id);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_CalcCustomRectUV(const ImFontAtlas* self, ImFontAtlasRect* r, ImVec2* out_uv_min, ImVec2* out_uv_max);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self, ImFont* font, unsigned short codepoint, int w, int h, float advance_x);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRectFontGlyphForSize(ImFontAtlas* self, ImFont* font, float font_size, unsigned short codepoint, int w, int h, float advance_x);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_operatorAssign(ImFontAtlas* self, ImFontAtlas* param1);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontDefaultWithFontCfg(ImFontAtlas* self, ImFontConfig* font_cfg);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontDefaultVectorWithFontCfg(ImFontAtlas* self, ImFontConfig* font_cfg);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontDefaultBitmapWithFontCfg(ImFontAtlas* self, ImFontConfig* font_cfg);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromFileTTF2(ImFontAtlas* self, const char* filename, float size_pixels);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromFileTTF3(ImFontAtlas* self, const char* filename, float size_pixels, ImFontConfig* font_cfg);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromFileTTF4(ImFontAtlas* self, const char* filename, float size_pixels, ImFontConfig* font_cfg, const unsigned short* glyph_ranges);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF2(ImFontAtlas* self, const char* compressed_font_data_base85, float size_pixels);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF3(ImFontAtlas* self, const char* compressed_font_data_base85, float size_pixels, ImFontConfig* font_cfg);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF4(ImFontAtlas* self, const char* compressed_font_data_base85, float size_pixels, ImFontConfig* font_cfg, const unsigned short* glyph_ranges);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_GetTexDataAsAlpha82(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height, int* out_bytes_per_pixel);
// 来源: 类方法
MIQT_EXPORT void ImFontAtlas_GetTexDataAsRGBA322(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height, int* out_bytes_per_pixel);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRect2(ImFontAtlas* self, int width, int height, ImFontAtlasRect* out_r);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRectFontGlyph2(ImFontAtlas* self, ImFont* font, unsigned short codepoint, int w, int h, float advance_x, ImVec2* offset);
// 来源: 类方法
MIQT_EXPORT int ImFontAtlas_AddCustomRectFontGlyphForSize2(ImFontAtlas* self, ImFont* font, float font_size, unsigned short codepoint, int w, int h, float advance_x, ImVec2* offset);

// 来源: 析构函数
MIQT_EXPORT void ImFontAtlas_delete(ImFontAtlas* self);

// 来源: 类 (ImFontBaked)
// 来源: 构造函数
MIQT_EXPORT ImFontBaked* ImFontBaked_new();
// 来源: 构造函数
MIQT_EXPORT ImFontBaked* ImFontBaked_new2(ImFontBaked* param1);
// 来源: 类方法
MIQT_EXPORT void* ImFontBaked_IndexAdvanceX(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setIndexAdvanceX(ImFontBaked* self, void* IndexAdvanceX);
// 来源: 类方法
MIQT_EXPORT float ImFontBaked_FallbackAdvanceX(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setFallbackAdvanceX(ImFontBaked* self, float FallbackAdvanceX);
// 来源: 类方法
MIQT_EXPORT float ImFontBaked_Size(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setSize(ImFontBaked* self, float Size);
// 来源: 类方法
MIQT_EXPORT float ImFontBaked_RasterizerDensity(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setRasterizerDensity(ImFontBaked* self, float RasterizerDensity);
// 来源: 类方法
MIQT_EXPORT void* ImFontBaked_IndexLookup(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setIndexLookup(ImFontBaked* self, void* IndexLookup);
// 来源: 类方法
MIQT_EXPORT void* ImFontBaked_Glyphs(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setGlyphs(ImFontBaked* self, void* Glyphs);
// 来源: 类方法
MIQT_EXPORT int ImFontBaked_FallbackGlyphIndex(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setFallbackGlyphIndex(ImFontBaked* self, int FallbackGlyphIndex);
// 来源: 类方法
MIQT_EXPORT float ImFontBaked_Ascent(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setAscent(ImFontBaked* self, float Ascent);
// 来源: 类方法
MIQT_EXPORT float ImFontBaked_Descent(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setDescent(ImFontBaked* self, float Descent);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontBaked_MetricsTotalSurface(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setMetricsTotalSurface(ImFontBaked* self, unsigned int MetricsTotalSurface);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontBaked_WantDestroy(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setWantDestroy(ImFontBaked* self, unsigned int WantDestroy);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontBaked_LoadNoFallback(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setLoadNoFallback(ImFontBaked* self, unsigned int LoadNoFallback);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontBaked_LoadNoRenderOnLayout(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setLoadNoRenderOnLayout(ImFontBaked* self, unsigned int LoadNoRenderOnLayout);
// 来源: 类方法
MIQT_EXPORT int ImFontBaked_LastUsedFrame(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setLastUsedFrame(ImFontBaked* self, int LastUsedFrame);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFontBaked_BakedId(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setBakedId(ImFontBaked* self, unsigned int BakedId);
// 来源: 类方法
MIQT_EXPORT ImFont* ImFontBaked_OwnerFont(const ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_setOwnerFont(ImFontBaked* self, ImFont* OwnerFont);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_ClearOutputData(ImFontBaked* self);
// 来源: 类方法
MIQT_EXPORT ImFontGlyph* ImFontBaked_FindGlyph(ImFontBaked* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT ImFontGlyph* ImFontBaked_FindGlyphNoFallback(ImFontBaked* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT float ImFontBaked_GetCharAdvance(ImFontBaked* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT bool ImFontBaked_IsGlyphLoaded(ImFontBaked* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT void ImFontBaked_operatorAssign(ImFontBaked* self, ImFontBaked* param1);

// 来源: 析构函数
MIQT_EXPORT void ImFontBaked_delete(ImFontBaked* self);

// 来源: 类 (ImFont)
// 来源: 构造函数
MIQT_EXPORT ImFont* ImFont_new();
// 来源: 构造函数
MIQT_EXPORT ImFont* ImFont_new2(ImFont* param1);
// 来源: 类方法
MIQT_EXPORT ImFontBaked* ImFont_LastBaked(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setLastBaked(ImFont* self, ImFontBaked* LastBaked);
// 来源: 类方法
MIQT_EXPORT ImFontAtlas* ImFont_OwnerAtlas(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setOwnerAtlas(ImFont* self, ImFontAtlas* OwnerAtlas);
// 来源: 类方法
MIQT_EXPORT int ImFont_Flags(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setFlags(ImFont* self, int Flags);
// 来源: 类方法
MIQT_EXPORT float ImFont_CurrentRasterizerDensity(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setCurrentRasterizerDensity(ImFont* self, float CurrentRasterizerDensity);
// 来源: 类方法
MIQT_EXPORT unsigned int ImFont_FontId(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setFontId(ImFont* self, unsigned int FontId);
// 来源: 类方法
MIQT_EXPORT float ImFont_LegacySize(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setLegacySize(ImFont* self, float LegacySize);
// 来源: 类方法
MIQT_EXPORT void* ImFont_Sources(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setSources(ImFont* self, void* Sources);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFont_EllipsisChar(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setEllipsisChar(ImFont* self, unsigned short EllipsisChar);
// 来源: 类方法
MIQT_EXPORT unsigned short ImFont_FallbackChar(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setFallbackChar(ImFont* self, unsigned short FallbackChar);
MIQT_EXPORT ImU8* ImFont_Used8kPagesMap(const ImFont* self);
MIQT_EXPORT void ImFont_setUsed8kPagesMap(ImFont* self, const ImU8* Used8kPagesMap);
// 来源: 类方法
MIQT_EXPORT bool ImFont_EllipsisAutoBake(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setEllipsisAutoBake(ImFont* self, bool EllipsisAutoBake);
// 来源: 类方法
MIQT_EXPORT ImGuiStorage* ImFont_RemapPairs(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setRemapPairs(ImFont* self, ImGuiStorage* RemapPairs);
// 来源: 类方法
MIQT_EXPORT float ImFont_Scale(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_setScale(ImFont* self, float Scale);
// 来源: 类方法
MIQT_EXPORT bool ImFont_IsGlyphInFont(ImFont* self, unsigned short c);
// 来源: 类方法
MIQT_EXPORT bool ImFont_IsLoaded(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT const char* ImFont_GetDebugName(const ImFont* self);
// 来源: 类方法
MIQT_EXPORT ImFontBaked* ImFont_GetFontBaked(ImFont* self, float font_size);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFont_CalcTextSizeA(ImFont* self, float size, float max_width, float wrap_width, const char* text_begin);
// 来源: 类方法
MIQT_EXPORT const char* ImFont_CalcWordWrapPosition(ImFont* self, float size, const char* text, const char* text_end, float wrap_width);
// 来源: 类方法
MIQT_EXPORT void ImFont_RenderChar(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, unsigned short c);
// 来源: 类方法
MIQT_EXPORT void ImFont_RenderText(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, ImVec4* clip_rect, const char* text_begin, const char* text_end);
// 来源: 类方法
MIQT_EXPORT const char* ImFont_CalcWordWrapPositionA(ImFont* self, float scale, const char* text, const char* text_end, float wrap_width);
// 来源: 类方法
MIQT_EXPORT void ImFont_ClearOutputData(ImFont* self);
// 来源: 类方法
MIQT_EXPORT void ImFont_AddRemapChar(ImFont* self, unsigned short from_codepoint, unsigned short to_codepoint);
// 来源: 类方法
MIQT_EXPORT bool ImFont_IsGlyphRangeUnused(ImFont* self, unsigned int c_begin, unsigned int c_last);
// 来源: 类方法
MIQT_EXPORT void ImFont_operatorAssign(ImFont* self, ImFont* param1);
// 来源: 类方法
MIQT_EXPORT ImFontBaked* ImFont_GetFontBaked2(ImFont* self, float font_size, float density);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFont_CalcTextSizeA2(ImFont* self, float size, float max_width, float wrap_width, const char* text_begin, const char* text_end);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImFont_CalcTextSizeA3(ImFont* self, float size, float max_width, float wrap_width, const char* text_begin, const char* text_end, const char** out_remaining);
// 来源: 类方法
MIQT_EXPORT void ImFont_RenderChar2(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, unsigned short c, ImVec4* cpu_fine_clip);
// 来源: 类方法
MIQT_EXPORT void ImFont_RenderText2(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, ImVec4* clip_rect, const char* text_begin, const char* text_end, float wrap_width);
// 来源: 类方法
MIQT_EXPORT void ImFont_RenderText3(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, ImVec4* clip_rect, const char* text_begin, const char* text_end, float wrap_width, int flags);

// 来源: 析构函数
MIQT_EXPORT void ImFont_delete(ImFont* self);

// 来源: 类 (ImGuiViewport)
// 来源: 构造函数
MIQT_EXPORT ImGuiViewport* ImGuiViewport_new();
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiViewport_ID(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setID(ImGuiViewport* self, unsigned int ID);
// 来源: 类方法
MIQT_EXPORT int ImGuiViewport_Flags(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setFlags(ImGuiViewport* self, int Flags);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_Pos(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setPos(ImGuiViewport* self, ImVec2* Pos);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_Size(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setSize(ImGuiViewport* self, ImVec2* Size);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_FramebufferScale(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setFramebufferScale(ImGuiViewport* self, ImVec2* FramebufferScale);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_WorkPos(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setWorkPos(ImGuiViewport* self, ImVec2* WorkPos);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_WorkSize(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiViewport_setWorkSize(ImGuiViewport* self, ImVec2* WorkSize);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_GetCenter(const ImGuiViewport* self);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiViewport_GetWorkCenter(const ImGuiViewport* self);

// 来源: 析构函数
MIQT_EXPORT void ImGuiViewport_delete(ImGuiViewport* self);

// 来源: 类 (ImGuiPlatformIO)
// 来源: 构造函数
MIQT_EXPORT ImGuiPlatformIO* ImGuiPlatformIO_new();
// 来源: 构造函数
MIQT_EXPORT ImGuiPlatformIO* ImGuiPlatformIO_new2(ImGuiPlatformIO* param1);
// 来源: 类方法
MIQT_EXPORT unsigned short ImGuiPlatformIO_Platform_LocaleDecimalPoint(const ImGuiPlatformIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_setPlatform_LocaleDecimalPoint(ImGuiPlatformIO* self, unsigned short Platform_LocaleDecimalPoint);
// 来源: 类方法
MIQT_EXPORT int ImGuiPlatformIO_Renderer_TextureMaxWidth(const ImGuiPlatformIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_setRenderer_TextureMaxWidth(ImGuiPlatformIO* self, int Renderer_TextureMaxWidth);
// 来源: 类方法
MIQT_EXPORT int ImGuiPlatformIO_Renderer_TextureMaxHeight(const ImGuiPlatformIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_setRenderer_TextureMaxHeight(ImGuiPlatformIO* self, int Renderer_TextureMaxHeight);
// 来源: 类方法
MIQT_EXPORT void* ImGuiPlatformIO_Textures(const ImGuiPlatformIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_setTextures(ImGuiPlatformIO* self, void* Textures);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_ClearPlatformHandlers(ImGuiPlatformIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_ClearRendererHandlers(ImGuiPlatformIO* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformIO_operatorAssign(ImGuiPlatformIO* self, ImGuiPlatformIO* param1);

// 来源: 析构函数
MIQT_EXPORT void ImGuiPlatformIO_delete(ImGuiPlatformIO* self);

// 来源: 类 (ImGuiPlatformImeData)
// 来源: 构造函数
MIQT_EXPORT ImGuiPlatformImeData* ImGuiPlatformImeData_new();
// 来源: 类方法
MIQT_EXPORT bool ImGuiPlatformImeData_WantVisible(const ImGuiPlatformImeData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformImeData_setWantVisible(ImGuiPlatformImeData* self, bool WantVisible);
// 来源: 类方法
MIQT_EXPORT bool ImGuiPlatformImeData_WantTextInput(const ImGuiPlatformImeData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformImeData_setWantTextInput(ImGuiPlatformImeData* self, bool WantTextInput);
// 来源: 类方法
MIQT_EXPORT ImVec2* ImGuiPlatformImeData_InputPos(const ImGuiPlatformImeData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformImeData_setInputPos(ImGuiPlatformImeData* self, ImVec2* InputPos);
// 来源: 类方法
MIQT_EXPORT float ImGuiPlatformImeData_InputLineHeight(const ImGuiPlatformImeData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformImeData_setInputLineHeight(ImGuiPlatformImeData* self, float InputLineHeight);
// 来源: 类方法
MIQT_EXPORT unsigned int ImGuiPlatformImeData_ViewportId(const ImGuiPlatformImeData* self);
// 来源: 类方法
MIQT_EXPORT void ImGuiPlatformImeData_setViewportId(ImGuiPlatformImeData* self, unsigned int ViewportId);

// 来源: 析构函数
MIQT_EXPORT void ImGuiPlatformImeData_delete(ImGuiPlatformImeData* self);

// 来源: 类 (ImGui_ImplDX11_RenderState)
// 来源: 类方法
MIQT_EXPORT void* ImGui_ImplDX11_RenderState_Device(const ImGui_ImplDX11_RenderState* self);
// 来源: 类方法
MIQT_EXPORT void ImGui_ImplDX11_RenderState_setDevice(ImGui_ImplDX11_RenderState* self, void* Device);
// 来源: 类方法
MIQT_EXPORT void* ImGui_ImplDX11_RenderState_DeviceContext(const ImGui_ImplDX11_RenderState* self);
// 来源: 类方法
MIQT_EXPORT void ImGui_ImplDX11_RenderState_setDeviceContext(ImGui_ImplDX11_RenderState* self, void* DeviceContext);
// 来源: 类方法
MIQT_EXPORT void* ImGui_ImplDX11_RenderState_VertexConstantBuffer(const ImGui_ImplDX11_RenderState* self);
// 来源: 类方法
MIQT_EXPORT void ImGui_ImplDX11_RenderState_setVertexConstantBuffer(ImGui_ImplDX11_RenderState* self, void* VertexConstantBuffer);

// 来源: 析构函数
MIQT_EXPORT void ImGui_ImplDX11_RenderState_delete(ImGui_ImplDX11_RenderState* self);

// 来源: 自由函数
MIQT_EXPORT void* cabi_ImGui__CreateContext(ImFontAtlas* shared_font_atlas);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__DestroyContext(void* ctx);
// 来源: 自由函数
MIQT_EXPORT void* cabi_ImGui__GetCurrentContext();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetCurrentContext(void* ctx);
// 来源: 自由函数
MIQT_EXPORT ImGuiIO* cabi_ImGui__GetIO();
// 来源: 自由函数
MIQT_EXPORT ImGuiPlatformIO* cabi_ImGui__GetPlatformIO();
// 来源: 自由函数
MIQT_EXPORT ImGuiStyle* cabi_ImGui__GetStyle();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__NewFrame();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndFrame();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Render();
// 来源: 自由函数
MIQT_EXPORT ImDrawData* cabi_ImGui__GetDrawData();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowDemoWindow(bool* p_open);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowMetricsWindow(bool* p_open);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowDebugLogWindow(bool* p_open);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowIDStackToolWindow(bool* p_open);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowAboutWindow(bool* p_open);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowStyleEditor(ImGuiStyle* ref);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ShowStyleSelector(const char* label);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowFontSelector(const char* label);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ShowUserGuide();
// 来源: 自由函数
MIQT_EXPORT const char* cabi_ImGui__GetVersion();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__StyleColorsDark(ImGuiStyle* dst);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__StyleColorsLight(ImGuiStyle* dst);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__StyleColorsClassic(ImGuiStyle* dst);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Begin(const char* name, bool* p_open, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__End();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginChild_1(const char* str_id, ImVec2* size, int child_flags, int window_flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginChild_2(int id, ImVec2* size, int child_flags, int window_flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndChild();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsWindowAppearing();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsWindowCollapsed();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsWindowFocused(int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsWindowHovered(int flags);
// 来源: 自由函数
MIQT_EXPORT ImDrawList* cabi_ImGui__GetWindowDrawList();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetWindowPos();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetWindowSize();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetWindowWidth();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetWindowHeight();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowPos(ImVec2* pos, int cond, ImVec2* pivot);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowSize(ImVec2* size, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowSizeConstraints(ImVec2* size_min, ImVec2* size_max, void* custom_callback, void* custom_callback_data);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowContentSize(ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowCollapsed(bool collapsed, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowFocus();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowScroll(ImVec2* scroll);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextWindowBgAlpha(float alpha);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowPos_1(ImVec2* pos, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowSize_1(ImVec2* size, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowCollapsed_1(bool collapsed, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowFocus_1();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowPos_2(const char* name, ImVec2* pos, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowSize_2(const char* name, ImVec2* size, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowCollapsed_2(const char* name, bool collapsed, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowFocus_2(const char* name);
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetScrollX();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetScrollY();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetScrollX(float scroll_x);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetScrollY(float scroll_y);
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetScrollMaxX();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetScrollMaxY();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetScrollHereX(float center_x_ratio);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetScrollHereY(float center_y_ratio);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetScrollFromPosX(float local_x, float center_x_ratio);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetScrollFromPosY(float local_y, float center_y_ratio);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushFont_1(ImFont* font, float font_size_base_unscaled);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopFont();
// 来源: 自由函数
MIQT_EXPORT ImFont* cabi_ImGui__GetFont();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetFontSize();
// 来源: 自由函数
MIQT_EXPORT ImFontBaked* cabi_ImGui__GetFontBaked();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushStyleColor_1(int idx, int col);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushStyleColor_2(int idx, ImVec4* col);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopStyleColor(int count);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushStyleVar_1(int idx, float val);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushStyleVar_2(int idx, ImVec2* val);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushStyleVarX(int idx, float val_x);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushStyleVarY(int idx, float val_y);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopStyleVar(int count);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushItemFlag(int option, bool enabled);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopItemFlag();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushItemWidth(float item_width);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopItemWidth();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextItemWidth(float item_width);
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__CalcItemWidth();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushTextWrapPos(float wrap_local_pos_x);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopTextWrapPos();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetFontTexUvWhitePixel();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetColorU32_1(int idx, float alpha_mul);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetColorU32_2(ImVec4* col);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetColorU32_3(int col, float alpha_mul);
// 来源: 自由函数
MIQT_EXPORT ImVec4* cabi_ImGui__GetStyleColorVec4(int idx);
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetCursorScreenPos();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetCursorScreenPos(ImVec2* pos);
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetContentRegionAvail();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetCursorPos();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetCursorPosX();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetCursorPosY();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetCursorPos(ImVec2* local_pos);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetCursorPosX(float local_x);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetCursorPosY(float local_y);
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetCursorStartPos();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Separator();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SameLine(float offset_from_start_x, float spacing);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__NewLine();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Spacing();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Dummy(ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Indent(float indent_w);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Unindent(float indent_w);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__BeginGroup();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndGroup();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__AlignTextToFramePadding();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetTextLineHeight();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetTextLineHeightWithSpacing();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetFrameHeight();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetFrameHeightWithSpacing();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushID_1(const char* str_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushID_2(const char* str_id_begin, const char* str_id_end);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushID_3(const void* ptr_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushID_4(int int_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopID();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetID_1(const char* str_id);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetID_2(const char* str_id_begin, const char* str_id_end);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetID_3(const void* ptr_id);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetID_4(int int_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TextUnformatted(const char* text, const char* text_end);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TextV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TextColoredV(ImVec4* col, const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TextDisabledV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TextWrappedV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LabelTextV(const char* label, const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__BulletTextV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SeparatorText(const char* label);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Button(const char* label, ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SmallButton(const char* label);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InvisibleButton(const char* str_id, ImVec2* size, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ArrowButton(const char* str_id, int dir);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Checkbox(const char* label, bool* v);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__CheckboxFlags_1(const char* label, int* flags, int flags_value);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__CheckboxFlags_2(const char* label, unsigned int* flags, unsigned int flags_value);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__RadioButton_1(const char* label, bool active);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__RadioButton_2(const char* label, int* v, int v_button);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ProgressBar(float fraction, ImVec2* size_arg, const char* overlay);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Bullet();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TextLink(const char* label);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TextLinkOpenURL(const char* label, const char* url);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Image_1(ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ImageWithBg(ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1, ImVec4* bg_col, ImVec4* tint_col);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ImageButton(const char* str_id, ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1, ImVec4* bg_col, ImVec4* tint_col);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginCombo(const char* label, const char* preview_value, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndCombo();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Combo_1(const char* label, int* current_item, const char** items, int items_count, int popup_max_height_in_items);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Combo_2(const char* label, int* current_item, const char* items_separated_by_zeros, int popup_max_height_in_items);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragFloat(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragFloat2(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragFloat3(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragFloat4(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragFloatRange2(const char* label, float* v_current_min, float* v_current_max, float v_speed, float v_min, float v_max, const char* format, const char* format_max, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragInt(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragInt2(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragInt3(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragInt4(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragIntRange2(const char* label, int* v_current_min, int* v_current_max, float v_speed, int v_min, int v_max, const char* format, const char* format_max, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragScalar(const char* label, int data_type, void* p_data, float v_speed, const void* p_min, const void* p_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DragScalarN(const char* label, int data_type, void* p_data, int components, float v_speed, const void* p_min, const void* p_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderFloat(const char* label, float* v, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderFloat2(const char* label, float* v, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderFloat3(const char* label, float* v, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderFloat4(const char* label, float* v, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderAngle(const char* label, float* v_rad, float v_degrees_min, float v_degrees_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderInt(const char* label, int* v, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderInt2(const char* label, int* v, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderInt3(const char* label, int* v, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderInt4(const char* label, int* v, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderScalar(const char* label, int data_type, void* p_data, const void* p_min, const void* p_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SliderScalarN(const char* label, int data_type, void* p_data, int components, const void* p_min, const void* p_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__VSliderFloat(const char* label, ImVec2* size, float* v, float v_min, float v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__VSliderInt(const char* label, ImVec2* size, int* v, int v_min, int v_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__VSliderScalar(const char* label, ImVec2* size, int data_type, void* p_data, const void* p_min, const void* p_max, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputText(const char* label, char* buf, size_t buf_size, int flags, void* callback, void* user_data);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputTextMultiline(const char* label, char* buf, size_t buf_size, ImVec2* size, int flags, void* callback, void* user_data);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputTextWithHint(const char* label, const char* hint, char* buf, size_t buf_size, int flags, void* callback, void* user_data);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputFloat(const char* label, float* v, float step, float step_fast, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputFloat2(const char* label, float* v, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputFloat3(const char* label, float* v, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputFloat4(const char* label, float* v, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputInt(const char* label, int* v, int step, int step_fast, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputInt2(const char* label, int* v, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputInt3(const char* label, int* v, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputInt4(const char* label, int* v, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputDouble(const char* label, double* v, double step, double step_fast, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputScalar(const char* label, int data_type, void* p_data, const void* p_step, const void* p_step_fast, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__InputScalarN(const char* label, int data_type, void* p_data, int components, const void* p_step, const void* p_step_fast, const char* format, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ColorEdit3(const char* label, float* col, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ColorEdit4(const char* label, float* col, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ColorPicker3(const char* label, float* col, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ColorPicker4(const char* label, float* col, int flags, const float* ref_col);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ColorButton(const char* desc_id, ImVec4* col, int flags, ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetColorEditOptions(int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNode(const char* label);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNodeV_1(const char* str_id, const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNodeV_2(const void* ptr_id, const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNodeEx(const char* label, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNodeExV_1(const char* str_id, int flags, const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNodeExV_2(const void* ptr_id, int flags, const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TreePush_1(const char* str_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TreePush_2(const void* ptr_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TreePop();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetTreeNodeToLabelSpacing();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__CollapsingHeader_1(const char* label, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__CollapsingHeader_2(const char* label, bool* p_visible, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextItemOpen(bool is_open, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextItemStorageID(int storage_id);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TreeNodeGetOpen(int storage_id);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Selectable_1(const char* label, bool selected, int flags, ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Selectable_2(const char* label, bool* p_selected, int flags, ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT ImGuiMultiSelectIO* cabi_ImGui__BeginMultiSelect(int flags, int selection_size, int items_count);
// 来源: 自由函数
MIQT_EXPORT ImGuiMultiSelectIO* cabi_ImGui__EndMultiSelect();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextItemSelectionUserData(int selection_user_data);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemToggledSelection();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginListBox(const char* label, ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndListBox();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__ListBox(const char* label, int* current_item, const char** items, int items_count, int height_in_items);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PlotLines(const char* label, const float* values, int values_count, int values_offset, const char* overlay_text, float scale_min, float scale_max, ImVec2* graph_size, int stride);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PlotHistogram(const char* label, const float* values, int values_count, int values_offset, const char* overlay_text, float scale_min, float scale_max, ImVec2* graph_size, int stride);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Value_1(const char* prefix, bool b);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Value_2(const char* prefix, int v);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Value_3(const char* prefix, unsigned int v);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Value_4(const char* prefix, float v, const char* float_format);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginMenuBar();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndMenuBar();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginMainMenuBar();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndMainMenuBar();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginMenu(const char* label, bool enabled);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndMenu();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__MenuItem_1(const char* label, const char* shortcut, bool selected, bool enabled);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__MenuItem_2(const char* label, const char* shortcut, bool* p_selected, bool enabled);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginTooltip();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndTooltip();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetTooltipV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginItemTooltip();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetItemTooltipV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginPopup(const char* str_id, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginPopupModal(const char* name, bool* p_open, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndPopup();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__OpenPopup_1(const char* str_id, int popup_flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__OpenPopup_2(int id, int popup_flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__OpenPopupOnItemClick(const char* str_id, int popup_flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__CloseCurrentPopup();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginPopupContextItem(const char* str_id, int popup_flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginPopupContextWindow(const char* str_id, int popup_flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginPopupContextVoid(const char* str_id, int popup_flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsPopupOpen(const char* str_id, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginTable(const char* str_id, int columns, int flags, ImVec2* outer_size, float inner_width);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndTable();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableNextRow(int row_flags, float min_row_height);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TableNextColumn();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TableSetColumnIndex(int column_n);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableSetupColumn(const char* label, int flags, float init_width_or_weight, int user_id);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableSetupScrollFreeze(int cols, int rows);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableHeader(const char* label);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableHeadersRow();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableAngledHeadersRow();
// 来源: 自由函数
MIQT_EXPORT ImGuiTableSortSpecs* cabi_ImGui__TableGetSortSpecs();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__TableGetColumnCount();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__TableGetColumnIndex();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__TableGetRowIndex();
// 来源: 自由函数
MIQT_EXPORT const char* cabi_ImGui__TableGetColumnName(int column_n);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__TableGetColumnFlags(int column_n);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableSetColumnEnabled(int column_n, bool v);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__TableGetHoveredColumn();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__TableSetBgColor(int target, int color, int column_n);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Columns(int count, const char* id, bool borders);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__NextColumn();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetColumnIndex();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetColumnWidth(int column_index);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetColumnWidth(int column_index, float width);
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui__GetColumnOffset(int column_index);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetColumnOffset(int column_index, float offset_x);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetColumnsCount();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginTabBar(const char* str_id, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndTabBar();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginTabItem(const char* label, bool* p_open, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndTabItem();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__TabItemButton(const char* label, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetTabItemClosed(const char* tab_or_docked_window_label);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LogToTTY(int auto_open_depth);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LogToFile(int auto_open_depth, const char* filename);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LogToClipboard(int auto_open_depth);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LogFinish();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LogButtons();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LogTextV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginDragDropSource(int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SetDragDropPayload(const char* type, const void* data, size_t sz, int cond);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndDragDropSource();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__BeginDragDropTarget();
// 来源: 自由函数
MIQT_EXPORT ImGuiPayload* cabi_ImGui__AcceptDragDropPayload(const char* type, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndDragDropTarget();
// 来源: 自由函数
MIQT_EXPORT ImGuiPayload* cabi_ImGui__GetDragDropPayload();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__BeginDisabled(bool disabled);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__EndDisabled();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushClipRect(ImVec2* clip_rect_min, ImVec2* clip_rect_max, bool intersect_with_current_clip_rect);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopClipRect();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetItemDefaultFocus();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetKeyboardFocusHere(int offset);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNavCursorVisible(bool visible);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextItemAllowOverlap();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemHovered(int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemActive();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemFocused();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemClicked(int mouse_button);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemVisible();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemEdited();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemActivated();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemDeactivated();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemDeactivatedAfterEdit();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsItemToggledOpen();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsAnyItemHovered();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsAnyItemActive();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsAnyItemFocused();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetItemID();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetItemRectMin();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetItemRectMax();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetItemRectSize();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetItemFlags();
// 来源: 自由函数
MIQT_EXPORT ImGuiViewport* cabi_ImGui__GetMainViewport();
// 来源: 自由函数
MIQT_EXPORT ImDrawList* cabi_ImGui__GetBackgroundDrawList();
// 来源: 自由函数
MIQT_EXPORT ImDrawList* cabi_ImGui__GetForegroundDrawList();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsRectVisible_1(ImVec2* size);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsRectVisible_2(ImVec2* rect_min, ImVec2* rect_max);
// 来源: 自由函数
MIQT_EXPORT double cabi_ImGui__GetTime();
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetFrameCount();
// 来源: 自由函数
MIQT_EXPORT void* cabi_ImGui__GetDrawListSharedData();
// 来源: 自由函数
MIQT_EXPORT const char* cabi_ImGui__GetStyleColorName(int idx);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetStateStorage(ImGuiStorage* storage);
// 来源: 自由函数
MIQT_EXPORT ImGuiStorage* cabi_ImGui__GetStateStorage();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__CalcTextSize(const char* text, const char* text_end, bool hide_text_after_double_hash, float wrap_width);
// 来源: 自由函数
MIQT_EXPORT ImVec4* cabi_ImGui__ColorConvertU32ToFloat4(int in);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__ColorConvertFloat4ToU32(ImVec4* in);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ColorConvertRGBtoHSV(float r, float g, float b, float* out_h, float* out_s, float* out_v);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ColorConvertHSVtoRGB(float h, float s, float v, float* out_r, float* out_g, float* out_b);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsKeyDown(int key);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsKeyPressed(int key, bool repeat);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsKeyReleased(int key);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsKeyChordPressed(int key_chord);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetKeyPressedAmount(int key, float repeat_delay, float rate);
// 来源: 自由函数
MIQT_EXPORT const char* cabi_ImGui__GetKeyName(int key);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__Shortcut(int key_chord, int flags);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextItemShortcut(int key_chord, int flags);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__SetItemKeyOwner(int key);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseDown(int button);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseClicked(int button, bool repeat);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseReleased(int button);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseDoubleClicked(int button);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseReleasedWithDelay(int button, float delay);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetMouseClickedCount(int button);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseHoveringRect(ImVec2* r_min, ImVec2* r_max, bool clip);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMousePosValid(ImVec2* mouse_pos);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsAnyMouseDown();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetMousePos();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetMousePosOnOpeningCurrentPopup();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__IsMouseDragging(int button, float lock_threshold);
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetMouseDragDelta(int button, float lock_threshold);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__ResetMouseDragDelta(int button);
// 来源: 自由函数
MIQT_EXPORT int cabi_ImGui__GetMouseCursor();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetMouseCursor(int cursor_type);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetNextFrameWantCaptureMouse(bool want_capture_mouse);
// 来源: 自由函数
MIQT_EXPORT const char* cabi_ImGui__GetClipboardText();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetClipboardText(const char* text);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LoadIniSettingsFromDisk(const char* ini_filename);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__LoadIniSettingsFromMemory(const char* ini_data, size_t ini_size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SaveIniSettingsToDisk(const char* ini_filename);
// 来源: 自由函数
MIQT_EXPORT const char* cabi_ImGui__SaveIniSettingsToMemory(size_t* out_ini_size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__DebugTextEncoding(const char* text);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__DebugFlashStyleColor(int idx);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__DebugStartItemPicker();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui__DebugCheckVersionAndDataLayout(const char* version_str, size_t sz_io, size_t sz_style, size_t sz_vec2, size_t sz_vec4, size_t sz_drawvert, size_t sz_drawidx);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__DebugLogV(const char* fmt, void* args);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetAllocatorFunctions(void* alloc_func, void* free_func, void* user_data);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__GetAllocatorFunctions(void** p_alloc_func, void** p_free_func, void** p_user_data);
// 来源: 自由函数
MIQT_EXPORT void* cabi_ImGui__MemAlloc(size_t size);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__MemFree(void* ptr);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushFont_2(ImFont* font);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__SetWindowFontScale(float scale);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__Image_2(ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1, ImVec4* tint_col, ImVec4* border_col);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushButtonRepeat(bool repeat);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopButtonRepeat();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PushTabStop(bool tab_stop);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui__PopTabStop();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetContentRegionMax();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetWindowContentRegionMin();
// 来源: 自由函数
MIQT_EXPORT ImVec2* cabi_ImGui__GetWindowContentRegionMax();
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui_ImplWin32_Init(void* hwnd);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui_ImplWin32_InitForOpenGL(void* hwnd);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplWin32_Shutdown();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplWin32_NewFrame();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplWin32_EnableDpiAwareness();
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui_ImplWin32_GetDpiScaleForHwnd(void* hwnd);
// 来源: 自由函数
MIQT_EXPORT float cabi_ImGui_ImplWin32_GetDpiScaleForMonitor(void* monitor);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplWin32_EnableAlphaCompositing(void* hwnd);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui_ImplDX11_Init(void* device, void* device_context);
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplDX11_Shutdown();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplDX11_NewFrame();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplDX11_RenderDrawData(ImDrawData* draw_data);
// 来源: 自由函数
MIQT_EXPORT bool cabi_ImGui_ImplDX11_CreateDeviceObjects();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplDX11_InvalidateDeviceObjects();
// 来源: 自由函数
MIQT_EXPORT void cabi_ImGui_ImplDX11_UpdateTexture(ImTextureData* tex);
// 来源: 自由函数
MIQT_EXPORT intptr_t cabi_ImGui_ImplWin32_WndProcHandler(void* hWnd, int msg, uintptr_t wParam, intptr_t lParam);
#endif // MIQT_TYPES_ONLY

#ifdef __cplusplus
} /* extern C */
#endif

#endif
