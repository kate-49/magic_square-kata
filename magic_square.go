package magic_square_kata

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type MagicSquare struct {
	Input        []float64
	Output       [][]float64
	Row1         []float64
	Row2         []float64
	Row3         []float64
	Diagonal1    []float64
	Diagonal2    []float64
	NonCorner1   []float64
	NonCorner2   []float64
	options      [][]float64
	finalOptions [][]float64
}

func CreateSquare() (MagicSquare, error) {
	ms := MagicSquare{
		Input: []float64{1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0},
	}
	ms.findAllOptions(ms.Input, 9)
	numberCount := ms.countIndividualElements()
	middleSquare := findCentreElement(numberCount)
	cornerElements := findCornerElements(numberCount)
	//diagonal must be 2 corner elements and middle square
	ms.Diagonal1, ms.Diagonal2 = ms.findDiagonals(cornerElements, middleSquare)
	ms.removeElementFromOptions(ms.Diagonal1)
	ms.removeElementFromOptions(ms.Diagonal2)
	ms.nonCorners(cornerElements)
	ms.removeElementFromOptions(ms.NonCorner1)
	ms.removeElementFromOptions(ms.NonCorner2)

	ms.Output = append(ms.Output, []float64{0, ms.NonCorner1[0], 0})
	ms.Output = append(ms.Output, []float64{0, ms.NonCorner1[1], 0})
	ms.Output = append(ms.Output, []float64{0, ms.NonCorner1[2], 0})
	ms.Output[1][0] = ms.NonCorner2[0]
	ms.Output[1][1] = ms.NonCorner2[1]
	ms.Output[1][2] = ms.NonCorner2[2]
	match1 := ms.findLine(ms.Output[0][1])
	fmt.Println(match1)
	match2 := ms.findLine(ms.Output[1][0])
	fmt.Println(match2)
	ms.Output[0][0] = findCommonElementFrom2Lines(match1, match2)
	ms.Output[0][2] = getRemainingElement(match1, ms.Output[0][0], ms.Output[0][1])
	ms.Output[2][0] = getRemainingElement(match2, ms.Output[0][0], ms.Output[1][0])

	match3 := ms.findLine(ms.Output[1][2])
	ms.Output[2][2] = getRemainingElement(match3, ms.Output[0][2], ms.Output[1][0])

	ms.Row1 = ms.Output[0]
	ms.Row2 = ms.Output[1]
	ms.Row3 = ms.Output[2]
	return ms, nil
}

func (ms *MagicSquare) findAllOptions(arr []float64, sum float64) [][]float64 {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := i + 1; k < len(arr); k++ {
				if (arr[i]+arr[j]+arr[k] == sum) && (arr[i] != arr[j]) && (arr[j] != arr[k]) && (arr[i] != arr[k]) {
					element, duplicate := ms.checkIfElementExistsAlready([]float64{arr[i], arr[j], arr[k]})
					if !duplicate {
						ms.options = append(ms.options, element)
					}
				}
			}
		}
	}
	return ms.options
}

func (ms *MagicSquare) checkIfElementExistsAlready(newElement []float64) ([]float64, bool) {
	sort.Float64s(newElement)
	for _, s := range ms.options {
		sort.Float64s(s)
		equal := reflect.DeepEqual(s, newElement)
		if equal {
			return nil, true
		}
	}
	return newElement, false
}

func (ms *MagicSquare) countIndividualElements() map[string]int {
	numberCount := map[string]int{}

	for _, s := range ms.options {
		for _, number := range s {
			stringNumber := fmt.Sprintf("%.1f", number)
			numberCount[stringNumber]++
		}
	}
	return numberCount
}

func findCentreElement(numberCount map[string]int) float64 {
	for k, v := range numberCount {
		if v == 4 {
			kAsInt, _ := strconv.ParseFloat(k, 64)
			return kAsInt
		}
	}
	return 0
}

func findCornerElements(numberCount map[string]int) []float64 {
	elementsToReturn := []float64{}
	for k, v := range numberCount {
		if v == 3 {
			kAsInt, _ := strconv.ParseFloat(k, 64)
			elementsToReturn = append(elementsToReturn, kAsInt)
		}
	}
	return elementsToReturn
}

func (ms *MagicSquare) findDiagonals(cornerElements []float64, centreSquare float64) ([]float64, []float64) {
	returnElement := [][]float64{}

	for _, s := range cornerElements {
		for _, t := range ms.options {
			if contains(t, s) {
				if contains(t, centreSquare) {
					returnElement = append(returnElement, t)
				}
			}
		}
	}
	return returnElement[0], returnElement[1]
}

func (ms *MagicSquare) removeElementFromOptions(element []float64) {
	newArray := [][]float64{}
	for _, s := range ms.options {
		equal := reflect.DeepEqual(s, element)
		if !equal {
			newArray = append(newArray, s)
		}
	}
	ms.options = newArray
}

func (ms *MagicSquare) nonCorners(cornerElements []float64) {
	var nonCorners [][]float64
	for _, t := range ms.options {
		containsCorner := false
		for _, s := range cornerElements {
			if contains(t, s) {
				containsCorner = true
			}
		}
		if containsCorner == false {
			nonCorners = append(nonCorners, t)
		}
	}
	ms.NonCorner1 = nonCorners[0]
	ms.NonCorner2 = nonCorners[1]
}

func (ms *MagicSquare) findLine(elementToMatch float64) []float64 {
	for _, t := range ms.options {
		if contains(t, elementToMatch) {
			return t
		}
	}
	return []float64{}
}

func findCommonElementFrom2Lines(line1 []float64, line2 []float64) float64 {
	for _, t := range line1 {
		for _, s := range line2 {
			if t == s {
				return t
			}
		}
	}
	return 0
}

func getRemainingElement(line1 []float64, element1 float64, element2 float64) float64 {
	for _, t := range line1 {
		if t != element1 && t != element2 {
			return t
		}
	}
	return 0
}

func contains(elems []float64, v float64) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
