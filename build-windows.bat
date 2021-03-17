@echo off

cd app
call npm run build
cd ..

echo Storing assets in bindata.go
go-bindata.exe .\app\dist\build.html

::go generate
::set /p B64html=<.\app\dist\build.html.b64
::go build -ldflags "-H windowsgui -X main.B64html=%B64html%" -o gex.exe
:: go build -ldflags "-X main.B64html=%B64html%" -o gex.exe
go build -ldflags "-H windowsgui" -o gex.exe

echo build complete