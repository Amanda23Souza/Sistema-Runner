package command

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// localJDKDir retorna o diretório de cache local do JDK provisionado.
func localJDKDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".assinador", "jdk")
}

// localJDKBin retorna o caminho do binário java no JDK local.
func localJDKBin() string {
	bin := "java"
	if runtime.GOOS == "windows" {
		bin = "java.exe"
	}
	return filepath.Join(localJDKDir(), "bin", bin)
}

// downloadAndProvisionJDK baixa o JDK 21 da Adoptium e extrai em ~/.assinador/jdk.
func downloadAndProvisionJDK(out io.Writer) error {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	adoptiumOS := goos
	if adoptiumOS == "darwin" {
		adoptiumOS = "mac"
	}

	adoptiumArch := "x64"
	if goarch == "arm64" {
		adoptiumArch = "aarch64"
	}

	ext := "tar.gz"
	if goos == "windows" {
		ext = "zip"
	}

	url := fmt.Sprintf(
		"https://api.adoptium.net/v3/binary/latest/21/ga/%s/%s/jdk/hotspot/normal/eclipse?project=jdk",
		adoptiumOS, adoptiumArch,
	)

	fmt.Fprintf(out, "Baixando JDK 21 da Adoptium (%s/%s) — pode demorar alguns minutos...\n", adoptiumOS, adoptiumArch)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("erro ao preparar download: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao baixar JDK de %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Adoptium retornou status %d para %s/%s", resp.StatusCode, adoptiumOS, adoptiumArch)
	}

	tmpFile, err := os.CreateTemp("", "jdk-download-*."+ext)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	written, err := io.Copy(tmpFile, resp.Body)
	tmpFile.Close()
	if err != nil {
		return fmt.Errorf("erro ao salvar JDK (%d bytes escritos): %w", written, err)
	}
	fmt.Fprintf(out, "Download concluído (%d MB). Extraindo...\n", written/1024/1024)

	jdkDir := localJDKDir()
	if err := os.MkdirAll(jdkDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório JDK %q: %w", jdkDir, err)
	}

	if ext == "zip" {
		return extractZip(tmpFile.Name(), jdkDir, out)
	}
	return extractTarGZ(tmpFile.Name(), jdkDir, out)
}

// extractTarGZ extrai um .tar.gz descartando o primeiro nível de diretório.
func extractTarGZ(archivePath, destDir string, out io.Writer) error {
	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("erro ao ler tar: %w", err)
		}

		// Descarta o primeiro componente do caminho (ex: "jdk-21.0.5+11/")
		idx := strings.Index(hdr.Name, "/")
		if idx < 0 {
			continue
		}
		rel := hdr.Name[idx+1:]
		if rel == "" {
			continue
		}

		target := filepath.Join(destDir, filepath.FromSlash(rel))

		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(hdr.Mode)|0111); err != nil {
				return err
			}
		case tar.TypeReg, tar.TypeRegA:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}
			w, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(hdr.Mode)|0644)
			if err != nil {
				return err
			}
			if _, err := io.Copy(w, tr); err != nil {
				w.Close()
				return err
			}
			w.Close()
		case tar.TypeSymlink:
			os.Remove(target)
			_ = os.Symlink(hdr.Linkname, target)
		case tar.TypeLink:
			// Hard link — resolve dentro do destino
			idx2 := strings.Index(hdr.Linkname, "/")
			if idx2 >= 0 {
				src := filepath.Join(destDir, filepath.FromSlash(hdr.Linkname[idx2+1:]))
				_ = os.Link(src, target)
			}
		}
	}

	fmt.Fprintf(out, "JDK extraído para: %s\n", destDir)
	return nil
}

// extractZip extrai um .zip descartando o primeiro nível de diretório (Windows).
func extractZip(archivePath, destDir string, out io.Writer) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return fmt.Errorf("erro ao abrir zip: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		idx := strings.Index(f.Name, "/")
		if idx < 0 {
			continue
		}
		rel := f.Name[idx+1:]
		if rel == "" {
			continue
		}

		target := filepath.Join(destDir, filepath.FromSlash(rel))

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(target, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		w, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			rc.Close()
			return err
		}

		_, err = io.Copy(w, rc)
		w.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(out, "JDK extraído para: %s\n", destDir)
	return nil
}
