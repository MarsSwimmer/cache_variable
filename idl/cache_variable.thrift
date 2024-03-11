namespace go cache_variable

struct GetReq {
    1: string Key (api.query="key"); // path 注解用来声明 url 中的路由参数
}
struct GetResp {
    1: string Value;
}

struct SetReq {
    1: string Key    (api.query="key");    // 需要设置的key
    2: string Value  (api.query="value");  // 需要设置的value
    3: i32    Expire (api.query="expire"); // 过期时间
}
struct SetResp {
}

struct DelReq {
    1: string Key    (api.query="key");    // 需要设置的key
}
struct DelResp {
}



service CacheVariableService{
    // 声明请求的路由
    GetResp Get(1:GetReq req) (api.get="/get_value", api.handler_path="get");
    SetResp Set(1:SetReq req) (api.get="/set_value", api.handler_path="get");
    DelResp Del(1:DelReq req) (api.get="/delete_value", api.handler_path="get");
}(
    api.base_domain="http://127.0.0.1:6366";
)
    

