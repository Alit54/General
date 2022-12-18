package gofunc

type Position struct {
	row    int
	column int
}

// Returns the minimun number of steps made from start to finish, using a UCS and a list of expanded nodes. (UCS + EXL)
/*func UCSexl(start Position, matrix map[Position]Node, goal []Position) int {
	var border []Position
	EXL := make(map[Position]bool)
	var pos, next Position
	border = append(border, start)
	for len(border) != 0 {
		pos = border[0]
		border = Alit.RemoveFirstElement(border)
		// ASCII problem: 'S' is << than 'a'
		if matrix[pos].high == 'S' {
			matrix[pos] = Node{'a', 0}
		}
		if !EXL[pos] {
			// Found a goal
			if matrix[pos].high == 'E' {
				return matrix[pos].step
			}
			EXL[pos] = true
			// look UP
			next = Position{row: pos.row, column: pos.column - 1}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
			// look DOWN
			next = Position{row: pos.row, column: pos.column + 1}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
			// look LEFT
			next = Position{row: pos.row - 1, column: pos.column}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
			// look RIGHT
			next = Position{row: pos.row + 1, column: pos.column}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
		}
	}
	return Alit.Min()
}*/
