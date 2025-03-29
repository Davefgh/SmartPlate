package repository

import (
	"fmt"
	"vehicle-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository  struct{
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

//create a new user
func (r *UserRepository) Create(user *models.User) error {
	tx := r.db.MustBegin()
    

    // Insert user
    err := tx.QueryRow(`
        INSERT INTO users (
            last_name, first_name, middle_name, email, password, role, status, lto_client_id
        ) VALUES (
            :last_name, :first_name, :middle_name, :email, :password, :role, :status, :lto_client_id
        )
        RETURNING user_id, created, updated
    `, user).Scan(&user.USER_ID, &user.CREATED, &user.UPDATED)
    
    if err != nil {
        tx.Rollback()
        return err
    }

    // Insert contact if any data exists
    _, err = tx.NamedExec(`
            INSERT INTO contacts (
                lto_client_id, telephone_number, int_area_code, mobile_number,
                emergency_contact_number, emergency_contact_name, emergency_contact_relationship, emergency_contact_address
            ) VALUES (
                :lto_client_id, :telephone_number, :int_area_code, :mobile_number,
                :emergency_contact_number, :emergency_contact_name, :emergency_contact_relationship, :emergency_contact_address
            )`,
			map[string]interface{}{
            "lto_client_id":                  user.LTO_CLIENT_ID,
            "telephone_number":               user.Contact.TELEPHONE_NUMBER,
            "int_area_code":                  user.Contact.INT_AREA_CODE,
            "mobile_number":                  user.Contact.MOBILE_NUMBER,
            "emergency_contact_number":       user.Contact.EMERGENCY_CONTACT_NUMBER,
            "emergency_contact_name":         user.Contact.EMERGENCY_CONTACT_NAME,
            "emergency_contact_relationship": user.Contact.EMERGENCY_CONTACT_RELATIONSHIP,
            "emergency_contact_address":      user.Contact.EMERGENCY_CONTACT_ADDRESS,
        	},
	)
	if err != nil {
        tx.Rollback()
        return err
    }
    

    return tx.Commit()
}


func (r *UserRepository) GetAll() ([]models.User, error) {
    const query = `
        SELECT 
            u.*,
            c.contact_id AS "contact.contact_id",
            c.lto_client_id AS "contact.lto_client_id",
            c.telephone_number AS "contact.telephone_number",
            c.int_area_code AS "contact.int_area_code",
            c.mobile_number AS "contact.mobile_number",
            c.emergency_contact_number AS "contact.emergency_contact_number",
            c.emergency_contact_name AS "contact.emergency_contact_name",
            c.emergency_contact_relationship AS "contact.emergency_contact_relationship",
            c.emergency_contact_address AS "contact.emergency_contact_address"
        FROM users u
        LEFT JOIN contacts c ON u.lto_client_id = c.lto_client_id
        ORDER BY u.user_id
    `
    var users []models.User
    err := r.db.Select(&users, query)
    return users, err
}

// GetByID with contact
// repository/user_repository.go
func (r *UserRepository) GetByID(user_id int) (models.User, error) {
    var user models.User
    query := `
        SELECT 
            u.*,
            c.contact_id AS "contact.contact_id",
            c.lto_client_id AS "contact.lto_client_id",
            c.telephone_number AS "contact.telephone_number",
            c.int_area_code AS "contact.int_area_code",
            c.mobile_number AS "contact.mobile_number",
            c.emergency_contact_number AS "contact.emergency_contact_number",
            c.emergency_contact_name AS "contact.emergency_contact_name",
            c.emergency_contact_relationship AS "contact.emergency_contact_relationship",
            c.emergency_contact_address AS "contact.emergency_contact_address"
        FROM users u
        LEFT JOIN contacts c ON u.lto_client_id = c.lto_client_id
        WHERE u.user_id = $1
    `
    err := r.db.Get(&user, query, user_id)
    return user, err
}
// Similarly update GetByEmail, GetByLTOClientID, etc.

// In your repository/user_repository.go
func (r *UserRepository) GetByLTOClientID(ltoClientID string) (models.User, error) {
    var user models.User
    err := r.db.Get(&user, "SELECT * FROM users WHERE lto_client_id = $1", ltoClientID)
    return user, err
}
//get user by email.l
func (r *UserRepository) GetByEmail(email string) (models.User, error){
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}

//delete by id use
func (r *UserRepository) Delete(user_id int) error {
    tx, err := r.db.Beginx()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }

    // Get LTO ID first
    var ltoID string
    if err := tx.Get(&ltoID, "SELECT lto_client_id FROM users WHERE user_id = $1", user_id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to get LTO ID: %w", err)
    }

    // Delete contact
    if _, err := tx.Exec("DELETE FROM contacts WHERE lto_client_id = $1", ltoID); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete contact: %w", err)
    }

    // Delete user
    if _, err := tx.Exec("DELETE FROM users WHERE user_id = $1", user_id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete user: %w", err)
    }

    return tx.Commit()
}
//update user
func (r *UserRepository) Update(user *models.User) error {
    tx, err := r.db.Beginx()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }

    // Update user
    userQuery := `
        UPDATE users SET
            last_name = :last_name,
            first_name = :first_name,
            middle_name = :middle_name,
            email = :email,
            password = :password,
            role = :role,
            status = :status,
            lto_client_id = :lto_client_id,
            updated = NOW()
        WHERE user_id = :user_id
    `
    if _, err = tx.NamedExec(userQuery, user); err != nil {
        tx.Rollback()
        return fmt.Errorf("user update failed: %w", err)
    }

    // Upsert contact
    contactQuery := `
        INSERT INTO contacts (
            lto_client_id, telephone_number, int_area_code, mobile_number,
            emergency_contact_number, emergency_contact_name, emergency_contact_relationship, emergency_contact_address
        ) VALUES (
            :lto_client_id, :telephone_number, :int_area_code, :mobile_number,
            :emergency_contact_number, :emergency_contact_name, :emergency_contact_relationship, :emergency_contact_address
        )
        ON CONFLICT (lto_client_id) DO UPDATE SET
            telephone_number = EXCLUDED.telephone_number,
            int_area_code = EXCLUDED.int_area_code,
            mobile_number = EXCLUDED.mobile_number,
            emergency_contact_number = EXCLUDED.emergency_contact_number,
            emergency_contact_name = EXCLUDED.emergency_contact_name,
            emergency_contact_relationship = EXCLUDED.emergency_contact_relationship,
            emergency_contact_address = EXCLUDED.emergency_contact_address
    `
    contactData := map[string]interface{}{
        "lto_client_id":                  user.LTO_CLIENT_ID,
        "telephone_number":               user.Contact.TELEPHONE_NUMBER,
        "int_area_code":                  user.Contact.INT_AREA_CODE,
        "mobile_number":                  user.Contact.MOBILE_NUMBER,
        "emergency_contact_number":       user.Contact.EMERGENCY_CONTACT_NUMBER,
        "emergency_contact_name":         user.Contact.EMERGENCY_CONTACT_NAME,
        "emergency_contact_relationship": user.Contact.EMERGENCY_CONTACT_RELATIONSHIP,
        "emergency_contact_address":      user.Contact.EMERGENCY_CONTACT_ADDRESS,
    }
    if _, err := tx.NamedExec(contactQuery, contactData); err != nil {
        tx.Rollback()
        return fmt.Errorf("contact upsert failed: %w", err)
    }

    return tx.Commit()
}