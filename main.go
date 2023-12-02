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

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		isMoveSafe["left"] = false

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		isMoveSafe["right"] = false

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		isMoveSafe["down"] = false

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		isMoveSafe["up"] = false
	}

	// Prevent from moving out of bounds
	if myHead.X == 10 {
		isMoveSafe["right"] = false
	} else if myHead.X == 0 {
		isMoveSafe["left"] = false
	}

	if myHead.Y == 10 {
		isMoveSafe["up"] = false
	} else if myHead.Y == 0 {
		isMoveSafe["down"] = false
	}

	// Prevent from colliding with itself
	mybody := state.You.Body
	for i := 0; i < len(mybody); i++ {
		if (mybody[i].X == myHead.X+1) && (mybody[i].Y == myHead.Y) {
			isMoveSafe["right"] = false
		} else if (mybody[i].X == myHead.X) && (mybody[i].Y == myHead.Y+1) {
			isMoveSafe["up"] = false
		} else if (mybody[i].X == myHead.X-1) && (mybody[i].Y == myHead.Y) {
			isMoveSafe["left"] = false
		} else if (mybody[i].X == myHead.X) && (mybody[i].Y == myHead.Y-1) {
			isMoveSafe["down"] = false
		}
	}

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	// opponents := state.Board.Snakes

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
