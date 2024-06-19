package middlewares

import (
	"bytes"
	"echo/systemSetting"
	"sync"
)

var bufPool = sync.Pool{New: func() interface{} {
	return new(bytes.Buffer)
}}

func GetBuf() *bytes.Buffer {
	if b, ok := bufPool.Get().(*bytes.Buffer); ok {
		return b
	}
	return new(bytes.Buffer)
}

func putBuf(b *bytes.Buffer) {
	b.Reset()
	bufPool.Put(b)
}

func checkIP(companyId int64, ip string) bool {
	m, has := systemSetting.IpWhitelists[companyId]
	if !has {
		return false
	}
	_, has = m[ip]
	return has
}
