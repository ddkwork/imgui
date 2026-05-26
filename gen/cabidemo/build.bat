@echo off
cd /d "%~dp0"
if exist build rmdir /s /q build
cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release . 2>&1
cmake --build build --config Release 2>&1 | powershell -Command "$input | Tee-Object -FilePath build\error.log"
