package utilities

import (
	"math"

	"tours.xws.com/model"
)

const earthRadiusInKm = 6371.0

func CalculateDistance(currentPointLatitude, currentPointLongitude float64, currentTouristPosition model.Position) float64 {

	currentTouristLatInRad := currentTouristPosition.Latitude * math.Pi / 180
	currentPointLatInRad := currentPointLatitude * math.Pi / 180

	deltaLatInRad := math.Abs(currentPointLatitude-currentTouristPosition.Latitude) * math.Pi / 180
	deltaLongInRad := math.Abs(currentPointLongitude-currentTouristPosition.Longitude) * math.Pi / 180

	a := math.Pow(math.Sin(deltaLatInRad/2), 2) +
		math.Pow(math.Sin(deltaLongInRad/2), 2)*
			math.Cos(currentTouristLatInRad)*math.Cos(currentPointLatInRad)
	return 2 * earthRadiusInKm * math.Asin(math.Sqrt(a))
}
