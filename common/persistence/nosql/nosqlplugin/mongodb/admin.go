package mongodb

import "io/ioutil"

var _ nosqlplugin.AdminDB = (*mdb)(nil)

const (
	testSchemaDir = "schema/mongodb/"
)

func (db *mdb) SetupTestDatabase(schemaBaseDir string) error {
	if schemaBaseDir == "" {
		var err error
		schemaBaseDir, err = nosqlplugin.GetDefaultTestSchemaDir(testSchemaDir)
		if err != nil {
			return err
		}
	}

	schemaFile := schemaBaseDir + "cadence/schema.json"
	byteValues, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		return err
	}
	var commands []interface{}
	err = bson.UnmarshalExtJSON(byteValues, false, &commands)
	if err != nil {
		return err
	}
	for _, cmd := range commands {
		result := db.dbConn.RunCommand(context.Background(), cmd)
		if result.Err() != nil {
			return err
		}
	}
	return nil
}

func (db *mdb) TeardownTestDatabase() error {
	result := db.dbConn.RunCommand(context.Background(), bson.D{{"dropDatabase", 1}})
	err := result.Err()
	return err
}
