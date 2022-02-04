package utils

import (
	"fmt"

	"github.com/vektra/gitreader"
)

func ReadFileFromTemplatesInGithub(filename string) ([]byte, error) {
	repo, err := gitreader.OpenRepo("github.com/nirsht/kube-sample")
	if err != nil {
		panic(err)
	}

	blob, err := repo.CatFile("HEAD", "kubernetes_supported_resources/template/"+filename)
	if err != nil {
		panic(err)
	}

	// WARNING: use Blob as an io.Reader instead if you can!
	bytes, err := blob.Bytes()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", bytes)

	repo.Close()

	return bytes, nil
}
