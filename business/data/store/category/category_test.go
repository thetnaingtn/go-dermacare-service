package category_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/category"
	"github.com/thetnaingtn/go-dermacare-service/business/data/test"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dbc = test.DBContainer{
	Image: "mongo:latest",
	Port:  "27017",
}

func TestCategory(t *testing.T) {
	db, teardown := test.NewUnit(t, dbc)
	t.Cleanup(teardown)

	store := category.NewStore(db)
	t.Log("Given the need to work with Category records.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen handling a single Category.", testID)
		{
			nc := category.NewCategory{
				Name:        "Test Category",
				Description: "Category create for unit test",
			}

			c, err := store.Create(nc)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to create category : %s.", test.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to create category.", test.Success, testID)

			objId, _ := primitive.ObjectIDFromHex(c.ID)
			saved, err := store.QueryById(objId)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to query category by ID: %s.", test.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to retrieve category by ID.", test.Success, testID)

			if diff := cmp.Diff(c, saved); diff != "" {
				t.Fatalf("\t%s\tTest %d:\tShould get back the same category. Diff:\n%s", test.Failed, testID, diff)
			}
			t.Logf("\t%s\tTest %d:\tShould get back the same category.", test.Success, testID)
		}
	}
}
