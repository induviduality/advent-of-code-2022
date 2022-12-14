/*

--- Day 5: Supply Stacks ---

The expedition can depart as soon as the final supplies have been unloaded
from the ships. Supplies are stored in stacks of marked crates, but because
the needed supplies are buried under many other crates, the crates need to
be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks.
To ensure none of the crates get crushed or fall over, the crane operator
will rearrange them in a series of carefully-planned steps. After the
crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate
procedure, but they forgot to ask her which crate will end up where, and
they want to be ready to unload them as soon as possible so they can
embark.

They do, however, have a drawing of the starting stacks of crates and the
rearrangement procedure (your puzzle input). For example:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

In this example, there are three stacks of crates. Stack 1 contains two
crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains
three crates; from bottom to top, they are crates M, C, and D. Finally,
stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure,
a quantity of crates is moved from one stack to a different stack. In the
first step of the above rearrangement procedure, one crate is moved from
stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3

In the second step, three crates are moved from stack 1 to stack 3. Crates
are moved one at a time, so the first crate to be moved (D) ends up below
the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3

Then, both crates are moved from stack 2 to stack 1. Again, because crates
are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3

Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3

The Elves just need to know which crate will end up on top of each stack;
in this example, the top crates are C in stack 1, M in stack 2, and Z
in stack 3, so you should combine these together and give the Elves the
message CMZ.

After the rearrangement procedure completes, what crate ends up on top of
each stack?

--- Part Two ---

As you watch the crane operator expertly rearrange the crates, you notice
the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly
wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air
conditioning, leather seats, an extra cup holder, and the ability to pick
up and move multiple crates at once.

Again considering the example above, the crates begin in the same
configuration:

    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 

However, the action of moving three crates from stack 1 to stack 3 means
that those three moved crates stay in the same order, resulting in this new
configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3

Next, as both crates are moved from stack 2 to stack 1, they retain their
order as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3

Finally, a single crate is still moved from stack 1 to stack 2, but now
it's crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3

In this example, the CrateMover 9001 has put the crates in a totally
different order: MCD.

Before the rearrangement process finishes, update your simulation so that
the Elves know where they should stand to be ready to unload the final
supplies. After the rearrangement procedure completes, what crate ends up
on top of each stack?

*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


// Stack is a stack implementation
type Stack struct {
	items []interface{}
}

// NewStack - creates a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push - Adds an item on top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop - Removes an item from the top of the  stack
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("Stack underflow")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek - returns the item at the top of the stack
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("Stack underflow")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear clears the stack
func (s *Stack) Clear() {
	s.items = nil
}

// Values returns the values in the stack
func (s *Stack) Values() []interface{} {
	return s.items
}

// reverseStack reverses the stack
func reverseStack(s *Stack) {
	var reversed Stack
	for !s.IsEmpty() {
		item, _ := s.Pop()
		reversed.Push(item)
	}
	*s = reversed
}

func reverseStacks(stacks *[]Stack) {
	for i := 0; i < len(*stacks); i++ {
		reverseStack(&(*stacks)[i])
	}
}

func getArguments(s string) (int, int, int) {
	parsed := strings.Split(s, " ")
	from, _ := strconv.Atoi(parsed[3])
	to, _ := strconv.Atoi(parsed[5])
	quantity, _ := strconv.Atoi(parsed[1])
	return quantity, from, to
}

func addBoxesByLines(line string, stacks *[]Stack) {
	// read every 4 characters
	// if second character is a space, then there's nothing there
	// if second character is a letter, then there's a box there
	for i := 0; i < len(line); i += 4 {
		if line[i + 1] != ' ' {
			(*stacks)[i / 4].Push(string(line[i + 1]))
		}
	}
}

func printStack(stack Stack) {
	fmt.Printf("Top -> ")
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Printf("%s ", item)
	}
	fmt.Println()
}

func printStacks(stacks []Stack) {
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("Stack %d: ", i + 1)
		printStack(stacks[i])
	}
}

// buildStacks builds the stacks from the raw input
// 
// The box for each stack takes 3 characters in the input.
// Every 3 characters are separated by a space.
// The number of character in each line of the input is 
// of the form 3 * n + (n - 1)
//
// 3 * n + (n - 1) = len(stackInput[0])
// 
// Hence, n = (len(stackInput[0]) + 1) / 4
// where n is the number of stacks in the input
func buildStacks(stackInput []string) []Stack {
	n := (len(stackInput[0]) + 1) / 4
	stacks := make([]Stack, n)
	for i := 0; i < n - 1; i++ {
		addBoxesByLines(stackInput[i], &stacks)
	}
	reverseStacks(&stacks)
	return stacks
}

func moveItems(quantity int, from int, to int, stacks *[]Stack) {
	for i := 0; i < quantity; i++ {
		item, _ := (*stacks)[from - 1].Pop()
		(*stacks)[to - 1].Push(item)
	}
}

func moveItemsWithoutReversing(quantity int, from int, to int, stacks *[]Stack) {
	auxStack := NewStack()
	for i := 0; i < quantity; i++ {
		item, _ := (*stacks)[from - 1].Pop()
		auxStack.Push(item)
	}

	for i := 0; i < quantity; i++ {
		item, _ := auxStack.Pop()
		(*stacks)[to - 1].Push(item)
	}
}

func rearrange(rearrangements []string, stacks *[]Stack, part int) {
	for _, rearrangement := range rearrangements {
		quantity, from, to := getArguments(rearrangement)
		if part == 1 {
			moveItems(quantity, from, to, stacks)
		} else if part == 2 {
			moveItemsWithoutReversing(quantity, from, to, stacks)
		} else {
			log.Fatal("Invalid part")
		}
	}

	// uncomment to see the rearrangement result
	// printStacks(*stacks)
}

func getTops(stacks *[]Stack) string {
	var tops string = ""
	for i := 0; i < len(*stacks); i++ {
		item, _ := (*stacks)[i].Peek()
		tops += item.(string)
	}
	return tops
}

func getSolution(rearrangements []string, stackInput []string, part int) string {
	stacks := buildStacks(stackInput)
	rearrange(rearrangements, &stacks, part)
	return getTops(&stacks)
}

func main() {
	file, err := os.Open("../../input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stackInput := []string{}
	rearrangements := []string{}

	// Reading till a blank lines is encountered,
	// since the stack input and the rearrangement
	// input are separated by a blank line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			stackInput = append(stackInput, line)
		}
	}

	// Reading the rearrangement input
	for scanner.Scan() {
		line := scanner.Text()
		rearrangements = append(rearrangements, line)
	}

	fmt.Println("Part One ------------------")
	fmt.Println("Tops of stacks: ", getSolution(rearrangements, stackInput, 1))
	fmt.Println("Part Two ------------------")
	fmt.Println("Tops of stacks: ", getSolution(rearrangements, stackInput, 2))
}
