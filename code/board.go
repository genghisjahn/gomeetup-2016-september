//Board contains all the holes that contain the pegs
type Board struct {
	Holes      []Hole
	MoveChart  []string
	SolveMoves int
	Rows       int
}
