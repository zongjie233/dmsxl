package main

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"

*/

func reverseLeftWords(s string, n int) string {
	res := []byte(s)
	//反转前n个字符
	reverse(&res, 0, n-1)
	//反转n到end字符
	reverse(&res, n, len(res)-1)
	//整个反转
	reverse(&res, 0, len(res)-1)
	return string(res)
}
