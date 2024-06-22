package auth
import (
  "errors"
  "time"
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "os"
)
// Retrieve secret key
var jwtKey = []byte(os.Getenv("JWT_SECRET"))
// Used to create JWT payload (data section)
// StandardClaims only contains the JWT's expiration date at the moment
type JWTClaim struct {
  Username string `json:"username"`
  Email    string `json:"email"`
  ID       uint   `json:"id"`
  jwt.StandardClaims
}
func GenerateJWT(email string, username string, id uint) (tokenString string, err error) {
  // Token expires in one week
  expirationTime := time.Now().Add(168 * time.Hour)
  claims:= &JWTClaim{
    Email: email,
    Username: username,
    ID: id,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expirationTime.Unix(),
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err = token.SignedString(jwtKey)
  return
}
func ValidateToken(signedToken string, context *gin.Context) (err error) {
  token, err := jwt.ParseWithClaims(
    signedToken,
    &JWTClaim{},
    func(token *jwt.Token) (interface{}, error) {
      return []byte(jwtKey), nil
    },
  )
  if err != nil {
    return
  }
  claims, ok := token.Claims.(*JWTClaim)
  if !ok {
    err = errors.New("couldn't parse claims")
    return
  }
  if claims.ExpiresAt < time.Now().Local().Unix() {
    err = errors.New("token expired")
    return
  }
  // Write Current User information to the context object
  context.Set("user_id", claims.ID)
  context.Set("user_email", claims.Email)
  context.Set("username", claims.Username)
  return
}