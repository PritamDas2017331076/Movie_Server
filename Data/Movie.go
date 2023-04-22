package Data

type Movie struct {
	Name      string `json: "name"`
	Id        int    `json:"id"`
	Directors []int  `json:"directors"`
	Genres    []int  `json:"genres"`
}
