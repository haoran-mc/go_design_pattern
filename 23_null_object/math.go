package null_object

type math struct {
	numberOfProfessors int
}

func (c *math) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *math) getName() string {
	return "math"
}
