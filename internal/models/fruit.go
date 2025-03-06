package models

type Fruit struct {
	Name       string `json:"name"`
	Family     string `json:"family"`
	Order      string `json:"order"`
	Genus      string `json:"genus"`
	Nutritions struct {
		Calories      int64   `json:"calories"`
		Fat           float64 `json:"fat"`
		Sugar         float64 `json:"sugar"`
		Carbohydrates float64 `json:"carbohydrates"`
		Protein       float64 `json:"protein"`
	} `json:"nutritions"`
}
