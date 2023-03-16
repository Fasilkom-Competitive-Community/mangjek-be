package main

import (
	configCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/config"
	firebaseCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/admin"
	firebaseAuthCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/auth"
	firebaseStgCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/firebase/storage"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	mapCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/map"
	pgCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/pg"
	uuidCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/uuid"
	oDelivery "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/delivery/order/http"

	dDelivery "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/delivery/driver/http"
	dRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver/pg"
	dUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/driver"

	oRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/order/pg"
	oUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/order"

	uDelivery "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/delivery/user/http"
	auRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/auth/firebase"
	uRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user/pg"
	uUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/user"

	fDelivery "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/delivery/file/http"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xendit/xendit-go"
)

func main() {
	cfg := configCommon.LoadConfig()
	store := pgCommon.New(cfg.DatabaseURL)
	defer store.Db.Close()

	xendit.Opt.SecretKey = cfg.XenditSecretKey

	app, err := firebaseCommon.NewFirebaseAdmin(cfg.CredentialType, cfg.CredentialValue)
	if err != nil {
		panic(err)
	}
	fAuth, err := firebaseAuthCommon.NewFirebaseAuth(app)
	if err != nil {
		panic(err)
	}
	_, err = firebaseStgCommon.NewFirebaseStorage(app, cfg.BucketName)
	if err != nil {
		panic(err)
	}

	gMap, err := mapCommon.NewMapCalculator(cfg.GMapAPIKey)
	if err != nil {
		panic(err)
	}

	uuid := uuidCommon.NewUUIDGenerator()

	h := httpCommon.NewHTTPServer()
	api := h.Router.Group("/api/v1", gin.Logger(), httpCommon.CORS())

	aur := auRepo.NewFirebaseAuthRepository(fAuth)

	ur := uRepo.NewPGUserRepository(store.Querier)
	uc := uUCase.NewUserUsecase(ur, aur)
	uDelivery.NewHTTPUserDelivery(api, uc, fAuth)

	dr := dRepo.NewPGDriverRepository(store.Querier)
	dc := dUCase.NewDriverUsecase(dr, aur)
	dDelivery.NewHTTPDriverDelivery(api, dc, fAuth)

	or := oRepo.NewPGOrderInquiryRepository(store)
	oc := oUCase.NewOrderUsecase(or, ur, dr, gMap, uuid)
	oDelivery.NewHTTPOrderDelivery(api, oc, fAuth)

	fDelivery.NewHTTPFileDelivery(api, fAuth)

	log.Fatal(h.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
