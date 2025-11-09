package git_test

import (
	"fmt"
	"os"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/git"
)

func ExampleGit() {
	// git status
	// Note: This would execute actual git commands
	if err := gloo.Run(Pipeline(
		git.Git("status"),
	)); err != nil {
		fmt.Fprintf(os.Stderr, "git: %v\n", err)
	}
}
