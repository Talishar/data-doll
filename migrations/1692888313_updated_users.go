package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add
		new_patreonAccessToken := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "0zyrwc9c",
			"name": "patreonAccessToken",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_patreonAccessToken)
		collection.Schema.AddField(new_patreonAccessToken)

		// add
		new_patreonRefreshToken := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gysohcxx",
			"name": "patreonRefreshToken",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_patreonRefreshToken)
		collection.Schema.AddField(new_patreonRefreshToken)

		// add
		new_fabraryId := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ansethir",
			"name": "fabraryId",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_fabraryId)
		collection.Schema.AddField(new_fabraryId)

		// add
		new_fabdbId := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zp5li16w",
			"name": "fabdbId",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_fabdbId)
		collection.Schema.AddField(new_fabdbId)

		// add
		new_patreonEnum := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "8d50bqy5",
			"name": "patreonEnum",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_patreonEnum)
		collection.Schema.AddField(new_patreonEnum)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("0zyrwc9c")

		// remove
		collection.Schema.RemoveField("gysohcxx")

		// remove
		collection.Schema.RemoveField("ansethir")

		// remove
		collection.Schema.RemoveField("zp5li16w")

		// remove
		collection.Schema.RemoveField("8d50bqy5")

		return dao.SaveCollection(collection)
	})
}
