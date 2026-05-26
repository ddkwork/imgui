package gen

import (
	"testing"

	"github.com/ddkwork/bindgen/cpp2c"
)

func TestImGuiFilterAllowClass(t *testing.T) {
	filter := NewImGuiFilter()

	tests := []struct {
		name      string
		className string
		expected  bool
	}{
		{"ImGui", "ImGui", true},
		{"ImVec2", "ImVec2", true},
		{"ImVec4", "ImVec4", true},
		{"ImDrawData", "ImDrawData", true},
		{"ImTextureID", "ImTextureID", true},
		{"ImU32", "ImU32", true},
		{"std::string", "std::string", false},
		{"QString", "QString", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.AllowClass(tt.className)
			if result != tt.expected {
				t.Errorf("AllowClass(%q) = %v, want %v", tt.className, result, tt.expected)
			}
		})
	}
}

func TestImGuiFilterAllowMethod(t *testing.T) {
	filter := NewImGuiFilter()

	tests := []struct {
		name      string
		className string
		method    cpp2c.CppMethod
		wantErr   bool
	}{
		{
			name:      "normal method",
			className: "ImGui",
			method: cpp2c.CppMethod{
				MethodName: "Begin",
			},
			wantErr: false,
		},
		{
			name:      "text method",
			className: "ImGui",
			method: cpp2c.CppMethod{
				MethodName: "Text",
			},
			wantErr: false,
		},
		{
			name:      "operator method",
			className: "ImVec2",
			method: cpp2c.CppMethod{
				MethodName: "operator+",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := filter.AllowMethod(tt.className, tt.method)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowMethod(%q.%q) error = %v, wantErr %v", tt.className, tt.method.MethodName, err, tt.wantErr)
			}
		})
	}
}

func TestImGuiFilterAllowVirtual(t *testing.T) {
	filter := NewImGuiFilter()

	tests := []struct {
		name     string
		method   cpp2c.CppMethod
		expected bool
	}{
		{
			name: "virtual method",
			method: cpp2c.CppMethod{
				MethodName: "someVirtual",
				IsVirtual:  true,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.AllowVirtual(tt.method)
			if result != tt.expected {
				t.Errorf("AllowVirtual(%q) = %v, want %v", tt.method.MethodName, result, tt.expected)
			}
		})
	}
}

func TestImGuiFilterAllowVirtualForClass(t *testing.T) {
	filter := NewImGuiFilter()

	tests := []struct {
		name      string
		className string
		expected  bool
	}{
		{"ImGui", "ImGui", true},
		{"ImVec2", "ImVec2", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.AllowVirtualForClass(tt.className)
			if result != tt.expected {
				t.Errorf("AllowVirtualForClass(%q) = %v, want %v", tt.className, result, tt.expected)
			}
		})
	}
}

func TestImGuiFilterAllowCtor(t *testing.T) {
	filter := NewImGuiFilter()

	tests := []struct {
		name      string
		className string
		method    cpp2c.CppMethod
		expected  bool
	}{
		{
			name:      "normal constructor",
			className: "ImVec2",
			method: cpp2c.CppMethod{
				MethodName: "ImVec2",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.AllowCtor(tt.className, tt.method)
			if result != tt.expected {
				t.Errorf("AllowCtor(%q) = %v, want %v", tt.className, result, tt.expected)
			}
		})
	}
}
