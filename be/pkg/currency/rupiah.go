package currency

import "strconv"

func Rupiah(amount int) string {
	if amount == 0 {
		return "Rp 0"
	}

	s := strconv.Itoa(amount)
	n := len(s)

	result := ""
	count := 0

	for i := n - 1; i >= 0; i-- {
		result = string(s[i]) + result
		count++

		if count%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}
