package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	validateErrorsOnGroup()
}

func validateErrorsOnGroup() {
	var g errgroup.Group
	var divisible []int32
	numbers := []int32{1, 2, 3, 4, 5, 6}

	for k := range numbers {
		num := numbers[k]
		g.Go(func() error {
			res, err := divideByTwo(num)
			if res > 0 {
				divisible = append(divisible, res)
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("An error ocurred:", err)
	}

	if len(divisible) > 0 {
		fmt.Println("Results:", divisible)
	}
}

func divideByTwo(num int32) (int32, error) {
	if num%2 != 0 {
		return 0, fmt.Errorf("%d is not divisible by 2", num)
	}
	return num / 2, nil
}
