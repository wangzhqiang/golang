package main

import (
	"io"
	"encoding/binary"
	"bytes"
	"hash/adler32"
)

//编码
var RPC_MAGIC = [4]byte{'p','y','x','i'}

func EncodePacket(w io.Writer, payload []byte) error {
	totalsize := uint32(len(payload) + 8)

	binary.Write(w,binary.BigEndian,totalsize)

	binary.Write(w,binary.BigEndian,RPC_MAGIC)

	w.Write(payload)


	var buf bytes.Buffer
	buf.Write(RPC_MAGIC[:])
	buf.Write(payload)
	checksum := adler32.Checksum(buf.Bytes())
	return binary.Write(w,binary.BigEndian,checksum)

}
