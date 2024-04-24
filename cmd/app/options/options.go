package options

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"sample-project/cmd/app/config"
	"sample-project/pkg/controller"
	"sample-project/pkg/db"
)

const (
	defaultConfigFile = "/etc/sample/config.yaml"
)

type ServerRunOptions struct {
	// The config file value
	ComponentConfig config.Config
	// config file path
	ConfigFile string
	// http engine
	HttpEngine *gin.Engine

	// mysql interface
	Factory db.ShareDaoFactory

	// controller interface
	Control controller.SampleInterface
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		ConfigFile: defaultConfigFile,
		HttpEngine: gin.Default(),
	}

	return &s
}

// Complete read configuration
func (s *ServerRunOptions) Complete() error {
	var componentConfig config.Config

	data, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &componentConfig); err != nil {
		return err
	}

	s.ComponentConfig = componentConfig

	return nil
}

// Registry register service
func (s *ServerRunOptions) Registry() error {
	if err := s.registerDatabase(); err != nil {
		return err
	}

	s.Control = controller.New(s.ComponentConfig, s.Factory)

	return nil
}

func (s *ServerRunOptions) registerDatabase() error {
	sqlConfig := s.ComponentConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.Name)

	opt := &gorm.Config{}
	if s.ComponentConfig.Default.Mode == "debug" {
		opt.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err := gorm.Open(mysql.Open(dsn), opt)
	if err != nil {
		return err
	}

	s.Factory, err = db.NewDaoFactory(DB, s.ComponentConfig.Default.AutoMigrate)
	if err != nil {
		return err
	}

	return nil
}

// BindFlags use cobra flags
func (s *ServerRunOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&s.ConfigFile, "file", s.ConfigFile, "The location of the sample configuration file")
}
