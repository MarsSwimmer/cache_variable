namespace go cache_variable_server

// 设置k-v
struct SetValueReq {
    1: string Key (api.query="key");       // 变量名称
    2: string Value (api.query="value");   // 变量值
    3: i32    expire (api.query="expire"); // 过期时间，单位秒
}
struct SetValueResp {
}

// 获取变量
struct GetValueReq {
    1: string Key (api.query="key");     // 变量名称
}
struct GetValueResp {
    1: string Value
}

// 删除key-value
struct DeleteValueReq {
    1: string Key (api.query="key");       // 变量名称
}
struct DeleteValueResp {
}

service CacheVariableServerService {
    // 设置k-v
    SetValueResp SetValue(1: SetValueReq req) (api.get="/set_value")
    
    // 获取变量值
    GetValueResp GetValue(1: GetValueReq req) (api.get="/get_value")

    // 删除key-value
    DeleteValueResp DeleteValue(1:DeleteValueReq req) (api.get="/delete_value")
}