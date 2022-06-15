package ex23

import (
	"testing"
	"fmt"
)

func TestRemove(t *testing.T){
	
	nums := []int{0,1,2,3,4,5,6,7,8,9,10}
	
	strs := []string{"start", "6756", "23123", "asvsd", "nbnvn", "fin"}
	
	res, _ := RemoveElem(nums, 5)
	fmt.Printf("1: %p\n2: %p\n", nums, res)
	fmt.Println(res)
	
	res1, _ := RemoveElem(strs, 5)
	fmt.Printf("1: %p\n2: %p\n", strs, res1)
	fmt.Println(res1)
}