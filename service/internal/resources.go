// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "esc -pkg internal -o resources.go templates/"; DO NOT EDIT.

package internal

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/component_header.html": {
		name:    "component_header.html",
		local:   "templates/component_header.html",
		size:    156,
		modtime: 1605957790,
		compressed: `
H4sIAAAAAAAC/1SMsQqDMBRFd7/iIq7q5lBiltKt9B8CPklQX6R1e9x/L6ZQ2vXcc65ZE3AZ0V3ztmcV
PW467TnpQVZmzZp0Kfs96VJQizTjw1uyAgAXB+8C4lPmsT4fydqbdY+wCen64F0fB19iWV/yF/54X0en
U3kHAAD//zT+SdCcAAAA
`,
	},

	"/templates/extensions_table.html": {
		name:    "extensions_table.html",
		local:   "templates/extensions_table.html",
		size:    353,
		modtime: 1605957790,
		compressed: `
H4sIAAAAAAAC/2SQwU7DMBBE7/2KlemRNJwjxxwQHDnwB248DRbOOnK2tGD531HTQIvqk1fzZjU7Wuw2
gCb5CmjVNiaHVE2j7Tz3DT0osyIiynltqWlp8xSHMTJYntmN0bOUsgDJcg9ap3jw7HC8n7+z5y0epgU7
oxX5HeETfMGv9NPTkv4i2e6jT3HPrqE7AEui8yaECbdWkzPYUXWlaHFkg++5VR1YkJTRlt4Tdq06HVfK
4zeOAp58ZLYD2pw3L/sQXu2AUpT5N+raGl2Lu0TRtaTfqsCulJWu52bNTwAAAP//sz5qjmEBAAA=
`,
	},

	"/templates/footer.html": {
		name:    "footer.html",
		local:   "templates/footer.html",
		size:    15,
		modtime: 1605957790,
		compressed: `
H4sIAAAAAAAC/7LRT8pPqbTjstHPKMnNsQMEAAD//wEFevAPAAAA
`,
	},

	"/templates/header.html": {
		name:    "header.html",
		local:   "templates/header.html",
		size:    467,
		modtime: 1605957790,
		compressed: `
H4sIAAAAAAAC/5TRMU8sIRAH8P4+BY/25eC9szGGxUItLIwW11giO7uMB8wG5rxsLvfdDdnTxNhoBeFP
fpnM3/y5fbzZPj/dicAp2pVph4guj52ELK0J4Hq7EkIIk4Cd8MGVCtzJPQ/rS3mOGDmCPR7Vtl1OJ6OX
lyWNmHeiQOxkDVTY71mgpyxFKDB0UuvD4aBogswQIQGXWSHpwb21Xwo9Sf1d4jlCDQD8wQTmqV5pPVDm
qkaiMYKbsCpPSTfpenAJ49w9OIaCLv6995Sr/AXtqQc1Aqc+tgn/qwv1T6czpzD3ONJ6wrxTCbPy9ROv
vuDEoocBiqjF/5RszGuV1uhFsCujl0bMC/Vz62vzZe1hY98DAAD//7qRGmLTAQAA
`,
	},

	"/templates/pipelines_table.html": {
		name:    "pipelines_table.html",
		local:   "templates/pipelines_table.html",
		size:    1946,
		modtime: 1605957790,
		compressed: `
H4sIAAAAAAAC/7SVwXLTMBCG7zyFxnRyIjVcU1scSpnhAMN0eAFZ2gRNlZVmJbdujd+dsWyrTp0LtL5k
rOjX/tlv/8hFEJUB5sOjgTKrLCmgrXdCajzs2MeMv2OMsSLQ8DAsFJPWeCew/MSE0QcsDewDLyr+tTbm
hzhCkVe8yIM6OcU3WHl3NXz+mS8W0oWBBAxAvcU3dHX49ejW9PheBxHAX1v09RHUFxHEim63IEHfA/kV
PX6SleC9XdXkpnGWwqKRIp/i07YXgu1Kdnltj84iYLhB5azG0HWjgAQegF2QfdCooPkQH+OZW/vgR9kg
3TK9Z3AP+Cyf7Y+5TdEW8u5Atka1Y+8BIOOzSmA8LI/ytgVUbDvb6VG1bW93OUW962Kn/wZxKpLC/Koq
Z+L6X/XGgWbDRGeETkcDMo0GZD+a+CNSil+AjLUF+02wL7M+AF33+clpB0YjoDhCuQC6eZJTPpIA5Mn3
dxpVSaNlxidFkQu+dK/oZSuAaj7VN0a1IUF0dZ6eIzvRc2QTvef/5yr4HNklPjd5Rn5Rcpbf2XbWJZhw
QeMmXNC4hCvdNKvQgsYtacFoGWFFxSvCNl+lu3HQFXl8JfO/AQAA//9We3KLmgcAAA==
`,
	},

	"/templates/properties_table.html": {
		name:    "properties_table.html",
		local:   "templates/properties_table.html",
		size:    420,
		modtime: 1605957790,
		compressed: `
H4sIAAAAAAAC/2SRwW7DIBBE7/6KVRr1VMc5u5gfqFT11Ds2U8sqWVuwqRoR/r1yTCpb4YAEO48ZDarV
MR7ezQkp1apqdaHEtA4U5OLQ7NrRW/gyTKYbuK/puNMFEVGMtB/Y4pfqho6UUr71hnvk0Qvt4XACyyw6
fPhxgpcBIasXoqThi/ADztRqOC8l/j+L6b57P57Z1vQEIEdZnoELeER1jGBL5WqixJJxQ89NBxZ4favg
nvTaQ95wSWnuQlVi9RrUz9yG6XXZr+vDg3TrsTX4NO6M2WLDVOLv3YJtSoWqbl+h/wIAAP//aLmk3KQB
AAA=
`,
	},

	"/templates": {
		name:  "templates",
		local: `templates/`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"templates/": {
		_escData["/templates/component_header.html"],
		_escData["/templates/extensions_table.html"],
		_escData["/templates/footer.html"],
		_escData["/templates/header.html"],
		_escData["/templates/pipelines_table.html"],
		_escData["/templates/properties_table.html"],
	},
}
