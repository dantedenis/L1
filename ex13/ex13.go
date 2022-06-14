package ex13

func Swap(first, second *int) {
	*first, *second = *second, *first
}

func SwapSum(first, second *int) {
	*first += *second
	*second = *first - *second
	*first = *first - *second
}
