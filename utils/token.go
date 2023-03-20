package utils

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

const (
	ATokenExpireDuration = 2 * time.Hour
	RTokenExpireDuration = 30 * 24 * time.Hour
	TokenIssuer          = ""
)

var (
	mySecret	= []byte("xxxx")
	ErrorInvalidToken = errors.New("verify Token Failed")
)

type MyClaims struct {
	UserID uint `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func getJWTTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return mySecret, nil
}

//Gentoken 颁发token access token and refresh token
func GenToken(UserID uint, Username string) (atoken, rtoken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: getJWTTime(ATokenExpireDuration),
		Issuer: TokenIssuer,
	}
	at := MyClaims{
		UserID,
		Username,
        rc,
	}
	
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, at).SignedString(mySecret)

	//refresh token 不需要保存任何用户信息
	rt := rc
	rt.ExpiresAt = getJWTTime(RTokenExpireDuration)
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rt).SignedString(mySecret)
	return

}

//VerifyToken 验证Token

func VerifyToken(tokenID string) (*MyClaims, error) {
	var myc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenID, myc, keyFunc)
	if err!= nil {
		return nil, err
	}
	if !token.Valid {
		err = ErrorInvalidToken
		return nil, err
	}
	return myc, nil
}

//RefreshToken 通过refresh token刷新atoken
func RefreshToken(atoken, rtoken string) (newAtoken, newRtoken string, err error) {
	//rtoken无直接返回
	if _, err = jwt.Parse(rtoken, keyFunc); err!= nil {
		return
	}

	//从旧access token中解析出claims数据
	var claim MyClaims
	_, err = jwt.ParseWithClaims(atoken, &claim, keyFunc)
	//判断是否因access token正常过期导致的错误
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claim.UserID, claim.Username)
	}
	return
}

