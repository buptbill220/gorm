package apaas

import (
	"fmt"
	"regexp"
	"strings"
)

const _formular_ = `\b[a-zA-Z_][a-zA-Z0-9_]*(\.[a-zA-Z_][a-zA-Z0-9_]*){1,}\b`
const RECORD_TAG = "record."
const RECORD_PREX = "record"

var _formular_regex = regexp.MustCompile(_formular_)

// extract all likes "a.b.c" sub field from formular
func ExtractNestedFields(s string) []string {
	return _formular_regex.FindAllString(s, -1)
}

// anchor_id.faction_name => anchor_id_faction_name
func FormatSQLColumnTag(tag string) string {
	return strings.Replace(tag, ".", "_", -1)
}

// anchor_id.faction_name => record.anchor_id_faction_name
func FormatByteScriptColumnTag(tag string) string {
	return fmt.Sprintf("%s%s", RECORD_TAG, FormatSQLColumnTag(tag))
}
