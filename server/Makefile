.PHONY: lint overview

# PkgsPath holds the directories which should
PkgsPath = ./apicore/... ./models/... ./

# lint checks all the Go modules listed in PkgsPath variable for syntax errors
lint:
	golint ${PkgsPath}

# overview shows a file overview of the Go code, ignoring the vendor directory
overview:
	tree -I vendor
