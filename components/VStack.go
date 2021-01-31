package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
)

func VStack(contents []FlexComponent) *BoxComponent {
	return &BoxComponent{
		Type:            FlexComponentTypeBox,
		Layout:          FlexBoxLayoutTypeVertical,
		Contents:        contents,
	}
}

func ColoredVStack(contents []FlexComponent, hexColor string) *BoxComponent {
	return &BoxComponent{
		Type:            FlexComponentTypeBox,
		Layout:          FlexBoxLayoutTypeVertical,
		Contents:        contents,
		BackgroundColor: hexColor,
	}
}

func CustomVStack(contents []FlexComponent, styling *BoxStyle) *BoxComponent {
	return &BoxComponent{
		Type:            FlexComponentTypeBox,
		Layout:          FlexBoxLayoutTypeVertical,
		Contents:        contents,
		Flex: 			 styling.flex,
		Spacing: 		 styling.spacing,
		Margin: 		 styling.margin,
		BackgroundColor: styling.color,
	}
}