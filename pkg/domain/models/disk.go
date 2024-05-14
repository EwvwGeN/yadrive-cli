package models

type DiskInfo struct {
	TrashSize     int64             `json:"trash_size"`
	TotalSpace    int64             `json:"total_space"`
	UsedSpace     int64             `json:"used_space"`
	SystemFolders map[string]string `json:"system_folders"`
}