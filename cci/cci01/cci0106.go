package cci01

import (
	"fmt"
	"strconv"
)



func compressString(S string) string {
	//直接遍历，最后一位处理

	l:=len(S)
	if l<=1{
		return S
	}
	res:=""
	cnt:=1  //记录连续个数，需重置为1
	for i:=0;i+cnt<l;{//i=i+cnt不能放在这
		if S[i]!=S[i+cnt]{//i+cnt可能越界  。后面循环前一层就要判断
			//输出cnt个s[i]
			//res=append(res,rune(S[i]),rune(cnt))//错误拼接
			res+=fmt.Sprintf("%c%d",S[i],cnt)
			// fmt.Println(res)
			i=i+cnt
			cnt=1
			if i==l-1{//不同到最后的处理
				res+=fmt.Sprintf("%c%d",S[i],cnt)
				// fmt.Println("不同时退出",res)//相同退出 [97 2 98 1 99 5 97 3]cci0106_test.go:21: compressString() = a╗b╔c║a╚, want a2b1c5a3
			}
			// cnt=1//TTTTTUUUUUTTTTTTTTBBBBBBBPPPPPPPPDDDDDDllllllhhhhhhvvvvvvvUQQQQvvvvvvH  -->T5U5T8B7P8D6l6h6v7U1Q4v6H6 .最后应该是H1先重置为1，不然输出最后一位时是前面连续的cnt
		}else{
			cnt++
			if i+cnt==l{//相同到最后,退出的处理
				res+=fmt.Sprintf("%c%d",S[i],cnt)
				// fmt.Println("相同退出",res)
			}
		}
	}
	if len(res)>=len(S){//"bb"  "b2"  没有变短
		return S
	}
	return string(res)
}

func compressString2(S string) string {
	le:=len(S)
	if le<=2{
		return S
	}

	res:=""
	for i :=0;i<le;i++{
		cnt:=1
		for i<le-1&&S[i+1]==S[i]{
			i++
			cnt++
		}
		res=res+string(S[i])+strconv.Itoa(cnt)
	}
	//不用出来再处理
	// if S[le-2]!=S[le-1]{
	//     res=res+string(S[le-1])+strconv.Itoa(cnt)
	// }

	// fmt.Println(res)
	if le<len(res){
		return S
	}else{
		return res
	}
}