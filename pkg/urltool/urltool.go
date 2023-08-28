package urltool

import (
	"errors"
	"net/url"
	"path"
)

func GetBasePath(URL string) (basePath string, err error) {
	myUrl, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("invalid host")
	}
	basePath = path.Base(myUrl.Path)
	return basePath, nil
}
