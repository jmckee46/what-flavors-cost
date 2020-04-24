package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func whatFlavors(cost []int32, money int32) {

	metaSlice := convertSlice(cost)

	// sort.Slice(metaSlice, func(i, j int) bool { return metaSlice[i][0] < metaSlice[j][0] })
	sortedMetaSlice := mergeSort(metaSlice)

	lenMetaSlice := int32(len(metaSlice))
	for i := int32(0); i < lenMetaSlice; i++ {
		x := money - sortedMetaSlice[i][0]

		index, secondHalf := binarySearch(x, i, sortedMetaSlice)

		if secondHalf > 0 && index != i {
			if sortedMetaSlice[i][1]+1 < secondHalf+1 {
				fmt.Println(sortedMetaSlice[i][1]+1, secondHalf+1)
			} else {
				fmt.Println(secondHalf+1, sortedMetaSlice[i][1]+1)
			}
			return
		}
	}
}

func convertSlice(input []int32) [][]int32 {
	lenInput := int32(len(input))
	output := make([][]int32, lenInput)

	for i := int32(0); i < lenInput; i++ {
		output[i] = make([]int32, 2)
		output[i][0] = input[i]
		output[i][1] = i
	}

	return output
}

func binarySearch(itemToSearchFor int32, check int32, sortedInfo [][]int32) (int32, int32) {

	lenSortedInfo := int32(len(sortedInfo))

	low := int32(0)
	high := int32(lenSortedInfo - 1)

	for low <= high {
		median := (low + high) / 2

		if sortedInfo[median][0] < itemToSearchFor {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == lenSortedInfo || sortedInfo[low][0] != itemToSearchFor {
		return -1, -1
	}

	if sortedInfo[low][1] == check {
		if sortedInfo[low+1][0] == itemToSearchFor {
			return low + 1, sortedInfo[low+1][1]
		}
		return low, sortedInfo[low][1]
	}

	return low, sortedInfo[low][1]

}

func mergeSort(arr [][]int32) [][]int32 {
	arrLength := int32(len(arr))

	if arrLength == 1 {
		return arr
	}

	arrMiddle := arrLength / 2
	var left = make([][]int32, arrMiddle)
	var right = make([][]int32, arrLength-arrMiddle)

	for i := int32(0); i < arrLength; i++ {
		if i < arrMiddle {
			left[i] = arr[i]
		} else {
			right[i-arrMiddle] = arr[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right [][]int32) [][]int32 {
	result := make([][]int32, len(left)+len(right))

	i := int32(0)

	for int32(len(left)) > 0 && int32(len(right)) > 0 {
		if left[0][0] <= right[0][0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := int32(0); j < int32(len(left)); j++ {
		result[i] = left[j]
		i++
	}

	for j := int32(0); j < int32(len(right)); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {
	file, err := os.Open("test-case-3-INPUT")
	checkError(err)

	reader := bufio.NewReaderSize(file, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
