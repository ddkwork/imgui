#include <windows.h>
#include <d3d11.h>
#include <cstdio>
#include <cstdlib>
#include <vector>
#include <string>
#include <filesystem>
#include <algorithm>

#define MIQT_TYPES_ONLY
#include "../cabi/gen_imgui_backends.h"

typedef LRESULT (*ImGui_ImplWin32_WndProcHandler_t)(HWND hWnd, UINT msg, WPARAM wParam, LPARAM lParam);

static HMODULE g_dll = nullptr;
static ImGui_ImplWin32_WndProcHandler_t g_WndProcHandler = nullptr;

#define DLL_FUNC(ret, name, ...) typedef ret (*name##_t)(__VA_ARGS__); static name##_t name = nullptr

DLL_FUNC(bool,      cabi_ImGui_ImplWin32_Init, void* hwnd);
DLL_FUNC(void,      cabi_ImGui_ImplWin32_Shutdown);
DLL_FUNC(void,      cabi_ImGui_ImplWin32_NewFrame);
DLL_FUNC(void,      cabi_ImGui_ImplWin32_EnableDpiAwareness);
DLL_FUNC(float,     cabi_ImGui_ImplWin32_GetDpiScaleForMonitor, void* monitor);
DLL_FUNC(bool,      cabi_ImGui_ImplDX11_Init, ID3D11Device* device, ID3D11DeviceContext* device_context);
DLL_FUNC(void,      cabi_ImGui_ImplDX11_Shutdown);
DLL_FUNC(void,      cabi_ImGui_ImplDX11_NewFrame);
DLL_FUNC(void,      cabi_ImGui_ImplDX11_RenderDrawData, ImDrawData* draw_data);
DLL_FUNC(ImGuiContext*, cabi_ImGui__CreateContext, ImFontAtlas* shared_font_atlas);
DLL_FUNC(void,           cabi_ImGui__DestroyContext, ImGuiContext* ctx);
DLL_FUNC(ImGuiIO*,       cabi_ImGui__GetIO);
DLL_FUNC(ImGuiStyle*,    cabi_ImGui__GetStyle);
DLL_FUNC(void,           cabi_ImGui__StyleColorsDark, ImGuiStyle* dst);
DLL_FUNC(void,           cabi_ImGui__NewFrame);
DLL_FUNC(void,           cabi_ImGui__Render);
DLL_FUNC(ImDrawData*,    cabi_ImGui__GetDrawData);
DLL_FUNC(const char*,    cabi_ImGui__GetVersion);
DLL_FUNC(bool,           cabi_ImGui__Begin, const char* name, bool* p_open, int flags);
DLL_FUNC(void,           cabi_ImGui__End);
DLL_FUNC(void,           cabi_ImGui__TextUnformatted, const char* text, const char* text_end);
DLL_FUNC(void,           cabi_ImGui__Separator);
DLL_FUNC(void,           cabi_ImGui__Indent, float indent_w);
DLL_FUNC(void,           cabi_ImGui__Unindent, float indent_w);
DLL_FUNC(void,           cabi_ImGui__PushID_1, const char* str_id);
DLL_FUNC(void,           cabi_ImGui__PopID);
DLL_FUNC(void,           cabi_ImGui__SetNextWindowSize, ImVec2* size, int cond);
DLL_FUNC(void,           cabi_ImGui__SetNextItemWidth, float item_width);
DLL_FUNC(bool,           cabi_ImGui__InputText, const char* label, char* buf, size_t buf_size, int flags, void* callback, void* user_data);
DLL_FUNC(bool,           cabi_ImGui__Button, const char* label, ImVec2* size);
DLL_FUNC(bool,           cabi_ImGui__IsItemHovered, int flags);
DLL_FUNC(bool,           cabi_ImGui__IsMouseDoubleClicked, int button);
DLL_FUNC(void,           cabi_ImGui__SameLine, float offset_from_start_x, float spacing);
DLL_FUNC(ImVec2*,        cabi_ImGui__GetContentRegionAvail);
DLL_FUNC(void,           cabi_ImGuiStyle_ScaleAllSizes, ImGuiStyle* self, float scale_factor);
DLL_FUNC(void,           cabi_ImGuiIO_setFontGlobalScale, ImGuiIO* self, float FontGlobalScale);
DLL_FUNC(bool,           cabi_ImGui__BeginTable, const char* str_id, int columns, int flags, ImVec2* outer_size, float inner_width);
DLL_FUNC(void,           cabi_ImGui__EndTable);
DLL_FUNC(void,           cabi_ImGui__TableNextRow, int flags, float min_row_height);
DLL_FUNC(bool,           cabi_ImGui__TableNextColumn);
DLL_FUNC(void,           cabi_ImGui__TableSetupColumn, const char* label, int flags, float init_width_or_weight, unsigned int user_id);
DLL_FUNC(void,           cabi_ImGui__TableHeadersRow);
DLL_FUNC(bool,           cabi_ImGui__TreeNodeEx, const char* label, int flags);
DLL_FUNC(void,           cabi_ImGui__TreePop);

#undef DLL_FUNC

static bool LoadImGuiDLL() {
    wchar_t dllPath[MAX_PATH];
    GetModuleFileNameW(nullptr, dllPath, MAX_PATH);
    wchar_t* p = wcsrchr(dllPath, L'\\');
    if (p) wcscpy_s(p + 1, MAX_PATH - (p - dllPath), L"imgui_bindings.dll");
    g_dll = LoadLibraryW(dllPath);
    if (!g_dll) {
        GetCurrentDirectoryW(MAX_PATH, dllPath);
        wcscat_s(dllPath, L"\\imgui_bindings.dll");
        g_dll = LoadLibraryW(dllPath);
    }
    if (!g_dll) return false;

#define LOAD(name) do { name = (name##_t)GetProcAddress(g_dll, #name); if (!name) { printf("Failed: " #name "\n"); return false; } } while(0)

    LOAD(cabi_ImGui_ImplWin32_EnableDpiAwareness);
    LOAD(cabi_ImGui_ImplWin32_Init); LOAD(cabi_ImGui_ImplWin32_NewFrame);
    LOAD(cabi_ImGui_ImplWin32_Shutdown);
    LOAD(cabi_ImGui_ImplWin32_GetDpiScaleForMonitor);
    LOAD(cabi_ImGui_ImplDX11_Init); LOAD(cabi_ImGui_ImplDX11_NewFrame);
    LOAD(cabi_ImGui_ImplDX11_Shutdown); LOAD(cabi_ImGui_ImplDX11_RenderDrawData);

    g_WndProcHandler = (ImGui_ImplWin32_WndProcHandler_t)GetProcAddress(g_dll, "cabi_ImGui_ImplWin32_WndProcHandler");

    LOAD(cabi_ImGui__NewFrame); LOAD(cabi_ImGui__Render);
    LOAD(cabi_ImGui__GetDrawData); LOAD(cabi_ImGui__GetVersion);
    LOAD(cabi_ImGui__CreateContext); LOAD(cabi_ImGui__StyleColorsDark); LOAD(cabi_ImGui__DestroyContext);
    LOAD(cabi_ImGui__GetIO); LOAD(cabi_ImGui__GetStyle);
    LOAD(cabi_ImGui__Begin); LOAD(cabi_ImGui__End);
    LOAD(cabi_ImGui__TextUnformatted); LOAD(cabi_ImGui__Separator);
    LOAD(cabi_ImGui__Indent); LOAD(cabi_ImGui__Unindent);
    LOAD(cabi_ImGui__PushID_1); LOAD(cabi_ImGui__PopID);
    LOAD(cabi_ImGui__SetNextWindowSize); LOAD(cabi_ImGui__SetNextItemWidth);
    LOAD(cabi_ImGui__InputText); LOAD(cabi_ImGui__Button);
    LOAD(cabi_ImGui__IsItemHovered); LOAD(cabi_ImGui__IsMouseDoubleClicked);
    LOAD(cabi_ImGui__SameLine); LOAD(cabi_ImGui__GetContentRegionAvail);
    LOAD(cabi_ImGuiStyle_ScaleAllSizes);
    LOAD(cabi_ImGuiIO_setFontGlobalScale);
    LOAD(cabi_ImGui__BeginTable); LOAD(cabi_ImGui__EndTable);
    LOAD(cabi_ImGui__TableNextRow); LOAD(cabi_ImGui__TableNextColumn);
    LOAD(cabi_ImGui__TableSetupColumn); LOAD(cabi_ImGui__TableHeadersRow);
    LOAD(cabi_ImGui__TreeNodeEx); LOAD(cabi_ImGui__TreePop);

#undef LOAD
    return true;
}

static ID3D11Device*           g_pd3dDevice = nullptr;
static ID3D11DeviceContext*    g_pd3dDeviceContext = nullptr;
static IDXGISwapChain*         g_pSwapChain = nullptr;
static bool                    g_SwapChainOccluded = false;
static ID3D11RenderTargetView* g_mainRenderTargetView = nullptr;

bool CreateDeviceD3D(HWND hWnd);
void CleanupDeviceD3D();
LRESULT WINAPI WndProc(HWND hWnd, UINT msg, WPARAM wParam, LPARAM lParam);

struct FileEntry {
    std::string name;
    std::string path;
    uint64_t size;
    bool isDir;
    std::vector<FileEntry> children;
    int depth;
};

static std::vector<FileEntry> g_fileTree;
static char g_pathBuf[256] = "C:\\";

static void PopulateDir(const std::string& dirPath, std::vector<FileEntry>& entries, int depth) {
    if (depth > 3) return;
    std::error_code ec;
    for (const auto& e : std::filesystem::directory_iterator(dirPath, std::filesystem::directory_options::skip_permission_denied, ec)) {
        FileEntry fe;
        fe.name = e.path().filename().string();
        fe.path = e.path().string();
        fe.isDir = e.is_directory(ec);
        fe.depth = depth;
        if (fe.isDir) {
            fe.name += "\\";
            PopulateDir(fe.path, fe.children, depth + 1);
        } else {
            fe.size = e.file_size(ec);
        }
        entries.push_back(fe);
    }
    std::sort(entries.begin(), entries.end(), [](const FileEntry& a, const FileEntry& b) {
        if (a.isDir != b.isDir) return a.isDir;
        std::string an, bn;
        std::transform(a.name.begin(), a.name.end(), std::back_inserter(an), ::tolower);
        std::transform(b.name.begin(), b.name.end(), std::back_inserter(bn), ::tolower);
        return an < bn;
    });
}

static void Refresh() {
    g_fileTree.clear();
    PopulateDir(g_pathBuf, g_fileTree, 0);
}

static void DrawTree(FileEntry& entry) {
    cabi_ImGui__TableNextRow(0, 0);
    cabi_ImGui__TableNextColumn();

    int flags = ImGuiTreeNodeFlags_OpenOnArrow | ImGuiTreeNodeFlags_SpanAllColumns;
    if (!entry.isDir) flags |= ImGuiTreeNodeFlags_Leaf | ImGuiTreeNodeFlags_Bullet;

    bool open = cabi_ImGui__TreeNodeEx(entry.path.c_str(), flags);

    if (open) {
        for (auto& c : entry.children) DrawTree(c);
        cabi_ImGui__TreePop();
    }

    cabi_ImGui__TableNextColumn();
    if (!entry.isDir) {
        char buf[64];
        if (entry.size < 1024) sprintf_s(buf, "%llu B", entry.size);
        else if (entry.size < 1024*1024) sprintf_s(buf, "%.1f KB", entry.size/1024.0f);
        else sprintf_s(buf, "%.1f MB", entry.size/(1024.0*1024.0));
        cabi_ImGui__TextUnformatted(buf, nullptr);
    }

    cabi_ImGui__TableNextColumn();
    cabi_ImGui__TextUnformatted(entry.isDir ? "Folder" : "File", nullptr);
}

int main(int, char**) {
    if (!LoadImGuiDLL()) {
        MessageBoxA(nullptr, "Failed to load imgui_bindings.dll", "Error", MB_ICONERROR);
        return 1;
    }

    printf("ImGui %s loaded from DLL\n", cabi_ImGui__GetVersion());

    cabi_ImGui_ImplWin32_EnableDpiAwareness();

    WNDCLASSEXW wc = { sizeof(wc), CS_CLASSDC, WndProc, 0L, 0L, GetModuleHandle(nullptr), nullptr, nullptr, nullptr, nullptr, L"ImGuiTreeTableDemo", nullptr };
    RegisterClassExW(&wc);
    HWND hwnd = CreateWindowW(wc.lpszClassName, L"ImGui Tree Table (LoadLibrary C ABI)", WS_OVERLAPPEDWINDOW, 100, 100, 1280, 800, nullptr, nullptr, wc.hInstance, nullptr);
    if (!CreateDeviceD3D(hwnd)) { CleanupDeviceD3D(); return 1; }

    ShowWindow(hwnd, SW_SHOWDEFAULT);
    UpdateWindow(hwnd);

    cabi_ImGui__CreateContext(nullptr);
    cabi_ImGui__StyleColorsDark(nullptr);

    float dpiScale = cabi_ImGui_ImplWin32_GetDpiScaleForMonitor(MonitorFromWindow(hwnd, MONITOR_DEFAULTTOPRIMARY));
    cabi_ImGuiStyle_ScaleAllSizes(cabi_ImGui__GetStyle(), dpiScale);
    cabi_ImGuiIO_setFontGlobalScale(cabi_ImGui__GetIO(), dpiScale);

    RECT rcWork;
    SystemParametersInfoW(SPI_GETWORKAREA, 0, &rcWork, 0);
    float winW = (rcWork.right - rcWork.left) * 0.85f;
    float winH = (rcWork.bottom - rcWork.top) * 0.85f;

    cabi_ImGui_ImplWin32_Init(hwnd);
    cabi_ImGui_ImplDX11_Init(g_pd3dDevice, g_pd3dDeviceContext);

    Refresh();

    bool done = false;
    while (!done) {
        MSG msg;
        while (PeekMessage(&msg, nullptr, 0U, 0U, PM_REMOVE)) {
            TranslateMessage(&msg);
            DispatchMessage(&msg);
            if (msg.message == WM_QUIT) done = true;
        }
        if (done) break;
        if (g_SwapChainOccluded && g_pSwapChain->Present(0, DXGI_PRESENT_TEST) == DXGI_STATUS_OCCLUDED) continue;
        g_SwapChainOccluded = false;

        cabi_ImGui_ImplDX11_NewFrame();
        cabi_ImGui_ImplWin32_NewFrame();
        cabi_ImGui__NewFrame();

        {
            ImVec2 winSize = { winW, winH };
            cabi_ImGui__SetNextWindowSize(&winSize, ImGuiCond_FirstUseEver);

            if (cabi_ImGui__Begin("File Explorer (LoadLibrary C ABI)", nullptr, 0)) {
                cabi_ImGui__TextUnformatted("ImGui version loaded from imgui_bindings.dll", nullptr);
                cabi_ImGui__Separator();

                cabi_ImGui__TextUnformatted("Path:  ", nullptr);
                cabi_ImGui__SameLine(0, 0);
                ImVec2* avail = cabi_ImGui__GetContentRegionAvail();
                float availX = avail->x;
                cabi_ImGui__SetNextItemWidth(availX - 60);
                if (cabi_ImGui__InputText("##path", g_pathBuf, sizeof(g_pathBuf), ImGuiInputTextFlags_EnterReturnsTrue, nullptr, nullptr)) Refresh();
                cabi_ImGui__SameLine(0, 0);
                ImVec2 btnSize = { 50, 0 };
                if (cabi_ImGui__Button("Go", &btnSize)) Refresh();

                cabi_ImGui__Separator();

                int tableFlags = ImGuiTableFlags_Resizable | ImGuiTableFlags_RowBg | ImGuiTableFlags_BordersV | ImGuiTableFlags_BordersOuterH | ImGuiTableFlags_ContextMenuInBody | ImGuiTableFlags_ScrollY | ImGuiTableFlags_NoBordersInBodyUntilResize;

                ImVec2 availSize = { 0, cabi_ImGui__GetContentRegionAvail()->y };
                if (cabi_ImGui__BeginTable("TreeTable", 3, tableFlags, &availSize, 0)) {
                    cabi_ImGui__TableSetupColumn("Name", 0, 0, 0);
                    cabi_ImGui__TableSetupColumn("Size", 0, 0, 0);
                    cabi_ImGui__TableSetupColumn("Type", 0, 0, 0);
                    cabi_ImGui__TableHeadersRow();
                    for (auto& e : g_fileTree) DrawTree(e);
                    cabi_ImGui__EndTable();
                }
            }
            cabi_ImGui__End();
        }

        cabi_ImGui__Render();
        g_pd3dDeviceContext->OMSetRenderTargets(1, &g_mainRenderTargetView, nullptr);
        float clear[4] = { 0.11f, 0.11f, 0.13f, 1.00f };
        g_pd3dDeviceContext->ClearRenderTargetView(g_mainRenderTargetView, clear);
        cabi_ImGui_ImplDX11_RenderDrawData(cabi_ImGui__GetDrawData());
        if (g_pSwapChain->Present(0, 0) == DXGI_STATUS_OCCLUDED) g_SwapChainOccluded = true;
    }

    cabi_ImGui_ImplDX11_Shutdown();
    cabi_ImGui_ImplWin32_Shutdown();
    cabi_ImGui__DestroyContext(nullptr);
    CleanupDeviceD3D();
    DestroyWindow(hwnd);
    UnregisterClassW(wc.lpszClassName, wc.hInstance);
    FreeLibrary(g_dll);
    return 0;
}

bool CreateDeviceD3D(HWND hWnd) {
    DXGI_SWAP_CHAIN_DESC sd = {};
    sd.BufferCount = 2;
    sd.BufferDesc.Format = DXGI_FORMAT_R8G8B8A8_UNORM;
    sd.BufferDesc.RefreshRate.Numerator = 60;
    sd.BufferDesc.RefreshRate.Denominator = 1;
    sd.Flags = DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH;
    sd.BufferUsage = DXGI_USAGE_RENDER_TARGET_OUTPUT;
    sd.OutputWindow = hWnd;
    sd.SampleDesc.Count = 1;
    sd.Windowed = TRUE;
    sd.SwapEffect = DXGI_SWAP_EFFECT_DISCARD;

    D3D_FEATURE_LEVEL fl;
    D3D_FEATURE_LEVEL fls[] = { D3D_FEATURE_LEVEL_11_0, D3D_FEATURE_LEVEL_10_0 };
    HRESULT r = D3D11CreateDeviceAndSwapChain(nullptr, D3D_DRIVER_TYPE_HARDWARE, nullptr, 0, fls, 2, D3D11_SDK_VERSION, &sd, &g_pSwapChain, &g_pd3dDevice, &fl, &g_pd3dDeviceContext);
    if (r == DXGI_ERROR_UNSUPPORTED) r = D3D11CreateDeviceAndSwapChain(nullptr, D3D_DRIVER_TYPE_WARP, nullptr, 0, fls, 2, D3D11_SDK_VERSION, &sd, &g_pSwapChain, &g_pd3dDevice, &fl, &g_pd3dDeviceContext);
    if (r != S_OK) return false;
    ID3D11Texture2D* pBack = nullptr;
    g_pSwapChain->GetBuffer(0, IID_PPV_ARGS(&pBack));
    g_pd3dDevice->CreateRenderTargetView(pBack, nullptr, &g_mainRenderTargetView);
    pBack->Release();
    return true;
}

void CleanupDeviceD3D() {
    if (g_mainRenderTargetView) { g_mainRenderTargetView->Release(); g_mainRenderTargetView = nullptr; }
    if (g_pSwapChain) { g_pSwapChain->Release(); g_pSwapChain = nullptr; }
    if (g_pd3dDeviceContext) { g_pd3dDeviceContext->Release(); g_pd3dDeviceContext = nullptr; }
    if (g_pd3dDevice) { g_pd3dDevice->Release(); g_pd3dDevice = nullptr; }
}

LRESULT WINAPI WndProc(HWND hWnd, UINT msg, WPARAM wParam, LPARAM lParam) {
    if (g_WndProcHandler && g_WndProcHandler(hWnd, msg, wParam, lParam)) return true;
    switch (msg) {
    case WM_SIZE:
        if (g_pd3dDevice && wParam != SIZE_MINIMIZED) {
            if (g_mainRenderTargetView) { g_mainRenderTargetView->Release(); g_mainRenderTargetView = nullptr; }
            g_pSwapChain->ResizeBuffers(0, LOWORD(lParam), HIWORD(lParam), DXGI_FORMAT_UNKNOWN, 0);
            ID3D11Texture2D* pBack = nullptr;
            g_pSwapChain->GetBuffer(0, IID_PPV_ARGS(&pBack));
            g_pd3dDevice->CreateRenderTargetView(pBack, nullptr, &g_mainRenderTargetView);
            pBack->Release();
        }
        return 0;
    case WM_DESTROY: PostQuitMessage(0); return 0;
    }
    return DefWindowProcW(hWnd, msg, wParam, lParam);
}
