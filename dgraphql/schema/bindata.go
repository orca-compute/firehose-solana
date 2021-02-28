// Code generated by go-bindata.
// sources:
// graphs.graphql
// query.graphql
// query_alpha.graphql
// schema.graphql
// serum_fill.graphql
// serum_market.graphql
// serum_order.graphql
// serum_order_tracker.graphql
// token.graphql
// DO NOT EDIT!

package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _graphsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x49\xcc\xcc\xa9\x0c\xcb\xcf\x29\xcd\x4d\x55\xa8\xe6\x52\x50\x50\x50\x48\x49\x2c\x49\xb5\x52\x08\xc9\xcc\x4d\x55\x04\xf3\xcb\x12\x73\x4a\x53\xad\x14\xdc\x72\xf2\x13\x4b\xcc\x4c\x14\xb9\x6a\x01\x01\x00\x00\xff\xff\x32\x54\xee\x0b\x38\x00\x00\x00")

func graphsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_graphsGraphql,
		"graphs.graphql",
	)
}

func graphsGraphql() (*asset, error) {
	bytes, err := graphsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "graphs.graphql", size: 56, mode: os.FileMode(420), modTime: time.Unix(1614181722, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\x4f\x6f\xdb\x48\x0c\xc5\xef\xfe\x14\x2f\xba\x6c\x16\x30\x82\x2c\xb0\xd8\x83\x8f\x9b\x34\x68\x0b\x34\xfd\xe3\xf4\x03\xd0\x12\x65\x11\x19\x0f\x5d\x92\x72\x6a\x04\xfd\xee\xc5\x8c\xa2\x34\x0d\x52\xe4\xa6\xf1\xbc\x47\x72\x7e\x8f\x8e\xe3\x9e\xf1\x79\x64\x3b\xe2\x7e\xb1\x00\x42\x6f\x39\xfb\x69\xab\x63\x8e\x15\xbe\x4a\x8e\xff\xfe\x5d\xa2\x1d\xcd\xd5\x56\x58\x87\x49\xde\xfe\xbd\xc2\x4d\x91\x5d\x68\xce\xdc\x86\x68\x3e\x79\xb4\x9e\x52\xd7\x19\xbb\xcf\xda\x93\x59\xfc\xa6\xdb\x72\x95\x35\x4d\xf3\x85\x63\xb4\x8c\x44\x1e\xf8\xe7\xfc\x1c\xbd\xa4\xe4\xe8\xd5\x10\x03\x63\x2b\x07\xce\x08\xa3\x8e\x6d\x89\x1d\xd9\x2d\xc7\x12\x6a\xd8\x68\x0c\x67\x98\xcc\x0e\x31\xe3\x03\x9b\xcb\x26\x31\x34\xa7\x23\x3a\x0a\x82\x2b\x92\x1c\x78\x3a\x64\x0d\x1c\x39\xb0\x27\x0b\x68\x0f\x82\x27\x0d\xc4\x40\x01\x71\x48\xae\x0d\x4d\x35\xe0\x1c\xb8\x93\x94\xaa\x67\xc3\xb0\xda\x86\x3b\xdc\x0d\x9c\xd1\x52\x4a\xdc\x9d\x35\x4d\xb3\x00\x9c\x6d\xdc\x5d\x49\x4a\x6f\xc5\x43\xed\x78\x3a\xcd\x3a\x3f\x79\x9e\xf9\x09\xae\xf5\xec\x78\x86\xac\x56\xfa\x50\xd5\xaf\x33\x5f\xff\x12\x3f\x2b\xf3\x88\xd4\x27\xd1\x5f\x8e\x49\x87\x8e\xe4\x01\xcc\x93\xd1\xa7\xbb\xcb\x72\x75\x49\x41\x2f\x45\xb6\x7e\x41\x77\xb2\xf8\xb1\x58\x34\x4d\x73\x51\x07\x73\x18\x7f\x1b\xc5\xb8\x43\x28\x5a\xcd\x21\x79\x64\xb0\xc4\xc0\x56\xb2\xbc\x23\xeb\x6a\x6a\xd4\xde\x96\x6f\x47\x6f\xba\x03\x21\x89\xd7\x30\xf6\xb4\x95\x4c\xc1\x1d\x38\xf1\x8e\x73\x78\x99\xb1\x2e\xe4\x27\xda\xf2\xbb\xdc\x2b\xee\x17\x40\x7d\xdf\xd4\xb4\xd8\x4a\x62\xbd\x98\xc7\x6c\x9b\x7f\x2c\x75\x97\x18\x9d\x21\x51\x66\x72\x26\x6b\x87\x39\x64\xdd\xef\xd5\x25\x18\x9d\xd8\x84\x6e\x22\x02\x78\x90\xc5\xc5\x6f\xb4\x2b\xd5\x17\xfa\xd6\x8d\x7d\xa5\xed\x8c\x62\x2e\xcf\xb9\xfb\x53\xf1\x4b\x0e\xb6\x9d\x64\x76\x48\x2d\x66\x5c\xd6\x92\x0a\x1a\x06\x1d\x48\x12\x95\xed\xa6\x3e\x78\xfa\x6f\x68\x2e\x9b\xb9\x37\x76\xce\x85\xdc\xe6\x88\x18\xc4\x27\x83\xe4\x5e\xe7\xae\x03\xf9\x35\x7f\x8f\x02\x72\x85\xff\x55\x13\x53\xae\x01\x7a\x4b\x89\xec\x61\xcb\xe6\xd3\x8d\xec\x78\xfe\xbe\x4a\x4a\x4f\xae\xde\xaf\x3f\x5e\xff\x0c\x00\x00\xff\xff\xb6\x56\x7f\x76\x25\x04\x00\x00")

func queryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphql,
		"query.graphql",
	)
}

func queryGraphql() (*asset, error) {
	bytes, err := queryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphql", size: 1061, mode: os.FileMode(420), modTime: time.Unix(1614181722, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _query_alphaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\x54\xa8\xae\xe5\x02\x04\x00\x00\xff\xff\x76\xca\x60\x3d\x0e\x00\x00\x00")

func query_alphaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_query_alphaGraphql,
		"query_alpha.graphql",
	)
}

func query_alphaGraphql() (*asset, error) {
	bytes, err := query_alphaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query_alpha.graphql", size: 14, mode: os.FileMode(420), modTime: time.Unix(1608159398, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x9e\xeb\xeb\x5e\x1c\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1611886197, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serum_fillGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\xcd\x6e\xdc\x36\x10\xbe\xeb\x29\x26\xba\xc4\x06\xdc\x5d\xa0\x4d\x73\xd0\x6d\xd3\xa4\x85\xe1\xda\x4d\xbc\x1b\xe4\x10\xe4\xc0\x25\x47\xd2\xd4\x14\x29\xf3\xc7\x1b\xa1\xe8\xbb\x17\x43\x4a\x96\xd6\xcd\x1a\x3e\x71\x25\xcd\xf7\x33\x1f\xc9\xd9\xa2\x2c\xcb\x5d\x8b\xf0\x9b\x35\x06\x65\x20\x6b\x20\x0c\x3d\x42\x6d\x1d\x08\xd8\xa2\x8b\xdd\xef\xa4\xf5\x0a\xb6\x88\x10\x5a\x84\xaf\x7f\x38\xd1\xb7\x9f\xfe\x04\x69\x1d\x82\xb4\x46\x62\x1f\xfc\xb7\xb3\x36\x84\xde\x57\xeb\xb5\xb2\xd2\xaf\x54\x1d\x3d\xae\xc8\xae\x9b\x48\x0a\xfd\x9a\x6b\x7f\x9a\x6a\xd7\x0d\x33\xdc\xeb\xf5\x79\x59\x96\x45\x52\x7b\xd4\x59\xd8\xf8\xa7\x00\x28\x37\xa0\xc9\x07\xb0\x35\xa0\x6a\xd0\x43\xb0\x73\x6d\x59\x40\x7e\x5b\xc1\xd7\xc7\x97\x1f\x54\x83\xdf\x5e\x15\x8c\xbd\x34\xb5\x75\x9d\xc8\x3d\x59\x10\xa4\xa0\x17\x0d\x99\xf4\x86\xc1\xbd\x68\x90\x8b\x2a\xf8\x38\xfe\x7a\x55\xfc\x5b\x70\x22\xc5\x06\x3c\x99\x46\x2f\x9c\x01\x6a\xec\xd0\x24\x2f\x9c\x83\x9c\x9d\xb2\xc5\x55\xf1\xff\x66\xd8\x4b\x6e\xa3\x2c\x3f\x45\x74\x03\xc8\xe8\xbc\x75\x17\x20\xb4\xb6\x07\x32\x0d\x0c\x36\xb2\x37\x69\x4d\x20\x13\x11\x6a\x0c\xb2\xe5\x0f\x0e\x7d\xd4\xc1\x03\x3e\xa0\x01\x51\x07\x74\x40\x26\xa0\x73\xb1\x4f\x9a\xb6\x66\xac\x5b\xd8\xb8\x80\x03\x85\xd6\xc6\x00\x1d\x79\x76\x0f\x02\xf6\x28\xc2\x8a\x8d\xc1\x28\x5d\xc1\x36\x38\x32\x4d\x4e\x28\x6f\xfd\xdc\xa2\xdd\xff\x8d\x72\x02\x18\xab\xb0\x9a\x3f\xa6\x68\x8e\xfb\x9b\x7a\x63\x12\xaf\x6d\x00\x13\xbb\x7d\x32\x0a\x87\x96\x64\x9b\x72\x5a\xb0\x4b\x19\x9d\x43\x35\xf2\x33\xe2\x26\x76\x15\x7c\x26\x13\xde\xbe\x59\x3a\x0a\x4e\x18\x2f\x52\x57\xaf\x3d\x90\x51\xf8\x3d\x35\x47\x26\x51\x26\xad\x43\x8b\x0e\x9f\x57\x58\xd0\x5c\x32\xc7\x8f\xa4\xc8\xf8\xe0\xe2\x49\xa9\x05\xc5\x4b\x14\x17\x6c\x27\x15\xad\x53\xe8\xa6\xa8\xa2\xa1\xfb\x88\x40\x0a\x4d\xa0\x9a\xd0\x4d\xe2\x02\x1a\xe2\xad\xef\x84\xbb\xc3\x69\x4b\x12\xf4\x54\x66\xd4\xa1\x0f\xa2\xeb\xa7\x03\xba\xd7\x56\xde\x2d\xf7\x82\xfc\x33\x51\x4d\xe8\x0a\x76\xd4\xe1\x93\xcd\x50\xe8\x5e\x7b\xe8\xe3\x5e\x93\x84\x3b\x1c\xd2\x74\x38\x26\xbc\x00\x5a\xe1\x2a\x09\x5b\xc3\x29\x89\x00\xfb\x38\x80\x75\xe0\x51\xeb\x19\x71\xd4\x50\xe6\xfe\xd1\xa1\xcc\x65\x09\xf6\xc2\x16\x32\x62\x3c\xb1\xd7\xe9\x61\x22\xfc\x92\xe0\x9e\x14\x3e\xe5\x20\xcf\x12\x17\x80\x14\x5a\x74\xf0\xee\xf2\x3d\x9c\xb1\x5f\x74\xe7\x6c\x7d\xb3\xbd\x82\xb3\x7d\x1c\xd0\x9d\x4f\xa7\x96\xd2\xad\x20\x85\xbb\xa1\x3f\xca\xe9\x3e\x0a\x13\x28\x0c\x29\x7f\x7b\x87\x26\x87\x70\x10\x1e\x1c\x4a\xa4\x07\x54\x30\x9e\xaa\x79\x84\x25\xce\x09\x79\x3b\x96\x55\xb0\x63\xfc\xa6\xb3\xd1\x84\x17\x4a\xf4\x3c\xde\xf8\xf2\x3f\x2f\xf1\x51\xd0\x69\xfa\xde\x91\xc4\xcc\xf4\xb8\x5d\xf8\x5d\xc8\xb0\xfc\x1f\xe0\x42\x85\xc6\x76\x24\xf3\x68\x25\x0f\x42\x1f\xc4\xc0\xb7\x07\xce\xf6\xc2\xe3\xe8\x6d\x0d\xf7\xd1\x86\xf1\xe9\x3c\xfb\x48\x12\xa7\x87\x10\xe2\x8e\xef\x80\xc6\x07\xd4\x4f\x8c\xd4\xac\x9e\x49\xea\x5c\x56\x1d\x81\xe6\xf9\xb4\xe8\x0e\xd2\x88\x7a\x10\x3a\xe2\x7c\x67\x00\x14\xf9\x5e\x8b\x61\xf2\xc1\x07\x91\x31\x63\x30\x4c\x84\x26\x76\xc7\x9e\x98\xe8\xdd\x66\xfb\xa1\x00\xd8\xde\x5e\xff\x9c\x97\x5f\xf2\xf2\x26\x2f\xbf\xe6\xe5\x6d\x01\x70\xbd\xbd\xbd\x9e\x79\xc6\xd3\x92\x38\x36\xdb\x2b\x66\xba\x7c\x5f\x00\x7c\xbe\xb9\xba\xf9\xeb\xcb\x0d\x17\xfe\x17\x00\x00\xff\xff\x66\x61\x75\x98\x88\x07\x00\x00")

func serum_fillGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_fillGraphql,
		"serum_fill.graphql",
	)
}

func serum_fillGraphql() (*asset, error) {
	bytes, err := serum_fillGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_fill.graphql", size: 1928, mode: os.FileMode(420), modTime: time.Unix(1612374293, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serum_marketGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x52\x4d\x6f\xd3\x40\x10\xbd\xfb\x57\xbc\xfa\x02\x48\x25\xbe\x20\x0e\xbe\x55\x0d\x20\x24\x90\xa8\xd2\x72\xa9\x7a\x98\x7a\xc7\xf6\xd2\xf5\x8e\xbb\x3b\x4b\x14\x21\xfe\x3b\x5a\xdb\x69\x12\xc8\x2f\xe0\x94\xec\xcc\xfb\x98\xf7\xe4\xb2\x2c\x6f\x7b\xc6\xb5\x78\xcf\x8d\x5a\xf1\xd0\xdd\xc8\x68\x25\x80\xb0\xe1\x90\x86\xaf\x14\x9e\x58\x57\xd8\x30\x43\x7b\xc6\xfd\xa7\x40\x63\x7f\xf3\x05\x8d\x04\x46\x23\xbe\xe1\x51\xe3\xc3\xeb\x5e\x75\x8c\x75\x55\x19\x69\xe2\xca\xb4\x29\xf2\xca\x4a\xd5\x25\x6b\x38\x56\x19\xfb\x76\x8f\xad\xba\xac\xf0\xec\xaa\x37\x65\x59\x16\x93\xdf\x91\xd3\xd1\x29\xbf\x0a\xa0\x2c\xcb\x2b\x38\x1b\x15\xd2\x82\x4d\xc7\x11\x2a\x33\xfe\xa3\x75\x2e\x2b\x60\x9e\xd7\xb8\x3f\x92\xf9\x60\x3a\x7e\xb8\x28\x66\x85\xcf\xbe\x95\x30\xd0\x9c\x4f\x40\xd6\x60\xa4\xce\xfa\x69\x32\x4b\x8c\xd4\x71\x86\xd5\xf8\xb6\xfc\xdb\x93\x6f\x45\xc9\xa1\x91\xe4\xa7\x1b\x86\x49\x1e\xda\x93\x62\x9b\x0b\x48\xce\x20\xb0\xa6\xb0\x28\x69\xc6\x5f\x67\x78\x8d\x3b\xeb\xf5\xfd\xbb\x8b\xe2\x77\x51\xe4\xe5\x15\xa2\xf5\x9d\x3b\xc9\x9b\x0f\x05\x3b\x1e\x78\xd6\xcf\x1d\x37\x87\x0e\x72\xf4\x55\x71\xae\xa8\x89\xb8\x54\x74\x93\x38\xec\xd0\xa4\x10\x25\x5c\x82\x9c\x93\xad\xf5\x1d\x76\x92\x72\xde\x46\xbc\x5a\x9f\x18\x2d\x6b\xd3\xe7\x45\xe0\x98\x9c\x46\xf0\x4f\xf6\xa0\x56\x39\xc0\x7a\xe5\x10\xd2\x38\xb9\x4a\x9b\xb9\xe1\xe8\x90\x4b\x6c\xad\xf6\x92\x14\x83\x8d\x39\x05\x08\x8f\x4c\xba\x9a\x43\xcf\xd6\x35\x36\x1a\xac\xef\x5e\xaa\xeb\x4f\x2e\x86\x3c\xfe\xe0\x66\x4f\xf1\x62\xb8\x3e\x5e\x4f\x35\xfd\x9d\x12\xfb\x88\x59\x6b\xee\xfe\x55\x04\x19\x13\x38\xc6\x45\x69\x79\x9d\x73\x7f\x61\x18\x1b\x47\x47\x3b\x78\x1a\x78\xf1\xa7\x81\xf7\x94\x7f\x18\x11\x8f\x14\x19\x2a\x4f\xec\x17\x97\x3c\xb8\xcd\xef\x1a\xd3\xcf\x19\xce\x73\x12\x3d\x25\x4d\x93\x13\xd6\x99\x8c\x6b\xb2\x6e\xb7\x26\xa5\xff\x2d\xeb\x81\x65\x72\x44\x18\x52\x5a\xc0\xd3\xe0\xbb\xb8\x34\xf0\xdd\x66\x5d\xe3\x7e\x7d\x18\x5c\x3c\xe4\x4f\xe1\x4f\x00\x00\x00\xff\xff\xe4\x75\x30\x33\x9b\x04\x00\x00")

func serum_marketGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_marketGraphql,
		"serum_market.graphql",
	)
}

func serum_marketGraphql() (*asset, error) {
	bytes, err := serum_marketGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_market.graphql", size: 1179, mode: os.FileMode(420), modTime: time.Unix(1614181722, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serum_orderGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xcb\x6a\xdb\x40\x14\x86\xf7\x7a\x8a\x3f\xda\x24\x81\xe0\x6e\x4a\x17\xda\xb9\x89\x4b\x45\xe3\x38\xc4\x2a\x25\x94\x62\x46\x9a\xe3\xea\x60\x69\x46\x99\x0b\xae\x28\x79\xf7\x32\xa3\xb8\x56\x1a\x85\x16\xba\x94\x35\xff\xf7\xcd\xb9\x58\x89\xeb\x3b\xc2\x9a\x8c\x6f\x57\x46\x92\xc1\xcf\x04\x48\xd3\xb4\xa8\x09\x3a\xfe\xa0\x7c\x5b\x92\x81\x57\xfc\xe0\x09\x2c\x49\x39\xde\x32\x99\x34\x4d\x13\x84\xb7\x19\x3e\xb3\x72\xef\xde\x9e\x24\xc7\xa8\x6d\xb4\x3b\x24\x59\x61\x5f\x73\x55\xc3\xd5\xcf\x4c\xba\xaa\xbc\x31\x24\x67\x03\x29\x44\x6e\xa6\x69\xce\x08\x65\x45\xe5\x58\xab\x53\x8b\x5a\xd8\x1a\x7b\x76\x35\xab\x88\x8c\xae\x7d\x4d\x86\xfe\x62\x18\x61\x3e\x0a\x5b\x67\x58\x3b\xc3\xea\xfb\xeb\x26\x56\x92\x7e\xfc\xaf\x2a\x0f\x90\xa9\xaa\x58\x59\x67\xfc\xab\xae\x11\xe2\x9f\x94\x23\xdc\xab\x4a\xc7\x2d\x59\x27\xda\x0e\x7a\x1b\x71\x65\xa3\xab\xdd\x78\x42\x6c\x07\xc5\x07\x6e\x9a\x17\x45\x1d\xd2\x19\x0a\x6e\xe9\x8f\xbe\x49\x32\xa7\x16\x9d\x2f\x1b\xae\xb0\xa3\x1e\x5b\x6d\x46\xc0\x78\xe7\x0b\xf0\x8c\x66\xd1\xac\x55\x28\x48\x38\x94\xbe\x87\x36\xb0\xd4\x34\xc7\x48\x2b\xcc\x8e\xdc\xb1\x99\x92\xcc\xd4\xb8\x86\x63\x31\x36\x51\xc3\x64\x9b\x86\x48\x36\x9c\x58\xc6\x87\x03\xf1\x4b\xcc\x5b\x96\xf4\x02\xc2\x36\x48\x2e\x40\xec\x6a\x32\x78\x9f\x5f\xe1\x2c\xdc\x98\xcc\x79\xb8\xfc\x7c\xfd\x09\x67\xa5\xef\xc9\x9c\x1f\xb6\x99\x25\x65\x58\xb3\xa4\xa2\xef\x9e\xb5\xaa\x33\x5c\x11\x6c\x47\x55\xf8\x1b\x49\x3c\x8d\xfb\x28\x9b\x21\x1c\x93\xa4\x74\xcb\x95\x88\xf3\x67\x0b\xd1\xec\x45\x1f\xd6\x04\x67\xa5\xb0\x04\xa7\x77\xa4\xf0\x06\x0f\x5e\xbb\xa7\xa7\xf3\x41\x1d\x05\x53\xdd\x92\x64\xd9\x90\xc4\x83\x17\xca\xb1\xeb\x5f\xaa\x07\xc0\xe1\x7d\x86\x22\x60\xe7\xad\xf6\xca\x8d\x41\xa3\xc6\x84\xcf\xc7\xd3\x94\xfa\x8e\x32\xac\xee\xae\x16\x77\x9b\xe2\xfe\x76\x31\x0e\x6c\xb9\x69\x2c\x84\xb5\xba\x62\xe1\x48\xc2\xe9\x61\x09\x8e\xd2\x78\x24\xc3\xd7\xdf\xdb\xf7\xed\x24\x79\x4c\x12\x52\xbe\x1d\x41\xe3\xe7\xe9\x3a\x5f\xe6\x45\x02\xe4\xcb\xe5\xe2\x2a\x9f\x17\x8b\xcd\xea\x6e\x73\x39\xbf\xb9\x5c\x5c\x27\xc0\xed\x6a\x5d\x6c\x56\x37\xd7\xf7\xc9\x63\xf2\x2b\x00\x00\xff\xff\xf3\x27\x17\x0a\xde\x04\x00\x00")

func serum_orderGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_orderGraphql,
		"serum_order.graphql",
	)
}

func serum_orderGraphql() (*asset, error) {
	bytes, err := serum_orderGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_order.graphql", size: 1246, mode: os.FileMode(420), modTime: time.Unix(1613413149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serum_order_trackerGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x41\x6f\xdb\x3c\x0c\xbd\xfb\x57\xb0\x39\x7f\xf8\x4e\xc3\x0e\x01\x76\x30\x52\x17\x33\x56\x38\x41\xe2\x6e\xbb\xad\x8a\xcd\xc6\xc4\x6c\x3a\x95\xa8\xa5\xc1\xb6\xff\x3e\xc8\xb1\x25\x27\xa9\xb7\x4b\x6f\x16\x49\xbd\xf7\xf4\x44\xd1\x51\x84\x6c\x1b\x58\xae\x6f\x93\xf5\xb7\x4d\x1e\xe7\x09\xfc\x8c\x00\x00\x1e\xb2\x4f\xd9\xf2\x4b\x16\x75\x8b\xd9\x6c\x06\x29\x97\x54\x28\x41\x03\x87\x0a\x19\xa4\x42\x68\x75\x89\x1a\x2a\x65\x60\x8b\xc8\x10\x17\x05\xee\x05\x4b\xd8\x1e\xbb\x34\xbe\x14\x95\xe2\x1d\xce\x66\xb3\x0e\x25\x5e\xad\xd6\xcb\xcf\xc9\x6d\xc0\x8c\xb9\xc7\x20\x03\xe4\x30\xc9\x80\x11\x25\xb6\x27\x21\x09\xe8\x45\xdb\xec\x6b\x14\xac\x8f\x50\x28\x2e\xb0\xae\xb1\x04\xc5\xe5\x48\x09\x19\xe0\x16\xea\x96\x77\xa8\xa1\xdd\x23\x0f\xcc\x8b\x38\x5b\x24\xf7\xf7\x23\xea\x18\xf6\x4a\x0b\xa9\x1a\x9e\xad\x62\x21\x39\x06\x26\x7c\xc1\xc2\x0a\x96\xff\xc1\xd6\x8a\x83\xd7\x08\x4a\x23\x18\xa1\xba\x06\x53\x29\x8d\xa6\x83\x07\xa9\x94\x74\xa9\x83\x22\x21\xde\x81\xb4\xb0\x45\x78\xa2\x4e\x9c\xcf\x2a\x63\xda\x82\x94\xb3\xe6\x40\x52\x05\xc5\xff\x0f\x02\x57\xf1\x3a\x4f\xe3\x7b\x2f\x2f\x77\xf6\xb1\x90\xc6\xa0\xaf\x7d\x02\xc5\x97\x9e\x8f\x5c\x39\xb1\x0e\x88\xc9\xd7\x64\xf1\x90\x27\xb7\xd1\xef\x28\x92\xe3\x1e\x61\x83\xda\x36\x4b\xb7\x7b\x23\x4a\xb0\xbf\xe6\x9e\xeb\x84\xca\xb6\xd9\xa2\x06\xcb\xf4\x6c\x11\xa8\x74\x0a\x9e\x08\xf5\x80\xc9\xb6\x99\xc3\x03\xb1\xbc\x7f\x77\xe3\xa5\xee\x35\xfe\xa0\xd6\x9a\x13\x2a\x99\xee\x74\x43\xd0\x89\x76\xeb\x40\xee\x8f\x7c\xb6\x6f\x3e\xee\xc0\x80\x5d\x58\xad\x91\xe5\x0c\xba\x8f\xc1\x77\x6e\x0f\xdc\x75\x0b\xfe\x85\x65\x8c\x30\x41\x22\x5a\xb1\x21\xa1\x96\x33\xd5\x5c\xd1\x84\xec\xc0\xd2\x45\x54\xe1\x42\x9e\xe6\x1c\x63\x20\xca\xd7\x71\xb6\x49\xf3\x74\x99\xf5\x6c\xa1\x6c\x0e\x9d\xce\xdc\x07\xdc\x3d\x59\x76\x34\x17\x09\xf8\x70\x19\x49\x99\x04\x7e\x5d\x46\xfd\xeb\xbb\xca\x2c\xfc\x73\xb9\x4a\xad\x4e\x8f\xa0\x3e\xde\xd1\xeb\x05\x49\xff\x16\xfa\x26\x7a\x4d\x89\x6f\xa4\xc7\xae\x8b\x1e\x07\x07\xc3\x6d\x40\x89\xa2\xa8\x36\xde\xae\xae\x70\x3e\xaa\xf0\x5d\x3a\x75\xa8\xb7\x20\xb9\x99\x62\x09\x06\x9d\x68\xfc\x7c\xd9\xd4\xad\x64\xe3\xa6\x3f\xcb\xe6\xa1\x11\x3e\x2a\x53\xcd\x61\x23\x9a\x78\x37\x5d\x95\x72\x89\x2f\x13\x60\x29\x1b\xd1\xf6\x9f\x65\xb1\xcc\x21\xa7\x06\x27\xcf\x72\x79\xa3\xde\xb8\x0c\x0f\xdd\x88\x38\x8d\xa5\x61\x6e\x8c\xc6\xe7\xe0\x9b\x2a\x4b\x2c\xdd\xee\xde\x3b\xf7\x39\x49\x37\xf4\xc7\xdb\x5d\xd0\xe8\x7f\x14\xde\x4f\x0f\x1f\x2f\x16\xc9\xca\x0d\xb5\xf3\xa1\xee\x56\x77\xa9\xff\x1c\x8d\xbe\x3f\x01\x00\x00\xff\xff\xc7\xca\x06\x7f\xdf\x06\x00\x00")

func serum_order_trackerGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_order_trackerGraphql,
		"serum_order_tracker.graphql",
	)
}

func serum_order_trackerGraphql() (*asset, error) {
	bytes, err := serum_order_trackerGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_order_tracker.graphql", size: 1759, mode: os.FileMode(420), modTime: time.Unix(1614304105, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tokenGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6f\xd4\x30\x10\x85\xef\xfe\x15\xaf\x39\x81\x54\x92\x0b\xe2\x90\x5b\x41\x08\x55\x02\x89\xaa\xe5\x54\xf5\xe0\xb5\x27\x89\xc1\xf1\x04\x7b\xdc\x55\x40\xfd\xef\xc8\xce\x2e\x1b\x24\x6e\xf1\x9b\x67\xbf\xef\x4d\x94\x6a\x9a\xe6\x61\x22\x7c\xe0\x10\xc8\x88\xe3\x00\x59\x17\xc2\xc0\x11\x1a\x0f\xfc\x83\x42\x8b\x7b\x22\xc8\x44\x78\xfc\x14\xf5\x32\xdd\x7d\x86\xe1\x48\x30\x1c\x0c\x2d\x92\x9e\x5e\x4d\x22\x4b\xea\xbb\xce\xb2\x49\xad\x1d\x72\xa2\xd6\x71\x37\x66\x67\x29\x75\xc5\xfb\xe6\xec\xed\xc6\xf2\xc2\x4f\xdf\xbd\x6e\x9a\x46\xd5\xa4\x9a\xb1\x8b\xff\xad\x80\xe6\x06\xde\x25\x01\x0f\x20\x3b\x52\x82\xf0\xe6\x43\xa3\xb0\x49\x3d\x1e\xab\xf2\xd1\x8e\xf4\x74\xa5\xca\xa5\xdb\x30\x70\x9c\xf5\x56\x82\xa1\x9d\xc5\xa2\x47\x17\xaa\x52\x2e\x2e\x7a\xa4\x62\xea\xf1\xf5\xf4\x75\xa5\x5e\xea\x0a\xd4\x0d\x92\x0b\xa3\x3f\xe1\x80\x3c\xcd\x14\x2a\x40\x29\x6e\x2e\x78\x85\xab\x55\xff\xd2\x17\x86\x8d\xbb\x69\xee\x32\xc5\x15\x26\xc7\xc4\xf1\x1a\xda\x7b\x3e\xba\x30\x62\xe5\x5c\x98\x0c\x07\x71\x21\x13\x06\x12\x33\x95\x41\xa4\x94\xbd\x24\xd0\x33\x05\xe8\x41\x28\xc2\x05\xa1\x18\xf3\x52\xf3\x78\x28\x77\xe3\x0e\xe1\x1a\x47\x27\x13\x67\xc1\xec\x52\xa1\x86\xc6\x81\xb4\xb4\x05\x0a\xa7\xe8\x1e\xf7\x12\x5d\x18\xb7\xcd\x6c\xff\x78\xab\xc6\x87\xef\x64\xce\xe6\xc0\x96\xfa\x6d\x50\x57\xb1\x2b\x55\x0b\x69\x6b\x23\xa5\x74\x79\x0d\x98\x5d\x90\x9b\x2c\x13\x47\x27\xeb\x7e\x30\x44\xa2\x5f\xf4\xdf\x51\xca\xcb\xe2\xd7\x1e\xdf\x5c\x90\x77\x6f\x8b\x62\xc9\xb8\x59\xfb\xd4\xe3\x36\x48\x11\x9e\x29\xba\xc1\x91\xed\xf1\x9e\xd9\x93\x0e\x35\x8b\x44\x9f\xf0\xbe\x90\xe8\x42\x78\x01\x2c\x4a\x85\x4c\xeb\x7c\x60\xbf\xcf\x0b\x7a\xa6\xfd\xd9\xf3\xc8\xe7\xb3\x02\x8e\x74\x48\x4e\xfe\x3a\xd4\x8b\xfa\x13\x00\x00\xff\xff\xa5\x9b\x3c\x1f\x08\x03\x00\x00")

func tokenGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_tokenGraphql,
		"token.graphql",
	)
}

func tokenGraphql() (*asset, error) {
	bytes, err := tokenGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "token.graphql", size: 776, mode: os.FileMode(420), modTime: time.Unix(1612206111, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"graphs.graphql": graphsGraphql,
	"query.graphql": queryGraphql,
	"query_alpha.graphql": query_alphaGraphql,
	"schema.graphql": schemaGraphql,
	"serum_fill.graphql": serum_fillGraphql,
	"serum_market.graphql": serum_marketGraphql,
	"serum_order.graphql": serum_orderGraphql,
	"serum_order_tracker.graphql": serum_order_trackerGraphql,
	"token.graphql": tokenGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"graphs.graphql": &bintree{graphsGraphql, map[string]*bintree{}},
	"query.graphql": &bintree{queryGraphql, map[string]*bintree{}},
	"query_alpha.graphql": &bintree{query_alphaGraphql, map[string]*bintree{}},
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"serum_fill.graphql": &bintree{serum_fillGraphql, map[string]*bintree{}},
	"serum_market.graphql": &bintree{serum_marketGraphql, map[string]*bintree{}},
	"serum_order.graphql": &bintree{serum_orderGraphql, map[string]*bintree{}},
	"serum_order_tracker.graphql": &bintree{serum_order_trackerGraphql, map[string]*bintree{}},
	"token.graphql": &bintree{tokenGraphql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

