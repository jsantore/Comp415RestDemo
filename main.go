package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Author: ")
	author, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	author = strings.TrimSpace(author)
	books, err := getBookData(author)
	if err != nil {
		fmt.Println("Error getting book data:", err)
		return
	}
	gui := buildBookGUI(books)
	gui.ShowAndRun()
}
