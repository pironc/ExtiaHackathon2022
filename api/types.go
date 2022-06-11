package main

type City struct {
	Name string `json:"city"`

	Rent float32 `json:"rent"`
	// insert data types that you wanna add in the correct format:
	// VARNAME TYPE `json:"NAME IN JSON"`
}
