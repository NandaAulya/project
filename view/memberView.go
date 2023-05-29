package view

import (
	"fmt"
	"project/entity"
	"project/model"
)

func Start(){
	rooms := []*entity.Room{
		{RoomType: "Standard", Available: 5},
		{RoomType: "Deluxe", Available: 3},
		{RoomType: "Suite", Available: 2},
	}

	employees := []*entity.Employee{
		{Username: "pegawai1", Password: "1"},
		{Username: "pegawai2", Password: "2"},
	}

	isLoggedIn := false
	var currentEmployee *entity.Employee

	for {
		fmt.Println("================================================")
		fmt.Println("|   Selamat Datang di Sistem Pemesanan Hotel   |")
		fmt.Println("================================================")
		fmt.Println("| 1. Login Pegawai                             |")
		fmt.Println("| 2. Daftar Customer                           |")
		fmt.Println("| 3. Login Customer                            |")
		fmt.Println("| 4. Memesan Kamar                             |")
		fmt.Println("| 5. Melihat Pemesanan                         |")
		fmt.Println("| 6. Daftar Semua Pemesanan (Pegawai)          |")
		fmt.Println("| 7. Logout (Pegawai)                          |")
		fmt.Println("| 8. Exit                                      |")
		fmt.Println("------------------------------------------------")
		fmt.Println()
		var choice int
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println()
			fmt.Println("---------------------")
			fmt.Println("|   Login Pegawai   |")
			fmt.Println("---------------------")
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			employee := model.FindEmployeeByUsername(employees, username) //controller
			if employee == nil || employee.Password != password {
				fmt.Println("Login gagal. Mohon periksa kembali username dan password Anda.")
			} else {
				fmt.Printf("Selamat datang, %s!\n", username)
				isLoggedIn = true
				currentEmployee = employee
			}
		case 2:
			fmt.Println()
			fmt.Println("==============")
			fmt.Println("|   Daftar   |")
			fmt.Println("==============")
			fmt.Print("Nama: ")
			var name string
			fmt.Scanln(&name)
			fmt.Print("Email: ")
			var email string
			fmt.Scanln(&email)
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			customer := &entity.Customer{
				Name:     name,
				Email:    email,
				Username: username,
				Password: password,
			}
			entity.Customer = append([]*Customer, customer)
		case 3:
			fmt.Println()
			fmt.Println("Pendaftaran berhasil!")
			fmt.Println("=============")
			fmt.Println("|   Login   |")
			fmt.Println("=============")
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			customer := model.FindCustomerByUsername(customers, username)
			if customer == nil || customer.Password != password {
				fmt.Println("Login gagal. Mohon periksa kembali username dan password Anda.")
			} else {
				fmt.Printf("Selamat datang, %s!\n", customer.Name)
				isLoggedIn = true
			}
		case 4:
			fmt.Println()
			fmt.Println("=====================")
			fmt.Println("|   Memesan Kamar   |")
			fmt.Println("=====================")
			if isLoggedIn {
				fmt.Print("Username: ")
				var username string
				fmt.Scanln(&username)

				customer := model.FindCustomerByUsername(customers, username)
				if customer != nil {
					fmt.Println("Pilih jenis kamar:")
					for i, room := range rooms {
						fmt.Printf("%d. %s (%d tersedia)\n", i+1, room.RoomType, room.Available)
					}

					var roomChoice int
					fmt.Print("Pilihan kamar: ")
					fmt.Scanln(&roomChoice)

					if roomChoice < 1 || roomChoice > len(rooms) {
						fmt.Println("Pilihan kamar tidak valid.")
					} else {
						roomIndex := roomChoice - 1
						room := rooms[roomIndex]
						if room.Available > 0 {
							fmt.Print("Durasi pemesanan (hari): ")
							var duration int
							fmt.Scanln(&duration)

							if duration <= 0 {
								fmt.Println("Durasi pemesanan tidak valid.")
							} else {
								booking := &entity.Booking{
									RoomType: room.RoomType,
									Duration: duration,
								}
								customer.Bookings = append(Customer.Bookings, booking)
								room.Available--
								fmt.Printf("Kamar %s telah dipesan selama %d hari.\n", room.RoomType, duration)
							}
						} else {
							fmt.Println("Kamar yang dipilih tidak tersedia.")
						}
					}
				} else {
					fmt.Println("Customer tidak ditemukan.")
				}
			} else {
				fmt.Println("Anda belum login.")
			}
		case 5:
			fmt.Println()
			fmt.Println("=========================")
			fmt.Println("|   Melihat Pemesanan   |")
			fmt.Println("=========================")
			if isLoggedIn {
				fmt.Print("Username: ")
				var username string
				fmt.Scanln(&username)

				customer := model.FindCustomerByUsername(customers, username)
				if customer != nil {
					model.ShowBooking(customer)
				} else {
					fmt.Println("Customer tidak ditemukan.")
				}
			} else {
				fmt.Println("Anda belum login.")
			}
		case 6:
			fmt.Println()
			fmt.Println("=========================")
			fmt.Println("|   Daftar Pemesanan   |")
			fmt.Println("=========================")
			if isLoggedIn && currentEmployee != nil {
				model.ShowAllBookings(customers) 
			} else {
				fmt.Println("Anda belum login sebagai pegawai.")
			}
		case 7:
			fmt.Println()
			fmt.Println("========================")
			fmt.Println("|   Logout (Pegawai)   |")
			fmt.Println("========================")
			isLoggedIn = false
			currentEmployee = nil
			fmt.Println("Anda telah logout sebagai pegawai.")
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		case 8:
			fmt.Println("Terima kasih telah menggunakan sistem pemesanan hotel. Sampai jumpa!")
			return
		}
	}
}