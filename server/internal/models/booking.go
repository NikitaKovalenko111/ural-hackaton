package models

type Booking struct {
	Id           uint64 `json:"booking_id"`
	BookingDate  string `json:"booking_date"`
	BookingZone  string `json:"booking_zone"`
	BookingSlots int    `json:"booking_slots"`
	UserId       uint64 `json:"user_id"`
}
