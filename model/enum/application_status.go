package enum

type ApplicationStatus string

const (
	NOT_ASSESSED            ApplicationStatus = "NOT_ASSESSED"
	PENDING_INTERVIEW       ApplicationStatus = "PENDING_INTERVIEW"
	ACCEPTED                ApplicationStatus = "ACCEPTED"
	REJECTED                ApplicationStatus = "REJECTED"
	ENROLLED                ApplicationStatus = "ENROLLED"
	IN_PROGRESS             ApplicationStatus = "IN_PROGRESS"
	FINISHED                ApplicationStatus = "FINISHED"
	DROPPED_OUT             ApplicationStatus = "DROPPED_OUT"
	INTRO_COURSE_PASSED     ApplicationStatus = "INTRO_COURSE_PASSED"
	INTRO_COURSE_NOT_PASSED ApplicationStatus = "INTRO_COURSE_NOT_PASSED"
)
