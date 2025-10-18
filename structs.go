package main

import (
	"errors"
	"math/rand"
)

type Points struct {
	x, y int
}

type Snake struct {
	body      []Points
	direction string
}

type Apple struct {
	x, y int
}

func GenerateNewApple(n *int, snake *Snake) Apple {
	a := Apple{rand.Intn(*n), rand.Intn(*n)}
	_, isHaveBodyOfSnake := snake.IsSnakeInThisPlace(a.x, a.y)
	for {
		if isHaveBodyOfSnake {
			a = Apple{rand.Intn(*n), rand.Intn(*n)}
			_, isHaveBodyOfSnake = snake.IsSnakeInThisPlace(a.x, a.y)
		} else {
			return a
		}
	}
}

func GenerateNewSnake(n *int) Snake {
	s := Snake{body: make([]Points, 0, *n**n)}
	center := *n / 2
	s.body = append(s.body, Points{center, center})
	return s
}

func (snake *Snake) IsSnakeInThisPlace(x, y int) (isHead bool, isBodyOfSnake bool) {
	if snake.body[0].x == x && snake.body[0].y == y {
		return true, true
	}
	for i := 1; i < len(snake.body); i++ {
		coordX := snake.body[i].x
		coordY := snake.body[i].y

		if coordX == x && coordY == y {
			return false, true
		}
	}
	return false, false
}

func (snake *Snake) Move(apple *Apple, numbersOfColumns *int) error {
	lastIndex := len(snake.body) - 1
	for i := lastIndex; i != 0; i-- {
		snake.body[i] = snake.body[i-1]
	}
	switch snake.direction {
	case "w":
		_, isHaveBodyOfSnake := snake.IsSnakeInThisPlace(snake.body[0].x, snake.body[0].y-1)
		if isHaveBodyOfSnake {
			return errors.New("The snake crashed into itself :(")
		} else if snake.body[0].y-1 < 0 {
			return errors.New("The snake went out of bounds :(")
		}
		snake.body[0].y -= 1
		if snake.body[0].x == apple.x && snake.body[0].y == apple.y {
			snake.body = append(snake.body, Points{snake.body[lastIndex].x, snake.body[lastIndex].y + 1})
			*apple = GenerateNewApple(numbersOfColumns, snake)
		}
	case "s":
		_, isHaveBodyOfSnake := snake.IsSnakeInThisPlace(snake.body[0].x, snake.body[0].y+1)
		if isHaveBodyOfSnake {
			return errors.New("The snake crashed into itself :(")
		} else if snake.body[0].y+2 > *numbersOfColumns {
			return errors.New("The snake went out of bounds :(")
		}
		snake.body[0].y += 1
		if snake.body[0].x == apple.x && snake.body[0].y == apple.y {
			snake.body = append(snake.body, Points{snake.body[lastIndex].x, snake.body[lastIndex].y - 1})
			*apple = GenerateNewApple(numbersOfColumns, snake)
		}
	case "a":
		_, isHaveBodyOfSnake := snake.IsSnakeInThisPlace(snake.body[0].x-1, snake.body[0].y)
		if isHaveBodyOfSnake {
			return errors.New("The snake crashed into itself :(")
		} else if snake.body[0].x-1 < 0 {
			return errors.New("The snake went out of bounds :(")
		}
		snake.body[0].x -= 1
		if snake.body[0].x == apple.x && snake.body[0].y == apple.y {
			snake.body = append(snake.body, Points{snake.body[lastIndex].x + 1, snake.body[lastIndex].y})
			*apple = GenerateNewApple(numbersOfColumns, snake)
		}
	case "d":
		_, isHaveBodyOfSnake := snake.IsSnakeInThisPlace(snake.body[0].x+1, snake.body[0].y)
		if isHaveBodyOfSnake {
			return errors.New("The snake crashed into itself :(")
		} else if snake.body[0].x+2 > *numbersOfColumns {
			return errors.New("The snake went out of bounds :(")
		}
		snake.body[0].x += 1
		if snake.body[0].x == apple.x && snake.body[0].y == apple.y {
			snake.body = append(snake.body, Points{snake.body[lastIndex].x - 1, snake.body[lastIndex].y})
			*apple = GenerateNewApple(numbersOfColumns, snake)
		}
	}
	return nil
}
