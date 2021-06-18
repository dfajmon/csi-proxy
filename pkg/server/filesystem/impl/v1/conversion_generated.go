// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubernetes-csi/csi-proxy/client/api/filesystem/v1"
	impl "github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl"
)

func autoConvert_v1_CreateSymlinkRequest_To_impl_CreateSymlinkRequest(in *v1.CreateSymlinkRequest, out *impl.CreateSymlinkRequest) error {
	out.SourcePath = in.SourcePath
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_v1_CreateSymlinkRequest_To_impl_CreateSymlinkRequest is an autogenerated conversion function.
func Convert_v1_CreateSymlinkRequest_To_impl_CreateSymlinkRequest(in *v1.CreateSymlinkRequest, out *impl.CreateSymlinkRequest) error {
	return autoConvert_v1_CreateSymlinkRequest_To_impl_CreateSymlinkRequest(in, out)
}

func autoConvert_impl_CreateSymlinkRequest_To_v1_CreateSymlinkRequest(in *impl.CreateSymlinkRequest, out *v1.CreateSymlinkRequest) error {
	out.SourcePath = in.SourcePath
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_impl_CreateSymlinkRequest_To_v1_CreateSymlinkRequest is an autogenerated conversion function.
func Convert_impl_CreateSymlinkRequest_To_v1_CreateSymlinkRequest(in *impl.CreateSymlinkRequest, out *v1.CreateSymlinkRequest) error {
	return autoConvert_impl_CreateSymlinkRequest_To_v1_CreateSymlinkRequest(in, out)
}

func autoConvert_v1_CreateSymlinkResponse_To_impl_CreateSymlinkResponse(in *v1.CreateSymlinkResponse, out *impl.CreateSymlinkResponse) error {
	return nil
}

// Convert_v1_CreateSymlinkResponse_To_impl_CreateSymlinkResponse is an autogenerated conversion function.
func Convert_v1_CreateSymlinkResponse_To_impl_CreateSymlinkResponse(in *v1.CreateSymlinkResponse, out *impl.CreateSymlinkResponse) error {
	return autoConvert_v1_CreateSymlinkResponse_To_impl_CreateSymlinkResponse(in, out)
}

func autoConvert_impl_CreateSymlinkResponse_To_v1_CreateSymlinkResponse(in *impl.CreateSymlinkResponse, out *v1.CreateSymlinkResponse) error {
	return nil
}

// Convert_impl_CreateSymlinkResponse_To_v1_CreateSymlinkResponse is an autogenerated conversion function.
func Convert_impl_CreateSymlinkResponse_To_v1_CreateSymlinkResponse(in *impl.CreateSymlinkResponse, out *v1.CreateSymlinkResponse) error {
	return autoConvert_impl_CreateSymlinkResponse_To_v1_CreateSymlinkResponse(in, out)
}

func autoConvert_v1_IsSymlinkRequest_To_impl_IsSymlinkRequest(in *v1.IsSymlinkRequest, out *impl.IsSymlinkRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1_IsSymlinkRequest_To_impl_IsSymlinkRequest is an autogenerated conversion function.
func Convert_v1_IsSymlinkRequest_To_impl_IsSymlinkRequest(in *v1.IsSymlinkRequest, out *impl.IsSymlinkRequest) error {
	return autoConvert_v1_IsSymlinkRequest_To_impl_IsSymlinkRequest(in, out)
}

func autoConvert_impl_IsSymlinkRequest_To_v1_IsSymlinkRequest(in *impl.IsSymlinkRequest, out *v1.IsSymlinkRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_impl_IsSymlinkRequest_To_v1_IsSymlinkRequest is an autogenerated conversion function.
func Convert_impl_IsSymlinkRequest_To_v1_IsSymlinkRequest(in *impl.IsSymlinkRequest, out *v1.IsSymlinkRequest) error {
	return autoConvert_impl_IsSymlinkRequest_To_v1_IsSymlinkRequest(in, out)
}

func autoConvert_v1_IsSymlinkResponse_To_impl_IsSymlinkResponse(in *v1.IsSymlinkResponse, out *impl.IsSymlinkResponse) error {
	out.IsSymlink = in.IsSymlink
	return nil
}

// Convert_v1_IsSymlinkResponse_To_impl_IsSymlinkResponse is an autogenerated conversion function.
func Convert_v1_IsSymlinkResponse_To_impl_IsSymlinkResponse(in *v1.IsSymlinkResponse, out *impl.IsSymlinkResponse) error {
	return autoConvert_v1_IsSymlinkResponse_To_impl_IsSymlinkResponse(in, out)
}

func autoConvert_impl_IsSymlinkResponse_To_v1_IsSymlinkResponse(in *impl.IsSymlinkResponse, out *v1.IsSymlinkResponse) error {
	out.IsSymlink = in.IsSymlink
	return nil
}

// Convert_impl_IsSymlinkResponse_To_v1_IsSymlinkResponse is an autogenerated conversion function.
func Convert_impl_IsSymlinkResponse_To_v1_IsSymlinkResponse(in *impl.IsSymlinkResponse, out *v1.IsSymlinkResponse) error {
	return autoConvert_impl_IsSymlinkResponse_To_v1_IsSymlinkResponse(in, out)
}

func autoConvert_v1_MkdirRequest_To_impl_MkdirRequest(in *v1.MkdirRequest, out *impl.MkdirRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1_MkdirRequest_To_impl_MkdirRequest is an autogenerated conversion function.
func Convert_v1_MkdirRequest_To_impl_MkdirRequest(in *v1.MkdirRequest, out *impl.MkdirRequest) error {
	return autoConvert_v1_MkdirRequest_To_impl_MkdirRequest(in, out)
}

func autoConvert_impl_MkdirRequest_To_v1_MkdirRequest(in *impl.MkdirRequest, out *v1.MkdirRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_impl_MkdirRequest_To_v1_MkdirRequest is an autogenerated conversion function.
func Convert_impl_MkdirRequest_To_v1_MkdirRequest(in *impl.MkdirRequest, out *v1.MkdirRequest) error {
	return autoConvert_impl_MkdirRequest_To_v1_MkdirRequest(in, out)
}

func autoConvert_v1_MkdirResponse_To_impl_MkdirResponse(in *v1.MkdirResponse, out *impl.MkdirResponse) error {
	return nil
}

// Convert_v1_MkdirResponse_To_impl_MkdirResponse is an autogenerated conversion function.
func Convert_v1_MkdirResponse_To_impl_MkdirResponse(in *v1.MkdirResponse, out *impl.MkdirResponse) error {
	return autoConvert_v1_MkdirResponse_To_impl_MkdirResponse(in, out)
}

func autoConvert_impl_MkdirResponse_To_v1_MkdirResponse(in *impl.MkdirResponse, out *v1.MkdirResponse) error {
	return nil
}

// Convert_impl_MkdirResponse_To_v1_MkdirResponse is an autogenerated conversion function.
func Convert_impl_MkdirResponse_To_v1_MkdirResponse(in *impl.MkdirResponse, out *v1.MkdirResponse) error {
	return autoConvert_impl_MkdirResponse_To_v1_MkdirResponse(in, out)
}

func autoConvert_v1_PathExistsRequest_To_impl_PathExistsRequest(in *v1.PathExistsRequest, out *impl.PathExistsRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1_PathExistsRequest_To_impl_PathExistsRequest is an autogenerated conversion function.
func Convert_v1_PathExistsRequest_To_impl_PathExistsRequest(in *v1.PathExistsRequest, out *impl.PathExistsRequest) error {
	return autoConvert_v1_PathExistsRequest_To_impl_PathExistsRequest(in, out)
}

func autoConvert_impl_PathExistsRequest_To_v1_PathExistsRequest(in *impl.PathExistsRequest, out *v1.PathExistsRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_impl_PathExistsRequest_To_v1_PathExistsRequest is an autogenerated conversion function.
func Convert_impl_PathExistsRequest_To_v1_PathExistsRequest(in *impl.PathExistsRequest, out *v1.PathExistsRequest) error {
	return autoConvert_impl_PathExistsRequest_To_v1_PathExistsRequest(in, out)
}

func autoConvert_v1_PathExistsResponse_To_impl_PathExistsResponse(in *v1.PathExistsResponse, out *impl.PathExistsResponse) error {
	out.Exists = in.Exists
	return nil
}

// Convert_v1_PathExistsResponse_To_impl_PathExistsResponse is an autogenerated conversion function.
func Convert_v1_PathExistsResponse_To_impl_PathExistsResponse(in *v1.PathExistsResponse, out *impl.PathExistsResponse) error {
	return autoConvert_v1_PathExistsResponse_To_impl_PathExistsResponse(in, out)
}

func autoConvert_impl_PathExistsResponse_To_v1_PathExistsResponse(in *impl.PathExistsResponse, out *v1.PathExistsResponse) error {
	out.Exists = in.Exists
	return nil
}

// Convert_impl_PathExistsResponse_To_v1_PathExistsResponse is an autogenerated conversion function.
func Convert_impl_PathExistsResponse_To_v1_PathExistsResponse(in *impl.PathExistsResponse, out *v1.PathExistsResponse) error {
	return autoConvert_impl_PathExistsResponse_To_v1_PathExistsResponse(in, out)
}

func autoConvert_v1_RmdirRequest_To_impl_RmdirRequest(in *v1.RmdirRequest, out *impl.RmdirRequest) error {
	out.Path = in.Path
	out.Force = in.Force
	return nil
}

// Convert_v1_RmdirRequest_To_impl_RmdirRequest is an autogenerated conversion function.
func Convert_v1_RmdirRequest_To_impl_RmdirRequest(in *v1.RmdirRequest, out *impl.RmdirRequest) error {
	return autoConvert_v1_RmdirRequest_To_impl_RmdirRequest(in, out)
}

func autoConvert_impl_RmdirRequest_To_v1_RmdirRequest(in *impl.RmdirRequest, out *v1.RmdirRequest) error {
	out.Path = in.Path
	out.Force = in.Force
	return nil
}

// Convert_impl_RmdirRequest_To_v1_RmdirRequest is an autogenerated conversion function.
func Convert_impl_RmdirRequest_To_v1_RmdirRequest(in *impl.RmdirRequest, out *v1.RmdirRequest) error {
	return autoConvert_impl_RmdirRequest_To_v1_RmdirRequest(in, out)
}

func autoConvert_v1_RmdirResponse_To_impl_RmdirResponse(in *v1.RmdirResponse, out *impl.RmdirResponse) error {
	return nil
}

// Convert_v1_RmdirResponse_To_impl_RmdirResponse is an autogenerated conversion function.
func Convert_v1_RmdirResponse_To_impl_RmdirResponse(in *v1.RmdirResponse, out *impl.RmdirResponse) error {
	return autoConvert_v1_RmdirResponse_To_impl_RmdirResponse(in, out)
}

func autoConvert_impl_RmdirResponse_To_v1_RmdirResponse(in *impl.RmdirResponse, out *v1.RmdirResponse) error {
	return nil
}

// Convert_impl_RmdirResponse_To_v1_RmdirResponse is an autogenerated conversion function.
func Convert_impl_RmdirResponse_To_v1_RmdirResponse(in *impl.RmdirResponse, out *v1.RmdirResponse) error {
	return autoConvert_impl_RmdirResponse_To_v1_RmdirResponse(in, out)
}