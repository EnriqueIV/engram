package dashboard

import "unicode/utf8"

// truncateContent truncates a string to max runes, appending "..." if needed.
func truncateContent(s string, max int) string {
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	runes := []rune(s)
	return string(runes[:max]) + "..."
}

// typeBadgeVariant returns a badge color variant for an observation type.
func typeBadgeVariant(obsType string) string {
	switch obsType {
	case "decision", "architecture":
		return "success"
	case "bugfix":
		return "danger"
	case "discovery", "learning":
		return "warning"
	default:
		return "muted"
	}
}
