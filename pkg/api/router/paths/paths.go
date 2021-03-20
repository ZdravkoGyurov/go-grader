package paths

const (
	Register = "/register"
	Login    = "/login"
	Logout   = "/logout"

	AssignmentIDParam = "id"
	Assignment        = "/assignment"
	AssignmentWithID  = "/assignment/{" + AssignmentIDParam + "}"

	Submission = "/submission"
)
