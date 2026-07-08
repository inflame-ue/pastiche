package tui

func hotkeyCode(s string) int {
	if len(s) < 5 || s[:4] != "ctrl" {
		return 0
	}

	// ctrl+letter
	if len(s) == 6 && s[4] == '+' {
		if c := s[5]; c >= 'a' && c <= 'z' {
			return int(c - 'a' + 1)
		}
	}

	return 0
}

func KeyName(code int) string {
	if code >= 1 && code <= 26 {
		return "Ctrl+" + string(rune('A'-1+code))
	}
	return ""
}