package bytes

import (
	"bytes"
)

func writeUInt16(buff []byte, data uint16) {
	for i := 0; i < 2; i++ {
		buff[i] = byte(data >> uint(i*8))
	}
}

func spilt(r rune) bool {
	if r == 'c' {
		return true
	}
	return false
}

func BytesBase() {

	println("hello world")
	buff1 := make([]byte, 2) // 创建一个切片
	writeUInt16(buff1, uint16(12))

	buff2 := make([]byte, 2) // 创建一个切片
	writeUInt16(buff2, uint16(12))

	// 比较两个字节数组切片
	res := bytes.Compare(buff1, buff2)
	println(res)

	// 字符串转换为字节数组
	buff3 := []byte("hello world hello world")
	seq := []byte("hello")

	// Count counts the number of non-overlapping instances of sep in s
	res = bytes.Count(buff3, seq)
	println(res)

	// Contains reports whether subslice is within b
	contains := bytes.Contains(buff3, seq) //true
	println(contains)

	res = bytes.Index(buff3, seq) // 0
	println(res)

	res = bytes.LastIndex(buff3, seq)
	println(res)

	/**
	Rune literals are just an integer value (as you've written).
	They are "mapped" to their unicode codepoint.
	*/
	a := rune('e')
	res = bytes.IndexRune(buff3, a) // -1
	println(res)

	println("------------")

	buff5 := []byte("abcabcabcabc")

	// SplitN 以 sep 为分隔符，将 s 切分成多个子串，结果中不包含 sep 本身
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表
	// 如果 s 中没有 sep，则将整个 s 作为 [][]byte 的第一个元素返回
	// 参数 n 表示最多切分出几个子串，超出的部分将不再切分
	// 如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分
	arr := bytes.SplitN(buff5, []byte("a"), 3)
	for _, v := range arr {
		for _, t := range v {
			print(t)
			print(",")
		}
		println("|")
	}

	// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
	// 空白字符有：\t, \n, \v, \f, \r, ' ', U+0085 (NEL), U+00A0 (NBSP)
	// 如果 s 中只包含空白字符，则返回一个空列表
	println("------------")
	buff6 := []byte("abc abc abc    abc")
	arr = bytes.Fields(buff6)
	for _, v := range arr {
		for _, t := range v {
			print(t)
			print(",")
		}
		println("|")
	}

	// FieldsFunc 以一个或多个连续的满足 f(rune) 的字符为分隔符，
	// 将 s 切分成多个子串，结果中不包含分隔符本身
	// 如果 s 中没有满足 f(rune) 的字符，则返回一个空列表
	println("------------")
	buff7 := []byte("abcabcabcabc")
	arr = bytes.FieldsFunc(buff7, spilt)
	for _, v := range arr {
		for _, t := range v {
			print(t)
			print(",")
		}
		println("|")
	}

	buff8 := []byte("我是中国人")
	// 将 s 切分为 Unicode 码点列表
	data := bytes.Runes(buff8)
	for _, elem := range data {
		println(string(elem))
	}

	// Title 将 s 中的所有单词的首字母修改为其 Title 格式
	buff9 := bytes.Title(buff7)
	println(string(buff9))

	// Map 将 s 中满足 mapping(rune) 的字符替换为 mapping(rune) 的返回值
	// 如果 mapping(rune) 返回负数，则相应的字符将被删除
	buff10 := bytes.Map(func(r rune) rune {
		if r == 'c' {
			return 'a'
		}
		return r
	}, buff7)
	println(string(buff10))
}
