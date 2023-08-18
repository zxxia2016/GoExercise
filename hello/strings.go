package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main7() {
	str := "hello world"
	fmt.Println("length:", len(str))

	has := strings.HasPrefix(str, "he")
	fmt.Println("HasPrefix:", has)

	has = strings.HasSuffix(str, "d")
	fmt.Println("HasSuffix:", has)

	contain := strings.Contains(str, "world")
	fmt.Println("Contains:", contain)

	index := strings.Index(str, "e")
	fmt.Println("Index:", index)

	lastIndex := strings.LastIndex(str, "o")
	fmt.Println("LastIndex:", lastIndex)

	repeat := strings.Repeat(str, 2)
	fmt.Println("repeat:", repeat)

	var str1 = "test"
	str1 = strings.ToUpper(str1)
	fmt.Println("ToUpper", str1)

	str1 = strings.ToLower(str1)
	fmt.Println("ToLower", str1)

	str1 = " 111 "
	str1 = strings.TrimSpace(str1)
	fmt.Println("TrimSpace", str1)

	str1 = "1112222"
	str1 = strings.TrimLeft(str1, "111")
	fmt.Println("TrimLeft", str1)

	str1 = "1112222"
	count := strings.Count(str1, "2")
	fmt.Println("Count", count)

	str = "The quick brown fox jumps over the lazy dog"
	strFields := strings.Fields(str)
	fmt.Println("Fields", strFields)

	str = strings.Join(strFields, ",")
	fmt.Println("Join", str)

	Reader := strings.NewReader(str)
	b, e := Reader.ReadByte()
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(b)
	}

	fmt.Println("strconv init size: ", strconv.IntSize)
	an, e := strconv.Atoi("eee")
	if e != nil {
		fmt.Println("Atoi", e.Error())
	} else {
		fmt.Println("666", an)
	}
	fmt.Println(strconv.Itoa(1))

}
