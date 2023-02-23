package storage

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

const storagePath = "/home/yuanyangen/HomeData/go/trader1024/data/datas"
const splitter = ","
const DateFormat = "2006-01-02"

type CsvStorage struct {
	allMarketData map[string]*MarketData
}

type MarketData struct {
	path        string
	DailyData   *LineStorage
	MinuteData  *LineStorage
	Minute5Data *LineStorage
}

type LineStorage struct {
	filePath string
	loaded   bool
	datas    map[int64]*model.KNode
}

func NewCsvStorage() *CsvStorage {
	allData := map[string]*MarketData{}
	allMarket := markets.GetAllFutureMarketIds()
	for _, mid := range allMarket {
		ms := &MarketData{
			path:        path.Join(storagePath, mid),
			DailyData:   &LineStorage{datas: map[int64]*model.KNode{}},
			MinuteData:  &LineStorage{datas: map[int64]*model.KNode{}},
			Minute5Data: &LineStorage{datas: map[int64]*model.KNode{}},
		}
		if _, err := os.Stat(ms.path); err != nil {
			e := os.MkdirAll(ms.path, 0755)
			if e != nil {
				panic(e)
			}
		}

		allData[mid] = ms
	}
	cs := &CsvStorage{
		allMarketData: allData,
	}
	return cs
}

func (cs *CsvStorage) SaveData(marketId string, t model.LineType, kdatas []*model.KNode) {
	sort.Slice(kdatas, func(i, j int) bool {
		return kdatas[i].TimeStamp < kdatas[j].TimeStamp
	})
	storage := cs.getStorage(marketId, t)
	res := make([]string, len(kdatas))
	for i, v := range kdatas {
		res[i] = convertKNodeToCSVLine(v)
	}
	content := strings.Join(res, "\n")
	err := os.WriteFile(storage.filePath, []byte(content), 0755)
	if err != nil {
		panic("write error")
	}
}

func (cs *CsvStorage) GetAllData(marketId string, t model.LineType) []*model.KNode {
	storage := cs.getAndInitFromFile(marketId, t)
	allData := []*model.KNode{}
	for _, v := range storage.datas {
		allData = append(allData, v)
	}
	sort.Slice(allData, func(i, j int) bool {
		return allData[i].TimeStamp < allData[j].TimeStamp
	})
	return allData
}

func (cs *CsvStorage) GetDataByTs(marketId string, t model.LineType, ts int64) *model.KNode {
	storage := cs.getAndInitFromFile(marketId, t)
	data, _ := storage.datas[utils.UnityDailyTimeStamp(ts)]
	return data
}

func (cs *CsvStorage) getStorage(marketId string, t model.LineType) *LineStorage {
	marketStorage := cs.allMarketData[marketId]
	if marketStorage == nil {
		panic("should not reach here")
	}
	var storage *LineStorage
	var fileName string
	switch t {
	case model.LineType_Day:
		storage = marketStorage.DailyData
		fileName = "daily"
	case model.LineType_Minite:
		storage = marketStorage.MinuteData
		fileName = "minute"

	case model.LineType_5Minite:
		storage = marketStorage.Minute5Data
		fileName = "minute5"
	default:
		panic("error")
	}

	storage.filePath = path.Join(marketStorage.path, fileName)
	return storage
}

func (cs *CsvStorage) getAndInitFromFile(marketId string, t model.LineType) *LineStorage {
	storage := cs.getStorage(marketId, t)
	if storage == nil || storage.loaded {
		return storage
	}

	content, err := os.ReadFile(storage.filePath)
	if err != nil {
		panic(err)
		//return
	}
	tmp := strings.Split(string(content), "\n")
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == "" {
			continue
		}
		node := convertCsvLineToKNode(tmp[i])
		storage.datas[utils.UnityDailyTimeStamp(node.TimeStamp)] = node
	}
	storage.loaded = true
	return storage
}

func convertKNodeToCSVLine(knode *model.KNode) string {
	tmp := make([]string, 11)
	tmp[0] = knode.Date
	tmp[1] = floatToStr(knode.High)
	tmp[2] = floatToStr(knode.Low)
	tmp[3] = floatToStr(knode.Open)
	tmp[4] = floatToStr(knode.Close)
	tmp[5] = floatToStr(knode.Volume)
	tmp[6] = floatToStr(knode.Turnover)
	tmp[7] = floatToStr(knode.Swing)
	tmp[8] = floatToStr(knode.Increase)
	tmp[9] = floatToStr(knode.IncreaseMount)
	tmp[10] = floatToStr(knode.Turnover)
	return strings.Join(tmp, splitter)
}

func convertCsvLineToKNode(in string) *model.KNode {
	tmp := strings.Split(in, splitter)
	if len(tmp) != 11 {
		panic("csv data format error")
	}
	knode := &model.KNode{}
	knode.Date = tmp[0]
	knode.High = strToFloat(tmp[1])
	knode.Low = strToFloat(tmp[2])
	knode.Open = strToFloat(tmp[3])
	knode.Close = strToFloat(tmp[4])
	knode.Volume = strToFloat(tmp[5])
	knode.Turnover = strToFloat(tmp[6])
	knode.Swing = strToFloat(tmp[7])
	knode.Increase = strToFloat(tmp[8])
	knode.IncreaseMount = strToFloat(tmp[9])
	knode.Turnover = strToFloat(tmp[10])
	t, err := time.Parse(DateFormat, knode.Date)
	if err != nil {
		panic("time format error")
	}
	knode.TimeStamp = t.Unix()
	return knode
}

func strToFloat(in string) float64 {
	f, err := strconv.ParseFloat(in, 64)
	if err != nil {
		panic("data error")
	}
	return f
}

func floatToStr(in float64) string {
	return fmt.Sprintf("%f", in)
}
