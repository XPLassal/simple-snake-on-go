package main

import (
	"fmt"
	"strings"
)

func RenderField(size *int, apple *Apple, snake *Snake) {
	ClearConsole()
	var field strings.Builder
	isDark := true
	var isHaveApple bool
	var isHaveSnake bool
	var isHeadOfSnake bool
	field.WriteString(DrawBordersForY(*size) + "\n")
	for y := 0; y < *size; y++ {
		field.WriteString(DrawBordersForX())
		for x := 0; x < *size; x++ {
			isHeadOfSnake, isHaveSnake = snake.IsSnakeInThisPlace(x, y)
			isHaveApple = y == apple.y && x == apple.x
			field.WriteString(DrawBg(isHaveApple, isDark, isHaveSnake, isHeadOfSnake))
			isDark = !isDark
		}
		field.WriteString(DrawBordersForX() + "\n")
	}
	field.WriteString(DrawBordersForY(*size))
	fmt.Println(field.String(), "\nTo exit, press q.")
}
