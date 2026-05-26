cmake_minimum_required(VERSION 3.16)
include("C:/Program Files/CMake/bin/ewdk.cmake")

project(imgui_bindings C CXX)

set(IMGUI_DIR ${CMAKE_CURRENT_SOURCE_DIR}/../cpp)

set(IMGUI_SOURCES
    ${IMGUI_DIR}/imgui.cpp
    ${IMGUI_DIR}/imgui_demo.cpp
    ${IMGUI_DIR}/imgui_draw.cpp
    ${IMGUI_DIR}/imgui_tables.cpp
    ${IMGUI_DIR}/imgui_widgets.cpp
    ${IMGUI_DIR}/backends/imgui_impl_win32.cpp
    ${IMGUI_DIR}/backends/imgui_impl_dx11.cpp
)

file(GLOB MIQT_GEN_SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/gen_*.cpp)

um_dll(imgui_bindings
    ${MIQT_GEN_SOURCES}
    ${IMGUI_SOURCES}
)

target_include_directories(imgui_bindings PUBLIC
    ${IMGUI_DIR}
    ${IMGUI_DIR}/backends
)

target_compile_definitions(imgui_bindings PRIVATE
    IMGUI_BUILDING_DLL
    IMGUI_USER_CONFIG=<imgui_dll_config.h>
)

target_link_libraries(imgui_bindings
    user32.lib
    d3d11.lib
    dxgi.lib
    dxguid.lib
    imm32.lib
)

set_target_properties(imgui_bindings PROPERTIES
    CXX_STANDARD 17
    CXX_STANDARD_REQUIRED ON
)
