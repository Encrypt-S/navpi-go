package fs

import (
	"archive/zip"
	"compress/gzip"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

// DownloadFile will download a url to a local file.
// efficiently writes as it downloads instead of loading file in memory
// io.TeeReader is passed into Copy() to report progress on the download
func DownloadFile(filepath string, url string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}

// DownloadExtract sets up and runs the functions
// needed for downloading and extracting of assets
func DownloadExtract(url string, assetName string) error {

	path, err := GetCurrentPath()

	downloadLocation := path + "/" + assetName

	extractPath := path + "/lib"

	Download(url, downloadLocation)

	Extract(assetName, downloadLocation, extractPath)

	if err != nil {
		return err
	}

	return nil

}

// Download performs file download of the given url
// this method provides no feedback to the system
func Download(url string, downloadTofileName string) {

	log.Println("Downloading", url)
	log.Println("Destination", downloadTofileName)
	log.Println("This could take a few mins :)")

	output, err := os.Create(downloadTofileName)

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}

	log.Println(n, "bytes downloaded")

}

func Extract(assetName string, downloadLocation string, extractPath string) {

	switch filepath.Ext(assetName) {
	case ".zip":
		Unzip(downloadLocation, extractPath)
	case ".gz":
		Ungzip(downloadLocation, extractPath)
	}

	log.Println("File extracted to " + extractPath)

}

func Unzip(src, dest string) error {

	log.Println("Unzip the zip file from " + dest)

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Ungzip uncompresses the given tar.gz file
func Ungzip(gzStream, dest string) error {

	log.Println("Ungzip the tar.gz from " + dest)

	gzReader, err := os.Open(gzStream)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	gzArchive, err := gzip.NewReader(gzReader)
	if err != nil {
		return err
	}
	defer gzArchive.Close()

	dest = filepath.Join(dest, gzArchive.Name)

	writer, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, gzArchive)

	return err
}

// GetCurrentPath gets the path of the go app
func GetCurrentPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	return exPath, nil
}

// Exists reports if the named file or directory exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
