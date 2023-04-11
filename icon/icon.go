package icon

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
	"gopkg.in/yaml.v2"
)

type Abbreviation struct {
	Name         string
	Abbreviation string
}
type Abbreviations map[string]string

func LoadAbbreviations(yamlPath string) Abbreviations {
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Fatal(err)
	}
	_abbrs := []Abbreviation{}
	if err = yaml.Unmarshal(yamlFile, &_abbrs); err != nil {
		log.Fatal(err)
	}
	abbrs := Abbreviations{}
	for _, _abbr := range _abbrs {
		abbrs[_abbr.Name] = _abbr.Abbreviation
	}
	return abbrs
}

func FindFilePaths(dirPattern, prefix, suffix string) []string {
	pattern := filepath.Join(dirPattern, prefix+"*"+suffix)
	paths, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

func LoadArchitectureIcons(
	wf *aw.Workflow,
	dirPattern,
	prefix,
	svgSuffix,
	pngSuffix string,
	abbrs Abbreviations,
) {
	svgPaths := FindFilePaths(dirPattern, prefix, svgSuffix)
	icons := map[string]bool{}
	for _, svgPath := range svgPaths {
		// name
		name := filepath.Base(svgPath)
		name = strings.Replace(name, prefix, "", 1)
		name = strings.Replace(name, svgSuffix, "", 1)
		name = strings.ReplaceAll(name, "-", " ")
		if _, ok := icons[name]; ok {
			continue
		}

		// match words
		match := name
		// abbreviation
		abbr := name
		if val, ok := abbrs[name]; ok {
			abbr = val
			match = abbr + " " + match
		}
		// png path
		pngPath := strings.Replace(svgPath, svgSuffix, pngSuffix, 1)

		// filter match words
		match = strings.ToLower(match)
		match = strings.ReplaceAll(match, "amazon ", "")
		match = strings.ReplaceAll(match, "aws ", "")
		match = strings.ReplaceAll(match, "(", "")
		match = strings.ReplaceAll(match, ")", "")

		// add to alfred
		wf.
			NewItem(abbr).
			Valid(true).
			Subtitle(name).
			Var("action", "run-script").
			Match(match).
			UID(name).
			Arg(svgPath).
			Icon(&aw.Icon{Value: pngPath})

		icons[name] = true
	}
}

func LoadResourceIcons(
	wf *aw.Workflow,
	dirPattern,
	prefix,
	svgSuffix,
	pngSuffix string,
	info string,
	abbrs Abbreviations,
) {
	svgPaths := FindFilePaths(dirPattern, prefix, svgSuffix)
	for _, svgPath := range svgPaths {
		// name
		name := filepath.Base(svgPath)
		name = strings.Replace(name, prefix, "", 1)
		name = strings.Replace(name, svgSuffix, "", 1)
		name = strings.ReplaceAll(name, "-", " ")
		abbr := name
		if strings.Contains(name, "_") {
			names := strings.Split(name, "_")
			name = strings.Join(names, " - ")
			if val, ok := abbrs[names[0]]; ok {
				names[0] = val
				abbr = strings.Join(names, " - ")
			} else {
				abbr = name
			}
		}
		// match words
		match := name
		if val, ok := abbrs[name]; ok {
			abbr = val
			match = abbr + " " + abbr + " " + abbr + " " + match
		}
		// png path
		pngPath := strings.Replace(svgPath, svgSuffix, pngSuffix, 1)

		// filter match words
		match = strings.ToLower(match)
		match = strings.ReplaceAll(match, "amazon ", "")
		match = strings.ReplaceAll(match, "aws ", "")
		match = strings.ReplaceAll(match, "(", "")
		match = strings.ReplaceAll(match, ")", "")

		// add to alfred
		wf.
			NewItem(abbr).
			Valid(true).
			Subtitle(name).
			Var("action", "run-script").
			Match(match).
			UID(name).
			Arg(svgPath).
			Icon(&aw.Icon{Value: pngPath})
	}
}
