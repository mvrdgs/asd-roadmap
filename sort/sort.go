package sort

func CocktailSort(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
		right--
		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		left++
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
