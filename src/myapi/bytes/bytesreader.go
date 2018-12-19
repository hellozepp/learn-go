package bytes

import (
	"bytes"
	"fmt"
	"io"
)

func BytesReader() {

	b1 := []byte("Hello World!")
	reader := bytes.NewReader(b1)

	// Read 方法
	buff := make([]byte, 5)
	count, err := reader.Read(buff)
	if err != nil {
		return
	}
	fmt.Printf("read count = %d,read data = %s\n", count, string(buff))

	// ReadAt 方法
	buff2 := make([]byte, 5)
	count, err = reader.ReadAt(buff2, 6)
	if err != nil {
		return
	}
	fmt.Printf("read count = %d,read data = %s\n", count, string(buff2))

	for {
		// 依次返回未被读取的字节
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}

	fmt.Println("--------")

	b2 := []byte("hello 世界！")

	reader2 := bytes.NewReader(b2)
	for {
		// 依次返回未被读取的rune
		r, _, err := reader2.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Println(string(r))
	}

	b3 := []byte("string builder")
	// Reset resets the Reader to be reading from b.
	reader2.Reset(b3)
	fmt.Println(reader2.Len())

	fmt.Println("----Seek----")

	reader3 := bytes.NewReader(b1)
	// Seek 设置下一次 Read 或 Write 的偏移量为 offset，它的解释取决于 whence：
	// 0 表示相对于文件的起始处，1 表示相对于当前的偏移，而 2 表示相对于其结尾处。
	// Seek 返回新的偏移量和一个错误，如果有的话。
	abs, err := reader3.Seek(-5, 2)
	fmt.Println(abs) //第7个下标
	b, _ := reader3.ReadByte()
	fmt.Println(string(b)) //倒数第五个
}
