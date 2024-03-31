package filex

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return errors.Is(err, fs.ErrNotExist)
}

func MkDir(src string) error {
	return os.MkdirAll(src, os.ModePerm)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return errors.Is(err, fs.ErrPermission)
}

func IsNotExistMkdir(src string) error {
	if notExist := CheckNotExist(src); notExist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %w", err)
	}

	src := filepath.Join(dir, filePath)

	permission := CheckPermission(src)
	if permission {
		return nil, fmt.Errorf("file.CheckPermission permission denied: %s", src)
	}

	err = IsNotExistMkdir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkdir src: %s, err: %w", src, err)
	}

	f, err := Open(filepath.Join(src, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %w", err)
	}

	return f, nil
}
