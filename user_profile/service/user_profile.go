package service

import (
	"context"
	pb "model/genproto"
	postgres "model/storage/postgres"
	// "gorm.io/driver/postgres"
)

type UserP struct {
	pb.UnimplementedUserProfilServerServer
	UserP *postgres.UserProfileRepository
}

func NewUserProfileService(userP *postgres.UserProfileRepository) *UserP {
	return &UserP{UserP: userP}
}

func (u *UserP) CreateUserProfiles(cntx context.Context, req *pb.UserRequest) (*pb.Void, error) {
	resp, err := u.UserP.CreateUserProfile(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (u *UserP) UpdateUserProfile(cntx context.Context, req *pb.UserRequest) (*pb.Void, error) {
	resp, err := u.UserP.UpdateUserProfile(req)
	if err != nil {
		return nil,  err
	}

	return resp, nil
}

func (u *UserP) DeleteUserProfile(cntx context.Context, req *pb.UserID) (*pb.Void, error) {
	resp, err := u.UserP.DeleteUserProfile(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (u *UserP) GetByUserID(cntx context.Context, req *pb.UserID) (*pb.GetIDUserRespons, error) {
	resp, err := u.UserP.GetUserProfileByIdS(req)
	if err != nil {
		return nil,  err
	}
	return resp, nil
}


func (u *UserP) GetAllUser(cntx context.Context, req *pb.UserFilter) (*pb.FilterUser, error)	{
	resp, err := u.UserP.GetUserProfile(req)
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}