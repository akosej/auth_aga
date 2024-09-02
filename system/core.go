package system

import (
	"crypto/md5"
	"encoding/base64"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/araddon/dateparse"
	"github.com/joho/godotenv"
)

func FindFiles(root, ext string) []string {
	var a []string
	_ = filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

// CheckFilesFolders -- Check system files and folders
func CheckFilesFolders() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	// Check folders
	if _, err := os.Stat(Path + "/modules"); os.IsNotExist(err) {
		_ = os.MkdirAll(Path+"/modules", 0777)
	}
	if _, err := os.Stat(Path + "/plugins"); os.IsNotExist(err) {
		_ = os.MkdirAll(Path+"/plugins", 0777)
	}
}

func ImgToB64(dni, opt string) string {
	// Abrir el archivo de imagen
	path := Path + "/tools/credentials/personal_photo/" + dni + ".jpg"
	if opt == "overlappin" {
		path = Path + "/tools/credentials/access_card/" + dni + ".jpg"
	}
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Leer el contenido del archivo de imagen
	fileInfo, _ := file.Stat()
	buffer := make([]byte, fileInfo.Size())
	file.Read(buffer)

	// Convertir la imagen en base64
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buffer)
}

// HowManyDaysAgo --------------------------------------------------------------------------------------
// -- I work with dates
func HowManyDaysAgo(date string) (string, int) {
	loc, _ := time.LoadLocation("America/Havana")
	time.Local = loc
	parse, _ := dateparse.ParseLocal(date)
	days := int(time.Since(parse).Hours()) / 24
	if days >= 90 {
		return "Expirado", days
	}
	return "Valido", days
}

// GetMD5Hash
// -- Generate md5 with ldap format
func GetMD5Hash(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return "{MD5}" + base64.StdEncoding.EncodeToString(cipherStr)
}
