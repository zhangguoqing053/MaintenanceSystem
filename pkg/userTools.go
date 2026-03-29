package pkg

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
    "errors"
)

var jwtKey = []byte("$2a$10$oEma9thLhzzteJ7IjUhhmOruqKeOTvrC3YgVi65efK9Lf./g7aC1.")

// 生成密码哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 校验密码
func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//生成JwtToken
func GenerateToken(userID uint, username string) (string, error) {

    claims := jwt.MapClaims{
        "user_id":  userID,                         // 自定义字段：用户ID
        "username": username,                       // 自定义字段：用户名
        "exp":      time.Now().Add(24 * time.Hour).Unix(), // 过期时间（24小时）
        "iat":      time.Now().Unix(),              // 签发时间（可选）
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil

}



func ParseToken(tokenString string) (jwt.MapClaims, error) {

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("invalid token")
}