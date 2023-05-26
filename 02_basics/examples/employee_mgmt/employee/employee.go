package employee

type Employee struct {
	FirstName          string
	LastName           string
	LeavesAllowedTotal int
	LeavesTaken        int
}

func (empl *Employee) TakeHolidays(days int) bool {

	if empl.limitExceeded(days) {
		return false
	}
	empl.LeavesTaken = empl.LeavesTaken + days
	return true
}

func (empl *Employee) limitExceeded(days int) bool {
	return days+empl.LeavesTaken > empl.LeavesAllowedTotal
}
