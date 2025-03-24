package dumper

import (
	"context"
	"dbsync/config"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func DBCreateDump(pairConfig config.PairConfig) (bool, error) {
	backupFileName := getBackupFileName(pairConfig)

	if err := os.Remove(backupFileName); err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("ошибка удаления старого дампа: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"mysqldump",
		"--host", pairConfig.Source.Host,
		"--port", pairConfig.Source.Port,
		"--user", pairConfig.Source.Username,
		"--password="+pairConfig.Source.Password,
		"--result-file="+backupFileName,
		pairConfig.Source.DBName,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("ошибка создания дампа: %w\n%s", err, string(output))
	}

	return true, nil
}

func DBRestoreDump(pairConfig config.PairConfig) (bool, error) {
	backupFileName := getBackupFileName(pairConfig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"mysql",
		"--host", pairConfig.Source.Host,
		"--port", pairConfig.Source.Port,
		"--user", pairConfig.Source.Username,
		"--password="+pairConfig.Source.Password,
		pairConfig.Target.DBName,
		"-e", fmt.Sprintf("SOURCE %s;", backupFileName),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("ошибка разворачивания дампа: %w\n%s", err, string(output))
	}

	return true, nil
}

func getBackupFileName(pairConfig config.PairConfig) string {
	return fmt.Sprintf("%s_dump.sql", pairConfig.Alias)
}
