package services

import (
	"ural-hackaton/internal/config"
	admin_service "ural-hackaton/internal/services/handlers/admin"
	hub_service "ural-hackaton/internal/services/handlers/hub"
	mentor_service "ural-hackaton/internal/services/handlers/mentor"
	user_service "ural-hackaton/internal/services/handlers/user"
	"ural-hackaton/internal/storage/repositories"
)

type Services struct {
	UserService   *user_service.UserService
	HubService    *hub_service.HubService
	AdminService  *admin_service.AdminService
	MentorService *mentor_service.MentorService
}

func Init(repos *repositories.Repositories, cfg *config.Config) *Services {
	return &Services{
		UserService:   user_service.Init(repos.UserRepository, cfg),
		HubService:    hub_service.Init(repos.HubRepository, cfg),
		AdminService:  admin_service.Init(repos.AdminRepository, cfg),
		MentorService: mentor_service.Init(repos.MentorRepository, cfg),
	}
}
