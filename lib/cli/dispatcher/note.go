package dispatcher

import (
	"fmt"
	"strings"

	"Platypus/lib/context"
	"Platypus/lib/util/log"
)

func (dispatcher Dispatcher) Note(args []string) {
	if len(args) < 2 {
		log.Error("Arguments error, use `Help Note` to get more information")
		dispatcher.NoteHelp([]string{})
		return
	}

	for _, server := range context.Ctx.Servers {
		for _, client := range (*server).GetAllTCPClients() {
			if strings.HasPrefix(client.Hash, strings.ToLower(args[0])) {
				client.Note = strings.Join(args[1:], " ")

				log.Success("The description of the note has changed: %s", client.Note)
				fmt.Println("[CLIENT]: ", client.Desc())
				return
			}
		}
	}

	log.Error("No such node")
}

func (dispatcher Dispatcher) NoteHelp(args []string) {
	fmt.Println("Usage of Note")
	fmt.Println("\tInfo [HASH] [NOTE]")
	fmt.Println("\tHASH\tThe hash of a node")
	fmt.Println("\tNOTE\tThe description of the node which is specified")
}

func (dispatcher Dispatcher) NoteDesc(args []string) {
	fmt.Println("Note")
	fmt.Println("\tChange the description of a note")
}
