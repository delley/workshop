package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var (
	reDayDir  = regexp.MustCompile(`^dia_\d{2}$`)
	reHeading = regexp.MustCompile(`^(#{1,6})\s+(.+?)\s*$`)
	reSpace   = regexp.MustCompile(`\s+`)
	reHyphens = regexp.MustCompile(`-{2,}`)

	projectOut = "README.md"
)

type Heading struct {
	Level int
	Title string
}

// slugger implementa o comportamento do GitHub para headings duplicados:
// primeiro uso: "titulo"
// duplicados: "titulo-1", "titulo-2", ...
type slugger struct {
	seen map[string]int
}

func newSlugger() *slugger {
	return &slugger{seen: make(map[string]int)}
}

func (s *slugger) Anchor(title string) string {
	base := ghSlugBase(title)
	if base == "" {
		base = "section"
	}
	// GitHub: primeiro uso sem sufixo; a partir do segundo, -1, -2, ...
	if n, ok := s.seen[base]; ok {
		n++
		s.seen[base] = n
		return fmt.Sprintf("%s-%d", base, n)
	}
	s.seen[base] = 0
	return base
}

func main() {
	root := "."
	dayDirs, err := listDayDirs(root)
	if err != nil {
		fatalf("falha ao listar diretórios dia_XX: %v", err)
	}

	var lines []string
	lines = append(lines, "# Workshop Go: do Zero à API", "")

	globalIdx := 1

	for _, dayDir := range dayDirs {
		dayLabel := dayLabelFromDir(dayDir)
		// Âncora do H1 "# Dia XX" dentro do README do dia
		dayAnchor := ghSlugBase(dayLabel)
		dayLink := fmt.Sprintf("../../tree/main/%s#%s", dayDir, dayAnchor)

		lines = append(lines, fmt.Sprintf("## [%s](%s)", dayLabel, dayLink), "")

		dayReadme := filepath.Join(root, dayDir, "README.md")
		if _, err := os.Stat(dayReadme); err != nil {
			if os.IsNotExist(err) {
				lines = append(lines, "_README.md não encontrado para este dia._", "")
				continue
			}
			fatalf("erro ao acessar %s: %v", dayReadme, err)
		}

		headings, err := readHeadings(dayReadme)
		if err != nil {
			fatalf("falha ao ler headings de %s: %v", dayReadme, err)
		}

		// Slugger por arquivo (igual ao GitHub)
		slugs := newSlugger()

		hasCurrentItem := false
		subIdx := 0

		for _, h := range headings {
			switch h.Level {
			case 1:
				// Ainda assim, registramos o H1 no slugger para garantir
				// consistência caso alguém crie links para ele.
				_ = slugs.Anchor(h.Title)
				continue

			case 2:
				anchor := slugs.Anchor(h.Title)
				link := fmt.Sprintf("../../tree/main/%s#%s", dayDir, anchor)
				lines = append(lines, fmt.Sprintf("%d. [%s](%s)", globalIdx, h.Title, link), "")
				globalIdx++
				subIdx = 0
				hasCurrentItem = true

			case 3:
				if !hasCurrentItem {
					// Ainda assim registra para manter contagem correta em duplicados
					_ = slugs.Anchor(h.Title)
					continue
				}
				subIdx++
				anchor := slugs.Anchor(h.Title)
				link := fmt.Sprintf("../../tree/main/%s#%s", dayDir, anchor)
				lines = append(lines, fmt.Sprintf("    %d. [%s](%s)", subIdx, h.Title, link))

			default:
				// Opcional: registrar para duplicados em níveis mais profundos
				_ = slugs.Anchor(h.Title)
				continue
			}
		}

		lines = append(lines, "")
	}

	out := strings.TrimRight(strings.Join(lines, "\n"), "\n") + "\n"
	if err := os.WriteFile(projectOut, []byte(out), 0o644); err != nil {
		fatalf("falha ao escrever %s: %v", projectOut, err)
	}

	fmt.Println("README.md atualizado com sucesso.")
}

func listDayDirs(root string) ([]string, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var dayDirs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if reDayDir.MatchString(name) {
			dayDirs = append(dayDirs, name)
		}
	}

	sort.Strings(dayDirs)
	return dayDirs, nil
}

func readHeadings(mdPath string) ([]Heading, error) {
	f, err := os.Open(mdPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var headings []Heading
	sc := bufio.NewScanner(f)

	inCodeBlock := false

	for sc.Scan() {
		line := sc.Text()
		trim := strings.TrimSpace(line)

		// Toggle bloco de código quando linha começa com ```
		if strings.HasPrefix(trim, "```") {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock {
			continue
		}

		m := reHeading.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		level := len(m[1])
		title := strings.TrimSpace(m[2])
		headings = append(headings, Heading{Level: level, Title: title})
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}
	return headings, nil
}

func dayLabelFromDir(dayDir string) string {
	// dia_02 -> "DIA 02"
	parts := strings.Split(dayDir, "_")
	n := ""
	if len(parts) == 2 {
		n = parts[1]
	}
	return "DIA " + n
}

func ghSlugBase(title string) string {
	// GitHub-like base:
	// - lowercase
	// - mantém acentos
	// - remove pontuação/símbolos
	// - espaços -> hífen
	// - colapsa hífens
	s := strings.ToLower(strings.TrimSpace(title))

	// Mantém: letras (incluindo acentuadas), dígitos, underscore, espaço, hífen
	var buf bytes.Buffer
	for _, r := range s {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			buf.WriteRune(r)
		case r == '_' || r == ' ' || r == '-':
			buf.WriteRune(r)
		default:
			// descarta pontuação/símbolos
		}
	}

	s = buf.String()
	s = reSpace.ReplaceAllString(s, "-")
	s = reHyphens.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
