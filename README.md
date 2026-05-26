# imgui-go

Go bindings for [Dear ImGui](https://github.com/ocornut/imgui) — 纯 Go 实现，无 CGo。

## 特性

- **零 CGo 依赖**：通过 `embed` + `syscall` 加载预编译 DLL，无需 C 编译器
- **D3D11 渲染后端**：纯 Go 调用 COM 接口（vtable 索引由 IDL 自动生成）
- **Win32 平台后端**：通过 syscall 直接调用 Win32 API
- **完整控件覆盖**：按钮、滑条、输入框、组合框、表格、绘图等

## 架构

```
imgui_bindings.dll  ← 预编译的 C ABI 封装层（embed 到 Go 二进制中）
        ↓
dll.go              ← embed DLL + syscall.LazyDLL 动态加载
imgui_backends.go   ← Go 绑定函数（自动生成）
example/
  main.go           ← 示例：D3D11 + Win32 + ImGui 全控件演示
  idl/              ← Windows SDK IDL 文件（d3d11.idl, dxgi.idl）
  idlgen/           ← IDL vtable 索引生成器
  vtable_gen.go     ← 自动生成的 COM vtable 常量
gen/                ← 绑定代码生成器
```

## 无 CGo 原理

1. **ImGui 核心库**：C++ 编译为 `imgui_bindings.dll`，通过 `//go:embed` 嵌入 Go 二进制
2. **DLL 加载**：运行时释放到缓存目录，用 `windows.LazyDLL` 加载
3. **COM 接口调用**：D3D11 等通过 vtable 手动调用，索引从 IDL 文件自动生成
4. **Win32 API**：直接 `syscall` 调用，无需 CGo

整个工具链：`Go 编译 → 单个 .exe`，不需要 MinGW/MSVC/CGo。

## IDL vtable 生成

升级 Windows SDK 时，重新生成 vtable 索引：

```bash
cd example/idlgen
go run . ../idl > ../vtable_gen.go
```

## 使用

```bash
cd example
go run .
```

## 依赖

- Go 1.26+
- Windows (x64)
- `golang.org/x/sys/windows`
