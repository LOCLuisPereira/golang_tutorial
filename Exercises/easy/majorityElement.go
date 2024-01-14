package main

import "fmt"


func majorityElement(nums []int) int {

	mem := make(map[int]int)

	for _, n := range(nums){
		mem[n] = mem[n] + 1
	}

	mx := -1
	mx_n := -1
	for n, count := range(mem){
		if mx < count {
			mx = count
			mx_n = n
		}
	}

	return mx_n
}

func main(){

	input := []int {1,2,2,3,4}

	output := majorityElement( input )

	fmt.Println( output )
}