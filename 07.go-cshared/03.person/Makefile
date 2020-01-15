#!MAKEFILE

all : lib bin

lib :
	make -f Makefile_lib.mk

bin :
	make -f Makefile_bin.mk

run:
	make -f Makefile_bin.mk run

clean:
	make -f Makefile_bin.mk clean
	make -f Makefile_lib.mk clean
