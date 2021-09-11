package downcase

import "regexp"

func Downcase(s string) (string, error) {
	r, _ := regexp.Compile("[A-Z]+")
	downcase := r.ReplaceAllFunc([]byte(s), func(bytes []byte) []byte {
		for i := range bytes {
			bytes[i] += 32
		}
		return bytes
	})
	return string(downcase), nil
}
