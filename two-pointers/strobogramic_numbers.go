package two_pointers

var strobogrammaticMap = map[string]string{
    "1": "1",
    "0": "0",
    "6": "9",
    "9": "6",
    "8": "8",
}

// isStrobogrammatic checks if a number is strobogrammatic using two pointers.
func isStrobogrammatic(num string) bool {
	for i, j := 0, len(num) - 1; i <= j; i,j = i + 1, j - 1 {
		rotated := strobogrammaticMap[string(num[i])]
		if rotated != string(num[j]) {
			return false
		}
	}

	// If all are strobogrammatic, then entire number supplied is strobogrammatic.
	return true
}
