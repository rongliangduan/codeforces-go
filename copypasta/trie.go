package copypasta

// 必要时禁止 GC，能加速不少
//func init() { debug.SetGCPercent(-1) }

/* 前缀树/字典树/单词查找树
另类解读：如果将字符串长度视作定值的话，trie 树是一种 O(n) 排序，O(1) 查询的数据结构
https://oi-wiki.org/string/trie/
https://www.quora.com/q/threadsiiithyderabad/Tutorial-on-Trie-and-example-problems
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieST.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieSET.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TST.java.html

模板题 LC208 https://leetcode-cn.com/problems/implement-trie-prefix-tree/
前缀和后缀搜索 周赛62D/LC745 https://leetcode-cn.com/problems/prefix-and-suffix-search/
回文对（配合 Manacher 可以做到线性复杂度）LC336 https://leetcode-cn.com/problems/palindrome-pairs/
todo https://codeforces.com/contest/455/problem/B
*/
type trieNode struct {
	son [26]*trieNode
	val int // 可以是个 []int，此时 cnt == len(val)
	cnt int

	// AC 自动机：当 o.son[i] 不能匹配文本串 text 中的某个字符时，o.fail 即为下一个应该查找的结点
	fail *trieNode
}

func (o *trieNode) empty() bool {
	for _, son := range o.son {
		if son != nil {
			return false
		}
	}
	return true
}

type trie struct{ root *trieNode }

func newTrie() *trie {
	// init with a root (empty string)
	return &trie{&trieNode{}}
}

func (trie) ord(c byte) byte { return c - 'a' }
func (trie) chr(v byte) byte { return v + 'a' }

// 插入字符串 s，附带值 val，返回插入后字符串末尾对应的节点
func (t *trie) put(s []byte, val int) *trieNode {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		//o.cnt++ // 统计子树字符串个数的写法
		//o.val = val // 更新 s 的所有前缀的值
	}
	o.cnt++
	o.val = val
	return o
}

// 查找字符串 s
func (t *trie) find(s []byte) *trieNode {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		// 未找到 s，且 s 不是任何字符串的前缀
		if o == nil {
			return nil
		}
	}
	// 未找到 s，但是 s 是某个字符串的前缀
	if o.cnt == 0 {
		return nil
	}
	return o
}

// 删除字符串 s，返回字符串末尾对应的节点
func (t *trie) delete(s []byte) *trieNode {
	fa := make([]*trieNode, len(s))
	o := t.root
	for i, b := range s {
		fa[i] = o
		o = o.son[t.ord(b)]
		if o == nil {
			return nil
		}
		//o.cnt-- // 对应 put 的写法
	}
	o.cnt--
	if o.cnt == 0 {
		for i := len(s) - 1; i >= 0; i-- {
			f := fa[i]
			f.son[t.ord(s[i])] = nil
			if !f.empty() {
				break
			}
		}
	}
	return o
}

// 求小于 s 的字符串个数
// 此时 o.cnt 保存子树字符串个数
func (t *trie) rank(s []byte) (k int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		for _, son := range o.son[:b] {
			if son != nil {
				k += son.cnt
			}
		}
		o = o.son[b]
		if o == nil {
			return
		}
	}
	//k += o.cnt // 这样写就是小于或等于 s 的字符串个数
	return
}

// 求第 k 小（k 从 0 开始，相当于有 k 个字符串小于返回的字符串 s）
// 此时 o.cnt 保存子树字符串个数
func (t *trie) kth(k int) (s []byte) {
	o := t.root
outer:
	for {
		for i, son := range o.son[:] {
			if son != nil {
				if k < son.cnt {
					o = son
					s = append(s, t.chr(byte(i)))
					continue outer
				}
				k -= son.cnt
			}
		}
		return
	}
}

// 结合 rank 和 kth，可以求出一个字符串的前驱和后继
// 见 bst.go 中的 prev 和 next

// 返回字符串 s 在 trie 中的前缀个数
// https://www.acwing.com/problem/content/144/
// https://codeforces.com/gym/101628/problem/K
func (t *trie) countPrefixOfString(s []byte) (cnt int) {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		if o == nil {
			return
		}
		cnt += o.cnt
	}
	return
}

// 返回 trie 中前缀为 p 的字符串个数
// 此时 o.cnt 保存子树字符串个数
// https://codeforces.com/gym/101628/problem/K
func (t *trie) countStringHasPrefix(p []byte) int {
	o := t.root
	for _, b := range p {
		o = o.son[t.ord(b)]
		if o == nil {
			return 0
		}
	}
	return o.cnt
}

// EXTRA: AC 自动机 Aho–Corasick algorithm / Deterministic Finite Automaton (DFA)
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
// https://en.wikipedia.org/wiki/Deterministic_finite_automaton
// 基础实现 https://zhuanlan.zhihu.com/p/80325757
// 基础实现 https://www.cnblogs.com/nullzx/p/7499397.html
// 改进实现 https://oi-wiki.org/string/ac-automaton/
// 应用 https://cp-algorithms.com/string/aho_corasick.html
//
// 模板题
// LC1032 https://leetcode-cn.com/problems/stream-of-characters/
// https://www.luogu.com.cn/problem/P3808
// https://www.luogu.com.cn/problem/P3796
// todo https://www.luogu.com.cn/problem/P5357 二次加强版
//
// todo https://codeforces.com/problemset/problem/963/D
func (t *trie) buildDFA() {
	q := []*trieNode{}
	for _, son := range t.root.son[:] {
		if son != nil {
			son.fail = t.root
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		if o.fail == nil {
			o.fail = t.root
		}
		for i, son := range o.son[:] {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
}

// 有多少个（编号）不同的模式串在文本串 text 里出现过
func (t *trie) sumCountAllPatterns(text []byte) (cnt int) {
	o := t.root
	for _, b := range text {
		o = o.son[t.ord(b)]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != nil && f.val > 0; f = f.fail {
			cnt += f.val
			f.val = 0
		}
	}
	return
}

// 返回所有模式串 patterns 的开头在文本串 text 的所有位置（未找到时对应数组为空）
// patterns 为模式串数组（无重复元素），为方便起见，patterns 从 1 开始
func (t *trie) acSearch(text []byte, patterns [][]byte) [][]int {
	pos := make([][]int, len(patterns))
	o := t.root
	for i, b := range text {
		o = o.son[t.ord(b)]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != nil; f = f.fail {
			if pid := f.val; pid != 0 {
				pos[pid] = append(pos[pid], i-len(patterns[pid])+1) // 也可以只记录 i，代表模式串末尾在文本的位置
			}
		}
	}
	return pos
}

//

// 可持久化 trie
// TODO https://oi-wiki.org/ds/persistent-trie/
// 模板题（最大异或和） https://www.luogu.com.cn/problem/P4735

//

// Suffix automaton (SAM)
// https://en.wikipedia.org/wiki/Suffix_automaton
//《后缀自动机》，陈立杰
//《后缀自动机在字典树上的拓展》，刘研绎
//《后缀自动机及其应用》，张天扬
// todo https://www.bilibili.com/video/av756051240/
// todo https://baobaobear.github.io/post/20200220-sam/
// todo https://codeforces.com/blog/entry/20861
// TODO https://oi-wiki.org/string/sam/
// TODO https://cp-algorithms.com/string/suffix-automaton.html
//      后缀树简介 https://eternalalexander.github.io/2019/10/31/%E5%90%8E%E7%BC%80%E6%A0%91%E7%AE%80%E4%BB%8B/
// 模板题 https://www.luogu.com.cn/problem/P3804

// 广义 SAM
// todo https://www.luogu.com.cn/problem/P6139
// todo https://codeforces.com/problemset/problem/1437/G

// 回文自动机 PAM
// todo https://baobaobear.github.io/post/20200416-pam/
//  https://www.luogu.com.cn/problem/P5496