package plugins

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"plugin"

	"github.com/robpike/filter"
)

var loadedPackage = map[string]([]Package){}

// GetPackage returns the list of
func GetPackage(locale string) []Package {
	return loadedPackage[locale]
}

func NewPackage(name string, plug *plugin.Plugin) Package {
	return Package{
		Plug:    plug,
		Modules: make(map[string]Module),
		Name:    name,
	}
}
func loader(f os.FileInfo, locale string, pchan chan interface{}) {
	name := f.Name()
	cr := exec.Command("go", "build", "-buildmode=plugin", "-o", name+".so")
	cr.Stdout = os.Stdout
	cr.Stderr = os.Stderr
	cr.Dir = "./package/" + name
	err := cr.Run()
	if err != nil {
		fmt.Printf("cr.Run() failed with %s\n", err)
		pchan <- err
		return
	}
	p, err := plugin.Open("package/" + name + "/" + name + ".so")
	if err != nil {
		fmt.Printf("plugin.Open(%s) failed with %s\n", name, err)
		pchan <- err
		return
	}

	pack := NewPackage(name, p)

	if err = pack.loadModules(locale); err != nil {
		fmt.Println(err)
		pchan <- err
		return
	}
	pchan <- pack
}

// LoadPackage load all the packages located in the "./pachage/" directory
func LoadPackage(locale string) []Package {

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
			//fmt.Println("loadModules failed : ", err)
		case Package:
			packages = append(packages, res.(Package))
		}
	}
	loadedPackage[locale] = packages[:]
	return GetPackage(locale)
}
