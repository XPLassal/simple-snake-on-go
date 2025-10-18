package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	var numbersOfColumns int
	var gameAcceleration string
	GetNumberOfColumns(&numbersOfColumns)
	fmt.Print("Do you want the game speed to increase as you progress? (y/n): ")
	fmt.Scan(&gameAcceleration)
	fmt.Println()
	baseMilliseconds := 200 * time.Millisecond
	ticker := time.NewTicker(baseMilliseconds)
	if gameAcceleration == "y" {
		go func() {
			for {
				time.Sleep(2 * time.Second)
				baseMilliseconds -= 1 * time.Millisecond
				ticker = time.NewTicker(baseMilliseconds)
			}
		}()
	}
	snake := GenerateNewSnake(&numbersOfColumns)
	apple := GenerateNewApple(&numbersOfColumns, &snake)

	if err := keyboard.Open(); err != nil {
		fmt.Println(Red + Bold + "Error: " + err.Error() + Reset)
		return
	}
	defer keyboard.Close()

	keyCh := make(chan rune)
	go func() {
		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				fmt.Println(Red + Bold + "Error: " + err.Error() + Reset)
				return
			}
			keyCh <- char
		}
	}()

	for {
		RenderField(&numbersOfColumns, &apple, &snake)
		select {
		case <-ticker.C:
			err := snake.Move(&apple, &numbersOfColumns)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(BrightGreen+Bold+"Your score: ", len(snake.body), Reset)
				return
			}
		case direction := <-keyCh:
			if direction == 'q' {
				fmt.Println("Game stopped, bye! 🙃")
				return
			}
			snake.direction = string(direction)
		}
	}
}
