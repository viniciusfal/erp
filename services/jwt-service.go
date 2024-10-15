package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey     string
	issuer        string
	refreshSecret string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey:     "generic-key",
		refreshSecret: "generic-refresh-key",
		issuer:        "erp-api",
	}
}

type Claim struct {
	Sub string `json:"sub"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id string) (string, string, error) {
	// Gerar o token de acesso
	accessClaim := &Claim{
		Sub: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)
	accessTokenString, err := accessToken.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", "", err
	}

	// Gerar o token de refresh
	refreshClaim := &Claim{
		Sub: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 semana
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.refreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *jwtService) ValidateRefreshToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid refresh token: %v", token)
		}

		return []byte(s.refreshSecret), nil
	})

	return err == nil
}

// SetTokenInCookie armazena o token JWT em um cookie
func (s *jwtService) SetTokenInCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,  // Para proteger contra XSS
		Secure:   false, // Use true se estiver usando HTTPS
	}
	http.SetCookie(w, &cookie)
}

// GetTokenFromCookie obtém o token JWT do cookie
func (s *jwtService) GetTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
