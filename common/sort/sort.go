package sort

import (
	"fmt"
)

func PrintObj(b Obj) {
	fmt.Printf(("%d "), b.GetValue())
}

func PrintObjList(b []Obj) {
	for i := 0; i < len(b); i++ {
		PrintObj(b[i])
	}
	fmt.Println("")
}

//根据需求修改getValue的返回值类型
type Obj interface {
	GetValue() int
	SetValue(a int)
}

//交换模板
func SwapT(a []Obj, i, j int) {
	a[i], a[j] = a[j], a[i]
}

//交换
func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

//根据需求比较getValue的返回值类型
func IsMax(a, b Obj) bool {
	//根据数据类型需要自定义比较
	if a.GetValue() > b.GetValue() {
		return true
	}
	return false
}

//冒泡
func BubbleSortT(a []Obj) []Obj {
	q := len(a)
	f := true
	for i := 0; i < q-1; i++ { //有多少个数字需要比较 注意不和自己比较所以减一个
		{
			for j := 0; j < q-i-1; j++ { //某个数字需要和别的多少个数字比较 排好序的不用比较
				if IsMax(a[j], a[j+1]) {
					SwapT(a, j, j+1)
					f = false
				}
			}
			if f == true {
				return a
			}
		}
	}
	return a
}

//冒泡
func BubbleSort(a []int) []int {
	q := len(a)
	f := true
	for i := 0; i < q-1; i++ { //有多少个数字需要比较 注意不和自己比较所以减一个
		{
			for j := 0; j < q-i-1; j++ { //某个数字需要和别的多少个数字比较 排好序的不用比较
				if a[j] > a[j+1] {
					Swap(a, j, j+1)
					f = false
				}
			}
			if f == true {
				return a
			}
		}
	}
	return a
}

//选择排序
func SelectionSortT(a []Obj) []Obj {
	l := len(a)
	for i := 0; i < l-1; i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		for j := i + 1; j < l; j++ {
			if IsMax(a[min], a[j]) {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

//选择排序
func SelectionSort(a []int) []int {
	l := len(a)
	for i := 0; i < l-1; i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		for j := i + 1; j < l; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

//插入排序
func InsertionSortT(s []Obj) []Obj {
	n := len(s)
	if n < 2 {
		return s
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && IsMax(s[j-1], s[j]); j-- {
			SwapT(s, j, j-1)
		}
	}
	return s
}

//插入排序
func InsertionSort(s []int) []int {
	n := len(s)
	if n < 2 {
		return s
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			Swap(s, j, j-1)
		}
	}
	return s
}

//快速
func QuickSortT(a []Obj, low, high int) []Obj {
	if low >= high {
		return a
	}
	start := a[low]
	i := low
	for j := low + 1; j <= high; j++ {
		if IsMax(start, a[j]) {
			i++
			if i != j {
				SwapT(a, i, j)
			}
		}
	}
	a[i], a[low] = a[low], a[i]
	QuickSortT(a, low, i-1)
	QuickSortT(a, i+1, high)
	return a
}

//快速
func QuickSort(a []int, low, high int) []int {
	if low >= high {
		return a
	}
	start := a[low]
	i := low
	for j := low + 1; j <= high; j++ {
		if a[j] <= start {
			i++
			if i != j {
				Swap(a, i, j)
			}
		}
	}
	a[i], a[low] = a[low], a[i]
	QuickSort(a, low, i-1)
	QuickSort(a, i+1, high)
	return a
}

//希尔排序
func ShellSortT(a []Obj) []Obj {
	length := len(a)

	gap := 1
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := a[i]
			j := i - gap
			for j >= 0 && IsMax(a[j], temp) {
				a[j+gap] = a[j]
				j -= gap
			}
			a[j+gap] = temp
		}
		//重新设置间隔
		gap = gap / 3
	}
	return a
}

//希尔排序
func ShellSort(a []int) []int {
	length := len(a)

	gap := 1
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := a[i]
			j := i - gap
			for j >= 0 && a[j] > temp {
				a[j+gap] = a[j]
				j -= gap
			}
			a[j+gap] = temp
		}
		//重新设置间隔
		gap = gap / 3
	}
	return a
}

//归并排序
func MergeSortT(a []Obj) []Obj {
	length := len(a)
	if length < 2 {
		return a
	}
	middle := length / 2
	left := a[0:middle]
	right := a[middle:]
	return mergeT(MergeSortT(left), MergeSortT(right))
}

//归并排序
func MergeSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	middle := length / 2
	left := a[0:middle]
	right := a[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

//归并
func mergeT(left []Obj, right []Obj) []Obj {
	var result []Obj
	for len(left) != 0 && len(right) != 0 {
		if IsMax(right[0], left[0]) {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

//归并
func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

//堆排序
func HeapSortT(a []Obj) []Obj {
	arrLen := len(a)
	buildMaxHeapT(a, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		SwapT(a, 0, i)
		arrLen -= 1
		heapifyT(a, 0, arrLen)
	}
	return a
}

//堆排序
func HeapSort(a []int) []int {
	arrLen := len(a)
	buildMaxHeap(a, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		Swap(a, 0, i)
		arrLen -= 1
		heapify(a, 0, arrLen)
	}
	return a
}

//建立大根堆
func buildMaxHeapT(a []Obj, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapifyT(a, i, arrLen)
	}
}

//建立大根堆
func buildMaxHeap(a []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(a, i, arrLen)
	}
}

func heapifyT(a []Obj, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && IsMax(a[left], a[largest]) {
		largest = left
	}
	if right < arrLen && IsMax(a[right], a[largest]) {
		largest = right
	}
	if largest != i {
		SwapT(a, i, largest)
		heapifyT(a, largest, arrLen)
	}
}

func heapify(a []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && a[left] > a[largest] {
		largest = left
	}
	if right < arrLen && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		Swap(a, i, largest)
		heapify(a, largest, arrLen)
	}
}

//计数排序
func CountingSortT(a []Obj, ini Obj) []Obj {
	var maxValue Obj
	maxValue = getMaxInArrT(a)
	bucketLen := maxValue.GetValue() + 1

	//借助这个新增的数组计数
	var bucket []int // 初始为0的数组
	for i := 0; i < bucketLen; i++ {
		bucket = append(bucket, 0)
	}

	sortedIndex := 0
	length := len(a)

	for i := 0; i < length; i++ {
		bucket[a[i].GetValue()] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			a[sortedIndex].SetValue(j)
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return a
}

//计数排序
func CountingSort(a []int) []int {
	var maxValue int
	maxValue = getMaxInArr(a)
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0
	length := len(a)

	for i := 0; i < length; i++ {
		bucket[a[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			a[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return a
}

//桶排序
func SortT(a []Obj) []Obj {
	//桶数
	num := len(a)
	//k（数组最大值）
	max := getMaxInArrT(a)
	//二维切片
	buckets := make([][]Obj, num)
	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = a[i].GetValue() * (num - 1) / max.GetValue() //分配桶index = value * (n-1) /k
		buckets[index] = append(buckets[index], a[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			//桶内可以调用不同排序算法
			InsertionSortT(buckets[i])
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return a
}

//桶排序
func Sort(a []int) []int {
	//桶数
	num := len(a)
	//k（数组最大值）
	max := getMaxInArr(a)
	//二维切片
	buckets := make([][]int, num)
	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max //分配桶index = value * (n-1) /k
		buckets[index] = append(buckets[index], a[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			//桶内可以调用不同排序算法
			InsertionSort(buckets[i])
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return a
}

//基数排序
func RadixSortT(data []Obj) []Obj {
	if len(data) < 2 {
		return data
	}
	max := data[0]
	dataLen := len(data)
	for i := 1; i < dataLen; i++ {
		if IsMax(data[i], max) {
			max = data[i]
		}
	}
	// 计算最大值的位数
	maxDigit := 0
	for max.GetValue() > 0 {
		max.SetValue(max.GetValue() / 10)
		maxDigit++
	}
	// 定义每一轮的除数，1,10,100...
	divisor := 1
	// 定义了10个桶，为了防止每一位都一样所以将每个桶的长度设为最大,与原数组大小相同
	bucket := [10][20]int{{0}}
	// 统计每个桶中实际存放的元素个数
	count := [10]int{0}
	// 获取元素中对应位上的数字，即装入那个桶
	var digit int
	// 经过maxDigit+1次装通操作，排序完成
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < dataLen; j++ {
			tmp := data[j].GetValue()
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp
			count[digit]++
		}
		// 被排序数组的下标
		k := 0
		// 从0到9号桶按照顺序取出
		for b := 0; b < 10; b++ {
			if count[b] == 0 {
				continue
			}
			for c := 0; c < count[b]; c++ {
				data[k].SetValue(bucket[b][c])
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}
	return data
}

//基数排序
func RadixSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	max := data[0]
	dataLen := len(data)
	for i := 1; i < dataLen; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	// 计算最大值的位数
	maxDigit := 0
	for max > 0 {
		max = max / 10
		maxDigit++
	}
	// 定义每一轮的除数，1,10,100...
	divisor := 1
	// 定义了10个桶，为了防止每一位都一样所以将每个桶的长度设为最大,与原数组大小相同
	bucket := [10][20]int{{0}}
	// 统计每个桶中实际存放的元素个数
	count := [10]int{0}
	// 获取元素中对应位上的数字，即装入那个桶
	var digit int
	// 经过maxDigit+1次装通操作，排序完成
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < dataLen; j++ {
			tmp := data[j]
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp
			count[digit]++
		}
		// 被排序数组的下标
		k := 0
		// 从0到9号桶按照顺序取出
		for b := 0; b < 10; b++ {
			if count[b] == 0 {
				continue
			}
			for c := 0; c < count[b]; c++ {
				data[k] = bucket[b][c]
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}
	return data
}

//************************公共函数
//获取最大值
func getMaxInArrT(a []Obj) Obj {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if IsMax(a[i], max) {
			max = a[i]
		}
	}
	return max
}

//获取最大值
func getMaxInArr(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
