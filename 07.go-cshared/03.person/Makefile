#!MAKEFILE

RM 		= /bin/rm -rf
RMDIR 	= /bin/rmdir

EXP_DIR	= libperson
SRCS 	= person.c
EXEC 	= ${EXP_DIR}/libperson.so
OPTS 	= -Wall -g -shared -fPIC

all : dir ${EXEC}

dir :
	@if [ ! -f ${EXP_DIR} ] ; then mkdir -p ${EXP_DIR}; fi

${EXEC}: 
	gcc -o ${EXEC} ${OPTS} ${SRCS}

clean:
	${RM} ${EXEC}
	${RMDIR} ${EXP_DIR}
