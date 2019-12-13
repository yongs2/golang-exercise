#!MAKEFILE

RM 		= /bin/rm -rf
RMDIR 	= /bin/rmdir

EXP_DIR	= libgoutil
SRCS 	= goutil.go
EXEC 	= ${EXP_DIR}/goutil.so
OBJS 	= ${EXP_DIR}/goutil.h
OPTS 	= -buildmode=c-shared

all : dir ${EXEC}

dir :
	@if [ ! -f ${EXP_DIR} ] ; then mkdir -p ${EXP_DIR}; fi

${EXEC}: 
	go build -o ${EXEC} ${OPTS} ${SRCS}

clean:
	${RM} ${OBJS}
	${RM} ${EXEC}
	${RMDIR} ${EXP_DIR}
