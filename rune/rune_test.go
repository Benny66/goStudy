package rune

/*
 * @Author: wenzhicong wenzhicong@jasonanime.com
 * @Date: 2023-01-03 13:52:36
 * @LastEditors: wenzhicong wenzhicong@jasonanime.com
 * @LastEditTime: 2023-01-03 16:26:12
 * @FilePath: /goStudy/rune/rune_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	//rune 是int32的别名，用于存放多字节字符，如占 3 字节的中文字符，返回的是字符 Unicode 码点值
	//byte 是unit8的别名，用于存放占 1 字节的 ASCII 字符，如英文字符，返回的是字符原始字节
	//string 字符串是由字符的组成，字符由字节组成，一个英文为1字节，一个中文为3字节
	str := "Go编程" //2+2*3 = 8字节
	byteLen := len(str)
	fmt.Println("字节长度，", byteLen)
	strRune := []rune(str)
	strLen := len(strRune)
	fmt.Println("字符长度，", strLen)
	fmt.Println("字符截取，", string(strRune[:2]))
}
