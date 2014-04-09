package main

import (
  "path"
  "errors"
  "strings"
  "time"
  "strconv"
  "os"
)

type ItemListHelper struct {
  Id int
  Key string
}

const TimeFormat = time.RFC1123

// Returns the location for the index
func (il *ItemListHelper) RepoLocation() string {
  return path.Join("repos",il.Key,strconv.FormatInt(int64(il.Id),10))
}

// Returns the location for the collection to index
func (il *ItemListHelper) DataLocation() string {
  return path.Join("data",il.Key,strconv.FormatInt(int64(il.Id),10))
}

// Returns the location for the config files
func (il *ItemListHelper) ConfigLocation() string {
  return path.Join("config",il.Key,strconv.FormatInt(int64(il.Id),10))
}

// Deletes the index directory for this collection 
func (il *ItemListHelper) RemoveRepo() (err error) {
  err = os.RemoveAll(il.RepoLocation())
  return
}

// Deletes the data directory for this collection
func (il *ItemListHelper) RemoveData() (err error) {
  err = os.RemoveAll(il.DataLocation())
  return
}

// Creates the data directory for this collection
func (il *ItemListHelper) MkdirData() (err error) {
  err = os.MkdirAll(il.DataLocation(),os.ModeDir | 0755)
  return
}

// Deletes the config directory for this collection
func (il *ItemListHelper) RemoveConfig() (err error) {
  err = os.RemoveAll(il.ConfigLocation())
  return
}

// Creates the config directory for this collection
func (il *ItemListHelper) MkdirConfig() (err error) {
  err = os.MkdirAll(il.ConfigLocation(),os.ModeDir | 0755)
  return
}

// Creates the repo directory for this collection
func (il *ItemListHelper) MkdirRepo() (err error) {
  err = os.MkdirAll(il.RepoLocation(),os.ModeDir | 0755)
  return
}


func (il *ItemListHelper) CreatedTime() (time string,err error) {
  fi, err := os.Lstat(path.Join(il.RepoLocation(),"manifest"))
  if err != nil {
    err = errors.New("No index found for this itemlist")
    return
  }

  time = fi.ModTime().Format(TimeFormat)
  return
}

func (il *ItemListHelper) docIdForFile(filename string) string {
  return strings.TrimPrefix(filename,path.Join(il.DataLocation()))[1:]
}

func (il *ItemListHelper) MakeReadyForDownload() (err error) {
  err = il.RemoveData()
  if err != nil {
    return
  }

  err = il.RemoveConfig()
  if err != nil {
    return
  }

  err = il.MkdirData()
  if err != nil {
    return
  }

  err = il.MkdirConfig()
  return
}

