package domain

type ExistsParams struct {
	Table            string
	IdColumnName     string
	IdValue          interface{}
	StatusColumnName *string
	StatusValue      *int
}
