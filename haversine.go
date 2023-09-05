package haversine

import (
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func radiansToDegrees(r float64) float64 {
	return r * 180 / math.Pi
}

func Bearing(position, target Coord) (bearing float64) {
	//dLon := radiansToDegrees(target.Lon - position.Lon)
	//y := math.Sin(dLon) * math.Cos(radiansToDegrees(target.Lat))
	//x := math.Cos(radiansToDegrees(position.Lat))*math.Sin(radiansToDegrees(target.Lat)) - math.Sin(radiansToDegrees(position.Lat))*math.Cos(dLon)
	//r := math.Atan2(y, x)
	//bearing = math.Mod(radiansToDegrees(r)+360, 360)

	dLon := (target.Lon - position.Lon) * math.Pi / 180
	lat1 := position.Lat * math.Pi / 180
	lat2 := target.Lat * math.Pi / 180

	y := math.Sin(dLon) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)
	bearing = math.Atan2(y, x) * 180 / math.Pi

	return math.Mod(bearing+360, 360) // in degrees
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q Coord) (mi, km float64) {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi = c * earthRadiusMi
	km = c * earthRaidusKm

	return mi, km
}
