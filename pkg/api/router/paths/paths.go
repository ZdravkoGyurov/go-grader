package paths

const (
	Register = "/register"
	Login    = "/login"
	Logout   = "/logout"

	IDParam = "id"

	User       = "/user"
	UserWithID = "/user/{" + IDParam + "}"

	Course       = "/course"
	CourseWithID = "/course/{" + IDParam + "}"

	Assignment       = "/assignment"
	AssignmentWithID = "/assignment/{" + IDParam + "}"

	Submission       = "/submission"
	SubmissionWithID = "/submission/{" + IDParam + "}"
)
