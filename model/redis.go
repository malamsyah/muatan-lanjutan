package model

import "fmt"

type RouteData struct {
	Key string
	Val []Transporter
}

type TruckData struct {
	Key string
	Val []Transporter
}

type Redis struct {
	RouteData []RouteData
	TruckData []TruckData
}

func (r *Redis) InsertRouteData(preference RoutePreference) {

	newKey := fmt.Sprint(preference.originLocation, "-", preference.destinationLocation, "-", preference.vehicleType)

	for i, v := range r.RouteData {
		if v.Key == newKey {
			r.RouteData[i].Val = append(r.RouteData[i].Val, preference.transporter)
			return
		}
	}

	r.RouteData = append(r.RouteData, RouteData{
		Key: newKey,
		Val: []Transporter{preference.transporter},
	})
}

func (r *Redis) GetRouteTransporter(key string) []Transporter {
	for _, v := range r.RouteData {
		if v.Key == key {
			return v.Val
		}
	}

	return nil
}

func (r *Redis) InsertTruckData(truck Truck) {

	newKey := fmt.Sprint(truck.currentLocation, "-", truck.vehicleType)

	for i, v := range r.TruckData {
		if v.Key == newKey {
			r.RouteData[i].Val = append(r.TruckData[i].Val, truck.lastTransporter)
			return
		}
	}

	r.TruckData = append(r.TruckData, TruckData{
		Key: newKey,
		Val: []Transporter{truck.lastTransporter},
	})
}

func (r *Redis) GetTruckTransporter(key string) []Transporter {
	for _, v := range r.TruckData {
		if v.Key == key {
			return v.Val
		}
	}

	return nil
}
