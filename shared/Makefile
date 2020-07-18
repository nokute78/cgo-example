${CC}=gcc
all:
	go build -buildmode=c-shared -o libshared.so .
	${CC} -L. c/*.c -lshared

clean:
	rm *.so a.out *.h
