package main

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/ddkwork/imgui"
	"golang.org/x/sys/windows"
)

var (
	user32   = windows.NewLazyDLL("user32.dll")
	kernel32 = windows.NewLazyDLL("kernel32.dll")
	d3d11    = windows.NewLazyDLL("d3d11.dll")

	procRegisterClassExW              = user32.NewProc("RegisterClassExW")
	procCreateWindowExW               = user32.NewProc("CreateWindowExW")
	procDefWindowProcW                = user32.NewProc("DefWindowProcW")
	procPeekMessageW                  = user32.NewProc("PeekMessageW")
	procTranslateMessage              = user32.NewProc("TranslateMessage")
	procDispatchMessageW              = user32.NewProc("DispatchMessageW")
	procShowWindow                    = user32.NewProc("ShowWindow")
	procPostQuitMessage               = user32.NewProc("PostQuitMessage")
	procAdjustWindowRect              = user32.NewProc("AdjustWindowRect")
	procGetModuleHandleW              = kernel32.NewProc("GetModuleHandleW")
	procGetTickCount64                = kernel32.NewProc("GetTickCount64")
	procD3D11CreateDeviceAndSwapChain = d3d11.NewProc("D3D11CreateDeviceAndSwapChain")
)

type (
	WNDCLASSEXW struct {
		CbSize        uint32
		Style         uint32
		LpfnWndProc   uintptr
		CbClsExtra    int32
		CbWndExtra    int32
		HInstance     uintptr
		HIcon         uintptr
		HCursor       uintptr
		HbrBackground uintptr
		LpszMenuName  *uint16
		LpszClassName *uint16
		HIconSm       uintptr
	}
	RECT struct{ Left, Top, Right, Bottom int32 }
	MSG  struct {
		Hwnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
	}
)

var (
	IID_IDXGIFactory1   = [16]byte{0x78, 0xae, 0x0a, 0x77, 0x6f, 0xf2, 0xba, 0x4d, 0xa8, 0x29, 0x25, 0x3c, 0x83, 0xd1, 0xb3, 0x87}
	IID_ID3D11Texture2D = [16]byte{0xf2, 0xaa, 0x15, 0x6f, 0x08, 0xd2, 0x89, 0x4e, 0x9a, 0xb4, 0x48, 0x95, 0x35, 0xd3, 0x4f, 0x9c}
)

const (
	DXGI_FORMAT_R8G8B8A8_UNORM             = 28
	DXGI_SWAP_EFFECT_DISCARD               = 0
	DXGI_USAGE_RENDER_TARGET_OUTPUT        = 0x20
	DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH = 0x4
	D3D11_SDK_VERSION                      = 7
	D3D_DRIVER_TYPE_HARDWARE               = 1
	D3D_FEATURE_LEVEL_11_0                 = 0xb000
	D3D_FEATURE_LEVEL_10_1                 = 0xa100
	D3D_FEATURE_LEVEL_10_0                 = 0xa000
	D3D_FEATURE_LEVEL_9_3                  = 0x9300
	WS_OVERLAPPEDWINDOW                    = 0x00090000
	SW_SHOWNORMAL                          = 1
	PM_REMOVE                              = 1
)

var (
	className   = windows.StringToUTF16Ptr("ImGuiDemoWindow")
	windowTitle = windows.StringToUTF16Ptr("ImGui Go Demo - All Widgets")
)

var (
	app                    *imgui.Imgui
	g_pd3dDevice           uintptr
	g_pd3dDeviceContext    uintptr
	g_pSwapChain           uintptr
	g_mainRenderTargetView uintptr
	g_hwnd                 uintptr
	windowClassRegistered  bool
	wndProcPtr             uintptr
)

func callVTable(this uintptr, index int, args ...uintptr) uintptr {
	vtable := *(*uintptr)(unsafe.Pointer(this))
	fn := *(*uintptr)(unsafe.Pointer(vtable + uintptr(index)*8))
	r, _, _ := syscall.SyscallN(fn, append([]uintptr{this}, args...)...)
	return r
}

func wndProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if app != nil {
		if app.CabiImGuiImplWin32WndProcHandler(unsafe.Pointer(hWnd), int32(msg), wParam, int(lParam)) != 0 {
			return 1
		}
	}
	switch msg {
	case 0x0002:
		procPostQuitMessage.Call(0)
		return 0
	case 0x0005:
		if g_pd3dDevice != 0 && wParam != 1 {
			cleanupRenderTarget()
			w := int(lParam & 0xFFFF)
			h := int((lParam >> 16) & 0xFFFF)
			if w > 0 && h > 0 {
				createRenderTarget(w, h)
			}
		}
		return 0
	}
	r, _, _ := procDefWindowProcW.Call(hWnd, uintptr(msg), wParam, lParam)
	return r
}

func strPtr(s string) *int8 {
	if s == "" {
		s = "\x00"
	}
	b := append([]byte(s), 0)
	return (*int8)(unsafe.Pointer(&b[0]))
}

func vec2(x, y float32) *imgui.ImVec2 {
	return &imgui.ImVec2{X: x, Y: y}
}

func createWindow(width, height int32) uintptr {
	if !windowClassRegistered {
		hInstance, _, _ := procGetModuleHandleW.Call(0)
		wndProcPtr = windows.NewCallback(wndProc)
		wc := WNDCLASSEXW{
			CbSize:        uint32(unsafe.Sizeof(WNDCLASSEXW{})),
			Style:         0x0001 | 0x0002,
			LpfnWndProc:   wndProcPtr,
			HInstance:     hInstance,
			HbrBackground: 6,
			LpszClassName: className,
		}
		if ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc))); ret == 0 {
			panic("RegisterClassExW failed")
		}
		windowClassRegistered = true
	}
	r := RECT{Right: width, Bottom: height}
	procAdjustWindowRect.Call(uintptr(unsafe.Pointer(&r)), WS_OVERLAPPEDWINDOW, 0)
	hInstance, _, _ := procGetModuleHandleW.Call(0)
	hwnd, _, _ := procCreateWindowExW.Call(0,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowTitle)),
		WS_OVERLAPPEDWINDOW,
		100, 100, uintptr(r.Right-r.Left), uintptr(r.Bottom-r.Top),
		0, 0, hInstance, 0)
	if hwnd == 0 {
		panic("CreateWindowExW failed")
	}
	procShowWindow.Call(hwnd, SW_SHOWNORMAL)
	return hwnd
}

type DXGI_SWAP_CHAIN_DESC struct {
	BufferDesc struct {
		Width            uint32
		Height           uint32
		RefreshRate      struct{ Numerator, Denominator uint32 }
		Format           uint32
		ScanlineOrdering uint32
		Scaling          uint32
	}
	SampleDesc struct {
		Count   uint32
		Quality uint32
	}
	BufferUsage  uint32
	BufferCount  uint32
	_            uint32
	OutputWindow uintptr
	Windowed     int32
	SwapEffect   uint32
	Flags        uint32
}

func initD3D(hwnd uintptr) {
	levels := []uint32{D3D_FEATURE_LEVEL_11_0, D3D_FEATURE_LEVEL_10_1, D3D_FEATURE_LEVEL_10_0, D3D_FEATURE_LEVEL_9_3}
	var device, context, swapChain uintptr
	var featureLevel uint32

	sd := DXGI_SWAP_CHAIN_DESC{
		BufferDesc: struct {
			Width            uint32
			Height           uint32
			RefreshRate      struct{ Numerator, Denominator uint32 }
			Format           uint32
			ScanlineOrdering uint32
			Scaling          uint32
		}{Format: DXGI_FORMAT_R8G8B8A8_UNORM, RefreshRate: struct{ Numerator, Denominator uint32 }{60, 1}},
		SampleDesc: struct {
			Count   uint32
			Quality uint32
		}{Count: 1},
		BufferUsage:  DXGI_USAGE_RENDER_TARGET_OUTPUT,
		BufferCount:  2,
		OutputWindow: hwnd,
		Windowed:     1,
		SwapEffect:   DXGI_SWAP_EFFECT_DISCARD,
		Flags:        DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH,
	}

	ret, _, _ := procD3D11CreateDeviceAndSwapChain.Call(
		0, D3D_DRIVER_TYPE_HARDWARE, 0,
		0,
		uintptr(unsafe.Pointer(&levels[0])), uintptr(len(levels)),
		D3D11_SDK_VERSION,
		uintptr(unsafe.Pointer(&sd)),
		uintptr(unsafe.Pointer(&swapChain)),
		uintptr(unsafe.Pointer(&device)),
		uintptr(unsafe.Pointer(&featureLevel)),
		uintptr(unsafe.Pointer(&context)),
	)
	if int32(ret) < 0 {
		panic(fmt.Sprintf("D3D11CreateDeviceAndSwapChain failed: 0x%08X", uint32(ret)))
	}
	g_pd3dDevice = device
	g_pd3dDeviceContext = context
	g_pSwapChain = swapChain
	g_hwnd = hwnd
	createRenderTarget(1280, 720)
}

func createRenderTarget(w, h int) {
	var bb uintptr
	ret := callVTable(g_pSwapChain, 9, 0, uintptr(unsafe.Pointer(&IID_ID3D11Texture2D)), uintptr(unsafe.Pointer(&bb)))
	if int32(ret) < 0 {
		return
	}
	var rtv uintptr
	ret = callVTable(g_pd3dDevice, 9, bb, 0, uintptr(unsafe.Pointer(&rtv)))
	callVTable(bb, 2)
	if int32(ret) >= 0 {
		g_mainRenderTargetView = rtv
	}
}

func cleanupRenderTarget() {
	if g_mainRenderTargetView != 0 {
		callVTable(g_mainRenderTargetView, 2)
		g_mainRenderTargetView = 0
	}
}

func cleanupD3D() {
	cleanupRenderTarget()
	if g_pSwapChain != 0 {
		callVTable(g_pSwapChain, 2)
	}
	if g_pd3dDeviceContext != 0 {
		callVTable(g_pd3dDeviceContext, 2)
	}
	if g_pd3dDevice != 0 {
		callVTable(g_pd3dDevice, 2)
	}
}

func initImGui() {
	app = new(imgui.Imgui)
	app.CabiImGuiCreateContext(nil)
	app.CabiImGuiStyleColorsDark(nil)
	app.CabiImGuiImplWin32Init(unsafe.Pointer(g_hwnd))
	app.CabiImGuiImplDX11Init(unsafe.Pointer(g_pd3dDevice), unsafe.Pointer(g_pd3dDeviceContext))
}

func renderFrame() {
	cc := [4]float32{0.1, 0.1, 0.1, 1.0}
	callVTable(g_pd3dDeviceContext, VT_ID3D11DeviceContext_ClearRenderTargetView, g_mainRenderTargetView, uintptr(unsafe.Pointer(&cc)))
	callVTable(g_pd3dDeviceContext, VT_ID3D11DeviceContext_OMSetRenderTargets, 1, uintptr(unsafe.Pointer(&g_mainRenderTargetView)), 0)

	app.CabiImGuiImplDX11NewFrame()
	app.CabiImGuiImplWin32NewFrame()
	app.CabiImGuiNewFrame()

	app.CabiImGuiSetNextWindowSize(vec2(1280, 800), 0)
	app.CabiImGuiSetNextWindowPos(vec2(0, 0), 0, vec2(0, 0))
	app.CabiImGuiBegin(strPtr("ImGui Widget Demo - All Controls"), nil, 0)

	if app.CabiImGuiCollapsingHeader1(strPtr("1. Text & Basic"), 0) {
		app.CabiImGuiTextUnformatted(strPtr("TextUnformatted - plain text display"), nil)
		app.CabiImGuiBullet()
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiTextUnformatted(strPtr("Bullet point with same-line text"), nil)
		app.CabiImGuiSeparator()
		app.CabiImGuiPushStyleColor1(0, -16776961) // 0xFF0000FF
		app.CabiImGuiTextUnformatted(strPtr("Red text"), nil)
		app.CabiImGuiPopStyleColor(1)
		app.CabiImGuiPushStyleColor1(0, -16711936) // 0xFF00FF00
		app.CabiImGuiTextUnformatted(strPtr("Green text"), nil)
		app.CabiImGuiPopStyleColor(1)
		app.CabiImGuiSeparatorText(strPtr("SeparatorText"))
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("2. Buttons"), 0) {
		app.CabiImGuiButton(strPtr("Normal Button"), vec2(0, 0))
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiSmallButton(strPtr("Small"))
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiInvisibleButton(strPtr("invis"), vec2(80, 20), 0)
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiTextUnformatted(strPtr("(invisible)"), nil)
		app.CabiImGuiArrowButton(strPtr("al"), 0)
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiArrowButton(strPtr("ar"), 1)
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiArrowButton(strPtr("au"), 2)
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiArrowButton(strPtr("ad"), 3)
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiTextUnformatted(strPtr("Arrow buttons (L,R,U,D)"), nil)
		app.CabiImGuiPushID1(strPtr("btns"))
		for i := 0; i < 4; i++ {
			if i > 0 {
				app.CabiImGuiSameLine(0, -1)
			}
			app.CabiImGuiButton(strPtr(fmt.Sprintf("Btn%d", i)), vec2(0, 0))
		}
		app.CabiImGuiPopID()
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("3. Checkboxes & Radio"), 0) {
		cb1 := new(bool)
		*cb1 = true
		app.CabiImGuiCheckbox(strPtr("Checkbox 1 (on)"), cb1)
		cb2 := new(bool)
		app.CabiImGuiCheckbox(strPtr("Checkbox 2 (off)"), cb2)
		radio := new(int32)
		app.CabiImGuiRadioButton2(strPtr("Radio A"), radio, 0)
		app.CabiImGuiRadioButton2(strPtr("Radio B"), radio, 1)
		app.CabiImGuiRadioButton2(strPtr("Radio C"), radio, 2)
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("4. Sliders & Drags"), 0) {
		sf := new(float32)
		*sf = 50
		app.CabiImGuiSliderFloat(strPtr("Slider Float"), sf, 0, 100, strPtr("%.2f"), 0)
		si := new(int32)
		*si = 5
		app.CabiImGuiSliderInt(strPtr("Slider Int"), si, 0, 10, strPtr("%d"), 0)
		f3 := []float32{25, 50, 75}
		app.CabiImGuiSliderFloat3(strPtr("Slider Float3"), &f3[0], 0, 100, strPtr("%.1f"), 0)
		df := new(float32)
		*df = 100
		app.CabiImGuiDragFloat(strPtr("Drag Float"), df, 1, 0, 500, strPtr("%.2f"), 0)
		di := new(int32)
		*di = 50
		app.CabiImGuiDragInt(strPtr("Drag Int"), di, 1, 0, 100, strPtr("%d"), 0)
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("5. Input"), 0) {
		ii := new(int32)
		*ii = 42
		app.CabiImGuiInputInt(strPtr("Input Int"), ii, 1, 100, 0)
		inf := new(float32)
		*inf = 3.14
		app.CabiImGuiInputFloat(strPtr("Input Float"), inf, 0.1, 1, strPtr("%.3f"), 0)
		id := new(float64)
		*id = 2.718
		app.CabiImGuiInputDouble(strPtr("Input Double"), id, 0.1, 1, strPtr("%.4f"), 0)
		i2 := []int32{10, 20}
		app.CabiImGuiInputInt2(strPtr("Input Int2"), &i2[0], 0)
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("6. Combo & ListBox"), 0) {
		if app.CabiImGuiBeginCombo(strPtr("Combo"), strPtr("Select..."), 0) {
			for _, item := range []string{"Option A", "Option B", "Option C", "Option D"} {
				app.CabiImGuiSelectable1(strPtr(item), item == "Option A", 0, vec2(0, 0))
			}
			app.CabiImGuiEndCombo()
		}
		if app.CabiImGuiBeginListBox(strPtr("ListBox"), vec2(0, 100)) {
			for _, item := range []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"} {
				app.CabiImGuiSelectable1(strPtr(item), item == "Item 1", 0, vec2(0, 0))
			}
			app.CabiImGuiEndListBox()
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("7. Color Controls"), 0) {
		c4 := []float32{0.5, 0.2, 0.8, 1.0}
		app.CabiImGuiColorEdit4(strPtr("ColorEdit4 RGBA"), &c4[0], 0)
		c3 := []float32{0.2, 0.6, 1.0}
		app.CabiImGuiColorEdit3(strPtr("ColorEdit3 RGB"), &c3[0], 0)
		app.CabiImGuiColorButton(strPtr("col"), (*imgui.ImVec4)(unsafe.Pointer(&c4[0])), 0, vec2(40, 20))
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiTextUnformatted(strPtr("ColorButton preview"), nil)
		app.CabiImGuiColorPicker4(strPtr("Color Picker"), &c4[0], 0, nil)
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("8. Progress Bar"), 0) {
		tick, _, _ := procGetTickCount64.Call()
		p := float32(float64(int64(tick)%3000) / 3000.0)
		app.CabiImGuiProgressBar(p, vec2(-1, 0), strPtr(fmt.Sprintf("%.0f%%", p*100)))
		app.CabiImGuiSeparator()
		app.CabiImGuiProgressBar(0.25, vec2(0, 0), strPtr("25%"))
		app.CabiImGuiProgressBar(0.5, vec2(0, 0), strPtr("50%"))
		app.CabiImGuiProgressBar(0.75, vec2(0, 0), strPtr("75%"))
		app.CabiImGuiProgressBar(1.0, vec2(0, 0), strPtr("100%"))
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("9. Tree Nodes"), 0) {
		if app.CabiImGuiTreeNode(strPtr("Parent Node 1")) {
			app.CabiImGuiTextUnformatted(strPtr("Child content"), nil)
			if app.CabiImGuiTreeNode(strPtr("Nested Child")) {
				app.CabiImGuiTextUnformatted(strPtr("Deeply nested"), nil)
				app.CabiImGuiTreePop()
			}
			app.CabiImGuiTreePop()
		}
		if app.CabiImGuiTreeNodeEx(strPtr("Parent Node 2"), 0) {
			for i := 0; i < 3; i++ {
				app.CabiImGuiTextUnformatted(strPtr(fmt.Sprintf("Item %d", i+1)), nil)
			}
			app.CabiImGuiTreePop()
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("10. Tables"), 0) {
		if app.CabiImGuiBeginTable(strPtr("t1"), 4, 0x10|0x40, vec2(0, 0), 0) {
			app.CabiImGuiTableSetupColumn(strPtr("ID"), 0, 0, 0)
			app.CabiImGuiTableSetupColumn(strPtr("Name"), 0, 0, 0)
			app.CabiImGuiTableSetupColumn(strPtr("Value"), 0, 0, 0)
			app.CabiImGuiTableSetupColumn(strPtr("Status"), 0, 0, 0)
			app.CabiImGuiTableHeadersRow()
			for row := 0; row < 5; row++ {
				app.CabiImGuiTableNextRow(0, 0)
				app.CabiImGuiTableNextColumn()
				app.CabiImGuiTextUnformatted(strPtr(fmt.Sprintf("%d", row+1)), nil)
				app.CabiImGuiTableNextColumn()
				app.CabiImGuiTextUnformatted(strPtr(fmt.Sprintf("Item %c", 'A'+row)), nil)
				app.CabiImGuiTableNextColumn()
				app.CabiImGuiTextUnformatted(strPtr(fmt.Sprintf("%.1f", float32(row+1)*1.5)), nil)
				app.CabiImGuiTableNextColumn()
				app.CabiImGuiTextUnformatted(strPtr([]string{"Active", "Pending", "Done", "Failed", "Skipped"}[row]), nil)
			}
			app.CabiImGuiEndTable()
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("11. Tabs"), 0) {
		if app.CabiImGuiBeginTabBar(strPtr("tb"), 0) {
			if app.CabiImGuiBeginTabItem(strPtr("Tab A"), nil, 0) {
				for i := 0; i < 3; i++ {
					app.CabiImGuiButton(strPtr(fmt.Sprintf("A%d", i+1)), vec2(0, 0))
				}
				app.CabiImGuiEndTabItem()
			}
			if app.CabiImGuiBeginTabItem(strPtr("Tab B"), nil, 0) {
				app.CabiImGuiProgressBar(0.7, vec2(-1, 0), strPtr("70%"))
				app.CabiImGuiEndTabItem()
			}
			if app.CabiImGuiBeginTabItem(strPtr("Tab C"), nil, 0) {
				cb := new(bool)
				*cb = true
				app.CabiImGuiCheckbox(strPtr("Opt1"), cb)
				app.CabiImGuiEndTabItem()
			}
			app.CabiImGuiEndTabBar()
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("12. Menu Bar"), 0) {
		if app.CabiImGuiBeginMenuBar() {
			if app.CabiImGuiBeginMenu(strPtr("File"), true) {
				app.CabiImGuiMenuItem1(strPtr("New"), strPtr("Ctrl+N"), false, true)
				app.CabiImGuiMenuItem1(strPtr("Exit"), strPtr("Alt+F4"), false, true)
				app.CabiImGuiEndMenu()
			}
			if app.CabiImGuiBeginMenu(strPtr("Edit"), true) {
				app.CabiImGuiMenuItem1(strPtr("Undo"), strPtr("Ctrl+Z"), false, true)
				app.CabiImGuiEndMenu()
			}
			app.CabiImGuiEndMenuBar()
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("13. Popups & Tooltips"), 0) {
		if app.CabiImGuiButton(strPtr("Open Modal"), vec2(0, 0)) {
			app.CabiImGuiOpenPopup1(strPtr("my_modal"), 0)
		}
		if app.CabiImGuiBeginPopupModal(strPtr("my_modal"), nil, 0) {
			if app.CabiImGuiButton(strPtr("Close"), vec2(0, 0)) {
				app.CabiImGuiCloseCurrentPopup()
			}
			app.CabiImGuiEndPopup()
		}
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiButton(strPtr("Hover for Tooltip"), vec2(0, 0))
		if app.CabiImGuiIsItemHovered(0) {
			app.CabiImGuiBeginTooltip()
			app.CabiImGuiTextUnformatted(strPtr("This is a tooltip!"), nil)
			app.CabiImGuiEndTooltip()
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("14. Selectables"), 0) {
		for _, item := range []string{"Item 1", "Item 2 (sel)", "Item 3"} {
			app.CabiImGuiSelectable1(strPtr(item), item == "Item 2 (sel)", 0, vec2(0, 0))
		}
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("15. Layout"), 0) {
		for i := 0; i < 5; i++ {
			if i > 0 {
				app.CabiImGuiSameLine(0, -1)
			}
			app.CabiImGuiSmallButton(strPtr(fmt.Sprintf("S%d", i+1)))
		}
		app.CabiImGuiDummy(vec2(0, 5))
		app.CabiImGuiBeginGroup()
		for i := 0; i < 3; i++ {
			app.CabiImGuiButton(strPtr(fmt.Sprintf("G%d", i+1)), vec2(100, 0))
		}
		app.CabiImGuiEndGroup()
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("16. Columns (legacy)"), 0) {
		app.CabiImGuiColumns(3, strPtr("cols"), true)
		app.CabiImGuiTextUnformatted(strPtr("Col1"), nil)
		app.CabiImGuiNextColumn()
		app.CabiImGuiTextUnformatted(strPtr("Col2"), nil)
		app.CabiImGuiNextColumn()
		app.CabiImGuiTextUnformatted(strPtr("Col3"), nil)
		app.CabiImGuiColumns(1, strPtr("cols"), false)
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("17. Disabled Section"), 0) {
		app.CabiImGuiBeginDisabled(true)
		app.CabiImGuiButton(strPtr("Disabled Button"), vec2(0, 0))
		app.CabiImGuiEndDisabled()
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("18. Child Windows"), 0) {
		app.CabiImGuiBeginChild1(strPtr("left"), vec2(200, 100), 0, 0)
		for i := 0; i < 3; i++ {
			app.CabiImGuiButton(strPtr(fmt.Sprintf("L%d", i+1)), vec2(0, 0))
		}
		app.CabiImGuiEndChild()
		app.CabiImGuiSameLine(0, -1)
		app.CabiImGuiBeginChild1(strPtr("right"), vec2(0, 100), 0, 0)
		app.CabiImGuiProgressBar(0.5, vec2(-1, 0), strPtr("50%"))
		app.CabiImGuiEndChild()
	}

	if app.CabiImGuiCollapsingHeader1(strPtr("19. Built-in Demo Window"), 0) {
		show := new(bool)
		*show = true
		app.CabiImGuiShowDemoWindow(show)
	}

	app.CabiImGuiEnd()
	app.CabiImGuiRender()
	app.CabiImGuiImplDX11RenderDrawData(app.CabiImGuiGetDrawData())
	callVTable(g_pSwapChain, 8, 1, 0)
}

func main() {
	runtime.LockOSThread()
	g_hwnd = createWindow(1280, 720)
	initD3D(g_hwnd)
	initImGui()
	var msg MSG
	for {
		for {
			r, _, _ := procPeekMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0, PM_REMOVE)
			if r == 0 {
				break
			}
			if msg.Message == 0x0012 {
				cleanupD3D()
				return
			}
			procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
			procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
		}
		renderFrame()
	}
}
