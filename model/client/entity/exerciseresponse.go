package client

type Exercise struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	BodyPart         string   `json:"bodyPart"`
	Equipment        string   `json:"equipment"`
	Target           string   `json:"target"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
	GifUrl           string   `json:"gifUrl"`
}
