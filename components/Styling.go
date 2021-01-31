package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
)

type BoxStyle struct {
	flex *int
	spacing FlexComponentSpacingType
	margin FlexComponentMarginType
	color string
}
