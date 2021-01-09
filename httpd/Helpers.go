package httpd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UntarUrlToSingleFile(url, outputfile string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("Error in getting S3-Url: " + url)
		return err
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)

	// create output file with original file mode
	file, err := os.OpenFile(outputfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	// untar each segment
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// determine proper file path info to check if is dir
		finfo := hdr.FileInfo()

		if finfo.Mode().IsDir() {
			continue
		}

		n, cpErr := io.Copy(file, tr)
		if cpErr != nil {
			return cpErr
		}
		if n != finfo.Size() {
			return fmt.Errorf("unexpected bytes written: wrote %d, want %d", n, finfo.Size())
		}
	}
	return nil
}
