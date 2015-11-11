package main
import (
    "fmt"
)

func max_element_position(a []int, size int) (int, int) {
    positionOfMaxElement := 0
    if size < 1 { return -1, -1 }
    for i := 0; i < size; i++ {
        if a[positionOfMaxElement] < a[i] {
            positionOfMaxElement = i
        }
    }
    return positionOfMaxElement, a[positionOfMaxElement]
}

func main() {
    var size, pos, value, i int
    a := make([]int, 10)
    fmt.Scanf("%d", &size)
    for i = 0; i < size; i++ {
        fmt.Scanf("%d", &a[i])
    }
    pos, value = max_element_position(a, size)
    if pos != -1 {
        fmt.Printf("%d\n%d", pos, value);
    }
}