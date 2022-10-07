package main

type IWritter interface {
	writeToDisk() error
}
