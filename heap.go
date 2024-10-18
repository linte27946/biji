package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 堆排序主函数
func heapSort(arr []int) {
	n := len(arr)

	// 建堆：从最后一个非叶子节点开始向上调整堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一次取出堆顶元素，将其与数组末尾元素交换
	for i := n - 1; i > 0; i-- {
		// 将堆顶元素和当前末尾元素交换
		arr[0], arr[i] = arr[i], arr[0]

		// 重新调整堆
		heapify(arr, i, 0)
	}
}

// 调整堆的函数：确保以 i 为根节点的子树为最大堆
func heapify(arr []int, n int, i int) {
	largest := i     // 假设当前节点是最大的
	left := 2*i + 1  // 左子节点索引
	right := 2*i + 2 // 右子节点索引

	// 如果左子节点存在且大于根节点，则更新最大节点
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在且大于当前最大节点，则更新最大节点
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大节点不是根节点，则交换，并递归调整
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换

		// 递归调整子树
		heapify(arr, n, largest)
	}
}

func main() {
	var arr []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		str := scanner.Text()
		number, _ := strconv.Atoi(str)
		arr = append(arr, number)
	}

	fmt.Println("原始数组:", arr)

	heapSort(arr)

	fmt.Println("排序后数组:", arr)
}
