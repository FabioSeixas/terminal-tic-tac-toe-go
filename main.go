package main

import (
	"bufio"
	_ "fmt"
	"gameterminal/presenter"
	"os"
	"os/exec"
)

const SPACES_TOTAL = 9

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

func (s *State) initState() {

	s.currentSpace = 0
	s.spaces = make(map[int]SpaceValue)
	for i := 0; i < SPACES_TOTAL; i++ {
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
		s.presenter.MovePlayer(s.currentSpace)
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
		s.presenter.MovePlayer(s.currentSpace)
	}
}

func (s *State) moveDown() {
	if s.currentSpace > 5 {
		return
	}
	s.currentSpace += 3
	s.presenter.MovePlayer(s.currentSpace)
}

func (s *State) moveUp() {
	if s.currentSpace < 3 {
		return
	}
	s.currentSpace -= 3
	s.presenter.MovePlayer(s.currentSpace)
}

func (s *State) markSpace() {
	if s.spaces[s.currentSpace] == 0 {
		s.spaces[s.currentSpace] = 1
		s.presenter.WriteX()
	}
}

func main() {
	input, state := setup()

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
			state.markSpace()
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
	state.initState()

	state.presenter.DrawGame()

	return input, &state
}
