package cci01

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	//字符串旋转  往后复制一遍
	//如何只检查一次？
	if len(s1)!=len(s2){
		return false
	}

	s1=s1+s1
	return strings.Contains(s1,s2)

}