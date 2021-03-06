package funcs

type Cell struct {
	Id   int
	Name string

	Alive bool

	Size  int
	Kills int
	Meals []Cell

	Pos []int
}

type Food struct {
	Id    int   `json:"id"`
	Pos   []int `json:"pos"`
	Value int   `json:"value"`
}
