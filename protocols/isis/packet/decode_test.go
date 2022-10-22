package packet

import (
	"bytes"
)

func convert(lsp []byte) {
	input2 := bytes.NewBuffer(lsp)
	b, _ := DecodeLSPDU(input2)
	println(b)
}
