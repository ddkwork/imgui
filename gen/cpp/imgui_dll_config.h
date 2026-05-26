#pragma once

#ifdef IMGUI_BUILDING_DLL
#define IMGUI_API __declspec(dllexport)
#else
#define IMGUI_API __declspec(dllimport)
#endif
