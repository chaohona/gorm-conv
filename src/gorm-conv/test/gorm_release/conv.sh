#!/bin/bash

./bin/gorm-conv --pb=true --sql=true -I=./conf/ -O=./conf/ --cpppath=./tables/ --codetype="client" --protoversion="3"

rm -rf ./tables/gorm_tables.h
rm -rf ./tables/gorm_table_field_map_define.cc

./bin/protoc --proto_path=./conf/  --cpp_out=./tables/ ./conf/gorm-db.proto --experimental_allow_proto3_optional
