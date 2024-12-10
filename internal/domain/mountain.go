package domain

type Mountain struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	ImagePath        string  `json:"image_path"`
	ShortDescription string  `json:"short_description"`
	LongDescription  string  `json:"long_description"`
	Elevation        int     `json:"elevation"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Region           string  `json:"region"`
	Prefectures      string  `json:"prefectures"`
}
