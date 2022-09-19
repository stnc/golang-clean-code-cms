package rbac

import (
	"fmt"
	"net/http"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncsession"
	repository "stncCms/app/domain/repository/cacheRepository"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
)

func RbacCheckRolePermission(RoleID int, permissionName string) bool {

	db := repository.DB
	appGrup := repository.PermissionRepositoryInit(db)
	userPermissionData, _ := appGrup.GetUserPermission(RoleID)

	words := []string{}
	for _, row := range userPermissionData {
		words = append(words, row.PermissionName)
	}

	_, found := stnccollection.FindSlice(words, permissionName)
	if found {
		return true
	} else {
		return false
	}
}

func RbacCheck(c *gin.Context, permissionName string) {
	stncsession.IsLoggedInRedirect(c)

	db := repository.DB
	appGrup := repository.UserRepositoryInit(db)
	userID := stncsession.GetUserID2(c)
	fmt.Println(userID)

	userData, _ := appGrup.GetUser(userID)
	roleID := userData.RoleID
	fmt.Println(roleID)
	found := RbacCheckRolePermission(roleID, permissionName)
	if found {
		viewDataPer := pongo2.Context{"title": "List"}
		c.HTML(
			http.StatusOK,
			"admin/roles/permission.html",
			viewDataPer,
		)
	}
	return
}
