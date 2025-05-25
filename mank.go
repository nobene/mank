package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	easy = 1
	hard = 4
)

var n [14]int
var d [14]string
var total int
var free int

func initGrid() {
	rnd := 0
	for i := 1; i < 14; i++ {
		rnd = rand.Intn(3) + 5
		if i != 7 {
			total = total + rnd
		}
		n[i] = rnd
	}
	n[13] = n[13] + 72 - total
	total = n[1] + n[2] + n[3] + n[4] + n[5] + n[6] + n[8] + n[9] + n[10] + n[11] + n[12] + n[13]
	n[7] = 0
	n[0] = 0
	for z := range 14 {
		d[z] = strconv.Itoa(n[z])
	}
	d[0] = ""
	d[7] = ""
}

func drawGrid() {
	//	fmt.Println(total)
	fmt.Println("                       ╔══════╦════╦════╦════╦════╦════╦════╦══════╗")
	fmt.Printf("                       ║      ║ %2s ║ %2s ║ %2s ║ %2s ║ %2s ║ %2s ║      ║ \n", d[1], d[2], d[3], d[4], d[5], d[6])
	fmt.Printf("                       ║  %2s  ╬════╬════╬════╬════╬════╬════╬  %2s  ║\n", d[0], d[7])
	fmt.Printf("                       ║      ║ %2s ║ %2s ║ %2s ║ %2s ║ %2s ║ %2s ║      ║\n", d[13], d[12], d[11], d[10], d[9], d[8])
	fmt.Println("                       ╚══════╩════╩════╩════╩════╩════╩════╩══════╝")
	fmt.Println("                                 A    B    C    D    E    F\n")
}

func turn(y int) int {
	hand := n[y]
	if hand == 0 {
		fmt.Println("No stones here, choose another pit (")
		return 0
	}
	n[y] = 0
	for {
		y++
		if y == 14 {
			y = 0
		}
		n[y] += 1
		hand--
		if hand < 1 {
			break
		}
	}
	if n[y] == 1 {
		switch y {
		case 13:
			if n[1] == 0 {
				break
			}
			n[7] += n[1]
			n[7]++
			n[13]--
			n[1] = 0
		case 12:
			if n[2] == 0 {
				break
			}
			n[7] += n[2]
			n[7]++
			n[12]--
			n[2] = 0
		case 11:
			if n[3] == 0 {
				break
			}
			n[7] += n[3]
			n[7]++
			n[11]--
			n[3] = 0
		case 10:
			if n[4] == 0 {
				break
			}
			n[7] += n[4]
			n[7]++
			n[10]--
			n[4] = 0
		case 9:
			if n[5] == 0 {
				break
			}
			n[7] += n[5]
			n[7]++
			n[9]--
			n[5] = 0
		case 8:
			if n[6] == 0 {
				break
			}
			n[7] += n[6]
			n[7]++
			n[8]--
			n[6] = 0
		default:
			break
		}
	}
	for z := range 14 {
		d[z] = strconv.Itoa(n[z])
		if n[z] == 0 {
			d[z] = ""
		}
	}
	drawGrid()
	if y == 7 {
		return 7
	}
	return 0
}

func compturn(x int) {
	hand := n[x]
	if hand == 0 {
        decide()
		return
	}
	n[x] = 0
	for {
		x++
		if x == 14 {
			x = 0
		}
		n[x] += 1
		hand--
		if hand < 1 {
			break
		}
	}
	for z := range 14 {
		d[z] = strconv.Itoa(n[z])
		if n[z] == 0 {
			d[z] = ""
		}
	}
	drawGrid()
	if x == 0 {
		decide()
	}
	return
}

func decide() {
	best := 1
	best = rand.Intn(6) + 1
	if best > 6 {
		best = 1
	}
	if best == 0 {
		best = 6
	}
	compturn(best)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initGrid()
	drawGrid()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("When entering moves, you can also enter Q to quit or S to start again.")
	moves := 0
outer:
	for {
		//		drawGrid()
		for {
			//			fmt.Println("Moves so far =", moves, "\n")
			fmt.Print("Enter move : ")
			scanner.Scan()
			move := strings.ToUpper(strings.TrimSpace(scanner.Text()))
			//			check(scanner.Err())
			switch move {
			case "A":
				free = turn(13)
				if free == 7 {
					free = 0
					continue outer
				}
				decide()
				moves++
				continue outer
			case "B":
				free = turn(12)
				if free == 7 {
					free = 0
					continue outer
				}
				decide()
				moves++
				continue outer
			case "C":
				free = turn(11)
				if free == 7 {
					free = 0
					continue outer
				}
				decide()
				moves++
				continue outer
			case "D":
				free = turn(10)
				if free == 7 {
					free = 0
					continue outer
				}
				decide()
				moves++
				continue outer
			case "E":
				free = turn(9)
				if free == 7 {
					free = 0
					continue outer
				}
				decide()
				moves++
				continue outer
			case "F":
				free = turn(8)
				if free == 7 {
					free = 0
					continue outer
				}
				decide()
				moves++
				continue outer
			case "Q":
				return
			case "S":
                total = 0
				initGrid()
				drawGrid()
				moves = 0
				continue outer
			default:
				fmt.Println("Invalid move, try again.")
			}
		}
	}
}
