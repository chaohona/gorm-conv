@echo off

set pwd=%~dp0



.\bin\gorm-conv.exe --pb=true --sql=true -I=.\conf\ -O=.\conf\ --cpppath=.\tables\ --codetype="client" --protoversion="3"

del .\tables\gorm_tables.h
del .\tables\gorm_table_field_map_define.cc

.\bin\protoc.exe --proto_path=.\conf\  --cpp_out=.\tables\ .\conf\gorm-db.proto --experimental_allow_proto3_optional
