package seeds

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type SeedsPostgres struct {
	db *sqlx.DB
}

func AddSeedsPostgres(db *sqlx.DB) *SeedsPostgres {
	return &SeedsPostgres{db: db}
}

func (s *SeedsPostgres) SeedAll() error {
	if err := s.SeedCompanies(); err != nil {
		return err
	}
	if err := s.SeedTours(); err != nil {
		return err
	}
	log.Println("All seeds applied successfully!")
	return nil
}

func (s *SeedsPostgres) SeedCompanies() error {
	query := `
	INSERT INTO company (id, name, phone_number, email) VALUES
	(102, 'Adventure Co.', '+12356789', 'test42@mail.com'),
	(103, 'Explorer Ltd.', '+98765432', 'contact@explorer.com'),
	(104, 'TravelWorld Inc.', '+11223344', 'info@travelworld.com'),
	(105, 'Globetrotters', '+99887766', 'hello@globetrotters.com')
	ON CONFLICT (id) DO NOTHING;
	`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Company seeds applied successfully!")
	return nil
}

func (s *SeedsPostgres) SeedTours() error {
	query := `
	INSERT INTO tour (id, company_id, title, start_time, end_time, group_size, languages, free_cancellation, cancellation_condition, description, meeting_place, 
	arrival_place, what_is_included, what_to_prepare, prohibitions, price, images)
	VALUES
	(3, 102, 'Mountain Adventure', '2024-11-01T08:00:00Z', '2024-11-01T18:00:00Z', 12, 
	'English, French', true, '{"policy": "Free cancellation up to 24 hours before departure."}', 
	'An exciting day exploring the mountains.', 'Central Station, Platform 5', 
	'Mountain Base Camp', 'Transport, Lunch, Guide', 'Hiking shoes, Water bottle, Sunscreen', 
	'No pets, No smoking', 150, 'image1.jpg,image2.jpg'),

	(4, 103, 'City Lights Tour', '2024-12-05T19:00:00Z', '2024-12-05T23:00:00Z', 25, 
	'English, Russian', true, '{"policy": "Free cancellation up to 12 hours before departure."}', 
	'Discover the charm of the city at night.', 'Downtown, Main Square', 
	'River Walk', 'Bus ride, Snacks', 'Comfortable shoes, Jacket', 
	'No alcohol, No pets', 75, 'city1.jpg,city2.jpg'),

	(5, 104, 'Beach Getaway', '2024-10-30T07:00:00Z', '2024-10-30T19:00:00Z', 15, 
	'English, Spanish', false, '{"policy": "Non-refundable."}', 
	'A full day of relaxation at the beach.', 'Beach Resort Lobby', 
	'Beach Resort', 'Transport, Lunch, Drinks', 'Swimsuit, Towel, Sunglasses', 
	'No smoking, No outside food', 120, 'beach1.jpg,beach2.jpg'),

	(6, 105, 'Cultural Heritage Tour', '2024-11-15T10:00:00Z', '2024-11-15T16:00:00Z', 10, 
	'English, German', true, '{"policy": "Free cancellation within 48 hours."}', 
	'Explore local culture and history.', 'Museum Entrance', 
	'Historic District', 'Tickets, Guide, Lunch', 'Comfortable clothes, Camera', 
	'No loud noises, Respect local customs', 90, 'heritage1.jpg,heritage2.jpg')

	ON CONFLICT (id) DO NOTHING;
	`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Tour seeds applied successfully!")
	return nil
}
