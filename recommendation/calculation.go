package recommendation

import (
	"fmt"
	"github.com/malamsyah/muatan-lanjutan/model"
)

type Service struct {
	redis model.Redis
}

func (s *Service) Calculate(shipment model.Shipment) []model.RecommendedJob {
	var recommendationJobs []model.RecommendedJob

	routeKey := fmt.Sprint(shipment.OriginLocation, "-", shipment.DestinationLocation, "-", shipment.VehicleType)
	routeList := s.redis.GetRouteTransporter(routeKey)

	truckKey := fmt.Sprint(shipment.OriginLocation, "-", shipment.VehicleType)
	truckList := s.redis.GetTruckTransporter(truckKey)

	transporterMap := make(map[string]bool)

	for _, v := range routeList {
		transporterMap[v.Name] = true
	}

	for _, v := range truckList {
		if transporterMap[v.Name] {
			newJob := model.RecommendedJob{
				Shipment:    shipment,
				Transporter: v,
			}
			recommendationJobs = append(recommendationJobs, newJob)
		}
	}

	return recommendationJobs
}
