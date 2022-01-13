package lib

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Seed(n int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(i+n))
	}

	return
}

func SeedWithinRange(n int, bottom, up int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(up-bottom)+bottom)
	}
	return
}

func CheckSorted(arr []int, n int) error {
	//fmt.Println("sorted, checking")
	if len(arr) != n {
		err := fmt.Sprintln("Length incorrect:", len(arr), n)
		return errors.New(err)
	}

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			err := fmt.Sprintln("Failed:", arr[i-1], arr[i])
			return errors.New(err)
		}
	}

	return nil
}
