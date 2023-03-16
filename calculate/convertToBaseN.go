package calculate

import "strconv"

func ConvertToBaseN(num, base int) string {
	if num == 0 {
		return "0"
	}
	tmp := num
	if num < 0 {
		tmp = -num
	}
	ret := ""
	for tmp > 0 {
		ret = strconv.Itoa(tmp%base) + ret
		tmp /= base
	}
	if num < 0 {
		ret = "-" + ret
	}
	return ret
}
