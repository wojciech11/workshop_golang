package main

import (
	"fmt"

	"github.com/wojciech11/employee_mgmt/employee"
)

func main() {
	empl := employee.Employee{
		FirstName:          "Natalia",
		LastName:           "Kowalski",
		LeavesAllowedTotal: 26}

	empl.TakeHolidays(10)
	fmt.Printf("left: %v", empl.LeavesTaken)
}
