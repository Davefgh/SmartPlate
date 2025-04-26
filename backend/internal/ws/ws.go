// internal/ws/ws.go
package ws

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"

    "smartplate-api/internal/models"
    "smartplate-api/internal/repository"
)

var Upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin:     func(r *http.Request) bool { return true },
}

type PlateCheckRequest struct {
    Plate     string `json:"plate"`
    Timestamp string `json:"timestamp"`
}

type PlateCheckResponse struct {
    Plate   string      `json:"plate"`
    Status  string      `json:"status"`              // valid, not_found, expired, error
    Details *DetailPack `json:"details,omitempty"`
}

type DetailPack struct {
    RegistrationForm *models.RegistrationForm `json:"registration_form,omitempty"`
    Plates           []models.Plate           `json:"plates,omitempty"`
    User             *models.User             `json:"user_record,omitempty"`
}

// ScannerWS now takes plateRepo, regFormRepo, and userRepo.
func ScannerWS(
    plateRepo   repository.PlateRepository,
    regFormRepo repository.RegistrationFormRepository,
    userRepo    *repository.UserRepository,
) echo.HandlerFunc {
    return func(c echo.Context) error {
        ws, err := Upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
        if err != nil {
            return err
        }
        defer ws.Close()

        for {
            _, msg, err := ws.ReadMessage()
            if err != nil {
                log.Println("ws read error:", err)
                break
            }

            var req PlateCheckRequest
            if err := json.Unmarshal(msg, &req); err != nil {
                log.Println("json unmarshal error:", err)
                ws.WriteJSON(PlateCheckResponse{Status: "bad_request"})
                continue
            }

            // 1) lookup a single plate record
            rec, err := plateRepo.GetByPlateNumber(c.Request().Context(), req.Plate)
            validity := "error"
            if err != nil {
                log.Println("db lookup error:", err)
            } else if rec == nil {
                validity = "not_found"
            } else if rec.PLATE_EXPIRATION_DATE.Before(time.Now()) {
                validity = "expired"
            } else {
                validity = "valid"
            }

            var details *DetailPack
            if rec != nil {
                // use the actual field name from your Plate model:
                vehicleID := rec.VEHICLE_ID

                // 2) fetch the registration form for this vehicle
                regForm, rErr := regFormRepo.GetByVehicleID(c.Request().Context(), vehicleID)
                if rErr != nil {
                    log.Println("registration form lookup err:", rErr)
                }

                // 3) fetch all plates for this same vehicle
                plates, pErr := plateRepo.GetPlatesByVehicleID(c.Request().Context(), vehicleID)
                if pErr != nil {
                    log.Println("plates lookup err:", pErr)
                }

                // 4) fetch the user via the form’s LTOClientID
                var usr *models.User
                if regForm != nil {
                    u, uErr := userRepo.GetByLTOClientID(regForm.LTOClientID)
                    if uErr != nil {
                        log.Println("user lookup err:", uErr)
                    } else {
                        usr = &u
                    }
                }
                details = &DetailPack{
                    RegistrationForm: regForm,
                    Plates:           plates,
                    User:             usr,
                }

				details = &DetailPack{
					RegistrationForm: regForm,
					Plates:           plates,
					User:             usr,
				}
            }

            resp := PlateCheckResponse{
                Plate:   req.Plate,
                Status:  validity,
                Details: details,
            }
            log.Printf("⟶ sending response: %+v\n", resp)

            if err := ws.WriteJSON(resp); err != nil {
                log.Println("ws write error:", err)
                break
            }
        }

        return nil
    }
}
