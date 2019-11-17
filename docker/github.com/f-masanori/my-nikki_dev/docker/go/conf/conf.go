package conf

type Database_ struct {
	Drivername string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
}

var Database Database_

func Init() {
	Database.Drivername = "mysql"
	Database.Host = "mysql_container"
	Database.Port = "3306"
	Database.User = "root"
	Database.Password = "mysql"
	Database.Dbname = "app"
}

// func ReadConf() {
// 	fmt.Println("config file read error")

// 設定ファイル名を記載
// viper.SetConfigName("config")

// // ファイルタイプ
// viper.SetConfigType("yml")

// // ファイルパスの設定
// viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "f-masanori", "my-nikki_dev", "docker", "go", "conf"))

// // 環境変数から設定値を上書きできるように設定
// viper.AutomaticEnv()

// conf読み取り
// if err := viper.ReadInConfig(); err != nil {
// 	fmt.Println("config file read error")
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// // UnmarshalしてCにマッピング
// if err := viper.Unmarshal(&Conf); err != nil {
// 	fmt.Println("config file Unmarshal error")
// 	fmt.Println(err)
// 	os.Exit(1)
// }
// }
