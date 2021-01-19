package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
)

func AnnounceCard(msg string, author string, from bool) *BubbleContainer {
	var (
		color   = "#FC0D85"
		header  = "to Somewhere"
		msgFlex = 0
		size    = FlexBubbleSizeTypeNano
	)
	if from {
		color = "#1DB446"
		header = "from Somewhere"
		size = FlexBubbleSizeTypeMicro
	}

	return &BubbleContainer{
		Type: FlexContainerTypeBubble,
		Size: size,
		Body: &BoxComponent{
			Type:   FlexComponentTypeBox,
			Layout: FlexBoxLayoutTypeVertical,
			Contents: []FlexComponent{
				&TextComponent{
					Type:   FlexComponentTypeText,
					Text:   "Announcer",
					Size:   FlexTextSizeTypeXxs,
					Wrap:   false,
					Weight: FlexTextWeightTypeBold,
					Color:  color,
				},
				&TextComponent{
					Type:  FlexComponentTypeText,
					Text:  header,
					Size:  FlexTextSizeTypeXxs,
					Wrap:  true,
					Color: "#aaaaaa",
				},
				&SeparatorComponent{
					Type:   FlexComponentTypeSeparator,
					Margin: FlexComponentMarginTypeSm,
				},
				&BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeVertical,
					Contents: []FlexComponent{
						&TextComponent{
							Type:  FlexComponentTypeText,
							Text:  msg,
							Flex:  &msgFlex,
							Size:  FlexTextSizeTypeXs,
							Wrap:  true,
							Color: "#555555",
						},
					},
					Spacing: FlexComponentSpacingTypeXs,
					Margin:  FlexComponentMarginTypeLg,
				},
				&SeparatorComponent{
					Type:   FlexComponentTypeSeparator,
					Margin: FlexComponentMarginTypeXl,
				},
				&BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeHorizontal,
					Contents: []FlexComponent{
						&TextComponent{
							Type:  FlexComponentTypeText,
							Text:  "Author",
							Flex:  &msgFlex,
							Size:  FlexTextSizeTypeXxs,
							Color: "#aaaaaa",
						},
						&TextComponent{
							Type:  FlexComponentTypeText,
							Text:  author,
							Align: FlexComponentAlignTypeEnd,
							Size:  FlexTextSizeTypeXxs,
							Color: "#aaaaaa",
						},
					},
					Margin: FlexComponentMarginTypeXs,
				},
			},
		},
		Footer: nil,
	}
}
