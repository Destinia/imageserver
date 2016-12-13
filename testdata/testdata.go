// Package testdata provides test images.
package testdata

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/pierrre/imageserver"
	imageserver_source "github.com/pierrre/imageserver/source"
)

var (
	// Dir is the path to the directory containing the test data.
	Dir = initDir()

	// Images contains all images by filename.
	Images = make(map[string]*imageserver.Image)

	// Server is an Image Server that uses filename as source.
	Server = imageserver.Server(imageserver.ServerFunc(func(params imageserver.Params) (*imageserver.Image, error) {
		source, err := params.GetString(imageserver_source.Param)
		if err != nil {
			return nil, err
		}
		im, err := Get(source)
		if err != nil {
			return nil, &imageserver.ParamError{Param: imageserver_source.Param, Message: err.Error()}
		}
		return im, nil
	}))
)

// Get returns an Image for a name.
// Get returns an Image for a name.
func Get(name string) (*imageserver.Image, error) {
	im, ok := Images[name]
	if !ok {
		loaded, err := loadImage(name, "jpeg")
		if err != nil {
			return nil, err
		}
		return loaded, nil
	}
	return im, nil
}

func initDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Dir(currentFile)
}

func loadImage(filename string, format string) (*imageserver.Image, error) {
	filePath := filepath.Join(Dir, filename)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unknown image \"%s\"", filename)
	}
	im := &imageserver.Image{
		Format: format,
		Data:   data,
	}
	Images[filename] = im
	return im, nil
}
