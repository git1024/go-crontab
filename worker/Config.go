package worker

import (
	"encoding/json"
	"io/ioutil"
)

// 程序配置
type Config struct {
	EtcdEndpoints         []string `json:"etcdEndpoints"`         // etcd集群地址
	EtcdDialTimeout       int      `json:"etcdDialTimeout"`       // etcd超时时间
	MongodbUri            string   `json:"mongodbUri"`            // mongo地址
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"` // mongo连接超时时间
	JobLogBatchSize       int      `json:"jobLogBatchSize"`       // 日志批量大小
	JobLogCommitTimeout   int      `json:"jobLogCommitTimeout"`   // 日志自动提交超时时间
}

var (
	// 单例
	G_config *Config
)

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	// 1,把配置文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2,做JSON反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3,赋值单例
	G_config = &conf
	return
}
