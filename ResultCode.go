package utils

const (
	SUCCESS         = 200
	SUCCESS_MESSAGE = "操作成功"

	PRODUCT_OPERATING         = 201
	PRODUCT_OPERATING_MESSAGE = "产品正在进行操作"

	LOGIN         = 307
	LOGIN_MESSAGE = "重新登录"

	TOKEN_IS_NIL         = 308
	TOKEN_IS_NIL_MESSAGE = "token不能为空"

	REQUEST_PARAMETERS         = 400
	REQUEST_PARAMETERS_MESSAGE = "请求参数不能为空"

	PARAMETERS_TYPE_ERROR         = 401
	PARAMETERS_TYPE_ERROR_MESSAGE = "参数类型不合法"

	FINANCIAL_NO_POWER         = 402
	FINANCIAL_NO_POWER_MESSAGE = "无财务权限"

	NO_CUSTOMER         = 403
	NO_CUSTOMER_MESSAGE = "此客户不存在"

	ERROR         = 500
	ERROR_MESSAGE = "服务器异常"

	RETURN_DATA_NULL         = 501
	RETURN_DATA_NULL_MESSAGE = "返回值为空"

	FAIL         = 502
	FAIL_MESSAGE = "操作失败"

	TOKEN_IS_NULL         = 600
	TOKEN_IS_NULL_MESSAGE = "获取token为空"

	//########################注册登录####################################
	REGISTER_SUCCESS         = 1000
	REGISTER_SUCCESS_MESSAGE = "注册成功"

	REGISTER_FAIL         = 1001
	REGISTER_FAIL_MESSAGE = "注册失败"

	NAME_NOT_NIL         = 1002
	NAME_NOT_NIL_MESSAGE = "用户名不合法"

	NAME_REGISTERD         = 1002
	NAME_REGISTERD_MESSAGE = "用户名已被注册"

	EMAIL_NOT_NIL         = 1003
	EMAIL_NOT_NIL_MESSAGE = "邮箱不合法"

	EMAIL_REGISTERD         = 1002
	EMAIL_REGISTERD_MESSAGE = "邮箱已被注册"

	PHONE_NOT_NIL         = 1004
	PHONE_NOT_NIL_MESSAGE = "手机号不合法"

	PHONE_REGISTERD         = 1002
	PHONE_REGISTERD_MESSAGE = "手机号已被注册"

	PASSWORD_NOT_NIL         = 1004
	PASSWORD_NOT_NIL_MESSAGE = "密码不合法"

	CONFIRMPASSWORD_NOT_NIL         = 1005
	CONFIRMPASSWORD_NOT_NIL_MESSAGE = "确认密码不合法"

	PASSWORD_NOT_SAME         = 1006
	PASSWORD_NOT_SAME_MESSAGE = "密码不相同"

	PHONECODE_NOT_NIL         = 1008
	PHONECODE_NOT_NIL_MESSAGE = "验证码不能为空"

	VERIFICATION_CODE_FAIL         = 1009
	VERIFICATION_CODE_FAIL_MESSAGE = "验证码失败"

	VERIFICATION_CODE_NOT_SAME         = 1010
	VERIFICATION_CODE_NOT_SAME_MESSAGE = "验证码不匹配"

	PASSWORD_MD5_FAIL         = 1011
	PASSWORD_MD5_FAIL_MESSAGE = "处理客户信息异常"

	RESEND_CODE         = 1012
	RESEND_CODE_MESSAGE = "请重新发送验证码"

	LOGIN_FAIL         = 1013
	LOGIN_FAIL_MESSAGE = "登录失败"

	PHONE_NOT_REGISTER         = 1014
	PHONE_NOT_REGISTER_MESSAGE = "手机号未注册"

	EMAIL_NOT_REGISTER         = 1015
	EMAIL_NOT_REGISTER_MESSAGE = "邮箱未注册"

	NOT_ALLOWED_LOGIN         = 1016
	NOT_ALLOWED_LOGIN_MESSAGE = "不允许登录"

	PASSWORD_FAULT         = 1017
	PASSWORD_FAULT_MESSAGE = "密码错误"

	UPDATE_PHONE_FAIL         = 1018
	UPDATE_PHONE_FAIL_MESSAGE = "修改手机号失败"

	UPDATE_EMAIL_FAIL         = 1019
	UPDATE_EMAIL_FAIL_MESSAGE = "修改邮箱失败"

	UPDATE_PASSWORD_FAIL         = 1020
	UPDATE_PASSWORD_FAIL_MESSAGE = "修改密码失败"

	SEND_CODE_FAIL         = 1021
	SEND_CODE_FAIL_MESSAGE = "发送验证码失败"

	GET_DATA_FAIL         = 1022
	GET_DATA_FAIL_MESSAGE = "解析数据失败"

	PAGING_CONDITION_NIL         = 1023
	PAGING_CONDITION_NIL_MESSAGE = "分页条件不能为空"

	CUSTOMERID_IS_NIL=1024
	CUSTOMERID_IS_NIL_MESSAGE="客户id为空"

	SEND_CODE_IS_NIL=1025
	SEND_CODE_IS_NIL_MESSAGE="是否短信通知不能为空"
	//############################################################

	//#################云主机#####################################
	ALIAS_NAME_IS_NIL         = 1100
	ALIAS_NAME_IS_NIL_MESSAGE = "主机别名不规范"

	REGIONID_IS_NIL         = 1101
	REGIONID_IS_NIL_MESSAGE = "区域id不能为空"

	REGIONNAME_IS_NIL         = 1102
	REGIONNAME_IS_NIL_MESSAGE = "区域名称不能为空"

	PROJECTID_IS_NIL         = 1103
	PROJECTID_IS_NIL_MESSAGE = "项目id不能为空"

	CPU_IS_NIL         = 1104
	CPU_IS_NIL_MESSAGE = "CPU不能为0"

	RAM_IS_NIL         = 1105
	RAM_IS_NIL_MESSAGE = "RAM不能为0"

	MIRRORID_IS_NIL         = 1106
	MIRRORID_IS_NIL_MESSAGE = "镜像id不能为空"

	MIRRORNAME_IS_NIL         = 1107
	MIRRORNAME_IS_NIL_MESSAGE = "镜像名称不能为空"

	SYSTEMCINDER_IS_NIL         = 1108
	SYSTEMCINDER_IS_NIL_MESSAGE = "系统盘大小不能为0"

	SECURITYGROUPID_IS_NIL         = 1109
	SECURITYGROUPID_IS_NIL_MESSAGE = "安全组id不能为空"

	SECURITYGROUPALISANAME_IS_NIL         = 1110
	SECURITYGROUPALISANAME_IS_NIL_MESSAGE = "安全组别名不能为空"

	SECURITYGROUPREALNAME_IS_NIL         = 1111
	SECURITYGROUPREALNAME_IS_NIL_MESSAGE = "安全组名称不能为空"

	PASSWORDSET_IS_NIL         = 1112
	PASSWORDSET_IS_NIL_MESSAGE = "是否设置密码不能为空"

	NUMBER_IS_NIL         = 1113
	NUMBER_IS_NIL_MESSAGE = "购买数量不能为0"

	RENEWTYPE_IS_NIL         = 1114
	RENEWTYPE_IS_NIL_MESSAGE = "续费方式不能为空"

	PAYTYPE_IS_NIL         = 1115
	PAYTYPE_IS_NIL_MESSAGE = "计费方式不能为空"

	PAYTIME_IS_NIL         = 1116
	PAYTIME_IS_NIL_MESSAGE = "计费时间不能为空"

	PAYMONTH_IS_NIL         = 1117
	PAYMONTH_IS_NIL_MESSAGE = "对应月不能为空"

	INTERNALIPNAME_IS_NIL         = 1118
	INTERNALIPNAME_IS_NIL_MESSAGE = "VPC名称不能为空"

	INTERNALIPCID_IS_NIL         = 1119
	INTERNALIPCID_IS_NIL_MESSAGE = "VPC客户id不能为空"

	INSERT_SERVER_FAIL         = 1120
	INSERT_SERVER_FAIL_MESSAGE = "创建主机失败"

	SERVERID_IS_NIL         = 1121
	SERVERID_IS_NIL_MESSAGE = "主机id不能为空"

	SERVER_STATUS_ERROR         = 1122
	SERVER_STATUS_ERROR_MESSAGE = "主机状态错误"

	SERVER_STATUS_ACTIVE         = 1123
	SERVER_STATUS_ACTIVE_MESSAGE = "主机是启动状态"

	SERVER_STATUS_NOT_ACTIVE         = 1124
	SERVER_STATUS_NOT_ACTIVE_MESSAGE = "主机不是运行状态"

	SERVER_STATUS_SHUTOFF         = 1125
	SERVER_STATUS_SHUTOFF_MESSAGE = "主机是关机状态"

	SERVER_START_FAIL         = 1126
	SERVER_START_FAIL_MESSAGE = "主机启动失败"

	SERVER_SHUTOFF_FAIL         = 1127
	SERVER_SHUTOFF_FAIL_MESSAGE = "主机关机失败"

	SERVER_REBOOT_FAIL         = 1128
	SERVER_REBOOT_FAIL_MESSAGE = "主机重启失败"

	SERVER_BIND_FLOATIP=1129
	SERVER_BIND_FLOATIP_MESSAGE="主机绑定浮动ip,请解绑"

	SERVER_BIND_CINDER=1130
	SERVER_BIND_CINDER_MESSAGE="主机挂载硬盘,请卸载"

	JUST_UPGRADE=1031
	JUST_UPGRADE_MESSAGE="配置只能升级"

	UPGRADE_STATUS_ING=1032
	UPGRADE_STATUS_ING_MESSAGE="正在升级"
	//############################################################

	//#################获取前端json数据错误#########################
	REQUESTBODY_FAIL         = 1009
	REQUESTBODY_FAIL_MESSAGE = "请求体获取失败"
	//############################################################

	//网络提示信息（1300-1400）
	PRIVATE_NNETWORK_EXIST         = 1300
	PRIVATE_NNETWORK_EXIST_MESSAGE = "此名字已存在"

	PRIVATE_NNETWORK_CREATE_FAIL         = 1301
	PRIVATE_NNETWORK_CREATE_FAIL_MESSAGE = "创建私有网络失败"

	VPC_CIDR_ILLEGALITY         = 1302
	VPC_CIDR_ILLEGALITY_MESSAGE = "VPC网段不合法"

	PRIVATE_NNETWORK_DELETE_FAIL         = 1303
	PRIVATE_NNETWORK_DELETE_FAIL_MESSAGE = "删除失败,已绑定云主机"

	PRIVATE_NNETWORK_HAVE_PORT         = 1304
	PRIVATE_NNETWORK_HAVE_PORT_MESSAGE = "该私网有内网ip，不可删除"

	PRIVATE_NNETWORK_NO_EXIST         = 1305
	PRIVATE_NNETWORK_NO_EXIST_MESSAGE = "该私网不存在"

	GET_PORTID_IP_FAIL         = 1306
	GET_PORTID_IP_FAIL_MESSAGE = "获取port和ip失败"

	RETURN_NULL         = 1307
	RETURN_NULL_MESSAGE = "调用创建vpc返回null"

	//项目成员（1400-1500）
	CREATE_PROJECT_ERROR         = 1400
	CREATE_PROJECT_ERROR_MESSAGE = "创建项目失败"

	CREATE_PROJECT_LIMIT         = 1401
	CREATE_PROJECT_LIMIT_MESSAGE = "项目已经达到上限，无法继续创建"

	EDITE_PROJECT_FAIL         = 1402
	EDITE_PROJECT_FAIL_MESSAGE = "项目修改失败"

	DELETE_PROJECT_FAIL         = 1403
	DELETE_PROJECT_FAIL_MESSAGE = "项目包含成员，请删除成员"

	DELETE_PROJECT_FAIL1         = 1404
	DELETE_PROJECT_FAIL1_MESSAGE = "项目包含产品，请删除产品"

	NEED_ADD_MEMBER         = 1405
	NEED_ADD_MEMBER_MESSAGE = "当前项目下没有可选成员，是否新建成员"

	USER_EXIT         = 1406
	USER_EXIT_MESSAGE = "用户已经存在，请重新输入"

	USER_ADD_SUCCESS         = 1407
	USER_ADD_SUCCESS_MESSAGE = "添加成员成功"

	USER_ADD_FAIL         = 1408
	USER_ADD_FAIL_MESSAGE = "添加成员失败"

	REMOVE_MEMBER_FAIL         = 1409
	REMOVE_MEMBER_FAIL_MESSAGE = "移除成员失败"

	NAME_EXIT         = 1410
	NAME_EXIT_MESSAGE = "用户名已经存在，不允许重复"

	MAIL_EXIT         = 1411
	MAIL_EXIT_MESSAGE = "邮箱已经存在，不允许重复"

	PHONE_EXIT         = 1412
	PHONE_EXIT_MESSAGE = "手机号已经存在，不允许重复"

	CREATE_MEM_SUCCESS         = 1413
	CREATE_MEM_SUCCESS_MESSAGE = "新建成员成功"

	CREATE_MEM_FAIL         = 1414
	CREATE_MEM_FAIL_MESSAGE = "新建成员失败"

	PROJECT_NOT_EXIT         = 1415
	PROJECT_NOT_EXIT_MESSAGE = "项目不存在"

	MEMBER_AND_PROJECT         = 1416
	MEMBER_AND_PROJECT_MESSAGE = "该成员不属于此项目"

	CUSTOMER_NOW_USING         = 1417
	CUSTOMER_NOW_USING_MESSAGE = "该用户登录状态正常,无需恢复"

	FAIL_TO_DISTRIBUTION         = 1418
	FAIL_TO_DISTRIBUTION_MESSAGE = "默认分配项目配额失败"

	DELETE_PROJECTS_FAIL         = 1419
	DELETE_PROJECTS_FAIL_MESSAGE = "删除项目失败"

	DELETE_DEFAULT_FAIL         = 1420
	DELETE_DEFAULT_FAIL_MESSAGE = "默认项目不允许删除"

	MEMBER_IN_PROJECT         = 1421
	MEMBER_IN_PROJECT_MESSAGE = "成员已添加,禁止重复添加"

	PASSWORD_NOT_EQUAL         = 1422
	PASSWORD_NOT_EQUAL_MESSAGE = "两次密码不统一"

	REMOVE_FAIL         = 1423
	REMOVE_FAIL_MESSAGE = "当前成员已分配项目,请在项目中移除"

	MEMBER_NOT_EXIT         = 1424
	MEMBER_NOT_EXIT_MESSAGE = "成员不存在"

	CREATE_LOANDBALANCE_FAIL         = 1425
	CREATE_LOANDBALANCE_FAIL_MESSAGE = "添加负载均衡的配额失败"

	DELETE_LOANDBALANCE_FAIL         = 1426
	DELETE_LOANDBALANCE_FAIL_MESSAGE = "删除负载均衡的配额失败"

	DELETE_DMEMBER_FAIL         = 1427
	DELETE_DMEMBER_FAIL_MESSAGE = "调用docker删除成员失败"

	DELETE_DOCKERMEMBER_FAIL         = 1428
	DELETE_DOCKERMEMBER_FAIL_MESSAGE = "该成员下含有容器产品无法删除"

	EDITE_PROJECT_SUCCESS         = 1429
	EDITE_PROJECT_SUCCESS_MESSAGE = "项目修改成功"

	//#################项目配额#########################
	CPU_QUOTA_INSUFFICIENT         = 1501
	CPU_QUOTA_INSUFFICIENT_MESSAGE = "cpu配额不足"

	RAM_QUOTA_INSUFFICIENT         = 1502
	RAM_QUOTA_INSUFFICIENT_MESSAGE = "ram配额不足"

	IP_QUOTA_INSUFFICIENT         = 1503
	IP_QUOTA_INSUFFICIENT_MESSAGE = "ip配额不足"

	IP_WIDTH_QUOTA_INSUFFICIENT         = 1504
	IP_WIDTH_QUOTA_INSUFFICIENT_MESSAGE = "带宽配额不足"

	CINDER_QUOTA_INSUFFICIENT         = 1505
	CINDER_QUOTA_INSUFFICIENT_MESSAGE = "硬盘数量配额不足"

	HOST_CREATE_LIMIT         = 1506
	HOST_CREATE_LIMIT_MESSAGE = "主机配额不足"

	CINDER_CREATE_LIMIT         = 1507
	CINDER_CREATE_LIMIT_MESSAGE = "硬盘配额不足"

	SNATSHOP_CREATE_LIMIT         = 1508
	SNATSHOP_CREATE_LIMIT_MESSAGE = "快照配额不足"

	SECURITY_GROUP_LIMIT         = 1509
	SECURITY_GROUP_LIMIT_MESSAGE = "安全组配额不足"

	DELETE_DEFAULTRULE_FAIL         = 1510
	DELETE_DEFAULTRULE_FAIL_MESSAGE = "默认安全组规则不可以删除"

	PRIVATENTWORK_CREATE_FAIL         = 1511
	PRIVATENTWORK_CREATE_FAIL_MESSAGE = "私有网络配额不足"

	LB_CREATE_LIMIT         = 1512
	LB_CREATE_LIMIT_MESSAGE = "负载均衡配额不足"

	QUOTA_CANNOTLESSTHN0         = 1513
	QUOTA_CANNOTLESSTHN0_MESSAGE = "配额数必须大于等于0"

	QUOTA_EDIT_FAILURE         = 1514
	QUOTA_EDIT_FAILURE_MESSAGE = "配额数修改失败"

	QUOTA_EDIT_CURRENTQUOTA                         = 1515
	QUOTA_EDIT_CURRENTQPROJECT_NOT_EXITUOTA_MESSAGE = "配额数必须大于当前配额"

	CLOUDHOST_EXIT         = 1516
	CLOUDHOST_EXIT_MESSAGE = "当前项目下含有主机,不允许删除"

	CLOUDCINDER_EXIT         = 1517
	CLOUDCINDER_EXIT_MESSAGE = "当前项目下含有硬盘,不允许删除"

	FLOATIP_EXIT         = 1518
	FLOATIP_EXIT_MESSAGE = "当前项目下含有IP,不允许删除"

	NAT64_EXIT         = 1519
	NAT64_EXIT_MESSAGE = "当前项目下含NAT64,不允许删除"

	PRIVATENETWORK_EXIT         = 1520
	PRIVATENETWORK_EXIT_MESSAGE = "当前项目下含privateNetwork,不允许删除"

	EQUIPMENT_DELETE_FAIL         = 1521
	EQUIPMENT_DELETE_FAIL_MESSAGE = "项目包含成员，请删除成员"

	REGION_RENAME_FAIL         = 1522
	REGION_RENAME_FAIL_MESSAGE = "区域名称已存在"

	REGION_DELETE_FAIL         = 1523
	REGION_DELETE_FAIL_MESSAGE = "存在子区域，请删除子区域"

	CPU_QUOTA_CANNOTLESSTHN0         = 1524
	CPU_QUOTA_CANNOTLESSTHN0_MESSAGE = "CPU配额数必须大于等于0"

	CLOUDHOST_QUOTA_CANNOTLESSTHN0         = 1525
	CLOUDHOST_QUOTA_CANNOTLESSTHN0_MESSAGE = "主机配额数必须大于等于0"

	CLOUDDISK_QUOTA_CANNOTLESSTHN0         = 1526
	CLOUDDISK_QUOTA_CANNOTLESSTHN0_MESSAGE = "云硬盘配额数必须大于等于0"

	CLOUDDISKSIZE_QUOTA_CANNOTLESSTHN0         = 1527
	CLOUDDISKSIZE_QUOTA_CANNOTLESSTHN0_MESSAGE = "云硬盘总量配额数必须大于等于0"

	SNAPSHOT_QUOTA_CANNOTLESSTHN0         = 1528
	SNAPSHOT_QUOTA_CANNOTLESSTHN0_MESSAGE = "快照数量配额数必须大于等于0"

	RAM_QUOTA_CANNOTLESSTHN0         = 1529
	RAM_QUOTA_CANNOTLESSTHN0_MESSAGE = "内存配额数必须大于等于0"

	SECURITYGROUPS_QUOTA_CANNOTLESSTHN0         = 1530
	SECURITYGROUPS_QUOTA_CANNOTLESSTHN0_MESSAGE = "安全组配额数必须大于等于0"

	SECURITYGROUPRULES_QUOTA_CANNOTLESSTHN0         = 1531
	SECURITYGROUPRULES_QUOTA_CANNOTLESSTHN0_MESSAGE = "安全组规则配额数必须大于等于0"

	FLOATIP_QUOTA_CANNOTLESSTHN0         = 1532
	FLOATIP_QUOTA_CANNOTLESSTHN0_MESSAGE = "浮动IP配额数必须大于等于0"

	NETWORK_QUOTA_CANNOTLESSTHN0         = 1533
	NETWORK_QUOTA_CANNOTLESSTHN0_MESSAGE = "内网配额数必须大于等于0"

	DATAPOOL_QUOTA_CANNOTLESSTHN0         = 1534
	DATAPOOL_QUOTA_CANNOTLESSTHN0_MESSAGE = "数据库集群配额数必须大于等于0"

	PRIVATENETWORK_QUOTA_CANNOTLESSTHN0         = 1535
	PRIVATENETWORK_QUOTA_CANNOTLESSTHN0_MESSAGE = "私有网络配额数必须大于等于0"

	LOADBALANCE_QUOTA_CANNOTLESSTHN0         = 1536
	LOADBALANCE_QUOTA_CANNOTLESSTHN0_MESSAGE = "负载均衡配额数必须大于等于0"

	PHONE_NOT_MATCH         = 1537
	PHONE_NOT_MATCH_MESSAGE = "手机格式有误"

	MAIL_NOT_MATCH         = 1538
	MAIL_NOT_MATCH_MESSAGE = "邮箱格式有误"

	REGION_NOT_EXIST         = 1539
	REGION_NOT_EXIST_MESSAGE = "区域不存在"

	EDITE_DADREGION_FAIL         = 1540
	EDITE_DADREGION_FAIL_MESSAGE = "修改父区域失败，区域名已存在"

	EDITE_SONREGION_FAIL         = 1541
	EDITE_SONREGION_FAIL_MESSAGE = "修改子区域失败，区域名已存在"
	//######################################################
)
