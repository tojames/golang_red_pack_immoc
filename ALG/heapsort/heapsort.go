package heapsort

var arr []interface{}

func heapSort(arr []int64) []int64 {
	arrLen := len(arr)
	buildMaxHeap(arr)
	for i := arrLen -1; i >=0; i -- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return []int64{}
}

func buildMaxHeap(arr []int64) []int64 {
	arrLen := len(arr)
	for i := arrLen / 2; i >= 0; i -- {
		heapify(arr, i, arrLen)
	}

	return []int64{}
}

func heapify(arr []int64, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, i, largest)
	}
}

func swap(arr []int64, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
