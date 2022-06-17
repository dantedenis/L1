package ex13

// свап за счет "синтаксического сахара" го
func Swap(first, second *int) {
	*first, *second = *second, *first
}

// просто с помощью 2х сумм и разности
func SwapSum(first, second *int) {
	*first += *second
	*second = *first - *second
	*first = *first - *second
}
