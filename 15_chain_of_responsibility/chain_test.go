package chain_of_responsibility

import "testing"

func TestChain(t *testing.T) {
	cashier := &cashier{}

	// Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	// Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	// Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}

	// Patient visiting
	reception.execute(patient)
}
