package stor

type Storage interface {
	Init() error
	Close() error
}
