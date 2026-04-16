package booking_service

import (
	"ural-hackaton/internal/config"
	booking_dto "ural-hackaton/internal/dto/booking"
	"ural-hackaton/internal/models"
	booking_storage "ural-hackaton/internal/storage/repositories/booking"
)

type BookingService struct {
	repo *booking_storage.BookingRepo
	cfg  *config.Config
}

func Init(bookingRepo *booking_storage.BookingRepo, cfg *config.Config) *BookingService {
	return &BookingService{
		repo: bookingRepo,
		cfg:  cfg,
	}
}

func (s *BookingService) CreateBooking(bookingDate string, bookingZone string, bookingSlots int, userId uint64) (*models.Booking, error) {
	bookingDto := &booking_dto.CreateBookingDto{
		BookingDate:  bookingDate,
		BookingZone:  bookingZone,
		BookingSlots: bookingSlots,
		UserId:       userId,
	}

	booking, err := s.repo.CreateBooking(bookingDto)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingService) GetBookingById(id uint64) (*models.Booking, error) {
	booking, err := s.repo.GetBookingById(id)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingService) GetBookingsByUserId(userId uint64) ([]models.Booking, error) {
	bookings, err := s.repo.GetBookingsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *BookingService) GetAllBookings() ([]models.Booking, error) {
	bookings, err := s.repo.GetAllBookings()
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *BookingService) UpdateBooking(bookingDate string, bookingZone string, bookingSlots int, userId uint64, id uint64) (*models.Booking, error) {
	bookingDto := &models.Booking{
		Id:           id,
		BookingDate:  bookingDate,
		BookingZone:  bookingZone,
		BookingSlots: bookingSlots,
		UserId:       userId,
	}

	booking, err := s.repo.UpdateBooking(bookingDto)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingService) DeleteBooking(id uint64) error {
	if err := s.repo.DeleteBooking(id); err != nil {
		return err
	}

	return nil
}
