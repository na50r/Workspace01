package action2

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)


const (
	// Folder names
	userFolder = "users"
)



var dataPath = "c:/Playground/Workspace/repo/data"



type User struct  {
	Name string `json:"name"`
	Age  int `json:"age"`
}


func loadEnv(t * testing.T) {
	envVal := os.Getenv("DATA_PATH")
	if envVal == "" {
		return
	}

	var err error
	dataPath, err =  filepath.Abs(envVal)
	require.NoError(t, err, "error getting absolute path")
}


func CheckUserData(t * testing.T, userId, expName string, expAge int) {
	loadEnv(t)
	userFolder := createFilePath(dataPath, userFolder)
	user := fetchUser(t, userFolder, userId+".json")
	require.Equal(t, expName, user.Name, "name not equal")
	require.Equal(t, expAge, user.Age, "age not equal")
} 



func fetchUser(t *testing.T, folderPath string, file string) (user *User) {
	filePath := createFilePath(folderPath, file)
	jsonFile, err := os.ReadFile(filePath)
	require.NoError(t, err, "error reading file:", filePath)

	err = json.Unmarshal(jsonFile, &user)
	require.NoError(t, err, "error unmarshalling json file:", filePath)

	return
}



func createFilePath(folderPath string, file string) (filePath string) {
	filePath = filepath.Join(folderPath, file)
	return
}
