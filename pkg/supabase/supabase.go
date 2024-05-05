package supabase

import (
	"io"
	"log"
	"mime/multipart"
	"path/filepath"

	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseStorageItf interface {
	Upload(bucket string, file *multipart.FileHeader) (string, error)
}

type SupabaseStorage struct {
	client *storage_go.Client
}

func NewSupabaseStorage(client *storage_go.Client) SupabaseStorageItf {
	return &SupabaseStorage{client: client}
}

func (s SupabaseStorage) Upload(bucket string, file *multipart.FileHeader) (string, error) {
	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	fileReader := io.ReadCloser(fileData)
	defer fileReader.Close()

	fileName := filepath.Base(file.Filename)
    log.Println(fileName)
	relativePath := fileName

	result, err := s.client.UploadFile(bucket, relativePath, fileReader,
		storage_go.FileOptions{ContentType: func() *string { s := "image/jpg"; return &s }(),
			Upsert: func() *bool { b := true; return &b }()})
	if err != nil {
		return "", err 
	}
    log.Println("sukses boskuh")
	return result.Key, nil
}