package binchunk

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

type binaryChunk struct {
	header                  // 头部
	sizeUpvalues byte       // 主函数upvalue数量
	mainFunc     *Prototype // 主函数原型
}

type header struct {
	signature       [4]byte // 签名，也就是常说的魔术Magic Number，主要用于快速识别文件格式
	version         byte    // 版本
	format          byte    // 格式
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

type Prototype struct {
	Source          string        // 源文件名
	LineDefined     uint32        // 函数在源文件中的起始行号
	LastLineDefined uint32        // 函数在源文件中的停止行号
	NumParams       byte          // 固定参数个数
	IsVararg        byte          // 是否是Vararg函数
	MaxStackSize    byte          // 寄存器数量
	Code            []uint32      // 指令表
	Constants       []interface{} // 常量表
	Upvalues        []Upvalue     // Upvalue表
	Protos          []*Prototype  // 子函数原型表
	LineInfo        []uint32      // 行号表
	LocVars         []LocVar      // 局部变量表
	UpvalueNames    []string      // Upvalue名列表
}

type Upvalue struct {
}

type LocVar struct {
}

// 解析二进制chunk
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()        // 校验头部
	reader.readByte()           // 跳过Upvalue数量
	return reader.readProto("") // 读取函数原型
}
