cmake_minimum_required(VERSION 3.16)
include("C:/Program Files/CMake/bin/ewdk.cmake")

project(imgui_tree_table_demo)

set(CABI_DIR ${CMAKE_CURRENT_SOURCE_DIR}/../cabi)

um_exe(imgui_tree_table_demo
    main.cpp
)

target_link_libraries(imgui_tree_table_demo
    user32.lib
    d3d11.lib
    dxgi.lib
)

set_target_properties(imgui_tree_table_demo PROPERTIES
    CXX_STANDARD 17
    CXX_STANDARD_REQUIRED ON
)

add_custom_command(TARGET imgui_tree_table_demo POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CABI_DIR}/build/imgui_bindings.dll"
        "$<TARGET_FILE_DIR:imgui_tree_table_demo>"
)
