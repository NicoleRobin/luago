package binchunk

import (
	"encoding/binary"
	"math"
)

type reader struct {
	data []byte
}

func (r *reader) readByte() byte {
	b := r.data[0]
	r.data = r.data[1:]
	return b
}

func (r *reader) readBytes(n uint) []byte {
	bytes := r.data[:n]
	r.data = r.data[n:]
	return bytes
}

func (r *reader) readUint32() uint32 {
	i := binary.LittleEndian.Uint32(r.data)
	r.data = r.data[4:]
	return i
}

func (r *reader) readUint64() uint64 {
	i := binary.LittleEndian.Uint64(r.data)
	r.data = r.data[8:]
	return i
}

func (r *reader) readLuaInteger() int64 {
	return int64(r.readUint64())
}

func (r *reader) readLuaNumber() float64 {
	return math.Float64frombits(r.readUint64())
}

func (r *reader) readString() string {
	size := uint(r.readByte())
	if size == 0 {
		// NULL字符串
		return ""
	}
	if size == 0xFF {
		size = uint(r.readUint64())
	}
	bytes := r.readBytes(size - 1)
	return string(bytes)
}

func (r *reader) checkHeader() {
	if string(r.readBytes(4)) != LUA_SIGNATURE {
		panic("not a precompiled chunk! ")
	} else if r.readByte() != LUAC_VERSION {
		panic("version mismatch! ")
	} else if r.readByte() != LUAC_FORMAT {
		panic("format mismatch! ")
	} else if string(r.readBytes(6)) != LUAC_DATA {
		panic("corrupted! ")
	} else if r.readByte() != CINT_SIZE {
		panic("int size mismatch! ")
	} else if r.readByte() != CSIZET_SIZE {
		panic("size_t size mismatch! ")
	} else if r.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch! ")
	} else if r.readByte() != LUA_INTEGER_SIZE {
		panic("lua_integer size mismatch! ")
	} else if r.readByte() != LUA_NUMBER_SIZE {
		panic("lua_number size mismatch! ")
	} else if r.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch! ")
	} else if r.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch! ")
	}
}
