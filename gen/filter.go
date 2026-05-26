package gen

import (
	"strings"

	"github.com/ddkwork/bindgen/cpp2c"
)

type ImGuiFilter struct{}

func NewImGuiFilter() *ImGuiFilter {
	return &ImGuiFilter{}
}

func (f *ImGuiFilter) AllowClass(className string) bool {
	if strings.Contains(className, "::") {
		return false
	}
	return cpp2c.DefaultAllowClass(className)
}

func (f *ImGuiFilter) AllowMethod(className string, mm cpp2c.CppMethod) error {
	return cpp2c.DefaultAllowMethod(className, mm)
}

func (f *ImGuiFilter) AllowFreeFunction(fn cpp2c.CppFreeFunction) error {
	err := cpp2c.DefaultAllowFreeFunction(fn)
	if err != nil && (strings.HasPrefix(fn.FunctionName, "ImGui_Impl") || strings.HasPrefix(fn.FunctionName, "ImGui::")) {
		return nil
	}
	return err
}

func (f *ImGuiFilter) AllowSignal(mm cpp2c.CppMethod) bool {
	return cpp2c.DefaultAllowSignal(mm)
}

func (f *ImGuiFilter) AllowInheritedParent(className string) bool {
	return true
}

func (f *ImGuiFilter) AllowVirtual(mm cpp2c.CppMethod) bool {
	return true
}

func (f *ImGuiFilter) AllowVirtualForClass(className string) bool {
	return true
}

func (f *ImGuiFilter) AllowCtor(className string, mm cpp2c.CppMethod) bool {
	return true
}

func (f *ImGuiFilter) AllowType(p cpp2c.CppParameter, isReturnType bool) error {
	if strings.Contains(p.ParameterType, "::") {
		return cpp2c.ErrTooComplex
	}
	return nil
}

func (f *ImGuiFilter) InsertTypedefs() {}

func (f *ImGuiFilter) HeaderPlatformRestriction(fullpath string) cpp2c.AllowedPlatformInfo {
	return nil
}

func (f *ImGuiFilter) GetPureVirtualDefaults(className string) []cpp2c.PureVirtualDefault {
	return nil
}

func (f *ImGuiFilter) ApplyQuirks(packageName, className string, mm *cpp2c.CppMethod) {}

func (f *ImGuiFilter) PreventStructDeclaration(className string) bool {
	return false
}

func (f *ImGuiFilter) ShouldEmitForwardDeclaration(className string) bool {
	return false
}

func (f *ImGuiFilter) ShouldEmitStructDefinition(className string) bool {
	switch className {
	case "ImVec2", "ImVec4":
		return true
	}
	return false
}

func (f *ImGuiFilter) ShouldEmitEnumDefinition(enumName string) bool {
	return strings.HasPrefix(enumName, "Im") || strings.HasPrefix(enumName, "ImGui")
}
