@echo off
title Pranor Sandbox Control Plane
cls

:: Check for Podman
where podman >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Podman is required but not installed or not in PATH.
    pause
    exit /b 1
)

if "%1"=="stop" goto do_stop
if "%1"=="down" goto do_stop
if "%1"=="logs" goto do_logs
if "%1"=="build" goto do_build
goto do_start

:do_stop
echo ==========================================================
echo       STOPPING SERVVERSE ECOSYSTEM SANDBOX
echo ==========================================================
echo [INFO] Tearing down all containers and networks...
podman compose down
echo [INFO] Sandbox stopped.
exit /b 0

:do_logs
echo ==========================================================
echo       SERVVERSE ECOSYSTEM SANDBOX LOGS
echo ==========================================================
podman compose logs
exit /b 0

:do_build
echo ==========================================================
echo       BUILDING & STARTING SERVVERSE ECOSYSTEM SANDBOX
echo ==========================================================
if "%2"=="" (
    echo [INFO] Compiling all 12 services and starting container stack...
    podman compose up -d --build
) else (
    echo [INFO] Compiling service '%2' and starting container stack...
    podman compose up -d --build %2
)
if %errorlevel% neq 0 (
    echo [ERROR] Failed to build/start container stack.
    exit /b 1
)
goto start_generator

:do_start
echo ==========================================================
echo       STARTING SERVVERSE ECOSYSTEM SANDBOX
echo ==========================================================
echo [INFO] Launching container stack...
podman compose up -d
if %errorlevel% neq 0 (
    echo [ERROR] Failed to start container stack.
    exit /b 1
)
goto start_generator

:start_generator
echo.
echo ==========================================================
echo       SANDBOX SERVICES ONLINE
echo ==========================================================
echo  * Pranor Console Dashboard : http://localhost:8083
echo  * Pranor Gate API Gateway  : http://localhost:8080
echo  * Pranor Hub Package  : http://localhost:8088
echo  * Pranor Trace Engine      : http://localhost:8090
echo ==========================================================
echo.
echo [INFO] Launching automated workload generator to simulate live traffic...
echo [INFO] Press Ctrl+C in this terminal window to stop the generator.
echo.

cd scripts
go run load_generator.go

