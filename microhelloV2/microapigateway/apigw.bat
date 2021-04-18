#注册外部服务到  micro-regsitry
# gin beego 等
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
set MICRO_API_NAMESPACE=sjfbjs.com.api
set MiCRO_API_HANDLER=rpc
micro api --namespace sjfbjs.com.api