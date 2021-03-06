package gchart

import (
	"github.com/bitly/go-simplejson"
	"github.com/zieckey/goini"
	"strconv"
	"strings"
	"time"
	"fmt"
)

const date_format = "2006-01-02"

type SplineChart int

func (c *SplineChart) Parse(ini *goini.INI, date string) (map[string]string, error) {
	args := make(map[string]string)

	datas := make([]interface{}, 0)

	kv, _ := ini.GetKvmap(goini.DefaultSection)

	today := time.Now().Format(date_format)
	_, err := time.Parse(date_format, date)
	if err == nil {
		today = date
	}
	fmt.Println("date", today)

	var temp float64
	for k, v := range kv {
		if !strings.HasPrefix(k, DataPrefix + today) {
			continue
		}
		temp = 0

		dd := strings.Split(v, ", ")
		df := make([]interface{}, 0)
		var total float64
		for i, d := range dd {
			val, err := strconv.ParseFloat(d, 64)
			if err == nil {
				if i == 0 {
					day, err := time.Parse(date_format, k[len(DataPrefix):])
					if err == nil {
						dayBefore := day.AddDate(0, 0, -1).Format(date_format)
						vv, ok := ini.Get(DataPrefix + dayBefore)
						if ok {
							index := strings.LastIndex(vv, ", ")
							if index > 0 {
								last, err := strconv.ParseFloat(vv[index + 2:], 64)
								if err == nil {
									fmt.Println("last", last)
									temp = last
								} else {
									fmt.Println(err)
								}
							}
						}
					}
				}
				if val > 0 && temp > 0 {
					total += (val - temp)
					df = append(df, val - temp)
				} else {
					df = append(df, 0)
				}
				temp = val
			}
		}
		args["TotalNum"] = strconv.FormatFloat(total, 'g', -1, 64)

		json := simplejson.New()
		json.Set("name", k[len(DataPrefix):])
		json.Set("data", df)
		datas = append(datas, json)
	}

	json := simplejson.New()
	json.Set("DataArray", datas)

	b, _ := json.Get("DataArray").Encode()
	println(string(b))
	args["DataArray"] = string(b)

	return args, nil
}

func (c *SplineChart) Template() string {
	return TemplateSplineHtml
}
