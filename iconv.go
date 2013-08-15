package iconv

// #cgo darwin LDFLAGS: -liconv
// #include <iconv.h>
// #include <errno.h>
import "C"

import (
	"bytes"
	"syscall"
	"unsafe"
)

const (
	DEFAULT_BUFFER_SIZE = 4096
)

type Iconv struct {
	iconv C.iconv_t
}

func Open(fromcode, tocode string) (*Iconv, error) {
	i, err := C.iconv_open(C.CString(tocode), C.CString(fromcode))
	if err != nil {
		return nil, err
	}
	return &Iconv{i}, nil
}

func (i *Iconv) Close() error {
	_, err := C.iconv_close(i.iconv)
	return err
}

func (i *Iconv) ConvBytes(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}
	var buf bytes.Buffer
	outbuf := make([]byte, DEFAULT_BUFFER_SIZE)
	input_bytes := C.size_t(len(input))
	input_buffer_ptr := &input[0]
	for input_bytes > 0 {
		out_bytes := C.size_t(len(outbuf))
		out_buffer_ptr := &outbuf[0]
		_, err := C.iconv(i.iconv,
			(**C.char)(unsafe.Pointer(&input_buffer_ptr)),
			&input_bytes,
			(**C.char)(unsafe.Pointer(&out_buffer_ptr)),
			&out_bytes)
		write_length := len(outbuf) - int(out_bytes)
		buf.Write(outbuf[:write_length])
		if err != nil && err != syscall.E2BIG {
			return buf.Bytes(), err
		}
	}
	return buf.Bytes(), nil
}

func (i *Iconv) ConvString(input string) (string, error) {
	b, err := i.ConvBytes([]byte(input))
	return string(b), err
}
