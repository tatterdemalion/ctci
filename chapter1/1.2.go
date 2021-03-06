package chapter1

// Implement a function void reverse(char* str) in C or C++ which
// reverses a null- terminated string.

// reverses given string pointer
func reverse(s *string) {

	arr := make([]uint8, len(*s))
	n := 0
	for _, r := range *s {
		arr[n] = uint8(r)
		n++
	}

	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	*s = string(arr)
}
