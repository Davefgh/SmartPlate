// internal/repository/vehicle_repository.go
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"smartplate-api/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VehicleRepository interface {
	CreateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error)
	GetVehicle(ctx context.Context, id int) (*models.Vehicle, error)
	UpdateVehicle(ctx context.Context, id int, vehicle *models.Vehicle) error
	DeleteVehicle(ctx context.Context, id int) error
}

type vehicleRepo struct {
	db *sqlx.DB
}

func NewVehicleRepository(db *sqlx.DB) VehicleRepository {
	return &vehicleRepo{db: db}
}

func (r *vehicleRepo) CreateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `INSERT INTO vehicles (
        vehicle_category, mv_file_number, vehicle_make, vehicle_series, vehicle_type,
        body_type, year_model, engine_model, engine_number, chassis_number,
        piston_displacement, number_of_cylinders, fuel_type, color, gvw,
        net_weight, shipping_weight, usage_classification, first_registration_date,
        late_renewal_date, registration_expiry_date, lto_office_code, classification,
        denomination, or_number, cr_number, lto_client_id
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
        $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
        $21, $22, $23, $24, $25, $26, $27
    ) RETURNING vehicle_id`


    err = tx.QueryRowContext(ctx, query,
        // Ensure all 27 parameters are included
        vehicle.VEHICLE_CATEGORY, vehicle.MV_FILE_NUMBER, vehicle.VEHICLE_MAKE, vehicle.VEHICLE_SERIES,
        vehicle.VEHICLE_TYPE, vehicle.BODY_TYPE, vehicle.YEAR_MODEL, vehicle.ENGINE_MODEL,
        vehicle.ENGINE_NUMBER, vehicle.CHASSIS_NUMBER, vehicle.PISTON_DISPLACEMENT,
        vehicle.NUMBER_OF_CYLINDERS, vehicle.FUEL_TYPE, vehicle.COLOR, vehicle.GVW,
        vehicle.NET_WEIGHT, vehicle.SHIPPING_WEIGHT, vehicle.USAGE_CLASSIFICATION,
        vehicle.FIRST_REGISTRATION_DATE, vehicle.LATE_RENEWAL_DATE,
        vehicle.REGISTRATION_EXPIRY_DATE, vehicle.LTO_OFFICE_CODE,
        vehicle.CLASSIFICATION, vehicle.DENOMINATION, vehicle.OR_NUMBER,
        vehicle.CR_NUMBER, vehicle.LTO_CLIENT_ID,  // Now 27 parameters
    ).Scan(&vehicle.VEHICLE_ID)
	
	if err != nil {
		return nil, err
	}

	// Insert plate
	plateQuery := `INSERT INTO plates (
		vehicle_id, plate_number, plate_type, plate_issue_date, 
		plate_expiration_date, status
	) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.ExecContext(ctx, plateQuery,
		vehicle.VEHICLE_ID,
		vehicle.PLATE.PLATE_NUMBER,
		vehicle.PLATE.PLATE_TYPE,
		vehicle.PLATE.PLATE_ISSUE_DATE,
		vehicle.PLATE.PLATE_EXPIRATION_DATE,
		vehicle.PLATE.STATUS,
	)
	if err != nil {
		return nil, err
	}
	formID := generateFormID() // Implement your ID generation logic
	inspectionID := generateInspectionID()
	// Insert registration form
	formQuery := `INSERT INTO registration_forms (
		registration_form_id, vehicle_id, submitted_date, status, lto_client_id
	) VALUES ($1, $2, $3, $4, $5)`
	
	_, err = tx.ExecContext(ctx, formQuery,
		formID,
		vehicle.VEHICLE_ID,
		vehicle.REGISTRATION_FORM.SUBMITTED_DATE,
		vehicle.REGISTRATION_FORM.STATUS,
		vehicle.REGISTRATION_FORM.LTO_CLIENT_ID,
	)
	if err != nil {
		return nil, err
	}

	
	// For inspection
	inspectionQuery := `INSERT INTO inspection (
		inspection_id, vehicle_id, lto_client_id, inspection_date, 
		inspection_status, notes
	) VALUES ($1, $2, $3, $4, $5, $6)`
	
	_, err = tx.ExecContext(ctx, inspectionQuery,
		inspectionID,
		vehicle.VEHICLE_ID,
		vehicle.INSPECTION.LTO_CLIENT_ID,
		vehicle.INSPECTION.INSPECTION_DATE,
		vehicle.INSPECTION.INSPECTION_STATUS,
		vehicle.INSPECTION.NOTES,
	)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (r *vehicleRepo) GetVehicle(ctx context.Context, id int) (*models.Vehicle, error) {
	vehicle := &models.Vehicle{}
	query := `SELECT * FROM vehicles WHERE vehicle_id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&vehicle.VEHICLE_ID, &vehicle.VEHICLE_CATEGORY, &vehicle.MV_FILE_NUMBER,
		&vehicle.VEHICLE_MAKE, &vehicle.VEHICLE_SERIES, &vehicle.VEHICLE_TYPE,
		&vehicle.BODY_TYPE, &vehicle.YEAR_MODEL, &vehicle.ENGINE_MODEL,
		&vehicle.ENGINE_NUMBER, &vehicle.CHASSIS_NUMBER, &vehicle.PISTON_DISPLACEMENT,
		&vehicle.NUMBER_OF_CYLINDERS, &vehicle.FUEL_TYPE, &vehicle.COLOR,
		&vehicle.GVW, &vehicle.NET_WEIGHT, &vehicle.SHIPPING_WEIGHT,
		&vehicle.USAGE_CLASSIFICATION, &vehicle.FIRST_REGISTRATION_DATE,
		&vehicle.LATE_RENEWAL_DATE, &vehicle.REGISTRATION_EXPIRY_DATE,
		&vehicle.LTO_OFFICE_CODE, &vehicle.CLASSIFICATION, &vehicle.DENOMINATION,
		&vehicle.OR_NUMBER, &vehicle.CR_NUMBER, &vehicle.LTO_CLIENT_ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("vehicle not found")
		}
		return nil, err
	}

	// Get plate
	plate := models.Plate{}
	plateQuery := `SELECT * FROM plates WHERE vehicle_id = $1`
	err = r.db.QueryRowContext(ctx, plateQuery, id).Scan(
		&plate.VEHICLE_ID, &plate.PLATE_NUMBER, &plate.PLATE_TYPE,
		&plate.PLATE_ISSUE_DATE, &plate.PLATE_EXPIRATION_DATE, &plate.STATUS,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	vehicle.PLATE = plate

	// Get registration form
	form := models.RegistrationForm{}
	formQuery := `SELECT * FROM registration_forms WHERE vehicle_id = $1`
	err = r.db.QueryRowContext(ctx, formQuery, id).Scan(
		&form.REGISTRATION_FORM_ID, &form.VEHICLE_ID, &form.SUBMITTED_DATE,
		&form.STATUS, &form.LTO_CLIENT_ID,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	vehicle.REGISTRATION_FORM = form

	// Get inspection
	inspection := models.Inspection{}
	inspectionQuery := `SELECT * FROM inspections WHERE vehicle_id = $1`
	err = r.db.QueryRowContext(ctx, inspectionQuery, id).Scan(
		&inspection.INSPECTION_ID, &inspection.VEHICLE_ID, &inspection.LTO_CLIENT_ID,
		&inspection.INSPECTION_DATE, &inspection.INSPECTION_STATUS, &inspection.NOTES,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	vehicle.INSPECTION = inspection

	return vehicle, nil
}

func (r *vehicleRepo) UpdateVehicle(ctx context.Context, id int, vehicle *models.Vehicle) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update vehicle
	query := `UPDATE vehicles SET 
		vehicle_category = $1, mv_file_number = $2, vehicle_make = $3, 
		vehicle_series = $4, vehicle_type = $5, body_type = $6, 
		year_model = $7, engine_model = $8, engine_number = $9, 
		chassis_number = $10, piston_displacement = $11, 
		number_of_cylinders = $12, fuel_type = $13, color = $14, 
		gvw = $15, net_weight = $16, shipping_weight = $17, 
		usage_classification = $18, first_registration_date = $19, 
		late_renewal_date = $20, registration_expiry_date = $21, 
		lto_office_code = $22, classification = $23, denomination = $24, 
		or_number = $25, cr_number = $26, lto_client_id = $27 
	WHERE vehicle_id = $28`
	_, err = tx.ExecContext(ctx, query,
		vehicle.VEHICLE_CATEGORY, vehicle.MV_FILE_NUMBER, vehicle.VEHICLE_MAKE,
		vehicle.VEHICLE_SERIES, vehicle.VEHICLE_TYPE, vehicle.BODY_TYPE,
		vehicle.YEAR_MODEL, vehicle.ENGINE_MODEL, vehicle.ENGINE_NUMBER,
		vehicle.CHASSIS_NUMBER, vehicle.PISTON_DISPLACEMENT,
		vehicle.NUMBER_OF_CYLINDERS, vehicle.FUEL_TYPE, vehicle.COLOR,
		vehicle.GVW, vehicle.NET_WEIGHT, vehicle.SHIPPING_WEIGHT,
		vehicle.USAGE_CLASSIFICATION, vehicle.FIRST_REGISTRATION_DATE,
		vehicle.LATE_RENEWAL_DATE, vehicle.REGISTRATION_EXPIRY_DATE,
		vehicle.LTO_OFFICE_CODE, vehicle.CLASSIFICATION, vehicle.DENOMINATION,
		vehicle.OR_NUMBER, vehicle.CR_NUMBER, vehicle.LTO_CLIENT_ID,
		id,
	)
	if err != nil {
		return err
	}

	// Update plate
	plateQuery := `UPDATE plates SET 
		plate_number = $1, plate_type = $2, plate_issue_date = $3, 
		plate_expiration_date = $4, status = $5 
	WHERE vehicle_id = $6`
	_, err = tx.ExecContext(ctx, plateQuery,
		vehicle.PLATE.PLATE_NUMBER, vehicle.PLATE.PLATE_TYPE,
		vehicle.PLATE.PLATE_ISSUE_DATE, vehicle.PLATE.PLATE_EXPIRATION_DATE,
		vehicle.PLATE.STATUS, id,
	)
	if err != nil {
		return err
	}

	// Update registration form
	formQuery := `UPDATE registration_forms SET 
		submitted_date = $1, status = $2, lto_client_id = $3 
	WHERE vehicle_id = $4`
	_, err = tx.ExecContext(ctx, formQuery,
		vehicle.REGISTRATION_FORM.SUBMITTED_DATE, vehicle.REGISTRATION_FORM.STATUS,
		vehicle.REGISTRATION_FORM.LTO_CLIENT_ID, id,
	)
	if err != nil {
		return err
	}

	// Update inspection
	inspectionQuery := `UPDATE inspections SET 
		lto_client_id = $1, inspection_date = $2, 
		inspection_status = $3, notes = $4 
	WHERE vehicle_id = $5`
	_, err = tx.ExecContext(ctx, inspectionQuery,
		vehicle.INSPECTION.LTO_CLIENT_ID, vehicle.INSPECTION.INSPECTION_DATE,
		vehicle.INSPECTION.INSPECTION_STATUS, vehicle.INSPECTION.NOTES, id,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *vehicleRepo) DeleteVehicle(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Delete related records first
	queries := []string{
		"DELETE FROM plates WHERE vehicle_id = $1",
		"DELETE FROM registration_forms WHERE vehicle_id = $1",
		"DELETE FROM inspections WHERE vehicle_id = $1",
		"DELETE FROM vehicles WHERE vehicle_id = $1",
	}

	for _, query := range queries {
		_, err = tx.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func generateFormID() string {
    return "RF-" + time.Now().Format("20060102") + "-" + uuid.New().String()[:8]
}

func generateInspectionID() string {
    return "INSP-" + time.Now().Format("20060102") + "-" + uuid.New().String()[:8]
}