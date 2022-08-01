package utils

import (
	"net/url"
	"sort"
)

type keyValue struct{ key, value string }

type OrderKeyValue []keyValue

func (p OrderKeyValue) Len() int      { return len(p) }
func (p OrderKeyValue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p OrderKeyValue) Less(i, j int) bool {
	return p[i].key < p[j].key
}

func (p OrderKeyValue) appendValues(values url.Values) OrderKeyValue {
	for k, vs := range values {
		if k == "sign" {
			continue
		}
		for _, v := range vs {
			p = append(p, keyValue{k, v})
		}
	}
	return p
}

func ApiCheckSign(urlParams url.Values, apiKey string) bool {

	params := make(OrderKeyValue, 0)
	params = params.appendValues(urlParams)
	sort.Sort(params)

	str := ""
	for _, kvObj := range params {
		str += kvObj.key + kvObj.value
	}

	str = apiKey + str + apiKey

	sign := GetMd5(str)
	reqSign := urlParams.Get("sign")

	return ToLower(sign) == ToLower(reqSign)
}

func ApiGenSign(urlParams url.Values, apiKey string) string {
	params := make(OrderKeyValue, 0)
	params = params.appendValues(urlParams)
	sort.Sort(params)

	str := ""
	for _, kvObj := range params {
		str += kvObj.key + kvObj.value
	}

	str = apiKey + str + apiKey

	sign := GetMd5(str)

	return ToLower(sign)

}
