package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gold-kou/go-housework/app/server/service"

	"github.com/gold-kou/go-housework/app/common"
	"github.com/gold-kou/go-housework/app/model/db"
	"github.com/gold-kou/go-housework/app/model/schemamodel"
	"github.com/gold-kou/go-housework/app/server/middleware"
	"github.com/gold-kou/go-housework/app/server/repository"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// ListFamilyMembers handler
func ListFamilyMembers(w http.ResponseWriter, r *http.Request) {
	// get jwt from header
	authHeader := r.Header.Get("Authorization")
	// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODM3MjIwNTMsImlhdCI6IjIwMjAtMDMtMDhUMTE6NDc6MzMuMTc4NjU5MyswOTowMCIsIm5hbWUiOiJ0ZXN0In0.YIyT1RJGcYbdynx1V4-6MhiosmTlHmKiyiG_GjxQeuw
	bearerToken := strings.Split(authHeader, " ")[1]

	// verify jwt
	authUser, err := middleware.VerifyToken(bearerToken)

	// service layer
	var f *db.Family
	var us []*db.User
	err = common.Transact(func(tx *gorm.DB) (err error) {
		familyRepo := repository.NewFamilyRepository(tx)
		userRepo := repository.NewUserRepository(tx)
		memberFamilyRepo := repository.NewMemberFamilyRepository(tx)
		f, us, err = service.NewListFamilyMembers(tx, familyRepo, userRepo, *memberFamilyRepo).Execute(authUser)
		return
	})

	// error handling
	switch err := err.(type) {
	case nil:
	case *common.BadRequestError:
		log.Warn(err)
		common.ResponseBadRequest(w, err.Message)
		return
	case *common.AuthorizationError:
		log.Warn(err)
		common.ResponseUnauthorized(w, err.Message)
		return
	default:
		log.Error(err)
		common.ResponseInternalServerError(w, err.Error())
		return
	}

	// http response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var members []schemamodel.Member
	for _, u := range us {
		m := schemamodel.Member{MemberId: int64(u.ID), MemberName: u.Name}
		members = append(members, m)
	}
	if err := json.NewEncoder(w).Encode(&schemamodel.ResponseListFamilyMembers{
		Family:  schemamodel.Family{FamilyId: int64(f.ID), FamilyName: f.Name},
		Members: members}); err != nil {
		log.Error(err)
		panic(err)
	}
}
