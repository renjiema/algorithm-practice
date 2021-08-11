package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{2, 5, 1, 34, 26, 3, 44, 22}
	Shuffle(arr)
	bubbleSort(arr)
	Shuffle(arr)
	selectionSort(arr)
	Shuffle(arr)
	insertionSort(arr)
	Shuffle(arr)
	shellSort(arr)
	Shuffle(arr)
	fmt.Printf("并归排序结果:%v \n", mergeSort(arr))
	Shuffle(arr)
	quickSort(arr)
}

// 快速排序
func quickSort(arr []int) {
	quickSortImpl(arr, 0, len(arr)-1)
	fmt.Printf("快速排序结果:%v \n", mergeSort(arr))
}

func quickSortImpl(arr []int, left, right int) {
	if left < right {
		mod := partion(arr, left, right)
		quickSortImpl(arr, left, mod-1)
		quickSortImpl(arr, mod+1, right)
	}
}

func partion(arr []int, left, right int) int {
	pivot := arr[left]
	i, j := left+1, right
	for {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] >= pivot {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[j] = arr[j], pivot
	return j
}

// 并归排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mod := len(arr) / 2
	left := arr[:mod]
	right := arr[mod:]
	return merge(mergeSort(left), mergeSort(right))
}
func merge(left, right []int) []int {
	res := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}

// 希尔排序,使用间隔的插入排序
func shellSort(arr []int) {
	for mod := len(arr) / 2; mod > 0; mod /= 2 {
		for i := mod; i < len(arr); i++ {
			j := i - mod
			cur := arr[i]
			for j >= 0 && arr[j] > cur {
				arr[j+mod] = arr[j]
				j -= mod
			}
			arr[j+mod] = cur
		}
	}
	fmt.Printf("希尔排序结果:%v \n", arr)
}

// 插入排序
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		cur := arr[i]
		for j >= 0 && arr[j] > cur {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = cur
	}
	fmt.Printf("插入排序结果:%v \n", arr)
}

// 选择排序:O(n^2)/O(n^2)/O(n^2)/O(1)/不稳定
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Printf("选择排序结果:%v \n", arr)
}

// 冒泡排序:O(n^2)/O(n^2)/O(n)/O(1)/稳定
func bubbleSort(arr []int) {
	flag := true
	for i := 0; i < len(arr)-1; i++ {
		flag = true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				flag = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		// 如果没有发生顺序交换说明已经排好序
		if flag {
			break
		}
	}
	fmt.Printf("冒泡排序结果:%v \n", arr)
}

func Shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())

	length := len(arr)
	for i := 0; i < length; i++ {
		index := rand.Intn(length - i)
		arr[i], arr[index+i] = arr[index+i], arr[i]
	}
	fmt.Printf("打乱后顺序:%v \n", arr)
}
