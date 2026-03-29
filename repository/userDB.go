package repository

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "MaintenanceSystem/config"
    "MaintenanceSystem/pkg"
)


type  UserDB  struct{
	db *gorm.DB
}

func InitDB() (*UserDB,error){
	dsn := "host=127.0.0.1 user=jobuser password=job1234 dbname=mydb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil,err
    }
      return &UserDB{db:db},nil
}



//创建用户
func ( d *UserDB)CreateUser( u  *config.User) error {
        // 加密密码
    hash,_ := pkg.HashPassword(u.Password)
    u.Password = hash
    
    return d.db.Create(u).Error
}


//查询用户
func (d *UserDB) GetByUsername(username string) (*config.User, error) {
    var user config.User
    err := d.db.Where("username = ?", username).First(&user).Error
    return &user, err
}




