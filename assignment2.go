package main

import (
	"fmt"
	"math"
)

type Employee struct {
	id int
	firstname string
	lastname string
	dob string
	leave_balance int
	rating string
	salaryDetails
	contactDetails
}

type salaryDetails struct{
	variable_payout float64
	basic float64
	hra float64
}

type contactDetails struct {
	city string
	mobile string
	email string
}

func main() {
	fmt.Println("Create Employee")
	emp:= createEmployee("Girish", "Warrier", "04121990", "B", "Pune", "1234567890",
		"girish.warrier@calsoftinc.com",80403, 27, 10000, 24000, 13000)
	fmt.Println("Read Employee Details")
	emp.readEmployee()
    fmt.Println("Update Employee Details")
	emp.updateEmployee(15000)
	emp.readEmployee()
	netSalary:= emp.calculateNetSalary(24)
	fmt.Println("Net Salary: ", netSalary)
	newSalary:= emp.calculateAppraisal()
	fmt.Println("New Salary after appraisal: ", newSalary)
	fmt.Printf("Emergency Details - Mobile: %s, City: %s", emp.getEmergencyDetails().mobile, emp.getEmergencyDetails().city)
	fmt.Println()
	frndsDict:= getFriends()
	fmt.Println("Friends list ", frndsDict)
	frndsRemove := []int{80500, 81234}
	fights(frndsDict, frndsRemove)
	fmt.Println("Updated Friends list after fight ", frndsDict)
}

func createEmployee(firstname, lastname, dob, rating, city, mobile, email string, id, leave_balance int, variable_payout, basic, hra float64) (Employee){
	employee:= Employee{id: id, firstname: firstname, lastname: lastname, dob: dob, leave_balance: leave_balance, rating: rating,
		salaryDetails:  salaryDetails{variable_payout: variable_payout, basic: basic, hra: hra},
		contactDetails: contactDetails{city: city, mobile: mobile, email: email},
	}
	return employee
}

func (e Employee) readEmployee () {
	fmt.Printf("Employee = %v \n", e)
}

func (e *Employee) updateEmployee (variable_payout float64) {
	e.variable_payout = variable_payout
}

func (e *Employee) calculateNetSalary (leaves float64) float64{
	lb:= float64(e.leave_balance) - leaves
	newlb:= math.Abs(lb)
	var newBasic float64
	if (lb < 0) {
		lb = 0
		newBasic = e.basic/30 * float64(30.0 - newlb)
	} else {
		newBasic = e.basic
	}
	e.leave_balance = int(lb)
	totalSalary:= e.variable_payout + newBasic + e.hra
	netSalary:= totalSalary - 200
	return netSalary
}

func (e Employee) calculateAppraisal () float64{
	var newBasic float64
	if (e.rating == "A"){
		newBasic = e.basic + e.basic * 0.2
	} else if (e.rating == "B") {
		newBasic = e.basic + e.basic * 0.1
	}
	newSalary:= newBasic + e.variable_payout + e.hra
	return newSalary
}

func (e Employee) getEmergencyDetails () contactDetails{
	return e.contactDetails
}

func getFriends () map[int]string{
	emp1:= createEmployee("Bobby", "Mohanty", "04121990", "A", "Pune", "1234567890",
		"bobby.mohanty@calsoftinc.com",80404, 27, 15000, 30000, 13000)
	emp2:= createEmployee("Amol", "Pawar", "04121990", "A", "Pune", "1234567890",
		"amol.pawar@calsoftinc.com",80405, 23, 14000, 25000, 11000)
	emp3:= createEmployee("Shahrin", "", "04121990", "A", "Pune", "1234567890",
		"shahrin.s@calsoftinc.com",80505, 30, 10000, 20000, 8000)
	emp4:= createEmployee("Madhav", "Mahajan", "04121990", "A", "Pune", "1234567890",
		"madhav.mahajan@calsoftinc.com",80610, 12, 15000, 35000, 13000)
	emp5:= createEmployee("Ameya", "Raut", "04121990", "A", "Pune", "1234567890",
		"ameya.raut@calsoftinc.com",81355, 27, 12000, 45000, 13000)
	emp6:= createEmployee("Pannag", "Swain", "04121990", "A", "Pune", "1234567890",
		"pannag.swain@calsoftinc.com",80435, 27, 14000, 25000, 11000)
	emp7:= createEmployee("Al", "A2", "04121990", "A", "Pune", "1234567890",
		"al.a2@calsoftinc.com",80500, 27, 14000, 25000, 11000)
	emp8:= createEmployee("A3", "A4", "04121990", "A", "Pune", "1234567890",
		"a3.a4@calsoftinc.com",80555, 27, 14000, 25000, 11000)
	emp9:= createEmployee("B1", "B2", "04121990", "A", "Pune", "1234567890",
		"b1.b2@calsoftinc.com",80445, 27, 14000, 25000, 11000)
	emp10:= createEmployee("c1", "c2", "04121990", "A", "Pune", "1234567890",
		"c1.c2@calsoftinc.com",81234, 27, 14000, 25000, 11000)

	friendsDetails := map[int]string{
		emp1.id: emp1.firstname + emp1.lastname,
		emp2.id: emp2.firstname + emp2.lastname,
		emp3.id: emp3.firstname + emp3.lastname,
		emp4.id: emp4.firstname + emp4.lastname,
		emp5.id: emp5.firstname + emp5.lastname,
		emp6.id: emp6.firstname + emp6.lastname,
		emp7.id: emp7.firstname + emp7.lastname,
		emp8.id: emp8.firstname + emp8.lastname,
		emp9.id: emp9.firstname + emp9.lastname,
		emp10.id: emp10.firstname + emp10.lastname,
	}
	return friendsDetails
}

func fights(m map[int]string, ids []int){
	for _, value:= range ids{
		delete(m, value)
	}
}