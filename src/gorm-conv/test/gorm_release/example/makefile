# ================================================================
# Makefile for gorm example
#
# Date:   2020-06-29
# ================================================================

GORM_HOME=$(PWD)/../
GORM_THIRD_PART=$(PWD)/thirdpart/

GORM_H=$(wildcard $(GORM_HOME)/include/*.h)
GORM_SRC=$(wildcard $(GORM_HOME)/src/*.cc)
TABLE_H=$(wildcard $(GORM_HOME)/tables/*.h)
TABLE_SRC=$(wildcard $(GORM_HOME)/tables/*.cc)

SHARE_LIB=-lprotobuf -lpthread  -lm -lrt -lssl -lcrypto -ldl

INC =-I$(GORM_HOME)/include/ -I$(GORM_HOME)/tables/ -I$(GORM_THIRD_PART)/protobuff/


.PHONY: all clean 

all:
	g++ -std=c++11 -lstdc++ -g -ggdb -rdynamic -o gorm_test main.cc $(GORM_SRC) $(TABLE_SRC) $(INC) ${SHARE_LIB}

clean:
	rm -f gorm_test
