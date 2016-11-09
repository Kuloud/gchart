package gchart

import (
	"net/http"
	"text/template"
)

var (
	ChartHandlers = make(map[string]ChartIf)
	ChartFiles    []string
)

func handler(w http.ResponseWriter, r *http.Request) {
	tt, err := Parse(ChartFiles[0], r.URL.Path[1:])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if t, err := template.New("foo").Parse(tt.tmpl); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		if err = t.ExecuteTemplate(w, "T", tt.args); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func ListenAndServe(addr string) error {
	http.HandleFunc("/", handler)

	var err error
	ChartFiles, err = LookupChartFiles(".")
	if err != nil {
		return err
	}

	ChartHandlers["column"] = new(SplineChart)

	return http.ListenAndServe(addr, nil)
}
