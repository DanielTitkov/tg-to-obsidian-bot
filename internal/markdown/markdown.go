package markdown

import (
	"strings"
	"time"
)

const (
	contentPlaceholder = "{{content}}"
	timePlaceholder    = "{{time}}"
	datePlaceholder    = "{{date}}"
	titlePlaceholder   = "{{title}}"
	dateFormat         = "2006-01-02"
	timeFormat         = "15:04"
)

func WrapWithMarkdown(text, template, title string) (string, error) {
	noteText := strings.ReplaceAll(template, contentPlaceholder, text)
	noteText = strings.ReplaceAll(noteText, titlePlaceholder, title)
	noteText = strings.ReplaceAll(noteText, datePlaceholder, time.Now().Format(dateFormat))
	noteText = strings.ReplaceAll(noteText, timePlaceholder, time.Now().Format(timeFormat))

	return noteText, nil
}
