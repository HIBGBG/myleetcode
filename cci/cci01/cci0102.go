package cci01

import (
	"fmt"
	"sort"
)

func CheckPermutation(s1 string, s2 string) bool {
	//method1
	//统计第一个、统计第二个，各字符次数一样
	//提前结束：长度不一样
	//特殊情况：无
	//提前退出2：遍历第二个的时候，开始扣除，不能扣到负数。
	//if len(s1)!=len(s2){
	//	return false
	//}
	//nums:=map[rune]int{}
	//for _,b:=range s1{
	//	nums[b]+=1
	//}
	//
	//for _,b:=range s2{
	//	nums[b]-=1
	//	if nums[b]<0{
	//		return false
	//	}
	//}
	//
	//return true


	//method2
	//排序后挨个比较
	if len(s1)!=len(s2){
		return false
	}
	a,b:=[]byte(s1),[]byte(s2)
	sort.Slice(a,func(i,j int)bool{return a[i]<a[j]})
	fmt.Println(a)
	sort.Slice(b,func(i,j int)bool{return b[i]<b[j]})
	fmt.Println(b)
	return string(a)==string(b)


}