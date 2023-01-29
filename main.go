package main

import (
	"fmt"
	"os"
	"os/exec"
	"project-euler/exercises"
	"time"
)

func main() {
	for {
		fmt.Println("Select an exercise to run:	")
		fmt.Println("    1 - Sum of 3 and 5 multiples")
		fmt.Println("    2 - Even Fibonacci numbers sum")

		var selectedExercise int
		_, err := fmt.Scanf("%d", &selectedExercise)

		if err != nil || selectedExercise < 1 {
			fmt.Println("Invalid exercise selected")
			return
		}

		switch selectedExercise {
		case 1:
			exercises.SumOf3And5Multiples()
		case 2:
			exercises.EvenFibonacciNumbersSum()
		}

		time.Sleep(10 * time.Second)

		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
