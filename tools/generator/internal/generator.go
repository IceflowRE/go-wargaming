package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Generator struct {
	moduleName string
	outputPath string
	verbose    bool
}

func NewGenerator(verbose bool) (*Generator, error) {
	return &Generator{
		outputPath: "./wargaming/",
		verbose:    verbose,
	}, nil
}

func (gen *Generator) Generate() (err error) {
	methodsDocs, realmDocs, err := getWgMethodsDoc()
	if err != nil {
		return err
	}
	var services []serviceData
	sectionsMap := map[string]*sectionData{}
	realmIdcs := map[string]string{}
	// method id to deprecated
	methodIds := map[string]bool{}

	err = gen.removeFiles(methodsDocs)
	if err != nil {
		return err
	}

	maxMethods := methodsCount(methodsDocs)
	processedMethods := 0
	for _, game := range methodsDocs {
		if len(game.Sections) == 0 {
			continue
		}
		// collect service
		services = append(services, serviceData{Name: snakeToCamel(game.Slug), LongName: game.Name})
		// init section
		sectionsMap[snakeToCamel(game.Slug)] = &sectionData{Name: snakeToCamel(game.Slug), LongName: game.Name}

		err = os.MkdirAll(filepath.Join(gen.outputPath, game.Slug), 0664)
		if err != nil {
			return err
		}
		for _, section := range game.Sections {
			for _, method := range section.Methods {
				processedMethods += 1
				gen.printf("(%03d/%03d) Generating %s\n", processedMethods, maxMethods, method.Key)
				doc, err := getWgMethodDoc(method.Key, "all")
				if err != nil {
					return err
				}
				// collect sections
				sec, _ := sectionsMap[snakeToCamel(game.Slug)]
				if sec.ApiUrlFormat != "" && sec.ApiUrlFormat != doc.ApiUrl {
					return fmt.Errorf("%s has multiple API URLs!? (%s | %s)", game.Slug, sec.ApiUrlFormat, doc.ApiUrl)
				}
				sec.ApiUrlFormat = doc.ApiUrl
				// collect realm indices
				for _, realmIdx := range doc.AvailableDisplayIndices {
					if _, ok := realmIdcs[realmIdx]; !ok {
						realmIdcs[realmIdx] = method.Key
					}
				}
				// collect method id
				methodIds[method.Key] = method.Deprecated

				ep, err := newEndpointFromDoc(doc, game.Slug, method.Key)
				if err != nil {
					return err
				}
				gen.print("    patching")
				patchEndpoint(ep)
				gen.print("    create method file")
				err = genMethod(ep, gen.outputPath)
				if err != nil {
					return err
				}
				if ep.OptionsType != nil || ep.MetaType != nil || ep.DataType != nil && ep.DataType.IsStruct() {
					gen.print("    create struct file")
					err = genStructs(ep, gen.outputPath)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	gen.print("Generating services.go")
	sort.Slice(services, func(i, k int) bool {
		return services[i].Name < services[k].Name
	})
	err = genServices(services, gen.outputPath)
	if err != nil {
		return err
	}

	gen.print("Generating sections.go")
	sections := make([]*sectionData, 0, len(sectionsMap))
	for _, sec := range sectionsMap {
		sections = append(sections, sec)
	}
	sort.Slice(sections, func(i, k int) bool {
		return sections[i].Name < sections[k].Name
	})
	err = genSections(sections, gen.outputPath)
	if err != nil {
		return err
	}

	gen.print("Generating realms.go")
	realms, err := collectRealms(realmIdcs, realmDocs)
	if err != nil {
		return err
	}
	err = genRealms(realms, gen.outputPath)
	if err != nil {
		return err
	}

	gen.print("Check client tests")
	err = gen.checkClientTest(methodIds)
	if err != nil {
		return err
	}

	gen.print("Check generator")
	gen.checkGeneratorOldIds(methodIds)
	return nil
}

func collectRealms(realmIdcs map[string]string, realmDocs []*realmDoc) ([]*realmData, error) {
	var realms []*realmData
	for idx, key := range realmIdcs {
		doc, err := getWgMethodDoc(key, idx)
		if err != nil {
			return nil, err
		}
		realm := &realmData{
			Tld:   rxApiUrl.FindAllStringSubmatch(doc.ApiUrl, 1)[0][2],
			Index: strings.ToLower(idx),
		}
		for _, doc := range realmDocs {
			if doc.Index == idx {
				realm.LongName = strings.Title(doc.Name) // lint:ignore
			}
		}
		if realm.LongName == "" {
			return nil, fmt.Errorf("realm long name not found (%s)", idx)
		}
		realms = append(realms, realm)
	}
	sort.Slice(realms, func(i, k int) bool {
		return realms[i].Index < realms[k].Index
	})
	return realms, nil
}

func methodsCount(methodsDocs []*wgMethodsDoc) int {
	methods := 0
	for _, game := range methodsDocs {
		for _, section := range game.Sections {
			methods += len(section.Methods)
		}
	}
	return methods
}

func (gen *Generator) removeFiles(methodsDocs []*wgMethodsDoc) error {
	gen.print("Removing files:")
	// prefix of files/ directories to remove
	prefixes := []string{"services"}
	for _, game := range methodsDocs {
		prefixes = append(prefixes, game.Slug)
	}

	dirs, err := os.ReadDir(gen.outputPath)
	if err != nil {
		return err
	}
DIRLOOP:
	for _, dirEntry := range dirs {
		// ignore wgnTime directory
		if dirEntry.Name() == "wgnTime" {
			continue
		}
		for _, prefix := range prefixes {
			if strings.HasPrefix(dirEntry.Name(), prefix) {
				gen.print("    -", gen.outputPath+dirEntry.Name())
				err = os.RemoveAll(gen.outputPath + dirEntry.Name())
				if err != nil {
					return err
				}
				continue DIRLOOP
			}
		}
	}
	return err
}

func (gen *Generator) printf(format string, a ...any) {
	if gen.verbose {
		fmt.Printf(format, a...)
	}
}

func (gen *Generator) print(a ...any) {
	if gen.verbose {
		fmt.Println(a...)
	}
}
