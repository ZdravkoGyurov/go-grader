# Mongo DB Model

<!-- ### Session
- ID     string
- UserID string
  - CreateSession(...)
  - ReadSession(sessionID string)
  - DeleteSession(sessionID string) -->

<!-- ### User
- ID             string
- Username       string
- GithubUsername string
- Fullname       string
- Password       string
- Permissions    []string
- Disabled       bool
- []CourseID     []string
  - CreateUser(...)
  - ReadUserByID(userID string)
  - ReadUserByUsername(username string)
  - ReadAllUsers(courseID string)
  - UpdateUser(userID string)
  - DeleteUser(userID string) -->

<!-- ### Submission
- ID           string
- Results      string?
- Status       string
- UserID       string <--+-< unique by user and assignment
- AssignmentID string <--+
  - CreateSubmission(...)
  - ReadSubmission(submissionID string)
  - ReadAllSubmissions(userID string, assignmentID string)
  - UpdateSubmission(submissionID string)
  - DeleteSubmission(submissionID string) -->

<!-- ### Assignment
- ID          string
- Name        string
- Description string
- CreatedOn   time.Duration
- DueDate     time.Duration
- CourseID    string
  - CreateAssignment(...)
  - ReadAssignment(assignmentID string)
  - ReadAllAssignments(courseID string)
  - UpdateAssignment(assignmentID string)
  - DeleteAssignment(assignmentID string) -->

### Course
- ID             string
- Name           string
- Description    string
- GithubRepoName string
  - CreateCourse(...)
  - ReadCourse(courseID string)
  - ReadAllCourses()
  - UpdateCourse(courseID string, ...)
  - DeleteCourse(courseID string)

### Request
- ID          string
- Name        string
- Description string
- Type        string - ask for permissions / ask to join course
- Permissions []string
- CourseID    string
- Status      string
- UserID
  - CreateRequest(...)
  - ReadRequest(requestID string)
  - ReadAllRequests(userID string)
  - ReadAllRequests()
  - UpdateRequest(requestID string)
  - DeleteRequest(requestID string)