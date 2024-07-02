package main

func CalculateRoomPrice(room_type string, nights int, person_count int) (price int) {
	switch room_type {
	case "standard":
		price = nights * 200000 * person_count
	case "vip":
		price = nights * 300000 * person_count
	case "double":
		price = nights * 250000 * person_count
	default:
		price = nights * 200000 * person_count

	}
	return price
}
