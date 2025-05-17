package store

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/pkg/utils"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func (c Controller) AddProduct(ctx context.Context, product dto.Product, file *multipart.FileHeader) error {
	imgUrl, err := c.saveFile(file)
	if err != nil {
		return fmt.Errorf("error with file saving: %w", err)
	}

	product.ImageURL = imgUrl

	err = c.r.CreateProduct(ctx, product)
	if err != nil {
		if err1 := os.Remove(imgUrl); err != nil {
			return fmt.Errorf("error with file deleting: %w caused after %w", err, err1)
		}
		return fmt.Errorf("error with file creating: %w", err)
	}

	return nil
}

func (c Controller) saveFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	fPath, err := utils.GetFilenameWithPrefix("frontend/img/store/products/", file.Filename)
	if err != nil {
		return "", err
	}

	frontPath, _ := strings.CutPrefix(fPath, "frontend/")

	f, err := os.Create(frontPath)
	if _, err := io.Copy(f, src); err != nil {
		return "", err
	}

	return fPath, nil
}
