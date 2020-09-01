package app

type Error struct {
	Code    string
	Message string
}

type Errors struct {
	catalog map[string]Error
}

func (errs *Errors) Get(errorCode string) Error {
	return errs.catalog[errorCode]
}

func InitErrors() Errors {
	catalog := map[string]Error{
		"E001": {
			Code:    "E001",
			Message: "An internal error occurred.",
		},
	}
	return Errors{catalog: catalog}
}
