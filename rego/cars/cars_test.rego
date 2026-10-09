package cars_test

import data.carsdata as hugedata

cardata := {
	"1000": {"make": "Toyota", "model": "Camry", "fuel": "petrol", "modelyear": 2015, "length": 4850},
	"1001": {"make": "Honda", "model": "Civic", "fuel": "petrol", "modelyear": 2018, "length": 4550},
	"1002": {"make": "Ford", "model": "Focus", "fuel": "diesel", "modelyear": 2012, "length": 4360},
	"1003": {"make": "Honda", "model": "CR-V", "fuel": "petrol", "modelyear": 2020, "length": 4600},
	"1004": {"make": "Renault", "model": "Clio", "fuel": "petrol", "modelyear": 2008, "length": 3990},
	"1005": {"make": "Nissan", "model": "Micra", "fuel": "diesel", "modelyear": 2005, "length": 3715},
	"1006": {"make": "Nissan", "model": "Qashqai", "fuel": "petrol", "modelyear": 2019, "length": 4394},
	"1007": {"make": "Renault", "model": "Captur", "fuel": "diesel", "modelyear": 2016, "length": 4122},
	"1008": {"make": "Renault", "model": "Scenic", "fuel": "electric", "modelyear": 2023, "length": 4470},
	"1009": {"make": "Nissan", "model": "Ariya", "fuel": "electric", "modelyear": 2022, "length": 4595},
	"1010": {"make": "Nissan", "model": "Leaf", "fuel": "electric", "modelyear": 2021, "length": 4490},
}

test_func_foobar if {
	# this line should show line with inlay information on arguments:
	# keys := object.keys(object: cardata)
	keys := object.keys(cardata)
	"1010" in keys
	not "1011" in keys
}

test_func_foobar2 if {
	# this line should show line with inlay information on arguments:
	# keys := object.keys(object: cardata)
	keys := object.keys(hugedata)
	"2010" in keys
	# not "2011" in keys
}
