package queens_attack

const (
	RightDiagonal       = 1
	LeftDiagonal        = 2
	HorizontalDirection = 3
	VerticalDirection   = 4
)

type ticker func(posX int32, posY int32, direction int32) (int32, int32)

// Complete the queensAttack function below.
func queensAttack(posX int32, posY int32, sizeX int32, sizeY int32, obstacles [][]int32) int32 {
	counter := func(posX int32, posY int32, sizeX int32, sizeY int32, direction int32, fn ticker) int32 {
		var counter int32
		for {
			posX, posY = fn(posX, posY, direction)

			if posX < 1 || posX > sizeX || posY < 1 || posY > sizeY {
				break
			}
			counter++
		}
		return counter
	}

	a := counter(posX, posY, sizeX, sizeY, RightDiagonal, incrementByDirection) +
		counter(posX, posY, sizeX, sizeY, LeftDiagonal, incrementByDirection) +
		counter(posX, posY, sizeX, sizeY, HorizontalDirection, incrementByDirection) +
		counter(posX, posY, sizeX, sizeY, VerticalDirection, incrementByDirection) +
		counter(posX, posY, sizeX, sizeY, RightDiagonal, decrementByDirection) +
		counter(posX, posY, sizeX, sizeY, LeftDiagonal, decrementByDirection) +
		counter(posX, posY, sizeX, sizeY, HorizontalDirection, decrementByDirection) +
		counter(posX, posY, sizeX, sizeY, VerticalDirection, decrementByDirection)

	return a
}

func incrementByDirection(posX int32, posY int32, direction int32) (int32, int32) {
	switch direction {
	case RightDiagonal:
		posX++
		posY++
	case LeftDiagonal:
		posX++
		posY--
	case HorizontalDirection:
		posX++
	case VerticalDirection:
		posY++
	}

	return posX, posY
}

func decrementByDirection(posX int32, posY int32, direction int32) (int32, int32) {
	switch direction {
	case RightDiagonal:
		posX--
		posY--
	case LeftDiagonal:
		posX--
		posY++
	case HorizontalDirection:
		posX--
	case VerticalDirection:
		posY--
	}

	return posX, posY
}
