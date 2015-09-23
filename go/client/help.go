package client

import (
	"github.com/keybase/cli"
)

func GetHelpTopics() []cli.HelpTopic {
	return []cli.HelpTopic{
		advancedHT,
		gpgHT,
	}
}

var advancedHT = cli.HelpTopic{
	Name:  "advanced",
	Usage: "description of advanced global options",
	Body: `Keybase commands can run with various command line flags to
configure global options.  To use, add the flags before the command name:

   {{ .Name }} [global options] command [command options]

GLOBAL OPTIONS:
   {{range .Flags}}{{ . }}
   {{end}}
	`,
}

var gpgHT = cli.HelpTopic{
	Name:  "gpg",
	Usage: "description of how keybase interacts with GPG",
	Body: `The Keybase client makes selective use of a local installation
of GnuPG (gpg) if it can find one. We realize that access to local gpg is
sensitive, so we try to tread lightly and only access the local gpg keyrings
in a few special cases.

On Device Provisioning - If you've previously set up a keybase account and host
  your secret key via gpg locally, the keybase app will need a signature signed with
  that key to delegate authority to a new local device key.  That device key will then
  do all of the non-PGP work for this machine going forward (e.g., signing tracker
  statements, provisioning new devices, etc.)

On signup - When you signup, you can also "select" a PGP key for use with keybase
  from your local GPG keyring.  See "keybase pgp select" below.

'keybase pgp gen' - Running "gen" will generate a new PGP key via the Go libraries.
  keybase will save the secret key locally, and will push the public half to the
  server, after it's been signed and provisioned with your local device key.
  Additionally, this command will export the secret and public keys to your local
  GPG keyring.

'keybase pgp pull' - All PGP keys pulled as a result of 'keybase pgp pull' are
  exported as public keys to your local gpg keyring.

'keybase pgp update' - Public PGP keys are exported from your local GPG keyring and
  sent to the keybase server, where they will supersede PGP keys that have been
  previously updated.  This feature is for updating PGP subkeys, identities, and
  signatures, but cannot be used to change PGP primary keys.

'keybase pgp select' - Pull a PGP key out of your GPG keyring, and "select" it
  for use on keybase.  This will: (1) sign it into your signature chain with your
  local device key; (2) copy a version of secret key to your local keybase keychain;
  and (3) encrypt this copy with your keybase passphrase via local-key security
  (see separate help topic). keybase takes steps (2) and (3) so that subsequent
  PGP operations (like sign and decrypt) don't need to access your GPG keyring again.
  Once running this command, you wind up in state similar to that following
  "keybase pgp gen" above.  The difference is that you've used keybase to generate
  the key rather than gpg.

'keybase pgp sign/encrypt/verify/decrypt' - These commands don't access your local
  GPG keyring, they only access the local keybase keyring as described in "select".

'keybase pgp import' - Doesn't access the GPG keyring at all, only works with
  keys passed in via standard input or files. A secret key is required for this
  operation, since a signature for the key is needed in the import process.

'keybase pgp export' - Doesn't access the GPG keyring at all, just outputs
  your currently provisioned PGP keys to standard output (or a file).
`,
}

// Custom help templates for cli package

func init() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.SubcommandHelpTemplate = SubcommandHelpTemplate
}

// AppHelpTemplate is used for `keybase help` or `keybase -h`.
var AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} [arguments...]
   {{if .Version}}
VERSION:
   {{.Version}}
   {{end}}{{if .Commands}}
COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .HelpTopics}}
ADDITIONAL HELP TOPICS:
   {{range .HelpTopics}}{{.Name}}{{ "\t\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}
`

// CommandHelpTemplate is used for `keybase help cmd` or
// `keybase cmd help subcmd`.
var CommandHelpTemplate = `NAME:
   keybase {{.FullName}} - {{.Usage}}

USAGE:
   keybase {{.FullName}}{{ if .Subcommands }} <command>{{ end }}{{if .Flags}} [command options]{{end}} {{ .ArgumentHelp }}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{ if .Subcommands }}

COMMANDS:
   {{range .Subcommands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`

// SubcommandHelpTemplate is used for `keybase cmd` with no
// other arguments when `cmd` has subcommands.
// Or for `keybase cmd help` when `cmd` has subcommands.
var SubcommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} <command> [arguments...]

COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}
`
