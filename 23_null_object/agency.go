package null_object

func createCollege1() *college {
	college := &college{}
	college.addDepartment("cs", 4)
	college.addDepartment("math", 5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("cs", 2)
	return college
}
