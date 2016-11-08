package gchart

import (
	"github.com/bitly/go-simplejson"
	"github.com/zieckey/goini"
	"strconv"
	"strings"
)

type SplineChart int

func (c *SplineChart) Parse(ini *goini.INI) (map[string]string, error) {
	args := make(map[string]string)
	args["ChartType"], _ = ini.Get("ChartType")
	args["Title"], _ = ini.Get("Title")
	args["SubTitle"], _ = ini.Get("SubTitle")
	args["YAxisText"], _ = ini.Get("YAxisText")
	args["XAxisNumbers"], _ = ini.Get("XAxisNumbers")
	args["ValueSuffix"], _ = ini.Get("ValueSuffix")

	datas := make([]interface{}, 0)
	datas2 := make([]interface{}, 0)

	kv, _ := ini.GetKvmap(goini.DefaultSection)

	var diff float64
	var min float64
	for k, v := range kv {
		if !strings.HasPrefix(k, DataPrefix) {
			continue
		}

		dd := strings.Split(v, ", ")
		jd := make([]interface{}, 0)
		diffjd := make([]interface{}, 0)
		for _, d := range dd {
			val, err := strconv.ParseFloat(d, 64)
			if err == nil {
				jd = append(jd, val)
				if val > 0 && min == 0 {
					min = val - 100
					args["Min"] = strconv.FormatFloat(min, 'g', -1, 64)
				}

				if diff > 0 && val > diff {
					diffjd = append(diffjd, val - diff)
				} else {
					diffjd = append(diffjd, 0)
				}
				diff = val
			}
		}
		json := simplejson.New()
		json.Set("name", k[len(DataPrefix):])
		json.Set("data", jd)
		datas = append(datas, json)

		json = simplejson.New()
		json.Set("name", k[len(DataPrefix):])
		json.Set("data", diffjd)
		datas2 = append(datas2, json)
	}

	json := simplejson.New()
	json.Set("DataArray", datas)

	b, _ := json.Get("DataArray").Encode()
	println(string(b))
	args["DataArray"] = string(b)

	json = simplejson.New()
	json.Set("DataArray2", datas2)

	b, _ = json.Get("DataArray2").Encode()
	println(string(b))
	args["DataArray2"] = string(b)
	return args, nil
}

func (c *SplineChart) Template() string {
	return TemplateSplineHtml
}
