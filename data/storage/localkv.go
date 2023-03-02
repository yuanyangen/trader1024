package storage

import (
    "encoding/json"
    "fmt"
    "github.com/boltdb/bolt"
    "github.com/yuanyangen/trader1024/engine/model"
    "github.com/yuanyangen/trader1024/engine/utils"
    "path"
    "sort"
    "github.com/spaolacci/murmur3"
)

const storageDataPath = "/home/yuanyangen/HomeData/go/trader1024/data/datas"
//
//第一层是：数据的来源，包括： main, eastMoney等，
//第二层是具体的DB文件，分成128个文件，文件名字是通过marketID hash得到的。
// bucket的名字是 market + dataType ，数据类型是枚举值： daily_k, minute5_k
// 每个bucket的k 是对应数据unix 时间戳， value是数据的json编码的结果。

//表示 存放 某Market的某个类型的在某个时间范围范围内的量
// 每一个文件de key的

type KVStorage struct {

    dbs []*bolt.DB
}

var allStorage = map[string]*KVStorage{}
func init() {
    InitStorage("eastmoney")
    InitStorage("main")
}

func EastMoneyStorage() *KVStorage {
    return allStorage["eastmoney"]
}

func MainStorage() *KVStorage {
    return allStorage["main"]
}

const dbSplitCount =128
func InitStorage(name string) *KVStorage {
    cs := &KVStorage{}
    dbs := make([]*bolt.DB, dbSplitCount)

    for i:=0;i < dbSplitCount; i++ {
        dbPath := path.Join(storageDataPath, name)
        utils.CreateDirIfNotExist(dbPath)

        filePath := path.Join(dbPath, fmt.Sprintf("%v_.db", i))
        db, err := bolt.Open(filePath, 0755, nil)
        if err != nil {
            panic("open  db error")
        }
        dbs[i]= db
    }
    cs.dbs  = dbs
    allStorage[name] = cs
	return cs
}

func (cs *KVStorage) getDbBycketName(marketId string) *bolt.DB {
    idx := murmur3.Sum32([]byte(marketId)) % dbSplitCount
    return cs.dbs[idx]
}

func (cs *KVStorage) getBycketName(tx *bolt.Tx, marketId string, t model.LineType) *bolt.Bucket {
    bucketName := fmt.Sprintf("%v_%v", marketId, t)
    bucket := tx.Bucket([]byte(bucketName))
    if bucket != nil {
        return bucket
    }
    var err error
	bucket, err = tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		panic(err)
	}

	return bucket
}

func (cs *KVStorage) SaveData(marketId string, t model.LineType, kdatas []*model.KNode) {
    db := cs.getDbBycketName(marketId)
	db.Update(func(tx *bolt.Tx) error {
		bucket := cs.getBycketName(tx, marketId, t)
		for _, knode := range kdatas {
			if knode.TimeStamp == 0 {
				panic("ts == 0")
			}
			key := fmt.Sprint(knode.TimeStamp)
			valueB, _ := json.Marshal(knode)
			err := bucket.Put([]byte(key), valueB)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
}

func (cs *KVStorage) GetAllData(marketId string, t model.LineType) []*model.KNode {
	allData := []*model.KNode{}
    db := cs.getDbBycketName(marketId)

	db.View(func(tx *bolt.Tx) error {
		bucket := cs.getBycketName(tx, marketId, t)
		bucket.ForEach(func(k, v []byte) error {
			knode := model.NewFromJson(v)
			allData = append(allData, knode)
			return nil
		})
		return nil
	})
	sort.Slice(allData, func(i, j int) bool {
		return allData[i].TimeStamp < allData[j].TimeStamp
	})
	return allData
}

func (cs *KVStorage) GetDataByTs(marketId string, t model.LineType, ts int64) *model.KNode {
	var res *model.KNode
    db := cs.getDbBycketName(marketId)

	db.View(func(tx *bolt.Tx) error {
		bucket := cs.getBycketName(tx, marketId, t)
		key := fmt.Sprint(ts)
		val := bucket.Get([]byte(key))
		res = model.NewFromJson(val)
		return nil
	})
	return res
}
