package entity

type Customer struct {
	Name     string
	Email    string
	Username string
	Password string
	Bookings *BookingList//linked list untuk menimpan pesanan
}

type Booking struct {
	RoomType string
	Duration int
	Next *Booking
}

type BookingList struct{
	Head *Booking //pointer ke node selanjutnya
}

type Room struct {
	RoomType  string
	Available int
}

type Employee struct {
	Username string
	Password string
}