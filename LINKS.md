## メモ

- S3からGCSへのデータ転送（AWSのIAM情報もらうパターン）

  https://note.com/ucwork/n/n5fe04382409b

- middleware を使用した mock

  https://go-to-k.hatenablog.com/entry/aws-sdk-go-v2-middleware-test

- interface を使用した mock

  https://aws.github.io/aws-sdk-go-v2/docs/migrating/

- Moto Documentation

  https://docs.getmoto.org/_/downloads/en/2.2.15/pdf/

- mitmproxy is a free and open source interactive HTTPS proxy.

  https://mitmproxy.org/

- Flutter web、amplify_flutter、Amazon Coginto、KeycloakでSAMLフェデレーション（SSO）

  https://zenn.dev/motu2119/articles/a70b611329d133

- golang 偽名住所ジェネレータ

  https://github.com/mattn/go-gimei

- mysql asdfインストール

```bash
$ asdf plugin-add mysql
$ asdf list-all mysql
$ asdf install mysql 8.0.32
$ asdf global mysql 8.0.32 
```

- mysql TCP接続

```bash
$ mysql -h localhost --port 3306 --protocol tcp -u root -ppass
```

- cobra sample

```golang
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "usage of command",
		Long:  "usage of command",
	}

	subCmd = &cobra.Command{
		Use:   "sub",
		Short: "usage of sub command",
		Long:  "usage of sub command",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cmd.PersistentFlags().Lookup("host").Value)
			fmt.Println(cmd.PersistentFlags().Lookup("port").Value)
			return nil
		},
	}

	versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of this command",
		Long:    "Print the version number of this command",
		Version: "0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("sample v" + cmd.Version)
			return nil
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	subCmd.PersistentFlags().String("host", "localhost", "hostname")
	subCmd.PersistentFlags().Int("port", 3000, "port number")
	rootCmd.AddCommand(subCmd)
	rootCmd.AddCommand(versionCmd)
}
```

- pprof sample

```golang
func saveMemoryProfile(filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	runtime.GC()    // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
```

```golang
saveMemoryProfile(filepath.Join("./prof", "mem.pprof"))
```

```bash
go tool pprof -inuse_space -top ./prof/mem.pprof | head -n 3
```
