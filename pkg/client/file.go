package client

//UploadFile uploads a file.
func UploadFile(masterHost, masterPort, name, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	url := "http://" + masterHost + ":" + masterPort + "/api/v1/files" + 
}