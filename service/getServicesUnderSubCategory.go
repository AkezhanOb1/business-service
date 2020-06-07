
package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-service/api"
	db "github.com/AkezhanOb1/business-service/repository"

)

//GetServicesUnderSubCategory is
func (*Server) GetServicesUnderSubCategory(ctx context.Context, request *pb.GetServicesUnderSubCategoryRequest) (*pb.GetServicesUnderSubCategoryResponse, error) {
	categoryID := request.GetSubCategoryID()

	services, err := db.GetServicesUnderSubCategoryRepository(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return services, nil
}

