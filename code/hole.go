type Hole struct {
	Row    int
	Col    int
	Peg    bool
	Status int
}

//Status is one of the values below
const (
	Dormant = iota
	Source
	Target
)
