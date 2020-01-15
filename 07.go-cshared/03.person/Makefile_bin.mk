#!MAKEFILE

RM 		= /bin/rm -rf
RMDIR 	= /bin/rmdir

SRCS 	= hello.c
EXEC 	= hello
OPTS 	= -L. -lperson

all : ${EXEC}

${EXEC}: 
	gcc -o ${EXEC} ${OPTS} ${SRCS}
	@echo "export LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:."

run : export LD_LIBRARY_PATH=.
run :
	./${EXEC}

clean:
	${RM} ${EXEC}
