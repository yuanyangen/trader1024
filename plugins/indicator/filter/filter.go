package filter

import (
	"fmt"
	"github.com/bytedance/sonic"
	"os/exec"
)

// https://github.com/TcheL/Road2Filter/blob/master/IIR/o4zpsbwlpf.c
func LowPassFilter(wn float64, x []float64) (y []float64) {
	inputStr, _ := sonic.MarshalString(x)
	pythonCode := fmt.Sprintf(LowPassFilterPythonCode, wn, inputStr)
	cmd := exec.Command("/home/yuanyangen/HomeData/install/miniconda/bin/python", "-c", pythonCode)
	outStr, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}
	var out []float64
	sonic.Unmarshal(outStr, &out)
	return out
}

var LowPassFilterPythonCode = `
import json
from scipy.signal import butter, filtfilt
import numpy as np


def filter_data(WK):
    if len(WK)<=15:
        return [0]*len(WK)
    N = 4  # Filter order
    Wn = %v  # Cutoff frequency
    B, A = butter(N, Wn, output='ba')
    smooth_data = filtfilt(B, A, WK)
    smooth_data = np.array(smooth_data).tolist()
    return smooth_data

beforeFilter = %v
afterFilter = filter_data(beforeFilter)
print (json.dumps( afterFilter))
`
