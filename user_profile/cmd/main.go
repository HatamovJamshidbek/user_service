package main

import (
	"fmt"
	"log"
	"model/config"
	pb "model/genproto"
	"model/service"
	"model/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cnf := config.Config{}
	db, err := postgres.ConnDB(&cnf)
	if err != nil {
		log.Fatal("error --> ", err)
		return
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("error --> ", err)
		return
	}
	fmt.Println("Server lesterner 500051...")
	lib := service.NewUserProfileService(postgres.NewUserProfileRepository(db))

	s := grpc.NewServer() // yangi grpc server yaratildi

	pb.RegisterUserProfilServerServer(s, lib) // serverda service ni metodlari ishlash uchuin register qilinadi

	if err = s.Serve(listener); err != nil {
		log.Fatal("error --> ", err.Error())
	}

}
