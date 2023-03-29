package fetch

type DataType interface {
	string | UserProfileData
}

type Resp[data DataType] struct {
	CurrentData data `json:"currentData"`
}
