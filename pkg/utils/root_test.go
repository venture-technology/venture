package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindGoModRoot(t *testing.T) {
	t.Run("encontra go.mod no diretório atual", func(t *testing.T) {
		tmp := t.TempDir()

		goModPath := filepath.Join(tmp, "go.mod")
		err := os.WriteFile(goModPath, []byte("module test"), 0644)
		if err != nil {
			t.Fatalf("falha ao criar go.mod: %v", err)
		}

		originalWD, _ := os.Getwd()
		defer os.Chdir(originalWD)
		os.Chdir(tmp)

		root, err := FindGoModRoot()
		if err != nil {
			t.Fatalf("esperava encontrar go.mod, mas deu erro: %v", err)
		}
		if root != tmp {
			t.Errorf("esperado %s, mas obteve %s", tmp, root)
		}
	})

	t.Run("encontra go.mod em um diretório pai", func(t *testing.T) {
		tmp := t.TempDir()

		subdir := filepath.Join(tmp, "a", "b", "c")
		err := os.MkdirAll(subdir, 0755)
		if err != nil {
			t.Fatalf("falha ao criar subdiretórios: %v", err)
		}

		goModPath := filepath.Join(tmp, "go.mod")
		err = os.WriteFile(goModPath, []byte("module test"), 0644)
		if err != nil {
			t.Fatalf("falha ao criar go.mod: %v", err)
		}

		originalWD, _ := os.Getwd()
		defer os.Chdir(originalWD)
		os.Chdir(subdir)

		root, err := FindGoModRoot()
		if err != nil {
			t.Fatalf("esperava encontrar go.mod, mas deu erro: %v", err)
		}
		if root != tmp {
			t.Errorf("esperado %s, mas obteve %s", tmp, root)
		}
	})

	t.Run("retorna erro quando go.mod não é encontrado", func(t *testing.T) {
		tmp := t.TempDir()

		originalWD, _ := os.Getwd()
		defer os.Chdir(originalWD)
		os.Chdir(tmp)

		_, err := FindGoModRoot()
		if err == nil {
			t.Fatal("esperava erro ao não encontrar go.mod, mas não ocorreu erro")
		}
	})
}
