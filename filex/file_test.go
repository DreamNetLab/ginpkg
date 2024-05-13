package filex

import (
	"os"
	"path/filepath"
	"testing"
)

// 测试CheckNotExist函数
func TestCheckNotExist(t *testing.T) {
	tempDir := t.TempDir()
	existingFile := filepath.Join(tempDir, "existing.txt")
	_, err := os.Create(existingFile)
	if err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"存在的文件", existingFile, false},
		{"不存在的文件", filepath.Join(tempDir, "non-existing.txt"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckNotExist(tt.src); got != tt.want {
				t.Errorf("CheckNotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 测试MkDir函数
func TestMkDir(t *testing.T) {
	tempDir := t.TempDir()
	targetDir := filepath.Join(tempDir, "newdir")

	err := MkDir(targetDir)
	if err != nil {
		t.Errorf("创建目录失败: %v", err)
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		t.Errorf("目录未创建")
	}
}

// 测试CheckPermission函数
func TestCheckPermission(t *testing.T) {
	// 由于权限问题在不同操作系统和环境下表现不同，这里仅提供基本思路
	t.Skip("权限测试在不同平台上结果可能不同，此测试被跳过")
}

// 测试IsNotExistMkdir函数
func TestIsNotExistMkdir(t *testing.T) {
	tempDir := t.TempDir()
	targetDir := filepath.Join(tempDir, "mkdirtest")

	err := IsNotExistMkdir(targetDir)
	if err != nil {
		t.Errorf("IsNotExistMkdir() error = %v", err)
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		t.Errorf("目录未创建")
	}
}

// 测试Open函数
func TestOpen(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "testfile.txt")

	// 测试文件创建
	f, err := Open(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}
	f.Close()

	// 测试文件是否真实存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("文件未创建")
	}
}

// 测试MustOpen函数
func TestMustOpen(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "///mustopen.txt"
	filePath := filepath.Join(tempDir, fileName)

	// 期望成功打开文件
	f, err := MustOpen(fileName, tempDir)
	if err != nil {
		t.Fatalf("MustOpen() error = %v", err)
	}
	f.Close()

	dir, _ := os.Getwd()
	// 检查文件是否存在
	if _, err := os.Stat(filepath.Join(dir, filePath)); os.IsNotExist(err) {
		t.Errorf("文件未创建")
	}
}
