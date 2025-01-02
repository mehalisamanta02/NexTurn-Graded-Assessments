package main

import (
	"errors"
	"fmt"
	"strings"
)

type CityData struct {
	Name     string
	AvgTemp  float64
	Rainfall float64
}

var cityData = []CityData{
	{"Pune", 24, 2213.4},
	{"Mumbai", 27.7, 2502.3},
	{"Kolkata", 26.8, 2423.6},
	{"Delhi", 25, 797.3},
	{"Bangalore", 29, 800},
}

// Find city with highest and lowest temperature
func findTemperatureExtremes(data []CityData) (CityData, CityData) {
	highest := data[0]
	lowest := data[0]
	for _, city := range data {
		if city.AvgTemp > highest.AvgTemp {
			highest = city
		}
		if city.AvgTemp < lowest.AvgTemp {
			lowest = city
		}
	}
	return highest, lowest
}

// Calculate average rainfall
func calculateAverageRainfall(data []CityData) float64 {
	totalRainfall := 0.0
	for _, city := range data {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(data))
}

// Filter cities by rainfall threshold
func filterCitiesByRainfall(data []CityData, threshold float64) []CityData {
	filtered := []CityData{}
	for _, city := range data {
		if city.Rainfall > threshold {
			filtered = append(filtered, city)
		}
	}
	return filtered
}

// Search for a city by name
func searchCityByName(data []CityData, name string) (CityData, error) {
	for _, city := range data {
		if strings.ToLower(city.Name) == strings.ToLower(name) {
			return city, nil
		}
	}
	return CityData{}, errors.New("city not found")
}

func main() {
	fmt.Println("Climate Data Analysis")

	// Find extremes
	highest, lowest := findTemperatureExtremes(cityData)
	fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highest.Name, highest.AvgTemp)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.AvgTemp)

	// Calculate average rainfall
	avgRainfall := calculateAverageRainfall(cityData)
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", avgRainfall)

	// Filter cities by rainfall
	var threshold float64
	fmt.Print("Enter rainfall threshold (mm): ")
	fmt.Scanln(&threshold)
	filteredCities := filterCitiesByRainfall(cityData, threshold)
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	for _, city := range filteredCities {
		fmt.Printf("- %s (%.2f mm)\n", city.Name, city.Rainfall)
	}

	// Search for a city
	var cityName string
	fmt.Print("Enter city name to search: ")
	fmt.Scanln(&cityName)
	city, err := searchCityByName(cityData, cityName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("City found: %s (Temperature: %.2f°C, Rainfall: %.2f mm)\n", city.Name, city.AvgTemp, city.Rainfall)
	}
}
