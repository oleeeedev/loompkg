package lib

import "loompkg/formulas"

type Tool struct {
	Name    string
	Install func()
	Version string
}

var tools = []Tool{
	{
		Name:    "awscli",
		Install: formulas.InstallAws,
		Version: "2.13.12",
	},
	{
		Name:    "nmap",
		Install: formulas.InstallNmap,
		Version: "7.93",
	},
	{
		Name:    "nuclei",
		Install: formulas.InstallNuclei,
		Version: "2.9.8",
	},
	{
		Name:    "k9s",
		Install: formulas.InstallK9s,
		Version: "0.27.0",
	},
	{
		Name:    "git",
		Install: formulas.InstallGit,
		Version: "2.41.0",
	},
	{
		Name:    "go",
		Install: formulas.InstallGo,
		Version: "1.21.4",
	},
	{
		Name:    "aircrack-ng",
		Install: formulas.InstallAircrackNg,
		Version: "1.6",
	},
	{
		Name:    "node",
		Install: formulas.InstallNode,
		Version: "20.5.1",
	},
	{
		Name:    "docker",
		Install: formulas.InstallDocker,
		Version: "24.0.6",
	},
	{
		Name:    "python3",
		Install: formulas.InstallPython3,
		Version: "3.11.6",
	},
	{
		Name:    "python2",
		Install: formulas.InstallPython2,
		Version: "2.7.18",
	},
	{
		Name:    "kubectl",
		Install: formulas.InstallKubectl,
		Version: "1.28.3",
	},
	{
		Name:    "rust",
		Install: formulas.InstallCargo,
		Version: "1.69.0",
	},
	{
		Name:    "java21",
		Install: formulas.InstallJava21,
		Version: "21.0.1",
	},
	{
		Name:    "dotnet",
		Install: formulas.InstallDotnet,
		Version: "8.0.100",
	},
	{
		Name:    "bazel",
		Install: formulas.InstallBazel,
		Version: "7.0.0",
	},
	{
		Name:    "bat",
		Install: formulas.InstallBat,
		Version: "0.24.0",
	},
}
