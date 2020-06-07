
package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-service/api"
	db "github.com/AkezhanOb1/business-service/repository"
)

//GetBusinessService is
func (*Server) GetBusinessService(ctx context.Context, request *pb.GetBusinessServiceRequest) (*pb.GetBusinessServiceResponse, error) {
	serviceID := request.GetBusinessServiceID()

	service, err := db.GetBusinessServiceRepository(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	return service, nil
}
