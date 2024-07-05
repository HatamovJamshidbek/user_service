package postgres

import (
	"model/config"
	pb "model/genproto"
	"reflect"
	"testing"
)

func TestRepoUser(t *testing.T) {
	cfg := config.Load()

	db, err := ConnDB(&cfg)
	if err != nil {
		return
	}

	repo := NewUserProfileRepository(db)

	user := pb.UserRequest{
		Id:        "",
		UserID:    "1",
		Fulname:   "Bobur",
		Bio:       "goland dasturchisi",
		Location:  "Uzbekiston, Surxondaryo",
		AvatarUrl: "bb",
		Website:   "aef",
		UserRole:  "",
		UpdatedAt: "",
	}

	// Create User method
	userRes, err := repo.CreateUserProfile(&user)
	if err != nil {
		t.Error("error while create user")
		return
	}
	if !reflect.DeepEqual(&user, userRes) {
		t.Error("error while create user", err)
		return
	}

	// Get User method
	id := pb.UserID{
		UserID: "",
	}
	userGet, err := repo.GetUserProfileByIdS(&id)
	if err != nil {
		t.Error("error while get user")
		return
	}

	if !reflect.DeepEqual(&userGet.Id, &id) {
		t.Error("error while get user")
		return
	}

	// Delete User method
	_, err = repo.DeleteUserProfile(&id)
	if err != nil {
		t.Error("error while delete user", err)
		return
	}
	
	userR := pb.UserRequest{ // create bo'lgan ma'lumotni kiritib tekshirib koramiz
		Id:        "",
		UserID:    "",
		Fulname:   "",
		Bio:       "",
		Location:  "",
		AvatarUrl: "",
		Website:   "",
		UserRole:  "",
		UpdatedAt: "",
	}

	_, err = repo.UpdateUserProfile(&userR)
	if err != nil {
		t.Error("error while update user", err)
	}

}
