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
		specialized = strings.ReplaceAll(specialized, "package jsort", fmt.Sprintf("package sort_%s", t))
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
		specialized = strings.ReplaceAll(specialized, "package jsort", fmt.Sprintf("package sort_%s", t))
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
package jsort

import (
    $IMPORTS
)

$EXPORTS
`

	var imports []string
	var exports []string

	var appendImport = func(pkg string) {
		imports = append(imports, fmt.Sprintf(`"github.com/openaphid/jsort/internal/%s"`, pkg))
	}

	var appendLine = func(line string) {
		exports = append(exports, line)
	}

	var appendSortExport = func(alias, pkg string) {
		appendLine(fmt.Sprintf("// %s", alias))
		appendLine(fmt.Sprintf("var %ss = %s.Sort", strings.Title(alias), pkg))
		appendLine(fmt.Sprintf("var %ssAreSorted = %s.IsSorted", strings.Title(alias), pkg))

		appendLine("")
	}

	var appendTypeExport = func(alias, pkg, typeName string) {
		appendLine(fmt.Sprintf("type %s = %s.%s", strings.Title(alias), pkg, typeName))
	}

	_ = appendTypeExport

	var appendFuncExport = func(alias, pkg, typeName string) {
		appendLine(fmt.Sprintf("var %s = %s.%s", strings.Title(alias), pkg, typeName))
	}

	_ = appendFuncExport

	for _, t := range allPrimitives {
		pkg := fmt.Sprintf("sort_%s", t)
		appendImport(pkg)
		appendSortExport(t, pkg)
	}
	appendSortExport("byte", "sort_uint8")
	appendSortExport("rune", "sort_int32")

	// `sort_slice_dps_ts`, `sort_slice_tim_ts`, `sort_slice_dps_go2`, and `sort_slice_tim_go2` are not exported
	appendImport("sort_slice_tim_interface")
	appendLine("// The following APIs are compatible with the ones in the built-in `sort` package")
	appendLine("// One difference is that all sort functions are stable by using timsort")
	appendFuncExport("Sort", "sort_slice_tim_interface", "Sort")
	appendFuncExport("Stable", "sort_slice_tim_interface", "Sort")
	appendFuncExport("Slice", "sort_slice_tim_interface", "Slice")
	appendFuncExport("SliceStable", "sort_slice_tim_interface", "Slice")
	appendFuncExport("SliceIsSorted", "sort_slice_tim_interface", "SliceIsSorted")

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
