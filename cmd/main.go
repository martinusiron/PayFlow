package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/martinusiron/PayFlow/internal/middleware"

	asdel "github.com/martinusiron/PayFlow/internal/adminsummary/delivery/http"
	asrepo "github.com/martinusiron/PayFlow/internal/adminsummary/repository"
	asuc "github.com/martinusiron/PayFlow/internal/adminsummary/usecase"

	atdel "github.com/martinusiron/PayFlow/internal/attendance/delivery/http"
	attrepo "github.com/martinusiron/PayFlow/internal/attendance/repository"
	atuc "github.com/martinusiron/PayFlow/internal/attendance/usecase"

	otdel "github.com/martinusiron/PayFlow/internal/overtime/delivery/http"
	otrepo "github.com/martinusiron/PayFlow/internal/overtime/repository"
	otuc "github.com/martinusiron/PayFlow/internal/overtime/usecase"

	redel "github.com/martinusiron/PayFlow/internal/reimbursement/delivery/http"
	reborepo "github.com/martinusiron/PayFlow/internal/reimbursement/repository"
	reuc "github.com/martinusiron/PayFlow/internal/reimbursement/usecase"

	paydel "github.com/martinusiron/PayFlow/internal/payroll/delivery/http"
	payrepo "github.com/martinusiron/PayFlow/internal/payroll/repository"
	payuc "github.com/martinusiron/PayFlow/internal/payroll/usecase"

	psdel "github.com/martinusiron/PayFlow/internal/payslip/delivery/http"
	psrepo "github.com/martinusiron/PayFlow/internal/payslip/repository"
	psuc "github.com/martinusiron/PayFlow/internal/payslip/usecase"

	alrepo "github.com/martinusiron/PayFlow/internal/auditlog/repository"
	aluc "github.com/martinusiron/PayFlow/internal/auditlog/usecase"

	userrepo "github.com/martinusiron/PayFlow/internal/user/repository"

	authdel "github.com/martinusiron/PayFlow/internal/auth/delivery/http"
	authrepo "github.com/martinusiron/PayFlow/internal/auth/repository"
	authuc "github.com/martinusiron/PayFlow/internal/auth/usecase"

	"github.com/martinusiron/PayFlow/internal/shared"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dbURL := os.Getenv("DB_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if dbURL == "" || jwtSecret == "" {
		log.Fatal("DB_URL and JWT_SECRET must be set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// ===== Repositories =====
	userRepo := userrepo.NewUserRepository(db)
	attRepo := attrepo.NewAttendanceRepository(db)
	otRepo := otrepo.NewOvertimeRepository(db)
	rebRepo := reborepo.NewReimbursementRepository(db)
	payRepo := payrepo.NewPayrollRepository(db)
	psRepo := psrepo.NewPayslipRepository(db)
	alRepo := alrepo.NewAuditLogRepository(db)
	asRepo := asrepo.NewAdminSummaryRepository(db)
	authRepo := authrepo.NewAuthRepository(db)

	// ===== Usecases =====
	auditUC := aluc.NewAuditLogUsecase(alRepo)
	authUC := authuc.NewAuthUsecase(authRepo, jwtSecret)
	sharedService := shared.NewService(auditUC, userRepo, attRepo, otRepo, rebRepo)

	attUC := atuc.NewAttendanceUsecase(attRepo, auditUC)
	otUC := otuc.NewOvertimeUsecase(otRepo, auditUC)
	rebUC := reuc.NewReimbursementUsecase(rebRepo, auditUC)
	payUC := payuc.NewPayrollUsecase(payRepo, sharedService)
	psUC := psuc.NewPayslipUsecase(psRepo)
	asUC := asuc.NewAdminSummaryUsecase(asRepo)

	// ===== Handlers =====
	authHandler := authdel.NewAuthHandler(authUC)
	attHandler := atdel.NewAttendanceHandlerr(attUC, sharedService)
	otHandler := otdel.NewOvertimeHandler(otUC)
	rebHandler := redel.NewReimbursementHandler(rebUC)
	payHandler := paydel.NewPayrollHandler(payUC)
	psHandler := psdel.NewPayslipHandler(psUC)
	asHandler := asdel.NewAdminSummaryHandler(asUC)

	// ===== Middleware =====
	authMiddleware := middleware.NewAuthMiddleware(authUC)

	// ===== Routing =====
	mux := http.NewServeMux()

	// === Auth routes (no middleware) ===
	// mux.HandleFunc("/api/auth/signup", authHandler.Signup)
	mux.HandleFunc("/api/auth/login", authHandler.Login)

	// === Protected Routes ===
	mux.Handle("/api/attendance/submit", authMiddleware.JWTAuth(http.HandlerFunc(attHandler.SubmitAttendance)))
	mux.Handle("/api/overtime/submit", authMiddleware.JWTAuth(http.HandlerFunc(otHandler.SubmitOvertime)))
	mux.Handle("/api/reimbursement/submit", authMiddleware.JWTAuth(http.HandlerFunc(rebHandler.SubmitReimbursement)))
	mux.Handle("/api/payroll/run", authMiddleware.JWTAuth(http.HandlerFunc(payHandler.RunPayroll)))
	mux.Handle("/api/payslip/get", authMiddleware.JWTAuth(http.HandlerFunc(psHandler.GetMyPayslips)))
	mux.Handle("/api/summary/admin", authMiddleware.JWTAuth(http.HandlerFunc(asHandler.GenerateSummary)))

	// Swagger doc (optional static)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs/swagger"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
