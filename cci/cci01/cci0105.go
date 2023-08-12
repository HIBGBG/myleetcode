package cci01

func oneEditAway(first string, second string) bool {
	//一次编辑距离
	//具体分析
	//	l1=l2  只能进行修改一次，不能添加或者删除
		//l1!=l2  只能长的删除一个，或者是短的增加一次。
	//l1,l2:=len(first),len(second)
	//
	//if l1-l2>=2 || l2-l1>=2{// 没有等号错误，边界值
	//	return false
	//}
	//if l1==l2{
	//	if idx,ok:=findFirstDifferent(first,second);!ok {
	//		//没找到返回false，说明相等
	//		return true
	//	}else{
	//		//找到了，就再看右边相不相等
	//		return first[idx+1:]==second[idx+1:]//会不会越界？
	//	}
	//}else{
	//	if l2>l1{
	//		return oneEditAway(second,first)//把长的放前面
	//	}else{
	//		if idx,ok:=findFirstDifferent(first,second);!ok{//肯定不ok的，不对!也可能ok。场景：长的字符串就多了最后一个字符不一样
	//			return true
	//		}else {
	//			return first[idx+1:] == second[idx:] //这个时候就是，短的不用+1
	//		}
	//	}
	//}



	//直接遍历
	l1,l2:=len(first),len(second)
	if l1-l2>=2 || l2-l1>=2 { // 没有等号错误，边界值
		return false
	}
	cnt:=0//次数，大于2就返回false
	//分相等遍历
	if l1==l2{
		i,j:=0,0
		for i<l1{
			if first[i]!=second[j]{
				cnt++ //失去一次机会
				if cnt>1{
					return false
				}
			}
			i++;j++//同时后移
		}
		return true
	}
	//不等遍历
	if l2>l1{
		return oneEditAway(second,first)//把长的放前面  让l1>l2
	}else{
		i,j:=0,0
		for j<l2{   //i<l1  不严格，可能不一样在多了最后一个，导致l2越界 .改成j<l2
			if first[i]!=second[j]{
				cnt++ //失去一次机会
				if cnt>1{
					return false
				}
				i++//1、长的当删除，j不加，继续匹配
			}else{
				i++;j++//2、正常还是同时移动
			}
			//i++//0：长的删除，j不加错误写法
		}
		return true
	}





}

func findFirstDifferent(s1,s2 string) (int,bool){
	//找到第一个不同的字符的索引位置,，没找到 返回false
	//找到返回true
	for i:=0;i<len(s2);i++{
		if s1[i]!=s2[i]{
			return i,true
		}
	}
	return -1,false//-1不存在
}