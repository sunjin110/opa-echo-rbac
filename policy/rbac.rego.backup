package rbac.authz

# inputのデータ
input_1 := {
    # ユーザーが所持している情報
    "user": "sunjin",
    "roles": ["admin"],
    "allow_resources": ["*"], # このユーザーが観覧を許可されているresource

    # 今回ユーザーがおこうアクション
    "method": "PUT",
    "path": "/apps/update/status",
    "access_resource": "resource_1" # 今回アクセスしようとしているresource
}

input_2 := {
    "user": "sunjin2",
    "roles": ["read-only"],
    "allow_resources": ["sunjin-app", "test-app"],


    "method": "GET",
    "path": "/apps/:id/huwahuwa/campaigns",
    "access_resource": "sunjin-app"
    
}

# 固定のデータ

# TODO 管理をシンプルにするために、Groupを作ってもいいかも

roles := {
    # 全ての権限を観覧できるrole
    "admin": [ 
        "*"
    ],
    # 観覧権限のみ所持しているrole
    "read-only": [
        "apps:list",
        "apps:detail",
        "huwahuwa:list",
        "hogehoge:list",
    ],
    # huwahuwa系の権限を全て所持しているユーザー
    "huwahuwa": [
        "huwahuwa:*"
    ]
}

# pathごとに、必要なaction権限を割り当てしていく
path_permissions := {
    "/login": {
        "GET": [],
        "POST": []
    },
    "/logout": {
        "GET": []
    },
    "/apps": {
        "GET": ["apps:list"],
        "POST": ["apps:create"],
        "PUT": ["apps:update"],
        "DELETE": ["apps:delete"],
    },
    "/apps/:id" : {
        "GET": ["apps:detail"],
    },
    "/apps/update/status": {
        "PUT": ["apps:activation"]
    },
    "/apps/:id/huwahuwa/campaigns": {
        "GET": ["huwahuwa:list"],
    }

}

# actionの種類
actions := {
    "apps": [
        "list",
        "create",
        "update",
        "delete",
        "activation", # 有効/無効にできるかどうか？
        "detail"
    ],
    "huwahuwa": [
        "list",
        "add",
        "update",
        "remove",
    ],
    "hogehoge": [
        "list",
        "add",
        "update",
        "remove"
    ],
    "account": [
        "list",
        "cresate",
        "update",
        "delete"
    ],
}

default allow = false
allow {
    # TODO
    true
}

# # logic that implements RBAC
# default allow = false
# allow {
#     # lookup the list of roles for the user
#     # ユーザーの役割リストを検索
#     roles := data.user_roles[input.user]

#     # for each role in that list
#     # 役割ごとに検証
#     r := roles[_]

#     # 格roleのpermissionを調べる
#     permissions := data.role_permissions[r]

#     # permissionごとにチェック
#     p := permissions[_]

#     # 権限のcheck
#     p == {"action": input.action, "object": input.object}
# }
