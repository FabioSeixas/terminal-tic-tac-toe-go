package presenter

import (
	"fmt"
	"time"
)

const MIN_SPACE_COLS = 10
const MIN_SPACE_ROWS = 5
const OFFSET = 1

type Presenter struct {
	cols         int
	rows         int
	spacesTotal  int
	spacesCoords map[int][2]int
}

func InitPresenter(cols int, rows int, spacesTotal int) Presenter {
	return Presenter{cols, rows, spacesTotal, make(map[int][2]int)}
}

func (p Presenter) initDisplay() {
	for i := 0; i < p.cols; i++ {
		fmt.Print("\n")
	}
}

func (p *Presenter) initSpaces() {
	topRow := (p.rows - (MIN_SPACE_ROWS * 3)) / 2
	leftMostCol := (p.cols - (MIN_SPACE_COLS * 3)) / 2
	delay()

	// first vertical line
	p.moveCursorDown(topRow)
	p.moveCursorRight(leftMostCol)
	p.moveCursorRight(MIN_SPACE_COLS)
	p.writeVerticalLine()
	p.writeVerticalLine()
	p.writeVerticalLine()

	// second vertical line
	p.moveCursorToReference()
	p.moveCursorDown(topRow)
	p.moveCursorRight(leftMostCol)
	p.moveCursorRight(MIN_SPACE_COLS)
	p.moveCursorRight(MIN_SPACE_COLS)
	p.moveCursorRight(OFFSET)
	p.writeVerticalLine()
	p.writeVerticalLine()
	p.writeVerticalLine()

	// first horizontal line
	p.moveCursorToReference()
	p.moveCursorDown(topRow)
	p.moveCursorDown(MIN_SPACE_ROWS)
	p.moveCursorRight(leftMostCol)
	p.writeHorizontalLine(MIN_SPACE_COLS)
	p.writeIntersection()
	p.writeHorizontalLine(MIN_SPACE_COLS)
	p.writeIntersection()
	p.writeHorizontalLine(MIN_SPACE_COLS)

	// second horizontal line
	p.moveCursorToReference()
	p.moveCursorDown(topRow)
	p.moveCursorDown(MIN_SPACE_ROWS)
	p.moveCursorDown(MIN_SPACE_ROWS)
	p.moveCursorDown(OFFSET)
	p.moveCursorRight(leftMostCol)
	p.writeHorizontalLine(MIN_SPACE_COLS)
	p.writeIntersection()
	p.writeHorizontalLine(MIN_SPACE_COLS)
	p.writeIntersection()
	p.writeHorizontalLine(MIN_SPACE_COLS)

	for i := 0; i < p.spacesTotal; i++ {
		p.moveCursorToReference()
		p.moveCursorDown(topRow)
		p.moveCursorRight(leftMostCol)

		if i < 3 {
			spaceStartX := leftMostCol + (MIN_SPACE_COLS * i) + (i + OFFSET)
			p.spacesCoords[i] = [2]int{spaceStartX, topRow}
		} else if i < 6 {
			factor := i - 3
			spaceStartX := leftMostCol + (MIN_SPACE_COLS * factor) + (factor + OFFSET)
			p.spacesCoords[i] = [2]int{spaceStartX, topRow + MIN_SPACE_ROWS + OFFSET}
		} else {
			factor := i - 6
			spaceStartX := leftMostCol + (MIN_SPACE_COLS * factor) + (factor + OFFSET)
			p.spacesCoords[i] = [2]int{spaceStartX, topRow + (MIN_SPACE_ROWS * 2) + (OFFSET * 2)}
		}
	}

	// Leave it here for future debug.
	// for i := 0; i < p.spacesTotal; i++ {
	// 	p.moveCursorToReference()
	// 	p.moveCursorRight(p.spacesCoords[i][0])
	// 	p.moveCursorDown(p.spacesCoords[i][1])
	// 	delay()
	// }
}

func (p Presenter) writeVerticalLine() {
	for i := 0; i < (MIN_SPACE_ROWS + OFFSET); i++ {
		fmt.Print("|")
		p.moveCursorDown(1)
		p.moveCursorLeft(1)
	}
}

func (p Presenter) writeHorizontalLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
}

func (p Presenter) writeIntersection() {
	fmt.Print("+")
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

func (p Presenter) MovePlayer(newPosition int) {
	p.moveCursorToReference()
	p.moveCursorRight(p.spacesCoords[newPosition][0])
	p.moveCursorDown(p.spacesCoords[newPosition][1])
}

func (p Presenter) WriteX() {
	p.moveCursorLeft(1)
	fmt.Print("**")

	for i := 0; i < 4; i++ {
		p.moveCursorDown(1)
		fmt.Print("**")
	}

	p.moveCursorUp(4)
	p.moveCursorLeft(2)
	fmt.Print("**")

	for i := 0; i < 4; i++ {
		p.moveCursorDown(1)
		p.moveCursorLeft(4)
		fmt.Print("**")
	}

	p.moveCursorUp(4)
	p.moveCursorLeft(1)
}

func (p Presenter) WriteO() {
	p.moveCursorRight(3)
	fmt.Print("**")
	p.moveCursorDown(1)
	fmt.Print("**")
	p.moveCursorDown(1)
	p.moveCursorLeft(1)
	fmt.Print("**")
	p.moveCursorLeft(3)
	p.moveCursorDown(1)
	fmt.Print("**")
	p.moveCursorLeft(4)
	p.moveCursorDown(1)
	fmt.Print("**")

	p.moveCursorLeft(4)
	p.moveCursorUp(1)
	fmt.Print("**")

	p.moveCursorLeft(3)
	p.moveCursorUp(1)
	fmt.Print("**")

	p.moveCursorLeft(1)
	p.moveCursorUp(1)
	fmt.Print("**")
}

func (p Presenter) AnnounceGameEnd() {
	fmt.Print("DRAW")
}

func (p Presenter) moveCursorToReference() {
	p.moveCursorLeft(999)
	p.moveCursorUp(999)
}

func (p Presenter) DrawGame() {

	p.initDisplay()
	p.moveCursorToReference()
	p.initSpaces()

	// start in the first position
	p.MovePlayer(0)
}

func delay() {
	time.Sleep(time.Millisecond * 500)
}
