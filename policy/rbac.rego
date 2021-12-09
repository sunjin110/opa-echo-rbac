package rbac.authz

# input format
# input := {
#     "user": "sunjin2",
#     "roles": ["read-only"],
#     "allow_resources": ["sunjin-app", "test-app"],
#     "method": "GET",
#     "path": "/apps/:id/huwahuwa/campaigns",
#     "access_resource": "sunjin-app"
# }



# 固定のデータ

# TODO 管理をシンプルにするために、Groupを作ってもいいかも

roles := {
    # 全ての権限を観覧できるrole
    "admin": [ 
        ".*"
    ],
    # 観覧権限のみ所持しているrole
    "read-only": [
        "apps:list",
        "apps:list:2",
        "apps:detail",
        "huwahuwa:list",
        "hogehoge:list",
    ],
    # huwahuwa系の権限を全て所持しているユーザー
    "huwahuwa": [
        "huwahuwa:*",
        "huwahuwa:list",
    ]
}

# pathごとに、必要なaction権限を割り当てしていく
# その権限を全て所持していないといけない
path_permissions := {
    "/login": {
        "GET": [],
        "POST": []
    },
    "/logout": {
        "GET": []
    },
    "/apps": {
        "GET": ["apps:list", "apps:list:2"],
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

# userのroleが存在することを確認する
eval_role_exists = true {
    # userのroleを一つずつ取得
    user_role := input.roles[_]

    # dataのroleに存在すればtrue
    roles[user_role]
}

# 今回アクセスするresourceに、ユーザーがアクセスする権限があるかどうかを確認
eval_resource_access = true {
    # userがアクセス可能なresoource
    allow_resource := input.allow_resources[_]


    # 正規表現でresourceにアクセスできるかどうかを確認する
    regex.match(allow_resource, input.access_resource)
}

# 今回アクセスするpathが存在することを確認する
eval_path_exists = true {
    # pathがあることを確認
    path := path_permissions[input.path]

    # methodがあるかどうか確認
    path[input.method]
}


default allow = false

# ユーザーができるactionをまとめる(sets)
user_actions[action] {
    user_role := input.roles[_]
    user_actions := roles[user_role]
    action := user_actions[_]
}

# 今回必要となるactionを取得する(sets)
require_actions[action] {
    path := path_permissions[input.path]
    actions := path[input.method]
    action := actions[_]
}

# 今回使うactionをuser_actionsから取得する
filter_user_actions[action] {
    require_action := require_actions[_]

    # iter
    some user_action
    user_actions[user_action]

    # 正規表現でのroleもサポートする
    result := regex.match(user_action, require_action)
    action := require_action
}

allow = true{
    eval_role_exists == true
    eval_resource_access == true
    eval_path_exists == true

    # 権限のcheck
    require_actions == filter_user_actions
}
