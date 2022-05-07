package entities

const (
	GlobalAlignment = "Global"
	LocalAlignment  = "Local"
)

type AlignConfig struct {
	Seq1 string
	Seq2 string
	Mode string
	Scheme
}

type Scheme struct {
	Match    int
	Mismatch int
	Gap      int
}
