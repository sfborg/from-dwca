package ds

type DataSource struct {
	ID           string
	GNID         int
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
