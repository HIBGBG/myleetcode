package cici01

func canPermutePalindrome(s string) bool {


	// 回文规律，两边都是成双的，中间可以有一个是单个的。
	//old
	// 偶数个
	// 奇数个只能有一个

	//cnt:=map[byte]int{}
	//valid:=0
	//for i:=range s{
	//	cnt[s[i]]++
	//	if cnt[s[i]]%2==0{//偶数合法//先减后加，永远为0
	//		valid++
	//	}else{
	//		valid--//奇数个，不合法  ，有多少奇数个元素，就减多少个1,顶多有一个
	//	}
	//}
	//fmt.Println(valid,cnt)
	//if valid>=-1{//允许有一个元素奇数个，放在中间
	//	return true
	//}else{
	//	return false
	//}

	//等价于  统计完个数，只能有一个是偶数的
	//统计
	m:=map[rune]int{}
	for _,i:=range s{
		m[i]++
	}
	//检查
	cnt:=0
	for _,v:=range m{
		//fmt.Println(k,v)
		if v%2!=0{
			cnt++
			if cnt>1{
				//fmt.Println(cnt)
				return false
			}
		}
	}
	return true


}