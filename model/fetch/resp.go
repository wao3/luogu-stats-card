package fetch

type DataType interface {
	UserProfileData
}

type Resp[data DataType] struct {
	CurrentData data `json:"currentData"`
}
