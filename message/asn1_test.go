package message

import (
  "testing"
  "bytes"
)

func TestSizeInt64(t *testing.T) {
  s := sizeInt64(0)
  if s != 1 {
    t.Errorf("computed size is %d, expected 1", s)
  }

  s = sizeInt64(127)
  if s != 1 {
    t.Errorf("computed size is %d, expected 1", s)
  }

  s = sizeInt64(128)
  if s != 2 {
    t.Errorf("computed size is %d, expected 2", s)
  }

  s = sizeInt64(50000)
  if s != 3 {
    t.Errorf("computed size is %d, expected 3", s)
  }

  s = sizeInt64(-12345)
  if s != 2 {
    t.Errorf("computed size is %d, expected 2", s)
  }
}

func TestWriteInt64(t *testing.T) {
  vtests := []int64{0, 127, 128, 50000, -12345}
  expsize := []int{1, 1, 2, 3, 2}
  expresult := [][]byte{{0x00}, {0x7F}, {0x00, 0x80}, {0x00, 0xc3, 0x50}, {0xcf, 0xc7}}

  for idx, v := range vtests {
    fs := sizeInt64(v)
    b := NewBytes(fs, make([]byte, fs))
    t.Log("computing", v)
    s := writeInt64(b, v)
    if s != expsize[idx] {
      t.Errorf("computed size is %d, expected %d", s, expsize[idx])
    }
    if !bytes.Equal(b.Bytes(), expresult[idx]) {
      t.Errorf("wrong computed bytes, got %v, expected %v", b.Bytes(), expresult[idx])
    }
    a, e := parseInt64(b.Bytes())
    t.Log("parse", a, e)
  }
}
