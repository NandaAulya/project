package controller

import (
	"project/model"
	"project/entity"
)

func CariUsernamePegawai(employees []*entity.Employee, username string) *entity.Employee{
	return model.FindEmployeeByUsername(employees, username)
}

func CariPelanngganUsername(customers []*entity.Customer, username string) *entity.Customer{
	return model.FindCustomerByUsername(customers, username)
}

func CariPesanan(booking []*entity.Booking, roomtype string) *entity.Booking{
	return model.FindBookingByRoomType(booking, roomtype)
}

func CariKamar(room []*entity.Room, roomtype string) *entity.Room{
	return model.FindRoomByRoomType(room, roomtype)
}

func TampilanPesanan(customer *entity.Customer) {
	model.ShowBooking(customer)
}

func TampilkanSemuaPesanan(customers []*entity.Customer) {
	model.ShowAllBookings(customers)
}