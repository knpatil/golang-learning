package main

import "fmt"

func main() {

	nums := []float64{140, 5, 1, -4, 20, 90, -1, 23}
	min, max, avg, success := minMaxAvg(nums)
	fmt.Println("-------------------------")
	if success {
		fmt.Printf("Min: %.2f, Max: %.2f, Avg: %.2f\n", min, max, avg)
	} else {
		fmt.Println("Empty slice provided.")
	}
	fmt.Println("-------------------------")
	fmt.Println(fib(45))
	fmt.Println("-------------------------")
}

// write a function minMaxAvg which accepts a float64 slice and returns three values,
// the minimum value in the slice, the maximum value in the slice,
// and the average of all the values
func minMaxAvg(nums []float64) (float64, float64, float64, bool) {

	if nums == nil {
		return 0, 0, 0, false
	}

	min, max, total := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
		total += nums[i]
	}
	return min, max, total / float64(len(nums)), true
}

// write a function fib that returns a slice containing the first n Fibonacci numbers
func fib(n int) []int {
	if n <= 1 {
		return nil
	}
	fibs := make([]int, n)
	fibs[0], fibs[1] = 1, 1
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}
