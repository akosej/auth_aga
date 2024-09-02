package polities

// func TwoFa(c *fiber.Ctx) error {
// 	if !core.TwoFactorPluginActive() {
// 		return c.Next()
// 	}
// 	// TODO Ahi que chequear si el usuario tiene activo el 2fa
// 	activeUser, _ := c.Context().UserValue("activeUser").(models.ActiveUser)
// 	if activeUser.TwoFactor.Status != "TRUE" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": core.GetTextMessage("2fa_err_002"),
// 		})
// 	}

// 	var data map[string]string
// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	code, ok := data["code"]
// 	if !ok {
// 		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": core.GetTextMessage("2fa_ok_003")})
// 	}

// 	claims, _ := controllers.GetClaims(c)

// 	var tfaSeed models.TwoFAUserSeed
// 	result := database.DB.Model(models.TwoFAUserSeed{}).Where("user_id = ?", claims["ID"].(string)).First(&tfaSeed)
// 	if result.RowsAffected == 0 {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": core.GetTextMessage("2fa_err_002"),
// 		})
// 	}

// 	interval := time.Now().Unix() / 30

// 	getTwoFa := core.ExtractFunctionsPlugins(
// 		"TwoFactor",
// 		"TwoFa",
// 		claims["ID"].(string), tfaSeed.Seed, interval)

// 	twoFa := reflect.ValueOf(getTwoFa).String()

// 	if twoFa == code {
// 		return c.Next()
// 	}

// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": core.GetTextMessage("login_err_004"),
// 	})
// }

// NOTE Este handler es para las url que se necesite verificar la seguridad
// USE Si el usuario tiene 2fa puede enviar "code" de lo cotrario tiene que enviar la "password"
// func Security(c *fiber.Ctx) error {
// 	// Se obtienen los datos enviados en el cuerpo de la solicitud como un mapa
// 	var data map[string]string
// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	// Se obtienen las reclamaciones del token JWT
// 	claims, _ := controllers.GetClaims(c)

// 	// Se obtiene el usuario activo del contexto de la solicitud
// 	activeUser, _ := c.Context().UserValue("activeUser").(models.ActiveUser)

// 	/*
// 		Si el plugin TwoFactorPlugin está activo y la 2fa está habilitada para el
// 		usuario y se proporciona un código de autenticación, se verifica el código
// 	*/
// 	if core.TwoFactorPluginActive() && activeUser.TwoFactor.Status == "TRUE" && data["code"] != "" {
// 		// Se obtiene el código de autenticación enviado por el usuario
// 		code, ok := data["code"]
// 		if !ok {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": core.GetTextMessage("2fa_ok_003")})
// 		}

// 		// Se obtiene la semilla de autenticación de dos factores del usuario
// 		var tfaSeed models.TwoFAUserSeed
// 		result := database.DB.Model(models.TwoFAUserSeed{}).Where("user_id = ?", claims["ID"].(string)).First(&tfaSeed)
// 		if result.RowsAffected == 0 {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": core.GetTextMessage("2fa_err_002"),
// 			})
// 		}

// 		// Se calcula el intervalo de tiempo actual para generar el código de autenticación
// 		interval := time.Now().Unix() / 30

// 		// Se obtiene el código de autenticación generado por el plugin TwoFactorPlugin
// 		getTwoFa := core.ExtractFunctionsPlugins(
// 			"TwoFactor",
// 			"TwoFa",
// 			claims["ID"].(string), tfaSeed.Seed, interval)
// 		twoFa := reflect.ValueOf(getTwoFa).String()

// 		// Si el código de autenticación es correcto, se salta al siguiente middleware
// 		if twoFa == code {
// 			return c.Next()
// 		}

// 		// Si el código de autenticación es incorrecto, se devuelve un error de autenticación
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": core.GetTextMessage("login_err_004"),
// 		})
// 	}

// 	// Si la autenticación de dos factores no está habilitada o no se proporciona un código de autenticación,
// 	// se verifica la contraseña del usuario
// 	password, ok := data["password"]
// 	if !ok {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": core.GetTextMessage("check_password_003")})
// 	}
// 	// Si la contraseña es correcta, se salta al siguiente middleware
// 	if claims["password"].(string) == password {
// 		c.Next()
// 	}

// 	// Si la contraseña es incorrecta, se devuelve un error de autenticación
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": core.GetTextMessage("check_password_004"),
// 	})
// }
