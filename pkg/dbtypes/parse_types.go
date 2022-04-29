package dbtypes

type TableCol struct {
	Name     string
	Type     string
	JsonName string
	Optional bool
	Comment  string
}

type TableDef struct {
	Name     string
	JsonName string
	Columns  []TableCol
}

type DbDef struct {
	Name   string
	Tables []TableDef
}

type OVNDbType interface {
	IsValid() bool
}
