package model

type Employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Status     string `json:"status"`
	JoinedAt   string `json:"joinedAt"`
}

type Attendance struct {
	Date    string `json:"date"`
	Present int    `json:"present"`
	Late    int    `json:"late"`
	Absent  int    `json:"absent"`
}

type LeaveRequest struct {
	ID       string `json:"id"`
	Employee string `json:"employee"`
	Type     string `json:"type"`
	Dates    string `json:"dates"`
	Status   string `json:"status"`
}

type PayrollSummary struct {
	Period    string `json:"period"`
	Employees int    `json:"employees"`
	Gross     string `json:"gross"`
	Status    string `json:"status"`
}
