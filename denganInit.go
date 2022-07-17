package main

import (
	"io/ioutil"
	"os"
)

func denganInit() {
	if len(os.Args) == 2 && os.Args[1] == "init" {

		_, errAndroJson := os.Stat("andro.json")
		if errAndroJson != nil {
			ioutil.WriteFile("andro.json", []byte(template()), 0777)
		}

		_, errGitignore := os.Stat(".gitignore")
		if errGitignore != nil {
			ioutil.WriteFile(".gitignore", []byte(""), 0777)
		}
		if errGitignore == nil {
			// 	let isiGitignore: string | string[] =
			// 	readFileSync(".gitignore").toString();
			// isiGitignore = isiGitignore.split("\n").filter((x) => x);

			// if (!isiGitignore.includes("andro")) {
			// 	isiGitignore.push("andro");
			// 	writeFileSync(".gitignore", isiGitignore.join("\n"));
			// }
			gitignoreFile, errGitignoreFile := ioutil.ReadFile(".gitignore")
			cek(errGitignoreFile)
			stringGitignore := string(gitignoreFile)

		}

	}
}
