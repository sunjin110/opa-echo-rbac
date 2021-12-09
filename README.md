# opa-echo-test

Open Policy Agentとechoで、  
Roll-Based Access Controlの実装がうまくできるかどうかの検証  
APIはRestAPI形式にしていく予定です  

# 書き方styleの言い訳
検証用なので、chk.SE()でpanicをさせるパターンです  
interface使って、依存関係のあれそれはやっていないです(検証用なため)

# テスト手順(scriptでは後で書きます...)

```sh
# 起動(デフォルトでread-onlyのユーザーになるようにしています)
go run main.go

# localhost:1234にアクセス
curl -X GET http://localhost:1234/apps
# "GET /apps"とアクセスできる

curl -X POST http://localhost:1234/apps
# "許可されていません"と表示されます

# /internal/echo/emiddleware/emiddleware.go内の
# userIDを2に替えると、admin権限になる(引数とかで制御するのめんどくさかったので、こうしてますごめんなさい)

# localhost:1234にアクセス
curl -X GET http://localhost:1234/apps
# "GET /apps"とアクセスできる

curl -X POST http://localhost:1234/apps
# "POST /apps"とアクセスできる


```