#!MAKEFILE

CC 	= /usr/bin/gcc
CXX	= /usr/bin/g++
LD 	= /usr/bin/ld
CP 	= /bin/cp -f
RM 	= /bin/rm -f

CFLAGS	=	-g -Wall -D_REENTRANT
INCDIR	= 	-I../libgoutil
LFLAGS	= 	-lpthread -lgoutil
LIBDIR	=	-L ../libgoutil	

OBJECT_PATH = .objs

EXEC	= client
SRCS 	= client1.c

.SUFFIXES: .cc .cpp .c

.c.o :
	${CC} -c $< ${CFLAGS} ${INCDIR}
.cc.o :
	${CXX} -c $(CXXFLAGS} ${INCDIR} -o $@ $<
.cpp.o :
	${CXX} -c $(CXXFLAGS} ${INCDIR} -o $@ $<

OBJS	= ${SRCS:%.c=${OBJECT_PATH}/%.o}

all : dir ${EXEC}

dir:
	if [ ! -f ${OBJECT_PATH} ] ; then mkdir -p ${OBJECT_PATH}; fi

${EXEC} : ${OBJS}
	${CC} -o ${EXEC} ${OBJS} ${LIBDIR} ${CFLAGS} ${LFLAGS}

${OBJECT_PATH}/%.o:%.c
	${CC} -c $< ${CFLAGS} ${INCDIR} -o $@

clean : 
	${RM} ${OBJS}
	${RM} -R ${OBJECT_PATH}
	${RM} ${EXEC} core
