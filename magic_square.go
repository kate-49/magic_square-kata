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
	CenterNumber float64
	NonCorners   []float64
	options      [][]float64
}

func CreateSquare() (MagicSquare, error) {
	ms := MagicSquare{
		Input: []float64{1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0},
	}
	ms.findAllOptions(ms.Input, 9)
	numberCount := ms.countIndividualElements()
	ms.findCentreElement(numberCount)
	ms.findNonCornerElements(numberCount)
	ms.removeElementFromOptions([]float64{ms.NonCorners[0], ms.CenterNumber, ms.NonCorners[1]})
	ms.removeElementFromOptions([]float64{ms.NonCorners[2], ms.CenterNumber, ms.NonCorners[3]})
	Row1 := ms.findLine(ms.NonCorners[0])
	Row3 := ms.findLine(ms.NonCorners[1])
	Col1 := ms.findLine(ms.NonCorners[2])
	Col3 := ms.findLine(ms.NonCorners[3])

	topLeft := findCommonElementFrom2Lines(Row1, Col1)
	topRight := findCommonElementFrom2Lines(Row1, Col3)

	ms.Output = append(ms.Output, []float64{topLeft, ms.NonCorners[0], topRight})
	ms.Output = append(ms.Output, []float64{ms.NonCorners[2], ms.CenterNumber, ms.NonCorners[3]})

	bottomLeft := findCommonElementFrom2Lines(Row3, Col1)
	bottomRight := findCommonElementFrom2Lines(Row3, Col3)
	ms.Output = append(ms.Output, []float64{bottomLeft, ms.NonCorners[1], bottomRight})
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

func (ms *MagicSquare) findCentreElement(numberCount map[string]int) {
	for k, v := range numberCount {
		if v == 4 {
			kAsInt, _ := strconv.ParseFloat(k, 64)
			ms.CenterNumber = kAsInt
		}
	}
}

func (ms *MagicSquare) findNonCornerElements(numberCount map[string]int) {
	elementsToReturn := []float64{}
	for k, v := range numberCount {
		if v == 2 {
			kAsInt, _ := strconv.ParseFloat(k, 64)
			elementsToReturn = append(elementsToReturn, kAsInt)
		}
	}
	ms.NonCorners = elementsToReturn
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
	var nonCorners []float64
	for _, t := range ms.options {
		containsCorner := false
		for _, s := range cornerElements {
			if contains(t, s) {
				containsCorner = true
			}
		}
		if containsCorner == false {
			nonCorners = append(nonCorners, t[0], t[2])
		}
	}

	ms.NonCorners = nonCorners
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

func getRemainingElementFromLine(line []float64, element1 float64, element2 float64) float64 {
	for _, t := range line {
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
