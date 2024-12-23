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

func matchRegex(regex, content string) (matches [][]string) {
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
	matches := matchRegex(`[*]{1}([^ *\n][\w\d ]+[^ *\n])[*]{1}`, content)
	for _, match := range matches {
		ogBold := strings.TrimSpace(match[0])
		boldTag := "<i>" + match[1] + "</i>"
		content = strings.Replace(content, ogBold, boldTag, 1)
	}
	matches = matchRegex(`[_]{1}([^ _\n][\w\d ]+[^ _\n])[_]{1}`, content)
	for _, match := range matches {
		ogBold := strings.TrimSpace(match[0])
		boldTag := "<i>" + match[1] + "</i>"
		content = strings.Replace(content, ogBold, boldTag, 1)
	}
	return content
}

func boldToHtml(content string) (modifiedContent string) {
	matches := matchRegex(`[*]{2}([^ *\n][\w\d ]+[^ *\n])[*]{2}`, content)
	for _, match := range matches {
		ogBold := strings.TrimSpace(match[0])
		boldTag := "<b>" + match[1] + "</b>"
		content = strings.Replace(content, ogBold, boldTag, 1)
	}
	matches = matchRegex(`[_]{2}([^ _\n][\w\d ]+[^ _\n])[_]{2}`, content)
	for _, match := range matches {
		ogBold := strings.TrimSpace(match[0])
		boldTag := "<b>" + match[1] + "</b>"
		content = strings.Replace(content, ogBold, boldTag, 1)
	}
	return content
}

func oListToHtml(content string) (modifiedContent string) {
	slicedContent := strings.Split(content, "\n")
	previousMatch := ""
	for i, line := range slicedContent {
		re := regexp.MustCompile(`^\d+\. .+ *$`)

		currentMatch := re.FindString(line)
		nextMatch := ""
		if i+1 < len(slicedContent) {
			nextMatch = re.FindString(slicedContent[i+1])
		}
		dotIdx := strings.Index(currentMatch, ".")

		if currentMatch == "" {
			previousMatch = currentMatch
			continue
		}
		if previousMatch != "" && nextMatch != "" {
			slicedContent[i] = "    <li>" + currentMatch[dotIdx+2:] + "</li>"
			previousMatch = currentMatch
			continue
		}
		if previousMatch == "" && nextMatch == "" {
			slicedContent[i] = "<ol>\n    <li value=\"" + currentMatch[:dotIdx] + "\">" + currentMatch[dotIdx+2:] + "</li>\n</ol>"
			previousMatch = currentMatch
			continue
		}
		if previousMatch != "" && nextMatch == "" {
			slicedContent[i] = "    <li>" + currentMatch[dotIdx+2:] + "</li>\n</ol>"
			previousMatch = currentMatch
			continue
		}
		if previousMatch == "" && nextMatch != "" {
			slicedContent[i] = "<ol>\n    <li value=\"" + currentMatch[:dotIdx] + "\">" + currentMatch[dotIdx+2:] + "</li>"
			previousMatch = currentMatch
			continue
		}
	}
	return strings.Join(slicedContent, "\n")
}
