@echo off

set pwd=%~dp0



%pwd%/gorm-conv.exe --pb=true --sql=true -I=./test/ -O=./test/ --cpppath=./test/ --codetype="client" --protoversion="3"
