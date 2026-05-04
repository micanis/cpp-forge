package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// TestCase holds a sample input/output pair.
type TestCase struct {
	Input  string
	Output string
}

// Problem holds the extracted data from a problem page.
type Problem struct {
	Title     string
	Statement string
	Cases     []TestCase
}

func main() {
	cookiePath := flag.String("cookies", "", "Netscape-format cookie file (default: ~/.config/forge/cookies.txt)")
	outDir := flag.String("dir", ".", "Directory to save test cases and problem statement")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: forge-dl [--cookies FILE] [--dir DIR] URL")
		os.Exit(1)
	}

	problemURL := flag.Arg(0)

	cookies, err := loadCookies(*cookiePath, problemURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[forge-dl] warning: could not load cookies: %v\n", err)
	}

	var prob *Problem
	switch detectJudge(problemURL) {
	case "atcoder":
		prob, err = scrapeAtCoder(problemURL, cookies)
	case "aoj":
		prob, err = scrapeAOJ(problemURL, cookies)
	default:
		prob, err = scrapeGeneric(problemURL, cookies)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "[forge-dl] error: %v\n", err)
		os.Exit(1)
	}

	if err := save(prob, *outDir); err != nil {
		fmt.Fprintf(os.Stderr, "[forge-dl] error saving: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[forge-dl] %d test case(s) saved to %s/test/\n", len(prob.Cases), *outDir)
	if prob.Statement != "" {
		fmt.Printf("[forge-dl] problem statement saved to %s/problem.txt\n", *outDir)
	}
}

// detectJudge returns the online judge name inferred from the URL.
func detectJudge(rawURL string) string {
	switch {
	case strings.Contains(rawURL, "atcoder.jp"):
		return "atcoder"
	case strings.Contains(rawURL, "u-aizu.ac.jp"):
		return "aoj"
	default:
		return "generic"
	}
}

// save writes test cases and the problem statement to outDir.
func save(prob *Problem, outDir string) error {
	testDir := filepath.Join(outDir, "test")
	if err := os.MkdirAll(testDir, 0o755); err != nil {
		return err
	}

	for i, tc := range prob.Cases {
		inFile := filepath.Join(testDir, fmt.Sprintf("sample-%d.in", i+1))
		outFile := filepath.Join(testDir, fmt.Sprintf("sample-%d.out", i+1))
		if err := os.WriteFile(inFile, []byte(tc.Input), 0o644); err != nil {
			return err
		}
		if err := os.WriteFile(outFile, []byte(tc.Output), 0o644); err != nil {
			return err
		}
	}

	if prob.Statement != "" {
		stmtFile := filepath.Join(outDir, "problem.txt")
		if err := os.WriteFile(stmtFile, []byte(prob.Statement), 0o644); err != nil {
			return err
		}
	}
	return nil
}

// loadCookies reads a Netscape-format cookie file and returns cookies matching
// the domain of targetURL. Falls back to ~/.config/forge/cookies.txt if path
// is empty.
func loadCookies(path, targetURL string) ([]*http.Cookie, error) {
	if path == "" {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, ".config", "forge", "cookies.txt")
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}
	domain := u.Hostname()

	var cookies []*http.Cookie
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Netscape format: domain TAB flag TAB path TAB secure TAB expiry TAB name TAB value
		fields := strings.Split(line, "\t")
		if len(fields) < 7 {
			continue
		}
		cookieDomain := strings.TrimPrefix(fields[0], ".")
		if !strings.HasSuffix(domain, cookieDomain) && domain != cookieDomain {
			continue
		}
		cookies = append(cookies, &http.Cookie{
			Name:  fields[5],
			Value: fields[6],
		})
	}
	return cookies, scanner.Err()
}

// fetch performs a GET request with session cookies attached.
func fetch(rawURL string, cookies []*http.Cookie) (*http.Response, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, err
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "ja,en;q=0.9")

	client := &http.Client{Timeout: 30 * time.Second}
	return client.Do(req)
}

// ---------- HTML helpers ----------

func findNode(root *html.Node, pred func(*html.Node) bool) *html.Node {
	if pred(root) {
		return root
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if n := findNode(c, pred); n != nil {
			return n
		}
	}
	return nil
}

func hasAttr(n *html.Node, key, val string) bool {
	for _, a := range n.Attr {
		if a.Key == key && a.Val == val {
			return true
		}
	}
	return false
}

// textContent extracts the plain text of a node tree.
func textContent(n *html.Node) string {
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		if n.Type == html.ElementNode && (n.Data == "br" || n.Data == "p") {
			b.WriteString("\n")
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.TrimSpace(b.String())
}

// findPreNear looks for a <pre> element adjacent to the given heading node:
// first among following siblings, then inside the parent element.
func findPreNear(heading *html.Node) string {
	for sib := heading.NextSibling; sib != nil; sib = sib.NextSibling {
		if sib.Type == html.ElementNode && sib.Data == "pre" {
			return textContent(sib)
		}
	}
	if heading.Parent != nil {
		if pre := findNode(heading.Parent, func(n *html.Node) bool {
			return n.Type == html.ElementNode && n.Data == "pre"
		}); pre != nil {
			return textContent(pre)
		}
	}
	return ""
}

func ensureNewline(s string) string {
	if s != "" && !strings.HasSuffix(s, "\n") {
		return s + "\n"
	}
	return s
}

// ---------- AtCoder ----------

var (
	acInputRe  = regexp.MustCompile(`(?i)(入力例|sample\s+input)\s*\d+`)
	acOutputRe = regexp.MustCompile(`(?i)(出力例|sample\s+output)\s*\d+`)
	acStmtRe   = regexp.MustCompile(`(?i)問題文|problem\s+statement`)
)

func scrapeAtCoder(pageURL string, cookies []*http.Cookie) (*Problem, error) {
	resp, err := fetch(pageURL, cookies)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse HTML: %w", err)
	}

	titleNode := findNode(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "title"
	})
	title := ""
	if titleNode != nil {
		title = textContent(titleNode)
	}

	taskDiv := findNode(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "div" && hasAttr(n, "id", "task-statement")
	})
	if taskDiv == nil {
		return nil, fmt.Errorf("could not find #task-statement — are cookies valid?")
	}

	// Prefer the Japanese section when both languages are present.
	langJA := findNode(taskDiv, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "span" && hasAttr(n, "class", "lang-ja")
	})
	root := taskDiv
	if langJA != nil {
		root = langJA
	}

	var inputs, outputs []string
	var stmtSections []string

	// Walk all headings (h2/h3) and classify them.
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "h2" || n.Data == "h3") {
			hText := textContent(n)
			preText := findPreNear(n)

			switch {
			case acInputRe.MatchString(hText):
				inputs = append(inputs, ensureNewline(preText))
			case acOutputRe.MatchString(hText):
				outputs = append(outputs, ensureNewline(preText))
			case acStmtRe.MatchString(hText):
				if n.Parent != nil {
					stmtSections = append(stmtSections, textContent(n.Parent))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(root)

	count := min(len(inputs), len(outputs))
	cases := make([]TestCase, count)
	for i := range cases {
		cases[i] = TestCase{Input: inputs[i], Output: outputs[i]}
	}

	return &Problem{
		Title:     title,
		Statement: strings.Join(stmtSections, "\n\n"),
		Cases:     cases,
	}, nil
}

// ---------- AOJ ----------

type aojHeader struct {
	ProblemID string `json:"problemId"`
	Testcases []struct {
		Serial int    `json:"serial"`
		Name   string `json:"name"`
	} `json:"testcases"`
}

type aojTestCaseData struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

func scrapeAOJ(pageURL string, cookies []*http.Cookie) (*Problem, error) {
	problemID := extractAOJID(pageURL)
	if problemID == "" {
		return nil, fmt.Errorf("could not extract AOJ problem ID from: %s", pageURL)
	}

	cases, err := fetchAOJTestCases(problemID, cookies)
	if err != nil {
		return nil, err
	}

	statement, err := fetchAOJStatement(problemID, cookies)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[forge-dl] warning: could not fetch problem statement: %v\n", err)
	}

	return &Problem{
		Title:     problemID,
		Statement: statement,
		Cases:     cases,
	}, nil
}

func extractAOJID(rawURL string) string {
	// https://onlinejudge.u-aizu.ac.jp/problems/ITP1_1_A
	if idx := strings.LastIndex(rawURL, "/problems/"); idx >= 0 {
		id := rawURL[idx+len("/problems/"):]
		return strings.Split(strings.TrimSpace(id), "?")[0]
	}
	// https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ITP1_1_A
	if u, err := url.Parse(rawURL); err == nil {
		if id := u.Query().Get("id"); id != "" {
			return id
		}
	}
	return ""
}

func fetchAOJTestCases(id string, cookies []*http.Cookie) ([]TestCase, error) {
	headerURL := fmt.Sprintf("https://judgedat.u-aizu.ac.jp/testcases/%s/header", id)
	resp, err := fetch(headerURL, cookies)
	if err != nil {
		return nil, fmt.Errorf("fetch AOJ header: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("AOJ API 403: test cases may not be public for %s", id)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("AOJ API HTTP %d", resp.StatusCode)
	}

	var header aojHeader
	if err := json.NewDecoder(resp.Body).Decode(&header); err != nil {
		return nil, fmt.Errorf("decode AOJ header: %w", err)
	}

	var cases []TestCase
	for _, tc := range header.Testcases {
		tcURL := fmt.Sprintf("https://judgedat.u-aizu.ac.jp/testcases/%s/%d", id, tc.Serial)
		r, err := fetch(tcURL, cookies)
		if err != nil {
			continue
		}
		var data aojTestCaseData
		json.NewDecoder(r.Body).Decode(&data) //nolint:errcheck
		r.Body.Close()

		if data.In == "" && data.Out == "" {
			continue
		}
		cases = append(cases, TestCase{
			Input:  ensureNewline(data.In),
			Output: ensureNewline(data.Out),
		})
	}
	return cases, nil
}

func fetchAOJStatement(id string, cookies []*http.Cookie) (string, error) {
	pageURL := fmt.Sprintf("https://onlinejudge.u-aizu.ac.jp/problems/%s", id)
	resp, err := fetch(pageURL, cookies)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	// AOJ renders content with Angular; try a few common selectors.
	for _, sel := range []struct{ attr, val string }{
		{"class", "problem-description"},
		{"id", "description"},
		{"class", "problem"},
	} {
		node := findNode(doc, func(n *html.Node) bool {
			return n.Type == html.ElementNode && hasAttr(n, sel.attr, sel.val)
		})
		if node != nil {
			return textContent(node), nil
		}
	}

	// Fallback: body text.
	if body := findNode(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "body"
	}); body != nil {
		return textContent(body), nil
	}
	return "", nil
}

// ---------- Generic ----------

func scrapeGeneric(pageURL string, cookies []*http.Cookie) (*Problem, error) {
	resp, err := fetch(pageURL, cookies)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse HTML: %w", err)
	}

	var inputs, outputs []string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "h2" || n.Data == "h3" || n.Data == "h4") {
			hText := textContent(n)
			preText := findPreNear(n)
			if preText != "" {
				switch {
				case acInputRe.MatchString(hText):
					inputs = append(inputs, ensureNewline(preText))
				case acOutputRe.MatchString(hText):
					outputs = append(outputs, ensureNewline(preText))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	count := min(len(inputs), len(outputs))
	cases := make([]TestCase, count)
	for i := range cases {
		cases[i] = TestCase{Input: inputs[i], Output: outputs[i]}
	}

	titleNode := findNode(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "title"
	})
	title := ""
	if titleNode != nil {
		title = textContent(titleNode)
	}

	bodyNode := findNode(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "body"
	})
	stmt := ""
	if bodyNode != nil {
		stmt = textContent(bodyNode)
	}

	return &Problem{Title: title, Statement: stmt, Cases: cases}, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
