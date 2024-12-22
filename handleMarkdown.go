package main

import (
	"fmt"
	"regexp"
	"strings"
)

func (f *File) ToHTML() {
	content := f.content
	content = imgToHtml(content)
	content = linkToHtml(content)
	content = allItalicAndBoldToHtml(content)
	content = boldToHtml(content)
	content = italicToHtml(content)

	oListToHtml(content)

	f.html = fmt.Sprintf("<article>%s</article>", content)
}

func matchRegex(regex, content string) [][]string {
	re := regexp.MustCompile(regex)
	return re.FindAllStringSubmatch(content, -1)
}

func imgToHtml(content string) (modifiedContent string) {
	matches := matchRegex(`(\!\[(.+)\]\((.+)\))`, content)
	for _, match := range matches {
		HTMLimg := fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", match[3], match[2])
		content = strings.Replace(content, match[1], HTMLimg, 1)
	}
	return content
}

func linkToHtml(content string) (modifiedContent string) {
	matches := matchRegex(`[^!](\[(.+)\]\((.+)\))`, content)
	for _, match := range matches {
		HTMLLink := fmt.Sprintf("<a href=\"%s\">%s</a>", match[3], match[2])
		content = strings.Replace(content, match[1], HTMLLink, 1)
	}
	return content
}

func allItalicAndBoldToHtml(content string) (modifiedContent string) {
	matches := matchRegex(`[*]{3}([^ ][^*\n]+[^ ])[*]{3}`, content)
	if matches == nil {
		return content
	}
	for _, match := range matches {
		mdText := match[0]
		ogText := match[1]
		textSlice := strings.Fields(ogText)
		if len(textSlice) == 0 {
			content = strings.Replace(content, mdText, "<hr></hr>", 1)
			continue
		}
		content = strings.Replace(content, mdText, fmt.Sprintf("<em><strong>%s</strong></em>", ogText), 1)
	}
	return content
}

func italicToHtml(content string) (modifiedContent string) {
	tag := "<i>"
	content = strings.Replace(content, "*", tag, -1)
	index := strings.Index(content, tag)
	bRepetitions := 0
	for {
		bRepetitions += 1
		if bRepetitions%2 == 0 {
			content = fmt.Sprintf("%s/%s", content[:index], content[index:])
		}
		if strings.Index(content[index+1:], tag) == -1 {
			break
		}
		index = strings.Index(content[index+1:], tag) + len(content[:index+2])
	}
	return content
}

func boldToHtml(content string) (modifiedContent string) {
	tag := "<b>"
	content = strings.Replace(content, "**", tag, -1)
	content = strings.Replace(content, "__", tag, -1)
	index := strings.Index(content, tag)
	bRepetitions := 0
	for {
		bRepetitions += 1
		if bRepetitions%2 == 0 {
			content = fmt.Sprintf("%s/%s", content[:index], content[index:])
		}
		if strings.Index(content[index+1:], tag) == -1 {
			break
		}
		index = strings.Index(content[index+1:], tag) + len(content[:index+2])
	}
	return content
}

func oListToHtml(content string) (modifiedContent string) {
	matches := matchRegex(`(\d+\. .+)([\n]\d+\. .+)*`, content)
	fmt.Println(matches)
	return content
}
