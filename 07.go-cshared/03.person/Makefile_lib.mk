#!MAKEFILE

RM 		= /bin/rm -rf
RMDIR 	= /bin/rmdir

SRCS 	= person.c
EXEC 	= libperson.so
OPTS 	= -Wall -g -shared -fPIC

all : ${EXEC}

${EXEC}: 
	gcc -o ${EXEC} ${OPTS} ${SRCS}

clean:
	${RM} ${EXEC}
