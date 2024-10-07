package govalidator

func HasLeastOneAlphabet(str string) bool {
	for _, charVariable := range str {
		if (charVariable >= 'a' && charVariable <= 'z') || (charVariable >= 'A' && charVariable <= 'Z') {
			return true
		}
	}
	return false
}

func HasLeastOneNumeric(str string) bool {
	for _, charVariable := range str {
		if charVariable >= '0' && charVariable <= '9' {
			return true
		}
	}
	return false
}

func HasLeastOneLowerCase(str string) bool {
	for _, charVariable := range str {
		if charVariable >= 'a' && charVariable <= 'z' {
			return true
		}
	}
	return false
}

func HasLeastOneUpperCase(str string) bool {
	for _, charVariable := range str {
		if charVariable >= 'A' && charVariable <= 'Z' {
			return true
		}
	}
	return false
}
