package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
)

func GameView(state [][]string, hexColor string, footer string, isWin bool) *BubbleContainer {
	var (
		header = "Burn the snakes"
		stateRows = rows(state)
		msgFlex = 0
	)
	if isWin {
		header = "Congrats"
	}
	return &BubbleContainer{
		Type:      FlexContainerTypeBubble,
		Size:      FlexBubbleSizeTypeKilo,
		Body:      &BoxComponent{
			Type:            FlexComponentTypeBox,
			Layout:          FlexBoxLayoutTypeVertical,
			Contents:        []FlexComponent{
				&TextComponent{
					Type:       FlexComponentTypeText,
					Text:       "Sokoban",
					Size:       FlexTextSizeTypeSm,
					Wrap:       false,
					Weight:     FlexTextWeightTypeBold,
					Color:      hexColor,
				},
				&TextComponent{
					Type:       FlexComponentTypeText,
					Text:       header,
					Size:       FlexTextSizeTypeXs,
					Wrap:       true,
					Color:      "#aaaaaa",
				},
				&SeparatorComponent{
					Type:   FlexComponentTypeSeparator,
					Margin: FlexComponentMarginTypeSm,
				},
				&BoxComponent{
					Type:            FlexComponentTypeBox,
					Layout:          FlexBoxLayoutTypeVertical,
					Contents:        ForEachString(stateRows, func(item string) FlexComponent {
						return &TextComponent{
							Type:       FlexComponentTypeText,
							Text:       item,
							Flex:       &msgFlex,
							Size:       FlexTextSizeTypeXs,
							Wrap: 		true,
						}
					}),
					Spacing:         FlexComponentSpacingTypeXs,
					Margin:          FlexComponentMarginTypeLg,
				},
				&SeparatorComponent{
					Type:   FlexComponentTypeSeparator,
					Margin: FlexComponentMarginTypeXl,
				},
				&BoxComponent{
					Type:            FlexComponentTypeBox,
					Layout:          FlexBoxLayoutTypeHorizontal,
					Contents:        []FlexComponent{
						&TextComponent{
							Type:       FlexComponentTypeText,
							Text:       footer,
							Flex:       &msgFlex,
							Size:       FlexTextSizeTypeXs,
							Color:      "#aaaaaa",
						},
					},
					Margin:          FlexComponentMarginTypeXs,
				},
			},
		},
	}
}

func rows(content [][]string) []string {
	var res []string
	for i := 0; i < len(content); i++ {
		var row = ""
		for j := 0; j < len(content[i]); j++ {
			row += content[i][j]
		}
		res = append(res, row)
	}
	return res
}
