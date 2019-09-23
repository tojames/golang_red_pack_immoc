package main

import "unsafe"

type HelloChan struct {
	qcount   uint
	dataqsiz uint
	buf      unsafe.Pointer
	elemsize uint16
	closed   uint32
	elemtype *_type
	sendx    uint
	recvx    uint
	recvq    waitq
	sendq    waitq

	lock mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}

func main() {

}
