package render

import (
	"strconv"
	"strings"

	. "github.com/XPLassal/simple-snake-on-go/structs"
)

func RenderField(numberOfColumns int, apple *Apple, snake *Snake, sb *strings.Builder) {
	sb.WriteString("\033[H")
	var isHaveApple, isHaveSnake, isHeadOfSnake, isDark bool

	sb.WriteString(Bold + "Your Score: " + strconv.Itoa(snake.GetLen()) + Reset + "\n")
	sb.WriteString(DrawBordersForY(numberOfColumns) + "\n")

	for y := range numberOfColumns {
		sb.WriteString(DrawBordersForX())
		for x := range numberOfColumns {
			coords := SetCoordinates(x, y)
			isHeadOfSnake, isHaveSnake = snake.Contains(coords)
			isHaveApple = apple.Contains(coords)
			sb.WriteString(DrawBg(isHaveApple, isDark, isHaveSnake, isHeadOfSnake))
			isDark = !isDark
		}
		sb.WriteString(DrawBordersForX() + "\n")
	}

	sb.WriteString(DrawBordersForY(numberOfColumns) + "\nTo exit, press q.\n")
}
