package master

import (
	"encoding/json"
	"io/ioutil"
)

// 程序配置
type Config struct {
	ApiPort               int      `json:"apiPort"`               // api端口
	ApiReadTimeout        int      `json:"apiReadTimeout"`        // api读超时时间
	ApiWriteTimeout       int      `json:"apiWriteTimeout"`       // api写超时时间
	EtcdEndpoints         []string `json:"etcdEndpoints"`         // etcd集群地址
	EtcdDialTimeout       int      `json:"etcdDialTimeout"`       // etcd超时时间
	WebRoot               string   `json:"webroot"`               // web文件目录
	MongodbUri            string   `json:"mongodbUri"`            // mongo地址
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"` // mongo连接超时时间
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

	// 2,做Json反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3,赋值单例
	G_config = &conf
	return
}
