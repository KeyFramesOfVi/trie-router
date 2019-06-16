package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type TrieSuite struct {
	suite.Suite
	trie *Trie
}

func (s *TrieSuite) SetupTest() {
	s.trie = NewTrie()
}

func (s *TrieSuite) TestTrie() {
	putCheck := s.trie.Put("/hello/world", 20)
	assert.Equal(s.T(), nil, putCheck)
	handler, getCheck := s.trie.Get("/hello/world")
	assert.Equal(s.T(), nil, getCheck)
	assert.Equal(s.T(), handler, 20)
	handler, getCheck = s.trie.Get("/")
	assert.Equal(s.T(), nil, handler)
}

func TestTrieSuite(t *testing.T) {
	suite.Run(t, &TrieSuite{trie: NewTrie()})
}
