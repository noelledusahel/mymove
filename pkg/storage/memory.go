package storage

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

// Memory is a storage backend that uses an in memory filesystem. It is intended only
// for use in development to avoid dependency on an external service.
type Memory struct {
	root    string
	webRoot string
	logger  *zap.Logger
	fs      *afero.Afero
}

// MemoryParams contains parameter for instantiating a Memory storage backend
type MemoryParams struct {
	root    string
	webRoot string
	logger  *zap.Logger
}

// DefaultMemoryParams returns default values for MemoryParams
func DefaultMemoryParams(logger *zap.Logger) MemoryParams {
	absTmpPath, err := filepath.Abs("tmp")
	if err != nil {
		log.Fatalln(errors.New("could not get absolute path for tmp"))
	}
	storagePath := path.Join(absTmpPath, "storage")
	webRoot := "/" + "storage"

	return MemoryParams{
		root:    storagePath,
		webRoot: webRoot,
		logger:  logger,
	}
}

// NewMemory creates a new Memory struct using the provided MemoryParams
func NewMemory(params MemoryParams) *Memory {
	var fs = afero.NewMemMapFs()

	return &Memory{
		root:    params.root,
		webRoot: params.webRoot,
		logger:  params.logger,
		fs:      &afero.Afero{Fs: fs},
	}
}

// Store stores the content from an io.ReadSeeker at the specified key.
func (fs *Memory) Store(key string, data io.ReadSeeker, checksum string) (*StoreResult, error) {
	if key == "" {
		return nil, errors.New("A valid StorageKey must be set before data can be uploaded")
	}

	joined := filepath.Join(fs.root, key)
	dir := filepath.Dir(joined)

	err := fs.fs.MkdirAll(dir, 0755)
	if err != nil {
		return nil, errors.Wrap(err, "could not create parent directory")
	}

	file, err := fs.fs.Create(joined)
	if err != nil {
		return nil, errors.Wrap(err, "could not open file")
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		return nil, errors.Wrap(err, "write to disk failed")
	}
	return &StoreResult{}, nil
}

// Delete deletes the file at the specified key
func (fs *Memory) Delete(key string) error {
	joined := filepath.Join(fs.root, key)
	return errors.Wrap(fs.fs.Remove(joined), "could not remove file")
}

// PresignedURL returns a URL that provides access to a file for 15 mintes.
func (fs *Memory) PresignedURL(key, contentType string) (string, error) {
	values := url.Values{}
	values.Add("contentType", contentType)
	url := fs.webRoot + "/" + key + "?" + values.Encode()
	return url, nil
}

// Fetch retrieves a copy of a file and stores it in a tempfile. The path to this
// file is returned.
//
// It is the caller's responsibility to delete the tempfile.
func (fs *Memory) Fetch(key string) (io.ReadCloser, error) {
	sourcePath := filepath.Join(fs.root, key)
	f, err := fs.fs.Open(sourcePath)
	return f, errors.Wrap(err, "could not open file")
}

// FileSystem returns the underlying afero filesystem
func (fs *Memory) FileSystem() *afero.Afero {
	return fs.fs
}

// NewMemoryHandler returns an Handler that adds a Content-Type header so that
// files are handled properly by the browser.
func NewMemoryHandler(root string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.URL.Query().Get("contentType")
		if contentType != "" {
			w.Header().Add("Content-Type", contentType)
		}

		input := filepath.Join(root, filepath.FromSlash(path.Clean("/"+r.URL.Path)))
		http.ServeFile(w, r, input)
	})
}
