package paths

const (
	Register = "/register"
	Login    = "/login"
	Logout   = "/logout"

	AssignmentsIDParam = "id"
	Assignments        = "/assignments"
	AssignmentsWithID  = "/assignments/{" + AssignmentsIDParam + "}"

	TestRun = "/testrun"
)
