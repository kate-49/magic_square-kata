package magic_square_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_E2E_Rows(t *testing.T) {
	ms, _ := CreateSquare()
	assert.Equal(t, 9.0, findArraySum(ms.Output[0]))
	assert.Equal(t, 9.0, findArraySum(ms.Output[1]))
	assert.Equal(t, 9.0, findArraySum(ms.Output[2]))
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

func findArraySum(arr []float64) float64 {
	res := 0.0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func createColumns(ms MagicSquare, element int) []float64 {
	column := []float64{}
	column = append(column, ms.Output[0][element])
	column = append(column, ms.Output[1][element])
	column = append(column, ms.Output[2][element])
	return column
}

func createDiagonal1(ms MagicSquare) []float64 {
	diagonal1 := []float64{}
	diagonal1 = append(diagonal1, ms.Output[0][0])
	diagonal1 = append(diagonal1, ms.Output[1][1])
	diagonal1 = append(diagonal1, ms.Output[2][2])

	return diagonal1
}

func createDiagonal2(ms MagicSquare) []float64 {
	diagonal2 := []float64{}
	diagonal2 = append(diagonal2, ms.Output[0][2])
	diagonal2 = append(diagonal2, ms.Output[1][1])
	diagonal2 = append(diagonal2, ms.Output[2][0])
	return diagonal2
}
