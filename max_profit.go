package solution

func Solution(A []int) int {
    max_profit := 0
    min_day := 200000 //as per the question maximum value
    for i := 0; i < len(A); i++ {
        day := A[i]
        min_day = Min(min_day, day)
        max_profit = Max(max_profit, day - min_day)
    }
    return max_profit
}

func Min(a int, b int) int {
    if a <= b {
        return a
    }
    return b
}

func Max(a int , b int) int {
    if a >= b {
        return a
    }
    return b
}
