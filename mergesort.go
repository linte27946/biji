package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	left := mergesort(arr[:middle])
	right := mergesort(arr[middle:])

	return merge(left, right)
}
func merge(left, right []int) []int {
	var ans []int
	var i, j int = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}
	ans = append(ans, left[i:]...)  // 如果left有剩余
	ans = append(ans, right[j:]...) // 如果right有剩余

	return ans
}
func main() {
	var arr []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)
		arr = append(arr, num)
	}
	fmt.Println("原始数组:", arr)

	arr = mergesort(arr)

	fmt.Println("排序后数组:", arr)
}
