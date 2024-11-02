package presenter

import (
	"fmt"
	"time"
)

type Presenter struct {
	MaxY      int
	SpaceSize int
}

func (p Presenter) initDisplay() {
	for i := 0; i < p.MaxY; i++ {
		fmt.Print("\n")
	}
}

func (p Presenter) writeVerticalLine() {
	for i := 0; i < p.MaxY; i++ {
		fmt.Print("|")
		p.moveCursorDown(1)
		p.moveCursorLeft(1)
	}
}

func (p Presenter) writeHorizontalLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
	p.moveCursorRight(1)
}

func (p Presenter) moveCursorLeft(n int) {
	fmt.Print("\u001b[", n, "D")
}

func (p Presenter) moveCursorRight(n int) {
	fmt.Print("\u001b[", n, "C")
}

func (p Presenter) moveCursorUp(n int) {
	fmt.Print("\u001b[", n, "A")
}

func (p Presenter) moveCursorDown(n int) {
	fmt.Print("\u001b[", n, "B")
}

func (p Presenter) MovePlayerLeft() {
	p.moveCursorLeft(11)
}

func (p Presenter) MovePlayerRight() {
	p.moveCursorRight(11)
}

func (p Presenter) MovePlayerUp() {
	p.moveCursorUp(6)
}

func (p Presenter) MovePlayerDown() {
	p.moveCursorDown(6)
}

func (p Presenter) WriteX() {
	p.moveCursorUp(2)
	p.moveCursorLeft(4)
	fmt.Print("**")
	p.moveCursorDown(1)
	fmt.Print("**")
	p.moveCursorDown(1)
	fmt.Print("**")
	p.moveCursorDown(1)
	fmt.Print("**")
	p.moveCursorDown(1)
	fmt.Print("**")

	p.moveCursorUp(4)
	p.moveCursorLeft(2)
	fmt.Print("**")
	p.moveCursorDown(1)
	p.moveCursorLeft(4)
	fmt.Print("**")
	p.moveCursorDown(1)
	p.moveCursorLeft(4)
	fmt.Print("**")
	p.moveCursorDown(1)
	p.moveCursorLeft(4)
	fmt.Print("**")
	p.moveCursorDown(1)
	p.moveCursorLeft(4)
	fmt.Print("**")

	// center in the space
	p.moveCursorUp(2)
	p.moveCursorRight(2)

}

func (p Presenter) DrawGame() {

	p.initDisplay()
	p.moveCursorUp(p.MaxY)
	p.moveCursorRight(10)
	p.writeVerticalLine()
	p.moveCursorRight(11)
	p.moveCursorUp(p.MaxY)
	p.writeVerticalLine()

	p.moveCursorLeft(999)
	p.moveCursorUp(p.MaxY)
	p.moveCursorDown(p.SpaceSize)
	p.writeHorizontalLine(10)
	p.writeHorizontalLine(10)
	p.writeHorizontalLine(10)
	p.moveCursorDown(1)
	p.moveCursorLeft(999)
	p.moveCursorDown(p.SpaceSize)
	p.writeHorizontalLine(10)
	p.writeHorizontalLine(10)
	p.writeHorizontalLine(10)

	p.moveCursorLeft(999)
	p.moveCursorUp(9)
	p.moveCursorRight(4)
}

func delay() {
	time.Sleep(time.Millisecond * 1000)
}
