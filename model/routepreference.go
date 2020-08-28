package model

type RoutePreference struct {
	transporter         Transporter
	originLocation      string
	destinationLocation string
	vehicleType         string
}
