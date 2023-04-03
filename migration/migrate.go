package migration

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/pflag"

	"github.com/test-scylla/config"
	"github.com/test-scylla/db"
)

var verbose = pflag.Bool("verbose", false, "output more info")

func init() {

	log.Println("Bootstrap database...")

	if *verbose {
		log.Printf("Configuration = %+v\n", config.Config())
	}

	createKeyspace()
	migrateKeyspace()
	printKeyspaceMetadata()
}

func createKeyspace() {
	ses, err := config.Session()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := ses.Query(db.KeySpaceQuery).Exec(); err != nil {
		log.Fatalln("ensure keyspace exists: ", err)
	}
}

func migrateKeyspace() {
	ses, err := config.Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	runMigration()
	// if err := migrate.Migrate(context.Background(), ses, "db/cql"); err != nil {
	// 	log.Fatalln("migrate: ", err)
	// }
}

func runMigration(sess string) {
	e := exec.Command("make", "migration_up", sess)
	var out bytes.Buffer
	e.Stdout = &out
	err := e.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %q\n", out.String())
}

func printKeyspaceMetadata() {
	ses, err := config.Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	m, err := ses.KeyspaceMetadata(db.KeySpace)
	if err != nil {
		log.Fatalln("keyspace metadata: ", err)
	}

	log.Printf("Keyspace metadata = %+v\n", *m)
}
