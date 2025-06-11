package attendance

import "errors"

var (
	ErrWeekendSubmission     = errors.New("attendance cannot be submitted on weekends")
	ErrDuplicateAttendance   = errors.New("attendance already submitted for this date")
	ErrAttendancePeriodUnset = errors.New("attendance period has not been set")
)
