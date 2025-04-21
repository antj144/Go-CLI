package main

type cmdFlags struct {
	Add string
	Del int
	Edit string
	List bool
}

func NewCmdFlags