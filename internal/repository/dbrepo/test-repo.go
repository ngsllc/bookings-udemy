package dbrepo

import (
	"bookings-udemy/internal/models"
	"errors"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// returns true if availability exists for room with ID roomID, false otherwise
func (m *testDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// returns slice of all available rooms, if any, for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("room id invalid")
	}
	return room, nil
}
