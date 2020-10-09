package fileutil

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//make directory
func MakeDirs(dir string) error {
	return os.MkdirAll(dir, 0777)
}

//fetch annotation
func ReadAllOfFile(fname string) (string, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}
	return string(b), err
}

//save text
func SaveText(fname, data string) error {
	w, err := os.Create(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer w.Close()

	_, err = w.WriteString(data)
	return err
}

//copy file
func CopyFile(dest, src string) error {
	r, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	w, err := os.Create(dest)
	if err != nil {
		log.Fatalln(err)
	}
	defer w.Close()

	//do the actual work
	_, err = io.Copy(w, r)
	return err
}
