package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Basic struct with JSON tags
type Bird struct {
	Species     string     `json:"birdType"`
	Description string     `json:"what it does,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"` // Pointer for nullable time
	Dimensions  Dimensions `json:"dimensions"`
	IsActive    bool       `json:"-"` // Field that will be ignored in JSON
}

// Custom type implementing Marshaler and Unmarshaler interfaces
type Dimensions struct {
	Height int
	Width  int
}

// Custom JSON unmarshaling
func (d *Dimensions) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}
	s := string(data)[1 : len(data)-1]
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}

	fmt.Sscanf(parts[0], "%d", &d.Height)
	fmt.Sscanf(parts[1], "%d", &d.Width)
	return nil
}

// Custom JSON marshaling
func (d Dimensions) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%dx%d"`, d.Height, d.Width)), nil
}

func main() {
	fmt.Println("1. Basic Unmarshaling:")
	// Basic unmarshaling
	basicJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(basicJson), &bird)
	fmt.Printf("Basic unmarshaling result: %+v\n\n", bird)

	fmt.Println("2. Array Unmarshaling:")
	// Array unmarshaling
	arrayJson := `[
		{"birdType": "pigeon","what it does": "likes to perch on rocks"},
		{"birdType": "eagle","what it does": "bird of prey"}
	]`
	var birds []Bird
	json.Unmarshal([]byte(arrayJson), &birds)
	fmt.Printf("Array unmarshaling result: %+v\n\n", birds)

	fmt.Println("3. Primitive Types:")
	// Primitive types
	numberJson := "42"
	floatJson := "3.14"
	stringJson := `"bird"`
	var n int
	var pi float64
	var str string
	json.Unmarshal([]byte(numberJson), &n)
	json.Unmarshal([]byte(floatJson), &pi)
	json.Unmarshal([]byte(stringJson), &str)
	fmt.Printf("Primitives: number=%d, float=%f, string=%s\n\n", n, pi, str)

	fmt.Println("4. Time Values:")
	// Time values
	timeJson := `{"birdType": "pigeon", "createdAt": "2023-10-26T12:00:00Z"}`
	var birdWithTime Bird
	json.Unmarshal([]byte(timeJson), &birdWithTime)
	fmt.Printf("Time unmarshaling: %+v\n\n", birdWithTime)

	fmt.Println("5. Custom Parsing (Dimensions):")
	// Custom parsing
	dimensionsJson := `{"birdType": "pigeon", "dimensions": "20x30"}`
	var birdWithDimensions Bird
	json.Unmarshal([]byte(dimensionsJson), &birdWithDimensions)
	fmt.Printf("Custom dimensions parsing: %+v\n\n", birdWithDimensions)

	fmt.Println("6. Unstructured Data (Maps):")
	// Unstructured data using maps
	unstructuredJson := `{
		"birds": {
			"pigeon": "likes to perch on rocks",
			"eagle": "bird of prey"
		},
		"animals": "none"
	}`
	var result map[string]any
	json.Unmarshal([]byte(unstructuredJson), &result)
	fmt.Printf("Unstructured data: %+v\n\n", result)

	fmt.Println("7. JSON Validation:")
	// JSON validation
	invalidJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
	if !json.Valid([]byte(invalidJson)) {
		fmt.Printf("Invalid JSON detected: %s\n\n", invalidJson)
	}

	fmt.Println("8. Marshaling:")
	// Marshaling struct to JSON
	newBird := Bird{
		Species:     "Falcon",
		Description: "Fast flyer",
		CreatedAt:   time.Now(),
		Dimensions:  Dimensions{Height: 30, Width: 40},
	}
	data, _ := json.Marshal(newBird)
	fmt.Printf("Marshaled struct: %s\n\n", string(data))

	fmt.Println("9. Pretty Printing:")
	// Pretty printing
	prettyData, _ := json.MarshalIndent(newBird, "", "  ")
	fmt.Printf("Pretty printed JSON:\n%s\n\n", string(prettyData))

	fmt.Println("10. Null Values:")
	// Null values
	nullBird := Bird{
		Species:     "Sparrow",
		Description: "", // This will be omitted due to omitempty
		UpdatedAt:   nil, // This will be encoded as null
	}
	nullData, _ := json.Marshal(nullBird)
	fmt.Printf("JSON with null values: %s\n\n", string(nullData))

	fmt.Println("11. Map Marshaling:")
	// Map marshaling
	birdMap := map[string]any{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squawk",
		},
		"total birds": 2,
		"nullValue":   nil,
	}
	mapData, _ := json.Marshal(birdMap)
	fmt.Printf("Marshaled map: %s\n", string(mapData))
}
