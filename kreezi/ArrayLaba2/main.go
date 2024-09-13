package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type Array []int
type CommandFunc func(args []string, arrays map[string]Array) error

func LoadArray(arrayName, fileName string, arrays map[string]Array) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	defer file.Close()

	array := arrays[arrayName]
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return err
		}

		array = append(array, number)
	}

	arrays[arrayName] = array

	return nil
}

func SaveArray(arrayName, fileName string, arrays map[string]Array) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	defer file.Close()

	array := arrays[arrayName]
	for _, element := range array {
		file.WriteString(strconv.Itoa(element) + "\n")
	}

	return nil
}

func RandArray(arrayName string, count, lb, rb int, arrays map[string]Array) error {
	array := arrays[arrayName]

	for i := 0; i < count; i++ {
		array = append(array, (count*i+lb)%rb)
	}

	arrays[arrayName] = array

	return nil
}

func ConcatArray(arrayName1, arrayName2 string, arrays map[string]Array) error {
	array1 := arrays[arrayName1]
	array2 := arrays[arrayName2]

	array1 = append(array1, array2...)

	arrays[arrayName1] = array1

	return nil
}

func FreeArray(arrayName string, arrays map[string]Array) error {
	arrays[arrayName] = Array{}

	return nil
}

func RemoveArray(arrayName string, index, count int, arrays map[string]Array) error {
	array := arrays[arrayName]

	if index < 0 || index > len(array) {
		fmt.Println("Error invalid index")
		return errors.New("invalid index")
	}

	if index+count > len(array) {
		array = append(array[:index], array[index:]...)
	} else {
		array = append(array[:index], array[index+count:]...)
	}

	arrays[arrayName] = array

	return nil
}

func CopyArray(arrayName1, arrayName2 string, lb, rb int, arrays map[string]Array) error {
	array1 := arrays[arrayName1]
	array2 := arrays[arrayName2]

	if lb > 0 && rb > 0 && lb < len(array1) && rb < len(array1) {
		array2 = append(array2, array1[lb:rb+1]...)
	} else {
		fmt.Println("Error invalid index")
		return errors.New("invalid index")
	}

	arrays[arrayName2] = array2

	return nil
}

func SortArray(arrayName string, arrays map[string]Array) error {
	array := arrays[string(arrayName[0])]
	sort.Ints(array)

	arrays[string(arrayName[0])] = array

	return nil
}

func shuffleArray(arrayName string, arrays map[string]Array) error {
	array := arrays[arrayName]

	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}

	arrays[arrayName] = array

	return nil
}

func statsArray(arrayName string, arrays map[string]Array) error {
	array := arrays[arrayName]

	fmt.Printf("Размер массива %s - %d\n", arrayName, len(array))

	maxEl, minEl := math.MinInt, math.MaxInt
	maxelIndex, minelIndex := 0, 0
	sum := 0
	frequency := make(map[int]int)
	maxCount := 0
	mostCountElement := array[0]
	for index, element := range array {
		if element > maxEl {
			maxEl = element
			maxelIndex = index
		}
		if element < minEl {
			minEl = element
			minelIndex = element
		}

		sum += element

		frequency[element]++
		if frequency[element] > maxCount {
			maxCount = frequency[element]
			mostCountElement = element
		} else if frequency[element] == maxCount && element > mostCountElement {
			mostCountElement = element
		}
	}

	mean := float64(sum) / float64(len(array))

	maxDeviation := 0.0
	for _, value := range array {
		deviation := math.Abs(float64(value) - mean)
		if deviation > maxDeviation {
			maxDeviation = deviation
		}
	}

	fmt.Printf("Минимальный элемент: %d (индекс %d)\n", minEl, minelIndex)
	fmt.Printf("Максимальный элемент: %d (индекс %d)\n", maxEl, maxelIndex)
	fmt.Printf("Наиболее часто встречающийся элемент: %d\n", mostCountElement)
	fmt.Printf("Среднее значение элементов: %.2f\n", mean)
	fmt.Printf("Максимальное отклонение от среднего значения: %.2f\n", maxDeviation)
	return nil
}

func PrintArray(arrayName string, index string, arrays map[string]Array) error {
	array := arrays[arrayName]

	if index == "all" {
		fmt.Printf("Массив %s - %v\n", arrayName, array)
		return nil
	} else {
		indexDigit, err := strconv.Atoi(index)
		if err != nil {
			fmt.Println("Error ATOI\n")
			return err
		}

		if indexDigit < 0 || indexDigit > len(array) {
			fmt.Println("Error invalid index\n")
			return errors.New("invalid index")
		}

		fmt.Printf("Элемент под индексом %s в массиве %s - %d\n", index, arrayName, array[indexDigit])
		return nil
	}
}

func PrintRangeArray(arrayName string, lb, rb int, arrays map[string]Array) error {
	array := arrays[arrayName]

	if rb+1 < len(array) {
		rb = rb + 1
	}

	if lb < 0 || rb < 0 || lb > len(array) || rb > len(array) {
		fmt.Println("Error invalid index")
		return errors.New("invalid index")
	}

	fmt.Printf("Элементы массива %s с %d по %d - %v\n", arrayName, lb, rb, array[lb:rb+1])
	return nil
}

func main() {
	var commandsMap = map[string]CommandFunc{
		"load": func(args []string, arrays map[string]Array) error { return LoadArray(args[0], args[1], arrays) },
		"save": func(args []string, arrays map[string]Array) error { return SaveArray(args[0], args[1], arrays) },
		"rand": func(args []string, arrays map[string]Array) error {
			count, _ := strconv.Atoi(args[1])
			lb, _ := strconv.Atoi(args[2])
			rb, _ := strconv.Atoi(args[3])
			return RandArray(args[0], count, lb, rb, arrays)
		},
		"concat": func(args []string, arrays map[string]Array) error { return ConcatArray(args[0], args[1], arrays) },
		"free":   func(args []string, arrays map[string]Array) error { return FreeArray(args[0], arrays) },
		"remove": func(args []string, arrays map[string]Array) error {
			index, _ := strconv.Atoi(args[1])
			count, _ := strconv.Atoi(args[2])
			return RemoveArray(args[0], index, count, arrays)
		},
		"copy": func(args []string, arrays map[string]Array) error {
			lb, _ := strconv.Atoi(args[1])
			rb, _ := strconv.Atoi(args[2])
			return CopyArray(args[0], args[3], lb, rb, arrays)
		},
		"sort":    func(args []string, arrays map[string]Array) error { return SortArray(args[0], arrays) },
		"shuffle": func(args []string, arrays map[string]Array) error { return shuffleArray(args[0], arrays) },
		"stats":   func(args []string, arrays map[string]Array) error { return statsArray(args[0], arrays) },
		"print": func(args []string, arrays map[string]Array) error {
			if len(args) == 2 {
				return PrintArray(args[0], args[1], arrays)
			} else {
				lb, _ := strconv.Atoi(args[1])
				rb, _ := strconv.Atoi(args[2])
				return PrintRangeArray(args[0], lb, rb, arrays)
			}
		},
	}

	var containerArrays map[string]Array = make(map[string]Array)

	var fileName string
	fmt.Print("Введите название файла: ")
	if _, err := fmt.Scan(&fileName); err != nil {
		fmt.Printf("Error scan Stdin: %v", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error open file: %v", err)
		return
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		var commandSplit []string
		var temp string

		for _, char := range line {
			if char == ' ' || char == ',' || char == ';' || char == '(' || char == ')' {
				if temp != "" {
					commandSplit = append(commandSplit, temp)
					temp = ""
				}
			} else {
				temp += string(unicode.ToLower(char))
			}
		}

		searchCommand(commandSplit, containerArrays, commandsMap)
	}

}

func searchCommand(command []string, arrays map[string]Array, commandsMap map[string]CommandFunc) {
	nameCommand := command[0]

	if function, exists := commandsMap[nameCommand]; exists {
		if err := function(command[1:], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else {
		fmt.Println("Invalid command!!!")
	}
}
