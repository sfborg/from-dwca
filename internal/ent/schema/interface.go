package schema

type SchemaSQL interface {
	GetSchema(tag string) error
}
