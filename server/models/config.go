package models

import "fmt"

type Config struct {
    DbHost string
    DbUser string
    DbPassword string
    DbName string
}

func NewConfig(DbHost string,
                DbUser string,
                DbPassword string,
                DbName string) Config {
    return Config{
        DbHost: DbHost,
        DbUser: DbUser,
        DbPassword: DbPassword,
        DbName: DbName,
    }
}

func (c *Config) GetDbConnOpts(extraOpts string) string {
    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s %s",
        c.DbHost,
        c.DbUser,
        c.DbPassword,
        c.DbName,
        extraOpts,
    )
}
