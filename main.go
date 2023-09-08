package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	b "github.com/cs-basics/binary_search"
	l "github.com/cs-basics/linked_list"
	"golang.org/x/exp/slog"
)

func main() {
	menu()
}

var menuItems = []string{"a: binary search", "b: singly linked list"}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range menuItems {
		fmt.Println(v)
	}

	char, _, err := reader.ReadRune()
	if err != nil {
		slog.Error("Issue with character input:", err)
	}

	char = unicode.ToUpper(char)
	if char == 'A' {
		digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		res := b.BinarySearch(digits, 5)
		slog.Info("Result:", res)
	}
	if char == 'B' {
		l.SinglyLinkedList()
	}
}
