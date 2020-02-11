// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 21:25

package common

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {

	// pflag
	pflag.Int("server.port", 8080, "ginchat server listen port")
	pflag.String("server.ip", "127.0.0.1", "ginchat server ip")
	pflag.String("server.log_path", "/var/log/", "server log path")
	pflag.Bool("server.log_stdout", false, "HTTP log output stdout")
	pflag.String("server.attach_path", "./mnt", "server upload path")
	pflag.Int("server.mode", 0, "server mode:0 INFINITY、 1 minio、2 INFINITY & minio")
	pflag.Bool("server.debug", false, "server debug")
	pflag.String("mysql.ip", "127.0.0.1", "mysql ip")
	pflag.Int("mysql.port", 3306, "mysql port")
	pflag.String("mysql.user", "root", "mysql user")
	pflag.String("mysql.pwd", "root", "mysql pwd")
	pflag.String("mysql.db_name", "chat", "mysql db_name")
	pflag.Bool("mysql.show_sql", false, "mysql show_sql")
	pflag.Int("mysql.max_open_conns", 2, "mysql max_open_conns")
	var configFile string
	pflag.StringVar(&configFile, "config", "./common/config.yaml", "platform server config file")
	pflag.Parse()

	if configFile != "" {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil && !os.IsNotExist(err) {
			log.Fatal("Fatal error config file: ", err)
		}
	}
	// bind pflag
	viper.BindPFlags(pflag.CommandLine)
	// bind env
	viper.SetEnvPrefix("ginchat")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// handle config
	ServerPort = viper.GetInt("server.port")
	ServerIp = viper.GetString("server.ip")
	ServerMode = viper.GetInt("server.mode")
	ServerDebug = viper.GetBool("server.debug")
	LogPath = viper.GetString("server.log_path")
	LogStdout = viper.GetBool("server.log_stdout")
	AttachPath = viper.GetString("server.attach_path")
	MysqlIp = viper.GetString("mysql.ip")
	MysqlPort = viper.GetInt("mysql.port")
	MysqlUser = viper.GetString("mysql.user")
	MysqlPwd = viper.GetString("mysql.pwd")
	DBName = viper.GetString("mysql.db_name")
	ShowSQL = viper.GetBool("mysql.show_sql")
	MaxOpenConns = viper.GetInt("mysql.max_open_conns")
}
