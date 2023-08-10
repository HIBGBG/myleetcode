package cici01


func replaceSpaces(S string, length int) string {
	// old
	//统计前length位的空格数量
	// cnt:=0
	// for i:=range S{
	//     if i<length&&S[i]==' '{
	//         cnt+=1
	//     }
	// }
	// b:=make([]byte,length+2*cnt)//实际长度
	// // fmt.Println(S[:length],cnt,b)
	// // for i,j:=len(S)-2*cnt-1,len(b)-1;i>=0;i--{//可能过长，S有多余空格
	// for i,j:=length-1,len(b)-1;i>=0;i--{
	//     if S[i]==' '{
	//         b[j]='0'
	//         b[j-1]='2'
	//         b[j-2]='%'
	//         j=j-3
	//     }else{
	//         b[j]=S[i]
	//         j=j-1
	//     }
	// }
	// return string(b)

	//length为总长度，
	//S包含空格   还有末尾空格，不能转化多了

	bt:=[]byte(S)
	r:=len(S)
	for i:=length-1;i>=0;i--{
		if bt[i]==' '{
			bt[r-3],bt[r-2],bt[r-1]='%','2','0'
			r=r-3
		}else{
			bt[r-1]=bt[i]
			r--
		}
	}
	return string(bt[r:])  //S>转换后长度



	//非0一起转化？
}