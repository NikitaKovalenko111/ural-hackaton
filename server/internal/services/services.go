package services

import (
	"ural-hackaton/internal/config"
	admin_service "ural-hackaton/internal/services/handlers/admin"
	auth_service "ural-hackaton/internal/services/handlers/auth"
	booking_service "ural-hackaton/internal/services/handlers/booking"
	"ural-hackaton/internal/services/handlers/email"
	event_service "ural-hackaton/internal/services/handlers/event"
	hub_service "ural-hackaton/internal/services/handlers/hub"
	mentor_service "ural-hackaton/internal/services/handlers/mentor"
	requests_service "ural-hackaton/internal/services/handlers/request"
	user_service "ural-hackaton/internal/services/handlers/user"
	"ural-hackaton/internal/storage/repositories"
)

type Services struct {
	UserService    *user_service.UserService
	HubService     *hub_service.HubService
	AdminService   *admin_service.AdminService
	MentorService  *mentor_service.MentorService
	RequestService *requests_service.RequestService
	EventService   *event_service.EventService

	AuthService *auth_service.AuthService

	BookingService *booking_service.BookingService
}

func Init(repos *repositories.Repositories, cfg *config.Config) *Services {
	emailSender := &email.EmailSender{

		SMTPHost: cfg.Host,
		SMTPPort: cfg.Port,

		SMTPUser: cfg.Username,
		SMTPPass: cfg.Password,

		FromEmail: cfg.FromEmail,
		FromName:  cfg.FromName,

		FrontendURL: cfg.AppHost,
	}

	return &Services{
		UserService:    user_service.Init(repos.UserRepository, cfg),
		HubService:     hub_service.Init(repos.HubRepository, cfg),
		AdminService:   admin_service.Init(repos.AdminRepository, cfg),
		MentorService:  mentor_service.Init(repos.MentorRepository, cfg),
		RequestService: requests_service.Init(repos.RequestRepository, cfg),
		EventService:   event_service.Init(repos.EventRepository, cfg),
		AuthService:    auth_service.Init(repos.UserRepository, repos.AuthTokenRepository, emailSender, []byte(cfg.SecretKey)),
		BookingService: booking_service.Init(repos.BookingRepository, cfg),
	}
}
