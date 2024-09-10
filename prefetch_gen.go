
package prefetch

// Autogenerated code from profile_vtypes.json. Do not edit.

import (
    "encoding/binary"
    "fmt"
    "bytes"
    "io"
    "sort"
    "strings"
    "unicode/utf16"
    "unicode/utf8"
)

var (
   // Depending on autogenerated code we may use this. Add a reference
   // to shut the compiler up.
   _ = bytes.MinRead
   _ = fmt.Sprintf
   _ = utf16.Decode
   _ = binary.LittleEndian
   _ = utf8.RuneError
   _ = sort.Strings
   _ = strings.Join
   _ = io.Copy
)

func indent(text string) string {
    result := []string{}
    lines := strings.Split(text,"\n")
    for _, line := range lines {
         result = append(result, "  " + line)
    }
    return strings.Join(result, "\n")
}


type PrefetchProfile struct {
    Off_FileInformationVista_FileMetricsOffset int64
    Off_FileInformationVista_NumberOfFileMetrics int64
    Off_FileInformationVista_TraceChainsArrayOffset int64
    Off_FileInformationVista_NumberOfTraceChains int64
    Off_FileInformationVista_FilenameOffset int64
    Off_FileInformationVista_FilenameSize int64
    Off_FileInformationVista_VolumesInformationOffset int64
    Off_FileInformationVista_NumberOfVolumes int64
    Off_FileInformationVista_VolumesInformationSize int64
    Off_FileInformationVista_LastRunTime int64
    Off_FileInformationVista_RunCount int64
    Off_FileInformationWin10_FileMetricsOffset int64
    Off_FileInformationWin10_NumberOfFileMetrics int64
    Off_FileInformationWin10_TraceChainsArrayOffset int64
    Off_FileInformationWin10_NumberOfTraceChains int64
    Off_FileInformationWin10_FilenameOffset int64
    Off_FileInformationWin10_FilenameSize int64
    Off_FileInformationWin10_VolumesInformationOffset int64
    Off_FileInformationWin10_NumberOfVolumes int64
    Off_FileInformationWin10_VolumesInformationSize int64
    Off_FileInformationWin10_LastRunTimes int64
    Off_FileInformationWin10_RunCount1 int64
    Off_FileInformationWin10_RunCount2 int64
    Off_FileInformationXP_FileMetricsOffset int64
    Off_FileInformationXP_NumberOfFileMetrics int64
    Off_FileInformationXP_TraceChainsArrayOffset int64
    Off_FileInformationXP_NumberOfTraceChains int64
    Off_FileInformationXP_FilenameOffset int64
    Off_FileInformationXP_FilenameSize int64
    Off_FileInformationXP_VolumesInformationOffset int64
    Off_FileInformationXP_NumberOfVolumes int64
    Off_FileInformationXP_VolumesInformationSize int64
    Off_FileInformationXP_LastRunTime int64
    Off_FileInformationXP_RunCount int64
    Off_FileMetricsEntryV17_FilenameOffset int64
    Off_FileMetricsEntryV17_FilenameLength int64
    Off_FileMetricsEntryV30_FilenameOffset int64
    Off_FileMetricsEntryV30_FilenameLength int64
    Off_FileMetricsEntryV30_MFTFileReference int64
    Off_MAMHeader_Signature int64
    Off_MAMHeader_UncompressedSize int64
    Off_SCCAHeader_Version int64
    Off_SCCAHeader_Signature int64
    Off_SCCAHeader_FileSize int64
    Off_SCCAHeader_Executable int64
    Off_SCCAHeader_Hash int64
}

func NewPrefetchProfile() *PrefetchProfile {
    // Specific offsets can be tweaked to cater for slight version mismatches.
    self := &PrefetchProfile{0,4,8,12,16,20,24,28,32,44,68,0,4,8,12,16,20,24,28,32,44,124,116,0,4,8,12,16,20,24,28,32,36,60,8,12,12,16,24,0,4,0,4,12,16,76}
    return self
}

func (self *PrefetchProfile) FileInformationVista(reader io.ReaderAt, offset int64) *FileInformationVista {
    return &FileInformationVista{Reader: reader, Offset: offset, Profile: self}
}

func (self *PrefetchProfile) FileInformationWin10(reader io.ReaderAt, offset int64) *FileInformationWin10 {
    return &FileInformationWin10{Reader: reader, Offset: offset, Profile: self}
}

func (self *PrefetchProfile) FileInformationXP(reader io.ReaderAt, offset int64) *FileInformationXP {
    return &FileInformationXP{Reader: reader, Offset: offset, Profile: self}
}

func (self *PrefetchProfile) FileMetricsEntryV17(reader io.ReaderAt, offset int64) *FileMetricsEntryV17 {
    return &FileMetricsEntryV17{Reader: reader, Offset: offset, Profile: self}
}

func (self *PrefetchProfile) FileMetricsEntryV30(reader io.ReaderAt, offset int64) *FileMetricsEntryV30 {
    return &FileMetricsEntryV30{Reader: reader, Offset: offset, Profile: self}
}

func (self *PrefetchProfile) MAMHeader(reader io.ReaderAt, offset int64) *MAMHeader {
    return &MAMHeader{Reader: reader, Offset: offset, Profile: self}
}

func (self *PrefetchProfile) SCCAHeader(reader io.ReaderAt, offset int64) *SCCAHeader {
    return &SCCAHeader{Reader: reader, Offset: offset, Profile: self}
}


type FileInformationVista struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *FileInformationVista) Size() int {
    return 156
}

func (self *FileInformationVista) FileMetricsOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_FileMetricsOffset + self.Offset)
}

func (self *FileInformationVista) NumberOfFileMetrics() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_NumberOfFileMetrics + self.Offset)
}

func (self *FileInformationVista) TraceChainsArrayOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_TraceChainsArrayOffset + self.Offset)
}

func (self *FileInformationVista) NumberOfTraceChains() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_NumberOfTraceChains + self.Offset)
}

func (self *FileInformationVista) FilenameOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_FilenameOffset + self.Offset)
}

func (self *FileInformationVista) FilenameSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_FilenameSize + self.Offset)
}

func (self *FileInformationVista) VolumesInformationOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_VolumesInformationOffset + self.Offset)
}

func (self *FileInformationVista) NumberOfVolumes() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_NumberOfVolumes + self.Offset)
}

func (self *FileInformationVista) VolumesInformationSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_VolumesInformationSize + self.Offset)
}

func (self *FileInformationVista) LastRunTime() *WinFileTime {
    return self.Profile.WinFileTime(self.Reader, self.Profile.Off_FileInformationVista_LastRunTime + self.Offset)
}

func (self *FileInformationVista) RunCount() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationVista_RunCount + self.Offset)
}
func (self *FileInformationVista) DebugString() string {
    result := fmt.Sprintf("struct FileInformationVista @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FileMetricsOffset: %#0x\n", self.FileMetricsOffset())
    result += fmt.Sprintf("  NumberOfFileMetrics: %#0x\n", self.NumberOfFileMetrics())
    result += fmt.Sprintf("  TraceChainsArrayOffset: %#0x\n", self.TraceChainsArrayOffset())
    result += fmt.Sprintf("  NumberOfTraceChains: %#0x\n", self.NumberOfTraceChains())
    result += fmt.Sprintf("  FilenameOffset: %#0x\n", self.FilenameOffset())
    result += fmt.Sprintf("  FilenameSize: %#0x\n", self.FilenameSize())
    result += fmt.Sprintf("  VolumesInformationOffset: %#0x\n", self.VolumesInformationOffset())
    result += fmt.Sprintf("  NumberOfVolumes: %#0x\n", self.NumberOfVolumes())
    result += fmt.Sprintf("  VolumesInformationSize: %#0x\n", self.VolumesInformationSize())
    result += fmt.Sprintf("  LastRunTime: {\n%v}\n", indent(self.LastRunTime().DebugString()))
    result += fmt.Sprintf("  RunCount: %#0x\n", self.RunCount())
    return result
}

type FileInformationWin10 struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *FileInformationWin10) Size() int {
    return 224
}

func (self *FileInformationWin10) FileMetricsOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_FileMetricsOffset + self.Offset)
}

func (self *FileInformationWin10) NumberOfFileMetrics() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_NumberOfFileMetrics + self.Offset)
}

func (self *FileInformationWin10) TraceChainsArrayOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_TraceChainsArrayOffset + self.Offset)
}

func (self *FileInformationWin10) NumberOfTraceChains() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_NumberOfTraceChains + self.Offset)
}

func (self *FileInformationWin10) FilenameOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_FilenameOffset + self.Offset)
}

func (self *FileInformationWin10) FilenameSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_FilenameSize + self.Offset)
}

func (self *FileInformationWin10) VolumesInformationOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_VolumesInformationOffset + self.Offset)
}

func (self *FileInformationWin10) NumberOfVolumes() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_NumberOfVolumes + self.Offset)
}

func (self *FileInformationWin10) VolumesInformationSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_VolumesInformationSize + self.Offset)
}

func (self *FileInformationWin10) LastRunTimes() []*WinFileTime {
   return ParseArray_WinFileTime(self.Profile, self.Reader, self.Profile.Off_FileInformationWin10_LastRunTimes + self.Offset, 8)
}

func (self *FileInformationWin10) RunCount1() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_RunCount1 + self.Offset)
}

func (self *FileInformationWin10) RunCount2() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationWin10_RunCount2 + self.Offset)
}
func (self *FileInformationWin10) DebugString() string {
    result := fmt.Sprintf("struct FileInformationWin10 @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FileMetricsOffset: %#0x\n", self.FileMetricsOffset())
    result += fmt.Sprintf("  NumberOfFileMetrics: %#0x\n", self.NumberOfFileMetrics())
    result += fmt.Sprintf("  TraceChainsArrayOffset: %#0x\n", self.TraceChainsArrayOffset())
    result += fmt.Sprintf("  NumberOfTraceChains: %#0x\n", self.NumberOfTraceChains())
    result += fmt.Sprintf("  FilenameOffset: %#0x\n", self.FilenameOffset())
    result += fmt.Sprintf("  FilenameSize: %#0x\n", self.FilenameSize())
    result += fmt.Sprintf("  VolumesInformationOffset: %#0x\n", self.VolumesInformationOffset())
    result += fmt.Sprintf("  NumberOfVolumes: %#0x\n", self.NumberOfVolumes())
    result += fmt.Sprintf("  VolumesInformationSize: %#0x\n", self.VolumesInformationSize())
    result += fmt.Sprintf("  RunCount1: %#0x\n", self.RunCount1())
    result += fmt.Sprintf("  RunCount2: %#0x\n", self.RunCount2())
    return result
}

type FileInformationXP struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *FileInformationXP) Size() int {
    return 68
}

func (self *FileInformationXP) FileMetricsOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_FileMetricsOffset + self.Offset)
}

func (self *FileInformationXP) NumberOfFileMetrics() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_NumberOfFileMetrics + self.Offset)
}

func (self *FileInformationXP) TraceChainsArrayOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_TraceChainsArrayOffset + self.Offset)
}

func (self *FileInformationXP) NumberOfTraceChains() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_NumberOfTraceChains + self.Offset)
}

func (self *FileInformationXP) FilenameOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_FilenameOffset + self.Offset)
}

func (self *FileInformationXP) FilenameSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_FilenameSize + self.Offset)
}

func (self *FileInformationXP) VolumesInformationOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_VolumesInformationOffset + self.Offset)
}

func (self *FileInformationXP) NumberOfVolumes() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_NumberOfVolumes + self.Offset)
}

func (self *FileInformationXP) VolumesInformationSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_VolumesInformationSize + self.Offset)
}

func (self *FileInformationXP) LastRunTime() *WinFileTime {
    return self.Profile.WinFileTime(self.Reader, self.Profile.Off_FileInformationXP_LastRunTime + self.Offset)
}

func (self *FileInformationXP) RunCount() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileInformationXP_RunCount + self.Offset)
}
func (self *FileInformationXP) DebugString() string {
    result := fmt.Sprintf("struct FileInformationXP @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FileMetricsOffset: %#0x\n", self.FileMetricsOffset())
    result += fmt.Sprintf("  NumberOfFileMetrics: %#0x\n", self.NumberOfFileMetrics())
    result += fmt.Sprintf("  TraceChainsArrayOffset: %#0x\n", self.TraceChainsArrayOffset())
    result += fmt.Sprintf("  NumberOfTraceChains: %#0x\n", self.NumberOfTraceChains())
    result += fmt.Sprintf("  FilenameOffset: %#0x\n", self.FilenameOffset())
    result += fmt.Sprintf("  FilenameSize: %#0x\n", self.FilenameSize())
    result += fmt.Sprintf("  VolumesInformationOffset: %#0x\n", self.VolumesInformationOffset())
    result += fmt.Sprintf("  NumberOfVolumes: %#0x\n", self.NumberOfVolumes())
    result += fmt.Sprintf("  VolumesInformationSize: %#0x\n", self.VolumesInformationSize())
    result += fmt.Sprintf("  LastRunTime: {\n%v}\n", indent(self.LastRunTime().DebugString()))
    result += fmt.Sprintf("  RunCount: %#0x\n", self.RunCount())
    return result
}

type FileMetricsEntryV17 struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *FileMetricsEntryV17) Size() int {
    return 20
}

func (self *FileMetricsEntryV17) FilenameOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileMetricsEntryV17_FilenameOffset + self.Offset)
}

func (self *FileMetricsEntryV17) FilenameLength() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileMetricsEntryV17_FilenameLength + self.Offset)
}
func (self *FileMetricsEntryV17) DebugString() string {
    result := fmt.Sprintf("struct FileMetricsEntryV17 @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FilenameOffset: %#0x\n", self.FilenameOffset())
    result += fmt.Sprintf("  FilenameLength: %#0x\n", self.FilenameLength())
    return result
}

type FileMetricsEntryV30 struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *FileMetricsEntryV30) Size() int {
    return 32
}

func (self *FileMetricsEntryV30) FilenameOffset() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileMetricsEntryV30_FilenameOffset + self.Offset)
}

func (self *FileMetricsEntryV30) FilenameLength() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileMetricsEntryV30_FilenameLength + self.Offset)
}

func (self *FileMetricsEntryV30) MFTFileReference() uint64 {
    return ParseUint64(self.Reader, self.Profile.Off_FileMetricsEntryV30_MFTFileReference + self.Offset)
}
func (self *FileMetricsEntryV30) DebugString() string {
    result := fmt.Sprintf("struct FileMetricsEntryV30 @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FilenameOffset: %#0x\n", self.FilenameOffset())
    result += fmt.Sprintf("  FilenameLength: %#0x\n", self.FilenameLength())
    result += fmt.Sprintf("  MFTFileReference: %#0x\n", self.MFTFileReference())
    return result
}

type MAMHeader struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *MAMHeader) Size() int {
    return 8
}


func (self *MAMHeader) Signature() string {
  return ParseString(self.Reader, self.Profile.Off_MAMHeader_Signature + self.Offset, 4)
}

func (self *MAMHeader) UncompressedSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_MAMHeader_UncompressedSize + self.Offset)
}
func (self *MAMHeader) DebugString() string {
    result := fmt.Sprintf("struct MAMHeader @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Signature: %v\n", string(self.Signature()))
    result += fmt.Sprintf("  UncompressedSize: %#0x\n", self.UncompressedSize())
    return result
}

type SCCAHeader struct {
    Reader io.ReaderAt
    Offset int64
    Profile *PrefetchProfile
}

func (self *SCCAHeader) Size() int {
    return 84
}

func (self *SCCAHeader) Version() *Enumeration {
   value := ParseUint32(self.Reader, self.Profile.Off_SCCAHeader_Version + self.Offset)
   name := "Unknown"
   switch value {

      case 17:
         name = "WinXP"

      case 23:
         name = "Vista"

      case 26:
         name = "Win8.1"

      case 30:
         name = "Win10"

      case 31:
         name = "Win11"
}
   return &Enumeration{Value: uint64(value), Name: name}
}



func (self *SCCAHeader) Signature() string {
  return ParseString(self.Reader, self.Profile.Off_SCCAHeader_Signature + self.Offset, 4)
}

func (self *SCCAHeader) FileSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_SCCAHeader_FileSize + self.Offset)
}


func (self *SCCAHeader) Executable() string {
  return ParseTerminatedUTF16String(self.Reader, self.Profile.Off_SCCAHeader_Executable + self.Offset)
}

func (self *SCCAHeader) Hash() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_SCCAHeader_Hash + self.Offset)
}
func (self *SCCAHeader) DebugString() string {
    result := fmt.Sprintf("struct SCCAHeader @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Version: %v\n", self.Version().DebugString())
    result += fmt.Sprintf("  Signature: %v\n", string(self.Signature()))
    result += fmt.Sprintf("  FileSize: %#0x\n", self.FileSize())
    result += fmt.Sprintf("  Executable: %v\n", string(self.Executable()))
    result += fmt.Sprintf("  Hash: %#0x\n", self.Hash())
    return result
}

type Enumeration struct {
    Value uint64
    Name  string
}

func (self Enumeration) DebugString() string {
    return fmt.Sprintf("%s (%d)", self.Name, self.Value)
}


func ParseArray_WinFileTime(profile *PrefetchProfile, reader io.ReaderAt, offset int64, count int) []*WinFileTime {
    result := make([]*WinFileTime, 0, count)
    for i:=0; i<count; i++ {
      value := profile.WinFileTime(reader, offset)
      result = append(result, value)
      offset += int64(value.Size())
    }
    return result
}

func ParseUint32(reader io.ReaderAt, offset int64) uint32 {
	var buf [4]byte
	data := buf[:]
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return binary.LittleEndian.Uint32(data)
}

func ParseUint64(reader io.ReaderAt, offset int64) uint64 {
	var buf [8]byte
	data := buf[:]
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return binary.LittleEndian.Uint64(data)
}

func ParseTerminatedString(reader io.ReaderAt, offset int64) string {
   var buf [1024]byte
   data := buf[:]
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
     return ""
   }
   idx := bytes.Index(data[:n], []byte{0})
   if idx < 0 {
      idx = n
   }
   return string(data[0:idx])
}

func ParseString(reader io.ReaderAt, offset int64, length int64) string {
   data := make([]byte, length)
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
      return ""
   }
   return string(data[:n])
}


func ParseTerminatedUTF16String(reader io.ReaderAt, offset int64) string {
   var buf [1024]byte
   data := buf[:]
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
     return ""
   }

   idx := bytes.Index(data[:n], []byte{0, 0})
   if idx < 0 {
      idx = n-1
   }
   if idx%2 != 0 {
      idx += 1
   }
   return UTF16BytesToUTF8(data[0:idx], binary.LittleEndian)
}

func ParseUTF16String(reader io.ReaderAt, offset int64, length int64) string {
   data := make([]byte, length)
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
     return ""
   }
   return UTF16BytesToUTF8(data[:n], binary.LittleEndian)
}

func UTF16BytesToUTF8(b []byte, o binary.ByteOrder) string {
	if len(b) < 2 {
		return ""
	}

	if b[0] == 0xff && b[1] == 0xfe {
		o = binary.BigEndian
		b = b[2:]
	} else if b[0] == 0xfe && b[1] == 0xff {
		o = binary.LittleEndian
		b = b[2:]
	}

	utf := make([]uint16, (len(b)+(2-1))/2)

	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}

	return string(utf16.Decode(utf))
}


