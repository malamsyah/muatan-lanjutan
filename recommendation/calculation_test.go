package recommendation

import (
	"github.com/malamsyah/muatan-lanjutan/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Calculate(t *testing.T) {
	t.Run("When data in redis is not present", func(t *testing.T) {
		redis := model.Redis{
			RouteData: nil,
			TruckData: nil,
		}
		service := Service{redis: redis}

		shipment := model.Shipment{
			OriginLocation:      "A",
			DestinationLocation: "B",
			VehicleType:         "FUSO",
		}

		actual := service.Calculate(shipment)

		assert.Equal(t, []model.RecommendedJob(nil), actual)
	})
	t.Run("When data data is present", func(t *testing.T) {
		transporterA := model.Transporter{Name: "TransporterA"}
		transporterB := model.Transporter{Name: "TransporterB"}
		routeData := []model.RouteData{
			{
				Key: "A-B-FUSO",
				Val: []model.Transporter{
					transporterA,
					transporterB,
				},
			},
		}

		truckData := []model.TruckData{
			{
				Key: "A-FUSO",
				Val: []model.Transporter{
					transporterA,
				},
			},
			{
				Key: "A-CDD",
				Val: []model.Transporter{
					transporterB,
				},
			},
		}

		redis := model.Redis{
			RouteData: routeData,
			TruckData: truckData,
		}

		shipment := model.Shipment{
			OriginLocation:      "A",
			DestinationLocation: "B",
			VehicleType:         "FUSO",
		}

		service := Service{redis: redis}

		actual := service.Calculate(shipment)
		expected := []model.RecommendedJob{
			{
				Shipment:    shipment,
				Transporter: transporterA,
			},
		}
		assert.Equal(t, expected, actual)
	})
}
