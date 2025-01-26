package deepseek

import (
	"bufio"
	"io"
)

type StreamReader struct {
	scanner *bufio.Scanner
	body    io.ReadCloser // 增加对响应体的引用
}
