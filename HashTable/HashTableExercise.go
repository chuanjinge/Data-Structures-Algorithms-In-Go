package main

import "fmt"

func IsAnagram(str1 string, str2 string) bool {
	size1 := len(str1)
	size2 := len(str2)
	if size1 != size2 {
		return false
	}
	cm := make(Counter)
	for _, ch := range str1 {
		cm.Add(ch)
	}
	for _, ch := range str2 {
		cm.Remove(ch)
	}
	return (len(cm) == 0)
}

func RemoveDuplicate(str string) string {
	input := []rune(str)
	hs := make(Set)
	var output []rune
	for _, ch := range input {
		if hs.Find(ch) == false {
			output = append(output, ch)
			hs.Add(ch)
		}
	}
	return string(output)
}

func FindMissing(arr []int, start int, end int) (int, bool) {
	hs := make(Set)
	for _, i := range arr {
		hs.Add(i)
	}
	for curr := start; curr <= end; curr++ {
		if hs.Find(curr) == false {
			return curr, true
		}
	}
	return 0, false
}

func PrintRepeating(arr []int) {
	hs := make(Set)
	fmt.Print("Repeating elements are :: ")
	for _, val := range arr {
		if hs.Find(val) {
			fmt.Print(val, "  ")
		} else {
			hs.Add(val)
		}
	}
	fmt.Println()
}

func PrintFirstRepeating(arr []int) {
	size := len(arr)
	hs := make(Counter)

	for i := 0; i < size; i++ {
		hs.Add(arr[i])
	}
	for i := 0; i < size; i++ {
		hs.Remove(arr[i])
		if hs.Find(arr[i]) {
			fmt.Println("First Repeating number is : ", arr[i])
			return
		}
	}
}

func hornerHash(key []rune, tableSize int) int {
	size := len(key)
	h := 0
	for i := 0; i < size; i++ {
		h = (32*h + (int)(key[i])) % tableSize
	}
	return h
}

func main() {
	var1 := "hello"
	var2 := "elloh"
	var3 := "world"

	fmt.Println("IsAnagram : ", IsAnagram(var1, var2))
	fmt.Println("IsAnagram : ", IsAnagram(var1, var3))

	fmt.Println(RemoveDuplicate(var1))

	arr := []int{1, 2, 3, 5, 6, 7, 9, 8, 10}
	fmt.Print("Missing number is :: ")
	fmt.Println(FindMissing(arr, 1, 10))

	arr1 := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	PrintRepeating(arr1)

	PrintFirstRepeating(arr1)
}

// ***********************
type Counter map[interface{}]int

func (s *Counter) Add(key interface{}) {
	(*s)[key]++
}
func (s *Counter) Find(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *Counter) Remove(key interface{}) {
	val, ok := (*s)[key]
	if ok == false {
		return
	} else if val == 1 {
		delete((*s), key)
		return
	}
	(*s)[key]--
}

func (s *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*s)[key]
	return val, ok
}



type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}
//**************************************