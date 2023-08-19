package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/vinceanalytics/vince/tools"
	_ "k8s.io/code-generator"
	_ "sigs.k8s.io/controller-tools/pkg/crd"
)

const (
	rootPackage = "github.com/vinceanalytics/vince/internal/v8s"
)

var (
	GENERATE_SCRIPT string
	root            string
)

func main() {
	println("### Generating k8s client for crds ###")
	root = tools.RootVince()
	println(">>> root:", root)
	make()
	buildCrd()
}

func make() {
	CODEGEN_PKG := tools.Root("k8s.io/code-generator")
	println(">>> using codegen: ", CODEGEN_PKG)
	GENERATE_SCRIPT = filepath.Join(CODEGEN_PKG, "generate-groups.sh")
	tools.ExecPlain("chmod", "+x", GENERATE_SCRIPT)
	println("##### Generating site client ######")
	generate("vince", "v1alpha1")
}

func generate(resource string, versions ...string) {
	dir, err := os.MkdirTemp("", resource)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	for _, v := range versions {
		tools.ExecPlain("rm", "-fv", filepath.Join(
			root, "internal", "v8s", "apis", resource, v,
			"zz_generated.deepcopy.go",
		))
	}
	tools.ExecPlain("rm", "-rfv", filepath.Join(
		root, "internal", "v8s", "gen", "client", resource,
	))
	tools.ExecPlain(GENERATE_SCRIPT, "all",
		path.Join(rootPackage, "/gen/client", resource),
		filepath.Join(rootPackage, "/apis"),
		resource+":"+strings.Join(versions, ","),
		"--go-header-file", filepath.Join(root, "internal", "v8s", "apis", "gen", "boilerplate.go.txt"),
		"--output-base", dir,
	)
	tools.ExecPlain("cp", "-r",
		filepath.Join(dir, rootPackage)+"/.",
		filepath.Join(root, "internal", "v8s"))
}

func buildCrd() {
	println("### Generating crds yaml ###")
	src := filepath.Join(root, "internal", "v8s", "apis", "vince", "v1alpha1")
	pkg := tools.Package("sigs.k8s.io/controller-tools")
	to := pkg.Path + "/cmd/controller-gen@" + pkg.Version
	println(">> using ", to)
	tools.ExecPlain("go", "install", to)
	out := filepath.Join(root, "manifests", "crds")
	tools.ExecPlain(
		"controller-gen",
		"crd",
		fmt.Sprintf("paths=%s", src),
		fmt.Sprintf("output:crd:artifacts:config=%s", out),
	)
}