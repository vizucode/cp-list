package rockpaperscissor

func RockPaperScissor(p1, p2 string) string {

	if p1 == p2 {
		return "Draw!"
	} else if p1 == "rock" && p2 == "scissor" {
		return "Player 1 won!"
	} else if p1 == "scissor" && p2 == "rock" {
		return "Player 2 won!"
	} else if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	} else if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	} else if p1 == "paper" && p2 == "scissor" {
		return "Player 2 won!"
	} else if p1 == "scissor" && p2 == "paper" {
		return "Player 1 won!"
	}

	return ""
}
