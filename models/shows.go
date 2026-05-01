package models

type Event struct {
	Name      string
	Date      string
	Genre     string
	Seats     int64
	Organizer string
}

func GetEvents() []Event {
	MultiEvent := []Event{
		{Name: "Concert A", Date: "2024-07-01", Genre: "Music", Seats: 100, Organizer: "Organizer A"},
		{Name: "play", Date: "2024-07-05", Genre: "Theater", Seats: 50, Organizer: "Organizer B"},
		{Name: "Comedy Show C", Date: "2024-07-10", Genre: "Comedy", Seats: 80, Organizer: "Organizer C"},
		{Name: "Dance Performance D", Date: "2024-07-15", Genre: "Dance", Seats: 60, Organizer: "Organizer D"},
	}
	return MultiEvent
}
