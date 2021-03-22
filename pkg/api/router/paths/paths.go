package paths

const (
	Register = "/register"
	Login    = "/login"
	Logout   = "/logout"

	IDParam = "id"

	Assignment       = "/assignment"
	AssignmentWithID = "/assignment/{" + IDParam + "}"

	Course       = "/course"
	CourseWithID = "/course/{" + IDParam + "}"

	User       = "/user"
	UserWithID = "/user/{" + IDParam + "}"

	Submission = "/submission"
)
