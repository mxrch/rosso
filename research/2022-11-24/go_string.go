package main

import (
	"fmt"

	"github.com/mxrch/rosso/protobuf"
)

/*
103:func (v Varint) encode(buf []byte, num Number) []byte {
53:func (f Fixed32) encode(buf []byte, num Number) []byte {
62:func (f Fixed64) encode(buf []byte, num Number) []byte {
71:func (r Raw) encode(buf []byte, num Number) []byte {
80:func (s Slice[T]) encode(buf []byte, num Number) []byte {
94:func (s String) encode(buf []byte, num Number) []byte {
*/
func main() {
	m := protobuf.Message{
		1: protobuf.Bytes{'A', 'B'},
	}
	fmt.Printf("%#v\n", m)
}
