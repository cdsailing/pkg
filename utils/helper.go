package utils

import (
	"encoding/binary"
	"fmt"
	"github.com/cdsailing/pkg/log"
	"os"
	"path/filepath"
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

func GetCurrentDirectory() string {
	//返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
