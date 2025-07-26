package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"gotrading/app/models"
	"gotrading/config"
)

var templates = template.Must(template.ParseFiles("app/views/chart.html"))

// func init() {
// 	fmt.Printf("Templates loaded: %+v\n", templates)
// }
//templtesにはメモリなどが格納される
// Templates loaded: &{escapeErr:<nil> text:0xc000226480 Tree:0xc00023ec60 nameSpace:0xc000228300}

func viewChartHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "chart.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)

}

var apiValidPath = regexp.MustCompile("^/api/candle/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func apiCandleHandler(w http.ResponseWriter, r *http.Request) {
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		APIError(w, "product_code is required", http.StatusBadRequest)
		return
	}
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000
	}

	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.Config.Durations[duration]

	df, _ := models.GetAllCandle(productCode, durationTime, limit)
	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartWebServer() error {
	http.HandleFunc("/api/candle/", apiMakeHandler(apiCandleHandler))
	http.HandleFunc("/chart", viewChartHandler)
	http.HandleFunc("/debug/", debugDataHandler) // 追加
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
	//dに8080を代入するという意味
}

func debugDataHandler(w http.ResponseWriter, r *http.Request) {
	limit := 10
	duration := "1m"
	durationTime := config.Config.Durations[duration]
	df, err := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	if df == nil || len(df.Candles) == 0 {
		fmt.Fprintf(w, "No data found\n")
		return
	}

	fmt.Fprintf(w, "Found %d candles:\n\n", len(df.Candles))

	for i, candle := range df.Candles {
		fmt.Fprintf(w, "%d. Time: %s\n", i+1, candle.Time.Format("2006-01-02 15:04:05"))
		fmt.Fprintf(w, "   Open: %.2f\n", candle.Open)
		fmt.Fprintf(w, "   High: %.2f\n", candle.High)
		fmt.Fprintf(w, "   Low: %.2f\n", candle.Low)
		fmt.Fprintf(w, "   Close: %.2f\n", candle.Close)
		fmt.Fprintf(w, "   Volume: %.2f\n", candle.Volume)
		fmt.Fprintf(w, "   Range: %.2f (High-Low)\n\n", candle.High-candle.Low)
	}
}
