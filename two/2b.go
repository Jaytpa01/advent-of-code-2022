package two

import (
	"log"
	"strings"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func B() int {
	strategyGuide := utils.OpenInputFile("./two/input.txt")

	rounds := strings.Split(strategyGuide, "\n")

	scoreTotal := 0

	for _, round := range rounds {
		shapes := strings.Split(round, " ")
		if len(shapes) != 2 {
			log.Fatalf("unexpected amount of shapes. expect 2, got %d", len(shapes))
		}

		opponentInput, outcome := oppenetInputToShape[shapes[0]], shapes[1]
		scoreTotal += DetermineScoreB(opponentInput, outcome)
	}

	return scoreTotal
}

func DetermineScoreB(opponentInput, outcome string) int {
	switch outcome {
	case "X": // you lose
		return LOSE_SCORE + shapeToScore[winMap[opponentInput]]

	case "Y": // you draw
		return DRAW_SCORE + shapeToScore[opponentInput]

	case "Z": // you win
		return WIN_SCORE + shapeToScore[loseMap[opponentInput]]

	default:
		return 0
	}
}

var loseMap = map[string]string{
	SCISSORS: ROCK,
	ROCK:     PAPER,
	PAPER:    SCISSORS,
}
