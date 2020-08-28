package model

var TranporterA Transporter
var TranporterB Transporter
var TruckAA Truck
var TruckAB Truck
var TruckBA Truck
var TruckBA2 Truck
var RoutePref []RoutePreference

func init() {
	TranporterA = Transporter{Name: "TransporterA"}
	TranporterB = Transporter{Name: "TransporterB"}

	TruckAA = Truck{
		lastTransporter: TranporterA,
		currentLocation: "A",
		vehicleType:     "FUSO",
	}

	TruckAB = Truck{
		lastTransporter: TranporterA,
		currentLocation: "B",
		vehicleType:     "CDD",
	}

	TruckBA = Truck{
		lastTransporter: TranporterB,
		currentLocation: "A",
		vehicleType:     "FUSO",
	}

	TruckBA2 = Truck{
		lastTransporter: TranporterB,
		currentLocation: "A",
		vehicleType:     "CDD",
	}

	RoutePref = []RoutePreference{
		{transporter: TranporterA, originLocation: "A", destinationLocation: "B", vehicleType: "FUSO"},
		{transporter: TranporterA, originLocation: "B", destinationLocation: "C", vehicleType: "CDD"},
		{transporter: TranporterB, originLocation: "A", destinationLocation: "B", vehicleType: "FUSO"},
		{transporter: TranporterB, originLocation: "A", destinationLocation: "B", vehicleType: "CDD"},
		{transporter: TranporterB, originLocation: "A", destinationLocation: "C", vehicleType: "FUSO"},
	}
}
