package polities

// PermissionLogged -->  Comprobar si la sección tiene todos los permisos se solicita
// usage -->  PermissionLogged ["permiso1@modulo1","permiso2@modulo1","permiso1@modulo2"]
// func PermissionLogged(permissionList []string) interface{} {
// 	wrappedFunction := func(c *fiber.Ctx) error {
// 		var activeUser models.ActiveUser
// 		activeUser, _ = c.Context().UserValue("activeUser").(models.ActiveUser)
// 		for _, modulePermission := range permissionList {
// 			values := strings.Split(modulePermission, "@")
// 			if len(values) != 2 {
// 				fmt.Println("Invalid permission name")
// 				continue
// 			}
// 			var found bool // By default, found is false
// 			for _, permission := range activeUser.Permissions {
// 				if permission.Module == values[1] && permission.Name == values[0] {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				return c.SendStatus(fiber.StatusForbidden)
// 			}
// 		}
// 		return c.Next()
// 	}
// 	return wrappedFunction
// }

// CheckPermission --> Comprobar si la sección tiene un permiso determinado en un modulo especifico
// func PermissionLoggedModule(module string, permissionName string) interface{} {
// 	wrappedFunction := func(c *fiber.Ctx) error {
// 		var activeUser models.ActiveUser
// 		activeUser, _ = c.Context().UserValue("activeUser").(models.ActiveUser)
// 		for _, permission := range activeUser.Permissions {
// 			if permission.Module == module && permission.Name == permissionName {
// 				return c.Next()
// 			}
// 		}
// 		return c.SendStatus(fiber.StatusForbidden)
// 	}
// 	return wrappedFunction
// }
