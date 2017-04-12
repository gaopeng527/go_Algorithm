// Algorithm project Algorithm.go
package Algorithm

// 冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	for i := n; i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// 选择排序
func SelectSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[k] {
				k = j
			}
		}
		a[i], a[k] = a[k], a[i]
	}
}

// 插入排序
func InsertionSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		temp := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > temp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = temp
	}
}

// 希尔排序
func ShellSort(a []int) {
	n := len(a)
	for d := n / 2; d >= 1; d /= 2 {
		for i := d; i < n; i++ {
			temp := a[i]
			j := i - d
			for ; j >= 0 && a[j] > temp; j -= d {
				a[j+d] = a[j]
			}
			a[j+d] = temp
		}
	}
}

// 快速排序的一次划分
func partition(a []int, s int, e int) int {
	temp := a[s]
	i := s
	j := e
	for i < j {
		for i < j && a[j] > temp {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] < temp {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = temp
	return i
}

// 快速排序
func QuickSort(a []int, s int, e int) {
	if s >= e {
		return
	}
	i := partition(a, s, e)
	QuickSort(a, s, i-1)
	QuickSort(a, i+1, e)
}

// 堆排序
func HeapSort(a []int) {
	n := len(a)
	// 建堆
	for i := n/2 - 1; i >= 0; i-- {
		k := i
		for 2*k+1 < n {
			j := 2*k + 1
			if j+1 < n && a[j] < a[j+1] {
				j++
			}
			if a[j] > a[k] {
				a[k], a[j] = a[j], a[k]
				k = j
			} else {
				break
			}
		}
	}
	// 调整堆
	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		k := 0
		for 2*k+1 < i {
			j := 2*k + 1
			if j+1 < i && a[j] < a[j+1] {
				j++
			}
			if a[j] > a[k] {
				a[k], a[j] = a[j], a[k]
				k = j
			} else {
				break
			}
		}
	}
}

// 合并一次
func mergeOne(a []int, b []int, n int, len int) {
	i := 0
	for i+len < n {
		j := i + 2*len - 1
		if j >= n {
			j = n - 1
		}
		m := i
		k := i
		l := i + len
		for i < k+len && l <= j {
			if a[i] <= a[l] {
				b[m] = a[i]
				m++
				i++
			} else {
				b[m] = a[l]
				m++
				l++
			}
		}
		for i < k+len {
			b[m] = a[i]
			m++
			i++
		}
		for l <= j {
			b[m] = a[l]
			m++
			l++
		}
		i = j + 1
	}
	if i < n {
		for ; i < n; i++ {
			b[i] = a[i]
		}
	}
}

// 归并排序
func MergeSort(a []int) {
	n := len(a)
	b := make([]int, n)
	len := 1
	flag := 0
	for len < n {
		if flag == 0 {
			mergeOne(a, b, n, len)
		}
		if flag == 1 {
			mergeOne(b, a, n, len)
		}
		flag = 1 - flag
		len *= 2
	}
	if flag == 1 {
		for i := 0; i < n; i++ {
			a[i] = b[i]
		}
	}
}
