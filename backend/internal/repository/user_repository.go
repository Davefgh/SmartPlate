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

    // Insert user with explicit parameter binding
    err := tx.QueryRow(`
        INSERT INTO users (
            last_name, first_name, middle_name, email, 
            password, role, status, lto_client_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING user_id, created, updated
    `,
    user.LAST_NAME,
    user.FIRST_NAME,
    user.MIDDLE_NAME,
    user.EMAIL,
    user.PASSWORD,
    user.ROLE,
    user.STATUS,
    user.LTO_CLIENT_ID,
    ).Scan(&user.USER_ID, &user.CREATED, &user.UPDATED)
    
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("user insertion failed: %w", err)
    }

    // Insert contact with proper null handling
    _, err = tx.NamedExec(`
        INSERT INTO contacts (
            lto_client_id, telephone_number, mobile_number,
            emergency_contact_number, emergency_contact_name,
            emergency_contact_relationship, emergency_contact_address
        ) VALUES (
            :lto_client_id, :telephone_number, :mobile_number,
            :emergency_contact_number, :emergency_contact_name,
            :emergency_contact_relationship, :emergency_contact_address
        )`,
        map[string]interface{}{
            "lto_client_id":                  user.LTO_CLIENT_ID,
            "telephone_number":               toNullString(user.Contact.TELEPHONE_NUMBER),
            "mobile_number":                 toNullString(user.Contact.MOBILE_NUMBER),
            "emergency_contact_number":      toNullString(user.Contact.EMERGENCY_CONTACT_NUMBER),
            "emergency_contact_name":        toNullString(user.Contact.EMERGENCY_CONTACT_NAME),
            "emergency_contact_relationship": toNullString(user.Contact.EMERGENCY_CONTACT_RELATIONSHIP),
            "emergency_contact_address":     toNullString(user.Contact.EMERGENCY_CONTACT_ADDRESS),
        })
     // address
     _, err = tx.NamedExec(`
     INSERT INTO addresses (
         lto_client_id, house_no, street, province, 
         city_municipality, barangay, zip_code
     ) VALUES (
         :lto_client_id, :house_no, :street, :province, 
         :city_municipality, :barangay, :zip_code
     )`,
     map[string]interface{}{
         "lto_client_id":      user.LTO_CLIENT_ID,
         "house_no":          toNullString(user.Address.HOUSE_NO),
         "street":            toNullString(user.Address.STREET),
         "province":          toNullString(user.Address.PROVINCE),
         "city_municipality": toNullString(user.Address.CITY_MUNICIPALITY),
         "barangay":          toNullString(user.Address.BARANGAY),
         "zip_code":          toNullString(user.Address.ZIP_CODE),
     })
     //medical informatio
    _, err = tx.NamedExec(`
    INSERT INTO medical_information(
        lto_client_id, gender, blood_type, complexion, eye_color, hair_color, weight, height, organ_donor)
        VALUES(:lto_client_id, :gender, :blood_type, :complexion, :eye_color, :hair_color, :weight, :height, :organ_donor

    )`,
    map[string]interface{}{
        "lto_client_id": user.LTO_CLIENT_ID,
        "gender": toNullString(user.MedicalInformation.GENDER),
        "blood_type": toNullString(user.MedicalInformation.BLOOD_TYPE),
        "complexion": toNullString(user.MedicalInformation.COMPLEXION),
        "eye_color": toNullString(user.MedicalInformation.EYE_COLOR),
        "hair_color": toNullString(user.MedicalInformation.HAIR_COLOR),
        "weight":       user.MedicalInformation.WEIGHT,
        "height":       user.MedicalInformation.HEIGHT,
        "organ_donor":  user.MedicalInformation.ORGAN_DONOR, 
    })
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("contact insertion failed: %w", err)
    }

    return tx.Commit()
}

// Helper function to handle null strings
func toNullString(s *string) interface{} {
    if s == nil || *s == "" {
        return nil
    }
    return *s
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
            c.emergency_contact_address AS "contact.emergency_contact_address",
            a.address_id AS "address.address_id",
            a.house_no AS "address.house_no",
            a.street AS "address.street",
            a.province AS "address.province",
            a.city_municipality AS "address.city_municipality",
            a.barangay AS "address.barangay",
            a.zip_code AS "address.zip_code",
            a.lto_client_id AS "address.lto_client_id",
            m.medical_id AS "medical_information.medical_id",
            m.gender AS "medical_information.gender",
            m.blood_type AS "medical_information.blood_type",
            m.complexion AS "medical_information.complexion",
            m.eye_color AS "medical_information.eye_color",
            m.hair_color AS "medical_information.hair_color",
            m.weight AS "medical_information.weight",
            m.height AS "medical_information.height",
            m.organ_donor AS "medical_information.organ_donor"
        FROM users u
        LEFT JOIN contacts c ON u.lto_client_id = c.lto_client_id
        LEFT JOIN addresses a ON u.lto_client_id = a.lto_client_id
        LEFT JOIN medical_information m ON u.lto_client_id = m.lto_client_id
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
            c.emergency_contact_address AS "contact.emergency_contact_address",
            a.address_id AS "address.address_id",
            a.house_no AS "address.house_no",
            a.street AS "address.street",
            a.province AS "address.province",
            a.city_municipality AS "address.city_municipality",
            a.barangay AS "address.barangay",
            a.zip_code AS "address.zip_code",
            a.lto_client_id AS "address.lto_client_id",
            m.medical_id AS "medical_information.medical_id",
            m.gender AS "medical_information.gender",
            m.blood_type AS "medical_information.blood_type",
            m.complexion AS "medical_information.complexion",
            m.eye_color AS "medical_information.eye_color",
            m.hair_color AS "medical_information.hair_color",
            m.weight AS "medical_information.weight",
            m.height AS "medical_information.height"
            m.organ_donor AS "medical_information.organ_donor"
        FROM users u
        LEFT JOIN contacts c ON u.lto_client_id = c.lto_client_id
        LEFT JOIN addresses a ON u.lto_client_id = a.lto_client_id
        LEFT JOIN medical_information m ON u.lto_client_id = m.lto_client_id
        WHERE u.user_id = $1
    `
    err := r.db.Get(&user, query, user_id)
    return user, err
}

func (r *UserRepository) GetByLTOClientID(ltoClientID string) (models.User, error) {
    var user models.User
    query := `
        SELECT 
            u.*,
            c.*,
            a.*,
            m.*
        FROM users u
        LEFT JOIN contacts c ON u.lto_client_id = c.lto_client_id
        LEFT JOIN addresses a ON u.lto_client_id = a.lto_client_id
        LEFT JOIN medical_information m ON u.lto_client_id = m.lto_client_id
        WHERE u.lto_client_id = $1
    `
    err := r.db.Get(&user, query, ltoClientID)
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

    if _, err := tx.Exec("DELETE FROM addresses WHERE lto_client_id = $1", ltoID); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete address: %w", err)
    }
    if _, err := tx.Exec("DELETE FROM medical_information WHERE lto_client_id = $1", ltoID); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete medical info: %w", err)
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
    addressQuery := `
        INSERT INTO addresses (
            lto_client_id, house_no, street, province, 
            city_municipality, barangay, zip_code
        ) VALUES (
            :lto_client_id, :house_no, :street, :province, 
            :city_municipality, :barangay, :zip_code
        )
        ON CONFLICT (lto_client_id) DO UPDATE SET
            house_no = EXCLUDED.house_no,
            street = EXCLUDED.street,
            province = EXCLUDED.province,
            city_municipality = EXCLUDED.city_municipality,
            barangay = EXCLUDED.barangay,
            zip_code = EXCLUDED.zip_code
    `
    _, err = tx.NamedExec(addressQuery, map[string]interface{}{
        "lto_client_id":      user.LTO_CLIENT_ID,
        "house_no":          toNullString(user.Address.HOUSE_NO),
        "street":            toNullString(user.Address.STREET),
        "province":          toNullString(user.Address.PROVINCE),
        "city_municipality": toNullString(user.Address.CITY_MUNICIPALITY),
        "barangay":          toNullString(user.Address.BARANGAY),
        "zip_code":          toNullString(user.Address.ZIP_CODE),
    })
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("address upsert failed: %w", err)
    }
    // In Update function
medicalQuery := `
INSERT INTO medical_information (
    lto_client_id, gender, blood_type, complexion, 
    eye_color, hair_color, weight, height, organ_donor
) VALUES (
    :medical_information.lto_client_id, 
    :medical_information.gender, 
    :medical_information.blood_type,
    :medical_information.complexion,
    :medical_information.eye_color,
    :medical_information.hair_color,
    :medical_information.weight,
    :medical_information.height,
    :medical_information.organ_donor
)
ON CONFLICT (lto_client_id) DO UPDATE SET
    gender = EXCLUDED.gender,
    blood_type = EXCLUDED.blood_type,
    complexion = EXCLUDED.complexion,
    eye_color = EXCLUDED.eye_color,
    hair_color = EXCLUDED.hair_color,
    weight = EXCLUDED.weight,
    height = EXCLUDED.height,
    organ_donor = EXCLUDED.organ_donor
`

_, err = tx.NamedExec(medicalQuery, user)
if err != nil {
tx.Rollback()
return fmt.Errorf("medical info upsert failed: %w", err)
}

    return tx.Commit()
}
func toNullInt(i *int) interface{} {
    if i == nil {
        return nil
    }
    return *i
}

func toNullBool(b *bool) interface{} {
    if b == nil {
        return nil
    }
    return *b
}