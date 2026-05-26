#pragma once

#include "imgui.h"

#ifndef IMGUI_DISABLE

#include "backends/imgui_impl_win32.h"
#include "backends/imgui_impl_dx11.h"

// ImGui_ImplWin32_WndProcHandler is inside #if 0 in the original header
// to avoid dragging <windows.h> dependencies. We need it visible for binding generation.
// Forward declare it here so clang can parse it.
#ifdef _WIN32
#include <windows.h>
extern IMGUI_IMPL_API LRESULT ImGui_ImplWin32_WndProcHandler(HWND hWnd, UINT msg, WPARAM wParam, LPARAM lParam);
#endif

#endif
