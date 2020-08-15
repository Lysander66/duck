# duck
sql to struct with gorm tag

## usage
`$ duck -h`
```
Specify the table name, generate struct according to the field

Usage:
  duck [flags]
  duck [command]

Available Commands:
  help        Help about any command
  init        Generate configuration file

Flags:
      --config string   config file (default is ./duck_config.toml)
  -h, --help            help for duck
  -n, --name string     table_name (default "country")
  -s, --schema string   table_schema (default "world")
```

e.g.

`$ duck -s world -n country`
```go
type Country struct {
    Code              string    `gorm:"column:Code"`
    Name              string    `gorm:"column:Name"`
    Continent         string    `gorm:"column:Continent"`
    Region            string    `gorm:"column:Region"`
    SurfaceArea       string    `gorm:"column:SurfaceArea"`
    IndepYear         int32     `gorm:"column:IndepYear"`
    Population        int64     `gorm:"column:Population"`
    LifeExpectancy    string    `gorm:"column:LifeExpectancy"`
    GNP               string    `gorm:"column:GNP"`
    GNPOld            string    `gorm:"column:GNPOld"`
    LocalName         string    `gorm:"column:LocalName"`
    GovernmentForm    string    `gorm:"column:GovernmentForm"`
    HeadOfState       string    `gorm:"column:HeadOfState"`
    Capital           int64     `gorm:"column:Capital"`
    Code2             string    `gorm:"column:Code2"`
}

func (t *Country) TableName() string {
	return "country"
}
```
