package main

import (
	"errors"
	"path/filepath"
)

var (
	ErrNotARepo = errors.New("Not a repo")
)

type Repository struct {
	dotDir FileSystemObject
}

func (repo Repository) discoverDotDir(root FileSystemObject) (FileSystemObject, error) {
	if root.isRoot() {
		return root, ErrNotARepo
	}
	result, err := root.DirectoryContains(".git")
	if nil != err {
		return root, err
	} else if result {
		return FileSystemObject(filepath.Join(root.String(), ".git")), nil
	}
	return repo.discoverDotDir(root.Parent())
}

func (repo *Repository) LoadOrInit(directory string) error {
	dot_dir, err := repo.discoverDotDir(FileSystemObject(directory))
	if nil != err {
		if ErrNotARepo == err {
			execute("git", "init")
		} else {
			return err
		}
	}
	repo.dotDir = dot_dir
	return nil
}
