package cci01



func rotate(matrix [][]int)  {
	//不占用额外空间，原地换，
	l:=len(matrix)

	for i:=0;i<l/2;i++{//圈数  l为奇数圈，中心值不用旋转 ： 3 一圈 4两圈，5两圈
		//移动每条边，边的一头带顶点，另一顶点不用
		for j:=i;j<l-1-i;j++ {   //l=7 ，第一圈边上的点为6，第二圈4  。不是j<l-1-2*i
			// 偶数层 最里层j<l-1-2*i没满足。没转
			a:=matrix[i][j]
			matrix[i][j]=matrix[l-1-j][i]
			matrix[l-1-j][i]=matrix[l-1-i][l-1-j]
			matrix[l-1-i][l-1-j] =matrix[j][l-1-i]
			matrix[j][l-1-i]=a
		}
	}
	// fmt.Println(matrix)
	//return matrix
}


