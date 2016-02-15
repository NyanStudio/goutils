package goutils

import (
    "bytes"
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
        return strings.Repeat(pad, pads) + s
    }

    if pads == padLen {
        return pad + s
    } else if pads < padLen {
        return pad[0:pads] + s
    } else {
        var buffer bytes.Buffer

        for i := 0; i < pads; i++ {
            idx := i % padLen
            buffer.WriteString(pad[idx:idx + 1])
        }

        return buffer.String() + s
    }
}

func (su StringUtils) Reverse(str string) string {
    r := []rune(str)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func (su StringUtils) RightPad(s string, size int, pad string) string {
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
        return s + strings.Repeat(pad, pads)
    }

    if pads == padLen {
        return s + pad
    } else if pads < padLen {
        return s + pad[0:pads]
    } else {
        var buffer bytes.Buffer

        for i := 0; i < pads; i++ {
            idx := i % padLen
            buffer.WriteString(pad[idx:idx + 1])
        }

        return s + buffer.String()
    }
}
