package dispatcher

import (
	"fmt"
	"strings"

	"Platypus/lib/context"
	"Platypus/lib/util/log"
)

func (dispatcher Dispatcher) Jump(args []string) {
	if len(args) != 1 {
		log.Error("Arguments error, use `Help Jump` to get more Jumprmation")
		dispatcher.JumpHelp([]string{})
		return
	}
	for _, server := range context.Ctx.Servers {
		for _, client := range (*server).GetAllTCPClients() {
			if strings.HasPrefix(client.Hash, strings.ToLower(args[0])) {
				context.Ctx.Current = client
				log.Success("The current interactive shell is set to: %s", client.Desc())
				return
			}
		}
	}
	log.Error("No such node")
}

func (dispatcher Dispatcher) JumpHelp(args []string) {
	fmt.Println("Usage of Jump")
	fmt.Println("\tJump [HASH]")
	fmt.Println("\tHASH\tThe hash of an node, node can be both a server or a client")
}

func (dispatcher Dispatcher) JumpDesc(args []string) {
	fmt.Println("Jump")
	fmt.Println("\tJump to a node, waiting to interactive with it")
}
