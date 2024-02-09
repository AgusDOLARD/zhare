package internal

import (
	"bytes"
	"fmt"
	go_qr "github.com/piglig/go-qr"
)

const (
	blackBlock = "\033[40m  \033[0m"
	whiteBlock = "\033[47m  \033[0m"
)

func GenerateQR(content string) error {
	qr, err := go_qr.EncodeText(content, go_qr.Low)
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	border := 4
	for y := -border; y < qr.GetSize()+border; y++ {
		for x := -border; x < qr.GetSize()+border; x++ {
			if !qr.GetModule(x, y) {
				buf.WriteString(blackBlock)
			} else {
				buf.WriteString(whiteBlock)
			}
		}
		buf.WriteString("\n")
	}
	fmt.Print(buf.String())
	return nil
}
