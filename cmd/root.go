/*
Copyright © 2022 Bora Kasmer <bora@borakasmer.com>

*/
//go get github.com/olekukonko/tablewriter
package cmd

import (
	"os"
	"runtime"
	"time"

	"github.com/borakasmer/fuel/parser"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
)

const _DATE_FORMAT_STRING = "2006.01.02 15:04:05"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fuel",
	Short: "Bu Cli Tool'u ile, PO'dan güncel petrol fiyatları İstanbul, Ankara ve İzmir için çekilir.",
	Long: `fuel

**** Herhangi bir tanımlı parametre yoktur.
------------------
Petrol Ofisi sitesi	üzerindeki fiyat listesi, anlık olarak Parse Edilerek ekrana, İl bazında basılır.
Örnek kullanım:
."fuel"
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		getFuel()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cityList []string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fuel.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringSliceVarP(&cityList, "city", "c", []string{"Ankara", "Istanbul", "Izmir"}, "City names seperated by comma")
}

func getFuel() {
	tableHeaders := make([]string, 0)
	tableRows := make([][]string, 0)

	tableHeaders = append(tableHeaders, "Yakıt Tipi")

	petrolRow := make([]string, 0)
	petrolRow = append(petrolRow, "Benzin =>")
	dieselRow := make([]string, 0)
	dieselRow = append(dieselRow, "Mazot =>")
	lpgRow := make([]string, 0)
	lpgRow = append(lpgRow, "LPG =>")

	// Repeating color pattern for table header
	var colorsPattern = []tablewriter.Colors{{
		tablewriter.Bold, tablewriter.BgGreenColor,
	},
		{
			tablewriter.Bold, tablewriter.BgYellowColor,
		},
		{
			tablewriter.Bold, tablewriter.BgBlueColor,
		}}

	colors := make([]tablewriter.Colors, 0)

	// Yakıt Tipi cell's Color
	colors = append(colors, tablewriter.Colors{
		tablewriter.Bold, tablewriter.BgMagentaColor,
	})

	for i, c := range cityList {
		item := parser.ParseWeb(c)
		tableHeaders = append(tableHeaders, item.City)
		petrolRow = append(petrolRow, item.Petrol+"₺")
		dieselRow = append(dieselRow, item.Diesel+"₺")
		lpgRow = append(lpgRow, item.Lpg+"₺")

		color := colorsPattern[i%len(colorsPattern)]
		colors = append(colors, color)
	}

	// Tarih cell's color
	colors = append(colors, tablewriter.Colors{
		tablewriter.Bold, tablewriter.BgRedColor,
	})

	tableHeaders = append(tableHeaders, "Tarih")
	petrolRow = append(petrolRow, time.Now().Format(_DATE_FORMAT_STRING))
	dieselRow = append(dieselRow, time.Now().Format(_DATE_FORMAT_STRING))
	lpgRow = append(lpgRow, time.Now().Format(_DATE_FORMAT_STRING))

	tableRows = append(tableRows, petrolRow)
	tableRows = append(tableRows, dieselRow)
	tableRows = append(tableRows, lpgRow)

	// Create Header of Table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders)
	//---------------------

	table.SetCaption(true, "Petrol Ofisinde, Şehir Bazlı Yakıt Fiyatları\n ®coderbora => www.borakasmer.com")
	table.AppendBulk(tableRows)
	// Set Table Color
	if !isWindows() { // Windows için Renkli Tablo başlıkları gözükmüyor...
		table.SetHeaderColor(colors...)
	}
	table.Render()
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
