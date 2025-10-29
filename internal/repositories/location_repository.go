package repositories

import (
	"database/sql"
	"github.com/rquiogue/travel-to-do-list/internal/models"
)

type LocationRepository struct {
	DB *sql.DB
}

func NewLocationRepository(db *sql.DB) *LocationRepository {
	return &LocationRepository{DB: db}
}

func (r *LocationRepository) GetAll() ([]models.Location, error) {
	rows, err := r.DB.Query("SELECT id, title, completed FROM locations ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Locations []models.Location
	for rows.Next() {
		var t models.Location
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			return nil, err
		}
		Locations = append(Locations, t)
	}
	return Locations, nil
}

func (r *LocationRepository) Create(Location *models.Location) error {
	return r.DB.QueryRow(
		"INSERT INTO locations (title, completed) VALUES ($1, $2) RETURNING id",
		Location.Title, Location.Completed,
	).Scan(&Location.ID)
}

func (r *LocationRepository) Update(Location *models.Location) error {
	_, err := r.DB.Exec(
		"UPDATE locations SET title=$1, completed=$2 WHERE id=$3",
		Location.Title, Location.Completed, Location.ID,
	)
	return err
}

func (r *LocationRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM locations WHERE id=$1", id)
	return err
}