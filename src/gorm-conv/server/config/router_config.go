package config

import (
	"fmt"
	"gorm-conv/common"
	"os"
)

func GeneralRoute(game common.XmlCfg, f *os.File) int {
	f.WriteString("table_route:\n")
	for _, table := range game.DB.Routes.TableList {
		f.WriteString("  - table_name: " + table.Name + "\n")
		f.WriteString("    router:\n")
		for _, db := range table.RoutesTableDBList {
			f.WriteString("      - " + db.Name + ":" + db.SplitNum + "\n")
		}
	}

	return 0
}

func GeneralDatabase(game common.XmlCfg, f *os.File) int {
	f.WriteString("databases:\n")
	for _, db := range game.DB.DBList.DBList {
		f.WriteString("  - name: " + db.Name + "\n")
		f.WriteString("    type: " + db.Type + "\n")
		f.WriteString("    host: " + db.Host + "\n")
		f.WriteString("    port: " + db.Port + "\n")
		f.WriteString("    user: " + db.User + "\n")
		f.WriteString("    passwd: " + db.PassWD + "\n")
		f.WriteString("    database: " + db.Database + "\n")
	}
	return 0
}

func General(game common.XmlCfg, outpath string) int {
	var ymlPath string = game.File[:len(game.File)-3]
	outfile := outpath + ymlPath + "yml"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	fmt.Println("begin to general route file:" + outfile)
	defer func() {
		f.Close()
		fmt.Println("end general golang codes file:" + outfile)
	}()
	f.Truncate(0)

	GeneralDatabase(game, f)

	return GeneralRoute(game, f)
}

func GeneralDBCfg(games []common.XmlCfg, outpath string) int {
	for _, game := range games {
		if 0 != General(game, outpath) {
			fmt.Println("general route file failed:", game.File)
			return -1
		}
	}

	return 0
}
