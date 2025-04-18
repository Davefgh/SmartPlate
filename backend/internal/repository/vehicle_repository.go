package repository

import (
	"context"
	"errors"
	"smartplate-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type VehicleRepository interface {
    CreateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error)
    GetAllVehicles(ctx context.Context) ([]*models.Vehicle, error)
    GetVehicle(ctx context.Context, id int) (*models.Vehicle, error)
    UpdateVehicle(ctx context.Context, id int, vehicle *models.Vehicle) error
    DeleteVehicle(ctx context.Context, id int) error
    UpdatePlate(ctx context.Context, plate *models.Plate) error
}


type vehicleRepo struct {
	db *sqlx.DB
}

// NewVehicleRepository creates a new VehicleRepository
func NewVehicleRepository(db *sqlx.DB) VehicleRepository {
	return &vehicleRepo{db: db}
}

func (r *vehicleRepo) GetAllVehicles(ctx context.Context) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	query := `SELECT * FROM vehicles`
	if err := r.db.SelectContext(ctx, &vehicles, query); err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (r *vehicleRepo) CreateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {
    tx, err := r.db.BeginTxx(ctx, nil)
    if err != nil {
        return nil, err
    }
    defer tx.Rollback()

    queryV := `INSERT INTO vehicles (
        vehicle_category, mv_file_number, vehicle_make, vehicle_series, vehicle_type,
        body_type, year_model, engine_model, engine_number, chassis_number,
        piston_displacement, number_of_cylinders, fuel_type, color,
        gvw, net_weight, shipping_weight, usage_classification,
        first_registration_date, late_renewal_date, registration_expiry_date,
        lto_office_code, classification, denomination, or_number, cr_number, lto_client_id
    ) VALUES (
        :n_vehicle_category, :n_mv_file_number, :n_vehicle_make, :n_vehicle_series, :n_vehicle_type,
        :n_body_type, :n_year_model, :n_engine_model, :n_engine_number, :n_chassis_number,
        :n_piston_displacement, :n_number_of_cylinders, :n_fuel_type, :n_color,
        :n_gvw, :n_net_weight, :n_shipping_weight, :n_usage_classification,
        :n_first_registration_date, :n_late_renewal_date, :n_registration_expiry_date,
        :n_lto_office_code, :n_classification, :n_denomination, :n_or_number, :n_cr_number, :n_lto_client_id
    ) RETURNING vehicle_id`

    namedParams := map[string]interface{}{
        "n_vehicle_category":         vehicle.VEHICLE_CATEGORY,
        "n_mv_file_number":           vehicle.MV_FILE_NUMBER,
        "n_vehicle_make":             vehicle.VEHICLE_MAKE,
        "n_vehicle_series":           vehicle.VEHICLE_SERIES,
        "n_vehicle_type":             vehicle.VEHICLE_TYPE,
        "n_body_type":                vehicle.BODY_TYPE,
        "n_year_model":               vehicle.YEAR_MODEL,
        "n_engine_model":             vehicle.ENGINE_MODEL,
        "n_engine_number":            vehicle.ENGINE_NUMBER,
        "n_chassis_number":           vehicle.CHASSIS_NUMBER,
        "n_piston_displacement":      vehicle.PISTON_DISPLACEMENT,
        "n_number_of_cylinders":      vehicle.NUMBER_OF_CYLINDERS,
        "n_fuel_type":                vehicle.FUEL_TYPE,
        "n_color":                    vehicle.COLOR,
        "n_gvw":                      vehicle.GVW,
        "n_net_weight":               vehicle.NET_WEIGHT,
        "n_shipping_weight":          vehicle.SHIPPING_WEIGHT,
        "n_usage_classification":     vehicle.USAGE_CLASSIFICATION,
        "n_first_registration_date":  vehicle.FIRST_REGISTRATION_DATE,
        "n_late_renewal_date":        vehicle.LATE_RENEWAL_DATE,
        "n_registration_expiry_date": vehicle.REGISTRATION_EXPIRY_DATE,
        "n_lto_office_code":          vehicle.LTO_OFFICE_CODE,
        "n_classification":           vehicle.CLASSIFICATION,
        "n_denomination":             vehicle.DENOMINATION,
        "n_or_number":                vehicle.OR_NUMBER,
        "n_cr_number":                vehicle.CR_NUMBER,
        "n_lto_client_id":            vehicle.LTO_CLIENT_ID,
    }

    // Use NamedQueryContext so sqlx can expand the map
       // 3) Expand into positional SQL + args
	   sqlStr, args, err := sqlx.Named(queryV, namedParams)
	   if err != nil {
		   return nil, err
	   }
	   // 4) Rebind to the driver’s bindvar style (Postgres wants $1,$2…)
	   sqlStr = r.db.Rebind(sqlStr)
   
	   // 5) Execute via the Tx
	   var newID int
	   row := tx.QueryRowxContext(ctx, sqlStr, args...)
	   if err := row.Scan(&newID); err != nil {
		   return nil, err
	   }
   
	   vehicle.VEHICLE_ID = &newID
   
	   if err := tx.Commit(); err != nil {
		   return nil, err
	   }
	   return vehicle, nil
}

func (r *vehicleRepo) GetVehicle(ctx context.Context, id int) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	query := `SELECT * FROM vehicles WHERE vehicle_id = $1`
	if err := r.db.GetContext(ctx, &vehicle, query, id); err != nil {
		return nil, errors.New("vehicle not found")
	}
	return &vehicle, nil
}

func (r *vehicleRepo) UpdateVehicle(ctx context.Context, id int, vehicle *models.Vehicle) error {
	query := `UPDATE vehicles SET
		vehicle_category = :vehicle_category,
		mv_file_number = :mv_file_number,
		vehicle_make = :vehicle_make,
		vehicle_series = :vehicle_series,
		vehicle_type = :vehicle_type,
		body_type = :body_type,
		year_model = :year_model,
		engine_model = :engine_model,
		engine_number = :engine_number,
		chassis_number = :chassis_number,
		piston_displacement = :piston_displacement,
		number_of_cylinders = :number_of_cylinders,
		fuel_type = :fuel_type,
		color = :color,
		gvw = :gvw,
		net_weight = :net_weight,
		shipping_weight = :shipping_weight,
		usage_classification = :usage_classification,
		first_registration_date = :first_registration_date,
		late_renewal_date = :late_renewal_date,
		registration_expiry_date = :registration_expiry_date,
		lto_office_code = :lto_office_code,
		classification = :classification,
		denomination = :denomination,
		or_number = :or_number,
		cr_number = :cr_number,
		lto_client_id = :lto_client_id
	WHERE vehicle_id = :vehicle_id`
	params := map[string]interface{}{
		"vehicle_id":                  vehicle.VEHICLE_ID,
		"vehicle_category":            vehicle.VEHICLE_CATEGORY,
		"mv_file_number":              vehicle.MV_FILE_NUMBER,
		"vehicle_make":                vehicle.VEHICLE_MAKE,
		"vehicle_series":              vehicle.VEHICLE_SERIES,
		"vehicle_type":                vehicle.VEHICLE_TYPE,
		"body_type":                   vehicle.BODY_TYPE,
		"year_model":                  vehicle.YEAR_MODEL,
		"engine_model":                vehicle.ENGINE_MODEL,
		"engine_number":               vehicle.ENGINE_NUMBER,
		"chassis_number":              vehicle.CHASSIS_NUMBER,
		"piston_displacement":         vehicle.PISTON_DISPLACEMENT,
		"number_of_cylinders":         vehicle.NUMBER_OF_CYLINDERS,
		"fuel_type":                   vehicle.FUEL_TYPE,
		"color":                       vehicle.COLOR,
		"gvw":                         vehicle.GVW,
		"net_weight":                  vehicle.NET_WEIGHT,
		"shipping_weight":             vehicle.SHIPPING_WEIGHT,
		"usage_classification":        vehicle.USAGE_CLASSIFICATION,
		"first_registration_date":     vehicle.FIRST_REGISTRATION_DATE,
		"late_renewal_date":           vehicle.LATE_RENEWAL_DATE,
		"registration_expiry_date":    vehicle.REGISTRATION_EXPIRY_DATE,
		"lto_office_code":             vehicle.LTO_OFFICE_CODE,
		"classification":              vehicle.CLASSIFICATION,
		"denomination":                vehicle.DENOMINATION,
		"or_number":                   vehicle.OR_NUMBER,
		"cr_number":                   vehicle.CR_NUMBER,
		"lto_client_id":               vehicle.LTO_CLIENT_ID,
	}
	// Using NamedExec for update
	result, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("no vehicle updated")
	}
	return nil
}

func (r *vehicleRepo) DeleteVehicle(ctx context.Context, id int) error {
	query := `DELETE FROM vehicles WHERE vehicle_id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("no vehicle deleted")
	}
	return nil
}

// Implementation:
func (r *vehicleRepo) UpdatePlate(ctx context.Context, plate *models.Plate) error {
    query := `
      UPDATE plate
         SET plate_number          = $1,
             plate_type            = $2,
             plate_issue_date      = $3,
             plate_expiration_date = $4,
             status                = $5
       WHERE vehicle_id = $6
    `
    _, err := r.db.ExecContext(
        ctx,
        query,
        plate.PLATE_NUMBER,
        plate.PLATE_TYPE,
        plate.PLATE_ISSUE_DATE,
        plate.PLATE_EXPIRATION_DATE,
        plate.STATUS,
        plate.VEHICLE_ID,
    )
    return err
}
