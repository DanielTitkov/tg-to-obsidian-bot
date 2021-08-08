package markdown

import (
	"regexp"
	"strings"
	"time"
)

const (
	contentPlaceholder = "{{content}}"
	timePlaceholder    = "{{time}}"
	datePlaceholder    = "{{date}}"
	titlePlaceholder   = "{{title}}"
	refPlaceholder     = "{{ref}}"
	dateFormat         = "2006-01-02"
	timeFormat         = "15:04"
	titleRegex         = `^#\s.*\n`
)

func WrapWithMarkdown(text, template, title, ref string) (string, error) {
	noteText := strings.ReplaceAll(template, contentPlaceholder, text)
	noteText = strings.ReplaceAll(noteText, titlePlaceholder, title)
	noteText = strings.ReplaceAll(noteText, datePlaceholder, time.Now().Format(dateFormat))
	noteText = strings.ReplaceAll(noteText, timePlaceholder, time.Now().Format(timeFormat))
	noteText = strings.ReplaceAll(noteText, refPlaceholder, ref)

	return noteText, nil
}

func ExtractTitle(text string) (string, string, error) {
	re := regexp.MustCompile(titleRegex)
	title := re.Find([]byte(text))
	if title == nil {
		return text, "", nil
	}

	text = strings.Replace(text, string(title), "", 1)
	titleText := strings.Replace(string(title), "#", "", -1)
	titleText = strings.Replace(string(titleText), "\n", "", -1)
	return text, titleText, nil
}
