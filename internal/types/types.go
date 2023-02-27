// Code generated by goctl. DO NOT EDIT.
package types

type ApiResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message" default:"success"`
	Data    interface{} `json:"data"`
}

type FileUploadReq struct {
	Types string `json:"types,optional"`
	Title string `json:"title,optional"`
	Ext   string `json:"ext,optional"`
	Size  int64  `json:"size,optional"`
	Path  string `json:"path,optional"`
	Hash  string `json:"hash,optional"`
	Bits  []byte `json:"bits,optional"`
}

type FileUploadResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Hash    string `json:"hash"`
	Link    string `json:"link"`
	Types   string `json:"types"`
	Ext     string `json:"ext"`
	AudioID int64  `json:"audio_id"`
}

type ListCourseReq struct {
	Title  string `form:"title,optional"`
	Author string `form:"author,optional"`
}

type CreateCourseReq struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Source string `json:"source,optional"`
}

type ListChapterReq struct {
	CourseID int64 `json:"course_id"`
}

type CreateChapterReq struct {
	Title    string `json:"title"`
	CourseID int64  `json:"course_id"`
	AudioID  int64  `json:"audio_id"`
	Hash     string `json:"hash"`
	Link     string `json:"link"`
	Ext      string `json:"ext"`
	Content  string `json:"content,optional"`
}
