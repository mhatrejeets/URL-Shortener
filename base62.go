package main

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


 
func encodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		result = string(charset[num%62]) + result
		num /= 62
	}
	return result
}

