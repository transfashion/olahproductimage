package olahproductimage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/agungdhewe/dwlog"
	"github.com/agungdhewe/dwpath"
)

func convertToDir(path string, targetdir string) error {
	fmt.Printf("converting %s\r\n", path)

	filename := filepath.Base(path)
	fileext := filepath.Ext(path)
	proddatatext := strings.TrimSuffix(filename, fileext)
	proddata := strings.Split(proddatatext, "~")

	var err error
	var art string
	var col string
	var seq string
	var targetbasename string

	if len(proddata) == 3 {
		art = proddata[0]
		col = proddata[1]
		seq = proddata[2]
		targetbasename = createSeqBaseName(seq, fileext)
		err = convertWithColor(path, art, col, seq, targetbasename, targetdir)
		if err != nil {
			dwlog.Error("error saat convert (demgan color) %s", filename)
			return err
		}
	} else if len(proddata) == 2 {
		art = proddata[0]
		seq = proddata[1]
		targetbasename = createSeqBaseName(seq, fileext)
		err = convertWithoutColor(path, art, seq, targetbasename, targetdir)
		if err != nil {
			dwlog.Error("error saat convert (tanpa color) %s", filename)
			return err
		}
	} else {
		err = fmt.Errorf("format nama file %s tidak diperbolehkan", filename)
		dwlog.Error(err.Error())
		return err
	}

	return nil
}

func convertWithColor(path string, art string, col string, seq string, targetbasename string, targetdir string) error {
	var err error

	// cek apakah direktori art ada di target dir
	artdir := filepath.Join(targetdir, art)
	err = createArtDir(artdir)
	if err != nil {
		return err
	}

	// cek apakah directori col ada di target dir
	coldir := filepath.Join(artdir, col)
	err = createColDir(coldir)
	if err != nil {
		return err
	}

	// copy file ke direktori target dengan padding 0

	return nil
}

func convertWithoutColor(path string, art string, seq string, targetbasename strring, targetdir string) error {
	var err error

	// cek apakah direktori art ada di target dir
	artdir := filepath.Join(targetdir, art)
	err = createArtDir(artdir)
	if err != nil {
		return err
	}

	return nil
}

func createArtDir(artdir string) error {
	exists, _ := dwpath.IsDirectoryExists(artdir)
	if !exists {
		// direktori belum ada, buat baru
		dwlog.Info("buat direktori artikel %s", artdir)
		err := os.Mkdir(artdir, os.ModePerm)
		if err != nil {
			dwlog.Error("tidak dapat membuat direktori %s", artdir)
			return err
		}
	}

	return nil
}

func createColDir(coldir string) error {
	exists, _ := dwpath.IsDirectoryExists(coldir)
	if !exists {
		// direktori belum ada, buat baru
		dwlog.Info("buat direktori color %s", coldir)
		err := os.Mkdir(coldir, os.ModePerm)
		if err != nil {
			dwlog.Error("tidak dapat membuat direktori %s", coldir)
			return err
		}
	}

	return nil
}

func createSeqBaseName(seq string, fileext string) string {
	return fmt.Sprintf("%s.%s", seq, fileext)
}
