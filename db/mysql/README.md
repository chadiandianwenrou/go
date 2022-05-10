mysql的简单使用

```
//只执行sql

//更新
prometheusRulesDatacenterService_mapCmd := fmt.Sprintf(`update prometheus_rules_datacenter_service_map SET 
  rule_id=%d,datacenter_group_id=%d,system_id=%d,module_id=%d,update_time='%s',update_by='%s' where id=%d`,
  ruleId,promRuleMap.DatacenterGroupId,promRuleMap.SystemId,promRuleMap.ModuleId,updateTime,updateBy,id)

_,err = global.Mysql.Exec(context.Background(), prometheusRulesDatacenterService_mapCmd)

  if err != nil{
    log.Error("%#v",err)
    err = errors.New("数据写入失败！")
    return
  }


//写入
  _,err = global.Mysql.Exec(context.Background(), prometheusRuleCmd)
  if err != nil{
    log.Error("写入失败 %#v",err)
    // err = errors.New("alert name 已存在！")
  }




//查询结果进切片
func GetCommunityList() (communityList []*rsp.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	if err := global.Mysql.Get(context.Background(), global.SQLFormat(sqlStr), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		s := rsp.Community{}
		json.Unmarshal(jsonString, &s)
		communityList = append(communityList, &s)
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

//查询单条数据
func (r *ListRule) Get() (rule RuleDB, err error) {
	// rule := RuleDB{}
	// cmd := "select * from prometheus_rule where is_del = 0 "
	field := "prometheus_rules_datacenter_service_map.id,prometheus_rules_datacenter_service_map.status,prometheus_rules_datacenter_service_map.create_time,prometheus_rules_datacenter_service_map.update_time,prometheus_rules_datacenter_service_map.is_super_admin,prometheus_rules_datacenter_service_map.is_del,prometheus_datacenter_group.datacenter_group_name,asura_service_system.system_name,asura_service_module.module_name,prometheus_rule.alert,prometheus_rule.expr,prometheus_rule.for,prometheus_rule.annotations,prometheus_rule.labels,asura_user.user_nick_name as create_by,a_user.user_nick_name as update_by"
	cmd := fmt.Sprintf(`SELECT %s
										 FROM prometheus_rules_datacenter_service_map
											LEFT JOIN     prometheus_datacenter_group ON  prometheus_datacenter_group.datacenter_group_id=prometheus_rules_datacenter_service_map.datacenter_group_id
											LEFT JOIN     asura_service_system   ON  asura_service_system.system_id=prometheus_rules_datacenter_service_map.system_id
											LEFT JOIN   	asura_service_module   ON asura_service_module.module_id=prometheus_rules_datacenter_service_map.module_id
											LEFT JOIN     prometheus_rule 			 ON prometheus_rule.id = prometheus_rules_datacenter_service_map.rule_id
											LEFT JOIN     asura_user  ON  asura_user.user_id = prometheus_rules_datacenter_service_map.create_by
											LEFT JOIN     asura_user  as a_user ON  a_user.user_id = prometheus_rules_datacenter_service_map.update_by
										WHERE 
											prometheus_rules_datacenter_service_map.is_del = 0`, field)

	if r.Id != 0 {
		cmd = fmt.Sprintf("%s and prometheus_rules_datacenter_service_map.id = %d", cmd, r.Id)
	}
	// if r.Alert != "" {
	// 	cmd = fmt.Sprintf("%s and alert = %s", cmd, r.Alert)
	// }

	if err := global.Mysql.Get(context.Background(), global.SQLFormat(cmd), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		json.Unmarshal(jsonString, &rule)
		// doc = append(doc,s)
		return nil
	}); err != nil {
		log.Error("%#v", err)

	}
	return
}


  //查询count
  var count int32
  cmdCount := strings.Replace(cmd, selectRuleField, "count(1) as count",1)
  if err := global.Mysql.Get(context.Background(), cmdCount, func(row map[string]interface{})error{ 
    for _,v := range row{
      c,_ := types.ToInt(v)
      count = int32(c)
    }
    return nil
  }); err != nil{
    log.Error("%#v",err)
  }

```
