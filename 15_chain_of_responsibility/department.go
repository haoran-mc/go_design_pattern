package chain_of_responsibility

type department interface {
	execute(*patient)
	setNext(department)
}
