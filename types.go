package runvert

type Score int

const (
	MatchNothing Score = 0
	MatchAll     Score = 10000
)

func CalcScore(validCount, totalCount int) Score {
	if totalCount <= 0 || validCount <= 0 {
		return MatchNothing
	} else if totalCount >= validCount {
		return MatchAll
	} else {
		return validCount * MatchAll / totalCount
	}
}

type Converter interface {
	Convert(data []byte) []byte
}

type Coder interface {
	Name() string
	Encode(data []byte) []byte
	Decode(data []byte) []byte

	Test(data []byte) Score
}

type CoderSet struct {
}

type Language interface {
	Name() string
	Aliases() []string
	Coders() []string

	TestCommon(text []rune) Score
}

type LanguageSet struct {
}

type CharsetConverter interface {
	Converter
	Coder

	Charset() string
}

type EncodeConverter interface {
	Converter
	Coder

	Coder() string
}

type encoder struct {
	coder Coder
}

func MakeEncoder(coder Coder) Converter {
	return &encoder{coder}
}

func (e *encoder) Convert(data []byte) []byte {
	return e.coder.Encode(data)
}

type decoder struct {
	coder Coder
}

func MakeDecoder(coder Coder) Converter {
	return &decoder{coder}
}

func (e *decoder) Convert(data []byte) []byte {
	return e.coder.Decode(data)
}

