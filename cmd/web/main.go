package main

import (
	configCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/config"
	firebaseCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/admin"
	firebaseAuthCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/auth"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	pgCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/pg"

	auRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/auth/firebase"
	uRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user/pg"

	uUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/user"

	uDelivery "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/delivery/user/http"

	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := configCommon.LoadConfig()
	pg, querier := pgCommon.New(cfg.DatabaseURL)
	defer pg.Close()

	app, err := firebaseCommon.NewFirebaseAdmin(cfg.CredentialType, cfg.CredentialValue)
	if err != nil {
		panic(err)
	}
	fAuth, err := firebaseAuthCommon.NewFirebaseAuth(app)
	if err != nil {
		panic(err)
	}
	//fStg, err := firebaseStgCommon.NewFirebaseStorage(app, cfg.BucketName)
	//if err != nil {
	//	panic(err)
	//}

	h := httpCommon.NewHTTPServer()
	//localhost:3000/api/v1/
	api := h.Router.Group("/api/v1", gin.Logger(), httpCommon.CORS())

	aur := auRepo.NewFirebaseAuthRepository(fAuth)

	ur := uRepo.NewPGUserRepository(querier)
	uc := uUCase.NewUserUsecase(ur, aur)
	uDelivery.NewHTTPUserDelivery(api, uc, fAuth)

	log.Fatal(h.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
