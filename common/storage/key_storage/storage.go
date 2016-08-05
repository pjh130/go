// Usage:
// import(
//   "github.com/pjh130/go/common/storage/key_storage"
// )
// 
// 

package key_storage

import(
	"fmt"
)

var provides = make(map[string]Storage)

// 目前支持注册的类型：
// 1: DriveHashStorage
// 2: DriveRedisStorage
func Register(provideName string, config string) {
	//检查是否已经注册了
	if _, dup := provides[provideName]; dup {
		panic("Register called twice for provider " + provideName)
	}
	
	//判断是否支持
	switch provideName {
		case DriveHashStorage:
			var hashspder = NewHashStorage()
			provides[provideName] = hashspder
		
		case DriveRedisStorage:
			var redisspder = NewRedisStorage(config)
			provides[provideName] = redisspder
			
		case DriveMysqlStorage:
			var mysqlspder = NewMysqlStorage(config)
			provides[provideName] = mysqlspder
			
		default:
		panic("Unsurport provide name")
	}
}

func GetStorage(provideName string)(Storage, error){
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("Unknown provide %q (forgotten import?)", provideName)
	}
	
	return  provider, nil
}


// Storage
type Storage interface {
	// Return storage drive name
	Driver() string

	// Get Value
	Get(key string, dst interface{}) error

	//Get raw value
	GetRaw(key string) (interface{}, error)

	// Set Value
	Set(key string, v interface{}) error

	GetBool(key string) (bool, error)

	GetInt(key string) (int, error)

	GetInt64(key string) (int64, error)

	GetString(key string) (string, error)

	GetFloat64(key string) (float64, error)

	// Delete Storage
	Del(key string)

	// Auto Delete Set
	SetExpire(key string, v interface{}, seconds int64) error
}