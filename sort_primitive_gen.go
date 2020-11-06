// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("sort_primitive.go")
	if err != nil {
		log.Panic(err)
	}

	codeTemplate := string(data)

	data, err = ioutil.ReadFile("sort_primitive_test.go")
	testCodeTemplate := string(data)

	allPrimitives := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"string",
	}

	const header = "// Code generated from sort_primitive.go using sort_primitive_gen.go; DO NOT EDIT.\n"

	// generate code files with specialized types
	for _, t := range allPrimitives {
		var specialized = header + codeTemplate
		specialized = strings.ReplaceAll(specialized, "// +build ignore\n", "")
		specialized = strings.ReplaceAll(specialized, "//go:generate go run sort_primitive_gen.go\n", "")
		specialized = strings.ReplaceAll(specialized, "package dualpivotsort", fmt.Sprintf("package sort_%s", t))
		specialized = strings.ReplaceAll(specialized, "type primitive = int64", fmt.Sprintf("type primitive = %s", t))

		output, err := format.Source([]byte(specialized))
		if err != nil {
			log.Panic(err)
		}

		dir := fmt.Sprintf("internal/sort_%s", t)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s/sort_%s.go", dir, t), output, 0644)
		if err != nil {
			log.Panic(err)
		}

		if t == "string" { // test for string type needs a different rand func, skip generation for it
			continue
		}

		// generate test code
		specialized = header + testCodeTemplate
		specialized = strings.ReplaceAll(specialized, "// +build ignore\n", "")
		specialized = strings.ReplaceAll(specialized, "//go:generate go run sort_primitive_gen.go\n", "")
		specialized = strings.ReplaceAll(specialized, "package dualpivotsort", fmt.Sprintf("package sort_%s", t))
		specialized = strings.ReplaceAll(specialized, "type primitive = int64", fmt.Sprintf("type primitive = %s", t))
		specialized = strings.ReplaceAll(specialized, "TestDpsPrimitive", fmt.Sprintf("TestDps%ss", strings.Title(t)))

		output, err = format.Source([]byte(specialized))
		if err != nil {
			log.Panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s/sort_%s_test.go", dir, t), output, 0644)
		if err != nil {
			log.Panic(err)
		}
	}

	const exportTemplate = header + `
package dualpivotsort

import (
    $IMPORTS
)

$EXPORTS
`

	var imports []string
	var exports []string

	var appendImport = func(typeName string) {
		imports = append(imports, fmt.Sprintf(`"github.com/openaphid/dualpivotsort/internal/sort_%s"`, typeName))
	}

	var appendExport = func(alias, typeName string, plural bool, hasSorted bool) {
		exports = append(exports, fmt.Sprintf("// %s", alias))
		if plural {
			exports = append(exports, fmt.Sprintf("var %ss = sort_%s.Sort", strings.Title(alias), typeName))
			if hasSorted {
				exports = append(exports, fmt.Sprintf("var %ssAreSorted = sort_%s.IsSorted", strings.Title(alias), typeName))
			}
		} else {
			exports = append(exports, fmt.Sprintf("var %s = sort_%s.Sort", strings.Title(alias), typeName))
			if hasSorted {
				exports = append(exports, fmt.Sprintf("var %sAreSorted = sort_%s.IsSorted", strings.Title(alias), typeName))
			}
		}
		exports = append(exports, "")
	}

	for _, t := range allPrimitives {
		appendImport(t)
		appendExport(t, t, true, true)
	}
	appendExport("byte", "uint8", true, true)
	appendExport("rune", "int32", true, true)

	for _, t := range []string{"slice_dps", "slice_tim"} {
		appendImport(t)
	}
	appendExport("Slice", "slice_dps", false, true)
	appendExport("SliceStable", "slice_tim", false, false)

	var exportCode = exportTemplate
	exportCode = strings.ReplaceAll(exportCode, "$IMPORTS", strings.Join(imports, "\n"))
	exportCode = strings.ReplaceAll(exportCode, "$EXPORTS", strings.Join(exports, "\n"))

	output, err := format.Source([]byte(exportCode))
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile("sort_export.go", output, 0644)
	if err != nil {
		log.Panic(err)
	}
}
