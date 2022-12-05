/*

--- Day 1: Calorie Counting ---

Santa's reindeer typically eat regular reindeer food, but they need a lot of magical energy to deliver presents on Christmas.
For that, their favorite snack is a special type of star fruit that only grows deep in the jungle.
The Elves have brought you on their annual expedition to the grove where the fruit grows.

To supply enough magical energy, the expedition needs to retrieve a minimum of fifty stars by December 25th.
Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar;
the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The jungle must be too overgrown and difficult to navigate in vehicles or access from the air;
the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies.
One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks,
rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory
from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

```
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
```

This list represents the Calories of the food carried by five Elves:

    The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
    The second Elf is carrying one food item with 4000 Calories.
    The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
    The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
    The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks, they need to know which Elf to ask:
they'd like to know how many Calories are being carried by the Elf carrying the most Calories.
In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?


--- Part Two ---

By the time you calculate the answer to the Elves' question, they've already realized that
the Elf carrying the most Calories of food might eventually run out of snacks.

To avoid this unacceptable situation, the Elves would instead like to know the total Calories
carried by the top three Elves carrying the most Calories. That way, even if one of those Elves
runs out of snacks, they still have two backups.

In the example above, the top three Elves are the fourth Elf (with 24000 Calories),
then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories).
The sum of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?

*/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Heap implementation taken from the docs & modified
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Leaving these functions in just to show that
// we don't need a heap if we just want to find the maximum calories
// and that it can be done in O(1) space rather than O(n) space (if heaps were used)
func updateMaxCalories(calories int, maxCalories *int) {
	if calories > *maxCalories {
		*maxCalories = calories
	}
}

func getMaxCalories(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var currentCalories int
	var maxCalories int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			updateMaxCalories(currentCalories, &maxCalories)
			currentCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += calories
		}
	}
	updateMaxCalories(currentCalories, &maxCalories)
	return maxCalories
}

func partOne(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	maxCalories := getMaxCalories(file)
	fmt.Println("Max calories:", maxCalories)
}

func getTopKSum(h *IntMaxHeap, k int) int {
	var sum int = 0
	for i := 0; i < k; i++ {
		sum += heap.Pop(h).(int)
	}
	return sum
}

func topKCaloriesSum(file *os.File, k int) int {
	scanner := bufio.NewScanner(file)

	// Use a max heap to keep track of the top K calories
	var currentCalories int = 0
	h := &IntMaxHeap{}
	heap.Init(h)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			heap.Push(h, currentCalories)
			currentCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += calories
		}
	}

	heap.Push(h, currentCalories)

	topKSum := getTopKSum(h, k)
	return topKSum
}

func partTwo(filePath string, k int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	topKCaloriesSum := topKCaloriesSum(file, k)

	fmt.Println("Top K calories:", topKCaloriesSum)
}

func main() {
	fmt.Println("Part One ----------------------------")
	partOne("../../input/day1.txt")
	fmt.Println("-------------------------------------")
	fmt.Println("Part 1, but with heaps --------------")
	partTwo("../../input/day1.txt", 1)
	fmt.Println("-------------------------------------")
	fmt.Println("Part Two ----------------------------")
	partTwo("../../input/day1.txt", 3)
	fmt.Println("-------------------------------------")
}
