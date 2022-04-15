mysql的简单使用


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
if err := global.Mysql.Get(context.Background(), global.SQLFormat(cmdPage), func(row map[string]interface{})error{  
    jsonString, _ := json.Marshal(row)
    s := RuleDB{}
    json.Unmarshal(jsonString, &s)
    doc = append(doc,s)
    return nil
  }); err != nil{
    log.Error("%#v",err)
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


