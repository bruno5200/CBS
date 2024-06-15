package util

import "strings"

func EncodeQueryParams(params map[string]string) string {
	query := ""
	for key, value := range params {
		query += key + "=" + value + "&"
	}
	query = strings.Replace(query, " ", "%20", -1)

	query = strings.Replace(query, ":", "%3A", -1)

	query = strings.Replace(query, ",", "%2C", -1)

	query = strings.Replace(query, "{", "%7B", -1)

	query = strings.Replace(query, "}", "%7D", -1)

	query = strings.Replace(query, "\"", "%22", -1)

	// query = strings.Replace(query, "/", "%2F", -1)
	// query = strings.Replace(query, "[", "%5B", -1)
	// query = strings.Replace(query, "]", "%5D", -1)
	// query = strings.Replace(query, "(", "%28", -1)
	// query = strings.Replace(query, ")", "%29", -1)
	// query = strings.Replace(query, "!", "%21", -1)
	// query = strings.Replace(query, "=", "%3D", -1)

	return query[:len(query)-1]
}
