package repository

import (
	pb "github.com/AkezhanOb1/business-service/api"
	"github.com/AkezhanOb1/business-service/config"
	"github.com/jackc/pgx/v4"
	"context"

)


//GetBusinessServicesRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessServicesRepository(ctx context.Context) (*pb.GetBusinessServicesResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)
	sqlQuery := `SELECT b.id, b.name, array_agg(bsc.id) FROM business_service b
				 INNER JOIN business_sub_categories_services
					ON b.id = business_sub_categories_services.business_services_id
				 INNER JOIN business_sub_category bsc
					ON business_sub_categories_services.sub_categories_id = bsc.id
	             GROUP BY b.id;`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	var services pb.GetBusinessServicesResponse

	for rows.Next() {
		var service pb.BusinessService

		err = rows.Scan(
			&service.BusinessServiceID,
			&service.BusinessServiceName,
			&service.SubCategories,
		)

		if err != nil {
			return nil, err
		}

		services.BusinessServices = append(services.BusinessServices, &service)
	}

	return &services, nil
}
