package Problem_101_200

func getRow(rowIndex int) []int {
	lastRow := []int{1}
	for i := 1; i <= rowIndex; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = lastRow[j-1] + lastRow[j]
		}
		lastRow = row
	}
	return lastRow
}
