package goos

/*
//#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

*/
import "C"
import (
	_"fmt"
)

func Getlogin() string {
	return C.GoString(C.getlogin())
}

func Getpgid(pid int) uint {
	return uint(C.getpgid(C.pid_t(pid)))
}

func Getpgrp() uint {
	return uint(C.getpgrp())
}
