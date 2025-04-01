package init

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
)

const defaultConfigTemplate = `basePackage: {{.BasePackage}}
outputDir: {{.OutputDir}}
testUtilDir: ./test
dbDriver: {{.DBDriver}}
webFramework: {{.WebFramework}}
databasePackage: {{.DatabasePackage}}
modelTemplate: ./templates/model.tmpl
modelTestTemplate: ./templates/model_test.tmpl
fixtureTemplate: ./templates/fixture.tmpl
repositoryTemplate: ./templates/repository.tmpl
repositoryTestTemplate: ./templates/repository_test.tmpl
usecaseTemplate: ./templates/usecase.tmpl
usecaseTestTemplate: ./templates/usecase_test.tmpl
handlerTemplate: ./templates/handler.tmpl
handlerTestTemplate: ./templates/handler_test.tmpl
`

type configInput struct {
	BasePackage     string
	OutputDir       string
	DBDriver        string
	WebFramework    string
	DatabasePackage string
}

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a gon.yaml config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := promptAndWriteConfigFile(); err != nil {
			fmt.Printf("💥 Failed to generate config file: %v\n", err)
			return err
		}
		return nil
	},
}

func promptAndWriteConfigFile() error {
	path := filepath.Join(".", "gon.yaml")

	if _, err := os.Stat(path); err == nil {
		fmt.Println("📄 gon.yaml already exists. Skipping generation.")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("🔧 Enter your base package (e.g. github.com/yourname/project): ")
	basePkg, _ := reader.ReadString('\n')
	basePkg = strings.TrimSpace(basePkg)
	if basePkg == "" {
		return fmt.Errorf("❌ Invalid base package specified: %s", basePkg)
	}

	fmt.Print("📁 Output directory (default: internal/domain): ")
	outDir, _ := reader.ReadString('\n')
	outDir = strings.TrimSpace(outDir)
	if outDir == "" {
		outDir = "internal/domain"
	}

	fmt.Print("🛢️  Default database driver (gorm): ")
	driver, _ := reader.ReadString('\n')
	driver = strings.TrimSpace(driver)
	if driver == "" {
		driver = config.DBDriverGorm.String()
	}
	if !config.IsValidDBDriver(driver) {
		return fmt.Errorf("❌ Invalid DB driver specified: %s", driver)
	}

	fmt.Print("🌐 Default web framework (echo): ")
	web, _ := reader.ReadString('\n')
	web = strings.TrimSpace(web)
	if web == "" {
		web = config.WebFrameworkEcho.String()
	}
	if !config.IsValidWebFramework(web) {
		return fmt.Errorf("❌ Invalid WEB framework specified: %s", web)
	}

	dbPkg := fmt.Sprintf("%s/internal/infra/storage/database", basePkg)

	cfg := configInput{
		BasePackage:     basePkg,
		OutputDir:       outDir,
		DBDriver:        driver,
		WebFramework:    web,
		DatabasePackage: dbPkg,
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	tmpl, err := template.New("config").Parse(defaultConfigTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, cfg); err != nil {
		return err
	}

	fmt.Println("✅ Config file gon.yaml created successfully.")
	return nil
}
