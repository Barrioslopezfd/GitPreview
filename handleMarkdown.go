package main

import (
    "fmt"
    "strings"
    "regexp"
)

func ConvertToHTML(content string) string {
    paragraphs:=strings.Split(content, "\n")
    content = strings.Join(paragraphs, "\n")
    content = convertLinks(content)
    content = convertBoldItalic(content, "**", "<b>")
    content = convertBoldItalic(content, "*", "<i>")

    return fmt.Sprintf("<div>%s</div>",content)
}

func convertLinks(content string) (modifiedContent string) {
    linkRegex := `(!?\[([^\]]+)\]\(([^\)]+)\))`
    re := regexp.MustCompile(linkRegex)
    matches := re.FindAllStringSubmatch(content, -1)
    if matches == nil {
        return content
    }
    for _, match:=range matches{
        if match[1][0] == 91 {
            content=linksToHtml(content, match)
        } else {
            content=imgToHtml(content, match)

        }
    }
    return content
}

func linksToHtml(content string, match []string) (modifiedContent string) {
    MDLink := match[1]
    textOnLink := match[2]
    link := match[3]
    HTMLLink := fmt.Sprintf("<a href=\"%s\">%s</a>", link, textOnLink)
    modifiedContent = strings.Replace(content, MDLink, HTMLLink, 1)
    return modifiedContent
}

func imgToHtml(content string, match []string) (modifiedContent string) {
    MDimg:=match[1]
    altText:=match[2]
    imgLink:=match[3]
    HTMLimg:=fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", imgLink, altText)
    modifiedContent=strings.Replace(content, MDimg, HTMLimg, 1)
    return modifiedContent
}

func convertBoldItalic(content string, boldOrItalic string, tag string) (modifiedContent string) {
    content=strings.Replace(content, boldOrItalic, tag, -1)
    index:=strings.Index(content, tag)
    bRepetitions:=0
    for {
        bRepetitions+=1
        if bRepetitions % 2 == 0 {
            content=fmt.Sprintf("%s/%s",content[:index],content[index:])
        }
        if strings.Index(content[index+1:], tag) == -1 {
            break
        }
        index = strings.Index(content[index+1:], tag) + len(content[:index+2])
    }
    return content
}

func convertOList(content string) (modifiedContent string) {
    OLRegex:=`
}
