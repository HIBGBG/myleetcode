package cci01
package cci01

func setZeroes(matrix [][]int)  {
	//时间复杂度，至少N*N
	//空间复杂地1、新建一个；2、用两个数组分别记录哪些行，哪些列是0的；3、利用现有数组

	//1、新建一个矩阵，遍历每个点，如果是0，将另一个表这个点所在的行和列，对应新数组置为1。再遍历一遍新矩阵，如果是0，就改成原矩阵的值，如果是1,就改成0
	//最坏每个点都是0 ，每个点都要设置一行一列，N*N*2*N

	//2
	//rows:=map[int]bool{}
	//cols:=map[int]bool{}
	//for i,m :=range matrix{
	//	for j,b:=range m{
	//		if b==0{
	//			rows[i]=true
	//			cols[j]=true
	//		}
	//	}
	//}
	//l:=len(matrix)  //l行
	//m:=len(matrix[0])  //m列
	//for _,r:=range rows{ //这个行
	//	for i := 0; i < m; i++ {
	//		matrix[r][i]=0
	//	}
	//}
	//for _,c:=range cols{//这个列
	//	for i := 0; i < l; i++ {
	//		matrix[i][c]=0
	//	}
	//}


	//3 原位记录，

	firstrow := false
	firstcol  := false
	for i:= range matrix{
		for j,b:=range matrix[i]{
			if b==0{
				if i==0{
					firstrow=true
				}else{
					matrix[i][0]=0  //表示i行后面都可以是0
				}
				if j==0{
					firstcol=true
				}else{
					matrix[0][j]=0  //表示j列下面都可以是0
				}
			}
		}
	}

	//循环第一行 即看每列第一个是不是0，是就更新处理那列
	for i:=1;i<len(matrix[0]);i++{
		if matrix[0][i]==0{
			fillZeroByCol(i,matrix)//更新整列
		}
	}
	//循环第一列 即看每行第一个是不是0，是获取更新处理那行
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0]==0{
			fillZeroByRow(i,matrix)
		}
	}

	//处理首行，处理首列
	if firstrow{
		fillZeroByRow(0,matrix)
	}
	if firstcol{
		fillZeroByCol(0,matrix)
	}

}
func fillZeroByRow(i int,matrix [][]int)  {  //整行填充0
	for j:=range matrix[i]{
		matrix[i][j]=0
	}
}

func fillZeroByCol(i int,matrix [][]int)  { //整列填充0
	for j:=range matrix{
		matrix[j][i]=0
	}
}


//func setZeroes2(matrix [][]int)  {
//	// le:=len(matrix)
//	// co:=len(matrix[0])
//	line_map:=map[int]bool{}
//	column_map:=map[int]bool{}
//	for i:=range matrix{
//		for j:=range matrix[i]{
//			if matrix[i][j]==0{
//				line_map[i]=true
//				column_map[j]=true
//			}
//		}
//	}
//	for i:=range line_map{
//		for j:=range matrix[i]{
//			matrix[i][j]=0
//		}
//	}
//	for j:=range column_map{
//		for i:=range matrix{
//			matrix[i][j]=0
//		}
//	}
//
//}