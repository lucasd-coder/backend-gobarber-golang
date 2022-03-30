package storage

type DiskStorageProvider struct{}

func NewDiskStorageProvider() *DiskStorageProvider {
	return &DiskStorageProvider{}
}

// func (disk *DiskStorageProvider) SaveFile(file string) string {
// 	out, err := os.Create("tmp/" + file)
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 	}
// 	defer out.Close()

// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
