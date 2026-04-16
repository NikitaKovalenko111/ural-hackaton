package bookings_storage

import (
	"database/sql"

	booking_dto "ural-hackaton/internal/dto/booking"
	"ural-hackaton/internal/models"
	"ural-hackaton/internal/storage"
)

type BookingRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) CreateBooking(booking *booking_dto.CreateBookingDto) (*models.Booking, error) {
	var createdBooking models.Booking

	err := r.db.Db.QueryRow(
		`INSERT INTO bookings (booking_date, booking_zone, booking_slots, user_id)
		 VALUES ($1::timestamptz, $2, $3, $4)
		 RETURNING booking_id, booking_date::text, booking_zone, booking_slots, user_id`,
		booking.BookingDate,
		booking.BookingZone,
		booking.BookingSlots,
		booking.UserId,
	).Scan(&createdBooking.Id, &createdBooking.BookingDate, &createdBooking.BookingZone, &createdBooking.BookingSlots, &createdBooking.UserId)

	if err != nil {
		return nil, err
	}

	return &createdBooking, nil
}

func (r *BookingRepo) GetBookingById(id uint64) (*models.Booking, error) {
	var booking models.Booking

	err := r.db.Db.QueryRow(
		`SELECT booking_id, booking_date::text, booking_zone, booking_slots, user_id FROM bookings WHERE booking_id = $1`,
		id,
	).Scan(&booking.Id, &booking.BookingDate, &booking.BookingZone, &booking.BookingSlots, &booking.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &booking, nil
}

func (r *BookingRepo) GetBookingsByUserId(userId uint64) ([]models.Booking, error) {
	rows, err := r.db.Db.Query(
		`SELECT booking_id, booking_date::text, booking_zone, booking_slots, user_id FROM bookings WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := make([]models.Booking, 0)
	for rows.Next() {
		var booking models.Booking
		if err = rows.Scan(&booking.Id, &booking.BookingDate, &booking.BookingZone, &booking.BookingSlots, &booking.UserId); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepo) GetAllBookings() ([]models.Booking, error) {
	rows, err := r.db.Db.Query(
		`SELECT booking_id, booking_date::text, booking_zone, booking_slots, user_id FROM bookings`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := make([]models.Booking, 0)
	for rows.Next() {
		var booking models.Booking
		if err = rows.Scan(&booking.Id, &booking.BookingDate, &booking.BookingZone, &booking.BookingSlots, &booking.UserId); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepo) UpdateBooking(booking *models.Booking) (*models.Booking, error) {
	var updatedBooking models.Booking

	err := r.db.Db.QueryRow(
		`UPDATE bookings
		 SET booking_date = $1::timestamptz, booking_zone = $2, booking_slots = $3, user_id = $4
		 WHERE booking_id = $5
		 RETURNING booking_id, booking_date::text, booking_zone, booking_slots, user_id`,
		booking.BookingDate,
		booking.BookingZone,
		booking.BookingSlots,
		booking.UserId,
		booking.Id,
	).Scan(&updatedBooking.Id, &updatedBooking.BookingDate, &updatedBooking.BookingZone, &updatedBooking.BookingSlots, &updatedBooking.UserId)

	if err != nil {
		return nil, err
	}

	return &updatedBooking, nil
}

func (r *BookingRepo) DeleteBooking(id uint64) error {
	_, err := r.db.Db.Exec(`DELETE FROM bookings WHERE booking_id = $1`, id)
	return err
}
