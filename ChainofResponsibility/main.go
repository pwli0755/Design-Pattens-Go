package main

func main() {
	cashier := &cashier{}

	// set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	// set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	// set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	// patient visiting
	reception.execute(patient)

}
