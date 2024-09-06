// package schema contains mappings to SQLite schema of the
// SFGarchive.
package schema

// DataSource data in SQLite schema for SFGarchive.
type DataSource struct {
	ID           string
	GnID         int
	LocalID      string
	GlobalID     string
	Title        string
	TitleShort   string
	Version      string
	RevisionDate string
	DOI          string
	Citation     string
	Authors      string
	Description  string
	WebsiteURL   string
	DataURL      string
	RecordCount  int
	UpdatedAt    int
}

// Core contains the bulk of a record data.
type Core struct {
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

// Vern provides data for a vernacular name record.
type Vern struct {
	TaxonID        string
	VernacularID   string
	VernacularName string
	Language       string
	LangCode       string
	LangInEnglish  string
	Locality       string
	CountryCode    string
}
