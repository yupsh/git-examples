package git_test

import (
	"fmt"
	"os"

	yup "github.com/gloo-foo/framework"
	"github.com/yupsh/git"
	. "github.com/gloo-foo/pipe"
)

func ExampleGit() {
	// git status
	// Note: This would execute actual git commands
	if err := yup.Run(Pipeline(
		git.Git("status"),
	)); err != nil {
		fmt.Fprintf(os.Stderr, "git: %v\n", err)
	}
}
