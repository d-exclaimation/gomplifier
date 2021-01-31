package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
)

func MiniBubbleCard(hexColor string, head []FlexComponent, body []FlexComponent) *BubbleContainer {
	return &BubbleContainer{
		Type:      FlexContainerTypeBubble,
		Size:      FlexBubbleSizeTypeNano,
		Header:    &BoxComponent{
			Type:            FlexComponentTypeBox,
			Layout:          FlexBoxLayoutTypeVertical,
			Contents:        head,
			BackgroundColor: hexColor,
		},
		Body:      &BoxComponent{
			Type:            FlexComponentTypeBox,
			Layout:          FlexBoxLayoutTypeVertical,
			Contents:        body,
			Spacing:         FlexComponentSpacingTypeMd,
		},
	}
}
