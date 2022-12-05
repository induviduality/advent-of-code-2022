/*

Navigation woes

This year, just like previous years, Santa Claus wants to visit all the houses
again to bring the presents to all the lovely children. Unfortunately, during
the preparations it turns out that his Santa Positioning System (KPS) is no
longer working and with the worldwide chip shortage, a new order is not possible.

One of the elves remembers the old-fashioned method and asks Santa to look in the
glove compartment of the sleigh. Here he finds a large magic scroll with navigation
instructions. This list is full of rules for walking, turning and jumping.

There is a small card on the roll with a separate instruction to always look towards
his house when he starts the navigation instructions. Of course, as everyone knows,
Santa lives in the North Pole because he is extremely allergic to penguins.

Suppose the following navigation instructions are on the list:

turn 90
walk 6
jump 2
turn -45
walk 2

There are five instructions in this example.

    According to the small map, Santa should start facing north.
    With the first instruction, Santa turns 90 degrees (to the right) and faces east.
    If he follows the second instruction he takes 6 steps (towards the east).
    Instruction number three makes Santa take a leap forward. The distance he jumps is equal to that of two steps.
    The fourth instruction makes it turn -45 degrees. This instruction causes Santa to face north-east.
    With the fifth and final instruction, Santa takes two steps to the northeast. If Santa takes
		one diagonal step, this corresponds in terms of distance to one step horizontally and one step vertically.

Santa breaks out in sweat. One of the most important functions of his KPS is missing for his
time schedule and that is the Manhattan distance between the starting point and the ending point.
In the example above, that's a Manhattan distance of 12.

Given the navigation instructions, find the Manhattan distance between the start and end points.
This distance is then the answer to part 1.

Enter the distance below.

infi@nerd-pc ▶ ~/aoc2022 ▶

Part 2 ----------------------------------------------------------------------------------------------------------------

While Santa was busy with all the navigation instructions, it just started snowing. Because of this,
Santa has left a trail of steps in the snow. As Santa leaves on his horse-drawn sleigh, he looks back
again and sees that his steps have made a pattern of letters.

Look closely at the tracks in the snow and look for the word that Santa left with his steps in the snow.
This word is the answer to part 2.

Enter the word below.

infi@nerd-pc ▶ ~/aoc2022 ▶

*/

/*

[Congrats message - Translated from Dutch]

On behalf of Santa: Thank you!

Santa Claus is very happy with all your help during navigation and thanks to you he managed
to bring a package under the Christmas tree to all children. Fortunately, no one needs to know
that Santa Claus is also suffering from the global chip shortage.

Santa also likes to show his gratitude. Among those who successfully solve the puzzle, we will
raffle five Infi packs. We have an appropriate gift card ready for the nicest/most original/best
solution. Follow the link below if you want to win.

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
)

type coordinates [2]int

type instruction struct {
	command string
	number  int
}

var directions = map[string]coordinates{
	"north":     {0, 1},
	"south":     {0, -1},
	"east":      {1, 0},
	"west":      {-1, 0},
	"northeast": {1, 1},
	"southeast": {1, -1},
	"southwest": {-1, -1},
	"northwest": {-1, 1},
}


var direction2degree = map[string]int{
	"north": 0,
	"south": 180,
	"east": 90,
	"west": 270,
	"northeast": 45,
	"northwest": 315,
	"southeast": 135,
	"southwest": 225,
}

var degrees2direction = map[int]string{
	0: "north",
	180: "south",
	90: "east",
	270: "west",
	45: "northeast",
	315: "northwest",
	135: "southeast",
	225: "southwest",
}

var allPlaces = []coordinates{}

func turn(currentDirection string, turn int) string {
	degrees := direction2degree[currentDirection]
	degrees += turn
	if degrees < 0 {
		degrees += 360
	}
	degrees = degrees % 360
	return degrees2direction[degrees]
}

func move(santa coordinates, direction string, number int) coordinates {
	directionCoordinates := directions[direction]
	santa[0] += directionCoordinates[0] * number
	santa[1] += directionCoordinates[1] * number
	allPlaces = append(allPlaces, santa)
	return santa
}

func partOne(instructions []instruction) {
	currentDirection := "north"
	santa := coordinates{0, 0}

	for _, instruction := range instructions {
		if instruction.command == "draai" {
			currentDirection = turn(currentDirection, instruction.number)
		} else if instruction.command == "loop" {
			santa = move(santa, currentDirection, instruction.number)
		} else if instruction.command == "spring" {
			// appending {-1, -1} to mark which steps are jumps
			// {-1, -1} will be removed in the python script
			allPlaces = append(allPlaces, coordinates{-1, -1})
			santa = move(santa, currentDirection, instruction.number)
		}
	}

	fmt.Println("Santa is at", santa)
	fmt.Println("Manhattan distance is", math.Abs(float64((santa[0])) + math.Abs(float64(santa[1] - 0))))
}

func main() {
	file, err := os.Open("challenge1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	var command string
	var number int
	instructions := []instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &command, &number)
		instructions = append(instructions, instruction{command, number})
	}

	partOne(instructions)

	// Used the below line to get the answer to part 2
	// by getting all the places Santa visited
	// Then I copied over the coordinates to the python code
	// and plotted the coordinates on a graph
	// fmt.Println(allPlaces)
}