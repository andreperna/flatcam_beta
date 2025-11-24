package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Caminho absoluto do próprio executável
	exePath, _ := os.Executable()
	baseDir := filepath.Dir(exePath)

	pythonPath := filepath.Join(baseDir, ".venv311", "Scripts", "python.exe")
	scriptPath := filepath.Join(baseDir, "flatcam.py")

	cmd := exec.Command(pythonPath, scriptPath)

	// Janela oculta no Windows
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	HideWindow: true,
	// }

	// Não herda console
	// cmd.Stdout = nil
	// cmd.Stderr = nil
	// cmd.Stdin = nil

	// Herda console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	_ = cmd.Start() //Janela oculta
}
