package gr_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestGeoBegin(t *testing.T) {
	log.Println("[Testing Geo]")
}

func TestGeoAdd(t *testing.T) {

	test := func() {

		//Add the locations to redis
		r, err := redis.GeoAdd("Sicily", "13.361389", "38.115556", "Palermo", "15.087269", "37.502669", "Catania")
		if err != nil || r != 2 {
			fmt.Println("Error Adding Locations to Redis")
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestGeoDist(t *testing.T) {

	test := func() {

		//Add the locations to redis
		r, err := redis.GeoAdd("Sicily", "13.361389", "38.115556", "Palermo", "15.087269", "37.502669", "Catania")
		if err != nil || r != 2 {
			fmt.Println("Error Adding Locations to Redis")
			t.Fail()
		}

		//Calculate the distance between 2 cities
		dist, err2 := redis.GeoDist("Sicily", "Palermo", "Catania", "")

		//Chek is the distance is correct
		if err2 != nil || !strings.Contains(dist, "166274.15") {
			fmt.Println("Error Calculating Distance", "Dist:", dist)
			fmt.Println(err2)
			t.Fail()
		}

		//Calculate the distance between 2 cities in Km
		dist, err2 = redis.GeoDist("Sicily", "Palermo", "Catania", "km")

		//Chek is the distance is correct
		if err2 != nil || !strings.Contains(dist, "166.274") {
			fmt.Println("Error Calculating Distance", "Dist:", dist)
			fmt.Println(err2)
			t.Fail()
		}

		//Calculate the distance between 2 cities in Miles
		dist, err2 = redis.GeoDist("Sicily", "Palermo", "Catania", "mi")

		//Chek is the distance is correct
		if err2 != nil || !strings.Contains(dist, "103.318") {
			fmt.Println("Error Calculating Distance", "Dist:", dist)
			fmt.Println(err2)
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestGeoHash(t *testing.T) {

	test := func() {

		//Add the locations to redis
		r, err := redis.GeoAdd("Sicily", "13.361389", "38.115556", "Palermo", "15.087269", "37.502669", "Catania")
		if err != nil || r != 2 {
			fmt.Println("Error Adding Locations to Redis")
			t.Fail()
		}

		//Chek the hash of the cities
		hashes, err2 := redis.GeoHash("Sicily", "Palermo", "Catania")
		if err2 != nil {
			fmt.Println("Error Getting hashes")
			t.Fail()
		}

		//Check that only 2 elements were returned
		if len(hashes) != 2 {
			fmt.Println("Invalid quantity of elements:", len(hashes))
			t.Fail()
		}

		//Check the hash of the first city
		if !(hashes[0] == "sqc8b49rny0" || hashes[1] == "sqc8b49rny0") {
			fmt.Println("Invalid hash for first city", hashes)
			t.Fail()
		}

		//Check the hash of the second city
		if !(hashes[0] == "sqdtr74hyu0" || hashes[1] == "sqdtr74hyu0") {
			fmt.Println("Invalid hash for second city", hashes)
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestGeoPos(t *testing.T) {

	test := func() {

		//Add the locations to redis
		r, err := redis.GeoAdd("Sicily", "13.361389", "38.115556", "Palermo", "15.087269", "37.502669", "Catania")
		if err != nil || r != 2 {
			fmt.Println("Error Adding Locations to Redis")
			t.Fail()
		}

		//Chek the position of the cities
		positions, err2 := redis.GeoPos("Sicily", "Palermo", "Catania")
		if err2 != nil {
			fmt.Println("Error Getting Positions")
			t.Fail()
		}

		//Check that only 5 elements were returned
		if len(positions) != 5 {
			fmt.Println("Invalid quantity of elements:", len(positions))
			t.Fail()
		}

		//Check lat and long of the first city
		if !(strings.Contains(positions[0], "13.3613")) || !(strings.Contains(positions[1], "38.1155")) {
			fmt.Println("Invalid positions for first city", positions)
			t.Fail()
		}

		//Check the hash of the second city
		if !(strings.Contains(positions[2], "15.0872")) || !(strings.Contains(positions[3], "37.5026")) {
			fmt.Println("Invalid positions for second city", positions)
			t.Fail()
		}

		//Check the nil position of the third city
		if positions[4] != "" {
			fmt.Println("Invalid nil city", positions)
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestGeoEnd(t *testing.T) {
	println("[OK]")
}
