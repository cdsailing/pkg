package filters

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful/v3"
	"strings"
)

var (
	sharedSecret = []byte("mySigningKey")
)

func ValidJwt(authHeader string) bool {
	if !strings.HasPrefix(authHeader, "JWT ") {
		return false
	}

	jwtToken := strings.Split(authHeader, " ")
	if len(jwtToken) < 2 {
		return false
	}
	parts := strings.Split(jwtToken[1], ".")
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], sharedSecret)
	if err != nil {
		return false
	}

	return true
}

func AuthJwt(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if req.Request.Method == "OPTIONS" {
		chain.ProcessFilter(req, resp)
		return
	}
	authHeader := req.HeaderParameter("Authorization")
	_, success := ValidateToken(authHeader)
	if !success {
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}

type Claims struct {
	jwt.StandardClaims
}

func (claims *Claims) CreateToken(user interface{}) (signedToken string, success bool) {
	//claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	//claims.Subject = strconv.Itoa(user.Id)
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signedToken, err := token.SignedString([]byte("secret"))
	//if err != nil {
	//	return signedToken, false
	//}
	//return signedToken, true
	return "", true
}
func ValidateToken(tokenHeader string) (claims *Claims, success bool) {
	token, err := jwt.ParseWithClaims(tokenHeader, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}
	return
}
