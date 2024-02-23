package helper

import "bytes"

func SplitText(b []byte, splitLength int) [][]byte {
	current := new(bytes.Buffer)
	var bSlice [][]byte
	counter := 0
	shouldTerminate := false

	for i, c := range b {
		if shouldTerminate == true && c == byte(32) {
			counter = 0
			shouldTerminate = false

			bSlice = append(bSlice, current.Bytes())
			current = new(bytes.Buffer)

			continue
		}

		counter++
		current.Write([]byte{c})

		if counter > splitLength {
			shouldTerminate = true
		}

		if i == len(b)-1 {
			bSlice = append(bSlice, current.Bytes())
		}
	}

	return bSlice
}
