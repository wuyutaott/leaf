package network

type IAgent interface {
	Run()
	OnClose()
}
