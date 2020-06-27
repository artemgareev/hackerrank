package queens_attack

const (
	RightDiagonal       = 1
	LeftDiagonal        = 2
	HorizontalDirection = 3
	VerticalDirection   = 4
)

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x1, x2 int32) int32 {
	if x1 > x2 {
		return x2
	}
	return x1
}

// Complete the queensAttack function below.
//https://www.hackerrank.com/challenges/queens-attack-2/problem
func queensAttack(sizeX int32, obsCount int32, posX int32, posY int32, obstacles [][]int32) int32 {
	if posY == 0 {
		posY = 1
	}
	if posX == 0 {
		posY = 1
	}

	//get all direction spaces from posX,posY
	rightSpace := sizeX - posX
	leftSpace := Abs(rightSpace-sizeX) - 1
	upSpace := sizeX - posY
	downSpace := Abs(upSpace-sizeX) - 1

	minSpace := Min(upSpace, rightSpace)
	rightDiagonalUp := point{
		posX + minSpace,
		posY + minSpace,
	}

	minSpace = Min(downSpace, leftSpace)
	rightDiagonalDown := point{
		posX - minSpace,
		posY - minSpace,
	}

	minSpace = Min(upSpace, leftSpace)
	leftDiagonalUp := point{
		posX - minSpace,
		posY + minSpace,
	}

	minSpace = Min(downSpace, rightSpace)
	leftDiagonalDown := point{
		posX + minSpace,
		posY - minSpace,
	}

	_, _ = leftDiagonalDown, leftDiagonalUp

	currentPoint := point{posX, posY}
	lines := []vector2d{
		{currentPoint, rightDiagonalDown, RightDiagonal, nil},
		{currentPoint, rightDiagonalUp, RightDiagonal, nil},
		{currentPoint, leftDiagonalDown, LeftDiagonal, nil},
		{currentPoint, leftDiagonalUp, LeftDiagonal, nil},
		{currentPoint, point{1, posY}, HorizontalDirection, nil},
		{currentPoint, point{sizeX, posY}, HorizontalDirection, nil},
		{currentPoint, point{posX, sizeX}, VerticalDirection, nil},
		{currentPoint, point{posX, 1}, VerticalDirection, nil},
	}

	for _, obstacle := range obstacles {
		obstaclePoint := point{obstacle[0], obstacle[1]}
		for index, line := range lines {
			if line.isObstacleOnVector(obstaclePoint) {
				line.addNewObstacle(obstaclePoint)
			}
			lines[index] = line
		}
	}

	var result int32
	for _, line := range lines {
		result += line.getSizeWithObstacles()
	}

	return result
}

type point struct {
	x int32
	y int32
}

type vector2d struct {
	x0        point
	x1        point
	lineType  int32
	obstacles []point
}

func (v *vector2d) getSizeWithObstacles() int32 {
	if len(v.obstacles) == 0 {
		size := v.getSize() - 1
		return size
	}
	var minDistToObstacle int32
	for _, obstaclePoint := range v.obstacles {
		var distFromPosToObstacle int32
		if v.lineType == VerticalDirection {
			distFromPosToObstacle = Abs(v.x0.y - obstaclePoint.y)
		} else {
			distFromPosToObstacle = Abs(v.x0.x - obstaclePoint.x)
		}

		if minDistToObstacle == 0 {
			minDistToObstacle = distFromPosToObstacle
		} else if distFromPosToObstacle < minDistToObstacle {
			minDistToObstacle = distFromPosToObstacle
		}
	}

	return minDistToObstacle - 1
}

func (v *vector2d) getSize() int32 {
	var size int32
	if v.lineType == VerticalDirection {
		size = Abs(v.x0.y-v.x1.y) + 1
	} else {
		size = Abs(v.x0.x-v.x1.x) + 1
	}
	return size
}

func (v *vector2d) isObstacleOnVector(oPoint point) bool {
	//check if obstacle between x0 and x1, y0 and y1
	obstacleOnX := (oPoint.x >= v.x0.x && oPoint.x <= v.x1.x) ||
		(oPoint.x >= v.x1.x && oPoint.x <= v.x0.x)
	obstacleOnY := (oPoint.y >= v.x0.y && oPoint.y <= v.x1.y) ||
		(oPoint.y >= v.x1.y && oPoint.y <= v.x0.y)
	posEqualObstacle := v.x0.x == oPoint.x && v.x0.y == oPoint.y
	if !obstacleOnX || !obstacleOnY || posEqualObstacle {
		return false
	}
	switch v.lineType {
	case RightDiagonal, LeftDiagonal:
		return Abs(v.x0.x-oPoint.x) == Abs(v.x0.y-oPoint.y)
	case HorizontalDirection:
		return v.x0.y == oPoint.y
	case VerticalDirection:
		return v.x0.x == oPoint.x
	default:
		panic("unknown vector type")
	}
}

func (v *vector2d) addNewObstacle(obstaclePoint point) {
	if v.obstacles == nil {
		v.obstacles = []point{}
	}
	v.obstacles = append(v.obstacles, obstaclePoint)
}
