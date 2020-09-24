// automatically generated by stateify.

package tmpfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *dentryList) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.dentryList"
}

func (x *dentryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *dentryList) beforeSave() {}

func (x *dentryList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *dentryList) afterLoad() {}

func (x *dentryList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *dentryEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.dentryEntry"
}

func (x *dentryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *dentryEntry) beforeSave() {}

func (x *dentryEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *dentryEntry) afterLoad() {}

func (x *dentryEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func (x *deviceFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.deviceFile"
}

func (x *deviceFile) StateFields() []string {
	return []string{
		"inode",
		"kind",
		"major",
		"minor",
	}
}

func (x *deviceFile) beforeSave() {}

func (x *deviceFile) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.inode)
	m.Save(1, &x.kind)
	m.Save(2, &x.major)
	m.Save(3, &x.minor)
}

func (x *deviceFile) afterLoad() {}

func (x *deviceFile) StateLoad(m state.Source) {
	m.Load(0, &x.inode)
	m.Load(1, &x.kind)
	m.Load(2, &x.major)
	m.Load(3, &x.minor)
}

func (x *directory) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.directory"
}

func (x *directory) StateFields() []string {
	return []string{
		"dentry",
		"inode",
		"childMap",
		"numChildren",
		"childList",
	}
}

func (x *directory) beforeSave() {}

func (x *directory) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dentry)
	m.Save(1, &x.inode)
	m.Save(2, &x.childMap)
	m.Save(3, &x.numChildren)
	m.Save(4, &x.childList)
}

func (x *directory) afterLoad() {}

func (x *directory) StateLoad(m state.Source) {
	m.Load(0, &x.dentry)
	m.Load(1, &x.inode)
	m.Load(2, &x.childMap)
	m.Load(3, &x.numChildren)
	m.Load(4, &x.childList)
}

func (x *directoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.directoryFD"
}

func (x *directoryFD) StateFields() []string {
	return []string{
		"fileDescription",
		"DirectoryFileDescriptionDefaultImpl",
		"iter",
		"off",
	}
}

func (x *directoryFD) beforeSave() {}

func (x *directoryFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileDescription)
	m.Save(1, &x.DirectoryFileDescriptionDefaultImpl)
	m.Save(2, &x.iter)
	m.Save(3, &x.off)
}

func (x *directoryFD) afterLoad() {}

func (x *directoryFD) StateLoad(m state.Source) {
	m.Load(0, &x.fileDescription)
	m.Load(1, &x.DirectoryFileDescriptionDefaultImpl)
	m.Load(2, &x.iter)
	m.Load(3, &x.off)
}

func (x *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.inodeRefs"
}

func (x *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (x *inodeRefs) beforeSave() {}

func (x *inodeRefs) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
}

func (x *inodeRefs) afterLoad() {}

func (x *inodeRefs) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
}

func (x *namedPipe) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.namedPipe"
}

func (x *namedPipe) StateFields() []string {
	return []string{
		"inode",
		"pipe",
	}
}

func (x *namedPipe) beforeSave() {}

func (x *namedPipe) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.inode)
	m.Save(1, &x.pipe)
}

func (x *namedPipe) afterLoad() {}

func (x *namedPipe) StateLoad(m state.Source) {
	m.Load(0, &x.inode)
	m.Load(1, &x.pipe)
}

func (x *regularFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.regularFile"
}

func (x *regularFile) StateFields() []string {
	return []string{
		"inode",
		"memFile",
		"memoryUsageKind",
		"mappings",
		"writableMappingPages",
		"data",
		"seals",
		"size",
	}
}

func (x *regularFile) beforeSave() {}

func (x *regularFile) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.inode)
	m.Save(1, &x.memFile)
	m.Save(2, &x.memoryUsageKind)
	m.Save(3, &x.mappings)
	m.Save(4, &x.writableMappingPages)
	m.Save(5, &x.data)
	m.Save(6, &x.seals)
	m.Save(7, &x.size)
}

func (x *regularFile) afterLoad() {}

func (x *regularFile) StateLoad(m state.Source) {
	m.Load(0, &x.inode)
	m.Load(1, &x.memFile)
	m.Load(2, &x.memoryUsageKind)
	m.Load(3, &x.mappings)
	m.Load(4, &x.writableMappingPages)
	m.Load(5, &x.data)
	m.Load(6, &x.seals)
	m.Load(7, &x.size)
}

func (x *regularFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.regularFileFD"
}

func (x *regularFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"off",
	}
}

func (x *regularFileFD) beforeSave() {}

func (x *regularFileFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileDescription)
	m.Save(1, &x.off)
}

func (x *regularFileFD) afterLoad() {}

func (x *regularFileFD) StateLoad(m state.Source) {
	m.Load(0, &x.fileDescription)
	m.Load(1, &x.off)
}

func (x *socketFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.socketFile"
}

func (x *socketFile) StateFields() []string {
	return []string{
		"inode",
		"ep",
	}
}

func (x *socketFile) beforeSave() {}

func (x *socketFile) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.inode)
	m.Save(1, &x.ep)
}

func (x *socketFile) afterLoad() {}

func (x *socketFile) StateLoad(m state.Source) {
	m.Load(0, &x.inode)
	m.Load(1, &x.ep)
}

func (x *symlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.symlink"
}

func (x *symlink) StateFields() []string {
	return []string{
		"inode",
		"target",
	}
}

func (x *symlink) beforeSave() {}

func (x *symlink) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.inode)
	m.Save(1, &x.target)
}

func (x *symlink) afterLoad() {}

func (x *symlink) StateLoad(m state.Source) {
	m.Load(0, &x.inode)
	m.Load(1, &x.target)
}

func (x *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.FilesystemType"
}

func (x *FilesystemType) StateFields() []string {
	return []string{}
}

func (x *FilesystemType) beforeSave() {}

func (x *FilesystemType) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *FilesystemType) afterLoad() {}

func (x *FilesystemType) StateLoad(m state.Source) {
}

func (x *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.filesystem"
}

func (x *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"memFile",
		"clock",
		"devMinor",
		"nextInoMinusOne",
	}
}

func (x *filesystem) beforeSave() {}

func (x *filesystem) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfs)
	m.Save(1, &x.memFile)
	m.Save(2, &x.clock)
	m.Save(3, &x.devMinor)
	m.Save(4, &x.nextInoMinusOne)
}

func (x *filesystem) afterLoad() {}

func (x *filesystem) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfs)
	m.Load(1, &x.memFile)
	m.Load(2, &x.clock)
	m.Load(3, &x.devMinor)
	m.Load(4, &x.nextInoMinusOne)
}

func (x *FilesystemOpts) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.FilesystemOpts"
}

func (x *FilesystemOpts) StateFields() []string {
	return []string{
		"RootFileType",
		"RootSymlinkTarget",
		"FilesystemType",
	}
}

func (x *FilesystemOpts) beforeSave() {}

func (x *FilesystemOpts) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.RootFileType)
	m.Save(1, &x.RootSymlinkTarget)
	m.Save(2, &x.FilesystemType)
}

func (x *FilesystemOpts) afterLoad() {}

func (x *FilesystemOpts) StateLoad(m state.Source) {
	m.Load(0, &x.RootFileType)
	m.Load(1, &x.RootSymlinkTarget)
	m.Load(2, &x.FilesystemType)
}

func (x *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.dentry"
}

func (x *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"parent",
		"name",
		"dentryEntry",
		"inode",
	}
}

func (x *dentry) beforeSave() {}

func (x *dentry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsd)
	m.Save(1, &x.parent)
	m.Save(2, &x.name)
	m.Save(3, &x.dentryEntry)
	m.Save(4, &x.inode)
}

func (x *dentry) afterLoad() {}

func (x *dentry) StateLoad(m state.Source) {
	m.Load(0, &x.vfsd)
	m.Load(1, &x.parent)
	m.Load(2, &x.name)
	m.Load(3, &x.dentryEntry)
	m.Load(4, &x.inode)
}

func (x *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.inode"
}

func (x *inode) StateFields() []string {
	return []string{
		"fs",
		"refs",
		"xattrs",
		"mode",
		"nlink",
		"uid",
		"gid",
		"ino",
		"atime",
		"ctime",
		"mtime",
		"locks",
		"watches",
		"impl",
	}
}

func (x *inode) beforeSave() {}

func (x *inode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fs)
	m.Save(1, &x.refs)
	m.Save(2, &x.xattrs)
	m.Save(3, &x.mode)
	m.Save(4, &x.nlink)
	m.Save(5, &x.uid)
	m.Save(6, &x.gid)
	m.Save(7, &x.ino)
	m.Save(8, &x.atime)
	m.Save(9, &x.ctime)
	m.Save(10, &x.mtime)
	m.Save(11, &x.locks)
	m.Save(12, &x.watches)
	m.Save(13, &x.impl)
}

func (x *inode) afterLoad() {}

func (x *inode) StateLoad(m state.Source) {
	m.Load(0, &x.fs)
	m.Load(1, &x.refs)
	m.Load(2, &x.xattrs)
	m.Load(3, &x.mode)
	m.Load(4, &x.nlink)
	m.Load(5, &x.uid)
	m.Load(6, &x.gid)
	m.Load(7, &x.ino)
	m.Load(8, &x.atime)
	m.Load(9, &x.ctime)
	m.Load(10, &x.mtime)
	m.Load(11, &x.locks)
	m.Load(12, &x.watches)
	m.Load(13, &x.impl)
}

func (x *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/tmpfs.fileDescription"
}

func (x *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
	}
}

func (x *fileDescription) beforeSave() {}

func (x *fileDescription) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfd)
	m.Save(1, &x.FileDescriptionDefaultImpl)
	m.Save(2, &x.LockFD)
}

func (x *fileDescription) afterLoad() {}

func (x *fileDescription) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfd)
	m.Load(1, &x.FileDescriptionDefaultImpl)
	m.Load(2, &x.LockFD)
}

func init() {
	state.Register((*dentryList)(nil))
	state.Register((*dentryEntry)(nil))
	state.Register((*deviceFile)(nil))
	state.Register((*directory)(nil))
	state.Register((*directoryFD)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*namedPipe)(nil))
	state.Register((*regularFile)(nil))
	state.Register((*regularFileFD)(nil))
	state.Register((*socketFile)(nil))
	state.Register((*symlink)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*FilesystemOpts)(nil))
	state.Register((*dentry)(nil))
	state.Register((*inode)(nil))
	state.Register((*fileDescription)(nil))
}