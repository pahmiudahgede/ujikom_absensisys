package middleware

import (
	"absensibe/utils"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type SessionData struct {
	UserID       string    `json:"user_id"`
	NISN         string    `json:"nisn"`
	Name         string    `json:"name"`
	Role         string    `json:"role,omitempty"`
	SessionID    string    `json:"session_id"`
	AccessToken  string    `json:"access_token"`
	CreatedAt    time.Time `json:"created_at"`
	LastActivity time.Time `json:"last_activity"`
	DeviceInfo   string    `json:"device_info,omitempty"`
	IPAddress    string    `json:"ip_address,omitempty"`
}

type JWTClaims struct {
	UserID    string `json:"user_id"`
	NISN      string `json:"nisn"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

type UserContext struct {
	UserID    string `json:"user_id"`
	NISN      string `json:"nisn"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	SessionID string `json:"session_id"`
}

func generateSessionID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func getSecretKey() []byte {
	return []byte(os.Getenv("SECRET_KEY"))
}

func getAPIKey() string {
	return os.Getenv("API_KEY")
}

func extractTokenFromHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}

	return parts[1]
}

func createAccessToken(userID, nisn, name, role, sessionID string) (string, error) {
	accessTokenExpiry, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		accessTokenExpiry = 24 * time.Hour
	}

	claims := JWTClaims{
		UserID:    userID,
		NISN:      nisn,
		Name:      name,
		Role:      role,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecretKey())
}

func parseAccessToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getSecretKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func CreateSession(userID, nisn, name, role string, c *fiber.Ctx) (*SessionData, error) {
	sessionID := generateSessionID()

	accessToken, err := createAccessToken(userID, nisn, name, role, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %v", err)
	}

	sessionData := &SessionData{
		UserID:       userID,
		NISN:         nisn,
		Name:         name,
		Role:         role,
		SessionID:    sessionID,
		AccessToken:  accessToken,
		CreatedAt:    time.Now(),
		LastActivity: time.Now(),
		DeviceInfo:   c.Get("User-Agent"),
		IPAddress:    c.IP(),
	}

	sessionKey := fmt.Sprintf("session:%s", sessionID)
	err = utils.Redis.SetCache(sessionKey, sessionData, utils.TTL_LONG)
	if err != nil {
		return nil, fmt.Errorf("failed to store session: %v", err)
	}

	userSessionsKey := fmt.Sprintf("user_sessions:%s", userID)
	utils.Redis.SAdd(userSessionsKey, sessionID)
	utils.Redis.Expire(userSessionsKey, utils.TTL_WEEK)

	return sessionData, nil
}

func GetSession(sessionID string) (*SessionData, error) {
	sessionKey := fmt.Sprintf("session:%s", sessionID)

	var sessionData SessionData
	err := utils.Redis.GetCache(sessionKey, &sessionData)
	if err != nil {
		return nil, fmt.Errorf("session not found or expired")
	}

	return &sessionData, nil
}

func UpdateSessionActivity(sessionID string) error {
	sessionKey := fmt.Sprintf("session:%s", sessionID)

	sessionData, err := GetSession(sessionID)
	if err != nil {
		return err
	}

	sessionData.LastActivity = time.Now()

	return utils.Redis.SetCache(sessionKey, sessionData, utils.TTL_LONG)
}

func DestroySession(sessionID string) error {
	sessionKey := fmt.Sprintf("session:%s", sessionID)

	sessionData, err := GetSession(sessionID)
	if err == nil {
		userSessionsKey := fmt.Sprintf("user_sessions:%s", sessionData.UserID)
		utils.Redis.SRem(userSessionsKey, sessionID)
	}

	return utils.Redis.Delete(sessionKey)
}

func DestroyAllUserSessions(userID string) error {
	userSessionsKey := fmt.Sprintf("user_sessions:%s", userID)

	sessionIDs, err := utils.Redis.SMembers(userSessionsKey)
	if err != nil {
		return err
	}

	for _, sessionID := range sessionIDs {
		sessionKey := fmt.Sprintf("session:%s", sessionID)
		utils.Redis.Delete(sessionKey)
	}

	return utils.Redis.Delete(userSessionsKey)
}

func APIKeyValidator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		expectedAPIKey := getAPIKey()

		if apiKey == "" {
			return utils.BadRequestResponse(c, "API key is required")
		}

		if apiKey != expectedAPIKey {
			return utils.UnauthorizedResponse(c, "Invalid API key")
		}

		return c.Next()
	}
}

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := extractTokenFromHeader(authHeader)

		if tokenString == "" {
			return utils.UnauthorizedResponse(c, "Access token is required")
		}

		claims, err := parseAccessToken(tokenString)
		if err != nil {
			return utils.UnauthorizedResponse(c, "Invalid or expired token")
		}

		sessionData, err := GetSession(claims.SessionID)
		if err != nil {
			return utils.UnauthorizedResponse(c, "Session not found or expired")
		}

		go UpdateSessionActivity(claims.SessionID)

		userContext := &UserContext{
			UserID:    sessionData.UserID,
			NISN:      sessionData.NISN,
			Name:      sessionData.Name,
			Role:      sessionData.Role,
			SessionID: sessionData.SessionID,
		}
		c.Locals("user", userContext)

		return c.Next()
	}
}

func RoleRequired(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := GetUserFromContext(c)
		if user == nil {
			return utils.UnauthorizedResponse(c, "Authentication required")
		}

		for _, role := range allowedRoles {
			if user.Role == role {
				return c.Next()
			}
		}

		return utils.ForbiddenResponse(c, "Insufficient permissions")
	}
}

func RateLimiter(maxRequests int64, window time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		count, err := utils.Redis.IncrementCounter(key, window)
		if err != nil {

			fmt.Printf("Rate limiter error: %v\n", err)
			return c.Next()
		}

		if count > maxRequests {

			return c.Status(fiber.StatusTooManyRequests).JSON(utils.APIResponse{
				Meta: utils.MetaResponse{
					Status:  429,
					Message: "Too many requests, please try again later",
				},
			})
		}

		c.Set("X-RateLimit-Limit", fmt.Sprintf("%d", maxRequests))
		c.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", maxRequests-count))

		return c.Next()
	}
}

func GetUserFromContext(c *fiber.Ctx) *UserContext {
	user := c.Locals("user")
	if user == nil {
		return nil
	}

	userContext, ok := user.(*UserContext)
	if !ok {
		return nil
	}

	return userContext
}

func IsAuthenticated(c *fiber.Ctx) bool {
	return GetUserFromContext(c) != nil
}

func GetCurrentUserID(c *fiber.Ctx) string {
	user := GetUserFromContext(c)
	if user == nil {
		return ""
	}
	return user.UserID
}

func GetCurrentUserRole(c *fiber.Ctx) string {
	user := GetUserFromContext(c)
	if user == nil {
		return ""
	}
	return user.Role
}
