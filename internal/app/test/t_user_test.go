package test

import (
	"net/http/httptest"
	"testing"

	"github.com/gotidy/ptr"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/pkg/helper/hash"
	uid "github.com/heromicro/omgind/pkg/helper/uid/ulid"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	const router = apiPrefix + "v1/users"
	var err error

	w := httptest.NewRecorder()

	// post /menus
	addMenuItem := &schema.Menu{
		Name:     uid.MustString(),
		IsShow:   ptr.Bool(true),
		IsActive: ptr.Bool(true),
	}
	engine.ServeHTTP(w, newPostRequest(apiPrefix+"v1/menus", addMenuItem))
	assert.Equal(t, 200, w.Code)
	var addMenuItemRes ResID
	err = parseReader(w.Body, &addMenuItemRes)
	assert.Nil(t, err)

	// post /roles
	addRoleItem := &schema.Role{
		Name:     uid.MustString(),
		IsActive: ptr.Bool(true),
		RoleMenus: schema.RoleMenus{
			&schema.RoleMenu{
				MenuID: addMenuItemRes.ID,
			},
		},
	}
	engine.ServeHTTP(w, newPostRequest(apiPrefix+"v1/roles", addRoleItem))
	assert.Equal(t, 200, w.Code)
	var addRoleItemRes ResID
	err = parseReader(w.Body, &addRoleItemRes)
	assert.Nil(t, err)

	// post /users
	addItem := &schema.User{
		UserName: uid.MustString(),
		RealName: uid.MustString(),
		IsActive: ptr.Bool(true),
		Password: hash.MD5String("test"),
		UserRoles: schema.UserRoles{
			&schema.UserRole{
				RoleID: addRoleItemRes.ID,
			},
		},
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)
	var addItemRes ResID
	err = parseReader(w.Body, &addItemRes)
	assert.Nil(t, err)

	// get /users/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	var getItem schema.User
	err = parseReader(w.Body, &getItem)
	assert.Nil(t, err)
	assert.Equal(t, addItem.UserName, getItem.UserName)
	assert.Equal(t, *addItem.IsActive, *getItem.IsActive)
	assert.NotEmpty(t, getItem.ID)

	// put /users/:id
	putItem := getItem
	putItem.UserName = uid.MustString()
	engine.ServeHTTP(w, newPutRequest("%s/%s", putItem, router, getItem.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// query /users
	engine.ServeHTTP(w, newGetRequest(router, newPageParam()))
	assert.Equal(t, 200, w.Code)
	var pageItems []*schema.User
	err = parsePageReader(w.Body, &pageItems)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(pageItems), 1)
	if len(pageItems) > 0 {
		assert.Equal(t, putItem.ID, pageItems[0].ID)
		assert.Equal(t, putItem.UserName, pageItems[0].UserName)
	}

	// delete /users/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%s", router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// delete /roles/:id
	engine.ServeHTTP(w, newDeleteRequest(apiPrefix+"v1/roles/%s", addRoleItemRes.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// delete /menus/:id
	engine.ServeHTTP(w, newDeleteRequest(apiPrefix+"v1/menus/%s", addMenuItemRes.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)
}
