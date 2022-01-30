package dashboard

import "embed"

//go:generate yarn build

//go:embed dist/*
var data embed.FS

func DashboardAsset(p string) ([]byte, error) {
	return data.ReadFile("dist" + p)
}
