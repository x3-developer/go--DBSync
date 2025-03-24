package main

import (
	"dbsync/config"
	"dbsync/internal/dumper"
	"fmt"
	"log"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		panic(cfgErr)
	}

	fmt.Println("Доступные проекты:")
	for i, pair := range cfg.Pairs {
		fmt.Printf("%d: %s\n", i, pair.Alias)
	}

	fmt.Print("Введите номер проекта, который хотите синхронизировать: ")
	var input string
	_, scanErr := fmt.Scanln(&input)
	if scanErr != nil {
		panic(scanErr)
	}

	var index int
	if _, inputErr := fmt.Sscanf(input, "%d", &index); inputErr != nil || index < 0 || index > len(cfg.Pairs) {
		log.Fatalf("некорректный ввод: ожидается число от 0 до %d", len(cfg.Pairs)-1)
	}

	createDumpResult, createErr := dumper.DBCreateDump(cfg.Pairs[index])
	if createErr != nil || !createDumpResult {
		log.Fatalf("неудалось создать дамп: %s", createErr)
	}

	restoreDumpResult, restoreErr := dumper.DBRestoreDump(cfg.Pairs[index])
	if restoreErr != nil || !restoreDumpResult {
		log.Fatalf("не удалось развернуть дамп: %v", restoreErr)
	}
}
