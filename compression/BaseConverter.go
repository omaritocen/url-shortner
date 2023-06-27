package compression

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// EncodeBase62 encodes an int64 to a base62 encoded string.
func EncodeBase62(number int64) string {
	return EncodeUint64(uint64(number))
}

// EncodeUint64 encodes a uint64 to a base62 encoded string.
func EncodeUint64(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := uint64(len(alphabet))

	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}
