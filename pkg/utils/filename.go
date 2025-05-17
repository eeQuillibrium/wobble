package utils

import uuid2 "github.com/google/uuid"

func GetFilenameWithPrefix(prefix string, fName string) (string, error) {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return "", err
	}

	return prefix + uuid.String() + fName, nil
}
