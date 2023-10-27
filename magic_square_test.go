package magic_square_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_E2E_Rows(t *testing.T) {
	ms, _ := CreateSquare()
	assert.Equal(t, 9, findArraySum(ms.Row1))
	assert.Equal(t, 9, findArraySum(ms.Row2))
	assert.Equal(t, 9, findArraySum(ms.Row3))
}

//func Test_E2E_Columns(t *testing.T) {
//	ms, _ := CreateSquare()
//	column1 := createColumns(ms, 1)
//	column2 := createColumns(ms, 2)
//	column3 := createColumns(ms, 3)
//	assert.Equal(t, 9, findArraySum(column1))
//	assert.Equal(t, 9, findArraySum(column2))
//	assert.Equal(t, 9, findArraySum(column3))
//}
//
//func Test_E2E_Diagonals(t *testing.T) {
//	ms, _ := CreateSquare()
//	diagonal1 := createDiagonal1(ms)
//	diagonal2 := createDiagonal2(ms)
//	assert.Equal(t, 9, findArraySum(diagonal1))
//	assert.Equal(t, 9, findArraySum(diagonal2))
//}

func findArraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func createColumns(ms MagicSquare, element int) []int {
	column := []int{}
	column = append(column, ms.Row1[element])
	column = append(column, ms.Row2[element])
	column = append(column, ms.Row3[element])
	return column
}

func createDiagonal1(ms MagicSquare) []int {
	diagonal1 := []int{}
	diagonal1 = append(diagonal1, ms.Row1[0])
	diagonal1 = append(diagonal1, ms.Row2[1])
	diagonal1 = append(diagonal1, ms.Row3[2])

	return diagonal1
}

func createDiagonal2(ms MagicSquare) []int {
	diagonal2 := []int{}
	diagonal2 = append(diagonal2, ms.Row1[2])
	diagonal2 = append(diagonal2, ms.Row2[1])
	diagonal2 = append(diagonal2, ms.Row3[0])
	return diagonal2
}
