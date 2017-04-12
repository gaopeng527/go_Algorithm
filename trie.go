// trie 字典树实现
package Algorithm

// 字典树节点
type TrieNode struct {
	children map[interface{}]*TrieNode
	isEnd    bool
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[interface{}]*TrieNode), isEnd: false}
}

// 字典树
type Trie struct {
	root *TrieNode
}

// 构造字典树
func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// 向字典树中插入一个单词
func (trie *Trie) Insert(word []interface{}) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

// 搜索字典树中是否存在指定单词
func (trie *Trie) Search(word []interface{}) bool {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			return false
		}
		node = node.children[word[i]]
	}
	return node.isEnd
}

// 判断字典树中是否有指定前缀的单词
func (trie *Trie) StartsWith(prefix []interface{}) bool {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}
