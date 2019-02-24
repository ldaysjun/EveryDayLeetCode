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

tips：注意p,q是否为空，每次循环开始重置x,y。

### [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters)

1. 设置字串array,结果result
2. 遍历s字符串，将value添加到array形成子串。
3. 遍历array，遍历到len(array)-1，遍历除value的所有字符，判断当前子串是否包含当加入的value。
4. 如果发现有相同的，更新子串array = array[i+1:]，从与value相同字符的下一位开始截取，直接形成新的子串。
5. 判断当前子串长度，如果大于结果result，更新result。
6. 重复进行，知道遍历到尾部。


