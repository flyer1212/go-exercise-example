package musiclib

import "errors"

/**
  很重要的Music基础数据管理类
*/

/**
  音乐对象
*/
type MusicEntry struct {
	Id     string // id
	Name   string // 音乐名
	Artist string //艺术家名
	Source string //音乐位置
	Type   string // 音乐文件类型
}

/**
  定义一个数组切片用于存储音乐
*/
type MusicManager struct {
	musics [] MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

/**
  音乐长度
*/
func (m *MusicManager) Len() int {
	return len(m.musics)
}

/**
  根据下标返回某一个音乐
*/
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	// 从数组切片中删除元素
	if index < len(m.musics)-1 {
		// 中间元素 通过切片的形式删除数据
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		// 删除仅有的一个元素
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}
