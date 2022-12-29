package main

import (
	"fmt"

	"github.com/krish-r/go-commit/color"
)

type ChangeType struct {
	id        int
	shortDesc string
	longDesc  string
}

func (c ChangeType) print() {
	id := color.Colorize(c.id, color.BrightGreen)
	longDesc := color.Colorize("("+c.longDesc+")", color.BrightBlackGray)

	fmt.Printf("%s: %s %s\n", id, c.shortDesc, longDesc)
}

var changeTypes = []ChangeType{
	{id: 1, shortDesc: "feat", longDesc: "A new feature"},
	{id: 2, shortDesc: "fix", longDesc: "A bug fix"},
	{id: 3, shortDesc: "docs", longDesc: "Documentation only changes"},
	{id: 4, shortDesc: "style", longDesc: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc.)"},
	{id: 5, shortDesc: "refactor", longDesc: "A code change that neither fixes a bug nor adds a feature"},
	{id: 6, shortDesc: "perf", longDesc: "A code change that improves the performance"},
	{id: 7, shortDesc: "test", longDesc: "Adding missing tests or correcting existing tests"},
	{id: 8, shortDesc: "build", longDesc: "Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"},
	{id: 9, shortDesc: "ci", longDesc: "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)"},
	{id: 10, shortDesc: "chore", longDesc: "Other changes that don't modify src or test files"},
	{id: 11, shortDesc: "revert", longDesc: "Reverts a previous commit"},
}
