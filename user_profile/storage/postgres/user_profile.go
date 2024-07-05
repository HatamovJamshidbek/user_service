package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "model/genproto"
	"model/helper"
	"time"
)

type UserProfileRepository struct {
	Db *sql.DB
}

func NewUserProfileRepository(db *sql.DB) *UserProfileRepository {
	return &UserProfileRepository{Db: db}
}

func (repo UserProfileRepository) CreateUserProfile(user_profile *pb.UserRequest) (*pb.Void, error) {
	query := `INSERT INTO user_profiles 
	          (user_id, bio, role, location, avatar_url, website, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := repo.Db.Exec(query, user_profile.UserId, user_profile.Bio, user_profile.UserRole, user_profile.Location, user_profile.AvatarUrl, user_profile.Website, time.Now())
	if err != nil {
		fmt.Println("++++++++", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (repo UserProfileRepository) UpdateUserProfile(user_profile *pb.UserRequest) (*pb.Void, error) {

	_, err := repo.Db.Exec("update user_profiles set user_id=$1,full_name=$2,bio=$3,role=$4,location=$5,avatar_url=$6,website=$7,updated_at=$8 where user_id=$9 and deleted_at=0)", user_profile.UserId, user_profile.FullName, user_profile.Bio, user_profile.UserRole, user_profile.Location, user_profile.AvatarUrl, user_profile.Website, time.Now(), user_profile.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo UserProfileRepository) DeleteUserProfile(id *pb.UserID) (*pb.Void, error) {
	_, err := repo.Db.Exec("update user_profiles set deleted_at=$1 where user_id=$2 and deleted_at is null)", time.Now(), id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo UserProfileRepository) GetUserProfileByIdS(id *pb.UserID) (*pb.GetIDUserRespons, error) {
	row := repo.Db.QueryRow("select user_id,full_name,bio,role,location,avatar_url,website from user_profiles  where user_id=$1 and deleted_at is null)", id)
	user_profile := pb.GetIDUserRespons{}
	err := row.Scan(&user_profile.Userid, &user_profile.FullName, &user_profile.Bio, &user_profile.Role, &user_profile.Location, &user_profile.AvatarUrl, &user_profile.Website)

	if err != nil {
		return nil, err
	}
	return &user_profile, nil
}

func (repo UserProfileRepository) GetUserProfile(user_profile *pb.UserFilter) (*pb.FilterUser, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(user_profile.UserID) > 0 {
		params["user_id"] = user_profile.UserID
		filter += " and user_id = :user_id "

	}
	if len(user_profile.FullName) > 0 {
		params["full_name"] = user_profile.FullName
		filter += " and full_name = :full_name "

	}

	if len(user_profile.UserRole) > 0 {
		params["role"] = user_profile.UserRole
		filter += " and role = :role "

	}
	if len(user_profile.Location) > 0 {
		params["location"] = user_profile.Location
		filter += " and location = :location "

	}
	// if len(user_profile.AvatarUrl) > 0 {
	// 	params["avatar_url"] = user_profile.AvatarUrl
	// 	filter += " and avatar_url = :avatar_url "

	// }
	// if len(user_profile.Website) > 0 {
	// 	params["website"] = user_profile.Website
	// 	filter += " and website = :website "

	// }

	if user_profile.Filter.Limit > 0 {
		params["limit"] = user_profile.Filter.Limit
		limit = ` LIMIT :limit`

	}
	if user_profile.Filter.Offset > 0 {
		params["offset"] = user_profile.Filter.Offset
		limit = ` OFFSET :offset`

	}
	query := "select user_id,full_name ,bio, role, location, avatar_url, website from user_profiles  where  deleted_at is null"

	query = query + filter + limit + offset
	query, arr = helper.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var user_profiles []*pb.GetIDUserRespons
	for rows.Next() {
		var user_profile pb.GetIDUserRespons
		err := rows.Scan(&user_profile.Id, &user_profile.Userid, &user_profile.FullName, &user_profile.Bio, &user_profile.Role, &user_profile.Location, &user_profile.AvatarUrl, &user_profile.Website)
		if err != nil {
			return nil, err
		}
		user_profiles = append(user_profiles, &user_profile)
	}
	return &pb.FilterUser{
		FilterRespons: user_profiles,
	}, nil
}
