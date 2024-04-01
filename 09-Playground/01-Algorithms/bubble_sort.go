package main

func BubbleSort(a []float32) []float32 {
	for j := 0; j < len(a)-1; j++ {
		for i := 0; i < len(a)-1-j; i++ {
			if a[i] > a[i+1] {
				// swap
				tmp := a[i]
				a[i] = a[i+1]
				a[i+1] = tmp
			}
		}
	}

	return a
}