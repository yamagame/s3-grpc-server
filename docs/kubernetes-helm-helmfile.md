# Kubernetes / helm / helmfile メモ

## vscode kubernetes extension

kubernetes を使用する場合の vscode のおすすめ機能拡張

https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools

## YAML extension

kubernetes では YAML ファイルを使用する

https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml

# YAML

YAML は JSON のスーパーセット。JSON で表現できるものは、YAML でも表現できる。

https://yaml.org/

# kubectl

kubernetes を操作するための基本のコマンド。

https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/

# golang template

helm / helmfile では golang の template を使用しているため、先に template 構文に慣れておく必要がある。

[Go Template、最高のプログラミング言語](https://qiita.com/Syuparn/items/4157d13c39b95185acfd)

gomplate を使用して golang の template 構文を試すことができる。

[gomplate : https://docs.gomplate.ca/](https://docs.gomplate.ca/)

```bash
# オブジェクトの値を埋め込む
$ echo '{{ .Env.HOSTNAME }}' | docker run -i --rm --hostname myhostname hairyhenderson/gomplate:stable
myhostname

# with 構文の間はオブジェクトのプロパティ名を省略できる
$ echo '{{ with .Env }}{{ .HOSTNAME }}{{ end }}' | docker run -i --rm --hostname myhostname hairyhenderson/gomplate:stable
myhostname

# パイプ 「|」を使用して値を渡す
$ echo '{{ "put" | printf "%s%s" "out" }}' | docker run -i --rm hairyhenderson/gomplate:stable

# 変数
$ echo '{{ $w := "world" }}Hello, {{ print $w }}!' | docker run -i --rm hairyhenderson/gomplate:stable

# if 文
$ echo '{{ $w := "" }}{{ if 1 }}{{ $w = "world" }}{{ end }}Hello, {{ print $w }}!' | docker run -i --rm hairyhenderson/gomplate:stable

# 変数スコープ、以下の例で$w変数未定義エラーとなる
$ echo '{{ if 1 }}{{ $w := "world" }}{{ end }}Hello, {{ print $w }}!' | docker run -i --rm hairyhenderson/gomplate:stable
```

# helm

helm は Kubernetes 用パッケージマネージャー

https://helm.sh/ja/

- [参考：Helmの概要とChart(チャート)の作り方](https://qiita.com/thinksphere/items/5f3e918015cf4e63a0bc)
- [参考：helmとは & helmの使い方](https://qiita.com/sheepland/items/75b647b71c34c7d38804)

```bash
# パッケージの作成
$ helm create sample-app
```

以下はパッケージに含まれるファイル、必要の応じて編集する。templatesディレクトリに含まれるファイルは kubectl で apply されるファイル。

- charts/
- templates/
  - _helpers.tpl
  - deployment.yaml
  - hpa.yaml
  - ingress.yaml
  - NOTES.txt
  - service.yaml
  - serviceaccount.yaml
- .helmignore
- Chart.yaml
- values.yaml

templates ディレクトリに作成されたファイルは golang の template になっている。
helm apply 時に values.yaml の値が埋め込まれる。

```bash
$ cat ./sample-app/templates/service.yaml 
apiVersion: v1
kind: Service
metadata:
  name: {{ include "sample-app.fullname" . }}
  labels:
    {{- include "sample-app.labels" . | nindent 4 }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: http
      protocol: TCP
      name: http
  selector:
    {{- include "sample-app.selectorLabels" . | nindent 4 }}
```

以下のように、values.yaml の値が埋め込まれる。

```bash
$ helm template -s templates/service.yaml ./sample-app/ 
---
# Source: sample-app/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: release-name-sample-app
  labels:
    helm.sh/chart: sample-app-0.1.0
    app.kubernetes.io/name: sample-app
    app.kubernetes.io/instance: release-name
    app.kubernetes.io/version: "1.16.0"
    app.kubernetes.io/managed-by: Helm
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: sample-app
    app.kubernetes.io/instance: release-name
```

values.yaml を -f オプションで変更して出力することも可能。また、-s オプションで特定の yaml ファイルを指定して処理することもできる。

```bash
$ helm template -f ./sample-app/values.yaml -s templates/service.yaml ./sample-app/
```

アプリケーションをデプロイする場合は、サブコマンドの apply または sync を実行する。

```bash
# マニフェストに差分がある場合に helm upgrade を実行する
$ helm apply ./sample-app/ 

# 毎回 upgrade を実行する、差分がなくても revision が上がっていく
$ helm sync ./sample-app/ 

# アプリケーションの削除
$ helm delete ./sample-app/

# インストール済みのリリースを一覧表示
$ helm list
```

# helmfile

helmfile は helm で使用する values.yaml をテンプレート化し、環境ごとに指定できる。

https://helmfile.readthedocs.io/en/latest/#configuration

下記は、helmfile.yaml の例

```yaml
# helmfile.yaml の例
releases:
  - name: sample-app
    chart: sample-app/
    version: 0.0.1
    values:
      - ./sample-app/values.yaml
```

この helmfile は以下のコマンドの実行と同じ意味となる。

```bash
$ helm install sample-app sample-app/ --version 0.0.1 -f ./sample-app/values.yaml
```

helmfile.yaml を準備しておけば、以下のコマンドで同じ結果を得られる。

```bash
$ helmfile template
```

helmfile.yaml はデフォルトのファイル名。-f オプションで変更することもできる。

```bash
$ helmfile -f ./sample-app.yaml template
```

アプリケーションをデプロイする場合は、サブコマンドの apply または sync を実行する。処理内容は helm ファイルと同様、以下のコマンドで helm が実行される。

```bash
# マニフェストに差分がある場合に helm upgrade を実行する
$ helmfile -f ./sample-app.yaml apply

# 毎回 upgrade を実行する、差分がなくても revision が上がっていく
$ helmfile -f ./sample-app.yaml sync

# アプリケーションの削除
$ helmfile -f ./sample-app.yaml destroy
```

# 環境ごとの helmfile

helmfile は golang の template 構文が使用できる。

```yaml
environments:
  development:
  staging:
  production:

---
releases:
  - name: sample-app
    chart: sample-app/
    version: 0.0.1
    values:
      - values-{{ .Environment.Name }}.yaml
```

valueファイルは環境ごとに作成しておく。

```bash
$ ls
./sample-app.yaml
./values-development.yaml
./values-staging.yaml
./values-production.yaml
```

staging 環境を指定して実行する。

```bash
$ helmfile -e staging -f ./sample-app.yaml template
```

golang の template 構文を使用しているファイルは作法として拡張子「.gotmpl」を使用する。

```yaml
releases:
  - name: sample-app
    chart: sample-app/
    version: 0.0.1
    values:
      - values-{{ .Environment.Name }}.yaml.gotmpl
```
