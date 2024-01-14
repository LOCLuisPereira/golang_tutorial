package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {

	// shift elements inside array 1
	for i := m-1; i >= 0; i -- {
		nums1[i+n] = nums1[i]
	}

	i := 0
	i1 := n
	i2 := 0
	for {

		if i1 == n+m && i2 == n {
			break
		}

		if i1 == n+m {

		}
		

		if i2 == n {
			
		}

		i ++
		j ++
	}


	fmt.Println(nums1, m)
	fmt.Println(nums2, n)

}

func main(){

	a := []int{1,3,6,0,0,0}
	ax := 3

	b := []int{2,4,5}
	bx := 3

	merge( a, ax, b, bx )

}