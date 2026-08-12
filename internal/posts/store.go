package posts

import "sync"

type Store struct {
	mu     sync.Mutex
	posts  []Post
	nextID int
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (s *Store) List() []Post {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.posts
}

func (s *Store) Create(title, body string) Post {
	s.mu.Lock()
	defer s.mu.Unlock()
	post := Post{ID: s.nextID, Title: title, Body: body}
	s.nextID++
	s.posts = append(s.posts, post)
	return post
}

func (s *Store) Get(id int) (Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.posts {
		if p.ID == id {
			return p, true
		}
	}
	return Post{}, false
}
