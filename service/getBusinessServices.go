
package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-service/api"
	db "github.com/AkezhanOb1/business-service/repository"
	"github.com/golang/protobuf/ptypes/empty"
)

//GetBusinessServices is
func (*Server) GetBusinessServices(ctx context.Context, emp *empty.Empty) (*pb.GetBusinessServicesResponse, error) {
	services, err := db.GetBusinessServicesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return services, nil
}
