package models

type Repeater struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `json:"name"`
    IP          string `json:"ip"`
    Frequency   string `json:"frequency"`
    SIPUser     string `json:"sip_user"`
    SIPPassword string `json:"sip_password"`
    AISPort     int    `json:"ais_port"`
    // Другие параметры из мануала DR600
}