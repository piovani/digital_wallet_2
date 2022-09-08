package cmd

func migrate() {
	InitConfig()
	CheckFatal(GetDatabase().Migrate())
}
