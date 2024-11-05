package main

import (
	"bufio"
	"fmt"
	"gameterminal/presenter"
	"math/rand"
	"os"
	"os/exec"
)

const SPACES_TOTAL = 9
const SPACES_SIDE = 3

type SpaceValue int

const (
	Blank SpaceValue = iota
	X
	O
)

type State struct {
	presenter          presenter.Presenter
	spaces             map[int]SpaceValue
	currentPlayerSpace int
}

func (s *State) initState() {

	s.currentPlayerSpace = 0
	s.spaces = make(map[int]SpaceValue)
	for i := 0; i < SPACES_TOTAL; i++ {
		s.spaces[i] = Blank
	}
}

func (s *State) updateCurrentSpace(n int) {
	s.currentPlayerSpace = n
}

func (s *State) moveLeft() {
	switch s.currentPlayerSpace {

	case 0:
		return
	case 3:
		return
	case 6:
		return
	default:
		s.currentPlayerSpace--
		s.presenter.MovePlayer(s.currentPlayerSpace)
	}
}

func (s *State) moveRight() {
	switch s.currentPlayerSpace {

	case 2:
		return
	case 5:
		return
	case 8:
		return
	default:
		s.currentPlayerSpace++
		s.presenter.MovePlayer(s.currentPlayerSpace)
	}
}

func (s *State) moveDown() {
	if s.currentPlayerSpace > 5 {
		return
	}
	s.currentPlayerSpace += 3
	s.presenter.MovePlayer(s.currentPlayerSpace)
}

func (s *State) moveUp() {
	if s.currentPlayerSpace < 3 {
		return
	}
	s.currentPlayerSpace -= 3
	s.presenter.MovePlayer(s.currentPlayerSpace)
}

func (s *State) opponentRound() {
	for {
		opponentSpace := rand.Intn(len(s.spaces))
		if s.spaces[opponentSpace] == 0 {
			s.spaces[opponentSpace] = 2
			s.presenter.MovePlayer(opponentSpace)
			s.presenter.WriteO()
			s.presenter.MovePlayer(s.currentPlayerSpace)
			return
		}
	}
}

func (s *State) markSpace() bool {
	if s.spaces[s.currentPlayerSpace] == 0 {
		s.spaces[s.currentPlayerSpace] = 1
		s.presenter.WriteX()
		return true
	}
	return false
}

func (s *State) isGameEnd() bool {
	draw := s.checkForDraw()
	if draw {
		s.presenter.AnnounceGameEnd()
		return true
	}
	winner := s.checkForWinner()
	if winner != 0 {
		s.presenter.AnnounceWinner(int(winner))
		return true
	}
	return false

}

func (s *State) checkForDraw() bool {
	for _, value := range s.spaces {
		if value == 0 {
			return false
		}
	}
	return true
}

func (s *State) checkForWinner() SpaceValue {
	/* It is only necessary to iterate starting at spaces 0, 1, 2, 3 and 6 */
	for i, value := range s.spaces {
		win := false

		if value == 0 {
			continue
		}
		if i < 3 {
			win = s.evaluateVerticalPath(i, value)
			if win {
				return value
			}
		}

		if i%SPACES_SIDE == 0 {
			win = s.evaluateHorizontalPath(i, value)
			if win {
				return value
			}
		}

		if i == 0 || i == 3 {
			// evaluate sloped path
		}

	}
	return 0
}

func (s *State) evaluateHorizontalPath(i int, firstValue SpaceValue) bool {
	for j := i + 1; j < (i + 3); j++ {
		value := s.spaces[j]
		if value == 0 {
			return false
		}
		if value != firstValue {
			return false
		}
	}
	return true
}

func (s *State) evaluateVerticalPath(i int, firstValue SpaceValue) bool {
	for j := i + 3; j < SPACES_TOTAL; j += 3 {
		value := s.spaces[j]
		if value == 0 {
			return false
		}
		if value != firstValue {
			return false
		}
	}
	return true
}

func gameLoop(input *bufio.Reader, state *State) {
	for {
		switch byte, _ := input.ReadByte(); byte {

		case 104:
			state.moveLeft()
		case 106:
			state.moveDown()
		case 107:
			state.moveUp()
		case 108:
			state.moveRight()
		case 10:
			canGoOn := state.markSpace()
			if !canGoOn {
				continue
			}
			end := state.isGameEnd()
			if end {
				return
			}
			state.opponentRound()
			end = state.isGameEnd()
			if end {
				return
			}
		}
	}
}

func main() {
	gameRunning := true
	input, state := setup()

	for {
		if gameRunning {
			state.initState()
			state.presenter.DrawGame()

			gameLoop(input, state)
		}

		gameRunning = false

		fmt.Print("Play again? (y/n)\n")

		switch byte, _ := input.ReadByte(); byte {

		case 121:
			gameRunning = true

		case 110:
			fmt.Print("Bye Bye!")
			os.Exit(0)
		}
	}
}

func setup() (*bufio.Reader, *State) {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	file, err := os.OpenFile("/dev/tty", os.O_RDONLY, 0)
	if err != nil {
		panic("error while opening device file")
	}

	cols, rows := get_term_size(file.Fd())

	input := bufio.NewReader(file)

	p := presenter.InitPresenter(cols, rows, SPACES_TOTAL)

	state := State{presenter: p}

	return input, &state
}
