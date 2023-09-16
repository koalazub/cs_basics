package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	b "github.com/cs-basics/binary_search"
	s "github.com/cs-basics/compare_strings"
	l "github.com/cs-basics/linked_list"
	"golang.org/x/exp/slog"
)

func main() {
	menu()
}

var menuItems = []string{"a: binary search", "b: singly linked list", "c: compare strings"}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range menuItems {
		fmt.Println(v)
	}

	char, _, err := reader.ReadRune()
	if err != nil {
		slog.Error("Issue with character input:\n", err)
	}

	char = unicode.ToUpper(char)
	if char == 'A' {
		digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		res := b.BinarySearch(digits, 5)
		slog.Info("Result:\n", res)
	}
	if char == 'B' {
		l.SinglyLinkedList()
	}
	if char == 'C' {
		prefixes := []string{"potato", "pot", "pointing"}
		lcp := s.CompareStrings(prefixes)
		slog.Info("lowest common prefix is:\n ", lcp, 0)
	}
}
