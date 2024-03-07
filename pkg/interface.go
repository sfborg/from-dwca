package fdwca

type FromDwCA interface {
	GetDwCA(fileDwCA string) error
	ExportData() error
	DumpData() error
}
