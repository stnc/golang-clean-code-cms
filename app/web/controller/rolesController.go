package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"
	"strconv"

	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"

	"stncCms/app/domain/helpers/rbac"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// Permission constructor
type Roles struct {
	permissionApp     services.PermissionAppInterface
	modulesApp        services.ModulesAppInterface
	roleApp           services.RoleAppInterface
	rolePermissionApp services.RolePermissionAppInterface
}

const viewPathPermission = "admin/roles/"

func InitRoles(perApp services.PermissionAppInterface, modApp services.ModulesAppInterface, rolesApp services.RoleAppInterface, rolePermApp services.RolePermissionAppInterface) *Roles {
	return &Roles{
		permissionApp:     perApp,
		modulesApp:        modApp,
		roleApp:           rolesApp,
		rolePermissionApp: rolePermApp,
	}
}

// Index list
func (access *Roles) Index(c *gin.Context) {
	rbac.RbacCheck(c, "post-index")
	// stncsession.IsLoggedInRedirect(c)
	var date stncdatetime.Inow
	var total int64
	access.roleApp.Count(&total)
	postsPerPage := 3
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	data, _ := access.roleApp.GetAllPagination(postsPerPage, offset)
	viewData := pongo2.Context{
		"paginator": paginator,
		"title":     "List",
		"dataList":  data,
		"date":      date,
		"csrf":      csrf.GetToken(c),
	}

	c.HTML(
		http.StatusOK,
		viewPathPermission+"index.html",
		viewData,
	)
}

// Create all list f
func (access *Roles) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var data []dto.ModulesAndPermission
	data, _ = access.modulesApp.GetAllModulesMerge()
	for num, v := range data {
		var list = []entity.Permission{}
		list, _ = access.permissionApp.GetAllPaginationermissionForModulID(int(v.ID))
		data[num].Permissions = list
	}

	// //#json formatter #stncjson
	// empJSON, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

	viewData := pongo2.Context{
		"paginator": paginator,
		"title":     "permissions",
		"datas":     data,
		"csrf":      csrf.GetToken(c),
	}

	c.HTML(
		http.StatusOK,
		viewPathPermission+"create.html",
		viewData,
	)

}

// store data
func (access *Roles) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	//once bi yere rolu kaydet
	//sonra kaydeidlern id yi alman lazim

	roleData := entity.Role{
		Title: c.PostForm("Title"),
		Slug:  c.PostForm("Title"),
	}
	saveRoleData, _ := access.roleApp.Save(&roleData)

	roleID := saveRoleData.ID

	var data []dto.ModulesAndPermission
	data, _ = access.modulesApp.GetAllModulesMerge()
	for num, v := range data {
		var list = []entity.Permission{}
		list, _ = access.permissionApp.GetAllPaginationermissionForModulID(int(v.ID))
		data[num].Permissions = list
	}

	for _, v := range data {
		for _, per := range v.Permissions {
			rolePermissondata := entity.RolePermisson{
				RoleID:       roleID,
				PermissionID: per.ID,
				Active:       0,
			}
			access.rolePermissionApp.Save(&rolePermissondata)
		}
	}

	names, _ := c.Request.PostForm["grant-caps[]"]
	for _, row := range names {
		grandPermissionID := stnccollection.StringToint(row)
		access.rolePermissionApp.UpdateActiveStatus(roleID, grandPermissionID, 1)
	}

	stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin/roles/edit/"+stnccollection.IntToString(roleID))
	return

	// grandPermissionID, e1 := strconv.Atoi(row)
	// fmt.Println(e1)
	// if e1 == nil {
	// 	fmt.Println("\nAfter:")
	// 	fmt.Printf("Type: %T ", grandPermissionID)
	// 	fmt.Printf("\nValue: %v", grandPermissionID)
	// }
	///*******************
	// names := c.PostFormMap("grant-caps")
	// for _, row := range names {
	// 	fmt.Println(row)
	// }
	// names := c.QueryMap("grant-caps[]")

	// names, _ := c.Request.PostForm["grant-caps[]"]
	// fmt.Println("grand list : ", names)
	// for _, row := range names {
	// 	fmt.Println("granst: ", row)

	// }

}

func (access *Roles) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var data []dto.ModulesAndPermissionRole

	if roleID, err := strconv.Atoi(c.Param("ID")); err == nil {
		data, _ = access.modulesApp.GetAllModulesMergePermission()
		roleData, _ := access.roleApp.GetByID(roleID) //TODO: bu veri  access.roleApp.EditList iicne de geliyor orada mi almak mantakli ??
		for num, v := range data {
			var list = []dto.RoleEditList{}
			list, _ = access.roleApp.EditList(v.ID, roleID)
			// fmt.Println(v.ModulName)
			data[num].RoleEditList = list
		}
		viewData := pongo2.Context{
			"title":    "permissions",
			"roleData": roleData,
			"datas":    data,
			"roleID":   roleID,
			"csrf":     csrf.GetToken(c),
		}

		c.HTML(
			http.StatusOK,
			viewPathPermission+"edit.html",
			viewData,
		)
	}

}

func (access *Roles) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	roleID := c.PostForm("roleID")
	roleIDint := stnccollection.StringToint(c.PostForm("roleID"))
	title := c.PostForm("Title")

	access.roleApp.UpdateTitle(roleIDint, title)

	//TODO: neden calismadi
	// titleSlug := stnchelper.Slugify(title, 15)
	// access.roleApp.Update(
	// 	&entity.Role{
	// 		ID:      roleIDint,
	// 		Title:   title,
	// 		Slug:    titleSlug,
	// 		Context: titleSlug,
	// 		Status:  1,
	// 	})

	grants, _ := c.Request.PostForm["grant-caps[]"]
	for _, row := range grants {
		grandPermissionID := stnccollection.StringToint(row)
		access.rolePermissionApp.UpdateActiveStatus(roleIDint, grandPermissionID, 1)
	}

	deny, _ := c.Request.PostForm["deny-caps[]"]
	for _, row := range deny {
		grandPermissionID := stnccollection.StringToint(row)
		access.rolePermissionApp.UpdateActiveStatus(roleIDint, grandPermissionID, 0)
	}

	stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin/roles/edit/"+roleID)
	return
}

func (access *Roles) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {

		access.roleApp.Delete(ID)
		stncsession.SetFlashMessage("delete", "success", c)

		c.Redirect(http.StatusMovedPermanently, "/"+viewPathPermission)
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (access *Roles) IndexKnockout(c *gin.Context) {
	// allpermission, err := access.permissionApp.GetAllPaginationermission()

	stncsession.IsLoggedInRedirect(c)
	var date stncdatetime.Inow

	var data []dto.ModulesAndPermission
	data, _ = access.modulesApp.GetAllModulesMerge()
	for num, v := range data {
		var list = []entity.Permission{}
		list, _ = access.permissionApp.GetAllPaginationermissionForModulID(int(v.ID))
		data[num].Permissions = list
	}

	// //#json formatter #stncjson https://github.com/TylerBrock/colorjson
	empJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

	// js, _ := json.Marshal(data)

	// fmt.Println((jsonData))

	viewData := pongo2.Context{
		"paginator": paginator,
		"title":     "permissions",
		"datas":     data,
		"json":      string(empJSON),
		"date":      date,
		"csrf":      csrf.GetToken(c),
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathPermission+"knockout.html",
		// Pass the data that the page uses
		viewData,
	)
}
