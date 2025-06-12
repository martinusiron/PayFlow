package shared

import (
	"context"
	"time"

	ac "github.com/martinusiron/PayFlow/internal/attendance/repository"
	au "github.com/martinusiron/PayFlow/internal/auditlog/usecase"
	od "github.com/martinusiron/PayFlow/internal/overtime/domain"
	or "github.com/martinusiron/PayFlow/internal/overtime/repository"
	pd "github.com/martinusiron/PayFlow/internal/payroll/domain"
	rd "github.com/martinusiron/PayFlow/internal/reimbursement/domain"
	rr "github.com/martinusiron/PayFlow/internal/reimbursement/repository"
	ur "github.com/martinusiron/PayFlow/internal/user/repository"
)

type Service struct {
	AuditUsecase      au.AuditLogService
	UserRepo          ur.UserRepository
	AttendanceRepo    ac.AttendanceRepository
	OvertimeRepo      or.OvertimeRepository
	ReimbursementRepo rr.ReimbursementRepository
}

type RequestContext struct {
	UserID    int
	RequestID string
	IPAddress string
}

func NewService(
	auditUc au.AuditLogService,
	userRepo ur.UserRepository,
	attendanceRepo ac.AttendanceRepository,
	overtimeRepo or.OvertimeRepository,
	reimbursementRepo rr.ReimbursementRepository,
) *Service {
	return &Service{
		AuditUsecase:      auditUc,
		UserRepo:          userRepo,
		AttendanceRepo:    attendanceRepo,
		OvertimeRepo:      overtimeRepo,
		ReimbursementRepo: reimbursementRepo,
	}
}

func (s *Service) CalculateAllEmployees(ctx context.Context, start, end time.Time) ([]pd.ProcessedPayroll, error) {
	users, err := s.UserRepo.GetAllEmployees(ctx)
	if err != nil {
		return nil, err
	}

	var summaries []pd.ProcessedPayroll

	for _, user := range users {
		// Calculate prorated salary
		totalWorkDays := countWeekdays(start, end)
		if totalWorkDays == 0 {
			continue
		}
		dailyRate := user.Salary / float64(totalWorkDays)

		attendedDays, err := s.AttendanceRepo.CountWeekdaysByUserID(ctx, user.ID, start, end)
		if err != nil {
			return nil, err
		}
		prorated := float64(attendedDays) * dailyRate

		// Calculate overtime
		overtimes, err := s.OvertimeRepo.GetOvertimeByUser(ctx, user.ID, start, end)
		if err != nil {
			return nil, err
		}

		// Calculate reimbursement
		reimbursements, err := s.ReimbursementRepo.GetReimbursementsByUser(ctx, user.ID, start, end)
		if err != nil {
			return nil, err
		}

		overtimeHours, overtimeAmount := s.calculateOvertimeTotal(overtimes, user.Salary/176)
		reimbursementTotal := s.calculateReimbursementTotal(reimbursements)

		summaries = append(summaries, pd.ProcessedPayroll{
			UserID:          user.ID,
			BaseSalary:      user.Salary,
			WorkdaysPresent: attendedDays,
			ProratedSalary:  prorated,
			OvertimeHours:   overtimeHours,
			OvertimePay:     overtimeAmount,
			Reimbursements:  reimbursementTotal,
			TotalTakeHome:   prorated + overtimeAmount + reimbursementTotal,
		})
	}

	return summaries, nil
}

// Count only weekdays in a date range
func countWeekdays(start, end time.Time) int {
	count := 0
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() >= time.Monday && d.Weekday() <= time.Friday {
			count++
		}
	}
	return count
}

func (s *Service) calculateOvertimeTotal(overtimes []od.Overtime, hourlyRate float64) (float64, float64) {
	var totalHours float64
	for _, ot := range overtimes {
		totalHours += ot.Hours
	}
	return totalHours, totalHours * hourlyRate * 2 // 2x multiplier for overtime
}

func (s *Service) calculateReimbursementTotal(reimbursements []rd.Reimbursement) float64 {
	var total float64
	for _, r := range reimbursements {
		total += r.Amount
	}

	return total
}
