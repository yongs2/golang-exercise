# go-cshared library

- [Build modes](https://golang.org/cmd/go/#hdr-Build_modes)

```
-buildmode=c-shared
	Build the listed main package, plus all packages it imports, into a C shared library. 
    The only callable symbols will be those functions exported using a cgo //export comment.
	Requires exactly one main package to be listed.

-buildmode=c-archive
	Build the listed main package, plus all packages it imports, into a C archive file. 
    The only callable symbols will be those functions exported using a cgo //export comment.
    Requires exactly one main package to be listed.
```


## 01.goutil

- dynamic library
- To make libgoutil/goutil.so

```sh
make clean;make
ls -la libgoutil/
```

- example_c

```sh
cd example_c
make
```

## 02.goutil_archive

- static library
- To make libgoutil/libgoutil.a

```sh
make clean;make
ls -la libgoutil/
```

- example_c

```sh
cd example_c
make
```

## 03.Using C Libraries From Go

- [blog post](https://www.thegoldfish.org/2019/04/using-c-libraries-from-go/)
- make so lib

```sh
gcc -o libperson.so -Wall -g -shared -fPIC person.c
```

- make binary

```sh
gcc -o hello -L. -lperson hello.c
```

- run example
```sh
make; make run
make clean;
```
