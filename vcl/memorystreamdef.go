package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

func (m *TMemoryStream) Read(count int32) (int32, []byte) {
    defer exceptionProc()
    return MemoryStream_Read(m.instance, count)
}

func (m *TMemoryStream) Write(buffer []byte) int32 {
    defer exceptionProc()
    return MemoryStream_Write(m.instance, buffer)
}

