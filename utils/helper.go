package utils

import (
	"encoding/binary"
	"fmt"
	"github.com/cdsailing/pkg/log"
	"strconv"
	"strings"
	"unsafe"
)

func ToHex(source int, size int) string {
	hexStr := strconv.FormatInt(int64(source), 16)
	hexStr = fmt.Sprintf("%0*s", size*2, hexStr)
	return hexStr
}

func Log(source string) {
	log.Info(source)
}

func ToSingle(buffer []byte) float32 {
	m := uintptr(binary.LittleEndian.Uint32(buffer))
	result := *(*float32)(unsafe.Pointer(&m))
	return result
}

func Join(bytes []byte, sep string) string {
	result := []string{}
	for _, item := range bytes {
		result = append(result, fmt.Sprintf("%v", item))
	}
	return strings.Join(result, sep)
}
