package null_object

// compute science
type cs struct {
	numberOfProfessors int
}

func (c *cs) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *cs) getName() string {
	return "cs"
}
