package magic_square_kata

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type MagicSquare struct {
	Input        []float64
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

	fmt.Println("final elements")
	fmt.Println(ms.options)
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
	fmt.Println("final element")
	fmt.Println(returnElement)
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

func contains(elems []float64, v float64) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
