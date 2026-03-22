package photo

import (
	"fmt"
	"os"
)

type Storage struct {
	bucket    string
	accessKey string
	secretKey string
	accountID string
	publicURL string
}

func NewStorage() *Storage {
	return &Storage{
		accountID: os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		accessKey: os.Getenv("CLOUDFLARE_R2_ACCESS_KEY"),
		secretKey: os.Getenv("CLOUDFLARE_R2_SECRET_KEY"),
		bucket:    os.Getenv("CLOUDFLARE_R2_BUCKET"),
		publicURL: os.Getenv("CLOUDFLARE_R2_PUBLIC_URL"),
	}
}

func (s *Storage) UploadPhoto(userID string, fileBytes []byte, filename string) (string, error) {
	// TODO: implement real Cloudflare R2 upload when credentials are ready
	// R2 is S3-compatible, use aws-sdk-go-v2 with custom endpoint:
	//    https://{accountID}.r2.cloudflarestorage.com
	return fmt.Sprintf("%s/%s/%s", s.publicURL, userID, filename), nil
}
