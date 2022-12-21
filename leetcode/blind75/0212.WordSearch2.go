package blind75

type TrieNode struct {
	word     string
	children [26]*TrieNode
	isWord   bool
}

func (t *TrieNode) add(word string) {
	cur := t
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = &TrieNode{}
		}
		cur = cur.children[c-'a']
	}
	cur.word = word
	cur.isWord = true
}

func findWords(board [][]byte, words []string) []string {
	t := &TrieNode{}

	for _, word := range words {
		t.add(word)
	}

	res := make([]string, 0)

	for r := range board {
		for c := range board[r] {
			findWordsDfs(r, c, t, board, &res)
		}
	}
	return res
}

func findWordsDfs(r, c int, t *TrieNode, board [][]byte, res *[]string) {
	rowInbound := 0 <= r && r < len(board)
	colInbound := 0 <= c && c < len(board[0])
	if !rowInbound || !colInbound || board[r][c] == '0' || t.children[board[r][c]-'a'] == nil {
		return
	}
	t = t.children[board[r][c]-'a']
	if t.isWord && len(t.word) != 0 {
		*res = append(*res, t.word)
		t.word = ""
	}
	b := board[r][c]
	board[r][c] = '0'
	findWordsDfs(r-1, c, t, board, res)
	findWordsDfs(r+1, c, t, board, res)
	findWordsDfs(r, c-1, t, board, res)
	findWordsDfs(r, c+1, t, board, res)
	board[r][c] = b
}
