package queens_attack

import "sync"

const (
	RightDiagonal       = 1
	LeftDiagonal        = 2
	HorizontalDirection = 3
	VerticalDirection   = 4
)

type ticker func(posX int32, posY int32, direction int32) (int32, int32)

// Complete the queensAttack function below.
//https://www.hackerrank.com/challenges/queens-attack-2/problem
func queensAttack(sizeX int32, obsCount int32, posX int32, posY int32, obstacles [][]int32) int32 {
	var wg sync.WaitGroup
	countersChan := make(chan int32, 8)
	sizeY := sizeX
	counter := func(posX int32, posY int32, sizeX int32, sizeY int32, direction int32, fn ticker) {
		wg.Add(1)
		defer wg.Done()

		if posX == 0 {
			posX = sizeX
		}
		if posY == 0 {
			posY = sizeY
		}
		var counter int32
	Exit:
		for {
			posX, posY = fn(posX, posY, direction)

			for _, obstacle := range obstacles {
				obX, obY := obstacle[0], obstacle[1]
				if posX == obX && posY == obY {
					break Exit
				}
			}
			if posX < 1 || posX > sizeX || posY < 1 || posY > sizeY {
				break Exit
			}
			counter++
		}
		countersChan <- counter
	}

	go counter(posX, posY, sizeX, sizeY, RightDiagonal, tickRightByDirection)
	go counter(posX, posY, sizeX, sizeY, RightDiagonal, tickLeftByDirection)
	go counter(posX, posY, sizeX, sizeY, LeftDiagonal, tickRightByDirection)
	go counter(posX, posY, sizeX, sizeY, LeftDiagonal, tickLeftByDirection)
	go counter(posX, posY, sizeX, sizeY, HorizontalDirection, tickRightByDirection)
	go counter(posX, posY, sizeX, sizeY, HorizontalDirection, tickLeftByDirection)
	go counter(posX, posY, sizeX, sizeY, VerticalDirection, tickRightByDirection)
	go counter(posX, posY, sizeX, sizeY, VerticalDirection, tickLeftByDirection)

	wg.Wait()

	var queenPossibleSteps int32
	for i := 0; i < 8; i++ {
		queenPossibleSteps += <-countersChan
	}
	close(countersChan)

	return queenPossibleSteps
}

func tickRightByDirection(posX int32, posY int32, direction int32) (int32, int32) {
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

func tickLeftByDirection(posX int32, posY int32, direction int32) (int32, int32) {
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
