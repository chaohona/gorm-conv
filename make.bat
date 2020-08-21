@echo off

set pwd=%~dp0
set GOPATH=%~dp0

if [%1] == [] goto:all

if %1==clean (
call:clean
) else (
call:build %1
)
goto:exit

:all
call:build gorm-conv

exit /b 0

:build
go install %1
if %errorlevel%==0 (
echo build %1 success!
) else (
echo build %1 error!
)
copy /y d:\work\codes\gorm-conv\bin\windows_386\gorm-conv.exe d:\work\codes\gorm\tools\bin\gorm-conv.exe
copy /y d:\work\codes\gorm-conv\bin\windows_386\gorm-conv.exe D:\work\codes\gorm\golang\test\bin\gorm-conv.exe
exit /b 0

:clean
rm pkg/* -rf
rm bin/*.exe -rf
echo clean ok!
exit /b 0

:exit