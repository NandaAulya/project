package model

import (
	"fmt"
	"project/entity"
)

// tugas model untuk mengambil data dari database
func FindCustomerByUsername(customers []*entity.Customer, username string) *entity.Customer {
	for _, customer := range customers {
		if customer.Username == username {
			return customer
		}
	}
	return nil
}

// employee
func FindEmployeeByUsername(employees []*entity.Employee, username string) *entity.Employee {
	for _, employee := range employees {
		if employee.Username == username {
			return employee
		}
	}
	return nil
}

// booking
func FindBookingByRoomType(booking []*entity.Booking, roomtype string) *entity.Booking {
	for _, booking := range booking {
		if booking.RoomType == roomtype {
			return booking
		}
	}
	return nil
}

// room
func FindRoomByRoomType(room []*entity.Room, roomtype string) *entity.Room {
	for _, room := range room {
		if room.RoomType == roomtype {
			return room
		}
	}
	return nil
}

// pemesanan
func ShowBooking(customer *entity.Customer) {
 }

// daftar semua pemesanan
func ShowAllBookings(customers []*entity.Customer) {
	if len(customers) > 0 {
		for _, customer := range customers {
			fmt.Printf("Pelanggan: %s\n", customer.Name)
			if len(customer.Bookings.Head.Next.RoomType) > 0 {
				fmt.Println("Daftar Pemesanan:")
				for i, booking := range customer.Bookings.Head.Next.RoomType {
					fmt.Printf("%d. Kamar %s, Durasi %d hari\n", i+1, booking)
				}
			} else {
				fmt.Println("Belum ada pemesanan.")
			}
			fmt.Println("------------------------------------------------")
		}
	} else {
		fmt.Println("Belum ada pelanggan atau pemesanan.")
	}
}