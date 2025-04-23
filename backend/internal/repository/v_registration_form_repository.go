package repository

import (
    "context"

    "github.com/jmoiron/sqlx"
    "smartplate-api/internal/models"
)

type RegistrationFormRepository interface {
    Create(ctx context.Context, p *models.CreateRegistrationFormParams) (*models.RegistrationForm, error)
    GetAll(ctx context.Context) ([]models.RegistrationForm, error)
    GetByID(ctx context.Context, id string) (*models.RegistrationForm, error)
    Update(ctx context.Context, f *models.RegistrationForm) error
    Delete(ctx context.Context, id string) error
}

type registrationFormRepo struct {
    db *sqlx.DB
}

func NewRegistrationFormRepository(db *sqlx.DB) RegistrationFormRepository {
    return &registrationFormRepo{db: db}
}

func (r *registrationFormRepo) Create(
    ctx context.Context,
    p *models.CreateRegistrationFormParams,
) (*models.RegistrationForm, error) {
    var full models.RegistrationForm
    err := r.db.
        QueryRowxContext(ctx, `
      INSERT INTO registration_form
        (lto_client_id, vehicle_id, status, registration_type)
      VALUES
        ($1, $2, $3, $4)
      RETURNING registration_form_id,
                lto_client_id,
                vehicle_id,
                submitted_date,
                status,
                registration_type
    `, p.LTOClientID, p.VehicleID, p.Status, p.RegistrationType).
        StructScan(&full)
    if err != nil {
        return nil, err
    }
    return &full, nil
}

func (r *registrationFormRepo) GetAll(ctx context.Context) ([]models.RegistrationForm, error) {
    var out []models.RegistrationForm
    err := r.db.SelectContext(ctx, &out, `
        SELECT registration_form_id,
               lto_client_id,
               vehicle_id,
               submitted_date,
               status,
               registration_type
          FROM registration_form
         ORDER BY submitted_date DESC
    `)
    return out, err
}

func (r *registrationFormRepo) GetByID(ctx context.Context, id string) (*models.RegistrationForm, error) {
    var f models.RegistrationForm
    err := r.db.GetContext(ctx, &f, `
        SELECT registration_form_id,
               lto_client_id,
               vehicle_id,
               submitted_date,
               status,
               registration_type
          FROM registration_form
         WHERE registration_form_id = $1
    `, id)
    if err != nil {
        return nil, err
    }
    return &f, nil
}

func (r *registrationFormRepo) Update(ctx context.Context, f *models.RegistrationForm) error {
    _, err := r.db.NamedExecContext(ctx, `
        UPDATE registration_form SET
          lto_client_id      = :lto_client_id,
          vehicle_id         = :vehicle_id,
          status             = :status,
          registration_type  = :registration_type
        WHERE registration_form_id = :registration_form_id
    `, f)
    return err
}

func (r *registrationFormRepo) Delete(ctx context.Context, id string) error {
    _, err := r.db.ExecContext(ctx, `
        DELETE FROM registration_form
         WHERE registration_form_id = $1
    `, id)
    return err
}
