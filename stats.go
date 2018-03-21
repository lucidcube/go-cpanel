package cpanel

// StatCollection is an array of StatType
type StatCollection []StatType

// QueryValue returns contained values as a single query parameter
func (sc StatCollection) QueryValue() string {
	if len(sc) == 0 {
		return ""
	}
	var (
		sep = []byte("|")
		out = make([]byte, 0, (1+len(sep))*len(sc))
	)
	for _, s := range sc {
		out = append(out, s...)
		out = append(out, sep...)
	}
	return string(out[:len(out)-len(sep)])
}

// StatType represents a type of account statistic
type StatType string

const (
	// FTPAccounts is the stat type for information about the account's FTP accounts.
	FTPAccounts StatType = "ftpaccounts"

	// PerlVersion is the stat type for the server's Perl version.
	PerlVersion StatType = "perlversion"

	// DedicatedIP is the stat type for cPanel account websites that use dedicated IP addresses.
	DedicatedIP StatType = "dedicatedip"

	// Hostname  is the stat type for The server's hostname.
	Hostname StatType = "hostname"

	// OperatingSystem is the stat type for The server's operating system
	OperatingSystem StatType = "operatingsystem"

	// SendmailPath is the stat type for The path to the system'ssendmail binary.
	SendmailPath StatType = "sendmailpath"

	// AutoResponders is the stat type for Information about the cPanel account's auto-responders.
	AutoResponders StatType = "autoresponders"

	// PerlPath is the stat type for The Perl binary's absolute oath.
	PerlPath StatType = "perlpath"

	// EmailForwarders is the stat type for Information about the cPanel account's forwarders.
	EmailForwarders StatType = "emailforwarders"

	// BandwidthUsage is the stat type for Information about the account's bandwidth usage.
	BandwidthUsage StatType = "bandwidthusage"

	// EmailFilters is the stat type for Information the cPanel account's email filters.
	EmailFilters StatType = "emailfilters"

	// MailingLists is the stat type for Information the cPanel account's mailing lists.
	MailingLists StatType = "mailinglists"

	// DiskUsage is the stat type for Information the account's disk space usage.
	DiskUsage StatType = "diskusage"

	// PHPVersion is the stat type for The server's PHP version.
	PHPVersion StatType = "phpversion"

	// SQLDatabases is the stat type for Information about the account's SQL databases.
	SQLDatabases StatType = "sqldatabases"

	// ApacheVersion is the stat type for The server's Apache version.
	ApacheVersion StatType = "apacheversion"

	// KernelVersion is the stat type for The operating system's kernel version.
	KernelVersion StatType = "kernelversion"

	// ShortHostname is the stat type for The short version of your server's hostname.
	ShortHostname StatType = "shorthostname"

	// ParkedDomains is the stat type for Information about the cPanel account's parked domains (aliases).
	ParkedDomains StatType = "parkeddomains"

	// CPanelBuild is the stat type for The server's cPanel build.
	CPanelBuild StatType = "cpanelbuild"

	// Theme is the stat type for The cPanel account's theme.
	Theme StatType = "theme"

	// AddonDomains is the stat type for Information about the cPanel account's addon domains.
	AddonDomains StatType = "addondomains"

	// CPanelRevision is the stat type for The build of cPanel that runs on the server.
	CPanelRevision StatType = "cpanelrevision"

	// MachineType is the stat type for The type of operating system your server runs.
	MachineType StatType = "machinetype"

	// CPanelVersion is the stat type for The server's cPanel version.
	CPanelVersion StatType = "cpanelversion"

	// MySQLDiskUsage is the stat type for The amount of disk space that the cPanel account's MySQL® databases use.
	MySQLDiskUsage StatType = "mysqldiskusage"

	// MySQLVersion is the stat type for The server's MySQL version.
	MySQLVersion StatType = "mysqlversion"

	// Subdomains is the stat type for Information about the cPanel account's subdomains.
	Subdomains StatType = "subdomains"

	// PostgresDiskUsage is the stat type for The amount of disk space that the cPanel account's PostgreSQL databases use.
	PostgresDiskUsage StatType = "postgresdiskusage"

	// SharedIP is the stat type for Any of the cPanel account's websites that use a shared IP address.
	SharedIP StatType = "sharedip"

	// HostingPackage is the stat type for The cPanel account's hosting package.
	HostingPackage StatType = "hostingpackage"

	// EmailAccounts is the stat type for Information about the cPanel account's email accounts.
	EmailAccounts StatType = "emailaccounts"
)
