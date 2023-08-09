package cici01

func isUnique(astr string) bool {

	if len(astr)>26 {
		return false
	}

	// map
	// m:=make(map[byte]bool)
	// bt:=[]byte(astr)
	// for _,b:=range bt{
	//     if m[b]==true{
	//         return false
	//     }else{
	//         m[b]=true
	//     }
	// }
	// return true

	// 类似位图 01010
	// bit:=[26]bool{}       //错误1、没有：2、没有{}
	// for _,b:=range astr{
	//     if bit[b-'a']==true{
	//         return false
	//     }else{
	//         bit[b-'a']=true //复制下来还是==
	//     }

	// }
	// return true



	//sort排序，临近比较 不用额外的数据结构？
	// b1:=[]byte(astr)
	// sort.Slice(b1,func(i,j int)bool{return b1[i]<b1[j]})
	// l:=len(astr)
	// for i:=1;i<l;i++{
	//     if b1[i]==b1[i-1]{
	//         return false
	//     }
	// }
	// return true

	//位运算
	num:=0
	for _,v:=range astr{
		bit:=v-'a'
		if num&(1<<bit)!=0{  //1<<bit  a=1 b=10 c=100 d=1000
			return false  //与运算,如果相应位置是1,表示出现过，返回false
		}else{//没有出现过，把这位放上去
			num=num|(1<<bit)//有1就1
		}
	}
	return  true

}
