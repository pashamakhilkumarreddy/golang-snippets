package models

type Spec struct { // exported struct
	Maker string // Field names should also be in capital-case to export
	Price int    // exported field
	// model string unexported field
}

type Name struct {
	FirstName, LastName string
}
