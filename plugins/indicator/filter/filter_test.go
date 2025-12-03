package filter

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/plugins/indicator/filter/butter"
	"math"
	"os/exec"
	"testing"
)

func TestFilter(t *testing.T) {
	in := []float64{
		2741, 2752, 2741, 2710, 2721, 2690, 2655, 2672, 2663, 2670, 2634, 2670, 2644, 2658, 2668, 2676, 2651, 2672, 2680, 2687, 2692, 2703, 2695, 2710, 2715, 2678, 2675, 2670, 2677, 2665, 2695, 2703, 2673, 2648, 2641, 2638, 2640, 2649, 2644, 2641, 2641, 2660, 2644, 2647, 2638, 2645, 2600, 2586, 2574, 2580, 2590, 2616, 2630, 2638, 2619, 2583, 2650, 2610, 2640, 2648, 2675, 2667, 2668, 2699, 2713, 2699, 2713, 2716, 2696, 2715, 2722, 2780, 2808, 2805, 2849, 2790, 2777, 2753, 2752, 2725, 2732, 2730, 2708, 2692, 2650, 2659, 2650, 2637, 2675, 2688, 2696, 2683, 2690, 2673, 2675, 2695, 2715, 2660, 2650, 2633, 2605, 2589, 2574, 2590, 2588, 2583, 2620, 2637, 2609, 2636, 2633, 2613, 2595, 2580, 2572, 2567, 2558, 2585, 2610, 2595, 2576, 2525, 2522, 2532, 2525, 2538, 2550, 2573, 2561, 2570, 2569, 2586, 2590, 2587, 2572, 2568, 2562, 2541, 2545, 2573, 2603, 2597, 2595, 2558, 2550, 2501, 2497, 2505, 2507, 2494, 2478, 2485, 2486, 2463, 2497, 2491, 2465, 2458, 2465, 2463, 2459, 2475, 2460, 2457, 2442, 2430, 2457, 2478, 2485, 2481, 2473, 2481, 2510, 2490, 2544, 2537, 2550, 2530, 2562, 2538, 2564, 2568, 2571, 2568, 2622, 2627, 2609, 2588, 2627, 2608, 2639, 2655, 2676, 2671, 2690, 2721, 2727, 2676, 2663, 2658, 2680, 2687, 2679, 2705, 2668, 2633, 2637, 2650, 2649, 2658, 2645, 2666, 2676, 2655, 2644, 2649, 2664, 2640, 2659, 2677, 2682, 2691, 2671, 2678, 2654, 2648, 2622, 2633, 2626, 2627, 2613, 2601, 2671, 2650, 2619, 2634, 2630, 2630, 2630, 2625, 2640,
	}
	//out := CalculateJMA(in, 5, 10, 0.2)
	//out := CalculateJMA(in, 20, 0, 0.8)
	//
	//out2 := CalculateJMA(in, 5, -50, 0.2)

	//out := talib.T3(in, 10, 0.7)
	//out2 := talib.Kama(in, 10)
	//out := talib.Trima(in, 40)
	out := talib.Sma(in, 10)
	//out := talib.Trima(in, 20)
	//out2 := talib.Dema(in, 40)
	//out2 := talib.Wma(in, 40)
	out2 := talib.MidPoint(in, 10)

	//out2, _ := talib.Mama(in, 0.5, 0.05)

	//out2 := CalculateJMA(in, 5, 0, 0.5)
	//out, out2 := talib.Mama(in, 1, 0.1)
	plotData(in, out, out2)
}

func calcButter(beforeFilter []float64) []float64 {
	wc := 0.2 // 7 hz cutoff
	lfp := butter.NewLowPass2(wc)

	fmt.Println("u  hp1  lp1  hp2  lp2  bp2  bs2  rl")
	out := []float64{}
	for i := 0; i < len(beforeFilter); i++ {
		// linear chirp 3..15 hz
		out = append(out, lfp.Next(beforeFilter[i]))
	}
	return out
}

// period [0-x]越大平滑效果越差
// phase  [-100,100]越大平滑效果越差
// power [0,1]越小，延时越高，越平滑,
func CalculateJMA(values []float64, period int, phase int, power float64) []float64 {
	if len(values) == 0 || len(values) < period {
		panic(fmt.Sprintf("[Calculate] values parameters is empty"))
	}

	beta := 0.45 * (float64(period) - 1) / (0.45*(float64(period)-1) + 2)
	alpha := math.Pow(beta, power)
	var ph float64
	if phase < -100 {
		ph = 0.5
	} else if phase > 100 {
		ph = 2.5
	} else {
		ph = float64(phase)/100 + 1.5
	}
	ma1, det0, det1, ma2, jma := 0.0, 0.0, 0.0, 0.0, 0.0
	var returnValues []float64
	for _, value := range values {
		if math.IsNaN(value) {
			panic(fmt.Sprintf("[Calculate] invalid value: %v", value))
		}
		if jma == 0 {
			jma = value
			ma1 = value
			ma2 = value
		} else {
			ma1 = (1-alpha)*value + alpha*ma1
			det0 = (value-ma1)*(1-beta) + beta*det0
			ma2 = ma1 + ph*det0
			det1 = math.Pow(1-alpha, 2)*(ma2-jma) + math.Pow(alpha, 2)*det1
			jma = jma + det1
		}

		returnValues = append(returnValues, jma)
	}

	//for i, j := 0, len(returnValues)-1; i < j; i, j = i+1, j-1 {
	//	returnValues[i], returnValues[j] = returnValues[j], returnValues[i]
	//}

	return returnValues
}

// https://github.com/TcheL/Road2Filter/blob/master/IIR/o4zpsbwlpf.c
func plotData(raw []float64, after, after2 []float64) {
	inputStr, _ := sonic.MarshalString(raw)
	afterStr, _ := sonic.MarshalString(after)
	afterStr2, _ := sonic.MarshalString(after2)
	pythonCode := fmt.Sprintf(plotpython, inputStr, afterStr, afterStr2)
	cmd := exec.Command("/home/yuanyangen/HomeData/install/miniconda/bin/python", "-c", pythonCode)
	outStr, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	var out []float64
	sonic.Unmarshal(outStr, &out)
}

var plotpython = `
import array

from scipy.signal import butter, filtfilt
import pandas as pd
from pylab import *
import mplcursors
import matplotlib.pyplot as plt
from matplotlib.ticker import MultipleLocator
#显示中文
mpl.rcParams['font.sans-serif'] = ['SimHei']

plt.subplot(1, 2, 1)
beforeFilter = %v
plt.title("滤波前", fontsize=8)
plt.plot(beforeFilter, c="y", label="CURRENT")

afterFilter = %v
plt.title("滤波后", fontsize=8)
plt.plot(afterFilter, c="b", label="CURRENT")

afterFilter2 = %v
plt.title("滤波后", fontsize=8)
plt.plot(afterFilter2, c="r", label="CURRENT")
plt.show()

`
