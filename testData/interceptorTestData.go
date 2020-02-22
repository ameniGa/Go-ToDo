package testData

// TTReqValidation represents the table test for interceptor
var TTReqValidation = []struct {
	Name       string
	Hash       string
	Title      string
	ReqType    string
	HasError   bool
	IsValidCtx bool
}{
	{
		Name:       "invalid context",
		Title:      "golang",
		ReqType:    "add",
		HasError:   false,
		IsValidCtx: false,
	},
	{
		Name:       "valid add request",
		Hash:       "",
		Title:      "golang",
		ReqType:    "add",
		HasError:   false,
		IsValidCtx: true,
	},
	{
		Name:       "empty Title in Add request",
		Title:      "",
		ReqType:    "add",
		HasError:   true,
		IsValidCtx: true,
	},
	{
		Name:       "empty Hash in Delete request",
		Hash:       "",
		ReqType:    "delete",
		HasError:   true,
		IsValidCtx: true,
	},
	{
		Name:       "empty Hash in Update request",
		Hash:       "",
		ReqType:    "update",
		HasError:   true,
		IsValidCtx: true,
	},
	{
		Name:       "unknown request",
		ReqType:    "unkown",
		HasError:   true,
		IsValidCtx: true,
	},
	{
		Name:       "nil request",
		ReqType:    "nil",
		HasError:   true,
		IsValidCtx: true,
	},
}
