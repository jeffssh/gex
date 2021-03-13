@echo off

cd app
call npm run build
cd ..

echo generating main.go and assets.go
go generate
set /p B64html=<.\app\build\index.html.b64
go build -ldflags "-H windowsgui -X main.B64html=%B64html%" -o gex.exe

echo build complete