// Code generated by protoc-gen-gohttp. DO NOT EDIT.
// source: helloworld/helloworld.proto

package helloworldpb

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	ioutil "io/ioutil"
	mime "mime"
	http "net/http"
	strings "strings"
)

// GreeterHTTPService is the server API for Greeter service.
type GreeterHTTPService interface {
	// SayHello says hello.
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// GreeterHTTPConverter has a function to convert GreeterHTTPService interface to http.HandlerFunc.
type GreeterHTTPConverter struct {
	srv GreeterHTTPService
}

// NewGreeterHTTPConverter returns GreeterHTTPConverter.
func NewGreeterHTTPConverter(srv GreeterHTTPService) *GreeterHTTPConverter {
	return &GreeterHTTPConverter{
		srv: srv,
	}
}

// SayHello returns GreeterHTTPService interface's SayHello converted to http.HandlerFunc.
//
// SayHello says hello.
func (h *GreeterHTTPConverter) SayHello(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type")); contentType {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					if err := json.NewEncoder(w).Encode(p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		arg := &HelloRequest{}
		contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if r.Method != http.MethodGet {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}

			switch contentType {
			case "application/protobuf", "application/x-protobuf":
				if err := proto.Unmarshal(body, arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			case "application/json":
				if err := jsonpb.Unmarshal(bytes.NewBuffer(body), arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
				cb(ctx, w, r, nil, nil, err)
				return
			}
		}

		n := len(interceptors)
		chained := func(ctx context.Context, arg interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return currentInter(currentCtx, currentReq, info, currentHandler)
				}
			}

			chainedHandler := handler
			for i := n - 1; i >= 0; i-- {
				chainedHandler = chainer(interceptors[i], chainedHandler)
			}
			return chainedHandler(ctx, arg)
		}

		info := &grpc.UnaryServerInfo{
			Server:     h.srv,
			FullMethod: "/helloworld.Greeter/SayHello",
		}

		handler := func(c context.Context, req interface{}) (interface{}, error) {
			return h.srv.SayHello(c, req.(*HelloRequest))
		}

		iret, err := chained(ctx, arg, info, handler)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		ret, ok := iret.(*HelloReply)
		if !ok {
			cb(ctx, w, r, arg, nil, fmt.Errorf("/helloworld.Greeter/SayHello: interceptors have not return HelloReply"))
			return
		}

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			m := jsonpb.Marshaler{
				EnumsAsInts:  true,
				EmitDefaults: true,
			}
			if err := m.Marshal(w, ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

// SayHelloWithName returns Service name, Method name and GreeterHTTPService interface's SayHello converted to http.HandlerFunc.
//
// SayHello says hello.
func (h *GreeterHTTPConverter) SayHelloWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) (string, string, http.HandlerFunc) {
	return "Greeter", "SayHello", h.SayHello(cb, interceptors...)
}
