package common

func FastFail(err error) {
	if err != nil {
		panic(err)
	}
}

func GetContainerWidth(cardWidth int) int {
	var paddingRight = 100
	return cardWidth - paddingRight
}
