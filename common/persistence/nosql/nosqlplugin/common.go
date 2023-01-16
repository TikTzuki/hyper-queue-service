package nosqlplugin

import (
	"os"
	"strings"
)

func getHyperQPackageDir() (string, error) {
	hyperQPackageDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cadenceIndex := strings.LastIndex(hyperQPackageDir, "/hyper-q/")
	hyperQPackageDir = hyperQPackageDir[:cadenceIndex+len("/hyper-q/")]
	return hyperQPackageDir, err
}

func GetDefaultTestSchemaDir(testSchemaRelativePath string) (string, error) {
	cadencePackageDir, err := getHyperQPackageDir()
	if err != nil {
		return "", err
	}
	return cadencePackageDir + testSchemaRelativePath, nil
}
