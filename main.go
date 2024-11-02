package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const maxY = 17
const squareSize = 5

func delay() {
	time.Sleep(time.Millisecond * 1000)
}

func initDisplay() {
	for i := 0; i < maxY; i++ {
		fmt.Print("\n")
	}
}

func writeVerticalLine() {
	for i := 0; i < maxY; i++ {
		fmt.Print("|")
		moveCursorDown(1)
		moveCursorLeft(1)
	}
}

func writeHorizontalLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
	moveCursorRight(1)
}

func moveCursorLeft(n int) {
	fmt.Print("\u001b[", n, "D")
}

func moveCursorRight(n int) {
	fmt.Print("\u001b[", n, "C")
}

func moveCursorUp(n int) {
	fmt.Print("\u001b[", n, "A")
}

func moveCursorDown(n int) {
	fmt.Print("\u001b[", n, "B")
}

func writeX() {
	moveCursorUp(2)
	moveCursorLeft(4)
	fmt.Print("**")
	moveCursorDown(1)
	fmt.Print("**")
	moveCursorDown(1)
	fmt.Print("**")
	moveCursorDown(1)
	fmt.Print("**")
	moveCursorDown(1)
	fmt.Print("**")

	moveCursorUp(4)
	moveCursorLeft(2)
	fmt.Print("**")
	moveCursorDown(1)
	moveCursorLeft(4)
	fmt.Print("**")
	moveCursorDown(1)
	moveCursorLeft(4)
	fmt.Print("**")
	moveCursorDown(1)
	moveCursorLeft(4)
	fmt.Print("**")
	moveCursorDown(1)
	moveCursorLeft(4)
	fmt.Print("**")

	centerCursor()
}

func centerCursor() {
	moveCursorLeft(999)
	moveCursorDown(999)

	moveCursorRight(15)
	moveCursorUp(9)
}

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	input := bufio.NewReader(os.Stdin)

	initDisplay()
	moveCursorUp(maxY)
	moveCursorRight(10)
	writeVerticalLine()
	moveCursorRight(11)
	moveCursorUp(maxY)
	writeVerticalLine()

	moveCursorLeft(999)
	moveCursorUp(maxY)
	moveCursorDown(squareSize)
	writeHorizontalLine(10)
	writeHorizontalLine(10)
	writeHorizontalLine(10)
	moveCursorDown(1)
	moveCursorLeft(999)
	moveCursorDown(squareSize)
	writeHorizontalLine(10)
	writeHorizontalLine(10)
	writeHorizontalLine(10)

	// buff := make([]byte, 8)

	centerCursor()

	for {
		switch byte, _ := input.ReadByte(); byte {

		case 104:
			moveCursorLeft(11)
		case 106:
			moveCursorDown(6)
		case 107:
			moveCursorUp(6)
		case 108:
			moveCursorRight(11)
		case 10:
			writeX()
		}
	}
}
