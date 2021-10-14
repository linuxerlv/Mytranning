func plusOne(digits []int) []int {
    lenth := len(digits) -1 
    
    if len == 0 {  //检测是否是空数组
        return []int{1}
    }

    for i := lenth; i >= 0; i-- {   //逐个入栈
        if(digits[i] < 10) {    // 过滤掉小于10的情况
            digits[i]++
            return digits
        } else {    // 将大于10的数字保留个位，然后进位
            digits[i] %= 10
            digits[i -1]++
        }
    }
    return append([]int{1},digits...)  //为最高位进位后补1
}
