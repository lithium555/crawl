package main

import (
	"fmt"
	"time"
)

type Ball struct{}

func player(n int, team string, field chan Ball, pass chan Ball, gate chan Ball) {
	for {
		var ball Ball
		select {
		case ball = <-field:
			fmt.Printf("%s-%d intercepts the ball!\n", team, n)
		case ball = <-pass:
			fmt.Printf("%s-%d received friendly pass!\n", team, n)
		}
		select {
		case pass <- ball:
			fmt.Printf("%s-%d passed the ball\n", team, n)
		case field <- ball:
			fmt.Printf("%s-%d lost the ball\n", team, n)
		case gate <- ball:
			fmt.Printf("%s-%d strikes the gate\n", team, n)
		}
	}
}

func main() {
	fmt.Println("\n[ Welcome ladies and gentlemans!! ]")
	const (
		teamA = "Red"
		teamB = "Blue"
	)
	fmt.Printf("[ Today is a match between %s vs %s ]\n", teamA, teamB)
	fmt.Println("let us prepare...\n")

	passA := make(chan Ball)
	passB := make(chan Ball)
	gateA := make(chan Ball)
	gateB := make(chan Ball)

	for i := 0; i < 6; i++ {
		team, pass, lost, gate := teamA, passA, passB, gateB
		if i%2 == 1 {
			team, pass, lost, gate = teamB, passB, passA, gateA
		}
		fmt.Printf("p-%d is in %s\n", i+1, team)
		go player(i+1, team, lost, pass, gate)
	}
	fmt.Println("\nlets GO!\n")
	giveBall := func(b Ball) {
		select {
		case passA <- b:
			fmt.Println(teamA, "receives the ball!")
		case passB <- b:
			fmt.Println(teamB, "receives the ball!")
		}
	}
	t := time.After(time.Millisecond / 8)
	scoreA := 0
	scoreB := 0
	giveBall(Ball{})
	for {
		select {
		case ball := <-gateA:
			scoreB++
			fmt.Printf("\nGOAL! score is %d:%d\n\n", scoreA, scoreB)
			giveBall(ball)
		case ball := <-gateB:
			scoreA++
			fmt.Printf("\nGOAL! score is %d:%d\n\n", scoreA, scoreB)
			giveBall(ball)
		case <-t:
			fmt.Println("\nTIMEOUT!!\n")
			select {
			case <-passA:
			case <-passB:
			}
			win := teamA + " wins!"
			if scoreA == scoreB {
				win = "its a DRAW!"
			} else if scoreA < scoreB {
				win = teamB + " wins!"
			}
			fmt.Printf("\nGAME OVER! %s final score is %d:%d\n", win, scoreA, scoreB)
			return
		}
	}
}
