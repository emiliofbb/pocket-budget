package migrations

import (
	"github.com/pocketbase/dbx"
    "github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"

    "github.com/pocketbase/pocketbase/models"
    "github.com/pocketbase/pocketbase/models/schema"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
        dao := daos.New(db)
        
        userCollection, err := dao.FindCollectionByNameOrId("users");

        if err != nil {
            return err
        }

        

        noteCollection := &models.Collection{
            Name: "note",
            Type: models.CollectionTypeBase,
            ListRule: nil,
            ViewRule: nil,
            CreateRule: nil,
            UpdateRule: nil,
            DeleteRule: nil,
            Schema: schema.NewSchema(
                &schema.SchemaField{
                    Name: "name",
                    Type: schema.FieldTypeText,
                    Required: true,
                },
                &schema.SchemaField{
                    Name: "image",
                    Type: schema.FieldTypeFile,
                    Required: false,
                    Options: &schema.FileOptions{
                        MimeTypes: types.JsonArray[string]{
                            "image/jpeg",
                            "image/png",
                            "image/svg+xml",
                            "image/gif",
                            "image/webp",
                        },
                        MaxSize: 500000000,
                    },
                },
                &schema.SchemaField{
                    Name: "notify_time",
                    Type: schema.FieldTypeDate,
                    Required: false,
                },
                &schema.SchemaField{
                    Name: "comment",
                    Type: schema.FieldTypeText,
                    Required: false,
                },
                &schema.SchemaField{
                    Name: "user",
                    Type: schema.FieldTypeRelation,
                    Required: true,
                    Options:  &schema.RelationOptions{
                        MaxSelect:     types.Pointer(1),
                        CollectionId:  userCollection.Id,
                        CascadeDelete: true,
                    },
                },
            ),
        }
        categoryCollection := &models.Collection{
            Name: "category",
            Type: models.CollectionTypeBase,
            ListRule: nil,
            ViewRule: nil,
            CreateRule: nil,
            UpdateRule: nil,
            DeleteRule: nil,
            Schema: schema.NewSchema(
                &schema.SchemaField{
                    Name: "name",
                    Type: schema.FieldTypeText,
                    Required: true,
                },
            ),
        }
        subcategoryCollection := &models.Collection{
            Name: "subcategory",
            Type: models.CollectionTypeBase,
            ListRule: nil,
            ViewRule: nil,
            CreateRule: nil,
            UpdateRule: nil,
            DeleteRule: nil,
            Schema: schema.NewSchema(
                &schema.SchemaField{
                    Name: "name",
                    Type: schema.FieldTypeText,
                    Required: true,
                },
                &schema.SchemaField{
                    Name: "color",
                    Type: schema.FieldTypeJson,
                    Required: true,
                },
                &schema.SchemaField{
                    Name: "category",
                    Type: schema.FieldTypeRelation,
                    Required: true,
                    Options:  &schema.RelationOptions{
                        MaxSelect:     types.Pointer(1),
                        CollectionId:  "category-pb",
                        CascadeDelete: true,
                    },
                },
            ),
        }
        moneyFlowCollection := &models.Collection{
            Name: "money_flow",
            Type: models.CollectionTypeBase,
            ListRule: nil,
            ViewRule: nil,
            CreateRule: nil,
            UpdateRule: nil,
            DeleteRule: nil,
            Schema: schema.NewSchema(
                &schema.SchemaField{
                    Name: "name",
                    Type: schema.FieldTypeText,
                    Required: true,
                },
                &schema.SchemaField{
                    Name: "value",
                    Type: schema.FieldTypeNumber,
                    Required: true,
                },
                &schema.SchemaField{
                    Name: "is_enty",
                    Type: schema.FieldTypeBool,
                    Required: true,
                },
                &schema.SchemaField{
                    Name: "image",
                    Type: schema.FieldTypeFile,
                    Required: false,
                    Options: &schema.FileOptions{
                        MimeTypes: types.JsonArray[string]{
                            "image/jpeg",
                            "image/png",
                            "image/svg+xml",
                            "image/gif",
                            "image/webp",
                        },
                        MaxSize: 500000000,
                    },
                },
                &schema.SchemaField{
                    Name: "comment",
                    Type: schema.FieldTypeText,
                    Required: false,
                },
                &schema.SchemaField{
                    Name: "subcategory",
                    Type: schema.FieldTypeRelation,
                    Required: true,
                    Options:  &schema.RelationOptions{
                        MaxSelect:     types.Pointer(1),
                        CollectionId:  "subcategory-pb",
                        CascadeDelete: true,
                    },
                },
                &schema.SchemaField{
                    Name: "user",
                    Type: schema.FieldTypeRelation,
                    Required: true,
                    Options:  &schema.RelationOptions{
                        MaxSelect:     types.Pointer(1),
                        CollectionId:  userCollection.Id,
                        CascadeDelete: true,
                    },
                },
            ),
        }
        noteCollection.SetId("note-pb")
        if err := dao.SaveCollection(noteCollection); err != nil {
            return err
        } 
        categoryCollection.SetId("category-pb")
        if err := dao.SaveCollection(categoryCollection); err != nil {
            return err
        } 
        subcategoryCollection.SetId("subcategory-pb")
        if err := dao.SaveCollection(subcategoryCollection); err != nil {
            return err
        } 
        moneyFlowCollection.SetId("money-flow-pb")
        if err := dao.SaveCollection(moneyFlowCollection); err != nil {
            return err
        } 
        return nil
	}, nil)
}
