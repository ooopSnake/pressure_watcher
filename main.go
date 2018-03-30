//+build linux

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const CpuFreqFilePatten = "/sys/devices/system/cpu/cpu%d/cpufreq/cpuinfo_cur_freq"
const CpuTempFile = "/sys/class/thermal/thermal_zone0/temp"
const ProcStatFile = "/proc/stat"

func GetCpuFreq() []string {
	ret := make([]string, 0)
	for i := 0; i < runtime.NumCPU(); i++ {
		fname := fmt.Sprintf(CpuFreqFilePatten, i)
		bytes, err := ioutil.ReadFile(fname)
		if err != nil {
			ret = make([]string, 0)
			return ret
		}
		line := string(bytes)
		freqStr := "0"
		freqInt, err := strconv.Atoi(strings.Trim(line, "\n"))
		if err != nil {
			fmt.Println(err)
			freqInt = 0
		}
		if freqInt != 0 {
			v := float64(freqInt) / 1e6
			freqStr = fmt.Sprintf("%.03f GHz", v)
		}
		ret = append(ret, freqStr)
	}
	return ret
}

func GetCpuTemp() string {
	ret := ""
	bytes, err := ioutil.ReadFile(CpuTempFile)
	if err != nil {
		return ret
	}
	line := string(bytes)
	tempStr := "0"
	tempInt, err := strconv.Atoi(strings.Trim(line, "\n"))
	if err != nil {
		tempInt = 0
	}
	if tempInt != 0 {
		v := float64(tempInt) / 1e3
		tempStr = fmt.Sprintf("%.01f °C", v)
	}
	ret = tempStr
	return ret
}

func ReadStat() *linuxproc.Stat {
	stat, _ := linuxproc.ReadStat(ProcStatFile)
	return stat
}

func CheckRoot() bool {
	return os.Getuid() == 0
}

func GetCpuUsage() []string {
	s1 := ReadStat()
	time.Sleep(time.Second * 1)
	s2 := ReadStat()
	cpuUsageAll := make([]string, 0, len(s1.CPUStats))
	if s1 == nil || s2 == nil || len(s1.CPUStats) != len(s2.CPUStats) {
		return cpuUsageAll
	}
	for idx := range s1.CPUStats {
		stat1 := s1.CPUStats[idx]
		stat2 := s2.CPUStats[idx]
		deltaU := (float64)(stat2.User - stat1.User)
		deltaN := (float64)(stat2.Nice - stat1.Nice)
		deltaS := (float64)(stat2.System - stat1.System)
		deltaI := (float64)(stat2.Idle - stat1.Idle)
		val := 100 * (deltaU + deltaN + deltaS) / (deltaU + deltaN + deltaS + deltaI)
		cpuUsageAll = append(cpuUsageAll, fmt.Sprintf("%.01f %%", val))
	}
	return cpuUsageAll
}

type CpuInfo struct {
	Id    string `json:"id"`
	Freq  string `json:"freq"`
	Usage string `json:"usage"`
}

var bindAddr = flag.String("addr", "", "http addr , eg : 127.0.0.1")
var bindPort = flag.String("port", "12345", "http listen port , eg : 8080")
var noHttp = flag.Bool("nohttp", false, "disable http server")

func genCpuStat() interface{} {
	coreNum := runtime.NumCPU()
	cpuFreq := GetCpuFreq()
	cpuUsage := GetCpuUsage()
	retJsonObj := make(map[string]interface{})
	if len(cpuFreq) != coreNum || len(cpuFreq) != len(cpuUsage) {
		return retJsonObj
	}
	cpuInfo := make([]CpuInfo, 0, coreNum)
	for i := 0; i < coreNum; i++ {
		cpuInfo = append(cpuInfo, CpuInfo{
			Id:    fmt.Sprintf("cpu %d", i),
			Freq:  cpuFreq[i],
			Usage: cpuUsage[i]})
	}
	retJsonObj["cpuInfo"] = cpuInfo
	retJsonObj["cpuTemp"] = GetCpuTemp()
	return retJsonObj
}

func main() {
	if !CheckRoot() {
		fmt.Println("require root !")
		os.Exit(1)
	}
	flag.Parse()
	if *noHttp {
		out, _ := json.MarshalIndent(genCpuStat(), "", "  ")
		fmt.Printf("%s\n", out)
	} else {
		StartServer(fmt.Sprintf("%s:%s", *bindAddr, *bindPort), genCpuStat)
	}
	//default exit 0
}
