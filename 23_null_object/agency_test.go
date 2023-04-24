package null_object

import (
	"fmt"
	"testing"
)

func TestNullObject(t *testing.T) {
	college1 := createCollege1()
	college2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"cs", "math", "civil", "electronics"}

	for _, deparmentName := range departmentArray {
		d := college1.getDepartment(deparmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	totalProfessors = 0
	for _, deparmentName := range departmentArray {
		d := college2.getDepartment(deparmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}
