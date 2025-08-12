package tradingalgo

/*
Tenkan = (9-period high + 9-period low) / 2
Kijun = (26-period high + 26-period low) / 2
Senkou A = (Tenkan + Kijun) / 2
Senkou B = (52-period high + 52-period low) / 2
Chikou = current close price
*/

func minMax(inReal []float64) (float64, float64) {
	min := inReal[0]
	max := inReal[0]
	for _, price := range inReal {
		if min > price {
			min = price
		}
		if max < price {
			max = price
		}
	}
	return min, max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func IchimokuCloud(inReal []float64) ([]float64, []float64, []float64, []float64, []float64) {
	length := len(inReal)
	tenkan := make([]float64, length)
	kijun := make([]float64, length)
	senkouA := make([]float64, length)
	senkouB := make([]float64, length)
	chikou := make([]float64, length)

	for i := range inReal {
		// Tenkan-sen (9-period)
		if i >= 8 {
			min, max := minMax(inReal[i-8 : i+1])
			tenkan[i] = (min + max) / 2
		}
		
		// Kijun-sen (26-period)
		if i >= 25 {
			min, max := minMax(inReal[i-25 : i+1])
			kijun[i] = (min + max) / 2
		}
		
		// Senkou Span A (Tenkan + Kijun) / 2
		if i >= 25 {
			senkouA[i] = (tenkan[i] + kijun[i]) / 2
		}
		
		// Senkou Span B (52-period)
		if i >= 51 {
			min, max := minMax(inReal[i-51 : i+1])
			senkouB[i] = (min + max) / 2
		}
		
		// Chikou Span (current close shifted back 26 periods)
		if i >= 25 {
			chikou[i-25] = inReal[i]
		}
	}
	return tenkan, kijun, senkouA, senkouB, chikou
}
