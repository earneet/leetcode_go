package Problem_901_1000

func addToArrayForm(num []int, k int) []int {
	carry := 0
	for i := 0; i < len(num); i++ {
		sum := num[len(num)-(i+1)] + carry + k%10
		carry = sum / 10
		num[len(num)-(i+1)] = sum % 10
		k /= 10
	}
	for carry != 0 || k != 0 {
		num = append([]int{(carry + k%10) % 10}, num...)
		carry = (carry + k%10) / 10
		k /= 10
	}
	return num
}
