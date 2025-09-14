package apaas

var extraChecker map[string]ExtraChecker = make(map[string]ExtraChecker, 2)

func RegisterExtraChecker(key string, checker ExtraChecker) {
	extraChecker[key] = checker
}

type ExtraChecker func(extra *ExtraMeta, dest, e map[string]any) error
