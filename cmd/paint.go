package main

import (
	"PaintTest/tools"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", FuncProvider)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}

func FuncProvider(writer http.ResponseWriter, request *http.Request) {
	rqMsg := strings.TrimPrefix(request.URL.Path, "/")
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	args := request.Form

	result := ""
	if rqMsg == "DrawER" {

		dct := make(map[string][]float64)
		for k, v := range args {
			dct[k] = lstStrToFloat(v)
		}

		result = tools.PaintCurrencyScheme(dct)
	}

	if _, err := writer.Write([]byte(result)); err != nil {
		panic(err)
	}

}

func lstStrToFloat(arg []string) []float64 {
	var floatLst []float64

	lst := strings.Split(arg[0], ",")
	for elem := range lst {
		num, _ := strconv.ParseFloat(lst[elem], 64)
		floatLst = append(floatLst, num)
	}
	return floatLst
}
