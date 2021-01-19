package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
	"math/rand"
)

func TaskItem(name string, description string) *BubbleContainer {
	var (
		colors = []string{"#27ACB2", "#FF6B6E", "#A17DF5"}
		color = colors[rand.Intn(len(colors))]
		textFlex = 1
	)

	// View Components
	return MiniBubbleCard(color, []FlexComponent{
		&TextComponent{
			Type:       FlexComponentTypeText,
			Text:       name,
			Margin:     FlexComponentMarginTypeLg,
			Size:       FlexTextSizeTypeMd,
			Align:      FlexComponentAlignTypeStart,
			Gravity:    FlexComponentGravityTypeCenter,
			Wrap:       false,
			Color:      "#ffffff",
		},
	}, []FlexComponent{
		&BoxComponent{
			Type:            FlexComponentTypeBox,
			Layout:          FlexBoxLayoutTypeHorizontal,
			Contents:        []FlexComponent{
				&TextComponent{
					Type:       FlexComponentTypeText,
					Text:       description,
					Margin:     FlexComponentMarginTypeLg,
					Size:       FlexTextSizeTypeXs,
					Align:      FlexComponentAlignTypeStart,
					Gravity:    FlexComponentGravityTypeCenter,
					Wrap:       false,
					Color:      "#8C8C8C",
				},
			},
			Flex:            &textFlex,
		},
	})
}
