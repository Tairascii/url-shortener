package pkg

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func EncodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var encoded []byte
	base := int64(len(base62Chars))

	for num > 0 {
		remainder := num % base
		encoded = append([]byte{base62Chars[remainder]}, encoded...)
		num = num / base
	}

	return string(encoded)
}
