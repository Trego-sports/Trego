package web

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"trego-backend/api-gateway/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	config      *config.Config
	oauthConfig *oauth2.Config
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(cfg *config.Config) *AuthHandler {
	oauthConfig := &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleRedirectURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &AuthHandler{
		config:      cfg,
		oauthConfig: oauthConfig,
	}
}

// GoogleLoginHandler initiates the Google OAuth flow
func (h *AuthHandler) GoogleLoginHandler(c *gin.Context) {
	// Generate a random state token for CSRF protection
	state, err := generateStateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate state token",
		})
		return
	}

	// Store state in session/cookie for validation in callback
	// For now, we'll just include it in the URL
	// In production, you'd want to store this in a session or encrypted cookie

	// Generate the OAuth URL
	url := h.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

	c.JSON(http.StatusOK, gin.H{
		"redirectUrl": url,
	})
}

// GoogleCallbackHandler handles the OAuth callback from Google
func (h *AuthHandler) GoogleCallbackHandler(c *gin.Context) {
	// Get the authorization code from query params
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization code not provided",
		})
		return
	}

	// Optional: Validate state parameter for CSRF protection
	// state := c.Query("state")
	// You would validate this against the stored state here

	// Exchange the authorization code for an access token
	token, err := h.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to exchange token",
		})
		return
	}

	// Fetch user info from Google
	userInfo, err := h.fetchGoogleUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user info",
		})
		return
	}

	// Log user info for debugging (remove in production)
	fmt.Printf("User authenticated: %+v\n", userInfo)

	// For now, just redirect to the frontend dashboard
	// In a real implementation, you would:
	// 1. Create or update user in database
	// 2. Create a session or JWT token
	// 3. Set authentication cookie
	c.Redirect(http.StatusFound, h.config.FrontendURL+"/dashboard")
}

// fetchGoogleUserInfo fetches user information from Google's userinfo endpoint
func (h *AuthHandler) fetchGoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo map[string]interface{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// generateStateToken generates a random state token for CSRF protection
func generateStateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
