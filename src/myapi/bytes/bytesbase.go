package bytes

import (
	"bytes"
	"fmt"
)

func writeUInt16(buff []byte, data uint16) {
	fmt.Println(buff, data)
	for i := 0; i < 2; i++ {
		buff[i] = byte(data >> uint(i*8))
		fmt.Println(buff[i])
	}
}

func spilt(r rune) bool {
	if r == 'c' {
		return true
	}
	return false
}

func BytesBase() {

	buff1 := make([]byte, 2) // 创建一个切片
	fmt.Printf("%08b\n", 12)
	writeUInt16(buff1, uint16(12))

	buff2 := make([]byte, 3) // 创建一个切片
	writeUInt16(buff2, uint16(12))

	res := bytes.Compare(buff1, buff2)
	fmt.Println(" 比较两个字节数组切片:", res) //-1 错误
	res = bytes.Compare(buff1, []byte{0, 0})
	fmt.Println(" 比较两个字节数组切片:", res) //1 不一样
	res = bytes.Compare(buff1, []byte{12, 0})
	fmt.Println(" 比较两个字节数组切片:", res) //0

	// 字符串转换为字节数组
	buff3 := []byte("hello world hello world")
	seq := []byte("hello")

	// Count counts the number of non-overlapping instances of sep in s 不重复子序列
	res = bytes.Count(buff3, seq)
	fmt.Println("Count:", res)

	// Contains reports whether subslice is within b
	contains := bytes.Contains(buff3, seq) //true
	fmt.Println(contains)

	res = bytes.Index(buff3, seq) // 0
	fmt.Println(res)

	res = bytes.LastIndex(buff3, seq) //12
	fmt.Println(res)

	/**
		Rune literals are just an integer value (as you've written).
		They are "mapped" to their unicode codepoint.
		符文只是一个整数值(正如您所写的)。
	它们被“映射”到它们的unicode码点。
	*/
	a := rune('e')
	res = bytes.IndexRune(buff3, a) // -1
	fmt.Println(res)

	fmt.Println("------------")

	buff5 := []byte("abcabcabcabc")

	// SplitN 以 sep 为分隔符，将 s 切分成多个子串，结果中不包含 sep 本身
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表
	// 如果 s 中没有 sep，则将整个 s 作为 [][]byte 的第一个元素返回
	// 参数 n 表示最多切分出几个子串，超出的部分将不再切分
	// 如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分
	arr := bytes.SplitN(buff5, []byte("a"), 3)
	for _, v := range arr {
		for index, t := range v {
			fmt.Print(t, "---->", index)
			fmt.Print(",")
		}
		fmt.Println("|")
	}

	// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
	// 空白字符有：\t, \n, \v, \f, \r, ' ', U+0085 (NEL), U+00A0 (NBSP)
	// 如果 s 中只包含空白字符，则返回一个空列表
	fmt.Println("------------")
	buff6 := []byte("abc abc abc    abc")
	arr = bytes.Fields(buff6)
	for _, v := range arr {
		for _, t := range v {
			fmt.Print(t)
			fmt.Print(",")
		}
		fmt.Println("|")
	}

	// FieldsFunc 以一个或多个连续的满足 f(rune) 的字符为分隔符，
	// 将 s 切分成多个子串，结果中不包含分隔符本身
	// 如果 s 中没有满足 f(rune) 的字符，则返回一个空列表
	fmt.Println("------------")
	buff7 := []byte("abcabcabcabc")
	arr = bytes.FieldsFunc(buff7, spilt)
	for _, v := range arr {
		for _, t := range v {
			fmt.Print(t)
			fmt.Print(",")
		}
		fmt.Println("|")
	}

	buff8 := []byte("我是中国人")
	// 将 s 切分为 Unicode 码点列表
	data := bytes.Runes(buff8)
	for _, elem := range data {
		fmt.Println(string(elem))
	}

	// Title 将 s 中的所有单词的首字母修改为其 Title 格式
	buff9 := bytes.Title(buff7)
	fmt.Println(string(buff9))

	// Map 将 s 中满足 mapping(rune) 的字符替换为 mapping(rune) 的返回值
	// 如果 mapping(rune) 返回负数，则相应的字符将被删除
	buff10 := bytes.Map(func(r rune) rune {
		if r == 'c' {
			return 'a'
		}
		return r
	}, buff7)
	fmt.Println(string(buff10))
}
