package rbac.authz

# roleのcheckが通ること
test_eval_role_exists {
    eval_role_exists with input as {"roles": ["admin"]}
}

# 存在しないroleの場合はcheckが通らないこと
test_not_eval_role_exists {
    not eval_role_exists with input as {"roles": ["not_exists_role"]}
}

# 今回アクセスするresourceに、ユーザーがアクセスする権限がある場合に許可されること
test_eval_resource_access {
    eval_resource_access with input as {"allow_resources": ["a", "b", "c"], "access_resource": "a"}
}

# 今回アクセスするresourceに、ユーザーがアクセスする権限がある場合に許可されないこと
test_not_eval_resource_access {
    not eval_resource_access with input as {"allow_resources": ["a", "b", "c"], "access_resource": "x"}
}

# 今回アクセスするpathが存在する場合に許可されることを確認
test_eval_path_exists {
    eval_path_exists with input as {"path": "/apps", "method": "GET"}
}

# 今回アクセスするpathが存在しない場合にfalseになることを確認
test_not_eval_path_exists {
    not eval_path_exists with input as {"path": "not_exists_path", "method": "GET"}
    not eval_path_exists with input as {"path": "/apps", "method": "NOT_EXISTS_METHOD"}
}

test_allow {
    allow with input as {
        # ユーザーの所持している権限情報
        "user": "user_name",
        "roles": ["read-only", "huwahuwa"],
        "allow_resources": ["test-1-app", "test-2-app"],

        # 今回のユーザーが行うアクション情報
        "method": "GET",
        "path": "/apps",
        "access_resource": "test-1-app"
    }
}

# read-only権限で、postしようとした時にerrorになることを確認
test_not_allow_post {
    not allow with input as {
        "user": "user_name",
        "roles": ["read-only"],
        "allow_resources": [".*"],

        "method": "POST",
        "path": "/apps",
        "access_resource": "test"
    }
}

test_allow_regex {
     allow with input as {
        # ユーザーの所持している権限情報
        "user": "user_name",
        "roles": ["admin"],
        "allow_resources": [".*"],

        # 今回のユーザーが行うアクション情報
        "method": "GET",
        "path": "/apps",
        "access_resource": "test-1-app"
    }
}
