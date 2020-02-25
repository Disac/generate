package main

import "git-admin.inyuapp.com/feature/generate/gen"

func main() {
	gen.NewGen(
		"git-admin.inyuapp.com/feature/generate/aaa",
		gen.ConfigPath("/etc/app", "toml"),
	).Config()
}
