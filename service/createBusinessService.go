
package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-service/api"
	db "github.com/AkezhanOb1/business-service/repository"

)

//CreateBusinessService is
func (*Server) CreateBusinessService(ctx context.Context, request *pb.CreateBusinessServiceRequest) (*pb.CreateBusinessServiceResponse, error) {

	var serviceName = request.GetBusinessServiceName()
	var subCategories = request.GetSubCategories()

	service, err := db.CreateBusinessService(ctx, serviceName, subCategories)
	if err != nil {
		return nil, err
	}

	return service, nil
}



