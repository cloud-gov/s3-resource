// This file was generated by counterfeiter
package fakes

import (
	"sync"

	s3resource "github.com/concourse/s3-resource"
)

type FakeS3Client struct {
	BucketFileVersionsStub        func(string, string) ([]string, error)
	bucketFileVersionsMutex       sync.RWMutex
	bucketFileVersionsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	bucketFileVersionsReturns struct {
		result1 []string
		result2 error
	}
	bucketFileVersionsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	BucketFilesStub        func(string, string) ([]string, error)
	bucketFilesMutex       sync.RWMutex
	bucketFilesArgsForCall []struct {
		arg1 string
		arg2 string
	}
	bucketFilesReturns struct {
		result1 []string
		result2 error
	}
	bucketFilesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	DeleteFileStub        func(string, string) error
	deleteFileMutex       sync.RWMutex
	deleteFileArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteFileReturns struct {
		result1 error
	}
	deleteFileReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteVersionedFileStub        func(string, string, string) error
	deleteVersionedFileMutex       sync.RWMutex
	deleteVersionedFileArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	deleteVersionedFileReturns struct {
		result1 error
	}
	deleteVersionedFileReturnsOnCall map[int]struct {
		result1 error
	}
	DownloadFileStub        func(string, string, string, string) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	downloadFileReturns struct {
		result1 error
	}
	downloadFileReturnsOnCall map[int]struct {
		result1 error
	}
	DownloadTagsStub        func(string, string, string, string) error
	downloadTagsMutex       sync.RWMutex
	downloadTagsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	downloadTagsReturns struct {
		result1 error
	}
	downloadTagsReturnsOnCall map[int]struct {
		result1 error
	}
	SetTagsStub        func(string, string, string, map[string]string) error
	setTagsMutex       sync.RWMutex
	setTagsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 map[string]string
	}
	setTagsReturns struct {
		result1 error
	}
	setTagsReturnsOnCall map[int]struct {
		result1 error
	}
	URLStub        func(string, string, bool, string) string
	uRLMutex       sync.RWMutex
	uRLArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
	}
	uRLReturns struct {
		result1 string
	}
	uRLReturnsOnCall map[int]struct {
		result1 string
	}
	UploadFileStub        func(string, string, string, s3resource.UploadFileOptions) (string, error)
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 s3resource.UploadFileOptions
	}
	uploadFileReturns struct {
		result1 string
		result2 error
	}
	uploadFileReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeS3Client) BucketFileVersions(arg1 string, arg2 string) ([]string, error) {
	fake.bucketFileVersionsMutex.Lock()
	ret, specificReturn := fake.bucketFileVersionsReturnsOnCall[len(fake.bucketFileVersionsArgsForCall)]
	fake.bucketFileVersionsArgsForCall = append(fake.bucketFileVersionsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("BucketFileVersions", []interface{}{arg1, arg2})
	fake.bucketFileVersionsMutex.Unlock()
	if fake.BucketFileVersionsStub != nil {
		return fake.BucketFileVersionsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.bucketFileVersionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeS3Client) BucketFileVersionsCallCount() int {
	fake.bucketFileVersionsMutex.RLock()
	defer fake.bucketFileVersionsMutex.RUnlock()
	return len(fake.bucketFileVersionsArgsForCall)
}

func (fake *FakeS3Client) BucketFileVersionsCalls(stub func(string, string) ([]string, error)) {
	fake.bucketFileVersionsMutex.Lock()
	defer fake.bucketFileVersionsMutex.Unlock()
	fake.BucketFileVersionsStub = stub
}

func (fake *FakeS3Client) BucketFileVersionsArgsForCall(i int) (string, string) {
	fake.bucketFileVersionsMutex.RLock()
	defer fake.bucketFileVersionsMutex.RUnlock()
	argsForCall := fake.bucketFileVersionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeS3Client) BucketFileVersionsReturns(result1 []string, result2 error) {
	fake.bucketFileVersionsMutex.Lock()
	defer fake.bucketFileVersionsMutex.Unlock()
	fake.BucketFileVersionsStub = nil
	fake.bucketFileVersionsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) BucketFileVersionsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.bucketFileVersionsMutex.Lock()
	defer fake.bucketFileVersionsMutex.Unlock()
	fake.BucketFileVersionsStub = nil
	if fake.bucketFileVersionsReturnsOnCall == nil {
		fake.bucketFileVersionsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.bucketFileVersionsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) BucketFiles(arg1 string, arg2 string) ([]string, error) {
	fake.bucketFilesMutex.Lock()
	ret, specificReturn := fake.bucketFilesReturnsOnCall[len(fake.bucketFilesArgsForCall)]
	fake.bucketFilesArgsForCall = append(fake.bucketFilesArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("BucketFiles", []interface{}{arg1, arg2})
	fake.bucketFilesMutex.Unlock()
	if fake.BucketFilesStub != nil {
		return fake.BucketFilesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.bucketFilesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeS3Client) BucketFilesCallCount() int {
	fake.bucketFilesMutex.RLock()
	defer fake.bucketFilesMutex.RUnlock()
	return len(fake.bucketFilesArgsForCall)
}

func (fake *FakeS3Client) BucketFilesCalls(stub func(string, string) ([]string, error)) {
	fake.bucketFilesMutex.Lock()
	defer fake.bucketFilesMutex.Unlock()
	fake.BucketFilesStub = stub
}

func (fake *FakeS3Client) BucketFilesArgsForCall(i int) (string, string) {
	fake.bucketFilesMutex.RLock()
	defer fake.bucketFilesMutex.RUnlock()
	argsForCall := fake.bucketFilesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeS3Client) BucketFilesReturns(result1 []string, result2 error) {
	fake.bucketFilesMutex.Lock()
	defer fake.bucketFilesMutex.Unlock()
	fake.BucketFilesStub = nil
	fake.bucketFilesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) BucketFilesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.bucketFilesMutex.Lock()
	defer fake.bucketFilesMutex.Unlock()
	fake.BucketFilesStub = nil
	if fake.bucketFilesReturnsOnCall == nil {
		fake.bucketFilesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.bucketFilesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) DeleteFile(arg1 string, arg2 string) error {
	fake.deleteFileMutex.Lock()
	ret, specificReturn := fake.deleteFileReturnsOnCall[len(fake.deleteFileArgsForCall)]
	fake.deleteFileArgsForCall = append(fake.deleteFileArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteFile", []interface{}{arg1, arg2})
	fake.deleteFileMutex.Unlock()
	if fake.DeleteFileStub != nil {
		return fake.DeleteFileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteFileReturns
	return fakeReturns.result1
}

func (fake *FakeS3Client) DeleteFileCallCount() int {
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	return len(fake.deleteFileArgsForCall)
}

func (fake *FakeS3Client) DeleteFileCalls(stub func(string, string) error) {
	fake.deleteFileMutex.Lock()
	defer fake.deleteFileMutex.Unlock()
	fake.DeleteFileStub = stub
}

func (fake *FakeS3Client) DeleteFileArgsForCall(i int) (string, string) {
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	argsForCall := fake.deleteFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeS3Client) DeleteFileReturns(result1 error) {
	fake.deleteFileMutex.Lock()
	defer fake.deleteFileMutex.Unlock()
	fake.DeleteFileStub = nil
	fake.deleteFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DeleteFileReturnsOnCall(i int, result1 error) {
	fake.deleteFileMutex.Lock()
	defer fake.deleteFileMutex.Unlock()
	fake.DeleteFileStub = nil
	if fake.deleteFileReturnsOnCall == nil {
		fake.deleteFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DeleteVersionedFile(arg1 string, arg2 string, arg3 string) error {
	fake.deleteVersionedFileMutex.Lock()
	ret, specificReturn := fake.deleteVersionedFileReturnsOnCall[len(fake.deleteVersionedFileArgsForCall)]
	fake.deleteVersionedFileArgsForCall = append(fake.deleteVersionedFileArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteVersionedFile", []interface{}{arg1, arg2, arg3})
	fake.deleteVersionedFileMutex.Unlock()
	if fake.DeleteVersionedFileStub != nil {
		return fake.DeleteVersionedFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteVersionedFileReturns
	return fakeReturns.result1
}

func (fake *FakeS3Client) DeleteVersionedFileCallCount() int {
	fake.deleteVersionedFileMutex.RLock()
	defer fake.deleteVersionedFileMutex.RUnlock()
	return len(fake.deleteVersionedFileArgsForCall)
}

func (fake *FakeS3Client) DeleteVersionedFileCalls(stub func(string, string, string) error) {
	fake.deleteVersionedFileMutex.Lock()
	defer fake.deleteVersionedFileMutex.Unlock()
	fake.DeleteVersionedFileStub = stub
}

func (fake *FakeS3Client) DeleteVersionedFileArgsForCall(i int) (string, string, string) {
	fake.deleteVersionedFileMutex.RLock()
	defer fake.deleteVersionedFileMutex.RUnlock()
	argsForCall := fake.deleteVersionedFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeS3Client) DeleteVersionedFileReturns(result1 error) {
	fake.deleteVersionedFileMutex.Lock()
	defer fake.deleteVersionedFileMutex.Unlock()
	fake.DeleteVersionedFileStub = nil
	fake.deleteVersionedFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DeleteVersionedFileReturnsOnCall(i int, result1 error) {
	fake.deleteVersionedFileMutex.Lock()
	defer fake.deleteVersionedFileMutex.Unlock()
	fake.DeleteVersionedFileStub = nil
	if fake.deleteVersionedFileReturnsOnCall == nil {
		fake.deleteVersionedFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteVersionedFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DownloadFile(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.downloadFileMutex.Lock()
	ret, specificReturn := fake.downloadFileReturnsOnCall[len(fake.downloadFileArgsForCall)]
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DownloadFile", []interface{}{arg1, arg2, arg3, arg4})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadFileReturns
	return fakeReturns.result1
}

func (fake *FakeS3Client) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakeS3Client) DownloadFileCalls(stub func(string, string, string, string) error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = stub
}

func (fake *FakeS3Client) DownloadFileArgsForCall(i int) (string, string, string, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	argsForCall := fake.downloadFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeS3Client) DownloadFileReturns(result1 error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DownloadFileReturnsOnCall(i int, result1 error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = nil
	if fake.downloadFileReturnsOnCall == nil {
		fake.downloadFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DownloadTags(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.downloadTagsMutex.Lock()
	ret, specificReturn := fake.downloadTagsReturnsOnCall[len(fake.downloadTagsArgsForCall)]
	fake.downloadTagsArgsForCall = append(fake.downloadTagsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DownloadTags", []interface{}{arg1, arg2, arg3, arg4})
	fake.downloadTagsMutex.Unlock()
	if fake.DownloadTagsStub != nil {
		return fake.DownloadTagsStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadTagsReturns
	return fakeReturns.result1
}

func (fake *FakeS3Client) DownloadTagsCallCount() int {
	fake.downloadTagsMutex.RLock()
	defer fake.downloadTagsMutex.RUnlock()
	return len(fake.downloadTagsArgsForCall)
}

func (fake *FakeS3Client) DownloadTagsCalls(stub func(string, string, string, string) error) {
	fake.downloadTagsMutex.Lock()
	defer fake.downloadTagsMutex.Unlock()
	fake.DownloadTagsStub = stub
}

func (fake *FakeS3Client) DownloadTagsArgsForCall(i int) (string, string, string, string) {
	fake.downloadTagsMutex.RLock()
	defer fake.downloadTagsMutex.RUnlock()
	argsForCall := fake.downloadTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeS3Client) DownloadTagsReturns(result1 error) {
	fake.downloadTagsMutex.Lock()
	defer fake.downloadTagsMutex.Unlock()
	fake.DownloadTagsStub = nil
	fake.downloadTagsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DownloadTagsReturnsOnCall(i int, result1 error) {
	fake.downloadTagsMutex.Lock()
	defer fake.downloadTagsMutex.Unlock()
	fake.DownloadTagsStub = nil
	if fake.downloadTagsReturnsOnCall == nil {
		fake.downloadTagsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadTagsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) SetTags(arg1 string, arg2 string, arg3 string, arg4 map[string]string) error {
	fake.setTagsMutex.Lock()
	ret, specificReturn := fake.setTagsReturnsOnCall[len(fake.setTagsArgsForCall)]
	fake.setTagsArgsForCall = append(fake.setTagsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 map[string]string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetTags", []interface{}{arg1, arg2, arg3, arg4})
	fake.setTagsMutex.Unlock()
	if fake.SetTagsStub != nil {
		return fake.SetTagsStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setTagsReturns
	return fakeReturns.result1
}

func (fake *FakeS3Client) SetTagsCallCount() int {
	fake.setTagsMutex.RLock()
	defer fake.setTagsMutex.RUnlock()
	return len(fake.setTagsArgsForCall)
}

func (fake *FakeS3Client) SetTagsCalls(stub func(string, string, string, map[string]string) error) {
	fake.setTagsMutex.Lock()
	defer fake.setTagsMutex.Unlock()
	fake.SetTagsStub = stub
}

func (fake *FakeS3Client) SetTagsArgsForCall(i int) (string, string, string, map[string]string) {
	fake.setTagsMutex.RLock()
	defer fake.setTagsMutex.RUnlock()
	argsForCall := fake.setTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeS3Client) SetTagsReturns(result1 error) {
	fake.setTagsMutex.Lock()
	defer fake.setTagsMutex.Unlock()
	fake.SetTagsStub = nil
	fake.setTagsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) SetTagsReturnsOnCall(i int, result1 error) {
	fake.setTagsMutex.Lock()
	defer fake.setTagsMutex.Unlock()
	fake.SetTagsStub = nil
	if fake.setTagsReturnsOnCall == nil {
		fake.setTagsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setTagsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) URL(arg1 string, arg2 string, arg3 bool, arg4 string) string {
	fake.uRLMutex.Lock()
	ret, specificReturn := fake.uRLReturnsOnCall[len(fake.uRLArgsForCall)]
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("URL", []interface{}{arg1, arg2, arg3, arg4})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uRLReturns
	return fakeReturns.result1
}

func (fake *FakeS3Client) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeS3Client) URLCalls(stub func(string, string, bool, string) string) {
	fake.uRLMutex.Lock()
	defer fake.uRLMutex.Unlock()
	fake.URLStub = stub
}

func (fake *FakeS3Client) URLArgsForCall(i int) (string, string, bool, string) {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	argsForCall := fake.uRLArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeS3Client) URLReturns(result1 string) {
	fake.uRLMutex.Lock()
	defer fake.uRLMutex.Unlock()
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeS3Client) URLReturnsOnCall(i int, result1 string) {
	fake.uRLMutex.Lock()
	defer fake.uRLMutex.Unlock()
	fake.URLStub = nil
	if fake.uRLReturnsOnCall == nil {
		fake.uRLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uRLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeS3Client) UploadFile(arg1 string, arg2 string, arg3 string, arg4 s3resource.UploadFileOptions) (string, error) {
	fake.uploadFileMutex.Lock()
	ret, specificReturn := fake.uploadFileReturnsOnCall[len(fake.uploadFileArgsForCall)]
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 s3resource.UploadFileOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("UploadFile", []interface{}{arg1, arg2, arg3, arg4})
	fake.uploadFileMutex.Unlock()
	if fake.UploadFileStub != nil {
		return fake.UploadFileStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.uploadFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeS3Client) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *FakeS3Client) UploadFileCalls(stub func(string, string, string, s3resource.UploadFileOptions) (string, error)) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = stub
}

func (fake *FakeS3Client) UploadFileArgsForCall(i int) (string, string, string, s3resource.UploadFileOptions) {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	argsForCall := fake.uploadFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeS3Client) UploadFileReturns(result1 string, result2 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) UploadFileReturnsOnCall(i int, result1 string, result2 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	if fake.uploadFileReturnsOnCall == nil {
		fake.uploadFileReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.uploadFileReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bucketFileVersionsMutex.RLock()
	defer fake.bucketFileVersionsMutex.RUnlock()
	fake.bucketFilesMutex.RLock()
	defer fake.bucketFilesMutex.RUnlock()
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	fake.deleteVersionedFileMutex.RLock()
	defer fake.deleteVersionedFileMutex.RUnlock()
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	fake.downloadTagsMutex.RLock()
	defer fake.downloadTagsMutex.RUnlock()
	fake.setTagsMutex.RLock()
	defer fake.setTagsMutex.RUnlock()
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeS3Client) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ s3resource.S3Client = new(FakeS3Client)
