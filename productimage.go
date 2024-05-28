package olahproductimage

import (
	"path/filepath"

	"github.com/agungdhewe/dwlog"
	"github.com/agungdhewe/dwpath"
)

type ProductImage struct {
	SourceDir string
	TargetDir string
}

func New(source string, target string) *ProductImage {
	pi := &ProductImage{
		SourceDir: source,
		TargetDir: target,
	}
	return pi
}

func (pi *ProductImage) Execute() error {
	var exists bool
	var err error

	// cek apakah direktori source exist
	exists, err = dwpath.IsDirectoryExists(pi.SourceDir)
	if !exists {
		dwlog.Error("direktori %s tidak ditemukan", pi.SourceDir)
		return err
	}

	// cek apakah direktori source target
	exists, err = dwpath.IsDirectoryExists(pi.TargetDir)
	if !exists {
		dwlog.Error("direktori %stidak ditemukan", pi.TargetDir)
		return err
	}

	// baca direktori sumber
	searchpattern := filepath.Join(pi.SourceDir, "*.jpg")
	dwlog.Info("Ambil file %s", searchpattern)
	entries, err := filepath.Glob(searchpattern)
	if err != nil {
		dwlog.Error("tidak bisa membaca file(s) %s", searchpattern)
		return err
	}

	for _, entry := range entries {
		err = convertToDir(entry, pi.TargetDir)
		if err != nil {
			dwlog.Error("error saat convert data %s", entry)
			return err
		}

	}

	return nil
}
