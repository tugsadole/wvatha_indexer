package coder

type Coder interface {
	Encode() interface{}
	Decode() interface{}
}
