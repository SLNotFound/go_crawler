package parser

import (
	"fmt"
	"go_crawler/engine"
	"go_crawler/model"
	"regexp"
	"strconv"
)

var ageReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
var marriageReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var xinzuoReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var heightReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)cm</div>`)
var weightReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)kg</div>`)
var workPlaceReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var incomeReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var workReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var educationReg = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)

func ParseUserList(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageReg))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageReg)
	profile.Xinzuo = extractString(contents, xinzuoReg)

	height, err := strconv.Atoi(extractString(contents, heightReg))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightReg))
	if err != nil {
		profile.Weight = weight
	}

	profile.WorkPlace = extractString(contents, workPlaceReg)
	profile.Income = extractString(contents, incomeReg)
	profile.Work = extractString(contents, workReg)
	profile.Education = extractString(contents, educationReg)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func TestParser(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageReg))
	if err != nil {
		profile.Age = age
	}
	fmt.Println(profile)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
