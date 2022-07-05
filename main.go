package main

import (
	"compareVersions/util"
	"fmt"
)

func main() {
	versions := []struct{ a, b string }{
		{"0.1", "1.1"},
		{"1.0.1", "1"},
		{"7.5.2.4", "7.5.3"},
		{"1.01", "1.001"},
		{"1.0", "1.0.0"},
	}

	for _, version := range versions {
		v := util.CompareVersion(version.a, version.b)
		fmt.Println(v)
	}
}
