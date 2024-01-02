package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ChaosTh3ori3/RangerEventManager.Api/config"
	"github.com/coreos/go-oidc"
)

type AuthenticateMiddleware struct {
	keyCloakSettings *config.KeyCloakSettings
}

func NewAuthMiddleware(keyCloakSettings *config.KeyCloakSettings) AuthenticateMiddleware {
	return AuthenticateMiddleware{
		keyCloakSettings: keyCloakSettings,
	}
}

func (am AuthenticateMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, am.keyCloakSettings.DiscoveryUrl)
		if err != nil {
			fmt.Printf("Error setting up OIDC provider: %v\n", err)
			os.Exit(1)
		}

		// Hier kannst du Logik für die Middleware hinzufügen
		fmt.Println("Middleware ausgeführt!")
		token, err := getTokenFromRequest(r)
		if err != nil {
			// Fehlerbehandlung: Redirect zum Keycloak-Login oder Fehlerseite
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Überprüfe den Token gegen Keycloak
		_, err = provider.Verifier(&oidc.Config{ClientID: am.keyCloakSettings.ClientId}).Verify(ctx, token)
		if err != nil {
			// Fehlerbehandlung: Redirect zum Keycloak-Login oder Fehlerseite
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token im Kontext speichern oder andere Verarbeitung, wenn gewünscht

		// Rufe die nächste Handler-Funktion auf
		next.ServeHTTP(w, r)
	})
}

func getTokenFromRequest(r *http.Request) (string, error) {
	// Extrahiere den Authorization-Header aus dem Request
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		// Wenn der Authorization-Header nicht vorhanden ist, gib einen Fehler zurück
		return "", fmt.Errorf("Authorization header not found")
	}

	// Der Authorization-Header sollte das Format "Bearer <Token>" haben
	// Trenne den Header, um das Token zu erhalten
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("Invalid Authorization header format")
	}

	return parts[1], nil
}
