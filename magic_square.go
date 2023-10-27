package magic_square_kata

import (
	"fmt"
	"sort"
)

type MagicSquare struct {
	Input   []float64
	Row1    []int
	Row2    []int
	Row3    []int
	options [][]float64
}

func CreateSquare() (MagicSquare, error) {
	ms := MagicSquare{
		Input: []float64{1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0},
	}
	options := ms.findAllOptions(ms.Input, 9)
	fmt.Println("end")
	fmt.Println(options)
	return ms, nil
}

func (ms *MagicSquare) findAllOptions(arr []float64, sum float64) [][]float64 {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := i + 1; k < len(arr); k++ {
				if (arr[i]+arr[j]+arr[k] == sum) && (arr[i] != arr[j]) && (arr[j] != arr[k]) && (arr[i] != arr[k]) {
					element, duplicate := ms.createNewElementAndCheckForDuplicates([]float64{arr[i], arr[j], arr[k]})
					if !duplicate {
						ms.options = append(ms.options, element)
					}
				}
			}
		}
	}
	return ms.options
}

func (ms *MagicSquare) createNewElementAndCheckForDuplicates(element []float64) ([]float64, bool) {
	sort.Float64s(element)
	for _, s := range ms.options {
		fmt.Println("element")
		fmt.Println(element)
		fmt.Println("s")
		fmt.Println(s)
		//if s == element {
		//	return element, true
		//}
	}
	return element, false
}
