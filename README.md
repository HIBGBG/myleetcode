# myleetcode





CCI:[程序员面试金典 - 学习计划 - 力扣（LeetCode）全球极客挚爱的技术成长平台](https://leetcode.cn/studyplan/cracking-the-coding-interview/)

plan :2/day 

0809:  2/109

0814:  expect 12/109  in fact 14

0821:  expect 26/109  

0828:  expect 40/109  in fact 

0904:  expect 54/109  in fact 

0911:  expect 68/109  in fact 

0918:  expect 82/109  in fact 

0925:  expect 96/109  in fact 

1002:  expect 108/109  in fact 



# readme

记录自己做题的过程

1、按数据结构建立文件夹

解法记录如下



array:



42接雨水：

```
分析：有凹槽才能接水；而且是凹下去的地方与低的那一边的差值才算数。
突破点：找到peak。  从两边往peak_index去。
例子：左边有凹槽，右边是斜坡
13254654321
特殊情况：
1234554321
12345
54321

```

75、012排序

```
分析：不能用排序函数，而且题目故意说一趟遍历法，那就是要考虑值得特殊性，只有0、1、2
思路：left  i  right 标志位
=0  left+1  i+1   呼唤位置
=2  right-1   互换位置
=1  i+1
202102010
```

78、子集全排列
列表直接用+ [nums[i]]
插入元素用append

88、归并两个有序数组  m+n 时间复杂度

1、先把nums1后面的空位填上
2、nums2大于等于nums1,就赋nums2的值（相等情况下nums先空 ）
3、循环要先进行一次，while true  .
4、循环内进行判断，一个数组空了之后，就退出
5、退出后，把另一个数组余下的放在nums1前面的位置

90题，有重复数字的排列组合
        [1,2,2] 2*3
        [],[1]    1可取0,1个
        [],[2],[2,2]  2可取 0,1,2个，再排列组合

1、先各数字出现的次数保存到dict
2、遍历dict  v次k的排列是v+1次：tmp=  [] [k] [k,k] [k,k,k]
    res  第一次就等于tmp
    后面的迭代，举例：
        res=[[],[1]]
        tmp=[[],[2],[2,2] k=2,v=2
    注意循环次数：
    for tmp里的每个非空元素ele：[2]或[2,2]
        对于res中的每个元素
            都加上2 或者2,2   res[i]+ele
        得到v倍
    一共是v+1倍


126 单词接龙2

