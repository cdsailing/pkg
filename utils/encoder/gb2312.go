package encoder

import (
	"fmt"
)

func readable(b byte) byte {
	if b <= 9 {
		return b + byte('0')
	}
	return b - 10 + byte('a')
}

func tohex(buf []byte) string {
	if len(buf) == 0 {
		return "[]"
	}

	var (
		i   int
		b   byte
		hex []byte = make([]byte, len(buf)*3+2)
	)

	hex[0] = byte('[')
	i = 1
	for _, b = range buf {
		hex[i] = readable((b & 0xf0) >> 4)
		i++
		hex[i] += readable(b & 0x0f)
		i++
		hex[i] = byte(' ')
		i++
	}
	// overwrite the space
	hex[i-1] = ']'
	hex[i] = ' '

	return string(hex)
}

func nearbygbk(buf []byte, tlen int, plen int) string {
	if tlen <= plen {
		return tohex(buf[0:tlen])
	}
	return tohex(buf[tlen-plen : tlen])
}

func nearbygbks(buf string, tlen int, plen int) string {
	if tlen <= plen {
		return tohex([]byte(buf[0:tlen]))
	}
	return tohex([]byte(buf[tlen-plen : tlen]))
}

// get a slice of utf-8 string, from tlen-plen, to tlen
func nearbyutf8(buf []byte, tlen int, plen int) string {
	if tlen <= plen {
		return tohex(buf) + string(buf)
	}

	start := tlen - plen
	if buf[start] < byte(0x7f) || buf[start] > byte(0xe0) {
		return tohex(buf[start:]) + string(buf[start:])
	}
	start = start - 1
	if buf[start] < byte(0x7f) || buf[start] > byte(0xe0) {
		return tohex(buf[start:]) + string(buf[start:])
	}
	if start < 1 {
		return ""
	}
	start = start - 1
	if buf[start] < byte(0x7f) || buf[start] > byte(0xe0) {
		return tohex(buf[start:]) + string(buf[start:])
	}
	return ""
}

// param: input: input bytes array
// return: output: output bytes array
//         err: error if there are errors when convert
//         ic: input has been converted
//         oc: output has been converted
func ConvertGB2312(input []byte) (output []byte, err error, ic int, oc int) {
	ilen := len(input)
	output = make([]byte, (ilen/2)*3+3)
	olen := 0
	i := 0
	for i < ilen-1 {
		if input[i] <= 0x7f {
			output[olen] = input[i]
			olen++
			i++
		} else {
			gb := int(input[i])<<8 | int(input[i+1])
			u8, ok := gb2312toutf8[gb]
			if !ok {
				err = fmt.Errorf("gb2312 has no character %x, at %d.\nnearby input: %s\nnearby output: %s",
					gb, ilen, nearbygbk(input[0:i], i, 20),
					nearbyutf8(output[0:olen], olen, 30))
				ic = i
				oc = olen
				return
			}
			if u8 >= 0x10000 {
				output[olen] = byte(u8 >> 16)
				olen++
				output[olen] = byte((u8 >> 8) & 0xff)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			} else {
				output[olen] = byte(u8 >> 8)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			}
			i = i + 2
		}
	}

	// the last character
	if i == ilen-1 {
		if byte(input[ilen-1]) <= 0x7f {
			output[olen] = input[ilen-1]
			olen++
			i++
		}
	}

	ilen = i
	output = output[0:olen]
	return output, nil, ilen, olen
}

func ConvertGB2312String(input string) (soutput string, err error, ic int, oc int) {
	ilen := len(input)
	output := make([]byte, (ilen/2)*3+3)
	olen := 0
	i := 0
	for i < ilen-1 {
		bi := byte(input[i])
		if bi <= 0x7f {
			output[olen] = bi
			olen++
			i++
		} else {
			bii := byte(input[i+1])
			gb := int(bi)<<8 | int(bii)
			u8, ok := gb2312toutf8[gb]
			if !ok {
				err = fmt.Errorf("gb2312 has no character %x, at %d\nnearby input: %s\nnearby output: %s",
					gb, ilen, nearbygbks(input[0:i], i, 20),
					nearbyutf8(output[0:olen], olen, 30))
				ic = i
				oc = olen
				return
			}
			if u8 >= 0x10000 {
				output[olen] = byte(u8 >> 16)
				olen++
				output[olen] = byte((u8 >> 8) & 0xff)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			} else {
				output[olen] = byte(u8 >> 8)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			}
			i = i + 2
		}
	}
	if i == ilen-1 {
		if byte(input[ilen-1]) <= 0x7f {
			output[olen] = input[ilen-1]
			olen++
			i++
		}
	}

	output = output[0:olen]
	soutput = string(output)
	return soutput, nil, ilen, olen
}
