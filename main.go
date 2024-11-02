package main

import (
	"bufio"
	"gameterminal/presenter"
	"os"
	"os/exec"
)

const maxY = 17
const squareSize = 5

type SpaceValue int

const (
	Blank SpaceValue = iota
	X
	O
)

type State struct {
	presenter    presenter.Presenter
	spaces       map[int]SpaceValue
	currentSpace int
}

func (s *State) initState(totalSpaces int) {

	s.currentSpace = 0
	s.spaces = make(map[int]SpaceValue)
	for i := 0; i < totalSpaces; i++ {
		s.spaces[i] = Blank
	}
}

func (s *State) updateCurrentSpace(n int) {
	s.currentSpace = n
}

func (s *State) moveLeft() {
	switch s.currentSpace {

	case 0:
		return
	case 3:
		return
	case 6:
		return
	default:
		s.currentSpace--
		s.presenter.MovePlayerLeft()
	}
}

func (s *State) moveRight() {
	switch s.currentSpace {

	case 2:
		return
	case 5:
		return
	case 8:
		return
	default:
		s.currentSpace++
		s.presenter.MovePlayerRight()
	}
}

func (s *State) moveDown() {
	if s.currentSpace > 5 {
		return
	}
	s.currentSpace += 3
	s.presenter.MovePlayerDown()
}

func (s *State) moveUp() {
	if s.currentSpace < 3 {
		return
	}
	s.currentSpace -= 3
	s.presenter.MovePlayerUp()
}

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	input := bufio.NewReader(os.Stdin)

	p := presenter.Presenter{MaxY: maxY, SpaceSize: squareSize}
	state := State{presenter: p}
	state.initState(6)

	state.presenter.DrawGame()

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
			// writeX()
		}
	}
}
