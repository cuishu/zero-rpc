syntax = "proto3";

package {{.ShortModule}};
option go_package="./proto";

// 请求参数
message Request {
  // field注释不能少
  string Arg = 1;
}

// 返回值
message Response {
  // field注释不能少
  string Ret = 1;
}

// 最好有注释
service {{.Service.Name}} {
  // method 注释必须要写
  rpc Hello(Request) returns(Response);
}
