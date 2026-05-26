package gen

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	RunImGui()
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:  "cabi",
		OutputDir:   "../",
		PackageName: "imgui",
		HeaderOrder: []string{
			"gen_cabi_imgui_backends.h",
		},
		BindDll: true,
		DllName: "imgui_bindings.dll",
		Predefined: `
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long long uint64_t;
typedef signed char int8_t;
typedef short int16_t;
typedef int int32_t;
typedef long long int64_t;
typedef unsigned long size_t;
typedef long long intptr_t;
typedef unsigned long long uintptr_t;
#define true 1
#define false 0
#define MIQT_EXPORT
#define MIQT_CALL
#define MIQT_EXPORT_H
#define MIQT_LIBMIQT_LIBMIQT_H
`,
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "ImGui") || strings.HasPrefix(name, "Im") || strings.HasPrefix(name, "cabi_")
		},
	}})
}

func TestCCbug(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:  "cabi",
		OutputDir:   "../",
		PackageName: "imgui",
		HeaderOrder: []string{
			"gen_cabi_imgui_backends.h",
		},
		BindDll: true,
		DllName: "imgui_bindings.dll",
		Predefined: `
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long long uint64_t;
typedef signed char int8_t;
typedef short int16_t;
typedef int int32_t;
typedef long long int64_t;
typedef unsigned long size_t;
typedef long long intptr_t;
typedef unsigned long long uintptr_t;
#define true 1
#define false 0
#define MIQT_EXPORT
#define MIQT_CALL
#define MIQT_EXPORT_H
#define MIQT_LIBMIQT_LIBMIQT_H
`,
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "ImGui") || strings.HasPrefix(name, "Im") || strings.HasPrefix(name, "cabi_")
		},
	}})
}
