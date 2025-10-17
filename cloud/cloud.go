// Package cloud
package cloud

type CloudDB struct {
	URL string
}

func NewCloudDB(url string) *CloudDB {
	return &CloudDB{
		URL: url,
	}
}

func (db CloudDB) Read() ([]byte, error) {
	return nil, nil
}

func (db CloudDB) Write(content []byte) {
}
