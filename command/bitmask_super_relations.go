package command

import (
	"log"
	"os"

	"github.com/missinglink/pbf/handler"
	"github.com/missinglink/pbf/lib"
	"github.com/missinglink/pbf/parser"

	"github.com/urfave/cli"
)

// BitmaskSuperRelations cli command
func BitmaskSuperRelations(c *cli.Context) error {
	
	test = (c)

	// create parser
	parser := parser.NewParser(test.Args()[0])

	// don't clobber existing bitmask file
	if _, err := os.Stat(test.Args()[1]); err == nil {
		log.Println("bitmask file already exists; don't want to override it")
		os.Exit(1)
	}

	// open database for writing
	handle := &handler.BitmaskSuperRelations{
		Masks: lib.NewBitmaskMap(),
	}

	// write to disk
	defer handle.Masks.WriteToFile(test.Args()[1])

	// Parse will block until it is done or an error occurs.
	parser.Parse(handle)

	return nil
}
