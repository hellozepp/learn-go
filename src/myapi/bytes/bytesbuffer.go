package bytes

import (
	"bytes"
	"fmt"
)

func BytesBuffer() {

	b1 := []byte("hello world!")
	buf := bytes.NewBuffer(b1)
	fmt.Printf("buff len=%d\n", buf.Len())
	fmt.Printf("buff cap=%d\n", buf.Cap())

	buf.Grow(100)
	fmt.Printf("buff len=%d\n", buf.Len())
	fmt.Printf("buff cap=%d\n", buf.Cap())

	b2 := make([]byte, 6)
	// Read reads the next len(p) bytes from the buffer or until the buffer is drained. The return value n is the number of bytes read.
	// If the buffer has no data to return, err is io.EOF (unless len(p) is zero); otherwise it is nil.
	buf.Read(b2)
	println(string(b2)) //hello

	b3 := buf.Next(5)
	println(string(b3)) //world

	b4 := buf.Next(3)
	println(string(b4))

	buf2 := bytes.NewBuffer(b1)
	// ReadBytes reads until the first occurrence of delim in the input,
	// returning a slice containing the data up to and including the delimiter.
	b5, _ := buf2.ReadBytes(byte(' '))
	println(len(b5))
	println(string(b5))

	b6 := []byte("go programming")
	buf3 := bytes.NewBuffer(b1)
	// Write appends the contents of p to the buffer, growing the buffer as
	// needed.
	buf3.Write(b6)
	println(string(buf3.Bytes()))
}
