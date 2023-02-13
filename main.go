package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	menuCounts := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var eaterID, menuID int
		_, err := fmt.Sscanf(scanner.Text(), "(%d, %d)", &eaterID, &menuID)
		if err != nil {
			fmt.Println("Error scanning line:", err)
			return
		}

		if _, ok := menuCounts[menuID]; ok {
			fmt.Println("Error: duplicate menu item consumed by same diner")
			return
		}
		menuCounts[menuID]++
	}

	topThree := make([]int, 0, 3)
	for menuID := range menuCounts {
		topThree = append(topThree, menuID)
		if len(topThree) > 2 {
			break
		}
	}

	fmt.Println("Top 3 menu items:")
	for i, menuID := range topThree {
		fmt.Printf("%d. %d\n", i+1, menuID)
	}
}
