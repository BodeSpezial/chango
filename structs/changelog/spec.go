package changelog

type Added struct {
	Message string
	Pr      int
}

type Removed struct {
	Message string
	Pr      int
}

type Changed struct {
	Message string
	Pr      int
}

type Fixed struct {
	Message string
	Pr      int
}
