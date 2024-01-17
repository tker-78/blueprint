package main

import (
	"os"
	"path/filepath"
)

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var authAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (url string, err error) {
	if avatarURL, ok := c.userData["avatar_url"]; ok {
		avatarURLStr := avatarURL.(string)
		url = avatarURLStr
		return url, nil
	}
	return "", err

}

type FileSystemAvatar struct{}

var fileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(c *client) (url string, err error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {

			// useridと一致する名称のファイル名を返す
			if files, err := os.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					if match, _ := filepath.Match(useridStr+"*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", err
}
