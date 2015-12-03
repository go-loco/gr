package gr

//GeoAdd Adds the specified geospatial items (latitude, longitude, name) to the specified key
func rGeoAdd(key string, fields ...string) ([][]byte, error) {
	return multiCompile(append([]string{"GEOADD", key}, fields...)...), nil
}

//GeoDist Return the distance between two members in the geospatial index represented by the sorted set.
func rGeoDist(key string, member1 string, member2 string, unit string) ([][]byte, error) {

	//Chek if we have to send the units
	if unit == "" {
		return multiCompile("GEODIST", key, member1, member2), nil
	}
	return multiCompile("GEODIST", key, member1, member2, unit), nil
}

//GeoHash return valid Geohash strings representing the position of one or more elements
//in a sorted set value representing a geospatial index (where elements were added using GEOADD).
func rGeoHash(key string, fields ...string) ([][]byte, error) {
	return multiCompile(append([]string{"GEOHASH", key}, fields...)...), nil
}

//GeoPos return the positions (longitude,latitude) of all the specified members
//of the geospatial index represented by the sorted set at key.
func rGeoPos(key string, fields ...string) ([][]byte, error) {
	return multiCompile(append([]string{"GEOPOS", key}, fields...)...), nil
}
