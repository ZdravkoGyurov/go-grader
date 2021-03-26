package api

const (
	CreateAssignmentPermission = "CREATE_ASSIGNMENT"
	ReadAssignmentPermission   = "READ_ASSIGNMENT"
	UpdateAssignmentPermission = "UPDATE_ASSIGNMENT"
	DeleteAssignmentPermission = "DELETE_ASSIGNMENT"

	CreateCoursePermission = "CREATE_COURSE"
	ReadCoursePermission   = "READ_COURSE"
	UpdateCoursePermission = "UPDATE_COURSE"
	DeleteCoursePermission = "DELETE_COURSE"

	CreateSubmissionPermission = "CREATE_SUBMISSION"
	ReadSubmissionPermission   = "READ_SUBMISSION"
	UpdateSubmissionPermission = "UPDATE_SUBMISSION"
	DeleteSubmissionPermission = "DELETE_SUBMISSION"

	ReadUserPermission   = "READ_USERS"
	UpdateUserPermission = "UPDATE_USER"
	DeleteUserPermission = "DELETE_USER"

	CreateRequestPermission = "CREATE_REQUEST"
	ReadRequestPermission   = "READ_REQUEST"
	UpdateRequestPermission = "UPDATE_REQUEST"
	DeleteRequestPermission = "DELETE_REQEUST"
)

var (
	TeacherPermissions = []string{
		CreateAssignmentPermission, ReadAssignmentPermission,
		UpdateAssignmentPermission, DeleteAssignmentPermission,
		CreateCoursePermission, ReadCoursePermission,
		UpdateCoursePermission, DeleteCoursePermission,
		CreateSubmissionPermission, ReadSubmissionPermission,
		UpdateSubmissionPermission, DeleteSubmissionPermission,
		ReadUserPermission, UpdateUserPermission, DeleteUserPermission,
		CreateRequestPermission, ReadRequestPermission,
		UpdateRequestPermission, DeleteRequestPermission,
	}

	StudentPermissions = []string{
		ReadAssignmentPermission, ReadCoursePermission,
		CreateSubmissionPermission, ReadSubmissionPermission,
		ReadUserPermission, UpdateUserPermission,
		CreateRequestPermission, ReadRequestPermission,
		UpdateRequestPermission, DeleteRequestPermission,
	}
)
