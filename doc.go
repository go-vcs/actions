// Package actions will hold the GitHub Actions workflow engine (parsing and
// planning: matrix, needs, expressions) and the server side of the gitea-runner
// protocol, reusable without the go-vcs forge.
//
// It is not started: see https://github.com/go-vcs/forge/blob/main/docs/design.md,
// milestone 5.
package actions
