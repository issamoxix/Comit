package auth

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetTokenPath() (string, error) {
	var basePath string

	switch runtime.GOOS {
	case "windows":

		if appData := os.Getenv("APPDATA"); appData != "" {
			basePath = filepath.Join(appData, "comit")
		} else {

			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			basePath = filepath.Join(home, "AppData", "Roaming", "comit")
		}
	case "darwin", "linux":

		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		basePath = filepath.Join(home, ".config", "comit")
	default:
		return "", os.ErrInvalid
	}

	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return "", err
	}

	return filepath.Join(basePath, "token.txt"), nil
}

func StoreToken(token string) error {
	tokenPath, err := GetTokenPath()
	if err != nil {
		return err
	}
	return os.WriteFile(tokenPath, []byte(token), 0600)
}

func GetToken() (string, error) {
	tokenPath, err := GetTokenPath()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
