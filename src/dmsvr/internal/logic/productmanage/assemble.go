package productmanagelogic

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/utils"
	mysql "github.com/i-Things/things/src/dmsvr/internal/repo/mysql"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

//func ToProductSchema(pt *schema.Info) *dm.ProductSchema {
//	return &dm.ProductSchema{
//		CreatedTime: pt.CreatedTime.Unix(),
//		ProductID:   pt.ProductID,
//		Schema:      pt.Schema,
//	}
//}

func ToProductInfo(pi *mysql.DmProductInfo) *dm.ProductInfo {
	var (
		tags map[string]string
	)

	_ = json.Unmarshal([]byte(pi.Tags), &tags)
	if pi.DeviceType == def.Unknown {
		pi.DeviceType = def.DeviceTypeDevice
	}
	if pi.NetType == def.Unknown {
		pi.NetType = def.NetOther
	}
	if pi.DataProto == def.Unknown {
		pi.DataProto = def.DataProtoCustom
	}
	if pi.AuthMode == def.Unknown {
		pi.AuthMode = def.AuthModePwd
	}
	if pi.AutoRegister == def.Unknown {
		pi.AutoRegister = def.AutoRegClose
	}
	dpi := &dm.ProductInfo{
		ProductID:    pi.ProductID,                          //产品id
		ProductName:  pi.ProductName,                        //产品名
		AuthMode:     pi.AuthMode,                           //认证方式:0:账密认证,1:秘钥认证
		DeviceType:   pi.DeviceType,                         //设备类型:0:设备,1:网关,2:子设备
		CategoryID:   pi.CategoryID,                         //产品品类
		NetType:      pi.NetType,                            //通讯方式:0:其他,1:wi-fi,2:2G/3G/4G,3:5G,4:BLE,5:LoRaWAN
		DataProto:    pi.DataProto,                          //数据协议:0:自定义,1:数据模板
		AutoRegister: pi.AutoRegister,                       //动态注册:0:关闭,1:打开,2:打开并自动创建设备
		Secret:       pi.Secret,                             //动态注册产品秘钥 只读
		Desc:         &wrappers.StringValue{Value: pi.Desc}, //描述
		CreatedTime:  pi.CreatedTime.Unix(),                 //创建时间
		Tags:         tags,                                  //产品tags
		//Model:     &wrappers.StringValue{Value: pi.Model},    //数据模板
	}
	return dpi
}

func ToProductSchemaRpc(info *mysql.DmProductSchema) *dm.ProductSchemaInfo {
	db := &dm.ProductSchemaInfo{
		ProductID:  info.ProductID,
		Tag:        info.Tag,
		Type:       info.Type,
		Identifier: info.Identifier,
		Name:       utils.ToRpcNullString(&info.Name),
		Desc:       utils.ToRpcNullString(&info.Desc),
		Required:   info.Required,
		Affordance: utils.ToRpcNullString(&info.Affordance),
	}
	return db
}

func ToProductSchemaPo(info *dm.ProductSchemaInfo) *mysql.DmProductSchema {
	db := &mysql.DmProductSchema{
		ProductID:  info.ProductID,
		Tag:        info.Tag,
		Type:       info.Type,
		Identifier: info.Identifier,
		Name:       info.Name.GetValue(),
		Desc:       info.Desc.GetValue(),
		Required:   info.Required,
		Affordance: info.Affordance.GetValue(),
	}
	return db
}
