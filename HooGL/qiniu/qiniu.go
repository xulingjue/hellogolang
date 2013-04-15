package qiniu

import (
	"fmt"
	qiniuConf "github.com/qiniu/api/conf"
	qiniuIo "github.com/qiniu/api/io"
	qiniuRs "github.com/qiniu/api/rs"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var policy qiniuRs.PutPolicy

func init() {
	qiniuConf.ACCESS_KEY = "i6F5TD_oZrApSI3PyAlucHX8dFVF9djQ3D5t4Rhf"
	qiniuConf.SECRET_KEY = "rg9AcJhLYTr29mR2wf5q6NEVGh-g9bpoqccQkg9T"

	policy = qiniuRs.PutPolicy{

		// [string] 必须, 指定授权上传的bucket
		Scope: "hellogolang",

		// [int] 表示有效时间为3600秒, 即一个小时
		Expires: 3600,

		// [string] 用于设置文件上传成功后，七牛云存储服务端要回调客户方的业务服务器地址
		CallbackUrl: "",

		// [string] 用于设置文件上传成功后，七牛云存储服务端向客户方的业务服务器发送回调请求的 `Content-Type`
		CallbackBodyType: "",

		// [string] 客户方终端用户（End User）的ID，该字段可以用来标示一个文件的属主，这在一些特殊场景下（比如给终端用户上传的图片打上名字水印）非常有用。
		Customer: "1",

		// [string] 用于设置文件上传成功后，执行指定的预转指令。
		// 参考 http://docs.qiniutek.com/v3/api/io/#uploadToken-asyncOps
		AsyncOps: "",

		// [uint16] 可选值 0 或者 1，缺省为 0 。值为 1 表示 callback 传递的自定义数据中允许存在转义符号 `$(VarExpression)
		// 参考 http://docs.qiniutek.com/v3/api/words/#VarExpression
		Escape: 0,

		// [uint16] 可选值 0 或者 1, 缺省为 0. 值为 1 表示在服务端自动识别上传文件类型.
		DetectMime: 0,
	}
}

func UploadAvatar(file *os.File, key string) error {
	// 生成 uploadToken, string类型
	uploadToken := policy.Token()

	// 声明 PutExtra
	extra := &qiniuIo.PutExtra{
		// [string] 必选, 指定上传的目标仓库
		Bucket: "hellogolang",

		// [string] 可选。在 uptoken 没有指定 DetectMime 时，用户客户端可自己指定 MimeType
		MimeType: "",

		// [string] 可选。用户自定义 Meta，不能超过 256 字节
		CustomMeta: "",

		// [string] 当 uptoken 指定了 CallbackUrl，则 CallbackParams 必须非空
		CallbackParams: "",
	}

	ret := new(qiniuIo.PutRet)
	err := qiniuIo.Put(nil, ret, uploadToken, key, file, extra) // PutRet, error
	if err != nil {
		// 上传失败
		fmt.Println(err)
		return err
	}
	return nil
}

func GetFk(userid int64) string {
	var fk [24]byte
	timeByte := []byte(strconv.FormatInt(time.Now().Unix(), 10))

	userByte := []byte(strconv.FormatInt(userid, 10))

	length := 0
	for length < 6 {
		fk[length] = byte(rand.Int31n(256))
		length++

	}

	for length < 14 {
		if len(userByte) > (length - 6) {
			fk[length] = userByte[length-6] - 48
			fmt.Println(fk[length])
		} else {
			fk[length] = '0' - 48
		}
		length++

	}

	for length < 24 {
		fk[length] = timeByte[length-14] - 48
		length++

	}

	result := fmt.Sprintf("%02x%02x%02x%02x%02x%02x%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d",
		fk[0], fk[1], fk[2], fk[3], fk[4], fk[5],
		fk[6], fk[7], fk[8], fk[9], fk[10], fk[11], fk[12], fk[13],
		fk[14], fk[15], fk[16], fk[17], fk[18], fk[19], fk[20], fk[21], fk[22], fk[23])

	return result
}

func GetExt(filename string) string {
	rep := regexp.MustCompile("\\.")

	point := []byte(filename)
	index := rep.FindIndex(point)

	if len(index) > 0 {
		return filename[index[1]:len(filename)]
	}
	return ""
}
