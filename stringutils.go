package goutils

import (
	"bytes"
	"fmt"
	"strings"
)

const PadLimit int = 8192
const SPACE string = " "

type StringUtils struct {

}

func (su StringUtils) LeftPad(s string, size int, pad string) string {
	if s == "" {
		return s
	}

	if pad == "" {
		pad = SPACE
	}

	padLen := len(pad)
	sLen := len(s)
	pads := size - sLen

	if pads <= 0 {
		// returns original String when possible
		return s
	}
	if padLen == 1 && pads <= PadLimit {
		fmt.Println("'1.'")
		return strings.Repeat(pad, pads) + s
	}

	if pads == padLen {
		fmt.Println("'2.'")
		return pad + s
	} else if pads < padLen {
		fmt.Println("'3.'")
		return pad[0:pads] + s
	} else {
		fmt.Println("'4.'")
		var buffer bytes.Buffer

		for i := 0; i < pads; i++ {
			idx := i % padLen
			buffer.WriteString(pad[idx:idx + 1])
		}

		return buffer.String() + s
	}
}