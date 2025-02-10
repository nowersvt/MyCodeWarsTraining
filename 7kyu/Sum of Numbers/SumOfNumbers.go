package main

func GetSum(a, b int) int {
    switch true {
    case a == b:
        return a

    case a > b:
        a, b = b, a
    }

    return (b - a + 1) * (a + b) / 2
}