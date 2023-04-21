

package v12

const (
	// UpgradeName is the shared upgrade plan name for mainnet
	UpgradeName = "v12.0.0"
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/evmos/evmos/releases/download/v12.0.0/evmos_12.0.0_Darwin_arm64.tar.gz","darwin/amd64":"https://github.com/evmos/evmos/releases/download/v12.0.0/evmos_12.0.0_Darwin_amd64.tar.gz","linux/arm64":"https://github.com/evmos/evmos/releases/download/v12.0.0/evmos_12.0.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/evmos/evmos/releases/download/v12.0.0/evmos_12.0.0_Linux_amd64.tar.gz","windows/x86_64":"https://github.com/evmos/evmos/releases/download/v12.0.0/evmos_12.0.0_Windows_x86_64.zip"}}'`
	// MaxRecover is the maximum amount of coins to be redistributed in the upgrade
	MaxRecover = "169580456887205410791936"
)