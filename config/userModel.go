package config

import "time"
//gorm会自己解析User ==users
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex;not null"`
    Password  string    `gorm:"not null"`
    Role      string    `gorm:"type:varchar(20)"`
    CreatedAt time.Time  `gorm:"autoCreateTime`
} 
