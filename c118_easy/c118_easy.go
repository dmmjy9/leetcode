package c118_easy

func generate(numRows int) [][]int {
	var rtn = make([][]int, numRows)

	for i := 0; i < numRows; i++ {

		rtn[i] = make([]int, i+1)

		for j := 0; j <= i; j++ {
			if i - 1 >= 0 && j - 1 >= 0 && j < i {
				rtn[i][j] = rtn[i-1][j-1] + rtn[i-1][j]
			} else {
				rtn[i][j] = 1
			}
		}
	}
	return rtn
}
