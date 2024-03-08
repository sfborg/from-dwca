package core

type Data struct {
	RecordID   string
	NameID     string
	Name       string
	Authorship string
	Year       int

	Cardinality     int
	CanonicalID     string
	Canonical       string
	CanonicalFullID string
	CanonicalFull   string

	CanonicalStemID     string
	CanonicalStem       string
	AcceptedNameUsageID string
	Classification      string
	ClassificationIDs   string
	ClassificationRanks string

	Rank             string
	Virus            bool
	Bacteria         bool
	Surrogate        bool
	NomeclaturalCode string

	ParseQuality int
}
