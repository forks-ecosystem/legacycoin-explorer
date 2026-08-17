package explorer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Bookmark struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Address   string `json:"address,omitempty"`
	Txid      string `json:"txid,omitempty"`
	Label     string `json:"label,omitempty"`
	CreatedAt int64  `json:"created_at"`
}

type BookmarkStore struct {
	path string
	mu   sync.RWMutex
	list []*Bookmark
}

func NewBookmarkStore(dir string) *BookmarkStore {
	os.MkdirAll(dir, 0755)
	path := filepath.Join(dir, "bookmarks.json")
	bs := &BookmarkStore{path: path}
	bs.load()
	return bs
}

func (bs *BookmarkStore) load() {
	data, err := os.ReadFile(bs.path)
	if err != nil {
		bs.list = []*Bookmark{}
		return
	}
	var bookmarks []*Bookmark
	if err := json.Unmarshal(data, &bookmarks); err != nil {
		bs.list = []*Bookmark{}
		return
	}
	bs.list = bookmarks
}

func (bs *BookmarkStore) save() {
	data, _ := json.MarshalIndent(bs.list, "", "  ")
	os.WriteFile(bs.path, data, 0644)
}

func (bs *BookmarkStore) List() []*Bookmark {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	out := make([]*Bookmark, len(bs.list))
	copy(out, bs.list)
	return out
}

func (bs *BookmarkStore) Add(bm *Bookmark) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bm.CreatedAt = time.Now().Unix()
	bs.list = append(bs.list, bm)
	bs.save()
}

func (bs *BookmarkStore) Delete(id string) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	for i, bm := range bs.list {
		if bm.ID == id {
			bs.list = append(bs.list[:i], bs.list[i+1:]...)
			bs.save()
			return true
		}
	}
	return false
}

func (bs *BookmarkStore) Get(id string) *Bookmark {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	for _, bm := range bs.list {
		if bm.ID == id {
			return bm
		}
	}
	return nil
}
