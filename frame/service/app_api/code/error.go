package code

var (
	ErrorNetwork OutputCode = &logicCode{
		Code: 1000,
		Msg:  "网络错误",
	}
	ErrorServer OutputCode = &logicCode{
		Code: 1001,
		Msg:  "服务器故障",
	}
	ErrorRateLimit OutputCode = &logicCode{
		Code: 1002,
		Msg:  "网络限流，请稍后再试",
	}
	ErrorAuth OutputCode = &logicCode{
		Code: 1003,
		Msg:  "登录失效",
	}
	ErrorParam OutputCode = &logicCode{
		Code: 1004,
		Msg:  "%s错误",
	}
	ErrorUpload OutputCode = &logicCode{
		Code: 1005,
		Msg:  "上传失败",
	}
	ErrorUploadTooLarge OutputCode = &logicCode{
		Code: 1006,
		Msg:  "上传内容过大",
	}
	ErrorUploadType OutputCode = &logicCode{
		Code: 1007,
		Msg:  "上传文件类型有误",
	}
	ErrorReqTooMany OutputCode = &logicCode{
		Code: 1008,
		Msg:  "请示过于频繁，请稍后重试",
	}
	ErrorNickname OutputCode = &logicCode{
		Code: 1009,
		Msg:  "昵称长度不能超过10位",
	}
	ErrorNicknameRepeat OutputCode = &logicCode{
		Code: 1010,
		Msg:  "昵称已被占用",
	}

	ErrorUploadContent OutputCode = &logicCode{
		Code: 1011,
		Msg:  "上传的图片内容违规",
	}

	ErrorNoAuth OutputCode = &logicCode{
		Code: 1012,
		Msg:  "您没有权限操作",
	}

	ErrorQueryFailure OutputCode = &logicCode{
		Code: 1013,
		Msg:  "查询失败，请重试",
	}
	ErrorParameter OutputCode = &logicCode{
		Code: 1014,
		Msg:  "参数不全",
	}
	ErrorPasswordAccord OutputCode = &logicCode{
		Code: 1015,
		Msg:  "两次密码不一致",
	}
)
