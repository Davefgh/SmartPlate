package models

import (
	"encoding/json"
	"fmt"
	"time"
)

const dateLayout = "2006-01-02"

type Vehicle struct {
	VEHICLE_ID               *int             `json:"vehicle_id,omitempty" db:"vehicle_id"`
	VEHICLE_CATEGORY         string           `json:"vehicle_category" db:"vehicle_category"`
	MV_FILE_NUMBER           string           `json:"mv_file_number" db:"mv_file_number"`
	VEHICLE_MAKE             string           `json:"vehicle_make" db:"vehicle_make"`
	VEHICLE_SERIES           string           `json:"vehicle_series" db:"vehicle_series"`
	VEHICLE_TYPE             string           `json:"vehicle_type" db:"vehicle_type"`
	BODY_TYPE                string           `json:"body_type" db:"body_type"`
	YEAR_MODEL               string           `json:"year_model" db:"year_model"`
	ENGINE_MODEL             string           `json:"engine_model" db:"engine_model"`
	ENGINE_NUMBER            string           `json:"engine_number" db:"engine_number"`
	CHASSIS_NUMBER           string           `json:"chassis_number" db:"chassis_number"`
	PISTON_DISPLACEMENT      string           `json:"piston_displacement" db:"piston_displacement"`
	NUMBER_OF_CYLINDERS      string           `json:"number_of_cylinders" db:"number_of_cylinders"`
	FUEL_TYPE                string           `json:"fuel_type" db:"fuel_type"`
	COLOR                    string           `json:"color" db:"color"`
	GVW                      string           `json:"gvw" db:"gvw"`
	NET_WEIGHT               string           `json:"net_weight" db:"net_weight"`
	SHIPPING_WEIGHT          string           `json:"shipping_weight" db:"shipping_weight"`
	USAGE_CLASSIFICATION     string           `json:"usage_classification" db:"usage_classification"`
	FIRST_REGISTRATION_DATE  string           `json:"first_registration_date" db:"first_registration_date"`
	LATE_RENEWAL_DATE        string           `json:"late_renewal_date" db:"late_renewal_date"`
	REGISTRATION_EXPIRY_DATE string           `json:"registration_expiry_date" db:"registration_expiry_date"`
	LTO_OFFICE_CODE          string           `json:"lto_office_code" db:"lto_office_code"`
	CLASSIFICATION           string           `json:"classification" db:"classification"`
	DENOMINATION             string           `json:"denomination" db:"denomination"`
	OR_NUMBER                string           `json:"or_number" db:"or_number"`
	CR_NUMBER                string           `json:"cr_number" db:"cr_number"`
	LTO_CLIENT_ID            string           `json:"lto_client_id,omitempty" db:"lto_client_id"`
	PLATE                    Plate            `json:"plate" db:"plate"`
}

type Plate struct {
    VEHICLE_ID            int        `json:"vehicle_id" db:"vehicle_id"`
    PLATE_NUMBER          string     `json:"plate_number" db:"plate_number"`
    PLATE_TYPE            string     `json:"plate_type" db:"plate_type"`
    PLATE_ISSUE_DATE      time.Time  `json:"plate_issue_date" db:"plate_issue_date"`
    PLATE_EXPIRATION_DATE time.Time  `json:"plate_expiration_date" db:"plate_expiration_date"`
    STATUS                string     `json:"status" db:"status"`
}

func (p *Plate) UnmarshalJSON(data []byte) error {
    type Alias Plate
    aux := &struct {
        PlateIssueDate      string `json:"plate_issue_date"`
        PlateExpirationDate string `json:"plate_expiration_date"`
        *Alias
    }{
        Alias: (*Alias)(p),
    }
    
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    
    var err error
    p.PLATE_ISSUE_DATE, err = time.Parse(dateLayout, aux.PlateIssueDate)
    if err != nil {
        return fmt.Errorf("invalid plate_issue_date format: %w", err)
    }
    
    p.PLATE_EXPIRATION_DATE, err = time.Parse(dateLayout, aux.PlateExpirationDate)
    if err != nil {
        return fmt.Errorf("invalid plate_expiration_date format: %w", err)
    }
    
    return nil
}

type RegistrationForm struct {
    REGISTRATION_FORM_ID string    `json:"registration_form_id" db:"registration_form_id"`
    VEHICLE_ID           int       `json:"vehicle_id" db:"vehicle_id"`
    SUBMITTED_DATE       time.Time `json:"submitted_date" db:"submitted_date"`
    STATUS               string    `json:"status" db:"status"`
    LTO_CLIENT_ID        string    `json:"lto_client_id,omitempty" db:"lto_client_id"`
}

func (rf *RegistrationForm) UnmarshalJSON(data []byte) error {
    type Alias RegistrationForm
    aux := &struct {
        SubmittedDate string `json:"submitted_date"`
        *Alias
    }{
        Alias: (*Alias)(rf),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    if aux.SubmittedDate != "" {
        t, err := time.Parse(dateLayout, aux.SubmittedDate)
        if err != nil {
            return fmt.Errorf("invalid submitted_date format: %w", err)
        }
        rf.SUBMITTED_DATE = t
    }
    // if empty, leave rf.SUBMITTED_DATE as the zero time
    return nil
}

type Inspection struct {
    INSPECTION_ID     string    `json:"inspection_id" db:"inspection_id"`
    VEHICLE_ID        int       `json:"vehicle_id" db:"vehicle_id"`
    LTO_CLIENT_ID     string    `json:"lto_client_id" db:"lto_client_id"`
    INSPECTION_DATE   time.Time `json:"inspection_date" db:"inspection_date"`
    INSPECTION_STATUS string    `json:"inspection_status" db:"inspection_status"`
    NOTES             string    `json:"notes" db:"notes"`
}

func (i *Inspection) UnmarshalJSON(data []byte) error {
    type Alias Inspection
    aux := &struct {
        InspectionDate string `json:"inspection_date"`
        *Alias
    }{
        Alias: (*Alias)(i),
    }
    
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    
    var err error
    i.INSPECTION_DATE, err = time.Parse(dateLayout, aux.InspectionDate)
    if err != nil {
        return fmt.Errorf("invalid inspection_date format: %w", err)
    }
    
    return nil
}

type Payment struct {
    PAYMENT_ID      string    `json:"payment_id" db:"payment_id"`
    VEHICLE_ID      int       `json:"vehicle_id" db:"vehicle_id"`
    AMOUNT          string    `json:"amount" db:"amount"`
    PAYMENT_DATE    time.Time `json:"payment_date" db:"payment_date"`
    PAYMENT_METHOD  string    `json:"payment_method" db:"payment_method"`
    STATUS          string    `json:"status" db:"status"`
    RECEIPT_NUMBER  string    `json:"receipt_number" db:"receipt_number"`
    LTO_CLIENT_ID   string    `json:"lto_client_id,omitempty" db:"lto_client_id"`
}

func (p *Payment) UnmarshalJSON(data []byte) error {
    type Alias Payment
    aux := &struct {
        PaymentDate string `json:"payment_date"`
        *Alias
    }{
        Alias: (*Alias)(p),
    }
    
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    
    var err error
    p.PAYMENT_DATE, err = time.Parse(dateLayout, aux.PaymentDate)
    if err != nil {
        return fmt.Errorf("invalid payment_date format: %w", err)
    }
    
    return nil
}