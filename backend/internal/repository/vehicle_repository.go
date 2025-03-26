package repository

import (
	"vehicle-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type VehicleRepository struct {
	db *sqlx.DB
}

func NewVehicleRepository(db *sqlx.DB) *VehicleRepository {
	return &VehicleRepository{db: db}
}

// Create a new vehicle
func (r *VehicleRepository) Create(vehicle *models.Vehicle) error {
	query := `INSERT INTO vehicles (make, model, plate) VALUES ($1, $2, $3) 
	          RETURNING id, created, updated`
	return r.db.QueryRow(query, vehicle.Make, vehicle.Model, vehicle.Plate).Scan(
		&vehicle.ID, &vehicle.Created, &vehicle.Updated)
}

// Get all vehicles
func (r *VehicleRepository) GetAll() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := r.db.Select(&vehicles, "SELECT * FROM vehicles ORDER BY id")
	return vehicles, err
}

// Get vehicle by ID
func (r *VehicleRepository) GetByID(id int) (models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.Get(&vehicle, "SELECT * FROM vehicles WHERE id = $1", id)
	return vehicle, err
}

// Update vehicle
func (r *VehicleRepository) Update(vehicle *models.Vehicle) error {
	query := `UPDATE vehicles SET make = $1, model = $2, plate = $3 
	          WHERE id = $4`
	_, err := r.db.Exec(query, vehicle.Make, vehicle.Model, vehicle.Plate, vehicle.ID)
	return err
}

// Delete vehicle
func (r *VehicleRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM vehicles WHERE id = $1", id)
	return err
}
