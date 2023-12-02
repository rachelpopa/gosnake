package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"log"
	"math/rand"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Salllllly", // TODO: Your Battlesnake username
		Color:      "#81D8D0",   // TODO: Choose color
		Head:       "rudolph",   // TODO: Choose head
		Tail:       "present",   // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

func upIsSafe(head Coord, badCoords []Coord) bool {
	if head.Y == 10 {
		return false
	}
	for i := 0; i < len(badCoords); i++ {
		if head.X == badCoords[i].X && head.Y+1 == badCoords[i].Y {
			return false
		}
	}
	return true
}

func downIsSafe(head Coord, badCoords []Coord) bool {
	if head.Y == 0 {
		return false
	}
	for i := 0; i < len(badCoords); i++ {
		if head.X == badCoords[i].X && head.Y-1 == badCoords[i].Y {
			return false
		}
	}
	return true
}

func rightIsSafe(head Coord, badCoords []Coord) bool {
	if head.X == 10 {
		return false
	}
	for i := 0; i < len(badCoords); i++ {
		if head.Y == badCoords[i].Y && head.X+1 == badCoords[i].X {
			return false
		}
	}
	return true
}

func leftIsSafe(head Coord, badCoords []Coord) bool {
	if head.X == 0 {
		return false
	}
	for i := 0; i < len(badCoords); i++ {
		if head.Y == badCoords[i].Y && head.X-1 == badCoords[i].X {
			return false
		}
	}
	return true
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {

	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	myHead := state.You.Head
	badSquares := state.You.Body

	for i := 0; i < len(state.Board.Snakes); i++ {
		badSquares = append(badSquares, state.Board.Snakes[i].Body...)
	}

	isMoveSafe["up"] = upIsSafe(myHead, badSquares)
	isMoveSafe["down"] = downIsSafe(myHead, badSquares)
	isMoveSafe["left"] = leftIsSafe(myHead, badSquares)
	isMoveSafe["right"] = rightIsSafe(myHead, badSquares)

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
