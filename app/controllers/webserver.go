package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"gotrading/app/models"
	"gotrading/config"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

// func init() {
// 	fmt.Printf("Templates loaded: %+v\n", templates)
// }
//templtesにはメモリなどが格納される
// Templates loaded: &{escapeErr:<nil> text:0xc000226480 Tree:0xc00023ec60 nameSpace:0xc000228300}

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	limit := 100
	duration := "1m"
	durationTime := config.Config.Durations[duration]
	df, err := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)

	// エラーハンドリングを追加
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	// データが空の場合の処理
	if df == nil || len(df.Candles) == 0 {
		http.Error(w, "No data available. Please wait for data collection.", http.StatusServiceUnavailable)
		return
	}

	err = templates.ExecuteTemplate(w, "google.html", df.Candles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
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
