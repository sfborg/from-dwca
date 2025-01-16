package config

import coldpcfg "github.com/sfborg/from-coldp/pkg/config"

func (c Config) ToColdpConfig() coldpcfg.Config {
	res := coldpcfg.New()
	res.GitRepo = c.GitRepo
	res.TempRepoDir = c.TempRepoDir
	res.CacheSfgaDir = c.CacheSfgaDir
	res.JobsNum = c.JobsNum
	res.BatchSize = c.BatchSize
	res.BadRow = c.BadRow
	return res
}
