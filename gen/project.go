package gen

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/bindgen/cpp2c"
)

func init() {
	ImGuiProject.Cflags = `--std=c++17 --target=x86_64-pc-windows-msvc -Wno-everything -I` +
		`cpp` + cpp2c.MsvcIncludeFlags()
}

func RunImGui() {
	cpp2c.SetActiveFilter(NewImGuiFilter())
	originalCacheDir := cpp2c.CacheDir
	cpp2c.CacheDir = ImGuiProject.CacheDir

	cpp2c.ProcessProjectHeaders(&ImGuiProject)

	cpp2c.CacheDir = originalCacheDir

	if ImGuiProject.GenerateCMake != nil {
		ImGuiProject.GenerateCMake(ImGuiProject.OutputCppDir, "imgui")
	}

	cpp2c.RunBuild(ImGuiProject.OutputCppDir, "imgui_bindings.dll")
}

var ImGuiProject = cpp2c.ProjectConfig{
	Name:          "imgui",
	SourceDir:     `cpp`,
	OutputCppDir:  `cabi`, // todo 从源头删除这个字段
	OutputCabiDir: `cabi`,
	CacheDir:      `cachedir/imgui`,
	Headers: []cpp2c.HeaderConfig{
		{Path: `cpp/imgui_backends.h`, AllowAll: true},
	},
	AllowedClassPrefix: []string{"Im", "ID3D11", "IDXGI", "ImGui"},
	GenerateCMake:      generateCMakeListsImGui,
	GenerateAuxiliary:  generateAuxiliaryImGui,
	Emit: cpp2c.EmitConfig{
		HeaderVersionMacro: "IMGUI_VERSION_NUM",
		AllowedEnumPrefix:  []string{"Im", "ImGui"},
		ValueTypes: map[string]string{
			"ImColor": "",
		},
		FreeFuncCabiPrefix: "cabi_",
		SkipLibmiqtInclude: true,
	},
}

func generateCMakeListsImGui(outDir, packageName string) {
	cmakePath := filepath.Join(outDir, "CMakeLists.txt")
	generateAuxiliaryImGui(outDir)

	var sb strings.Builder
	sb.WriteString(`cmake_minimum_required(VERSION 3.16)
include("C:/Program Files/CMake/bin/ewdk.cmake")

project(` + packageName + `_bindings C CXX)

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

um_dll(` + packageName + `_bindings
    ${MIQT_GEN_SOURCES}
    ${IMGUI_SOURCES}
)

target_include_directories(` + packageName + `_bindings PUBLIC
    ${IMGUI_DIR}
    ${IMGUI_DIR}/backends
)

target_compile_definitions(` + packageName + `_bindings PRIVATE
    IMGUI_BUILDING_DLL
    IMGUI_USER_CONFIG=<imgui_dll_config.h>
)

target_link_libraries(` + packageName + `_bindings
    user32.lib
    d3d11.lib
    dxgi.lib
    dxguid.lib
    imm32.lib
)

set_target_properties(` + packageName + `_bindings PROPERTIES
    CXX_STANDARD 17
    CXX_STANDARD_REQUIRED ON
)
`)

	if err := os.WriteFile(cmakePath, []byte(sb.String()), 0o644); err != nil {
		panic(err)
	}
	cpp2c.GenerateBuildBat(outDir)
}

func generateAuxiliaryImGui(outDir string) {
	cpp2c.GenerateCommonFiles(outDir)
}
