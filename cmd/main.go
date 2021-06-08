package main

import (
	"fmt"
	"log"

	"github.com/tugsadole/wvatha_indexer/pkg/config"
)

const (
	A = 0x1
	B = 0x2
	C = 0x3
	D = 0x4
	E = 0x5
	F = 0x6
	G = 0x7
	H = 0x8
	I = 0x9
	J = 0xa
	K = 0xb
	L = 0xc
	M = 0xd
	N = 0xe
	O = 0xf
	P = 0x10
	Q = 0x11
	R = 0x12
	S = 0x13
	T = 0x14
	U = 0x15
	V = 0x16
	W = 0x18
	X = 0x19
	Y = 0x1a
	Z = 0x1b
)

const (
	OneBitOn   = 0x1
	TwoBitOn   = 0x3
	ThreeBitOn = 0x7
	FourBitOn  = 0xf
	FiveBitOn  = 0x1f
)

// 00000000-00000000-00000000-00000000

// 000000-000000-000000-000000-000000-00

// 00000-00000-00000-00000-00000-00000-00

func main() {

	var seq uint32 = (0x1f & A) << 0
	seq |= (0x1f & B) << 5
	seq |= (0x1f & C) << 10
	seq |= (0x1f & D) << 15
	seq |= (0x1f & E) << 20
	seq |= (0x1f & F) << 25

	println(seq)

	//ret := crawler.Add(2, 3)
	conf, err := config.GetAppConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(conf)
}
