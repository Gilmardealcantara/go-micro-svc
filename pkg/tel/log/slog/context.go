package slog

import (
	"context"
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
	httputil "github.com/Gilmardealcantara/go-micro-svc/pkg/tel/http"
)

type logCtxBuilder struct {
	ctx context.Context
}

func NewContextBuilder(ctx context.Context) *logCtxBuilder {
	return &logCtxBuilder{ctx: ctx}
}

const LogHttpRequestContextKey string = "logHttpRequestContext"
const LogHttpResponseContextKey string = "logHttpResponseContext"
const LogAccountContextKey string = "logAccountContext"
const LogCustomContextKey string = "logCustomContext"

func (lc *logCtxBuilder) AddRequestInfo(r *http.Request) *logCtxBuilder {
	info := GetRequestInfo(r)
	lc.ctx = context.WithValue(lc.ctx, LogHttpRequestContextKey, info)
	return lc
}

func (lc *logCtxBuilder) AddResponseInfo(rww httputil.ResponseWriterWrapper) *logCtxBuilder {
	info := getResponseInfo(rww)
	lc.ctx = context.WithValue(lc.ctx, LogHttpResponseContextKey, info)
	return lc
}

func (lc *logCtxBuilder) AddAccountInfo(rww httputil.ResponseWriterWrapper) *logCtxBuilder {
	info := rww.Account()
	if info != nil {
		lc.ctx = context.WithValue(lc.ctx, LogAccountContextKey, info)
	}
	return lc
}

func (lc *logCtxBuilder) AddCustomInfo(info data.CustomDataType) *logCtxBuilder {
	lc.ctx = context.WithValue(lc.ctx, LogCustomContextKey, info)
	return lc
}

func (lc *logCtxBuilder) AddCustomInfoProp(k string, v any) *logCtxBuilder {
	customInfo, _ := CustomInfoFromContext(lc.ctx)
	if customInfo == nil {
		customInfo = make(data.CustomDataType)
	}
	customInfo[k] = v
	lc.ctx = context.WithValue(lc.ctx, LogCustomContextKey, customInfo)
	return lc
}

func (lc *logCtxBuilder) Build() context.Context {
	return lc.ctx
}

func HttpRequestInfoFromContext(ctx context.Context) (data.HttpRequestInfo, bool) {
	info, ok := ctx.Value(LogHttpRequestContextKey).(data.HttpRequestInfo)
	return info, ok
}

func HttpResponseInfoFromContext(ctx context.Context) (data.HttpResponseInfo, bool) {
	info, ok := ctx.Value(LogHttpResponseContextKey).(data.HttpResponseInfo)
	return info, ok
}

func AccountFromContext(ctx context.Context) (*data.Account, bool) {
	info, ok := ctx.Value(LogAccountContextKey).(*data.Account)
	return info, ok
}

func CustomInfoFromContext(ctx context.Context) (data.CustomDataType, bool) {
	info, ok := ctx.Value(LogCustomContextKey).(data.CustomDataType)
	return info, ok
}

func GetRequestInfo(r *http.Request) data.HttpRequestInfo {
	schoolId := r.Header.Get("X-SCHOOL-ID")
	if schoolId == "" {
		schoolId = r.PathValue("school_id")
	}
	accountId := r.Header.Get("X-ACCOUNT-ID")

	userId := r.Header.Get("X-USER-ID")

	return data.HttpRequestInfo{
		ClientService:  r.Header.Get("X-CLIENT-SERVICE"),
		IPAddress:      r.RemoteAddr,
		Method:         r.Method,
		Path:           r.URL.Path,
		UserAgent:      r.UserAgent(),
		Origin:         r.Header.Get("Origin"),
		XforwardedHost: r.Header.Get("X-FORWARDED-HOST"),
		SchoolId:       schoolId,
		AccountId:      accountId,
		UserId:         userId,
		RawQuery:       r.URL.RawQuery,
	}
}

func getResponseInfo(rww httputil.ResponseWriterWrapper) data.HttpResponseInfo {
	var body, errMsg string
	code := rww.Code()
	if code > 299 {
		body = rww.Body().String()
	}
	err := rww.Error()
	if err != nil {
		errMsg = err.Error()
	}
	return data.HttpResponseInfo{
		StatusCode: code,
		ErrorMsg:   errMsg,
		Body:       body,
	}
}
