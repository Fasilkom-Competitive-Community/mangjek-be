package main

import (
	"context"

	configCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/config"
	firebaseCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/admin"
	firebaseAuthCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/auth"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := configCommon.LoadConfig()
	ctx := context.Background()

	app, err := firebaseCommon.NewFirebaseAdmin(cfg.CredentialType, cfg.CredentialValue)
	if err != nil {
		panic(err)
	}
	fAuth, err := firebaseAuthCommon.NewFirebaseAuth(app)
	if err != nil {
		panic(err)
	}

	fAuth.SetCustomUserClaims(ctx, "WX7TVelcU4VILtCnkECGdrXBNx42", map[string]interface{}{
		"role": "admin",
	})
}
