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
	// Read从缓冲区读取下一个len(p)字节，或者直到缓冲区耗尽。返回值n是读取的字节数。
	//如果缓冲区没有数据返回，err是io。EOF(除非len(p)为零);否则就是nil。
	hasRead, _ := buf.Read(b2)
	fmt.Println(hasRead, string(b2)) //hello

	b3 := buf.Next(5)
	fmt.Println(string(b3)) //world

	b4 := buf.Next(3)
	fmt.Println(string(b4))

	buf2 := bytes.NewBuffer(b1)
	// ReadBytes reads until the first occurrence of delim in the input,
	// returning a slice containing the data up to and including the delimiter.
	//读取直到输入中第一次出现delim，
	//返回一个包含分隔符之前和包括分隔符的数据片。
	b5, _ := buf2.ReadBytes(byte(' '))
	fmt.Println(len(b5))
	fmt.Println("string(b5):", string(b5))

	b6 := []byte("go programming")
	buf3 := bytes.NewBuffer(b1)
	// Write appends the contents of p to the buffer, growing the buffer as
	// needed.
	buf3.Write(b6)
	fmt.Println(string(buf3.Bytes()))
}
