package service

import "strings"

// Look at the message and replace the placeholder with the key
// The message will be in the form of {{@placeholder_key}}


func ReplacePlaceholders(message string, placeholders map[string]string) string {
	for key, value := range placeholders {
		message = strings.ReplaceAll(message, "{{@"+key+"}}", value)
	}
	return message
}
