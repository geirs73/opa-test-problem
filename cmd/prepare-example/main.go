package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	outputPath = "rego/carsdata/carsdata.json"
	startIndex = 2000
	minYear    = 2000
	maxYear    = 2040
	minLength  = 3300
	maxLength  = 5500
)

type car struct {
	Make      string `json:"make"`
	Model     string `json:"model"`
	Fuel      string `json:"fuel"`
	ModelYear int    `json:"modelyear"`
	Length    int    `json:"length"`
}

var fuels = []string{"petrol", "diesel", "electric", "hybrid"}

var models = []car{
	{Make: "Toyota", Model: "Camry"}, {Make: "Toyota", Model: "Corolla"},
	{Make: "Toyota", Model: "RAV4"}, {Make: "Toyota", Model: "Yaris"},
	{Make: "Honda", Model: "Civic"}, {Make: "Honda", Model: "CR-V"},
	{Make: "Honda", Model: "Accord"}, {Make: "Honda", Model: "Jazz"},
	{Make: "Ford", Model: "Focus"}, {Make: "Ford", Model: "Fiesta"},
	{Make: "Ford", Model: "Mustang"}, {Make: "Ford", Model: "Kuga"},
	{Make: "Renault", Model: "Clio"}, {Make: "Renault", Model: "Captur"},
	{Make: "Renault", Model: "Scenic"}, {Make: "Renault", Model: "Megane"},
	{Make: "Nissan", Model: "Micra"}, {Make: "Nissan", Model: "Qashqai"},
	{Make: "Nissan", Model: "Ariya"}, {Make: "Nissan", Model: "Leaf"},
	{Make: "Volkswagen", Model: "Golf"}, {Make: "Volkswagen", Model: "Polo"},
	{Make: "Volkswagen", Model: "Passat"}, {Make: "Volkswagen", Model: "Tiguan"},
	{Make: "BMW", Model: "3 Series"}, {Make: "BMW", Model: "5 Series"},
	{Make: "BMW", Model: "X3"}, {Make: "BMW", Model: "i4"},
	{Make: "Mercedes-Benz", Model: "A-Class"}, {Make: "Mercedes-Benz", Model: "C-Class"},
	{Make: "Mercedes-Benz", Model: "E-Class"}, {Make: "Mercedes-Benz", Model: "GLC"},
	{Make: "Audi", Model: "A3"}, {Make: "Audi", Model: "A4"},
	{Make: "Audi", Model: "Q5"}, {Make: "Audi", Model: "e-tron"},
	{Make: "Volvo", Model: "XC40"}, {Make: "Volvo", Model: "XC60"},
	{Make: "Volvo", Model: "V60"}, {Make: "Volvo", Model: "EX30"},
	{Make: "Hyundai", Model: "i20"}, {Make: "Hyundai", Model: "i30"},
	{Make: "Hyundai", Model: "Tucson"}, {Make: "Hyundai", Model: "Ioniq 5"},
	{Make: "Kia", Model: "Ceed"}, {Make: "Kia", Model: "Sportage"},
	{Make: "Kia", Model: "Niro"}, {Make: "Kia", Model: "EV6"},
	{Make: "Skoda", Model: "Octavia"}, {Make: "Skoda", Model: "Fabia"},
	{Make: "Skoda", Model: "Kodiaq"}, {Make: "Skoda", Model: "Enyaq"},
	{Make: "Peugeot", Model: "208"}, {Make: "Peugeot", Model: "308"},
	{Make: "Peugeot", Model: "3008"}, {Make: "Peugeot", Model: "508"},
	{Make: "Mazda", Model: "3"}, {Make: "Mazda", Model: "CX-5"},
	{Make: "Mazda", Model: "MX-5"}, {Make: "Mazda", Model: "CX-30"},
	{Make: "Subaru", Model: "Outback"}, {Make: "Subaru", Model: "Forester"},
	{Make: "Subaru", Model: "Impreza"}, {Make: "Subaru", Model: "Solterra"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <count>\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Fprintf(os.Stderr, "error: count must be a positive integer, got %q\n", os.Args[1])
		os.Exit(2)
	}
	maxUnique := len(models) * len(fuels) * (maxYear - minYear + 1) * (maxLength - minLength + 1)
	if n > maxUnique {
		fmt.Fprintf(os.Stderr, "error: at most %d unique make/model/fuel/modelyear/length combinations are available, got %d\n", maxUnique, n)
		os.Exit(2)
	}

	if err := writeData(n); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %d records to %s\n", n, outputPath)
}

// writeData emits one record per line, keyed from startIndex, in index order.
func writeData(n int) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, "{")
	numYears := maxYear - minYear + 1
	numLengths := maxLength - minLength + 1
	for i := 0; i < n; i++ {
		// model/fuel/year cycle with period 10496; length (stride 137, coprime with
		// 2201 lengths) only repeats together with them after 10496*2201 records,
		// so the full tuple is unique while i < maxUnique
		c := models[i%len(models)]
		c.Fuel = fuels[(i/len(models))%len(fuels)]
		c.ModelYear = minYear + (i/(len(models)*len(fuels)))%numYears
		c.Length = minLength + (i*137)%numLengths
		rec, err := json.Marshal(c)
		if err != nil {
			return err
		}
		comma := ","
		if i == n-1 {
			comma = ""
		}
		fmt.Fprintf(f, "  %q: %s%s\n", strconv.Itoa(startIndex+i), rec, comma)
	}
	fmt.Fprintln(f, "}")
	return f.Close()
}
