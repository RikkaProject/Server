package proto

import (
	"encoding/binary"
	"fmt"
	"math"
)

/*
Buffer []byte读写类

	Byte []byte数据
	pos 暂时没用
	offset 读写偏移
	p Proto类
*/
type Buffer struct {
	Byte   []byte
	pos    int
	offset int
	p      *Proto
}

func (bb *Buffer) ResetOffset() {
	bb.offset = 0
}

func NewBuffer(size int, proto *Proto) *Buffer {
	b := make([]byte, size)
	return &Buffer{Byte: b, p: proto}
}

func NewBufferFrom(b []byte) *Buffer {
	return &Buffer{Byte: b}
}

func (bb *Buffer) ReadInt() int {
	value := binary.LittleEndian.Uint32(bb.Byte[bb.offset : bb.offset+4])
	bb.offset += 4
	return int(value)
}

func (bb *Buffer) ReadString() string {
	length := bb.ReadInt()
	str := string(bb.Byte[bb.offset : bb.offset+length])
	bb.offset += length
	return str
}

func (bb *Buffer) ReadByte() byte {
	value := bb.Byte[bb.offset : bb.offset+1]
	bb.offset += 1
	return value[0]
}

func (bb *Buffer) ReadUint16() uint16 {
	value := binary.LittleEndian.Uint16(bb.Byte[bb.offset : bb.offset+2])
	bb.offset += 2
	return value
}

func (bb *Buffer) ReadLong() uint64 {
	value := binary.LittleEndian.Uint64(bb.Byte[bb.offset : bb.offset+8])
	bb.offset += 8
	return value
}

// ReadFloat 从[]byte中读取float32
func (bb *Buffer) ReadFloat(buf []byte) float32 {
	bits := binary.LittleEndian.Uint32(bb.Byte[bb.offset : bb.offset+4])
	bb.offset += 4
	return math.Float32frombits(bits)
}

// ----------------------------------- 读写分割线 -----------------------------------

func (bb *Buffer) PutInt(value int) {
	binary.LittleEndian.PutUint32(bb.Byte[bb.offset:], uint32(value))
	bb.offset += 4
}

func (bb *Buffer) PutByte(value byte) {
	bb.Byte[bb.offset] = value
	bb.offset += 1
}

func (bb *Buffer) PutUint16(value int) {
	binary.LittleEndian.PutUint16(bb.Byte[bb.offset:], uint16(value))
	bb.offset += 2
}

func (bb *Buffer) PutLong(value int) {
	binary.LittleEndian.PutUint64(bb.Byte[bb.offset:], uint64(value))
	bb.offset += 8
}

func (bb *Buffer) PutString(value string) {
	length := len(value)
	bb.PutInt(length)
	copy(bb.Byte[bb.offset:], value)
	bb.offset += length
}

// PutFloat 将float32转换为字节序列并存储到[]byte中
func (bb *Buffer) PutFloat(val float32) {
	bits := math.Float32bits(val)
	binary.LittleEndian.PutUint32(bb.Byte[bb.offset:], bits)
	bb.offset += 4
}

func (bb *Buffer) Finish(crc bool) {
	if crc {
		newBuf := make([]byte, len(bb.Byte))
		copy(newBuf[6:], bb.Byte[:bb.offset])
		bb.offset += 6
		length := bb.offset
		binary.LittleEndian.PutUint16(newBuf[0:2], uint16(length))
		binary.LittleEndian.PutUint32(newBuf[2:6], bb.p.Key)
		fmt.Println(bb.p.Key)
		bb.Byte = newBuf[:bb.offset]
		bb.p.Key += 1
	} else {
		newBuf := make([]byte, len(bb.Byte))
		copy(newBuf[2:], bb.Byte[:bb.offset])
		bb.offset += 2
		length := bb.offset
		binary.LittleEndian.PutUint16(newBuf[0:2], uint16(length))
		bb.Byte = newBuf[:bb.offset]
	}
}
