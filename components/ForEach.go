package components

import (
	. "github.com/line/line-bot-sdk-go/linebot"
)


func ForEachText(rows []string) []FlexComponent {
	var res []FlexComponent

	for i := 0; i < len(rows); i++ {
		res = append(res, &TextComponent{
			Type:       FlexComponentTypeText,
			Text:       rows[i],
			Size:       FlexTextSizeTypeSm,
			Wrap:       false,
			Weight:     FlexTextWeightTypeBold,
		})
	}

	return res
}

// ForEach Primitives

func ForEachString(rows []string, content func(item string) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(rows); i++ {
		res = append(res, content(rows[i]))
	}
	return res
}

func ForEachInt(items []int, maker func(item int) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

func ForEachFloat(items []float64, maker func(item float64) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

func ForEachBool(items []bool, maker func(item bool) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

// ForEach Slices

func ForEachIntSlice(items [][]int, maker func(item []int) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

func ForEachStringSlice(items [][]string, maker func(item []string) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

func ForEachFloatSlice(items [][]float64, maker func(item []float64) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

func ForEachBoolSlice(items [][]bool, maker func(item []bool) FlexComponent) []FlexComponent {
	var res []FlexComponent
	for i := 0; i < len(items); i++ {
		res = append(res, maker(items[i]))
	}
	return res
}

