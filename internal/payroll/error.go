package payroll

import "errors"

var (
	ErrPayrollExists = errors.New("payroll already processed for this period")
)
