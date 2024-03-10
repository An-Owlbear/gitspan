package util

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
)

// Retrieves the list of git repositories in a given directory
func GetRepos(path string, depth int) []string {
	var repos []string
	dirs := getDirs(path)
	for _, e := range dirs {
		if isGitRepo(e) {
			repos = append(repos, e)
		} else if depth != 0 {
			repos = append(repos, GetRepos(e, depth-1)...)
		}
	}
	return repos
}

// Returns a list of directories in folder
func getDirs(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, filepath.Join(path, e.Name()))
		}
	}

	return dirs
}

// Determines whether the directory at a given path is a git repo
func isGitRepo(path string) bool {
	entries, _ := os.ReadDir(path)
	return slices.ContainsFunc(entries, func(e fs.DirEntry) bool {
		return e.Name() == ".git" && e.IsDir()
	})
}
