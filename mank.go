// 2025 Copyright Mank developers, licensed under APACHE license.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	easy = 1
	hard = 4
)

var n, k [14]int
var d [14]string
var r [9010]string
var pc, hum [9010]int
var total, over int
var free, moves int
var result string

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
	who_won()
	if y == 7 {
		return 7
	}
	return 0
}

func compturn(x int) int {
	hand := n[x]
	if hand == 0 {
		who_won()
		return 1
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
	if n[x] == 1 {
		switch x {
		case 1:
			if n[13] == 0 {
				break
			}
			n[0] += n[13]
			n[0]++
			n[1]--
			n[13] = 0
		case 2:
			if n[12] == 0 {
				break
			}
			n[0] += n[12]
			n[0]++
			n[2]--
			n[12] = 0
		case 3:
			if n[11] == 0 {
				break
			}
			n[0] += n[11]
			n[0]++
			n[3]--
			n[11] = 0
		case 4:
			if n[10] == 0 {
				break
			}
			n[0] += n[10]
			n[0]++
			n[4]--
			n[10] = 0
		case 5:
			if n[9] == 0 {
				break
			}
			n[0] += n[9]
			n[0]++
			n[5]--
			n[9] = 0
		case 6:
			if n[8] == 0 {
				break
			}
			n[0] += n[8]
			n[0]++
			n[6]--
			n[8] = 0
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
	who_won()
	return 0
}

func decide1() {
	who_won()
	best := 1
	best = pc_can_steal()
    fmt.Println("pc_can_steal gave ", best)
	if best != 0 {
		if n[best] != 0 {
			compturn(best)
			fmt.Println("decide() exited... after can_steal() gave best != 0")
			fmt.Println("best from can_steal() ", best)
			return
		}
		//		best = rand.Intn(6) + 1
	}
    if n[13] == 0 { best = 1 }
    if n[12] == 0 { best = 2 }
    if n[11] == 0 { best = 3 }
    if n[10] == 0 { best = 4 }
    if n[9] == 0 { best = 5 }
    if n[8] == 0 { best = 6 }
    if  best != 0 {
         if n[best] != 0 {
             compturn(best)
             fmt.Println("saved from stealing best was ", best)
             return
         }
    }
	max := slices.Max(n[1:7])
	//    fmt.Println(max, n[1:7])
	for t := range n {
		if n[t] == max {
			compturn(t)
			fmt.Println("max best choosen in decide() was ", t)
			return
		}
	}
	fmt.Println(best)
	for {
		best = rand.Intn(6) + 1
		if n[best] != 0 {
			compturn(best)
			//     		best = rand.Intn(6) + 1
			fmt.Println("in decide() final best was ", best)
			fmt.Println("decide() exited...")
			return
		}
	}
}

func can_steal() int {
	for t := range 6 {
		t++
		for v := range 14 {
			k[v] = n[v]
		}
		foot := k[t]
		k[t] = 0
		for {
			t++
			if t == 14 {
				t = 0
			}
			k[t] += 1
			foot--
			if foot < 1 {
				break
			}
		}
		if k[t] == 1 {
			switch t {
			case 1:
				if k[13] != 0 {
					return 1
				}
				k[0] += k[13]
				k[0]++
				k[1]--
				k[13] = 0
			case 2:
				if k[12] != 0 {
					return 2
				}
				k[0] += k[12]
				k[0]++
				k[2]--
				k[12] = 0
			case 3:
				if k[11] != 0 {
					return 3
				}
				k[0] += k[11]
				k[0]++
				k[3]--
				k[11] = 0
			case 4:
				if k[10] != 0 {
					return 4
				}
				k[0] += k[10]
				k[0]++
				k[4]--
				k[10] = 0
			case 5:
				if k[9] != 0 {
					return 5
				}
				k[0] += k[9]
				k[0]++
				k[5]--
				k[9] = 0
			case 6:
				if k[8] != 0 {
					return 6
				}
				k[0] += k[8]
				k[0]++
				k[6]--
				k[8] = 0
			default:
				return 0
			}
		}
	}
	//  os.Exit(0)
	return 0
}

func who_won() {
	var top, bottom int
	top = n[1] + n[2] + n[3] + n[4] + n[5] + n[6]
	bottom = n[8] + n[9] + n[10] + n[11] + n[12] + n[13]
	if top != 0 {
		if bottom != 0 {
			return
		}
		n[7] += bottom
		n[0] += top
		gameover()
		return
	}
	n[7] += bottom
	n[0] = n[0] + top
	gameover()
	return
}

func gameover() {
	fmt.Println("                    In " + strconv.Itoa(moves) + " moves: GAME OVER !!!  PC: " + strconv.Itoa(n[0]) + "    Human: " + strconv.Itoa(n[7]))
	left := n[0]
	right := n[7]
	for b := range 14 {
		n[b] = 0
	}
	n[0] = left
	n[7] = right
	if left == right {
		result = "DRAW"
	}
	if left < right {
		result = "HUMAN WON"
	}
	if left > right {
		result = "PC WON"
	}
	for m := range 14 {
		d[m] = strconv.Itoa(n[m])
		if n[m] == 0 {
			d[m] = ""
		}
	}
	drawGrid()
	over = 1
	return
	// os.Exit(0)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initGrid()
	drawGrid()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("When entering moves, you can also enter Q to quit or S to start again.")
	// moves := 0
outer:
	for {
		for {
			//			fmt.Println("Moves so far =", moves, "\n")
			fmt.Print("Enter move : ")
			scanner.Scan()
			move := strings.ToUpper(strings.TrimSpace(scanner.Text()))
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
			case "R":
				run()
				//				os.Exit(0)
				continue outer
			case "S":
				total = 0
				over = 0
				initGrid()
				drawGrid()
				moves = 0
				continue outer
			case "T":
				autoplay()
				total = 0
				over = 0
				initGrid()
				drawGrid()
				moves = 0
				continue outer
				//				os.Exit(0)
			default:
				fmt.Println("Invalid move, try again.")
			}
		}
	}
}

func autoplay() {
	for {
		fmt.Println("calling fakehuman()...")
		fakehuman()
//        decide()
		moves++
		fmt.Println("calling decide()...")
		decide()
//        fakehuman()
		moves++
		fmt.Println("calling who_won()...")
		who_won()
		if over == 1 {
			break
		}
	}
	fmt.Println("     ------------- Game over ----------------")
	fmt.Println()
	return
}

func fakehuman() {
	who_won()
	best := 1
	best = can_steal()
	if best == 0 {
		best = rand.Intn(6) + 8
	}
	if best > 13 {
		best = 8
	}
	if best < 8 {
		best = 13
	}
	fmt.Println("best was ", best)
	for {
		zero := autoturn(best)
		if zero != 0 {
			break
		}
	}
	fmt.Println("fakehuman() exited...")
	who_won()
	return
}

func autoturn(y int) int {
	hand := n[y]
	if hand == 0 {
		return 1
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
	who_won()
	fmt.Println("autoturn() exited...")
	return 0
}

func run() {
	initGrid()
	drawGrid()
	for q := 1; q < 9000; q++ {
		autoplay()
		pc[q] = n[0]
		hum[q] = n[7]
		r[q] = result
		//        print_stats()
		//        os.Exit(0)
		over = 0
		total = 0
		initGrid()
		drawGrid()
		moves = 0
	}
	print_stats()
	return
}

func print_stats() {
	fmt.Println("          ===================================================================")
	fmt.Println()
	var wins, losses, draws int
	for g := range r {
		if r[g] == "DRAW" {
			draws++
		}
		if r[g] == "HUMAN WON" {
			wins++
		}
		if r[g] == "PC WON" {
			losses++
		}
	}
	fmt.Println("Human won: " + strconv.Itoa(wins) + " |  PC won: " + strconv.Itoa(losses) + " |  Draws: " + strconv.Itoa(draws))
}

func pc_can_steal() int {
	for t := range 6 {
		t++
		for v := range 14 {
			k[v] = n[v]
		}
		foot := k[t]
		k[t] = 0
		for {
			t++
			if t == 14 {
				t = 0
			}
			k[t] += 1
			foot--
			if foot < 1 {
				break
			}
		}
		if k[t] == 1 {
			switch t {
			case 13:
				if k[1] != 0 {
					return 13
				}
				k[7] += k[1]
				k[7]++
				k[13]--
				k[1] = 0
			case 12:
				if k[2] != 0 {
					return 12
				}
				k[7] += k[2]
				k[7]++
				k[12]--
				k[2] = 0
			case 11:
				if k[3] != 0 {
					return 11
				}
				k[7] += k[3]
				k[7]++
				k[11]--
				k[3] = 0
			case 10:
				if k[4] != 0 {
					return 10
				}
				k[7] += k[4]
				k[7]++
				k[10]--
				k[4] = 0
			case 9:
				if k[5] != 0 {
					return 9
				}
				k[7] += k[5]
				k[7]++
				k[9]--
				k[5] = 0
			case 8:
				if k[6] != 0 {
					return 8
				}
				k[7] += k[6]
				k[7]++
				k[8]--
				k[6] = 0
			default:
				return 0
			}
		}
	}
	//  os.Exit(0)
	return 0
}

func decide() {
	who_won()
	best := 1
	best = pc_can_steal()
	if best == 0 {
		best = rand.Intn(6) + 1
	}
	if best > 6 {
		best = 1
	}
	if best < 1 {
		best = 6
	}
	fmt.Println("best was ", best)
	for {
		zero := compturn(best)
		if zero != 0 {
			break
		}
	}
	fmt.Println("decide() exited...")
	who_won()
	return
}
