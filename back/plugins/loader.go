package plugins

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"plugin"

	"github.com/PoCFrance/e/myutil"
	"github.com/robpike/filter"
)

// Package records all the necessary to load runtime package
type Package struct {
	Plug *plugin.Plugin
	IO   myutil.IOMod
	Name string
}

var loadedPackage = map[string]([]Package){}

// GetPackage returns the list of
func GetPackage(locale string) []Package {
	return loadedPackage[locale]
}

func loader(f os.FileInfo, local string, pchan chan interface{}) {
	name := f.Name()
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", name+".so")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "./package/" + name
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
		pchan <- err
		return
	}
	p, err := plugin.Open("package/" + name + "/" + name + ".so")
	if err != nil {
		fmt.Printf("plugin.Open(%s) failed with %s\n", name, err)
		pchan <- err
		return
	}

	io, err := myutil.SerializeIO(name, local)
	if err != nil {
		fmt.Println("SerializingIO failed : ", err)
		pchan <- err
		return
	}

	pchan <- Package{
		Plug: p,
		IO:   io,
		Name: name}
}

// LoadPlugins load all the packages located in the "./pachage/" directory
func LoadPlugins(locale string) []Package {

	files, err := ioutil.ReadDir("./package")
	if err != nil {
		log.Fatal(err)
	}

	dirs := filter.Choose(files, func(f os.FileInfo) bool {
		return f.IsDir()
	}).([]os.FileInfo)

	packages := make([]Package, len(dirs))
	packchan := make(chan interface{})

	for _, f := range dirs {
		go loader(f, locale, packchan)
	}

	for range dirs {
		res := <-packchan
		switch res.(type) {
		case error:
			// fmt.Println(res.(error))
		case Package:
			packages = append(packages, res.(Package))
		}
	}
	loadedPackage[locale] = packages[:]
	return GetPackage(locale)
}
