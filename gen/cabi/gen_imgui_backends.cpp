#include <imgui_backends.h>
#include "gen_imgui_backends.h"
#include <new>

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
} /* extern C */
#endif

ImVec2* ImVec2_new() {
	return new (std::nothrow) ImVec2();
}

ImVec2* ImVec2_new2(float _x, float _y) {
	return new (std::nothrow) ImVec2(static_cast<float>(_x), static_cast<float>(_y));
}

ImVec2* ImVec2_new3(ImVec2* param1) {
	return new (std::nothrow) ImVec2(*param1);
}

float ImVec2_x(const ImVec2* self) {
	return self->x;
}

void ImVec2_setX(ImVec2* self, float x) {
	self->x = static_cast<float>(x);
}

float ImVec2_y(const ImVec2* self) {
	return self->y;
}

void ImVec2_setY(ImVec2* self, float y) {
	self->y = static_cast<float>(y);
}

float* ImVec2_operatorSubscript(ImVec2* self, unsigned long long idx) {
	return &self->operator[](static_cast<unsigned long long>(idx));
}

float ImVec2_operatorSubscriptWithIdx(const ImVec2* self, unsigned long long idx) {
	return self->operator[](static_cast<unsigned long long>(idx));
}

void ImVec2_operatorAssign(ImVec2* self, ImVec2* param1) {
	self->operator=(*param1);
}

void ImVec2_delete(ImVec2* self) {
	delete self;
}

ImVec4* ImVec4_new() {
	return new (std::nothrow) ImVec4();
}

ImVec4* ImVec4_new2(float _x, float _y, float _z, float _w) {
	return new (std::nothrow) ImVec4(static_cast<float>(_x), static_cast<float>(_y), static_cast<float>(_z), static_cast<float>(_w));
}

ImVec4* ImVec4_new3(ImVec4* param1) {
	return new (std::nothrow) ImVec4(*param1);
}

float ImVec4_x(const ImVec4* self) {
	return self->x;
}

void ImVec4_setX(ImVec4* self, float x) {
	self->x = static_cast<float>(x);
}

float ImVec4_y(const ImVec4* self) {
	return self->y;
}

void ImVec4_setY(ImVec4* self, float y) {
	self->y = static_cast<float>(y);
}

float ImVec4_z(const ImVec4* self) {
	return self->z;
}

void ImVec4_setZ(ImVec4* self, float z) {
	self->z = static_cast<float>(z);
}

float ImVec4_w(const ImVec4* self) {
	return self->w;
}

void ImVec4_setW(ImVec4* self, float w) {
	self->w = static_cast<float>(w);
}

void ImVec4_operatorAssign(ImVec4* self, ImVec4* param1) {
	self->operator=(*param1);
}

void ImVec4_delete(ImVec4* self) {
	delete self;
}

ImTextureRef* ImTextureRef_new() {
	return new (std::nothrow) ImTextureRef();
}

ImTextureRef* ImTextureRef_new2(unsigned long long tex_id) {
	return new (std::nothrow) ImTextureRef(static_cast<unsigned long long>(tex_id));
}

ImTextureRef* ImTextureRef_new3(void* tex_id) {
	return new (std::nothrow) ImTextureRef(static_cast<void*>(tex_id));
}

ImTextureRef* ImTextureRef_new4(ImTextureRef* param1) {
	return new (std::nothrow) ImTextureRef(*param1);
}

unsigned long long ImTextureRef_GetTexID(const ImTextureRef* self) {
	unsigned long long _ret = self->GetTexID();
	return static_cast<unsigned long long>(_ret);
}

ImTextureData* ImTextureRef__TexData(const ImTextureRef* self) {
	return self->_TexData;
}

void ImTextureRef_set_TexData(ImTextureRef* self, ImTextureData* _TexData) {
	self->_TexData = _TexData;
}

unsigned long long ImTextureRef__TexID(const ImTextureRef* self) {
	return self->_TexID;
}

void ImTextureRef_set_TexID(ImTextureRef* self, unsigned long long _TexID) {
	self->_TexID = static_cast<unsigned long long>(_TexID);
}

void ImTextureRef_operatorAssign(ImTextureRef* self, ImTextureRef* param1) {
	self->operator=(*param1);
}

void ImTextureRef_delete(ImTextureRef* self) {
	delete self;
}

ImGuiTableSortSpecs* ImGuiTableSortSpecs_new() {
	return new (std::nothrow) ImGuiTableSortSpecs();
}

ImGuiTableColumnSortSpecs* ImGuiTableSortSpecs_Specs(const ImGuiTableSortSpecs* self) {
	return (ImGuiTableColumnSortSpecs*) self->Specs;
}

void ImGuiTableSortSpecs_setSpecs(ImGuiTableSortSpecs* self, ImGuiTableColumnSortSpecs* Specs) {
	self->Specs = Specs;
}

int ImGuiTableSortSpecs_SpecsCount(const ImGuiTableSortSpecs* self) {
	return self->SpecsCount;
}

void ImGuiTableSortSpecs_setSpecsCount(ImGuiTableSortSpecs* self, int SpecsCount) {
	self->SpecsCount = static_cast<int>(SpecsCount);
}

bool ImGuiTableSortSpecs_SpecsDirty(const ImGuiTableSortSpecs* self) {
	return self->SpecsDirty;
}

void ImGuiTableSortSpecs_setSpecsDirty(ImGuiTableSortSpecs* self, bool SpecsDirty) {
	self->SpecsDirty = SpecsDirty;
}

void ImGuiTableSortSpecs_delete(ImGuiTableSortSpecs* self) {
	delete self;
}

ImGuiTableColumnSortSpecs* ImGuiTableColumnSortSpecs_new() {
	return new (std::nothrow) ImGuiTableColumnSortSpecs();
}

unsigned int ImGuiTableColumnSortSpecs_ColumnUserID(const ImGuiTableColumnSortSpecs* self) {
	return self->ColumnUserID;
}

void ImGuiTableColumnSortSpecs_setColumnUserID(ImGuiTableColumnSortSpecs* self, unsigned int ColumnUserID) {
	self->ColumnUserID = static_cast<unsigned int>(ColumnUserID);
}

short ImGuiTableColumnSortSpecs_ColumnIndex(const ImGuiTableColumnSortSpecs* self) {
	return self->ColumnIndex;
}

void ImGuiTableColumnSortSpecs_setColumnIndex(ImGuiTableColumnSortSpecs* self, short ColumnIndex) {
	self->ColumnIndex = static_cast<short>(ColumnIndex);
}

short ImGuiTableColumnSortSpecs_SortOrder(const ImGuiTableColumnSortSpecs* self) {
	return self->SortOrder;
}

void ImGuiTableColumnSortSpecs_setSortOrder(ImGuiTableColumnSortSpecs* self, short SortOrder) {
	self->SortOrder = static_cast<short>(SortOrder);
}

int ImGuiTableColumnSortSpecs_SortDirection(const ImGuiTableColumnSortSpecs* self) {
	ImGuiSortDirection SortDirection_ret = self->SortDirection;
	return static_cast<int>(SortDirection_ret);
}

void ImGuiTableColumnSortSpecs_setSortDirection(ImGuiTableColumnSortSpecs* self, int SortDirection) {
	self->SortDirection = static_cast<ImGuiSortDirection>(SortDirection);
}

void ImGuiTableColumnSortSpecs_delete(ImGuiTableColumnSortSpecs* self) {
	delete self;
}

void ImNewWrapper_delete(ImNewWrapper* self) {
	delete self;
}

ImGuiStyle* ImGuiStyle_new() {
	return new (std::nothrow) ImGuiStyle();
}

float ImGuiStyle_FontSizeBase(const ImGuiStyle* self) {
	return self->FontSizeBase;
}

void ImGuiStyle_setFontSizeBase(ImGuiStyle* self, float FontSizeBase) {
	self->FontSizeBase = static_cast<float>(FontSizeBase);
}

float ImGuiStyle_FontScaleMain(const ImGuiStyle* self) {
	return self->FontScaleMain;
}

void ImGuiStyle_setFontScaleMain(ImGuiStyle* self, float FontScaleMain) {
	self->FontScaleMain = static_cast<float>(FontScaleMain);
}

float ImGuiStyle_FontScaleDpi(const ImGuiStyle* self) {
	return self->FontScaleDpi;
}

void ImGuiStyle_setFontScaleDpi(ImGuiStyle* self, float FontScaleDpi) {
	self->FontScaleDpi = static_cast<float>(FontScaleDpi);
}

float ImGuiStyle_Alpha(const ImGuiStyle* self) {
	return self->Alpha;
}

void ImGuiStyle_setAlpha(ImGuiStyle* self, float Alpha) {
	self->Alpha = static_cast<float>(Alpha);
}

float ImGuiStyle_DisabledAlpha(const ImGuiStyle* self) {
	return self->DisabledAlpha;
}

void ImGuiStyle_setDisabledAlpha(ImGuiStyle* self, float DisabledAlpha) {
	self->DisabledAlpha = static_cast<float>(DisabledAlpha);
}

ImVec2* ImGuiStyle_WindowPadding(const ImGuiStyle* self) {
	return new ImVec2(self->WindowPadding);
}

void ImGuiStyle_setWindowPadding(ImGuiStyle* self, ImVec2* WindowPadding) {
	self->WindowPadding = *WindowPadding;
}

float ImGuiStyle_WindowRounding(const ImGuiStyle* self) {
	return self->WindowRounding;
}

void ImGuiStyle_setWindowRounding(ImGuiStyle* self, float WindowRounding) {
	self->WindowRounding = static_cast<float>(WindowRounding);
}

float ImGuiStyle_WindowBorderSize(const ImGuiStyle* self) {
	return self->WindowBorderSize;
}

void ImGuiStyle_setWindowBorderSize(ImGuiStyle* self, float WindowBorderSize) {
	self->WindowBorderSize = static_cast<float>(WindowBorderSize);
}

float ImGuiStyle_WindowBorderHoverPadding(const ImGuiStyle* self) {
	return self->WindowBorderHoverPadding;
}

void ImGuiStyle_setWindowBorderHoverPadding(ImGuiStyle* self, float WindowBorderHoverPadding) {
	self->WindowBorderHoverPadding = static_cast<float>(WindowBorderHoverPadding);
}

ImVec2* ImGuiStyle_WindowMinSize(const ImGuiStyle* self) {
	return new ImVec2(self->WindowMinSize);
}

void ImGuiStyle_setWindowMinSize(ImGuiStyle* self, ImVec2* WindowMinSize) {
	self->WindowMinSize = *WindowMinSize;
}

ImVec2* ImGuiStyle_WindowTitleAlign(const ImGuiStyle* self) {
	return new ImVec2(self->WindowTitleAlign);
}

void ImGuiStyle_setWindowTitleAlign(ImGuiStyle* self, ImVec2* WindowTitleAlign) {
	self->WindowTitleAlign = *WindowTitleAlign;
}

int ImGuiStyle_WindowMenuButtonPosition(const ImGuiStyle* self) {
	ImGuiDir WindowMenuButtonPosition_ret = self->WindowMenuButtonPosition;
	return static_cast<int>(WindowMenuButtonPosition_ret);
}

void ImGuiStyle_setWindowMenuButtonPosition(ImGuiStyle* self, int WindowMenuButtonPosition) {
	self->WindowMenuButtonPosition = static_cast<ImGuiDir>(WindowMenuButtonPosition);
}

float ImGuiStyle_ChildRounding(const ImGuiStyle* self) {
	return self->ChildRounding;
}

void ImGuiStyle_setChildRounding(ImGuiStyle* self, float ChildRounding) {
	self->ChildRounding = static_cast<float>(ChildRounding);
}

float ImGuiStyle_ChildBorderSize(const ImGuiStyle* self) {
	return self->ChildBorderSize;
}

void ImGuiStyle_setChildBorderSize(ImGuiStyle* self, float ChildBorderSize) {
	self->ChildBorderSize = static_cast<float>(ChildBorderSize);
}

float ImGuiStyle_PopupRounding(const ImGuiStyle* self) {
	return self->PopupRounding;
}

void ImGuiStyle_setPopupRounding(ImGuiStyle* self, float PopupRounding) {
	self->PopupRounding = static_cast<float>(PopupRounding);
}

float ImGuiStyle_PopupBorderSize(const ImGuiStyle* self) {
	return self->PopupBorderSize;
}

void ImGuiStyle_setPopupBorderSize(ImGuiStyle* self, float PopupBorderSize) {
	self->PopupBorderSize = static_cast<float>(PopupBorderSize);
}

ImVec2* ImGuiStyle_FramePadding(const ImGuiStyle* self) {
	return new ImVec2(self->FramePadding);
}

void ImGuiStyle_setFramePadding(ImGuiStyle* self, ImVec2* FramePadding) {
	self->FramePadding = *FramePadding;
}

float ImGuiStyle_FrameRounding(const ImGuiStyle* self) {
	return self->FrameRounding;
}

void ImGuiStyle_setFrameRounding(ImGuiStyle* self, float FrameRounding) {
	self->FrameRounding = static_cast<float>(FrameRounding);
}

float ImGuiStyle_FrameBorderSize(const ImGuiStyle* self) {
	return self->FrameBorderSize;
}

void ImGuiStyle_setFrameBorderSize(ImGuiStyle* self, float FrameBorderSize) {
	self->FrameBorderSize = static_cast<float>(FrameBorderSize);
}

ImVec2* ImGuiStyle_ItemSpacing(const ImGuiStyle* self) {
	return new ImVec2(self->ItemSpacing);
}

void ImGuiStyle_setItemSpacing(ImGuiStyle* self, ImVec2* ItemSpacing) {
	self->ItemSpacing = *ItemSpacing;
}

ImVec2* ImGuiStyle_ItemInnerSpacing(const ImGuiStyle* self) {
	return new ImVec2(self->ItemInnerSpacing);
}

void ImGuiStyle_setItemInnerSpacing(ImGuiStyle* self, ImVec2* ItemInnerSpacing) {
	self->ItemInnerSpacing = *ItemInnerSpacing;
}

ImVec2* ImGuiStyle_CellPadding(const ImGuiStyle* self) {
	return new ImVec2(self->CellPadding);
}

void ImGuiStyle_setCellPadding(ImGuiStyle* self, ImVec2* CellPadding) {
	self->CellPadding = *CellPadding;
}

ImVec2* ImGuiStyle_TouchExtraPadding(const ImGuiStyle* self) {
	return new ImVec2(self->TouchExtraPadding);
}

void ImGuiStyle_setTouchExtraPadding(ImGuiStyle* self, ImVec2* TouchExtraPadding) {
	self->TouchExtraPadding = *TouchExtraPadding;
}

float ImGuiStyle_IndentSpacing(const ImGuiStyle* self) {
	return self->IndentSpacing;
}

void ImGuiStyle_setIndentSpacing(ImGuiStyle* self, float IndentSpacing) {
	self->IndentSpacing = static_cast<float>(IndentSpacing);
}

float ImGuiStyle_ColumnsMinSpacing(const ImGuiStyle* self) {
	return self->ColumnsMinSpacing;
}

void ImGuiStyle_setColumnsMinSpacing(ImGuiStyle* self, float ColumnsMinSpacing) {
	self->ColumnsMinSpacing = static_cast<float>(ColumnsMinSpacing);
}

float ImGuiStyle_ScrollbarSize(const ImGuiStyle* self) {
	return self->ScrollbarSize;
}

void ImGuiStyle_setScrollbarSize(ImGuiStyle* self, float ScrollbarSize) {
	self->ScrollbarSize = static_cast<float>(ScrollbarSize);
}

float ImGuiStyle_ScrollbarRounding(const ImGuiStyle* self) {
	return self->ScrollbarRounding;
}

void ImGuiStyle_setScrollbarRounding(ImGuiStyle* self, float ScrollbarRounding) {
	self->ScrollbarRounding = static_cast<float>(ScrollbarRounding);
}

float ImGuiStyle_ScrollbarPadding(const ImGuiStyle* self) {
	return self->ScrollbarPadding;
}

void ImGuiStyle_setScrollbarPadding(ImGuiStyle* self, float ScrollbarPadding) {
	self->ScrollbarPadding = static_cast<float>(ScrollbarPadding);
}

float ImGuiStyle_GrabMinSize(const ImGuiStyle* self) {
	return self->GrabMinSize;
}

void ImGuiStyle_setGrabMinSize(ImGuiStyle* self, float GrabMinSize) {
	self->GrabMinSize = static_cast<float>(GrabMinSize);
}

float ImGuiStyle_GrabRounding(const ImGuiStyle* self) {
	return self->GrabRounding;
}

void ImGuiStyle_setGrabRounding(ImGuiStyle* self, float GrabRounding) {
	self->GrabRounding = static_cast<float>(GrabRounding);
}

float ImGuiStyle_LogSliderDeadzone(const ImGuiStyle* self) {
	return self->LogSliderDeadzone;
}

void ImGuiStyle_setLogSliderDeadzone(ImGuiStyle* self, float LogSliderDeadzone) {
	self->LogSliderDeadzone = static_cast<float>(LogSliderDeadzone);
}

float ImGuiStyle_ImageRounding(const ImGuiStyle* self) {
	return self->ImageRounding;
}

void ImGuiStyle_setImageRounding(ImGuiStyle* self, float ImageRounding) {
	self->ImageRounding = static_cast<float>(ImageRounding);
}

float ImGuiStyle_ImageBorderSize(const ImGuiStyle* self) {
	return self->ImageBorderSize;
}

void ImGuiStyle_setImageBorderSize(ImGuiStyle* self, float ImageBorderSize) {
	self->ImageBorderSize = static_cast<float>(ImageBorderSize);
}

float ImGuiStyle_TabRounding(const ImGuiStyle* self) {
	return self->TabRounding;
}

void ImGuiStyle_setTabRounding(ImGuiStyle* self, float TabRounding) {
	self->TabRounding = static_cast<float>(TabRounding);
}

float ImGuiStyle_TabBorderSize(const ImGuiStyle* self) {
	return self->TabBorderSize;
}

void ImGuiStyle_setTabBorderSize(ImGuiStyle* self, float TabBorderSize) {
	self->TabBorderSize = static_cast<float>(TabBorderSize);
}

float ImGuiStyle_TabMinWidthBase(const ImGuiStyle* self) {
	return self->TabMinWidthBase;
}

void ImGuiStyle_setTabMinWidthBase(ImGuiStyle* self, float TabMinWidthBase) {
	self->TabMinWidthBase = static_cast<float>(TabMinWidthBase);
}

float ImGuiStyle_TabMinWidthShrink(const ImGuiStyle* self) {
	return self->TabMinWidthShrink;
}

void ImGuiStyle_setTabMinWidthShrink(ImGuiStyle* self, float TabMinWidthShrink) {
	self->TabMinWidthShrink = static_cast<float>(TabMinWidthShrink);
}

float ImGuiStyle_TabCloseButtonMinWidthSelected(const ImGuiStyle* self) {
	return self->TabCloseButtonMinWidthSelected;
}

void ImGuiStyle_setTabCloseButtonMinWidthSelected(ImGuiStyle* self, float TabCloseButtonMinWidthSelected) {
	self->TabCloseButtonMinWidthSelected = static_cast<float>(TabCloseButtonMinWidthSelected);
}

float ImGuiStyle_TabCloseButtonMinWidthUnselected(const ImGuiStyle* self) {
	return self->TabCloseButtonMinWidthUnselected;
}

void ImGuiStyle_setTabCloseButtonMinWidthUnselected(ImGuiStyle* self, float TabCloseButtonMinWidthUnselected) {
	self->TabCloseButtonMinWidthUnselected = static_cast<float>(TabCloseButtonMinWidthUnselected);
}

float ImGuiStyle_TabBarBorderSize(const ImGuiStyle* self) {
	return self->TabBarBorderSize;
}

void ImGuiStyle_setTabBarBorderSize(ImGuiStyle* self, float TabBarBorderSize) {
	self->TabBarBorderSize = static_cast<float>(TabBarBorderSize);
}

float ImGuiStyle_TabBarOverlineSize(const ImGuiStyle* self) {
	return self->TabBarOverlineSize;
}

void ImGuiStyle_setTabBarOverlineSize(ImGuiStyle* self, float TabBarOverlineSize) {
	self->TabBarOverlineSize = static_cast<float>(TabBarOverlineSize);
}

float ImGuiStyle_TableAngledHeadersAngle(const ImGuiStyle* self) {
	return self->TableAngledHeadersAngle;
}

void ImGuiStyle_setTableAngledHeadersAngle(ImGuiStyle* self, float TableAngledHeadersAngle) {
	self->TableAngledHeadersAngle = static_cast<float>(TableAngledHeadersAngle);
}

ImVec2* ImGuiStyle_TableAngledHeadersTextAlign(const ImGuiStyle* self) {
	return new ImVec2(self->TableAngledHeadersTextAlign);
}

void ImGuiStyle_setTableAngledHeadersTextAlign(ImGuiStyle* self, ImVec2* TableAngledHeadersTextAlign) {
	self->TableAngledHeadersTextAlign = *TableAngledHeadersTextAlign;
}

int ImGuiStyle_TreeLinesFlags(const ImGuiStyle* self) {
	return self->TreeLinesFlags;
}

void ImGuiStyle_setTreeLinesFlags(ImGuiStyle* self, int TreeLinesFlags) {
	self->TreeLinesFlags = static_cast<int>(TreeLinesFlags);
}

float ImGuiStyle_TreeLinesSize(const ImGuiStyle* self) {
	return self->TreeLinesSize;
}

void ImGuiStyle_setTreeLinesSize(ImGuiStyle* self, float TreeLinesSize) {
	self->TreeLinesSize = static_cast<float>(TreeLinesSize);
}

float ImGuiStyle_TreeLinesRounding(const ImGuiStyle* self) {
	return self->TreeLinesRounding;
}

void ImGuiStyle_setTreeLinesRounding(ImGuiStyle* self, float TreeLinesRounding) {
	self->TreeLinesRounding = static_cast<float>(TreeLinesRounding);
}

float ImGuiStyle_DragDropTargetRounding(const ImGuiStyle* self) {
	return self->DragDropTargetRounding;
}

void ImGuiStyle_setDragDropTargetRounding(ImGuiStyle* self, float DragDropTargetRounding) {
	self->DragDropTargetRounding = static_cast<float>(DragDropTargetRounding);
}

float ImGuiStyle_DragDropTargetBorderSize(const ImGuiStyle* self) {
	return self->DragDropTargetBorderSize;
}

void ImGuiStyle_setDragDropTargetBorderSize(ImGuiStyle* self, float DragDropTargetBorderSize) {
	self->DragDropTargetBorderSize = static_cast<float>(DragDropTargetBorderSize);
}

float ImGuiStyle_DragDropTargetPadding(const ImGuiStyle* self) {
	return self->DragDropTargetPadding;
}

void ImGuiStyle_setDragDropTargetPadding(ImGuiStyle* self, float DragDropTargetPadding) {
	self->DragDropTargetPadding = static_cast<float>(DragDropTargetPadding);
}

float ImGuiStyle_ColorMarkerSize(const ImGuiStyle* self) {
	return self->ColorMarkerSize;
}

void ImGuiStyle_setColorMarkerSize(ImGuiStyle* self, float ColorMarkerSize) {
	self->ColorMarkerSize = static_cast<float>(ColorMarkerSize);
}

int ImGuiStyle_ColorButtonPosition(const ImGuiStyle* self) {
	ImGuiDir ColorButtonPosition_ret = self->ColorButtonPosition;
	return static_cast<int>(ColorButtonPosition_ret);
}

void ImGuiStyle_setColorButtonPosition(ImGuiStyle* self, int ColorButtonPosition) {
	self->ColorButtonPosition = static_cast<ImGuiDir>(ColorButtonPosition);
}

ImVec2* ImGuiStyle_ButtonTextAlign(const ImGuiStyle* self) {
	return new ImVec2(self->ButtonTextAlign);
}

void ImGuiStyle_setButtonTextAlign(ImGuiStyle* self, ImVec2* ButtonTextAlign) {
	self->ButtonTextAlign = *ButtonTextAlign;
}

ImVec2* ImGuiStyle_SelectableTextAlign(const ImGuiStyle* self) {
	return new ImVec2(self->SelectableTextAlign);
}

void ImGuiStyle_setSelectableTextAlign(ImGuiStyle* self, ImVec2* SelectableTextAlign) {
	self->SelectableTextAlign = *SelectableTextAlign;
}

float ImGuiStyle_SeparatorSize(const ImGuiStyle* self) {
	return self->SeparatorSize;
}

void ImGuiStyle_setSeparatorSize(ImGuiStyle* self, float SeparatorSize) {
	self->SeparatorSize = static_cast<float>(SeparatorSize);
}

float ImGuiStyle_SeparatorTextBorderSize(const ImGuiStyle* self) {
	return self->SeparatorTextBorderSize;
}

void ImGuiStyle_setSeparatorTextBorderSize(ImGuiStyle* self, float SeparatorTextBorderSize) {
	self->SeparatorTextBorderSize = static_cast<float>(SeparatorTextBorderSize);
}

ImVec2* ImGuiStyle_SeparatorTextAlign(const ImGuiStyle* self) {
	return new ImVec2(self->SeparatorTextAlign);
}

void ImGuiStyle_setSeparatorTextAlign(ImGuiStyle* self, ImVec2* SeparatorTextAlign) {
	self->SeparatorTextAlign = *SeparatorTextAlign;
}

ImVec2* ImGuiStyle_SeparatorTextPadding(const ImGuiStyle* self) {
	return new ImVec2(self->SeparatorTextPadding);
}

void ImGuiStyle_setSeparatorTextPadding(ImGuiStyle* self, ImVec2* SeparatorTextPadding) {
	self->SeparatorTextPadding = *SeparatorTextPadding;
}

ImVec2* ImGuiStyle_DisplayWindowPadding(const ImGuiStyle* self) {
	return new ImVec2(self->DisplayWindowPadding);
}

void ImGuiStyle_setDisplayWindowPadding(ImGuiStyle* self, ImVec2* DisplayWindowPadding) {
	self->DisplayWindowPadding = *DisplayWindowPadding;
}

ImVec2* ImGuiStyle_DisplaySafeAreaPadding(const ImGuiStyle* self) {
	return new ImVec2(self->DisplaySafeAreaPadding);
}

void ImGuiStyle_setDisplaySafeAreaPadding(ImGuiStyle* self, ImVec2* DisplaySafeAreaPadding) {
	self->DisplaySafeAreaPadding = *DisplaySafeAreaPadding;
}

float ImGuiStyle_MouseCursorScale(const ImGuiStyle* self) {
	return self->MouseCursorScale;
}

void ImGuiStyle_setMouseCursorScale(ImGuiStyle* self, float MouseCursorScale) {
	self->MouseCursorScale = static_cast<float>(MouseCursorScale);
}

bool ImGuiStyle_AntiAliasedLines(const ImGuiStyle* self) {
	return self->AntiAliasedLines;
}

void ImGuiStyle_setAntiAliasedLines(ImGuiStyle* self, bool AntiAliasedLines) {
	self->AntiAliasedLines = AntiAliasedLines;
}

bool ImGuiStyle_AntiAliasedLinesUseTex(const ImGuiStyle* self) {
	return self->AntiAliasedLinesUseTex;
}

void ImGuiStyle_setAntiAliasedLinesUseTex(ImGuiStyle* self, bool AntiAliasedLinesUseTex) {
	self->AntiAliasedLinesUseTex = AntiAliasedLinesUseTex;
}

bool ImGuiStyle_AntiAliasedFill(const ImGuiStyle* self) {
	return self->AntiAliasedFill;
}

void ImGuiStyle_setAntiAliasedFill(ImGuiStyle* self, bool AntiAliasedFill) {
	self->AntiAliasedFill = AntiAliasedFill;
}

float ImGuiStyle_CurveTessellationTol(const ImGuiStyle* self) {
	return self->CurveTessellationTol;
}

void ImGuiStyle_setCurveTessellationTol(ImGuiStyle* self, float CurveTessellationTol) {
	self->CurveTessellationTol = static_cast<float>(CurveTessellationTol);
}

float ImGuiStyle_CircleTessellationMaxError(const ImGuiStyle* self) {
	return self->CircleTessellationMaxError;
}

void ImGuiStyle_setCircleTessellationMaxError(ImGuiStyle* self, float CircleTessellationMaxError) {
	self->CircleTessellationMaxError = static_cast<float>(CircleTessellationMaxError);
}

ImVec4* ImGuiStyle_Colors(const ImGuiStyle* self) {
	return (ImVec4*)self->Colors;
}

void ImGuiStyle_setColors(ImGuiStyle* self, const ImVec4* Colors) {
	memcpy(self->Colors, Colors, sizeof(self->Colors));
}

float ImGuiStyle_HoverStationaryDelay(const ImGuiStyle* self) {
	return self->HoverStationaryDelay;
}

void ImGuiStyle_setHoverStationaryDelay(ImGuiStyle* self, float HoverStationaryDelay) {
	self->HoverStationaryDelay = static_cast<float>(HoverStationaryDelay);
}

float ImGuiStyle_HoverDelayShort(const ImGuiStyle* self) {
	return self->HoverDelayShort;
}

void ImGuiStyle_setHoverDelayShort(ImGuiStyle* self, float HoverDelayShort) {
	self->HoverDelayShort = static_cast<float>(HoverDelayShort);
}

float ImGuiStyle_HoverDelayNormal(const ImGuiStyle* self) {
	return self->HoverDelayNormal;
}

void ImGuiStyle_setHoverDelayNormal(ImGuiStyle* self, float HoverDelayNormal) {
	self->HoverDelayNormal = static_cast<float>(HoverDelayNormal);
}

int ImGuiStyle_HoverFlagsForTooltipMouse(const ImGuiStyle* self) {
	return self->HoverFlagsForTooltipMouse;
}

void ImGuiStyle_setHoverFlagsForTooltipMouse(ImGuiStyle* self, int HoverFlagsForTooltipMouse) {
	self->HoverFlagsForTooltipMouse = static_cast<int>(HoverFlagsForTooltipMouse);
}

int ImGuiStyle_HoverFlagsForTooltipNav(const ImGuiStyle* self) {
	return self->HoverFlagsForTooltipNav;
}

void ImGuiStyle_setHoverFlagsForTooltipNav(ImGuiStyle* self, int HoverFlagsForTooltipNav) {
	self->HoverFlagsForTooltipNav = static_cast<int>(HoverFlagsForTooltipNav);
}

float ImGuiStyle__MainScale(const ImGuiStyle* self) {
	return self->_MainScale;
}

void ImGuiStyle_set_MainScale(ImGuiStyle* self, float _MainScale) {
	self->_MainScale = static_cast<float>(_MainScale);
}

float ImGuiStyle__NextFrameFontSizeBase(const ImGuiStyle* self) {
	return self->_NextFrameFontSizeBase;
}

void ImGuiStyle_set_NextFrameFontSizeBase(ImGuiStyle* self, float _NextFrameFontSizeBase) {
	self->_NextFrameFontSizeBase = static_cast<float>(_NextFrameFontSizeBase);
}

void ImGuiStyle_ScaleAllSizes(ImGuiStyle* self, float scale_factor) {
	self->ScaleAllSizes(static_cast<float>(scale_factor));
}

void ImGuiStyle_delete(ImGuiStyle* self) {
	delete self;
}

ImGuiKeyData* ImGuiKeyData_new(ImGuiKeyData* param1) {
	return new (std::nothrow) ImGuiKeyData(*param1);
}

bool ImGuiKeyData_Down(const ImGuiKeyData* self) {
	return self->Down;
}

void ImGuiKeyData_setDown(ImGuiKeyData* self, bool Down) {
	self->Down = Down;
}

float ImGuiKeyData_DownDuration(const ImGuiKeyData* self) {
	return self->DownDuration;
}

void ImGuiKeyData_setDownDuration(ImGuiKeyData* self, float DownDuration) {
	self->DownDuration = static_cast<float>(DownDuration);
}

float ImGuiKeyData_DownDurationPrev(const ImGuiKeyData* self) {
	return self->DownDurationPrev;
}

void ImGuiKeyData_setDownDurationPrev(ImGuiKeyData* self, float DownDurationPrev) {
	self->DownDurationPrev = static_cast<float>(DownDurationPrev);
}

float ImGuiKeyData_AnalogValue(const ImGuiKeyData* self) {
	return self->AnalogValue;
}

void ImGuiKeyData_setAnalogValue(ImGuiKeyData* self, float AnalogValue) {
	self->AnalogValue = static_cast<float>(AnalogValue);
}

void ImGuiKeyData_operatorAssign(ImGuiKeyData* self, ImGuiKeyData* param1) {
	self->operator=(*param1);
}

void ImGuiKeyData_delete(ImGuiKeyData* self) {
	delete self;
}

ImGuiIO* ImGuiIO_new() {
	return new (std::nothrow) ImGuiIO();
}

ImGuiIO* ImGuiIO_new2(ImGuiIO* param1) {
	return new (std::nothrow) ImGuiIO(*param1);
}

int ImGuiIO_ConfigFlags(const ImGuiIO* self) {
	return self->ConfigFlags;
}

void ImGuiIO_setConfigFlags(ImGuiIO* self, int ConfigFlags) {
	self->ConfigFlags = static_cast<int>(ConfigFlags);
}

int ImGuiIO_BackendFlags(const ImGuiIO* self) {
	return self->BackendFlags;
}

void ImGuiIO_setBackendFlags(ImGuiIO* self, int BackendFlags) {
	self->BackendFlags = static_cast<int>(BackendFlags);
}

ImVec2* ImGuiIO_DisplaySize(const ImGuiIO* self) {
	return new ImVec2(self->DisplaySize);
}

void ImGuiIO_setDisplaySize(ImGuiIO* self, ImVec2* DisplaySize) {
	self->DisplaySize = *DisplaySize;
}

ImVec2* ImGuiIO_DisplayFramebufferScale(const ImGuiIO* self) {
	return new ImVec2(self->DisplayFramebufferScale);
}

void ImGuiIO_setDisplayFramebufferScale(ImGuiIO* self, ImVec2* DisplayFramebufferScale) {
	self->DisplayFramebufferScale = *DisplayFramebufferScale;
}

float ImGuiIO_DeltaTime(const ImGuiIO* self) {
	return self->DeltaTime;
}

void ImGuiIO_setDeltaTime(ImGuiIO* self, float DeltaTime) {
	self->DeltaTime = static_cast<float>(DeltaTime);
}

float ImGuiIO_IniSavingRate(const ImGuiIO* self) {
	return self->IniSavingRate;
}

void ImGuiIO_setIniSavingRate(ImGuiIO* self, float IniSavingRate) {
	self->IniSavingRate = static_cast<float>(IniSavingRate);
}

const char* ImGuiIO_IniFilename(const ImGuiIO* self) {
	return (const char*) self->IniFilename;
}

void ImGuiIO_setIniFilename(ImGuiIO* self, const char* IniFilename) {
	self->IniFilename = IniFilename;
}

const char* ImGuiIO_LogFilename(const ImGuiIO* self) {
	return (const char*) self->LogFilename;
}

void ImGuiIO_setLogFilename(ImGuiIO* self, const char* LogFilename) {
	self->LogFilename = LogFilename;
}

ImFontAtlas* ImGuiIO_Fonts(const ImGuiIO* self) {
	return self->Fonts;
}

void ImGuiIO_setFonts(ImGuiIO* self, ImFontAtlas* Fonts) {
	self->Fonts = Fonts;
}

ImFont* ImGuiIO_FontDefault(const ImGuiIO* self) {
	return self->FontDefault;
}

void ImGuiIO_setFontDefault(ImGuiIO* self, ImFont* FontDefault) {
	self->FontDefault = FontDefault;
}

bool ImGuiIO_FontAllowUserScaling(const ImGuiIO* self) {
	return self->FontAllowUserScaling;
}

void ImGuiIO_setFontAllowUserScaling(ImGuiIO* self, bool FontAllowUserScaling) {
	self->FontAllowUserScaling = FontAllowUserScaling;
}

bool ImGuiIO_ConfigNavSwapGamepadButtons(const ImGuiIO* self) {
	return self->ConfigNavSwapGamepadButtons;
}

void ImGuiIO_setConfigNavSwapGamepadButtons(ImGuiIO* self, bool ConfigNavSwapGamepadButtons) {
	self->ConfigNavSwapGamepadButtons = ConfigNavSwapGamepadButtons;
}

bool ImGuiIO_ConfigNavMoveSetMousePos(const ImGuiIO* self) {
	return self->ConfigNavMoveSetMousePos;
}

void ImGuiIO_setConfigNavMoveSetMousePos(ImGuiIO* self, bool ConfigNavMoveSetMousePos) {
	self->ConfigNavMoveSetMousePos = ConfigNavMoveSetMousePos;
}

bool ImGuiIO_ConfigNavCaptureKeyboard(const ImGuiIO* self) {
	return self->ConfigNavCaptureKeyboard;
}

void ImGuiIO_setConfigNavCaptureKeyboard(ImGuiIO* self, bool ConfigNavCaptureKeyboard) {
	self->ConfigNavCaptureKeyboard = ConfigNavCaptureKeyboard;
}

bool ImGuiIO_ConfigNavEscapeClearFocusItem(const ImGuiIO* self) {
	return self->ConfigNavEscapeClearFocusItem;
}

void ImGuiIO_setConfigNavEscapeClearFocusItem(ImGuiIO* self, bool ConfigNavEscapeClearFocusItem) {
	self->ConfigNavEscapeClearFocusItem = ConfigNavEscapeClearFocusItem;
}

bool ImGuiIO_ConfigNavEscapeClearFocusWindow(const ImGuiIO* self) {
	return self->ConfigNavEscapeClearFocusWindow;
}

void ImGuiIO_setConfigNavEscapeClearFocusWindow(ImGuiIO* self, bool ConfigNavEscapeClearFocusWindow) {
	self->ConfigNavEscapeClearFocusWindow = ConfigNavEscapeClearFocusWindow;
}

bool ImGuiIO_ConfigNavCursorVisibleAuto(const ImGuiIO* self) {
	return self->ConfigNavCursorVisibleAuto;
}

void ImGuiIO_setConfigNavCursorVisibleAuto(ImGuiIO* self, bool ConfigNavCursorVisibleAuto) {
	self->ConfigNavCursorVisibleAuto = ConfigNavCursorVisibleAuto;
}

bool ImGuiIO_ConfigNavCursorVisibleAlways(const ImGuiIO* self) {
	return self->ConfigNavCursorVisibleAlways;
}

void ImGuiIO_setConfigNavCursorVisibleAlways(ImGuiIO* self, bool ConfigNavCursorVisibleAlways) {
	self->ConfigNavCursorVisibleAlways = ConfigNavCursorVisibleAlways;
}

bool ImGuiIO_MouseDrawCursor(const ImGuiIO* self) {
	return self->MouseDrawCursor;
}

void ImGuiIO_setMouseDrawCursor(ImGuiIO* self, bool MouseDrawCursor) {
	self->MouseDrawCursor = MouseDrawCursor;
}

bool ImGuiIO_ConfigMacOSXBehaviors(const ImGuiIO* self) {
	return self->ConfigMacOSXBehaviors;
}

void ImGuiIO_setConfigMacOSXBehaviors(ImGuiIO* self, bool ConfigMacOSXBehaviors) {
	self->ConfigMacOSXBehaviors = ConfigMacOSXBehaviors;
}

bool ImGuiIO_ConfigInputTrickleEventQueue(const ImGuiIO* self) {
	return self->ConfigInputTrickleEventQueue;
}

void ImGuiIO_setConfigInputTrickleEventQueue(ImGuiIO* self, bool ConfigInputTrickleEventQueue) {
	self->ConfigInputTrickleEventQueue = ConfigInputTrickleEventQueue;
}

bool ImGuiIO_ConfigInputTextCursorBlink(const ImGuiIO* self) {
	return self->ConfigInputTextCursorBlink;
}

void ImGuiIO_setConfigInputTextCursorBlink(ImGuiIO* self, bool ConfigInputTextCursorBlink) {
	self->ConfigInputTextCursorBlink = ConfigInputTextCursorBlink;
}

bool ImGuiIO_ConfigInputTextEnterKeepActive(const ImGuiIO* self) {
	return self->ConfigInputTextEnterKeepActive;
}

void ImGuiIO_setConfigInputTextEnterKeepActive(ImGuiIO* self, bool ConfigInputTextEnterKeepActive) {
	self->ConfigInputTextEnterKeepActive = ConfigInputTextEnterKeepActive;
}

bool ImGuiIO_ConfigDragClickToInputText(const ImGuiIO* self) {
	return self->ConfigDragClickToInputText;
}

void ImGuiIO_setConfigDragClickToInputText(ImGuiIO* self, bool ConfigDragClickToInputText) {
	self->ConfigDragClickToInputText = ConfigDragClickToInputText;
}

bool ImGuiIO_ConfigWindowsResizeFromEdges(const ImGuiIO* self) {
	return self->ConfigWindowsResizeFromEdges;
}

void ImGuiIO_setConfigWindowsResizeFromEdges(ImGuiIO* self, bool ConfigWindowsResizeFromEdges) {
	self->ConfigWindowsResizeFromEdges = ConfigWindowsResizeFromEdges;
}

bool ImGuiIO_ConfigWindowsMoveFromTitleBarOnly(const ImGuiIO* self) {
	return self->ConfigWindowsMoveFromTitleBarOnly;
}

void ImGuiIO_setConfigWindowsMoveFromTitleBarOnly(ImGuiIO* self, bool ConfigWindowsMoveFromTitleBarOnly) {
	self->ConfigWindowsMoveFromTitleBarOnly = ConfigWindowsMoveFromTitleBarOnly;
}

bool ImGuiIO_ConfigWindowsCopyContentsWithCtrlC(const ImGuiIO* self) {
	return self->ConfigWindowsCopyContentsWithCtrlC;
}

void ImGuiIO_setConfigWindowsCopyContentsWithCtrlC(ImGuiIO* self, bool ConfigWindowsCopyContentsWithCtrlC) {
	self->ConfigWindowsCopyContentsWithCtrlC = ConfigWindowsCopyContentsWithCtrlC;
}

bool ImGuiIO_ConfigScrollbarScrollByPage(const ImGuiIO* self) {
	return self->ConfigScrollbarScrollByPage;
}

void ImGuiIO_setConfigScrollbarScrollByPage(ImGuiIO* self, bool ConfigScrollbarScrollByPage) {
	self->ConfigScrollbarScrollByPage = ConfigScrollbarScrollByPage;
}

float ImGuiIO_ConfigMemoryCompactTimer(const ImGuiIO* self) {
	return self->ConfigMemoryCompactTimer;
}

void ImGuiIO_setConfigMemoryCompactTimer(ImGuiIO* self, float ConfigMemoryCompactTimer) {
	self->ConfigMemoryCompactTimer = static_cast<float>(ConfigMemoryCompactTimer);
}

float ImGuiIO_MouseDoubleClickTime(const ImGuiIO* self) {
	return self->MouseDoubleClickTime;
}

void ImGuiIO_setMouseDoubleClickTime(ImGuiIO* self, float MouseDoubleClickTime) {
	self->MouseDoubleClickTime = static_cast<float>(MouseDoubleClickTime);
}

float ImGuiIO_MouseDoubleClickMaxDist(const ImGuiIO* self) {
	return self->MouseDoubleClickMaxDist;
}

void ImGuiIO_setMouseDoubleClickMaxDist(ImGuiIO* self, float MouseDoubleClickMaxDist) {
	self->MouseDoubleClickMaxDist = static_cast<float>(MouseDoubleClickMaxDist);
}

float ImGuiIO_MouseDragThreshold(const ImGuiIO* self) {
	return self->MouseDragThreshold;
}

void ImGuiIO_setMouseDragThreshold(ImGuiIO* self, float MouseDragThreshold) {
	self->MouseDragThreshold = static_cast<float>(MouseDragThreshold);
}

float ImGuiIO_KeyRepeatDelay(const ImGuiIO* self) {
	return self->KeyRepeatDelay;
}

void ImGuiIO_setKeyRepeatDelay(ImGuiIO* self, float KeyRepeatDelay) {
	self->KeyRepeatDelay = static_cast<float>(KeyRepeatDelay);
}

float ImGuiIO_KeyRepeatRate(const ImGuiIO* self) {
	return self->KeyRepeatRate;
}

void ImGuiIO_setKeyRepeatRate(ImGuiIO* self, float KeyRepeatRate) {
	self->KeyRepeatRate = static_cast<float>(KeyRepeatRate);
}

bool ImGuiIO_ConfigErrorRecovery(const ImGuiIO* self) {
	return self->ConfigErrorRecovery;
}

void ImGuiIO_setConfigErrorRecovery(ImGuiIO* self, bool ConfigErrorRecovery) {
	self->ConfigErrorRecovery = ConfigErrorRecovery;
}

bool ImGuiIO_ConfigErrorRecoveryEnableAssert(const ImGuiIO* self) {
	return self->ConfigErrorRecoveryEnableAssert;
}

void ImGuiIO_setConfigErrorRecoveryEnableAssert(ImGuiIO* self, bool ConfigErrorRecoveryEnableAssert) {
	self->ConfigErrorRecoveryEnableAssert = ConfigErrorRecoveryEnableAssert;
}

bool ImGuiIO_ConfigErrorRecoveryEnableDebugLog(const ImGuiIO* self) {
	return self->ConfigErrorRecoveryEnableDebugLog;
}

void ImGuiIO_setConfigErrorRecoveryEnableDebugLog(ImGuiIO* self, bool ConfigErrorRecoveryEnableDebugLog) {
	self->ConfigErrorRecoveryEnableDebugLog = ConfigErrorRecoveryEnableDebugLog;
}

bool ImGuiIO_ConfigErrorRecoveryEnableTooltip(const ImGuiIO* self) {
	return self->ConfigErrorRecoveryEnableTooltip;
}

void ImGuiIO_setConfigErrorRecoveryEnableTooltip(ImGuiIO* self, bool ConfigErrorRecoveryEnableTooltip) {
	self->ConfigErrorRecoveryEnableTooltip = ConfigErrorRecoveryEnableTooltip;
}

bool ImGuiIO_ConfigDebugIsDebuggerPresent(const ImGuiIO* self) {
	return self->ConfigDebugIsDebuggerPresent;
}

void ImGuiIO_setConfigDebugIsDebuggerPresent(ImGuiIO* self, bool ConfigDebugIsDebuggerPresent) {
	self->ConfigDebugIsDebuggerPresent = ConfigDebugIsDebuggerPresent;
}

bool ImGuiIO_ConfigDebugHighlightIdConflicts(const ImGuiIO* self) {
	return self->ConfigDebugHighlightIdConflicts;
}

void ImGuiIO_setConfigDebugHighlightIdConflicts(ImGuiIO* self, bool ConfigDebugHighlightIdConflicts) {
	self->ConfigDebugHighlightIdConflicts = ConfigDebugHighlightIdConflicts;
}

bool ImGuiIO_ConfigDebugHighlightIdConflictsShowItemPicker(const ImGuiIO* self) {
	return self->ConfigDebugHighlightIdConflictsShowItemPicker;
}

void ImGuiIO_setConfigDebugHighlightIdConflictsShowItemPicker(ImGuiIO* self, bool ConfigDebugHighlightIdConflictsShowItemPicker) {
	self->ConfigDebugHighlightIdConflictsShowItemPicker = ConfigDebugHighlightIdConflictsShowItemPicker;
}

bool ImGuiIO_ConfigDebugBeginReturnValueOnce(const ImGuiIO* self) {
	return self->ConfigDebugBeginReturnValueOnce;
}

void ImGuiIO_setConfigDebugBeginReturnValueOnce(ImGuiIO* self, bool ConfigDebugBeginReturnValueOnce) {
	self->ConfigDebugBeginReturnValueOnce = ConfigDebugBeginReturnValueOnce;
}

bool ImGuiIO_ConfigDebugBeginReturnValueLoop(const ImGuiIO* self) {
	return self->ConfigDebugBeginReturnValueLoop;
}

void ImGuiIO_setConfigDebugBeginReturnValueLoop(ImGuiIO* self, bool ConfigDebugBeginReturnValueLoop) {
	self->ConfigDebugBeginReturnValueLoop = ConfigDebugBeginReturnValueLoop;
}

bool ImGuiIO_ConfigDebugIgnoreFocusLoss(const ImGuiIO* self) {
	return self->ConfigDebugIgnoreFocusLoss;
}

void ImGuiIO_setConfigDebugIgnoreFocusLoss(ImGuiIO* self, bool ConfigDebugIgnoreFocusLoss) {
	self->ConfigDebugIgnoreFocusLoss = ConfigDebugIgnoreFocusLoss;
}

bool ImGuiIO_ConfigDebugIniSettings(const ImGuiIO* self) {
	return self->ConfigDebugIniSettings;
}

void ImGuiIO_setConfigDebugIniSettings(ImGuiIO* self, bool ConfigDebugIniSettings) {
	self->ConfigDebugIniSettings = ConfigDebugIniSettings;
}

const char* ImGuiIO_BackendPlatformName(const ImGuiIO* self) {
	return (const char*) self->BackendPlatformName;
}

void ImGuiIO_setBackendPlatformName(ImGuiIO* self, const char* BackendPlatformName) {
	self->BackendPlatformName = BackendPlatformName;
}

const char* ImGuiIO_BackendRendererName(const ImGuiIO* self) {
	return (const char*) self->BackendRendererName;
}

void ImGuiIO_setBackendRendererName(ImGuiIO* self, const char* BackendRendererName) {
	self->BackendRendererName = BackendRendererName;
}

void ImGuiIO_AddKeyEvent(ImGuiIO* self, int key, bool down) {
	self->AddKeyEvent(static_cast<ImGuiKey>(key), down);
}

void ImGuiIO_AddKeyAnalogEvent(ImGuiIO* self, int key, bool down, float v) {
	self->AddKeyAnalogEvent(static_cast<ImGuiKey>(key), down, static_cast<float>(v));
}

void ImGuiIO_AddMousePosEvent(ImGuiIO* self, float x, float y) {
	self->AddMousePosEvent(static_cast<float>(x), static_cast<float>(y));
}

void ImGuiIO_AddMouseButtonEvent(ImGuiIO* self, int button, bool down) {
	self->AddMouseButtonEvent(static_cast<int>(button), down);
}

void ImGuiIO_AddMouseWheelEvent(ImGuiIO* self, float wheel_x, float wheel_y) {
	self->AddMouseWheelEvent(static_cast<float>(wheel_x), static_cast<float>(wheel_y));
}

void ImGuiIO_AddMouseSourceEvent(ImGuiIO* self, int source) {
	self->AddMouseSourceEvent(static_cast<ImGuiMouseSource>(source));
}

void ImGuiIO_AddFocusEvent(ImGuiIO* self, bool focused) {
	self->AddFocusEvent(focused);
}

void ImGuiIO_AddInputCharacter(ImGuiIO* self, unsigned int c) {
	self->AddInputCharacter(static_cast<unsigned int>(c));
}

void ImGuiIO_AddInputCharacterUTF16(ImGuiIO* self, unsigned short c) {
	self->AddInputCharacterUTF16(static_cast<unsigned short>(c));
}

void ImGuiIO_AddInputCharactersUTF8(ImGuiIO* self, const char* str) {
	self->AddInputCharactersUTF8(str);
}

void ImGuiIO_SetKeyEventNativeData(ImGuiIO* self, int key, int native_keycode, int native_scancode) {
	self->SetKeyEventNativeData(static_cast<ImGuiKey>(key), static_cast<int>(native_keycode), static_cast<int>(native_scancode));
}

void ImGuiIO_SetAppAcceptingEvents(ImGuiIO* self, bool accepting_events) {
	self->SetAppAcceptingEvents(accepting_events);
}

void ImGuiIO_ClearEventsQueue(ImGuiIO* self) {
	self->ClearEventsQueue();
}

void ImGuiIO_ClearInputKeys(ImGuiIO* self) {
	self->ClearInputKeys();
}

void ImGuiIO_ClearInputMouse(ImGuiIO* self) {
	self->ClearInputMouse();
}

bool ImGuiIO_WantCaptureMouse(const ImGuiIO* self) {
	return self->WantCaptureMouse;
}

void ImGuiIO_setWantCaptureMouse(ImGuiIO* self, bool WantCaptureMouse) {
	self->WantCaptureMouse = WantCaptureMouse;
}

bool ImGuiIO_WantCaptureKeyboard(const ImGuiIO* self) {
	return self->WantCaptureKeyboard;
}

void ImGuiIO_setWantCaptureKeyboard(ImGuiIO* self, bool WantCaptureKeyboard) {
	self->WantCaptureKeyboard = WantCaptureKeyboard;
}

bool ImGuiIO_WantTextInput(const ImGuiIO* self) {
	return self->WantTextInput;
}

void ImGuiIO_setWantTextInput(ImGuiIO* self, bool WantTextInput) {
	self->WantTextInput = WantTextInput;
}

bool ImGuiIO_WantSetMousePos(const ImGuiIO* self) {
	return self->WantSetMousePos;
}

void ImGuiIO_setWantSetMousePos(ImGuiIO* self, bool WantSetMousePos) {
	self->WantSetMousePos = WantSetMousePos;
}

bool ImGuiIO_WantSaveIniSettings(const ImGuiIO* self) {
	return self->WantSaveIniSettings;
}

void ImGuiIO_setWantSaveIniSettings(ImGuiIO* self, bool WantSaveIniSettings) {
	self->WantSaveIniSettings = WantSaveIniSettings;
}

bool ImGuiIO_NavActive(const ImGuiIO* self) {
	return self->NavActive;
}

void ImGuiIO_setNavActive(ImGuiIO* self, bool NavActive) {
	self->NavActive = NavActive;
}

bool ImGuiIO_NavVisible(const ImGuiIO* self) {
	return self->NavVisible;
}

void ImGuiIO_setNavVisible(ImGuiIO* self, bool NavVisible) {
	self->NavVisible = NavVisible;
}

float ImGuiIO_Framerate(const ImGuiIO* self) {
	return self->Framerate;
}

void ImGuiIO_setFramerate(ImGuiIO* self, float Framerate) {
	self->Framerate = static_cast<float>(Framerate);
}

int ImGuiIO_MetricsRenderVertices(const ImGuiIO* self) {
	return self->MetricsRenderVertices;
}

void ImGuiIO_setMetricsRenderVertices(ImGuiIO* self, int MetricsRenderVertices) {
	self->MetricsRenderVertices = static_cast<int>(MetricsRenderVertices);
}

int ImGuiIO_MetricsRenderIndices(const ImGuiIO* self) {
	return self->MetricsRenderIndices;
}

void ImGuiIO_setMetricsRenderIndices(ImGuiIO* self, int MetricsRenderIndices) {
	self->MetricsRenderIndices = static_cast<int>(MetricsRenderIndices);
}

int ImGuiIO_MetricsRenderWindows(const ImGuiIO* self) {
	return self->MetricsRenderWindows;
}

void ImGuiIO_setMetricsRenderWindows(ImGuiIO* self, int MetricsRenderWindows) {
	self->MetricsRenderWindows = static_cast<int>(MetricsRenderWindows);
}

int ImGuiIO_MetricsActiveWindows(const ImGuiIO* self) {
	return self->MetricsActiveWindows;
}

void ImGuiIO_setMetricsActiveWindows(ImGuiIO* self, int MetricsActiveWindows) {
	self->MetricsActiveWindows = static_cast<int>(MetricsActiveWindows);
}

ImVec2* ImGuiIO_MouseDelta(const ImGuiIO* self) {
	return new ImVec2(self->MouseDelta);
}

void ImGuiIO_setMouseDelta(ImGuiIO* self, ImVec2* MouseDelta) {
	self->MouseDelta = *MouseDelta;
}

void* ImGuiIO_Ctx(const ImGuiIO* self) {
	return self->Ctx;
}

void ImGuiIO_setCtx(ImGuiIO* self, void* Ctx) {
	self->Ctx = static_cast<ImGuiContext*>(Ctx);
}

ImVec2* ImGuiIO_MousePos(const ImGuiIO* self) {
	return new ImVec2(self->MousePos);
}

void ImGuiIO_setMousePos(ImGuiIO* self, ImVec2* MousePos) {
	self->MousePos = *MousePos;
}

bool* ImGuiIO_MouseDown(const ImGuiIO* self) {
	return (bool*)self->MouseDown;
}

void ImGuiIO_setMouseDown(ImGuiIO* self, const bool* MouseDown) {
	memcpy(self->MouseDown, MouseDown, sizeof(self->MouseDown));
}

float ImGuiIO_MouseWheel(const ImGuiIO* self) {
	return self->MouseWheel;
}

void ImGuiIO_setMouseWheel(ImGuiIO* self, float MouseWheel) {
	self->MouseWheel = static_cast<float>(MouseWheel);
}

float ImGuiIO_MouseWheelH(const ImGuiIO* self) {
	return self->MouseWheelH;
}

void ImGuiIO_setMouseWheelH(ImGuiIO* self, float MouseWheelH) {
	self->MouseWheelH = static_cast<float>(MouseWheelH);
}

int ImGuiIO_MouseSource(const ImGuiIO* self) {
	ImGuiMouseSource MouseSource_ret = self->MouseSource;
	return static_cast<int>(MouseSource_ret);
}

void ImGuiIO_setMouseSource(ImGuiIO* self, int MouseSource) {
	self->MouseSource = static_cast<ImGuiMouseSource>(MouseSource);
}

bool ImGuiIO_KeyCtrl(const ImGuiIO* self) {
	return self->KeyCtrl;
}

void ImGuiIO_setKeyCtrl(ImGuiIO* self, bool KeyCtrl) {
	self->KeyCtrl = KeyCtrl;
}

bool ImGuiIO_KeyShift(const ImGuiIO* self) {
	return self->KeyShift;
}

void ImGuiIO_setKeyShift(ImGuiIO* self, bool KeyShift) {
	self->KeyShift = KeyShift;
}

bool ImGuiIO_KeyAlt(const ImGuiIO* self) {
	return self->KeyAlt;
}

void ImGuiIO_setKeyAlt(ImGuiIO* self, bool KeyAlt) {
	self->KeyAlt = KeyAlt;
}

bool ImGuiIO_KeySuper(const ImGuiIO* self) {
	return self->KeySuper;
}

void ImGuiIO_setKeySuper(ImGuiIO* self, bool KeySuper) {
	self->KeySuper = KeySuper;
}

int ImGuiIO_KeyMods(const ImGuiIO* self) {
	return self->KeyMods;
}

void ImGuiIO_setKeyMods(ImGuiIO* self, int KeyMods) {
	self->KeyMods = static_cast<int>(KeyMods);
}

ImGuiKeyData* ImGuiIO_KeysData(const ImGuiIO* self) {
	return (ImGuiKeyData*)self->KeysData;
}

void ImGuiIO_setKeysData(ImGuiIO* self, const ImGuiKeyData* KeysData) {
	memcpy(self->KeysData, KeysData, sizeof(self->KeysData));
}

bool ImGuiIO_WantCaptureMouseUnlessPopupClose(const ImGuiIO* self) {
	return self->WantCaptureMouseUnlessPopupClose;
}

void ImGuiIO_setWantCaptureMouseUnlessPopupClose(ImGuiIO* self, bool WantCaptureMouseUnlessPopupClose) {
	self->WantCaptureMouseUnlessPopupClose = WantCaptureMouseUnlessPopupClose;
}

ImVec2* ImGuiIO_MousePosPrev(const ImGuiIO* self) {
	return new ImVec2(self->MousePosPrev);
}

void ImGuiIO_setMousePosPrev(ImGuiIO* self, ImVec2* MousePosPrev) {
	self->MousePosPrev = *MousePosPrev;
}

ImVec2* ImGuiIO_MouseClickedPos(const ImGuiIO* self) {
	return (ImVec2*)self->MouseClickedPos;
}

void ImGuiIO_setMouseClickedPos(ImGuiIO* self, const ImVec2* MouseClickedPos) {
	memcpy(self->MouseClickedPos, MouseClickedPos, sizeof(self->MouseClickedPos));
}

double* ImGuiIO_MouseClickedTime(const ImGuiIO* self) {
	return (double*)self->MouseClickedTime;
}

void ImGuiIO_setMouseClickedTime(ImGuiIO* self, const double* MouseClickedTime) {
	memcpy(self->MouseClickedTime, MouseClickedTime, sizeof(self->MouseClickedTime));
}

bool* ImGuiIO_MouseClicked(const ImGuiIO* self) {
	return (bool*)self->MouseClicked;
}

void ImGuiIO_setMouseClicked(ImGuiIO* self, const bool* MouseClicked) {
	memcpy(self->MouseClicked, MouseClicked, sizeof(self->MouseClicked));
}

bool* ImGuiIO_MouseDoubleClicked(const ImGuiIO* self) {
	return (bool*)self->MouseDoubleClicked;
}

void ImGuiIO_setMouseDoubleClicked(ImGuiIO* self, const bool* MouseDoubleClicked) {
	memcpy(self->MouseDoubleClicked, MouseDoubleClicked, sizeof(self->MouseDoubleClicked));
}

ImU16* ImGuiIO_MouseClickedCount(const ImGuiIO* self) {
	return (ImU16*)self->MouseClickedCount;
}

void ImGuiIO_setMouseClickedCount(ImGuiIO* self, const ImU16* MouseClickedCount) {
	memcpy(self->MouseClickedCount, MouseClickedCount, sizeof(self->MouseClickedCount));
}

ImU16* ImGuiIO_MouseClickedLastCount(const ImGuiIO* self) {
	return (ImU16*)self->MouseClickedLastCount;
}

void ImGuiIO_setMouseClickedLastCount(ImGuiIO* self, const ImU16* MouseClickedLastCount) {
	memcpy(self->MouseClickedLastCount, MouseClickedLastCount, sizeof(self->MouseClickedLastCount));
}

bool* ImGuiIO_MouseReleased(const ImGuiIO* self) {
	return (bool*)self->MouseReleased;
}

void ImGuiIO_setMouseReleased(ImGuiIO* self, const bool* MouseReleased) {
	memcpy(self->MouseReleased, MouseReleased, sizeof(self->MouseReleased));
}

double* ImGuiIO_MouseReleasedTime(const ImGuiIO* self) {
	return (double*)self->MouseReleasedTime;
}

void ImGuiIO_setMouseReleasedTime(ImGuiIO* self, const double* MouseReleasedTime) {
	memcpy(self->MouseReleasedTime, MouseReleasedTime, sizeof(self->MouseReleasedTime));
}

bool* ImGuiIO_MouseDownOwned(const ImGuiIO* self) {
	return (bool*)self->MouseDownOwned;
}

void ImGuiIO_setMouseDownOwned(ImGuiIO* self, const bool* MouseDownOwned) {
	memcpy(self->MouseDownOwned, MouseDownOwned, sizeof(self->MouseDownOwned));
}

bool* ImGuiIO_MouseDownOwnedUnlessPopupClose(const ImGuiIO* self) {
	return (bool*)self->MouseDownOwnedUnlessPopupClose;
}

void ImGuiIO_setMouseDownOwnedUnlessPopupClose(ImGuiIO* self, const bool* MouseDownOwnedUnlessPopupClose) {
	memcpy(self->MouseDownOwnedUnlessPopupClose, MouseDownOwnedUnlessPopupClose, sizeof(self->MouseDownOwnedUnlessPopupClose));
}

bool ImGuiIO_MouseWheelRequestAxisSwap(const ImGuiIO* self) {
	return self->MouseWheelRequestAxisSwap;
}

void ImGuiIO_setMouseWheelRequestAxisSwap(ImGuiIO* self, bool MouseWheelRequestAxisSwap) {
	self->MouseWheelRequestAxisSwap = MouseWheelRequestAxisSwap;
}

bool ImGuiIO_MouseCtrlLeftAsRightClick(const ImGuiIO* self) {
	return self->MouseCtrlLeftAsRightClick;
}

void ImGuiIO_setMouseCtrlLeftAsRightClick(ImGuiIO* self, bool MouseCtrlLeftAsRightClick) {
	self->MouseCtrlLeftAsRightClick = MouseCtrlLeftAsRightClick;
}

float* ImGuiIO_MouseDownDuration(const ImGuiIO* self) {
	return (float*)self->MouseDownDuration;
}

void ImGuiIO_setMouseDownDuration(ImGuiIO* self, const float* MouseDownDuration) {
	memcpy(self->MouseDownDuration, MouseDownDuration, sizeof(self->MouseDownDuration));
}

float* ImGuiIO_MouseDownDurationPrev(const ImGuiIO* self) {
	return (float*)self->MouseDownDurationPrev;
}

void ImGuiIO_setMouseDownDurationPrev(ImGuiIO* self, const float* MouseDownDurationPrev) {
	memcpy(self->MouseDownDurationPrev, MouseDownDurationPrev, sizeof(self->MouseDownDurationPrev));
}

float* ImGuiIO_MouseDragMaxDistanceSqr(const ImGuiIO* self) {
	return (float*)self->MouseDragMaxDistanceSqr;
}

void ImGuiIO_setMouseDragMaxDistanceSqr(ImGuiIO* self, const float* MouseDragMaxDistanceSqr) {
	memcpy(self->MouseDragMaxDistanceSqr, MouseDragMaxDistanceSqr, sizeof(self->MouseDragMaxDistanceSqr));
}

float ImGuiIO_PenPressure(const ImGuiIO* self) {
	return self->PenPressure;
}

void ImGuiIO_setPenPressure(ImGuiIO* self, float PenPressure) {
	self->PenPressure = static_cast<float>(PenPressure);
}

bool ImGuiIO_AppFocusLost(const ImGuiIO* self) {
	return self->AppFocusLost;
}

void ImGuiIO_setAppFocusLost(ImGuiIO* self, bool AppFocusLost) {
	self->AppFocusLost = AppFocusLost;
}

bool ImGuiIO_AppAcceptingEvents(const ImGuiIO* self) {
	return self->AppAcceptingEvents;
}

void ImGuiIO_setAppAcceptingEvents(ImGuiIO* self, bool AppAcceptingEvents) {
	self->AppAcceptingEvents = AppAcceptingEvents;
}

unsigned short ImGuiIO_InputQueueSurrogate(const ImGuiIO* self) {
	return self->InputQueueSurrogate;
}

void ImGuiIO_setInputQueueSurrogate(ImGuiIO* self, unsigned short InputQueueSurrogate) {
	self->InputQueueSurrogate = static_cast<unsigned short>(InputQueueSurrogate);
}

void* ImGuiIO_InputQueueCharacters(const ImGuiIO* self) {
	return const_cast<void*>(static_cast<const void*>(&self->InputQueueCharacters));
}

void ImGuiIO_setInputQueueCharacters(ImGuiIO* self, void* InputQueueCharacters) {
	self->InputQueueCharacters = *reinterpret_cast<ImVector<unsigned short>*>(InputQueueCharacters);
}

float ImGuiIO_FontGlobalScale(const ImGuiIO* self) {
	return self->FontGlobalScale;
}

void ImGuiIO_setFontGlobalScale(ImGuiIO* self, float FontGlobalScale) {
	self->FontGlobalScale = static_cast<float>(FontGlobalScale);
}

void ImGuiIO_operatorAssign(ImGuiIO* self, ImGuiIO* param1) {
	self->operator=(*param1);
}

void ImGuiIO_SetKeyEventNativeData2(ImGuiIO* self, int key, int native_keycode, int native_scancode, int native_legacy_index) {
	self->SetKeyEventNativeData(static_cast<ImGuiKey>(key), static_cast<int>(native_keycode), static_cast<int>(native_scancode), static_cast<int>(native_legacy_index));
}

void ImGuiIO_delete(ImGuiIO* self) {
	delete self;
}

ImGuiInputTextCallbackData* ImGuiInputTextCallbackData_new() {
	return new (std::nothrow) ImGuiInputTextCallbackData();
}

void* ImGuiInputTextCallbackData_Ctx(const ImGuiInputTextCallbackData* self) {
	return self->Ctx;
}

void ImGuiInputTextCallbackData_setCtx(ImGuiInputTextCallbackData* self, void* Ctx) {
	self->Ctx = static_cast<ImGuiContext*>(Ctx);
}

int ImGuiInputTextCallbackData_EventFlag(const ImGuiInputTextCallbackData* self) {
	return self->EventFlag;
}

void ImGuiInputTextCallbackData_setEventFlag(ImGuiInputTextCallbackData* self, int EventFlag) {
	self->EventFlag = static_cast<int>(EventFlag);
}

int ImGuiInputTextCallbackData_Flags(const ImGuiInputTextCallbackData* self) {
	return self->Flags;
}

void ImGuiInputTextCallbackData_setFlags(ImGuiInputTextCallbackData* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

unsigned int ImGuiInputTextCallbackData_ID(const ImGuiInputTextCallbackData* self) {
	return self->ID;
}

void ImGuiInputTextCallbackData_setID(ImGuiInputTextCallbackData* self, unsigned int ID) {
	self->ID = static_cast<unsigned int>(ID);
}

int ImGuiInputTextCallbackData_EventKey(const ImGuiInputTextCallbackData* self) {
	ImGuiKey EventKey_ret = self->EventKey;
	return static_cast<int>(EventKey_ret);
}

void ImGuiInputTextCallbackData_setEventKey(ImGuiInputTextCallbackData* self, int EventKey) {
	self->EventKey = static_cast<ImGuiKey>(EventKey);
}

unsigned short ImGuiInputTextCallbackData_EventChar(const ImGuiInputTextCallbackData* self) {
	return self->EventChar;
}

void ImGuiInputTextCallbackData_setEventChar(ImGuiInputTextCallbackData* self, unsigned short EventChar) {
	self->EventChar = static_cast<unsigned short>(EventChar);
}

bool ImGuiInputTextCallbackData_EventActivated(const ImGuiInputTextCallbackData* self) {
	return self->EventActivated;
}

void ImGuiInputTextCallbackData_setEventActivated(ImGuiInputTextCallbackData* self, bool EventActivated) {
	self->EventActivated = EventActivated;
}

bool ImGuiInputTextCallbackData_BufDirty(const ImGuiInputTextCallbackData* self) {
	return self->BufDirty;
}

void ImGuiInputTextCallbackData_setBufDirty(ImGuiInputTextCallbackData* self, bool BufDirty) {
	self->BufDirty = BufDirty;
}

char* ImGuiInputTextCallbackData_Buf(const ImGuiInputTextCallbackData* self) {
	return self->Buf;
}

void ImGuiInputTextCallbackData_setBuf(ImGuiInputTextCallbackData* self, char* Buf) {
	self->Buf = Buf;
}

int ImGuiInputTextCallbackData_BufTextLen(const ImGuiInputTextCallbackData* self) {
	return self->BufTextLen;
}

void ImGuiInputTextCallbackData_setBufTextLen(ImGuiInputTextCallbackData* self, int BufTextLen) {
	self->BufTextLen = static_cast<int>(BufTextLen);
}

int ImGuiInputTextCallbackData_BufSize(const ImGuiInputTextCallbackData* self) {
	return self->BufSize;
}

void ImGuiInputTextCallbackData_setBufSize(ImGuiInputTextCallbackData* self, int BufSize) {
	self->BufSize = static_cast<int>(BufSize);
}

int ImGuiInputTextCallbackData_CursorPos(const ImGuiInputTextCallbackData* self) {
	return self->CursorPos;
}

void ImGuiInputTextCallbackData_setCursorPos(ImGuiInputTextCallbackData* self, int CursorPos) {
	self->CursorPos = static_cast<int>(CursorPos);
}

int ImGuiInputTextCallbackData_SelectionStart(const ImGuiInputTextCallbackData* self) {
	return self->SelectionStart;
}

void ImGuiInputTextCallbackData_setSelectionStart(ImGuiInputTextCallbackData* self, int SelectionStart) {
	self->SelectionStart = static_cast<int>(SelectionStart);
}

int ImGuiInputTextCallbackData_SelectionEnd(const ImGuiInputTextCallbackData* self) {
	return self->SelectionEnd;
}

void ImGuiInputTextCallbackData_setSelectionEnd(ImGuiInputTextCallbackData* self, int SelectionEnd) {
	self->SelectionEnd = static_cast<int>(SelectionEnd);
}

void ImGuiInputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self, int pos, int bytes_count) {
	self->DeleteChars(static_cast<int>(pos), static_cast<int>(bytes_count));
}

void ImGuiInputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self, int pos, const char* text) {
	self->InsertChars(static_cast<int>(pos), text);
}

void ImGuiInputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self) {
	self->SelectAll();
}

void ImGuiInputTextCallbackData_SetSelection(ImGuiInputTextCallbackData* self, int s, int e) {
	self->SetSelection(static_cast<int>(s), static_cast<int>(e));
}

void ImGuiInputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self) {
	self->ClearSelection();
}

bool ImGuiInputTextCallbackData_HasSelection(const ImGuiInputTextCallbackData* self) {
	return self->HasSelection();
}

void ImGuiInputTextCallbackData_InsertChars2(ImGuiInputTextCallbackData* self, int pos, const char* text, const char* text_end) {
	self->InsertChars(static_cast<int>(pos), text, text_end);
}

void ImGuiInputTextCallbackData_delete(ImGuiInputTextCallbackData* self) {
	delete self;
}

ImVec2* ImGuiSizeCallbackData_Pos(const ImGuiSizeCallbackData* self) {
	return new ImVec2(self->Pos);
}

void ImGuiSizeCallbackData_setPos(ImGuiSizeCallbackData* self, ImVec2* Pos) {
	self->Pos = *Pos;
}

ImVec2* ImGuiSizeCallbackData_CurrentSize(const ImGuiSizeCallbackData* self) {
	return new ImVec2(self->CurrentSize);
}

void ImGuiSizeCallbackData_setCurrentSize(ImGuiSizeCallbackData* self, ImVec2* CurrentSize) {
	self->CurrentSize = *CurrentSize;
}

ImVec2* ImGuiSizeCallbackData_DesiredSize(const ImGuiSizeCallbackData* self) {
	return new ImVec2(self->DesiredSize);
}

void ImGuiSizeCallbackData_setDesiredSize(ImGuiSizeCallbackData* self, ImVec2* DesiredSize) {
	self->DesiredSize = *DesiredSize;
}

void ImGuiSizeCallbackData_delete(ImGuiSizeCallbackData* self) {
	delete self;
}

ImGuiPayload* ImGuiPayload_new() {
	return new (std::nothrow) ImGuiPayload();
}

int ImGuiPayload_DataSize(const ImGuiPayload* self) {
	return self->DataSize;
}

void ImGuiPayload_setDataSize(ImGuiPayload* self, int DataSize) {
	self->DataSize = static_cast<int>(DataSize);
}

unsigned int ImGuiPayload_SourceId(const ImGuiPayload* self) {
	return self->SourceId;
}

void ImGuiPayload_setSourceId(ImGuiPayload* self, unsigned int SourceId) {
	self->SourceId = static_cast<unsigned int>(SourceId);
}

unsigned int ImGuiPayload_SourceParentId(const ImGuiPayload* self) {
	return self->SourceParentId;
}

void ImGuiPayload_setSourceParentId(ImGuiPayload* self, unsigned int SourceParentId) {
	self->SourceParentId = static_cast<unsigned int>(SourceParentId);
}

int ImGuiPayload_DataFrameCount(const ImGuiPayload* self) {
	return self->DataFrameCount;
}

void ImGuiPayload_setDataFrameCount(ImGuiPayload* self, int DataFrameCount) {
	self->DataFrameCount = static_cast<int>(DataFrameCount);
}

bool ImGuiPayload_Preview(const ImGuiPayload* self) {
	return self->Preview;
}

void ImGuiPayload_setPreview(ImGuiPayload* self, bool Preview) {
	self->Preview = Preview;
}

bool ImGuiPayload_Delivery(const ImGuiPayload* self) {
	return self->Delivery;
}

void ImGuiPayload_setDelivery(ImGuiPayload* self, bool Delivery) {
	self->Delivery = Delivery;
}

void ImGuiPayload_Clear(ImGuiPayload* self) {
	self->Clear();
}

bool ImGuiPayload_IsDataType(const ImGuiPayload* self, const char* type) {
	return self->IsDataType(type);
}

bool ImGuiPayload_IsPreview(const ImGuiPayload* self) {
	return self->IsPreview();
}

bool ImGuiPayload_IsDelivery(const ImGuiPayload* self) {
	return self->IsDelivery();
}

void ImGuiPayload_delete(ImGuiPayload* self) {
	delete self;
}

ImGuiOnceUponAFrame* ImGuiOnceUponAFrame_new() {
	return new (std::nothrow) ImGuiOnceUponAFrame();
}

int ImGuiOnceUponAFrame_RefFrame(const ImGuiOnceUponAFrame* self) {
	return self->RefFrame;
}

void ImGuiOnceUponAFrame_setRefFrame(ImGuiOnceUponAFrame* self, int RefFrame) {
	self->RefFrame = static_cast<int>(RefFrame);
}

bool ImGuiOnceUponAFrame_ToBool(const ImGuiOnceUponAFrame* self) {
	return self->operator bool();
}

void ImGuiOnceUponAFrame_delete(ImGuiOnceUponAFrame* self) {
	delete self;
}

ImGuiTextFilter* ImGuiTextFilter_new() {
	return new (std::nothrow) ImGuiTextFilter();
}

ImGuiTextFilter* ImGuiTextFilter_new2(ImGuiTextFilter* param1) {
	return new (std::nothrow) ImGuiTextFilter(*param1);
}

ImGuiTextFilter* ImGuiTextFilter_new3(const char* default_filter) {
	return new (std::nothrow) ImGuiTextFilter(default_filter);
}

bool ImGuiTextFilter_Draw(ImGuiTextFilter* self) {
	return self->Draw();
}

bool ImGuiTextFilter_PassFilter(const ImGuiTextFilter* self, const char* text) {
	return self->PassFilter(text);
}

void ImGuiTextFilter_Build(ImGuiTextFilter* self) {
	self->Build();
}

void ImGuiTextFilter_Clear(ImGuiTextFilter* self) {
	self->Clear();
}

bool ImGuiTextFilter_IsActive(const ImGuiTextFilter* self) {
	return self->IsActive();
}

int ImGuiTextFilter_CountGrep(const ImGuiTextFilter* self) {
	return self->CountGrep;
}

void ImGuiTextFilter_setCountGrep(ImGuiTextFilter* self, int CountGrep) {
	self->CountGrep = static_cast<int>(CountGrep);
}

void ImGuiTextFilter_operatorAssign(ImGuiTextFilter* self, ImGuiTextFilter* param1) {
	self->operator=(*param1);
}

bool ImGuiTextFilter_DrawWithLabel(ImGuiTextFilter* self, const char* label) {
	return self->Draw(label);
}

bool ImGuiTextFilter_Draw2(ImGuiTextFilter* self, const char* label, float width) {
	return self->Draw(label, static_cast<float>(width));
}

bool ImGuiTextFilter_PassFilter2(const ImGuiTextFilter* self, const char* text, const char* text_end) {
	return self->PassFilter(text, text_end);
}

void ImGuiTextFilter_delete(ImGuiTextFilter* self) {
	delete self;
}

ImGuiTextBuffer* ImGuiTextBuffer_new() {
	return new (std::nothrow) ImGuiTextBuffer();
}

ImGuiTextBuffer* ImGuiTextBuffer_new2(ImGuiTextBuffer* param1) {
	return new (std::nothrow) ImGuiTextBuffer(*param1);
}

void* ImGuiTextBuffer_Buf(const ImGuiTextBuffer* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Buf));
}

void ImGuiTextBuffer_setBuf(ImGuiTextBuffer* self, void* Buf) {
	self->Buf = *reinterpret_cast<ImVector<char>*>(Buf);
}

char ImGuiTextBuffer_operatorSubscript(const ImGuiTextBuffer* self, int i) {
	return self->operator[](static_cast<int>(i));
}

const char* ImGuiTextBuffer_begin(const ImGuiTextBuffer* self) {
	return (const char*) self->begin();
}

const char* ImGuiTextBuffer_end(const ImGuiTextBuffer* self) {
	return (const char*) self->end();
}

int ImGuiTextBuffer_size(const ImGuiTextBuffer* self) {
	return self->size();
}

bool ImGuiTextBuffer_empty(const ImGuiTextBuffer* self) {
	return self->empty();
}

void ImGuiTextBuffer_clear(ImGuiTextBuffer* self) {
	self->clear();
}

void ImGuiTextBuffer_resize(ImGuiTextBuffer* self, int size) {
	self->resize(static_cast<int>(size));
}

void ImGuiTextBuffer_reserve(ImGuiTextBuffer* self, int capacity) {
	self->reserve(static_cast<int>(capacity));
}

const char* ImGuiTextBuffer_cStr(const ImGuiTextBuffer* self) {
	return (const char*) self->c_str();
}

void ImGuiTextBuffer_append(ImGuiTextBuffer* self, const char* str) {
	self->append(str);
}

void ImGuiTextBuffer_operatorAssign(ImGuiTextBuffer* self, ImGuiTextBuffer* param1) {
	self->operator=(*param1);
}

void ImGuiTextBuffer_append2(ImGuiTextBuffer* self, const char* str, const char* str_end) {
	self->append(str, str_end);
}

void ImGuiTextBuffer_delete(ImGuiTextBuffer* self) {
	delete self;
}

ImGuiStoragePair* ImGuiStoragePair_new(unsigned int _key, int _val) {
	return new (std::nothrow) ImGuiStoragePair(static_cast<unsigned int>(_key), static_cast<int>(_val));
}

ImGuiStoragePair* ImGuiStoragePair_new2(unsigned int _key, float _val) {
	return new (std::nothrow) ImGuiStoragePair(static_cast<unsigned int>(_key), static_cast<float>(_val));
}

ImGuiStoragePair* ImGuiStoragePair_new3(unsigned int _key, void* _val) {
	return new (std::nothrow) ImGuiStoragePair(static_cast<unsigned int>(_key), static_cast<void*>(_val));
}

unsigned int ImGuiStoragePair_key(const ImGuiStoragePair* self) {
	return self->key;
}

void ImGuiStoragePair_setKey(ImGuiStoragePair* self, unsigned int key) {
	self->key = static_cast<unsigned int>(key);
}

void ImGuiStoragePair_delete(ImGuiStoragePair* self) {
	delete self;
}

ImGuiStorage* ImGuiStorage_new(ImGuiStorage* param1) {
	return new (std::nothrow) ImGuiStorage(*param1);
}

void* ImGuiStorage_Data(const ImGuiStorage* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Data));
}

void ImGuiStorage_setData(ImGuiStorage* self, void* Data) {
	self->Data = *reinterpret_cast<ImVector<ImGuiStoragePair>*>(Data);
}

void ImGuiStorage_Clear(ImGuiStorage* self) {
	self->Clear();
}

int ImGuiStorage_GetInt(const ImGuiStorage* self, unsigned int key) {
	return self->GetInt(static_cast<unsigned int>(key));
}

void ImGuiStorage_SetInt(ImGuiStorage* self, unsigned int key, int val) {
	self->SetInt(static_cast<unsigned int>(key), static_cast<int>(val));
}

bool ImGuiStorage_GetBool(const ImGuiStorage* self, unsigned int key) {
	return self->GetBool(static_cast<unsigned int>(key));
}

void ImGuiStorage_SetBool(ImGuiStorage* self, unsigned int key, bool val) {
	self->SetBool(static_cast<unsigned int>(key), val);
}

float ImGuiStorage_GetFloat(const ImGuiStorage* self, unsigned int key) {
	return self->GetFloat(static_cast<unsigned int>(key));
}

void ImGuiStorage_SetFloat(ImGuiStorage* self, unsigned int key, float val) {
	self->SetFloat(static_cast<unsigned int>(key), static_cast<float>(val));
}

void* ImGuiStorage_GetVoidPtr(const ImGuiStorage* self, unsigned int key) {
	return self->GetVoidPtr(static_cast<unsigned int>(key));
}

int* ImGuiStorage_GetIntRef(ImGuiStorage* self, unsigned int key) {
	return self->GetIntRef(static_cast<unsigned int>(key));
}

bool* ImGuiStorage_GetBoolRef(ImGuiStorage* self, unsigned int key) {
	return self->GetBoolRef(static_cast<unsigned int>(key));
}

float* ImGuiStorage_GetFloatRef(ImGuiStorage* self, unsigned int key) {
	return self->GetFloatRef(static_cast<unsigned int>(key));
}

void ImGuiStorage_BuildSortByKey(ImGuiStorage* self) {
	self->BuildSortByKey();
}

void ImGuiStorage_SetAllInt(ImGuiStorage* self, int val) {
	self->SetAllInt(static_cast<int>(val));
}

void ImGuiStorage_operatorAssign(ImGuiStorage* self, ImGuiStorage* param1) {
	self->operator=(*param1);
}

int ImGuiStorage_GetInt2(const ImGuiStorage* self, unsigned int key, int default_val) {
	return self->GetInt(static_cast<unsigned int>(key), static_cast<int>(default_val));
}

bool ImGuiStorage_GetBool2(const ImGuiStorage* self, unsigned int key, bool default_val) {
	return self->GetBool(static_cast<unsigned int>(key), default_val);
}

float ImGuiStorage_GetFloat2(const ImGuiStorage* self, unsigned int key, float default_val) {
	return self->GetFloat(static_cast<unsigned int>(key), static_cast<float>(default_val));
}

int* ImGuiStorage_GetIntRef2(ImGuiStorage* self, unsigned int key, int default_val) {
	return self->GetIntRef(static_cast<unsigned int>(key), static_cast<int>(default_val));
}

bool* ImGuiStorage_GetBoolRef2(ImGuiStorage* self, unsigned int key, bool default_val) {
	return self->GetBoolRef(static_cast<unsigned int>(key), default_val);
}

float* ImGuiStorage_GetFloatRef2(ImGuiStorage* self, unsigned int key, float default_val) {
	return self->GetFloatRef(static_cast<unsigned int>(key), static_cast<float>(default_val));
}

void ImGuiStorage_delete(ImGuiStorage* self) {
	delete self;
}

ImGuiListClipper* ImGuiListClipper_new() {
	return new (std::nothrow) ImGuiListClipper();
}

int ImGuiListClipper_DisplayStart(const ImGuiListClipper* self) {
	return self->DisplayStart;
}

void ImGuiListClipper_setDisplayStart(ImGuiListClipper* self, int DisplayStart) {
	self->DisplayStart = static_cast<int>(DisplayStart);
}

int ImGuiListClipper_DisplayEnd(const ImGuiListClipper* self) {
	return self->DisplayEnd;
}

void ImGuiListClipper_setDisplayEnd(ImGuiListClipper* self, int DisplayEnd) {
	self->DisplayEnd = static_cast<int>(DisplayEnd);
}

int ImGuiListClipper_UserIndex(const ImGuiListClipper* self) {
	return self->UserIndex;
}

void ImGuiListClipper_setUserIndex(ImGuiListClipper* self, int UserIndex) {
	self->UserIndex = static_cast<int>(UserIndex);
}

int ImGuiListClipper_ItemsCount(const ImGuiListClipper* self) {
	return self->ItemsCount;
}

void ImGuiListClipper_setItemsCount(ImGuiListClipper* self, int ItemsCount) {
	self->ItemsCount = static_cast<int>(ItemsCount);
}

float ImGuiListClipper_ItemsHeight(const ImGuiListClipper* self) {
	return self->ItemsHeight;
}

void ImGuiListClipper_setItemsHeight(ImGuiListClipper* self, float ItemsHeight) {
	self->ItemsHeight = static_cast<float>(ItemsHeight);
}

int ImGuiListClipper_Flags(const ImGuiListClipper* self) {
	return self->Flags;
}

void ImGuiListClipper_setFlags(ImGuiListClipper* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

double ImGuiListClipper_StartPosY(const ImGuiListClipper* self) {
	return self->StartPosY;
}

void ImGuiListClipper_setStartPosY(ImGuiListClipper* self, double StartPosY) {
	self->StartPosY = static_cast<double>(StartPosY);
}

double ImGuiListClipper_StartSeekOffsetY(const ImGuiListClipper* self) {
	return self->StartSeekOffsetY;
}

void ImGuiListClipper_setStartSeekOffsetY(ImGuiListClipper* self, double StartSeekOffsetY) {
	self->StartSeekOffsetY = static_cast<double>(StartSeekOffsetY);
}

void* ImGuiListClipper_Ctx(const ImGuiListClipper* self) {
	return self->Ctx;
}

void ImGuiListClipper_setCtx(ImGuiListClipper* self, void* Ctx) {
	self->Ctx = static_cast<ImGuiContext*>(Ctx);
}

void ImGuiListClipper_Begin(ImGuiListClipper* self, int items_count) {
	self->Begin(static_cast<int>(items_count));
}

void ImGuiListClipper_End(ImGuiListClipper* self) {
	self->End();
}

bool ImGuiListClipper_Step(ImGuiListClipper* self) {
	return self->Step();
}

void ImGuiListClipper_IncludeItemByIndex(ImGuiListClipper* self, int item_index) {
	self->IncludeItemByIndex(static_cast<int>(item_index));
}

void ImGuiListClipper_IncludeItemsByIndex(ImGuiListClipper* self, int item_begin, int item_end) {
	self->IncludeItemsByIndex(static_cast<int>(item_begin), static_cast<int>(item_end));
}

void ImGuiListClipper_SeekCursorForItem(ImGuiListClipper* self, int item_index) {
	self->SeekCursorForItem(static_cast<int>(item_index));
}

void ImGuiListClipper_Begin2(ImGuiListClipper* self, int items_count, float items_height) {
	self->Begin(static_cast<int>(items_count), static_cast<float>(items_height));
}

void ImGuiListClipper_delete(ImGuiListClipper* self) {
	delete self;
}

ImColor* ImColor_new() {
	return new (std::nothrow) ImColor();
}

ImColor* ImColor_new2(float r, float g, float b) {
	return new (std::nothrow) ImColor(static_cast<float>(r), static_cast<float>(g), static_cast<float>(b));
}

ImColor* ImColor_new3(ImVec4* col) {
	return new (std::nothrow) ImColor(*col);
}

ImColor* ImColor_new4(int r, int g, int b) {
	return new (std::nothrow) ImColor(static_cast<int>(r), static_cast<int>(g), static_cast<int>(b));
}

ImColor* ImColor_new5(unsigned int rgba) {
	return new (std::nothrow) ImColor(static_cast<unsigned int>(rgba));
}

ImColor* ImColor_new6(ImColor* param1) {
	return new (std::nothrow) ImColor(*param1);
}

ImColor* ImColor_new7(float r, float g, float b, float a) {
	return new (std::nothrow) ImColor(static_cast<float>(r), static_cast<float>(g), static_cast<float>(b), static_cast<float>(a));
}

ImColor* ImColor_new8(int r, int g, int b, int a) {
	return new (std::nothrow) ImColor(static_cast<int>(r), static_cast<int>(g), static_cast<int>(b), static_cast<int>(a));
}

ImVec4* ImColor_Value(const ImColor* self) {
	return new ImVec4(self->Value);
}

void ImColor_setValue(ImColor* self, ImVec4* Value) {
	self->Value = *Value;
}

unsigned int ImColor_ToUnsignedInt(const ImColor* self) {
	unsigned int _ret = self->operator unsigned int();
	return static_cast<unsigned int>(_ret);
}

ImVec4* ImColor_ToImVec4(const ImColor* self) {
	return new ImVec4(self->operator ImVec4());
}

void ImColor_SetHSV(ImColor* self, float h, float s, float v) {
	self->SetHSV(static_cast<float>(h), static_cast<float>(s), static_cast<float>(v));
}

ImColor* ImColor_HSV(float h, float s, float v) {
	return new ImColor(ImColor::HSV(static_cast<float>(h), static_cast<float>(s), static_cast<float>(v)));
}

void ImColor_SetHSV2(ImColor* self, float h, float s, float v, float a) {
	self->SetHSV(static_cast<float>(h), static_cast<float>(s), static_cast<float>(v), static_cast<float>(a));
}

ImColor* ImColor_HSV2(float h, float s, float v, float a) {
	return new ImColor(ImColor::HSV(static_cast<float>(h), static_cast<float>(s), static_cast<float>(v), static_cast<float>(a)));
}

void ImColor_delete(ImColor* self) {
	delete self;
}

ImGuiMultiSelectIO* ImGuiMultiSelectIO_new(ImGuiMultiSelectIO* param1) {
	return new (std::nothrow) ImGuiMultiSelectIO(*param1);
}

void* ImGuiMultiSelectIO_Requests(const ImGuiMultiSelectIO* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Requests));
}

void ImGuiMultiSelectIO_setRequests(ImGuiMultiSelectIO* self, void* Requests) {
	self->Requests = *reinterpret_cast<ImVector<ImGuiSelectionRequest>*>(Requests);
}

long long ImGuiMultiSelectIO_RangeSrcItem(const ImGuiMultiSelectIO* self) {
	return self->RangeSrcItem;
}

void ImGuiMultiSelectIO_setRangeSrcItem(ImGuiMultiSelectIO* self, long long RangeSrcItem) {
	self->RangeSrcItem = static_cast<long long>(RangeSrcItem);
}

long long ImGuiMultiSelectIO_NavIdItem(const ImGuiMultiSelectIO* self) {
	return self->NavIdItem;
}

void ImGuiMultiSelectIO_setNavIdItem(ImGuiMultiSelectIO* self, long long NavIdItem) {
	self->NavIdItem = static_cast<long long>(NavIdItem);
}

bool ImGuiMultiSelectIO_NavIdSelected(const ImGuiMultiSelectIO* self) {
	return self->NavIdSelected;
}

void ImGuiMultiSelectIO_setNavIdSelected(ImGuiMultiSelectIO* self, bool NavIdSelected) {
	self->NavIdSelected = NavIdSelected;
}

bool ImGuiMultiSelectIO_RangeSrcReset(const ImGuiMultiSelectIO* self) {
	return self->RangeSrcReset;
}

void ImGuiMultiSelectIO_setRangeSrcReset(ImGuiMultiSelectIO* self, bool RangeSrcReset) {
	self->RangeSrcReset = RangeSrcReset;
}

int ImGuiMultiSelectIO_ItemsCount(const ImGuiMultiSelectIO* self) {
	return self->ItemsCount;
}

void ImGuiMultiSelectIO_setItemsCount(ImGuiMultiSelectIO* self, int ItemsCount) {
	self->ItemsCount = static_cast<int>(ItemsCount);
}

void ImGuiMultiSelectIO_operatorAssign(ImGuiMultiSelectIO* self, ImGuiMultiSelectIO* param1) {
	self->operator=(*param1);
}

void ImGuiMultiSelectIO_delete(ImGuiMultiSelectIO* self) {
	delete self;
}

int ImGuiSelectionRequest_Type(const ImGuiSelectionRequest* self) {
	ImGuiSelectionRequestType Type_ret = self->Type;
	return static_cast<int>(Type_ret);
}

void ImGuiSelectionRequest_setType(ImGuiSelectionRequest* self, int Type) {
	self->Type = static_cast<ImGuiSelectionRequestType>(Type);
}

bool ImGuiSelectionRequest_Selected(const ImGuiSelectionRequest* self) {
	return self->Selected;
}

void ImGuiSelectionRequest_setSelected(ImGuiSelectionRequest* self, bool Selected) {
	self->Selected = Selected;
}

signed char ImGuiSelectionRequest_RangeDirection(const ImGuiSelectionRequest* self) {
	return self->RangeDirection;
}

void ImGuiSelectionRequest_setRangeDirection(ImGuiSelectionRequest* self, signed char RangeDirection) {
	self->RangeDirection = static_cast<signed char>(RangeDirection);
}

long long ImGuiSelectionRequest_RangeFirstItem(const ImGuiSelectionRequest* self) {
	return self->RangeFirstItem;
}

void ImGuiSelectionRequest_setRangeFirstItem(ImGuiSelectionRequest* self, long long RangeFirstItem) {
	self->RangeFirstItem = static_cast<long long>(RangeFirstItem);
}

long long ImGuiSelectionRequest_RangeLastItem(const ImGuiSelectionRequest* self) {
	return self->RangeLastItem;
}

void ImGuiSelectionRequest_setRangeLastItem(ImGuiSelectionRequest* self, long long RangeLastItem) {
	self->RangeLastItem = static_cast<long long>(RangeLastItem);
}

void ImGuiSelectionRequest_delete(ImGuiSelectionRequest* self) {
	delete self;
}

ImGuiSelectionBasicStorage* ImGuiSelectionBasicStorage_new() {
	return new (std::nothrow) ImGuiSelectionBasicStorage();
}

int ImGuiSelectionBasicStorage_Size(const ImGuiSelectionBasicStorage* self) {
	return self->Size;
}

void ImGuiSelectionBasicStorage_setSize(ImGuiSelectionBasicStorage* self, int Size) {
	self->Size = static_cast<int>(Size);
}

bool ImGuiSelectionBasicStorage_PreserveOrder(const ImGuiSelectionBasicStorage* self) {
	return self->PreserveOrder;
}

void ImGuiSelectionBasicStorage_setPreserveOrder(ImGuiSelectionBasicStorage* self, bool PreserveOrder) {
	self->PreserveOrder = PreserveOrder;
}

int ImGuiSelectionBasicStorage__SelectionOrder(const ImGuiSelectionBasicStorage* self) {
	return self->_SelectionOrder;
}

void ImGuiSelectionBasicStorage_set_SelectionOrder(ImGuiSelectionBasicStorage* self, int _SelectionOrder) {
	self->_SelectionOrder = static_cast<int>(_SelectionOrder);
}

ImGuiStorage* ImGuiSelectionBasicStorage__Storage(const ImGuiSelectionBasicStorage* self) {
	return new ImGuiStorage(self->_Storage);
}

void ImGuiSelectionBasicStorage_set_Storage(ImGuiSelectionBasicStorage* self, ImGuiStorage* _Storage) {
	self->_Storage = *_Storage;
}

void ImGuiSelectionBasicStorage_ApplyRequests(ImGuiSelectionBasicStorage* self, ImGuiMultiSelectIO* ms_io) {
	self->ApplyRequests(ms_io);
}

bool ImGuiSelectionBasicStorage_Contains(const ImGuiSelectionBasicStorage* self, unsigned int id) {
	return self->Contains(static_cast<unsigned int>(id));
}

void ImGuiSelectionBasicStorage_Clear(ImGuiSelectionBasicStorage* self) {
	self->Clear();
}

void ImGuiSelectionBasicStorage_Swap(ImGuiSelectionBasicStorage* self, ImGuiSelectionBasicStorage* r) {
	self->Swap(*r);
}

void ImGuiSelectionBasicStorage_SetItemSelected(ImGuiSelectionBasicStorage* self, unsigned int id, bool selected) {
	self->SetItemSelected(static_cast<unsigned int>(id), selected);
}

unsigned int ImGuiSelectionBasicStorage_GetStorageIdFromIndex(ImGuiSelectionBasicStorage* self, int idx) {
	unsigned int _ret = self->GetStorageIdFromIndex(static_cast<int>(idx));
	return static_cast<unsigned int>(_ret);
}

void ImGuiSelectionBasicStorage_delete(ImGuiSelectionBasicStorage* self) {
	delete self;
}

ImGuiSelectionExternalStorage* ImGuiSelectionExternalStorage_new() {
	return new (std::nothrow) ImGuiSelectionExternalStorage();
}

void ImGuiSelectionExternalStorage_ApplyRequests(ImGuiSelectionExternalStorage* self, ImGuiMultiSelectIO* ms_io) {
	self->ApplyRequests(ms_io);
}

void ImGuiSelectionExternalStorage_delete(ImGuiSelectionExternalStorage* self) {
	delete self;
}

ImDrawCmd* ImDrawCmd_new() {
	return new (std::nothrow) ImDrawCmd();
}

ImVec4* ImDrawCmd_ClipRect(const ImDrawCmd* self) {
	return new ImVec4(self->ClipRect);
}

void ImDrawCmd_setClipRect(ImDrawCmd* self, ImVec4* ClipRect) {
	self->ClipRect = *ClipRect;
}

ImTextureRef* ImDrawCmd_TexRef(const ImDrawCmd* self) {
	return new ImTextureRef(self->TexRef);
}

void ImDrawCmd_setTexRef(ImDrawCmd* self, ImTextureRef* TexRef) {
	self->TexRef = *TexRef;
}

unsigned int ImDrawCmd_VtxOffset(const ImDrawCmd* self) {
	return self->VtxOffset;
}

void ImDrawCmd_setVtxOffset(ImDrawCmd* self, unsigned int VtxOffset) {
	self->VtxOffset = static_cast<unsigned int>(VtxOffset);
}

unsigned int ImDrawCmd_IdxOffset(const ImDrawCmd* self) {
	return self->IdxOffset;
}

void ImDrawCmd_setIdxOffset(ImDrawCmd* self, unsigned int IdxOffset) {
	self->IdxOffset = static_cast<unsigned int>(IdxOffset);
}

unsigned int ImDrawCmd_ElemCount(const ImDrawCmd* self) {
	return self->ElemCount;
}

void ImDrawCmd_setElemCount(ImDrawCmd* self, unsigned int ElemCount) {
	self->ElemCount = static_cast<unsigned int>(ElemCount);
}

int ImDrawCmd_UserCallbackDataSize(const ImDrawCmd* self) {
	return self->UserCallbackDataSize;
}

void ImDrawCmd_setUserCallbackDataSize(ImDrawCmd* self, int UserCallbackDataSize) {
	self->UserCallbackDataSize = static_cast<int>(UserCallbackDataSize);
}

int ImDrawCmd_UserCallbackDataOffset(const ImDrawCmd* self) {
	return self->UserCallbackDataOffset;
}

void ImDrawCmd_setUserCallbackDataOffset(ImDrawCmd* self, int UserCallbackDataOffset) {
	self->UserCallbackDataOffset = static_cast<int>(UserCallbackDataOffset);
}

unsigned long long ImDrawCmd_GetTexID(const ImDrawCmd* self) {
	unsigned long long _ret = self->GetTexID();
	return static_cast<unsigned long long>(_ret);
}

void ImDrawCmd_delete(ImDrawCmd* self) {
	delete self;
}

ImVec2* ImDrawVert_pos(const ImDrawVert* self) {
	return new ImVec2(self->pos);
}

void ImDrawVert_setPos(ImDrawVert* self, ImVec2* pos) {
	self->pos = *pos;
}

ImVec2* ImDrawVert_uv(const ImDrawVert* self) {
	return new ImVec2(self->uv);
}

void ImDrawVert_setUv(ImDrawVert* self, ImVec2* uv) {
	self->uv = *uv;
}

unsigned int ImDrawVert_col(const ImDrawVert* self) {
	return self->col;
}

void ImDrawVert_setCol(ImDrawVert* self, unsigned int col) {
	self->col = static_cast<unsigned int>(col);
}

void ImDrawVert_delete(ImDrawVert* self) {
	delete self;
}

ImDrawCmdHeader* ImDrawCmdHeader_new(ImDrawCmdHeader* param1) {
	return new (std::nothrow) ImDrawCmdHeader(*param1);
}

ImVec4* ImDrawCmdHeader_ClipRect(const ImDrawCmdHeader* self) {
	return new ImVec4(self->ClipRect);
}

void ImDrawCmdHeader_setClipRect(ImDrawCmdHeader* self, ImVec4* ClipRect) {
	self->ClipRect = *ClipRect;
}

ImTextureRef* ImDrawCmdHeader_TexRef(const ImDrawCmdHeader* self) {
	return new ImTextureRef(self->TexRef);
}

void ImDrawCmdHeader_setTexRef(ImDrawCmdHeader* self, ImTextureRef* TexRef) {
	self->TexRef = *TexRef;
}

unsigned int ImDrawCmdHeader_VtxOffset(const ImDrawCmdHeader* self) {
	return self->VtxOffset;
}

void ImDrawCmdHeader_setVtxOffset(ImDrawCmdHeader* self, unsigned int VtxOffset) {
	self->VtxOffset = static_cast<unsigned int>(VtxOffset);
}

void ImDrawCmdHeader_operatorAssign(ImDrawCmdHeader* self, ImDrawCmdHeader* param1) {
	self->operator=(*param1);
}

void ImDrawCmdHeader_delete(ImDrawCmdHeader* self) {
	delete self;
}

ImDrawChannel* ImDrawChannel_new(ImDrawChannel* param1) {
	return new (std::nothrow) ImDrawChannel(*param1);
}

void* ImDrawChannel__CmdBuffer(const ImDrawChannel* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_CmdBuffer));
}

void ImDrawChannel_set_CmdBuffer(ImDrawChannel* self, void* _CmdBuffer) {
	self->_CmdBuffer = *reinterpret_cast<ImVector<ImDrawCmd>*>(_CmdBuffer);
}

void* ImDrawChannel__IdxBuffer(const ImDrawChannel* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_IdxBuffer));
}

void ImDrawChannel_set_IdxBuffer(ImDrawChannel* self, void* _IdxBuffer) {
	self->_IdxBuffer = *reinterpret_cast<ImVector<unsigned short>*>(_IdxBuffer);
}

void ImDrawChannel_operatorAssign(ImDrawChannel* self, ImDrawChannel* param1) {
	self->operator=(*param1);
}

void ImDrawChannel_delete(ImDrawChannel* self) {
	delete self;
}

ImDrawListSplitter* ImDrawListSplitter_new() {
	return new (std::nothrow) ImDrawListSplitter();
}

ImDrawListSplitter* ImDrawListSplitter_new2(ImDrawListSplitter* param1) {
	return new (std::nothrow) ImDrawListSplitter(*param1);
}

int ImDrawListSplitter__Current(const ImDrawListSplitter* self) {
	return self->_Current;
}

void ImDrawListSplitter_set_Current(ImDrawListSplitter* self, int _Current) {
	self->_Current = static_cast<int>(_Current);
}

int ImDrawListSplitter__Count(const ImDrawListSplitter* self) {
	return self->_Count;
}

void ImDrawListSplitter_set_Count(ImDrawListSplitter* self, int _Count) {
	self->_Count = static_cast<int>(_Count);
}

void* ImDrawListSplitter__Channels(const ImDrawListSplitter* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_Channels));
}

void ImDrawListSplitter_set_Channels(ImDrawListSplitter* self, void* _Channels) {
	self->_Channels = *reinterpret_cast<ImVector<ImDrawChannel>*>(_Channels);
}

void ImDrawListSplitter_Clear(ImDrawListSplitter* self) {
	self->Clear();
}

void ImDrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self) {
	self->ClearFreeMemory();
}

void ImDrawListSplitter_Split(ImDrawListSplitter* self, ImDrawList* draw_list, int count) {
	self->Split(draw_list, static_cast<int>(count));
}

void ImDrawListSplitter_Merge(ImDrawListSplitter* self, ImDrawList* draw_list) {
	self->Merge(draw_list);
}

void ImDrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self, ImDrawList* draw_list, int channel_idx) {
	self->SetCurrentChannel(draw_list, static_cast<int>(channel_idx));
}

void ImDrawListSplitter_operatorAssign(ImDrawListSplitter* self, ImDrawListSplitter* param1) {
	self->operator=(*param1);
}

void ImDrawListSplitter_delete(ImDrawListSplitter* self) {
	delete self;
}

ImDrawList* ImDrawList_new(void* shared_data) {
	return new (std::nothrow) ImDrawList(static_cast<ImDrawListSharedData*>(shared_data));
}

ImDrawList* ImDrawList_new2(ImDrawList* param1) {
	return new (std::nothrow) ImDrawList(*param1);
}

void* ImDrawList_CmdBuffer(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->CmdBuffer));
}

void ImDrawList_setCmdBuffer(ImDrawList* self, void* CmdBuffer) {
	self->CmdBuffer = *reinterpret_cast<ImVector<ImDrawCmd>*>(CmdBuffer);
}

void* ImDrawList_IdxBuffer(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->IdxBuffer));
}

void ImDrawList_setIdxBuffer(ImDrawList* self, void* IdxBuffer) {
	self->IdxBuffer = *reinterpret_cast<ImVector<unsigned short>*>(IdxBuffer);
}

void* ImDrawList_VtxBuffer(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->VtxBuffer));
}

void ImDrawList_setVtxBuffer(ImDrawList* self, void* VtxBuffer) {
	self->VtxBuffer = *reinterpret_cast<ImVector<ImDrawVert>*>(VtxBuffer);
}

int ImDrawList_Flags(const ImDrawList* self) {
	return self->Flags;
}

void ImDrawList_setFlags(ImDrawList* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

unsigned int ImDrawList__VtxCurrentIdx(const ImDrawList* self) {
	return self->_VtxCurrentIdx;
}

void ImDrawList_set_VtxCurrentIdx(ImDrawList* self, unsigned int _VtxCurrentIdx) {
	self->_VtxCurrentIdx = static_cast<unsigned int>(_VtxCurrentIdx);
}

void* ImDrawList__Data(const ImDrawList* self) {
	return self->_Data;
}

void ImDrawList_set_Data(ImDrawList* self, void* _Data) {
	self->_Data = static_cast<ImDrawListSharedData*>(_Data);
}

ImDrawVert* ImDrawList__VtxWritePtr(const ImDrawList* self) {
	return self->_VtxWritePtr;
}

void ImDrawList_set_VtxWritePtr(ImDrawList* self, ImDrawVert* _VtxWritePtr) {
	self->_VtxWritePtr = _VtxWritePtr;
}

unsigned short* ImDrawList__IdxWritePtr(const ImDrawList* self) {
	unsigned short* _IdxWritePtr_ret = self->_IdxWritePtr;
	return static_cast<unsigned short*>(_IdxWritePtr_ret);
}

void ImDrawList_set_IdxWritePtr(ImDrawList* self, unsigned short* _IdxWritePtr) {
	self->_IdxWritePtr = _IdxWritePtr;
}

void* ImDrawList__Path(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_Path));
}

void ImDrawList_set_Path(ImDrawList* self, void* _Path) {
	self->_Path = *reinterpret_cast<ImVector<ImVec2>*>(_Path);
}

ImDrawCmdHeader* ImDrawList__CmdHeader(const ImDrawList* self) {
	return new ImDrawCmdHeader(self->_CmdHeader);
}

void ImDrawList_set_CmdHeader(ImDrawList* self, ImDrawCmdHeader* _CmdHeader) {
	self->_CmdHeader = *_CmdHeader;
}

ImDrawListSplitter* ImDrawList__Splitter(const ImDrawList* self) {
	return new ImDrawListSplitter(self->_Splitter);
}

void ImDrawList_set_Splitter(ImDrawList* self, ImDrawListSplitter* _Splitter) {
	self->_Splitter = *_Splitter;
}

void* ImDrawList__ClipRectStack(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_ClipRectStack));
}

void ImDrawList_set_ClipRectStack(ImDrawList* self, void* _ClipRectStack) {
	self->_ClipRectStack = *reinterpret_cast<ImVector<ImVec4>*>(_ClipRectStack);
}

void* ImDrawList__TextureStack(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_TextureStack));
}

void ImDrawList_set_TextureStack(ImDrawList* self, void* _TextureStack) {
	self->_TextureStack = *reinterpret_cast<ImVector<ImTextureRef>*>(_TextureStack);
}

void* ImDrawList__CallbacksDataBuf(const ImDrawList* self) {
	return const_cast<void*>(static_cast<const void*>(&self->_CallbacksDataBuf));
}

void ImDrawList_set_CallbacksDataBuf(ImDrawList* self, void* _CallbacksDataBuf) {
	self->_CallbacksDataBuf = *reinterpret_cast<ImVector<unsigned char>*>(_CallbacksDataBuf);
}

float ImDrawList__FringeScale(const ImDrawList* self) {
	return self->_FringeScale;
}

void ImDrawList_set_FringeScale(ImDrawList* self, float _FringeScale) {
	self->_FringeScale = static_cast<float>(_FringeScale);
}

const char* ImDrawList__OwnerName(const ImDrawList* self) {
	return (const char*) self->_OwnerName;
}

void ImDrawList_set_OwnerName(ImDrawList* self, const char* _OwnerName) {
	self->_OwnerName = _OwnerName;
}

void ImDrawList_PushClipRect(ImDrawList* self, ImVec2* clip_rect_min, ImVec2* clip_rect_max) {
	self->PushClipRect(*clip_rect_min, *clip_rect_max);
}

void ImDrawList_PushClipRectFullScreen(ImDrawList* self) {
	self->PushClipRectFullScreen();
}

void ImDrawList_PopClipRect(ImDrawList* self) {
	self->PopClipRect();
}

void ImDrawList_PushTexture(ImDrawList* self, ImTextureRef* tex_ref) {
	self->PushTexture(*tex_ref);
}

void ImDrawList_PopTexture(ImDrawList* self) {
	self->PopTexture();
}

ImVec2* ImDrawList_GetClipRectMin(const ImDrawList* self) {
	return new ImVec2(self->GetClipRectMin());
}

ImVec2* ImDrawList_GetClipRectMax(const ImDrawList* self) {
	return new ImVec2(self->GetClipRectMax());
}

void ImDrawList_AddLine(ImDrawList* self, ImVec2* p1, ImVec2* p2, unsigned int col) {
	self->AddLine(*p1, *p2, static_cast<unsigned int>(col));
}

void ImDrawList_AddLineH(ImDrawList* self, float min_x, float max_x, float y, unsigned int col) {
	self->AddLineH(static_cast<float>(min_x), static_cast<float>(max_x), static_cast<float>(y), static_cast<unsigned int>(col));
}

void ImDrawList_AddLineV(ImDrawList* self, float x, float min_y, float max_y, unsigned int col) {
	self->AddLineV(static_cast<float>(x), static_cast<float>(min_y), static_cast<float>(max_y), static_cast<unsigned int>(col));
}

void ImDrawList_AddRect(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col) {
	self->AddRect(*p_min, *p_max, static_cast<unsigned int>(col));
}

void ImDrawList_AddRectFilled(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col) {
	self->AddRectFilled(*p_min, *p_max, static_cast<unsigned int>(col));
}

void ImDrawList_AddRectFilledMultiColor(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col_upr_left, unsigned int col_upr_right, unsigned int col_bot_right, unsigned int col_bot_left) {
	self->AddRectFilledMultiColor(*p_min, *p_max, static_cast<unsigned int>(col_upr_left), static_cast<unsigned int>(col_upr_right), static_cast<unsigned int>(col_bot_right), static_cast<unsigned int>(col_bot_left));
}

void ImDrawList_AddQuad(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col) {
	self->AddQuad(*p1, *p2, *p3, *p4, static_cast<unsigned int>(col));
}

void ImDrawList_AddQuadFilled(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col) {
	self->AddQuadFilled(*p1, *p2, *p3, *p4, static_cast<unsigned int>(col));
}

void ImDrawList_AddTriangle(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col) {
	self->AddTriangle(*p1, *p2, *p3, static_cast<unsigned int>(col));
}

void ImDrawList_AddTriangleFilled(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col) {
	self->AddTriangleFilled(*p1, *p2, *p3, static_cast<unsigned int>(col));
}

void ImDrawList_AddCircle(ImDrawList* self, ImVec2* center, float radius, unsigned int col) {
	self->AddCircle(*center, static_cast<float>(radius), static_cast<unsigned int>(col));
}

void ImDrawList_AddCircleFilled(ImDrawList* self, ImVec2* center, float radius, unsigned int col) {
	self->AddCircleFilled(*center, static_cast<float>(radius), static_cast<unsigned int>(col));
}

void ImDrawList_AddNgon(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments) {
	self->AddNgon(*center, static_cast<float>(radius), static_cast<unsigned int>(col), static_cast<int>(num_segments));
}

void ImDrawList_AddNgonFilled(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments) {
	self->AddNgonFilled(*center, static_cast<float>(radius), static_cast<unsigned int>(col), static_cast<int>(num_segments));
}

void ImDrawList_AddEllipse(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col) {
	self->AddEllipse(*center, *radius, static_cast<unsigned int>(col));
}

void ImDrawList_AddEllipseFilled(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col) {
	self->AddEllipseFilled(*center, *radius, static_cast<unsigned int>(col));
}

void ImDrawList_AddText(ImDrawList* self, ImVec2* pos, unsigned int col, const char* text_begin) {
	self->AddText(*pos, static_cast<unsigned int>(col), text_begin);
}

void ImDrawList_AddText2(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin) {
	self->AddText(font, static_cast<float>(font_size), *pos, static_cast<unsigned int>(col), text_begin);
}

void ImDrawList_AddBezierCubic(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col, float thickness) {
	self->AddBezierCubic(*p1, *p2, *p3, *p4, static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddBezierQuadratic(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col, float thickness) {
	self->AddBezierQuadratic(*p1, *p2, *p3, static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddPolyline(ImDrawList* self, ImVec2* points, int num_points, unsigned int col, float thickness) {
	self->AddPolyline(points, static_cast<int>(num_points), static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddConvexPolyFilled(ImDrawList* self, ImVec2* points, int num_points, unsigned int col) {
	self->AddConvexPolyFilled(points, static_cast<int>(num_points), static_cast<unsigned int>(col));
}

void ImDrawList_AddConcavePolyFilled(ImDrawList* self, ImVec2* points, int num_points, unsigned int col) {
	self->AddConcavePolyFilled(points, static_cast<int>(num_points), static_cast<unsigned int>(col));
}

void ImDrawList_AddImage(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max) {
	self->AddImage(*tex_ref, *p_min, *p_max);
}

void ImDrawList_AddImageQuad(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4) {
	self->AddImageQuad(*tex_ref, *p1, *p2, *p3, *p4);
}

void ImDrawList_AddImageRounded(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max, unsigned int col, float rounding) {
	self->AddImageRounded(*tex_ref, *p_min, *p_max, *uv_min, *uv_max, static_cast<unsigned int>(col), static_cast<float>(rounding));
}

void ImDrawList_PathClear(ImDrawList* self) {
	self->PathClear();
}

void ImDrawList_PathLineTo(ImDrawList* self, ImVec2* pos) {
	self->PathLineTo(*pos);
}

void ImDrawList_PathLineToMergeDuplicate(ImDrawList* self, ImVec2* pos) {
	self->PathLineToMergeDuplicate(*pos);
}

void ImDrawList_PathFillConvex(ImDrawList* self, unsigned int col) {
	self->PathFillConvex(static_cast<unsigned int>(col));
}

void ImDrawList_PathFillConcave(ImDrawList* self, unsigned int col) {
	self->PathFillConcave(static_cast<unsigned int>(col));
}

void ImDrawList_PathStroke(ImDrawList* self, unsigned int col) {
	self->PathStroke(static_cast<unsigned int>(col));
}

void ImDrawList_PathArcTo(ImDrawList* self, ImVec2* center, float radius, float a_min, float a_max) {
	self->PathArcTo(*center, static_cast<float>(radius), static_cast<float>(a_min), static_cast<float>(a_max));
}

void ImDrawList_PathArcToFast(ImDrawList* self, ImVec2* center, float radius, int a_min_of_12, int a_max_of_12) {
	self->PathArcToFast(*center, static_cast<float>(radius), static_cast<int>(a_min_of_12), static_cast<int>(a_max_of_12));
}

void ImDrawList_PathEllipticalArcTo(ImDrawList* self, ImVec2* center, ImVec2* radius, float rot, float a_min, float a_max) {
	self->PathEllipticalArcTo(*center, *radius, static_cast<float>(rot), static_cast<float>(a_min), static_cast<float>(a_max));
}

void ImDrawList_PathBezierCubicCurveTo(ImDrawList* self, ImVec2* p2, ImVec2* p3, ImVec2* p4) {
	self->PathBezierCubicCurveTo(*p2, *p3, *p4);
}

void ImDrawList_PathBezierQuadraticCurveTo(ImDrawList* self, ImVec2* p2, ImVec2* p3) {
	self->PathBezierQuadraticCurveTo(*p2, *p3);
}

void ImDrawList_PathRect(ImDrawList* self, ImVec2* rect_min, ImVec2* rect_max) {
	self->PathRect(*rect_min, *rect_max);
}

void ImDrawList_AddDrawCmd(ImDrawList* self) {
	self->AddDrawCmd();
}

ImDrawList* ImDrawList_CloneOutput(const ImDrawList* self) {
	return self->CloneOutput();
}

void ImDrawList_ChannelsSplit(ImDrawList* self, int count) {
	self->ChannelsSplit(static_cast<int>(count));
}

void ImDrawList_ChannelsMerge(ImDrawList* self) {
	self->ChannelsMerge();
}

void ImDrawList_ChannelsSetCurrent(ImDrawList* self, int n) {
	self->ChannelsSetCurrent(static_cast<int>(n));
}

void ImDrawList_PrimReserve(ImDrawList* self, int idx_count, int vtx_count) {
	self->PrimReserve(static_cast<int>(idx_count), static_cast<int>(vtx_count));
}

void ImDrawList_PrimUnreserve(ImDrawList* self, int idx_count, int vtx_count) {
	self->PrimUnreserve(static_cast<int>(idx_count), static_cast<int>(vtx_count));
}

void ImDrawList_PrimRect(ImDrawList* self, ImVec2* a, ImVec2* b, unsigned int col) {
	self->PrimRect(*a, *b, static_cast<unsigned int>(col));
}

void ImDrawList_PrimRectUV(ImDrawList* self, ImVec2* a, ImVec2* b, ImVec2* uv_a, ImVec2* uv_b, unsigned int col) {
	self->PrimRectUV(*a, *b, *uv_a, *uv_b, static_cast<unsigned int>(col));
}

void ImDrawList_PrimQuadUV(ImDrawList* self, ImVec2* a, ImVec2* b, ImVec2* c, ImVec2* d, ImVec2* uv_a, ImVec2* uv_b, ImVec2* uv_c, ImVec2* uv_d, unsigned int col) {
	self->PrimQuadUV(*a, *b, *c, *d, *uv_a, *uv_b, *uv_c, *uv_d, static_cast<unsigned int>(col));
}

void ImDrawList_PrimWriteVtx(ImDrawList* self, ImVec2* pos, ImVec2* uv, unsigned int col) {
	self->PrimWriteVtx(*pos, *uv, static_cast<unsigned int>(col));
}

void ImDrawList_PrimWriteIdx(ImDrawList* self, unsigned short idx) {
	self->PrimWriteIdx(static_cast<unsigned short>(idx));
}

void ImDrawList_PrimVtx(ImDrawList* self, ImVec2* pos, ImVec2* uv, unsigned int col) {
	self->PrimVtx(*pos, *uv, static_cast<unsigned int>(col));
}

void ImDrawList_AddRect2(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, int flags, float thickness) {
	self->AddRect(*p_min, *p_max, static_cast<unsigned int>(col), static_cast<float>(rounding), static_cast<int>(flags), static_cast<float>(thickness));
}

void ImDrawList_AddPolyline2(ImDrawList* self, ImVec2* points, int num_points, unsigned int col, int flags, float thickness) {
	self->AddPolyline(points, static_cast<int>(num_points), static_cast<unsigned int>(col), static_cast<int>(flags), static_cast<float>(thickness));
}

void ImDrawList_PathStroke2(ImDrawList* self, unsigned int col, int flags, float thickness) {
	self->PathStroke(static_cast<unsigned int>(col), static_cast<int>(flags), static_cast<float>(thickness));
}

void ImDrawList_PushTextureID(ImDrawList* self, ImTextureRef* tex_ref) {
	self->PushTextureID(*tex_ref);
}

void ImDrawList_PopTextureID(ImDrawList* self) {
	self->PopTextureID();
}

void ImDrawList__SetDrawListSharedData(ImDrawList* self, void* data) {
	self->_SetDrawListSharedData(static_cast<ImDrawListSharedData*>(data));
}

void ImDrawList__ResetForNewFrame(ImDrawList* self) {
	self->_ResetForNewFrame();
}

void ImDrawList__ClearFreeMemory(ImDrawList* self) {
	self->_ClearFreeMemory();
}

void ImDrawList__PopUnusedDrawCmd(ImDrawList* self) {
	self->_PopUnusedDrawCmd();
}

void ImDrawList__TryMergeDrawCmds(ImDrawList* self) {
	self->_TryMergeDrawCmds();
}

void ImDrawList__OnChangedClipRect(ImDrawList* self) {
	self->_OnChangedClipRect();
}

void ImDrawList__OnChangedTexture(ImDrawList* self) {
	self->_OnChangedTexture();
}

void ImDrawList__OnChangedVtxOffset(ImDrawList* self) {
	self->_OnChangedVtxOffset();
}

void ImDrawList__SetTexture(ImDrawList* self, ImTextureRef* tex_ref) {
	self->_SetTexture(*tex_ref);
}

int ImDrawList__CalcCircleAutoSegmentCount(const ImDrawList* self, float radius) {
	return self->_CalcCircleAutoSegmentCount(static_cast<float>(radius));
}

void ImDrawList__PathArcToFastEx(ImDrawList* self, ImVec2* center, float radius, int a_min_sample, int a_max_sample, int a_step) {
	self->_PathArcToFastEx(*center, static_cast<float>(radius), static_cast<int>(a_min_sample), static_cast<int>(a_max_sample), static_cast<int>(a_step));
}

void ImDrawList__PathArcToN(ImDrawList* self, ImVec2* center, float radius, float a_min, float a_max, int num_segments) {
	self->_PathArcToN(*center, static_cast<float>(radius), static_cast<float>(a_min), static_cast<float>(a_max), static_cast<int>(num_segments));
}

void ImDrawList_operatorAssign(ImDrawList* self, ImDrawList* param1) {
	self->operator=(*param1);
}

void ImDrawList_PushClipRect2(ImDrawList* self, ImVec2* clip_rect_min, ImVec2* clip_rect_max, bool intersect_with_current_clip_rect) {
	self->PushClipRect(*clip_rect_min, *clip_rect_max, intersect_with_current_clip_rect);
}

void ImDrawList_AddLine2(ImDrawList* self, ImVec2* p1, ImVec2* p2, unsigned int col, float thickness) {
	self->AddLine(*p1, *p2, static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddLineH2(ImDrawList* self, float min_x, float max_x, float y, unsigned int col, float thickness) {
	self->AddLineH(static_cast<float>(min_x), static_cast<float>(max_x), static_cast<float>(y), static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddLineV2(ImDrawList* self, float x, float min_y, float max_y, unsigned int col, float thickness) {
	self->AddLineV(static_cast<float>(x), static_cast<float>(min_y), static_cast<float>(max_y), static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddRect3(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding) {
	self->AddRect(*p_min, *p_max, static_cast<unsigned int>(col), static_cast<float>(rounding));
}

void ImDrawList_AddRect4(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, float thickness) {
	self->AddRect(*p_min, *p_max, static_cast<unsigned int>(col), static_cast<float>(rounding), static_cast<float>(thickness));
}

void ImDrawList_AddRect5(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, float thickness, int flags) {
	self->AddRect(*p_min, *p_max, static_cast<unsigned int>(col), static_cast<float>(rounding), static_cast<float>(thickness), static_cast<int>(flags));
}

void ImDrawList_AddRectFilled2(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding) {
	self->AddRectFilled(*p_min, *p_max, static_cast<unsigned int>(col), static_cast<float>(rounding));
}

void ImDrawList_AddRectFilled3(ImDrawList* self, ImVec2* p_min, ImVec2* p_max, unsigned int col, float rounding, int flags) {
	self->AddRectFilled(*p_min, *p_max, static_cast<unsigned int>(col), static_cast<float>(rounding), static_cast<int>(flags));
}

void ImDrawList_AddQuad2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col, float thickness) {
	self->AddQuad(*p1, *p2, *p3, *p4, static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddTriangle2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col, float thickness) {
	self->AddTriangle(*p1, *p2, *p3, static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_AddCircle2(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments) {
	self->AddCircle(*center, static_cast<float>(radius), static_cast<unsigned int>(col), static_cast<int>(num_segments));
}

void ImDrawList_AddCircle3(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments, float thickness) {
	self->AddCircle(*center, static_cast<float>(radius), static_cast<unsigned int>(col), static_cast<int>(num_segments), static_cast<float>(thickness));
}

void ImDrawList_AddCircleFilled2(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments) {
	self->AddCircleFilled(*center, static_cast<float>(radius), static_cast<unsigned int>(col), static_cast<int>(num_segments));
}

void ImDrawList_AddNgon2(ImDrawList* self, ImVec2* center, float radius, unsigned int col, int num_segments, float thickness) {
	self->AddNgon(*center, static_cast<float>(radius), static_cast<unsigned int>(col), static_cast<int>(num_segments), static_cast<float>(thickness));
}

void ImDrawList_AddEllipse2(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot) {
	self->AddEllipse(*center, *radius, static_cast<unsigned int>(col), static_cast<float>(rot));
}

void ImDrawList_AddEllipse3(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot, int num_segments) {
	self->AddEllipse(*center, *radius, static_cast<unsigned int>(col), static_cast<float>(rot), static_cast<int>(num_segments));
}

void ImDrawList_AddEllipse4(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot, int num_segments, float thickness) {
	self->AddEllipse(*center, *radius, static_cast<unsigned int>(col), static_cast<float>(rot), static_cast<int>(num_segments), static_cast<float>(thickness));
}

void ImDrawList_AddEllipseFilled2(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot) {
	self->AddEllipseFilled(*center, *radius, static_cast<unsigned int>(col), static_cast<float>(rot));
}

void ImDrawList_AddEllipseFilled3(ImDrawList* self, ImVec2* center, ImVec2* radius, unsigned int col, float rot, int num_segments) {
	self->AddEllipseFilled(*center, *radius, static_cast<unsigned int>(col), static_cast<float>(rot), static_cast<int>(num_segments));
}

void ImDrawList_AddText3(ImDrawList* self, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end) {
	self->AddText(*pos, static_cast<unsigned int>(col), text_begin, text_end);
}

void ImDrawList_AddText4(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end) {
	self->AddText(font, static_cast<float>(font_size), *pos, static_cast<unsigned int>(col), text_begin, text_end);
}

void ImDrawList_AddText5(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end, float wrap_width) {
	self->AddText(font, static_cast<float>(font_size), *pos, static_cast<unsigned int>(col), text_begin, text_end, static_cast<float>(wrap_width));
}

void ImDrawList_AddText6(ImDrawList* self, ImFont* font, float font_size, ImVec2* pos, unsigned int col, const char* text_begin, const char* text_end, float wrap_width, ImVec4* cpu_fine_clip_rect) {
	self->AddText(font, static_cast<float>(font_size), *pos, static_cast<unsigned int>(col), text_begin, text_end, static_cast<float>(wrap_width), cpu_fine_clip_rect);
}

void ImDrawList_AddBezierCubic2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, unsigned int col, float thickness, int num_segments) {
	self->AddBezierCubic(*p1, *p2, *p3, *p4, static_cast<unsigned int>(col), static_cast<float>(thickness), static_cast<int>(num_segments));
}

void ImDrawList_AddBezierQuadratic2(ImDrawList* self, ImVec2* p1, ImVec2* p2, ImVec2* p3, unsigned int col, float thickness, int num_segments) {
	self->AddBezierQuadratic(*p1, *p2, *p3, static_cast<unsigned int>(col), static_cast<float>(thickness), static_cast<int>(num_segments));
}

void ImDrawList_AddPolyline3(ImDrawList* self, ImVec2* points, int num_points, unsigned int col, float thickness, int flags) {
	self->AddPolyline(points, static_cast<int>(num_points), static_cast<unsigned int>(col), static_cast<float>(thickness), static_cast<int>(flags));
}

void ImDrawList_AddImage2(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min) {
	self->AddImage(*tex_ref, *p_min, *p_max, *uv_min);
}

void ImDrawList_AddImage3(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max) {
	self->AddImage(*tex_ref, *p_min, *p_max, *uv_min, *uv_max);
}

void ImDrawList_AddImage4(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max, unsigned int col) {
	self->AddImage(*tex_ref, *p_min, *p_max, *uv_min, *uv_max, static_cast<unsigned int>(col));
}

void ImDrawList_AddImageQuad2(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1) {
	self->AddImageQuad(*tex_ref, *p1, *p2, *p3, *p4, *uv1);
}

void ImDrawList_AddImageQuad3(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2) {
	self->AddImageQuad(*tex_ref, *p1, *p2, *p3, *p4, *uv1, *uv2);
}

void ImDrawList_AddImageQuad4(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2, ImVec2* uv3) {
	self->AddImageQuad(*tex_ref, *p1, *p2, *p3, *p4, *uv1, *uv2, *uv3);
}

void ImDrawList_AddImageQuad5(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2, ImVec2* uv3, ImVec2* uv4) {
	self->AddImageQuad(*tex_ref, *p1, *p2, *p3, *p4, *uv1, *uv2, *uv3, *uv4);
}

void ImDrawList_AddImageQuad6(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p1, ImVec2* p2, ImVec2* p3, ImVec2* p4, ImVec2* uv1, ImVec2* uv2, ImVec2* uv3, ImVec2* uv4, unsigned int col) {
	self->AddImageQuad(*tex_ref, *p1, *p2, *p3, *p4, *uv1, *uv2, *uv3, *uv4, static_cast<unsigned int>(col));
}

void ImDrawList_AddImageRounded2(ImDrawList* self, ImTextureRef* tex_ref, ImVec2* p_min, ImVec2* p_max, ImVec2* uv_min, ImVec2* uv_max, unsigned int col, float rounding, int flags) {
	self->AddImageRounded(*tex_ref, *p_min, *p_max, *uv_min, *uv_max, static_cast<unsigned int>(col), static_cast<float>(rounding), static_cast<int>(flags));
}

void ImDrawList_PathStroke3(ImDrawList* self, unsigned int col, float thickness) {
	self->PathStroke(static_cast<unsigned int>(col), static_cast<float>(thickness));
}

void ImDrawList_PathStroke4(ImDrawList* self, unsigned int col, float thickness, int flags) {
	self->PathStroke(static_cast<unsigned int>(col), static_cast<float>(thickness), static_cast<int>(flags));
}

void ImDrawList_PathArcTo2(ImDrawList* self, ImVec2* center, float radius, float a_min, float a_max, int num_segments) {
	self->PathArcTo(*center, static_cast<float>(radius), static_cast<float>(a_min), static_cast<float>(a_max), static_cast<int>(num_segments));
}

void ImDrawList_PathEllipticalArcTo2(ImDrawList* self, ImVec2* center, ImVec2* radius, float rot, float a_min, float a_max, int num_segments) {
	self->PathEllipticalArcTo(*center, *radius, static_cast<float>(rot), static_cast<float>(a_min), static_cast<float>(a_max), static_cast<int>(num_segments));
}

void ImDrawList_PathBezierCubicCurveTo2(ImDrawList* self, ImVec2* p2, ImVec2* p3, ImVec2* p4, int num_segments) {
	self->PathBezierCubicCurveTo(*p2, *p3, *p4, static_cast<int>(num_segments));
}

void ImDrawList_PathBezierQuadraticCurveTo2(ImDrawList* self, ImVec2* p2, ImVec2* p3, int num_segments) {
	self->PathBezierQuadraticCurveTo(*p2, *p3, static_cast<int>(num_segments));
}

void ImDrawList_PathRect2(ImDrawList* self, ImVec2* rect_min, ImVec2* rect_max, float rounding) {
	self->PathRect(*rect_min, *rect_max, static_cast<float>(rounding));
}

void ImDrawList_PathRect3(ImDrawList* self, ImVec2* rect_min, ImVec2* rect_max, float rounding, int flags) {
	self->PathRect(*rect_min, *rect_max, static_cast<float>(rounding), static_cast<int>(flags));
}

void ImDrawList_delete(ImDrawList* self) {
	delete self;
}

ImDrawData* ImDrawData_new() {
	return new (std::nothrow) ImDrawData();
}

ImDrawData* ImDrawData_new2(ImDrawData* param1) {
	return new (std::nothrow) ImDrawData(*param1);
}

bool ImDrawData_Valid(const ImDrawData* self) {
	return self->Valid;
}

void ImDrawData_setValid(ImDrawData* self, bool Valid) {
	self->Valid = Valid;
}

int ImDrawData_CmdListsCount(const ImDrawData* self) {
	return self->CmdListsCount;
}

void ImDrawData_setCmdListsCount(ImDrawData* self, int CmdListsCount) {
	self->CmdListsCount = static_cast<int>(CmdListsCount);
}

int ImDrawData_TotalIdxCount(const ImDrawData* self) {
	return self->TotalIdxCount;
}

void ImDrawData_setTotalIdxCount(ImDrawData* self, int TotalIdxCount) {
	self->TotalIdxCount = static_cast<int>(TotalIdxCount);
}

int ImDrawData_TotalVtxCount(const ImDrawData* self) {
	return self->TotalVtxCount;
}

void ImDrawData_setTotalVtxCount(ImDrawData* self, int TotalVtxCount) {
	self->TotalVtxCount = static_cast<int>(TotalVtxCount);
}

void* ImDrawData_CmdLists(const ImDrawData* self) {
	return const_cast<void*>(static_cast<const void*>(&self->CmdLists));
}

void ImDrawData_setCmdLists(ImDrawData* self, void* CmdLists) {
	self->CmdLists = *reinterpret_cast<ImVector<ImDrawList *>*>(CmdLists);
}

ImVec2* ImDrawData_DisplayPos(const ImDrawData* self) {
	return new ImVec2(self->DisplayPos);
}

void ImDrawData_setDisplayPos(ImDrawData* self, ImVec2* DisplayPos) {
	self->DisplayPos = *DisplayPos;
}

ImVec2* ImDrawData_DisplaySize(const ImDrawData* self) {
	return new ImVec2(self->DisplaySize);
}

void ImDrawData_setDisplaySize(ImDrawData* self, ImVec2* DisplaySize) {
	self->DisplaySize = *DisplaySize;
}

ImVec2* ImDrawData_FramebufferScale(const ImDrawData* self) {
	return new ImVec2(self->FramebufferScale);
}

void ImDrawData_setFramebufferScale(ImDrawData* self, ImVec2* FramebufferScale) {
	self->FramebufferScale = *FramebufferScale;
}

ImGuiViewport* ImDrawData_OwnerViewport(const ImDrawData* self) {
	return self->OwnerViewport;
}

void ImDrawData_setOwnerViewport(ImDrawData* self, ImGuiViewport* OwnerViewport) {
	self->OwnerViewport = OwnerViewport;
}

void* ImDrawData_Textures(const ImDrawData* self) {
	return self->Textures;
}

void ImDrawData_setTextures(ImDrawData* self, void* Textures) {
	self->Textures = reinterpret_cast<ImVector<ImTextureData *>*>(Textures);
}

void ImDrawData_Clear(ImDrawData* self) {
	self->Clear();
}

void ImDrawData_AddDrawList(ImDrawData* self, ImDrawList* draw_list) {
	self->AddDrawList(draw_list);
}

void ImDrawData_DeIndexAllBuffers(ImDrawData* self) {
	self->DeIndexAllBuffers();
}

void ImDrawData_ScaleClipRects(ImDrawData* self, ImVec2* fb_scale) {
	self->ScaleClipRects(*fb_scale);
}

void ImDrawData_operatorAssign(ImDrawData* self, ImDrawData* param1) {
	self->operator=(*param1);
}

void ImDrawData_delete(ImDrawData* self) {
	delete self;
}

ImTextureRect* ImTextureRect_new(ImTextureRect* param1) {
	return new (std::nothrow) ImTextureRect(*param1);
}

ImTextureRect* ImTextureRect_new2() {
	return new (std::nothrow) ImTextureRect();
}

unsigned short ImTextureRect_x(const ImTextureRect* self) {
	return self->x;
}

void ImTextureRect_setX(ImTextureRect* self, unsigned short x) {
	self->x = static_cast<unsigned short>(x);
}

unsigned short ImTextureRect_y(const ImTextureRect* self) {
	return self->y;
}

void ImTextureRect_setY(ImTextureRect* self, unsigned short y) {
	self->y = static_cast<unsigned short>(y);
}

unsigned short ImTextureRect_w(const ImTextureRect* self) {
	return self->w;
}

void ImTextureRect_setW(ImTextureRect* self, unsigned short w) {
	self->w = static_cast<unsigned short>(w);
}

unsigned short ImTextureRect_h(const ImTextureRect* self) {
	return self->h;
}

void ImTextureRect_setH(ImTextureRect* self, unsigned short h) {
	self->h = static_cast<unsigned short>(h);
}

void ImTextureRect_operatorAssign(ImTextureRect* self, ImTextureRect* param1) {
	self->operator=(*param1);
}

void ImTextureRect_delete(ImTextureRect* self) {
	delete self;
}

ImTextureData* ImTextureData_new() {
	return new (std::nothrow) ImTextureData();
}

ImTextureData* ImTextureData_new2(ImTextureData* param1) {
	return new (std::nothrow) ImTextureData(*param1);
}

int ImTextureData_UniqueID(const ImTextureData* self) {
	return self->UniqueID;
}

void ImTextureData_setUniqueID(ImTextureData* self, int UniqueID) {
	self->UniqueID = static_cast<int>(UniqueID);
}

int ImTextureData_Status(const ImTextureData* self) {
	ImTextureStatus Status_ret = self->Status;
	return static_cast<int>(Status_ret);
}

void ImTextureData_setStatus(ImTextureData* self, int Status) {
	self->Status = static_cast<ImTextureStatus>(Status);
}

unsigned long long ImTextureData_TexID(const ImTextureData* self) {
	return self->TexID;
}

void ImTextureData_setTexID(ImTextureData* self, unsigned long long TexID) {
	self->TexID = static_cast<unsigned long long>(TexID);
}

int ImTextureData_Format(const ImTextureData* self) {
	ImTextureFormat Format_ret = self->Format;
	return static_cast<int>(Format_ret);
}

void ImTextureData_setFormat(ImTextureData* self, int Format) {
	self->Format = static_cast<ImTextureFormat>(Format);
}

int ImTextureData_Width(const ImTextureData* self) {
	return self->Width;
}

void ImTextureData_setWidth(ImTextureData* self, int Width) {
	self->Width = static_cast<int>(Width);
}

int ImTextureData_Height(const ImTextureData* self) {
	return self->Height;
}

void ImTextureData_setHeight(ImTextureData* self, int Height) {
	self->Height = static_cast<int>(Height);
}

int ImTextureData_BytesPerPixel(const ImTextureData* self) {
	return self->BytesPerPixel;
}

void ImTextureData_setBytesPerPixel(ImTextureData* self, int BytesPerPixel) {
	self->BytesPerPixel = static_cast<int>(BytesPerPixel);
}

unsigned char* ImTextureData_Pixels(const ImTextureData* self) {
	return self->Pixels;
}

void ImTextureData_setPixels(ImTextureData* self, unsigned char* Pixels) {
	self->Pixels = Pixels;
}

ImTextureRect* ImTextureData_UsedRect(const ImTextureData* self) {
	return new ImTextureRect(self->UsedRect);
}

void ImTextureData_setUsedRect(ImTextureData* self, ImTextureRect* UsedRect) {
	self->UsedRect = *UsedRect;
}

ImTextureRect* ImTextureData_UpdateRect(const ImTextureData* self) {
	return new ImTextureRect(self->UpdateRect);
}

void ImTextureData_setUpdateRect(ImTextureData* self, ImTextureRect* UpdateRect) {
	self->UpdateRect = *UpdateRect;
}

void* ImTextureData_Updates(const ImTextureData* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Updates));
}

void ImTextureData_setUpdates(ImTextureData* self, void* Updates) {
	self->Updates = *reinterpret_cast<ImVector<ImTextureRect>*>(Updates);
}

int ImTextureData_UnusedFrames(const ImTextureData* self) {
	return self->UnusedFrames;
}

void ImTextureData_setUnusedFrames(ImTextureData* self, int UnusedFrames) {
	self->UnusedFrames = static_cast<int>(UnusedFrames);
}

unsigned short ImTextureData_RefCount(const ImTextureData* self) {
	return self->RefCount;
}

void ImTextureData_setRefCount(ImTextureData* self, unsigned short RefCount) {
	self->RefCount = static_cast<unsigned short>(RefCount);
}

bool ImTextureData_UseColors(const ImTextureData* self) {
	return self->UseColors;
}

void ImTextureData_setUseColors(ImTextureData* self, bool UseColors) {
	self->UseColors = UseColors;
}

bool ImTextureData_WantDestroyNextFrame(const ImTextureData* self) {
	return self->WantDestroyNextFrame;
}

void ImTextureData_setWantDestroyNextFrame(ImTextureData* self, bool WantDestroyNextFrame) {
	self->WantDestroyNextFrame = WantDestroyNextFrame;
}

void ImTextureData_Create(ImTextureData* self, int format, int w, int h) {
	self->Create(static_cast<ImTextureFormat>(format), static_cast<int>(w), static_cast<int>(h));
}

void ImTextureData_DestroyPixels(ImTextureData* self) {
	self->DestroyPixels();
}

void* ImTextureData_GetPixels(ImTextureData* self) {
	return self->GetPixels();
}

void* ImTextureData_GetPixelsAt(ImTextureData* self, int x, int y) {
	return self->GetPixelsAt(static_cast<int>(x), static_cast<int>(y));
}

int ImTextureData_GetSizeInBytes(const ImTextureData* self) {
	return self->GetSizeInBytes();
}

int ImTextureData_GetPitch(const ImTextureData* self) {
	return self->GetPitch();
}

ImTextureRef* ImTextureData_GetTexRef(ImTextureData* self) {
	return new ImTextureRef(self->GetTexRef());
}

unsigned long long ImTextureData_GetTexID(const ImTextureData* self) {
	unsigned long long _ret = self->GetTexID();
	return static_cast<unsigned long long>(_ret);
}

void ImTextureData_SetTexID(ImTextureData* self, unsigned long long tex_id) {
	self->SetTexID(static_cast<unsigned long long>(tex_id));
}

void ImTextureData_SetStatus(ImTextureData* self, int status) {
	self->SetStatus(static_cast<ImTextureStatus>(status));
}

void ImTextureData_operatorAssign(ImTextureData* self, ImTextureData* param1) {
	self->operator=(*param1);
}

void ImTextureData_delete(ImTextureData* self) {
	delete self;
}

ImFontConfig* ImFontConfig_new() {
	return new (std::nothrow) ImFontConfig();
}

int ImFontConfig_FontDataSize(const ImFontConfig* self) {
	return self->FontDataSize;
}

void ImFontConfig_setFontDataSize(ImFontConfig* self, int FontDataSize) {
	self->FontDataSize = static_cast<int>(FontDataSize);
}

bool ImFontConfig_FontDataOwnedByAtlas(const ImFontConfig* self) {
	return self->FontDataOwnedByAtlas;
}

void ImFontConfig_setFontDataOwnedByAtlas(ImFontConfig* self, bool FontDataOwnedByAtlas) {
	self->FontDataOwnedByAtlas = FontDataOwnedByAtlas;
}

bool ImFontConfig_MergeMode(const ImFontConfig* self) {
	return self->MergeMode;
}

void ImFontConfig_setMergeMode(ImFontConfig* self, bool MergeMode) {
	self->MergeMode = MergeMode;
}

bool ImFontConfig_PixelSnapH(const ImFontConfig* self) {
	return self->PixelSnapH;
}

void ImFontConfig_setPixelSnapH(ImFontConfig* self, bool PixelSnapH) {
	self->PixelSnapH = PixelSnapH;
}

signed char ImFontConfig_OversampleH(const ImFontConfig* self) {
	return self->OversampleH;
}

void ImFontConfig_setOversampleH(ImFontConfig* self, signed char OversampleH) {
	self->OversampleH = static_cast<signed char>(OversampleH);
}

signed char ImFontConfig_OversampleV(const ImFontConfig* self) {
	return self->OversampleV;
}

void ImFontConfig_setOversampleV(ImFontConfig* self, signed char OversampleV) {
	self->OversampleV = static_cast<signed char>(OversampleV);
}

unsigned short ImFontConfig_EllipsisChar(const ImFontConfig* self) {
	return self->EllipsisChar;
}

void ImFontConfig_setEllipsisChar(ImFontConfig* self, unsigned short EllipsisChar) {
	self->EllipsisChar = static_cast<unsigned short>(EllipsisChar);
}

float ImFontConfig_SizePixels(const ImFontConfig* self) {
	return self->SizePixels;
}

void ImFontConfig_setSizePixels(ImFontConfig* self, float SizePixels) {
	self->SizePixels = static_cast<float>(SizePixels);
}

const unsigned short* ImFontConfig_GlyphRanges(const ImFontConfig* self) {
	const unsigned short* GlyphRanges_ret = self->GlyphRanges;
	return static_cast<const unsigned short*>(GlyphRanges_ret);
}

void ImFontConfig_setGlyphRanges(ImFontConfig* self, const unsigned short* GlyphRanges) {
	self->GlyphRanges = GlyphRanges;
}

const unsigned short* ImFontConfig_GlyphExcludeRanges(const ImFontConfig* self) {
	const unsigned short* GlyphExcludeRanges_ret = self->GlyphExcludeRanges;
	return static_cast<const unsigned short*>(GlyphExcludeRanges_ret);
}

void ImFontConfig_setGlyphExcludeRanges(ImFontConfig* self, const unsigned short* GlyphExcludeRanges) {
	self->GlyphExcludeRanges = GlyphExcludeRanges;
}

ImVec2* ImFontConfig_GlyphOffset(const ImFontConfig* self) {
	return new ImVec2(self->GlyphOffset);
}

void ImFontConfig_setGlyphOffset(ImFontConfig* self, ImVec2* GlyphOffset) {
	self->GlyphOffset = *GlyphOffset;
}

float ImFontConfig_GlyphMinAdvanceX(const ImFontConfig* self) {
	return self->GlyphMinAdvanceX;
}

void ImFontConfig_setGlyphMinAdvanceX(ImFontConfig* self, float GlyphMinAdvanceX) {
	self->GlyphMinAdvanceX = static_cast<float>(GlyphMinAdvanceX);
}

float ImFontConfig_GlyphMaxAdvanceX(const ImFontConfig* self) {
	return self->GlyphMaxAdvanceX;
}

void ImFontConfig_setGlyphMaxAdvanceX(ImFontConfig* self, float GlyphMaxAdvanceX) {
	self->GlyphMaxAdvanceX = static_cast<float>(GlyphMaxAdvanceX);
}

float ImFontConfig_GlyphExtraAdvanceX(const ImFontConfig* self) {
	return self->GlyphExtraAdvanceX;
}

void ImFontConfig_setGlyphExtraAdvanceX(ImFontConfig* self, float GlyphExtraAdvanceX) {
	self->GlyphExtraAdvanceX = static_cast<float>(GlyphExtraAdvanceX);
}

unsigned int ImFontConfig_FontNo(const ImFontConfig* self) {
	return self->FontNo;
}

void ImFontConfig_setFontNo(ImFontConfig* self, unsigned int FontNo) {
	self->FontNo = static_cast<unsigned int>(FontNo);
}

unsigned int ImFontConfig_FontLoaderFlags(const ImFontConfig* self) {
	return self->FontLoaderFlags;
}

void ImFontConfig_setFontLoaderFlags(ImFontConfig* self, unsigned int FontLoaderFlags) {
	self->FontLoaderFlags = static_cast<unsigned int>(FontLoaderFlags);
}

float ImFontConfig_RasterizerMultiply(const ImFontConfig* self) {
	return self->RasterizerMultiply;
}

void ImFontConfig_setRasterizerMultiply(ImFontConfig* self, float RasterizerMultiply) {
	self->RasterizerMultiply = static_cast<float>(RasterizerMultiply);
}

float ImFontConfig_RasterizerDensity(const ImFontConfig* self) {
	return self->RasterizerDensity;
}

void ImFontConfig_setRasterizerDensity(ImFontConfig* self, float RasterizerDensity) {
	self->RasterizerDensity = static_cast<float>(RasterizerDensity);
}

float ImFontConfig_ExtraSizeScale(const ImFontConfig* self) {
	return self->ExtraSizeScale;
}

void ImFontConfig_setExtraSizeScale(ImFontConfig* self, float ExtraSizeScale) {
	self->ExtraSizeScale = static_cast<float>(ExtraSizeScale);
}

int ImFontConfig_Flags(const ImFontConfig* self) {
	return self->Flags;
}

void ImFontConfig_setFlags(ImFontConfig* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

ImFont* ImFontConfig_DstFont(const ImFontConfig* self) {
	return self->DstFont;
}

void ImFontConfig_setDstFont(ImFontConfig* self, ImFont* DstFont) {
	self->DstFont = DstFont;
}

const ImFontLoader* ImFontConfig_FontLoader(const ImFontConfig* self) {
	return (const ImFontLoader*) self->FontLoader;
}

void ImFontConfig_setFontLoader(ImFontConfig* self, const ImFontLoader* FontLoader) {
	self->FontLoader = FontLoader;
}

bool ImFontConfig_PixelSnapV(const ImFontConfig* self) {
	return self->PixelSnapV;
}

void ImFontConfig_setPixelSnapV(ImFontConfig* self, bool PixelSnapV) {
	self->PixelSnapV = PixelSnapV;
}

void ImFontConfig_delete(ImFontConfig* self) {
	delete self;
}

ImFontGlyph* ImFontGlyph_new() {
	return new (std::nothrow) ImFontGlyph();
}

unsigned int ImFontGlyph_Colored(const ImFontGlyph* self) {
	return self->Colored;
}

void ImFontGlyph_setColored(ImFontGlyph* self, unsigned int Colored) {
	self->Colored = static_cast<unsigned int>(Colored);
}

unsigned int ImFontGlyph_Visible(const ImFontGlyph* self) {
	return self->Visible;
}

void ImFontGlyph_setVisible(ImFontGlyph* self, unsigned int Visible) {
	self->Visible = static_cast<unsigned int>(Visible);
}

unsigned int ImFontGlyph_SourceIdx(const ImFontGlyph* self) {
	return self->SourceIdx;
}

void ImFontGlyph_setSourceIdx(ImFontGlyph* self, unsigned int SourceIdx) {
	self->SourceIdx = static_cast<unsigned int>(SourceIdx);
}

unsigned int ImFontGlyph_Codepoint(const ImFontGlyph* self) {
	return self->Codepoint;
}

void ImFontGlyph_setCodepoint(ImFontGlyph* self, unsigned int Codepoint) {
	self->Codepoint = static_cast<unsigned int>(Codepoint);
}

float ImFontGlyph_AdvanceX(const ImFontGlyph* self) {
	return self->AdvanceX;
}

void ImFontGlyph_setAdvanceX(ImFontGlyph* self, float AdvanceX) {
	self->AdvanceX = static_cast<float>(AdvanceX);
}

float ImFontGlyph_X0(const ImFontGlyph* self) {
	return self->X0;
}

void ImFontGlyph_setX0(ImFontGlyph* self, float X0) {
	self->X0 = static_cast<float>(X0);
}

float ImFontGlyph_Y0(const ImFontGlyph* self) {
	return self->Y0;
}

void ImFontGlyph_setY0(ImFontGlyph* self, float Y0) {
	self->Y0 = static_cast<float>(Y0);
}

float ImFontGlyph_X1(const ImFontGlyph* self) {
	return self->X1;
}

void ImFontGlyph_setX1(ImFontGlyph* self, float X1) {
	self->X1 = static_cast<float>(X1);
}

float ImFontGlyph_Y1(const ImFontGlyph* self) {
	return self->Y1;
}

void ImFontGlyph_setY1(ImFontGlyph* self, float Y1) {
	self->Y1 = static_cast<float>(Y1);
}

float ImFontGlyph_U0(const ImFontGlyph* self) {
	return self->U0;
}

void ImFontGlyph_setU0(ImFontGlyph* self, float U0) {
	self->U0 = static_cast<float>(U0);
}

float ImFontGlyph_V0(const ImFontGlyph* self) {
	return self->V0;
}

void ImFontGlyph_setV0(ImFontGlyph* self, float V0) {
	self->V0 = static_cast<float>(V0);
}

float ImFontGlyph_U1(const ImFontGlyph* self) {
	return self->U1;
}

void ImFontGlyph_setU1(ImFontGlyph* self, float U1) {
	self->U1 = static_cast<float>(U1);
}

float ImFontGlyph_V1(const ImFontGlyph* self) {
	return self->V1;
}

void ImFontGlyph_setV1(ImFontGlyph* self, float V1) {
	self->V1 = static_cast<float>(V1);
}

int ImFontGlyph_PackId(const ImFontGlyph* self) {
	return self->PackId;
}

void ImFontGlyph_setPackId(ImFontGlyph* self, int PackId) {
	self->PackId = static_cast<int>(PackId);
}

void ImFontGlyph_delete(ImFontGlyph* self) {
	delete self;
}

ImFontGlyphRangesBuilder* ImFontGlyphRangesBuilder_new() {
	return new (std::nothrow) ImFontGlyphRangesBuilder();
}

ImFontGlyphRangesBuilder* ImFontGlyphRangesBuilder_new2(ImFontGlyphRangesBuilder* param1) {
	return new (std::nothrow) ImFontGlyphRangesBuilder(*param1);
}

void* ImFontGlyphRangesBuilder_UsedChars(const ImFontGlyphRangesBuilder* self) {
	return const_cast<void*>(static_cast<const void*>(&self->UsedChars));
}

void ImFontGlyphRangesBuilder_setUsedChars(ImFontGlyphRangesBuilder* self, void* UsedChars) {
	self->UsedChars = *reinterpret_cast<ImVector<unsigned int>*>(UsedChars);
}

void ImFontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self) {
	self->Clear();
}

bool ImFontGlyphRangesBuilder_GetBit(const ImFontGlyphRangesBuilder* self, unsigned long long n) {
	return self->GetBit(static_cast<unsigned long long>(n));
}

void ImFontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self, unsigned long long n) {
	self->SetBit(static_cast<unsigned long long>(n));
}

void ImFontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self, unsigned short c) {
	self->AddChar(static_cast<unsigned short>(c));
}

void ImFontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self, const char* text) {
	self->AddText(text);
}

void ImFontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self, const unsigned short* ranges) {
	self->AddRanges(ranges);
}

void ImFontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self, void* out_ranges) {
	self->BuildRanges(reinterpret_cast<ImVector<ImWchar>*>(out_ranges));
}

void ImFontGlyphRangesBuilder_operatorAssign(ImFontGlyphRangesBuilder* self, ImFontGlyphRangesBuilder* param1) {
	self->operator=(*param1);
}

void ImFontGlyphRangesBuilder_AddText2(ImFontGlyphRangesBuilder* self, const char* text, const char* text_end) {
	self->AddText(text, text_end);
}

void ImFontGlyphRangesBuilder_delete(ImFontGlyphRangesBuilder* self) {
	delete self;
}

ImFontAtlasRect* ImFontAtlasRect_new() {
	return new (std::nothrow) ImFontAtlasRect();
}

ImFontAtlasRect* ImFontAtlasRect_new2(ImFontAtlasRect* param1) {
	return new (std::nothrow) ImFontAtlasRect(*param1);
}

unsigned short ImFontAtlasRect_x(const ImFontAtlasRect* self) {
	return self->x;
}

void ImFontAtlasRect_setX(ImFontAtlasRect* self, unsigned short x) {
	self->x = static_cast<unsigned short>(x);
}

unsigned short ImFontAtlasRect_y(const ImFontAtlasRect* self) {
	return self->y;
}

void ImFontAtlasRect_setY(ImFontAtlasRect* self, unsigned short y) {
	self->y = static_cast<unsigned short>(y);
}

unsigned short ImFontAtlasRect_w(const ImFontAtlasRect* self) {
	return self->w;
}

void ImFontAtlasRect_setW(ImFontAtlasRect* self, unsigned short w) {
	self->w = static_cast<unsigned short>(w);
}

unsigned short ImFontAtlasRect_h(const ImFontAtlasRect* self) {
	return self->h;
}

void ImFontAtlasRect_setH(ImFontAtlasRect* self, unsigned short h) {
	self->h = static_cast<unsigned short>(h);
}

ImVec2* ImFontAtlasRect_uv0(const ImFontAtlasRect* self) {
	return new ImVec2(self->uv0);
}

void ImFontAtlasRect_setUv0(ImFontAtlasRect* self, ImVec2* uv0) {
	self->uv0 = *uv0;
}

ImVec2* ImFontAtlasRect_uv1(const ImFontAtlasRect* self) {
	return new ImVec2(self->uv1);
}

void ImFontAtlasRect_setUv1(ImFontAtlasRect* self, ImVec2* uv1) {
	self->uv1 = *uv1;
}

void ImFontAtlasRect_operatorAssign(ImFontAtlasRect* self, ImFontAtlasRect* param1) {
	self->operator=(*param1);
}

void ImFontAtlasRect_delete(ImFontAtlasRect* self) {
	delete self;
}

ImFontAtlas* ImFontAtlas_new() {
	return new (std::nothrow) ImFontAtlas();
}

ImFontAtlas* ImFontAtlas_new2(ImFontAtlas* param1) {
	return new (std::nothrow) ImFontAtlas(*param1);
}

ImFont* ImFontAtlas_AddFont(ImFontAtlas* self, ImFontConfig* font_cfg) {
	return self->AddFont(font_cfg);
}

ImFont* ImFontAtlas_AddFontDefault(ImFontAtlas* self) {
	return self->AddFontDefault();
}

ImFont* ImFontAtlas_AddFontDefaultVector(ImFontAtlas* self) {
	return self->AddFontDefaultVector();
}

ImFont* ImFontAtlas_AddFontDefaultBitmap(ImFontAtlas* self) {
	return self->AddFontDefaultBitmap();
}

ImFont* ImFontAtlas_AddFontFromFileTTF(ImFontAtlas* self, const char* filename) {
	return self->AddFontFromFileTTF(filename);
}

ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self, const char* compressed_font_data_base85) {
	return self->AddFontFromMemoryCompressedBase85TTF(compressed_font_data_base85);
}

void ImFontAtlas_RemoveFont(ImFontAtlas* self, ImFont* font) {
	self->RemoveFont(font);
}

void ImFontAtlas_Clear(ImFontAtlas* self) {
	self->Clear();
}

void ImFontAtlas_ClearFonts(ImFontAtlas* self) {
	self->ClearFonts();
}

void ImFontAtlas_CompactCache(ImFontAtlas* self) {
	self->CompactCache();
}

void ImFontAtlas_SetFontLoader(ImFontAtlas* self, const ImFontLoader* font_loader) {
	self->SetFontLoader(font_loader);
}

void ImFontAtlas_ClearInputData(ImFontAtlas* self) {
	self->ClearInputData();
}

void ImFontAtlas_ClearTexData(ImFontAtlas* self) {
	self->ClearTexData();
}

bool ImFontAtlas_Build(ImFontAtlas* self) {
	return self->Build();
}

void ImFontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height) {
	self->GetTexDataAsAlpha8(out_pixels, out_width, out_height);
}

void ImFontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height) {
	self->GetTexDataAsRGBA32(out_pixels, out_width, out_height);
}

void ImFontAtlas_SetTexID(ImFontAtlas* self, unsigned long long id) {
	self->SetTexID(static_cast<unsigned long long>(id));
}

void ImFontAtlas_SetTexIDWithId(ImFontAtlas* self, ImTextureRef* id) {
	self->SetTexID(*id);
}

bool ImFontAtlas_IsBuilt(const ImFontAtlas* self) {
	return self->IsBuilt();
}

const unsigned short* ImFontAtlas_GetGlyphRangesDefault(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesDefault();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesGreek(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesGreek();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesKorean(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesKorean();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesJapanese();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesChineseFull();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesChineseSimplifiedCommon();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesCyrillic();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesThai(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesThai();
	return static_cast<const unsigned short*>(_ret);
}

const unsigned short* ImFontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self) {
	const unsigned short* _ret = self->GetGlyphRangesVietnamese();
	return static_cast<const unsigned short*>(_ret);
}

int ImFontAtlas_AddCustomRect(ImFontAtlas* self, int width, int height) {
	int _ret = self->AddCustomRect(static_cast<int>(width), static_cast<int>(height));
	return static_cast<int>(_ret);
}

void ImFontAtlas_RemoveCustomRect(ImFontAtlas* self, int id) {
	self->RemoveCustomRect(static_cast<int>(id));
}

bool ImFontAtlas_GetCustomRect(const ImFontAtlas* self, int id, ImFontAtlasRect* out_r) {
	return self->GetCustomRect(static_cast<int>(id), out_r);
}

int ImFontAtlas_Flags(const ImFontAtlas* self) {
	return self->Flags;
}

void ImFontAtlas_setFlags(ImFontAtlas* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

int ImFontAtlas_TexDesiredFormat(const ImFontAtlas* self) {
	ImTextureFormat TexDesiredFormat_ret = self->TexDesiredFormat;
	return static_cast<int>(TexDesiredFormat_ret);
}

void ImFontAtlas_setTexDesiredFormat(ImFontAtlas* self, int TexDesiredFormat) {
	self->TexDesiredFormat = static_cast<ImTextureFormat>(TexDesiredFormat);
}

int ImFontAtlas_TexGlyphPadding(const ImFontAtlas* self) {
	return self->TexGlyphPadding;
}

void ImFontAtlas_setTexGlyphPadding(ImFontAtlas* self, int TexGlyphPadding) {
	self->TexGlyphPadding = static_cast<int>(TexGlyphPadding);
}

int ImFontAtlas_TexMinWidth(const ImFontAtlas* self) {
	return self->TexMinWidth;
}

void ImFontAtlas_setTexMinWidth(ImFontAtlas* self, int TexMinWidth) {
	self->TexMinWidth = static_cast<int>(TexMinWidth);
}

int ImFontAtlas_TexMinHeight(const ImFontAtlas* self) {
	return self->TexMinHeight;
}

void ImFontAtlas_setTexMinHeight(ImFontAtlas* self, int TexMinHeight) {
	self->TexMinHeight = static_cast<int>(TexMinHeight);
}

int ImFontAtlas_TexMaxWidth(const ImFontAtlas* self) {
	return self->TexMaxWidth;
}

void ImFontAtlas_setTexMaxWidth(ImFontAtlas* self, int TexMaxWidth) {
	self->TexMaxWidth = static_cast<int>(TexMaxWidth);
}

int ImFontAtlas_TexMaxHeight(const ImFontAtlas* self) {
	return self->TexMaxHeight;
}

void ImFontAtlas_setTexMaxHeight(ImFontAtlas* self, int TexMaxHeight) {
	self->TexMaxHeight = static_cast<int>(TexMaxHeight);
}

ImTextureData* ImFontAtlas_TexData(const ImFontAtlas* self) {
	return self->TexData;
}

void ImFontAtlas_setTexData(ImFontAtlas* self, ImTextureData* TexData) {
	self->TexData = TexData;
}

void* ImFontAtlas_TexList(const ImFontAtlas* self) {
	return const_cast<void*>(static_cast<const void*>(&self->TexList));
}

void ImFontAtlas_setTexList(ImFontAtlas* self, void* TexList) {
	self->TexList = *reinterpret_cast<ImVector<ImTextureData *>*>(TexList);
}

bool ImFontAtlas_Locked(const ImFontAtlas* self) {
	return self->Locked;
}

void ImFontAtlas_setLocked(ImFontAtlas* self, bool Locked) {
	self->Locked = Locked;
}

bool ImFontAtlas_RendererHasTextures(const ImFontAtlas* self) {
	return self->RendererHasTextures;
}

void ImFontAtlas_setRendererHasTextures(ImFontAtlas* self, bool RendererHasTextures) {
	self->RendererHasTextures = RendererHasTextures;
}

bool ImFontAtlas_TexIsBuilt(const ImFontAtlas* self) {
	return self->TexIsBuilt;
}

void ImFontAtlas_setTexIsBuilt(ImFontAtlas* self, bool TexIsBuilt) {
	self->TexIsBuilt = TexIsBuilt;
}

bool ImFontAtlas_TexPixelsUseColors(const ImFontAtlas* self) {
	return self->TexPixelsUseColors;
}

void ImFontAtlas_setTexPixelsUseColors(ImFontAtlas* self, bool TexPixelsUseColors) {
	self->TexPixelsUseColors = TexPixelsUseColors;
}

ImVec2* ImFontAtlas_TexUvScale(const ImFontAtlas* self) {
	return new ImVec2(self->TexUvScale);
}

void ImFontAtlas_setTexUvScale(ImFontAtlas* self, ImVec2* TexUvScale) {
	self->TexUvScale = *TexUvScale;
}

ImVec2* ImFontAtlas_TexUvWhitePixel(const ImFontAtlas* self) {
	return new ImVec2(self->TexUvWhitePixel);
}

void ImFontAtlas_setTexUvWhitePixel(ImFontAtlas* self, ImVec2* TexUvWhitePixel) {
	self->TexUvWhitePixel = *TexUvWhitePixel;
}

void* ImFontAtlas_Fonts(const ImFontAtlas* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Fonts));
}

void ImFontAtlas_setFonts(ImFontAtlas* self, void* Fonts) {
	self->Fonts = *reinterpret_cast<ImVector<ImFont *>*>(Fonts);
}

void* ImFontAtlas_Sources(const ImFontAtlas* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Sources));
}

void ImFontAtlas_setSources(ImFontAtlas* self, void* Sources) {
	self->Sources = *reinterpret_cast<ImVector<ImFontConfig>*>(Sources);
}

ImVec4* ImFontAtlas_TexUvLines(const ImFontAtlas* self) {
	return (ImVec4*)self->TexUvLines;
}

void ImFontAtlas_setTexUvLines(ImFontAtlas* self, const ImVec4* TexUvLines) {
	memcpy(self->TexUvLines, TexUvLines, sizeof(self->TexUvLines));
}

int ImFontAtlas_TexNextUniqueID(const ImFontAtlas* self) {
	return self->TexNextUniqueID;
}

void ImFontAtlas_setTexNextUniqueID(ImFontAtlas* self, int TexNextUniqueID) {
	self->TexNextUniqueID = static_cast<int>(TexNextUniqueID);
}

int ImFontAtlas_FontNextUniqueID(const ImFontAtlas* self) {
	return self->FontNextUniqueID;
}

void ImFontAtlas_setFontNextUniqueID(ImFontAtlas* self, int FontNextUniqueID) {
	self->FontNextUniqueID = static_cast<int>(FontNextUniqueID);
}

void* ImFontAtlas_DrawListSharedDatas(const ImFontAtlas* self) {
	return const_cast<void*>(static_cast<const void*>(&self->DrawListSharedDatas));
}

void ImFontAtlas_setDrawListSharedDatas(ImFontAtlas* self, void* DrawListSharedDatas) {
	self->DrawListSharedDatas = *reinterpret_cast<ImVector<ImDrawListSharedData *>*>(DrawListSharedDatas);
}

void* ImFontAtlas_Builder(const ImFontAtlas* self) {
	return self->Builder;
}

void ImFontAtlas_setBuilder(ImFontAtlas* self, void* Builder) {
	self->Builder = static_cast<ImFontAtlasBuilder*>(Builder);
}

const ImFontLoader* ImFontAtlas_FontLoader(const ImFontAtlas* self) {
	return (const ImFontLoader*) self->FontLoader;
}

void ImFontAtlas_setFontLoader(ImFontAtlas* self, const ImFontLoader* FontLoader) {
	self->FontLoader = FontLoader;
}

const char* ImFontAtlas_FontLoaderName(const ImFontAtlas* self) {
	return (const char*) self->FontLoaderName;
}

void ImFontAtlas_setFontLoaderName(ImFontAtlas* self, const char* FontLoaderName) {
	self->FontLoaderName = FontLoaderName;
}

unsigned int ImFontAtlas_FontLoaderFlags(const ImFontAtlas* self) {
	return self->FontLoaderFlags;
}

void ImFontAtlas_setFontLoaderFlags(ImFontAtlas* self, unsigned int FontLoaderFlags) {
	self->FontLoaderFlags = static_cast<unsigned int>(FontLoaderFlags);
}

int ImFontAtlas_RefCount(const ImFontAtlas* self) {
	return self->RefCount;
}

void ImFontAtlas_setRefCount(ImFontAtlas* self, int RefCount) {
	self->RefCount = static_cast<int>(RefCount);
}

void* ImFontAtlas_OwnerContext(const ImFontAtlas* self) {
	return self->OwnerContext;
}

void ImFontAtlas_setOwnerContext(ImFontAtlas* self, void* OwnerContext) {
	self->OwnerContext = static_cast<ImGuiContext*>(OwnerContext);
}

ImFontAtlasRect* ImFontAtlas_TempRect(const ImFontAtlas* self) {
	return new ImFontAtlasRect(self->TempRect);
}

void ImFontAtlas_setTempRect(ImFontAtlas* self, ImFontAtlasRect* TempRect) {
	self->TempRect = *TempRect;
}

int ImFontAtlas_AddCustomRectRegular(ImFontAtlas* self, int w, int h) {
	int _ret = self->AddCustomRectRegular(static_cast<int>(w), static_cast<int>(h));
	return static_cast<int>(_ret);
}

ImFontAtlasRect* ImFontAtlas_GetCustomRectByIndex(ImFontAtlas* self, int id) {
	return (ImFontAtlasRect*) self->GetCustomRectByIndex(static_cast<int>(id));
}

void ImFontAtlas_CalcCustomRectUV(const ImFontAtlas* self, ImFontAtlasRect* r, ImVec2* out_uv_min, ImVec2* out_uv_max) {
	self->CalcCustomRectUV(r, out_uv_min, out_uv_max);
}

int ImFontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self, ImFont* font, unsigned short codepoint, int w, int h, float advance_x) {
	int _ret = self->AddCustomRectFontGlyph(font, static_cast<unsigned short>(codepoint), static_cast<int>(w), static_cast<int>(h), static_cast<float>(advance_x));
	return static_cast<int>(_ret);
}

int ImFontAtlas_AddCustomRectFontGlyphForSize(ImFontAtlas* self, ImFont* font, float font_size, unsigned short codepoint, int w, int h, float advance_x) {
	int _ret = self->AddCustomRectFontGlyphForSize(font, static_cast<float>(font_size), static_cast<unsigned short>(codepoint), static_cast<int>(w), static_cast<int>(h), static_cast<float>(advance_x));
	return static_cast<int>(_ret);
}

void ImFontAtlas_operatorAssign(ImFontAtlas* self, ImFontAtlas* param1) {
	self->operator=(*param1);
}

ImFont* ImFontAtlas_AddFontDefaultWithFontCfg(ImFontAtlas* self, ImFontConfig* font_cfg) {
	return self->AddFontDefault(font_cfg);
}

ImFont* ImFontAtlas_AddFontDefaultVectorWithFontCfg(ImFontAtlas* self, ImFontConfig* font_cfg) {
	return self->AddFontDefaultVector(font_cfg);
}

ImFont* ImFontAtlas_AddFontDefaultBitmapWithFontCfg(ImFontAtlas* self, ImFontConfig* font_cfg) {
	return self->AddFontDefaultBitmap(font_cfg);
}

ImFont* ImFontAtlas_AddFontFromFileTTF2(ImFontAtlas* self, const char* filename, float size_pixels) {
	return self->AddFontFromFileTTF(filename, static_cast<float>(size_pixels));
}

ImFont* ImFontAtlas_AddFontFromFileTTF3(ImFontAtlas* self, const char* filename, float size_pixels, ImFontConfig* font_cfg) {
	return self->AddFontFromFileTTF(filename, static_cast<float>(size_pixels), font_cfg);
}

ImFont* ImFontAtlas_AddFontFromFileTTF4(ImFontAtlas* self, const char* filename, float size_pixels, ImFontConfig* font_cfg, const unsigned short* glyph_ranges) {
	return self->AddFontFromFileTTF(filename, static_cast<float>(size_pixels), font_cfg, glyph_ranges);
}

ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF2(ImFontAtlas* self, const char* compressed_font_data_base85, float size_pixels) {
	return self->AddFontFromMemoryCompressedBase85TTF(compressed_font_data_base85, static_cast<float>(size_pixels));
}

ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF3(ImFontAtlas* self, const char* compressed_font_data_base85, float size_pixels, ImFontConfig* font_cfg) {
	return self->AddFontFromMemoryCompressedBase85TTF(compressed_font_data_base85, static_cast<float>(size_pixels), font_cfg);
}

ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF4(ImFontAtlas* self, const char* compressed_font_data_base85, float size_pixels, ImFontConfig* font_cfg, const unsigned short* glyph_ranges) {
	return self->AddFontFromMemoryCompressedBase85TTF(compressed_font_data_base85, static_cast<float>(size_pixels), font_cfg, glyph_ranges);
}

void ImFontAtlas_GetTexDataAsAlpha82(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height, int* out_bytes_per_pixel) {
	self->GetTexDataAsAlpha8(out_pixels, out_width, out_height, out_bytes_per_pixel);
}

void ImFontAtlas_GetTexDataAsRGBA322(ImFontAtlas* self, unsigned char** out_pixels, int* out_width, int* out_height, int* out_bytes_per_pixel) {
	self->GetTexDataAsRGBA32(out_pixels, out_width, out_height, out_bytes_per_pixel);
}

int ImFontAtlas_AddCustomRect2(ImFontAtlas* self, int width, int height, ImFontAtlasRect* out_r) {
	int _ret = self->AddCustomRect(static_cast<int>(width), static_cast<int>(height), out_r);
	return static_cast<int>(_ret);
}

int ImFontAtlas_AddCustomRectFontGlyph2(ImFontAtlas* self, ImFont* font, unsigned short codepoint, int w, int h, float advance_x, ImVec2* offset) {
	int _ret = self->AddCustomRectFontGlyph(font, static_cast<unsigned short>(codepoint), static_cast<int>(w), static_cast<int>(h), static_cast<float>(advance_x), *offset);
	return static_cast<int>(_ret);
}

int ImFontAtlas_AddCustomRectFontGlyphForSize2(ImFontAtlas* self, ImFont* font, float font_size, unsigned short codepoint, int w, int h, float advance_x, ImVec2* offset) {
	int _ret = self->AddCustomRectFontGlyphForSize(font, static_cast<float>(font_size), static_cast<unsigned short>(codepoint), static_cast<int>(w), static_cast<int>(h), static_cast<float>(advance_x), *offset);
	return static_cast<int>(_ret);
}

void ImFontAtlas_delete(ImFontAtlas* self) {
	delete self;
}

ImFontBaked* ImFontBaked_new() {
	return new (std::nothrow) ImFontBaked();
}

ImFontBaked* ImFontBaked_new2(ImFontBaked* param1) {
	return new (std::nothrow) ImFontBaked(*param1);
}

void* ImFontBaked_IndexAdvanceX(const ImFontBaked* self) {
	return const_cast<void*>(static_cast<const void*>(&self->IndexAdvanceX));
}

void ImFontBaked_setIndexAdvanceX(ImFontBaked* self, void* IndexAdvanceX) {
	self->IndexAdvanceX = *reinterpret_cast<ImVector<float>*>(IndexAdvanceX);
}

float ImFontBaked_FallbackAdvanceX(const ImFontBaked* self) {
	return self->FallbackAdvanceX;
}

void ImFontBaked_setFallbackAdvanceX(ImFontBaked* self, float FallbackAdvanceX) {
	self->FallbackAdvanceX = static_cast<float>(FallbackAdvanceX);
}

float ImFontBaked_Size(const ImFontBaked* self) {
	return self->Size;
}

void ImFontBaked_setSize(ImFontBaked* self, float Size) {
	self->Size = static_cast<float>(Size);
}

float ImFontBaked_RasterizerDensity(const ImFontBaked* self) {
	return self->RasterizerDensity;
}

void ImFontBaked_setRasterizerDensity(ImFontBaked* self, float RasterizerDensity) {
	self->RasterizerDensity = static_cast<float>(RasterizerDensity);
}

void* ImFontBaked_IndexLookup(const ImFontBaked* self) {
	return const_cast<void*>(static_cast<const void*>(&self->IndexLookup));
}

void ImFontBaked_setIndexLookup(ImFontBaked* self, void* IndexLookup) {
	self->IndexLookup = *reinterpret_cast<ImVector<unsigned short>*>(IndexLookup);
}

void* ImFontBaked_Glyphs(const ImFontBaked* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Glyphs));
}

void ImFontBaked_setGlyphs(ImFontBaked* self, void* Glyphs) {
	self->Glyphs = *reinterpret_cast<ImVector<ImFontGlyph>*>(Glyphs);
}

int ImFontBaked_FallbackGlyphIndex(const ImFontBaked* self) {
	return self->FallbackGlyphIndex;
}

void ImFontBaked_setFallbackGlyphIndex(ImFontBaked* self, int FallbackGlyphIndex) {
	self->FallbackGlyphIndex = static_cast<int>(FallbackGlyphIndex);
}

float ImFontBaked_Ascent(const ImFontBaked* self) {
	return self->Ascent;
}

void ImFontBaked_setAscent(ImFontBaked* self, float Ascent) {
	self->Ascent = static_cast<float>(Ascent);
}

float ImFontBaked_Descent(const ImFontBaked* self) {
	return self->Descent;
}

void ImFontBaked_setDescent(ImFontBaked* self, float Descent) {
	self->Descent = static_cast<float>(Descent);
}

unsigned int ImFontBaked_MetricsTotalSurface(const ImFontBaked* self) {
	return self->MetricsTotalSurface;
}

void ImFontBaked_setMetricsTotalSurface(ImFontBaked* self, unsigned int MetricsTotalSurface) {
	self->MetricsTotalSurface = static_cast<unsigned int>(MetricsTotalSurface);
}

unsigned int ImFontBaked_WantDestroy(const ImFontBaked* self) {
	return self->WantDestroy;
}

void ImFontBaked_setWantDestroy(ImFontBaked* self, unsigned int WantDestroy) {
	self->WantDestroy = static_cast<unsigned int>(WantDestroy);
}

unsigned int ImFontBaked_LoadNoFallback(const ImFontBaked* self) {
	return self->LoadNoFallback;
}

void ImFontBaked_setLoadNoFallback(ImFontBaked* self, unsigned int LoadNoFallback) {
	self->LoadNoFallback = static_cast<unsigned int>(LoadNoFallback);
}

unsigned int ImFontBaked_LoadNoRenderOnLayout(const ImFontBaked* self) {
	return self->LoadNoRenderOnLayout;
}

void ImFontBaked_setLoadNoRenderOnLayout(ImFontBaked* self, unsigned int LoadNoRenderOnLayout) {
	self->LoadNoRenderOnLayout = static_cast<unsigned int>(LoadNoRenderOnLayout);
}

int ImFontBaked_LastUsedFrame(const ImFontBaked* self) {
	return self->LastUsedFrame;
}

void ImFontBaked_setLastUsedFrame(ImFontBaked* self, int LastUsedFrame) {
	self->LastUsedFrame = static_cast<int>(LastUsedFrame);
}

unsigned int ImFontBaked_BakedId(const ImFontBaked* self) {
	return self->BakedId;
}

void ImFontBaked_setBakedId(ImFontBaked* self, unsigned int BakedId) {
	self->BakedId = static_cast<unsigned int>(BakedId);
}

ImFont* ImFontBaked_OwnerFont(const ImFontBaked* self) {
	return self->OwnerFont;
}

void ImFontBaked_setOwnerFont(ImFontBaked* self, ImFont* OwnerFont) {
	self->OwnerFont = OwnerFont;
}

void ImFontBaked_ClearOutputData(ImFontBaked* self) {
	self->ClearOutputData();
}

ImFontGlyph* ImFontBaked_FindGlyph(ImFontBaked* self, unsigned short c) {
	return self->FindGlyph(static_cast<unsigned short>(c));
}

ImFontGlyph* ImFontBaked_FindGlyphNoFallback(ImFontBaked* self, unsigned short c) {
	return self->FindGlyphNoFallback(static_cast<unsigned short>(c));
}

float ImFontBaked_GetCharAdvance(ImFontBaked* self, unsigned short c) {
	return self->GetCharAdvance(static_cast<unsigned short>(c));
}

bool ImFontBaked_IsGlyphLoaded(ImFontBaked* self, unsigned short c) {
	return self->IsGlyphLoaded(static_cast<unsigned short>(c));
}

void ImFontBaked_operatorAssign(ImFontBaked* self, ImFontBaked* param1) {
	self->operator=(*param1);
}

void ImFontBaked_delete(ImFontBaked* self) {
	delete self;
}

ImFont* ImFont_new() {
	return new (std::nothrow) ImFont();
}

ImFont* ImFont_new2(ImFont* param1) {
	return new (std::nothrow) ImFont(*param1);
}

ImFontBaked* ImFont_LastBaked(const ImFont* self) {
	return self->LastBaked;
}

void ImFont_setLastBaked(ImFont* self, ImFontBaked* LastBaked) {
	self->LastBaked = LastBaked;
}

ImFontAtlas* ImFont_OwnerAtlas(const ImFont* self) {
	return self->OwnerAtlas;
}

void ImFont_setOwnerAtlas(ImFont* self, ImFontAtlas* OwnerAtlas) {
	self->OwnerAtlas = OwnerAtlas;
}

int ImFont_Flags(const ImFont* self) {
	return self->Flags;
}

void ImFont_setFlags(ImFont* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

float ImFont_CurrentRasterizerDensity(const ImFont* self) {
	return self->CurrentRasterizerDensity;
}

void ImFont_setCurrentRasterizerDensity(ImFont* self, float CurrentRasterizerDensity) {
	self->CurrentRasterizerDensity = static_cast<float>(CurrentRasterizerDensity);
}

unsigned int ImFont_FontId(const ImFont* self) {
	return self->FontId;
}

void ImFont_setFontId(ImFont* self, unsigned int FontId) {
	self->FontId = static_cast<unsigned int>(FontId);
}

float ImFont_LegacySize(const ImFont* self) {
	return self->LegacySize;
}

void ImFont_setLegacySize(ImFont* self, float LegacySize) {
	self->LegacySize = static_cast<float>(LegacySize);
}

void* ImFont_Sources(const ImFont* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Sources));
}

void ImFont_setSources(ImFont* self, void* Sources) {
	self->Sources = *reinterpret_cast<ImVector<ImFontConfig *>*>(Sources);
}

unsigned short ImFont_EllipsisChar(const ImFont* self) {
	return self->EllipsisChar;
}

void ImFont_setEllipsisChar(ImFont* self, unsigned short EllipsisChar) {
	self->EllipsisChar = static_cast<unsigned short>(EllipsisChar);
}

unsigned short ImFont_FallbackChar(const ImFont* self) {
	return self->FallbackChar;
}

void ImFont_setFallbackChar(ImFont* self, unsigned short FallbackChar) {
	self->FallbackChar = static_cast<unsigned short>(FallbackChar);
}

ImU8* ImFont_Used8kPagesMap(const ImFont* self) {
	return (ImU8*)self->Used8kPagesMap;
}

void ImFont_setUsed8kPagesMap(ImFont* self, const ImU8* Used8kPagesMap) {
	memcpy(self->Used8kPagesMap, Used8kPagesMap, sizeof(self->Used8kPagesMap));
}

bool ImFont_EllipsisAutoBake(const ImFont* self) {
	return self->EllipsisAutoBake;
}

void ImFont_setEllipsisAutoBake(ImFont* self, bool EllipsisAutoBake) {
	self->EllipsisAutoBake = EllipsisAutoBake;
}

ImGuiStorage* ImFont_RemapPairs(const ImFont* self) {
	return new ImGuiStorage(self->RemapPairs);
}

void ImFont_setRemapPairs(ImFont* self, ImGuiStorage* RemapPairs) {
	self->RemapPairs = *RemapPairs;
}

float ImFont_Scale(const ImFont* self) {
	return self->Scale;
}

void ImFont_setScale(ImFont* self, float Scale) {
	self->Scale = static_cast<float>(Scale);
}

bool ImFont_IsGlyphInFont(ImFont* self, unsigned short c) {
	return self->IsGlyphInFont(static_cast<unsigned short>(c));
}

bool ImFont_IsLoaded(const ImFont* self) {
	return self->IsLoaded();
}

const char* ImFont_GetDebugName(const ImFont* self) {
	return (const char*) self->GetDebugName();
}

ImFontBaked* ImFont_GetFontBaked(ImFont* self, float font_size) {
	return self->GetFontBaked(static_cast<float>(font_size));
}

ImVec2* ImFont_CalcTextSizeA(ImFont* self, float size, float max_width, float wrap_width, const char* text_begin) {
	return new ImVec2(self->CalcTextSizeA(static_cast<float>(size), static_cast<float>(max_width), static_cast<float>(wrap_width), text_begin));
}

const char* ImFont_CalcWordWrapPosition(ImFont* self, float size, const char* text, const char* text_end, float wrap_width) {
	return (const char*) self->CalcWordWrapPosition(static_cast<float>(size), text, text_end, static_cast<float>(wrap_width));
}

void ImFont_RenderChar(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, unsigned short c) {
	self->RenderChar(draw_list, static_cast<float>(size), *pos, static_cast<unsigned int>(col), static_cast<unsigned short>(c));
}

void ImFont_RenderText(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, ImVec4* clip_rect, const char* text_begin, const char* text_end) {
	self->RenderText(draw_list, static_cast<float>(size), *pos, static_cast<unsigned int>(col), *clip_rect, text_begin, text_end);
}

const char* ImFont_CalcWordWrapPositionA(ImFont* self, float scale, const char* text, const char* text_end, float wrap_width) {
	return (const char*) self->CalcWordWrapPositionA(static_cast<float>(scale), text, text_end, static_cast<float>(wrap_width));
}

void ImFont_ClearOutputData(ImFont* self) {
	self->ClearOutputData();
}

void ImFont_AddRemapChar(ImFont* self, unsigned short from_codepoint, unsigned short to_codepoint) {
	self->AddRemapChar(static_cast<unsigned short>(from_codepoint), static_cast<unsigned short>(to_codepoint));
}

bool ImFont_IsGlyphRangeUnused(ImFont* self, unsigned int c_begin, unsigned int c_last) {
	return self->IsGlyphRangeUnused(static_cast<unsigned int>(c_begin), static_cast<unsigned int>(c_last));
}

void ImFont_operatorAssign(ImFont* self, ImFont* param1) {
	self->operator=(*param1);
}

ImFontBaked* ImFont_GetFontBaked2(ImFont* self, float font_size, float density) {
	return self->GetFontBaked(static_cast<float>(font_size), static_cast<float>(density));
}

ImVec2* ImFont_CalcTextSizeA2(ImFont* self, float size, float max_width, float wrap_width, const char* text_begin, const char* text_end) {
	return new ImVec2(self->CalcTextSizeA(static_cast<float>(size), static_cast<float>(max_width), static_cast<float>(wrap_width), text_begin, text_end));
}

ImVec2* ImFont_CalcTextSizeA3(ImFont* self, float size, float max_width, float wrap_width, const char* text_begin, const char* text_end, const char** out_remaining) {
	return new ImVec2(self->CalcTextSizeA(static_cast<float>(size), static_cast<float>(max_width), static_cast<float>(wrap_width), text_begin, text_end, out_remaining));
}

void ImFont_RenderChar2(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, unsigned short c, ImVec4* cpu_fine_clip) {
	self->RenderChar(draw_list, static_cast<float>(size), *pos, static_cast<unsigned int>(col), static_cast<unsigned short>(c), cpu_fine_clip);
}

void ImFont_RenderText2(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, ImVec4* clip_rect, const char* text_begin, const char* text_end, float wrap_width) {
	self->RenderText(draw_list, static_cast<float>(size), *pos, static_cast<unsigned int>(col), *clip_rect, text_begin, text_end, static_cast<float>(wrap_width));
}

void ImFont_RenderText3(ImFont* self, ImDrawList* draw_list, float size, ImVec2* pos, unsigned int col, ImVec4* clip_rect, const char* text_begin, const char* text_end, float wrap_width, int flags) {
	self->RenderText(draw_list, static_cast<float>(size), *pos, static_cast<unsigned int>(col), *clip_rect, text_begin, text_end, static_cast<float>(wrap_width), static_cast<int>(flags));
}

void ImFont_delete(ImFont* self) {
	delete self;
}

ImGuiViewport* ImGuiViewport_new() {
	return new (std::nothrow) ImGuiViewport();
}

unsigned int ImGuiViewport_ID(const ImGuiViewport* self) {
	return self->ID;
}

void ImGuiViewport_setID(ImGuiViewport* self, unsigned int ID) {
	self->ID = static_cast<unsigned int>(ID);
}

int ImGuiViewport_Flags(const ImGuiViewport* self) {
	return self->Flags;
}

void ImGuiViewport_setFlags(ImGuiViewport* self, int Flags) {
	self->Flags = static_cast<int>(Flags);
}

ImVec2* ImGuiViewport_Pos(const ImGuiViewport* self) {
	return new ImVec2(self->Pos);
}

void ImGuiViewport_setPos(ImGuiViewport* self, ImVec2* Pos) {
	self->Pos = *Pos;
}

ImVec2* ImGuiViewport_Size(const ImGuiViewport* self) {
	return new ImVec2(self->Size);
}

void ImGuiViewport_setSize(ImGuiViewport* self, ImVec2* Size) {
	self->Size = *Size;
}

ImVec2* ImGuiViewport_FramebufferScale(const ImGuiViewport* self) {
	return new ImVec2(self->FramebufferScale);
}

void ImGuiViewport_setFramebufferScale(ImGuiViewport* self, ImVec2* FramebufferScale) {
	self->FramebufferScale = *FramebufferScale;
}

ImVec2* ImGuiViewport_WorkPos(const ImGuiViewport* self) {
	return new ImVec2(self->WorkPos);
}

void ImGuiViewport_setWorkPos(ImGuiViewport* self, ImVec2* WorkPos) {
	self->WorkPos = *WorkPos;
}

ImVec2* ImGuiViewport_WorkSize(const ImGuiViewport* self) {
	return new ImVec2(self->WorkSize);
}

void ImGuiViewport_setWorkSize(ImGuiViewport* self, ImVec2* WorkSize) {
	self->WorkSize = *WorkSize;
}

ImVec2* ImGuiViewport_GetCenter(const ImGuiViewport* self) {
	return new ImVec2(self->GetCenter());
}

ImVec2* ImGuiViewport_GetWorkCenter(const ImGuiViewport* self) {
	return new ImVec2(self->GetWorkCenter());
}

void ImGuiViewport_delete(ImGuiViewport* self) {
	delete self;
}

ImGuiPlatformIO* ImGuiPlatformIO_new() {
	return new (std::nothrow) ImGuiPlatformIO();
}

ImGuiPlatformIO* ImGuiPlatformIO_new2(ImGuiPlatformIO* param1) {
	return new (std::nothrow) ImGuiPlatformIO(*param1);
}

unsigned short ImGuiPlatformIO_Platform_LocaleDecimalPoint(const ImGuiPlatformIO* self) {
	return self->Platform_LocaleDecimalPoint;
}

void ImGuiPlatformIO_setPlatform_LocaleDecimalPoint(ImGuiPlatformIO* self, unsigned short Platform_LocaleDecimalPoint) {
	self->Platform_LocaleDecimalPoint = static_cast<unsigned short>(Platform_LocaleDecimalPoint);
}

int ImGuiPlatformIO_Renderer_TextureMaxWidth(const ImGuiPlatformIO* self) {
	return self->Renderer_TextureMaxWidth;
}

void ImGuiPlatformIO_setRenderer_TextureMaxWidth(ImGuiPlatformIO* self, int Renderer_TextureMaxWidth) {
	self->Renderer_TextureMaxWidth = static_cast<int>(Renderer_TextureMaxWidth);
}

int ImGuiPlatformIO_Renderer_TextureMaxHeight(const ImGuiPlatformIO* self) {
	return self->Renderer_TextureMaxHeight;
}

void ImGuiPlatformIO_setRenderer_TextureMaxHeight(ImGuiPlatformIO* self, int Renderer_TextureMaxHeight) {
	self->Renderer_TextureMaxHeight = static_cast<int>(Renderer_TextureMaxHeight);
}

void* ImGuiPlatformIO_Textures(const ImGuiPlatformIO* self) {
	return const_cast<void*>(static_cast<const void*>(&self->Textures));
}

void ImGuiPlatformIO_setTextures(ImGuiPlatformIO* self, void* Textures) {
	self->Textures = *reinterpret_cast<ImVector<ImTextureData *>*>(Textures);
}

void ImGuiPlatformIO_ClearPlatformHandlers(ImGuiPlatformIO* self) {
	self->ClearPlatformHandlers();
}

void ImGuiPlatformIO_ClearRendererHandlers(ImGuiPlatformIO* self) {
	self->ClearRendererHandlers();
}

void ImGuiPlatformIO_operatorAssign(ImGuiPlatformIO* self, ImGuiPlatformIO* param1) {
	self->operator=(*param1);
}

void ImGuiPlatformIO_delete(ImGuiPlatformIO* self) {
	delete self;
}

ImGuiPlatformImeData* ImGuiPlatformImeData_new() {
	return new (std::nothrow) ImGuiPlatformImeData();
}

bool ImGuiPlatformImeData_WantVisible(const ImGuiPlatformImeData* self) {
	return self->WantVisible;
}

void ImGuiPlatformImeData_setWantVisible(ImGuiPlatformImeData* self, bool WantVisible) {
	self->WantVisible = WantVisible;
}

bool ImGuiPlatformImeData_WantTextInput(const ImGuiPlatformImeData* self) {
	return self->WantTextInput;
}

void ImGuiPlatformImeData_setWantTextInput(ImGuiPlatformImeData* self, bool WantTextInput) {
	self->WantTextInput = WantTextInput;
}

ImVec2* ImGuiPlatformImeData_InputPos(const ImGuiPlatformImeData* self) {
	return new ImVec2(self->InputPos);
}

void ImGuiPlatformImeData_setInputPos(ImGuiPlatformImeData* self, ImVec2* InputPos) {
	self->InputPos = *InputPos;
}

float ImGuiPlatformImeData_InputLineHeight(const ImGuiPlatformImeData* self) {
	return self->InputLineHeight;
}

void ImGuiPlatformImeData_setInputLineHeight(ImGuiPlatformImeData* self, float InputLineHeight) {
	self->InputLineHeight = static_cast<float>(InputLineHeight);
}

unsigned int ImGuiPlatformImeData_ViewportId(const ImGuiPlatformImeData* self) {
	return self->ViewportId;
}

void ImGuiPlatformImeData_setViewportId(ImGuiPlatformImeData* self, unsigned int ViewportId) {
	self->ViewportId = static_cast<unsigned int>(ViewportId);
}

void ImGuiPlatformImeData_delete(ImGuiPlatformImeData* self) {
	delete self;
}

void* ImGui_ImplDX11_RenderState_Device(const ImGui_ImplDX11_RenderState* self) {
	return self->Device;
}

void ImGui_ImplDX11_RenderState_setDevice(ImGui_ImplDX11_RenderState* self, void* Device) {
	self->Device = static_cast<ID3D11Device*>(Device);
}

void* ImGui_ImplDX11_RenderState_DeviceContext(const ImGui_ImplDX11_RenderState* self) {
	return self->DeviceContext;
}

void ImGui_ImplDX11_RenderState_setDeviceContext(ImGui_ImplDX11_RenderState* self, void* DeviceContext) {
	self->DeviceContext = static_cast<ID3D11DeviceContext*>(DeviceContext);
}

void* ImGui_ImplDX11_RenderState_VertexConstantBuffer(const ImGui_ImplDX11_RenderState* self) {
	return self->VertexConstantBuffer;
}

void ImGui_ImplDX11_RenderState_setVertexConstantBuffer(ImGui_ImplDX11_RenderState* self, void* VertexConstantBuffer) {
	self->VertexConstantBuffer = static_cast<ID3D11Buffer*>(VertexConstantBuffer);
}

void ImGui_ImplDX11_RenderState_delete(ImGui_ImplDX11_RenderState* self) {
	delete self;
}

void* cabi_ImGui__CreateContext(ImFontAtlas* shared_font_atlas) {
	return ImGui::CreateContext(shared_font_atlas);
}

void cabi_ImGui__DestroyContext(void* ctx) {
	ImGui::DestroyContext(static_cast<ImGuiContext*>(ctx));
}

void* cabi_ImGui__GetCurrentContext() {
	return ImGui::GetCurrentContext();
}

void cabi_ImGui__SetCurrentContext(void* ctx) {
	ImGui::SetCurrentContext(static_cast<ImGuiContext*>(ctx));
}

ImGuiIO* cabi_ImGui__GetIO() {
	ImGuiIO& _ret = ImGui::GetIO();
	// Cast returned reference into pointer
	return &_ret;
}

ImGuiPlatformIO* cabi_ImGui__GetPlatformIO() {
	ImGuiPlatformIO& _ret = ImGui::GetPlatformIO();
	// Cast returned reference into pointer
	return &_ret;
}

ImGuiStyle* cabi_ImGui__GetStyle() {
	ImGuiStyle& _ret = ImGui::GetStyle();
	// Cast returned reference into pointer
	return &_ret;
}

void cabi_ImGui__NewFrame() {
	ImGui::NewFrame();
}

void cabi_ImGui__EndFrame() {
	ImGui::EndFrame();
}

void cabi_ImGui__Render() {
	ImGui::Render();
}

ImDrawData* cabi_ImGui__GetDrawData() {
	return ImGui::GetDrawData();
}

void cabi_ImGui__ShowDemoWindow(bool* p_open) {
	ImGui::ShowDemoWindow(p_open);
}

void cabi_ImGui__ShowMetricsWindow(bool* p_open) {
	ImGui::ShowMetricsWindow(p_open);
}

void cabi_ImGui__ShowDebugLogWindow(bool* p_open) {
	ImGui::ShowDebugLogWindow(p_open);
}

void cabi_ImGui__ShowIDStackToolWindow(bool* p_open) {
	ImGui::ShowIDStackToolWindow(p_open);
}

void cabi_ImGui__ShowAboutWindow(bool* p_open) {
	ImGui::ShowAboutWindow(p_open);
}

void cabi_ImGui__ShowStyleEditor(ImGuiStyle* ref) {
	ImGui::ShowStyleEditor(ref);
}

bool cabi_ImGui__ShowStyleSelector(const char* label) {
	return ImGui::ShowStyleSelector(label);
}

void cabi_ImGui__ShowFontSelector(const char* label) {
	ImGui::ShowFontSelector(label);
}

void cabi_ImGui__ShowUserGuide() {
	ImGui::ShowUserGuide();
}

const char* cabi_ImGui__GetVersion() {
	return (const char*) ImGui::GetVersion();
}

void cabi_ImGui__StyleColorsDark(ImGuiStyle* dst) {
	ImGui::StyleColorsDark(dst);
}

void cabi_ImGui__StyleColorsLight(ImGuiStyle* dst) {
	ImGui::StyleColorsLight(dst);
}

void cabi_ImGui__StyleColorsClassic(ImGuiStyle* dst) {
	ImGui::StyleColorsClassic(dst);
}

bool cabi_ImGui__Begin(const char* name, bool* p_open, int flags) {
	return ImGui::Begin(name, p_open, static_cast<ImGuiWindowFlags>(flags));
}

void cabi_ImGui__End() {
	ImGui::End();
}

bool cabi_ImGui__BeginChild_1(const char* str_id, ImVec2* size, int child_flags, int window_flags) {
	return ImGui::BeginChild(str_id, *size, static_cast<ImGuiChildFlags>(child_flags), static_cast<ImGuiWindowFlags>(window_flags));
}

bool cabi_ImGui__BeginChild_2(int id, ImVec2* size, int child_flags, int window_flags) {
	return ImGui::BeginChild(static_cast<ImGuiID>(id), *size, static_cast<ImGuiChildFlags>(child_flags), static_cast<ImGuiWindowFlags>(window_flags));
}

void cabi_ImGui__EndChild() {
	ImGui::EndChild();
}

bool cabi_ImGui__IsWindowAppearing() {
	return ImGui::IsWindowAppearing();
}

bool cabi_ImGui__IsWindowCollapsed() {
	return ImGui::IsWindowCollapsed();
}

bool cabi_ImGui__IsWindowFocused(int flags) {
	return ImGui::IsWindowFocused(static_cast<ImGuiFocusedFlags>(flags));
}

bool cabi_ImGui__IsWindowHovered(int flags) {
	return ImGui::IsWindowHovered(static_cast<ImGuiHoveredFlags>(flags));
}

ImDrawList* cabi_ImGui__GetWindowDrawList() {
	return ImGui::GetWindowDrawList();
}

ImVec2* cabi_ImGui__GetWindowPos() {
	return new ImVec2(ImGui::GetWindowPos());
}

ImVec2* cabi_ImGui__GetWindowSize() {
	return new ImVec2(ImGui::GetWindowSize());
}

float cabi_ImGui__GetWindowWidth() {
	return ImGui::GetWindowWidth();
}

float cabi_ImGui__GetWindowHeight() {
	return ImGui::GetWindowHeight();
}

void cabi_ImGui__SetNextWindowPos(ImVec2* pos, int cond, ImVec2* pivot) {
	ImGui::SetNextWindowPos(*pos, static_cast<ImGuiCond>(cond), *pivot);
}

void cabi_ImGui__SetNextWindowSize(ImVec2* size, int cond) {
	ImGui::SetNextWindowSize(*size, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetNextWindowSizeConstraints(ImVec2* size_min, ImVec2* size_max, void* custom_callback, void* custom_callback_data) {
	ImGui::SetNextWindowSizeConstraints(*size_min, *size_max, (ImGuiSizeCallback)(custom_callback), static_cast<void*>(custom_callback_data));
}

void cabi_ImGui__SetNextWindowContentSize(ImVec2* size) {
	ImGui::SetNextWindowContentSize(*size);
}

void cabi_ImGui__SetNextWindowCollapsed(bool collapsed, int cond) {
	ImGui::SetNextWindowCollapsed(collapsed, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetNextWindowFocus() {
	ImGui::SetNextWindowFocus();
}

void cabi_ImGui__SetNextWindowScroll(ImVec2* scroll) {
	ImGui::SetNextWindowScroll(*scroll);
}

void cabi_ImGui__SetNextWindowBgAlpha(float alpha) {
	ImGui::SetNextWindowBgAlpha(static_cast<float>(alpha));
}

void cabi_ImGui__SetWindowPos_1(ImVec2* pos, int cond) {
	ImGui::SetWindowPos(*pos, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetWindowSize_1(ImVec2* size, int cond) {
	ImGui::SetWindowSize(*size, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetWindowCollapsed_1(bool collapsed, int cond) {
	ImGui::SetWindowCollapsed(collapsed, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetWindowFocus_1() {
	ImGui::SetWindowFocus();
}

void cabi_ImGui__SetWindowPos_2(const char* name, ImVec2* pos, int cond) {
	ImGui::SetWindowPos(name, *pos, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetWindowSize_2(const char* name, ImVec2* size, int cond) {
	ImGui::SetWindowSize(name, *size, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetWindowCollapsed_2(const char* name, bool collapsed, int cond) {
	ImGui::SetWindowCollapsed(name, collapsed, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetWindowFocus_2(const char* name) {
	ImGui::SetWindowFocus(name);
}

float cabi_ImGui__GetScrollX() {
	return ImGui::GetScrollX();
}

float cabi_ImGui__GetScrollY() {
	return ImGui::GetScrollY();
}

void cabi_ImGui__SetScrollX(float scroll_x) {
	ImGui::SetScrollX(static_cast<float>(scroll_x));
}

void cabi_ImGui__SetScrollY(float scroll_y) {
	ImGui::SetScrollY(static_cast<float>(scroll_y));
}

float cabi_ImGui__GetScrollMaxX() {
	return ImGui::GetScrollMaxX();
}

float cabi_ImGui__GetScrollMaxY() {
	return ImGui::GetScrollMaxY();
}

void cabi_ImGui__SetScrollHereX(float center_x_ratio) {
	ImGui::SetScrollHereX(static_cast<float>(center_x_ratio));
}

void cabi_ImGui__SetScrollHereY(float center_y_ratio) {
	ImGui::SetScrollHereY(static_cast<float>(center_y_ratio));
}

void cabi_ImGui__SetScrollFromPosX(float local_x, float center_x_ratio) {
	ImGui::SetScrollFromPosX(static_cast<float>(local_x), static_cast<float>(center_x_ratio));
}

void cabi_ImGui__SetScrollFromPosY(float local_y, float center_y_ratio) {
	ImGui::SetScrollFromPosY(static_cast<float>(local_y), static_cast<float>(center_y_ratio));
}

void cabi_ImGui__PushFont_1(ImFont* font, float font_size_base_unscaled) {
	ImGui::PushFont(font, static_cast<float>(font_size_base_unscaled));
}

void cabi_ImGui__PopFont() {
	ImGui::PopFont();
}

ImFont* cabi_ImGui__GetFont() {
	return ImGui::GetFont();
}

float cabi_ImGui__GetFontSize() {
	return ImGui::GetFontSize();
}

ImFontBaked* cabi_ImGui__GetFontBaked() {
	return ImGui::GetFontBaked();
}

void cabi_ImGui__PushStyleColor_1(int idx, int col) {
	ImGui::PushStyleColor(static_cast<ImGuiCol>(idx), static_cast<ImU32>(col));
}

void cabi_ImGui__PushStyleColor_2(int idx, ImVec4* col) {
	ImGui::PushStyleColor(static_cast<ImGuiCol>(idx), *col);
}

void cabi_ImGui__PopStyleColor(int count) {
	ImGui::PopStyleColor(static_cast<int>(count));
}

void cabi_ImGui__PushStyleVar_1(int idx, float val) {
	ImGui::PushStyleVar(static_cast<ImGuiStyleVar>(idx), static_cast<float>(val));
}

void cabi_ImGui__PushStyleVar_2(int idx, ImVec2* val) {
	ImGui::PushStyleVar(static_cast<ImGuiStyleVar>(idx), *val);
}

void cabi_ImGui__PushStyleVarX(int idx, float val_x) {
	ImGui::PushStyleVarX(static_cast<ImGuiStyleVar>(idx), static_cast<float>(val_x));
}

void cabi_ImGui__PushStyleVarY(int idx, float val_y) {
	ImGui::PushStyleVarY(static_cast<ImGuiStyleVar>(idx), static_cast<float>(val_y));
}

void cabi_ImGui__PopStyleVar(int count) {
	ImGui::PopStyleVar(static_cast<int>(count));
}

void cabi_ImGui__PushItemFlag(int option, bool enabled) {
	ImGui::PushItemFlag(static_cast<ImGuiItemFlags>(option), enabled);
}

void cabi_ImGui__PopItemFlag() {
	ImGui::PopItemFlag();
}

void cabi_ImGui__PushItemWidth(float item_width) {
	ImGui::PushItemWidth(static_cast<float>(item_width));
}

void cabi_ImGui__PopItemWidth() {
	ImGui::PopItemWidth();
}

void cabi_ImGui__SetNextItemWidth(float item_width) {
	ImGui::SetNextItemWidth(static_cast<float>(item_width));
}

float cabi_ImGui__CalcItemWidth() {
	return ImGui::CalcItemWidth();
}

void cabi_ImGui__PushTextWrapPos(float wrap_local_pos_x) {
	ImGui::PushTextWrapPos(static_cast<float>(wrap_local_pos_x));
}

void cabi_ImGui__PopTextWrapPos() {
	ImGui::PopTextWrapPos();
}

ImVec2* cabi_ImGui__GetFontTexUvWhitePixel() {
	return new ImVec2(ImGui::GetFontTexUvWhitePixel());
}

int cabi_ImGui__GetColorU32_1(int idx, float alpha_mul) {
	return ImGui::GetColorU32(static_cast<ImGuiCol>(idx), static_cast<float>(alpha_mul));
}

int cabi_ImGui__GetColorU32_2(ImVec4* col) {
	return ImGui::GetColorU32(*col);
}

int cabi_ImGui__GetColorU32_3(int col, float alpha_mul) {
	return ImGui::GetColorU32(static_cast<ImU32>(col), static_cast<float>(alpha_mul));
}

ImVec4* cabi_ImGui__GetStyleColorVec4(int idx) {
	const ImVec4& _ret = ImGui::GetStyleColorVec4(static_cast<ImGuiCol>(idx));
	// Cast returned reference into pointer
	return const_cast<ImVec4*>(&_ret);
}

ImVec2* cabi_ImGui__GetCursorScreenPos() {
	return new ImVec2(ImGui::GetCursorScreenPos());
}

void cabi_ImGui__SetCursorScreenPos(ImVec2* pos) {
	ImGui::SetCursorScreenPos(*pos);
}

ImVec2* cabi_ImGui__GetContentRegionAvail() {
	return new ImVec2(ImGui::GetContentRegionAvail());
}

ImVec2* cabi_ImGui__GetCursorPos() {
	return new ImVec2(ImGui::GetCursorPos());
}

float cabi_ImGui__GetCursorPosX() {
	return ImGui::GetCursorPosX();
}

float cabi_ImGui__GetCursorPosY() {
	return ImGui::GetCursorPosY();
}

void cabi_ImGui__SetCursorPos(ImVec2* local_pos) {
	ImGui::SetCursorPos(*local_pos);
}

void cabi_ImGui__SetCursorPosX(float local_x) {
	ImGui::SetCursorPosX(static_cast<float>(local_x));
}

void cabi_ImGui__SetCursorPosY(float local_y) {
	ImGui::SetCursorPosY(static_cast<float>(local_y));
}

ImVec2* cabi_ImGui__GetCursorStartPos() {
	return new ImVec2(ImGui::GetCursorStartPos());
}

void cabi_ImGui__Separator() {
	ImGui::Separator();
}

void cabi_ImGui__SameLine(float offset_from_start_x, float spacing) {
	ImGui::SameLine(static_cast<float>(offset_from_start_x), static_cast<float>(spacing));
}

void cabi_ImGui__NewLine() {
	ImGui::NewLine();
}

void cabi_ImGui__Spacing() {
	ImGui::Spacing();
}

void cabi_ImGui__Dummy(ImVec2* size) {
	ImGui::Dummy(*size);
}

void cabi_ImGui__Indent(float indent_w) {
	ImGui::Indent(static_cast<float>(indent_w));
}

void cabi_ImGui__Unindent(float indent_w) {
	ImGui::Unindent(static_cast<float>(indent_w));
}

void cabi_ImGui__BeginGroup() {
	ImGui::BeginGroup();
}

void cabi_ImGui__EndGroup() {
	ImGui::EndGroup();
}

void cabi_ImGui__AlignTextToFramePadding() {
	ImGui::AlignTextToFramePadding();
}

float cabi_ImGui__GetTextLineHeight() {
	return ImGui::GetTextLineHeight();
}

float cabi_ImGui__GetTextLineHeightWithSpacing() {
	return ImGui::GetTextLineHeightWithSpacing();
}

float cabi_ImGui__GetFrameHeight() {
	return ImGui::GetFrameHeight();
}

float cabi_ImGui__GetFrameHeightWithSpacing() {
	return ImGui::GetFrameHeightWithSpacing();
}

void cabi_ImGui__PushID_1(const char* str_id) {
	ImGui::PushID(str_id);
}

void cabi_ImGui__PushID_2(const char* str_id_begin, const char* str_id_end) {
	ImGui::PushID(str_id_begin, str_id_end);
}

void cabi_ImGui__PushID_3(const void* ptr_id) {
	ImGui::PushID(ptr_id);
}

void cabi_ImGui__PushID_4(int int_id) {
	ImGui::PushID(static_cast<int>(int_id));
}

void cabi_ImGui__PopID() {
	ImGui::PopID();
}

int cabi_ImGui__GetID_1(const char* str_id) {
	return ImGui::GetID(str_id);
}

int cabi_ImGui__GetID_2(const char* str_id_begin, const char* str_id_end) {
	return ImGui::GetID(str_id_begin, str_id_end);
}

int cabi_ImGui__GetID_3(const void* ptr_id) {
	return ImGui::GetID(ptr_id);
}

int cabi_ImGui__GetID_4(int int_id) {
	return ImGui::GetID(static_cast<int>(int_id));
}

void cabi_ImGui__TextUnformatted(const char* text, const char* text_end) {
	ImGui::TextUnformatted(text, text_end);
}

void cabi_ImGui__TextV(const char* fmt, void* args) {
	ImGui::TextV(fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__TextColoredV(ImVec4* col, const char* fmt, void* args) {
	ImGui::TextColoredV(*col, fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__TextDisabledV(const char* fmt, void* args) {
	ImGui::TextDisabledV(fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__TextWrappedV(const char* fmt, void* args) {
	ImGui::TextWrappedV(fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__LabelTextV(const char* label, const char* fmt, void* args) {
	ImGui::LabelTextV(label, fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__BulletTextV(const char* fmt, void* args) {
	ImGui::BulletTextV(fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__SeparatorText(const char* label) {
	ImGui::SeparatorText(label);
}

bool cabi_ImGui__Button(const char* label, ImVec2* size) {
	return ImGui::Button(label, *size);
}

bool cabi_ImGui__SmallButton(const char* label) {
	return ImGui::SmallButton(label);
}

bool cabi_ImGui__InvisibleButton(const char* str_id, ImVec2* size, int flags) {
	return ImGui::InvisibleButton(str_id, *size, static_cast<ImGuiButtonFlags>(flags));
}

bool cabi_ImGui__ArrowButton(const char* str_id, int dir) {
	return ImGui::ArrowButton(str_id, static_cast<ImGuiDir>(dir));
}

bool cabi_ImGui__Checkbox(const char* label, bool* v) {
	return ImGui::Checkbox(label, v);
}

bool cabi_ImGui__CheckboxFlags_1(const char* label, int* flags, int flags_value) {
	return ImGui::CheckboxFlags(label, flags, static_cast<int>(flags_value));
}

bool cabi_ImGui__CheckboxFlags_2(const char* label, unsigned int* flags, unsigned int flags_value) {
	return ImGui::CheckboxFlags(label, flags, static_cast<unsigned int>(flags_value));
}

bool cabi_ImGui__RadioButton_1(const char* label, bool active) {
	return ImGui::RadioButton(label, active);
}

bool cabi_ImGui__RadioButton_2(const char* label, int* v, int v_button) {
	return ImGui::RadioButton(label, v, static_cast<int>(v_button));
}

void cabi_ImGui__ProgressBar(float fraction, ImVec2* size_arg, const char* overlay) {
	ImGui::ProgressBar(static_cast<float>(fraction), *size_arg, overlay);
}

void cabi_ImGui__Bullet() {
	ImGui::Bullet();
}

bool cabi_ImGui__TextLink(const char* label) {
	return ImGui::TextLink(label);
}

bool cabi_ImGui__TextLinkOpenURL(const char* label, const char* url) {
	return ImGui::TextLinkOpenURL(label, url);
}

void cabi_ImGui__Image_1(ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1) {
	ImGui::Image(*tex_ref, *image_size, *uv0, *uv1);
}

void cabi_ImGui__ImageWithBg(ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1, ImVec4* bg_col, ImVec4* tint_col) {
	ImGui::ImageWithBg(*tex_ref, *image_size, *uv0, *uv1, *bg_col, *tint_col);
}

bool cabi_ImGui__ImageButton(const char* str_id, ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1, ImVec4* bg_col, ImVec4* tint_col) {
	return ImGui::ImageButton(str_id, *tex_ref, *image_size, *uv0, *uv1, *bg_col, *tint_col);
}

bool cabi_ImGui__BeginCombo(const char* label, const char* preview_value, int flags) {
	return ImGui::BeginCombo(label, preview_value, static_cast<ImGuiComboFlags>(flags));
}

void cabi_ImGui__EndCombo() {
	ImGui::EndCombo();
}

bool cabi_ImGui__Combo_1(const char* label, int* current_item, const char** items, int items_count, int popup_max_height_in_items) {
	return ImGui::Combo(label, current_item, items, static_cast<int>(items_count), static_cast<int>(popup_max_height_in_items));
}

bool cabi_ImGui__Combo_2(const char* label, int* current_item, const char* items_separated_by_zeros, int popup_max_height_in_items) {
	return ImGui::Combo(label, current_item, items_separated_by_zeros, static_cast<int>(popup_max_height_in_items));
}

bool cabi_ImGui__DragFloat(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags) {
	return ImGui::DragFloat(label, v, static_cast<float>(v_speed), static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragFloat2(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags) {
	return ImGui::DragFloat2(label, v, static_cast<float>(v_speed), static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragFloat3(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags) {
	return ImGui::DragFloat3(label, v, static_cast<float>(v_speed), static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragFloat4(const char* label, float* v, float v_speed, float v_min, float v_max, const char* format, int flags) {
	return ImGui::DragFloat4(label, v, static_cast<float>(v_speed), static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragFloatRange2(const char* label, float* v_current_min, float* v_current_max, float v_speed, float v_min, float v_max, const char* format, const char* format_max, int flags) {
	return ImGui::DragFloatRange2(label, v_current_min, v_current_max, static_cast<float>(v_speed), static_cast<float>(v_min), static_cast<float>(v_max), format, format_max, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragInt(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags) {
	return ImGui::DragInt(label, v, static_cast<float>(v_speed), static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragInt2(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags) {
	return ImGui::DragInt2(label, v, static_cast<float>(v_speed), static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragInt3(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags) {
	return ImGui::DragInt3(label, v, static_cast<float>(v_speed), static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragInt4(const char* label, int* v, float v_speed, int v_min, int v_max, const char* format, int flags) {
	return ImGui::DragInt4(label, v, static_cast<float>(v_speed), static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragIntRange2(const char* label, int* v_current_min, int* v_current_max, float v_speed, int v_min, int v_max, const char* format, const char* format_max, int flags) {
	return ImGui::DragIntRange2(label, v_current_min, v_current_max, static_cast<float>(v_speed), static_cast<int>(v_min), static_cast<int>(v_max), format, format_max, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragScalar(const char* label, int data_type, void* p_data, float v_speed, const void* p_min, const void* p_max, const char* format, int flags) {
	return ImGui::DragScalar(label, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), static_cast<float>(v_speed), p_min, p_max, format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__DragScalarN(const char* label, int data_type, void* p_data, int components, float v_speed, const void* p_min, const void* p_max, const char* format, int flags) {
	return ImGui::DragScalarN(label, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), static_cast<int>(components), static_cast<float>(v_speed), p_min, p_max, format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderFloat(const char* label, float* v, float v_min, float v_max, const char* format, int flags) {
	return ImGui::SliderFloat(label, v, static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderFloat2(const char* label, float* v, float v_min, float v_max, const char* format, int flags) {
	return ImGui::SliderFloat2(label, v, static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderFloat3(const char* label, float* v, float v_min, float v_max, const char* format, int flags) {
	return ImGui::SliderFloat3(label, v, static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderFloat4(const char* label, float* v, float v_min, float v_max, const char* format, int flags) {
	return ImGui::SliderFloat4(label, v, static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderAngle(const char* label, float* v_rad, float v_degrees_min, float v_degrees_max, const char* format, int flags) {
	return ImGui::SliderAngle(label, v_rad, static_cast<float>(v_degrees_min), static_cast<float>(v_degrees_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderInt(const char* label, int* v, int v_min, int v_max, const char* format, int flags) {
	return ImGui::SliderInt(label, v, static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderInt2(const char* label, int* v, int v_min, int v_max, const char* format, int flags) {
	return ImGui::SliderInt2(label, v, static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderInt3(const char* label, int* v, int v_min, int v_max, const char* format, int flags) {
	return ImGui::SliderInt3(label, v, static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderInt4(const char* label, int* v, int v_min, int v_max, const char* format, int flags) {
	return ImGui::SliderInt4(label, v, static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderScalar(const char* label, int data_type, void* p_data, const void* p_min, const void* p_max, const char* format, int flags) {
	return ImGui::SliderScalar(label, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), p_min, p_max, format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__SliderScalarN(const char* label, int data_type, void* p_data, int components, const void* p_min, const void* p_max, const char* format, int flags) {
	return ImGui::SliderScalarN(label, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), static_cast<int>(components), p_min, p_max, format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__VSliderFloat(const char* label, ImVec2* size, float* v, float v_min, float v_max, const char* format, int flags) {
	return ImGui::VSliderFloat(label, *size, v, static_cast<float>(v_min), static_cast<float>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__VSliderInt(const char* label, ImVec2* size, int* v, int v_min, int v_max, const char* format, int flags) {
	return ImGui::VSliderInt(label, *size, v, static_cast<int>(v_min), static_cast<int>(v_max), format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__VSliderScalar(const char* label, ImVec2* size, int data_type, void* p_data, const void* p_min, const void* p_max, const char* format, int flags) {
	return ImGui::VSliderScalar(label, *size, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), p_min, p_max, format, static_cast<ImGuiSliderFlags>(flags));
}

bool cabi_ImGui__InputText(const char* label, char* buf, size_t buf_size, int flags, void* callback, void* user_data) {
	return ImGui::InputText(label, buf, static_cast<size_t>(buf_size), static_cast<ImGuiInputTextFlags>(flags), (ImGuiInputTextCallback)(callback), static_cast<void*>(user_data));
}

bool cabi_ImGui__InputTextMultiline(const char* label, char* buf, size_t buf_size, ImVec2* size, int flags, void* callback, void* user_data) {
	return ImGui::InputTextMultiline(label, buf, static_cast<size_t>(buf_size), *size, static_cast<ImGuiInputTextFlags>(flags), (ImGuiInputTextCallback)(callback), static_cast<void*>(user_data));
}

bool cabi_ImGui__InputTextWithHint(const char* label, const char* hint, char* buf, size_t buf_size, int flags, void* callback, void* user_data) {
	return ImGui::InputTextWithHint(label, hint, buf, static_cast<size_t>(buf_size), static_cast<ImGuiInputTextFlags>(flags), (ImGuiInputTextCallback)(callback), static_cast<void*>(user_data));
}

bool cabi_ImGui__InputFloat(const char* label, float* v, float step, float step_fast, const char* format, int flags) {
	return ImGui::InputFloat(label, v, static_cast<float>(step), static_cast<float>(step_fast), format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputFloat2(const char* label, float* v, const char* format, int flags) {
	return ImGui::InputFloat2(label, v, format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputFloat3(const char* label, float* v, const char* format, int flags) {
	return ImGui::InputFloat3(label, v, format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputFloat4(const char* label, float* v, const char* format, int flags) {
	return ImGui::InputFloat4(label, v, format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputInt(const char* label, int* v, int step, int step_fast, int flags) {
	return ImGui::InputInt(label, v, static_cast<int>(step), static_cast<int>(step_fast), static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputInt2(const char* label, int* v, int flags) {
	return ImGui::InputInt2(label, v, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputInt3(const char* label, int* v, int flags) {
	return ImGui::InputInt3(label, v, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputInt4(const char* label, int* v, int flags) {
	return ImGui::InputInt4(label, v, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputDouble(const char* label, double* v, double step, double step_fast, const char* format, int flags) {
	return ImGui::InputDouble(label, v, static_cast<double>(step), static_cast<double>(step_fast), format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputScalar(const char* label, int data_type, void* p_data, const void* p_step, const void* p_step_fast, const char* format, int flags) {
	return ImGui::InputScalar(label, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), p_step, p_step_fast, format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__InputScalarN(const char* label, int data_type, void* p_data, int components, const void* p_step, const void* p_step_fast, const char* format, int flags) {
	return ImGui::InputScalarN(label, static_cast<ImGuiDataType>(data_type), static_cast<void*>(p_data), static_cast<int>(components), p_step, p_step_fast, format, static_cast<ImGuiInputTextFlags>(flags));
}

bool cabi_ImGui__ColorEdit3(const char* label, float* col, int flags) {
	return ImGui::ColorEdit3(label, col, static_cast<ImGuiColorEditFlags>(flags));
}

bool cabi_ImGui__ColorEdit4(const char* label, float* col, int flags) {
	return ImGui::ColorEdit4(label, col, static_cast<ImGuiColorEditFlags>(flags));
}

bool cabi_ImGui__ColorPicker3(const char* label, float* col, int flags) {
	return ImGui::ColorPicker3(label, col, static_cast<ImGuiColorEditFlags>(flags));
}

bool cabi_ImGui__ColorPicker4(const char* label, float* col, int flags, const float* ref_col) {
	return ImGui::ColorPicker4(label, col, static_cast<ImGuiColorEditFlags>(flags), ref_col);
}

bool cabi_ImGui__ColorButton(const char* desc_id, ImVec4* col, int flags, ImVec2* size) {
	return ImGui::ColorButton(desc_id, *col, static_cast<ImGuiColorEditFlags>(flags), *size);
}

void cabi_ImGui__SetColorEditOptions(int flags) {
	ImGui::SetColorEditOptions(static_cast<ImGuiColorEditFlags>(flags));
}

bool cabi_ImGui__TreeNode(const char* label) {
	return ImGui::TreeNode(label);
}

bool cabi_ImGui__TreeNodeV_1(const char* str_id, const char* fmt, void* args) {
	return ImGui::TreeNodeV(str_id, fmt, *static_cast<va_list*>(args));
}

bool cabi_ImGui__TreeNodeV_2(const void* ptr_id, const char* fmt, void* args) {
	return ImGui::TreeNodeV(ptr_id, fmt, *static_cast<va_list*>(args));
}

bool cabi_ImGui__TreeNodeEx(const char* label, int flags) {
	return ImGui::TreeNodeEx(label, static_cast<ImGuiTreeNodeFlags>(flags));
}

bool cabi_ImGui__TreeNodeExV_1(const char* str_id, int flags, const char* fmt, void* args) {
	return ImGui::TreeNodeExV(str_id, static_cast<ImGuiTreeNodeFlags>(flags), fmt, *static_cast<va_list*>(args));
}

bool cabi_ImGui__TreeNodeExV_2(const void* ptr_id, int flags, const char* fmt, void* args) {
	return ImGui::TreeNodeExV(ptr_id, static_cast<ImGuiTreeNodeFlags>(flags), fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__TreePush_1(const char* str_id) {
	ImGui::TreePush(str_id);
}

void cabi_ImGui__TreePush_2(const void* ptr_id) {
	ImGui::TreePush(ptr_id);
}

void cabi_ImGui__TreePop() {
	ImGui::TreePop();
}

float cabi_ImGui__GetTreeNodeToLabelSpacing() {
	return ImGui::GetTreeNodeToLabelSpacing();
}

bool cabi_ImGui__CollapsingHeader_1(const char* label, int flags) {
	return ImGui::CollapsingHeader(label, static_cast<ImGuiTreeNodeFlags>(flags));
}

bool cabi_ImGui__CollapsingHeader_2(const char* label, bool* p_visible, int flags) {
	return ImGui::CollapsingHeader(label, p_visible, static_cast<ImGuiTreeNodeFlags>(flags));
}

void cabi_ImGui__SetNextItemOpen(bool is_open, int cond) {
	ImGui::SetNextItemOpen(is_open, static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__SetNextItemStorageID(int storage_id) {
	ImGui::SetNextItemStorageID(static_cast<ImGuiID>(storage_id));
}

bool cabi_ImGui__TreeNodeGetOpen(int storage_id) {
	return ImGui::TreeNodeGetOpen(static_cast<ImGuiID>(storage_id));
}

bool cabi_ImGui__Selectable_1(const char* label, bool selected, int flags, ImVec2* size) {
	return ImGui::Selectable(label, selected, static_cast<ImGuiSelectableFlags>(flags), *size);
}

bool cabi_ImGui__Selectable_2(const char* label, bool* p_selected, int flags, ImVec2* size) {
	return ImGui::Selectable(label, p_selected, static_cast<ImGuiSelectableFlags>(flags), *size);
}

ImGuiMultiSelectIO* cabi_ImGui__BeginMultiSelect(int flags, int selection_size, int items_count) {
	return ImGui::BeginMultiSelect(static_cast<ImGuiMultiSelectFlags>(flags), static_cast<int>(selection_size), static_cast<int>(items_count));
}

ImGuiMultiSelectIO* cabi_ImGui__EndMultiSelect() {
	return ImGui::EndMultiSelect();
}

void cabi_ImGui__SetNextItemSelectionUserData(int selection_user_data) {
	ImGui::SetNextItemSelectionUserData(static_cast<ImGuiSelectionUserData>(selection_user_data));
}

bool cabi_ImGui__IsItemToggledSelection() {
	return ImGui::IsItemToggledSelection();
}

bool cabi_ImGui__BeginListBox(const char* label, ImVec2* size) {
	return ImGui::BeginListBox(label, *size);
}

void cabi_ImGui__EndListBox() {
	ImGui::EndListBox();
}

bool cabi_ImGui__ListBox(const char* label, int* current_item, const char** items, int items_count, int height_in_items) {
	return ImGui::ListBox(label, current_item, items, static_cast<int>(items_count), static_cast<int>(height_in_items));
}

void cabi_ImGui__PlotLines(const char* label, const float* values, int values_count, int values_offset, const char* overlay_text, float scale_min, float scale_max, ImVec2* graph_size, int stride) {
	ImGui::PlotLines(label, values, static_cast<int>(values_count), static_cast<int>(values_offset), overlay_text, static_cast<float>(scale_min), static_cast<float>(scale_max), *graph_size, static_cast<int>(stride));
}

void cabi_ImGui__PlotHistogram(const char* label, const float* values, int values_count, int values_offset, const char* overlay_text, float scale_min, float scale_max, ImVec2* graph_size, int stride) {
	ImGui::PlotHistogram(label, values, static_cast<int>(values_count), static_cast<int>(values_offset), overlay_text, static_cast<float>(scale_min), static_cast<float>(scale_max), *graph_size, static_cast<int>(stride));
}

void cabi_ImGui__Value_1(const char* prefix, bool b) {
	ImGui::Value(prefix, b);
}

void cabi_ImGui__Value_2(const char* prefix, int v) {
	ImGui::Value(prefix, static_cast<int>(v));
}

void cabi_ImGui__Value_3(const char* prefix, unsigned int v) {
	ImGui::Value(prefix, static_cast<unsigned int>(v));
}

void cabi_ImGui__Value_4(const char* prefix, float v, const char* float_format) {
	ImGui::Value(prefix, static_cast<float>(v), float_format);
}

bool cabi_ImGui__BeginMenuBar() {
	return ImGui::BeginMenuBar();
}

void cabi_ImGui__EndMenuBar() {
	ImGui::EndMenuBar();
}

bool cabi_ImGui__BeginMainMenuBar() {
	return ImGui::BeginMainMenuBar();
}

void cabi_ImGui__EndMainMenuBar() {
	ImGui::EndMainMenuBar();
}

bool cabi_ImGui__BeginMenu(const char* label, bool enabled) {
	return ImGui::BeginMenu(label, enabled);
}

void cabi_ImGui__EndMenu() {
	ImGui::EndMenu();
}

bool cabi_ImGui__MenuItem_1(const char* label, const char* shortcut, bool selected, bool enabled) {
	return ImGui::MenuItem(label, shortcut, selected, enabled);
}

bool cabi_ImGui__MenuItem_2(const char* label, const char* shortcut, bool* p_selected, bool enabled) {
	return ImGui::MenuItem(label, shortcut, p_selected, enabled);
}

bool cabi_ImGui__BeginTooltip() {
	return ImGui::BeginTooltip();
}

void cabi_ImGui__EndTooltip() {
	ImGui::EndTooltip();
}

void cabi_ImGui__SetTooltipV(const char* fmt, void* args) {
	ImGui::SetTooltipV(fmt, *static_cast<va_list*>(args));
}

bool cabi_ImGui__BeginItemTooltip() {
	return ImGui::BeginItemTooltip();
}

void cabi_ImGui__SetItemTooltipV(const char* fmt, void* args) {
	ImGui::SetItemTooltipV(fmt, *static_cast<va_list*>(args));
}

bool cabi_ImGui__BeginPopup(const char* str_id, int flags) {
	return ImGui::BeginPopup(str_id, static_cast<ImGuiWindowFlags>(flags));
}

bool cabi_ImGui__BeginPopupModal(const char* name, bool* p_open, int flags) {
	return ImGui::BeginPopupModal(name, p_open, static_cast<ImGuiWindowFlags>(flags));
}

void cabi_ImGui__EndPopup() {
	ImGui::EndPopup();
}

void cabi_ImGui__OpenPopup_1(const char* str_id, int popup_flags) {
	ImGui::OpenPopup(str_id, static_cast<ImGuiPopupFlags>(popup_flags));
}

void cabi_ImGui__OpenPopup_2(int id, int popup_flags) {
	ImGui::OpenPopup(static_cast<ImGuiID>(id), static_cast<ImGuiPopupFlags>(popup_flags));
}

void cabi_ImGui__OpenPopupOnItemClick(const char* str_id, int popup_flags) {
	ImGui::OpenPopupOnItemClick(str_id, static_cast<ImGuiPopupFlags>(popup_flags));
}

void cabi_ImGui__CloseCurrentPopup() {
	ImGui::CloseCurrentPopup();
}

bool cabi_ImGui__BeginPopupContextItem(const char* str_id, int popup_flags) {
	return ImGui::BeginPopupContextItem(str_id, static_cast<ImGuiPopupFlags>(popup_flags));
}

bool cabi_ImGui__BeginPopupContextWindow(const char* str_id, int popup_flags) {
	return ImGui::BeginPopupContextWindow(str_id, static_cast<ImGuiPopupFlags>(popup_flags));
}

bool cabi_ImGui__BeginPopupContextVoid(const char* str_id, int popup_flags) {
	return ImGui::BeginPopupContextVoid(str_id, static_cast<ImGuiPopupFlags>(popup_flags));
}

bool cabi_ImGui__IsPopupOpen(const char* str_id, int flags) {
	return ImGui::IsPopupOpen(str_id, static_cast<ImGuiPopupFlags>(flags));
}

bool cabi_ImGui__BeginTable(const char* str_id, int columns, int flags, ImVec2* outer_size, float inner_width) {
	return ImGui::BeginTable(str_id, static_cast<int>(columns), static_cast<ImGuiTableFlags>(flags), *outer_size, static_cast<float>(inner_width));
}

void cabi_ImGui__EndTable() {
	ImGui::EndTable();
}

void cabi_ImGui__TableNextRow(int row_flags, float min_row_height) {
	ImGui::TableNextRow(static_cast<ImGuiTableRowFlags>(row_flags), static_cast<float>(min_row_height));
}

bool cabi_ImGui__TableNextColumn() {
	return ImGui::TableNextColumn();
}

bool cabi_ImGui__TableSetColumnIndex(int column_n) {
	return ImGui::TableSetColumnIndex(static_cast<int>(column_n));
}

void cabi_ImGui__TableSetupColumn(const char* label, int flags, float init_width_or_weight, int user_id) {
	ImGui::TableSetupColumn(label, static_cast<ImGuiTableColumnFlags>(flags), static_cast<float>(init_width_or_weight), static_cast<ImGuiID>(user_id));
}

void cabi_ImGui__TableSetupScrollFreeze(int cols, int rows) {
	ImGui::TableSetupScrollFreeze(static_cast<int>(cols), static_cast<int>(rows));
}

void cabi_ImGui__TableHeader(const char* label) {
	ImGui::TableHeader(label);
}

void cabi_ImGui__TableHeadersRow() {
	ImGui::TableHeadersRow();
}

void cabi_ImGui__TableAngledHeadersRow() {
	ImGui::TableAngledHeadersRow();
}

ImGuiTableSortSpecs* cabi_ImGui__TableGetSortSpecs() {
	return ImGui::TableGetSortSpecs();
}

int cabi_ImGui__TableGetColumnCount() {
	return ImGui::TableGetColumnCount();
}

int cabi_ImGui__TableGetColumnIndex() {
	return ImGui::TableGetColumnIndex();
}

int cabi_ImGui__TableGetRowIndex() {
	return ImGui::TableGetRowIndex();
}

const char* cabi_ImGui__TableGetColumnName(int column_n) {
	return (const char*) ImGui::TableGetColumnName(static_cast<int>(column_n));
}

int cabi_ImGui__TableGetColumnFlags(int column_n) {
	return ImGui::TableGetColumnFlags(static_cast<int>(column_n));
}

void cabi_ImGui__TableSetColumnEnabled(int column_n, bool v) {
	ImGui::TableSetColumnEnabled(static_cast<int>(column_n), v);
}

int cabi_ImGui__TableGetHoveredColumn() {
	return ImGui::TableGetHoveredColumn();
}

void cabi_ImGui__TableSetBgColor(int target, int color, int column_n) {
	ImGui::TableSetBgColor(static_cast<ImGuiTableBgTarget>(target), static_cast<ImU32>(color), static_cast<int>(column_n));
}

void cabi_ImGui__Columns(int count, const char* id, bool borders) {
	ImGui::Columns(static_cast<int>(count), id, borders);
}

void cabi_ImGui__NextColumn() {
	ImGui::NextColumn();
}

int cabi_ImGui__GetColumnIndex() {
	return ImGui::GetColumnIndex();
}

float cabi_ImGui__GetColumnWidth(int column_index) {
	return ImGui::GetColumnWidth(static_cast<int>(column_index));
}

void cabi_ImGui__SetColumnWidth(int column_index, float width) {
	ImGui::SetColumnWidth(static_cast<int>(column_index), static_cast<float>(width));
}

float cabi_ImGui__GetColumnOffset(int column_index) {
	return ImGui::GetColumnOffset(static_cast<int>(column_index));
}

void cabi_ImGui__SetColumnOffset(int column_index, float offset_x) {
	ImGui::SetColumnOffset(static_cast<int>(column_index), static_cast<float>(offset_x));
}

int cabi_ImGui__GetColumnsCount() {
	return ImGui::GetColumnsCount();
}

bool cabi_ImGui__BeginTabBar(const char* str_id, int flags) {
	return ImGui::BeginTabBar(str_id, static_cast<ImGuiTabBarFlags>(flags));
}

void cabi_ImGui__EndTabBar() {
	ImGui::EndTabBar();
}

bool cabi_ImGui__BeginTabItem(const char* label, bool* p_open, int flags) {
	return ImGui::BeginTabItem(label, p_open, static_cast<ImGuiTabItemFlags>(flags));
}

void cabi_ImGui__EndTabItem() {
	ImGui::EndTabItem();
}

bool cabi_ImGui__TabItemButton(const char* label, int flags) {
	return ImGui::TabItemButton(label, static_cast<ImGuiTabItemFlags>(flags));
}

void cabi_ImGui__SetTabItemClosed(const char* tab_or_docked_window_label) {
	ImGui::SetTabItemClosed(tab_or_docked_window_label);
}

void cabi_ImGui__LogToTTY(int auto_open_depth) {
	ImGui::LogToTTY(static_cast<int>(auto_open_depth));
}

void cabi_ImGui__LogToFile(int auto_open_depth, const char* filename) {
	ImGui::LogToFile(static_cast<int>(auto_open_depth), filename);
}

void cabi_ImGui__LogToClipboard(int auto_open_depth) {
	ImGui::LogToClipboard(static_cast<int>(auto_open_depth));
}

void cabi_ImGui__LogFinish() {
	ImGui::LogFinish();
}

void cabi_ImGui__LogButtons() {
	ImGui::LogButtons();
}

void cabi_ImGui__LogTextV(const char* fmt, void* args) {
	ImGui::LogTextV(fmt, *static_cast<va_list*>(args));
}

bool cabi_ImGui__BeginDragDropSource(int flags) {
	return ImGui::BeginDragDropSource(static_cast<ImGuiDragDropFlags>(flags));
}

bool cabi_ImGui__SetDragDropPayload(const char* type, const void* data, size_t sz, int cond) {
	return ImGui::SetDragDropPayload(type, data, static_cast<size_t>(sz), static_cast<ImGuiCond>(cond));
}

void cabi_ImGui__EndDragDropSource() {
	ImGui::EndDragDropSource();
}

bool cabi_ImGui__BeginDragDropTarget() {
	return ImGui::BeginDragDropTarget();
}

ImGuiPayload* cabi_ImGui__AcceptDragDropPayload(const char* type, int flags) {
	return (ImGuiPayload*) ImGui::AcceptDragDropPayload(type, static_cast<ImGuiDragDropFlags>(flags));
}

void cabi_ImGui__EndDragDropTarget() {
	ImGui::EndDragDropTarget();
}

ImGuiPayload* cabi_ImGui__GetDragDropPayload() {
	return (ImGuiPayload*) ImGui::GetDragDropPayload();
}

void cabi_ImGui__BeginDisabled(bool disabled) {
	ImGui::BeginDisabled(disabled);
}

void cabi_ImGui__EndDisabled() {
	ImGui::EndDisabled();
}

void cabi_ImGui__PushClipRect(ImVec2* clip_rect_min, ImVec2* clip_rect_max, bool intersect_with_current_clip_rect) {
	ImGui::PushClipRect(*clip_rect_min, *clip_rect_max, intersect_with_current_clip_rect);
}

void cabi_ImGui__PopClipRect() {
	ImGui::PopClipRect();
}

void cabi_ImGui__SetItemDefaultFocus() {
	ImGui::SetItemDefaultFocus();
}

void cabi_ImGui__SetKeyboardFocusHere(int offset) {
	ImGui::SetKeyboardFocusHere(static_cast<int>(offset));
}

void cabi_ImGui__SetNavCursorVisible(bool visible) {
	ImGui::SetNavCursorVisible(visible);
}

void cabi_ImGui__SetNextItemAllowOverlap() {
	ImGui::SetNextItemAllowOverlap();
}

bool cabi_ImGui__IsItemHovered(int flags) {
	return ImGui::IsItemHovered(static_cast<ImGuiHoveredFlags>(flags));
}

bool cabi_ImGui__IsItemActive() {
	return ImGui::IsItemActive();
}

bool cabi_ImGui__IsItemFocused() {
	return ImGui::IsItemFocused();
}

bool cabi_ImGui__IsItemClicked(int mouse_button) {
	return ImGui::IsItemClicked(static_cast<ImGuiMouseButton>(mouse_button));
}

bool cabi_ImGui__IsItemVisible() {
	return ImGui::IsItemVisible();
}

bool cabi_ImGui__IsItemEdited() {
	return ImGui::IsItemEdited();
}

bool cabi_ImGui__IsItemActivated() {
	return ImGui::IsItemActivated();
}

bool cabi_ImGui__IsItemDeactivated() {
	return ImGui::IsItemDeactivated();
}

bool cabi_ImGui__IsItemDeactivatedAfterEdit() {
	return ImGui::IsItemDeactivatedAfterEdit();
}

bool cabi_ImGui__IsItemToggledOpen() {
	return ImGui::IsItemToggledOpen();
}

bool cabi_ImGui__IsAnyItemHovered() {
	return ImGui::IsAnyItemHovered();
}

bool cabi_ImGui__IsAnyItemActive() {
	return ImGui::IsAnyItemActive();
}

bool cabi_ImGui__IsAnyItemFocused() {
	return ImGui::IsAnyItemFocused();
}

int cabi_ImGui__GetItemID() {
	return ImGui::GetItemID();
}

ImVec2* cabi_ImGui__GetItemRectMin() {
	return new ImVec2(ImGui::GetItemRectMin());
}

ImVec2* cabi_ImGui__GetItemRectMax() {
	return new ImVec2(ImGui::GetItemRectMax());
}

ImVec2* cabi_ImGui__GetItemRectSize() {
	return new ImVec2(ImGui::GetItemRectSize());
}

int cabi_ImGui__GetItemFlags() {
	return ImGui::GetItemFlags();
}

ImGuiViewport* cabi_ImGui__GetMainViewport() {
	return ImGui::GetMainViewport();
}

ImDrawList* cabi_ImGui__GetBackgroundDrawList() {
	return ImGui::GetBackgroundDrawList();
}

ImDrawList* cabi_ImGui__GetForegroundDrawList() {
	return ImGui::GetForegroundDrawList();
}

bool cabi_ImGui__IsRectVisible_1(ImVec2* size) {
	return ImGui::IsRectVisible(*size);
}

bool cabi_ImGui__IsRectVisible_2(ImVec2* rect_min, ImVec2* rect_max) {
	return ImGui::IsRectVisible(*rect_min, *rect_max);
}

double cabi_ImGui__GetTime() {
	return ImGui::GetTime();
}

int cabi_ImGui__GetFrameCount() {
	return ImGui::GetFrameCount();
}

void* cabi_ImGui__GetDrawListSharedData() {
	return ImGui::GetDrawListSharedData();
}

const char* cabi_ImGui__GetStyleColorName(int idx) {
	return (const char*) ImGui::GetStyleColorName(static_cast<ImGuiCol>(idx));
}

void cabi_ImGui__SetStateStorage(ImGuiStorage* storage) {
	ImGui::SetStateStorage(storage);
}

ImGuiStorage* cabi_ImGui__GetStateStorage() {
	return ImGui::GetStateStorage();
}

ImVec2* cabi_ImGui__CalcTextSize(const char* text, const char* text_end, bool hide_text_after_double_hash, float wrap_width) {
	return new ImVec2(ImGui::CalcTextSize(text, text_end, hide_text_after_double_hash, static_cast<float>(wrap_width)));
}

ImVec4* cabi_ImGui__ColorConvertU32ToFloat4(int in) {
	return new ImVec4(ImGui::ColorConvertU32ToFloat4(static_cast<ImU32>(in)));
}

int cabi_ImGui__ColorConvertFloat4ToU32(ImVec4* in) {
	return ImGui::ColorConvertFloat4ToU32(*in);
}

void cabi_ImGui__ColorConvertRGBtoHSV(float r, float g, float b, float* out_h, float* out_s, float* out_v) {
	ImGui::ColorConvertRGBtoHSV(static_cast<float>(r), static_cast<float>(g), static_cast<float>(b), *out_h, *out_s, *out_v);
}

void cabi_ImGui__ColorConvertHSVtoRGB(float h, float s, float v, float* out_r, float* out_g, float* out_b) {
	ImGui::ColorConvertHSVtoRGB(static_cast<float>(h), static_cast<float>(s), static_cast<float>(v), *out_r, *out_g, *out_b);
}

bool cabi_ImGui__IsKeyDown(int key) {
	return ImGui::IsKeyDown(static_cast<ImGuiKey>(key));
}

bool cabi_ImGui__IsKeyPressed(int key, bool repeat) {
	return ImGui::IsKeyPressed(static_cast<ImGuiKey>(key), repeat);
}

bool cabi_ImGui__IsKeyReleased(int key) {
	return ImGui::IsKeyReleased(static_cast<ImGuiKey>(key));
}

bool cabi_ImGui__IsKeyChordPressed(int key_chord) {
	return ImGui::IsKeyChordPressed(static_cast<ImGuiKeyChord>(key_chord));
}

int cabi_ImGui__GetKeyPressedAmount(int key, float repeat_delay, float rate) {
	return ImGui::GetKeyPressedAmount(static_cast<ImGuiKey>(key), static_cast<float>(repeat_delay), static_cast<float>(rate));
}

const char* cabi_ImGui__GetKeyName(int key) {
	return (const char*) ImGui::GetKeyName(static_cast<ImGuiKey>(key));
}

void cabi_ImGui__SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) {
	ImGui::SetNextFrameWantCaptureKeyboard(want_capture_keyboard);
}

bool cabi_ImGui__Shortcut(int key_chord, int flags) {
	return ImGui::Shortcut(static_cast<ImGuiKeyChord>(key_chord), static_cast<ImGuiInputFlags>(flags));
}

void cabi_ImGui__SetNextItemShortcut(int key_chord, int flags) {
	ImGui::SetNextItemShortcut(static_cast<ImGuiKeyChord>(key_chord), static_cast<ImGuiInputFlags>(flags));
}

bool cabi_ImGui__SetItemKeyOwner(int key) {
	return ImGui::SetItemKeyOwner(static_cast<ImGuiKey>(key));
}

bool cabi_ImGui__IsMouseDown(int button) {
	return ImGui::IsMouseDown(static_cast<ImGuiMouseButton>(button));
}

bool cabi_ImGui__IsMouseClicked(int button, bool repeat) {
	return ImGui::IsMouseClicked(static_cast<ImGuiMouseButton>(button), repeat);
}

bool cabi_ImGui__IsMouseReleased(int button) {
	return ImGui::IsMouseReleased(static_cast<ImGuiMouseButton>(button));
}

bool cabi_ImGui__IsMouseDoubleClicked(int button) {
	return ImGui::IsMouseDoubleClicked(static_cast<ImGuiMouseButton>(button));
}

bool cabi_ImGui__IsMouseReleasedWithDelay(int button, float delay) {
	return ImGui::IsMouseReleasedWithDelay(static_cast<ImGuiMouseButton>(button), static_cast<float>(delay));
}

int cabi_ImGui__GetMouseClickedCount(int button) {
	return ImGui::GetMouseClickedCount(static_cast<ImGuiMouseButton>(button));
}

bool cabi_ImGui__IsMouseHoveringRect(ImVec2* r_min, ImVec2* r_max, bool clip) {
	return ImGui::IsMouseHoveringRect(*r_min, *r_max, clip);
}

bool cabi_ImGui__IsMousePosValid(ImVec2* mouse_pos) {
	return ImGui::IsMousePosValid(mouse_pos);
}

bool cabi_ImGui__IsAnyMouseDown() {
	return ImGui::IsAnyMouseDown();
}

ImVec2* cabi_ImGui__GetMousePos() {
	return new ImVec2(ImGui::GetMousePos());
}

ImVec2* cabi_ImGui__GetMousePosOnOpeningCurrentPopup() {
	return new ImVec2(ImGui::GetMousePosOnOpeningCurrentPopup());
}

bool cabi_ImGui__IsMouseDragging(int button, float lock_threshold) {
	return ImGui::IsMouseDragging(static_cast<ImGuiMouseButton>(button), static_cast<float>(lock_threshold));
}

ImVec2* cabi_ImGui__GetMouseDragDelta(int button, float lock_threshold) {
	return new ImVec2(ImGui::GetMouseDragDelta(static_cast<ImGuiMouseButton>(button), static_cast<float>(lock_threshold)));
}

void cabi_ImGui__ResetMouseDragDelta(int button) {
	ImGui::ResetMouseDragDelta(static_cast<ImGuiMouseButton>(button));
}

int cabi_ImGui__GetMouseCursor() {
	return ImGui::GetMouseCursor();
}

void cabi_ImGui__SetMouseCursor(int cursor_type) {
	ImGui::SetMouseCursor(static_cast<ImGuiMouseCursor>(cursor_type));
}

void cabi_ImGui__SetNextFrameWantCaptureMouse(bool want_capture_mouse) {
	ImGui::SetNextFrameWantCaptureMouse(want_capture_mouse);
}

const char* cabi_ImGui__GetClipboardText() {
	return (const char*) ImGui::GetClipboardText();
}

void cabi_ImGui__SetClipboardText(const char* text) {
	ImGui::SetClipboardText(text);
}

void cabi_ImGui__LoadIniSettingsFromDisk(const char* ini_filename) {
	ImGui::LoadIniSettingsFromDisk(ini_filename);
}

void cabi_ImGui__LoadIniSettingsFromMemory(const char* ini_data, size_t ini_size) {
	ImGui::LoadIniSettingsFromMemory(ini_data, static_cast<size_t>(ini_size));
}

void cabi_ImGui__SaveIniSettingsToDisk(const char* ini_filename) {
	ImGui::SaveIniSettingsToDisk(ini_filename);
}

const char* cabi_ImGui__SaveIniSettingsToMemory(size_t* out_ini_size) {
	return (const char*) ImGui::SaveIniSettingsToMemory(out_ini_size);
}

void cabi_ImGui__DebugTextEncoding(const char* text) {
	ImGui::DebugTextEncoding(text);
}

void cabi_ImGui__DebugFlashStyleColor(int idx) {
	ImGui::DebugFlashStyleColor(static_cast<ImGuiCol>(idx));
}

void cabi_ImGui__DebugStartItemPicker() {
	ImGui::DebugStartItemPicker();
}

bool cabi_ImGui__DebugCheckVersionAndDataLayout(const char* version_str, size_t sz_io, size_t sz_style, size_t sz_vec2, size_t sz_vec4, size_t sz_drawvert, size_t sz_drawidx) {
	return ImGui::DebugCheckVersionAndDataLayout(version_str, static_cast<size_t>(sz_io), static_cast<size_t>(sz_style), static_cast<size_t>(sz_vec2), static_cast<size_t>(sz_vec4), static_cast<size_t>(sz_drawvert), static_cast<size_t>(sz_drawidx));
}

void cabi_ImGui__DebugLogV(const char* fmt, void* args) {
	ImGui::DebugLogV(fmt, *static_cast<va_list*>(args));
}

void cabi_ImGui__SetAllocatorFunctions(void* alloc_func, void* free_func, void* user_data) {
	ImGui::SetAllocatorFunctions((ImGuiMemAllocFunc)(alloc_func), (ImGuiMemFreeFunc)(free_func), static_cast<void*>(user_data));
}

void cabi_ImGui__GetAllocatorFunctions(void** p_alloc_func, void** p_free_func, void** p_user_data) {
	ImGui::GetAllocatorFunctions(reinterpret_cast<ImGuiMemAllocFunc*>(p_alloc_func), reinterpret_cast<ImGuiMemFreeFunc*>(p_free_func), p_user_data);
}

void* cabi_ImGui__MemAlloc(size_t size) {
	return ImGui::MemAlloc(static_cast<size_t>(size));
}

void cabi_ImGui__MemFree(void* ptr) {
	ImGui::MemFree(static_cast<void*>(ptr));
}

void cabi_ImGui__PushFont_2(ImFont* font) {
	ImGui::PushFont(font);
}

void cabi_ImGui__SetWindowFontScale(float scale) {
	ImGui::SetWindowFontScale(static_cast<float>(scale));
}

void cabi_ImGui__Image_2(ImTextureRef* tex_ref, ImVec2* image_size, ImVec2* uv0, ImVec2* uv1, ImVec4* tint_col, ImVec4* border_col) {
	ImGui::Image(*tex_ref, *image_size, *uv0, *uv1, *tint_col, *border_col);
}

void cabi_ImGui__PushButtonRepeat(bool repeat) {
	ImGui::PushButtonRepeat(repeat);
}

void cabi_ImGui__PopButtonRepeat() {
	ImGui::PopButtonRepeat();
}

void cabi_ImGui__PushTabStop(bool tab_stop) {
	ImGui::PushTabStop(tab_stop);
}

void cabi_ImGui__PopTabStop() {
	ImGui::PopTabStop();
}

ImVec2* cabi_ImGui__GetContentRegionMax() {
	return new ImVec2(ImGui::GetContentRegionMax());
}

ImVec2* cabi_ImGui__GetWindowContentRegionMin() {
	return new ImVec2(ImGui::GetWindowContentRegionMin());
}

ImVec2* cabi_ImGui__GetWindowContentRegionMax() {
	return new ImVec2(ImGui::GetWindowContentRegionMax());
}

bool cabi_ImGui_ImplWin32_Init(void* hwnd) {
	return ImGui_ImplWin32_Init(static_cast<void*>(hwnd));
}

bool cabi_ImGui_ImplWin32_InitForOpenGL(void* hwnd) {
	return ImGui_ImplWin32_InitForOpenGL(static_cast<void*>(hwnd));
}

void cabi_ImGui_ImplWin32_Shutdown() {
	ImGui_ImplWin32_Shutdown();
}

void cabi_ImGui_ImplWin32_NewFrame() {
	ImGui_ImplWin32_NewFrame();
}

void cabi_ImGui_ImplWin32_EnableDpiAwareness() {
	ImGui_ImplWin32_EnableDpiAwareness();
}

float cabi_ImGui_ImplWin32_GetDpiScaleForHwnd(void* hwnd) {
	return ImGui_ImplWin32_GetDpiScaleForHwnd(static_cast<void*>(hwnd));
}

float cabi_ImGui_ImplWin32_GetDpiScaleForMonitor(void* monitor) {
	return ImGui_ImplWin32_GetDpiScaleForMonitor(static_cast<void*>(monitor));
}

void cabi_ImGui_ImplWin32_EnableAlphaCompositing(void* hwnd) {
	ImGui_ImplWin32_EnableAlphaCompositing(static_cast<void*>(hwnd));
}

bool cabi_ImGui_ImplDX11_Init(void* device, void* device_context) {
	return ImGui_ImplDX11_Init(static_cast<ID3D11Device*>(device), static_cast<ID3D11DeviceContext*>(device_context));
}

void cabi_ImGui_ImplDX11_Shutdown() {
	ImGui_ImplDX11_Shutdown();
}

void cabi_ImGui_ImplDX11_NewFrame() {
	ImGui_ImplDX11_NewFrame();
}

void cabi_ImGui_ImplDX11_RenderDrawData(ImDrawData* draw_data) {
	ImGui_ImplDX11_RenderDrawData(draw_data);
}

bool cabi_ImGui_ImplDX11_CreateDeviceObjects() {
	return ImGui_ImplDX11_CreateDeviceObjects();
}

void cabi_ImGui_ImplDX11_InvalidateDeviceObjects() {
	ImGui_ImplDX11_InvalidateDeviceObjects();
}

void cabi_ImGui_ImplDX11_UpdateTexture(ImTextureData* tex) {
	ImGui_ImplDX11_UpdateTexture(tex);
}

intptr_t cabi_ImGui_ImplWin32_WndProcHandler(void* hWnd, int msg, uintptr_t wParam, intptr_t lParam) {
	return ImGui_ImplWin32_WndProcHandler(reinterpret_cast<HWND>(hWnd), static_cast<UINT>(msg), static_cast<WPARAM>(wParam), static_cast<LPARAM>(lParam));
}

