//
// A simple rating implementation
//
// TODO(kashomon): Consider moving up a directory.
package simple

import "github.com/Kashomon/gorating"

// A simple player instance. Most practical systems will want to use more complex
// players. However, this can be useful for scripts. or for tests.
type Player struct {
	// ID of the player.
	id string

	// Rating of the player.
	score float64
}

// Get the UnqiueId, to make this a CompareablePlayer.
func (t *Player) UnqiueId() string {
	return t.id
}

// Get the PlayerRating, to make this a CompareablePlayer.
func (t *Player) NumericScore() float64 {
	return t.score
}

// Ensure that the Simple player satisfies the Player interface
var _ gorating.PlayerRating = &Player{}

type Game struct {
	p1     *Player
	p2     *Player
	result float64
}

func (t *Game) PlayerOne() gorating.PlayerRating {
	return gorating.PlayerRating(t.p1)
}

func (t *Game) PlayerTwo() gorating.PlayerRating {
	return gorating.PlayerRating(t.p2)
}

func (t *Game) GameResult() float64 {
	return t.result
}

var _ gorating.Game = &Game{}

type SimpleRatingSystem struct {
}

func (t *SimpleRatingSystem) RateAll([]gorating.Game) ([]gorating.PlayerRating, error) {
	return nil, nil
}

func (t *SimpleRatingSystem) Rate(gorating.Player, []gorating.Game) (gorating.PlayerRating, error) {
	return nil, nil
}

var _ gorating.RatingSystem = &SimpleRatingSystem{}
