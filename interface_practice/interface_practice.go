package interface_practice

import (
	"io"
)

type LimitedReader struct {
	reader io.Reader
	//  запоминаем количество Считанных байт
	left int
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{reader: r, left: n}
}

func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(p) {
		p = p[0:r.left]
	}
	n, err := r.reader.Read(p)
	r.left -= n
	return n, err
}
