package main

type Expression interface {
	Reduce(string) IMoney
}
