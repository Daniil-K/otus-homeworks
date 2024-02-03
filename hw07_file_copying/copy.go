package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// Copy копирует исходный файл частично или полностью.
func Copy(fromPath, toPath string, limit, offset int64) error {
	// Открытие исходного файла для чтения.
	sourceFile, err := os.Open(fromPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer sourceFile.Close()

	// Проверка размера сдвига и сдвиг в исходном файле.
	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	sourceFileSize := sourceFileInfo.Size()
	if sourceFileSize <= offset {
		return ErrOffsetExceedsFileSize
	}
	if offset > int64(io.SeekStart) {
		_, err = sourceFile.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	// Обертка для вывода прогресса копирования.
	barSize := sourceFileSize - offset
	if limit != 0 && limit < barSize {
		barSize = limit
	}
	bar := pb.New64(barSize)
	sourceReader := bar.NewProxyReader(sourceFile)

	// Копирование исходного файла.
	targetFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if barSize != 0 {
		_, err = io.CopyN(targetFile, sourceReader, barSize)
	} else {
		_, err = io.Copy(targetFile, sourceReader)
	}

	bar.Finish()

	if err != nil {
		return err
	}

	return nil
}
