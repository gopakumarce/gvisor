// automatically generated by stateify.

package vfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *anonFilesystemType) StateTypeName() string {
	return "pkg/sentry/vfs.anonFilesystemType"
}

func (x *anonFilesystemType) StateFields() []string {
	return []string{}
}

func (x *anonFilesystemType) beforeSave() {}

func (x *anonFilesystemType) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *anonFilesystemType) afterLoad() {}

func (x *anonFilesystemType) StateLoad(m state.Source) {
}

func (x *anonFilesystem) StateTypeName() string {
	return "pkg/sentry/vfs.anonFilesystem"
}

func (x *anonFilesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"devMinor",
	}
}

func (x *anonFilesystem) beforeSave() {}

func (x *anonFilesystem) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfs)
	m.Save(1, &x.devMinor)
}

func (x *anonFilesystem) afterLoad() {}

func (x *anonFilesystem) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfs)
	m.Load(1, &x.devMinor)
}

func (x *anonDentry) StateTypeName() string {
	return "pkg/sentry/vfs.anonDentry"
}

func (x *anonDentry) StateFields() []string {
	return []string{
		"vfsd",
		"name",
	}
}

func (x *anonDentry) beforeSave() {}

func (x *anonDentry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsd)
	m.Save(1, &x.name)
}

func (x *anonDentry) afterLoad() {}

func (x *anonDentry) StateLoad(m state.Source) {
	m.Load(0, &x.vfsd)
	m.Load(1, &x.name)
}

func (x *Dentry) StateTypeName() string {
	return "pkg/sentry/vfs.Dentry"
}

func (x *Dentry) StateFields() []string {
	return []string{
		"dead",
		"mounts",
		"impl",
	}
}

func (x *Dentry) beforeSave() {}

func (x *Dentry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dead)
	m.Save(1, &x.mounts)
	m.Save(2, &x.impl)
}

func (x *Dentry) afterLoad() {}

func (x *Dentry) StateLoad(m state.Source) {
	m.Load(0, &x.dead)
	m.Load(1, &x.mounts)
	m.Load(2, &x.impl)
}

func (x *DeviceKind) StateTypeName() string {
	return "pkg/sentry/vfs.DeviceKind"
}

func (x *DeviceKind) StateFields() []string {
	return nil
}

func (x *devTuple) StateTypeName() string {
	return "pkg/sentry/vfs.devTuple"
}

func (x *devTuple) StateFields() []string {
	return []string{
		"kind",
		"major",
		"minor",
	}
}

func (x *devTuple) beforeSave() {}

func (x *devTuple) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.kind)
	m.Save(1, &x.major)
	m.Save(2, &x.minor)
}

func (x *devTuple) afterLoad() {}

func (x *devTuple) StateLoad(m state.Source) {
	m.Load(0, &x.kind)
	m.Load(1, &x.major)
	m.Load(2, &x.minor)
}

func (x *registeredDevice) StateTypeName() string {
	return "pkg/sentry/vfs.registeredDevice"
}

func (x *registeredDevice) StateFields() []string {
	return []string{
		"dev",
		"opts",
	}
}

func (x *registeredDevice) beforeSave() {}

func (x *registeredDevice) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dev)
	m.Save(1, &x.opts)
}

func (x *registeredDevice) afterLoad() {}

func (x *registeredDevice) StateLoad(m state.Source) {
	m.Load(0, &x.dev)
	m.Load(1, &x.opts)
}

func (x *RegisterDeviceOptions) StateTypeName() string {
	return "pkg/sentry/vfs.RegisterDeviceOptions"
}

func (x *RegisterDeviceOptions) StateFields() []string {
	return []string{
		"GroupName",
	}
}

func (x *RegisterDeviceOptions) beforeSave() {}

func (x *RegisterDeviceOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.GroupName)
}

func (x *RegisterDeviceOptions) afterLoad() {}

func (x *RegisterDeviceOptions) StateLoad(m state.Source) {
	m.Load(0, &x.GroupName)
}

func (x *EpollInstance) StateTypeName() string {
	return "pkg/sentry/vfs.EpollInstance"
}

func (x *EpollInstance) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"q",
		"interest",
		"ready",
	}
}

func (x *EpollInstance) beforeSave() {}

func (x *EpollInstance) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfd)
	m.Save(1, &x.FileDescriptionDefaultImpl)
	m.Save(2, &x.DentryMetadataFileDescriptionImpl)
	m.Save(3, &x.NoLockFD)
	m.Save(4, &x.q)
	m.Save(5, &x.interest)
	m.Save(6, &x.ready)
}

func (x *EpollInstance) afterLoad() {}

func (x *EpollInstance) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfd)
	m.Load(1, &x.FileDescriptionDefaultImpl)
	m.Load(2, &x.DentryMetadataFileDescriptionImpl)
	m.Load(3, &x.NoLockFD)
	m.Load(4, &x.q)
	m.Load(5, &x.interest)
	m.Load(6, &x.ready)
}

func (x *epollInterestKey) StateTypeName() string {
	return "pkg/sentry/vfs.epollInterestKey"
}

func (x *epollInterestKey) StateFields() []string {
	return []string{
		"file",
		"num",
	}
}

func (x *epollInterestKey) beforeSave() {}

func (x *epollInterestKey) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.file)
	m.Save(1, &x.num)
}

func (x *epollInterestKey) afterLoad() {}

func (x *epollInterestKey) StateLoad(m state.Source) {
	m.Load(0, &x.file)
	m.Load(1, &x.num)
}

func (x *epollInterest) StateTypeName() string {
	return "pkg/sentry/vfs.epollInterest"
}

func (x *epollInterest) StateFields() []string {
	return []string{
		"epoll",
		"key",
		"waiter",
		"mask",
		"ready",
		"epollInterestEntry",
		"userData",
	}
}

func (x *epollInterest) beforeSave() {}

func (x *epollInterest) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.epoll)
	m.Save(1, &x.key)
	m.Save(2, &x.waiter)
	m.Save(3, &x.mask)
	m.Save(4, &x.ready)
	m.Save(5, &x.epollInterestEntry)
	m.Save(6, &x.userData)
}

func (x *epollInterest) afterLoad() {}

func (x *epollInterest) StateLoad(m state.Source) {
	m.Load(0, &x.epoll)
	m.Load(1, &x.key)
	m.Load(2, &x.waiter)
	m.Load(3, &x.mask)
	m.Load(4, &x.ready)
	m.Load(5, &x.epollInterestEntry)
	m.Load(6, &x.userData)
}

func (x *epollInterestList) StateTypeName() string {
	return "pkg/sentry/vfs.epollInterestList"
}

func (x *epollInterestList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *epollInterestList) beforeSave() {}

func (x *epollInterestList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *epollInterestList) afterLoad() {}

func (x *epollInterestList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *epollInterestEntry) StateTypeName() string {
	return "pkg/sentry/vfs.epollInterestEntry"
}

func (x *epollInterestEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *epollInterestEntry) beforeSave() {}

func (x *epollInterestEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *epollInterestEntry) afterLoad() {}

func (x *epollInterestEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func (x *eventList) StateTypeName() string {
	return "pkg/sentry/vfs.eventList"
}

func (x *eventList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *eventList) beforeSave() {}

func (x *eventList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *eventList) afterLoad() {}

func (x *eventList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *eventEntry) StateTypeName() string {
	return "pkg/sentry/vfs.eventEntry"
}

func (x *eventEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *eventEntry) beforeSave() {}

func (x *eventEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *eventEntry) afterLoad() {}

func (x *eventEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func (x *FileDescription) StateTypeName() string {
	return "pkg/sentry/vfs.FileDescription"
}

func (x *FileDescription) StateFields() []string {
	return []string{
		"FileDescriptionRefs",
		"statusFlags",
		"asyncHandler",
		"epolls",
		"vd",
		"opts",
		"readable",
		"writable",
		"usedLockBSD",
		"impl",
	}
}

func (x *FileDescription) beforeSave() {}

func (x *FileDescription) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.FileDescriptionRefs)
	m.Save(1, &x.statusFlags)
	m.Save(2, &x.asyncHandler)
	m.Save(3, &x.epolls)
	m.Save(4, &x.vd)
	m.Save(5, &x.opts)
	m.Save(6, &x.readable)
	m.Save(7, &x.writable)
	m.Save(8, &x.usedLockBSD)
	m.Save(9, &x.impl)
}

func (x *FileDescription) afterLoad() {}

func (x *FileDescription) StateLoad(m state.Source) {
	m.Load(0, &x.FileDescriptionRefs)
	m.Load(1, &x.statusFlags)
	m.Load(2, &x.asyncHandler)
	m.Load(3, &x.epolls)
	m.Load(4, &x.vd)
	m.Load(5, &x.opts)
	m.Load(6, &x.readable)
	m.Load(7, &x.writable)
	m.Load(8, &x.usedLockBSD)
	m.Load(9, &x.impl)
}

func (x *FileDescriptionOptions) StateTypeName() string {
	return "pkg/sentry/vfs.FileDescriptionOptions"
}

func (x *FileDescriptionOptions) StateFields() []string {
	return []string{
		"AllowDirectIO",
		"DenyPRead",
		"DenyPWrite",
		"UseDentryMetadata",
	}
}

func (x *FileDescriptionOptions) beforeSave() {}

func (x *FileDescriptionOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.AllowDirectIO)
	m.Save(1, &x.DenyPRead)
	m.Save(2, &x.DenyPWrite)
	m.Save(3, &x.UseDentryMetadata)
}

func (x *FileDescriptionOptions) afterLoad() {}

func (x *FileDescriptionOptions) StateLoad(m state.Source) {
	m.Load(0, &x.AllowDirectIO)
	m.Load(1, &x.DenyPRead)
	m.Load(2, &x.DenyPWrite)
	m.Load(3, &x.UseDentryMetadata)
}

func (x *Dirent) StateTypeName() string {
	return "pkg/sentry/vfs.Dirent"
}

func (x *Dirent) StateFields() []string {
	return []string{
		"Name",
		"Type",
		"Ino",
		"NextOff",
	}
}

func (x *Dirent) beforeSave() {}

func (x *Dirent) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Name)
	m.Save(1, &x.Type)
	m.Save(2, &x.Ino)
	m.Save(3, &x.NextOff)
}

func (x *Dirent) afterLoad() {}

func (x *Dirent) StateLoad(m state.Source) {
	m.Load(0, &x.Name)
	m.Load(1, &x.Type)
	m.Load(2, &x.Ino)
	m.Load(3, &x.NextOff)
}

func (x *FileDescriptionDefaultImpl) StateTypeName() string {
	return "pkg/sentry/vfs.FileDescriptionDefaultImpl"
}

func (x *FileDescriptionDefaultImpl) StateFields() []string {
	return []string{}
}

func (x *FileDescriptionDefaultImpl) beforeSave() {}

func (x *FileDescriptionDefaultImpl) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *FileDescriptionDefaultImpl) afterLoad() {}

func (x *FileDescriptionDefaultImpl) StateLoad(m state.Source) {
}

func (x *DirectoryFileDescriptionDefaultImpl) StateTypeName() string {
	return "pkg/sentry/vfs.DirectoryFileDescriptionDefaultImpl"
}

func (x *DirectoryFileDescriptionDefaultImpl) StateFields() []string {
	return []string{}
}

func (x *DirectoryFileDescriptionDefaultImpl) beforeSave() {}

func (x *DirectoryFileDescriptionDefaultImpl) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *DirectoryFileDescriptionDefaultImpl) afterLoad() {}

func (x *DirectoryFileDescriptionDefaultImpl) StateLoad(m state.Source) {
}

func (x *DentryMetadataFileDescriptionImpl) StateTypeName() string {
	return "pkg/sentry/vfs.DentryMetadataFileDescriptionImpl"
}

func (x *DentryMetadataFileDescriptionImpl) StateFields() []string {
	return []string{}
}

func (x *DentryMetadataFileDescriptionImpl) beforeSave() {}

func (x *DentryMetadataFileDescriptionImpl) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *DentryMetadataFileDescriptionImpl) afterLoad() {}

func (x *DentryMetadataFileDescriptionImpl) StateLoad(m state.Source) {
}

func (x *StaticData) StateTypeName() string {
	return "pkg/sentry/vfs.StaticData"
}

func (x *StaticData) StateFields() []string {
	return []string{
		"Data",
	}
}

func (x *StaticData) beforeSave() {}

func (x *StaticData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Data)
}

func (x *StaticData) afterLoad() {}

func (x *StaticData) StateLoad(m state.Source) {
	m.Load(0, &x.Data)
}

func (x *DynamicBytesFileDescriptionImpl) StateTypeName() string {
	return "pkg/sentry/vfs.DynamicBytesFileDescriptionImpl"
}

func (x *DynamicBytesFileDescriptionImpl) StateFields() []string {
	return []string{
		"data",
		"buf",
		"off",
		"lastRead",
	}
}

func (x *DynamicBytesFileDescriptionImpl) beforeSave() {}

func (x *DynamicBytesFileDescriptionImpl) StateSave(m state.Sink) {
	x.beforeSave()
	var buf []byte = x.saveBuf()
	m.SaveValue(1, buf)
	m.Save(0, &x.data)
	m.Save(2, &x.off)
	m.Save(3, &x.lastRead)
}

func (x *DynamicBytesFileDescriptionImpl) afterLoad() {}

func (x *DynamicBytesFileDescriptionImpl) StateLoad(m state.Source) {
	m.Load(0, &x.data)
	m.Load(2, &x.off)
	m.Load(3, &x.lastRead)
	m.LoadValue(1, new([]byte), func(y interface{}) { x.loadBuf(y.([]byte)) })
}

func (x *LockFD) StateTypeName() string {
	return "pkg/sentry/vfs.LockFD"
}

func (x *LockFD) StateFields() []string {
	return []string{
		"locks",
	}
}

func (x *LockFD) beforeSave() {}

func (x *LockFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.locks)
}

func (x *LockFD) afterLoad() {}

func (x *LockFD) StateLoad(m state.Source) {
	m.Load(0, &x.locks)
}

func (x *NoLockFD) StateTypeName() string {
	return "pkg/sentry/vfs.NoLockFD"
}

func (x *NoLockFD) StateFields() []string {
	return []string{}
}

func (x *NoLockFD) beforeSave() {}

func (x *NoLockFD) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *NoLockFD) afterLoad() {}

func (x *NoLockFD) StateLoad(m state.Source) {
}

func (x *FileDescriptionRefs) StateTypeName() string {
	return "pkg/sentry/vfs.FileDescriptionRefs"
}

func (x *FileDescriptionRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (x *FileDescriptionRefs) beforeSave() {}

func (x *FileDescriptionRefs) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
}

func (x *FileDescriptionRefs) afterLoad() {}

func (x *FileDescriptionRefs) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
}

func (x *Filesystem) StateTypeName() string {
	return "pkg/sentry/vfs.Filesystem"
}

func (x *Filesystem) StateFields() []string {
	return []string{
		"FilesystemRefs",
		"vfs",
		"fsType",
		"impl",
	}
}

func (x *Filesystem) beforeSave() {}

func (x *Filesystem) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.FilesystemRefs)
	m.Save(1, &x.vfs)
	m.Save(2, &x.fsType)
	m.Save(3, &x.impl)
}

func (x *Filesystem) afterLoad() {}

func (x *Filesystem) StateLoad(m state.Source) {
	m.Load(0, &x.FilesystemRefs)
	m.Load(1, &x.vfs)
	m.Load(2, &x.fsType)
	m.Load(3, &x.impl)
}

func (x *PrependPathAtVFSRootError) StateTypeName() string {
	return "pkg/sentry/vfs.PrependPathAtVFSRootError"
}

func (x *PrependPathAtVFSRootError) StateFields() []string {
	return []string{}
}

func (x *PrependPathAtVFSRootError) beforeSave() {}

func (x *PrependPathAtVFSRootError) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *PrependPathAtVFSRootError) afterLoad() {}

func (x *PrependPathAtVFSRootError) StateLoad(m state.Source) {
}

func (x *PrependPathAtNonMountRootError) StateTypeName() string {
	return "pkg/sentry/vfs.PrependPathAtNonMountRootError"
}

func (x *PrependPathAtNonMountRootError) StateFields() []string {
	return []string{}
}

func (x *PrependPathAtNonMountRootError) beforeSave() {}

func (x *PrependPathAtNonMountRootError) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *PrependPathAtNonMountRootError) afterLoad() {}

func (x *PrependPathAtNonMountRootError) StateLoad(m state.Source) {
}

func (x *PrependPathSyntheticError) StateTypeName() string {
	return "pkg/sentry/vfs.PrependPathSyntheticError"
}

func (x *PrependPathSyntheticError) StateFields() []string {
	return []string{}
}

func (x *PrependPathSyntheticError) beforeSave() {}

func (x *PrependPathSyntheticError) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *PrependPathSyntheticError) afterLoad() {}

func (x *PrependPathSyntheticError) StateLoad(m state.Source) {
}

func (x *FilesystemRefs) StateTypeName() string {
	return "pkg/sentry/vfs.FilesystemRefs"
}

func (x *FilesystemRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (x *FilesystemRefs) beforeSave() {}

func (x *FilesystemRefs) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
}

func (x *FilesystemRefs) afterLoad() {}

func (x *FilesystemRefs) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
}

func (x *registeredFilesystemType) StateTypeName() string {
	return "pkg/sentry/vfs.registeredFilesystemType"
}

func (x *registeredFilesystemType) StateFields() []string {
	return []string{
		"fsType",
		"opts",
	}
}

func (x *registeredFilesystemType) beforeSave() {}

func (x *registeredFilesystemType) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fsType)
	m.Save(1, &x.opts)
}

func (x *registeredFilesystemType) afterLoad() {}

func (x *registeredFilesystemType) StateLoad(m state.Source) {
	m.Load(0, &x.fsType)
	m.Load(1, &x.opts)
}

func (x *RegisterFilesystemTypeOptions) StateTypeName() string {
	return "pkg/sentry/vfs.RegisterFilesystemTypeOptions"
}

func (x *RegisterFilesystemTypeOptions) StateFields() []string {
	return []string{
		"AllowUserMount",
		"AllowUserList",
		"RequiresDevice",
	}
}

func (x *RegisterFilesystemTypeOptions) beforeSave() {}

func (x *RegisterFilesystemTypeOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.AllowUserMount)
	m.Save(1, &x.AllowUserList)
	m.Save(2, &x.RequiresDevice)
}

func (x *RegisterFilesystemTypeOptions) afterLoad() {}

func (x *RegisterFilesystemTypeOptions) StateLoad(m state.Source) {
	m.Load(0, &x.AllowUserMount)
	m.Load(1, &x.AllowUserList)
	m.Load(2, &x.RequiresDevice)
}

func (x *EventType) StateTypeName() string {
	return "pkg/sentry/vfs.EventType"
}

func (x *EventType) StateFields() []string {
	return nil
}

func (x *Inotify) StateTypeName() string {
	return "pkg/sentry/vfs.Inotify"
}

func (x *Inotify) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"id",
		"events",
		"scratch",
		"nextWatchMinusOne",
		"watches",
	}
}

func (x *Inotify) beforeSave() {}

func (x *Inotify) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfd)
	m.Save(1, &x.FileDescriptionDefaultImpl)
	m.Save(2, &x.DentryMetadataFileDescriptionImpl)
	m.Save(3, &x.NoLockFD)
	m.Save(4, &x.id)
	m.Save(5, &x.events)
	m.Save(6, &x.scratch)
	m.Save(7, &x.nextWatchMinusOne)
	m.Save(8, &x.watches)
}

func (x *Inotify) afterLoad() {}

func (x *Inotify) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfd)
	m.Load(1, &x.FileDescriptionDefaultImpl)
	m.Load(2, &x.DentryMetadataFileDescriptionImpl)
	m.Load(3, &x.NoLockFD)
	m.Load(4, &x.id)
	m.Load(5, &x.events)
	m.Load(6, &x.scratch)
	m.Load(7, &x.nextWatchMinusOne)
	m.Load(8, &x.watches)
}

func (x *Watches) StateTypeName() string {
	return "pkg/sentry/vfs.Watches"
}

func (x *Watches) StateFields() []string {
	return []string{
		"ws",
	}
}

func (x *Watches) beforeSave() {}

func (x *Watches) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.ws)
}

func (x *Watches) afterLoad() {}

func (x *Watches) StateLoad(m state.Source) {
	m.Load(0, &x.ws)
}

func (x *Watch) StateTypeName() string {
	return "pkg/sentry/vfs.Watch"
}

func (x *Watch) StateFields() []string {
	return []string{
		"owner",
		"wd",
		"target",
		"mask",
		"expired",
	}
}

func (x *Watch) beforeSave() {}

func (x *Watch) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.owner)
	m.Save(1, &x.wd)
	m.Save(2, &x.target)
	m.Save(3, &x.mask)
	m.Save(4, &x.expired)
}

func (x *Watch) afterLoad() {}

func (x *Watch) StateLoad(m state.Source) {
	m.Load(0, &x.owner)
	m.Load(1, &x.wd)
	m.Load(2, &x.target)
	m.Load(3, &x.mask)
	m.Load(4, &x.expired)
}

func (x *Event) StateTypeName() string {
	return "pkg/sentry/vfs.Event"
}

func (x *Event) StateFields() []string {
	return []string{
		"eventEntry",
		"wd",
		"mask",
		"cookie",
		"len",
		"name",
	}
}

func (x *Event) beforeSave() {}

func (x *Event) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.eventEntry)
	m.Save(1, &x.wd)
	m.Save(2, &x.mask)
	m.Save(3, &x.cookie)
	m.Save(4, &x.len)
	m.Save(5, &x.name)
}

func (x *Event) afterLoad() {}

func (x *Event) StateLoad(m state.Source) {
	m.Load(0, &x.eventEntry)
	m.Load(1, &x.wd)
	m.Load(2, &x.mask)
	m.Load(3, &x.cookie)
	m.Load(4, &x.len)
	m.Load(5, &x.name)
}

func (x *FileLocks) StateTypeName() string {
	return "pkg/sentry/vfs.FileLocks"
}

func (x *FileLocks) StateFields() []string {
	return []string{
		"bsd",
		"posix",
	}
}

func (x *FileLocks) beforeSave() {}

func (x *FileLocks) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.bsd)
	m.Save(1, &x.posix)
}

func (x *FileLocks) afterLoad() {}

func (x *FileLocks) StateLoad(m state.Source) {
	m.Load(0, &x.bsd)
	m.Load(1, &x.posix)
}

func (x *Mount) StateTypeName() string {
	return "pkg/sentry/vfs.Mount"
}

func (x *Mount) StateFields() []string {
	return []string{
		"vfs",
		"fs",
		"root",
		"ID",
		"Flags",
		"key",
		"ns",
		"refs",
		"children",
		"umounted",
		"writers",
	}
}

func (x *Mount) beforeSave() {}

func (x *Mount) StateSave(m state.Sink) {
	x.beforeSave()
	var key VirtualDentry = x.saveKey()
	m.SaveValue(5, key)
	m.Save(0, &x.vfs)
	m.Save(1, &x.fs)
	m.Save(2, &x.root)
	m.Save(3, &x.ID)
	m.Save(4, &x.Flags)
	m.Save(6, &x.ns)
	m.Save(7, &x.refs)
	m.Save(8, &x.children)
	m.Save(9, &x.umounted)
	m.Save(10, &x.writers)
}

func (x *Mount) afterLoad() {}

func (x *Mount) StateLoad(m state.Source) {
	m.Load(0, &x.vfs)
	m.Load(1, &x.fs)
	m.Load(2, &x.root)
	m.Load(3, &x.ID)
	m.Load(4, &x.Flags)
	m.Load(6, &x.ns)
	m.Load(7, &x.refs)
	m.Load(8, &x.children)
	m.Load(9, &x.umounted)
	m.Load(10, &x.writers)
	m.LoadValue(5, new(VirtualDentry), func(y interface{}) { x.loadKey(y.(VirtualDentry)) })
}

func (x *MountNamespace) StateTypeName() string {
	return "pkg/sentry/vfs.MountNamespace"
}

func (x *MountNamespace) StateFields() []string {
	return []string{
		"MountNamespaceRefs",
		"Owner",
		"root",
		"mountpoints",
	}
}

func (x *MountNamespace) beforeSave() {}

func (x *MountNamespace) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.MountNamespaceRefs)
	m.Save(1, &x.Owner)
	m.Save(2, &x.root)
	m.Save(3, &x.mountpoints)
}

func (x *MountNamespace) afterLoad() {}

func (x *MountNamespace) StateLoad(m state.Source) {
	m.Load(0, &x.MountNamespaceRefs)
	m.Load(1, &x.Owner)
	m.Load(2, &x.root)
	m.Load(3, &x.mountpoints)
}

func (x *umountRecursiveOptions) StateTypeName() string {
	return "pkg/sentry/vfs.umountRecursiveOptions"
}

func (x *umountRecursiveOptions) StateFields() []string {
	return []string{
		"eager",
		"disconnectHierarchy",
	}
}

func (x *umountRecursiveOptions) beforeSave() {}

func (x *umountRecursiveOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.eager)
	m.Save(1, &x.disconnectHierarchy)
}

func (x *umountRecursiveOptions) afterLoad() {}

func (x *umountRecursiveOptions) StateLoad(m state.Source) {
	m.Load(0, &x.eager)
	m.Load(1, &x.disconnectHierarchy)
}

func (x *MountNamespaceRefs) StateTypeName() string {
	return "pkg/sentry/vfs.MountNamespaceRefs"
}

func (x *MountNamespaceRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (x *MountNamespaceRefs) beforeSave() {}

func (x *MountNamespaceRefs) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
}

func (x *MountNamespaceRefs) afterLoad() {}

func (x *MountNamespaceRefs) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
}

func (x *GetDentryOptions) StateTypeName() string {
	return "pkg/sentry/vfs.GetDentryOptions"
}

func (x *GetDentryOptions) StateFields() []string {
	return []string{
		"CheckSearchable",
	}
}

func (x *GetDentryOptions) beforeSave() {}

func (x *GetDentryOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.CheckSearchable)
}

func (x *GetDentryOptions) afterLoad() {}

func (x *GetDentryOptions) StateLoad(m state.Source) {
	m.Load(0, &x.CheckSearchable)
}

func (x *MkdirOptions) StateTypeName() string {
	return "pkg/sentry/vfs.MkdirOptions"
}

func (x *MkdirOptions) StateFields() []string {
	return []string{
		"Mode",
		"ForSyntheticMountpoint",
	}
}

func (x *MkdirOptions) beforeSave() {}

func (x *MkdirOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Mode)
	m.Save(1, &x.ForSyntheticMountpoint)
}

func (x *MkdirOptions) afterLoad() {}

func (x *MkdirOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Mode)
	m.Load(1, &x.ForSyntheticMountpoint)
}

func (x *MknodOptions) StateTypeName() string {
	return "pkg/sentry/vfs.MknodOptions"
}

func (x *MknodOptions) StateFields() []string {
	return []string{
		"Mode",
		"DevMajor",
		"DevMinor",
		"Endpoint",
	}
}

func (x *MknodOptions) beforeSave() {}

func (x *MknodOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Mode)
	m.Save(1, &x.DevMajor)
	m.Save(2, &x.DevMinor)
	m.Save(3, &x.Endpoint)
}

func (x *MknodOptions) afterLoad() {}

func (x *MknodOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Mode)
	m.Load(1, &x.DevMajor)
	m.Load(2, &x.DevMinor)
	m.Load(3, &x.Endpoint)
}

func (x *MountFlags) StateTypeName() string {
	return "pkg/sentry/vfs.MountFlags"
}

func (x *MountFlags) StateFields() []string {
	return []string{
		"NoExec",
		"NoATime",
		"NoDev",
		"NoSUID",
	}
}

func (x *MountFlags) beforeSave() {}

func (x *MountFlags) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.NoExec)
	m.Save(1, &x.NoATime)
	m.Save(2, &x.NoDev)
	m.Save(3, &x.NoSUID)
}

func (x *MountFlags) afterLoad() {}

func (x *MountFlags) StateLoad(m state.Source) {
	m.Load(0, &x.NoExec)
	m.Load(1, &x.NoATime)
	m.Load(2, &x.NoDev)
	m.Load(3, &x.NoSUID)
}

func (x *MountOptions) StateTypeName() string {
	return "pkg/sentry/vfs.MountOptions"
}

func (x *MountOptions) StateFields() []string {
	return []string{
		"Flags",
		"ReadOnly",
		"GetFilesystemOptions",
		"InternalMount",
	}
}

func (x *MountOptions) beforeSave() {}

func (x *MountOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Flags)
	m.Save(1, &x.ReadOnly)
	m.Save(2, &x.GetFilesystemOptions)
	m.Save(3, &x.InternalMount)
}

func (x *MountOptions) afterLoad() {}

func (x *MountOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Flags)
	m.Load(1, &x.ReadOnly)
	m.Load(2, &x.GetFilesystemOptions)
	m.Load(3, &x.InternalMount)
}

func (x *OpenOptions) StateTypeName() string {
	return "pkg/sentry/vfs.OpenOptions"
}

func (x *OpenOptions) StateFields() []string {
	return []string{
		"Flags",
		"Mode",
		"FileExec",
	}
}

func (x *OpenOptions) beforeSave() {}

func (x *OpenOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Flags)
	m.Save(1, &x.Mode)
	m.Save(2, &x.FileExec)
}

func (x *OpenOptions) afterLoad() {}

func (x *OpenOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Flags)
	m.Load(1, &x.Mode)
	m.Load(2, &x.FileExec)
}

func (x *ReadOptions) StateTypeName() string {
	return "pkg/sentry/vfs.ReadOptions"
}

func (x *ReadOptions) StateFields() []string {
	return []string{
		"Flags",
	}
}

func (x *ReadOptions) beforeSave() {}

func (x *ReadOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Flags)
}

func (x *ReadOptions) afterLoad() {}

func (x *ReadOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Flags)
}

func (x *RenameOptions) StateTypeName() string {
	return "pkg/sentry/vfs.RenameOptions"
}

func (x *RenameOptions) StateFields() []string {
	return []string{
		"Flags",
		"MustBeDir",
	}
}

func (x *RenameOptions) beforeSave() {}

func (x *RenameOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Flags)
	m.Save(1, &x.MustBeDir)
}

func (x *RenameOptions) afterLoad() {}

func (x *RenameOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Flags)
	m.Load(1, &x.MustBeDir)
}

func (x *SetStatOptions) StateTypeName() string {
	return "pkg/sentry/vfs.SetStatOptions"
}

func (x *SetStatOptions) StateFields() []string {
	return []string{
		"Stat",
		"NeedWritePerm",
	}
}

func (x *SetStatOptions) beforeSave() {}

func (x *SetStatOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Stat)
	m.Save(1, &x.NeedWritePerm)
}

func (x *SetStatOptions) afterLoad() {}

func (x *SetStatOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Stat)
	m.Load(1, &x.NeedWritePerm)
}

func (x *BoundEndpointOptions) StateTypeName() string {
	return "pkg/sentry/vfs.BoundEndpointOptions"
}

func (x *BoundEndpointOptions) StateFields() []string {
	return []string{
		"Addr",
	}
}

func (x *BoundEndpointOptions) beforeSave() {}

func (x *BoundEndpointOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Addr)
}

func (x *BoundEndpointOptions) afterLoad() {}

func (x *BoundEndpointOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Addr)
}

func (x *GetXattrOptions) StateTypeName() string {
	return "pkg/sentry/vfs.GetXattrOptions"
}

func (x *GetXattrOptions) StateFields() []string {
	return []string{
		"Name",
		"Size",
	}
}

func (x *GetXattrOptions) beforeSave() {}

func (x *GetXattrOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Name)
	m.Save(1, &x.Size)
}

func (x *GetXattrOptions) afterLoad() {}

func (x *GetXattrOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Name)
	m.Load(1, &x.Size)
}

func (x *SetXattrOptions) StateTypeName() string {
	return "pkg/sentry/vfs.SetXattrOptions"
}

func (x *SetXattrOptions) StateFields() []string {
	return []string{
		"Name",
		"Value",
		"Flags",
	}
}

func (x *SetXattrOptions) beforeSave() {}

func (x *SetXattrOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Name)
	m.Save(1, &x.Value)
	m.Save(2, &x.Flags)
}

func (x *SetXattrOptions) afterLoad() {}

func (x *SetXattrOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Name)
	m.Load(1, &x.Value)
	m.Load(2, &x.Flags)
}

func (x *StatOptions) StateTypeName() string {
	return "pkg/sentry/vfs.StatOptions"
}

func (x *StatOptions) StateFields() []string {
	return []string{
		"Mask",
		"Sync",
	}
}

func (x *StatOptions) beforeSave() {}

func (x *StatOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Mask)
	m.Save(1, &x.Sync)
}

func (x *StatOptions) afterLoad() {}

func (x *StatOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Mask)
	m.Load(1, &x.Sync)
}

func (x *UmountOptions) StateTypeName() string {
	return "pkg/sentry/vfs.UmountOptions"
}

func (x *UmountOptions) StateFields() []string {
	return []string{
		"Flags",
	}
}

func (x *UmountOptions) beforeSave() {}

func (x *UmountOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Flags)
}

func (x *UmountOptions) afterLoad() {}

func (x *UmountOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Flags)
}

func (x *WriteOptions) StateTypeName() string {
	return "pkg/sentry/vfs.WriteOptions"
}

func (x *WriteOptions) StateFields() []string {
	return []string{
		"Flags",
	}
}

func (x *WriteOptions) beforeSave() {}

func (x *WriteOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Flags)
}

func (x *WriteOptions) afterLoad() {}

func (x *WriteOptions) StateLoad(m state.Source) {
	m.Load(0, &x.Flags)
}

func (x *AccessTypes) StateTypeName() string {
	return "pkg/sentry/vfs.AccessTypes"
}

func (x *AccessTypes) StateFields() []string {
	return nil
}

func (x *ResolvingPath) StateTypeName() string {
	return "pkg/sentry/vfs.ResolvingPath"
}

func (x *ResolvingPath) StateFields() []string {
	return []string{
		"vfs",
		"root",
		"mount",
		"start",
		"pit",
		"flags",
		"mustBeDir",
		"mustBeDirOrig",
		"symlinks",
		"symlinksOrig",
		"curPart",
		"numOrigParts",
		"creds",
		"nextMount",
		"nextStart",
		"absSymlinkTarget",
		"parts",
		"origParts",
	}
}

func (x *ResolvingPath) beforeSave() {}

func (x *ResolvingPath) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfs)
	m.Save(1, &x.root)
	m.Save(2, &x.mount)
	m.Save(3, &x.start)
	m.Save(4, &x.pit)
	m.Save(5, &x.flags)
	m.Save(6, &x.mustBeDir)
	m.Save(7, &x.mustBeDirOrig)
	m.Save(8, &x.symlinks)
	m.Save(9, &x.symlinksOrig)
	m.Save(10, &x.curPart)
	m.Save(11, &x.numOrigParts)
	m.Save(12, &x.creds)
	m.Save(13, &x.nextMount)
	m.Save(14, &x.nextStart)
	m.Save(15, &x.absSymlinkTarget)
	m.Save(16, &x.parts)
	m.Save(17, &x.origParts)
}

func (x *ResolvingPath) afterLoad() {}

func (x *ResolvingPath) StateLoad(m state.Source) {
	m.Load(0, &x.vfs)
	m.Load(1, &x.root)
	m.Load(2, &x.mount)
	m.Load(3, &x.start)
	m.Load(4, &x.pit)
	m.Load(5, &x.flags)
	m.Load(6, &x.mustBeDir)
	m.Load(7, &x.mustBeDirOrig)
	m.Load(8, &x.symlinks)
	m.Load(9, &x.symlinksOrig)
	m.Load(10, &x.curPart)
	m.Load(11, &x.numOrigParts)
	m.Load(12, &x.creds)
	m.Load(13, &x.nextMount)
	m.Load(14, &x.nextStart)
	m.Load(15, &x.absSymlinkTarget)
	m.Load(16, &x.parts)
	m.Load(17, &x.origParts)
}

func (x *resolveMountRootOrJumpError) StateTypeName() string {
	return "pkg/sentry/vfs.resolveMountRootOrJumpError"
}

func (x *resolveMountRootOrJumpError) StateFields() []string {
	return []string{}
}

func (x *resolveMountRootOrJumpError) beforeSave() {}

func (x *resolveMountRootOrJumpError) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *resolveMountRootOrJumpError) afterLoad() {}

func (x *resolveMountRootOrJumpError) StateLoad(m state.Source) {
}

func (x *resolveMountPointError) StateTypeName() string {
	return "pkg/sentry/vfs.resolveMountPointError"
}

func (x *resolveMountPointError) StateFields() []string {
	return []string{}
}

func (x *resolveMountPointError) beforeSave() {}

func (x *resolveMountPointError) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *resolveMountPointError) afterLoad() {}

func (x *resolveMountPointError) StateLoad(m state.Source) {
}

func (x *resolveAbsSymlinkError) StateTypeName() string {
	return "pkg/sentry/vfs.resolveAbsSymlinkError"
}

func (x *resolveAbsSymlinkError) StateFields() []string {
	return []string{}
}

func (x *resolveAbsSymlinkError) beforeSave() {}

func (x *resolveAbsSymlinkError) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *resolveAbsSymlinkError) afterLoad() {}

func (x *resolveAbsSymlinkError) StateLoad(m state.Source) {
}

func (x *VirtualFilesystem) StateTypeName() string {
	return "pkg/sentry/vfs.VirtualFilesystem"
}

func (x *VirtualFilesystem) StateFields() []string {
	return []string{
		"mounts",
		"mountpoints",
		"lastMountID",
		"anonMount",
		"devices",
		"anonBlockDevMinorNext",
		"anonBlockDevMinor",
		"fsTypes",
		"filesystems",
	}
}

func (x *VirtualFilesystem) beforeSave() {}

func (x *VirtualFilesystem) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.mounts)
	m.Save(1, &x.mountpoints)
	m.Save(2, &x.lastMountID)
	m.Save(3, &x.anonMount)
	m.Save(4, &x.devices)
	m.Save(5, &x.anonBlockDevMinorNext)
	m.Save(6, &x.anonBlockDevMinor)
	m.Save(7, &x.fsTypes)
	m.Save(8, &x.filesystems)
}

func (x *VirtualFilesystem) afterLoad() {}

func (x *VirtualFilesystem) StateLoad(m state.Source) {
	m.Load(0, &x.mounts)
	m.Load(1, &x.mountpoints)
	m.Load(2, &x.lastMountID)
	m.Load(3, &x.anonMount)
	m.Load(4, &x.devices)
	m.Load(5, &x.anonBlockDevMinorNext)
	m.Load(6, &x.anonBlockDevMinor)
	m.Load(7, &x.fsTypes)
	m.Load(8, &x.filesystems)
}

func (x *PathOperation) StateTypeName() string {
	return "pkg/sentry/vfs.PathOperation"
}

func (x *PathOperation) StateFields() []string {
	return []string{
		"Root",
		"Start",
		"Path",
		"FollowFinalSymlink",
	}
}

func (x *PathOperation) beforeSave() {}

func (x *PathOperation) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Root)
	m.Save(1, &x.Start)
	m.Save(2, &x.Path)
	m.Save(3, &x.FollowFinalSymlink)
}

func (x *PathOperation) afterLoad() {}

func (x *PathOperation) StateLoad(m state.Source) {
	m.Load(0, &x.Root)
	m.Load(1, &x.Start)
	m.Load(2, &x.Path)
	m.Load(3, &x.FollowFinalSymlink)
}

func (x *VirtualDentry) StateTypeName() string {
	return "pkg/sentry/vfs.VirtualDentry"
}

func (x *VirtualDentry) StateFields() []string {
	return []string{
		"mount",
		"dentry",
	}
}

func (x *VirtualDentry) beforeSave() {}

func (x *VirtualDentry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.mount)
	m.Save(1, &x.dentry)
}

func (x *VirtualDentry) afterLoad() {}

func (x *VirtualDentry) StateLoad(m state.Source) {
	m.Load(0, &x.mount)
	m.Load(1, &x.dentry)
}

func init() {
	state.Register((*anonFilesystemType)(nil))
	state.Register((*anonFilesystem)(nil))
	state.Register((*anonDentry)(nil))
	state.Register((*Dentry)(nil))
	state.Register((*DeviceKind)(nil))
	state.Register((*devTuple)(nil))
	state.Register((*registeredDevice)(nil))
	state.Register((*RegisterDeviceOptions)(nil))
	state.Register((*EpollInstance)(nil))
	state.Register((*epollInterestKey)(nil))
	state.Register((*epollInterest)(nil))
	state.Register((*epollInterestList)(nil))
	state.Register((*epollInterestEntry)(nil))
	state.Register((*eventList)(nil))
	state.Register((*eventEntry)(nil))
	state.Register((*FileDescription)(nil))
	state.Register((*FileDescriptionOptions)(nil))
	state.Register((*Dirent)(nil))
	state.Register((*FileDescriptionDefaultImpl)(nil))
	state.Register((*DirectoryFileDescriptionDefaultImpl)(nil))
	state.Register((*DentryMetadataFileDescriptionImpl)(nil))
	state.Register((*StaticData)(nil))
	state.Register((*DynamicBytesFileDescriptionImpl)(nil))
	state.Register((*LockFD)(nil))
	state.Register((*NoLockFD)(nil))
	state.Register((*FileDescriptionRefs)(nil))
	state.Register((*Filesystem)(nil))
	state.Register((*PrependPathAtVFSRootError)(nil))
	state.Register((*PrependPathAtNonMountRootError)(nil))
	state.Register((*PrependPathSyntheticError)(nil))
	state.Register((*FilesystemRefs)(nil))
	state.Register((*registeredFilesystemType)(nil))
	state.Register((*RegisterFilesystemTypeOptions)(nil))
	state.Register((*EventType)(nil))
	state.Register((*Inotify)(nil))
	state.Register((*Watches)(nil))
	state.Register((*Watch)(nil))
	state.Register((*Event)(nil))
	state.Register((*FileLocks)(nil))
	state.Register((*Mount)(nil))
	state.Register((*MountNamespace)(nil))
	state.Register((*umountRecursiveOptions)(nil))
	state.Register((*MountNamespaceRefs)(nil))
	state.Register((*GetDentryOptions)(nil))
	state.Register((*MkdirOptions)(nil))
	state.Register((*MknodOptions)(nil))
	state.Register((*MountFlags)(nil))
	state.Register((*MountOptions)(nil))
	state.Register((*OpenOptions)(nil))
	state.Register((*ReadOptions)(nil))
	state.Register((*RenameOptions)(nil))
	state.Register((*SetStatOptions)(nil))
	state.Register((*BoundEndpointOptions)(nil))
	state.Register((*GetXattrOptions)(nil))
	state.Register((*SetXattrOptions)(nil))
	state.Register((*StatOptions)(nil))
	state.Register((*UmountOptions)(nil))
	state.Register((*WriteOptions)(nil))
	state.Register((*AccessTypes)(nil))
	state.Register((*ResolvingPath)(nil))
	state.Register((*resolveMountRootOrJumpError)(nil))
	state.Register((*resolveMountPointError)(nil))
	state.Register((*resolveAbsSymlinkError)(nil))
	state.Register((*VirtualFilesystem)(nil))
	state.Register((*PathOperation)(nil))
	state.Register((*VirtualDentry)(nil))
}