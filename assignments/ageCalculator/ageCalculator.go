package ageCalculator

import (
	"fmt"
	"time"

	"github.com/bearbin/go-age"
)

func AgeCalculator(dob string) int {

	dob_, err := time.Parse(time.DateOnly, dob)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	a := age.Age(dob_)
	return a
}
