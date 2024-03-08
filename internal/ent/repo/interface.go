package repo

type Repo interface {
	FetchSchema() ([]byte, error)
}
