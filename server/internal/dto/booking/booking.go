package booking_dto

type CreateBookingDto struct {
	BookingDate  string `json:"booking_date"`
	BookingZone  string `json:"booking_zone"`
	BookingSlots int    `json:"booking_slots"`
	UserId       uint64 `json:"user_id"`
}
