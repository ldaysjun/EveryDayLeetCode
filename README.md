# EveryDayLeetCode



### [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

与选择排序原理相同

1. 每一趟循环比较，相加后与target对比。
2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。 
3. 针对所有的元素重复以上的步骤，除了最后一个。
4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

### [2. 两数相加](https://leetcode-cn.com/problems/add-two-numbers/)

1. 设置头指针，设置游标（curry），指向头指针。
2. 设置光标，p,q分别指向两个链表，进行遍历
3. 设置进位（carry），用于相加进位
4. 循环遍历，l1,l2。两节点val值x,y与carry相加
5. 重置进位，游标下移，p,q光标下移
6. 遍历结束，如果carry>0，最后在加一位。

tips：注意p,q是否为空，每次循环开始重置x,y